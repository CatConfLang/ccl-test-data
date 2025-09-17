package stats

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/tylerbu/ccl-test-lib/types"
)

// FeatureCategories maps feature names to categories
var FeatureCategories = map[string]string{
	"parsing":             "core-parsing",
	"processing":          "advanced-processing",
	"comments":            "advanced-processing",
	"object_construction": "object-construction",
	"dotted_keys":         "object-construction",
	"typed_parsing":       "type-system",
	"pretty_printing":     "output-validation",
	"error_handling":      "output-validation",
}

// FileStats represents statistics for a single test file
type FileStats struct {
	Tests      int `json:"tests"`
	Assertions int `json:"assertions"`
}

// CategoryStats represents statistics for a feature category
type CategoryStats struct {
	Total      int                   `json:"total"`
	Assertions int                   `json:"assertions"`
	Files      map[string]*FileStats `json:"files"`
}

// Statistics represents the complete test suite statistics
type Statistics struct {
	Structure       string                    `json:"structure"`
	Categories      map[string]*CategoryStats `json:"categories"`
	TotalTests      int                       `json:"totalTests"`
	TotalAssertions int                       `json:"totalAssertions"`
	TotalFiles      int                       `json:"totalFiles"`
}

// Collector handles statistics collection from test files
type Collector struct {
	testDir string
}

// NewCollector creates a new statistics collector
func NewCollector(testDir string) *Collector {
	return &Collector{testDir: testDir}
}

// findTestFiles finds all JSON test files in the test directory
func (c *Collector) findTestFiles() ([]string, error) {
	var files []string

	err := filepath.WalkDir(c.testDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, "schema.json") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// countAssertions counts assertions in validation data
func countAssertions(validationData interface{}) int {
	if validationData == nil {
		return 0
	}

	// Handle map structure (with count field or error format)
	if m, ok := validationData.(map[string]interface{}); ok {
		// Check for explicit count field
		if count, exists := m["count"]; exists {
			if countFloat, ok := count.(float64); ok {
				return int(countFloat)
			}
		}
		// Check for error format (single assertion)
		if _, exists := m["error"]; exists {
			return 1
		}
	}

	// If no count field, assume single assertion
	return 1
}

// SourceTest represents a single test in source format
type SourceTest struct {
	Name      string                 `json:"name"`
	Input     string                 `json:"input"`
	Tests     []SourceTestValidation `json:"tests"`
	Level     int                    `json:"level,omitempty"`
	Features  []string               `json:"features,omitempty"`
	Behaviors []string               `json:"behaviors,omitempty"`
	Variants  []string               `json:"variants,omitempty"`
	Conflicts map[string][]string    `json:"conflicts,omitempty"`
}

// SourceTestValidation represents a single validation in source format
type SourceTestValidation struct {
	Function string      `json:"function"`
	Expect   interface{} `json:"expect"`
	Args     []string    `json:"args,omitempty"`
	Error    bool        `json:"error,omitempty"`
}

// analyzeTestFile analyzes a single test file and returns its statistics
func (c *Collector) analyzeTestFile(filePath string) (*FileStats, string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("reading file %s: %w", filePath, err)
	}

	// Try to parse as source format (array of tests)
	var sourceTests []SourceTest
	if err := json.Unmarshal(data, &sourceTests); err != nil {
		// Fallback to flat format (TestSuite)
		var testSuite types.TestSuite
		if err := json.Unmarshal(data, &testSuite); err != nil {
			return nil, "", fmt.Errorf("parsing JSON in %s: %w", filePath, err)
		}
		return c.analyzeTestSuite(testSuite)
	}

	return c.analyzeSourceTests(sourceTests)
}

// analyzeSourceTests analyzes source format tests
func (c *Collector) analyzeSourceTests(sourceTests []SourceTest) (*FileStats, string, error) {
	if len(sourceTests) == 0 {
		return &FileStats{}, "", nil
	}

	// Get feature from first test that has one
	feature := ""
	for _, test := range sourceTests {
		if len(test.Features) > 0 {
			feature = test.Features[0] // Use first feature
			break
		}
	}

	// Count tests and assertions
	totalTests := len(sourceTests)
	totalAssertions := 0

	for _, test := range sourceTests {
		// Each validation in the tests array is one assertion
		totalAssertions += len(test.Tests)
	}

	return &FileStats{
		Tests:      totalTests,
		Assertions: totalAssertions,
	}, feature, nil
}

// analyzeTestSuite analyzes flat format test suite (fallback)
func (c *Collector) analyzeTestSuite(testSuite types.TestSuite) (*FileStats, string, error) {
	if len(testSuite.Tests) == 0 {
		return &FileStats{}, "", nil
	}

	// Get feature from first test's metadata
	feature := ""
	if len(testSuite.Tests) > 0 {
		feature = testSuite.Tests[0].Meta.Feature
	}

	// Count assertions
	totalAssertions := 0
	for _, test := range testSuite.Tests {
		// Use reflection to iterate over validation fields
		validationData := map[string]interface{}{
			"parse":           test.Validations.Parse,
			"parse_value":     test.Validations.ParseValue,
			"filter":          test.Validations.Filter,
			"combine":         test.Validations.Combine,
			"expand_dotted":   test.Validations.ExpandDotted,
			"build_hierarchy": test.Validations.BuildHierarchy,
			"get_string":      test.Validations.GetString,
			"get_int":         test.Validations.GetInt,
			"get_bool":        test.Validations.GetBool,
			"get_float":       test.Validations.GetFloat,
			"get_list":        test.Validations.GetList,
			"pretty_print":    test.Validations.PrettyPrint,
			"round_trip":      test.Validations.RoundTrip,
			"associativity":   test.Validations.Associativity,
		}

		for _, validation := range validationData {
			if validation != nil {
				totalAssertions += countAssertions(validation)
			}
		}
	}

	return &FileStats{
		Tests:      len(testSuite.Tests),
		Assertions: totalAssertions,
	}, feature, nil
}

// categorizeByFeature maps a feature to its category
func categorizeByFeature(feature string) string {
	if feature == "" {
		return "other"
	}

	if category, exists := FeatureCategories[feature]; exists {
		return category
	}

	return "other"
}

// CollectStats collects statistics from all test files
func (c *Collector) CollectStats() (*Statistics, error) {
	testFiles, err := c.findTestFiles()
	if err != nil {
		return nil, fmt.Errorf("finding test files: %w", err)
	}

	stats := &Statistics{
		Structure: "feature-based",
		Categories: map[string]*CategoryStats{
			"core-parsing":        {Files: make(map[string]*FileStats)},
			"advanced-processing": {Files: make(map[string]*FileStats)},
			"object-construction": {Files: make(map[string]*FileStats)},
			"type-system":         {Files: make(map[string]*FileStats)},
			"output-validation":   {Files: make(map[string]*FileStats)},
			"other":               {Files: make(map[string]*FileStats)},
		},
	}

	for _, filePath := range testFiles {
		fileStats, feature, err := c.analyzeTestFile(filePath)
		if err != nil {
			// Log error but continue processing other files
			fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
			continue
		}

		if fileStats.Tests == 0 {
			continue
		}

		category := categorizeByFeature(feature)
		fileName := strings.TrimSuffix(filepath.Base(filePath), ".json")

		// Update category stats
		categoryStats := stats.Categories[category]
		categoryStats.Files[fileName] = fileStats
		categoryStats.Total += fileStats.Tests
		categoryStats.Assertions += fileStats.Assertions

		// Update totals
		stats.TotalTests += fileStats.Tests
		stats.TotalAssertions += fileStats.Assertions
		stats.TotalFiles++
	}

	return stats, nil
}
