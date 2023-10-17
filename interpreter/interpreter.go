package interpreter

import (
	"github.com/joshbochu/lox-go/expr"
	"github.com/joshbochu/lox-go/token"
)

type Interpreter struct{}

func (i *Interpreter) VisitLiteralExpr(expr *expr.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr *expr.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitBinaryExpr(expr *expr.Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	// Attempt to convert to float64
	leftNum, leftIsNum := left.(float64)
	rightNum, rightIsNum := right.(float64)

	// Attempt to convert to string if not a number
	leftStr, leftIsStr := left.(string)
	rightStr, rightIsStr := right.(string)

	if !(leftIsNum && rightIsNum) && !(leftIsStr && rightIsStr) {
		panic("Expected float64 or string values for binary expr operation")
	}

	switch expr.Operator.Type {
	case token.BANG_EQUAL:
		// return isEqual()
	case token.EQUAL_EQUAL:
		// return isEqual()
	case token.GREATER:
		return leftNum > rightNum
	case token.GREATER_EQUAL:
		return leftNum >= rightNum
	case token.LESS:
		return leftNum < rightNum
	case token.LESS_EQUAL:
		return leftNum <= rightNum
	case token.MINUS:
		return leftNum - rightNum
	case token.SLASH:
		return leftNum / rightNum
	case token.STAR:
		return leftNum * rightNum
	case token.PLUS:
		if leftIsNum && rightIsNum {
			return leftNum + rightNum
		}
		return leftStr + rightStr
	}
	// unreachable
	return nil
}

func (i *Interpreter) VisitUnaryExpr(expr *expr.Unary) interface{} {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case token.BANG:
		return !isTruthy(right)
	case token.MINUS:
		num, ok := right.(float64)
		if !ok {
			panic("Expected a float64 value for unary minus")
		}
		return -num
	}
	// unreachable
	return nil
}

func (i *Interpreter) evaluate(expr expr.Expr) interface{} {
	return expr.Accept(i)
}

func isTruthy(object interface{}) bool {
	if object == nil {
		return false
	}
	val, ok := object.(bool)
	if ok {
		return val
	}
	return true
}
