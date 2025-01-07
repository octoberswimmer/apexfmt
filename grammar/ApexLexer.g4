/*
 [The "BSD licence"]
 Copyright (c) 2013 Terence Parr, Sam Harwell
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions
 are met:
 1. Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.
 2. Redistributions in binary form must reproduce the above copyright
    notice, this list of conditions and the following disclaimer in the
    documentation and/or other materials provided with the distribution.
 3. The name of the author may not be used to endorse or promote products
    derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
 OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
 IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
 INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
 NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

/**
 *  An Apexcode grammar derived from Java 1.7 grammar for ANTLR v4.
 *  Uses ANTLR v4's left-recursive expression notation.
 *
 *  @maintainer: Andrey Gavrikov
 *
 *  You can test with
 *
 *  $ antlr4 ApexGrammar.g4
 *  $ javac *.java
 *  $ grun Apexcode compilationUnit *.cls
 */
lexer grammar ApexLexer;

channels {
    WHITESPACE_CHANNEL,
    COMMENT_CHANNEL
}

// Keywords
ABSTRACT     : A B S T R A C T;
AFTER        : A F T E R;
BEFORE       : B E F O R E;
BREAK        : B R E A K;
CATCH        : C A T C H;
CLASS        : C L A S S;
CONTINUE     : C O N T I N U E;
DELETE       : D E L E T E;
DO           : D O;
ELSE         : E L S E;
ENUM         : E N U M;
EXTENDS      : E X T E N D S;
FINAL        : F I N A L;
FINALLY      : F I N A L L Y;
FOR          : F O R;
GET          : G E T;
GLOBAL       : G L O B A L;
IF           : I F;
IMPLEMENTS   : I M P L E M E N T S;
INHERITED    : I N H E R I T E D;
INSERT       : I N S E R T;
INSTANCEOF   : I N S T A N C E O F;
INTERFACE    : I N T E R F A C E;
MERGE        : M E R G E;
NEW          : N E W;
NULL         : N U L L;
ON           : O N;
OVERRIDE     : O V E R R I D E;
PRIVATE      : P R I V A T E;
PROTECTED    : P R O T E C T E D;
PUBLIC       : P U B L I C;
RETURN       : R E T U R N;
SYSTEMRUNAS  : S Y S T E M '.' R U N A S;
SET          : S E T;
SHARING      : S H A R I N G;
STATIC       : S T A T I C;
SUPER        : S U P E R;
SWITCH       : S W I T C H;
TESTMETHOD   : T E S T M E T H O D;
THIS         : T H I S;
THROW        : T H R O W;
TRANSIENT    : T R A N S I E N T;
TRIGGER      : T R I G G E R;
TRY          : T R Y;
UNDELETE     : U N D E L E T E;
UPDATE       : U P D A T E;
UPSERT       : U P S E R T;
VIRTUAL      : V I R T U A L;
VOID         : V O I D;
WEBSERVICE   : W E B S E R V I C E;
WHEN         : W H E N;
WHILE        : W H I L E;
WITH         : W I T H;
WITHOUT      : W I T H O U T;

// Apex generic types, Set is defined above for properties
LIST          : L I S T;
MAP           : M A P;

// Soql specific keywords
SELECT            : S E L E C T;
COUNT             : C O U N T;
FROM              : F R O M;
AS                : A S;
USING             : U S I N G;
SCOPE             : S C O P E;
WHERE             : W H E R E;
ORDER             : O R D E R;
BY                : B Y;
LIMIT             : L I M I T;
SOQLAND           : A N D;
SOQLOR            : O R;
NOT               : N O T;
AVG               : A V G;
COUNT_DISTINCT    : C O U N T '_' D I S T I N C T;
MIN               : M I N;
MAX               : M A X;
SUM               : S U M;
TYPEOF            : T Y P E O F;
END               : E N D;
THEN              : T H E N;
LIKE              : L I K E;
IN                : I N;
INCLUDES          : I N C L U D E S;
EXCLUDES          : E X C L U D E S;
ASC               : A S C;
DESC              : D E S C;
NULLS             : N U L L S;
FIRST             : F I R S T;
LAST              : L A S T;
GROUP             : G R O U P;
ALL               : A L L;
ROWS              : R O W S;
VIEW              : V I E W;
HAVING            : H A V I N G;
ROLLUP            : R O L L U P;
TOLABEL           : T O L A B E L;
OFFSET            : O F F S E T;
DATA              : D A T A;
CATEGORY          : C A T E G O R Y;
AT                : A T;
ABOVE             : A B O V E;
BELOW             : B E L O W;
ABOVE_OR_BELOW    : A B O V E '_' O R '_' B E L O W;
SECURITY_ENFORCED : S E C U R I T Y '_' E N F O R C E D;
REFERENCE         : R E F E R E N C E;
CUBE              : C U B E;
FORMAT            : F O R M A T;
TRACKING          : T R A C K I N G;
VIEWSTAT          : V I E W S T A T;
CUSTOM            : C U S T O M;
STANDARD          : S T A N D A R D;
DISTANCE          : D I S T A N C E;
GEOLOCATION       : G E O L O C A T I O N;
GROUPING          : G R O U P I N G;
CONVERT_CURRENCY  : C O N V E R T C U R R E N C Y;

// SOQL Date functions
CALENDAR_MONTH      : C A L E N D A R '_' M O N T H;
CALENDAR_QUARTER    : C A L E N D A R '_' Q U A R T E R;
CALENDAR_YEAR       : C A L E N D A R '_' Y E A R;
DAY_IN_MONTH        : D A Y '_' I N '_' M O N T H;
DAY_IN_WEEK         : D A Y '_' I N '_' W E E K;
DAY_IN_YEAR         : D A Y '_' I N '_' Y E A R;
DAY_ONLY            : D A Y '_' O N L Y;
FISCAL_MONTH        : F I S C A L '_' M O N T H;
FISCAL_QUARTER      : F I S C A L '_' Q U A R T E R;
FISCAL_YEAR         : F I S C A L '_' Y E A R;
HOUR_IN_DAY         : H O U R '_' I N '_' D A Y;
WEEK_IN_MONTH       : W E E K '_' I N '_' M O N T H;
WEEK_IN_YEAR        : W E E K '_' I N '_' Y E A R;
CONVERT_TIMEZONE    : C O N V E R T '_' T I M E Z O N E;

// SOQL Date formulas
YESTERDAY                : Y E S T E R D A Y;
TODAY                    : T O D A Y;
TOMORROW                 : T O M O R R O W;
LAST_WEEK                : L A S T '_' W E E K;
THIS_WEEK                : T H I S '_' W E E K;
NEXT_WEEK                : N E X T '_' W E E K;
LAST_MONTH               : L A S T '_' M O N T H;
THIS_MONTH               : T H I S '_' M O N T H;
NEXT_MONTH               : N E X T '_' M O N T H;
LAST_90_DAYS             : L A S T '_90_' D A Y S;
NEXT_90_DAYS             : N E X T '_90_' D A Y S;
LAST_N_DAYS              : L A S T '_' N '_' D A Y S;
NEXT_N_DAYS              : N E X T '_' N '_' D A Y S;
NEXT_N_WEEKS             : N E X T '_' N '_' W E E K S;
LAST_N_WEEKS             : L A S T '_' N '_' W E E K S;
NEXT_N_MONTHS            : N E X T '_' N '_' M O N T H S;
LAST_N_MONTHS            : L A S T '_' N '_' M O N T H S;
THIS_QUARTER             : T H I S '_' Q U A R T E R;
LAST_QUARTER             : L A S T '_' Q U A R T E R;
NEXT_QUARTER             : N E X T '_' Q U A R T E R;
NEXT_N_QUARTERS          : N E X T '_' N '_' Q U A R T E R S;
LAST_N_QUARTERS          : L A S T '_' N '_' Q U A R T E R S;
THIS_YEAR                : T H I S '_' Y E A R;
LAST_YEAR                : L A S T '_' Y E A R;
NEXT_YEAR                : N E X T '_' Y E A R;
NEXT_N_YEARS             : N E X T '_' N '_' Y E A R S;
LAST_N_YEARS             : L A S T '_' N '_' Y E A R S;
THIS_FISCAL_QUARTER      : T H I S '_' F I S C A L '_' Q U A R T E R;
LAST_FISCAL_QUARTER      : L A S T '_' F I S C A L '_' Q U A R T E R;
NEXT_FISCAL_QUARTER      : N E X T '_' F I S C A L '_' Q U A R T E R;
NEXT_N_FISCAL_QUARTERS   : N E X T '_' N '_' F I S C A L '_' Q U A R T E R S;
LAST_N_FISCAL_QUARTERS   : L A S T '_' N '_' F I S C A L '_' Q U A R T E R S;
THIS_FISCAL_YEAR         : T H I S '_' F I S C A L '_' Y E A R;
LAST_FISCAL_YEAR         : L A S T '_' F I S C A L '_' Y E A R;
NEXT_FISCAL_YEAR         : N E X T '_' F I S C A L '_' Y E A R;
NEXT_N_FISCAL_YEARS      : N E X T '_' N '_' F I S C A L '_' Y E A R S;
LAST_N_FISCAL_YEARS      : L A S T '_' N '_' F I S C A L '_' Y E A R S;
N_DAYS_AGO               : N '_' D A Y S '_' A G O;
N_WEEKS_AGO              : N '_' W E E K S '_' A G O;
N_MONTHS_AGO             : N '_' M O N T H S '_' A G O;
N_QUARTERS_AGO           : N '_' Q U A R T E R S '_' A G O;
N_FISCAL_QUARTERS_AGO    : N '_' F I S C A L '_' Q U A R T E R S '_' A G O;
N_YEARS_AGO              : N '_' Y E A R S '_' A G O;
N_FISCAL_YEARS_AGO       : N '_' F I S C A L '_' Y E A R S '_' A G O;

// SOQL Date literal
DateLiteral: Digit Digit Digit Digit '-' Digit Digit '-' Digit Digit;
TimeLiteral: Digit Digit ':' Digit Digit ':' Digit Digit ('.' Digit+ )? (Z | (('+' | '-') Digit+ ( ':' Digit+)? ));
DateTimeLiteral: DateLiteral T TimeLiteral;

// SOQL Currency literal
// (NOTE: this is also a valid Identifier)
IntegralCurrencyLiteral: [a-zA-Z] [a-zA-Z] [a-zA-Z] Digit+;

// SOSL Keywords
FIND             : F I N D;
EMAIL            : E M A I L;
NAME             : N A M E;
PHONE            : P H O N E;
SIDEBAR          : S I D E B A R;
FIELDS           : F I E L D S;
METADATA         : M E T A D A T A;
PRICEBOOKID      : P R I C E B O O K I D;
NETWORK          : N E T W O R K;
SNIPPET          : S N I P P E T;
TARGET_LENGTH    : T A R G E T '_' L E N G T H;
DIVISION         : D I V I S I O N;
RETURNING        : R E T U R N I N G;
LISTVIEW         : L I S T V I E W;

FindLiteral
    :   '[' WS? F I N D WS '\'' FindCharacters? '\''
    ;

fragment
FindCharacters
    :   FindCharacter+
    ;

fragment
FindCharacter
    :   ~['\\]
    |   FindEscapeSequence
    ;

FindLiteralAlt
    :   '[' WS? F I N D WS '{' FindCharactersAlt? '}'
    ;

fragment
FindCharactersAlt
    :   FindCharacterAlt+
    ;

fragment
FindCharacterAlt
    :   ~[}\\]
    |   FindEscapeSequence
    ;

fragment
FindEscapeSequence
    :   '\\' [+\-&|!(){}^"~*?:'\\]
    ;

// §3.10.1 Integer Literals

IntegerLiteral
    :   Digit Digit*
    ;

LongLiteral
    : Digit Digit* [lL]
    ;

NumberLiteral
    :   Digit* '.' Digit Digit* [dD]?
    ;

fragment
HexCharacter
    :   Digit | A | B | C | D | E | F
    ;

fragment
Digit
    :   [0-9]
    ;

// §3.10.3 Boolean Literals

BooleanLiteral
    :   T R U E
    |   F A L S E
    ;

// §3.10.5 String Literals

StringLiteral
    :   '\'' StringCharacters? '\''
    ;

fragment
StringCharacters
    :   StringCharacter+
    ;

fragment
StringCharacter
    :   ~['\\]
    |   EscapeSequence
    ;

// §3.10.6 Escape Sequences for Character and String Literals

fragment
EscapeSequence
    :   '\\' [btnfr"'\\]
    |   '\\u' HexCharacter HexCharacter HexCharacter HexCharacter
    ;

// §3.10.7 The Null Literal

NullLiteral
    :   NULL
    ;


// §3.11 Separators

LPAREN          : '(';
RPAREN          : ')';
LBRACE          : '{';
RBRACE          : '}';
LBRACK          : '[';
RBRACK          : ']';
SEMI            : ';';
COMMA           : ',';
DOT             : '.';

// §3.12 Operators

ASSIGN          : '=';
GT              : '>';
LT              : '<';
BANG            : '!';
TILDE           : '~';
QUESTIONDOT     : '?.';
QUESTION        : '?';
COLON           : ':';
EQUAL           : '==';
TRIPLEEQUAL     : '===';
NOTEQUAL        : '!=';
LESSANDGREATER  : '<>';
TRIPLENOTEQUAL  : '!==';
AND             : '&&';
OR              : '||';
COAL            : '??';
INC             : '++';
DEC             : '--';
ADD             : '+';
SUB             : '-';
MUL             : '*';
DIV             : '/';
BITAND          : '&';
BITOR           : '|';
CARET           : '^';
MOD             : '%';
MAPTO           : '=>';

ADD_ASSIGN      : '+=';
SUB_ASSIGN      : '-=';
MUL_ASSIGN      : '*=';
DIV_ASSIGN      : '/=';
AND_ASSIGN      : '&=';
OR_ASSIGN       : '|=';
XOR_ASSIGN      : '^=';
MOD_ASSIGN      : '%=';
LSHIFT_ASSIGN   : '<<=';
RSHIFT_ASSIGN   : '>>=';
URSHIFT_ASSIGN  : '>>>=';

//
// Additional symbols not defined in the lexical specification
//

ATSIGN : '@';


// §3.8 Identifiers (must appear after all keywords in the grammar)

Identifier
    :   JavaLetter JavaLetterOrDigit*
    ;

// Apex identifiers don't support non-ascii but we maintain Java rules here and post-validate
// so we can give better error messages
fragment
JavaLetter
    :   [a-zA-Z$_] // these are the "java letters" below 0xFF
    |   // covers all characters above 0xFF which are not a surrogate
        ~[\u0000-\u00FF\uD800-\uDBFF]
    |   // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
        [\uD800-\uDBFF] [\uDC00-\uDFFF]
    ;

fragment
JavaLetterOrDigit
    :   [a-zA-Z0-9$_] // these are the "java letters or digits" below 0xFF
    |   // covers all characters above 0xFF which are not a surrogate
        ~[\u0000-\u00FF\uD800-\uDBFF]
    |   // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
        [\uD800-\uDBFF] [\uDC00-\uDFFF]
    ;

fragment A : [aA]; // match either an 'a' or 'A'
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

//
// Whitespace and comments
//

DOC_COMMENT
    :  [ \t\r\n\u000C]* '/**' [\r\n] .*? '*/' [ \t\r\n\u000C]* -> channel(COMMENT_CHANNEL)
    ;

COMMENT
    :  [ \t\r\n\u000C]* '/*' .*? '*/' [ \t\r\n\u000C]* -> channel(COMMENT_CHANNEL)
    ;

LINE_COMMENT
    :  [ \t\r\n\u000C]* '//' ~[\r\n]* -> channel(COMMENT_CHANNEL)
    ;

WS  :  [ \t\r\n\u000C]+ -> channel(WHITESPACE_CHANNEL)
    ;
