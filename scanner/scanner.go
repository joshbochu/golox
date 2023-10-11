package scanner

import (
	"strconv"

	"github.com/joshbochu/lox-go/token"
	"github.com/joshbochu/lox-go/util"
)

type Scanner struct {
	source   string
	tokens   []token.Token
	start    int
	current  int
	line     int
	keywords map[string]token.TokenType
}

func NewScanner(source string) *Scanner {
	keywords := map[string]token.TokenType{
		"and":    token.AND,
		"class":  token.CLASS,
		"else":   token.ELSE,
		"false":  token.FALSE,
		"for":    token.FOR,
		"fun":    token.FUN,
		"if":     token.IF,
		"nil":    token.NIL,
		"or":     token.OR,
		"print":  token.PRINT,
		"return": token.RETURN,
		"super":  token.SUPER,
		"this":   token.THIS,
		"true":   token.TRUE,
		"var":    token.VAR,
		"while":  token.WHILE,
	}

	return &Scanner{
		source:   source,
		tokens:   make([]token.Token, 0),
		start:    0,
		current:  0,
		line:     1,
		keywords: keywords,
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
	case "/":
		// nextIsNonNewLine && nextInRange
		if s.match(("/")) {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}

	case "\n":
		s.line++
	case " ", "\t", "\r":
		// Ignore whitespace
	case "\"":
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identififer()
		} else {
			util.ErrorLine(s.line, "Unexpected character.")
		}
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
	if s.isAtEnd() {
		return false
	}

	c := string(s.source[s.current])
	if c != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\x00"
	}
	return string(s.source[s.current])
}

func (s *Scanner) string() {
	// nextIsNonTerminal && nextInRange
	for s.peek() != "\"" && !s.isAtEnd() {
		// nextIsNewLine
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		util.ErrorLine(s.line, "Unterminated string")
	}

	// is terminal quote character "
	s.advance()

	stringLiteral := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(token.STRING, stringLiteral)
}

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == "." && isDigit(s.peekNext()) {
		// skip .
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	num, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addTokenWithLiteral(token.NUMBER, num)
}

func (s *Scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return "\x00"
	}
	return string(s.source[s.current+1])
}

func isAlpha(c string) bool {
	isLower := 'a' <= c[0] && c[0] <= 'z'
	isUpper := 'A' <= c[0] && c[0] <= 'Z'
	isUnderScore := c[0] == '_'
	return isLower || isUpper || isUnderScore
}

func isAlphaNumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner) identififer() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	lexeme := s.source[s.start:s.current]
	tokenType, isKeyword := s.keywords[lexeme]

	if isKeyword {
		s.addToken(tokenType)
	} else {
		s.addToken(token.IDENTIFIER)
	}
}
