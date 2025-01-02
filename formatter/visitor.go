package formatter

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

const (
	WHITESPACE_CHANNEL = 2
	COMMENTS_CHANNEL   = 3
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

func (v *FormatVisitor) VisitRule(node antlr.RuleNode) interface{} {
	return v.visitRule(node)
}

func (v *FormatVisitor) visitRule(node antlr.RuleNode) interface{} {
	start := node.(antlr.ParserRuleContext).GetStart()
	var beforeWhitespace, beforeComments []antlr.Token
	if start != nil && len(v.tokens.GetAllTokens()) > 0 {
		beforeWhitespace = v.tokens.GetHiddenTokensToLeft(start.GetTokenIndex(), WHITESPACE_CHANNEL)
		beforeComments = v.tokens.GetHiddenTokensToLeft(start.GetTokenIndex(), COMMENTS_CHANNEL)
	}
	result := node.Accept(v)
	if result == nil {
		panic(fmt.Sprintf("MISSING VISIT FUNCTION FOR %T", node))
	}
	commentsWithNewlines := commentsWithTrailingNewlines(beforeComments, beforeWhitespace)
	hasComments := beforeComments != nil && len(beforeComments) > 0
	if beforeComments != nil {
		comments := []string{}
		for _, c := range beforeComments {
			index := c.GetTokenIndex()
			if _, seen := v.commentsOutput[index]; !seen {
				// Mark the start and end of comments so we can remove indentation
				// added to multi-line comments, preserving the whitespace within
				// them.  See removeIndentationFromComment.
				if n, exists := commentsWithNewlines[index]; exists {
					comments = append(comments, "\uFFFA"+c.GetText()+"\uFFFB"+strings.Repeat("\n", n))
				} else {
					comments = append(comments, "\uFFFA"+c.GetText()+"\uFFFB")
				}
				v.commentsOutput[index] = struct{}{}
			}
		}
		if len(comments) > 0 {
			allComments := strings.Join(comments, "")
			containsNewline := strings.Contains(allComments, "\n")
			if !containsNewline {
				result = fmt.Sprintf("%s %s", strings.TrimSuffix(strings.TrimPrefix(allComments, "\uFFFA"), "\uFFFB"), result)
			} else {
				result = fmt.Sprintf("%s%s", allComments, result)
			}
		}
	}
	if beforeWhitespace != nil {
		injectNewline := false
		for _, c := range beforeWhitespace {
			if !hasComments && len(strings.Split(c.GetText(), "\n")) > 2 {
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
	stop := node.(antlr.ParserRuleContext).GetStop()
	var afterWhitespace, afterComments []antlr.Token
	if stop == nil {
		return result
	}
	if len(v.tokens.GetAllTokens()) > 0 {
		afterWhitespace = v.tokens.GetHiddenTokensToRight(stop.GetTokenIndex(), WHITESPACE_CHANNEL)
		afterComments = v.tokens.GetHiddenTokensToRight(stop.GetTokenIndex(), COMMENTS_CHANNEL)
	}

	if afterComments == nil {
		return result
	}
	afterCommentsWithLeadingNewlines := commentsWithLeadingNewlines(afterComments, afterWhitespace)
	comments := []string{}

	for _, c := range afterComments {
		index := c.GetTokenIndex()
		if _, seen := v.commentsOutput[index]; !seen {
			// Mark the start and end of comments so we can remove indentation
			// added to multi-line comments, preserving the whitespace within
			// them.  See removeIndentationFromComment.
			if n, exists := afterCommentsWithLeadingNewlines[index]; exists {
				comments = append(comments, strings.Repeat("\n", n)+"\uFFFA"+c.GetText()+"\uFFFB")
			} else {
				comments = append(comments, "\uFFFA"+c.GetText()+"\uFFFB")
			}
			v.commentsOutput[index] = struct{}{}
		}
	}

	if len(comments) > 0 {
		allComments := strings.Join(comments, "")
		containsNewline := strings.Contains(allComments, "\n")
		if !containsNewline {
			result = fmt.Sprintf("%s %s", result, strings.TrimSuffix(strings.TrimPrefix(allComments, "\uFFFA"), "\uFFFB"))
		} else {
			result = fmt.Sprintf("%s%s", result, allComments)
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

// Find comments that have trailing newlines
func commentsWithTrailingNewlines(comments []antlr.Token, whitespace []antlr.Token) map[int]int {
	result := make(map[int]int)

	whitespaceMap := make(map[int]antlr.Token)
	for _, ws := range whitespace {
		whitespaceMap[ws.GetTokenIndex()] = ws
	}

	for _, comment := range comments {
		commentIndex := comment.GetTokenIndex()

		// Find the immediate next token index
		nextTokenIndex := commentIndex + 1

		// Check if the next token is whitespace
		if ws, exists := whitespaceMap[nextTokenIndex]; exists {
			// Check if the whitespace contains a newline
			if strings.Contains(ws.GetText(), "\n") {
				result[commentIndex] = len(strings.Split(ws.GetText(), "\n")) - 1
			}
		}
	}

	return result
}

// Find comments that have leading newlines
func commentsWithLeadingNewlines(comments []antlr.Token, whitespace []antlr.Token) map[int]int {
	result := make(map[int]int)

	whitespaceMap := make(map[int]antlr.Token)
	for _, ws := range whitespace {
		whitespaceMap[ws.GetTokenIndex()] = ws
	}

	for _, comment := range comments {
		commentIndex := comment.GetTokenIndex()

		// Find the immediate previous token index
		prevTokenIndex := commentIndex - 1

		// Check if the next token is whitespace
		if ws, exists := whitespaceMap[prevTokenIndex]; exists {
			// Check if the whitespace contains a newline
			if strings.Contains(ws.GetText(), "\n") {
				result[commentIndex] = len(strings.Split(ws.GetText(), "\n")) - 1
			}
		}
	}

	return result
}
