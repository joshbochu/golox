package scanner

import (
	"github.com/joshbochu/lox-go/pkg/token"
	"github.com/joshbochu/lox-go/pkg/util"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  make([]token.Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: s.line})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	switch c := s.advance(); c {
	case "(":
		s.addToken(token.LEFT_PAREN)
	case ")":
		s.addToken(token.RIGHT_PAREN)
	case "{":
		s.addToken(token.LEFT_BRACE)
	case "}":
		s.addToken(token.RIGHT_BRACE)
	case ",":
		s.addToken(token.COMMA)
	case ".":
		s.addToken(token.DOT)
	case "-":
		s.addToken(token.MINUS)
	case "+":
		s.addToken(token.PLUS)
	case ";":
		s.addToken(token.SEMICOLON)
	case "*":
		s.addToken(token.STAR)
	case "!":
		if s.match("=") {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case "=":
		if s.match("=") {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case "<":
		if s.match("=") {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case ">":
		if s.match("=") {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	default:
		util.Error(s.line, "Unexpected character.")
	}
}

func (s *Scanner) advance() string {
	c := string(s.source[s.current])
	s.current++
	return c
}

func (s *Scanner) addToken(tokenType token.TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType token.TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) match(expected string) bool {
	if !s.isAtEnd() {
		return false
	}

	c := string(s.source[s.current])
	if c != expected {
		return false
	}

	s.current++
	return true
}
