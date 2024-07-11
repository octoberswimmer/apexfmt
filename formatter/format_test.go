package formatter

import (
	"strconv"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"

	log "github.com/sirupsen/logrus"
)

type testErrorListener struct {
	*antlr.DefaultErrorListener
	t *testing.T
}

func (e *testErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	e.t.Error("Parse Error: line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

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
				`Opportunity o = new Opportunity(
Name = 'My Opportunity',
   AccountId = a.Id,
   StageName = 'Contract Requested/verbal',
   Amount = 1,
   CloseDate = Date.today() + 10
);`,
				// VisitArguments
				`Opportunity o = new Opportunity(
	Name = 'My Opportunity',
	AccountId = a.Id,
	StageName = 'Contract Requested/verbal',
	Amount = 1,
	CloseDate = Date.today() + 10
);`},

			{
				`Psychological__c psyc = Fixtures.psychological(inq).put(Psychological__c.RecordTypeId, Schema.SObjectType.Psychological__c.getRecordTypeInfosByDeveloperName().get('ICD_10').getRecordTypeId()).put(Psychological__c.Diagnosis_Lookup__c, newDiagnosis[0].Id).save();`,
				`Psychological__c psyc = Fixtures.psychological(inq)
	.put(Psychological__c.RecordTypeId, Schema.SObjectType.Psychological__c.getRecordTypeInfosByDeveloperName()
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
			{
				// Don't wrap at what might be an inner class
				`CRC_Inquiry__c inquiry1 = Fixtures.InquiryFactory.inquiry(program1).standardInquiry().patient(patient1).save();`,
				`CRC_Inquiry__c inquiry1 = Fixtures.InquiryFactory.inquiry(program1)
	.standardInquiry()
	.patient(patient1)
	.save();`},

			{
				`return 'lorem' + 'ipsum' + '\n' +
					'lorem' + 'ipsum' + '\n' +
					'lorem' + 'ipsum' + '\n' +
					'lorem' + 'ipsum' + '\n' +
					'lorem' + 'ipsum' + '\n' +
					'lorem' + 'ipsum';
					`,
				`return 'lorem' + 'ipsum' + '\n' + 'lorem' +
	'ipsum' +
	'\n' +
	'lorem' +
	'ipsum' +
	'\n' +
	'lorem' +
	'ipsum' +
	'\n' +
	'lorem' +
	'ipsum' +
	'\n' +
	'lorem' +
	'ipsum';`},

			{
				`Id originalGroupId = ql.SBQQ__RenewedSubscription__r.SBQQ__QuoteLine__r.SBQQ__Group__c != null ?  ql.SBQQ__RenewedSubscription__r.SBQQ__QuoteLine__r.SBQQ__Group__c : ql.SBQQ__UpgradedSubscription__r.SBQQ__QuoteLine__r.SBQQ__Group__c;`,
				`Id originalGroupId = ql.SBQQ__RenewedSubscription__r.SBQQ__QuoteLine__r.SBQQ__Group__c != null ?
	ql.SBQQ__RenewedSubscription__r.SBQQ__QuoteLine__r.SBQQ__Group__c :
	ql.SBQQ__UpgradedSubscription__r.SBQQ__QuoteLine__r.SBQQ__Group__c;`},

			{
				`List<SBQQ__QuoteLineGroup__c> originalGroups = Database.query('SELECT ' + String.join(new List<String>(qlgfields.keySet()), ',') + ' FROM SBQQ__QuoteLineGroup__c WHERE Id IN :originalGroupIds');`,
				`List<SBQQ__QuoteLineGroup__c> originalGroups = Database.query('SELECT ' +
	String.join(new List<String>(qlgfields.keySet()), ',') +
	' FROM SBQQ__QuoteLineGroup__c WHERE Id IN :originalGroupIds');`},

			{
				`public static final List<Schema.SObjectField> MY_IMPORTANT_FIELDS = new List<Schema.SObjectField>{ My_Object__c.The_Field__c, My_Object__c.The_Better_Field__c };`,
				`public static final List<Schema.SObjectField> MY_IMPORTANT_FIELDS = new List<Schema.SObjectField>{
	My_Object__c.The_Field__c,
	My_Object__c.The_Better_Field__c
};`},

			{
				`Error__c[] errorLogs = new Error__c[0];`,
				`Error__c[] errorLogs = new Error__c[0];`},
			{
				`upsert myAccount External_Id__c;`,
				`upsert myAccount External_Id__c;`},
			{
				`List<SObjectField> fs =Schema.getGlobalDescribe().get('MemberPlan').getDescribe().fields.getMap().values();`,
				`List<SObjectField> fs = Schema.getGlobalDescribe()
	.get('MemberPlan')
	.getDescribe().fields.getMap()
	.values();`},
			{
				`Account a=[SELECT Id FROM Account WHERE Id = '001000000FAKEID']??defaultAccount;`,
				`Account a = [SELECT Id FROM Account WHERE Id = '001000000FAKEID'] ?? defaultAccount;`},
			{
				`BatchScheduleManager manager = new BatchScheduleManager(
OneDayDischargeFollowUp.class.getName(),
OneDayDischargeFollowUp.batchJobName,
OneDayDischargeFollowUp.twoHoursDelay,
1
);`,
				`BatchScheduleManager manager = new BatchScheduleManager(
	OneDayDischargeFollowUp.class.getName(),
	OneDayDischargeFollowUp.batchJobName,
	OneDayDischargeFollowUp.twoHoursDelay,
	1
);`},
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
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
		}
	}

}

func TestMemberDeclaration(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`Boolean hasCompleteAddress { get {
  return !String.isBlank(this.upcomingClinic.Location__r.Location_Street_Address__c) && !String.isBlank(this.upcomingClinic.Location__r.Location_City__c) && !String.isBlank(this.upcomingClinic.Location__r.Location_State__c);
}}`,
				`Boolean hasCompleteAddress {
	get {
		return !String.isBlank(this.upcomingClinic.Location__r.Location_Street_Address__c) &&
			!String.isBlank(this.upcomingClinic.Location__r.Location_City__c) &&
			!String.isBlank(this.upcomingClinic.Location__r.Location_State__c);
	}
}`},
			{
				`List<SObjectField> memberPlanFields {
	get {
		if (memberPlanFields == null) {
			List<SObjectField> fs =Schema.getGlobalDescribe()
				.get('MemberPlan')
				.getDescribe().fields.getMap()
				.values();
			List<SObjectField> editable =new List<SObjectField>();

			for (SObjectField f : fs) {
				if (f != MemberPlan.LastViewedDate &&
					f != MemberPlan.LastReferencedDate) {
					editable.add(f);
				}
			}
			memberPlanfields = editable;
		}
	}
	set;
}`, `List<SObjectField> memberPlanFields {
	get {
		if (memberPlanFields == null) {
			List<SObjectField> fs = Schema.getGlobalDescribe()
				.get('MemberPlan')
				.getDescribe().fields.getMap()
				.values();
			List<SObjectField> editable = new List<SObjectField>();

			for (SObjectField f : fs) {
				if (f != MemberPlan.LastViewedDate &&
					f != MemberPlan.LastReferencedDate) {
					editable.add(f);
				}
			}
			memberPlanfields = editable;
		}
	}
	set;
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
		out, ok := v.visitRule(p.MemberDeclaration()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
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
			{
				`public class MyClass { public static void noop() {}}`,
				`public class MyClass {
	public static void noop() {}
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
		out, ok := v.visitRule(p.CompilationUnit()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
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
			{
				`[ SELECT
           Id
        FROM
           Location__c
        WHERE Id IN (
           SELECT
              Location__c
           FROM
              Clinic__c
           WHERE
              Clinic_Type__c IN ('Clinic', 'Clinic - Remote NP') AND
              Status__c = 'Confirmed' AND
              Location__c != null AND
              Start__c = YESTERDAY
        )
     ]`,
				`[
	SELECT
		Id
	FROM
		Location__c
	WHERE
		Id IN (
			SELECT
				Location__c
			FROM
				Clinic__c
			WHERE
				Clinic_Type__c IN ('Clinic', 'Clinic - Remote NP') AND
				Status__c = 'Confirmed' AND
				Location__c != null AND
				Start__c = YESTERDAY
		)
]`},

			{
				`[
	SELECT
		Name
	FROM
		My_Object__c
	WHERE
		Type__c = 'Virtual' AND
		(
			Start__c = TODAY OR
			Start__c = N_DAYS_AGO:7 OR
			Start__c = N_DAYS_AGO:14 OR
			Start__c = N_DAYS_AGO:21 OR
			Start__c <= N_DAYS_AGO:28
		) AND
		Status__c = 'Confirmed'
	ORDER BY
		Start__c
];`, `[
	SELECT
		Name
	FROM
		My_Object__c
	WHERE
		Type__c = 'Virtual' AND
		(
			Start__c = TODAY OR
			Start__c = N_DAYS_AGO:7 OR
			Start__c = N_DAYS_AGO:14 OR
			Start__c = N_DAYS_AGO:21 OR
			Start__c <= N_DAYS_AGO:28
		) AND
		Status__c = 'Confirmed'
	ORDER BY
		Start__c
]`},
			{
				`[SELECT Id FROM ClinicalEncounter WHERE Id = :encounters[0].Id ALL ROWS]`,
				`[SELECT Id FROM ClinicalEncounter WHERE Id = :encounters[0].Id ALL ROWS]`,
			},
			{
				`[SELECT Id, SBQQ__Quote__c FROM SBQQ__QuoteLineGroup__c WHERE SBQQ__Quote__c IN :quoteIds ORDER BY SBQQ__Quote__c, SBQQ__Number__c]`,
				`[
	SELECT
		Id,
		SBQQ__Quote__c
	FROM
		SBQQ__QuoteLineGroup__c
	WHERE
		SBQQ__Quote__c IN :quoteIds
	ORDER BY
		SBQQ__Quote__c,
		SBQQ__Number__c
]`},
		}
	for _, tt := range tests {
		input := antlr.NewInputStream(tt.input)
		lexer := parser.NewApexLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewApexParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(&testErrorListener{t: t})

		v := NewFormatVisitor(stream)
		out, ok := v.visitRule(p.SoqlLiteral()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
		}
	}
}
