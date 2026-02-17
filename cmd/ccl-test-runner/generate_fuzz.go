package main

import (
	"fmt"

	"github.com/catconflang/ccl-test-data/fuzz"
	"github.com/catconflang/ccl-test-data/internal/styles"
	"github.com/urfave/cli/v2"
)

func generateFuzzAction(ctx *cli.Context) error {
	seed := ctx.Int64("seed")
	count := ctx.Int("count")
	outputDir := ctx.String("output")
	validate := ctx.Bool("validate")

	styles.Status("🎲", fmt.Sprintf("Generating %d fuzz tests (seed: %d)...", count, seed))

	gen := fuzz.NewGenerator(fuzz.GeneratorOptions{
		Seed:  seed,
		Count: count,
	})

	suite, err := gen.Generate()
	if err != nil {
		return fmt.Errorf("error generating fuzz tests: %w", err)
	}

	if validate {
		styles.InfoLite("Validating generated tests against mock implementation...")
		errors := gen.Validate(suite)
		for _, e := range errors {
			styles.Warning("  %s", e)
		}
		if len(errors) > 0 {
			return fmt.Errorf("%d validation errors found", len(errors))
		}
		styles.InfoLite("All %d tests validated successfully", len(suite.Tests))
	}

	if err := gen.WriteToFile(suite, outputDir); err != nil {
		return fmt.Errorf("error writing fuzz tests: %w", err)
	}

	styles.Success("✅ Generated %d fuzz tests in %s", len(suite.Tests), outputDir)
	return nil
}
