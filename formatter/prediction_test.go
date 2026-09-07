package formatter

import (
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
)

func tokenStream(src string) *antlr.CommonTokenStream {
	lexer := parser.NewApexLexer(antlr.NewInputStream(src))
	return antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
}

// SLL prediction resolves the conflict between a method call and an
// identifier followed by another expression to the methodCall alternative,
// which is listed first in the expression rule.
func Test_sll_prediction_accepts_unqualified_method_calls(t *testing.T) {
	tests := []struct {
		name string
		src  string
	}{
		{
			name: "return with method call",
			src:  "public class Test {\n\tpublic Object run(String s) {\n\t\treturn calc(s);\n\t}\n}\n",
		},
		{
			name: "method call statement",
			src:  "public class Test {\n\tpublic void run(String s) {\n\t\tcalc(s);\n\t}\n}\n",
		},
		{
			name: "method call initializer",
			src:  "public class Test {\n\tpublic void run(String s) {\n\t\tObject o = calc(s);\n\t}\n}\n",
		},
		{
			name: "super constructor call",
			src:  "public class Test extends Base {\n\tpublic Test(String s) {\n\t\tsuper(s);\n\t}\n}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, ok := parseSLL(tokenStream(tt.src)); !ok {
				t.Errorf("expected SLL prediction to parse %q", tt.src)
			}
		})
	}
}

// A merge target followed by a parenthesized expression is the one place
// where the methodCall alternative is the wrong choice.
func Test_sll_prediction_reports_failure_on_merge_with_parenthesized_duplicate(t *testing.T) {
	src := "public class Test {\n\tvoid run(Account a, Account b) {\n\t\tmerge a (b);\n\t}\n}\n"
	if _, ok := parseSLL(tokenStream(src)); ok {
		t.Errorf("expected SLL prediction to fail on %q", src)
	}
}

func Test_formatter_falls_back_to_ll_prediction_when_sll_fails(t *testing.T) {
	src := "public class Test { void run(Account a, Account b) { merge a (b); } }"
	want := "public class Test {\n\tvoid run(Account a, Account b) {\n\t\tmerge a (b);\n\t}\n}\n"
	f := NewFormatter("", strings.NewReader(src))
	got, err := f.Formatted()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func Test_formatter_accepts_method_calls_as_merge_operands(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "unqualified method call target",
			src:  "public class Test { void run(Account a, Account b) { merge a(b) b; } }",
			want: "public class Test {\n\tvoid run(Account a, Account b) {\n\t\tmerge a(b) b;\n\t}\n}\n",
		},
		{
			name: "qualified method call operands",
			src:  "public class Test { void run(List<Account> a) { merge a.get(0) a.get(1); } }",
			want: "public class Test {\n\tvoid run(List<Account> a) {\n\t\tmerge a.get(0) a.get(1);\n\t}\n}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter("", strings.NewReader(tt.src))
			got, err := f.Formatted()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_formatter_reports_syntax_errors_once_after_ll_fallback(t *testing.T) {
	// The missing semicolon makes both passes fail; only the LL pass
	// reports errors.
	src := "public class Test { void run(Account a, Account b) { merge a (b) } }"
	f := NewFormatter("", strings.NewReader(src))
	err := f.Format()
	if err == nil {
		t.Fatal("expected a syntax error")
	}
	if n := strings.Count(err.Error(), "line "); n != 1 {
		t.Errorf("expected one reported error, got %d: %q", n, err.Error())
	}
}
