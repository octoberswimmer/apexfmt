package formatter

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"

	log "github.com/sirupsen/logrus"
)

func TestComments(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{

			{
				`for (Referral__c ref : [SELECT Summary_Name__c, Name FROM Referral__c WHERE Id IN :referralIdSet]) {
	/* Multi-line comments
	preserve white-space.
		So they can contain
			{ code }
				or poetry
		for
			example
	*/
  System.assertEquals(ref.Name, ref.Summary_Name__c);
}`,
				`for (Referral__c ref : [
	SELECT
		Summary_Name__c,
		Name
	FROM
		Referral__c
	WHERE
		Id IN :referralIdSet
]) {
	/* Multi-line comments
	preserve white-space.
		So they can contain
			{ code }
				or poetry
		for
			example
	*/
	System.assertEquals(ref.Name, ref.Summary_Name__c);
}`},
		}
	for _, tt := range tests {
		input := antlr.NewInputStream(tt.input)
		lexer := parser.NewApexLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewApexParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(&testErrorListener{t: t})

		v := NewFormatVisitor(stream)
		out, ok := v.visitRule(p.Statement()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		out = removeExtraCommentIndentation(out)
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%q\ngot:\n%q\n", tt.output, out)
		}
	}

}
