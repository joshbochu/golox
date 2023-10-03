package util

import (
	"fmt"
	"os"
)

var hadError bool = false

func SetHadError(value bool) {
	hadError = value
}

func GetHadError() bool {
	return hadError
}

func Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	SetHadError(true)
}
