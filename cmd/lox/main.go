package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joshbochu/lox-go/pkg/scanner"
	"github.com/joshbochu/lox-go/pkg/util"
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
	if util.GetHadError() {
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
		if util.GetHadError() {
			util.SetHadError(false)
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
