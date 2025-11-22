package formatter

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/apexfmt/parser"
	log "github.com/sirupsen/logrus"
)

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
	GROUP BY
		Account.Name
]`},
			{
				`[SELECT Account.Name, count(Id) FROM Contact GROUP BY Account.Name HAVING COUNT(Id) > 10]`,
				`[
	SELECT
		Account.Name,
		COUNT(Id)
	FROM
		Contact
	GROUP BY
		Account.Name
	HAVING
		COUNT(Id) > 10
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
			{
				`[SELECT OBJ1__c O1, OBJ2__c O2, OBJ3__c O3, SUM(OBJ4__c) O4, GROUPING(OBJ1__c) O1Group, GROUPING(OBJ2__c) O2Group, GROUPING(OBJ3__c) O3Group FROM OBJ4__c GROUP BY ROLLUP(OBJ1__c, OBJ2__c, OBJ3__c)]`,
				`[
	SELECT
		OBJ1__c O1,
		OBJ2__c O2,
		OBJ3__c O3,
		SUM(OBJ4__c) O4,
		GROUPING(OBJ1__c) O1Group,
		GROUPING(OBJ2__c) O2Group,
		GROUPING(OBJ3__c) O3Group
	FROM
		OBJ4__c
	GROUP BY ROLLUP(OBJ1__c, OBJ2__c, OBJ3__c)
]`},
			{
				`[SELECT Name, (SELECT Id, (SELECT Id, (SELECT Id, (SELECT Id FROM Child4 ) FROM Child3 ) FROM Child2 ) FROM Child1) FROM Parent]`,
				`[
	SELECT
		Name,
		(SELECT
			Id,
			(SELECT
				Id,
				(SELECT
					Id,
					(SELECT
						Id
					FROM
						Child4)
				FROM
					Child3)
			FROM
				Child2)
		FROM
			Child1)
	FROM
		Parent
]`},
			{
				`[ SELECT convertCurrency(Amount) FROM Opportunity ]`,
				`[SELECT CONVERTCURRENCY(Amount) FROM Opportunity]`,
			},
			{
				`[SELECT CALENDAR_MONTH(CONVERTTIMEZONE(CreatedDate)) month FROM Opportunity]`,
				`[SELECT CALENDAR_MONTH(convertTimezone(CreatedDate)) month FROM Opportunity]`,
			},
			{
				`[SELECT Amount, FORMAT(amount) Amt, convertCurrency(amount) convertedAmount,
						 FORMAT(convertCurrency(amount)) convertedCurrency FROM Opportunity where Id = '006R00000024gDtIAI']`,
				`[
	SELECT
		Amount,
		FORMAT(amount) Amt,
		CONVERTCURRENCY(amount) convertedAmount,
		FORMAT(CONVERTCURRENCY(amount)) convertedCurrency
	FROM
		Opportunity
	WHERE
		Id = '006R00000024gDtIAI'
]`,
			},
			{
				`[SELECT FORMAT(MIN(closedate)) Amt FROM opportunity]`,
				`[SELECT FORMAT(MIN(closedate)) Amt FROM opportunity]`,
			},
			{
				`[SELECT Name, DISTANCE(BillingAddress, GEOLOCATION(37.7749, -122.4194), 'mi') dist FROM Account]`,
				`[SELECT Name, DISTANCE(BillingAddress, GEOLOCATION(37.7749, -122.4194), 'mi') dist FROM Account]`,
			},
			{
				`[SELECT Name FROM Account WHERE DISTANCE(BillingAddress, GEOLOCATION(37.7749, -122.4194), 'mi') < 100]`,
				`[SELECT Name FROM Account WHERE DISTANCE(BillingAddress, GEOLOCATION(37.7749, -122.4194), 'mi') < 100]`,
			},
			{
				`[SELECT Name, Id, DISTANCE(ShippingAddress, GEOLOCATION(32.7157, -117.1611), 'km') dist FROM Account ORDER BY DISTANCE(ShippingAddress, GEOLOCATION(32.7157, -117.1611), 'km')]`,
				`[
	SELECT
		Name,
		Id,
		DISTANCE(ShippingAddress, GEOLOCATION(32.7157, -117.1611), 'km') dist
	FROM
		Account
	ORDER BY
		DISTANCE(ShippingAddress, GEOLOCATION(32.7157, -117.1611), 'km')
]`,
			},
			{
				`[SELECT Name, DISTANCE(BillingAddress, GEOLOCATION(:lat, :lng), 'mi') dist FROM Account WHERE Id = :accountId]`,
				`[
	SELECT
		Name,
		DISTANCE(BillingAddress, GEOLOCATION(:lat, :lng), 'mi') dist
	FROM
		Account
	WHERE
		Id = :accountId
]`,
			},
			{
				`[SELECT Name FROM Account WITH USER_MODE]`,
				`[SELECT Name FROM Account WITH USER_MODE]`,
			},
			{
				`[SELECT Name FROM Account WHERE Id = :accountId WITH SYSTEM_MODE]`,
				`[SELECT Name FROM Account WHERE Id = :accountId WITH SYSTEM_MODE]`,
			},
			{
				`[SELECT StageName, COUNT(Id) cnt FROM opportunity GROUP BY StageName HAVING COUNT(Id) > 1]`,
				`[
	SELECT
		StageName,
		COUNT(Id) cnt
	FROM
		opportunity
	GROUP BY
		StageName
	HAVING
		COUNT(Id) > 1
]`,
			},
			{
				`[SELECT Id FROM Account GROUP BY Id HAVING COUNT(Id) > 1]`,
				`[
	SELECT
		Id
	FROM
		Account
	GROUP BY
		Id
	HAVING
		COUNT(Id) > 1
]`,
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
		out, ok := v.visitRule(p.SoqlLiteral()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
		}
	}
}

func TestSOSLWithModeClause(t *testing.T) {
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`[FIND 'Acme' RETURNING Account(Name) WITH USER_MODE]`,
				`[FIND 'Acme'
RETURNING Account(
	Name)
WITH USER_MODE]`,
			},
			{
				`[FIND 'Acme' IN ALL FIELDS RETURNING Account(Name) WITH SYSTEM_MODE]`,
				`[FIND 'Acme'
IN ALL FIELDS
RETURNING Account(
	Name)
WITH SYSTEM_MODE]`,
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
		out, ok := v.visitRule(p.SoslLiteral()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing apex")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
		}
	}
}
