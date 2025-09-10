package generator

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ccl-test-data/test-runner/internal/styles"
	"github.com/ccl-test-data/test-runner/internal/types"
)

// Options configures test generation behavior
type Options struct {
	SkipDisabled bool     // Skip tests with disabled feature tags
	SkipTags     []string // Additional tags to skip
	RunOnly      []string // Only run tests with these tags (overrides skip behavior)
}

// AssertionStats tracks assertion counts from test generation
type AssertionStats struct {
	TotalTests      int
	TotalAssertions int
	SkippedTests    int
	SkippedAssertions int
	TestCounts      map[string]int // test name -> assertion count
}

// Generator handles test file generation from JSON test data
type Generator struct {
	inputDir  string
	outputDir string
	options   Options
	stats     AssertionStats
}

// New creates a new generator instance with default options
func New(inputDir, outputDir string) *Generator {
	return &Generator{
		inputDir:  inputDir,
		outputDir: outputDir,
		options: Options{
			SkipDisabled: true, // Default behavior
		},
		stats: AssertionStats{
			TestCounts: make(map[string]int),
		},
	}
}

// NewWithOptions creates a new generator instance with custom options
func NewWithOptions(inputDir, outputDir string, options Options) *Generator {
	return &Generator{
		inputDir:  inputDir,
		outputDir: outputDir,
		options:   options,
		stats: AssertionStats{
			TestCounts: make(map[string]int),
		},
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

// generateTestFile generates a Go test file from a JSON test suite
func (g *Generator) generateTestFile(jsonFile string) error {
	// Read and parse JSON file
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", jsonFile, err)
	}

	var testSuite types.TestSuite
	if err := json.Unmarshal(data, &testSuite); err != nil {
		return fmt.Errorf("failed to parse JSON in %s: %w", jsonFile, err)
	}

	// Generate test file content
	testContent, err := g.generateTestContent(testSuite, jsonFile)
	if err != nil {
		return fmt.Errorf("failed to generate test content: %w", err)
	}

	// Determine output file path
	outputPath := g.getOutputPath(testSuite, jsonFile)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write test file
	if err := os.WriteFile(outputPath, []byte(testContent), 0644); err != nil {
		return fmt.Errorf("failed to write test file: %w", err)
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
	
	// Determine level and feature from the first test or suite name
	level := g.inferLevel(testSuite)
	feature := g.inferFeature(testSuite)
	
	// Create directory structure: level-X-feature/
	dirName := fmt.Sprintf("level%d_%s", level, strings.ReplaceAll(feature, "-", "_"))
	
	return filepath.Join(g.outputDir, dirName, baseName+"_test.go")
}

// getPackageName generates the package name for the test file
func (g *Generator) getPackageName(testSuite types.TestSuite) string {
	level := g.inferLevel(testSuite)
	feature := g.inferFeature(testSuite)
	return fmt.Sprintf("level%d_%s", level, strings.ReplaceAll(feature, "-", "_"))
}

// inferLevel attempts to determine the CCL level from the test suite
func (g *Generator) inferLevel(testSuite types.TestSuite) int {
	// Look at the first test's metadata
	if len(testSuite.Tests) > 0 {
		return testSuite.Tests[0].Meta.Level
	}
	
	// Fall back to parsing suite name
	suiteName := strings.ToLower(testSuite.Suite)
	if strings.Contains(suiteName, "typed") {
		return 4
	} else if strings.Contains(suiteName, "object") {
		return 3
	} else if strings.Contains(suiteName, "processing") || strings.Contains(suiteName, "dotted") {
		return 2
	}
	return 1
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
		return "typed-parsing"
	} else if strings.Contains(suiteName, "object") {
		return "object-construction"
	} else if strings.Contains(suiteName, "dotted") {
		return "dotted-keys"
	} else if strings.Contains(suiteName, "comment") {
		return "comments"
	} else if strings.Contains(suiteName, "processing") {
		return "processing"
	}
	return "parsing"
}