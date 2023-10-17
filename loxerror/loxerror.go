package loxerror

import (
	"fmt"
	"os"

	"github.com/joshbochu/lox-go/interpreter"
	"github.com/joshbochu/lox-go/token"
)

// LoxError is the global instance of the ErrorHandler.
var LoxError = &ErrorHandler{}

type ErrorHandler struct {
	HadError        bool
	HadRuntimeError bool
}

func (e *ErrorHandler) report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	e.HadError = true
}

// ErrorLine reports an error on a specific line.
func ErrorLine(line int, message string) {
	LoxError.report(line, "", message)
}

// ErrorToken reports an error at a specific token.
func ErrorToken(t token.Token, message string) {
	if t.Type == token.EOF {
		LoxError.report(t.Line, " at end", message)
	} else {
		LoxError.report(t.Line, " at '"+t.Lexeme+"'", message)
	}
}

// ErrorRuntime sets the HadError flag for runtime errors.
func ErrorRuntime(error interpreter.RuntimeError) {
	fmt.Fprintf(os.Stderr, "%s\n[line %d]", error.Message, error.Token.Line)
	LoxError.HadError = true
}
