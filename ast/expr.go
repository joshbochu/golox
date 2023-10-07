package ast

import (
	"github.com/joshbochu/lox-go/token"
)

type Expr interface {
}

type BinaryExpr struct {
	left     Expr
	operator token.Token
	right    Expr
}

type GroupingExpr struct {
	expression Expr
}

type LiteralExpr struct {
	value interface{}
}

type UnaryExpr struct {
	operator token.Token
	right    Expr
}
