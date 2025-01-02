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
			{
				`/* preserve newline at end of this comment */
System.debug('I am on a separate line!');`,
				`/* preserve newline at end of this comment */
System.debug('I am on a separate line!');`,
			},
			{
				`return String.isBlank(contact.MailingStreet) &&
	String.isBlank(contact.MailingCity) &&
	String.isBlank(contact.MailingState) &&
	String.isBlank(contact.MailingPostalCode) &&
	// Comment
	(String.isBlank(contact.MailingCountry) ||
	// Country default value
	contact.MailingCountry == 'United States');`,
				`return String.isBlank(contact.MailingStreet) &&
	String.isBlank(contact.MailingCity) &&
	String.isBlank(contact.MailingState) &&
	String.isBlank(contact.MailingPostalCode) &&
	// Comment
	(String.isBlank(contact.MailingCountry) ||
		// Country default value
		contact.MailingCountry == 'United States');`,
			},
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

func TestTrailingComments(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`private class T1Exception extends Exception {} //test`,
				`private class T1Exception extends Exception {} //test`,
			},
			{
				`public class MyClass { public static void noop() {}
	// Comment Inside Compilation Unit
	// Line 2
}`,
				`public class MyClass {
	public static void noop() {}
	// Comment Inside Compilation Unit
	// Line 2
}`},
			{
				`public class MyClass { public static void noop() {}}
// Comment Outside Compilation Unit Moved Inside
// Line 2`,
				`public class MyClass {
	public static void noop() {}
	// Comment Outside Compilation Unit Moved Inside
	// Line 2
}`},
			{
				`
/* comment with whitespace before */
private class T1Exception {}`,
				`/* comment with whitespace before */
private class T1Exception {}`,
			},
			{
				`/* comment with whitespace after */

private class T1Exception {}`,
				`/* comment with whitespace after */

private class T1Exception {}`,
			},
			{
				`class TestClass {
	private void test() {
		statement1();

		// details about statement2
		statement2();
	}
}`,
				`class TestClass {
	private void test() {
		statement1();

		// details about statement2
		statement2();
	}
}`,
			},
		}
	for _, tt := range tests {
		input := antlr.NewInputStream(tt.input)
		lexer := parser.NewApexLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewApexParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(&testErrorListener{t: t})

		v := NewFormatVisitor(stream)
		out, ok := v.visitRule(p.CompilationUnit()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		out = removeExtraCommentIndentation(out)
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%q\ngot:\n%q\n", tt.output, out)
		}
	}
}
