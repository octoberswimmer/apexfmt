package formatter

import (
	"bufio"
	"bytes"
	"testing"
)

func TestIndent(t *testing.T) {
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				"abc",
				"\tabc",
			},
			{
				"\uFFFA" + `

	/**
	 * Doc string after newline
	 * line 2
	 */` + "\uFFFB",
				"\t\uFFFA" + `

		/**
		 * Doc string after newline
		 * line 2
		 */` + "\uFFFB",
			},
			{
				"public class B {\n\t\ufffa\n\t/**\n\t\t\t */\n\ufffb\tpublic X(Y client) {}\n}",
				"\tpublic class B {\n\t\t\ufffa\n\t\t/**\n\t\t\t\t */\ufffb\n\t\tpublic X(Y client) {}\n\t}",
			},
		}
	for _, tt := range tests {
		out := indent(tt.input)
		if out != tt.output {
			t.Errorf("unexpected indent format.  expected:\n%q\ngot:\n%q\n", tt.output, out)
		}
	}
}

func TestRemoveIndentation(t *testing.T) {
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				"\tabc",
				"\tabc",
			},
			{
				"\t\uFFFA" + `

	/**
		 * Doc string after newline
		 * line 2
		 */` + "\uFFFB",
				`

	/**
	 * Doc string after newline
	 * line 2
	 */`,
			},
			{
				"\t\ufffa\n\n\t/* comment\n\t\t */\n\ufffb",
				"\n\n\t/* comment\n\t */\n",
			},
		}
	for _, tt := range tests {
		out := removeIndentationFromComment(tt.input)
		if out != tt.output {
			t.Errorf("unexpected indent format.  expected:\n%q\ngot:\n%q\n", tt.output, out)
		}
	}
}

func TestSplitLeadingFFFAOrFFFBOrNewline(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Basic class declaration",
			input: "public class B {\n}",
			expected: []string{
				"public class B {",
				"}",
			},
		},
		{
			name:  "Single \ufffa delimiter at start of line",
			input: "public class B {\n\t\ufffa\n}",
			expected: []string{
				"public class B {",
				"\t\ufffa",
				"}",
			},
		},
		{
			name:  "Multiple delimiters with comments",
			input: "public class B {\n\t\ufffa\n\t/**\n\t\t */\n\ufffb\n}",
			expected: []string{
				"public class B {",
				"\t\ufffa",
				"\t/**",
				"\t\t */",
				"\ufffb",
				"}",
			},
		},
		{
			name:  "Delimiter without newline at EOF",
			input: "public class B {\n\t\ufffa",
			expected: []string{
				"public class B {",
				"\t\ufffa",
			},
		},
		{
			name:  "Indented \ufffb delimiter at end of line",
			input: "public class B {\n    \ufffb\n}",
			expected: []string{
				"public class B {",
				"    \ufffb",
				"}",
			},
		},
		{
			name:  "Delimiter with content on the same line",
			input: "public class B {\n\t\ufffa some content\n}",
			expected: []string{
				"public class B {",
				"\t\ufffa",
				" some content",
				"}",
			},
		},
		{
			name:  "Multiple newlines and delimiters",
			input: "public class B {\n\n\t\ufffa\n/**\n\t\t */\n\n\ufffb\n}",
			expected: []string{
				"public class B {",
				"",
				"\t\ufffa",
				"/**",
				"\t\t */",
				"",
				"\ufffb",
				"}",
			},
		},
		{
			name:  "Delimiter preceded by spaces and tab",
			input: "public class B {\n \t\ufffa\n}",
			expected: []string{
				"public class B {",
				" \t\ufffa",
				"}",
			},
		},
		{
			name:  "No delimiters, multiple lines",
			input: "public class B {\n\tpublic X(Y client) {}\n}",
			expected: []string{
				"public class B {",
				"\tpublic X(Y client) {}",
				"}",
			},
		},
		{
			name:  "Delimiter only",
			input: "\ufffa",
			expected: []string{
				"\ufffa",
			},
		},
		{
			name:  "Delimiter with newline immediately after",
			input: "\ufffa\n",
			expected: []string{
				"\ufffa",
			},
		},
		{
			name:  "Multiple delimiters in a row",
			input: "public class B {\n\ufffa\n\ufffb\n}",
			expected: []string{
				"public class B {",
				"\ufffa",
				"\ufffb",
				"}",
			},
		},
		{
			name:  "Delimiters with mixed indentation",
			input: "public class B {\n\t\ufffa\n\t\ufffb\tpublic X(Y client) {}\n}",
			expected: []string{
				"public class B {",
				"\t\ufffa",
				"\t\ufffb",
				"\tpublic X(Y client) {}",
				"}",
			},
		},
		{
			name:  "Delimiter in the middle of the line (should split)",
			input: "public class B {\n\tpublic \ufffa X(Y client) {}\n}",
			expected: []string{
				"public class B {",
				"\tpublic ",
				"\ufffa",
				" X(Y client) {}",
				"}",
			},
		},
		{
			name:  "Delimiter at the end with no newline",
			input: "public class B {\n\ufffa",
			expected: []string{
				"public class B {",
				"\ufffa",
			},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: []string{
				// No tokens expected
			},
		},
		{
			name:  "Only newline",
			input: "\n",
			expected: []string{
				"",
			},
		},
		{
			name:  "Only whitespace",
			input: "   \t  ",
			expected: []string{
				"   \t  ",
			},
		},
		{
			name:  "Delimiter preceded by multiple whitespace characters",
			input: "public class B {\n\t  \ufffa\n}",
			expected: []string{
				"public class B {",
				"\t  \ufffa",
				"}",
			},
		},
		{
			name:  "Delimiter \ufffb on the same line with content (should split)",
			input: "public class B {\n\ufffb\tpublic X(Y client) {}\n}",
			expected: []string{
				"public class B {",
				"\ufffb",
				"\tpublic X(Y client) {}",
				"}",
			},
		},
		{
			name:  "Multiple \ufffb delimiters on separate lines",
			input: "public class B {\n\ufffb\n\ufffb\n}",
			expected: []string{
				"public class B {",
				"\ufffb",
				"\ufffb",
				"}",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			scanner := bufio.NewScanner(bytes.NewReader([]byte(tc.input)))
			scanner.Split(SplitLeadingFFFAOrFFFBOrNewline)
			var tokens []string
			for scanner.Scan() {
				tokens = append(tokens, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				t.Fatalf("Scanner error: %v", err)
			}
			if len(tokens) != len(tc.expected) {
				t.Errorf("Expected %d tokens, got %d", len(tc.expected), len(tokens))
				t.Errorf("Expected tokens: %v", tc.expected)
				t.Errorf("Got tokens: %v", tokens)
				return
			}
			for i, expected := range tc.expected {
				if tokens[i] != expected {
					t.Errorf("Token %d: expected %q, got %q", i, expected, tokens[i])
				}
			}
		})
	}
}
