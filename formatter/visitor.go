package formatter

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

type FormatVisitor struct {
	tokens         *antlr.CommonTokenStream
	commentsOutput map[int]struct{}
	newlinesOutput map[int]struct{}
	parser.BaseApexParserVisitor
	wrap bool
}

func NewFormatVisitor(tokens *antlr.CommonTokenStream) *FormatVisitor {
	return &FormatVisitor{
		tokens:         tokens,
		commentsOutput: make(map[int]struct{}),
		newlinesOutput: make(map[int]struct{}),
	}
}

func (v *FormatVisitor) visitRule(node antlr.RuleNode) interface{} {
	start := node.(antlr.ParserRuleContext).GetStart()
	beforeWhitespace := v.tokens.GetHiddenTokensToLeft(start.GetTokenIndex(), 2)
	beforeComments := v.tokens.GetHiddenTokensToLeft(start.GetTokenIndex(), 3)
	result := node.Accept(v)
	if result == nil {
		panic(fmt.Sprintf("MISSING VISIT FUNCTION FOR %T", node))
	}
	if beforeComments != nil {
		comments := []string{}
		for _, c := range beforeComments {
			if _, seen := v.commentsOutput[c.GetTokenIndex()]; !seen {
				comments = append(comments, cleanWhitespace(c.GetText()))
				v.commentsOutput[c.GetTokenIndex()] = struct{}{}
			}
		}
		if len(comments) > 0 {
			result = fmt.Sprintf("%s\n%s", strings.Join(comments, "\n"), result)
		}
	}
	if beforeWhitespace != nil {
		injectNewline := false
		for _, c := range beforeWhitespace {
			if len(strings.Split(c.GetText(), "\n")) > 2 {
				if _, seen := v.newlinesOutput[c.GetTokenIndex()]; !seen {
					v.newlinesOutput[c.GetTokenIndex()] = struct{}{}
					injectNewline = true
				}
			}
		}
		if injectNewline {
			result = fmt.Sprintf("\n%s", result)
		}
	}
	return result
}

func (v *FormatVisitor) Modifiers(ctxs []parser.IModifierContext) string {
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

func (v *FormatVisitor) indent(text string) string {
	return v.indentTo(text, 1)
}

func (v *FormatVisitor) indentTo(text string, indents int) string {
	var indentedText strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(text))
	isFirstLine := true

	for scanner.Scan() {
		if isFirstLine {
			isFirstLine = false
		} else {
			indentedText.WriteString("\n")
		}
		if scanner.Text() != "" {
			indentedText.WriteString(strings.Repeat("\t", indents) + scanner.Text())
		} else {
			indentedText.WriteString(scanner.Text())
		}
	}

	return indentedText.String()
}

// Remove leading tabs
func cleanWhitespace(input string) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimRight(strings.TrimLeft(line, "\t"), " \t")
	}

	return strings.Join(lines, "\n")
}

func restoreWrap(v *FormatVisitor, reset bool) *FormatVisitor {
	v.wrap = reset
	return v
}

func wrap(v *FormatVisitor) (*FormatVisitor, bool) {
	old := v.wrap
	v.wrap = true
	return v, old
}

func unwrap(v *FormatVisitor) (*FormatVisitor, bool) {
	old := v.wrap
	v.wrap = false
	return v, old
}
