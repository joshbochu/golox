package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		programName := os.Args[0]
		fmt.Fprintf(os.Stderr, "Usage: %s <output_directory>\n", programName)
		os.Exit(64)
	}
	outputDir := os.Args[1]
	// TODO Actaual AST Generation Stuff
	fmt.Printf("AST generated successfully to directory %s\n", outputDir)
}
