package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type TreeShapeListener struct {
	*parser.BaseApexParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println(ctx.GetText())
	ctxType := reflect.TypeOf(ctx)
	fmt.Println(ctxType, ctx.GetStart().GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewApexParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.CompilationUnit())
}
