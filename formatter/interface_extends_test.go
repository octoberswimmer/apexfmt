package formatter

import (
	"strings"
	"testing"
)

func Test_formatter_writes_one_space_between_interface_extends_and_body(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "one parent interface",
			src:  "public interface Child extends Parent {\n\tvoid run();\n}\n",
			want: "public interface Child extends Parent {\n\tvoid run();\n}\n",
		},
		{
			name: "several parent interfaces",
			src:  "public interface Child extends First,Second {\n\tvoid run();\n}\n",
			want: "public interface Child extends First, Second {\n\tvoid run();\n}\n",
		},
		{
			name: "extra spaces in the source",
			src:  "public interface Child extends Parent   {\n\tvoid run();\n}\n",
			want: "public interface Child extends Parent {\n\tvoid run();\n}\n",
		},
		{
			name: "no parent interface",
			src:  "public interface Child {\n\tvoid run();\n}\n",
			want: "public interface Child {\n\tvoid run();\n}\n",
		},
		{
			name: "nested interface",
			src:  "public class Outer {\n\tpublic interface Child extends Parent {\n\t\tvoid run();\n\t}\n}\n",
			want: "public class Outer {\n\tpublic interface Child extends Parent {\n\t\tvoid run();\n\t}\n}\n",
		},
	}
	format := func(t *testing.T, src string) string {
		t.Helper()
		got, err := NewFormatter("", strings.NewReader(src)).Formatted()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		return got
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := format(t, tt.src)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if again := format(t, got); again != got {
				t.Errorf("formatting again changed the output: got %q, want %q", again, got)
			}
		})
	}
}
