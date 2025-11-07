package formatter

import (
	"strings"
	"testing"
)

func TestFormatterFormatValidSource(t *testing.T) {
	input := "public class Test { void run() { Integer i = 0; } }"
	f := NewFormatter("", strings.NewReader(input))

	if err := f.Format(); err != nil {
		t.Fatalf("Format() returned error for valid source: %v", err)
	}

	got := string(f.formatted)
	want := "public class Test {\n\tvoid run() {\n\t\tInteger i = 0;\n\t}\n}\n"

	if got != want {
		t.Fatalf("formatted output mismatch\nwant:\n%s\ngot:\n%s", want, got)
	}
}

func TestFormatterFormatInvalidSource(t *testing.T) {
	input := "public class Test { void run() { Integer i = 0 } }" // missing semicolon
	f := NewFormatter("", strings.NewReader(input))

	if err := f.Format(); err == nil {
		t.Fatalf("Format() should return error for invalid source")
	}

	if f.formatted != nil {
		t.Fatalf("expected formatted output to be nil when formatting fails")
	}
}
