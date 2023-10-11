package parser

import (
	"reflect"
	"testing"

	"github.com/joshbochu/lox-go/expr"
	"github.com/joshbochu/lox-go/scanner"
	"github.com/joshbochu/lox-go/token"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		expected  expr.Expr
		expectErr bool
	}{
		{
			name:     "Unary operation",
			source:   "-5",
			expected: &expr.Unary{Operator: token.NewToken(token.MINUS, "-", nil, 1), Right: &expr.Literal{Value: float64(5)}},
		},
		{
			name:     "Binary operation",
			source:   "5 + 3",
			expected: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.PLUS, "+", nil, 1), Right: &expr.Literal{Value: float64(3)}},
		},
		{
			name:     "Grouping",
			source:   "(5 + 3)",
			expected: &expr.Grouping{Expression: &expr.Binary{Left: &expr.Literal{Value: float64(5)}, Operator: token.NewToken(token.PLUS, "+", nil, 1), Right: &expr.Literal{Value: float64(3)}}},
		},
		{
			name:     "String Literal",
			source:   "\"Hello\"",
			expected: &expr.Literal{Value: "Hello"},
		},
		{
			name:     "Number Literal",
			source:   "42",
			expected: &expr.Literal{Value: float64(42)},
		},
		{
			name:     "Boolean Literal",
			source:   "true",
			expected: &expr.Literal{Value: true},
		},
		{
			name:     "Nil Literal",
			source:   "nil",
			expected: &expr.Literal{Value: nil},
		},
		{
			name:   "Nested Binary Operations",
			source: "4 + 5 * 3 - 2",
			expected: &expr.Binary{
				Left: &expr.Binary{
					Left:     &expr.Literal{Value: float64(4)},
					Operator: token.Token{Type: token.PLUS, Lexeme: "+", Line: 1},
					Right: &expr.Binary{
						Left:     &expr.Literal{Value: float64(5)},
						Operator: token.Token{Type: token.STAR, Lexeme: "*", Line: 1},
						Right:    &expr.Literal{Value: float64(3)},
					},
				},
				Operator: token.Token{Type: token.MINUS, Lexeme: "-", Line: 1},
				Right:    &expr.Literal{Value: float64(2)},
			},
		},
		{
			name:   "Nested Unary Operations",
			source: "-!-5",
			expected: &expr.Unary{
				Operator: token.Token{Type: token.MINUS, Lexeme: "-", Line: 1},
				Right: &expr.Unary{
					Operator: token.Token{Type: token.BANG, Lexeme: "!", Line: 1},
					Right: &expr.Unary{
						Operator: token.Token{Type: token.MINUS, Lexeme: "-", Line: 1},
						Right:    &expr.Literal{Value: float64(5)},
					},
				},
			},
		},
		{
			name:   "Mixed Binary and Unary Operations",
			source: "-5 + 3",
			expected: &expr.Binary{
				Left: &expr.Unary{
					Operator: token.Token{Type: token.MINUS, Lexeme: "-", Line: 1},
					Right:    &expr.Literal{Value: float64(5)},
				},
				Operator: token.Token{Type: token.PLUS, Lexeme: "+", Line: 1},
				Right:    &expr.Literal{Value: float64(3)},
			},
		},
		{
			name:   "Grouping with Mixed Operations",
			source: "-(5 + 3)",
			expected: &expr.Unary{
				Operator: token.Token{Type: token.MINUS, Lexeme: "-", Line: 1},
				Right: &expr.Grouping{
					Expression: &expr.Binary{
						Left:     &expr.Literal{Value: float64(5)},
						Operator: token.Token{Type: token.PLUS, Lexeme: "+", Line: 1},
						Right:    &expr.Literal{Value: float64(3)},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokens := scanner.NewScanner(test.source).ScanTokens()
			parser := NewParser(tokens)
			result, err := parser.Parse()

			if test.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none for source: %v", test.source)
					return
				}
			} else if err != nil {
				t.Errorf("Didn't expect error for source: %v but got %v", test.source, err)
				return
			}

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
