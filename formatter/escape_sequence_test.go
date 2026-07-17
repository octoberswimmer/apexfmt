package formatter

import (
	"strings"
	"testing"
)

// Apex accepts the escape letter in either case: '\F' is a form feed just
// like '\f', and the unicode marker '\U0041' works like 'A'. apexfmt
// must lex both cases without a token recognition error and must preserve the
// escape sequence verbatim — a formatter must never change the runtime value
// of a string literal by decoding or re-casing an escape.
func TestFormatter_EscapeSequencesInStringLiteral(t *testing.T) {
	escapes := []string{
		`\b`, `\t`, `\n`, `\f`, `\r`, `\s`,
		`\B`, `\T`, `\N`, `\F`, `\R`, `\S`,
		`\'`, `\"`, `\\`,
		`A`, `\U0041`, `J`, `\U004A`,
	}

	for _, esc := range escapes {
		t.Run(esc, func(t *testing.T) {
			literal := "'x" + esc + "y'"
			input := "public class Foo {\n" +
				"\tpublic static String s() {\n" +
				"\t\treturn " + literal + ";\n" +
				"\t}\n" +
				"}\n"

			got, err := NewFormatter("", strings.NewReader(input)).Formatted()
			if err != nil {
				t.Fatalf("formatting %q failed: %v", esc, err)
			}
			if !strings.Contains(got, literal) {
				t.Fatalf("escape %q not preserved verbatim\ninput literal: %q\noutput:\n%s", esc, literal, got)
			}
		})
	}
}

// The same escapes must survive inside a triple-quoted text block, whose
// interior is preserved verbatim.
func TestFormatter_EscapeSequencesInTextBlock(t *testing.T) {
	escapes := []string{
		`\F`, `\N`, `\S`, `\U0041`,
		`\f`, `\n`, `\s`, `A`,
	}

	for _, esc := range escapes {
		t.Run(esc, func(t *testing.T) {
			input := "public class Foo {\n" +
				"\tpublic static String s() {\n" +
				"\t\treturn '''\n" +
				"a" + esc + "b''';\n" +
				"\t}\n" +
				"}\n"

			got, err := NewFormatter("", strings.NewReader(input)).Formatted()
			if err != nil {
				t.Fatalf("formatting text block with %q failed: %v", esc, err)
			}
			if !strings.Contains(got, "a"+esc+"b") {
				t.Fatalf("text block escape %q not preserved verbatim\noutput:\n%s", esc, got)
			}
		})
	}
}

// Re-running the formatter on its own output must be a no-op for escaped
// literals, guarding against a second pass decoding or re-casing an escape.
func TestFormatter_EscapeSequencesAreIdempotent(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String s() {\n" +
		"\t\treturn 'a\\Fb\\Nc\\U0041d';\n" +
		"\t}\n" +
		"}\n"

	first, err := NewFormatter("", strings.NewReader(input)).Formatted()
	if err != nil {
		t.Fatalf("first Format() failed: %v", err)
	}
	second, err := NewFormatter("", strings.NewReader(first)).Formatted()
	if err != nil {
		t.Fatalf("second Format() failed: %v", err)
	}
	if first != second {
		t.Fatalf("formatter not idempotent\nfirst:\n%q\nsecond:\n%q", first, second)
	}
}
