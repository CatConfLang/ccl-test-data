package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/catconflang/ccl-test-data/fuzz"
	"github.com/catconflang/ccl-test-data/internal/styles"
	"github.com/urfave/cli/v2"
)

func compareFuzzAction(ctx *cli.Context) error {
	count := ctx.Int("count")

	// Parse seeds from positional args
	args := ctx.Args().Slice()
	if len(args) < 2 {
		return fmt.Errorf("need at least 2 seeds, e.g.: compare-fuzz 42 99 7")
	}

	seeds := make([]int64, len(args))
	for i, arg := range args {
		v, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid seed %q: %w", arg, err)
		}
		seeds[i] = v
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
