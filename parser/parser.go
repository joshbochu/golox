package parser

import (
	"github.com/joshbochu/lox-go/token"
)

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
