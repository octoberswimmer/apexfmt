package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type TreeShapeListener struct {
	indentLevel int
	*parser.BaseApexParserListener
}

type Visitor struct {
	indentLevel int
	parser.BaseApexParserVisitor
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
	return node.Accept(v)
}

func (v *Visitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	fmt.Fprintln(os.Stderr, "HERE WE GO!")
	t := ctx.TypeDeclaration()
	// TODO: Handle the typeDeclaration modifiers
	switch {
	case t.ClassDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "IT'S A CLASS")
		return v.visitRule(t.ClassDeclaration()).(string)
	case t.InterfaceDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "IT'S AN INTERFACE")
	case t.EnumDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "IT'S AN ENUM")
	}
	return ""
}

func (v *Visitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN THE CLASS DECLARATION")
	// TODO: handle extends and implements
	return fmt.Sprintf("class %s {\n%s\n}", ctx.Id().GetText(), v.visitRule(ctx.ClassBody()))
}

func (v *Visitor) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	fmt.Fprintln(os.Stderr, "NEED TO DEAL WITH THE CLASS BODY")
	var cb []string
	for _, b := range ctx.AllClassBodyDeclaration() {
		cb = append(cb, v.visitRule(b).(string))
	}
	return strings.Join(cb, "\n")
}

func (v *Visitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) interface{} {
	switch {
	case ctx.SEMI() != nil:
		return ";"
	case ctx.Block() != nil:
		// TODO: Handle static
		fmt.Fprintln(os.Stderr, "GOT A BLOCK")
	case ctx.MemberDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "IT'S A MEMBER")
		mods := []string{}
		for _, m := range ctx.AllModifier() {
			mods = append(mods, m.GetText())
		}
		return fmt.Sprintf("%s %s", strings.Join(mods, " "), v.visitRule(ctx.MemberDeclaration()))
	}
	return ""
}

func (v *Visitor) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "MEMBER DECLARATION")
	return ""
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewApexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewApexParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// fmt.Println(TreesIndentedStringTree(p.CompilationUnit(), "", nil, p))
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.CompilationUnit())
	v := NewVisitor()
	out, ok := p.CompilationUnit().Accept(v).(string)
	if !ok {
	}
	fmt.Println(out)
}
