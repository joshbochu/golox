package scanner

type Token string

type Scanner struct {
	source string
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
	}
}

func (s *Scanner) ScanTokens() []Token {
	return []Token{"TOKEN"}
}
