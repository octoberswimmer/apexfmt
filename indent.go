package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func TreesIndentedStringTree(tree antlr.Tree, indent string, ruleNames []string, recog antlr.Recognizer) string {

	if recog != nil {
		ruleNames = recog.GetRuleNames()
	}

	s := antlr.TreesGetNodeText(tree, ruleNames, nil)
	fmt.Println("GOT", s)

	if _, ok := tree.(antlr.RuleNode); ok {
		fmt.Println("RULE NODE")
		s = " " + s
	}

	if _, ok := tree.(antlr.TerminalNode); ok {
		fmt.Println("TERMINAL NODE")
		s = " " + s
	}
	fmt.Printf("\n\n")

	s = antlr.EscapeWhitespace(s, false)
	c := tree.GetChildCount()
	if c == 0 {
		return s
	}
	res := ""
	if len(indent) > 0 {
		res = res + "\n"
	}
	res = res + indent + "(" + s + " "
	if c > 0 {
		s = TreesIndentedStringTree(tree.GetChild(0), indent+" ", ruleNames, nil)
		res += s
	}
	for i := 1; i < c; i++ {
		s = TreesIndentedStringTree(tree.GetChild(i), indent+" ", ruleNames, nil)
		res += s
	}
	if len(indent) > 0 {
		indent = indent[:len(indent)-1]
	}

	res += ")" + "\n" + indent
	return res
}
