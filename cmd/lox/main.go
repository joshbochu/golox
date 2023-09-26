package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 { // ./lox fileName extraStuff
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 { // ./lox fileName
		// runFile()
	} else { // ./lox (interactive mode)
		// runPrompt()
	}
}
