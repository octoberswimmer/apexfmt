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
	f.formatted = append([]byte(out), '\n')
	return nil
}

func (f *Formatter) Write() error {
	if f.formatted == nil {
		return fmt.Errorf("No formatted source found")
	}
	return writeFile(f.filename, f.formatted)
}

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
	return firstLine + modifiedComment
}

// Comments are annotated in FormatVisitor.visitRule.  We preserve whitespace
// within multi-line comments by removing the indentation added within the
// comment.
func removeExtraCommentIndentation(input string) string {
	commentPattern := regexp.MustCompile(`(?s)\t*` + "\uFFFA" + `.*?` + "\uFFFB")
	return commentPattern.ReplaceAllStringFunc(input, removeIndentationFromComment)
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
