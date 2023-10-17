package main // to test stuff

import (
	"fmt"

	"github.com/joshbochu/golox/astprinter"
	"github.com/joshbochu/golox/expr"
	"github.com/joshbochu/golox/token"
)

func main() {
	expression := &expr.Binary{
		Left:     &expr.Unary{Operator: token.NewToken(token.MINUS, "-", nil, 1), Right: &expr.Literal{Value: 123}},
		Operator: token.NewToken(token.STAR, "*", nil, 1),
		Right:    &expr.Grouping{Expression: &expr.Literal{Value: 45.67}},
	}

	printer := &astprinter.Printer{}
	fmt.Println(expression.Accept(printer))
}
