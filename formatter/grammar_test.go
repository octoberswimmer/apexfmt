package formatter

import (
	"strings"
	"testing"
)

// formatMethodBody formats body as the statements of a method and returns
// the statements as formatted, without the enclosing class and method.
func formatMethodBody(t *testing.T, body string) string {
	t.Helper()
	src := "public class Test {\n\tvoid run(Long x, Account a, Account b) {\n\t\t" + body + "\n\t}\n}\n"
	f := NewFormatter("", strings.NewReader(src))
	got, err := f.Formatted()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	prefix := "public class Test {\n\tvoid run(Long x, Account a, Account b) {\n\t\t"
	suffix := "\n\t}\n}\n"
	if !strings.HasPrefix(got, prefix) || !strings.HasSuffix(got, suffix) {
		t.Fatalf("unexpected method wrapper in %q", got)
	}
	out := strings.TrimSuffix(strings.TrimPrefix(got, prefix), suffix)
	return strings.ReplaceAll(out, "\n\t\t", "\n")
}

// assertFormatsStably formats src, compares the result with want, then
// formats the result again and checks that it does not change.
func assertFormatsStably(t *testing.T, src, want string) {
	t.Helper()
	got := formatMethodBody(t, src)
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
	if again := formatMethodBody(t, got); again != got {
		t.Errorf("not idempotent: got %q, then %q", got, again)
	}
}

func Test_formatter_formats_signed_numeric_when_literals(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "negative long",
			src:  "switch on x { when -5L { } }",
			want: "switch on x {\n\twhen -5L {}\n}",
		},
		{
			name: "positive integer",
			src:  "switch on x { when +5 { } }",
			want: "switch on x {\n\twhen +5 {}\n}",
		},
		{
			name: "positive long",
			src:  "switch on x { when + 5L { } }",
			want: "switch on x {\n\twhen +5L {}\n}",
		},
		{
			name: "negative integer with space after the sign",
			src:  "switch on x { when - 5 { } }",
			want: "switch on x {\n\twhen -5 {}\n}",
		},
		{
			name: "repeated minus signs stay separated",
			src:  "switch on x { when - -5 { } }",
			want: "switch on x {\n\twhen - -5 {}\n}",
		},
		{
			name: "repeated plus signs stay separated",
			src:  "switch on x { when + +5L { } }",
			want: "switch on x {\n\twhen + +5L {}\n}",
		},
		{
			name: "mixed signs",
			src:  "switch on x { when +-5, -+6L { } }",
			want: "switch on x {\n\twhen +-5, -+6L {}\n}",
		},
		{
			name: "parenthesized signed long",
			src:  "switch on x { when (-5L) { } }",
			want: "switch on x {\n\twhen (-5L) {}\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFormatsStably(t, tt.src, tt.want)
		})
	}
}

func Test_formatter_formats_merge_with_access_level(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "user mode",
			src:  "merge as user a b;",
			want: "merge as user a b;",
		},
		{
			name: "system mode",
			src:  "merge as system a b;",
			want: "merge as system a b;",
		},
		{
			name: "upper case keywords",
			src:  "MERGE AS USER a b;",
			want: "merge as user a b;",
		},
		{
			name: "method call operands",
			src:  "merge as system a.clone() b;",
			want: "merge as system a.clone() b;",
		},
		{
			name: "variables named as and user",
			src:  "merge as user;",
			want: "merge as user;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFormatsStably(t, tt.src, tt.want)
		})
	}
}

func Test_sll_prediction_accepts_merge_with_access_level_and_signed_when_literals(t *testing.T) {
	src := "public class Test {\n\tvoid run(Long x, Account a, Account b) {\n\t\tmerge as user a b;\n\t\tswitch on x {\n\t\t\twhen -5L, +5 {}\n\t\t}\n\t}\n}\n"
	if _, ok := parseSLL(testEngine(t), tokenStream(src)); !ok {
		t.Errorf("expected SLL prediction to parse %q", src)
	}
}

func Test_formatter_formats_functions_in_sosl_field_lists(t *testing.T) {
	tests := []struct {
		name   string
		fields string
		want   string
	}{
		{
			name:   "toLabel",
			fields: "Id, toLabel(Name)",
			want:   "Id,\n\tTOLABEL(Name)",
		},
		{
			name:   "convertCurrency",
			fields: "Name, convertCurrency(Amount)",
			want:   "Name,\n\tCONVERTCURRENCY(Amount)",
		},
		{
			name:   "FORMAT of convertCurrency",
			fields: "AnnualRevenue, FORMAT(convertCurrency(AnnualRevenue))",
			want:   "AnnualRevenue,\n\tFORMAT(CONVERTCURRENCY(AnnualRevenue))",
		},
		{
			name:   "FORMAT of an aggregate function",
			fields: "AnnualRevenue, FORMAT(MIN(CloseDate))",
			want:   "AnnualRevenue,\n\tFORMAT(MIN(CloseDate))",
		},
		{
			name:   "FORMAT of a field",
			fields: "FORMAT(LastModifiedDate)",
			want:   "FORMAT(LastModifiedDate)",
		},
		{
			name:   "function followed by more fields",
			fields: "toLabel(Name), Id, convertCurrency(Amount), Phone",
			want:   "TOLABEL(Name),\n\tId,\n\tCONVERTCURRENCY(Amount),\n\tPhone",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := "List<List<SObject>> r = [FIND 'Acme' RETURNING Account(" + tt.fields + ")];"
			want := "List<List<SObject>> r = [FIND 'Acme'\nRETURNING Account(\n\t" + tt.want + ")];"
			assertFormatsStably(t, src, want)
		})
	}
}

func Test_formatter_rejects_an_alias_after_a_function_in_a_sosl_field_list(t *testing.T) {
	// Salesforce rejects these with "Expecting ')' but was: '<alias>'".
	tests := []string{
		"TOLABEL(Industry) industryLabel, Name",
		"CONVERTCURRENCY(Amount) convertedAmount, Name",
		"FORMAT(CloseDate) closeText, Name",
	}
	for _, fields := range tests {
		t.Run(fields, func(t *testing.T) {
			src := "public class Test {\n\tvoid run() {\n\t\tList<List<SObject>> r = [FIND 'Acme' RETURNING Opportunity(" + fields + ")];\n\t}\n}\n"
			f := NewFormatter("", strings.NewReader(src))
			if got, err := f.Formatted(); err == nil {
				t.Errorf("expected a syntax error, got %q", got)
			}
		})
	}
}
