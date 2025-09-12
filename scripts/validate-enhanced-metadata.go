package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// EnhancedTestFile represents the new v2.1 schema with llm_metadata
type EnhancedTestFile struct {
	Schema      string      `json:"$schema"`
	Suite       string      `json:"suite"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	LLMMetadata LLMMetadata `json:"llm_metadata"`
	Tests       []TestCase  `json:"tests"`
}

// LLMMetadata contains AI-friendly metadata for test files
type LLMMetadata struct {
	LLMDescription      string           `json:"llm_description"`
	ComplexityLevel     string           `json:"complexity_level"`
	PrerequisiteTests   []string         `json:"prerequisite_tests"`
	RelatedFunctions    []string         `json:"related_functions"`
	ImplementationNotes string           `json:"implementation_notes"`
	TestCount           int              `json:"test_count"`
	AssertionCount      int              `json:"assertion_count"`
	CrossReferences     *CrossReferences `json:"cross_references,omitempty"`
	LearningPath        string           `json:"learning_path,omitempty"`
	FeatureFlag         string           `json:"feature_flag,omitempty"`
}

// CrossReferences links to related implementations and documentation
type CrossReferences struct {
	GleamImplementation string `json:"gleam_implementation,omitempty"`
	GoReference         string `json:"go_reference,omitempty"`
	Documentation       string `json:"documentation,omitempty"`
}

// TestCase represents individual test with metadata
type TestCase struct {
	Name        string                 `json:"name"`
	Input       string                 `json:"input"`
	Validations map[string]interface{} `json:"validations"`
	Meta        TestMeta               `json:"meta"`
}

// TestMeta contains test-level metadata
type TestMeta struct {
	Tags        []string     `json:"tags"`
	Level       int          `json:"level"`
	Feature     string       `json:"feature"`
	LLMGuidance *LLMGuidance `json:"llm_guidance,omitempty"`
}

// LLMGuidance provides test-level AI guidance
type LLMGuidance struct {
	TestPurpose          string                `json:"test_purpose"`
	ImplementationFocus  string                `json:"implementation_focus"`
	CommonPitfalls       []string              `json:"common_pitfalls"`
	DebuggingHints       []string              `json:"debugging_hints,omitempty"`
	RelatedTests         []string              `json:"related_tests,omitempty"`
	ComplexityIndicators *ComplexityIndicators `json:"complexity_indicators,omitempty"`
}

// ComplexityIndicators provide implementation difficulty guidance
type ComplexityIndicators struct {
	ParsingDifficulty string   `json:"parsing_difficulty"`
	EdgeCases         string   `json:"edge_cases"`
	Prerequisites     []string `json:"prerequisites"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run validate-enhanced-metadata.go <test-files-pattern>")
	}

	pattern := os.Args[1]
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("Error globbing files: %v", err)
	}

	if len(files) == 0 {
		log.Fatalf("No files found matching pattern: %s", pattern)
	}

	fmt.Printf("Validating %d enhanced metadata files...\n", len(files))

	var totalErrors int
	for _, file := range files {
		errors := validateFile(file)
		totalErrors += errors
	}

	if totalErrors == 0 {
		fmt.Println("✅ All files passed enhanced metadata validation!")
	} else {
		fmt.Printf("❌ Found %d validation errors across all files\n", totalErrors)
		os.Exit(1)
	}
}

func validateFile(filename string) int {
	fmt.Printf("\n📁 Validating: %s\n", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ Error reading file: %v\n", err)
		return 1
	}

	var testFile EnhancedTestFile
	if err := json.Unmarshal(data, &testFile); err != nil {
		fmt.Printf("❌ Error parsing JSON: %v\n", err)
		return 1
	}

	var errors int

	// Validate schema version
	if testFile.Version != "2.1" {
		fmt.Printf("❌ Expected version 2.1, got: %s\n", testFile.Version)
		errors++
	} else {
		fmt.Printf("✅ Schema version: %s\n", testFile.Version)
	}

	// Validate required LLM metadata fields
	metadata := testFile.LLMMetadata
	if metadata.LLMDescription == "" {
		fmt.Printf("❌ Missing llm_description\n")
		errors++
	} else {
		fmt.Printf("✅ LLM description: %s\n", truncateString(metadata.LLMDescription, 60))
	}

	if metadata.ComplexityLevel == "" {
		fmt.Printf("❌ Missing complexity_level\n")
		errors++
	} else if !isValidComplexityLevel(metadata.ComplexityLevel) {
		fmt.Printf("❌ Invalid complexity_level: %s\n", metadata.ComplexityLevel)
		errors++
	} else {
		fmt.Printf("✅ Complexity level: %s\n", metadata.ComplexityLevel)
	}

	if len(metadata.RelatedFunctions) == 0 {
		fmt.Printf("❌ Missing related_functions\n")
		errors++
	} else {
		fmt.Printf("✅ Related functions: %v\n", metadata.RelatedFunctions)
	}

	if metadata.ImplementationNotes == "" {
		fmt.Printf("❌ Missing implementation_notes\n")
		errors++
	} else {
		fmt.Printf("✅ Implementation notes: %s\n", truncateString(metadata.ImplementationNotes, 60))
	}

	// Validate test count accuracy
	actualTestCount := len(testFile.Tests)
	if metadata.TestCount != actualTestCount {
		fmt.Printf("❌ Test count mismatch: metadata=%d, actual=%d\n", metadata.TestCount, actualTestCount)
		errors++
	} else {
		fmt.Printf("✅ Test count: %d\n", actualTestCount)
	}

	// Validate assertion count (approximate)
	totalAssertions := calculateAssertions(testFile.Tests)
	if metadata.AssertionCount > 0 && abs(metadata.AssertionCount-totalAssertions) > 10 {
		fmt.Printf("⚠️ Assertion count may be inaccurate: metadata=%d, calculated=%d\n", metadata.AssertionCount, totalAssertions)
	} else {
		fmt.Printf("✅ Assertion count: %d\n", metadata.AssertionCount)
	}

	// Validate prerequisite test file references
	for _, prereq := range metadata.PrerequisiteTests {
		if prereq != "" && !strings.HasSuffix(prereq, ".json") {
			fmt.Printf("❌ Invalid prerequisite test format: %s (should end with .json)\n", prereq)
			errors++
		}
	}
	if len(metadata.PrerequisiteTests) > 0 {
		fmt.Printf("✅ Prerequisites: %v\n", metadata.PrerequisiteTests)
	}

	// Validate cross-references format
	if refs := metadata.CrossReferences; refs != nil {
		validateCrossReferences(refs, &errors)
	}

	// Validate test-level metadata (sample)
	testsWithGuidance := 0
	for _, test := range testFile.Tests {
		if test.Meta.LLMGuidance != nil {
			testsWithGuidance++
			validateTestGuidance(test.Name, test.Meta.LLMGuidance, &errors)
		}
	}

	if testsWithGuidance > 0 {
		fmt.Printf("✅ Tests with LLM guidance: %d/%d\n", testsWithGuidance, len(testFile.Tests))
	}

	if errors == 0 {
		fmt.Printf("✅ %s: All validations passed\n", filepath.Base(filename))
	} else {
		fmt.Printf("❌ %s: %d validation errors\n", filepath.Base(filename), errors)
	}

	return errors
}

func isValidComplexityLevel(level string) bool {
	validLevels := []string{
		"Level 1", "Level 1+", "Level 2", "Level 2 Feature",
		"Level 3", "Level 3 Feature", "Level 4", "Level 5",
	}
	for _, valid := range validLevels {
		if level == valid {
			return true
		}
	}
	return false
}

func calculateAssertions(tests []TestCase) int {
	total := 0
	for _, test := range tests {
		for _, validation := range test.Validations {
			if validationMap, ok := validation.(map[string]interface{}); ok {
				if count, exists := validationMap["count"]; exists {
					if countFloat, ok := count.(float64); ok {
						total += int(countFloat)
					}
				}
			}
		}
	}
	return total
}

func validateCrossReferences(refs *CrossReferences, errors *int) {
	if refs.GleamImplementation != "" {
		fmt.Printf("✅ Gleam reference: %s\n", refs.GleamImplementation)
	}
	if refs.GoReference != "" {
		fmt.Printf("✅ Go reference: %s\n", refs.GoReference)
	}
	if refs.Documentation != "" && strings.HasPrefix(refs.Documentation, "https://") {
		fmt.Printf("✅ Documentation: %s\n", refs.Documentation)
	} else if refs.Documentation != "" {
		fmt.Printf("❌ Invalid documentation URL: %s\n", refs.Documentation)
		*errors++
	}
}

func validateTestGuidance(testName string, guidance *LLMGuidance, errors *int) {
	if guidance.TestPurpose == "" {
		fmt.Printf("❌ Test %s: Missing test_purpose in llm_guidance\n", testName)
		*errors++
	}
	if guidance.ImplementationFocus == "" {
		fmt.Printf("❌ Test %s: Missing implementation_focus in llm_guidance\n", testName)
		*errors++
	}

	if indicators := guidance.ComplexityIndicators; indicators != nil {
		validDifficulties := []string{"trivial", "easy", "moderate", "complex", "advanced"}
		validEdgeCases := []string{"none", "few", "moderate", "many"}

		if !contains(validDifficulties, indicators.ParsingDifficulty) {
			fmt.Printf("❌ Test %s: Invalid parsing_difficulty: %s\n", testName, indicators.ParsingDifficulty)
			*errors++
		}
		if !contains(validEdgeCases, indicators.EdgeCases) {
			fmt.Printf("❌ Test %s: Invalid edge_cases: %s\n", testName, indicators.EdgeCases)
			*errors++
		}
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
