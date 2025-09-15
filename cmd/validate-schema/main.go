package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/validate-schema/main.go <directory-or-files...>")
		os.Exit(1)
	}

	// Load schema
	schemaLoader := gojsonschema.NewReferenceLoader("file://tests/schema.json")

	var testFiles []string

	for _, arg := range os.Args[1:] {
		info, err := os.Stat(arg)
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", arg, err)
			continue
		}

		if info.IsDir() {
			// Find all api_*.json files in the directory (excluding schema.json)
			err := filepath.Walk(arg, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.HasPrefix(filepath.Base(path), "api_") &&
					strings.HasSuffix(path, ".json") && filepath.Base(path) != "schema.json" {
					testFiles = append(testFiles, path)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Error walking directory %s: %v\n", arg, err)
			}
		} else {
			// It's a file, add it directly
			testFiles = append(testFiles, arg)
		}
	}

	for _, testFile := range testFiles {
		// Load test file
		testLoader := gojsonschema.NewReferenceLoader("file://" + testFile)

		// Validate
		result, err := gojsonschema.Validate(schemaLoader, testLoader)
		if err != nil {
			fmt.Printf("✗ %s - Error: %v\n", testFile, err)
			continue
		}

		if result.Valid() {
			fmt.Printf("✓ %s - Valid\n", testFile)
		} else {
			fmt.Printf("✗ %s - Invalid:\n", testFile)
			for _, desc := range result.Errors() {
				fmt.Printf("  - %s\n", desc)
			}
		}
	}
}
