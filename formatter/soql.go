package formatter

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type SOQLFormatter struct {
	source    []byte
	formatted []byte
}

func NewSOQLFormatter() *SOQLFormatter {
	return &SOQLFormatter{}
}

func (f *SOQLFormatter) Formatted() (string, error) {
	if f.formatted == nil {
		err := f.Format()
		if err != nil {
			return "", err
		}
	}
	return string(f.formatted), nil
}

func (f *SOQLFormatter) Changed() (bool, error) {
	if f.formatted == nil {
		err := f.Format()
		if err != nil {
			return false, err
		}
	}
	return !bytes.Equal(f.source, f.formatted), nil
}

func (f *SOQLFormatter) Format() error {
	if f.source == nil {
		src, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("Failed to read in query: %w", err)
		}
		f.source = src
	}
	input := antlr.NewInputStream(string(f.source))
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewApexParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&errorListener{})

	v := NewFormatVisitor(stream)
	out, ok := v.visitRule(p.Query()).(string)
	if !ok {
		return fmt.Errorf("Unexpected result parsing apex")
	}
	f.formatted = append([]byte(out), '\n')
	return nil
}
