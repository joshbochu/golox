package interpreter

import (
	"fmt"

	"github.com/joshbochu/golox/expr"
	"github.com/joshbochu/golox/loxerror"
	"github.com/joshbochu/golox/stmt"
	"github.com/joshbochu/golox/token"
)

type Interpreter struct{}

// VisitVarStmt implements stmt.StmtVisitor.
func (*Interpreter) VisitVarStmt(expr *stmt.Var) (interface{}, error) {
	panic("unimplemented")
}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Interpret(statements []stmt.Stmt) {
	for _, statement := range statements {
		i.execute(statement)
	}

}

func (i *Interpreter) execute(stmt stmt.Stmt) (interface{}, error) {
	return stmt.Accept(i)
}

// TODO
func stringify(object interface{}) string {
	if object == nil {
		return "nil"
	}

	switch v := object.(type) {
	case float64:
		txt := fmt.Sprintf("%f", v)
		if txt[len(txt)-2:] == ".0" {
			return txt[:len(txt)-2]
		}
	}

	return fmt.Sprintf("%v", object)
}

func (i *Interpreter) evaluate(expr expr.Expr) (interface{}, error) {
	return expr.Accept(i)
}

func (i *Interpreter) VisitExpressionStmt(stmt *stmt.Expression) (interface{}, error) {
	i.evaluate(stmt.Expression)
	return nil, nil
}

func (i *Interpreter) VisitPrintStmt(stmt *stmt.Print) (interface{}, error) {
	v, _ := i.evaluate(stmt.Expression)
	fmt.Println(stringify(v))
	return nil, nil
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

		return nil, loxerror.NewRuntimeError(expr.Operator, "operands must be two numbers or two strings for + operator.")
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
		return 0.0, loxerror.NewRuntimeError(operator, "operand must be a number")
	}
	return num, nil
}

func checkNumberOperands(operator token.Token, leftOperand interface{}, rightOperand interface{}) (float64, float64, error) {
	leftNum, leftOk := leftOperand.(float64)
	rightNum, rightOk := rightOperand.(float64)
	if !(leftOk && rightOk) {
		return 0.0, 0.0, loxerror.NewRuntimeError(operator, "operands must be numbers")
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
