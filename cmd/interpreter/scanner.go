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
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case '/':
		if s.match('/') {
			// Line comment: scanned and discarded, runs to end of line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}

	// whitespace
	case ' ', '\t', '\r':
		// discard, emit nothing

	// increment s.line when emcountering new line character
	case '\n':
		s.line++

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

// match is peek plus a conditional advance: consumes the expected
// character only if it's actually there. This is what lets "=" and
// "==" be told apart with one character of lookahead.
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() || s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}
