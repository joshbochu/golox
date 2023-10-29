package astprinter

import (
	"fmt"
	"strings"

	"github.com/joshbochu/golox/expr"
)

type Printer struct{}

// VisitVariableExpr implements expr.ExprVisitor.
func (*Printer) VisitVariableExpr(expr *expr.Variable) (interface{}, error) {
	panic("unimplemented")
}

func (p *Printer) VisitBinaryExpr(expr *expr.Binary) (interface{}, error) {
	v, _ := p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
	return v, nil
}

func (p *Printer) VisitGroupingExpr(expr *expr.Grouping) (interface{}, error) {
	v, _ := p.parenthesize("grouping", expr.Expression)
	return v, nil
}

func (p *Printer) VisitLiteralExpr(expr *expr.Literal) (interface{}, error) {
	if expr.Value == nil {
		return "nil", nil
	}
	return fmt.Sprintf("%v", expr.Value), nil
}

func (p *Printer) VisitUnaryExpr(expr *expr.Unary) (interface{}, error) {
	v, _ := p.parenthesize(expr.Operator.Lexeme, expr.Right)
	return v, nil
}

func (p *Printer) parenthesize(name string, exprs ...expr.Expr) (string, error) {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		v, _ := expr.Accept(p)
		builder.WriteString(v.(string))
	}
	builder.WriteString(")")
	return builder.String(), nil
}
