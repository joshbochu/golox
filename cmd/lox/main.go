package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joshbochu/lox-go/pkg/scanner"
)

var hadError = false

func main() {
	switch len(os.Args) {
	case 1: // "./main"
		runPrompt()
	case 2: // "./main fileName"
		runFile(os.Args[1])
	default: // "./main fileName ..."
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
		os.Exit(65)
	}
	source := string(bytes)
	run(source)
	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		run(line)
		if hadError {
			hadError = false
		}
		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while processing input: %v\n", err)
	}
}

func run(source string) {
	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	hadError = true
}
