package parser

import (
	"sync"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func parseWith(t *testing.T, p *ApexParser) string {
	t.Helper()
	p.RemoveErrorListeners()
	tree := p.CompilationUnit()
	if p.HasError() {
		t.Fatal("parse failed")
	}
	return tree.ToStringTree(nil, p)
}

func Test_engine_parser_produces_the_same_tree_as_the_shared_parser(t *testing.T) {
	src := "public class Test {\n\tpublic String run(String s) {\n\t\treturn [SELECT Id FROM Account WHERE Name = :s LIMIT 1].Id + s;\n\t}\n}\n"
	want := parseWith(t, NewApexParser(antlr.NewCommonTokenStream(NewApexLexer(antlr.NewInputStream(src)), antlr.TokenDefaultChannel)))
	engine, release := AcquireParserEngine()
	defer release()
	got := parseWith(t, engine.NewApexParser(antlr.NewCommonTokenStream(NewApexLexer(antlr.NewInputStream(src)), antlr.TokenDefaultChannel)))
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func Test_sequential_acquires_reuse_the_first_engine(t *testing.T) {
	first, release := AcquireParserEngine()
	release()
	for i := 0; i < 3; i++ {
		engine, release := AcquireParserEngine()
		if engine != first {
			t.Errorf("acquire %d returned a different engine", i)
		}
		release()
	}
}

func Test_concurrent_acquires_spread_over_at_most_max_engines(t *testing.T) {
	var releases []func()
	seen := map[*ParserEngine]int{}
	for i := 0; i < 3*maxParserEngines; i++ {
		engine, release := AcquireParserEngine()
		seen[engine]++
		releases = append(releases, release)
	}
	if len(seen) != maxParserEngines {
		t.Errorf("got %d engines, want %d", len(seen), maxParserEngines)
	}
	for _, n := range seen {
		if n != 3 {
			t.Errorf("engine used by %d parsers, want 3", n)
		}
	}
	for _, release := range releases {
		release()
	}
	if engine, release := AcquireParserEngine(); engine.users != 1 {
		t.Errorf("engine has %d users after release, want 1", engine.users)
		release()
	} else {
		release()
	}
}

func Test_engines_parse_concurrently(t *testing.T) {
	src := "public class Test {\n\tpublic Integer run(Integer n) {\n\t\treturn n + 1;\n\t}\n}\n"
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				engine, release := AcquireParserEngine()
				p := engine.NewApexParser(antlr.NewCommonTokenStream(NewApexLexer(antlr.NewInputStream(src)), antlr.TokenDefaultChannel))
				p.RemoveErrorListeners()
				p.CompilationUnit()
				if p.HasError() {
					t.Error("parse failed")
				}
				release()
			}
		}()
	}
	wg.Wait()
}
