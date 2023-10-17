package scanner

import (
	"testing"

	"github.com/joshbochu/golox/token"
)

func TestScanner_ScanTokens(t *testing.T) {
	tests := []struct {
		name   string
		source string
		tokens []token.TokenType
	}{
		{"Single characters", "()", []token.TokenType{token.LEFT_PAREN, token.RIGHT_PAREN, token.EOF}},
		{"Math operators", "+-*/", []token.TokenType{token.PLUS, token.MINUS, token.STAR, token.SLASH, token.EOF}},
		{"Comparison", "!=", []token.TokenType{token.BANG_EQUAL, token.EOF}},
		{"Number", "123.456", []token.TokenType{token.NUMBER, token.EOF}},
		{"String", "\"test string\"", []token.TokenType{token.STRING, token.EOF}},
		{"Variable declaration", "var x = 5;", []token.TokenType{token.VAR, token.IDENTIFIER, token.EQUAL, token.NUMBER, token.SEMICOLON, token.EOF}},
		{"For loop", "for (i = 0; i < 10; i = i + 1) {}", []token.TokenType{
			token.FOR, token.LEFT_PAREN, token.IDENTIFIER, token.EQUAL, token.NUMBER, token.SEMICOLON,
			token.IDENTIFIER, token.LESS, token.NUMBER, token.SEMICOLON,
			token.IDENTIFIER, token.EQUAL, token.IDENTIFIER, token.PLUS, token.NUMBER,
			token.RIGHT_PAREN, token.LEFT_BRACE, token.RIGHT_BRACE, token.EOF}},
		{"Multiline", `
			var x = 10;
			print x;
			x = x * 2;
		`, []token.TokenType{
			token.VAR, token.IDENTIFIER, token.EQUAL, token.NUMBER, token.SEMICOLON,
			token.PRINT, token.IDENTIFIER, token.SEMICOLON,
			token.IDENTIFIER, token.EQUAL, token.IDENTIFIER, token.STAR, token.NUMBER, token.SEMICOLON,
			token.EOF}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := NewScanner(test.source)
			tokens := scanner.ScanTokens()
			for i, tt := range test.tokens {
				if len(tokens) <= i {
					t.Errorf("Not enough tokens produced for test: %s", test.name)
					break
				}
				if tokens[i].Type != tt {
					t.Errorf("Expected token %v but got %v", tt, tokens[i].Type)
				}
			}
		})
	}
}
