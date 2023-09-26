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
		runFile(os.Args[1])
	} else { // ./lox (interactive mode)
		runPrompt()
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
		os.Exit(65)
	}
	run(string(bytes))
}

func runPrompt() {
	fmt.Println("Running in interactive mode")
}

func run(source string) {
	fmt.Println("Running source...")
}
