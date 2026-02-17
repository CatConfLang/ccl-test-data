package main

import (
	"fmt"
	"strings"

	"github.com/catconflang/ccl-test-data/fuzz"
	"github.com/catconflang/ccl-test-data/internal/styles"
	"github.com/urfave/cli/v2"
)

func compareFuzzAction(ctx *cli.Context) error {
	seeds := ctx.Int64Slice("seeds")
	count := ctx.Int("count")

	if len(seeds) < 2 {
		return fmt.Errorf("need at least 2 seeds (got %d), e.g.: --seeds 42 --seeds 99", len(seeds))
	}

	seedStrs := make([]string, len(seeds))
	for i, s := range seeds {
		seedStrs[i] = fmt.Sprintf("%d", s)
	}
	styles.Status("🔍", fmt.Sprintf("Comparing %d fuzz datasets [seeds: %s] (%d tests each)...",
		len(seeds), strings.Join(seedStrs, ", "), count))

	report, err := fuzz.Compare(seeds, count)
	if err != nil {
		return fmt.Errorf("comparison failed: %w", err)
	}

	fmt.Println()
	fmt.Print(report.FormatReport())

	return nil
}
