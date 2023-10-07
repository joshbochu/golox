package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		programName := os.Args[0]
		fmt.Fprintf(os.Stderr, "Usage: %s <output_directory>\n", programName)
		os.Exit(64)
	}
	outputDir := os.Args[1]
	err := defineAst(outputDir, "Expr", []string{
		"Binary   : Expr left, Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : Object value",
		"Unary    : Token operator, Expr right",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating AST definition: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("AST definition generated successfully to directory %s\n", outputDir)
}

func defineAst(outputDir string, baseName string, types []string) error {
	// Create file
	path := filepath.Join(outputDir, strings.ToLower(baseName)+".go")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	var builder strings.Builder

	// Define package
	builder.WriteString("package ast\n")
	builder.WriteString("import (\n")
	// TODO add imports
	builder.WriteString(")\n")

	builder.WriteString(fmt.Sprintf("type %s interface {\n", baseName))
	// TODO add interface definition
	builder.WriteString("}\n")

	for _, typeDef := range types {
		builder.WriteString(fmt.Sprintf("type %s%s struct {\n", typeDef, baseName))
		parts := strings.SplitN(typeDef, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid type definition %s", typeDef)
		}
		name := strings.TrimSpace(parts[0])
		// fields := strings.TrimSpace(parts[1])
		builder.WriteString("}\n")
	}

	_, err = file.WriteString(builder.String())
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}
