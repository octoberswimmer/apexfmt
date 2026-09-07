package formatter

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

const visitorCacheSource = `public class Test {
	// leading comment

	/* block
	   comment */
	public String run(List<Account> accounts, Integer n) { // trailing
		String s = accounts.get(0).Name + accounts.get(1).Name; /* inline */
		if (n > 1 && s.length() > 40 || accounts.isEmpty()) {
			return [SELECT Id FROM Account WHERE Name = :s LIMIT 1].Id;
		}

		return n > 0 ? s : 'é' + s;
	}
}
`

func parseForVisitor(t *testing.T, src string) (*FormatVisitor, antlr.ParseTree) {
	t.Helper()
	stream := tokenStream(src)
	p := parser.NewApexParser(stream)
	p.RemoveErrorListeners()
	tree := p.CompilationUnit()
	stream.Fill()
	return NewFormatVisitor(stream), tree
}

func walk(node antlr.Tree, visit func(antlr.Tree)) {
	visit(node)
	for _, child := range node.GetChildren() {
		walk(child, visit)
	}
}

func Test_text_len_matches_the_length_of_get_text(t *testing.T) {
	v, tree := parseForVisitor(t, visitorCacheSource)
	nodes := 0
	walk(tree, func(node antlr.Tree) {
		nodes++
		want := len(node.(antlr.ParseTree).GetText())
		if got := v.textLen(node); got != want {
			t.Errorf("%T: got %d, want %d", node, got, want)
		}
		// A second call answers from the cache.
		if got := v.textLen(node); got != want {
			t.Errorf("%T cached: got %d, want %d", node, got, want)
		}
	})
	if nodes < 50 {
		t.Errorf("walked only %d nodes", nodes)
	}
}

func Test_hidden_tokens_match_the_token_stream_lookup(t *testing.T) {
	v, _ := parseForVisitor(t, visitorCacheSource)
	stream := v.tokens
	for _, token := range stream.GetAllTokens() {
		i := token.GetTokenIndex()
		wantBefore := interleaveHiddenTokens(
			stream.GetHiddenTokensToLeft(i, WHITESPACE_CHANNEL),
			stream.GetHiddenTokensToLeft(i, COMMENTS_CHANNEL),
		)
		wantAfter := interleaveHiddenTokens(
			stream.GetHiddenTokensToRight(i, WHITESPACE_CHANNEL),
			stream.GetHiddenTokensToRight(i, COMMENTS_CHANNEL),
		)
		compareTokens(t, "before", i, v.hiddenTokens(token, HiddenTokenDirectionBefore), wantBefore)
		compareTokens(t, "after", i, v.hiddenTokens(token, HiddenTokenDirectionAfter), wantAfter)
	}
}

func compareTokens(t *testing.T, direction string, index int, got, want []antlr.Token) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("token %d %s: got %d tokens, want %d", index, direction, len(got), len(want))
		return
	}
	for j := range want {
		if got[j].GetTokenIndex() != want[j].GetTokenIndex() {
			t.Errorf("token %d %s: position %d got token %d, want %d", index, direction, j, got[j].GetTokenIndex(), want[j].GetTokenIndex())
		}
	}
}
