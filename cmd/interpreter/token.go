package main

import "fmt"

// TokenType categorizes a lexeme so downstream phases can switch on a
// category instead of comparing raw strings.
type TokenType string

const (
	// grouping
<<<<<<< HEAD
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"
	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"

	// arithmetic
	PLUS    TokenType = "PLUS"
	MINUS   TokenType = "MINUS"
	STAR    TokenType = "STAR"
	SLASH   TokenType = "SLASH"
	PERCENT TokenType = "PERCENT"

	// assignment
	EQUAL TokenType = "EQUAL"

	// comparison
	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"
	BANG_EQUAL    TokenType = "BANG_EQUAL"
	LESS          TokenType = "LESS"
	LESS_EQUAL    TokenType = "LESS_EQUAL"
	GREATER       TokenType = "GREATER"
	GREATER_EQUAL TokenType = "GREATER_EQUAL"

	// logical
	BANG TokenType = "BANG"
	AND  TokenType = "AND"
	OR   TokenType = "OR"

	// indexing / access
	LEFT_BRACKET  TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"
	DOT           TokenType = "DOT"
	COMMA         TokenType = "COMMA"

	// literals
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

	// keywords (Amadeus vocabulary)
	RES            TokenType = "RES"            // res   — variable declaration
	DMAIL          TokenType = "DMAIL"          // dmail — print/output
	DIV            TokenType = "DIV"            // div   — if
	CON            TokenType = "CON"            // con   — else
	TIMELEAP       TokenType = "TIMELEAP"       // timeleap — while
	OPERATION      TokenType = "OPERATION"      // operation — function declaration
	ELPSY          TokenType = "ELPSY"          // elpsy — return
	READINGSTEINER TokenType = "READINGSTEINER" // readingsteiner — execution-history value
	TRUE           TokenType = "TRUE"
	FALSE          TokenType = "FALSE"
	NULL           TokenType = "NULL"

	EOF TokenType = "EOF"
=======
	TokenTypeLEFT_PAREN  TokenType = "LEFT_PAREN"
	TokenTypeRIGHT_PAREN TokenType = "RIGHT_PAREN"
	TokenTypeLEFT_BRACE  TokenType = "LEFT_BRACE"
	TokenTypeRIGHT_BRACE TokenType = "RIGHT_BRACE"

	// arithmetic
	TokenTypePLUS  TokenType = "PLUS"
	TokenTypeMINUS TokenType = "MINUS"
	TokenTypeSTAR  TokenType = "STAR"
	TokenTypeSLASH TokenType = "SLASH"
	TokenTypePERCENT TokenType = "PERCENT"

	// assignment
	TokenTypeEQUAL TokenType = "EQUAL"

	// comparison
	TokenTypeEQUAL_EQUAL   TokenType = "EQUAL_EQUAL"
	TokenTypeBANG_EQUAL    TokenType = "BANG_EQUAL"
	TokenTypeLESS          TokenType = "LESS"
	TokenTypeLESS_EQUAL    TokenType = "LESS_EQUAL"
	TokenTypeGREATER       TokenType = "GREATER"
	TokenTypeGREATER_EQUAL TokenType = "GREATER_EQUAL"

	// logical
	TokenTypeBANG TokenType = "BANG"
	TokenTypeAND  TokenType = "AND"
	TokenTypeOR   TokenType = "OR"

	// indexing / access
	TokenTypeLEFT_BRACKET  TokenType = "LEFT_BRACKET"
	TokenTypeRIGHT_BRACKET TokenType = "RIGHT_BRACKET"
	TokenTypeDOT           TokenType = "DOT"
  TokenTypeCOMMA         TokenType = "COMMA"

	// literals
	TokenTypeIDENTIFIER TokenType = "IDENTIFIER"
	TokenTypeSTRING     TokenType = "STRING"
	TokenTypeNUMBER     TokenType = "NUMBER"

	// keywords (Amadeus vocabulary)
	TokenTypeRES            TokenType = "RES"            // res   — variable declaration
	TokenTypeDMAIL          TokenType = "DMAIL"          // dmail — print/output
	TokenTypeDIV            TokenType = "DIV"            // div   — if
	TokenTypeCON            TokenType = "CON"            // con   — else
	TokenTypeTIMELEAP       TokenType = "TIMELEAP"       // timeleap — while
	TokenTypeOPERATION      TokenType = "OPERATION"      // operation — function declaration
	TokenTypeELPSY          TokenType = "ELPSY"          // elpsy — return
	TokenTypeREADINGSTEINER TokenType = "READINGSTEINER" // readingsteiner — execution-history value
	TokenTypeTRUE           TokenType = "TRUE"
	TokenTypeFALSE          TokenType = "FALSE"
	TokenTypeNULL           TokenType = "NULL"

	TokenTypeEOF TokenType = "EOF"
>>>>>>> 6b24176a9135ca3e08f131b400c2455adcea7709
)

// Token bundles a lexeme with the information downstream phases need.
type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func (t Token) String() string {
	lit := "null"
	if t.Literal != nil {
		lit = fmt.Sprintf("%v", t.Literal)
	}
	return fmt.Sprintf("Token(type=%s, lexeme=%s, literal=%s, line=%d)",
		t.Type, t.Lexeme, lit, t.Line)
}
