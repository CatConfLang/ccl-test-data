// Package stats provides comprehensive test suite analysis and statistics collection.
//
// This package analyzes JSON test files to provide detailed metrics about test coverage,
// feature usage, assertion counts, and implementation requirements. It supports both
// basic and enhanced statistics collection with feature-based categorization.
//
// Key Features:
//   - Enhanced statistics with function, feature, behavior, and variant analysis
//   - Object pooling to reduce memory allocations during collection
//   - Structured tagging system analysis (function:*, feature:*, behavior:*, variant:*)
//   - Conflict relationship tracking for mutually exclusive implementation choices
//   - JSON and pretty-print output formats for both human and machine consumption
//
// Statistics Categories:
//   - Function Requirements: Which CCL functions are tested (parse, build-hierarchy, etc.)
//   - Language Features: Optional features like comments, unicode, multiline support
//   - Behavioral Choices: Implementation-specific behaviors (tabs-preserve vs tabs-to-spaces)
//   - Implementation Variants: Specification variants (proposed-behavior vs reference-compliant)
//
// Example Usage:
//
//	collector := stats.NewEnhancedCollector("tests")
//	statistics, err := collector.CollectEnhancedStats()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	stats.PrintEnhancedStats(statistics)
package stats

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/catconflang/ccl-test-data/types"
)

// String slice pool for reducing allocations during stats collection
var stringSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]string, 0, 16) // Pre-allocate capacity
	},
}

// getStringSlice returns a reusable string slice from the pool
func getStringSlice() []string {
	slice := stringSlicePool.Get().([]string)
	return slice[:0] // Reset length but keep capacity
}

// putStringSlice returns a string slice to the pool
func putStringSlice(slice []string) {
	if slice != nil && cap(slice) <= 256 { // Avoid holding onto huge slices
		stringSlicePool.Put(slice)
	}
}

// Enhanced statistics structures
type FunctionStats struct {
	Tests      int      `json:"tests"`
	Assertions int      `json:"assertions"`
	Files      []string `json:"files"`
}

type FeatureStats struct {
	Tests      int      `json:"tests"`
	Assertions int      `json:"assertions"`
	Files      []string `json:"files"`
}

type BehaviorStats struct {
	Tests      int            `json:"tests"`
	Assertions int            `json:"assertions"`
	Files      []string       `json:"files"`
	Conflicts  map[string]int `json:"conflicts"`
}

type VariantStats struct {
	Tests      int      `json:"tests"`
	Assertions int      `json:"assertions"`
	Files      []string `json:"files"`
}

type ConflictGroup struct {
	Tags  []string `json:"tags"`
	Count int      `json:"count"`
}

type EnhancedStatistics struct {
	Structure       string `json:"structure"`
	TotalTests      int    `json:"totalTests"`
	TotalAssertions int    `json:"totalAssertions"`
	TotalFiles      int    `json:"totalFiles"`

	// Feature-based analysis
	Functions map[string]*FunctionStats `json:"functions"`
	Features  map[string]*FeatureStats  `json:"features"`
	Behaviors map[string]*BehaviorStats `json:"behaviors"`
	Variants  map[string]*VariantStats  `json:"variants"`

	// Conflict analysis
	ConflictPairs     map[string][]string `json:"conflictPairs"`
	ConflictGroups    []ConflictGroup     `json:"conflictGroups"`
	MutuallyExclusive int                 `json:"mutuallyExclusiveTests"`

	// Legacy compatibility
	Categories map[string]*CategoryStats `json:"categories"`
}

// EnhancedCollector provides detailed statistics about the new tagging system
type EnhancedCollector struct {
	testDir string
}

// NewEnhancedCollector creates a new enhanced statistics collector
func NewEnhancedCollector(testDir string) *EnhancedCollector {
	return &EnhancedCollector{testDir: testDir}
}

// parseTag extracts category and name from structured tags like "function:parse"
func parseTag(tag string) (category, name string) {
	parts := strings.SplitN(tag, ":", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return "legacy", tag
}

// analyzeEnhancedTestFile analyzes a test file with the new tagging system
func (c *EnhancedCollector) analyzeEnhancedTestFile(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", filePath, err)
	}

	// Parse as new source format (object with $schema and tests)
	var sourceTestFile SourceTestFile
	if err := json.Unmarshal(data, &sourceTestFile); err == nil && len(sourceTestFile.Tests) > 0 {
		return c.analyzeSourceTests(sourceTestFile.Tests, filePath)
	}

	// Fallback to flat format (TestSuite)
	var testSuite types.TestSuite
	if err := json.Unmarshal(data, &testSuite); err != nil {
		return nil, fmt.Errorf("parsing JSON in %s: %w", filePath, err)
	}
	return c.analyzeTestSuite(testSuite, filePath)
}

// analyzeSourceTests analyzes source format tests for enhanced stats
func (c *EnhancedCollector) analyzeSourceTests(sourceTests []SourceTest, filePath string) (map[string]interface{}, error) {
	fileName := strings.TrimSuffix(filepath.Base(filePath), ".json")
	result := map[string]interface{}{
		"fileName": fileName,
		"tests":    []map[string]interface{}{},
	}

	for _, test := range sourceTests {
		// Count assertions for this test (each validation is one assertion)
		assertions := len(test.Tests)

		// Build tag arrays from top-level fields and convert to structured tags
		functions := getStringSlice()
		features := getStringSlice()
		behaviors := getStringSlice()
		variants := getStringSlice()

		// Ensure slices are returned to pool after use
		defer func() {
			putStringSlice(functions)
			putStringSlice(features)
			putStringSlice(behaviors)
			putStringSlice(variants)
		}()

		// Extract functions from the test validations
		for _, validation := range test.Tests {
			functions = append(functions, validation.Function)
		}

		// Copy features, behaviors, variants from top-level fields
		features = append(features, test.Features...)
		behaviors = append(behaviors, test.Behaviors...)
		variants = append(variants, test.Variants...)

		// Create copies of slices before returning them to pool
		functionsCopy := make([]string, len(functions))
		copy(functionsCopy, functions)
		featuresCopy := make([]string, len(features))
		copy(featuresCopy, features)
		behaviorsCopy := make([]string, len(behaviors))
		copy(behaviorsCopy, behaviors)
		variantsCopy := make([]string, len(variants))
		copy(variantsCopy, variants)

		// Convert conflicts to string slice (flatten all conflict categories)
		var conflictSlice []string
		if test.Conflicts != nil {
			for _, conflicts := range test.Conflicts {
				conflictSlice = append(conflictSlice, conflicts...)
			}
		}

		// Determine feature for categorization (use first feature or derive from filename)
		feature := ""
		if len(test.Features) > 0 {
			feature = test.Features[0]
		} else {
			// Derive feature from filename
			if strings.Contains(fileName, "parsing") {
				feature = "parsing"
			} else if strings.Contains(fileName, "object") {
				feature = "object_construction"
			} else if strings.Contains(fileName, "typed") {
				feature = "typed_parsing"
			} else if strings.Contains(fileName, "comments") {
				feature = "comments"
			} else if strings.Contains(fileName, "processing") {
				feature = "processing"
			}
		}

		testData := map[string]interface{}{
			"name":       test.Name,
			"feature":    feature,
			"assertions": assertions,
			"functions":  functionsCopy,
			"features":   featuresCopy,
			"behaviors":  behaviorsCopy,
			"variants":   variantsCopy,
		}

		// Only include conflicts field if it has actual values
		if len(conflictSlice) > 0 {
			testData["conflicts"] = conflictSlice
		}

		result["tests"] = append(result["tests"].([]map[string]interface{}), testData)
	}

	return result, nil
}

// analyzeTestSuite analyzes flat format test suite (fallback)
func (c *EnhancedCollector) analyzeTestSuite(testSuite types.TestSuite, filePath string) (map[string]interface{}, error) {

	fileName := strings.TrimSuffix(filepath.Base(filePath), ".json")
	result := map[string]interface{}{
		"fileName": fileName,
		"tests":    []map[string]interface{}{},
	}

	for _, test := range testSuite.Tests {
		// Count assertions for this test
		assertions := 0
		validationData := map[string]interface{}{
			"parse":               test.Validations.Parse,
			"parse_indented":      test.Validations.ParseIndented,
			"filter":              test.Validations.Filter,
			"combine":             test.Validations.Combine,
			"expand_dotted":       test.Validations.ExpandDotted,
			"build_hierarchy":     test.Validations.BuildHierarchy,
			"get_string":          test.Validations.GetString,
			"get_int":             test.Validations.GetInt,
			"get_bool":            test.Validations.GetBool,
			"get_float":           test.Validations.GetFloat,
			"get_list":            test.Validations.GetList,
			"pretty_print":        test.Validations.PrettyPrint,
			"round_trip":          test.Validations.RoundTrip,
			"compose_associative": test.Validations.ComposeAssociative,
			"identity_left":       test.Validations.IdentityLeft,
			"identity_right":      test.Validations.IdentityRight,
		}

		for _, validation := range validationData {
			if validation != nil {
				assertions += countAssertions(validation)
			}
		}

		// Analyze tags using pooled slices
		functions := getStringSlice()
		features := getStringSlice()
		behaviors := getStringSlice()
		variants := getStringSlice()

		// Ensure slices are returned to pool after use
		defer func() {
			putStringSlice(functions)
			putStringSlice(features)
			putStringSlice(behaviors)
			putStringSlice(variants)
		}()

		for _, tag := range test.Meta.Tags {
			category, name := parseTag(tag)
			switch category {
			case "function":
				functions = append(functions, name)
			case "feature":
				features = append(features, name)
			case "behavior":
				behaviors = append(behaviors, name)
			case "variant":
				variants = append(variants, name)
			}
		}

		// Create copies of slices before returning them to pool
		functionsCopy := make([]string, len(functions))
		copy(functionsCopy, functions)
		featuresCopy := make([]string, len(features))
		copy(featuresCopy, features)
		behaviorsCopy := make([]string, len(behaviors))
		copy(behaviorsCopy, behaviors)
		variantsCopy := make([]string, len(variants))
		copy(variantsCopy, variants)

		testData := map[string]interface{}{
			"name":       test.Name,
			"feature":    test.Meta.Feature,
			"assertions": assertions,
			"functions":  functionsCopy,
			"features":   featuresCopy,
			"behaviors":  behaviorsCopy,
			"variants":   variantsCopy,
		}

		// Only include conflicts field if it has actual values
		if test.Conflicts != nil && (len(test.Conflicts.Behaviors) > 0 || len(test.Conflicts.Variants) > 0 || len(test.Conflicts.Features) > 0) {
			testData["conflicts"] = test.Conflicts
		}

		result["tests"] = append(result["tests"].([]map[string]interface{}), testData)
	}

	return result, nil
}

// CollectEnhancedStats collects enhanced statistics with the new tagging system
func (c *EnhancedCollector) CollectEnhancedStats() (*EnhancedStatistics, error) {
	testFiles, err := c.findTestFiles()
	if err != nil {
		return nil, fmt.Errorf("finding test files: %w", err)
	}

	stats := &EnhancedStatistics{
		Structure:     "feature-based-enhanced",
		Functions:     make(map[string]*FunctionStats),
		Features:      make(map[string]*FeatureStats),
		Behaviors:     make(map[string]*BehaviorStats),
		Variants:      make(map[string]*VariantStats),
		ConflictPairs: make(map[string][]string),
		Categories:    make(map[string]*CategoryStats),
	}

	// Initialize legacy categories for compatibility
	stats.Categories = map[string]*CategoryStats{
		"core-parsing":        {Files: make(map[string]*FileStats)},
		"advanced-processing": {Files: make(map[string]*FileStats)},
		"object-construction": {Files: make(map[string]*FileStats)},
		"type-system":         {Files: make(map[string]*FileStats)},
		"output-validation":   {Files: make(map[string]*FileStats)},
		"other":               {Files: make(map[string]*FileStats)},
	}

	conflictPairs := make(map[string]map[string]int)
	allConflicts := make(map[string]map[string]bool) // behavior -> set of conflicts

	// First pass: collect all conflict relationships from test data
	// Only record conflicts between behaviors that share a common prefix
	// (e.g., list_coercion_enabled conflicts with list_coercion_disabled)
	for _, filePath := range testFiles {
		fileData, err := c.analyzeEnhancedTestFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
			continue
		}

		tests := fileData["tests"].([]map[string]interface{})
		for _, testData := range tests {
			behaviors := testData["behaviors"].([]string)

			if conflictSlice, ok := testData["conflicts"].([]string); ok && len(conflictSlice) > 0 {
				for _, behavior := range behaviors {
					if allConflicts[behavior] == nil {
						allConflicts[behavior] = make(map[string]bool)
					}
					for _, conflict := range conflictSlice {
						// Only record conflict if they share a common prefix
						// This prevents cross-contamination between different behavior types
						if haveSameBehaviorPrefix(behavior, conflict) {
							allConflicts[behavior][conflict] = true
						}
					}
				}
			}
		}
	}

	// Infer groups: behaviors that conflict with each other bidirectionally are in the same group
	behaviorGroups := make(map[string][]string)
	processed := make(map[string]bool)

	for behavior := range allConflicts {
		if processed[behavior] {
			continue
		}

		group := []string{behavior}
		processed[behavior] = true

		// Find all other behaviors that conflict with this one bidirectionally
		for otherBehavior := range allConflicts {
			if processed[otherBehavior] {
				continue
			}
			// Both must conflict with each other
			if allConflicts[behavior][otherBehavior] && allConflicts[otherBehavior][behavior] {
				group = append(group, otherBehavior)
				processed[otherBehavior] = true
			}
		}

		// Store the group for all members
		if len(group) > 1 {
			for _, b := range group {
				behaviorGroups[b] = group
			}
		}
	}

	// Second pass: collect stats and build conflict pairs
	for _, filePath := range testFiles {
		fileData, err := c.analyzeEnhancedTestFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
			continue
		}

		fileName := fileData["fileName"].(string)
		tests := fileData["tests"].([]map[string]interface{})

		if len(tests) == 0 {
			continue
		}

		fileStats := &FileStats{Tests: len(tests), Assertions: 0}

		for _, testData := range tests {
			assertions := testData["assertions"].(int)
			feature := testData["feature"].(string)
			functions := testData["functions"].([]string)
			features := testData["features"].([]string)
			behaviors := testData["behaviors"].([]string)
			variants := testData["variants"].([]string)

			fileStats.Assertions += assertions
			stats.TotalTests++
			stats.TotalAssertions += assertions

			// Function stats
			for _, fn := range functions {
				if stats.Functions[fn] == nil {
					stats.Functions[fn] = &FunctionStats{Files: []string{}}
				}
				fnStats := stats.Functions[fn]
				fnStats.Tests++
				fnStats.Assertions += assertions
				if !contains(fnStats.Files, fileName) {
					fnStats.Files = append(fnStats.Files, fileName)
				}
			}

			// Feature stats
			for _, feat := range features {
				if stats.Features[feat] == nil {
					stats.Features[feat] = &FeatureStats{Files: []string{}}
				}
				featStats := stats.Features[feat]
				featStats.Tests++
				featStats.Assertions += assertions
				if !contains(featStats.Files, fileName) {
					featStats.Files = append(featStats.Files, fileName)
				}
			}

			// Behavior stats
			for _, behav := range behaviors {
				if stats.Behaviors[behav] == nil {
					stats.Behaviors[behav] = &BehaviorStats{Files: []string{}, Conflicts: make(map[string]int)}
				}
				behavStats := stats.Behaviors[behav]
				behavStats.Tests++
				behavStats.Assertions += assertions
				if !contains(behavStats.Files, fileName) {
					behavStats.Files = append(behavStats.Files, fileName)
				}
			}

			// Variant stats
			for _, variant := range variants {
				if stats.Variants[variant] == nil {
					stats.Variants[variant] = &VariantStats{Files: []string{}}
				}
				varStats := stats.Variants[variant]
				varStats.Tests++
				varStats.Assertions += assertions
				if !contains(varStats.Files, fileName) {
					varStats.Files = append(varStats.Files, fileName)
				}
			}

			// Conflict analysis
			if conflictSlice, ok := testData["conflicts"].([]string); ok && len(conflictSlice) > 0 {
				// Count this as a mutually exclusive test
				stats.MutuallyExclusive++

				// Track conflicts only for behaviors that are in the same conflict group
				// This prevents incorrectly attributing array_order conflicts to list_coercion behaviors
				for _, behavior := range behaviors {
					for _, conflict := range conflictSlice {
						// Only attribute conflict if they're in the same bidirectional group
						if isInSameConflictGroup(behavior, conflict, behaviorGroups) {
							fullTag := "behavior:" + behavior
							if conflictPairs[fullTag] == nil {
								conflictPairs[fullTag] = make(map[string]int)
							}
							conflictPairs[fullTag][conflict]++
						}
					}
				}

				// Handle variant conflicts (these don't need group filtering)
				for _, variant := range variants {
					for _, conflict := range conflictSlice {
						// Check if the conflict is with another variant
						if isInSameConflictGroup(variant, conflict, behaviorGroups) {
							fullTag := "variant:" + variant
							if conflictPairs[fullTag] == nil {
								conflictPairs[fullTag] = make(map[string]int)
							}
							conflictPairs[fullTag][conflict]++
						}
					}
				}
			}

			// Legacy category stats for compatibility
			category := categorizeByFeature(feature)
			if stats.Categories[category] == nil {
				stats.Categories[category] = &CategoryStats{Files: make(map[string]*FileStats)}
			}
			categoryStats := stats.Categories[category]
			if categoryStats.Files[fileName] == nil {
				categoryStats.Files[fileName] = &FileStats{}
			}
			categoryStats.Files[fileName].Tests++
			categoryStats.Files[fileName].Assertions += assertions
		}

		// Update legacy category totals
		for _, categoryStats := range stats.Categories {
			if categoryStats.Files[fileName] != nil {
				categoryStats.Total += fileStats.Tests
				categoryStats.Assertions += fileStats.Assertions
			}
		}

		stats.TotalFiles++
	}

	// Convert conflict pairs to slice format
	for tag, conflicts := range conflictPairs {
		conflictList := []string{}
		for conflict := range conflicts {
			conflictList = append(conflictList, conflict)
		}
		stats.ConflictPairs[tag] = conflictList
	}

	return stats, nil
}

// findTestFiles finds all JSON test files (reuse from base collector)
func (c *EnhancedCollector) findTestFiles() ([]string, error) {
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

// Helper function to check if slice contains string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// isInSameConflictGroup checks if two behaviors are in the same bidirectional conflict group
func isInSameConflictGroup(behavior, conflict string, behaviorGroups map[string][]string) bool {
	group, exists := behaviorGroups[behavior]
	if !exists {
		return false
	}
	for _, member := range group {
		if member == conflict {
			return true
		}
	}
	return false
}

// haveSameBehaviorPrefix checks if two behavior names share a common prefix
// indicating they are the same type of behavior (e.g., list_coercion_enabled and list_coercion_disabled)
func haveSameBehaviorPrefix(a, b string) bool {
	// Known behavior prefixes - behaviors with same prefix are alternatives
	prefixes := []string{
		"list_coercion_",
		"array_order_",
		"tabs_as_",
		"crlf_",
		"toplevel_indent_",
		"boolean_",
	}

	for _, prefix := range prefixes {
		aHas := strings.HasPrefix(a, prefix)
		bHas := strings.HasPrefix(b, prefix)
		if aHas && bHas {
			return true
		}
		// If one has the prefix but the other doesn't, they're different types
		if aHas || bHas {
			return false
		}
	}

	// If neither matches any known prefix, fall back to checking if they share
	// a common underscore-separated prefix
	aParts := strings.Split(a, "_")
	bParts := strings.Split(b, "_")
	if len(aParts) >= 2 && len(bParts) >= 2 {
		return aParts[0] == bParts[0]
	}

	return false
}

// PrintEnhancedStats prints enhanced statistics in a human-readable format
func PrintEnhancedStats(stats *EnhancedStatistics) {
	fmt.Printf("📊 Enhanced CCL Test Suite Statistics\n\n")

	// Overall summary
	fmt.Printf("🔍 Overview:\n")
	fmt.Printf("  Total Tests: %d\n", stats.TotalTests)
	fmt.Printf("  Total Assertions: %d\n", stats.TotalAssertions)
	fmt.Printf("  Total Files: %d\n", stats.TotalFiles)
	fmt.Printf("  Mutually Exclusive Tests: %d\n\n", stats.MutuallyExclusive)

	// Function requirements
	fmt.Printf("⚙️  Function Requirements:\n")
	functionKeys := make([]string, 0, len(stats.Functions))
	for k := range stats.Functions {
		functionKeys = append(functionKeys, k)
	}
	sort.Strings(functionKeys)

	for _, fn := range functionKeys {
		fnStats := stats.Functions[fn]
		fmt.Printf("  function:%s: %d tests (%d assertions) across %d files\n",
			fn, fnStats.Tests, fnStats.Assertions, len(fnStats.Files))
	}
	fmt.Println()

	// Language features
	if len(stats.Features) > 0 {
		fmt.Printf("🎨 Language Features:\n")
		featureKeys := make([]string, 0, len(stats.Features))
		for k := range stats.Features {
			featureKeys = append(featureKeys, k)
		}
		sort.Strings(featureKeys)

		for _, feat := range featureKeys {
			featStats := stats.Features[feat]
			fmt.Printf("  feature:%s: %d tests (%d assertions) across %d files\n",
				feat, featStats.Tests, featStats.Assertions, len(featStats.Files))
		}
		fmt.Println()
	}

	// Behavioral choices
	if len(stats.Behaviors) > 0 {
		fmt.Printf("🔧 Behavioral Choices:\n")
		behaviorKeys := make([]string, 0, len(stats.Behaviors))
		for k := range stats.Behaviors {
			behaviorKeys = append(behaviorKeys, k)
		}
		sort.Strings(behaviorKeys)

		for _, behav := range behaviorKeys {
			behavStats := stats.Behaviors[behav]
			fmt.Printf("  behavior:%s: %d tests (%d assertions)\n",
				behav, behavStats.Tests, behavStats.Assertions)
		}
		fmt.Println()
	}

	// Implementation variants
	if len(stats.Variants) > 0 {
		fmt.Printf("🔀 Implementation Variants:\n")
		variantKeys := make([]string, 0, len(stats.Variants))
		for k := range stats.Variants {
			variantKeys = append(variantKeys, k)
		}
		sort.Strings(variantKeys)

		for _, variant := range variantKeys {
			varStats := stats.Variants[variant]
			fmt.Printf("  variant:%s: %d tests (%d assertions)\n",
				variant, varStats.Tests, varStats.Assertions)
		}
		fmt.Println()
	}

	// Conflict relationships - display as bidirectional pairs
	if len(stats.ConflictPairs) > 0 {
		fmt.Printf("⚔️  Conflict Relationships (mutually exclusive pairs):\n")

		// ANSI color codes
		const (
			colorCyan   = "\033[36m"
			colorYellow = "\033[33m"
			colorReset  = "\033[0m"
		)

		// Collect unique bidirectional pairs
		type conflictPair struct {
			category string // "behavior" or "variant"
			first    string
			second   string
		}
		seenPairs := make(map[string]bool)
		var pairs []conflictPair

		for tag, conflicts := range stats.ConflictPairs {
			// Parse the tag (e.g., "behavior:array_order_insertion")
			parts := strings.SplitN(tag, ":", 2)
			if len(parts) != 2 {
				continue
			}
			category := parts[0]
			name := parts[1]

			for _, conflict := range conflicts {
				// Create a canonical key for the pair (alphabetically sorted)
				var pairKey string
				if name < conflict {
					pairKey = category + ":" + name + "<->" + conflict
				} else {
					pairKey = category + ":" + conflict + "<->" + name
				}

				if !seenPairs[pairKey] {
					seenPairs[pairKey] = true
					// Store with alphabetically first name first
					if name < conflict {
						pairs = append(pairs, conflictPair{category, name, conflict})
					} else {
						pairs = append(pairs, conflictPair{category, conflict, name})
					}
				}
			}
		}

		// Sort pairs by category then by first name
		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].category != pairs[j].category {
				return pairs[i].category < pairs[j].category
			}
			return pairs[i].first < pairs[j].first
		})

		// Print each pair on a single line with colors
		for _, pair := range pairs {
			fmt.Printf("  %s: %s%s%s ↔ %s%s%s\n",
				pair.category,
				colorCyan, pair.first, colorReset,
				colorYellow, pair.second, colorReset)
		}
		fmt.Println()
	}
}
