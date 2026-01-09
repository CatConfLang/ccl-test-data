// Package main contains CLI commands for the CCL test runner.
// This file provides a thin wrapper around ccl-test-lib's flat generation functionality.
package main

import (
	"fmt"
	"path/filepath"

	"github.com/catconflang/ccl-test-data/generator"
	"github.com/catconflang/ccl-test-data/internal/styles"
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
	schemasDir := ctx.String("schemas")
	autoConflicts := ctx.Bool("auto-conflicts")
	validate := ctx.Bool("validate")
	verbose := ctx.Bool("verbose")

	styles.Status("⚡", fmt.Sprintf("Generating flat tests from %s to %s...", sourceDir, generatedDir))

	// Resolve schemas directory relative to source directory if not absolute
	if schemasDir != "" && !filepath.IsAbs(schemasDir) {
		// Try relative to current directory first
		if _, err := filepath.Abs(schemasDir); err != nil {
			// Fall back to relative to source directory parent
			schemasDir = filepath.Join(filepath.Dir(sourceDir), schemasDir)
		}
	}

	// DELEGATION: Create ccl-test-lib generator with compact format support
	// The actual logic lives in ccl-test-lib/generator.NewFlatGenerator()
	// This CLI command simply provides a convenient interface to that functionality
	flatGen := generator.NewFlatGenerator(sourceDir, generatedDir, generator.GenerateOptions{
		Verbose:               verbose,
		SkipPropertyTests:     false, // Include property-based tests in output
		SchemasDir:            schemasDir,
		AutoGenerateConflicts: autoConflicts,
		ValidateSourceTests:   validate,
	})

	// Show metadata status
	if flatGen.BehaviorMetadata != nil {
		styles.InfoLite("Loaded behavior metadata with %d behaviors", len(flatGen.BehaviorMetadata.Behaviors))
		if autoConflicts {
			styles.InfoLite("Auto-generating conflicts from behavior metadata")
		}
		if validate {
			styles.InfoLite("Validating source tests against behavior metadata")
		}
	} else if schemasDir != "" {
		styles.Warning("Could not load behavior metadata from %s", schemasDir)
	}

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
