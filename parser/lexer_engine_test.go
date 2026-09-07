package parser

import (
	"sync"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

const lexerEngineSource = "public class Test {\n\t// comment\n\tpublic String run(String s) {\n\t\treturn [SELECT Id FROM Account WHERE Name = :s LIMIT 1].Id + 'x';\n\t}\n}\n"

func tokenTexts(t *testing.T, lexer *ApexLexer) []string {
	t.Helper()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()
	var texts []string
	for _, tok := range stream.GetAllTokens() {
		texts = append(texts, tok.GetText())
	}
	return texts
}

func Test_private_atn_lexer_produces_the_same_tokens_as_the_shared_lexer(t *testing.T) {
	want := tokenTexts(t, NewApexLexer(antlr.NewInputStream(lexerEngineSource)))
	lexer, release := NewApexLexerWithPrivateATN(antlr.NewInputStream(lexerEngineSource))
	defer release()
	got := tokenTexts(t, lexer)
	if len(got) != len(want) {
		t.Fatalf("got %d tokens, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("token %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func Test_private_atn_lexers_run_concurrently(t *testing.T) {
	want := tokenTexts(t, NewApexLexer(antlr.NewInputStream(lexerEngineSource)))
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				lexer, release := NewApexLexerWithPrivateATN(antlr.NewInputStream(lexerEngineSource))
				got := tokenTexts(t, lexer)
				release()
				if len(got) != len(want) {
					t.Errorf("got %d tokens, want %d", len(got), len(want))
				}
			}
		}()
	}
	wg.Wait()
}
