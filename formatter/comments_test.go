package formatter

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cwarden/go-diff/diffmatchpatch"
	"github.com/octoberswimmer/apexfmt/parser"

	log "github.com/sirupsen/logrus"
)

func TestComments(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.TraceLevel)
		log.SetFormatter(&log.TextFormatter{
			DisableQuote: true,
		})
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
			{
				`// Test trailing whitespace 
				go();`,
				`// Test trailing whitespace
go();`,
			},
			{
				`if (true) {
	if (true) {
		go();
	} // line comment
}`,
				`if (true) {
	if (true) {
		go();
	} // line comment
}`,
			},
			{
				`if (true) {
	if (true) {
		go(); 			// line comment
	}
}`,
				`if (true) {
	if (true) {
		go(); // line comment
	}
}`,
			},
			{
				`
/**
* comment 1
*/
/*
 * comment 2
 */
go();`,
				`
/**
* comment 1
*/
/*
 * comment 2
 */
go();`,
			},
		}
	dmp := diffmatchpatch.New()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
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
				diffs := dmp.DiffMain(tt.output, out, false)
				t.Errorf("unexpected format.  expected:\n%q\ngot:\n%q\ndiff:\n%s\n", tt.output, out, dmp.DiffPrettyText(diffs))
			}
		})
	}

}

func TestTrailingComments(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.TraceLevel)
		log.SetFormatter(&log.TextFormatter{
			DisableQuote: true,
		})
	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`private class T1Exception extends Exception {} //test`,
				"private class T1Exception extends Exception {} //test\n",
			},
			{
				`public class MyClass { public static void noop() {}
	// Comment Inside Compilation Unit
	// Line 2 Inside Compilation Unit
}`,
				`public class MyClass {
	public static void noop() {}
	// Comment Inside Compilation Unit
	// Line 2 Inside Compilation Unit
}`},
			{
				`public class MyClass { public static void noop() {}}
// Comment Outside Compilation Unit Not Moved Inside
// Line 2
`,
				`public class MyClass {
	public static void noop() {}
}
// Comment Outside Compilation Unit Not Moved Inside
// Line 2
`},
			{
				`
/* comment with whitespace before */
private class T1Exception {}`,
				`
/* comment with whitespace before */
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
			{
				`class TestClass {

	// Blank line before comment
	private Integer i;
}`,
				`class TestClass {

	// Blank line before comment
	private Integer i;
}`,
			},
			{
				`class TestClass {
	public static void go() {
	// First Comment

	// Second Comment
go();}}`,
				`class TestClass {
	public static void go() {
		// First Comment

		// Second Comment
		go();
	}
}`,
			},
		}
	dmp := diffmatchpatch.New()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
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
				diffs := dmp.DiffMain(tt.output, out, false)
				t.Errorf("unexpected format.  expected:\n%q\ngot:\n%q\ndiff:\n%s\n", tt.output, out, dmp.DiffPrettyText(diffs))
			}
		})
	}
}
