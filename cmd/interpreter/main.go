package main

import (
	"fmt"
	"os"
	"unicode"
)

// --- TOKEN TYPES & STRUCT ---

type TokenType string

const (
	TokenTypeVAR        TokenType = "VAR"
	TokenTypeIDENTIFIER TokenType = "IDENTIFIER"
	TokenTypeEQUAL      TokenType = "EQUAL"
	TokenTypeSTRING     TokenType = "STRING"
	TokenTypePRINT      TokenType = "PRINT"
	TokenTypeEOF        TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}



func main() {
	args := os.Args[1:]

	if len(args) >= 2 && args[0] == "--tokenize" {
		runFileTokenize(args[1])
		return
	}
}

func runFileTokenize(path string) {
	source, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lab1: cannot read '%s': %v\n", path, err)
		os.Exit(65)
	}

	scanner := NewScanner(string(source)) //Have yet to create the function
	tokens, hasError := scanner.ScanTokens() //Have to make this too

	// Exit code 65 on error with clean stdout
	if hasError {
		os.Exit(65)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}