package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joshbochu/lox-go/astprinter"
	"github.com/joshbochu/lox-go/interpreter"
	"github.com/joshbochu/lox-go/loxerror"
	"github.com/joshbochu/lox-go/parser"
	"github.com/joshbochu/lox-go/scanner"
)

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
	if loxerror.LoxError.HadError {
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
		if loxerror.LoxError.HadError {
			loxerror.LoxError.HadError = false
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
	parser := parser.NewParser(tokens)
	expression, err := parser.Parse()
	if err != nil && loxerror.LoxError.HadError {
		return
	}
	printer := &astprinter.Printer{}
	fmt.Println(expression.Accept(printer))

	interpreter := &interpreter.Interpreter{}
	interpreter.Interpret(expression)
}
