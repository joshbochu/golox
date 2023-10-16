package interpreter

import (
	"github.com/joshbochu/lox-go/expr"
)

type Interpreter struct{}

func (i *Interpreter) VisitLiteralExpr(expr *expr.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr *expr.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitBinaryExpr(expr *expr.Binary) interface{} {
	return nil
}

func (i *Interpreter) VisitUnaryExpr(expr *expr.Unary) interface{} {
	return nil
}

func (i *Interpreter) evaluate(expr expr.Expr) interface{} {
	return expr.Accept(i)
}
