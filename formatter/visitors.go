package formatter

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
	log "github.com/sirupsen/logrus"
)

func (v *FormatVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	if trigger := ctx.TriggerUnit(); trigger != nil {
		return v.visitRule(trigger)
	}
	t := ctx.TypeDeclaration()
	switch {
	case t.ClassDeclaration() != nil:
		return v.Modifiers(t.AllModifier()) + v.visitRule(t.ClassDeclaration()).(string)
	case t.InterfaceDeclaration() != nil:
		return v.Modifiers(t.AllModifier()) + v.visitRule(t.InterfaceDeclaration()).(string)
	case t.EnumDeclaration() != nil:
		return v.Modifiers(t.AllModifier()) + v.visitRule(t.EnumDeclaration()).(string)
	}
	return ""
}

func (v *FormatVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	var class strings.Builder
	class.WriteString("class " + v.visitRule(ctx.Id()).(string))
	if ctx.EXTENDS() != nil {
		class.WriteString(" extends " + v.visitRule(ctx.TypeRef()).(string))
	}
	if ctx.IMPLEMENTS() != nil {
		class.WriteString(" implements " + v.visitRule(ctx.TypeList()).(string))
	}
	class.WriteString(" " + v.visitRule(ctx.ClassBody()).(string))
	return class.String()
}

func (v *FormatVisitor) VisitTriggerUnit(ctx *parser.TriggerUnitContext) interface{} {
	triggerCases := []string{}
	for _, t := range ctx.AllTriggerCase() {
		triggerCases = append(triggerCases, v.visitRule(t).(string))
	}
	return "trigger " + v.visitRule(ctx.Id(0)).(string) + " on " + v.visitRule(ctx.Id(1)).(string) + " (" +
		strings.Join(triggerCases, ", ") + ") " +
		v.visitRule(ctx.TriggerBlock()).(string)
}

func (v *FormatVisitor) VisitTriggerBlock(ctx *parser.TriggerBlockContext) interface{} {
	statements := []string{}
	for _, stmt := range ctx.AllTriggerStatement() {
		statements = append(statements, v.visitRule(stmt).(string))
	}
	return "{\n" + indent(strings.Join(statements, "\n")) + "\n}"
}

func (v *FormatVisitor) VisitTriggerStatement(ctx *parser.TriggerStatementContext) interface{} {
	return v.visitRule(ctx.GetChild(0).(antlr.RuleNode))
}

func (v *FormatVisitor) VisitTriggerCase(ctx *parser.TriggerCaseContext) interface{} {
	return ctx.GetChild(0).(antlr.TerminalNode).GetText() + " " + ctx.GetChild(1).(antlr.TerminalNode).GetText()
}

func (v *FormatVisitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) interface{} {
	if ctx.EnumConstants() == nil {
		return "enum " + v.visitRule(ctx.Id()).(string) + " {}"
	}
	enumConstants := v.visitRule(ctx.EnumConstants()).(string)
	if strings.Contains(enumConstants, "\n") {
		return "enum " + v.visitRule(ctx.Id()).(string) + " {\n" + indent(enumConstants) + "\n}"
	}
	if v.wrap {
		return "enum " + v.visitRule(ctx.Id()).(string) + " {\n" + indent(enumConstants) + "\n}"
	}
	return "enum " + v.visitRule(ctx.Id()).(string) + " { " + enumConstants + " }"
}

func (v *FormatVisitor) VisitEnumConstants(ctx *parser.EnumConstantsContext) interface{} {
	ids := []string{}
	for _, i := range ctx.AllId() {
		ids = append(ids, v.visitRule(i).(string))
	}
	if v.wrap {
		return indent(strings.Join(ids, ",\n"))
	}
	return strings.Join(ids, ", ")
}

func (v *FormatVisitor) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) interface{} {
	extends := ""
	if ctx.EXTENDS() != nil {
		extends = " extends " + v.visitRule(ctx.TypeList()).(string) + " "
	}
	return "interface " + ctx.Id().GetText() + extends + " {\n" + indent(v.visitRule(ctx.InterfaceBody()).(string)) + "\n}"
}

func (v *FormatVisitor) VisitInterfaceBody(ctx *parser.InterfaceBodyContext) interface{} {
	declarations := []string{}
	for _, d := range ctx.AllInterfaceMethodDeclaration() {
		declarations = append(declarations, v.visitRule(d).(string))
	}
	return strings.Join(declarations, "\n")
}

func (v *FormatVisitor) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	var cb []string
	declarations := ctx.AllClassBodyDeclaration()
	if len(declarations) == 0 {
		return "{}"
	}
	for _, b := range declarations {
		cb = append(cb, v.visitRule(b).(string))
	}
	return "{\n" + indent(strings.Join(cb, "\n")) + "\n}"
}

func (v *FormatVisitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) interface{} {
	switch {
	case ctx.SEMI() != nil:
		return ";"
	case ctx.Block() != nil:
		static := ""
		if ctx.STATIC() != nil {
			static = "static "
		}
		return static + v.visitRule(ctx.Block()).(string)
	case ctx.MemberDeclaration() != nil:
		return v.Modifiers(ctx.AllModifier()) + v.visitRule(ctx.MemberDeclaration()).(string)
	}
	return ""
}

func (v *FormatVisitor) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	return v.visitRule(ctx.GetChild(0).(antlr.RuleNode))
}

func (v *FormatVisitor) VisitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) interface{} {
	returnType := "void"
	if ctx.TypeRef() != nil {
		returnType = v.visitRule(ctx.TypeRef()).(string)
	}
	return v.Modifiers(ctx.AllModifier()) + returnType + " " + ctx.MethodId().GetText() + v.visitRule(ctx.FormalParameters()).(string) + ";"
}

func (v *FormatVisitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	return v.visitRule(ctx.TypeRef()).(string) + " " + v.visitRule(ctx.VariableDeclarators()).(string) + ";"
}

func (v *FormatVisitor) VisitPropertyDeclaration(ctx *parser.PropertyDeclarationContext) interface{} {
	propertyBlocks := []string{}
	if ctx.AllPropertyBlock() != nil {
		for _, p := range ctx.AllPropertyBlock() {
			propertyBlocks = append(propertyBlocks, v.visitRule(p).(string))
		}
	}
	// Flatten empty getter/setter
	if len(strings.Join(propertyBlocks, "")) == 8 {
		return v.visitRule(ctx.TypeRef()).(string) + " " + ctx.Id().GetText() + " {" + strings.Join(propertyBlocks, " ") + "}"
	}
	sep := "\n"
	return v.visitRule(ctx.TypeRef()).(string) + " " + ctx.Id().GetText() + " {" + sep + indent(strings.Join(propertyBlocks, sep)) + sep + "}"
}

func (v *FormatVisitor) VisitPropertyBlock(ctx *parser.PropertyBlockContext) interface{} {
	if v.textLen(ctx) > 40 {
		defer restoreWrap(wrap(v))
	}
	if ctx.Getter() != nil {
		return v.Modifiers(ctx.AllModifier()) + v.visitRule(ctx.Getter()).(string)
	} else {
		return v.Modifiers(ctx.AllModifier()) + v.visitRule(ctx.Setter()).(string)
	}
}

func (v *FormatVisitor) VisitGetter(ctx *parser.GetterContext) interface{} {
	if ctx.SEMI() != nil {
		return "get;"
	} else {
		return "get " + v.visitRule(ctx.Block()).(string)
	}
}

func (v *FormatVisitor) VisitSetter(ctx *parser.SetterContext) interface{} {
	if ctx.SEMI() != nil {
		return "set;"
	} else {
		return "set " + v.visitRule(ctx.Block()).(string)
	}
}

func (v *FormatVisitor) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) interface{} {
	return v.visitRule(ctx.QualifiedName()).(string) + v.visitRule(ctx.FormalParameters()).(string) + " " + v.visitRule(ctx.Block()).(string)
}

func (v *FormatVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	statements := []string{}
	for _, stmt := range ctx.AllStatement() {
		statements = append(statements, v.visitRule(stmt).(string))
	}
	if len(statements) == 0 {
		return "{}"
	}
	return "{\n" + indent(strings.Join(statements, "\n")) + "\n}"
}

func (v *FormatVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	child := ctx.GetChild(0).(antlr.RuleNode)
	return v.visitRule(child)
}

func (v *FormatVisitor) VisitBlockMemberDeclaration(ctx *parser.BlockMemberDeclarationContext) interface{} {
	return v.Modifiers(ctx.AllModifier()) + v.visitRule(ctx.MemberDeclaration()).(string)
}

func (v *FormatVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	var out strings.Builder
	if block := ctx.Statement(0).Block(); block != nil {
		out.WriteString("if " + v.visitRule(ctx.ParExpression()).(string) + " " +
			v.visitRule(ctx.Statement(0)).(string))
	} else {
		out.WriteString("if " + v.visitRule(ctx.ParExpression()).(string) + " {\n" +
			indent(v.visitRule(ctx.Statement(0)).(string)) + "\n}")
	}
	if ctx.ELSE() != nil {
		if block := ctx.Statement(1).Block(); block != nil {
			out.WriteString(" else " + v.visitRule(ctx.Statement(1)).(string))
		} else if ifStatement := ctx.Statement(1).IfStatement(); ifStatement != nil {
			out.WriteString(" else " + v.visitRule(ifStatement).(string))
		} else {
			out.WriteString(" else {\n" + indent(v.visitRule(ctx.Statement(1)).(string)) + "\n}")
		}
	}
	return out.String()
}

func (v *FormatVisitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	if s := ctx.Statement(); s == nil {
		return "while " + v.visitRule(ctx.ParExpression()).(string) + ";"
	}
	if block := ctx.Statement().Block(); block != nil {
		return "while " + v.visitRule(ctx.ParExpression()).(string) + " " + v.visitRule(ctx.Statement()).(string)
	} else {
		return "while " + v.visitRule(ctx.ParExpression()).(string) + " {\n" + indent(v.visitRule(ctx.Statement()).(string)) + "\n}"
	}
}

func (v *FormatVisitor) VisitDoWhileStatement(ctx *parser.DoWhileStatementContext) interface{} {
	if block := ctx.Statement().Block(); block != nil {
		return "do " + v.visitRule(ctx.Statement()).(string) + " while " + v.visitRule(ctx.ParExpression()).(string) + ";"
	} else {
		return "do {\n" + indent(v.visitRule(ctx.Statement()).(string)) + "\n} while " + v.visitRule(ctx.ParExpression()).(string) + ";"
	}
}

func (v *FormatVisitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	if statement := ctx.Statement(); statement != nil {
		if statement.Block() != nil {
			return "for (" + v.visitRule(ctx.ForControl()).(string) + ") " + v.visitRule(ctx.Statement()).(string)
		} else {
			return "for (" + v.visitRule(ctx.ForControl()).(string) + ") {\n" + indent(v.visitRule(ctx.Statement()).(string)) + "\n}"
		}
	} else {
		return "for (" + v.visitRule(ctx.ForControl()).(string) + ");"
	}
}

func (v *FormatVisitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	when := []string{}
	for _, w := range ctx.AllWhenControl() {
		when = append(when, v.visitRule(w).(string))
	}
	return "switch on " + v.visitRule(ctx.Expression()).(string) + " {\n" + indent(strings.Join(when, "\n")) + "\n}"
}

func (v *FormatVisitor) VisitWhenControl(ctx *parser.WhenControlContext) interface{} {
	return "when " + v.visitRule(ctx.WhenValue()).(string) + " " + v.visitRule(ctx.Block()).(string)
}

func (v *FormatVisitor) VisitWhenValue(ctx *parser.WhenValueContext) interface{} {
	switch {
	case ctx.ELSE() != nil:
		return "else"
	case ctx.TypeRef() != nil:
		return v.visitRule(ctx.TypeRef()).(string) + " " + v.visitRule(ctx.Id(0)).(string)
	case len(ctx.AllId()) == 2:
		return v.visitRule(ctx.Id(0)).(string) + " " + v.visitRule(ctx.Id(1)).(string)
	default:
		whenLiterals := []string{}
		for _, w := range ctx.AllWhenLiteral() {
			whenLiterals = append(whenLiterals, v.visitRule(w).(string))
		}
		return strings.Join(whenLiterals, ", ")
	}
}

func (v *FormatVisitor) VisitWhenLiteral(ctx *parser.WhenLiteralContext) interface{} {
	if w := ctx.WhenLiteral(); w != nil {
		return "(" + v.visitRule(w).(string) + ")"
	}
	if i := ctx.Id(); i != nil {
		return v.visitRule(i)
	}
	return ctx.GetText()
}

func (v *FormatVisitor) VisitTryStatement(ctx *parser.TryStatementContext) interface{} {
	if len(ctx.AllCatchClause()) > 0 {
		catchClauses := []string{}
		for _, c := range ctx.AllCatchClause() {
			catchClauses = append(catchClauses, v.visitRule(c).(string))
		}
		finally := ""
		if f := ctx.FinallyBlock(); f != nil {
			finally = " " + v.visitRule(f).(string)
		}
		return "try " + v.visitRule(ctx.Block()).(string) + " " + strings.Join(catchClauses, " ") + finally
	} else {
		return "try " + v.visitRule(ctx.Block()).(string) + " " + v.visitRule(ctx.FinallyBlock()).(string)
	}
}

func (v *FormatVisitor) VisitCatchClause(ctx *parser.CatchClauseContext) interface{} {
	return "catch (" +
		v.Modifiers(ctx.AllModifier()) +
		v.visitRule(ctx.QualifiedName()).(string) + " " +
		v.visitRule(ctx.Id()).(string) + ") " +
		v.visitRule(ctx.Block()).(string)
}

func (v *FormatVisitor) VisitFinallyBlock(ctx *parser.FinallyBlockContext) interface{} {
	return "finally " + v.visitRule(ctx.Block()).(string)
}

func (v *FormatVisitor) VisitThrowStatement(ctx *parser.ThrowStatementContext) interface{} {
	return "throw " + v.visitRule(ctx.Expression()).(string) + ";"
}

func (v *FormatVisitor) VisitRunAsStatement(ctx *parser.RunAsStatementContext) interface{} {
	expressionList := ""
	if e := ctx.ExpressionList(); e != nil {
		expressionList = v.visitRule(e).(string)
	}
	return "System.runAs(" + expressionList + ") " + v.visitRule(ctx.Block()).(string)
}

func (v *FormatVisitor) VisitForControl(ctx *parser.ForControlContext) interface{} {
	if enhancedForControl := ctx.EnhancedForControl(); enhancedForControl != nil {
		return v.visitRule(enhancedForControl)
	}
	var init strings.Builder
	if forInit := ctx.ForInit(); forInit != nil {
		init.WriteString(v.visitRule(forInit).(string))
	}
	init.WriteString(";")
	if expression := ctx.Expression(); expression != nil {
		init.WriteString(" " + v.visitRule(expression).(string))
	}
	init.WriteString(";")
	if forUpdate := ctx.ForUpdate(); forUpdate != nil {
		init.WriteString(" " + v.visitRule(forUpdate).(string))
	}
	return init.String()
}

func (v *FormatVisitor) VisitEnhancedForControl(ctx *parser.EnhancedForControlContext) interface{} {
	var out strings.Builder
	out.WriteString(v.visitRule(ctx.TypeRef()).(string) + " " + v.visitRule(ctx.Id()).(string) + " : ")
	out.WriteString(v.visitRule(ctx.Expression()).(string))
	return out.String()
}

func (v *FormatVisitor) VisitForInit(ctx *parser.ForInitContext) interface{} {
	return v.visitRule(ctx.GetChild(0).(antlr.RuleNode))
}

func (v *FormatVisitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	return "continue;"
}

func (v *FormatVisitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	return "break;"
}

func (v *FormatVisitor) VisitForUpdate(ctx *parser.ForUpdateContext) interface{} {
	return v.visitRule(ctx.ExpressionList())
}

func (v *FormatVisitor) VisitLocalVariableDeclarationStatement(ctx *parser.LocalVariableDeclarationStatementContext) interface{} {
	return v.visitRule(ctx.LocalVariableDeclaration()).(string) + ";"
}

func (v *FormatVisitor) VisitInsertStatement(ctx *parser.InsertStatementContext) interface{} {
	accessMode := ""
	if ctx.AS() != nil {
		if ctx.SYSTEM() != nil {
			accessMode = " as system"
		} else if ctx.USER() != nil {
			accessMode = " as user"
		}
	}
	return "insert" + accessMode + " " + v.visitRule(ctx.Expression()).(string) + ";"
}

func (v *FormatVisitor) VisitUpdateStatement(ctx *parser.UpdateStatementContext) interface{} {
	accessMode := ""
	if ctx.AS() != nil {
		if ctx.SYSTEM() != nil {
			accessMode = " as system"
		} else if ctx.USER() != nil {
			accessMode = " as user"
		}
	}
	return "update" + accessMode + " " + v.visitRule(ctx.Expression()).(string) + ";"
}

func (v *FormatVisitor) VisitUpsertStatement(ctx *parser.UpsertStatementContext) interface{} {
	accessMode := ""
	if ctx.AS() != nil {
		if ctx.SYSTEM() != nil {
			accessMode = " as system"
		} else if ctx.USER() != nil {
			accessMode = " as user"
		}
	}
	if q := ctx.QualifiedName(); q != nil {
		return "upsert" + accessMode + " " + v.visitRule(ctx.Expression()).(string) + " " + v.visitRule(q).(string) + ";"
	} else {
		return "upsert" + accessMode + " " + v.visitRule(ctx.Expression()).(string) + ";"
	}
}

func (v *FormatVisitor) VisitMergeStatement(ctx *parser.MergeStatementContext) interface{} {
	return "merge " + v.visitRule(ctx.Expression(0)).(string) + " " + v.visitRule(ctx.Expression(1)).(string) + ";"
}

func (v *FormatVisitor) VisitDeleteStatement(ctx *parser.DeleteStatementContext) interface{} {
	accessMode := ""
	if ctx.AS() != nil {
		if ctx.SYSTEM() != nil {
			accessMode = " as system"
		} else if ctx.USER() != nil {
			accessMode = " as user"
		}
	}
	return "delete" + accessMode + " " + v.visitRule(ctx.Expression()).(string) + ";"
}

func (v *FormatVisitor) VisitUndeleteStatement(ctx *parser.UndeleteStatementContext) interface{} {
	accessMode := ""
	if ctx.AS() != nil {
		if ctx.SYSTEM() != nil {
			accessMode = " as system"
		} else if ctx.USER() != nil {
			accessMode = " as user"
		}
	}
	return "undelete" + accessMode + " " + v.visitRule(ctx.Expression()).(string) + ";"
}

func (v *FormatVisitor) VisitLocalVariableDeclaration(ctx *parser.LocalVariableDeclarationContext) interface{} {
	return v.Modifiers(ctx.AllModifier()) + v.visitRule(ctx.TypeRef()).(string) + " " + v.visitRule(ctx.VariableDeclarators()).(string)
}

func (v *FormatVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return "return " + v.visitRule(e).(string) + ";"
	}
	return "return;"
}

func (v *FormatVisitor) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return "(" + v.visitRule(ctx.Expression()).(string) + ")"
}

func (v *FormatVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	return v.visitRule(ctx.Expression()).(string) + ";"
}

func (v *FormatVisitor) VisitAssignExpression(ctx *parser.AssignExpressionContext) interface{} {
	assignmentToken := ctx.GetChild(1).(antlr.TerminalNode)
	return v.visitRule(ctx.Expression(0)).(string) + " " + assignmentToken.GetText() + " " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitCondExpression(ctx *parser.CondExpressionContext) interface{} {
	if v.textLen(ctx.Expression(0))+v.textLen(ctx.Expression(1))+v.textLen(ctx.Expression(2)) <= 60 {
		return v.visitRule(ctx.Expression(0)).(string) + " ? " + v.visitRule(ctx.Expression(1)).(string) + " : " + v.visitRule(ctx.Expression(2)).(string)
	}
	return v.visitRule(ctx.Expression(0)).(string) + " ?\n" +
		indent(v.visitRule(ctx.Expression(1)).(string)) + " :\n" +
		indent(v.visitRule(ctx.Expression(2)).(string))
}

func (v *FormatVisitor) VisitLogAndExpression(ctx *parser.LogAndExpressionContext) interface{} {
	i := NewChainVisitor()
	if v.textLen(ctx) < 40 {
		defer restoreWrap(unwrap(v))
	} else if i.visitRule(ctx.Expression(0)).(int)+i.visitRule(ctx.Expression(1)).(int) > 2 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return v.visitRule(ctx.Expression(0)).(string) + " &&\n" + indent(v.visitRule(ctx.Expression(1)).(string))
	}
	return v.visitRule(ctx.Expression(0)).(string) + " && " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitLogOrExpression(ctx *parser.LogOrExpressionContext) interface{} {
	i := NewChainVisitor()
	if v.textLen(ctx) < 40 {
		defer restoreWrap(unwrap(v))
	} else if i.visitRule(ctx.Expression(0)).(int)+i.visitRule(ctx.Expression(1)).(int) > 2 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return v.visitRule(ctx.Expression(0)).(string) + " ||\n" + indent(v.visitRule(ctx.Expression(1)).(string))
	}
	return v.visitRule(ctx.Expression(0)).(string) + " || " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitCoalExpression(ctx *parser.CoalExpressionContext) interface{} {
	i := NewChainVisitor()
	if i.visitRule(ctx.Expression(0)).(int)+i.visitRule(ctx.Expression(1)).(int) > 2 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return v.visitRule(ctx.Expression(0)).(string) + " ??\n" + indent(v.visitRule(ctx.Expression(1)).(string))
	}
	return v.visitRule(ctx.Expression(0)).(string) + " ?? " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	return v.visitRule(ctx.Expression(0)).(string) + " & " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	return v.visitRule(ctx.Expression(0)).(string) + " | " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitBitNotExpression(ctx *parser.BitNotExpressionContext) interface{} {
	return v.visitRule(ctx.Expression(0)).(string) + " ^ " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitBitExpression(ctx *parser.BitExpressionContext) interface{} {
	operators := []string{}
	for i := 1; i < ctx.GetChildCount()-1; i++ {
		if token, ok := ctx.GetChild(i).(antlr.TerminalNode); ok {
			operators = append(operators, token.GetText())
		}
	}
	left := v.visitRule(ctx.Expression(0)).(string)
	right := v.visitRule(ctx.Expression(1)).(string)
	return left + " " + strings.Join(operators, "") + " " + right
}

func (v *FormatVisitor) VisitArth1Expression(ctx *parser.Arth1ExpressionContext) interface{} {
	return v.visitRule(ctx.Expression(0)).(string) + " " + ctx.GetChild(1).(antlr.TerminalNode).GetText() + " " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitArth2Expression(ctx *parser.Arth2ExpressionContext) interface{} {
	sep := " "
	i := NewChainVisitor()
	left := i.visitRule(ctx.Expression(0)).(int)
	right := i.visitRule(ctx.Expression(1)).(int)
	textLen := v.textLen(ctx)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("LEFT %d: %s ", left, ctx.Expression(0).GetText())
		log.Debugf("RIGHT %d: %s ", right, ctx.Expression(1).GetText())
		log.Debugf("TEXT %d: %s ", textLen, ctx.GetText())
	}
	wrap := v.wrap || (left+right > 2 && textLen > 40) || textLen > 60
	if wrap {
		sep = "\n\t"
		defer restoreWrap(unwrap(v))
	}
	return v.visitRule(ctx.Expression(0)).(string) + " " + ctx.GetChild(1).(antlr.TerminalNode).GetText() + sep + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitNegExpression(ctx *parser.NegExpressionContext) interface{} {
	return ctx.GetChild(0).(antlr.TerminalNode).GetText() + v.visitRule(ctx.Expression()).(string)
}

func (v *FormatVisitor) VisitPreOpExpression(ctx *parser.PreOpExpressionContext) interface{} {
	return ctx.GetChild(0).(antlr.TerminalNode).GetText() + v.visitRule(ctx.Expression()).(string)
}

func (v *FormatVisitor) VisitPostOpExpression(ctx *parser.PostOpExpressionContext) interface{} {
	return v.visitRule(ctx.Expression()).(string) + ctx.GetChild(1).(antlr.TerminalNode).GetText()
}

func (v *FormatVisitor) VisitSubExpression(ctx *parser.SubExpressionContext) interface{} {
	return "(" + v.visitRule(ctx.Expression()).(string) + ")"
}

func (v *FormatVisitor) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	return "(" + v.visitRule(ctx.TypeRef()).(string) + ")" + v.visitRule(ctx.Expression()).(string)
}

func (v *FormatVisitor) VisitNewInstanceExpression(ctx *parser.NewInstanceExpressionContext) interface{} {
	return "new " + v.visitRule(ctx.Creator()).(string)
}

func (v *FormatVisitor) VisitArrayExpression(ctx *parser.ArrayExpressionContext) interface{} {
	if v.textLen(ctx) < 20 {
		defer restoreWrap(unwrap(v))
	}
	return v.visitRule(ctx.Expression(0)).(string) + "[" + v.visitRule(ctx.Expression(1)).(string) + "]"
}

func (v *FormatVisitor) VisitDotExpression(ctx *parser.DotExpressionContext) interface{} {
	i := NewChainVisitor()
	depth := i.visitRule(ctx.Expression()).(int)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("depth is %d: %s", depth, ctx.GetText())
	}
	if depth > 1 {
		defer restoreWrap(wrap(v))
	}
	expr := v.visitRule(ctx.Expression())
	dot := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	switch {
	case ctx.DotMethodCall() != nil:
		i := NewIndentVisitor()
		depth := i.visitRule(ctx.Expression()).(int)
		if v.wrap {
			if depth == 0 {
				depth = 1
			}
			switch left := ctx.Expression().(type) {
			case *parser.PrimaryExpressionContext:
				log.Debugf("NOT wrapping after between %q (%T)", expr, ctx.Expression())
			case *parser.DotExpressionContext:
				if left.DotMethodCall() != nil {
					if log.IsLevelEnabled(log.DebugLevel) {
						log.Debugf("%q is method call; safe to wrap before %q", expr, ctx.DotMethodCall().GetText())
					}
					return expr.(string) + "\n" + indentTo(dot+v.visitRule(ctx.DotMethodCall()).(string), depth)
				}
			default:
				if log.IsLevelEnabled(log.DebugLevel) {
					log.Debugf("Wrapping in between %q (%T) and %q", expr, ctx.Expression(), ctx.DotMethodCall().GetText())
				}
				return expr.(string) + "\n" + indentTo(dot+v.visitRule(ctx.DotMethodCall()).(string), depth)
			}
		}

		return expr.(string) + dot + v.visitRule(ctx.DotMethodCall()).(string)
	case ctx.AnyId() != nil:
		return expr.(string) + dot + v.visitRule(ctx.AnyId()).(string)
	}
	return ""
}

func (v *FormatVisitor) VisitDotMethodCall(ctx *parser.DotMethodCallContext) interface{} {
	if v.wrap {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("Visitor says to wrap in VisitDotMethodCall; not wrapping individual expressions: %s", ctx.GetText())
		}
		defer restoreWrap(unwrap(v))
	}
	expressionList := ""
	if l := ctx.ExpressionList(); l != nil {
		expressionList = v.visitRule(l).(string)
	}
	return v.visitRule(ctx.AnyId()).(string) + "(" + expressionList + ")"
}

func (v *FormatVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	textLen := v.textLen(ctx)
	wrap := v.wrap || (textLen > 40 && len(ctx.AllExpression()) > 3) || textLen > 150

	expressions := []string{}
	for i, p := range ctx.AllExpression() {
		// We want to indent method argument expressions, but not new instance arguments
		switch p.(type) {
		case *parser.AssignExpressionContext:
			defer restoreWrap(unwrap(v))
			expressions = append(expressions, v.visitRule(p).(string))
		default:
			if wrap && i > 0 && !v.wrap {
				expressions = append(expressions, indent(v.visitRule(p).(string)))
			} else {
				expressions = append(expressions, v.visitRule(p).(string))
			}
		}
	}

	if wrap {
		return strings.Join(expressions, ",\n")
	}
	return strings.Join(expressions, ", ")
}

func (v *FormatVisitor) VisitAnyId(ctx *parser.AnyIdContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	switch e := ctx.Primary().(type) {
	case *parser.ThisPrimaryContext:
		return "this"
	case *parser.SuperPrimaryContext:
		return "super"
	case *parser.LiteralPrimaryContext:
		return v.visitRule(e)
	case *parser.TypeRefPrimaryContext:
		return v.visitRule(e)
	case *parser.IdPrimaryContext:
		return v.visitRule(e)
	case *parser.SoqlPrimaryContext:
		return v.visitRule(e)
	case *parser.SoslPrimaryContext:
		return v.visitRule(e)
	default:
		return fmt.Sprintf("UNHANDLED PRIMARY EXPRESSION: %T %s", e, e.GetText())
	}
}

func (v *FormatVisitor) VisitIdPrimary(ctx *parser.IdPrimaryContext) interface{} {
	return v.visitRule(ctx.Id())
}

func (v *FormatVisitor) VisitLiteralPrimary(ctx *parser.LiteralPrimaryContext) interface{} {
	return v.visitRule(ctx.Literal())
}

func (v *FormatVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitMethodCallExpression(ctx *parser.MethodCallExpressionContext) interface{} {
	return v.visitRule(ctx.MethodCall())
}

func (v *FormatVisitor) VisitMethodCall(ctx *parser.MethodCallContext) interface{} {
	var f string
	switch e := ctx.GetChild(0).(type) {
	case *parser.IdContext:
		f = v.visitRule(e).(string)
	case antlr.TerminalNode:
		f = strings.ToLower(e.GetText())
	}
	expressionList := ""
	if el := ctx.ExpressionList(); el != nil {
		expressionList = v.visitRule(el).(string)
	}
	return f + "(" + expressionList + ")"
}

func (v *FormatVisitor) VisitSoslPrimary(ctx *parser.SoslPrimaryContext) interface{} {
	return v.visitRule(ctx.SoslLiteral())
}

func (v *FormatVisitor) VisitSoqlPrimary(ctx *parser.SoqlPrimaryContext) interface{} {
	return v.visitRule(ctx.SoqlLiteral())
}

func (v *FormatVisitor) VisitSoqlLiteral(ctx *parser.SoqlLiteralContext) interface{} {
	// Check whether we should wrap this SOQL Query based on query complexity
	i := NewChainVisitor()
	n := i.visitRule(ctx.Query()).(int)
	if n > 3 {
		defer restoreWrap(wrap(v))
		return "[\n" + indent(v.visitRule(ctx.Query()).(string)) + "\n]"
	}
	return "[" + v.visitRule(ctx.Query()).(string) + "]"
}

func (v *FormatVisitor) VisitQuery(ctx *parser.QueryContext) interface{} {
	i := NewChainVisitor()
	n := i.visitRule(ctx).(int)
	if n > 3 {
		defer restoreWrap(wrap(v))
	}
	sep := " "
	indent := 0
	if v.wrap {
		sep = "\n"
		indent = 1
	}
	var query strings.Builder
	query.WriteString("SELECT")
	query.WriteString(sep)
	query.WriteString(indentTo(v.visitRule(ctx.SelectList()).(string), indent))
	query.WriteString(sep)
	query.WriteString("FROM")
	query.WriteString(sep)
	query.WriteString(indentTo(v.visitRule(ctx.FromNameList()).(string), indent))
	if scope := ctx.UsingScope(); scope != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(scope).(string))
	}
	if where := ctx.WhereClause(); where != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(where).(string))
	}
	if withClause := ctx.WithClause(); withClause != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(withClause).(string))
	}
	if groupBy := ctx.GroupByClause(); groupBy != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(groupBy).(string))
	}
	if orderBy := ctx.OrderByClause(); orderBy != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(orderBy).(string))
	}
	if limit := ctx.LimitClause(); limit != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(limit).(string))
	}
	if offset := ctx.OffsetClause(); offset != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(offset).(string))
	}
	if ctx.AllRowsClause() != nil {
		query.WriteString(sep)
		query.WriteString("ALL ROWS")
	}
	forClauses := v.visitRule(ctx.ForClauses())
	if forClauses != "" {
		query.WriteString(sep)
		query.WriteString(forClauses.(string))
	}
	if update := ctx.UpdateList(); update != nil {
		query.WriteString(sep)
		query.WriteString("UPDATE " + v.visitRule(update).(string))
	}
	if setOptions := ctx.SetOptionsClause(); setOptions != nil {
		query.WriteString(sep)
		query.WriteString(v.visitRule(setOptions).(string))
	}
	return query.String()
}

// VisitSetOptionsClause formats a SET OPTIONS clause, which passes a
// Database.QueryOptions bind variable to the query.
func (v *FormatVisitor) VisitSetOptionsClause(ctx *parser.SetOptionsClauseContext) interface{} {
	return "SET OPTIONS " + v.visitRule(ctx.BoundExpression()).(string)
}

func (v *FormatVisitor) VisitSubQuery(ctx *parser.SubQueryContext) interface{} {
	var query strings.Builder
	query.WriteString("SELECT\n" +
		indent(v.visitRule(ctx.SubFieldList()).(string)) + "\nFROM\n" +
		indent(v.visitRule(ctx.FromNameList()).(string)),
	)
	if where := ctx.WhereClause(); where != nil {
		query.WriteString("\n" + v.visitRule(where).(string))
	}
	if orderBy := ctx.OrderByClause(); orderBy != nil {
		query.WriteString("\n" + v.visitRule(orderBy).(string))
	}
	if limit := ctx.LimitClause(); limit != nil {
		query.WriteString("\n" + v.visitRule(limit).(string))
	}
	forClauses := v.visitRule(ctx.ForClauses())
	if forClauses != "" {
		query.WriteString("\n" + forClauses.(string))
	}
	if update := ctx.UpdateList(); update != nil {
		query.WriteString("\nUPDATE " + v.visitRule(update).(string))
	}
	return query.String()
}

func (v *FormatVisitor) VisitFromNameList(ctx *parser.FromNameListContext) interface{} {
	fieldNames := []string{}
	for _, p := range ctx.AllFieldNameAlias() {
		fieldNames = append(fieldNames, v.visitRule(p).(string))
	}
	return strings.Join(fieldNames, ",\n")
}

func (v *FormatVisitor) VisitUpdateList(ctx *parser.UpdateListContext) interface{} {
	updateList := ""
	if u := ctx.UpdateList(); u != nil {
		updateList = ", " + v.visitRule(u).(string)
	}
	return ctx.UpdateType().GetText() + updateList
}

func (v *FormatVisitor) VisitFieldNameAlias(ctx *parser.FieldNameAliasContext) interface{} {
	alias := ""
	if s := ctx.SoqlSelectAlias(); s != nil {
		alias = " " + s.GetText()
	}
	return v.visitRule(ctx.FieldName()).(string) + alias
}

func (v *FormatVisitor) VisitSelectList(ctx *parser.SelectListContext) interface{} {
	sep := ", "
	if v.wrap {
		sep = ",\n"
	}
	selectEntries := []string{}
	for _, p := range ctx.AllSelectEntry() {
		selectEntries = append(selectEntries, v.visitRule(p).(string))
	}
	return strings.Join(selectEntries, sep)
}

func (v *FormatVisitor) VisitSubFieldList(ctx *parser.SubFieldListContext) interface{} {
	selectEntries := []string{}
	for _, p := range ctx.AllSubFieldEntry() {
		selectEntries = append(selectEntries, v.visitRule(p).(string))
	}
	return strings.Join(selectEntries, ",\n")
}

func (v *FormatVisitor) VisitSelectEntry(ctx *parser.SelectEntryContext) interface{} {
	alias := ""
	if s := ctx.SoqlSelectAlias(); s != nil {
		alias = " " + s.GetText()
	}
	switch {
	case ctx.FieldName() != nil:
		return v.visitRule(ctx.FieldName()).(string) + alias
	case ctx.SoqlFunction() != nil:
		return v.visitRule(ctx.SoqlFunction()).(string) + alias
	case ctx.SubQuery() != nil:
		return "(" + v.visitRule(ctx.SubQuery()).(string) + ")" + alias
	case ctx.TypeOf() != nil:
		return v.visitRule(ctx.TypeOf()).(string)
	}
	panic("Unexpected selectEntry")
}

func (v *FormatVisitor) VisitSubFieldEntry(ctx *parser.SubFieldEntryContext) interface{} {
	alias := ""
	if s := ctx.SoqlSelectAlias(); s != nil {
		alias = " " + s.GetText()
	}
	switch {
	case ctx.FieldName() != nil:
		return v.visitRule(ctx.FieldName()).(string) + alias
	case ctx.SoqlFunction() != nil:
		return v.visitRule(ctx.SoqlFunction()).(string) + alias
	case ctx.SubQuery() != nil:
		return "(" + v.visitRule(ctx.SubQuery()).(string) + ")" + alias
	case ctx.TypeOf() != nil:
		return v.visitRule(ctx.TypeOf()).(string)
	}
	panic("Unexpected selectEntry")
}

func (v *FormatVisitor) VisitFieldName(ctx *parser.FieldNameContext) interface{} {
	ids := []string{}
	for _, t := range ctx.AllSoqlId() {
		ids = append(ids, t.GetText())
	}
	return strings.Join(ids, ".")
}

func (v *FormatVisitor) VisitFieldNameList(ctx *parser.FieldNameListContext) interface{} {
	fieldNames := []string{}
	for _, p := range ctx.AllFieldName() {
		fieldNames = append(fieldNames, v.visitRule(p).(string))
	}
	return strings.Join(fieldNames, ",\n")
}

func (v *FormatVisitor) VisitTypeOf(ctx *parser.TypeOfContext) interface{} {
	whenClauses := []string{}
	for _, w := range ctx.AllWhenClause() {
		whenClauses = append(whenClauses, v.visitRule(w).(string))
	}
	elseClause := ""
	if e := ctx.ElseClause(); e != nil {
		elseClause = v.visitRule(e).(string) + "\n"
	}

	return "TYPEOF " +
		v.visitRule(ctx.FieldName()).(string) + "\n" +
		strings.Join(whenClauses, "\n") + "\n" +
		elseClause + "END"

}

func (v *FormatVisitor) VisitForClauses(ctx *parser.ForClausesContext) interface{} {
	forClauses := []string{}
	for _, f := range ctx.AllForClause() {
		forClauses = append(forClauses, v.visitRule(f).(string))
	}
	return strings.Join(forClauses, " ")
}

func (v *FormatVisitor) VisitForClause(ctx *parser.ForClauseContext) interface{} {
	return "FOR " + ctx.GetChild(1).(antlr.TerminalNode).GetText()
}

func (v *FormatVisitor) VisitWhenClause(ctx *parser.WhenClauseContext) interface{} {
	sep := " "
	indent := 0
	if v.wrap {
		sep = "\n"
		indent = 1
	}
	var clause strings.Builder
	clause.WriteString("WHEN")
	clause.WriteString(sep)
	clause.WriteString(indentTo(v.visitRule(ctx.FieldName()).(string), indent))
	clause.WriteString(sep)
	clause.WriteString("THEN")
	clause.WriteString(sep)
	clause.WriteString(indentTo(v.visitRule(ctx.FieldNameList()).(string), indent))
	return clause.String()
}

func (v *FormatVisitor) VisitElseClause(ctx *parser.ElseClauseContext) interface{} {
	sep := " "
	indent := 0
	if v.wrap {
		sep = "\n"
		indent = 1
	}
	var clause strings.Builder
	clause.WriteString("ELSE")
	clause.WriteString(sep)
	clause.WriteString(indentTo(v.visitRule(ctx.FieldNameList()).(string), indent))
	return clause.String()
}

func (v *FormatVisitor) VisitWithClause(ctx *parser.WithClauseContext) interface{} {
	if logical := ctx.LogicalExpression(); logical != nil {
		sep := " "
		indent := 0
		if v.wrap {
			sep = "\n"
			indent = 1
		}
		return "WITH" + sep + indentTo(v.visitRule(logical).(string), indent)
	}
	if ctx.FilteringExpression() != nil {
		return "WITH DATA CATEGORY " + v.visitRule(ctx.FilteringExpression()).(string)
	}
	if ctx.SECURITY_ENFORCED() != nil {
		return "WITH SECURITY_ENFORCED"
	}
	if ctx.USER_MODE() != nil {
		return "WITH USER_MODE"
	}
	if ctx.SYSTEM_MODE() != nil {
		return "WITH SYSTEM_MODE"
	}
	if ctx.RECORDVISIBILITYCONTEXT() != nil {
		params := make([]string, 0, len(ctx.AllRecordVisibilityContextParam()))
		for _, param := range ctx.AllRecordVisibilityContextParam() {
			params = append(params, v.visitRule(param).(string))
		}
		return "WITH RecordVisibilityContext (" + strings.Join(params, ", ") + ")"
	}
	return "WITH"
}

func (v *FormatVisitor) VisitFilteringExpression(ctx *parser.FilteringExpressionContext) interface{} {
	selections := make([]string, 0, len(ctx.AllDataCategorySelection()))
	for _, selection := range ctx.AllDataCategorySelection() {
		selections = append(selections, v.visitRule(selection).(string))
	}
	return strings.Join(selections, " AND ")
}

func (v *FormatVisitor) VisitDataCategorySelection(ctx *parser.DataCategorySelectionContext) interface{} {
	return ctx.SoqlId().GetText() + " " +
		v.visitRule(ctx.FilteringSelector()).(string) + " " +
		v.visitRule(ctx.DataCategoryName()).(string)
}

func (v *FormatVisitor) VisitDataCategoryName(ctx *parser.DataCategoryNameContext) interface{} {
	ids := make([]string, 0, len(ctx.AllSoqlId()))
	for _, id := range ctx.AllSoqlId() {
		ids = append(ids, id.GetText())
	}
	if ctx.LPAREN() != nil {
		return "(" + strings.Join(ids, ", ") + ")"
	}
	return strings.Join(ids, ", ")
}

func (v *FormatVisitor) VisitFilteringSelector(ctx *parser.FilteringSelectorContext) interface{} {
	switch {
	case ctx.ABOVE_OR_BELOW() != nil:
		return "ABOVE_OR_BELOW"
	case ctx.ABOVE() != nil:
		return "ABOVE"
	case ctx.BELOW() != nil:
		return "BELOW"
	default:
		return "AT"
	}
}

func (v *FormatVisitor) VisitRecordVisibilityContextParam(ctx *parser.RecordVisibilityContextParamContext) interface{} {
	if intValue := ctx.IntegerLiteral(); intValue != nil {
		return "maxDescriptorPerRecord=" + intValue.GetText()
	}
	if boolValue := ctx.BooleanLiteral(); boolValue != nil {
		switch {
		case ctx.SUPPORTSDOMAINS() != nil:
			return "supportsDomains=" + boolValue.GetText()
		case ctx.SUPPORTSDELEGATES() != nil:
			return "supportsDelegates=" + boolValue.GetText()
		}
	}
	return ctx.GetText()
}

func (v *FormatVisitor) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	sep := " "
	indent := 0
	if v.wrap {
		sep = "\n"
		indent = 1
	}
	var clause strings.Builder
	clause.WriteString("WHERE")
	clause.WriteString(sep)
	clause.WriteString(indentTo(v.visitRule(ctx.LogicalExpression()).(string), indent))
	return clause.String()
}

func (v *FormatVisitor) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	if e := ctx.BoundExpression(); e != nil {
		return "LIMIT " + v.visitRule(ctx.BoundExpression()).(string)
	}
	return "LIMIT " + ctx.IntegerLiteral().GetText()
}

func (v *FormatVisitor) VisitOffsetClause(ctx *parser.OffsetClauseContext) interface{} {
	if e := ctx.BoundExpression(); e != nil {
		return "OFFSET " + v.visitRule(ctx.BoundExpression()).(string)
	}
	return "OFFSET " + ctx.IntegerLiteral().GetText()
}

func (v *FormatVisitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) interface{} {
	switch {
	case ctx.NOT() != nil:
		return "NOT " + v.visitRule(ctx.ConditionalExpression(0)).(string)
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

func (v *FormatVisitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) interface{} {
	switch {
	case ctx.LogicalExpression() != nil:
		return "(\n" + indent(v.visitRule(ctx.LogicalExpression()).(string)) + "\n)"
	case ctx.FieldExpression() != nil:
		return v.visitRule(ctx.FieldExpression())
	}
	panic("Unexpected conditionalExpression")
}

func (v *FormatVisitor) VisitFieldExpression(ctx *parser.FieldExpressionContext) interface{} {
	switch {
	case ctx.FieldName() != nil:
		// TODO: Format IN/NOT IN
		return v.visitRule(ctx.FieldName()).(string) + " " + v.visitRule(ctx.ComparisonOperator()).(string) + " " + v.visitRule(ctx.Value()).(string)
	case ctx.SoqlFunction() != nil:
		return v.visitRule(ctx.SoqlFunction()).(string) + " " + v.visitRule(ctx.ComparisonOperator()).(string) + " " + v.visitRule(ctx.Value()).(string)
	}
	panic("Unexpected fieldExpression")
}

func (v *FormatVisitor) VisitComparisonOperator(ctx *parser.ComparisonOperatorContext) interface{} {
	if ctx.NOT() != nil {
		return "NOT IN"
	}
	return ctx.GetText()
}

func (v *FormatVisitor) VisitSoqlFunction(ctx *parser.SoqlFunctionContext) interface{} {
	param := ""
	switch {
	case ctx.FieldName() != nil:
		param = v.visitRule(ctx.FieldName()).(string)
	case ctx.COUNT() != nil:
		return "COUNT()"
	case ctx.SoqlFunction() != nil:
		param = v.visitRule(ctx.SoqlFunction()).(string)
	case ctx.DateFieldName() != nil:
		param = v.visitRule(ctx.DateFieldName()).(string)
	case ctx.SoqlFieldsParameter() != nil:
		param = v.visitRule(ctx.SoqlFieldsParameter()).(string)
	case ctx.DISTANCE() != nil:
		locationValues := ctx.AllLocationValue()
		loc1 := v.visitRule(locationValues[0]).(string)
		loc2 := v.visitRule(locationValues[1]).(string)
		unit := ctx.StringLiteral().GetText()
		return "DISTANCE(" + loc1 + ", " + loc2 + ", " + unit + ")"
	case ctx.FORMULA() != nil:
		return "FORMULA(" + ctx.StringLiteral().GetText() + ")"
	default:
		panic("Unexpected parameter type for soqlFunction")
	}
	return strings.ToUpper(ctx.GetChild(0).(antlr.TerminalNode).GetText()) + "(" + param + ")"
}

func (v *FormatVisitor) VisitSoqlFieldsParameter(ctx *parser.SoqlFieldsParameterContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitDateFieldName(ctx *parser.DateFieldNameContext) interface{} {
	if ctx.CONVERT_TIMEZONE() != nil {
		return "convertTimezone(" + v.visitRule(ctx.FieldName()).(string) + ")"
	}
	return v.visitRule(ctx.FieldName())
}

func (v *FormatVisitor) VisitLocationValue(ctx *parser.LocationValueContext) interface{} {
	switch {
	case ctx.FieldName() != nil:
		return v.visitRule(ctx.FieldName()).(string)
	case ctx.BoundExpression() != nil:
		return v.visitRule(ctx.BoundExpression()).(string)
	case ctx.GEOLOCATION() != nil:
		coords := []string{}
		for _, coord := range ctx.AllCoordinateValue() {
			coords = append(coords, v.visitRule(coord).(string))
		}
		return "GEOLOCATION(" + coords[0] + ", " + coords[1] + ")"
	default:
		panic("Unexpected location value type")
	}
}

func (v *FormatVisitor) VisitCoordinateValue(ctx *parser.CoordinateValueContext) interface{} {
	switch {
	case ctx.SignedNumber() != nil:
		return v.visitRule(ctx.SignedNumber()).(string)
	case ctx.BoundExpression() != nil:
		return v.visitRule(ctx.BoundExpression()).(string)
	default:
		panic("Unexpected coordinate value type")
	}
}

func (v *FormatVisitor) VisitNullValue(ctx *parser.NullValueContext) interface{} {
	return "null"
}

func (v *FormatVisitor) VisitBooleanLiteralValue(ctx *parser.BooleanLiteralValueContext) interface{} {
	return strings.ToLower(ctx.GetText())
}

func (v *FormatVisitor) VisitSignedNumberValue(ctx *parser.SignedNumberValueContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitStringLiteralValue(ctx *parser.StringLiteralValueContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitDateLiteralValue(ctx *parser.DateLiteralValueContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitDateTimeLiteralValue(ctx *parser.DateTimeLiteralValueContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitTimeLiteralValue(ctx *parser.TimeLiteralValueContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitDateFormulaValue(ctx *parser.DateFormulaValueContext) interface{} {
	return v.visitRule(ctx.DateFormula())
}

func (v *FormatVisitor) VisitCurrencyValueValue(ctx *parser.CurrencyValueValueContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitSubQueryValue(ctx *parser.SubQueryValueContext) interface{} {
	return "(\n" + indent(v.visitRule(ctx.SubQuery()).(string)) + "\n)"
}

func (v *FormatVisitor) VisitValueListValue(ctx *parser.ValueListValueContext) interface{} {
	return v.visitRule(ctx.ValueList())
}

func (v *FormatVisitor) VisitBoundExpressionValue(ctx *parser.BoundExpressionValueContext) interface{} {
	return v.visitRule(ctx.BoundExpression())
}

func (v *FormatVisitor) VisitDateFormula(ctx *parser.DateFormulaContext) interface{} {
	if ctx.SignedInteger() != nil {
		return ctx.GetChild(0).(antlr.TerminalNode).GetText() + ":" + v.visitRule(ctx.SignedInteger()).(string)
	}
	return ctx.GetChild(0).(antlr.TerminalNode).GetText()
}

func (v *FormatVisitor) VisitSignedInteger(ctx *parser.SignedIntegerContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitSignedNumber(ctx *parser.SignedNumberContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitValueList(ctx *parser.ValueListContext) interface{} {
	values := []string{}
	for _, i := range ctx.AllValue() {
		values = append(values, v.visitRule(i).(string))
	}
	return "(" + strings.Join(values, ", ") + ")"
}

func (v *FormatVisitor) VisitGroupByClause(ctx *parser.GroupByClauseContext) interface{} {
	fieldNames := []string{}
	for _, i := range ctx.AllGroupByField() {
		fieldNames = append(fieldNames, v.visitRule(i).(string))
	}
	sep := " "
	indent := 0
	if v.wrap {
		sep = "\n"
		indent = 1
	}
	switch {
	case ctx.ROLLUP() != nil:
		return v.formatGroupedClause("ROLLUP", fieldNames, ctx.LogicalExpression(), sep, indent)
	case ctx.CUBE() != nil:
		return v.formatGroupedClause("CUBE", fieldNames, ctx.LogicalExpression(), sep, indent)
	default:
		var clause strings.Builder
		clause.WriteString("GROUP BY")
		clause.WriteString(sep)
		clause.WriteString(indentTo(v.visitRule(ctx.SelectList()).(string), indent))
		if l := ctx.LogicalExpression(); l != nil {
			clause.WriteString(sep)
			clause.WriteString("HAVING")
			clause.WriteString(sep)
			clause.WriteString(indentTo(v.visitRule(l).(string), indent))
		}
		return clause.String()
	}
}

func (v *FormatVisitor) VisitUsingScope(ctx *parser.UsingScopeContext) interface{} {
	return "USING SCOPE " + ctx.SoqlId().GetText()
}

func (v *FormatVisitor) formatGroupedClause(keyword string, fields []string, having parser.ILogicalExpressionContext, sep string, indent int) string {
	var clause strings.Builder
	if v.wrap {
		clause.WriteString("GROUP BY")
		clause.WriteString(sep)
		clause.WriteString(indentTo(keyword+"("+strings.Join(fields, ", ")+")", indent))
	} else {
		clause.WriteString("GROUP BY " + keyword + "(" + strings.Join(fields, ", ") + ")")
	}
	if having != nil {
		clause.WriteString(sep)
		clause.WriteString("HAVING")
		clause.WriteString(sep)
		clause.WriteString(indentTo(v.visitRule(having).(string), indent))
	}
	return clause.String()
}

func (v *FormatVisitor) VisitGroupByField(ctx *parser.GroupByFieldContext) interface{} {
	switch {
	case ctx.SoqlFunction() != nil:
		return v.visitRule(ctx.SoqlFunction()).(string)
	case ctx.FieldName() != nil:
		return v.visitRule(ctx.FieldName()).(string)
	default:
		return ""
	}
}

func (v *FormatVisitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) interface{} {
	sep := " "
	indent := 0
	if v.wrap {
		sep = "\n"
		indent = 1
	}
	var clause strings.Builder
	clause.WriteString("ORDER BY")
	clause.WriteString(sep)
	clause.WriteString(indentTo(v.visitRule(ctx.FieldOrderList()).(string), indent))
	return clause.String()
}

func (v *FormatVisitor) VisitFieldOrderList(ctx *parser.FieldOrderListContext) interface{} {
	fields := []string{}
	for _, i := range ctx.AllFieldOrder() {
		fields = append(fields, v.visitRule(i).(string))
	}
	sep := ", "
	if v.wrap {
		sep = ",\n"
	}
	return strings.Join(fields, sep)
}

func (v *FormatVisitor) VisitFieldOrder(ctx *parser.FieldOrderContext) interface{} {
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

func (v *FormatVisitor) VisitBoundExpression(ctx *parser.BoundExpressionContext) interface{} {
	return ":" + v.visitRule(ctx.Expression()).(string)
}

func (v *FormatVisitor) VisitCreator(ctx *parser.CreatorContext) interface{} {
	return v.visitRule(ctx.CreatedName()).(string) + v.visitRule(ctx.GetChild(1).(antlr.RuleNode)).(string)
}

func (v *FormatVisitor) VisitCreatedName(ctx *parser.CreatedNameContext) interface{} {
	namePairs := []string{}
	for _, i := range ctx.AllIdCreatedNamePair() {
		namePairs = append(namePairs, v.visitRule(i).(string))
	}
	return strings.Join(namePairs, ".")
}

func (v *FormatVisitor) VisitIdCreatedNamePair(ctx *parser.IdCreatedNamePairContext) interface{} {
	if typeList := ctx.TypeList(); typeList != nil {
		return v.visitRule(ctx.AnyId()).(string) + "<" + v.visitRule(typeList).(string) + ">"
	}
	return v.visitRule(ctx.AnyId())
}

func (v *FormatVisitor) VisitNoRest(ctx *parser.NoRestContext) interface{} {
	return "{}"
}

func (v *FormatVisitor) VisitId(ctx *parser.IdContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitClassCreatorRest(ctx *parser.ClassCreatorRestContext) interface{} {
	return v.visitRule(ctx.Arguments())
}

func (v *FormatVisitor) VisitArrayCreatorRest(ctx *parser.ArrayCreatorRestContext) interface{} {
	if expression := ctx.Expression(); expression != nil {
		return "[" + v.visitRule(expression).(string) + "]"
	}

	dimensions := len(ctx.AllLBRACK())
	if dimensions == 0 {
		dimensions = 1
	}

	var out strings.Builder
	for i := 0; i < dimensions; i++ {
		out.WriteString("[]")
	}

	if arrayInitializer := ctx.ArrayInitializer(); arrayInitializer != nil {
		out.WriteString(v.visitRule(arrayInitializer).(string))
	}

	return out.String()
}

func (v *FormatVisitor) VisitMapCreatorRest(ctx *parser.MapCreatorRestContext) interface{} {
	pairs := []string{}
	for _, i := range ctx.AllMapCreatorRestPair() {
		pairs = append(pairs, v.visitRule(i).(string))
	}
	if len(pairs) > 1 {
		return "{\n" + indent(strings.Join(pairs, ",\n")) + "\n}"
	}
	return "{ " + strings.Join(pairs, ", ") + " }"
}

func (v *FormatVisitor) VisitMapCreatorRestPair(ctx *parser.MapCreatorRestPairContext) interface{} {
	return v.visitRule(ctx.Expression(0)).(string) + " => " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitSetCreatorRest(ctx *parser.SetCreatorRestContext) interface{} {
	expressions := []string{}
	for _, i := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(i).(string))
	}
	if v.textLen(ctx) > 50 {
		return "{\n" + indent(strings.Join(expressions, ",\n")) + "\n}"
	}
	return "{ " + strings.Join(expressions, ", ") + " }"
}

func (v *FormatVisitor) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	expressions := []string{}
	for _, i := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(i).(string))
	}
	if len(expressions) == 0 {
		return "{}"
	}

	hasNestedInitializer := false
	for _, expression := range expressions {
		if strings.Contains(expression, "{") {
			hasNestedInitializer = true
			break
		}
	}

	wrap := v.wrap || v.textLen(ctx) > 50 || (hasNestedInitializer && len(expressions) > 1)
	if wrap {
		return "{\n" + indent(strings.Join(expressions, ",\n")) + "\n}"
	}
	return "{ " + strings.Join(expressions, ", ") + " }"
}

// Class instance arguments, e.g. (Name = 'Acme', BillingCity = 'Los Angeles') in Account(Name = 'Acme', BillingCity = 'Los Angeles')
func (v *FormatVisitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	expressionList := ctx.ExpressionList()
	if expressionList == nil {
		return "()"
	}
	if v.wrap {
		log.Debug("Visitor says to wrap in VisitArguments")
	}
	if v.textLen(expressionList) > 40 {
		defer restoreWrap(wrap(v))
		return "(\n" + indent(v.visitRule(expressionList).(string)) + "\n)"
	}
	return "(" + v.visitRule(expressionList).(string) + ")"
}

func (v *FormatVisitor) VisitCmpExpression(ctx *parser.CmpExpressionContext) interface{} {
	cmpToken := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	if ctx.ASSIGN() != nil {
		cmpToken += "="
	}
	return v.visitRule(ctx.Expression(0)).(string) + " " + cmpToken + " " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	defer restoreWrap(unwrap(v))
	cmpToken := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	return v.visitRule(ctx.Expression(0)).(string) + " " + cmpToken + " " + v.visitRule(ctx.Expression(1)).(string)
}

func (v *FormatVisitor) VisitInstanceOfExpression(ctx *parser.InstanceOfExpressionContext) interface{} {
	return v.visitRule(ctx.Expression()).(string) + " instanceof " + v.visitRule(ctx.TypeRef()).(string)
}

func (v *FormatVisitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	types := []string{}
	for _, p := range ctx.AllTypeRef() {
		types = append(types, v.visitRule(p).(string))
	}
	sep := ", "
	if v.textLen(ctx) > 80 {
		sep = ",\n"
	}
	return strings.Join(types, sep)
}

func (v *FormatVisitor) VisitFormalParameters(ctx *parser.FormalParametersContext) interface{} {
	params := []string{}
	list := ctx.FormalParameterList()
	if list == nil {
		return "()"
	}
	textLen := v.textLen(ctx)
	wrap := v.wrap || (textLen > 40 && len(list.AllFormalParameter()) > 2) || textLen > 60
	for _, p := range list.AllFormalParameter() {
		if wrap {
			params = append(params, indent(v.visitRule(p).(string)))
		} else {
			params = append(params, v.visitRule(p).(string))
		}
	}
	if wrap {
		return "(\n" + strings.Join(params, ",\n") + ")"
	} else {
		return "(" + strings.Join(params, ", ") + ")"
	}
}

func (v *FormatVisitor) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	args := ""
	if ctx.LPAREN() != nil {
		vals := ""
		if ctx.ElementValuePairs() != nil {
			vals = v.visitRule(ctx.ElementValuePairs()).(string)
		} else {
			vals = v.visitRule(ctx.ElementValue()).(string)
		}
		args = "(" + vals + ")"
	}
	return "@" + v.visitRule(ctx.QualifiedName()).(string) + args
}

func (v *FormatVisitor) VisitElementValuePairs(ctx *parser.ElementValuePairsContext) interface{} {
	pairs := []string{v.visitRule(ctx.ElementValuePair()).(string)}
	for _, p := range ctx.AllDelimitedElementValuePair() {
		pairs = append(pairs, v.visitRule(p).(string))
	}
	return strings.Join(pairs, "")
}

func (v *FormatVisitor) VisitDelimitedElementValuePair(ctx *parser.DelimitedElementValuePairContext) interface{} {
	delimiter := " "
	if ctx.COMMA() != nil {
		delimiter = ", "
	}
	return delimiter + v.visitRule(ctx.ElementValuePair()).(string)
}

func (v *FormatVisitor) VisitElementValuePair(ctx *parser.ElementValuePairContext) interface{} {
	return v.visitRule(ctx.Id()).(string) + "=" + v.visitRule(ctx.ElementValue()).(string)
}

func (v *FormatVisitor) VisitElementValue(ctx *parser.ElementValueContext) interface{} {
	return v.visitRule(ctx.GetChild(0).(antlr.RuleNode))
}

func (v *FormatVisitor) VisitElementValueArrayInitializer(ctx *parser.ElementValueArrayInitializerContext) interface{} {
	values := []string{}
	for _, val := range ctx.AllElementValue() {
		values = append(values, v.visitRule(val).(string))
	}
	trailingComma := ""
	if ctx.TrailingComma() != nil {
		trailingComma = ","
	}
	return "(" + strings.Join(values, ", ") + trailingComma + ")"
}

func (v *FormatVisitor) VisitFormalParameter(ctx *parser.FormalParameterContext) interface{} {
	return v.Modifiers(ctx.AllModifier()) + v.visitRule(ctx.TypeRef()).(string) + " " + ctx.Id().GetText()
}

func (v *FormatVisitor) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	ids := []string{}
	for _, i := range ctx.AllId() {
		ids = append(ids, i.GetText())
	}
	return strings.Join(ids, ".")
}

func (v *FormatVisitor) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	vars := []string{}
	for _, vd := range ctx.AllVariableDeclarator() {
		vars = append(vars, v.visitRule(vd).(string))
	}
	return strings.Join(vars, ", ")
}

func (v *FormatVisitor) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	decl := ctx.Id().GetText()
	if ctx.Expression() == nil {
		return decl
	}
	if v.wrap {
		return decl + " = " + v.visitRule(ctx.Expression()).(string)
	}
	return decl + " = " + v.visitRule(ctx.Expression()).(string)
}

func (v *FormatVisitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	returnType := "void"
	if ctx.TypeRef() != nil {
		returnType = v.visitRule(ctx.TypeRef()).(string)
	}
	body := ";"
	if ctx.Block() != nil {
		body = " " + v.visitRule(ctx.Block()).(string)
	}
	return returnType + " " + ctx.MethodId().GetText() +
		v.visitRule(ctx.FormalParameters()).(string) +
		body
}

func (v *FormatVisitor) VisitTypeRefPrimary(ctx *parser.TypeRefPrimaryContext) interface{} {
	return v.visitRule(ctx.TypeRef()).(string) + ".class"
}

func (v *FormatVisitor) VisitTypeRef(ctx *parser.TypeRefContext) interface{} {
	typeNames := []string{}
	for _, t := range ctx.AllTypeName() {
		typeNames = append(typeNames, v.visitRule(t).(string))
	}

	val := strings.Join(typeNames, ".") + ctx.ArraySubscripts().GetText()
	return val
}

func (v *FormatVisitor) VisitTypeName(ctx *parser.TypeNameContext) interface{} {
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
	return typeName + typeArguments
}

func (v *FormatVisitor) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	return "<" + v.visitRule(ctx.TypeList()).(string) + ">"
}

func (v *FormatVisitor) VisitSoslLiteral(ctx *parser.SoslLiteralContext) interface{} {
	if ctx.BoundExpression() != nil {
		return "[\n" +
			indent("FIND\n"+indent(v.visitRule(ctx.BoundExpression()).(string))+v.visitRule(ctx.SoslClauses()).(string)) + "]"

	}
	return ctx.GetChild(0).(antlr.TerminalNode).GetText() + v.visitRule(ctx.SoslClauses()).(string) + "]"
}

func (v *FormatVisitor) VisitSoslClauses(ctx *parser.SoslClausesContext) interface{} {
	var clauses strings.Builder
	if i := ctx.InSearchGroup(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.ReturningFieldSpecList(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithDivisionAssign(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithDataCategory(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithSnippet(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithHighlight(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithSpellCorrection(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithNetworkIn(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithNetworkAssign(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithPricebookIdAssign(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithMetadataAssign(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.WithModeClause(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.LimitClause(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.UpdateListClause(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	return clauses.String()
}

func (v *FormatVisitor) VisitInSearchGroup(ctx *parser.InSearchGroupContext) interface{} {
	return "IN " + v.visitRule(ctx.SearchGroup()).(string)
}

func (v *FormatVisitor) VisitSearchGroup(ctx *parser.SearchGroupContext) interface{} {
	return strings.ToUpper(ctx.GetChild(0).(antlr.TerminalNode).GetText()) + " FIELDS"
}

func (v *FormatVisitor) VisitReturningFieldSpecList(ctx *parser.ReturningFieldSpecListContext) interface{} {
	return "RETURNING " + v.visitRule(ctx.FieldSpecList()).(string)
}

func (v *FormatVisitor) VisitFieldSpecList(ctx *parser.FieldSpecListContext) interface{} {
	list := []string{v.visitRule(ctx.FieldSpec()).(string)}
	for _, f := range ctx.AllFieldSpecList() {
		list = append(list, v.visitRule(f).(string))
	}
	return strings.Join(list, ",\n")
}

func (v *FormatVisitor) VisitFieldSpec(ctx *parser.FieldSpecContext) interface{} {
	if ctx.FieldSpecClauses() == nil {
		return v.visitRule(ctx.SoslId())
	}
	return v.visitRule(ctx.SoslId()).(string) + v.visitRule(ctx.FieldSpecClauses()).(string)
}

func (v *FormatVisitor) VisitWithModeClause(ctx *parser.WithModeClauseContext) interface{} {
	if ctx.USER_MODE() != nil {
		return "WITH USER_MODE"
	}
	return "WITH SYSTEM_MODE"
}

func (v *FormatVisitor) VisitWithHighlight(_ *parser.WithHighlightContext) interface{} {
	return "WITH HIGHLIGHT"
}

func (v *FormatVisitor) VisitWithSpellCorrection(ctx *parser.WithSpellCorrectionContext) interface{} {
	if b := ctx.BooleanLiteral(); b != nil {
		return "WITH SPELL_CORRECTION = " + strings.ToLower(b.GetText())
	}
	return "WITH SPELL_CORRECTION"
}

func (v *FormatVisitor) VisitFieldSpecClauses(ctx *parser.FieldSpecClausesContext) interface{} {
	var clauses strings.Builder
	clauses.WriteString("(\n" + indent(v.visitRule(ctx.FieldList()).(string)))
	if i := ctx.LogicalExpression(); i != nil {
		clauses.WriteString("\nWHERE\n" + indent(v.visitRule(i).(string)))
	}
	if i := ctx.SoslId(); i != nil {
		clauses.WriteString("\nUSING LISTVIEW =  " + v.visitRule(i).(string))
	}
	if i := ctx.FieldOrderList(); i != nil {
		clauses.WriteString("\nORDER BY " + v.visitRule(i).(string))
	}
	if i := ctx.LimitClause(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	if i := ctx.OffsetClause(); i != nil {
		clauses.WriteString("\n" + v.visitRule(i).(string))
	}
	clauses.WriteString(")")
	return clauses.String()
}

func (v *FormatVisitor) VisitFieldList(ctx *parser.FieldListContext) interface{} {
	list := []string{v.visitRule(ctx.SoslId()).(string)}
	for _, f := range ctx.AllFieldList() {
		list = append(list, v.visitRule(f).(string))
	}
	return strings.Join(list, ",\n")
}

func (v *FormatVisitor) VisitSoslId(ctx *parser.SoslIdContext) interface{} {
	list := []string{v.visitRule(ctx.Id()).(string)}
	for _, f := range ctx.AllSoslId() {
		list = append(list, v.visitRule(f).(string))
	}
	return strings.Join(list, ".")
}
