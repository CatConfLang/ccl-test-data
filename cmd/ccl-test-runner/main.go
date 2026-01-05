package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tylerbutler/ccl-test-data/internal/benchmark"
	"github.com/tylerbutler/ccl-test-data/internal/config"
	"github.com/tylerbutler/ccl-test-data/internal/generator"
	"github.com/tylerbutler/ccl-test-data/internal/stats"
	"github.com/tylerbutler/ccl-test-data/internal/styles"
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
		Usage: "Generate and run Go tests from flat JSON test data",
		Description: `A test runner for CCL (Categorical Configuration Language) implementations.
		
This tool consumes flat JSON test files and generates corresponding Go test files
with proper organization by function and feature. Provides enhanced output formatting
for both human and machine consumption.`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"gen", "g"},
				Usage:   "Generate Go test files from flat JSON test data",
				Description: `Generate Go test files from flat JSON test suite data.
				
This command reads flat JSON test files and generates corresponding Go test files
with proper organization by function and feature. Uses configuration-based filtering
to exclude tests incompatible with implementation choices.`,
				Action: generateAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Value:   "generated_tests",
						Usage:   "Input directory containing flat JSON test files",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "go_tests",
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
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "Filter by tags (not yet implemented)",
					},
					&cli.StringSliceFlag{
						Name:  "features",
						Usage: "Filter by features (comments, parsing, objects, etc)",
					},
					&cli.StringSliceFlag{
						Name:  "skip",
						Usage: "Skip specific tests by name pattern (e.g., --skip TestKeyWithNewlineBeforeEqualsParse)",
					},
					&cli.BoolFlag{
						Name:  "basic-only",
						Usage: "Run only basic tests, skipping known failing edge cases",
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
				Description: `Analyze flat JSON test files and display comprehensive statistics.
				
This command scans flat JSON test files and provides detailed statistics including
test counts, assertion counts, and categorization by feature areas.`,
				Action: statsAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Value:   "generated_tests",
						Usage:   "Input directory containing flat JSON test files",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Value:   "pretty",
						Usage:   "Output format (pretty, json)",
					},
				},
			},
			{
				Name:    "benchmark",
				Aliases: []string{"bench", "b"},
				Usage:   "Run performance benchmarks on core operations",
				Description: `Run performance benchmarks to measure and track the performance of 
core operations including test generation, statistics collection, and parsing.

This command measures execution time and memory usage, comparing against historical 
results to detect performance regressions.`,
				Action: benchmarkAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Value:   "generated_tests",
						Usage:   "Input directory containing flat JSON test files",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "go_tests",
						Usage:   "Output directory for generated test files",
					},
					&cli.StringFlag{
						Name:    "results",
						Aliases: []string{"r"},
						Value:   "benchmarks/results.json",
						Usage:   "File to save benchmark results",
					},
					&cli.StringFlag{
						Name:    "compare",
						Aliases: []string{"c"},
						Usage:   "Historical results file to compare against",
					},
					&cli.Float64Flag{
						Name:  "threshold",
						Value: 10.0,
						Usage: "Regression threshold percentage (default: 10%)",
					},
				},
			},
			{
				Name:    "generate-flat",
				Aliases: []string{"flat"},
				Usage:   "Generate flat format tests from source format (delegates to ccl-test-lib)",
				Description: `Generate implementation-friendly flat format tests from maintainable 
source format tests. This command is a thin CLI wrapper around ccl-test-lib.

ARCHITECTURE: The actual conversion logic lives in ccl-test-lib/generator where it belongs.
This CLI command provides convenient access while maintaining proper separation of concerns.

Each source test with multiple validations becomes multiple flat tests (one per validation),
creating a simple, uniform format that's easy for test runners to process.`,
				Action: generateFlatAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "source",
						Aliases: []string{"s"},
						Value:   "source_tests",
						Usage:   "Source directory with source format tests",
					},
					&cli.StringFlag{
						Name:    "generated",
						Aliases: []string{"g"},
						Value:   "generated_tests",
						Usage:   "Output directory for flat format tests",
					},
					&cli.StringFlag{
						Name:  "schemas",
						Value: "schemas",
						Usage: "Directory containing source-format.json with behavior metadata for filtering",
					},
					&cli.BoolFlag{
						Name:  "auto-conflicts",
						Value: true,
						Usage: "Auto-generate behavior conflicts from metadata (default: true)",
					},
					&cli.BoolFlag{
						Name:  "validate",
						Value: false,
						Usage: "Validate source tests against behavior metadata",
					},
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Value:   true,
						Usage:   "Verbose output",
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

	// Load centralized configuration with validation
	cfg := config.DefaultConfig()

	// Override test filtering settings from CLI flags
	cfg.TestFiltering.SkipDisabled = skipDisabled
	cfg.TestFiltering.SkipTags = skipTags
	cfg.TestFiltering.RunOnlyFunctions = runOnly

	// Validate configuration (will error if required choices aren't made)
	if err := cfg.Validate(); err != nil {
		styles.Error("Configuration validation failed:")
		styles.Error("  %v", err)
		styles.Info("\nTo fix this, all mutually exclusive behavioral choices must be explicitly configured.")
		styles.Info("Check internal/config/runner_config.go DefaultConfig() for required settings.")
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Create generator with validated configuration
	gen, err := generator.NewWithConfig(inputDir, outputDir, cfg)
	if err != nil {
		return fmt.Errorf("failed to create generator: %w", err)
	}

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
	tags := ctx.StringSlice("tags")
	features := ctx.StringSlice("features")
	skipTests := ctx.StringSlice("skip")
	basicOnly := ctx.Bool("basic-only")
	listOnly := ctx.Bool("list")
	verbose := ctx.Bool("verbose")

	if verbose {
		format = "verbose"
	}

	// Add basic-only exclusions
	if basicOnly {
		defaultSkipTests := []string{
			"TestKeyWithNewlineBeforeEqualsParse",
			"TestComplexMultiNewlineWhitespaceParse",
			"TestRoundTripWhitespaceNormalizationParse",
		}
		skipTests = append(skipTests, defaultSkipTests...)
	}

	packages := buildPackagePatterns(tags, features)

	if listOnly {
		styles.Info("📋 Available test packages:")
		if len(packages) == 0 {
			// List all packages
			if matches, err := filepath.Glob("./go_tests/*"); err == nil {
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
	return runTestsWithGotestsum(format, tags, features, skipTests, ctx.Args().Slice())
}

func statsAction(ctx *cli.Context) error {
	inputDir := ctx.String("input")
	format := ctx.String("format")

	// Only show status message for pretty format
	if format != "json" {
		styles.Status("📊", "Collecting test statistics...")
	}

	collector := stats.NewEnhancedCollector(inputDir)
	statistics, err := collector.CollectEnhancedStats()
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
		// Use the enhanced stats printer
		stats.PrintEnhancedStats(statistics)
	}

	return nil
}

func benchmarkAction(ctx *cli.Context) error {
	inputDir := ctx.String("input")
	outputDir := ctx.String("output")
	resultsFile := ctx.String("results")
	compareFile := ctx.String("compare")
	threshold := ctx.Float64("threshold")

	styles.Status("🚀", "Running performance benchmarks...")

	// Create benchmark tracker
	tracker := benchmark.NewTracker()

	// Benchmark 1: Test Generation
	tracker.StartBenchmark("test-generation")
	gen := generator.New(inputDir, outputDir)
	if err := gen.GenerateAll(); err != nil {
		return fmt.Errorf("benchmark failed during test generation: %w", err)
	}
	genResult := tracker.EndBenchmark("test-generation")

	// Benchmark 2: Statistics Collection
	tracker.StartBenchmark("stats-collection")
	collector := stats.NewEnhancedCollector(inputDir)
	if _, err := collector.CollectEnhancedStats(); err != nil {
		return fmt.Errorf("benchmark failed during stats collection: %w", err)
	}
	statsResult := tracker.EndBenchmark("stats-collection")

	// Display results
	results := tracker.GetAllResults()
	benchmark.PrintResults(results)

	// Save results
	if err := os.MkdirAll(filepath.Dir(resultsFile), 0755); err != nil {
		return fmt.Errorf("failed to create benchmark results directory: %w", err)
	}

	if err := tracker.SaveResults(resultsFile); err != nil {
		return fmt.Errorf("failed to save benchmark results: %w", err)
	}

	styles.Success("✅ Benchmark results saved to %s", resultsFile)

	// Compare with historical results if provided
	if compareFile != "" {
		if historical, err := benchmark.LoadResults(compareFile); err == nil {
			alerts := benchmark.CompareResults(results, historical, threshold)
			benchmark.PrintRegressionAlerts(alerts)

			if len(alerts) > 0 {
				styles.Warning("⚠️  Performance regressions detected!")
				return fmt.Errorf("performance regression threshold exceeded")
			}
		} else {
			styles.Warning("⚠️  Could not load historical results from %s: %v", compareFile, err)
		}
	}

	styles.InfoLite("Test Generation: %v (%d bytes allocated)",
		genResult.Duration, genResult.MemAllocBytes)
	styles.InfoLite("Stats Collection: %v (%d bytes allocated)",
		statsResult.Duration, statsResult.MemAllocBytes)

	return nil
}

func runTestsWithGotestsum(format string, tags []string, features []string, skipTests []string, extraArgs []string) error {
	// Check if gotestsum is available
	if _, err := exec.LookPath("gotestsum"); err != nil {
		styles.Warning("⚠️  gotestsum not found, falling back to go test")
		return runWithGoTest(tags, features, skipTests, extraArgs)
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
	packages := buildPackagePatterns(tags, features)

	// Add separator for go test args
	cmd.Args = append(cmd.Args, "--")

	// Add skip patterns if specified
	if len(skipTests) > 0 {
		skipPattern := strings.Join(skipTests, "|")
		cmd.Args = append(cmd.Args, "-skip", skipPattern)
	}

	// Add package patterns
	if len(packages) == 0 {
		cmd.Args = append(cmd.Args, "./go_tests/...")
	} else {
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

func runWithGoTest(tags []string, features []string, skipTests []string, extraArgs []string) error {
	cmd := exec.Command("go", "test")

	// Add skip patterns if specified
	if len(skipTests) > 0 {
		skipPattern := strings.Join(skipTests, "|")
		cmd.Args = append(cmd.Args, "-skip", skipPattern)
	}

	// Add package patterns
	packages := buildPackagePatterns(tags, features)

	if len(packages) == 0 {
		cmd.Args = append(cmd.Args, "./go_tests/...")
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

func buildPackagePatterns(tags []string, features []string) []string {
	var packages []string

	// If features are specified, filter by feature names
	if len(features) > 0 {
		for _, feature := range features {
			pattern := fmt.Sprintf("go_tests/*%s*", feature)
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
