package formatter

import (
	"strings"
	"testing"
)

// instanceof types are parsed by instanceOfTypeRef rather than typeRef so
// that the typeArguments decision in typeName is not ambiguous with a
// comparison; they must still format exactly like other types.
func Test_formatter_formats_instanceof_types_like_other_types(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "simple type",
			src:  "if (o instanceof Account) { }",
			want: "if (o instanceof Account) {}",
		},
		{
			name: "generic type",
			src:  "if (o instanceof List < String >) { }",
			want: "if (o instanceof List<String>) {}",
		},
		{
			name: "nested generic type",
			src:  "if (o instanceof Map<String,List<Account>>) { }",
			want: "if (o instanceof Map<String, List<Account>>) {}",
		},
		{
			name: "dotted and array type",
			src:  "if (o instanceof Schema.SObjectType[]) { }",
			want: "if (o instanceof Schema.SObjectType[]) {}",
		},
		{
			name: "followed by a logical operator",
			src:  "if (o instanceof Account && n < 3) { }",
			want: "if (o instanceof Account && n < 3) {}",
		},
		{
			name: "comment before the type",
			src:  "if (o instanceof /* c */ Account) { }",
			want: "if (o instanceof /* c */ Account) {}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := "public class Test {\n\tvoid run(Object o, Integer n) {\n\t\t" + tt.src + "\n\t}\n}\n"
			want := "public class Test {\n\tvoid run(Object o, Integer n) {\n\t\t" + strings.ReplaceAll(tt.want, "\n", "\n\t\t") + "\n\t}\n}\n"
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
