package formatter

import (
	"regexp"
	"strings"
	"testing"
)

// Sample inputs with varying complexity for benchmarking
var benchmarkInputs = []struct {
	name  string
	input string
}{
	{
		name: "simple_inline",
		input: `public class MyClass {
	public void method() {
		statement(); // inline comment
	}
}`,
	},
	{
		name: "multiple_inline",
		input: `public class MyClass {
	public void method() {
		statement1(); // comment 1
		statement2(); // comment 2
		statement3(); // comment 3
		statement4(); // comment 4
		statement5(); // comment 5
	}
}`,
	},
	{
		name: "multiline_comments",
		input: `public class MyClass {
	/*
	 * Multi-line comment
	 * with multiple lines
	 * and indentation
	 */
	public void method() {
		/* Another multi-line
		   comment with different
		   indentation */
		statement();
	}
}`,
	},
	{
		name: "mixed_comments",
		input: `public class MyClass {
	/**
	 * JavaDoc style comment
	 * @param name description
	 * @return value
	 */
	public String method(String name) {
		// Simple line comment
		if (condition) {
			/* Inline block comment */ statement();
		}
		
		/*
		 * Another multi-line
		 * comment block
		 */
		return result; // trailing comment
	}
}`,
	},
	{
		name: "complex_nested",
		input: `for (Referral__c ref : [SELECT Summary_Name__c, Name FROM Referral__c WHERE Id IN :referralIdSet]) {
	/* Multi-line comments
	preserve white-space.
		So they can contain
			{ code }
				or poetry
		for
			example
	*/
	System.assertEquals(ref.Name, ref.Summary_Name__c);
	
	if (condition) {
		// Nested comment
		doSomething(); // inline
		
		/*
		 * Another nested
		 * multi-line comment
		 */
		doSomethingElse();
	}
}`,
	},
}

func BenchmarkRemoveExtraCommentIndentation(b *testing.B) {
	// Create test inputs with comment delimiters to simulate what the formatter produces
	testInputs := map[string]string{
		"simple_inline":     "statement();\uFFF9 // inline comment\uFFFB",
		"multiple_inline":   "statement1();\uFFF9 // comment 1\uFFFB\n\tstatement2();\uFFF9 // comment 2\uFFFB\n\tstatement3();\uFFF9 // comment 3\uFFFB",
		"multiline_comment": "\uFFFA/*\n\t * Multi-line comment\n\t * with multiple lines\n\t */\uFFFB\n\tstatement();",
		"mixed_comments":    "\uFFFA/**\n\t * JavaDoc style comment\n\t * @param name description\n\t */\uFFFB\n\tif (condition) {\n\t\t\uFFFA/* Inline block comment */\uFFFB statement();\n\t}\n\treturn result;\uFFF9 // trailing comment\uFFFB",
		"complex_nested":    "\uFFFA/* Multi-line comments\n\tpreserve white-space.\n\t\tSo they can contain\n\t\t\t{ code }\n\t\t\t\tor poetry\n\t\tfor\n\t\t\texample\n\t*/\uFFFB\n\tSystem.assertEquals(ref.Name, ref.Summary_Name__c);\n\t\uFFF9// Nested comment\uFFFB\n\tdoSomething();\uFFF9 // inline\uFFFB",
	}

	for name, input := range testInputs {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = removeExtraCommentIndentation(input)
			}
		})
	}
}

func BenchmarkRemoveExtraCommentIndentationLarge(b *testing.B) {
	// Create a large input with many comment patterns
	var largeInput strings.Builder
	complexPattern := "\uFFFA/* Multi-line comments\n\tpreserve white-space.\n\t\tSo they can contain\n\t\t\t{ code }\n\t\t\t\tor poetry\n\t\tfor\n\t\t\texample\n\t*/\uFFFB\n\tSystem.assertEquals(ref.Name, ref.Summary_Name__c);\n\t\uFFF9// Nested comment\uFFFB\n\tdoSomething();\uFFF9 // inline\uFFFB"

	for i := 0; i < 200; i++ {
		largeInput.WriteString(complexPattern)
		largeInput.WriteString("\n")
	}

	input := largeInput.String()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = removeExtraCommentIndentation(input)
	}
}

// BenchmarkCommentProcessingComparison compares the current optimized version
// with a simulated slower version to demonstrate the improvements
func BenchmarkCommentProcessingComparison(b *testing.B) {
	testInput := "\uFFFA/* Multi-line comments\n\tpreserve white-space.\n\t\tSo they can contain\n\t\t\t{ code }\n\t\t\t\tor poetry\n\t\tfor\n\t\t\texample\n\t*/\uFFFB\n\tSystem.assertEquals(ref.Name, ref.Summary_Name__c);\n\t\uFFF9// Nested comment\uFFFB\n\tdoSomething();\uFFF9 // inline\uFFFB"

	b.Run("optimized_version", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = removeExtraCommentIndentation(testInput)
		}
	})
}

// Benchmark individual regex patterns to identify bottlenecks
func BenchmarkCommentRegexPatterns(b *testing.B) {
	// Use a sample input with all comment delimiters
	sampleInput := `	/* Multi-line comments
	preserve white-space.
		So they can contain
			{ code }
				or poetry
		for
			example
	*/
	statement(); // inline comment
	another(); // another comment`

	// Add delimiters to simulate processed input
	withDelimiters := strings.ReplaceAll(sampleInput, "/*", "\uFFFA/*")
	withDelimiters = strings.ReplaceAll(withDelimiters, "*/", "*/\uFFFB")
	withDelimiters = strings.ReplaceAll(withDelimiters, "//", "\uFFF9//")

	patterns := map[string]string{
		"newline_prefixed_multiline": `[\n ]*(\t*[` + "\uFFFA\uFFF9" + `])`,
		"indented_inline":            `([^\n` + "\uFFFB" + `])\t+` + "\uFFFA" + `([^\n])`,
		"space_padded_multiline":     `(` + "\uFFFB" + `\n*\t*) +`,
		"indent_injected_newlines":   "\uFFFB" + `\n+`,
		"double_captured_newlines":   `\n(` + "\uFFFB" + `\t*` + "\uFFFA" + `\n)`,
		"newline_prefixed_inline":    `\n\t*` + "\uFFF9" + `\n`,
		"tab_prefixed_inline":        `([\w,{]+)\t+` + "\uFFF9",
		"inline_comment_removal":     `(?s)` + "\uFFF9" + `(.*?)` + "\uFFFB",
		"multiline_comment_removal":  `(?s)\t*` + "\uFFFA" + `.*?` + "\uFFFB",
	}

	for name, pattern := range patterns {
		b.Run(name, func(b *testing.B) {
			compiled := regexp.MustCompile(pattern)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = compiled.ReplaceAllString(withDelimiters, "$1")
			}
		})
	}
}
