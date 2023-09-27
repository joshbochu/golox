package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joshbochu/lox-go/pkg/scanner"
)

func main() {
	switch len(os.Args) {
	case 1: // "./main"
		runPrompt()
	case 2: // "./main fileName"
		runFile(os.Args[1])
	default: // "./main fileName extra1 extra2 ..."
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
