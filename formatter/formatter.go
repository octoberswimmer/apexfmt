package formatter

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

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
	switch {
	case t.ClassDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "IT'S A CLASS")
		return fmt.Sprintf("%s%s", modifiers(t.AllModifier()), v.visitRule(t.ClassDeclaration()).(string))
	case t.InterfaceDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "IT'S AN INTERFACE")
		return fmt.Sprintf("%s%s", modifiers(t.AllModifier()), v.visitRule(t.InterfaceDeclaration()).(string))
	case t.EnumDeclaration() != nil:
		enum := t.EnumDeclaration()
		constants := []string{}
		if enum.EnumConstants() != nil {
			for _, e := range enum.EnumConstants().AllId() {
				constants = append(constants, e.GetText())
			}
		}
		fmt.Fprintln(os.Stderr, "IT'S AN ENUM")
		return fmt.Sprintf("enum %s {%s}", enum.Id().GetText(), strings.Join(constants, ", "))
	}
	return ""
}

func (v *Visitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN THE CLASS DECLARATION")
	extends := ""
	if ctx.EXTENDS() != nil {
		extends = fmt.Sprintf(" extends %s ", v.VisitTypeRef(ctx.TypeRef().(*parser.TypeRefContext)))
	}
	implements := ""
	if ctx.IMPLEMENTS() != nil {
		extends = fmt.Sprintf(" implements %s ", v.VisitTypeList(ctx.TypeList().(*parser.TypeListContext)))
	}
	return fmt.Sprintf("class %s%s%s{\n%s\n}\n", ctx.Id().GetText(),
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
	fmt.Fprintln(os.Stderr, "IN THE INTERFACE DECLARATION")
	extends := ""
	if ctx.EXTENDS() != nil {
		extends = fmt.Sprintf(" extends %s ", v.VisitTypeList(ctx.TypeList().(*parser.TypeListContext)))
	}
	return fmt.Sprintf("interface %s%s {\n%s\n}\n", ctx.Id().GetText(), extends, indent(v.visitRule(ctx.InterfaceBody()).(string)))
}

func (v *Visitor) VisitInterfaceBody(ctx *parser.InterfaceBodyContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN THE INTERFACE BODY")
	declarations := []string{}
	for _, d := range ctx.AllInterfaceMethodDeclaration() {
		declarations = append(declarations, v.visitRule(d).(string))
	}
	return strings.Join(declarations, "\n")
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
		static := ""
		if ctx.STATIC() != nil {
			static = "static "
		}
		fmt.Fprintln(os.Stderr, "GOT A BLOCK")
		return fmt.Sprintf("%s%s", static, indent(v.VisitBlock(ctx.Block().(*parser.BlockContext)).(string)))
	case ctx.MemberDeclaration() != nil:
		return fmt.Sprintf("%s%s", modifiers(ctx.AllModifier()), v.visitRule(ctx.MemberDeclaration()))
	}
	return ""
}

func (v *Visitor) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	// fmt.Fprintln(os.Stderr, "IN MEMBER DECLARATION")
	switch {
	case ctx.MethodDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "FOUND METHOD DECLARATION")
		return v.VisitMethodDeclaration(ctx.MethodDeclaration().(*parser.MethodDeclarationContext))
	case ctx.FieldDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "FOUND FIELD DECLARATION")
		return v.VisitFieldDeclaration(ctx.FieldDeclaration().(*parser.FieldDeclarationContext))
	case ctx.ConstructorDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "FOUND CONSTRUCTOR DECLARATION")
		return v.VisitConstructorDeclaration(ctx.ConstructorDeclaration().(*parser.ConstructorDeclarationContext))
	case ctx.InterfaceDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "FOUND INTERFACE DECLARATION")
		return v.VisitInterfaceDeclaration(ctx.InterfaceDeclaration().(*parser.InterfaceDeclarationContext))
	case ctx.ClassDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "FOUND CLASS DECLARATION")
		return v.VisitClassDeclaration(ctx.ClassDeclaration().(*parser.ClassDeclarationContext))
	case ctx.EnumDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "FOUND ENUM DECLARATION")
		return v.VisitEnumDeclaration(ctx.EnumDeclaration().(*parser.EnumDeclarationContext))
	case ctx.PropertyDeclaration() != nil:
		fmt.Fprintln(os.Stderr, "FOUND PROPERTY DECLARATION")
		return v.visitRule(ctx.PropertyDeclaration())
	}
	fmt.Fprintln(os.Stderr, "FOUND UNEXPECTED DECLARATION")
	return ""
}

func (v *Visitor) VisitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN INTERFACE METHOD DECLARATION")
	returnType := "void"
	if ctx.TypeRef() != nil {
		returnType = v.visitRule(ctx.TypeRef()).(string)
	}
	return fmt.Sprintf("%s%s %s%s;", modifiers(ctx.AllModifier()), returnType, ctx.Id().GetText(), v.visitRule(ctx.FormalParameters()))
}

func (v *Visitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN FIELD DECLARATION")
	return fmt.Sprintf("%s %s;", v.visitRule(ctx.TypeRef()), v.visitRule(ctx.VariableDeclarators()))
}

func (v *Visitor) VisitPropertyDeclaration(ctx *parser.PropertyDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN PROPERTY DECLARATION")
	propertyBlocks := []string{}
	if ctx.AllPropertyBlock() != nil {
		for _, p := range ctx.AllPropertyBlock() {
			propertyBlocks = append(propertyBlocks, v.VisitPropertyBlock(p.(*parser.PropertyBlockContext)).(string))
		}
	}
	return fmt.Sprintf("%s %s {\n%s}\n", v.visitRule(ctx.TypeRef()), ctx.Id().GetText(), strings.Join(propertyBlocks, "\n"))
}

func (v *Visitor) VisitPropertyBlock(ctx *parser.PropertyBlockContext) interface{} {
	if ctx.Getter() != nil {
		return fmt.Sprintf("%s%s", modifiers(ctx.AllModifier()), v.VisitGetter(ctx.Getter().(*parser.GetterContext)))
	} else {
		return fmt.Sprintf("%s%s", modifiers(ctx.AllModifier()), v.VisitSetter(ctx.Setter().(*parser.SetterContext)))
	}
}

func (v *Visitor) VisitGetter(ctx *parser.GetterContext) interface{} {
	if ctx.SEMI() != nil {
		return "get;"
	} else {
		return fmt.Sprintf("get %s", v.VisitBlock(ctx.Block().(*parser.BlockContext)))
	}
}

func (v *Visitor) VisitSetter(ctx *parser.SetterContext) interface{} {
	if ctx.SEMI() != nil {
		return "set;"
	} else {
		return fmt.Sprintf("set %s", v.VisitBlock(ctx.Block().(*parser.BlockContext)))
	}
}

func (v *Visitor) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN CONSTRUCTOR DECLARATION")
	return fmt.Sprintf("%s%s %s\n", v.visitRule(ctx.QualifiedName()), v.visitRule(ctx.FormalParameters()), v.visitRule(ctx.Block()).(string))
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	statements := []string{}
	for _, stmt := range ctx.AllStatement() {
		statements = append(statements, v.VisitStatement(stmt.(*parser.StatementContext)).(string))
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
		return v.VisitBlock(s)
	case *parser.IfStatementContext:
		return v.VisitIfStatement(s)
	case *parser.ForStatementContext:
		return v.VisitForStatement(s)
	case *parser.ExpressionStatementContext:
		return fmt.Sprintf("%s;", v.VisitExpressionStatement(s))
	case *parser.ReturnStatementContext:
		return fmt.Sprintf("%s", v.VisitReturnStatement(s))
	case *parser.LocalVariableDeclarationStatementContext:
		return fmt.Sprintf("%s", v.VisitLocalVariableDeclarationStatement(s))
	}
	return fmt.Sprintf("UNHANDLED STATEMENT: %T %s", ctx.GetChild(0).(antlr.RuleNode), ctx.GetText())
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	elseStatement := ""
	if ctx.ELSE() != nil {
		elseStatement = " } else { " + v.VisitStatement(ctx.Statement(1).(*parser.StatementContext)).(string)
	}
	return fmt.Sprintf("if %s {\n%s}%s", v.VisitParExpression(ctx.ParExpression().(*parser.ParExpressionContext)),
		v.VisitStatement(ctx.Statement(0).(*parser.StatementContext)),
		elseStatement)
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
	return fmt.Sprintf("%s %s : %s", v.VisitTypeRef(ctx.TypeRef().(*parser.TypeRefContext)), v.visitRule(ctx.Id()), v.visitRule(ctx.Expression()))
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
	return fmt.Sprintf("%s%s %s", modifiers(ctx.AllModifier()), v.visitRule(ctx.TypeRef()), v.visitRule(ctx.VariableDeclarators()))
}

func (v *Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return fmt.Sprintf("return %s;", v.visitRule(e))
	}
	return "return;"
}

func (v *Visitor) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return fmt.Sprintf("(|%T|%s)", ctx.Expression(), v.visitRule(ctx.Expression()))
}

func (v *Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	return v.visitRule(ctx.Expression())
	/*
		switch e := ctx.Expression().(type) {
		case *parser.AssignExpressionContext:
			return v.VisitAssignExpression(e)
		default:
			return fmt.Sprintf("UNHANDLED EXPRESSION TYPE %T: %s", e, e.GetText())
		}
	*/
}

func (v *Visitor) VisitAssignExpression(ctx *parser.AssignExpressionContext) interface{} {
	assignmentToken := ctx.GetChild(1).(antlr.TerminalNode)
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), assignmentToken.GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *Visitor) VisitCondExpresssion(ctx *parser.CondExpressionContext) interface{} {
	return fmt.Sprintf("%s ? %s : %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)), v.visitRule(ctx.Expression(2)))
}

func (v *Visitor) VisitLogAndExpression(ctx *parser.LogAndExpressionContext) interface{} {
	return fmt.Sprintf("%s && %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
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

func (v *Visitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return "GOT UNEXPECTED VISIT TO EXPRESSION"
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
	return fmt.Sprintf("TODO: IMPLEMENT SOQL PRIMARY")
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
	fmt.Fprintln(os.Stderr, "IN FORMAL PARAMETERS")
	params := []string{}
	list := ctx.FormalParameterList()
	if list == nil {
		return "()"
	}
	for _, p := range list.AllFormalParameter() {
		params = append(params, v.visitRule(p).(string))
	}
	val := fmt.Sprintf("(%s)", strings.Join(params, ", "))
	fmt.Fprintf(os.Stderr, "FORMAL PARAMETERS:|%s|\n", val)
	return val
}

func modifiers(ctxs []parser.IModifierContext) string {
	mods := []string{}
	for _, m := range ctxs {
		mods = append(mods, m.GetText())
	}
	modifiers := strings.Join(mods, " ")
	if modifiers != "" {
		modifiers += " "
	}
	return modifiers
}

func (v *Visitor) VisitFormalParameter(ctx *parser.FormalParameterContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN FORMAL PARAMETER")
	return fmt.Sprintf("%s%s %s", modifiers(ctx.AllModifier()), v.visitRule(ctx.TypeRef()), ctx.Id().GetText())
}

func (v *Visitor) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN QUALIFIED NAME")
	ids := []string{}
	for _, i := range ctx.AllId() {
		ids = append(ids, i.GetText())
	}
	return strings.Join(ids, ".")
}

func (v *Visitor) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN VARIABLE DECLARATORS")
	vars := []string{}
	for _, vd := range ctx.AllVariableDeclarator() {
		vars = append(vars, v.visitRule(vd).(string))
	}
	return strings.Join(vars, ", ")
}

func (v *Visitor) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN VARIABLE DECLARATOR")
	decl := ctx.Id().GetText()
	if ctx.Expression() != nil {
		decl = fmt.Sprintf("%s = %s", decl, v.visitRule(ctx.Expression()))
	}
	return decl
}

func (v *Visitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	fmt.Fprintln(os.Stderr, "IN METHOD DECLARATION")
	returnType := "void"
	if ctx.TypeRef() != nil {
		returnType = v.visitRule(ctx.TypeRef()).(string)
	}
	body := ";"
	if ctx.Block() != nil {
		body = " " + v.VisitBlock(ctx.Block().(*parser.BlockContext)).(string)
	}
	return fmt.Sprintf("%s %s%s%s\n", returnType, ctx.Id().GetText(),
		v.VisitFormalParameters(ctx.FormalParameters().(*parser.FormalParametersContext)),
		body)
}

func (v *Visitor) VisitTypeRef(ctx *parser.TypeRefContext) interface{} {
	typeNames := []string{}
	for _, t := range ctx.AllTypeName() {
		typeNames = append(typeNames, t.GetText())
	}

	fmt.Fprintf(os.Stderr, "TYPE NAMES:|%s|\n", strings.Join(typeNames, "."))

	val := fmt.Sprintf("%s%s", strings.Join(typeNames, "."), ctx.ArraySubscripts().GetText())
	return val
}
