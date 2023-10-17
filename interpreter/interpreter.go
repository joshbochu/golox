package interpreter

import (
	"fmt"

	"github.com/joshbochu/lox-go/expr"
	"github.com/joshbochu/lox-go/token"
)

type RuntimeError struct {
	token   token.Token
	message string
}

func NewRuntimeError(token token.Token, message string) *RuntimeError {
	return &RuntimeError{token: token, message: message}
}

func (e *RuntimeError) Error() string {
	return e.message
}

type Interpreter struct{}

func (i *Interpreter) Interpret(expr expr.Expr) {
	val, err := i.evaluate(expr)
	if err != nil {
		// reportruntimeerror
	}
	fmt.Println(stringify(val))
}

// TODO
func stringify(object interface{}) string {
	return ""
}

func (i *Interpreter) evaluate(expr expr.Expr) (interface{}, error) {
	return expr.Accept(i)
}

func (i *Interpreter) VisitLiteralExpr(expr *expr.Literal) (interface{}, error) {
	return expr.Value, nil
}

func (i *Interpreter) VisitGroupingExpr(expr *expr.Grouping) (interface{}, error) {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitBinaryExpr(expr *expr.Binary) (interface{}, error) {
	leftObj, leftErr := i.evaluate(expr.Left)
	if leftErr != nil {
		return nil, leftErr
	}

	rightObj, rightErr := i.evaluate(expr.Right)
	if rightErr != nil {
		return nil, rightErr
	}

	switch expr.Operator.Type {
	case token.BANG_EQUAL:
		return !isEqual(leftObj, rightObj), nil
	case token.EQUAL_EQUAL:
		return isEqual(leftObj, rightObj), nil
	case token.GREATER:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err != nil {
			return nil, err
		}
		return left > right, nil
	case token.GREATER_EQUAL:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err != nil {
			return nil, err
		}
		return left >= right, nil
	case token.LESS:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err != nil {
			return nil, err
		}
		return left < right, nil
	case token.LESS_EQUAL:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err != nil {
			return nil, err
		}
		return left <= right, nil
	case token.MINUS:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err != nil {
			return nil, err
		}
		return left - right, nil
	case token.SLASH:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err != nil {
			return nil, err
		}
		return left / right, nil
	case token.STAR:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err != nil {
			return nil, err
		}
		return left * right, nil
	case token.PLUS:
		left, right, err := checkNumberOperands(expr.Operator, leftObj, rightObj)
		if err == nil {
			return left + right, nil
		}

		leftStr, leftStrOk := leftObj.(string)
		rightStr, rightStrOk := rightObj.(string)
		if leftStrOk && rightStrOk {
			return leftStr + rightStr, nil
		}

		return nil, NewRuntimeError(expr.Operator, "operands must be two numbers or two strings for + operator.")
	}

	return nil, nil
}

func (i *Interpreter) VisitUnaryExpr(expr *expr.Unary) (interface{}, error) {
	rightObj, _ := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case token.BANG:
		return !isTruthy(rightObj), nil
	case token.MINUS:
		num, err := checkNumberOperand(expr.Operator, rightObj)
		if err != nil {
			return nil, err
		}
		return -num, nil
	}
	// unreachable
	return nil, nil
}

func checkNumberOperand(operator token.Token, operand interface{}) (float64, error) {
	num, ok := operand.(float64)
	if !ok {
		return 0.0, NewRuntimeError(operator, "operand must be a number")
	}
	return num, nil
}

func checkNumberOperands(operator token.Token, leftOperand interface{}, rightOperand interface{}) (float64, float64, error) {
	leftNum, leftOk := leftOperand.(float64)
	rightNum, rightOk := rightOperand.(float64)
	if !(leftOk && rightOk) {
		return 0.0, 0.0, NewRuntimeError(operator, "operands must be numbers")
	}
	return leftNum, rightNum, nil

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
