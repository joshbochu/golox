package parser

import (
	"github.com/joshbochu/lox-go/expr"
	"github.com/joshbochu/lox-go/token"
)

/* Eval Order
expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
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

// expression     → equality ;
func (p *Parser) expression() expr.Expr {
	return p.equality()
}

// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
func (p *Parser) equality() expr.Expr {
	left := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left
}

// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
func (p *Parser) comparison() expr.Expr {
	left := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left
}

// term           → factor ( ( "-" | "+" ) factor )* ;
func (p *Parser) term() expr.Expr {
	left := p.factor()
	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.factor()
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left
}

// factor         → unary ( ( "/" | "*" ) unary )* ;
func (p *Parser) factor() expr.Expr {
	left := p.unary()
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left
}

// unary          → ( "!" | "-" ) unary
func (p *Parser) unary() expr.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &expr.Unary{Operator: operator, Right: right}
	}
	return p.primary()
}

// primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
func (p *Parser) primary() expr.Expr {
	if p.match(token.FALSE) {
		return &expr.Literal{Value: false}
	} else if p.match(token.TRUE) {
		return &expr.Literal{Value: false}
	} else if p.match(token.NIL) {
		return &expr.Literal{Value: nil}
	} else if p.match(token.NUMBER, token.STRING) {
		return &expr.Literal{Value: p.previous().Literal}
	} else if p.match(token.LEFT_PAREN) {
		expression := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return &expr.Grouping{Expression: expression}
	}

	return nil
}

// TODO
func (p *Parser) consume(tokenType token.TokenType, messsage string) {
}

func (p *Parser) match(tokenTypes ...token.TokenType) bool {
	for _, t := range tokenTypes {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}
