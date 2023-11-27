package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type TreeShapeListener struct {
	indentLevel int
	*parser.BaseApexParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (l *TreeShapeListener) indent() {
	for i := 0; i < l.indentLevel; i++ {
		fmt.Printf("\t")
	}
}

func (l *TreeShapeListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	l.indent()
	fmt.Printf("class %s ", ctx.Id().GetText())
}

func (l *TreeShapeListener) EnterClassBody(ctx *parser.ClassBodyContext) {
	fmt.Printf("{\n")
	l.indentLevel++
}

func (l *TreeShapeListener) ExitClassBody(ctx *parser.ClassBodyContext) {
	l.indent()
	fmt.Printf("}\n")
	l.indentLevel--
}

func (l *TreeShapeListener) EnterFieldDeclaration(ctx *parser.FieldDeclarationContext) {
	l.indent()
	fmt.Printf("%s", ctx.TypeRef().GetText())
}

func (l *TreeShapeListener) EnterVariableDeclarator(ctx *parser.VariableDeclaratorContext) {
	fmt.Printf(" %s", ctx.Id().GetText())
	e := ctx.Expression()
	if e != nil {
		fmt.Printf(" = %s", e.GetText())
	}
}

func (l *TreeShapeListener) ExitFieldDeclaration(ctx *parser.FieldDeclarationContext) {
	fmt.Printf(";\n")
}

func (l *TreeShapeListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	t := ""
	if ctx.TypeRef() != nil {
		t = ctx.TypeRef().GetText()
	}
	l.indent()
	fmt.Printf("%s %s ", t, ctx.Id().GetText())
}

func (l *TreeShapeListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	l.indent()
	fmt.Printf("}\n")
	l.indentLevel--
}

func (l *TreeShapeListener) EnterBlock(ctx *parser.BlockContext) {
	l.indent()
	fmt.Printf("{\n")
	l.indentLevel++
}

func (l *TreeShapeListener) ExitBlock(ctx *parser.BlockContext) {
	l.indent()
	fmt.Printf("}\n")
	l.indentLevel--
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	ctxType := reflect.TypeOf(ctx)
	// fmt.Println(ctxType, ctx.GetStart().GetText())
	// fmt.Println(ctx.GetText(), "\n\n")
	switch ctx.(type) {
	default:
		fmt.Fprintf(os.Stderr, "(%s) %s\n", ctxType, ctx.GetStart().GetText())
	}
	_ = ctxType
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewApexParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// fmt.Println(TreesIndentedStringTree(p.CompilationUnit(), "", nil, p))
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.CompilationUnit())
}
