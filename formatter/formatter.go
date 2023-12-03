package formatter

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type Visitor struct {
	tokens         *antlr.CommonTokenStream
	commentsOutput map[int]struct{}
	parser.BaseApexParserVisitor
}

func NewVisitor(tokens *antlr.CommonTokenStream) *Visitor {
	return &Visitor{
		tokens:         tokens,
		commentsOutput: make(map[int]struct{}),
	}
}

func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
	start := node.(antlr.ParserRuleContext).GetStart()
	beforeComments := v.tokens.GetHiddenTokensToLeft(start.GetTokenIndex(), 3)
	result := node.Accept(v)
	if result == nil {
		panic(fmt.Sprintf("MISSING VISIT FUNCTION FOR %T", node))
	}
	if beforeComments != nil {
		comments := []string{}
		for _, c := range beforeComments {
			if _, seen := v.commentsOutput[c.GetTokenIndex()]; !seen {
				comments = append(comments, c.GetText())
				v.commentsOutput[c.GetTokenIndex()] = struct{}{}
			}
		}
		result = fmt.Sprintf("%s\n%s", strings.Join(comments, "\n"), result)
	}
	return result
}

func (v *Visitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	t := ctx.TypeDeclaration()
	switch {
	case t.ClassDeclaration() != nil:
		return fmt.Sprintf("%s%s", v.Modifiers(t.AllModifier()), v.visitRule(t.ClassDeclaration()).(string))
	case t.InterfaceDeclaration() != nil:
		return fmt.Sprintf("%s%s", v.Modifiers(t.AllModifier()), v.visitRule(t.InterfaceDeclaration()).(string))
	case t.EnumDeclaration() != nil:
		enum := t.EnumDeclaration()
		constants := []string{}
		if enum.EnumConstants() != nil {
			for _, e := range enum.EnumConstants().AllId() {
				constants = append(constants, e.GetText())
			}
		}
		return fmt.Sprintf("enum %s {%s}", enum.Id().GetText(), strings.Join(constants, ", "))
	}
	return ""
}

func (v *Visitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	extends := ""
	if ctx.EXTENDS() != nil {
		extends = fmt.Sprintf(" extends %s ", v.visitRule(ctx.TypeRef()))
	}
	implements := ""
	if ctx.IMPLEMENTS() != nil {
		extends = fmt.Sprintf(" implements %s ", v.visitRule(ctx.TypeList()))
	}
	return fmt.Sprintf("class %s%s%s {\n%s\n}\n", ctx.Id().GetText(),
		extends,
		implements,
		indent(v.visitRule(ctx.ClassBody()).(string)))
}

func indent(text string) string {
	var indentedText strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		indentedText.WriteString("\t" + scanner.Text() + "\n")
	}

	return indentedText.String()
}

func (v *Visitor) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) interface{} {
	extends := ""
	if ctx.EXTENDS() != nil {
		extends = fmt.Sprintf(" extends %s ", v.visitRule(ctx.TypeList()))
	}
	return fmt.Sprintf("interface %s%s {\n%s\n}\n", ctx.Id().GetText(), extends, indent(v.visitRule(ctx.InterfaceBody()).(string)))
}

func (v *Visitor) VisitInterfaceBody(ctx *parser.InterfaceBodyContext) interface{} {
	declarations := []string{}
	for _, d := range ctx.AllInterfaceMethodDeclaration() {
		declarations = append(declarations, v.visitRule(d).(string))
	}
	return strings.Join(declarations, "\n")
}

func (v *Visitor) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
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
		static := ""
		if ctx.STATIC() != nil {
			static = "static "
		}
		return fmt.Sprintf("%s%s", static, indent(v.visitRule(ctx.Block()).(string)))
	case ctx.MemberDeclaration() != nil:
		return fmt.Sprintf("%s%s", v.Modifiers(ctx.AllModifier()), v.visitRule(ctx.MemberDeclaration()))
	}
	return ""
}

func (v *Visitor) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	return v.visitRule(ctx.GetChild(0).(antlr.RuleNode))
}

func (v *Visitor) VisitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) interface{} {
	returnType := "void"
	if ctx.TypeRef() != nil {
		returnType = v.visitRule(ctx.TypeRef()).(string)
	}
	return fmt.Sprintf("%s%s %s%s;", v.Modifiers(ctx.AllModifier()), returnType, ctx.Id().GetText(), v.visitRule(ctx.FormalParameters()))
}

func (v *Visitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	return fmt.Sprintf("%s %s;", v.visitRule(ctx.TypeRef()), v.visitRule(ctx.VariableDeclarators()))
}

func (v *Visitor) VisitPropertyDeclaration(ctx *parser.PropertyDeclarationContext) interface{} {
	propertyBlocks := []string{}
	if ctx.AllPropertyBlock() != nil {
		for _, p := range ctx.AllPropertyBlock() {
			propertyBlocks = append(propertyBlocks, v.visitRule(p).(string))
		}
	}
	return fmt.Sprintf("%s %s {\n%s}\n", v.visitRule(ctx.TypeRef()), ctx.Id().GetText(), indent(strings.Join(propertyBlocks, "\n")))
}

func (v *Visitor) VisitPropertyBlock(ctx *parser.PropertyBlockContext) interface{} {
	if ctx.Getter() != nil {
		return fmt.Sprintf("%s%s", v.Modifiers(ctx.AllModifier()), v.visitRule(ctx.Getter()))
	} else {
		return fmt.Sprintf("%s%s", v.Modifiers(ctx.AllModifier()), v.visitRule(ctx.Setter()))
	}
}

func (v *Visitor) VisitGetter(ctx *parser.GetterContext) interface{} {
	if ctx.SEMI() != nil {
		return "get;"
	} else {
		return fmt.Sprintf("get %s", v.visitRule(ctx.Block()))
	}
}

func (v *Visitor) VisitSetter(ctx *parser.SetterContext) interface{} {
	if ctx.SEMI() != nil {
		return "set;"
	} else {
		return fmt.Sprintf("set %s", v.visitRule(ctx.Block()))
	}
}

func (v *Visitor) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) interface{} {
	return fmt.Sprintf("%s%s %s\n", v.visitRule(ctx.QualifiedName()), v.visitRule(ctx.FormalParameters()), v.visitRule(ctx.Block()).(string))
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	statements := []string{}
	for _, stmt := range ctx.AllStatement() {
		statements = append(statements, v.visitRule(stmt).(string))
	}
	return fmt.Sprintf("{\n%s}", indent(strings.Join(statements, "\n")))
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.GetChild(0) == nil {
		return "NIL STATEMENT?" + ctx.GetText()
	}
	if errNode, ok := ctx.GetChild(0).(antlr.ErrorNode); ok {
		return fmt.Sprintf("ERROR: %+v", errNode)
	}
	child := ctx.GetChild(0).(antlr.RuleNode)
	switch s := child.(type) {
	case *parser.BlockContext:
		return v.visitRule(s)
	case *parser.IfStatementContext:
		return v.visitRule(s)
	case *parser.ForStatementContext:
		return v.visitRule(s)
	case *parser.ExpressionStatementContext:
		return fmt.Sprintf("%s;", v.visitRule(s))
	case *parser.ReturnStatementContext:
		return fmt.Sprintf("%s", v.visitRule(s))
	case *parser.LocalVariableDeclarationStatementContext:
		return fmt.Sprintf("%s", v.visitRule(s))
	}
	return fmt.Sprintf("UNHANDLED STATEMENT: %T %s", ctx.GetChild(0).(antlr.RuleNode), ctx.GetText())
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	elseStatement := ""
	if ctx.ELSE() != nil {
		if block := ctx.Statement(1).Block(); block != nil {
			elseStatement = " else " + v.visitRule(ctx.Statement(1)).(string)
		} else if ifStatement := ctx.Statement(1).IfStatement(); ifStatement != nil {
			elseStatement = fmt.Sprintf(" else %s", v.visitRule(ifStatement))
		} else {
			elseStatement = fmt.Sprintf(" else {\n%s}", indent(v.visitRule(ctx.Statement(1)).(string)))
		}
	}
	if block := ctx.Statement(0).Block(); block != nil {
		return fmt.Sprintf("if %s %s%s", v.visitRule(ctx.ParExpression()),
			v.visitRule(ctx.Statement(0)),
			elseStatement)
	} else {
		return fmt.Sprintf("if %s {\n%s}%s", v.visitRule(ctx.ParExpression()),
			v.visitRule(ctx.Statement(0)),
			elseStatement)
	}
}

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	if statement := ctx.Statement(); statement != nil {
		if statement.Block() != nil {
			return fmt.Sprintf("for (%s) %s", v.visitRule(ctx.ForControl()), v.visitRule(ctx.Statement()))
		} else {
			return fmt.Sprintf("for (%s) {\n%s}\n", v.visitRule(ctx.ForControl()), indent(v.visitRule(ctx.Statement()).(string)))
		}
	} else {
		return fmt.Sprintf("for (%s);", v.visitRule(ctx.ForControl()))
	}
}

func (v *Visitor) VisitForControl(ctx *parser.ForControlContext) interface{} {
	if enhancedForControl := ctx.EnhancedForControl(); enhancedForControl != nil {
		return v.visitRule(enhancedForControl)
	}
	parts := []string{}
	if forInit := ctx.ForInit(); forInit != nil {
		parts = append(parts, v.visitRule(forInit).(string))
	}
	if expression := ctx.Expression(); expression != nil {
		parts = append(parts, v.visitRule(expression).(string))
	}
	if forUpdate := ctx.ForUpdate(); forUpdate != nil {
		parts = append(parts, v.visitRule(forUpdate).(string))
	}
	return strings.Join(parts, "; ")
}

func (v *Visitor) VisitEnhancedForControl(ctx *parser.EnhancedForControlContext) interface{} {
	return fmt.Sprintf("%s %s : %s", v.visitRule(ctx.TypeRef()), v.visitRule(ctx.Id()), v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitForInit(ctx *parser.ForInitContext) interface{} {
	return v.visitRule(ctx.GetChild(0).(antlr.RuleNode))
}

func (v *Visitor) VisitForUpdate(ctx *parser.ForUpdateContext) interface{} {
	return v.visitRule(ctx.ExpressionList())
}

func (v *Visitor) VisitLocalVariableDeclarationStatement(ctx *parser.LocalVariableDeclarationStatementContext) interface{} {
	return fmt.Sprintf("%s;", v.visitRule(ctx.LocalVariableDeclaration()))
}

func (v *Visitor) VisitLocalVariableDeclaration(ctx *parser.LocalVariableDeclarationContext) interface{} {
	return fmt.Sprintf("%s%s %s", v.Modifiers(ctx.AllModifier()), v.visitRule(ctx.TypeRef()), v.visitRule(ctx.VariableDeclarators()))
}

func (v *Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return fmt.Sprintf("return %s;", v.visitRule(e))
	}
	return "return;"
}

func (v *Visitor) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return fmt.Sprintf("(%s)", v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *Visitor) VisitAssignExpression(ctx *parser.AssignExpressionContext) interface{} {
	assignmentToken := ctx.GetChild(1).(antlr.TerminalNode)
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), assignmentToken.GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitCondExpression(ctx *parser.CondExpressionContext) interface{} {
	return fmt.Sprintf("%s ? %s : %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)), v.visitRule(ctx.Expression(2)))
}

func (v *Visitor) VisitLogAndExpression(ctx *parser.LogAndExpressionContext) interface{} {
	// TODO: Wrap long expressions
	return fmt.Sprintf("%s && %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitLogOrExpression(ctx *parser.LogOrExpressionContext) interface{} {
	// TODO: Wrap long expressions
	return fmt.Sprintf("%s || %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	return fmt.Sprintf("%s & %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	return fmt.Sprintf("%s | %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitBitNotExpression(ctx *parser.BitNotExpressionContext) interface{} {
	return fmt.Sprintf("%s ^ %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitBitExpression(ctx *parser.BitExpressionContext) interface{} {
	return fmt.Sprintf("TODO: IMPLEMENT BIT EXPRESSION")
}

func (v *Visitor) VisitArth1Expression(ctx *parser.Arth1ExpressionContext) interface{} {
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), ctx.GetChild(1).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitArth2Expression(ctx *parser.Arth2ExpressionContext) interface{} {
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), ctx.GetChild(1).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitNegExpression(ctx *parser.NegExpressionContext) interface{} {
	return fmt.Sprintf("%s%s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitPreOpExpression(ctx *parser.PreOpExpressionContext) interface{} {
	return fmt.Sprintf("%s%s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitPostOpExpression(ctx *parser.PostOpExpressionContext) interface{} {
	return fmt.Sprintf("%s%s", v.visitRule(ctx.Expression()), ctx.GetChild(1).(antlr.TerminalNode).GetText())
}

func (v *Visitor) VisitSubExpression(ctx *parser.SubExpressionContext) interface{} {
	return fmt.Sprintf("(%s)", v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	return fmt.Sprintf("(%s)%s", v.visitRule(ctx.TypeRef()), v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitNewInstanceExpression(ctx *parser.NewInstanceExpressionContext) interface{} {
	return fmt.Sprintf("new %s", v.visitRule(ctx.Creator()))
}

func (v *Visitor) VisitArrayExpression(ctx *parser.ArrayExpressionContext) interface{} {
	return fmt.Sprintf("%s[%s]", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitDotExpression(ctx *parser.DotExpressionContext) interface{} {
	expr := v.visitRule(ctx.Expression())
	dot := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	switch {
	case ctx.DotMethodCall() != nil:
		return fmt.Sprintf("%s%s%s", expr, dot, v.visitRule(ctx.DotMethodCall()))
	case ctx.AnyId() != nil:
		return fmt.Sprintf("%s%s%s", expr, dot, v.visitRule(ctx.AnyId()))
	}
	return ""
}

func (v *Visitor) VisitDotMethodCall(ctx *parser.DotMethodCallContext) interface{} {
	expressionList := ""
	if l := ctx.ExpressionList(); l != nil {
		expressionList = v.visitRule(l).(string)
	}
	return fmt.Sprintf("%s(%s)", v.visitRule(ctx.AnyId()), expressionList)
}

func (v *Visitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	expressions := []string{}
	for _, p := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(p).(string))
	}
	return strings.Join(expressions, ", ")
}

func (v *Visitor) VisitAnyId(ctx *parser.AnyIdContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	switch e := ctx.Primary().(type) {
	case *parser.ThisPrimaryContext:
		return "this"
	case *parser.SuperPrimaryContext:
		return "super"
	case *parser.LiteralPrimaryContext:
		return e.GetText()
	case *parser.TypeRefPrimaryContext:
		return fmt.Sprintf("%s.class", v.visitRule(e))
	case *parser.IdPrimaryContext:
		return e.GetText()
	case *parser.SoqlPrimaryContext:
		return v.visitRule(e)
	case *parser.SoslPrimaryContext:
		return v.visitRule(e)
	default:
		return fmt.Sprintf("UNHANDLED PRIMARY EXPRESSION: %T %s", e, e.GetText())
	}
}

func (v *Visitor) VisitMethodCallExpression(ctx *parser.MethodCallExpressionContext) interface{} {
	return v.visitRule(ctx.MethodCall())
}

func (v *Visitor) VisitMethodCall(ctx *parser.MethodCallContext) interface{} {
	var f string
	switch e := ctx.GetChild(0).(type) {
	case *parser.IdContext:
		f = e.GetText()
	case antlr.TerminalNode:
		f = strings.ToLower(e.GetText())
	}
	expressionList := ""
	if el := ctx.ExpressionList(); el != nil {
		expressionList = v.visitRule(el).(string)
	}
	return fmt.Sprintf("%s(%s)", f, expressionList)
}

func (v *Visitor) VisitSoslPrimary(ctx *parser.SoslPrimaryContext) interface{} {
	return fmt.Sprintf("TODO: IMPLEMENT SOSL PRIMARY")
}

func (v *Visitor) VisitSoqlPrimary(ctx *parser.SoqlPrimaryContext) interface{} {
	return v.visitRule(ctx.SoqlLiteral())
}

func (v *Visitor) VisitSoqlLiteral(ctx *parser.SoqlLiteralContext) interface{} {
	return fmt.Sprintf("[\n%s]", indent(v.visitRule(ctx.Query()).(string)))
}

func (v *Visitor) VisitQuery(ctx *parser.QueryContext) interface{} {
	usingScope := ""
	if scope := ctx.UsingScope(); scope != nil {
		usingScope = fmt.Sprintf("\n%s", v.visitRule(scope).(string))
	}
	whereClause := ""
	if where := ctx.WhereClause(); where != nil {
		whereClause = fmt.Sprintf("%s", v.visitRule(where).(string))
	}
	return fmt.Sprintf("SELECT\n%sFROM\n%s%s%sTODO: FINISH VisitQuery",
		indent(v.visitRule(ctx.SelectList()).(string)),
		indent(v.visitRule(ctx.FromNameList()).(string)),
		usingScope,
		whereClause,
	)
}

func (v *Visitor) VisitSubQuery(ctx *parser.SubQueryContext) interface{} {
	return fmt.Sprintf("SELECT\n%sTODO: FINISH VisitSubQuery",
		indent(v.visitRule(ctx.SubFieldList()).(string)),
	)
}

func (v *Visitor) VisitFromNameList(ctx *parser.FromNameListContext) interface{} {
	fieldNames := []string{}
	for _, p := range ctx.AllFieldNameAlias() {
		fieldNames = append(fieldNames, v.visitRule(p).(string))
	}
	return strings.Join(fieldNames, ",\n")
}

func (v *Visitor) VisitFieldNameAlias(ctx *parser.FieldNameAliasContext) interface{} {
	soqlId := ""
	if s := ctx.SoqlId(); s != nil {
		soqlId = " " + s.GetText()
	}
	return fmt.Sprintf("%s%s", v.visitRule(ctx.FieldName()), soqlId)
}

func (v *Visitor) VisitSelectList(ctx *parser.SelectListContext) interface{} {
	selectEntries := []string{}
	for _, p := range ctx.AllSelectEntry() {
		selectEntries = append(selectEntries, v.visitRule(p).(string))
	}
	return strings.Join(selectEntries, ",\n")
}

func (v *Visitor) VisitSubFieldList(ctx *parser.SubFieldListContext) interface{} {
	selectEntries := []string{}
	for _, p := range ctx.AllSubFieldEntry() {
		selectEntries = append(selectEntries, v.visitRule(p).(string))
	}
	return strings.Join(selectEntries, ",\n")
}

func (v *Visitor) VisitSelectEntry(ctx *parser.SelectEntryContext) interface{} {
	soqlId := ""
	if s := ctx.SoqlId(); s != nil {
		soqlId = " " + s.GetText()
	}
	switch {
	case ctx.FieldName() != nil:
		return fmt.Sprintf("%s%s", v.visitRule(ctx.FieldName()), soqlId)
	case ctx.SoqlFunction() != nil:
		return fmt.Sprintf("%s%s", v.visitRule(ctx.SoqlFunction()), soqlId)
	case ctx.SubQuery() != nil:
		return fmt.Sprintf("(%s)%s", v.visitRule(ctx.SubQuery()), soqlId)
	case ctx.TypeOf() != nil:
		return fmt.Sprintf("%s", v.visitRule(ctx.TypeOf()))
	}
	panic("Unexpected selectEntry")
}

func (v *Visitor) VisitSubFieldEntry(ctx *parser.SubFieldEntryContext) interface{} {
	soqlId := ""
	if s := ctx.SoqlId(); s != nil {
		soqlId = " " + s.GetText()
	}
	switch {
	case ctx.FieldName() != nil:
		return fmt.Sprintf("%s%s", v.visitRule(ctx.FieldName()), soqlId)
	case ctx.SoqlFunction() != nil:
		return fmt.Sprintf("%s%s", v.visitRule(ctx.SoqlFunction()), soqlId)
	case ctx.TypeOf() != nil:
		return fmt.Sprintf("%s", v.visitRule(ctx.TypeOf()))
	}
	panic("Unexpected selectEntry")
}

func (v *Visitor) VisitFieldName(ctx *parser.FieldNameContext) interface{} {
	ids := []string{}
	for _, t := range ctx.AllSoqlId() {
		ids = append(ids, t.GetText())
	}
	return strings.Join(ids, ".")
}

func (v *Visitor) VisitFieldNameList(ctx *parser.FieldNameListContext) interface{} {
	fieldNames := []string{}
	for _, p := range ctx.AllFieldName() {
		fieldNames = append(fieldNames, v.visitRule(p).(string))
	}
	return strings.Join(fieldNames, ",\n")
}

func (v *Visitor) VisitTypeOf(ctx *parser.TypeOfContext) interface{} {
	whenClauses := []string{}
	for _, w := range ctx.AllWhenClause() {
		whenClauses = append(whenClauses, v.visitRule(w).(string))
	}
	elseClause := ""
	if e := ctx.ElseClause(); e != nil {
		elseClause = fmt.Sprintf("ELSE %s", v.visitRule(e))
	}

	return fmt.Sprintf("TYPEOF %s\n%s\n%sEND",
		v.visitRule(ctx.FieldName()),
		strings.Join(whenClauses, "\n"),
		elseClause,
	)
}

func (v *Visitor) VisitWhenClause(ctx *parser.WhenClauseContext) interface{} {
	return fmt.Sprintf("WHEN\n%sTHEN\n%s", indent(v.visitRule(ctx.FieldName()).(string)), indent(v.visitRule(ctx.FieldNameList()).(string)))
}

func (v *Visitor) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return fmt.Sprintf("WHERE\n%s", indent(v.visitRule(ctx.LogicalExpression()).(string)))
}

func (v *Visitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) interface{} {
	switch {
	case ctx.NOT() != nil:
		return fmt.Sprintf("NOT %s", ctx.ConditionalExpression(0))
	case len(ctx.AllSOQLOR()) > 0:
		conditions := []string{}
		for _, cond := range ctx.AllConditionalExpression() {
			conditions = append(conditions, v.visitRule(cond).(string))
		}
		return strings.Join(conditions, "OR ")
	case len(ctx.AllSOQLAND()) > 0:
		conditions := []string{}
		for _, cond := range ctx.AllConditionalExpression() {
			conditions = append(conditions, v.visitRule(cond).(string))
		}
		return strings.Join(conditions, "AND ")
	default:
		// Only a single condition
		return v.visitRule(ctx.ConditionalExpression(0))
	}
}

func (v *Visitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) interface{} {
	switch {
	case ctx.LogicalExpression() != nil:
		return fmt.Sprintf("(%s)", v.visitRule(ctx.LogicalExpression()))
	case ctx.FieldExpression() != nil:
		return v.visitRule(ctx.FieldExpression())
	}
	panic("Unexpected conditionalExpression")
}

func (v *Visitor) VisitFieldExpression(ctx *parser.FieldExpressionContext) interface{} {
	switch {
	case ctx.FieldName() != nil:
		return fmt.Sprintf("%s %s %s", v.visitRule(ctx.FieldName()), ctx.ComparisonOperator().GetText(), v.visitRule(ctx.Value()))
	case ctx.SoqlFunction() != nil:
		return fmt.Sprintf("%s %s %s", v.visitRule(ctx.SoqlFunction()), ctx.ComparisonOperator().GetText(), v.visitRule(ctx.Value()))
	}
	panic("Unexpected fieldExpression")
}

func (v *Visitor) VisitSoqlFunction(ctx *parser.SoqlFunctionContext) interface{} {
	param := ""
	switch {
	case ctx.FieldName() != nil:
		param = v.visitRule(ctx.FieldName()).(string)
	case ctx.DateFieldName() != nil:
		param = v.visitRule(ctx.DateFieldName()).(string)
	case ctx.SoqlFieldsParameter() != nil:
		param = v.visitRule(ctx.SoqlFieldsParameter()).(string)
	default:
		panic("Unexpected parameter type for soqlFunction")
	}
	ctx.AVG()
	return fmt.Sprintf("%s(%s)", ctx.GetChild(0).(antlr.TerminalNode).GetText(), param)
}

func (v *Visitor) VisitValue(ctx *parser.ValueContext) interface{} {
	return "TODO: IMPLEMENT VisitValue\n"
}

func (v *Visitor) VisitUsingScope(ctx *parser.UsingScopeContext) interface{} {
	return fmt.Sprintf("USING SCOPE %s", ctx.SoqlId().GetText())
}

func (v *Visitor) VisitCreator(ctx *parser.CreatorContext) interface{} {
	return fmt.Sprintf("%s%s", v.visitRule(ctx.CreatedName()), v.visitRule(ctx.GetChild(1).(antlr.RuleNode)))
}

func (v *Visitor) VisitCreatedName(ctx *parser.CreatedNameContext) interface{} {
	namePairs := []string{}
	for _, i := range ctx.AllIdCreatedNamePair() {
		namePairs = append(namePairs, v.visitRule(i).(string))
	}
	return strings.Join(namePairs, ".")
}

func (v *Visitor) VisitIdCreatedNamePair(ctx *parser.IdCreatedNamePairContext) interface{} {
	if typeList := ctx.TypeList(); typeList != nil {
		return fmt.Sprintf("%s<%s>", v.visitRule(ctx.AnyId()), v.visitRule(typeList))
	}
	return v.visitRule(ctx.AnyId())
}

func (v *Visitor) VisitNoRest(ctx *parser.NoRestContext) interface{} {
	return "{}"
}

func (v *Visitor) VisitId(ctx *parser.IdContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitClassCreatorRest(ctx *parser.ClassCreatorRestContext) interface{} {
	return v.visitRule(ctx.Arguments())
}

func (v *Visitor) VisitArrayCreatorRest(ctx *parser.ArrayCreatorRestContext) interface{} {
	if expression := ctx.Expression(); expression != nil {
		return fmt.Sprintf("{ %s }", v.visitRule(expression))
	} else if arrayInitializer := ctx.ArrayInitializer(); arrayInitializer != nil {
		return fmt.Sprintf("{}%s", v.visitRule(arrayInitializer))
	}
	return "{}"
}

func (v *Visitor) VisitMapCreatorRest(ctx *parser.MapCreatorRestContext) interface{} {
	pairs := []string{}
	for _, i := range ctx.AllMapCreatorRestPair() {
		pairs = append(pairs, v.visitRule(i).(string))
	}
	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}

func (v *Visitor) VisitMapCreatorRestPair(ctx *parser.MapCreatorRestPairContext) interface{} {
	return fmt.Sprintf("%s => %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitSetCreatorRest(ctx *parser.SetCreatorRestContext) interface{} {
	expressions := []string{}
	for _, i := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(i).(string))
	}
	return fmt.Sprintf("{ %s }", strings.Join(expressions, ", "))
}

func (v *Visitor) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	return fmt.Sprintf("TODO: IMPLEMENT ARRAY INITIALIZER")
}

func (v *Visitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	if expressionList := ctx.ExpressionList(); expressionList != nil {
		return fmt.Sprintf("(%s)", v.visitRule(expressionList))
	}
	return "()"
}

func (v *Visitor) VisitCmpExpression(ctx *parser.CmpExpressionContext) interface{} {
	cmpToken := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	if ctx.ASSIGN() != nil {
		cmpToken += "="
	}
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), cmpToken, v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	cmpToken := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), cmpToken, v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitInstanceOfExpression(ctx *parser.InstanceOfExpressionContext) interface{} {
	return fmt.Sprintf("%s instanceof %s", v.visitRule(ctx.Expression()), v.visitRule(ctx.TypeRef()))
}

func (v *Visitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	types := []string{}
	for _, p := range ctx.AllTypeRef() {
		types = append(types, v.visitRule(p).(string))
	}
	return strings.Join(types, ", ")
}

func (v *Visitor) VisitFormalParameters(ctx *parser.FormalParametersContext) interface{} {
	params := []string{}
	list := ctx.FormalParameterList()
	if list == nil {
		return "()"
	}
	for _, p := range list.AllFormalParameter() {
		params = append(params, v.visitRule(p).(string))
	}
	val := fmt.Sprintf("(%s)", strings.Join(params, ", "))
	return val
}

func (v *Visitor) Modifiers(ctxs []parser.IModifierContext) string {
	mods := []string{}
	for _, m := range ctxs {
		if m.Annotation() != nil {
			mods = append(mods, fmt.Sprintf("%s\n", v.visitRule(m.Annotation())))
		} else {
			mods = append(mods, fmt.Sprintf("%s ", m.GetText()))
		}
	}
	modifiers := strings.Join(mods, "")
	return modifiers
}

func (v *Visitor) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	args := ""
	if ctx.LPAREN() != nil {
		vals := ""
		if ctx.ElementValuePairs() != nil {
			vals = v.visitRule(ctx.ElementValuePairs()).(string)
		} else {
			vals = v.visitRule(ctx.ElementValue()).(string)
		}
		args = fmt.Sprintf("(%s)", vals)
	}
	return fmt.Sprintf("@%s%s", v.visitRule(ctx.QualifiedName()), args)
}

func (v *Visitor) VisitElementValuePairs(ctx *parser.ElementValuePairsContext) interface{} {
	pairs := []string{v.visitRule(ctx.ElementValuePair()).(string)}
	for _, p := range ctx.AllDelimitedElementValuePair() {
		pairs = append(pairs, v.visitRule(p).(string))
	}
	return strings.Join(pairs, "")
}

func (v *Visitor) VisitDelimitedElementValuePair(ctx *parser.DelimitedElementValuePairContext) interface{} {
	delimiter := " "
	if ctx.COMMA() != nil {
		delimiter = ", "
	}
	return fmt.Sprintf("%s%s", delimiter, v.visitRule(ctx.ElementValuePair()))
}

func (v *Visitor) VisitElementValuePair(ctx *parser.ElementValuePairContext) interface{} {
	return fmt.Sprintf("%s = %s", v.visitRule(ctx.Id()), v.visitRule(ctx.ElementValue()))
}

func (v *Visitor) VisitElementValue(ctx *parser.ElementValueContext) interface{} {
	return v.visitRule(ctx.GetChild(0).(antlr.RuleNode))
}

func (v *Visitor) VisitElementValueArrayInitializer(ctx *parser.ElementValueArrayInitializerContext) interface{} {
	values := []string{}
	for _, val := range ctx.AllElementValue() {
		values = append(values, v.visitRule(val).(string))
	}
	trailingComma := ""
	if ctx.TrailingComma() != nil {
		trailingComma = ","
	}
	return fmt.Sprintf("(%s%s)", strings.Join(values, ", "), trailingComma)
}

func (v *Visitor) VisitFormalParameter(ctx *parser.FormalParameterContext) interface{} {
	return fmt.Sprintf("%s%s %s", v.Modifiers(ctx.AllModifier()), v.visitRule(ctx.TypeRef()), ctx.Id().GetText())
}

func (v *Visitor) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	ids := []string{}
	for _, i := range ctx.AllId() {
		ids = append(ids, i.GetText())
	}
	return strings.Join(ids, ".")
}

func (v *Visitor) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	vars := []string{}
	for _, vd := range ctx.AllVariableDeclarator() {
		vars = append(vars, v.visitRule(vd).(string))
	}
	return strings.Join(vars, ", ")
}

func (v *Visitor) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	decl := ctx.Id().GetText()
	if ctx.Expression() != nil {
		decl = fmt.Sprintf("%s = %s", decl, v.visitRule(ctx.Expression()))
	}
	return decl
}

func (v *Visitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	returnType := "void"
	if ctx.TypeRef() != nil {
		returnType = v.visitRule(ctx.TypeRef()).(string)
	}
	body := ";"
	if ctx.Block() != nil {
		body = " " + v.visitRule(ctx.Block()).(string)
	}
	return fmt.Sprintf("%s %s%s%s\n", returnType, ctx.Id().GetText(),
		v.visitRule(ctx.FormalParameters()),
		body)
}

func (v *Visitor) VisitTypeRef(ctx *parser.TypeRefContext) interface{} {
	typeNames := []string{}
	for _, t := range ctx.AllTypeName() {
		typeNames = append(typeNames, t.GetText())
	}

	val := fmt.Sprintf("%s%s", strings.Join(typeNames, "."), ctx.ArraySubscripts().GetText())
	return val
}
