package formatter

import (
	"strings"
	"testing"
)

// TestFormatter_TextBlockPreservesContent ensures apexfmt preserves the
// interior of a triple-quoted text block verbatim. Any stray indentation
// injected inside the block would change the runtime string value, and
// Salesforce rejects a triple-quoted block that has been transformed into
// an ordinary single-quoted literal with embedded newlines.
func TestFormatter_TextBlockPreservesContent(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String s() {\n" +
		"\t\treturn '''\n" +
		"{\n" +
		"\t\"k\": \"v\"\n" +
		"}''';\n" +
		"\t}\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("text block content was altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_TextBlockIsIdempotent double-checks that re-running the
// formatter on its own output is a no-op.
func TestFormatter_TextBlockIsIdempotent(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String s() {\n" +
		"\t\treturn '''\n" +
		"first\n" +
		"\tsecond\n" +
		"third''';\n" +
		"\t}\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("first Format() failed: %v", err)
	}
	first := string(f.formatted)

	f2 := NewFormatter("", strings.NewReader(first))
	if err := f2.Format(); err != nil {
		t.Fatalf("second Format() failed: %v", err)
	}
	second := string(f2.formatted)
	if first != second {
		t.Fatalf("formatter not idempotent\nfirst:\n%q\nsecond:\n%q", first, second)
	}
}

// TestFormatter_TextBlockAlongsideRegularCode ensures indentation of
// surrounding statements is unchanged when a text block is present.
func TestFormatter_TextBlockAlongsideRegularCode(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String s() {\n" +
		"\t\tInteger before = 1;\n" +
		"\t\tString json = '''\n" +
		"{\n" +
		"\t\"x\": 1\n" +
		"}''';\n" +
		"\t\tInteger after = 2;\n" +
		"\t\treturn json;\n" +
		"\t}\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("surrounding indentation altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestUpdateTextBlockState exercises the scanner that decides whether the
// next line begins inside a text block, particularly the edge cases around
// comments and plain single-quoted strings that must not toggle the state.
func TestUpdateTextBlockState(t *testing.T) {
	cases := []struct {
		name        string
		line        string
		startInside bool
		wantInside  bool
	}{
		{"empty line outside", ``, false, false},
		{"plain code outside", `Integer x = 1;`, false, false},
		{"single-quoted string outside", `String s = 'hello';`, false, false},
		{"two empty strings outside", `String s = '' + '';`, false, false},
		{"opening ''' at end of line", `String s = ''';`, false, true},
		{"opening and closing on same line", `String s = '''abc''';`, false, false},
		{"closing ''' on line inside", `}''';`, true, false},
		{"line-comment containing '''", `// Note about ''' in text`, false, false},
		// When we are already inside a text block, everything on the line
		// is string content; a ''' still closes the block regardless of
		// what other characters (including // that would look like a line
		// comment outside) are on the same line.
		{"''' while inside closes block", `// Note about ''' text`, true, false},
		{"block comment spanning no text block", `/* foo ''' bar */`, false, false},
		{"text block inside continuation", `more content`, true, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := updateTextBlockState(tc.line, tc.startInside); got != tc.wantInside {
				t.Errorf("updateTextBlockState(%q, %v) = %v, want %v", tc.line, tc.startInside, got, tc.wantInside)
			}
		})
	}
}

// TestFormatter_TextBlockDelimiterInCommentDoesNotUnindent guards against
// the regression where a // comment mentioning ”' flipped indentTo's
// text-block state and left subsequent lines un-indented.
func TestFormatter_TextBlockDelimiterInCommentDoesNotUnindent(t *testing.T) {
	input := "public class Foo {\n" +
		"\t// The opening ''' must be followed by a line terminator; that\n" +
		"\t// terminator is NOT part of the string value.\n" +
		"\tpublic static Integer x = 1;\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("comments containing ''' should not affect indentation\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_TextBlockPreservesTrailingWhitespace asserts apexfmt does
// not right-trim content lines inside a text block. The Apex compiler
// applies Java text block semantics (stripping per-line trailing
// whitespace) at compile time; the formatter must not silently change
// source bytes that belong to a string literal, and any pre-build tool
// relying on byte-identical round-tripping would otherwise drop data.
func TestFormatter_TextBlockPreservesTrailingWhitespace(t *testing.T) {
	// Three trailing spaces after `hello`, two trailing tabs after `world`.
	input := "public class Foo {\n" +
		"\tpublic static String s = '''\n" +
		"hello   \n" +
		"world\t\t\n" +
		"''';\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("trailing whitespace inside text block was altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_TextBlockPreservesIncidentalLeadingWhitespace asserts
// apexfmt does not alter leading whitespace inside a text block. The
// Apex compiler performs "incidental whitespace" stripping at compile
// time based on the minimum common indent, but the formatter operates
// purely on source bytes and must hand the text block through unchanged.
func TestFormatter_TextBlockPreservesIncidentalLeadingWhitespace(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String s = '''\n" +
		"    indented-with-spaces\n" +
		"\t\tindented-with-tabs\n" +
		"''';\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("leading whitespace inside text block was altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_TextBlockPreservesIndentedClosingDelimiter asserts that
// whitespace that precedes the closing ”' on its own line survives
// formatting. The compiler uses this column to compute the incidental
// whitespace strip; the formatter must leave it alone.
func TestFormatter_TextBlockPreservesIndentedClosingDelimiter(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String s = '''\n" +
		"\t\thello\n" +
		"\t\t''';\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("indented closing delimiter was altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_TextBlockPreservesBlankLines asserts blank lines inside
// a text block — both immediately after the opening ”' and between
// content lines — pass through without being collapsed or indented.
func TestFormatter_TextBlockPreservesBlankLines(t *testing.T) {
	// Two blank lines after opening, three blank lines between content.
	input := "public class Foo {\n" +
		"\tpublic static String s = '''\n" +
		"\n" +
		"\n" +
		"a\n" +
		"\n" +
		"\n" +
		"\n" +
		"b''';\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("blank lines inside text block were altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_EmptyTextBlockRoundTrips asserts that a text block whose
// value is empty (opening ”', LF, closing ”') survives formatting.
func TestFormatter_EmptyTextBlockRoundTrips(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String s = '''\n" +
		"''';\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("empty text block was altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_AdjacentTextBlocksStayIndependent asserts back-to-back
// text blocks — where the first block's closing ”' and the second
// block's opening ”' appear on separate statements — are each handled
// independently without the state leaking between them.
func TestFormatter_AdjacentTextBlocksStayIndependent(t *testing.T) {
	input := "public class Foo {\n" +
		"\tpublic static String a = '''\n" +
		"first\n" +
		"''';\n" +
		"\tpublic static String b = '''\n" +
		"second\n" +
		"''';\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != input {
		t.Fatalf("adjacent text blocks were altered\nwant:\n%q\n got:\n%q", input, got)
	}
}

// TestFormatter_TextBlockInsideIndentedMethodBody is the end-to-end
// regression: a text block nested inside two levels of block indentation
// must come out with the surrounding code re-indented normally but its
// own body untouched.
func TestFormatter_TextBlockInsideIndentedMethodBody(t *testing.T) {
	// Note: the statement before/after the text block is intentionally
	// left-aligned in the input so apexfmt has to add indentation to
	// them while leaving the text block body alone.
	input := "public class Foo {\npublic static String build() {\nInteger n = 1;\nString json = '''\n{\n\t\"n\": 1\n}''';\nreturn json;\n}\n}\n"
	want := "public class Foo {\n" +
		"\tpublic static String build() {\n" +
		"\t\tInteger n = 1;\n" +
		"\t\tString json = '''\n" +
		"{\n" +
		"\t\"n\": 1\n" +
		"}''';\n" +
		"\t\treturn json;\n" +
		"\t}\n" +
		"}\n"
	f := NewFormatter("", strings.NewReader(input))
	if err := f.Format(); err != nil {
		t.Fatalf("Format() failed: %v", err)
	}
	got := string(f.formatted)
	if got != want {
		t.Fatalf("nested text block formatting mismatch\nwant:\n%q\n got:\n%q", want, got)
	}
}
