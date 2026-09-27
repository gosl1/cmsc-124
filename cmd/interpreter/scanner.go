package main

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	//Gives a pointer to Scanner and creates the variable implicitly
	return &Scanner{
		source: source,
		tokens: []Token{},
		line:   1,
	}
}

func (s *Scanner) ScanTokens() ([]Token, bool) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: s.line})
	//Ethan go SETUP ERROR//
	return s.tokens, false
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case '[':
		s.addToken(LEFT_BRACKET, nil)
	case ']':
		s.addToken(RIGHT_BRACKET, nil)
	case '.':
		s.addToken(DOT, nil)
	case '+':
		s.addToken(PLUS, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '*':
		s.addToken(STAR, nil)
	case '%':
		s.addToken(PERCENT, nil)
	//case '/':
	default:
	}
}
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// advance consumes and returns the current character.
func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) addToken(tokenType TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line})
}

// peek reads the current character WITHOUT consuming it.
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}
