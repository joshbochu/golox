package interpreter

import (
	"testing"

	"github.com/joshbochu/golox/expr"
	"github.com/joshbochu/golox/token"
)

func TestInterpreter(t *testing.T) {
	interpreter := NewInterpreter()
	tests := []struct {
		name       string
		expression expr.Expr
		expected   interface{}
		expectErr  bool
		errMsg     string
	}{
		{
			name:       "5 -> 5",
			expression: &expr.Literal{Value: float64(5)},
			expected:   float64(5),
		},
		{
			name:       `"Hello" -> "Hello"`,
			expression: &expr.Literal{Value: "Hello"},
			expected:   "Hello",
		},
		{
			name:       "-5 -> -5",
			expression: &expr.Unary{Operator: token.NewToken(token.MINUS, "-", nil, 1), Right: &expr.Literal{Value: float64(5)}},
			expected:   float64(-5),
		},
		{
			name:       "!true -> false",
			expression: &expr.Unary{Operator: token.NewToken(token.BANG, "!", nil, 1), Right: &expr.Literal{Value: true}},
			expected:   false,
		},
		{
			name:       "5 + 3 -> 8",
			expression: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.PLUS, "+", nil, 1), Right: &expr.Literal{Value: float64(3)}},
			expected:   float64(8),
		},
		{
			name:       "5 * 3 -> 15",
			expression: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.STAR, "*", nil, 1), Right: &expr.Literal{Value: float64(3)}},
			expected:   float64(15),
		},
		{
			name:       "(5 + 3) -> 8",
			expression: &expr.Grouping{Expression: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.PLUS, "+", nil, 1), Right: &expr.Literal{Value: float64(3)}}},
			expected:   float64(8),
		},
		{
			name: "4 + (5 * 3) - 2 -> 17",
			expression: &expr.Binary{
				Left: &expr.Binary{
					Left:     &expr.Literal{Value: float64(4)},
					Operator: token.NewToken(token.PLUS, "+", nil, 1),
					Right: &expr.Binary{
						Left:     &expr.Literal{Value: float64(5)},
						Operator: token.NewToken(token.STAR, "*", nil, 1),
						Right:    &expr.Literal{Value: float64(3)},
					},
				},
				Operator: token.NewToken(token.MINUS, "-", nil, 1),
				Right:    &expr.Literal{Value: float64(2)},
			},
			expected: float64(17),
		},
		{
			name: "-!-5 -> Error",
			expression: &expr.Unary{
				Operator: token.NewToken(token.MINUS, "-", nil, 1),
				Right: &expr.Unary{
					Operator: token.NewToken(token.BANG, "!", nil, 1),
					Right: &expr.Unary{
						Operator: token.NewToken(token.MINUS, "-", nil, 1),
						Right:    &expr.Literal{Value: float64(5)},
					},
				},
			},
			expectErr: true,
			errMsg:    "operand must be a number",
		},
		{
			name: "-5 + 3 -> -2",
			expression: &expr.Binary{
				Left: &expr.Unary{
					Operator: token.NewToken(token.MINUS, "-", nil, 1),
					Right:    &expr.Literal{Value: float64(5)},
				},
				Operator: token.NewToken(token.PLUS, "+", nil, 1),
				Right:    &expr.Literal{Value: float64(3)},
			},
			expected: float64(-2),
		},
		{
			name: "-(5 + 3) -> -8",
			expression: &expr.Unary{
				Operator: token.NewToken(token.MINUS, "-", nil, 1),
				Right: &expr.Grouping{
					Expression: &expr.Binary{
						Left:     &expr.Literal{Value: float64(5)},
						Operator: token.NewToken(token.PLUS, "+", nil, 1),
						Right:    &expr.Literal{Value: float64(3)},
					},
				},
			},
			expected: float64(-8),
		},
		{
			name:       "Addition of number and string",
			expression: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.PLUS, "+", nil, 1), Right: &expr.Literal{Value: "hello"}},
			expectErr:  true,
			errMsg:     "operands must be two numbers or two strings for + operator.",
		},
		{
			name:       "Multiplication of number and string",
			expression: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.STAR, "*", nil, 1), Right: &expr.Literal{Value: "hello"}},
			expectErr:  true,
			errMsg:     "operands must be numbers",
		},
		{
			name:       "Comparison of number and string",
			expression: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.GREATER, ">", nil, 1), Right: &expr.Literal{Value: "hello"}},
			expectErr:  true,
			errMsg:     "operands must be numbers",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := interpreter.evaluate(test.expression)

			// Check if error was expected but didn't occur
			if test.expectErr && err == nil {
				t.Errorf("Expected an error but got none for expr: %v", test.expression)
				return
			}

			// Check if error was not expected but did occur
			if !test.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check if the error message is what was expected
			if test.expectErr && err.Error() != test.errMsg {
				t.Errorf("Expected error message: %q, but got: %q", test.errMsg, err.Error())
				return
			}

			// Check if the returned value matches the expected value
			if val != test.expected {
				t.Errorf("Expected value %v, but got %v", test.expected, val)
			}
		})
	}
}
