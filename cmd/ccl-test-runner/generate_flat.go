package main

import (
	"fmt"

	"github.com/ccl-test-data/test-runner/internal/styles"
	"github.com/tylerbu/ccl-test-lib/generator"
	"github.com/urfave/cli/v2"
)

func generateFlatAction(ctx *cli.Context) error {
	sourceDir := ctx.String("source")
	generatedDir := ctx.String("generated")

	styles.Status("⚡", fmt.Sprintf("Generating flat tests from %s to %s...", sourceDir, generatedDir))

	// Create ccl-test-lib generator with compact format support
	flatGen := generator.NewFlatGenerator(sourceDir, generatedDir, generator.GenerateOptions{
		Verbose:           true,
		SourceFormat:      generator.FormatCompact, // Use compact format
		SkipPropertyTests: false,
	})

	// Generate all flat format files
	err := flatGen.GenerateAll()
	if err != nil {
		return fmt.Errorf("error generating flat tests: %w", err)
	}

	styles.Success("✅ Flat test generation completed successfully!")
	styles.InfoLite("Implementation-friendly tests saved in: %s", generatedDir)
	return nil
}
