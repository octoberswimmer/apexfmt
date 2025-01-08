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

	if result.(string) == "{}" {
		inbetweenTokens := interleaveHiddenTokens(
			getHiddenTokensBetween(v.tokens, start, stop),
		)
		result = fmt.Sprintf("%s}", appendHiddenTokens(v, "{", inbetweenTokens, PositionAfter))
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

func indent(text string) string {
	return indentTo(text, 1)
}

// SplitLeadingFFFAOrFFFBOrNewline splits the input data like SplitLines, with
// special handling for comments.
//
// Multi-line comments delimited by \uFFFA and \uFFFB are handled as follows:
// \uFFFA and \uFFFB should never have leading text other than whitespace.
//
// \uFFFA can have trailing text.
//
// \uFFFB cannot have trailing text.
//
// Inline comments delimited by \uFFF9 and \uFFFB should always be returned
// unbroken.
func SplitLeadingFFFAOrFFFBOrNewline(data []byte, atEOF bool) (advance int, token []byte, err error) {
	log.Trace(fmt.Sprintf("SPLITTING: %q", string(data)))
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
		log.Trace(fmt.Sprintf("NO NEWLINE IN: %q", string(data)))
		if !atEOF {
			// No newline, not at EOF => request more data
			log.Trace(fmt.Sprintf("REQUESTING MORE DATA: %q", string(data)))
			return 0, nil, nil
		}

		log.Trace(fmt.Sprintf("AT EOF IN: %q", string(data)))
		line := data

		if hasInlineComment(line, inlineCommentStart, fffb) {
			log.Trace(fmt.Sprintf("HAS INLINE COMMENT: %q", string(line)))
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
				log.Trace(fmt.Sprintf("HAS ONLY DELIMITER: %q", string(line)))
				// Line contains only the delimiter
				return len(data), line, nil
			}

			// Line starts with the delimiter and has additional content
			delimiterIdx := bytes.Index(line, delimiter)
			if delimiterIdx != -1 {
				// Include leading whitespace in delimiter token
				delimiterEnd := delimiterIdx + len(delimiter)
				log.Trace(fmt.Sprintf("DELIMITER+: %q", string(line[:delimiterEnd])))
				return delimiterEnd, line[:delimiterEnd], nil
			}
		}

		// Otherwise, no delimiters => return the entire line
		log.Trace(fmt.Sprintf("NO DELIMITERS: %q", string(data)))
		return len(data), data, nil
	}
	log.Trace(fmt.Sprintf("FOUND NEWLINE IN: %q", string(data)))

	// ----------------------------------------------------------------
	// 2. WE FOUND A NEWLINE => Extract the line
	// ----------------------------------------------------------------
	line := data[:newlineIdx]

	// --- Inline comment check ---
	// If line has \ufff9 and \ufffb in the correct order, keep it as one token.
	if hasInlineComment(line, inlineCommentStart, fffb) {
		log.Trace(fmt.Sprintf("INLINE COMMENT: %q", string(data)))
		return newlineIdx + 1, line, nil
	}
	// --------------------------------------

	trimmed := bytes.TrimLeft(line, " \t")

	// ----------------------------------------------------------------
	// If line starts with \ufffb
	// ----------------------------------------------------------------
	if bytes.HasPrefix(trimmed, fffb) {
		var delimiter []byte
		delimiter = fffb
		delimiterLen := len(delimiter)

		if bytes.Equal(trimmed, delimiter) {
			log.Trace(fmt.Sprintf("HAS DELIMITER ONLY: %q", string(line)))
			// Line contains only the delimiter
			return newlineIdx + 1, line, nil
		}

		// Line starts with delimiter but has more content
		delimiterIdx := bytes.Index(line, delimiter)
		if delimiterIdx != -1 {
			delimiterEnd := delimiterIdx + delimiterLen
			log.Trace(fmt.Sprintf("\\uFFFB+: %q", string(line[:delimiterEnd])))
			return delimiterEnd, line[:delimiterEnd], nil
		}
	}

	fffaIdx := bytes.Index(line, fffa)
	fffbIdx := bytes.Index(line, fffb)

	// ----------------------------------------------------------------
	// If line starts with \ufffa
	// ----------------------------------------------------------------
	if bytes.HasPrefix(trimmed, fffa) {
		delimiter := fffa

		if bytes.Equal(trimmed, delimiter) {
			log.Trace(fmt.Sprintf("HAS DELIMITER ONLY: %q", string(line)))
			// Line contains only the delimiter
			return newlineIdx + 1, line, nil
		}

		if fffbIdx != -1 {
			// \ufffb is before the newline
			delimiterEnd := fffbIdx + len("\ufffb")
			log.Trace(fmt.Sprintf("RETURNING UP TO \\uFFFB: %q", string(line[:delimiterEnd])))
			// Advance past the newline after \uFFFB
			return delimiterEnd + 1, line[:delimiterEnd], nil
		}
		if f := bytes.Index(data, fffb); f == newlineIdx+1 {
			delimiterLen := len(fffb)
			log.Trace(fmt.Sprintf("RETURNING UP TO NEWLINE WITH \\uFFFB: %q", string(data[:f+delimiterLen])))
			return f + delimiterLen, data[:f+delimiterLen], nil
		}

		// Line starts with delimiter but has more content
		delimiterIdx := bytes.Index(line, delimiter)
		if delimiterIdx != -1 {
			log.Trace(fmt.Sprintf("\\uFFFA+: %q", string(line)))
			return newlineIdx + 1, line, nil
		}
	}

	// ----------------------------------------------------------------
	// Delimiter Elsewhere in the Line
	// ----------------------------------------------------------------

	if fffaIdx != -1 && (fffbIdx == -1 || fffaIdx < fffbIdx) {
		// Split BEFORE the delimiter
		log.Trace(fmt.Sprintf("HAS \\uFFFA IN LINE: %q", string(line[:fffaIdx])))
		return fffaIdx, line[:fffaIdx], nil
	}
	if fffbIdx != -1 && (fffaIdx == -1 || fffbIdx < fffaIdx) {
		delimiterLen := len(fffb)
		// Split AFTER the delimiter
		log.Trace(fmt.Sprintf("HAS \\uFFFB IN LINE: %q", string(line[:fffbIdx+delimiterLen])))
		advance := 0
		if bytes.IndexByte(line[:fffbIdx+delimiterLen], '\n') == 0 {
			// Advance past the newline after \uFFFB
			advance = 1
		}
		return fffbIdx + delimiterLen + advance, line[:fffbIdx+delimiterLen], nil
	}

	// ----------------------------------------------------------------
	// 2c. No Delimiters => Return Entire Line
	// ----------------------------------------------------------------
	log.Trace(fmt.Sprintf("NO DELIMITER: %q", string(line)))
	var fffbFollowsNewlines = regexp.MustCompile(`(s?)^` + "\n+\uFFFB")
	if len(line) > 0 && fffbFollowsNewlines.Match(data[newlineIdx:]) {
		// \uFFFB follows newline.  We want to keep the newline by returning an
		// extra empty line so we don't advance over the newline.
		return newlineIdx, line, nil
	}
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
		t := scanner.Text()
		log.Trace(fmt.Sprintf("INDENTING LINE: %q\n", t))
		if t == "\uFFFB" {
			indentedText.WriteString(t)
			continue
		}
		if isFirstLine {
			isFirstLine = false
		} else if !strings.HasPrefix(t, "\uFFFA") && !strings.HasPrefix(t, "\uFFF9") {
			indentedText.WriteString("\n")
		}
		if scanner.Text() == "" {
			indentedText.WriteString(scanner.Text())
			continue
		}
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
				if n := countNewlines(leadingWhitespace); n > 1 {
					leading = strings.Repeat("\n", 2)
				} else if countNewlines(leadingWhitespace) == 1 {
					leading = "\n"
				} else if len(leadingWhitespace) > 0 && position == PositionAfter {
					leading = " "
				}
				// Strip leading whitespace so the comment can be indented to the right location
				containsNewline := strings.Contains(text, "\n")
				text = strings.TrimSpace(text)
				lineComment := strings.HasPrefix(text, "//")

				if n := countNewlines(trailingWhitespace); n > 0 {
					trailing = strings.Repeat("\n", n)
				} else if len(trailingWhitespace) > 0 && !lineComment {
					trailing = " "
				}

				text = fmt.Sprintf("%s%s%s", leading, text, trailing)
				log.Trace(fmt.Sprintf("NORMALIZED COMMENT: %q\n", text))
				if containsNewline {
					text = "\uFFFA" + text + "\uFFFB" + "\n"
				} else if lineComment {
					text = "\uFFF9" + text + "\n\uFFFB"
				} else {
					text = "\uFFF9" + text + "\uFFFB"
				}
				log.Trace(fmt.Sprintf("WRAPPED COMMENT: %q\n\n", text))
			} else if token.GetChannel() == WHITESPACE_CHANNEL && countNewlines(text) > 1 {
				text = "\n" // Replace multiple blank lines with a single blank line
			} else {
				// whitespace to ignore
				continue
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

func getHiddenTokensBetween(tokens *antlr.CommonTokenStream, start, stop antlr.Token) ([]antlr.Token, []antlr.Token) {
	if start == nil || stop == nil || len(tokens.GetAllTokens()) == 0 {
		return nil, nil
	}
	after := tokens.GetHiddenTokensToRight(start.GetTokenIndex(), WHITESPACE_CHANNEL)
	before := tokens.GetHiddenTokensToLeft(stop.GetTokenIndex(), WHITESPACE_CHANNEL)
	inAfter := make(map[int]struct{})
	for _, t := range after {
		inAfter[t.GetTokenIndex()] = struct{}{}
	}
	whitespaceTokens := []antlr.Token{}
	for _, t := range before {
		if _, exists := inAfter[t.GetTokenIndex()]; exists {
			whitespaceTokens = append(whitespaceTokens, t)
		}
	}

	after = tokens.GetHiddenTokensToRight(start.GetTokenIndex(), COMMENTS_CHANNEL)
	before = tokens.GetHiddenTokensToLeft(stop.GetTokenIndex(), COMMENTS_CHANNEL)
	inAfter = make(map[int]struct{})
	for _, t := range after {
		inAfter[t.GetTokenIndex()] = struct{}{}
	}
	commentTokens := []antlr.Token{}
	for _, t := range before {
		if _, exists := inAfter[t.GetTokenIndex()]; exists {
			commentTokens = append(commentTokens, t)
		}
	}
	return whitespaceTokens, commentTokens
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
