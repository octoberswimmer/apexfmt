package formatter

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// The regexp forms that the scanning functions replaced. They are kept here
// as the reference for the differential tests below.
var (
	referenceNewlinePrefixedMultilineComment = regexp.MustCompile("[\n ]*(\t*[￺￹])")
	referenceIndentedInlineComment           = regexp.MustCompile("([^\n￻])\t+￺([^\n])")
	referenceTabPrefixedInlineComment        = regexp.MustCompile(`([\w,{]+)` + "\t+￹")
	referenceInlineCommentPattern            = regexp.MustCompile(`(?s)` + "￹" + `(.*?)` + "￻")
	referenceMultilineCommentPattern         = regexp.MustCompile(`(?s)\t*` + "￺" + `.*?` + "￻")
)

func referenceRemoveIndentationFromComment(comment string) string {
	startIndex := strings.Index(comment, "￺")
	endIndex := strings.LastIndex(comment, "￻")
	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		return comment
	}
	firstLine := comment[:startIndex]
	leadingTabs := strings.Count(firstLine, "\t")
	commentBody := comment[startIndex+len("￺") : endIndex]
	re := regexp.MustCompile(`\n\t{` + strconv.Itoa(leadingTabs) + `}`)
	modifiedComment := re.ReplaceAllString(commentBody, "\n")
	unindented := regexp.MustCompile("^\t*").ReplaceAllString(firstLine, "") + modifiedComment
	return regexp.MustCompile(`(?s)^(\s*)(\S)`).ReplaceAllString(unindented, "$1"+strings.Repeat("\t", leadingTabs)+"$2")
}

var cleanupPasses = []struct {
	name      string
	reference func(string) string
	scan      func(string) string
}{
	{
		name:      "newlines before comment start",
		reference: func(s string) string { return referenceNewlinePrefixedMultilineComment.ReplaceAllString(s, "$1") },
		scan:      removeNewlinesBeforeCommentStart,
	},
	{
		name:      "tabs before indented multiline comment",
		reference: func(s string) string { return referenceIndentedInlineComment.ReplaceAllString(s, "$1￺$2") },
		scan:      removeTabsBeforeIndentedMultilineComment,
	},
	{
		name:      "tabs before inline comment",
		reference: func(s string) string { return referenceTabPrefixedInlineComment.ReplaceAllString(s, "$1￹") },
		scan:      removeTabsBeforeInlineComment,
	},
	{
		name:      "inline comment markers",
		reference: func(s string) string { return referenceInlineCommentPattern.ReplaceAllString(s, "$1") },
		scan:      removeInlineCommentMarkers,
	},
	{
		name: "multiline comments",
		reference: func(s string) string {
			return referenceMultilineCommentPattern.ReplaceAllStringFunc(s, referenceRemoveIndentationFromComment)
		},
		scan: restoreMultilineComments,
	},
}

// Pieces from which test inputs are assembled. Multi-byte pieces exercise
// the rune handling around the markers.
var cleanupPieces = []string{
	"a", "b", "_", "1", ",", "{", "}", ".", ";", "/", "*", "é", "\v",
	" ", "  ", "\t", "\t\t", "\t\t\t", "\n", "\n\n", "\r", "\f",
	"￹", "￺", "￻",
}

func Test_comment_cleanup_scans_match_the_regexp_forms(t *testing.T) {
	inputs := []string{
		"",
		"￺",
		"￻",
		"\t\t￺/* a */￻",
		"x\n\n\t￺/* a\n\t\t * b\n\t\t */￻\n",
		"foo();\t￹// c\n￻",
		"a\t￺\t￺b",
		"{\n\n\t￹// c\n￻\n}",
		"￻\t￺x",
		"\n\t￺\n\t\t * b￻",
	}
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 20000; i++ {
		n := 1 + r.Intn(12)
		var b strings.Builder
		for j := 0; j < n; j++ {
			b.WriteString(cleanupPieces[r.Intn(len(cleanupPieces))])
		}
		inputs = append(inputs, b.String())
	}
	for _, pass := range cleanupPasses {
		t.Run(pass.name, func(t *testing.T) {
			for _, in := range inputs {
				want := pass.reference(in)
				got := pass.scan(in)
				if got != want {
					t.Errorf("input %q: got %q, want %q", in, got, want)
				}
			}
		})
	}
}

func Test_remove_indentation_from_comment_matches_the_regexp_form(t *testing.T) {
	inputs := []string{
		"￺￻",
		"\t￺/* a */￻",
		"\t\t￺\n\t\t * a\n\t\t\t * b\n\t\t */￻",
		"\t￺ \n\t￻",
		"\t￺\t\n\t\t\n￻",
		"\t￺\v\n\t\tx￻",
		"\t￺\r\n\tx￻",
		"no markers",
	}
	r := rand.New(rand.NewSource(2))
	for i := 0; i < 20000; i++ {
		var b strings.Builder
		b.WriteString(strings.Repeat("\t", r.Intn(4)))
		b.WriteString("￺")
		n := r.Intn(10)
		for j := 0; j < n; j++ {
			b.WriteString(cleanupPieces[r.Intn(len(cleanupPieces)-3)])
		}
		b.WriteString("￻")
		inputs = append(inputs, b.String())
	}
	for _, in := range inputs {
		want := referenceRemoveIndentationFromComment(in)
		got := removeIndentationFromComment(in)
		if got != want {
			t.Errorf("input %q: got %q, want %q", in, got, want)
		}
	}
}
