package formatter

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

func (v *Visitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	if trigger := ctx.TriggerUnit(); trigger != nil {
		return v.visitRule(trigger)
	}
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
		return fmt.Sprintf("enum %s {%s}", v.visitRule(enum.Id()), strings.Join(constants, ", "))
	}
	return ""
}

func (v *Visitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	var class strings.Builder
	class.WriteString(fmt.Sprintf("class %s", v.visitRule(ctx.Id())))
	if ctx.EXTENDS() != nil {
		class.WriteString(fmt.Sprintf(" extends %s", v.visitRule(ctx.TypeRef())))
	}
	if ctx.IMPLEMENTS() != nil {
		class.WriteString(fmt.Sprintf(" implements %s", v.visitRule(ctx.TypeList())))
	}
	class.WriteString(fmt.Sprintf(" {\n%s\n}", indent(v.visitRule(ctx.ClassBody()).(string))))
	return class.String()
}

func (v *Visitor) VisitTriggerUnit(ctx *parser.TriggerUnitContext) interface{} {
	triggerCases := []string{}
	for _, t := range ctx.AllTriggerCase() {
		triggerCases = append(triggerCases, v.visitRule(t).(string))
	}
	return fmt.Sprintf("trigger %s on %s (%s) %s", v.visitRule(ctx.Id(0)), v.visitRule(ctx.Id(1)),
		strings.Join(triggerCases, ", "),
		v.visitRule(ctx.Block()))
}

func (v *Visitor) VisitTriggerCase(ctx *parser.TriggerCaseContext) interface{} {
	return fmt.Sprintf("%s %s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), ctx.GetChild(1).(antlr.TerminalNode).GetText())
}

func (v *Visitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) interface{} {
	enumConstants := ""
	if ctx.EnumConstants() != nil {
		enumConstants = v.visitRule(ctx.EnumConstants()).(string)
	}
	return fmt.Sprintf("enum %s { %s }", v.visitRule(ctx.Id()), enumConstants)
}

func (v *Visitor) VisitEnumConstants(ctx *parser.EnumConstantsContext) interface{} {
	ids := []string{}
	for _, t := range ctx.AllId() {
		ids = append(ids, t.GetText())
	}
	return strings.Join(ids, ", ")
}

func (v *Visitor) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) interface{} {
	extends := ""
	if ctx.EXTENDS() != nil {
		extends = fmt.Sprintf(" extends %s ", v.visitRule(ctx.TypeList()))
	}
	return fmt.Sprintf("interface %s%s {\n%s\n}", ctx.Id().GetText(), extends, indent(v.visitRule(ctx.InterfaceBody()).(string)))
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
		return fmt.Sprintf("%s%s", static, v.visitRule(ctx.Block()).(string))
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
	// TODO: collapse if empty getter and setter
	return fmt.Sprintf("%s %s {\n%s\n}", v.visitRule(ctx.TypeRef()), ctx.Id().GetText(), indent(strings.Join(propertyBlocks, "\n")))
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
	return fmt.Sprintf("%s%s %s", v.visitRule(ctx.QualifiedName()), v.visitRule(ctx.FormalParameters()), v.visitRule(ctx.Block()).(string))
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	statements := []string{}
	for _, stmt := range ctx.AllStatement() {
		statements = append(statements, v.visitRule(stmt).(string))
	}
	return fmt.Sprintf("{\n%s\n}", indent(strings.Join(statements, "\n")))
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	/*
		if ctx.GetChild(0) == nil {
			return "NIL STATEMENT?" + ctx.GetText()
		}
		if errNode, ok := ctx.GetChild(0).(antlr.ErrorNode); ok {
			return fmt.Sprintf("ERROR: %+v", errNode)
		}
	*/
	child := ctx.GetChild(0).(antlr.RuleNode)
	return v.visitRule(child)
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	var out strings.Builder
	if block := ctx.Statement(0).Block(); block != nil {
		out.WriteString(fmt.Sprintf("if %s %s", v.visitRule(ctx.ParExpression()),
			v.visitRule(ctx.Statement(0))))
	} else {
		out.WriteString(fmt.Sprintf("if %s {\n%s\n}", v.visitRule(ctx.ParExpression()),
			indent(v.visitRule(ctx.Statement(0)).(string))))
	}
	if ctx.ELSE() != nil {
		if block := ctx.Statement(1).Block(); block != nil {
			out.WriteString(fmt.Sprintf(" else %s", v.visitRule(ctx.Statement(1)).(string)))
		} else if ifStatement := ctx.Statement(1).IfStatement(); ifStatement != nil {
			out.WriteString(fmt.Sprintf(" else %s", v.visitRule(ifStatement)))
		} else {
			out.WriteString(fmt.Sprintf(" else {\n%s}", indent(v.visitRule(ctx.Statement(1)).(string))))
		}
	}
	return out.String()
}

func (v *Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	if s := ctx.Statement; s == nil {
		return fmt.Sprintf("while %s;", v.visitRule(ctx.ParExpression()))
	}
	if block := ctx.Statement().Block(); block != nil {
		return fmt.Sprintf("while %s %s", v.visitRule(ctx.ParExpression()), v.visitRule(ctx.Statement()))
	} else {
		return fmt.Sprintf("while %s {\n%s}", v.visitRule(ctx.ParExpression()), v.visitRule(ctx.Statement()))
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

func (v *Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	when := []string{}
	for _, w := range ctx.AllWhenControl() {
		when = append(when, v.visitRule(w).(string))
	}
	return fmt.Sprintf("switch on %s {\n%s}", v.visitRule(ctx.Expression()), indent(strings.Join(when, "\n")))
}

func (v *Visitor) VisitWhenControl(ctx *parser.WhenControlContext) interface{} {
	return fmt.Sprintf("when %s %s", v.visitRule(ctx.WhenValue()), v.visitRule(ctx.Block()))
}

func (v *Visitor) VisitWhenValue(ctx *parser.WhenValueContext) interface{} {
	switch {
	case ctx.ELSE() != nil:
		return "else"
	case len(ctx.AllId()) == 2:
		return fmt.Sprintf("%s %s", v.visitRule(ctx.Id(0)), v.visitRule(ctx.Id(1)))
	default:
		whenLiterals := []string{}
		for _, w := range ctx.AllWhenLiteral() {
			whenLiterals = append(whenLiterals, v.visitRule(w).(string))
		}
		return strings.Join(whenLiterals, ", ")
	}
}

func (v *Visitor) VisitWhenLiteral(ctx *parser.WhenLiteralContext) interface{} {
	if w := ctx.WhenLiteral(); w != nil {
		return fmt.Sprintf("(%s)", v.visitRule(w))
	}
	if i := ctx.Id(); i != nil {
		return v.visitRule(i)
	}
	return ctx.GetText()
}

func (v *Visitor) VisitTryStatement(ctx *parser.TryStatementContext) interface{} {
	if len(ctx.AllCatchClause()) > 0 {
		catchClauses := []string{}
		for _, c := range ctx.AllCatchClause() {
			catchClauses = append(catchClauses, v.visitRule(c).(string))
		}
		finally := ""
		if f := ctx.FinallyBlock(); f != nil {
			finally = fmt.Sprintf("\n%s", v.visitRule(f).(string))
		}
		return fmt.Sprintf("try %s %s%s", v.visitRule(ctx.Block()), strings.Join(catchClauses, "\n"), finally)
	} else {
		return fmt.Sprintf("try %s %s", v.visitRule(ctx.Block()), v.visitRule(ctx.FinallyBlock()))
	}
}

func (v *Visitor) VisitCatchClause(ctx *parser.CatchClauseContext) interface{} {
	return fmt.Sprintf("catch (%s%s %s) %s",
		v.Modifiers(ctx.AllModifier()),
		v.visitRule(ctx.QualifiedName()),
		v.visitRule(ctx.Id()),
		v.visitRule(ctx.Block()))
}

func (v *Visitor) VisitFinallyBlock(ctx *parser.FinallyBlockContext) interface{} {
	return fmt.Sprintf("finally %s", v.visitRule(ctx.Block()))
}

func (v *Visitor) VisitThrowStatement(ctx *parser.ThrowStatementContext) interface{} {
	return fmt.Sprintf("throw %s;", v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitRunAsStatement(ctx *parser.RunAsStatementContext) interface{} {
	expressionList := ""
	if e := ctx.ExpressionList(); e != nil {
		expressionList = v.visitRule(e).(string)
	}
	return fmt.Sprintf("System.runAs(%s) %s", expressionList, v.visitRule(ctx.Block()))
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

func (v *Visitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	return "continue;"
}

func (v *Visitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	return "break;"
}

func (v *Visitor) VisitForUpdate(ctx *parser.ForUpdateContext) interface{} {
	return v.visitRule(ctx.ExpressionList())
}

func (v *Visitor) VisitLocalVariableDeclarationStatement(ctx *parser.LocalVariableDeclarationStatementContext) interface{} {
	return fmt.Sprintf("%s;", v.visitRule(ctx.LocalVariableDeclaration()))
}

func (v *Visitor) VisitInsertStatement(ctx *parser.InsertStatementContext) interface{} {
	return fmt.Sprintf("insert %s;", v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitUpdateStatement(ctx *parser.UpdateStatementContext) interface{} {
	return fmt.Sprintf("update %s;", v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitUpsertStatement(ctx *parser.UpsertStatementContext) interface{} {
	return fmt.Sprintf("upsert %s;", v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitMergeStatement(ctx *parser.MergeStatementContext) interface{} {
	return fmt.Sprintf("merge %s %s;", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitDeleteStatement(ctx *parser.DeleteStatementContext) interface{} {
	return fmt.Sprintf("delete %s;", v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitUndeleteStatement(ctx *parser.UndeleteStatementContext) interface{} {
	return fmt.Sprintf("undelete %s;", v.visitRule(ctx.Expression()))
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
	return fmt.Sprintf("%s;", v.visitRule(ctx.Expression()))
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
	ops := []string{}
	for _, o := range ctx.AllGT() {
		ops = append(ops, o.GetText())
	}
	for _, o := range ctx.AllLT() {
		ops = append(ops, o.GetText())
	}
	return strings.Join(ops, "")
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
	// TODO: Figure out rule(s) for wrapping
	if len(expressions) > 3 {
		return indent("\n" + strings.Join(expressions, ",\n"))
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
		return v.visitRule(e)
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
	return v.visitRule(ctx.SoslLiteral())
}

func (v *Visitor) VisitSoqlPrimary(ctx *parser.SoqlPrimaryContext) interface{} {
	return v.visitRule(ctx.SoqlLiteral())
}

func (v *Visitor) VisitSoqlLiteral(ctx *parser.SoqlLiteralContext) interface{} {
	return fmt.Sprintf("[\n%s\n]", indent(v.visitRule(ctx.Query()).(string)))
}

func (v *Visitor) VisitQuery(ctx *parser.QueryContext) interface{} {
	var query strings.Builder
	query.WriteString(fmt.Sprintf("SELECT\n%s\nFROM\n%s",
		indent(v.visitRule(ctx.SelectList()).(string)),
		indent(v.visitRule(ctx.FromNameList()).(string))))
	if scope := ctx.UsingScope(); scope != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(scope).(string)))
	}
	if where := ctx.WhereClause(); where != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(where).(string)))
	}
	if groupBy := ctx.GroupByClause(); groupBy != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(groupBy).(string)))
	}
	if orderBy := ctx.OrderByClause(); orderBy != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(orderBy).(string)))
	}
	if limit := ctx.LimitClause(); limit != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(limit).(string)))
	}
	if offset := ctx.OffsetClause(); offset != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(offset).(string)))
	}
	if ctx.OffsetClause() != nil {
		query.WriteString("\nALL ROWS")
	}
	forClauses := v.visitRule(ctx.ForClauses())
	if forClauses != "" {
		query.WriteString(fmt.Sprintf("\n%s", forClauses))
	}
	if update := ctx.UpdateList(); update != nil {
		query.WriteString(fmt.Sprintf("\nUPDATE %s", v.visitRule(update).(string)))
	}
	return query.String()
}

func (v *Visitor) VisitSubQuery(ctx *parser.SubQueryContext) interface{} {
	var query strings.Builder
	query.WriteString(fmt.Sprintf("SELECT\n%s\nFROM\n%s",
		indent(v.visitRule(ctx.SubFieldList()).(string)),
		indent(v.visitRule(ctx.FromNameList()).(string)),
	))
	if where := ctx.WhereClause(); where != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(where).(string)))
	}
	if orderBy := ctx.OrderByClause(); orderBy != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(orderBy).(string)))
	}
	if limit := ctx.LimitClause(); limit != nil {
		query.WriteString(fmt.Sprintf("\n%s", v.visitRule(limit).(string)))
	}
	forClauses := v.visitRule(ctx.ForClauses())
	if forClauses != "" {
		query.WriteString(fmt.Sprintf("\n%s", forClauses))
	}
	if update := ctx.UpdateList(); update != nil {
		query.WriteString(fmt.Sprintf("\nUPDATE %s", v.visitRule(update).(string)))
	}
	return query.String()
}

func (v *Visitor) VisitFromNameList(ctx *parser.FromNameListContext) interface{} {
	fieldNames := []string{}
	for _, p := range ctx.AllFieldNameAlias() {
		fieldNames = append(fieldNames, v.visitRule(p).(string))
	}
	return strings.Join(fieldNames, ",\n")
}

func (v *Visitor) VisitUpdateList(ctx *parser.UpdateListContext) interface{} {
	updateList := ""
	if u := ctx.UpdateList(); u != nil {
		updateList = fmt.Sprintf(", %s", v.visitRule(u).(string))
	}
	return fmt.Sprintf("%s%s", ctx.UpdateType().GetText(), updateList)
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

func (v *Visitor) VisitForClauses(ctx *parser.ForClausesContext) interface{} {
	forClauses := []string{}
	for _, f := range ctx.AllForClause() {
		forClauses = append(forClauses, v.visitRule(f).(string))
	}
	return strings.Join(forClauses, " ")
}

func (v *Visitor) VisitForClause(ctx *parser.ForClauseContext) interface{} {
	return fmt.Sprintf("FOR %s", ctx.GetChild(1).(antlr.TerminalNode).GetText())
}

func (v *Visitor) VisitWhenClause(ctx *parser.WhenClauseContext) interface{} {
	return fmt.Sprintf("WHEN\n%s\nTHEN\n%s", indent(v.visitRule(ctx.FieldName()).(string)), indent(v.visitRule(ctx.FieldNameList()).(string)))
}

func (v *Visitor) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return fmt.Sprintf("WHERE\n%s", indent(v.visitRule(ctx.LogicalExpression()).(string)))
}

func (v *Visitor) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	if e := ctx.BoundExpression(); e != nil {
		return fmt.Sprintf("LIMIT %s", v.visitRule(ctx.BoundExpression()))
	}
	return fmt.Sprintf("LIMIT %s", ctx.IntegerLiteral().GetText())
}

func (v *Visitor) VisitOffsetClause(ctx *parser.OffsetClauseContext) interface{} {
	if e := ctx.BoundExpression(); e != nil {
		return fmt.Sprintf("OFFSET %s", v.visitRule(ctx.BoundExpression()))
	}
	return fmt.Sprintf("OFFSET %s", ctx.IntegerLiteral().GetText())
}

func (v *Visitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) interface{} {
	switch {
	case ctx.NOT() != nil:
		return fmt.Sprintf("NOT %s", v.visitRule(ctx.ConditionalExpression(0)))
	case len(ctx.AllSOQLOR()) > 0:
		conditions := []string{}
		for _, cond := range ctx.AllConditionalExpression() {
			conditions = append(conditions, v.visitRule(cond).(string))
		}
		return strings.Join(conditions, " OR\n")
	case len(ctx.AllSOQLAND()) > 0:
		conditions := []string{}
		for _, cond := range ctx.AllConditionalExpression() {
			conditions = append(conditions, v.visitRule(cond).(string))
		}
		return strings.Join(conditions, " AND\n")
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
		// TODO: Format IN/NOT IN
		return fmt.Sprintf("%s %s %s", v.visitRule(ctx.FieldName()), v.visitRule(ctx.ComparisonOperator()), v.visitRule(ctx.Value()))
	case ctx.SoqlFunction() != nil:
		return fmt.Sprintf("%s %s %s", v.visitRule(ctx.SoqlFunction()), v.visitRule(ctx.ComparisonOperator()), v.visitRule(ctx.Value()))
	}
	panic("Unexpected fieldExpression")
}

func (v *Visitor) VisitComparisonOperator(ctx *parser.ComparisonOperatorContext) interface{} {
	if ctx.NOT() != nil {
		return "NOT IN"
	}
	return ctx.GetText()
}

func (v *Visitor) VisitSoqlFunction(ctx *parser.SoqlFunctionContext) interface{} {
	param := ""
	switch {
	case ctx.COUNT() != nil:
		return "COUNT()"
	case ctx.FieldName() != nil:
		param = v.visitRule(ctx.FieldName()).(string)
	case ctx.DateFieldName() != nil:
		param = v.visitRule(ctx.DateFieldName()).(string)
	case ctx.SoqlFieldsParameter() != nil:
		param = v.visitRule(ctx.SoqlFieldsParameter()).(string)
	default:
		panic("Unexpected parameter type for soqlFunction")
	}
	return fmt.Sprintf("%s(%s)", ctx.GetChild(0).(antlr.TerminalNode).GetText(), param)
}

func (v *Visitor) VisitSoqlFieldsParameter(ctx *parser.SoqlFieldsParameterContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitNullValue(ctx *parser.NullValueContext) interface{} {
	return "null"
}

func (v *Visitor) VisitBooleanLiteralValue(ctx *parser.BooleanLiteralValueContext) interface{} {
	return strings.ToLower(ctx.GetText())
}

func (v *Visitor) VisitSignedNumberValue(ctx *parser.SignedNumberValueContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitStringLiteralValue(ctx *parser.StringLiteralValueContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitDateLiteralValue(ctx *parser.DateLiteralValueContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitDateTimeLiteralValue(ctx *parser.DateTimeLiteralValueContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitDateFormulaValue(ctx *parser.DateFormulaValueContext) interface{} {
	return v.visitRule(ctx.DateFormula())
}

func (v *Visitor) VisitCurrencyValueValue(ctx *parser.CurrencyValueValueContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitSubQueryValue(ctx *parser.SubQueryValueContext) interface{} {
	return fmt.Sprintf("(%s)", v.visitRule(ctx.SubQuery()))
}

func (v *Visitor) VisitValueListValue(ctx *parser.ValueListValueContext) interface{} {
	return v.visitRule(ctx.ValueList())
}

func (v *Visitor) VisitBoundExpressionValue(ctx *parser.BoundExpressionValueContext) interface{} {
	return v.visitRule(ctx.BoundExpression())
}

func (v *Visitor) VisitDateFormula(ctx *parser.DateFormulaContext) interface{} {
	if ctx.SignedInteger() != nil {
		return fmt.Sprintf("%s:%s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), v.visitRule(ctx.SignedInteger()))
	}
	return ctx.GetChild(0).(antlr.TerminalNode).GetText()
}

func (v *Visitor) VisitSignedInteger(ctx *parser.SignedIntegerContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitSignedNumber(ctx *parser.SignedNumberContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitValueList(ctx *parser.ValueListContext) interface{} {
	values := []string{}
	for _, i := range ctx.AllValue() {
		values = append(values, v.visitRule(i).(string))
	}
	return fmt.Sprintf("(%s)", strings.Join(values, ", "))
}

func (v *Visitor) VisitGroupByClause(ctx *parser.GroupByClauseContext) interface{} {
	fieldNames := []string{}
	for _, i := range ctx.AllFieldName() {
		fieldNames = append(fieldNames, v.visitRule(i).(string))
	}
	switch {
	case ctx.ROLLUP() != nil:
		return fmt.Sprintf("GROUP BY ROLLUP (%s)", strings.Join(fieldNames, ", "))
	case ctx.CUBE() != nil:
		return fmt.Sprintf("GROUP BY CUBE (%s)", strings.Join(fieldNames, ", "))
	default:
		having := ""
		if l := ctx.LogicalExpression(); l != nil {
			having = fmt.Sprintf("HAVING %s", v.visitRule(l))
		}
		return fmt.Sprintf("GROUP BY %s%s", v.visitRule(ctx.SelectList()), having)
	}
}

func (v *Visitor) VisitUsingScope(ctx *parser.UsingScopeContext) interface{} {
	return fmt.Sprintf("USING SCOPE %s", ctx.SoqlId().GetText())
}

func (v *Visitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) interface{} {
	return fmt.Sprintf("ORDER BY\n%s", indent(v.visitRule(ctx.FieldOrderList()).(string)))
}

func (v *Visitor) VisitFieldOrderList(ctx *parser.FieldOrderListContext) interface{} {
	fields := []string{}
	for _, i := range ctx.AllFieldOrder() {
		fields = append(fields, v.visitRule(i).(string))
	}
	return strings.Join(fields, ", ")
}

func (v *Visitor) VisitFieldOrder(ctx *parser.FieldOrderContext) interface{} {
	var field strings.Builder
	if f := ctx.FieldName(); f != nil {
		field.WriteString(v.visitRule(f).(string))
	} else if s := ctx.SoqlFunction(); s != nil {
		field.WriteString(v.visitRule(s).(string))
	}
	if ctx.ASC() != nil {
		field.WriteString(" ASC")
	} else if ctx.DESC() != nil {
		field.WriteString(" DESC")
	}
	if ctx.NULLS() != nil {
		field.WriteString(" NULLS")
		if ctx.FIRST() != nil {
			field.WriteString(" FIRST")
		} else {
			field.WriteString(" LAST")
		}
	}
	return field.String()
}

func (v *Visitor) VisitBoundExpression(ctx *parser.BoundExpressionContext) interface{} {
	return fmt.Sprintf(":%s", v.visitRule(ctx.Expression()))
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
		return fmt.Sprintf("[ %s ]", v.visitRule(expression))
	} else if arrayInitializer := ctx.ArrayInitializer(); arrayInitializer != nil {
		return fmt.Sprintf("[]%s", v.visitRule(arrayInitializer))
	}
	return "[]"
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
	expressions := []string{}
	for _, i := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(i).(string))
	}
	return fmt.Sprintf("{ %s }", strings.Join(expressions, ", "))
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
	return fmt.Sprintf("%s %s%s%s", returnType, ctx.Id().GetText(),
		v.visitRule(ctx.FormalParameters()),
		body)
}

func (v *Visitor) VisitTypeRefPrimary(ctx *parser.TypeRefPrimaryContext) interface{} {
	return fmt.Sprintf("%s.class", v.visitRule(ctx.TypeRef()))
}

func (v *Visitor) VisitTypeRef(ctx *parser.TypeRefContext) interface{} {
	typeNames := []string{}
	for _, t := range ctx.AllTypeName() {
		typeNames = append(typeNames, v.visitRule(t).(string))
	}

	val := fmt.Sprintf("%s%s", strings.Join(typeNames, "."), ctx.ArraySubscripts().GetText())
	return val
}

func (v *Visitor) VisitTypeName(ctx *parser.TypeNameContext) interface{} {
	typeName := ""
	if id := ctx.Id(); id != nil {
		typeName = v.visitRule(id).(string)
	} else {
		typeName = ctx.GetChild(0).(antlr.TerminalNode).GetText()
	}
	typeArguments := ""
	if args := ctx.TypeArguments(); args != nil {
		typeArguments = v.visitRule(args).(string)
	}
	return fmt.Sprintf("%s%s", typeName, typeArguments)
}

func (v *Visitor) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	return fmt.Sprintf("<%s>", v.visitRule(ctx.TypeList()))
}

func (v *Visitor) VisitSoslLiteral(ctx *parser.SoslLiteralContext) interface{} {
	if ctx.BoundExpression() != nil {
		return fmt.Sprintf("[\n%s]",
			indent(fmt.Sprintf("FIND\n%s%s", indent(v.visitRule(ctx.BoundExpression()).(string)), v.visitRule(ctx.SoslClauses()))),
		)
	}
	return fmt.Sprintf("%s %s ]", ctx.GetChild(0).(antlr.TerminalNode).GetText(), v.visitRule(ctx.SoslClauses()))
}

func (v *Visitor) VisitSoslClauses(ctx *parser.SoslClausesContext) interface{} {
	var clauses strings.Builder
	if i := ctx.InSearchGroup(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.ReturningFieldSpecList(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.WithDivisionAssign(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.WithDataCategory(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.WithSnippet(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.WithNetworkIn(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.WithNetworkAssign(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.WithPricebookIdAssign(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.WithMetadataAssign(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.LimitClause(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.UpdateListClause(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	return clauses.String()
}

func (v *Visitor) VisitInSearchGroup(ctx *parser.InSearchGroupContext) interface{} {
	return fmt.Sprintf("IN %s", v.visitRule(ctx.SearchGroup()))
}

func (v *Visitor) VisitSearchGroup(ctx *parser.SearchGroupContext) interface{} {
	return fmt.Sprintf("%s FIELDS", strings.ToUpper(ctx.GetChild(0).(antlr.TerminalNode).GetText()))
}

func (v *Visitor) VisitReturningFieldSpecList(ctx *parser.ReturningFieldSpecListContext) interface{} {
	return fmt.Sprintf("RETURNING %s", v.visitRule(ctx.FieldSpecList()))
}

func (v *Visitor) VisitFieldSpecList(ctx *parser.FieldSpecListContext) interface{} {
	list := []string{v.visitRule(ctx.FieldSpec()).(string)}
	for _, f := range ctx.AllFieldSpecList() {
		list = append(list, v.visitRule(f).(string))
	}
	return strings.Join(list, ",\n")
}

func (v *Visitor) VisitFieldSpec(ctx *parser.FieldSpecContext) interface{} {
	if ctx.FieldSpecClauses() == nil {
		return v.visitRule(ctx.SoslId())
	}
	return fmt.Sprintf("%s%s", v.visitRule(ctx.SoslId()), v.visitRule(ctx.FieldSpecClauses()))
}

func (v *Visitor) VisitFieldSpecClauses(ctx *parser.FieldSpecClausesContext) interface{} {
	var clauses strings.Builder
	clauses.WriteString(fmt.Sprintf("(\n%s", indent(v.visitRule(ctx.FieldList()).(string))))
	if i := ctx.LogicalExpression(); i != nil {
		clauses.WriteString(fmt.Sprintf("\nWHERE\n%s", indent(v.visitRule(i).(string))))
	}
	if i := ctx.SoslId(); i != nil {
		clauses.WriteString(fmt.Sprintf("\nUSING LISTVIEW =  %s", v.visitRule(i)))
	}
	if i := ctx.FieldOrderList(); i != nil {
		clauses.WriteString(fmt.Sprintf("\nORDER BY %s", v.visitRule(i)))
	}
	if i := ctx.LimitClause(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	if i := ctx.OffsetClause(); i != nil {
		clauses.WriteString(fmt.Sprintf("\n%s", v.visitRule(i)))
	}
	clauses.WriteString(")")
	return clauses.String()
}

func (v *Visitor) VisitFieldList(ctx *parser.FieldListContext) interface{} {
	list := []string{v.visitRule(ctx.SoslId()).(string)}
	for _, f := range ctx.AllFieldList() {
		list = append(list, v.visitRule(f).(string))
	}
	return strings.Join(list, ",\n")
}

func (v *Visitor) VisitSoslId(ctx *parser.SoslIdContext) interface{} {
	list := []string{v.visitRule(ctx.Id()).(string)}
	for _, f := range ctx.AllSoslId() {
		list = append(list, v.visitRule(f).(string))
	}
	return strings.Join(list, ".")
}
