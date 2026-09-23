package main

// TokenType categorizes a lexeme so downstream phases can switch on a
// category instead of comparing raw strings.
type TokenType string

const (
	// grouping
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
)

// Token bundles a lexeme with the information downstream phases need.
type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}