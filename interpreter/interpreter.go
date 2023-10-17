package interpreter

import (
	"fmt"

	"github.com/joshbochu/lox-go/expr"
	"github.com/joshbochu/lox-go/token"
)

type Interpreter struct{}

func (i *Interpreter) evaluate(expr expr.Expr) interface{} {
	return expr.Accept(i)
}

func (i *Interpreter) VisitLiteralExpr(expr *expr.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr *expr.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitBinaryExpr(expr *expr.Binary) interface{} {
	leftObj := i.evaluate(expr.Left)
	rightObj := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.BANG_EQUAL:
		return !isEqual(leftObj, rightObj)
	case token.EQUAL_EQUAL:
		return isEqual(leftObj, rightObj)
	case token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL, token.MINUS, token.SLASH, token.STAR:
		left, leftOk := leftObj.(float64)
		right, rightOk := rightObj.(float64)

		if !leftOk || !rightOk {
			panic(fmt.Sprintf("Expected float64 operands but got (%T, %T).", leftObj, rightObj))
		}

		switch expr.Operator.Type {
		case token.GREATER:
			return left > right
		case token.GREATER_EQUAL:
			return left >= right
		case token.LESS:
			return left < right
		case token.LESS_EQUAL:
			return left <= right
		case token.MINUS:
			return left - right
		case token.SLASH:
			return left / right
		case token.STAR:
			return left * right
		}
	case token.PLUS:
		leftNum, leftIsNum := leftObj.(float64)
		rightNum, rightIsNum := rightObj.(float64)
		if leftIsNum && rightIsNum {
			return leftNum + rightNum
		}

		leftStr, leftIsStr := leftObj.(string)
		rightStr, rightIsStr := rightObj.(string)
		if leftIsStr && rightIsStr {
			return leftStr + rightStr
		}

		panic("Unsupported types for + operator.")
	default:
		panic(fmt.Sprintf("Unexpected token type: %v.", expr.Operator.Type))
	}

	return nil
}

func isEqual(leftObj interface{}, rightObj interface{}) bool {
	if leftObj == nil && rightObj == nil {
		return true
	}
	if leftObj == nil {
		return false
	}
	return leftObj == rightObj
}

func (i *Interpreter) VisitUnaryExpr(expr *expr.Unary) interface{} {
	rightObj := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case token.BANG:
		return !isTruthy(rightObj)
	case token.MINUS:
		num, ok := rightObj.(float64)
		if !ok {
			panic(fmt.Sprintf("Expected float64 operand but got %T", rightObj))
		}
		return -num
	}
	// unreachable
	return nil
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
