package scanner

import (
	"testing"
)

func TestScanner_isAtEnd(t *testing.T) {
	tests := []struct {
		input    string
		position int
		expected bool
	}{
		{"hello", 5, true},  // At the end
		{"hello", 4, false}, // Not at the end
		{"", 0, true},       // Empty string, should be at the end
	}

	for _, tt := range tests {
		scanner := NewScanner(tt.input)
		scanner.current = tt.position

		if result := scanner.isAtEnd(); result != tt.expected {
			t.Errorf("For input %q at position %d, expected isAtEnd() to be %v, but got %v", tt.input, tt.position, tt.expected, result)
		}
	}
}
