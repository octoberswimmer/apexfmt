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

func indent(text string) string {
	var indentedText strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		indentedText.WriteString("\t" + scanner.Text() + "\n")
	}

	return indentedText.String()
}
