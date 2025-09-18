// Package generator provides test file generation capabilities for CCL implementations.
//
// This package transforms JSON test suites into executable Go test files, supporting
// feature-based tagging, object pooling for performance, and comprehensive assertion
// tracking. The generator organizes tests by CCL implementation levels (1-5) and
// features (parsing, comments, objects, etc.).
//
// Key Features:
//   - Feature-based test selection via structured tagging
//   - Object pooling to reduce memory allocations during generation
//   - Template-based Go test file generation with proper package organization
//   - Assertion counting and statistics collection for test suite analysis
//   - Support for mock implementation development with selective test generation
//
// Example Usage:
//
//	gen := generator.New("tests", "go_tests")
//	if err := gen.GenerateAll(); err != nil {
//	    log.Fatal(err)
//	}
//	stats := gen.GetStats()
//	fmt.Printf("Generated %d tests with %d assertions\n", stats.TotalTests, stats.TotalAssertions)
package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ccl-test-data/test-runner/internal/config"
	"github.com/ccl-test-data/test-runner/internal/styles"
	"github.com/tylerbu/ccl-test-lib/loader"
	"github.com/tylerbu/ccl-test-lib/types"
)

// Options configures test generation behavior
type Options struct {
	SkipDisabled    bool     // Skip tests with disabled feature tags
	SkipTags        []string // Additional tags to skip
	SkipTestsByName []string // Skip specific tests by name
	RunOnly         []string // Only run tests with these tags (overrides skip behavior)
}

// AssertionStats tracks assertion counts from test generation
type AssertionStats struct {
	TotalTests        int
	TotalAssertions   int
	SkippedTests      int
	SkippedAssertions int
	TestCounts        map[string]int // test name -> assertion count
}

// Generator handles test file generation from JSON test data
type Generator struct {
	inputDir  string
	outputDir string
	options   Options
	config    *config.RunnerConfig // Centralized configuration with behavioral choices
	stats     AssertionStats
	pool      *Pool // Object pool for memory optimization
}

// New creates a new generator instance with default options and configuration
func New(inputDir, outputDir string) *Generator {
	cfg := config.DefaultConfig()

	return &Generator{
		inputDir:  inputDir,
		outputDir: outputDir,
		options: Options{
			SkipDisabled: true, // Default behavior
		},
		config: cfg,
		stats: AssertionStats{
			TestCounts: make(map[string]int),
		},
		pool: NewPool(),
	}
}

// NewWithConfig creates a new generator instance with custom configuration
func NewWithConfig(inputDir, outputDir string, cfg *config.RunnerConfig) (*Generator, error) {
	// Validate configuration before using it
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &Generator{
		inputDir:  inputDir,
		outputDir: outputDir,
		options: Options{
			SkipDisabled:    cfg.TestFiltering.SkipDisabled,
			RunOnly:         cfg.TestFiltering.RunOnlyFunctions,
			SkipTags:        cfg.TestFiltering.SkipTags,
			SkipTestsByName: cfg.TestFiltering.SkipTestsByName,
		},
		config: cfg,
		stats: AssertionStats{
			TestCounts: make(map[string]int),
		},
		pool: NewPool(),
	}, nil
}

// NewWithOptions creates a new generator instance with custom options
func NewWithOptions(inputDir, outputDir string, options Options) *Generator {
	cfg := config.DefaultConfig()

	return &Generator{
		inputDir:  inputDir,
		outputDir: outputDir,
		options:   options,
		config:    cfg,
		stats: AssertionStats{
			TestCounts: make(map[string]int),
		},
		pool: NewPool(),
	}
}

// GetStats returns the assertion statistics
func (g *Generator) GetStats() AssertionStats {
	return g.stats
}

// GenerateAll generates test files for all JSON test suites
func (g *Generator) GenerateAll() error {
	// Ensure output directory exists
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Find all JSON test files
	testFiles, err := g.findTestFiles()
	if err != nil {
		return fmt.Errorf("failed to find test files: %w", err)
	}

	styles.InfoLite("Found %d test files to process", len(testFiles))

	// Process each test file
	for _, file := range testFiles {
		if err := g.generateTestFile(file); err != nil {
			return fmt.Errorf("failed to generate test file for %s: %w", file, err)
		}
		styles.FileProcessed(filepath.Base(file))
	}

	return nil
}

// findTestFiles discovers all JSON test files in the input directory
func (g *Generator) findTestFiles() ([]string, error) {
	var files []string

	err := filepath.WalkDir(g.inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, "schema.json") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// generateTestFile generates a Go test file from a flat format JSON test file
func (g *Generator) generateTestFile(jsonFile string) error {
	// Convert centralized config to ccl-test-lib format
	impl := g.config.ToImplementationConfig()

	// Add conflicting tags to skip list for behavior/variant filtering
	conflictingTags := g.config.GetConflictingTags()
	allSkipTags := append(g.options.SkipTags, conflictingTags...)

	// Create custom filter function for tag-based filtering
	customFilter := func(test types.TestCase) bool {
		// Check if any test tags are in the skip list
		for _, testTag := range test.Meta.Tags {
			for _, skipTag := range allSkipTags {
				if testTag == skipTag {
					return false // Skip this test
				}
			}
		}

		// If run-only tags are specified, check if test has any of them
		if len(g.options.RunOnly) > 0 {
			hasRunOnlyTag := false
			for _, testTag := range test.Meta.Tags {
				for _, runOnlyTag := range g.options.RunOnly {
					if testTag == runOnlyTag {
						hasRunOnlyTag = true
						break
					}
				}
				if hasRunOnlyTag {
					break
				}
			}
			if !hasRunOnlyTag {
				return false // Skip if no run-only tag found
			}
		}

		return true // Include this test
	}

	// Use ccl-test-lib loader to load the test file from flat format
	testLoader := loader.NewTestLoader(".", impl)
	testSuite, err := testLoader.LoadTestFile(jsonFile, loader.LoadOptions{
		Format:       loader.FormatFlat,
		FilterMode:   loader.FilterCustom,
		CustomFilter: customFilter,
	})
	if err != nil {
		return fmt.Errorf("failed to load flat format test file %s: %w", jsonFile, err)
	}

	// Generate test file content
	testContent, err := g.generateTestContent(*testSuite, jsonFile)
	if err != nil {
		return fmt.Errorf("failed to generate test content for %s: %w", filepath.Base(jsonFile), err)
	}

	// Determine output file path
	outputPath := g.getOutputPath(*testSuite, jsonFile)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", filepath.Dir(outputPath), err)
	}

	// Write test file
	if err := os.WriteFile(outputPath, []byte(testContent), 0644); err != nil {
		return fmt.Errorf("failed to write test file %s: %w", outputPath, err)
	}

	return nil
}

// generateTestContent creates the Go test file content
func (g *Generator) generateTestContent(testSuite types.TestSuite, sourceFile string) (string, error) {
	return g.generateTestContentFromTemplate(testSuite, sourceFile)
}

// getOutputPath determines where to place the generated test file
func (g *Generator) getOutputPath(testSuite types.TestSuite, sourceFile string) string {
	// Extract base filename without extension
	baseName := strings.TrimSuffix(filepath.Base(sourceFile), ".json")

	// Determine feature from the first test or suite name
	feature := g.inferFeature(testSuite)

	// Create directory structure: feature/
	dirName := strings.ReplaceAll(feature, "-", "_")

	return filepath.Join(g.outputDir, dirName, baseName+"_test.go")
}

// getPackageName generates the package name for the test file
func (g *Generator) getPackageName(testSuite types.TestSuite) string {
	feature := g.inferFeature(testSuite)
	return strings.ReplaceAll(feature, "-", "_")
}

// inferFeature attempts to determine the feature from the test suite
func (g *Generator) inferFeature(testSuite types.TestSuite) string {
	// Look at the first test's metadata
	if len(testSuite.Tests) > 0 && testSuite.Tests[0].Meta.Feature != "" {
		return testSuite.Tests[0].Meta.Feature
	}

	// Fall back to parsing suite name
	suiteName := strings.ToLower(testSuite.Suite)
	if strings.Contains(suiteName, "typed") {
		return "typed_parsing"
	} else if strings.Contains(suiteName, "object") {
		return "object_construction"
	} else if strings.Contains(suiteName, "dotted") {
		return "dotted_keys"
	} else if strings.Contains(suiteName, "comment") {
		return "comments"
	} else if strings.Contains(suiteName, "processing") {
		return "processing"
	}
	return "parsing"
}
