package util

import (
	"fmt"
	"os"

	"github.com/joshbochu/lox-go/token"
)

var hadError bool = false

func SetHadError(value bool) {
	hadError = value
}

func GetHadError() bool {
	return hadError
}

func ErrorLine(line int, message string) {
	report(line, "", message)
}

func ErrorToken(t token.Token, message string) {
	if t.Type == token.EOF {
		report(t.Line, " at end", message)
	} else {
		report(t.Line, " at '"+t.Lexeme+"'", message)
	}
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	SetHadError(true)
}
