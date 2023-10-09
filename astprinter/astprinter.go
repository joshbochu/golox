package astprinter

import (
	"fmt"
	"strings"

	"github.com/joshbochu/lox-go/expr"
)

type Printer struct{}

func (p *Printer) VisitBinaryExpr(expr *expr.Binary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p *Printer) VisitGroupingExpr(expr *expr.Grouping) interface{} {
	return p.parenthesize("grouping", expr.Expression)
}

func (p *Printer) VisitLiteralExpr(expr *expr.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (p *Printer) VisitUnaryExpr(expr *expr.Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (p *Printer) parenthesize(name string, exprs ...expr.Expr) string {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.Accept(p).(string))
	}
	builder.WriteString(")")
	return builder.String()
}
