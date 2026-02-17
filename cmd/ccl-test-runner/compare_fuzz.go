package main

import (
	"fmt"

	"github.com/catconflang/ccl-test-data/fuzz"
	"github.com/catconflang/ccl-test-data/internal/styles"
	"github.com/urfave/cli/v2"
)

func compareFuzzAction(ctx *cli.Context) error {
	seedA := ctx.Int64("seed-a")
	seedB := ctx.Int64("seed-b")
	count := ctx.Int("count")

	styles.Status("🔍", fmt.Sprintf("Comparing fuzz datasets: seed %d vs seed %d (%d tests each)...", seedA, seedB, count))

	report, err := fuzz.Compare(seedA, seedB, count)
	if err != nil {
		return fmt.Errorf("comparison failed: %w", err)
	}

	fmt.Println()
	fmt.Print(report.FormatReport())

	return nil
}
