// Package main contains CLI commands for the CCL test runner.
// This file provides a thin wrapper around ccl-test-lib's flat generation functionality.
package main

import (
	"fmt"

	"github.com/CatConfLang/ccl-test-lib/generator"
	"github.com/ccl-test-data/test-runner/internal/styles"
	"github.com/urfave/cli/v2"
)

// generateFlatAction is a thin CLI wrapper around ccl-test-lib's flat generation functionality.
//
// ARCHITECTURE NOTE: This function exists as a convenience wrapper to maintain CLI consistency,
// but the actual flat generation logic lives in ccl-test-lib where it belongs. This follows
// the separation of concerns principle:
//
//   - ccl-test-lib: Contains the core flat generation logic
//   - ccl-test-runner: Provides CLI interface for convenient access
//
// The flat generation process converts source format tests (designed for maintainability)
// into flat format tests (designed for implementation simplicity). Each source test with
// multiple validations becomes multiple flat tests, one per validation.
//
// This delegation pattern allows:
//  1. Other projects to use ccl-test-lib directly without CLI dependencies
//  2. CLI users to access functionality conveniently
//  3. Logic to remain in the appropriate library layer
func generateFlatAction(ctx *cli.Context) error {
	sourceDir := ctx.String("source")
	generatedDir := ctx.String("generated")

	styles.Status("⚡", fmt.Sprintf("Generating flat tests from %s to %s...", sourceDir, generatedDir))

	// DELEGATION: Create ccl-test-lib generator with compact format support
	// The actual logic lives in ccl-test-lib/generator.NewFlatGenerator()
	// This CLI command simply provides a convenient interface to that functionality
	flatGen := generator.NewFlatGenerator(sourceDir, generatedDir, generator.GenerateOptions{
		Verbose:           true,
		SkipPropertyTests: false, // Include property-based tests in output
	})

	// DELEGATION: Execute the flat generation using ccl-test-lib
	// All conversion logic, validation, and file writing happens in the library
	err := flatGen.GenerateAll()
	if err != nil {
		return fmt.Errorf("error generating flat tests: %w", err)
	}

	styles.Success("✅ Flat test generation completed successfully!")
	styles.InfoLite("Implementation-friendly tests saved in: %s", generatedDir)
	styles.InfoLite("Note: Conversion logic provided by ccl-test-lib/generator")
	return nil
}
