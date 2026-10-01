package formatter

import (
	"strings"
	"testing"
)

// A method may be named void, so void must be accepted as the member name of a
// dot expression.
func Test_formatter_formats_calls_to_a_method_named_void(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "instance call",
			src:  "w.void( 'a' );",
			want: "w.void('a');",
		},
		{
			name: "static call",
			src:  "String s = Test.void( 1 );",
			want: "String s = Test.void(1);",
		},
		{
			name: "safe navigation call",
			src:  "w?.void( 'a' );",
			want: "w?.void('a');",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := "public class Test {\n\tvoid run(Widget w) {\n\t\t" + tt.src + "\n\t}\n}\n"
			want := "public class Test {\n\tvoid run(Widget w) {\n\t\t" + tt.want + "\n\t}\n}\n"
			f := NewFormatter("", strings.NewReader(src))
			got, err := f.Formatted()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
