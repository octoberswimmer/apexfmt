package formatter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
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
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parse with SLL prediction first. It is faster than LL prediction but
	// cannot resolve every ambiguity in the grammar, for example a call
	// statement such as "return name(arg);" or "super(arg);". When it
	// stops at a syntax error, rewind the token stream and parse again with
	// full LL prediction, which reports any real syntax errors.
	parseTree, ok := parseSLL(stream)
	if !ok {
		stream.Seek(0)
		p := parser.NewApexParser(stream)
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

// bailErrorStrategy stops the parser at the first syntax error without
// reporting it. antlr's BailErrorStrategy passes the cancellation to
// DefaultErrorStrategy.ReportError, which prints "unknown recognition error
// type" to stdout, so ReportError is overridden to record the failure only.
type bailErrorStrategy struct {
	*antlr.BailErrorStrategy
	failed bool
}

func (b *bailErrorStrategy) ReportError(_ antlr.Parser, _ antlr.RecognitionException) {
	b.failed = true
}

// parseSLL parses the token stream with SLL prediction. It returns false when
// the parser hit a syntax error, which may or may not be a real error in the
// source; the caller must then parse again with LL prediction.
func parseSLL(stream *antlr.CommonTokenStream) (parser.ICompilationUnitContext, bool) {
	p := parser.NewApexParser(stream)
	p.RemoveErrorListeners()
	strategy := &bailErrorStrategy{BailErrorStrategy: antlr.NewBailErrorStrategy()}
	p.SetErrorHandler(strategy)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := p.CompilationUnit()
	return tree, !strategy.failed
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

// removeNewlinesBeforeCommentStart removes newlines and spaces that precede
// the tabs before a comment start marker.
func removeNewlinesBeforeCommentStart(s string) string {
	var b strings.Builder
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
		b.WriteString(s[pos:w])
		b.WriteString(s[t : m+markerLen])
		pos = m + markerLen
	}
	if pos == 0 {
		return s
	}
	b.WriteString(s[pos:])
	return b.String()
}

// removeTabsBeforeIndentedMultilineComment removes the tabs between a
// multi-line comment start marker and the preceding character when that
// character is not a newline or a comment end marker and the marker is not at
// the end of a line. The preceding character may itself be a tab, so when the
// character before the run of tabs does not qualify, the first tab of a run
// of two or more takes its place and stays.
func removeTabsBeforeIndentedMultilineComment(s string) string {
	var b strings.Builder
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
			if t > guard && prev != '\n' && prev != '\uFFFB' {
				keep = t
			} else if m-t >= 2 {
				keep = t + 1
			}
		}
		if keep >= 0 {
			b.WriteString(s[pos:keep])
			b.WriteString(multilineCommentStart)
			_, size := utf8.DecodeRuneInString(s[next:])
			guard = next + size
		} else {
			b.WriteString(s[pos:next])
		}
		pos = next
	}
	if pos == 0 {
		return s
	}
	b.WriteString(s[pos:])
	return b.String()
}

// removeTabsBeforeInlineComment removes the tabs between an inline comment
// start marker and a preceding word character, comma or opening brace.
func removeTabsBeforeInlineComment(s string) string {
	var b strings.Builder
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
			b.WriteString(s[pos:t])
		} else {
			b.WriteString(s[pos:m])
		}
		b.WriteString(inlineCommentStart)
		pos = m + markerLen
	}
	if pos == 0 {
		return s
	}
	b.WriteString(s[pos:])
	return b.String()
}

// removeInlineCommentMarkers removes each inline comment start marker together
// with the next comment end marker.
func removeInlineCommentMarkers(s string) string {
	var b strings.Builder
	pos := 0
	for {
		m := nextMarker(s, pos, inlineCommentStart)
		if m < 0 {
			break
		}
		e := nextMarker(s, m+markerLen, commentEnd)
		if e < 0 {
			break
		}
		b.WriteString(s[pos:m])
		b.WriteString(s[m+markerLen : e])
		pos = e + markerLen
	}
	if pos == 0 {
		return s
	}
	b.WriteString(s[pos:])
	return b.String()
}

// restoreMultilineComments applies removeIndentationFromComment to each
// multi-line comment, including the tabs that indent it.
func restoreMultilineComments(s string) string {
	var b strings.Builder
	pos := 0
	for {
		m := nextMarker(s, pos, multilineCommentStart)
		if m < 0 {
			break
		}
		e := nextMarker(s, m+markerLen, commentEnd)
		if e < 0 {
			break
		}
		t := m
		for t > pos && s[t-1] == '\t' {
			t--
		}
		b.WriteString(s[pos:t])
		b.WriteString(removeIndentationFromComment(s[t : e+markerLen]))
		pos = e + markerLen
	}
	if pos == 0 {
		return s
	}
	b.WriteString(s[pos:])
	return b.String()
}

var (
	spacePaddedMultilineComment  = regexp.MustCompile(`(` + "\uFFFB\n*\t*" + `) +`)
	indentInjectedNewlines       = regexp.MustCompile("\uFFFB\n+")
	doubleCapturedNewlines       = regexp.MustCompile("\n(\uFFFB\t*\uFFFA\n)")
	newlinePrefixedInlineComment = regexp.MustCompile("\n\t*\uFFF9\n")
	whitespaceInBraces           = regexp.MustCompile(`(?s){\n+}`)
)

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
// The passes that would need a regexp without a literal prefix are written as
// scans for the marker instead; package regexp tries such a pattern at every
// position of the input.
func removeExtraCommentIndentation(input string) string {
	if strings.IndexByte(input, markerFirstByte) < 0 {
		// No comments: only the brace cleanup applies.
		return whitespaceInBraces.ReplaceAllString(input, "{}")
	}
	input = removeNewlinesBeforeCommentStart(input)
	input = removeTabsBeforeIndentedMultilineComment(input)
	input = spacePaddedMultilineComment.ReplaceAllString(input, "$1")
	input = indentInjectedNewlines.ReplaceAllString(input, "\uFFFB\n")

	// Simple string replacements are faster than regex for exact matches
	input = strings.ReplaceAll(input, "\n\uFFFB\n", "\n\uFFFB")

	input = doubleCapturedNewlines.ReplaceAllString(input, "$1")
	input = newlinePrefixedInlineComment.ReplaceAllString(input, "\uFFF9\n")
	input = removeTabsBeforeInlineComment(input)

	// Simple string replacement
	input = strings.ReplaceAll(input, " \uFFF9 ", "\uFFF9 ")

	// Remove inline comment delimiters
	input = removeInlineCommentMarkers(input)

	// Remove newlines in braces
	input = whitespaceInBraces.ReplaceAllString(input, "{}")

	// Restore formatting of indented multi-line comments
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
