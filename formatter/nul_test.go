package formatter

import (
	"strings"
	"testing"
)

// Some editors introduce spurious NUL (\x00) bytes into source files.
// Salesforce ignores them; apexfmt must lex them as whitespace rather than
// failing with a token recognition error, and must not emit them in the
// formatted output.
func TestNulBytesAreIgnored(t *testing.T) {
	clean := "@IsTest\n" +
		"private class NulSample {\n" +
		"\t@IsTest\n" +
		"\tstatic void t() {\n" +
		"\t\tSystem.assert(true);\n" +
		"\t}\n" +
		"}\n"

	tests := []struct {
		name  string
		input string
	}{
		{"trailing", clean + "\x00\x00\x00"},
		{"leading", "\x00" + clean},
		{"between_tokens", strings.Replace(clean, "class NulSample", "class\x00NulSample", 1)},
	}

	// Baseline: the clean source formats successfully; every NUL variant must
	// produce the identical output.
	want, err := NewFormatter("", strings.NewReader(clean)).Formatted()
	if err != nil {
		t.Fatalf("formatting clean source failed: %v", err)
	}
	if strings.Contains(want, "\x00") {
		t.Fatalf("clean formatted output unexpectedly contains a NUL byte")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFormatter("", strings.NewReader(tt.input)).Formatted()
			if err != nil {
				t.Fatalf("formatting NUL-containing source failed: %v", err)
			}
			if strings.Contains(got, "\x00") {
				t.Errorf("formatted output still contains a NUL byte: %q", got)
			}
			if got != want {
				t.Errorf("NUL bytes changed formatting.\n want: %q\n  got: %q", want, got)
			}
		})
	}
}
