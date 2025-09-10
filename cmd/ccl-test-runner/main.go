package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ccl-test-data/test-runner/internal/generator"
	"github.com/ccl-test-data/test-runner/internal/stats"
	"github.com/ccl-test-data/test-runner/internal/styles"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	app := &cli.App{
		Name:  "ccl-test-runner",
		Usage: "Generate and run tests for CCL implementations",
		Description: `A test runner for CCL (Categorical Configuration Language) implementations.
		
This tool generates Go test files from JSON test data and provides enhanced
output formatting for both human and machine consumption.`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"gen", "g"},
				Usage:   "Generate Go test files from JSON test data",
				Description: `Generate Go test files from the JSON test suite data.
				
This command reads the JSON test files and generates corresponding Go test files
with proper organization by level and feature.`,
				Action: generateAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Value:   "tests",
						Usage:   "Input directory containing JSON test files",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "generated_tests",
						Usage:   "Output directory for generated test files",
					},
					&cli.BoolFlag{
						Name:  "skip-disabled",
						Value: true,
						Usage: "Skip tests with tags that indicate disabled features (default: true)",
					},
					&cli.StringSliceFlag{
						Name:  "skip-tags",
						Usage: "Additional tags to skip (e.g., --skip-tags multiline,error)",
					},
					&cli.StringSliceFlag{
						Name:  "run-only",
						Usage: "Only generate tests with these tags (overrides skip behavior)",
					},
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Run generated tests with enhanced output",
				Description: `Run the generated Go tests with beautiful, human-readable output.
				
This command uses gotestsum to provide enhanced formatting while maintaining
compatibility with standard go test flags.`,
				Action: testAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Value:   "pretty",
						Usage:   "Output format (pretty, table, verbose)",
					},
					&cli.IntSliceFlag{
						Name:    "levels",
						Aliases: []string{"l"},
						Usage:   "Filter by CCL levels (1,2,3,4)",
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "Filter by tags (not yet implemented)",
					},
					&cli.StringSliceFlag{
						Name:  "features",
						Usage: "Filter by features (comments, parsing, objects, etc)",
					},
					&cli.BoolFlag{
						Name:  "list",
						Usage: "List available test packages without running",
					},
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "Verbose output (same as --format verbose)",
					},
				},
			},
			{
				Name:    "stats",
				Aliases: []string{"statistics", "s"},
				Usage:   "Collect and display test suite statistics",
				Description: `Analyze test files and display comprehensive statistics about the test suite.
				
This command scans all JSON test files and provides detailed statistics including
test counts, assertion counts, and categorization by feature areas.`,
				Action: statsAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Value:   "tests",
						Usage:   "Input directory containing JSON test files",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Value:   "pretty",
						Usage:   "Output format (pretty, json)",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		styles.Error("Error: %v", err)
		os.Exit(1)
	}
}

// Action functions for CLI commands
func generateAction(ctx *cli.Context) error {
	styles.Status("🚀", "Generating test files...")

	inputDir := ctx.String("input")
	outputDir := ctx.String("output")
	skipDisabled := ctx.Bool("skip-disabled")
	skipTags := ctx.StringSlice("skip-tags")
	runOnly := ctx.StringSlice("run-only")

	// Create generator options
	options := generator.Options{
		SkipDisabled: skipDisabled,
		SkipTags:     skipTags,
		RunOnly:      runOnly,
	}

	gen := generator.NewWithOptions(inputDir, outputDir, options)
	if err := gen.GenerateAll(); err != nil {
		return fmt.Errorf("failed to generate tests: %w", err)
	}

	// Display assertion statistics
	stats := gen.GetStats()
	totalTests := stats.TotalTests
	activeTests := totalTests - stats.SkippedTests

	styles.InfoLite("Generated %d tests with %d total assertions", totalTests, stats.TotalAssertions+stats.SkippedAssertions)
	styles.InfoLite("Active tests: %d (with %d assertions)", activeTests, stats.TotalAssertions)
	if stats.SkippedTests > 0 {
		styles.InfoLite("Skipped tests: %d (with %d assertions)", stats.SkippedTests, stats.SkippedAssertions)
	}

	styles.Success("✅ Test generation completed successfully")
	return nil
}

func testAction(ctx *cli.Context) error {
	format := ctx.String("format")
	levels := ctx.IntSlice("levels")
	tags := ctx.StringSlice("tags")
	features := ctx.StringSlice("features")
	listOnly := ctx.Bool("list")
	verbose := ctx.Bool("verbose")

	if verbose {
		format = "verbose"
	}

	packages := buildPackagePatterns(levels, tags, features)

	if listOnly {
		styles.Info("📋 Available test packages:")
		if len(packages) == 0 {
			// List all packages
			if matches, err := filepath.Glob("./generated_tests/*"); err == nil {
				for _, match := range matches {
					styles.InfoLite("  %s", match)
				}
			}
		} else {
			for _, pkg := range packages {
				styles.InfoLite("  %s", pkg)
			}
		}
		return nil
	}

	styles.Status("🧪", "Running tests...")
	return runTestsWithGotestsum(format, levels, tags, features, ctx.Args().Slice())
}

func statsAction(ctx *cli.Context) error {
	inputDir := ctx.String("input")
	format := ctx.String("format")

	// Only show status message for pretty format
	if format != "json" {
		styles.Status("📊", "Collecting test statistics...")
	}

	collector := stats.NewCollector(inputDir)
	statistics, err := collector.CollectStats()
	if err != nil {
		return fmt.Errorf("failed to collect statistics: %w", err)
	}

	switch format {
	case "json":
		// Output pure JSON for machine consumption (no status messages)
		jsonData, err := json.MarshalIndent(statistics, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal statistics: %w", err)
		}
		fmt.Println(string(jsonData))

	default: // "pretty"
		// Display human-readable summary
		styles.Info("🔍 CCL Test Suite Statistics")
		styles.InfoLite("")

		styles.InfoLite("Feature-Based Structure:")
		styles.InfoLite("  Core Parsing: %d tests (%d assertions)",
			statistics.Categories["core-parsing"].Total,
			statistics.Categories["core-parsing"].Assertions)
		styles.InfoLite("  Advanced Processing: %d tests (%d assertions)",
			statistics.Categories["advanced-processing"].Total,
			statistics.Categories["advanced-processing"].Assertions)
		styles.InfoLite("  Object Construction: %d tests (%d assertions)",
			statistics.Categories["object-construction"].Total,
			statistics.Categories["object-construction"].Assertions)
		styles.InfoLite("  Type System: %d tests (%d assertions)",
			statistics.Categories["type-system"].Total,
			statistics.Categories["type-system"].Assertions)
		styles.InfoLite("  Output & Validation: %d tests (%d assertions)",
			statistics.Categories["output-validation"].Total,
			statistics.Categories["output-validation"].Assertions)

		if statistics.Categories["other"].Total > 0 {
			styles.InfoLite("  Other: %d tests (%d assertions)",
				statistics.Categories["other"].Total,
				statistics.Categories["other"].Assertions)
		}

		styles.InfoLite("")
		styles.Success("Total: %d tests (%d assertions) across %d files",
			statistics.TotalTests, statistics.TotalAssertions, statistics.TotalFiles)
	}

	return nil
}

func runTestsWithGotestsum(format string, levels []int, tags []string, features []string, extraArgs []string) error {
	// Check if gotestsum is available
	if _, err := exec.LookPath("gotestsum"); err != nil {
		styles.Warning("⚠️  gotestsum not found, falling back to go test")
		return runWithGoTest(levels, tags, features, extraArgs)
	}

	// Build gotestsum command
	cmd := exec.Command("gotestsum")

	// Add format flag
	switch format {
	case "json":
		cmd.Args = append(cmd.Args, "--format", "standard-verbose")
		cmd.Args = append(cmd.Args, "--jsonfile", "/dev/stdout")
	case "table":
		cmd.Args = append(cmd.Args, "--format", "dots-v2")
	case "verbose":
		cmd.Args = append(cmd.Args, "--format", "standard-verbose")
	default: // pretty
		cmd.Args = append(cmd.Args, "--format", "testname")
	}

	// Add package patterns based on filters
	packages := buildPackagePatterns(levels, tags, features)

	if len(packages) == 0 {
		cmd.Args = append(cmd.Args, "--", "./generated_tests/...")
	} else {
		cmd.Args = append(cmd.Args, "--")
		cmd.Args = append(cmd.Args, packages...)
	}

	// Add extra args from user
	cmd.Args = append(cmd.Args, extraArgs...)

	// Set up command output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	styles.Command(strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func runWithGoTest(levels []int, tags []string, features []string, extraArgs []string) error {
	cmd := exec.Command("go", "test")

	// Add package patterns
	packages := buildPackagePatterns(levels, tags, features)

	if len(packages) == 0 {
		cmd.Args = append(cmd.Args, "./generated_tests/...")
	} else {
		cmd.Args = append(cmd.Args, packages...)
	}

	// Add extra args
	cmd.Args = append(cmd.Args, extraArgs...)

	// Set up command output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	styles.Command(strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func buildPackagePatterns(levels []int, tags []string, features []string) []string {
	var packages []string

	// If levels are specified, filter by level
	if len(levels) > 0 {
		for _, level := range levels {
			pattern := fmt.Sprintf("generated_tests/level%d*", level)
			if matches, err := filepath.Glob(pattern); err == nil {
				for _, match := range matches {
					// Ensure the package pattern is correct for go test
					packages = append(packages, "./"+match)
				}
			}
		}
	}

	// If features are specified, filter by feature names
	if len(features) > 0 {
		for _, feature := range features {
			pattern := fmt.Sprintf("generated_tests/*%s*", feature)
			if matches, err := filepath.Glob(pattern); err == nil {
				for _, match := range matches {
					packages = append(packages, "./"+match)
				}
			}
		}
	}

	// TODO: Tag filtering would require parsing test files or metadata
	// For now, we'll just warn the user
	if len(tags) > 0 {
		styles.Warning("⚠️  Tag filtering not yet implemented")
	}

	return packages
}
