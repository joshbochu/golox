package parser

import (
	"github.com/joshbochu/lox-go/expr"
	"github.com/joshbochu/lox-go/token"
	"github.com/joshbochu/lox-go/util"
)

type ParseError struct {
	message string
}

func NewParseError(message string) *ParseError {
	return &ParseError{message: message}
}

func (e *ParseError) Error() string {
	return e.message
}

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

func (p *Parser) Parse() (expr.Expr, error) {
	expression, err := p.expression()
	if err != nil {
		return nil, nil
	}
	return expression, nil
}

// expression     → equality ;
func (p *Parser) expression() (expr.Expr, error) {
	return p.equality()
}

// equality       → comparison ( ( "!=" | "==" ) comparison )* ;
func (p *Parser) equality() (expr.Expr, error) {
	left, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left, nil
}

// comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
func (p *Parser) comparison() (expr.Expr, error) {
	left, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left, nil
}

// term           → factor ( ( "-" | "+" ) factor )* ;
func (p *Parser) term() (expr.Expr, error) {
	left, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left, nil
}

// factor         → unary ( ( "/" | "*" ) unary )* ;
func (p *Parser) factor() (expr.Expr, error) {
	left, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		left = &expr.Binary{Left: left, Operator: operator, Right: right}
	}
	return left, nil
}

// unary          → ( "!" | "-" ) unary
func (p *Parser) unary() (expr.Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &expr.Unary{Operator: operator, Right: right}, nil
	}
	return p.primary()
}

// primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
func (p *Parser) primary() (expr.Expr, error) {
	if p.match(token.FALSE) {
		return &expr.Literal{Value: false}, nil
	} else if p.match(token.TRUE) {
		return &expr.Literal{Value: true}, nil
	} else if p.match(token.NIL) {
		return &expr.Literal{Value: nil}, nil
	} else if p.match(token.NUMBER, token.STRING) {
		return &expr.Literal{Value: p.previous().Literal}, nil
	} else if p.match(token.LEFT_PAREN) {
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return &expr.Grouping{Expression: expression}, nil
	}

	return nil, p.error(p.peek(), "Expression Expected")
}

func (p *Parser) consume(tokenType token.TokenType, messsage string) (token.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}
	return token.Token{}, p.error(p.peek(), messsage)
}

func (p *Parser) error(token token.Token, message string) error {
	util.ErrorToken(token, message)
	return NewParseError(message)
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

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case token.CLASS, token.FUN, token.VAR, token.FOR, token.IF, token.WHILE, token.PRINT, token.RETURN:
			return
		}

		p.advance()
	}
}
