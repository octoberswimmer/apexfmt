package main

import (
	"bufio"
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
	return fmt.Sprintf("class %s%s%s{\n%s\n}", ctx.Id().GetText(),
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
	return fmt.Sprintf("interface %s%s {\n%s\n}", ctx.Id().GetText(), extends, indent(v.visitRule(ctx.InterfaceBody()).(string)))
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
	return fmt.Sprintf("%s%s %s", v.visitRule(ctx.QualifiedName()), v.visitRule(ctx.FormalParameters()), v.visitRule(ctx.Block()).(string))
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return fmt.Sprintf("{\n%s}", indent("BLOCK"))
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
	return fmt.Sprintf("%s %s%s%s", returnType, ctx.Id().GetText(),
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
