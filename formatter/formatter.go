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

func (v *Visitor) Modifiers(ctxs []parser.IModifierContext) string {
	mods := []string{}
	annotations := []string{}
	for _, m := range ctxs {
		if m.Annotation() != nil {
			annotations = append(annotations, v.visitRule(m.Annotation()).(string))
		} else {
			for _, word := range m.GetChildren() {
				mods = append(mods, word.(antlr.TerminalNode).GetText())
			}
		}
	}
	var m strings.Builder
	if len(annotations) > 0 {
		m.WriteString(strings.Join(annotations, "\n") + "\n")
	}
	if len(mods) > 0 {
		m.WriteString(strings.Join(mods, " ") + " ")
	}
	return m.String()
}

func indent(text string) string {
	var indentedText strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(text))
	isFirstLine := true

	for scanner.Scan() {
		if isFirstLine {
			isFirstLine = false
		}
		if scanner.Text() != "" {
			if !isFirstLine {
				indentedText.WriteString("\n")
			}
			indentedText.WriteString("\t" + scanner.Text())
		}
	}

	return indentedText.String()
}
