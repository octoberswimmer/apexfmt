package formatter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type IndentVisitor struct {
	parser.BaseApexParserVisitor
}

func NewIndentVisitor() *IndentVisitor {
	return &IndentVisitor{}
}

func (v *IndentVisitor) visitRule(node antlr.RuleNode) interface{} {
	result := node.Accept(v)
	if r, ok := result.(int); ok {
		return r
	}
	return 0
}

func (v *IndentVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	indent := 0
	for _, p := range ctx.AllExpression() {
		n := v.visitRule(p).(int)
		if n > indent {
			indent = n
		}
	}
	return indent
}

func (v *IndentVisitor) VisitDotExpression(ctx *parser.DotExpressionContext) interface{} {
	switch {
	case ctx.DotMethodCall() != nil:
		switch ctx.Expression().(type) {
		case *parser.PrimaryExpressionContext:
			return 0
		case *parser.NewInstanceExpressionContext:
			return 2
		default:
			return v.visitRule(ctx.Expression())
		}
	}
	return 0
}

func (v *IndentVisitor) VisitNewInstanceExpression(ctx *parser.NewInstanceExpressionContext) interface{} {
	return 2
}
