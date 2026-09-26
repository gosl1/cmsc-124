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

}

func (s *Scanner) scanToken() {

}

func (s *Scanner) isAtEnd() bool {

}
