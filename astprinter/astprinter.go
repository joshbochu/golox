package astprinter

import (
	"strings"

	"github.com/joshbochu/lox-go/expr"
)

type AstPrinter struct{}

func (p *AstPrinter) VisitBinaryExpr(expr *expr.BinaryExpr) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p *AstPrinter) VisitGroupingExpr(expr *expr.GroupingExpr) interface{} {
	return p.parenthesize("grouping", expr, expr.Expression)
}

func (p *AstPrinter) VisitLiteralExpr(expr *expr.LiteralExpr) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return expr.Value
}

func (p *AstPrinter) VisitUnaryExpr(expr *expr.UnaryExpr) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (p *AstPrinter) parenthesize(name string, exprs ...expr.Expr) string {
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
