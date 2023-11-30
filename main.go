package main

import (
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/formatter"
	"github.com/octoberswimmer/apexfmt/parser"
)

func main() {
	src, err := readFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %s\n", err.Error())
		os.Exit(1)
	}
	input := antlr.NewInputStream(string(src))
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewApexParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	v := formatter.NewVisitor(stream)
	out, ok := p.CompilationUnit().Accept(v).(string)
	if !ok {
	}
	fmt.Println(out)
	if string(src) != out {
		fmt.Fprintf(os.Stderr, "Content changed\n")
		os.Exit(1)
	}
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
