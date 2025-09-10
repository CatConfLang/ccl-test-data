package main

import (
	"fmt"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/validate-schema/main.go <test-file>")
		os.Exit(1)
	}

	// Load schema
	schemaLoader := gojsonschema.NewReferenceLoader("file://tests/schema.json")

	for _, testFile := range os.Args[1:] {
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
