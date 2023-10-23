package main

import (
	"fmt"
	"os"
	"os/exec"
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
		"Binary   : Expr Left, token.Token Operator, Expr Right",
		"Grouping : Expr Expression",
		"Literal  : Object Value",
		"Unary    : token.Token Operator, Expr Right",
	})
	defineAst(outputDir, "Stmt", []string{
		"Expression : expr.Expr expression",
		"Print : expr.Expr expression",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating AST definition: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("AST definition generated successfully to directory %s\n", outputDir)
}

func defineAst(baseOutputDir string, baseName string, types []string) error {
	// Create directory for the file (e.g., "expr" or "stmt")
	dirPath := filepath.Join(baseOutputDir, strings.ToLower(baseName))
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Create file
	path := filepath.Join(dirPath, strings.ToLower(baseName)+".go")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	var builder strings.Builder

	imports := []string{"github.com/joshbochu/golox/token", "github.com/joshbochu/golox/expr"}

	// Extract package name from the directory's base name
	packageName := filepath.Base(dirPath)
	builder.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	builder.WriteString("import (\n")
	for _, pkg := range imports {
		builder.WriteString(fmt.Sprintf("\t\"%s\"\n", pkg))
	}
	builder.WriteString(")\n\n")

	builder.WriteString(fmt.Sprintf("type %s interface {\n", baseName))
	builder.WriteString(fmt.Sprintf("\tAccept(visitor %sVisitor) (interface{}, error)\n", baseName))
	builder.WriteString("}\n\n")

	// ExprVisitor Interface
	builder.WriteString(fmt.Sprintf("type %sVisitor interface {\n", baseName))
	for _, typeDef := range types {
		parts := strings.SplitN(typeDef, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid type definition %s", typeDef)
		}
		typeName := strings.TrimSpace(parts[0])
		builder.WriteString(fmt.Sprintf("\tVisit%s%s(expr *%s) (interface{}, error)\n", typeName, baseName, typeName))
	}
	builder.WriteString("}\n\n")

	for _, typeDef := range types {
		parts := strings.SplitN(typeDef, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid type definition %s", typeDef)
		}
		typeName := strings.TrimSpace(parts[0])
		builder.WriteString(fmt.Sprintf("type %s struct {\n", typeName))

		fields := strings.Split(strings.TrimSpace(parts[1]), ",")
		for _, field := range fields {
			field = strings.TrimSpace(field)
			fieldParts := strings.Split(field, " ")
			fieldType := fieldParts[0]
			if fieldType == "Object" {
				fieldType = "interface{}"
			}
			fieldName := fieldParts[1]
			builder.WriteString(fmt.Sprintf("\t%s %s\n", fieldName, fieldType))
		}
		builder.WriteString("}\n\n")
		builder.WriteString(fmt.Sprintf("func (e *%s) Accept(visitor %sVisitor) (interface{}, error) {\n", typeName, baseName))
		builder.WriteString(fmt.Sprintf("\tval, err := visitor.Visit%s%s(e)\n", typeName, baseName))
		builder.WriteString("\tif err != nil {\n\t\treturn nil, err\n\t}\n")
		builder.WriteString("\treturn val, nil\n")
		builder.WriteString("}\n\n")
	}

	_, err = file.WriteString(builder.String())
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	cmd := exec.Command("go", "fmt", path)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to format file: %v", err)
	}

	return nil
}
