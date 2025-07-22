package formatter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

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
	filename string
}

func (e *errorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	if e.filename == "" {
		_, _ = fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, e.filename+" line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
	}
	os.Exit(1)
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

	p := parser.NewApexParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&errorListener{filename: f.filename})
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

	v := NewFormatVisitor(stream)
	out, ok := v.visitRule(p.CompilationUnit()).(string)
	if !ok {
		return fmt.Errorf("Unexpected result parsing apex")
	}
	out = removeExtraCommentIndentation(out)
	out = strings.TrimRight(out, " \t\n")
	f.formatted = append([]byte(out), '\n')
	return nil
}

func (f *Formatter) Write() error {
	if f.formatted == nil {
		return fmt.Errorf("No formatted source found")
	}
	return writeFile(f.filename, f.formatted)
}

// removeIndentationFromComment removes extra tabs that were introduced during
// formatting from a single multi-line comment.
func removeIndentationFromComment(comment string) string {
	// Find the position of the initial \uFFFA and the final \uFFFB
	startIndex := strings.Index(comment, "\uFFFA")
	endIndex := strings.LastIndex(comment, "\uFFFB")
	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		// \uFFFA or \uFFFB not found, or the indices are invalid, return the original comment
		return comment
	}

	// Determine the indentation level from the first line
	firstLine := comment[:startIndex]
	leadingTabs := strings.Count(firstLine, "\t")

	// Extract the content between \uFFFA and \uFFFB
	commentBody := comment[startIndex+len("\uFFFA") : endIndex]

	// Create a regex to match the leading tabs in subsequent lines
	tabPattern := fmt.Sprintf(`\n\t{%d}`, leadingTabs)
	re := regexp.MustCompile(tabPattern)

	// Replace the matched pattern with a newline, effectively removing the leading tabs
	modifiedComment := re.ReplaceAllString(commentBody, "\n")

	// Return the modified comment, reattaching any text outside the comment block if necessary
	leadingTabPattern := regexp.MustCompile("^\t*")
	unindented := leadingTabPattern.ReplaceAllString(firstLine, "") + modifiedComment
	// Add leading tab
	firstCharPattern := regexp.MustCompile(`(?s)^(\s*)(\S)`)
	unindented = firstCharPattern.ReplaceAllString(unindented, "$1"+strings.Repeat("\t", leadingTabs)+"$2")

	return unindented
}

var (
	newlinePrefixedMultilineComment = regexp.MustCompile("[\n ]*(\t*[\uFFFA\uFFF9])")
	indentedInlineComment           = regexp.MustCompile("([^\n\uFFFB])\t+\uFFFA([^\n])")
	spacePaddedMultilineComment     = regexp.MustCompile(`(` + "\uFFFB\n*\t*" + `) +`)
	indentInjectedNewlines          = regexp.MustCompile("\uFFFB\n+")
	doubleCapturedNewlines          = regexp.MustCompile("\n(\uFFFB\t*\uFFFA\n)")
	newlinePrefixedInlineComment    = regexp.MustCompile("\n\t*\uFFF9\n")
	tabPrefixedInlineComment        = regexp.MustCompile(`([\w,{]+)` + "\t+\uFFF9")
	inlineCommentPattern            = regexp.MustCompile(`(?s)` + "\uFFF9" + `(.*?)` + "\uFFFB")
	whitespaceInBraces              = regexp.MustCompile(`(?s){\n+}`)
	multilineCommentPattern         = regexp.MustCompile(`(?s)\t*` + "\uFFFA" + `.*?` + "\uFFFB")
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
func removeExtraCommentIndentation(input string) string {
	// Apply regex transformations in sequence using pre-compiled patterns
	input = newlinePrefixedMultilineComment.ReplaceAllString(input, "$1")
	input = indentedInlineComment.ReplaceAllString(input, "$1\uFFFA$2")
	input = spacePaddedMultilineComment.ReplaceAllString(input, "$1")
	input = indentInjectedNewlines.ReplaceAllString(input, "\uFFFB\n")

	// Simple string replacements are faster than regex for exact matches
	input = strings.ReplaceAll(input, "\n\uFFFB\n", "\n\uFFFB")

	input = doubleCapturedNewlines.ReplaceAllString(input, "$1")
	input = newlinePrefixedInlineComment.ReplaceAllString(input, "\uFFF9\n")
	input = tabPrefixedInlineComment.ReplaceAllString(input, "$1\uFFF9")

	// Simple string replacement
	input = strings.ReplaceAll(input, " \uFFF9 ", "\uFFF9 ")

	// Remove inline comment delimiters
	input = inlineCommentPattern.ReplaceAllString(input, "$1")

	// Remove newlines in braces
	input = whitespaceInBraces.ReplaceAllString(input, "{}")

	// Restore formatting of indented multi-line comments
	return multilineCommentPattern.ReplaceAllStringFunc(input, removeIndentationFromComment)
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
