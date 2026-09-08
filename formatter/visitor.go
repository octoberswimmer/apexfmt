package formatter

import (
	"bytes"
	"fmt"
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
	// hiddenBefore[i] and hiddenAfter[i] hold the whitespace and comment
	// tokens adjacent to token i, up to the nearest token on the default
	// channel. They are built once from the token stream on first use.
	hiddenBefore [][]antlr.Token
	hiddenAfter  [][]antlr.Token
	// textLens caches the length of each rule node's text, which the
	// wrapping decisions consult for nested expressions.
	textLens map[antlr.RuleNode]int
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
		textLens:       make(map[antlr.RuleNode]int),
	}
}

// textLen returns len(node.GetText()) without building the text: the length
// of a rule node is the sum of its children's lengths, and each rule node's
// length is cached, so nested expressions that check their length at every
// level do not rebuild the same text repeatedly.
func (v *FormatVisitor) textLen(node antlr.Tree) int {
	switch n := node.(type) {
	case antlr.TerminalNode:
		return len(n.GetText())
	case antlr.RuleNode:
		if l, ok := v.textLens[n]; ok {
			return l
		}
		l := 0
		for _, child := range n.GetChildren() {
			l += v.textLen(child)
		}
		v.textLens[n] = l
		return l
	}
	return len(node.(antlr.ParseTree).GetText())
}

// indexHiddenTokens records, for every token index, the whitespace and
// comment tokens between it and the nearest default-channel token on each
// side, in stream order. This is what interleaving the results of
// GetHiddenTokensToLeft and GetHiddenTokensToRight for the whitespace and
// comment channels produces, computed once instead of four times per node.
func (v *FormatVisitor) indexHiddenTokens() {
	all := v.tokens.GetAllTokens()
	v.hiddenBefore = make([][]antlr.Token, len(all))
	v.hiddenAfter = make([][]antlr.Token, len(all))
	hidden := func(from, to int) []antlr.Token {
		var run []antlr.Token
		for i := from; i <= to; i++ {
			if c := all[i].GetChannel(); c == WHITESPACE_CHANNEL || c == COMMENTS_CHANNEL {
				run = append(run, all[i])
			}
		}
		return run
	}
	// runStart is the index after the last default-channel token seen.
	runStart := 0
	for i := range all {
		if i > runStart {
			v.hiddenBefore[i] = hidden(runStart, i-1)
		}
		if all[i].GetChannel() == antlr.TokenDefaultChannel {
			runStart = i + 1
		}
	}
	// runEnd is the index before the next default-channel token seen.
	runEnd := len(all) - 1
	for i := len(all) - 1; i >= 0; i-- {
		if i < runEnd {
			v.hiddenAfter[i] = hidden(i+1, runEnd)
		}
		if all[i].GetChannel() == antlr.TokenDefaultChannel {
			runEnd = i - 1
		}
	}
}

// hiddenTokens returns the whitespace and comment tokens adjacent to token
// in the given direction, in stream order.
func (v *FormatVisitor) hiddenTokens(token antlr.Token, direction HiddenTokenDirection) []antlr.Token {
	if token == nil {
		return nil
	}
	if v.hiddenBefore == nil {
		v.indexHiddenTokens()
	}
	index := token.GetTokenIndex()
	if index < 0 || index >= len(v.hiddenBefore) {
		return nil
	}
	if direction == HiddenTokenDirectionBefore {
		return v.hiddenBefore[index]
	}
	return v.hiddenAfter[index]
}

func (v *FormatVisitor) VisitRule(node antlr.RuleNode) interface{} {
	return v.visitRule(node)
}

func (v *FormatVisitor) visitRule(node antlr.RuleNode) interface{} {
	start, stop := getStartStop(node)

	// Collect comments and whitespace before the node
	beforeHiddenTokens := v.hiddenTokens(start, HiddenTokenDirectionBefore)

	result := node.Accept(v)
	if result == nil {
		panic(fmt.Sprintf("MISSING VISIT FUNCTION FOR %T", node))
	}

	result = appendHiddenTokens(v, result, beforeHiddenTokens, PositionBefore)

	// Check for empty block before adding after tokens
	handledEmptyBlock := false
	if result.(string) == "{}" {
		inbetweenTokens := interleaveHiddenTokens(
			getHiddenTokensBetween(v.tokens, start, stop),
		)
		tokenText := appendHiddenTokens(v, "", inbetweenTokens, PositionAfter).(string)
		// Strip trailing whitespace, accounting for comment markers
		tokenText = strings.TrimRight(tokenText, " \t\n")
		if strings.HasSuffix(tokenText, "\uFFFB") {
			tokenText = strings.TrimRight(tokenText[:len(tokenText)-len("\uFFFB")], " \t\n") + "\uFFFB"
		}
		if strings.TrimSpace(tokenText) != "" {
			result = "{\n" + indent(tokenText) + "\n}"
			handledEmptyBlock = true
		}
	}

	// Collect comments and whitespace after the node
	afterHiddenTokens := v.hiddenTokens(stop, HiddenTokenDirectionAfter)

	_ = afterHiddenTokens
	if !handledEmptyBlock {
		result = appendHiddenTokens(v, result, afterHiddenTokens, PositionAfter)
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
	traceEnabled := log.IsLevelEnabled(log.TraceLevel)
	if traceEnabled {
		log.Tracef("SPLITTING: %q", string(data))
	}
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
		if traceEnabled {
			log.Tracef("NO NEWLINE IN: %q", string(data))
		}
		if !atEOF {
			// No newline, not at EOF => request more data
			if traceEnabled {
				log.Tracef("REQUESTING MORE DATA: %q", string(data))
			}
			return 0, nil, nil
		}

		if traceEnabled {
			log.Tracef("AT EOF IN: %q", string(data))
		}
		line := data

		if hasInlineComment(line, inlineCommentStart, fffb) {
			if traceEnabled {
				log.Tracef("HAS INLINE COMMENT: %q", string(line))
			}
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
				if traceEnabled {
					log.Tracef("HAS ONLY DELIMITER: %q", string(line))
				}
				// Line contains only the delimiter
				return len(data), line, nil
			}

			// Line starts with the delimiter and has additional content
			delimiterIdx := bytes.Index(line, delimiter)
			if delimiterIdx != -1 {
				// Include leading whitespace in delimiter token
				delimiterEnd := delimiterIdx + len(delimiter)
				if traceEnabled {
					log.Tracef("DELIMITER+: %q", string(line[:delimiterEnd]))
				}
				return delimiterEnd, line[:delimiterEnd], nil
			}
		}

		// Otherwise, no delimiters => return the entire line
		if traceEnabled {
			log.Tracef("NO DELIMITERS: %q", string(data))
		}
		return len(data), data, nil
	}
	if traceEnabled {
		log.Tracef("FOUND NEWLINE IN: %q", string(data))
	}

	// ----------------------------------------------------------------
	// 2. WE FOUND A NEWLINE => Extract the line
	// ----------------------------------------------------------------
	line := data[:newlineIdx]

	// --- Inline comment check ---
	// If line has \ufff9 and \ufffb in the correct order, keep it as one token.
	if hasInlineComment(line, inlineCommentStart, fffb) {
		if traceEnabled {
			log.Tracef("INLINE COMMENT: %q", string(data))
		}
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
			if traceEnabled {
				log.Tracef("HAS DELIMITER ONLY: %q", string(line))
			}
			// Line contains only the delimiter
			return newlineIdx + 1, line, nil
		}

		// Line starts with delimiter but has more content
		delimiterIdx := bytes.Index(line, delimiter)
		if delimiterIdx != -1 {
			delimiterEnd := delimiterIdx + delimiterLen
			if traceEnabled {
				log.Tracef("\\uFFFB+: %q", string(line[:delimiterEnd]))
			}
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
			if traceEnabled {
				log.Tracef("HAS DELIMITER ONLY: %q", string(line))
			}
			// Line contains only the delimiter
			return newlineIdx + 1, line, nil
		}

		if fffbIdx != -1 {
			// \ufffb is before the newline
			delimiterEnd := fffbIdx + len("\ufffb")
			if traceEnabled {
				log.Tracef("RETURNING UP TO \\uFFFB: %q", string(line[:delimiterEnd]))
			}
			// Advance past the newline after \uFFFB
			return delimiterEnd + 1, line[:delimiterEnd], nil
		}
		if f := bytes.Index(data, fffb); f == newlineIdx+1 {
			delimiterLen := len(fffb)
			if traceEnabled {
				log.Tracef("RETURNING UP TO NEWLINE WITH \\uFFFB: %q", string(data[:f+delimiterLen]))
			}
			return f + delimiterLen, data[:f+delimiterLen], nil
		}

		// Line starts with delimiter but has more content
		delimiterIdx := bytes.Index(line, delimiter)
		if delimiterIdx != -1 {
			if traceEnabled {
				log.Tracef("\\uFFFA+: %q", string(line))
			}
			return newlineIdx + 1, line, nil
		}
	}

	// ----------------------------------------------------------------
	// Delimiter Elsewhere in the Line
	// ----------------------------------------------------------------

	if fffaIdx != -1 && (fffbIdx == -1 || fffaIdx < fffbIdx) {
		// Split BEFORE the delimiter
		if traceEnabled {
			log.Tracef("HAS \\uFFFA IN LINE: %q", string(line[:fffaIdx]))
		}
		return fffaIdx, line[:fffaIdx], nil
	}
	if fffbIdx != -1 && (fffaIdx == -1 || fffbIdx < fffaIdx) {
		delimiterLen := len(fffb)
		// Split AFTER the delimiter
		if traceEnabled {
			log.Tracef("HAS \\uFFFB IN LINE: %q", string(line[:fffbIdx+delimiterLen]))
		}
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
	if traceEnabled {
		log.Tracef("NO DELIMITER: %q", string(line))
	}
	if len(line) > 0 && fffbFollowsNewlines(data[newlineIdx:]) {
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
	indentedText.Grow(len(text) + indents*(strings.Count(text, "\n")+1))
	tabs := strings.Repeat("\t", indents)
	isFirstLine := true
	// Track whether the scanner is currently inside a triple-quoted text
	// block literal. Content inside must not receive indentation; any
	// whitespace prepended would become part of the string value.
	insideTextBlock := false

	log.Debugf("INDENTING: %q\n", text)

	// Split the whole text at once rather than through a bufio.Scanner. The
	// scanner stops at the first token longer than its 64 KB buffer limit,
	// which silently dropped the rest of the text after a long line. The
	// split function returns each token as a prefix of the data it is
	// given, so the token is also a substring of text at the same offset.
	data := []byte(text)
	for pos := 0; pos < len(data); {
		advance, token, err := SplitLeadingFFFAOrFFFBOrNewline(data[pos:], true)
		if err != nil || advance <= 0 {
			break
		}
		t := text[pos : pos+len(token)]
		pos += advance
		if token == nil {
			continue
		}
		log.Tracef("INDENTING LINE: %q\n", t)
		if t == "\uFFFB" {
			indentedText.WriteString(t)
			continue
		}
		// Decide indentation based on the state entering this line, then
		// update the state for the next line based on any triple-quoted
		// text block transitions that happen on this one. Content inside
		// a text block is preserved verbatim; everything else is indented
		// normally.
		applyIndent := !insideTextBlock
		insideTextBlock = updateTextBlockState(t, insideTextBlock)
		if isFirstLine {
			isFirstLine = false
		} else if !strings.HasPrefix(t, "\uFFFA") && !strings.HasPrefix(t, "\uFFF9") {
			indentedText.WriteString("\n")
		}
		if t == "" {
			continue
		}
		if applyIndent {
			indentedText.WriteString(tabs)
		}
		indentedText.WriteString(t)
	}
	log.Debugf("INDENTED:  %q\n\n", indentedText.String())

	return indentedText.String()
}

// updateTextBlockState walks the line to determine whether, at its end, we
// are still inside a triple-quoted text block. When entering inside a text
// block, it looks only for the first ”' that closes the block. When
// starting outside, it skips over // line comments, /* ... */ block
// comments, and ordinary '...' string literals so those cannot
// spuriously trigger a text block.
func updateTextBlockState(line string, insideTextBlock bool) bool {
	i := 0
	for i < len(line) {
		if insideTextBlock {
			// Walk forward char by char so backslash-escaped delimiters
			// (\''' per JEP 378) do not spuriously close the block.
			for i < len(line) {
				if line[i] == '\\' && i+1 < len(line) {
					i += 2
					continue
				}
				if i+2 < len(line) && line[i] == '\'' && line[i+1] == '\'' && line[i+2] == '\'' {
					i += 3
					insideTextBlock = false
					break
				}
				i++
			}
			if insideTextBlock {
				return insideTextBlock
			}
			continue
		}
		c := line[i]
		switch {
		case c == '/' && i+1 < len(line) && line[i+1] == '/':
			// Line comment — nothing after this matters for the state.
			return insideTextBlock
		case c == '/' && i+1 < len(line) && line[i+1] == '*':
			i += 2
			if end := strings.Index(line[i:], "*/"); end >= 0 {
				i += end + 2
			} else {
				// Unterminated block comment on this line; nothing
				// further can affect the text block state.
				return insideTextBlock
			}
		case c == '\'' && i+2 < len(line) && line[i+1] == '\'' && line[i+2] == '\'':
			// Opening triple-quote delimiter.
			insideTextBlock = true
			i += 3
		case c == '\'':
			// Ordinary single-quoted string literal; skip to the
			// matching closing quote, respecting backslash escapes so
			// we do not terminate on a \' in the middle.
			i++
			for i < len(line) {
				if line[i] == '\\' && i+1 < len(line) {
					i += 2
					continue
				}
				if line[i] == '\'' {
					i++
					break
				}
				i++
			}
		default:
			i++
		}
	}
	return insideTextBlock
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
				} else if len(leadingWhitespace) > 0 {
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

				text = leading + text + trailing
				log.Tracef("NORMALIZED COMMENT: %q\n", text)
				if containsNewline {
					text = "\uFFFA" + text + "\uFFFB" + "\n"
				} else if lineComment {
					text = "\uFFF9" + text + "\n\uFFFB"
				} else {
					text = "\uFFF9" + text + "\uFFFB"
				}
				log.Tracef("WRAPPED COMMENT: %q\n\n", text)
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
			result = tokenText + result.(string)
		case PositionAfter:
			result = result.(string) + tokenText
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

// fffbFollowsNewlines reports whether data starts with one or more newlines
// followed immediately by \uFFFB.
func fffbFollowsNewlines(data []byte) bool {
	i := 0
	for i < len(data) && data[i] == '\n' {
		i++
	}
	return i > 0 && bytes.HasPrefix(data[i:], []byte("\uFFFB"))
}
