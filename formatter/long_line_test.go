package formatter

import (
	"strings"
	"testing"
)

// bufio.Scanner refuses tokens longer than 64 KB. indentTo used it to walk
// each block's lines and ignored the error, so everything from a long line
// to the end of the block was dropped from the output.
func Test_formatter_preserves_lines_longer_than_64_kilobytes(t *testing.T) {
	literal := strings.Repeat("x", 70000)
	src := "public class Test {\n\tpublic void run() {\n\t\tString s = '" + literal + "';\n\t\tSystem.debug(s);\n\t}\n}\n"
	f := NewFormatter("", strings.NewReader(src))
	got, err := f.Formatted()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != src {
		if len(got) < 200 {
			t.Fatalf("got %q, want the input unchanged", got)
		}
		t.Fatalf("got %d bytes, want %d bytes unchanged", len(got), len(src))
	}
}
