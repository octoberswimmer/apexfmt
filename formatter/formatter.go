package formatter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type Formatter struct {
	filename  string
	source    []byte
	formatted []byte
}

type errorListener struct {
	*antlr.DefaultErrorListener
	filename string
}

func (e *errorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	_, _ = fmt.Fprintln(os.Stderr, e.filename+" line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}

func NewFormatter(filename string) *Formatter {
	return &Formatter{
		filename: filename,
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
		src, err := readFile(f.filename)
		if err != nil {
			return fmt.Errorf("Failed to read file %s: %w", f.filename, err)
		}
		f.source = src
	}
	input := antlr.NewInputStream(string(f.source))
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewApexParser(stream)
	p.AddErrorListener(&errorListener{filename: f.filename})
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

	v := NewVisitor(stream)
	out, ok := p.CompilationUnit().Accept(v).(string)
	if !ok {
		return fmt.Errorf("Unexpected result parsing apex")
	}
	f.formatted = []byte(out)
	return nil
}

func (f *Formatter) Write() error {
	if f.formatted == nil {
		return fmt.Errorf("No formatted source found")
	}
	return writeFile(f.filename, f.formatted)
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	src, err := io.ReadAll(f)
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
	if len(contents) > 0 && contents[len(contents)-1] != '\n' {
		contents = append(contents, '\n')
	}
	n, err := fout.Write(contents)
	if err == nil && int64(n) < size {
		err = fout.Truncate(int64(n))
	}
	return err
}
