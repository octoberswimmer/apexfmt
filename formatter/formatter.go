package formatter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type Formatter struct {
	filename  string
	reader    io.Reader
	source    []byte
	formatted []byte
}

func (f *Formatter) SourceName() string {
	if f.filename != "" {
		return f.filename
	}
	return "<stdin>"
}

type errorListener struct {
	*antlr.DefaultErrorListener
	filename       string
	errors         []string
	suppressStderr bool
}

func (e *errorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	var errMsg string
	if e.filename == "" {
		errMsg = fmt.Sprintf("line %d:%d %s", line, column, msg)
	} else {
		errMsg = fmt.Sprintf("%s: line %d:%d %s", e.filename, line, column, msg)
	}
	e.errors = append(e.errors, errMsg)
	// Still print to stderr for backwards compatibility, but don't exit
	if !e.suppressStderr {
		fmt.Fprintln(os.Stderr, errMsg)
	}
}

func (e *errorListener) HasErrors() bool {
	return len(e.errors) > 0
}

func (e *errorListener) GetError() error {
	if len(e.errors) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(e.errors, "\n"))
}

func NewFormatter(filename string, reader io.Reader) *Formatter {
	if filename != "" {
		return &Formatter{
			filename: filename,
		}
	}
	return &Formatter{
		reader: reader,
	}
}

func (f *Formatter) Formatted() (string, error) {
	if f.formatted == nil {
		err := f.Format()
		if err != nil {
			return "", err
		}
	}
	return string(f.formatted), nil
}

func (f *Formatter) Changed() (bool, error) {
	if f.formatted == nil {
		err := f.Format()
		if err != nil {
			return false, err
		}
	}
	return !bytes.Equal(f.source, f.formatted), nil
}

func (f *Formatter) Format() error {
	if f.source == nil {
		src, err := readFile(f.filename, f.reader)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", f.SourceName(), err)
		}
		f.source = src
	}
	input := antlr.NewInputStream(string(f.source))
	lexer, release := parser.NewApexLexerWithPrivateATN(input)
	defer release()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parse with SLL prediction first. It is faster than LL prediction but
	// cannot resolve every ambiguity in the grammar, for example a call
	// statement such as "return name(arg);" or "super(arg);". When it
	// stops at a syntax error, rewind the token stream and parse again with
	// full LL prediction, which reports any real syntax errors.
	engine, release := parser.AcquireParserEngine()
	defer release()
	parseTree, ok := parseSLL(engine, stream)
	if !ok {
		stream.Seek(0)
		p := engine.NewApexParser(stream)
		p.RemoveErrorListeners()
		errListener := &errorListener{filename: f.filename}
		p.AddErrorListener(errListener)
		// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

		parseTree = p.CompilationUnit()

		// Check if there were any syntax errors during parsing
		if errListener.HasErrors() {
			return errListener.GetError()
		}
	}

	v := NewFormatVisitor(stream)
	out, ok := v.visitRule(parseTree).(string)

	if !ok {
		return fmt.Errorf("Unexpected result parsing apex")
	}
	out = removeExtraCommentIndentation(out)
	out = strings.TrimRight(out, " \t\n")
	f.formatted = append([]byte(out), '\n')
	return nil
}

// errBail is the panic value that unwinds an SLL parse at its first syntax
// error.
type errBail struct{}

// bailErrorStrategy abandons the parse at the first syntax error. antlr's own
// BailErrorStrategy is not enough in the Go runtime: the generated rule
// functions clear the parser's error after calling Recover, and since the
// strategy consumes no tokens a rule invoked from a loop such as
// classBodyDeclaration* can fail at the same token forever. Panicking from
// the strategy unwinds the whole parse; parseSLL recovers the panic.
type bailErrorStrategy struct {
	*antlr.DefaultErrorStrategy
}

func (b *bailErrorStrategy) ReportError(_ antlr.Parser, _ antlr.RecognitionException) {
	panic(errBail{})
}

func (b *bailErrorStrategy) Recover(_ antlr.Parser, _ antlr.RecognitionException) {
	panic(errBail{})
}

func (b *bailErrorStrategy) RecoverInline(_ antlr.Parser) antlr.Token {
	panic(errBail{})
}

func (b *bailErrorStrategy) Sync(_ antlr.Parser) {}

// parseSLL parses the token stream with SLL prediction. It returns false when
// the parser hit a syntax error, which may or may not be a real error in the
// source; the caller must then parse again with LL prediction.
func parseSLL(engine *parser.ParserEngine, stream *antlr.CommonTokenStream) (tree parser.ICompilationUnitContext, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			if _, bailed := r.(errBail); !bailed {
				panic(r)
			}
			tree, ok = nil, false
		}
	}()
	p := engine.NewApexParser(stream)
	p.RemoveErrorListeners()
	p.SetErrorHandler(&bailErrorStrategy{DefaultErrorStrategy: antlr.NewDefaultErrorStrategy()})
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	return p.CompilationUnit(), true
}

func (f *Formatter) Write() error {
	if f.formatted == nil {
		return fmt.Errorf("No formatted source found")
	}
	return writeFile(f.filename, f.formatted)
}

const (
	inlineCommentStart    = "\uFFF9"
	multilineCommentStart = "\uFFFA"
	commentEnd            = "\uFFFB"
	// The three markers share the same first two bytes in UTF-8, so a scan
	// for the first byte finds candidates for any of them.
	markerLen       = len(commentEnd)
	markerFirstByte = byte(0xEF)
)

// nextMarker returns the index of the first comment marker at or after pos
// whose text is one of markers, or -1.
func nextMarker(s string, pos int, markers ...string) int {
	for pos < len(s) {
		i := strings.IndexByte(s[pos:], markerFirstByte)
		if i < 0 {
			return -1
		}
		i += pos
		for _, m := range markers {
			if strings.HasPrefix(s[i:], m) {
				return i
			}
		}
		pos = i + 1
	}
	return -1
}

// removeIndentationFromComment removes extra tabs that were introduced during
// formatting from a single multi-line comment. The comment consists of the
// tabs that indent it, a \uFFFA marker, the comment text, and a \uFFFB marker.
// The markers are removed.
func removeIndentationFromComment(comment string) string {
	// Find the position of the initial \uFFFA and the final \uFFFB
	startIndex := strings.Index(comment, multilineCommentStart)
	endIndex := strings.LastIndex(comment, commentEnd)
	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		// \uFFFA or \uFFFB not found, or the indices are invalid, return the original comment
		return comment
	}

	// Determine the indentation level from the first line
	firstLine := comment[:startIndex]
	leadingTabs := strings.Count(firstLine, "\t")
	tabs := strings.Repeat("\t", leadingTabs)

	// Extract the content between \uFFFA and \uFFFB
	commentBody := comment[startIndex+len(multilineCommentStart) : endIndex]

	// Remove the leading tabs from each subsequent line
	if leadingTabs > 0 {
		commentBody = strings.ReplaceAll(commentBody, "\n"+tabs, "\n")
	}

	unindented := strings.TrimLeft(firstLine, "\t") + commentBody
	// Add leading tab before the first non-whitespace character
	if i := strings.IndexFunc(unindented, isNotRegexpSpace); i >= 0 {
		unindented = unindented[:i] + tabs + unindented[i:]
	}

	return unindented
}

// isNotRegexpSpace reports whether r is outside the character class matched
// by \s in package regexp.
func isNotRegexpSpace(r rune) bool {
	switch r {
	case '\t', '\n', '\f', '\r', ' ':
		return false
	}
	return true
}

func isWordCommaOrBrace(c byte) bool {
	return c == '_' || c == ',' || c == '{' ||
		(c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// editor applies deletions and replacements to a string, copying it only once
// the first edit is made and returning the original string when there are no
// edits. Every cleanup pass over the formatted output uses it, so a pass that
// finds nothing to change costs a scan and no allocation.
type editor struct {
	s       string
	b       strings.Builder
	pos     int
	changed bool
}

func newEditor(s string) *editor {
	return &editor{s: s}
}

// replace substitutes with for s[from:to]. Edits must be made in order.
func (e *editor) replace(from, to int, with string) {
	if !e.changed {
		e.changed = true
		e.b.Grow(len(e.s))
	}
	e.b.WriteString(e.s[e.pos:from])
	e.b.WriteString(with)
	e.pos = to
}

func (e *editor) cut(from, to int) {
	e.replace(from, to, "")
}

func (e *editor) String() string {
	if !e.changed {
		return e.s
	}
	e.b.WriteString(e.s[e.pos:])
	return e.b.String()
}

// removeNewlinesBeforeCommentStart removes newlines and spaces that precede
// the tabs before a comment start marker.
func removeNewlinesBeforeCommentStart(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, multilineCommentStart, inlineCommentStart)
		if m < 0 {
			break
		}
		t := m
		for t > pos && s[t-1] == '\t' {
			t--
		}
		w := t
		for w > pos && (s[w-1] == '\n' || s[w-1] == ' ') {
			w--
		}
		if w < t {
			e.cut(w, t)
		}
		pos = m + markerLen
	}
	return e.String()
}

// removeTabsBeforeIndentedMultilineComment removes the tabs between a
// multi-line comment start marker and the preceding character when that
// character is not a newline or a comment end marker and the marker is not at
// the end of a line. The preceding character may itself be a tab, so when the
// character before the run of tabs does not qualify, the first tab of a run
// of two or more takes its place and stays.
func removeTabsBeforeIndentedMultilineComment(s string) string {
	e := newEditor(s)
	pos := 0
	// The character after the marker belongs to the previous match, so the
	// next match cannot begin before it.
	guard := 0
	for {
		m := nextMarker(s, pos, multilineCommentStart)
		if m < 0 {
			break
		}
		next := m + markerLen
		t := m
		for t > guard && s[t-1] == '\t' {
			t--
		}
		keep := -1
		if t < m && next < len(s) && s[next] != '\n' {
			prev := utf8.RuneError
			if t > guard {
				prev, _ = utf8.DecodeLastRuneInString(s[:t])
			}
			if t > guard && prev != '\n' && prev != '￻' {
				keep = t
			} else if m-t >= 2 {
				keep = t + 1
			}
		}
		if keep >= 0 {
			e.cut(keep, m)
			_, size := utf8.DecodeRuneInString(s[next:])
			guard = next + size
		}
		pos = next
	}
	return e.String()
}

// removeSpacesAfterCommentEnd removes the spaces that follow a comment end
// marker and any newlines and tabs after it, as the regexp
// "(￻\n*\t*) +" replaced with "$1" did.
func removeSpacesAfterCommentEnd(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, commentEnd)
		if m < 0 {
			break
		}
		i := m + markerLen
		for i < len(s) && s[i] == '\n' {
			i++
		}
		for i < len(s) && s[i] == '\t' {
			i++
		}
		j := i
		for j < len(s) && s[j] == ' ' {
			j++
		}
		if j > i {
			e.cut(i, j)
		}
		pos = j
	}
	return e.String()
}

// collapseNewlinesAfterCommentEnd keeps one newline of a run that follows a
// comment end marker, as the regexp "￻\n+" replaced with "￻\n" did.
func collapseNewlinesAfterCommentEnd(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, commentEnd)
		if m < 0 {
			break
		}
		i := m + markerLen
		j := i
		for j < len(s) && s[j] == '\n' {
			j++
		}
		if j-i >= 2 {
			e.cut(i+1, j)
		}
		pos = j
	}
	return e.String()
}

// removeNewlineBeforeAdjacentComments removes the newline before a comment
// end marker that is followed, after tabs, by a multi-line comment start
// marker and a newline, as the regexp "\n(￻\t*￺\n)" replaced with
// "$1" did.
func removeNewlineBeforeAdjacentComments(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, commentEnd)
		if m < 0 {
			break
		}
		next := m + markerLen
		i := next
		for i < len(s) && s[i] == '\t' {
			i++
		}
		if m > pos && s[m-1] == '\n' && strings.HasPrefix(s[i:], multilineCommentStart) && i+markerLen < len(s) && s[i+markerLen] == '\n' {
			e.cut(m-1, m)
			pos = i + markerLen + 1
			continue
		}
		pos = next
	}
	return e.String()
}

// removeNewlineBeforeInlineCommentLine removes the newline and tabs before an
// inline comment start marker that ends its line, as the regexp
// "\n\t*￹\n" replaced with "￹\n" did.
func removeNewlineBeforeInlineCommentLine(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, inlineCommentStart)
		if m < 0 {
			break
		}
		next := m + markerLen
		t := m
		for t > pos && s[t-1] == '\t' {
			t--
		}
		if t > pos && s[t-1] == '\n' && next < len(s) && s[next] == '\n' {
			e.cut(t-1, m)
			pos = next + 1
			continue
		}
		pos = next
	}
	return e.String()
}

// removeTabsBeforeInlineComment removes the tabs between an inline comment
// start marker and a preceding word character, comma or opening brace.
func removeTabsBeforeInlineComment(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, inlineCommentStart)
		if m < 0 {
			break
		}
		t := m
		for t > pos && s[t-1] == '\t' {
			t--
		}
		if t < m && t > pos && isWordCommaOrBrace(s[t-1]) {
			e.cut(t, m)
		}
		pos = m + markerLen
	}
	return e.String()
}

// removeInlineCommentMarkers removes each inline comment start marker together
// with the next comment end marker.
func removeInlineCommentMarkers(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, inlineCommentStart)
		if m < 0 {
			break
		}
		end := nextMarker(s, m+markerLen, commentEnd)
		if end < 0 {
			break
		}
		e.cut(m, m+markerLen)
		e.cut(end, end+markerLen)
		pos = end + markerLen
	}
	return e.String()
}

// removeNewlinesInEmptyBraces turns a brace pair holding only newlines into
// "{}", as the regexp "{\n+}" did.
func removeNewlinesInEmptyBraces(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		i := strings.IndexByte(s[pos:], '{')
		if i < 0 {
			break
		}
		i += pos
		j := i + 1
		for j < len(s) && s[j] == '\n' {
			j++
		}
		if j > i+1 && j < len(s) && s[j] == '}' {
			e.cut(i+1, j)
			pos = j + 1
			continue
		}
		pos = i + 1
	}
	return e.String()
}

// restoreMultilineComments applies removeIndentationFromComment to each
// multi-line comment, including the tabs that indent it.
func restoreMultilineComments(s string) string {
	e := newEditor(s)
	pos := 0
	for {
		m := nextMarker(s, pos, multilineCommentStart)
		if m < 0 {
			break
		}
		end := nextMarker(s, m+markerLen, commentEnd)
		if end < 0 {
			break
		}
		t := m
		for t > pos && s[t-1] == '\t' {
			t--
		}
		e.replace(t, end+markerLen, removeIndentationFromComment(s[t:end+markerLen]))
		pos = end + markerLen
	}
	return e.String()
}

// removeExtraCommentIndentation cleans up the formatting of comments after the
// formatter has run.
//
// The antlr lexer pulls comments into a separate token stream so we don't need
// to check for comments in every visit function. Instead, we look for
// comments, each represented as a single token, before the start of or after
// the end of the current parser node. Then we reinject the comments as we're
// visiting each node.
//
// The visitor functions don't know about the comments so they introduce
// whitespace around them when formatting and indenting the code. We need to
// ensure that the comments don't end up mangled. We wrap the comments in
// delimiters so we can easily identify the comments and clean up after
// formatter runs. This code cleans up the whitespace and removes the comment
// delimiters.
//
// Each pass scans for the marker it concerns and copies the output only when
// it changes something; package regexp would copy the output twice per pass
// whether or not a pattern matched.
func removeExtraCommentIndentation(input string) string {
	if strings.IndexByte(input, markerFirstByte) < 0 {
		// No comments: only the brace cleanup applies.
		return removeNewlinesInEmptyBraces(input)
	}
	input = removeNewlinesBeforeCommentStart(input)
	input = removeTabsBeforeIndentedMultilineComment(input)
	input = removeSpacesAfterCommentEnd(input)
	input = collapseNewlinesAfterCommentEnd(input)
	input = strings.ReplaceAll(input, "\n￻\n", "\n￻")
	input = removeNewlineBeforeAdjacentComments(input)
	input = removeNewlineBeforeInlineCommentLine(input)
	input = removeTabsBeforeInlineComment(input)
	input = strings.ReplaceAll(input, " ￹ ", "￹ ")
	input = removeInlineCommentMarkers(input)
	input = removeNewlinesInEmptyBraces(input)
	return restoreMultilineComments(input)
}

func readFile(filename string, reader io.Reader) ([]byte, error) {
	r := reader
	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r = f
	}
	src, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func writeFile(filename string, contents []byte) error {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %s\n", err.Error())
		os.Exit(1)
	}
	perm := info.Mode().Perm()
	size := info.Size()
	fout, err := os.OpenFile(filename, os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	defer fout.Close()
	n, err := fout.Write(contents)
	if err == nil && int64(n) < size {
		err = fout.Truncate(int64(n))
	}
	return err
}
