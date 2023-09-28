package scanner

import (
	"github.com/joshbochu/lox-go/pkg/token"
)

type Scanner struct {
	source string
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	return []token.Token{}
}
