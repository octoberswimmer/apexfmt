package formatter

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
	log "github.com/sirupsen/logrus"
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

type HiddenTokenDirection int

const (
	HiddenTokenDirectionBefore HiddenTokenDirection = iota
	HiddenTokenDirectionAfter
)

type CommentPosition int

const (
	PositionBefore CommentPosition = iota
	PositionAfter
)

var manyNewlines = regexp.MustCompile(`\n{3,}`)

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
	start, stop := getStartStop(node)

	// Collect and interleave comments and whitespace before the node
	beforeHiddenTokens := interleaveHiddenTokens(
		getHiddenTokens(v.tokens, start, HiddenTokenDirectionBefore),
	)

	result := node.Accept(v)
	if result == nil {
		panic(fmt.Sprintf("MISSING VISIT FUNCTION FOR %T", node))
	}

	result = appendHiddenTokens(v, result, beforeHiddenTokens, PositionBefore)

	// Collect and interleave comments and whitespace after the node
	afterHiddenTokens := interleaveHiddenTokens(
		getHiddenTokens(v.tokens, stop, HiddenTokenDirectionAfter),
	)

	_ = afterHiddenTokens
	result = appendHiddenTokens(v, result, afterHiddenTokens, PositionAfter)

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

func indent(text string) string {
	return indentTo(text, 1)
}

// SplitLeadingFFFAOrFFFBOrNewline splits the input data into tokens based on the following rules:
//
//  1. If a line starts with \ufffa or \ufffb (possibly preceded by whitespace),
//     and the line contains only the delimiter, return the entire line as a separate token.
//
//  2. If a line starts with \ufffa or \ufffb (possibly preceded by whitespace),
//     and the line contains additional content after the delimiter,
//     split the line into two tokens:
//     - The delimiter (\ufffa or \ufffb) including any leading whitespace.
//     - The remaining content.
//
//  3. If \ufffa or \ufffb appear anywhere else in a line,
//     split the line into two tokens:
//     - Content before the delimiter.
//     - The delimiter and the remaining content.
//
//  4. Otherwise, split lines normally based on newline characters.
//
//  5. If \ufff9 and \ufffb appear in the same line (with \ufff9 < \ufffb),
//     treat that line as a single token, preserving inline comment.
//
//  6. At EOF, return any remaining data as the final token.
//
// This function ensures that \ufffa, \ufffb, and inline \ufff9-\ufffb comments
// are handled correctly based on their positions.
func SplitLeadingFFFAOrFFFBOrNewline(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Define \ufffa and \ufffb
	fffa := []byte("\ufffa")
	fffb := []byte("\ufffb")
	inlineCommentStart := []byte("\ufff9")

	// Handle empty input
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the first newline
	newlineIdx := bytes.IndexByte(data, '\n')

	// ----------------------------------------------------------------
	// 1. NO NEWLINE FOUND BUT EOF => Return Last Line
	// ----------------------------------------------------------------
	if newlineIdx == -1 {
		if atEOF {
			line := data

			// --- Inline comment check (rule #5) ---
			if hasInlineComment(line, inlineCommentStart, fffb) {
				// Return the entire line as a single token
				return len(data), data, nil
			}
			// --------------------------------------

			trimmed := bytes.TrimLeft(line, " \t")

			// If line starts with \ufffa or \ufffb and contains only that delimiter
			if bytes.HasPrefix(trimmed, fffa) || bytes.HasPrefix(trimmed, fffb) {
				var delimiter []byte
				if bytes.HasPrefix(trimmed, fffa) {
					delimiter = fffa
				} else {
					delimiter = fffb
				}

				if bytes.Equal(trimmed, delimiter) {
					// Line contains only the delimiter
					return len(data), line, nil
				}

				// Line starts with the delimiter and has additional content
				delimiterIdx := bytes.Index(line, delimiter)
				if delimiterIdx != -1 {
					// Include leading whitespace in delimiter token
					delimiterEnd := delimiterIdx + len(delimiter)
					return delimiterEnd, line[:delimiterEnd], nil
				}
			}

			// Otherwise, no delimiters => return the entire line
			return len(data), data, nil
		}

		// No newline, not at EOF => request more data
		return 0, nil, nil
	}

	// ----------------------------------------------------------------
	// 2. WE FOUND A NEWLINE => Extract the line
	// ----------------------------------------------------------------
	line := data[:newlineIdx]

	// --- Inline comment check (rule #5) ---
	// If line has \ufff9 and \ufffb in the correct order, keep it as one token.
	if hasInlineComment(line, inlineCommentStart, fffb) {
		return newlineIdx + 1, line, nil
	}
	// --------------------------------------

	trimmed := bytes.TrimLeft(line, " \t")

	// ----------------------------------------------------------------
	// 2a. If line starts with \ufffa or \ufffb
	// ----------------------------------------------------------------
	if bytes.HasPrefix(trimmed, fffa) || bytes.HasPrefix(trimmed, fffb) {
		var delimiter []byte
		if bytes.HasPrefix(trimmed, fffa) {
			delimiter = fffa
		} else {
			delimiter = fffb
		}
		delimiterLen := len(delimiter)

		if bytes.Equal(trimmed, delimiter) {
			// Line contains only the delimiter
			return newlineIdx + 1, line, nil
		}

		// Line starts with delimiter but has more content
		delimiterIdx := bytes.Index(line, delimiter)
		if delimiterIdx != -1 {
			delimiterEnd := delimiterIdx + delimiterLen
			return delimiterEnd, line[:delimiterEnd], nil
		}
	}

	// ----------------------------------------------------------------
	// 2b. Delimiter Elsewhere in the Line
	// ----------------------------------------------------------------
	fffaIdx := bytes.Index(line, fffa)
	fffbIdx := bytes.Index(line, fffb)

	firstIdx := -1
	if fffaIdx != -1 && (fffbIdx == -1 || fffaIdx < fffbIdx) {
		firstIdx = fffaIdx
	}
	if fffbIdx != -1 && (firstIdx == -1 || fffbIdx < fffaIdx) {
		firstIdx = fffbIdx
	}

	if firstIdx != -1 {
		// Split BEFORE the delimiter
		return firstIdx, line[:firstIdx], nil
	}

	// ----------------------------------------------------------------
	// 2c. No Delimiters => Return Entire Line
	// ----------------------------------------------------------------
	return newlineIdx + 1, line, nil
}

// hasInlineComment checks if the line has \ufff9 and \ufffb in that order
// indicating an inline comment that should remain together.
func hasInlineComment(line, inlineCommentStart, fffb []byte) bool {
	idx9 := bytes.Index(line, inlineCommentStart)
	idxFB := bytes.Index(line, fffb)
	return idx9 != -1 && idxFB != -1 && idx9 < idxFB
}

func indentTo(text string, indents int) string {
	var indentedText strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(text))
	isFirstLine := true

	scanner.Split(SplitLeadingFFFAOrFFFBOrNewline)

	log.Debug(fmt.Sprintf("INDENTING: %q\n", text))

	for scanner.Scan() {
		log.Debug(fmt.Sprintf("INDENTING LINE: %q\n", scanner.Text()))
		if scanner.Text() == "\uFFFB" {
			indentedText.WriteString(scanner.Text())
			continue
		}
		if isFirstLine {
			isFirstLine = false
		} else {
			indentedText.WriteString("\n")
		}
		if scanner.Text() == "" {
			indentedText.WriteString(scanner.Text())
			continue
		}
		t := scanner.Text()
		t = strings.Repeat("\t", indents) + t
		indentedText.WriteString(t)
	}
	log.Debug(fmt.Sprintf("INDENTED:  %q\n\n", indentedText.String()))

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

func interleaveHiddenTokens(whitespace []antlr.Token, comments []antlr.Token) []antlr.Token {
	interleaved := []antlr.Token{}
	allTokens := append(whitespace, comments...)

	// Sort tokens by their position in the stream
	sort.Slice(allTokens, func(i, j int) bool {
		return allTokens[i].GetTokenIndex() < allTokens[j].GetTokenIndex()
	})

	interleaved = append(interleaved, allTokens...)
	return interleaved
}

func appendHiddenTokens(v *FormatVisitor, result interface{}, tokens []antlr.Token, position CommentPosition) interface{} {
	var tokenLines []string
	for _, token := range tokens {
		index := token.GetTokenIndex()
		if _, seen := v.commentsOutput[index]; !seen {
			v.commentsOutput[index] = struct{}{}

			text := token.GetText()
			if token.GetChannel() == COMMENTS_CHANNEL {
				leadingWhitespace := getLeadingWhitespace(text)
				trailingWhitespace := getTrailingWhitespace(text)
				leading := ""
				trailing := ""
				if n := countNewlines(leadingWhitespace); n > 0 {
					leading = strings.Repeat("\n", n)
				} else if len(leadingWhitespace) > 0 && position == PositionAfter {
					leading = " "
				}
				if n := countNewlines(trailingWhitespace); n > 0 {
					trailing = strings.Repeat("\n", n)
				} else if len(trailingWhitespace) > 0 {
					trailing = " "
				}
				// Strip leading whitespace so the comment can be indented to the right location
				text = strings.TrimSpace(text)
				containsNewline := strings.Contains(text, "\n")

				text = fmt.Sprintf("%s%s%s", leading, text, trailing)
				log.Debug(fmt.Sprintf("NORMALIZED COMMENT: %q\n", text))
				if containsNewline {
					text = "\uFFFA" + text + "\uFFFB"
				} else {
					text = "\uFFF9" + text + "\uFFFB"
				}
				log.Debug(fmt.Sprintf("WRAPPED COMMENT: %q\n\n", text))
			} else if token.GetChannel() == WHITESPACE_CHANNEL && countNewlines(text) > 1 {
				text = "\n" // Replace multiple blank lines with a single blank line
			} else {
				// whitespace to ignore
				text = ""
			}

			tokenLines = append(tokenLines, text)
		}
	}

	if len(tokenLines) > 0 {
		tokenText := strings.Join(tokenLines, "")
		switch position {
		case PositionBefore:
			result = fmt.Sprintf("%s%s", tokenText, result)
		case PositionAfter:
			result = fmt.Sprintf("%s%s", result, tokenText)
		}
	}

	return any(result)
}

func countNewlines(text string) int {
	return strings.Count(text, "\n")
}

func getStartStop(node antlr.RuleNode) (start, stop antlr.Token) {
	ctx := node.(antlr.ParserRuleContext)
	return ctx.GetStart(), ctx.GetStop()
}

func getHiddenTokens(tokens *antlr.CommonTokenStream, token antlr.Token, direction HiddenTokenDirection) ([]antlr.Token, []antlr.Token) {
	if token == nil || len(tokens.GetAllTokens()) == 0 {
		return nil, nil
	}

	switch direction {
	case HiddenTokenDirectionBefore:
		return tokens.GetHiddenTokensToLeft(token.GetTokenIndex(), WHITESPACE_CHANNEL),
			tokens.GetHiddenTokensToLeft(token.GetTokenIndex(), COMMENTS_CHANNEL)
	case HiddenTokenDirectionAfter:
		return tokens.GetHiddenTokensToRight(token.GetTokenIndex(), WHITESPACE_CHANNEL),
			tokens.GetHiddenTokensToRight(token.GetTokenIndex(), COMMENTS_CHANNEL)
	default:
		return nil, nil
	}
}

func getLeadingWhitespace(s string) string {
	var i int
	for i = 0; i < len(s); i++ {
		if !unicode.IsSpace(rune(s[i])) {
			break
		}
	}
	return s[:i]
}

func getTrailingWhitespace(s string) string {
	var i int
	for i = len(s) - 1; i >= 0; i-- {
		if !unicode.IsSpace(rune(s[i])) {
			break
		}
	}
	return s[i+1:]
}
