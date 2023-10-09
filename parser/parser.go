package parser

import (
	"github.com/joshbochu/lox-go/expr"
	"github.com/joshbochu/lox-go/token"
)

/* Eval Order
✅ expression     → equality ;
✅ equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil"
               | "(" expression ")" ;
*/

type Parser struct {
	current int
	tokens  []token.Token
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		current: 0,
		tokens:  tokens,
	}
}

func expression() expr.Expr {
	return equality()
}

func equality() expr.Expr {
	left := comparison()
	for match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := previous()
		right := comparison()
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left
}

func comparison() expr.Expr {
	return nil
}

func previous() token.Token {
	return token.Token{
		Type:    0,
		Lexeme:  "",
		Literal: nil,
		Line:    0,
	}
}

func match(tokenTypes ...token.TokenType) bool {
	return false
}

func check(tokenType token.TokenType) bool {
	return false
}

func advance() token.Token {
	return token.Token{
		Type:    0,
		Lexeme:  "",
		Literal: nil,
		Line:    0,
	}
}

func isAtEnd() bool {
	return false
}

func peek() token.Token {
	return token.Token{
		Type:    0,
		Lexeme:  "",
		Literal: nil,
		Line:    0,
	}
}
