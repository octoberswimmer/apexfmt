package formatter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
	log "github.com/sirupsen/logrus"
)

type ChainVisitor struct {
	parser.BaseApexParserVisitor
}

func NewChainVisitor() *ChainVisitor {
	return &ChainVisitor{}
}

func (v *ChainVisitor) visitRule(node antlr.RuleNode) interface{} {
	result := node.Accept(v)
	if r, ok := result.(int); ok {
		return r
	}
	if result == nil {
		log.Debug(fmt.Sprintf("missing ChainVisitor function for %T", node))
	}
	return 0
}

func (v *ChainVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	child := ctx.GetChild(0).(antlr.RuleNode)
	return v.visitRule(child)
}

func (v *ChainVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *ChainVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	switch e := ctx.Primary().(type) {
	case *parser.ThisPrimaryContext:
		return 0
	case *parser.SuperPrimaryContext:
		return 0
	case *parser.LiteralPrimaryContext:
		return 0
	case *parser.TypeRefPrimaryContext:
		return 0
	case *parser.IdPrimaryContext:
		return 0
	case *parser.SoqlPrimaryContext:
		return 1
	case *parser.SoslPrimaryContext:
		return 1
	default:
		return fmt.Sprintf("UNHANDLED PRIMARY EXPRESSION: %T %s", e, e.GetText())
	}
}

func (v *ChainVisitor) VisitDotExpression(ctx *parser.DotExpressionContext) interface{} {
	if ctx.DotMethodCall() != nil {
		return 1 + v.visitRule(ctx.Expression()).(int)
	}
	return v.visitRule(ctx.Expression())
}

func (v *ChainVisitor) VisitLogAndExpression(ctx *parser.LogAndExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitLogOrExpression(ctx *parser.LogOrExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitSubExpression(ctx *parser.SubExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression()).(int)
}

func (v *ChainVisitor) VisitQuery(ctx *parser.QueryContext) interface{} {
	score := v.visitRule(ctx.SelectList()).(int) + v.visitRule(ctx.FromNameList()).(int)
	if scope := ctx.UsingScope(); scope != nil {
		score++
	}
	if where := ctx.WhereClause(); where != nil {
		score += v.visitRule(where).(int)
	}
	if groupBy := ctx.GroupByClause(); groupBy != nil {
		score += v.visitRule(groupBy).(int)
	}
	if orderBy := ctx.OrderByClause(); orderBy != nil {
		score += v.visitRule(orderBy).(int)
	}
	if limit := ctx.LimitClause(); limit != nil {
		score++
	}
	if offset := ctx.OffsetClause(); offset != nil {
		score++
	}
	score += v.visitRule(ctx.ForClauses()).(int)
	if update := ctx.UpdateList(); update != nil {
		score++
	}
	return score
}

func (v *ChainVisitor) VisitGroupByClause(ctx *parser.GroupByClauseContext) interface{} {
	score := 1
	if ctx.LogicalExpression() != nil {
		score += v.visitRule(ctx.LogicalExpression()).(int)
	}
	return score
}

func (v *ChainVisitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) interface{} {
	return len(ctx.FieldOrderList().AllFieldOrder())
}

func (v *ChainVisitor) VisitSelectList(ctx *parser.SelectListContext) interface{} {
	score := 0
	for _, p := range ctx.AllSelectEntry() {
		score += v.visitRule(p).(int)
	}
	return score
}

func (v *ChainVisitor) VisitSelectEntry(ctx *parser.SelectEntryContext) interface{} {
	if ctx.SubQuery() != nil {
		// estimate complexity; probably close enough
		return 5
	}
	if ctx.TypeOf() != nil {
		return 3
	}
	return 1
}

func (v *ChainVisitor) VisitFromNameList(ctx *parser.FromNameListContext) interface{} {
	return len(ctx.AllFieldNameAlias())
}

func (v *ChainVisitor) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return v.visitRule(ctx.LogicalExpression())
}

func (v *ChainVisitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) interface{} {
	return len(ctx.AllSOQLOR()) + len(ctx.AllSOQLAND()) + v.visitRule(ctx.ConditionalExpression(0)).(int)
}

func (v *ChainVisitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) interface{} {
	switch {
	case ctx.LogicalExpression() != nil:
		return fmt.Sprintf("(%s)", v.visitRule(ctx.LogicalExpression()))
	case ctx.FieldExpression() != nil:
		return v.visitRule(ctx.FieldExpression())
	}
	panic("Unexpected conditionalExpression")
}

func (v *ChainVisitor) VisitFieldExpression(ctx *parser.FieldExpressionContext) interface{} {
	switch {
	case ctx.FieldName() != nil:
		switch ctx.ComparisonOperator().GetText() {
		case "IN":
			return 1 + v.visitRule(ctx.Value()).(int)
		case "NOTIN":
			return 1 + v.visitRule(ctx.Value()).(int)
		default:
			return 1
		}
	case ctx.SoqlFunction() != nil:
		return 1 + v.visitRule(ctx.Value()).(int)
	}
	panic("Unexpected fieldExpression")
}

func (v *ChainVisitor) VisitSubQueryValue(ctx *parser.SubQueryValueContext) interface{} {
	// estimate complexity; probably close enough
	return 5
}

func (v *ChainVisitor) VisitForClauses(ctx *parser.ForClausesContext) interface{} {
	return len(ctx.AllForClause())
}

func (v *ChainVisitor) VisitArth2Expression(ctx *parser.Arth2ExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}
