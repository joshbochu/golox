package expr

import (
	"github.com/joshbochu/golox/token"
)

type Expr interface {
	Accept(visitor ExprVisitor) (interface{}, error)
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *Binary) (interface{}, error)
	VisitGroupingExpr(expr *Grouping) (interface{}, error)
	VisitLiteralExpr(expr *Literal) (interface{}, error)
	VisitUnaryExpr(expr *Unary) (interface{}, error)
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (e *Binary) Accept(visitor ExprVisitor) (interface{}, error) {
	val, err := visitor.VisitBinaryExpr(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type Grouping struct {
	Expression Expr
}

func (e *Grouping) Accept(visitor ExprVisitor) (interface{}, error) {
	val, err := visitor.VisitGroupingExpr(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type Literal struct {
	Value interface{}
}

func (e *Literal) Accept(visitor ExprVisitor) (interface{}, error) {
	val, err := visitor.VisitLiteralExpr(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (e *Unary) Accept(visitor ExprVisitor) (interface{}, error) {
	val, err := visitor.VisitUnaryExpr(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}
