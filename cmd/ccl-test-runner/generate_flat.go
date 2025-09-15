package main

import (
	"fmt"

	"github.com/ccl-test-data/test-runner/internal/generator"
	"github.com/ccl-test-data/test-runner/internal/styles"
	"github.com/urfave/cli/v2"
)

func generateFlatAction(ctx *cli.Context) error {
	sourceDir := ctx.String("source")
	generatedDir := ctx.String("generated")

	styles.Status("⚡", fmt.Sprintf("Generating flat tests from %s to %s...", sourceDir, generatedDir))

	// Load source format tests
	sourceTests, err := generator.LoadSourceTests(sourceDir)
	if err != nil {
		return fmt.Errorf("error loading source tests: %w", err)
	}

	styles.InfoLite("Loaded %d source tests", len(sourceTests))

	// Generate flat format tests
	flatTests, err := generator.GenerateFlatTests(sourceTests)
	if err != nil {
		return fmt.Errorf("error generating flat tests: %w", err)
	}

	styles.InfoLite("Generated %d flat tests", len(flatTests))

	// Save flat format tests
	err = generator.SaveFlatTests(flatTests, generatedDir)
	if err != nil {
		return fmt.Errorf("error saving flat tests: %w", err)
	}

	styles.Success("✅ Flat test generation completed successfully!")
	styles.InfoLite("Implementation-friendly tests saved in: %s", generatedDir)
	return nil
}
