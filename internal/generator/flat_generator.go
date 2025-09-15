package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SourceTestSuite represents the top-level test suite structure
type SourceTestSuite struct {
	Schema       string            `json:"$schema"`
	Suite        string            `json:"suite"`
	Version      string            `json:"version"`
	Description  string            `json:"description"`
	LLMMetadata  *json.RawMessage  `json:"llm_metadata,omitempty"`
	FeatureFlag  string            `json:"feature_flag,omitempty"`
	Tests        []SourceTest      `json:"tests"`
}

// SourceTest represents a single test case in the original format
type SourceTest struct {
	Name        string                    `json:"name"`
	Input       string                    `json:"input"`
	Validations map[string]interface{}    `json:"validations"`
	Meta        SourceTestMeta            `json:"meta"`
}

// SourceTestMeta represents the metadata structure with structured tags
type SourceTestMeta struct {
	Tags      []string  `json:"tags"`
	Level     int       `json:"level,omitempty"`
	Feature   string    `json:"feature,omitempty"`
	Conflicts []string  `json:"conflicts,omitempty"`
}

// FlatTest represents the generated flat format for implementations
type FlatTest struct {
	Name        string              `json:"name"`
	Input       string              `json:"input"`
	Validation  string              `json:"validation"`
	Expected    ExpectedResult      `json:"expected"`
	Args        []string            `json:"args,omitempty"`
	Functions   []string            `json:"functions,omitempty"`
	Behaviors   []string            `json:"behaviors,omitempty"`
	Variants    []string            `json:"variants,omitempty"`
	Features    []string            `json:"features,omitempty"`
	Conflicts   ConflictsByCategory `json:"conflicts,omitempty"`
	Requires    []string            `json:"requires,omitempty"`
	Level       int                 `json:"level,omitempty"`
	SourceTest  string              `json:"source_test,omitempty"`
	ExpectError bool                `json:"expect_error,omitempty"`
	ErrorType   string              `json:"error_type,omitempty"`
}

// ConflictsByCategory represents conflicts organized by tag type
type ConflictsByCategory struct {
	Functions []string `json:"functions,omitempty"`
	Behaviors []string `json:"behaviors,omitempty"`
	Variants  []string `json:"variants,omitempty"`
	Features  []string `json:"features,omitempty"`
}

// ExpectedResult standardized result format
type ExpectedResult struct {
	Count   int         `json:"count"`
	Entries []Entry     `json:"entries,omitempty"`
	Object  interface{} `json:"object,omitempty"`
	Value   interface{} `json:"value,omitempty"`
	List    []interface{} `json:"list,omitempty"`
}

// Entry represents a key-value pair
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// FunctionDependencies maps CCL functions to their prerequisites
var FunctionDependencies = map[string][]string{
	"parse":            {},
	"parse_value":      {"parse"},
	"filter":           {"parse"},
	"expand_dotted":    {"parse"},
	"build_hierarchy":  {"parse"},
	"get_string":       {"parse", "build_hierarchy"},
	"get_int":          {"parse", "build_hierarchy"},
	"get_bool":         {"parse", "build_hierarchy"},
	"get_float":        {"parse", "build_hierarchy"},
	"get_list":         {"parse", "build_hierarchy"},
	"load":             {"parse"},
	"round_trip":       {"parse"},
	"canonical_format": {"parse"},
	"associativity":    {"parse"},
}

// GenerateFlatTests converts source format to flat format tests
func GenerateFlatTests(sourceTests []SourceTest) ([]FlatTest, error) {
	var flatTests []FlatTest

	for _, sourceTest := range sourceTests {
		for validationName, validationData := range sourceTest.Validations {
			flatTest, err := convertToFlatTest(sourceTest, validationName, validationData)
			if err != nil {
				return nil, fmt.Errorf("error converting test %s validation %s: %w", 
					sourceTest.Name, validationName, err)
			}
			flatTests = append(flatTests, flatTest)
		}
	}

	return flatTests, nil
}

// convertToFlatTest converts a single source validation to flat format
func convertToFlatTest(source SourceTest, validationName string, validationData interface{}) (FlatTest, error) {
	// Parse validation data
	validationMap, ok := validationData.(map[string]interface{})
	if !ok {
		return FlatTest{}, fmt.Errorf("validation data must be an object")
	}

	// Extract expected result and other fields
	var expectedData interface{}
	var args []string
	var expectError bool
	var errorMessage string

	// Handle different expected result formats
	if exp, exists := validationMap["expected"]; exists {
		expectedData = exp
	} else if count, exists := validationMap["count"]; exists {
		// For count-only format, create a standardized result
		expectedData = map[string]interface{}{
			"count": count,
		}
	}
	if cases, exists := validationMap["cases"]; exists {
		// Handle cases format for typed access functions
		if casesList, ok := cases.([]interface{}); ok && len(casesList) > 0 {
			if firstCase, ok := casesList[0].(map[string]interface{}); ok {
				if argsInterface, exists := firstCase["args"]; exists {
					if argsList, ok := argsInterface.([]interface{}); ok {
						for _, arg := range argsList {
							if argStr, ok := arg.(string); ok {
								args = append(args, argStr)
							}
						}
					}
				}
				if exp, exists := firstCase["expected"]; exists {
					expectedData = exp
				}
			}
		}
	}
	if err, exists := validationMap["error"]; exists {
		if errBool, ok := err.(bool); ok {
			expectError = errBool
		}
	}
	if errMsg, exists := validationMap["error_message"]; exists {
		if errMsgStr, ok := errMsg.(string); ok {
			errorMessage = errMsgStr
		}
	}

	expected, err := standardizeExpected(validationName, expectedData)
	if err != nil {
		return FlatTest{}, fmt.Errorf("error standardizing expected result: %w", err)
	}

	// Parse structured tags into separate fields
	functions, behaviors, variants, features := parseStructuredTags(source.Meta.Tags)
	
	// Add the validation function if not already present
	if !contains(functions, validationName) {
		functions = append(functions, validationName)
	}

	// Parse conflicts into categorized structure
	conflictsByCategory := parseConflicts(source.Meta.Conflicts)

	flatTest := FlatTest{
		Name:        fmt.Sprintf("%s_%s", source.Name, validationName),
		Input:       source.Input,
		Validation:  validationName,
		Expected:    expected,
		Args:        args,
		Functions:   functions,
		Behaviors:   behaviors,
		Variants:    variants,
		Features:    features,
		Conflicts:   conflictsByCategory,
		Requires:    FunctionDependencies[validationName],
		Level:       source.Meta.Level,
		SourceTest:  source.Name,
		ExpectError: expectError,
		ErrorType:   errorMessage,
	}

	return flatTest, nil
}

// standardizeExpected converts various expected formats to standardized format
func standardizeExpected(function string, expect interface{}) (ExpectedResult, error) {
	// Handle nil expected values (error cases)
	if expect == nil {
		return ExpectedResult{Count: 0}, nil
	}

	// Handle count-only format common in original tests
	if expectMap, ok := expect.(map[string]interface{}); ok {
		if count, exists := expectMap["count"]; exists {
			if countInt, ok := count.(float64); ok {
				return ExpectedResult{Count: int(countInt)}, nil
			}
		}
	}

	switch function {
	case "parse", "parse_value", "filter", "expand_dotted":
		return standardizeEntriesResult(expect)
	case "build_hierarchy":
		return standardizeObjectResult(expect)
	case "get_string", "get_int", "get_bool", "get_float":
		return standardizeValueResult(expect)
	case "get_list":
		return standardizeListResult(expect)
	case "load", "round_trip", "canonical_format", "associativity":
		return standardizeValueResult(expect)
	default:
		return ExpectedResult{}, fmt.Errorf("unknown function: %s", function)
	}
}

// standardizeEntriesResult standardizes array-based results
func standardizeEntriesResult(expect interface{}) (ExpectedResult, error) {
	entries, ok := expect.([]interface{})
	if !ok {
		return ExpectedResult{}, fmt.Errorf("expected array for entries result, got %T", expect)
	}

	var standardEntries []Entry
	for _, item := range entries {
		entryMap, ok := item.(map[string]interface{})
		if !ok {
			return ExpectedResult{}, fmt.Errorf("expected entry object, got %T", item)
		}

		key, keyOk := entryMap["key"].(string)
		value, valueOk := entryMap["value"].(string)
		if !keyOk || !valueOk {
			return ExpectedResult{}, fmt.Errorf("entry must have string key and value")
		}

		standardEntries = append(standardEntries, Entry{Key: key, Value: value})
	}

	return ExpectedResult{
		Count:   len(standardEntries),
		Entries: standardEntries,
	}, nil
}

// standardizeObjectResult standardizes object-based results
func standardizeObjectResult(expect interface{}) (ExpectedResult, error) {
	return ExpectedResult{
		Count:  1,
		Object: expect,
	}, nil
}

// standardizeValueResult standardizes single value results
func standardizeValueResult(expect interface{}) (ExpectedResult, error) {
	return ExpectedResult{
		Count: 1,
		Value: expect,
	}, nil
}

// standardizeListResult standardizes list-based results
func standardizeListResult(expect interface{}) (ExpectedResult, error) {
	list, ok := expect.([]interface{})
	if !ok {
		return ExpectedResult{}, fmt.Errorf("expected array for list result, got %T", expect)
	}

	return ExpectedResult{
		Count: len(list),
		List:  list,
	}, nil
}

// LoadSourceTests loads source format tests from a directory
func LoadSourceTests(sourceDir string) ([]SourceTest, error) {
	var allTests []SourceTest

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", path, err)
		}

		var testSuite SourceTestSuite
		if err := json.Unmarshal(data, &testSuite); err != nil {
			return fmt.Errorf("error parsing JSON in %s: %w", path, err)
		}

		allTests = append(allTests, testSuite.Tests...)
		return nil
	})

	return allTests, err
}

// SaveFlatTests saves flat format tests to files
func SaveFlatTests(flatTests []FlatTest, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	// Group tests by original source for better organization
	testsBySource := make(map[string][]FlatTest)
	for _, test := range flatTests {
		source := test.SourceTest
		if source == "" {
			source = "unknown"
		}
		testsBySource[source] = append(testsBySource[source], test)
	}

	// Save each source group to its own file
	for sourceName, tests := range testsBySource {
		filename := fmt.Sprintf("%s-flat.json", sourceName)
		filepath := filepath.Join(outputDir, filename)

		data, err := json.MarshalIndent(tests, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling tests for %s: %w", sourceName, err)
		}

		if err := os.WriteFile(filepath, data, 0644); err != nil {
			return fmt.Errorf("error writing file %s: %w", filepath, err)
		}
	}

	return nil
}

// parseStructuredTags parses structured tags into separate categories
func parseStructuredTags(tags []string) (functions, behaviors, variants, features []string) {
	for _, tag := range tags {
		parts := strings.SplitN(tag, ":", 2)
		if len(parts) != 2 {
			continue
		}
		
		category := parts[0]
		value := parts[1]
		
		switch category {
		case "function":
			functions = append(functions, value)
		case "behavior":
			behaviors = append(behaviors, value)
		case "variant":
			variants = append(variants, value)
		case "feature":
			features = append(features, value)
		}
	}
	return
}

// parseConflicts parses conflict tags into categorized structure
func parseConflicts(conflicts []string) ConflictsByCategory {
	var result ConflictsByCategory
	
	for _, conflict := range conflicts {
		parts := strings.SplitN(conflict, ":", 2)
		if len(parts) != 2 {
			continue
		}
		
		category := parts[0]
		value := parts[1]
		
		switch category {
		case "function":
			result.Functions = append(result.Functions, value)
		case "behavior":
			result.Behaviors = append(result.Behaviors, value)
		case "variant":
			result.Variants = append(result.Variants, value)
		case "feature":
			result.Features = append(result.Features, value)
		}
	}
	return result
}

// contains checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}