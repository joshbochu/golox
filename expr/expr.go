package ast

import (
	"github.com/joshbochu/lox-go/token"
)

type Expr interface {
	Accept(visitor ExprVisitor) interface{}
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *BinaryExpr) interface{}
	VisitGroupingExpr(expr *GroupingExpr) interface{}
	VisitLiteralExpr(expr *LiteralExpr) interface{}
	VisitUnaryExpr(expr *UnaryExpr) interface{}
}

type BinaryExpr struct {
	left     Expr
	operator token.Token
	right    Expr
}

func (e *BinaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(e)
}

type GroupingExpr struct {
	expression Expr
}

func (e *GroupingExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(e)
}

type LiteralExpr struct {
	value interface{}
}

func (e *LiteralExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(e)
}

type UnaryExpr struct {
	operator token.Token
	right    Expr
}

func (e *UnaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(e)
}
