package formatter

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"

	log "github.com/sirupsen/logrus"
)

func TestStatement(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`Account a = new Account(Name='Acme', RecordTypeId=myRecordTypeId, BillingCity='Los Angeles', BillingState = 'CA');`,
				`Account a = new Account(
	Name = 'Acme',
	RecordTypeId = myRecordTypeId,
	BillingCity = 'Los Angeles',
	BillingState = 'CA'
);`},

			{
				`System.runAs(u) {
   facility = Fixtures.account().put(Schema.Account.RecordTypeId, facilityRecordType).put(Schema.Account.HealthCloudGA__SourceSystemId__c, '0001')
      .put(Schema.Account.Patient_Owner__c, u.Id)
      .save();
}`, `System.runAs(u) {
	facility = Fixtures.account()
		.put(Schema.Account.RecordTypeId, facilityRecordType)
		.put(Schema.Account.HealthCloudGA__SourceSystemId__c, '0001')
		.put(Schema.Account.Patient_Owner__c, u.Id)
		.save();
}`},

			{

				`System.assertEquals(UserInfo.getUserId(), [SELECT OwnerId FROM Account WHERE Id = :person.Id].OwnerId, 'Account should be owned by correct user');`,
				`System.assertEquals(UserInfo.getUserId(), [SELECT OwnerId FROM Account WHERE Id = :person.Id].OwnerId, 'Account should be owned by correct user');`},
			{
				`System.assert(lsr[0].getErrors()[0].getMessage().contains(constants.ERR_MSG_NO_CLIENT_DEMOGRAPHICS), 'error message');`,
				`System.assert(lsr[0].getErrors()[0].getMessage().contains(constants.ERR_MSG_NO_CLIENT_DEMOGRAPHICS), 'error message');`,
			},

			{
				`RecordType referralType = [ SELECT Id FROM RecordType WHERE SobjectType = 'Contact' AND DeveloperName = 'Referral_Contact' ];`,
				`RecordType referralType = [
	SELECT
		Id
	FROM
		RecordType
	WHERE
		SobjectType = 'Contact' AND
		DeveloperName = 'Referral_Contact'
];`},

			{
				`update [SELECT Id FROM Territory_Coverage__c WHERE Named_Account__c IN :accountIds];`,
				`update [SELECT Id FROM Territory_Coverage__c WHERE Named_Account__c IN :accountIds];`,
			},

			{
				`for (Referral__c ref : [SELECT Summary_Name__c, Name FROM Referral__c WHERE Id IN :referralIdSet]) {
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
	System.assertEquals(ref.Name, ref.Summary_Name__c);
}`},

			{
				`for (Referral__c ref : [SELECT Name FROM Referral__c WHERE Id IN :referralIdSet]) {
  System.assertEquals('test', ref.Summary_Name__c);
}`,
				`for (Referral__c ref : [SELECT Name FROM Referral__c WHERE Id IN :referralIdSet]) {
	System.assertEquals('test', ref.Summary_Name__c);
}`},

			{
				`if (!r.isSuccess()) {
   throw new BenefitCheckNotificationException(
      'Failed to send Benefit Check notification.  First error: ' +
         r.getErrors()[0].getMessage()
   );
}`,
				`if (!r.isSuccess()) {
	throw new BenefitCheckNotificationException(
		'Failed to send Benefit Check notification.  First error: ' +
			r.getErrors()[0].getMessage()
	);
}`},

			{
				`return new List<CountryZip>{
   new CountryZip( new Territory_Zip_Lookup__c( Id = zip.Id, Name = zip.Name, City__c = zip.City__c, State_2_Letter_Code__c = zip.State_2_Letter_Code__c, Country__c = zip.Country__c))
};`,
				`return new List<CountryZip>{
	new CountryZip(
		new Territory_Zip_Lookup__c(
			Id = zip.Id,
			Name = zip.Name,
			City__c = zip.City__c,
			State_2_Letter_Code__c = zip.State_2_Letter_Code__c,
			Country__c = zip.Country__c
		)
	)
};`},

			{
				`Psychological__c psyc = Fixtures.psychological(inq).put(Psychological__c.RecordTypeId, Schema.SObjectType.Psychological__c.getRecordTypeInfosByDeveloperName().get('ICD_10').getRecordTypeId()).put(Psychological__c.Diagnosis_Lookup__c, newDiagnosis[0].Id).save();`,
				`Psychological__c psyc = Fixtures.psychological(inq)
	.put(Psychological__c.RecordTypeId, Schema.SObjectType.Psychological__c
		.getRecordTypeInfosByDeveloperName()
		.get('ICD_10')
		.getRecordTypeId())
	.put(Psychological__c.Diagnosis_Lookup__c, newDiagnosis[0].Id)
	.save();`},

			{
				`this.assertPassed(Assert.isNumericallyInner(2, '~0.5', 2.4, null));`,
				`this.assertPassed(Assert.isNumericallyInner(2, '~0.5', 2.4, null));`},

			{
				`assertFailed(Assert.consistOfInner(new List<Object>{ a, b }, new List<Object>{ b, c }, 'doodle'), 'doodle: expected (1, 2) to consist of (2, 3)\nextra elements:\n\t(1)\nmissing elements:\n\t(3)');`,
				`assertFailed(Assert.consistOfInner(new List<Object>{ a, b }, new List<Object>{ b, c }, 'doodle'),
	'doodle: expected (1, 2) to consist of (2, 3)\nextra elements:\n\t(1)\nmissing elements:\n\t(3)');`},

			{
				`if (cl_record.Last_Placement__c == true &&
					(Trigger.isInsert || (Trigger.isUpdate && cl_record.Last_Placement__c != Trigger.OldMap.get(cl_record.Id).Last_Placement__c))) {x=1;}`,
				`if (cl_record.Last_Placement__c == true &&
	(Trigger.isInsert ||
	(Trigger.isUpdate &&
	cl_record.Last_Placement__c != Trigger.OldMap.get(cl_record.Id).Last_Placement__c))) {
	x = 1;
}`},
		}
	for _, tt := range tests {
		input := antlr.NewInputStream(tt.input)
		lexer := parser.NewApexLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewApexParser(stream)
		p.RemoveErrorListeners()

		v := NewFormatVisitor(stream)
		out, ok := v.visitRule(p.Statement()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s;\ngot:\n%s\n", tt.output, out)
		}
	}

}

func TestCompilationUnit(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`private class T1Exception extends Exception {}`,
				`private class T1Exception extends Exception {}`},
		}
	for _, tt := range tests {
		input := antlr.NewInputStream(tt.input)
		lexer := parser.NewApexLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewApexParser(stream)
		p.RemoveErrorListeners()

		v := NewFormatVisitor(stream)
		out, ok := v.visitRule(p.CompilationUnit()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s;\ngot:\n%s\n", tt.output, out)
		}
	}
}

func TestSOQL(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`[SELECT Account.Name, count(Id) FROM Contact WHERE AccountId IN : accounts.keySet() GROUP BY Account.Name]`,
				`[
	SELECT
		Account.Name,
		COUNT(Id)
	FROM
		Contact
	WHERE
		AccountId IN :accounts.keySet()
	GROUP BY Account.Name
]`},
		}
	for _, tt := range tests {
		input := antlr.NewInputStream(tt.input)
		lexer := parser.NewApexLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewApexParser(stream)
		p.RemoveErrorListeners()

		v := NewFormatVisitor(stream)
		out, ok := v.visitRule(p.SoqlLiteral()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s;\ngot:\n%s\n", tt.output, out)
		}
	}
}
