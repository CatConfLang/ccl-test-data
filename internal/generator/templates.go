package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/tylerbutler/ccl-test-data/types"
)

const testFileTemplate = `package {{.PackageName}}_test

import (
	"testing"
	{{if .HasActiveTests}}
	"github.com/tylerbutler/ccl-test-data/internal/mock"{{if .HasAssertions}}
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"{{end}}{{end}}
)

// Generated from {{.SourceFile}}
// Suite: {{.Suite}}
// Version: {{.Version}}
{{if .Description}}// Description: {{.Description}}{{end}}

{{range .Tests}}
{{.}}
{{end}}
`

const testCaseTemplate = `// {{.Name}} - {{.TagsString}}
func Test{{.TestFuncName}}(t *testing.T) {
	{{if .ShouldSkip}}t.Skip("{{.SkipReason}}"){{else}}

	ccl := mock.New()
	{{if .IsSingleInput}}input := {{index .InputStrings 0}}{{end}}
	{{if .IsMultiInput}}{{range $i, $s := .InputStrings}}input{{$i}} := {{$s}}
	{{end}}{{end}}
	{{if .HasValidations}}// Declare variables for reuse across validations
	{{if .NeedsParseResult}}var parseResult []mock.Entry{{end}}
	{{if .NeedsObjectResult}}var objectResult map[string]interface{}{{end}}
	{{if .NeedsFilterResult}}var filterResult []mock.Entry{{end}}
	var err error
	{{end}}
{{range .Validations}}	{{.}}
{{end}}{{if not .HasValidations}}	// TODO: Implement test validations
	_ = ccl // Prevent unused variable warning
{{if .IsSingleInput}}	_ = input // Prevent unused variable warning
{{end}}{{if .IsMultiInput}}{{range $i, $s := .Inputs}}	_ = input{{$i}} // Prevent unused variable warning
{{end}}{{end}}{{end}}{{end}}
}
`

// TemplateData holds data for generating test files
type TemplateData struct {
	PackageName    string
	SourceFile     string
	Suite          string
	Version        string
	Description    string
	Tests          []string
	HasActiveTests bool // Whether any tests are not skipped
	HasAssertions  bool // Whether any active tests have assertions
}

// TestCaseData holds data for generating individual test cases
type TestCaseData struct {
	Name              string
	TestFuncName      string
	TagsString        string
	ShouldSkip        bool
	SkipReason        string
	Inputs            []string // CCL input text(s) - single-input tests use 1-element array
	InputStrings      []string // Escaped Go strings for each input
	HasInputs         bool     // True if inputs array exists and has elements
	IsSingleInput     bool     // True if exactly one input (common case)
	IsMultiInput      bool     // True if more than one input (algebraic tests)
	Validations       []string
	HasValidations    bool
	NeedsParseResult  bool
	NeedsObjectResult bool
	NeedsFilterResult bool
}

// generateTestContentFromTemplate creates Go test content using templates
func (g *Generator) generateTestContentFromTemplate(testSuite types.TestSuite, sourceFile string) (string, error) {
	// Generate individual test cases
	var testCases []string
	hasActiveTests := false
	hasAssertions := false

	for _, test := range testSuite.Tests {
		testCase, err := g.generateTestCase(test)
		if err != nil {
			return "", fmt.Errorf("failed to generate test case %s: %w", test.Name, err)
		}
		testCases = append(testCases, testCase)

		// Count assertions and track statistics (flat format has 1 assertion per test)
		assertionCount := 1 // Each flat test case is exactly 1 assertion
		g.stats.TestCounts[test.Name] = assertionCount
		g.stats.TotalTests++

		// Check if this test is not skipped using generator options
		// For flat format, use Functions field instead of Meta.Tags
		tags := g.getTestTags(test)
		isSkipped := g.shouldSkipTestByName(test.Name, tags)
		if isSkipped {
			g.stats.SkippedTests++
			g.stats.SkippedAssertions += assertionCount
		} else {
			hasActiveTests = true
			g.stats.TotalAssertions += assertionCount
			// Check if this test has implemented assertions (flat format always has assertions)
			if test.Validation != "" {
				hasAssertions = true
			}
		}
	}

	// Create template data
	data := TemplateData{
		PackageName:    g.getPackageName(testSuite),
		SourceFile:     sourceFile,
		Suite:          testSuite.Suite,
		Version:        testSuite.Version,
		Description:    testSuite.Description,
		Tests:          testCases,
		HasActiveTests: hasActiveTests,
		HasAssertions:  hasAssertions,
	}

	// Execute template
	tmpl, err := template.New("testfile").Parse(testFileTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// generateTestCase creates a single test case
func (g *Generator) generateTestCase(test types.TestCase) (string, error) {
	// Build escaped input strings
	inputStrings := make([]string, len(test.Inputs))
	for i, input := range test.Inputs {
		inputStrings[i] = escapeGoString(input)
	}

	data := TestCaseData{
		Name:          test.Name,
		TestFuncName:  toPascalCase(test.Name),
		TagsString:    strings.Join(g.getTestTags(test), " "),
		Inputs:        test.Inputs,
		InputStrings:  inputStrings,
		HasInputs:     len(test.Inputs) > 0,
		IsSingleInput: len(test.Inputs) == 1,
		IsMultiInput:  len(test.Inputs) > 1,
	}

	// Check if test should be skipped using generator options
	tags := g.getTestTags(test)
	if g.shouldSkipTestByName(test.Name, tags) {
		data.ShouldSkip = true
		data.SkipReason = g.getSkipReasonByName(test.Name, tags)
	}

	// Generate actual validation for flat format
	validation, err := g.generateFlatFormatValidation(test)
	if err != nil {
		return "", fmt.Errorf("failed to generate flat format validation: %w", err)
	}
	data.Validations = []string{validation}

	// Flat format has real validations
	data.HasValidations = true

	// No variables needed for TODO validations
	data.NeedsParseResult = false
	data.NeedsObjectResult = false
	data.NeedsFilterResult = false

	// Execute template
	tmpl, err := template.New("testcase").Parse(testCaseTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse test case template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute test case template: %w", err)
	}

	return buf.String(), nil
}

// generateValidations creates validation assertions for a test case
func (g *Generator) generateValidations(validations *types.ValidationSet) ([]string, error) {
	if validations == nil {
		return []string{}, nil
	}
	var assertions []string

	// Handle each validation type
	if validations.Parse != nil {
		assertion, err := g.generateParseValidation(validations.Parse)
		if err != nil {
			return nil, err
		}
		if assertion != "" {
			assertions = append(assertions, assertion)
		}
	}

	if validations.Filter != nil {
		assertion, err := g.generateFilterValidation(validations.Filter)
		if err != nil {
			return nil, err
		}
		if assertion != "" {
			assertions = append(assertions, assertion)
		}
	}

	if validations.BuildHierarchy != nil {
		assertion, err := g.generateBuildHierarchyValidation(validations.BuildHierarchy)
		if err != nil {
			return nil, err
		}
		if assertion != "" {
			assertions = append(assertions, assertion)
		}
	}

	if validations.GetString != nil {
		assertion, err := g.generateTypedAccessValidation("GetString", validations.GetString)
		if err != nil {
			return nil, err
		}
		if assertion != "" {
			assertions = append(assertions, assertion)
		}
	}

	if validations.GetInt != nil {
		assertion, err := g.generateTypedAccessValidation("GetInt", validations.GetInt)
		if err != nil {
			return nil, err
		}
		if assertion != "" {
			assertions = append(assertions, assertion)
		}
	}

	if validations.GetBool != nil {
		assertion, err := g.generateTypedAccessValidation("GetBool", validations.GetBool)
		if err != nil {
			return nil, err
		}
		if assertion != "" {
			assertions = append(assertions, assertion)
		}
	}

	if validations.GetFloat != nil {
		assertion, err := g.generateTypedAccessValidation("GetFloat", validations.GetFloat)
		if err != nil {
			return nil, err
		}
		if assertion != "" {
			assertions = append(assertions, assertion)
		}
	}

	return assertions, nil
}

// generateParseValidation creates assertion for parse validation
func (g *Generator) generateParseValidation(validation interface{}) (string, error) {
	// Try to parse as array of entries (simple format without count field)
	if entries, ok := g.parseAsEntryArray(validation); ok {
		inputVar := "input"
		// Special case for empty input
		if len(entries) == 0 {
			inputVar = `""`
		}
		return fmt.Sprintf(`// Parse validation
	parseResult, err = ccl.Parse(%s)
	require.NoError(t, err)
	expectedParse := %s
	assert.Equal(t, expectedParse, parseResult)`, inputVar, formatEntryArray(entries)), nil
	}

	// Try to parse as validation with expected field
	if countedValidation, ok := g.parseAsCountedValidation(validation); ok {
		inputVar := "input"
		if len(countedValidation.Expected) == 0 {
			inputVar = `""`
		}
		return fmt.Sprintf(`// Parse validation
	parseResult, err = ccl.Parse(%s)
	require.NoError(t, err)
	expectedParse := %s
	assert.Equal(t, expectedParse, parseResult)`, inputVar, formatEntryArray(countedValidation.Expected)), nil
	}

	// Try to parse as validation with count or error format
	return g.generateComplexValidation("Parse", validation)
}

// generateFilterValidation creates assertion for filter validation
func (g *Generator) generateFilterValidation(validation interface{}) (string, error) {
	// Try to parse as validation with expected field (current schema format)
	if countedValidation, ok := g.parseAsCountedValidation(validation); ok {
		inputVar := "input"
		if len(countedValidation.Expected) == 0 {
			inputVar = `""`
		}
		return fmt.Sprintf(`// Filter validation
	parseResult, err = ccl.Parse(%s)
	require.NoError(t, err)
	filterResult = ccl.Filter(parseResult)
	expectedFilter := %s
	assert.Equal(t, expectedFilter, filterResult)`, inputVar, formatEntryArray(countedValidation.Expected)), nil
	}

	// Try to parse as array of entries (legacy format)
	if entries, ok := g.parseAsEntryArray(validation); ok {
		inputVar := "input"
		// Special case for empty input
		if len(entries) == 0 {
			inputVar = `""`
		}
		return fmt.Sprintf(`// Filter validation
	parseResult, err = ccl.Parse(%s)
	require.NoError(t, err)
	filterResult = ccl.Filter(parseResult)
	expectedFilter := %s
	assert.Equal(t, expectedFilter, filterResult)`, inputVar, formatEntryArray(entries)), nil
	}

	return g.generateComplexValidation("Filter", validation)
}

// generateBuildHierarchyValidation creates assertion for make_objects validation
func (g *Generator) generateBuildHierarchyValidation(validation interface{}) (string, error) {
	// Try to parse as validation with count field first
	if countedValidation, ok := g.parseAsCountedObjectValidation(validation); ok {
		return fmt.Sprintf(`// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := %s
	assert.Equal(t, expectedObjects, objectResult)`, formatGoValue(countedValidation.Expected)), nil
	}

	// Try to parse as direct object (legacy format)
	if obj, ok := validation.(map[string]interface{}); ok {
		// Check if it has count field (which would make it a validation we missed)
		if _, hasCount := obj["count"]; hasCount {
			return g.generateComplexValidation("BuildHierarchy", validation)
		}
		return fmt.Sprintf(`// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := %s
	assert.Equal(t, expectedObjects, objectResult)`, formatGoValue(obj)), nil
	}

	return g.generateComplexValidation("BuildHierarchy", validation)
}

// generateTypedAccessValidation creates assertion for typed access validation
func (g *Generator) generateTypedAccessValidation(method string, validation interface{}) (string, error) {
	// Try to parse as single test case
	if testCase, ok := g.parseAsTypedAccessCase(validation); ok {
		if testCase.Error {
			// Generate error case
			return fmt.Sprintf(`// %s validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	_, err = ccl.%s(objectResult, %s)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "%s")`,
				method, method, formatStringArray(testCase.Args), testCase.ErrorMessage), nil
		} else {
			// Generate success case
			return fmt.Sprintf(`// %s validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	%sResult, err := ccl.%s(objectResult, %s)
	require.NoError(t, err)
	assert.Equal(t, %s, %sResult)`,
				method, strings.ToLower(method), method, formatStringArray(testCase.Args), formatGoValue(testCase.Expected), strings.ToLower(method)), nil
		}
	}

	return g.generateComplexValidation(method, validation)
}

// Helper methods

func (g *Generator) parseAsEntryArray(validation interface{}) ([]map[string]string, bool) {
	// Convert to JSON and back to handle the interface{} properly
	jsonData, err := json.Marshal(validation)
	if err != nil {
		return nil, false
	}

	var entries []map[string]string
	if err := json.Unmarshal(jsonData, &entries); err != nil {
		return nil, false
	}

	// Validate that all entries have key and value
	for _, entry := range entries {
		if _, hasKey := entry["key"]; !hasKey {
			return nil, false
		}
		if _, hasValue := entry["value"]; !hasValue {
			return nil, false
		}
	}

	return entries, true
}

type TypedAccessCase struct {
	Args         []string
	Expected     interface{}
	Error        bool
	ErrorMessage string
}

// CountedValidation represents a validation with count field
type CountedValidation struct {
	Count    int                 `json:"count"`
	Expected []map[string]string `json:"expected"`
}

// CountedObjectValidation represents an object validation with count field
type CountedObjectValidation struct {
	Count    int                    `json:"count"`
	Expected map[string]interface{} `json:"expected"`
}

// CountedTypedAccessValidation represents typed access validation with count field
type CountedTypedAccessValidation struct {
	Count int                      `json:"count"`
	Cases []map[string]interface{} `json:"cases"`
}

func (g *Generator) parseAsCountedValidation(validation interface{}) (*CountedValidation, bool) {
	jsonData, err := json.Marshal(validation)
	if err != nil {
		return nil, false
	}

	var counted CountedValidation
	if err := json.Unmarshal(jsonData, &counted); err != nil {
		return nil, false
	}

	// Validate that it has the required fields
	if counted.Count == 0 && len(counted.Expected) == 0 {
		return nil, false
	}

	return &counted, true
}

func (g *Generator) parseAsCountedObjectValidation(validation interface{}) (*CountedObjectValidation, bool) {
	jsonData, err := json.Marshal(validation)
	if err != nil {
		return nil, false
	}

	var counted CountedObjectValidation
	if err := json.Unmarshal(jsonData, &counted); err != nil {
		return nil, false
	}

	// Validate that it has the required fields
	if counted.Count == 0 && len(counted.Expected) == 0 {
		return nil, false
	}

	return &counted, true
}

func (g *Generator) parseAsCountedTypedAccessValidation(validation interface{}) (*CountedTypedAccessValidation, bool) {
	jsonData, err := json.Marshal(validation)
	if err != nil {
		return nil, false
	}

	var counted CountedTypedAccessValidation
	if err := json.Unmarshal(jsonData, &counted); err != nil {
		return nil, false
	}

	// Validate that it has the required fields
	if counted.Count == 0 && len(counted.Cases) == 0 {
		return nil, false
	}

	return &counted, true
}

func (g *Generator) parseAsTypedAccessCase(validation interface{}) (*TypedAccessCase, bool) {
	jsonData, err := json.Marshal(validation)
	if err != nil {
		return nil, false
	}

	var testCase struct {
		Args         []string    `json:"args"`
		Expected     interface{} `json:"expected"`
		Error        bool        `json:"error"`
		ErrorMessage string      `json:"error_message"`
	}
	if err := json.Unmarshal(jsonData, &testCase); err != nil {
		return nil, false
	}

	if len(testCase.Args) == 0 {
		return nil, false
	}

	return &TypedAccessCase{
		Args:         testCase.Args,
		Expected:     testCase.Expected,
		Error:        testCase.Error,
		ErrorMessage: testCase.ErrorMessage,
	}, true
}

func (g *Generator) generateComplexValidation(method string, validation interface{}) (string, error) {
	// For now, generate a placeholder for complex validations
	// Escape any problematic characters in the validation data
	validationStr := fmt.Sprintf("%v", validation)
	// Replace newlines and other problematic characters in comments
	validationStr = strings.ReplaceAll(validationStr, "\n", "\\n")
	validationStr = strings.ReplaceAll(validationStr, "\r", "\\r")
	validationStr = strings.ReplaceAll(validationStr, "\t", "\\t")

	return fmt.Sprintf(`// TODO: Implement %s validation
	// Validation data: %s`, method, validationStr), nil
}

// Utility functions

func toPascalCase(input string) string {
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})

	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}

	return strings.Join(parts, "")
}

func escapeGoString(s string) string {
	if s == "" {
		return `""`
	}
	// If string contains CRLF sequences, use quoted string to preserve them
	if strings.Contains(s, "\r\n") {
		// Use %q to create a properly escaped quoted string
		return fmt.Sprintf("%q", s)
	}
	return fmt.Sprintf("`%s`", s)
}

// generateFlatFormatValidation creates validation code for flat format tests
func (g *Generator) generateFlatFormatValidation(test types.TestCase) (string, error) {
	switch test.Validation {
	case "parse":
		return g.generateFlatParseValidation(test)
	case "build_hierarchy":
		return g.generateFlatBuildHierarchyValidation(test)
	case "get_string", "get_int", "get_bool", "get_float", "get_list":
		return g.generateFlatTypedAccessValidation(test, test.Validation)
	default:
		// For unimplemented validations, generate a safe comment with variable usage
		// Handle both single-input and multi-input tests
		var inputVarLines string
		if len(test.Inputs) == 1 {
			inputVarLines = "\t_ = input // Prevent unused variable warning"
		} else if len(test.Inputs) > 1 {
			var lines []string
			for i := range test.Inputs {
				lines = append(lines, fmt.Sprintf("\t_ = input%d // Prevent unused variable warning", i))
			}
			inputVarLines = strings.Join(lines, "\n")
		}
		return fmt.Sprintf(`// TODO: Implement %s validation
	_ = ccl // Prevent unused variable warning
%s
	_ = err // Prevent unused variable warning`, test.Validation, inputVarLines), nil
	}
}

// generateFlatParseValidation generates parse validation for flat format
func (g *Generator) generateFlatParseValidation(test types.TestCase) (string, error) {
	// Handle case where Expected is directly an array of entries (loader returns this format)
	if entriesArray, ok := test.Expected.([]interface{}); ok {
		// Convert to Go-formatted entry array
		var goEntries []string
		for _, entry := range entriesArray {
			entryMap, ok := entry.(map[string]interface{})
			if !ok {
				continue
			}
			key, _ := entryMap["key"].(string)
			value, _ := entryMap["value"].(string)
			goEntries = append(goEntries, fmt.Sprintf(`mock.Entry{Key: %q, Value: %q}`, key, value))
		}

		entryArrayStr := "[]mock.Entry{" + strings.Join(goEntries, ", ") + "}"

		return fmt.Sprintf(`// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := %s
	assert.Equal(t, expected, parseResult)`, entryArrayStr), nil
	}

	// Handle case where Expected is a map with count/entries/error fields (JSON schema format)
	expectedMap, ok := test.Expected.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("expected map or array for parse validation, got %T", test.Expected)
	}

	// Check for error expectation first
	if errorExpected, ok := expectedMap["error"]; ok {
		if errorBool, ok := errorExpected.(bool); ok && errorBool {
			return `// Parse validation (expects error)
	_, err := ccl.Parse(input)
	require.Error(t, err)`, nil
		}
	}

	// Handle normal case with entries or empty result
	if entries, ok := expectedMap["entries"]; ok {
		entriesArray, ok := entries.([]interface{})
		if !ok {
			return "", fmt.Errorf("expected entries array, got %T", entries)
		}

		// Convert to Go-formatted entry array
		var goEntries []string
		for _, entry := range entriesArray {
			entryMap, ok := entry.(map[string]interface{})
			if !ok {
				continue
			}
			key, _ := entryMap["key"].(string)
			value, _ := entryMap["value"].(string)
			goEntries = append(goEntries, fmt.Sprintf(`mock.Entry{Key: %q, Value: %q}`, key, value))
		}

		entryArrayStr := "[]mock.Entry{" + strings.Join(goEntries, ", ") + "}"

		return fmt.Sprintf(`// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := %s
	assert.Equal(t, expected, parseResult)`, entryArrayStr), nil
	} else {
		// Handle case with only count (empty result) - schema says count is always required
		return `// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)`, nil
	}
}

// generateFlatBuildHierarchyValidation generates build_hierarchy validation for flat format
func (g *Generator) generateFlatBuildHierarchyValidation(test types.TestCase) (string, error) {
	// For build_hierarchy, expected is usually a nested object
	expectedMap, ok := test.Expected.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("expected map for build_hierarchy validation, got %T", test.Expected)
	}

	// Check for error expectation first
	if errorExpected, ok := expectedMap["error"]; ok {
		if errorBool, ok := errorExpected.(bool); ok && errorBool {
			return `// BuildHierarchy validation (expects error)
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	_, err = ccl.BuildHierarchy(parseResult)
	require.Error(t, err)`, nil
		}
	}

	// Handle normal case with object result
	if object, ok := expectedMap["object"]; ok {
		return fmt.Sprintf(`// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := %s
	assert.Equal(t, expected, objectResult)`, formatGoValue(object)), nil
	} else {
		// Handle case with only count (empty result)
		return `// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)`, nil
	}
}

// generateFlatTypedAccessValidation generates typed access validation for flat format
func (g *Generator) generateFlatTypedAccessValidation(test types.TestCase, validation string) (string, error) {
	// Handle case where Expected is directly the value (loader returns this format for typed access)
	if test.Expected != nil {
		// Check if it's a simple value (string, int, bool, float, or array for lists)
		switch v := test.Expected.(type) {
		case string, int, float64, bool, []interface{}:
			// Use args if provided, otherwise use a default key
			args := test.Args
			if len(args) == 0 {
				args = []string{"title"} // Common default for typed access tests
			}

			return g.generateTypedAccessForDirectValue(validation, args, v)
		}
	}

	// Handle case where Expected is a map with value/error/count fields (JSON schema format)
	expectedMap, ok := test.Expected.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("expected must be a map or simple value for typed access validation, got %T", test.Expected)
	}

	// Check for error expectation first
	if errorExpected, ok := expectedMap["error"]; ok {
		if errorBool, ok := errorExpected.(bool); ok && errorBool {
			// Use args if provided, otherwise use a default key
			args := test.Args
			if len(args) == 0 {
				args = []string{"nonexistent"}
			}

			return fmt.Sprintf(`// %s validation (expects error)
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	_, err = ccl.%s(hierarchy, %s)
	require.Error(t, err)`, validation, validation, formatArgs(args)), nil
		}
	}

	// Determine args and expected value based on schema
	var args []string
	var expectedValue interface{}

	// Use explicit args if provided
	if len(test.Args) > 0 {
		args = test.Args
	} else {
		// Infer args based on expected result fields
		if _, ok := expectedMap["value"]; ok {
			args = []string{"key"} // Generic fallback for single values
		} else if _, ok := expectedMap["list"]; ok {
			args = []string{"servers"} // Common list key
		} else {
			// For cases with only count=0, this is testing non-existent keys
			args = []string{"nonexistent"}
		}
	}

	// Get expected value from appropriate field
	if value, ok := expectedMap["value"]; ok {
		expectedValue = value
	} else if list, ok := expectedMap["list"]; ok {
		expectedValue = list
	} else {
		// Only count field - this means empty/error result
		expectedValue = nil
	}

	var template string
	switch validation {
	case "get_string":
		if expectedValue == nil {
			template = `	result, err := ccl.GetString(hierarchy, %s)
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, "", result)
	}`
		} else {
			template = `	result, err := ccl.GetString(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
		}
	case "get_int":
		if expectedValue == nil {
			template = `	result, err := ccl.GetInt(hierarchy, %s)
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, 0, result)
	}`
		} else {
			template = `	result, err := ccl.GetInt(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
		}
	case "get_bool":
		if expectedValue == nil {
			template = `	result, err := ccl.GetBool(hierarchy, %s)
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}`
		} else {
			template = `	result, err := ccl.GetBool(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
		}
	case "get_float":
		if expectedValue == nil {
			template = `	result, err := ccl.GetFloat(hierarchy, %s)
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, 0.0, result)
	}`
		} else {
			template = `	result, err := ccl.GetFloat(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
		}
	case "get_list":
		if expectedValue == nil {
			template = `	result, err := ccl.GetList(hierarchy, %s)
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Empty(t, result)
	}`
			return fmt.Sprintf(`// %s validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
%s`, validation, fmt.Sprintf(template, formatArgs(args))), nil
		} else {
			template = `	result, err := ccl.GetList(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
		}
	default:
		return "", fmt.Errorf("unknown typed access validation: %s", validation)
	}

	if expectedValue == nil {
		return fmt.Sprintf(`// %s validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
%s`, validation, fmt.Sprintf(template, formatArgs(args))), nil
	}

	return fmt.Sprintf(`// %s validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
%s`, validation, fmt.Sprintf(template, formatArgs(args), expectedValue)), nil
}

// generateTypedAccessForDirectValue generates validation for when the expected value is returned directly
func (g *Generator) generateTypedAccessForDirectValue(validation string, args []string, expectedValue interface{}) (string, error) {
	var template string
	switch validation {
	case "get_string":
		template = `	result, err := ccl.GetString(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
	case "get_int":
		template = `	result, err := ccl.GetInt(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
	case "get_bool":
		template = `	result, err := ccl.GetBool(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
	case "get_float":
		template = `	result, err := ccl.GetFloat(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
	case "get_list":
		template = `	result, err := ccl.GetList(hierarchy, %s)
	require.NoError(t, err)
	assert.Equal(t, %#v, result)`
	default:
		return "", fmt.Errorf("unknown typed access validation: %s", validation)
	}

	return fmt.Sprintf(`// %s validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
%s`, validation, fmt.Sprintf(template, formatArgs(args), expectedValue)), nil
}

// getTestTags converts flat format test fields to structured tags for filtering
func (g *Generator) getTestTags(test types.TestCase) []string {
	var tags []string

	// Convert Functions to function: tags
	for _, fn := range test.Functions {
		tags = append(tags, "function:"+fn)
	}

	// Convert Features to feature: tags
	for _, feat := range test.Features {
		tags = append(tags, "feature:"+feat)
	}

	// Convert Behaviors to behavior: tags
	for _, behav := range test.Behaviors {
		tags = append(tags, "behavior:"+behav)
	}

	// Convert Variants to variant: tags
	for _, variant := range test.Variants {
		tags = append(tags, "variant:"+variant)
	}

	// Also include any existing Meta.Tags
	tags = append(tags, test.Meta.Tags...)

	return tags
}

// shouldSkipTest determines if a test should be skipped based on tags and generator options
func (g *Generator) shouldSkipTest(tags []string) bool {
	// If runOnly is specified, only run tests with those tags
	if len(g.options.RunOnly) > 0 {
		hasRunOnlyTag := false
		for _, runTag := range g.options.RunOnly {
			for _, tag := range tags {
				if tag == runTag {
					hasRunOnlyTag = true
					break
				}
			}
			if hasRunOnlyTag {
				break
			}
		}
		if !hasRunOnlyTag {
			return true
		}
	}

	// Check additional skip tags from options
	for _, skipTag := range g.options.SkipTags {
		for _, tag := range tags {
			if tag == skipTag {
				return true
			}
		}
	}

	// Skip disabled features if option is enabled
	if g.options.SkipDisabled {
		for _, tag := range tags {
			// Skip tests that need specific features not implemented in mock
			if strings.HasPrefix(tag, "needs-") {
				return true
			}
			// Skip complex parsing features not in our simple mock
			if tag == "multiline" || tag == "nested" || tag == "continuation" {
				return true
			}
			// Skip error handling tests as our mock doesn't handle errors properly
			if tag == "error" || tag == "incomplete" {
				return true
			}
			// Skip advanced features
			if tag == "whitespace" || tag == "tabs" || tag == "newlines" {
				return true
			}
			// Skip algebraic and property tests (advanced features)
			if tag == "algebraic" || tag == "round-trip" || tag == "semigroup" || tag == "monoid" {
				return true
			}
		}
	}
	return false
}

// shouldSkipTestByName determines if a test should be skipped based on tags, name, and generator options
// shouldSkipTestByName determines if a test should be skipped based on tags, name, and generator options
// shouldSkipTestByName determines if a test should be skipped based on name first, then tags and generator options
func (g *Generator) shouldSkipTestByName(testName string, tags []string) bool {
	// Check if test name is in skip list FIRST (highest precedence)
	for _, skipName := range g.options.SkipTestsByName {
		if testName == skipName {
			return true
		}
	}

	// Delegate to existing tag-based filtering
	return g.shouldSkipTest(tags)
}

// getSkipReason determines the reason for skipping a test based on its tags and options
func (g *Generator) getSkipReason(tags []string) string {
	// Check if skipped due to run-only filter
	if len(g.options.RunOnly) > 0 {
		hasRunOnlyTag := false
		for _, runTag := range g.options.RunOnly {
			for _, tag := range tags {
				if tag == runTag {
					hasRunOnlyTag = true
					break
				}
			}
			if hasRunOnlyTag {
				break
			}
		}
		if !hasRunOnlyTag {
			return fmt.Sprintf("Test does not match run-only filter: %v", g.options.RunOnly)
		}
	}

	// Check if skipped due to additional skip tags
	for _, skipTag := range g.options.SkipTags {
		for _, tag := range tags {
			if tag == skipTag {
				return fmt.Sprintf("Test skipped due to tag filter: %s", skipTag)
			}
		}
	}

	// Default skip reasons for disabled features
	for _, tag := range tags {
		if strings.HasPrefix(tag, "needs-") {
			return fmt.Sprintf("Requires feature: %s", strings.TrimPrefix(tag, "needs-"))
		}
		// Provide specific skip reasons for different categories
		switch tag {
		case "multiline", "nested", "continuation":
			return "Complex parsing not implemented in mock CCL"
		case "error", "incomplete":
			return "Error handling not implemented in mock CCL"
		case "whitespace", "tabs", "newlines":
			return "Whitespace handling not fully implemented in mock CCL"
		case "algebraic", "round-trip", "semigroup", "monoid":
			return "Advanced algebraic properties not implemented in mock CCL"
		}
	}
	return "Skipped test"
}

// getSkipReasonByName determines the reason for skipping a test based on name, tags and options
func (g *Generator) getSkipReasonByName(testName string, tags []string) string {
	// Check if skipped by name first
	for _, skipName := range g.options.SkipTestsByName {
		if testName == skipName {
			return fmt.Sprintf("Test skipped by name filter: %s", skipName)
		}
	}

	// Delegate to existing tag-based skip reason logic
	return g.getSkipReason(tags)
}

// hasValidations checks if the ValidationSet has any non-nil validations
func hasValidations(validations types.ValidationSet) bool {
	return validations.Parse != nil ||
		validations.ParseIndented != nil ||
		validations.Filter != nil ||
		validations.Combine != nil ||
		validations.ExpandDotted != nil ||
		validations.BuildHierarchy != nil ||
		validations.GetString != nil ||
		validations.GetInt != nil ||
		validations.GetBool != nil ||
		validations.GetFloat != nil ||
		validations.GetList != nil ||
		validations.PrettyPrint != nil ||
		validations.RoundTrip != nil ||
		validations.ComposeAssociative != nil ||
		validations.IdentityLeft != nil ||
		validations.IdentityRight != nil
}

// hasImplementedValidations checks if the ValidationSet has any validations that generate actual assertions (not just TODOs)
func (g *Generator) hasImplementedValidations(validations *types.ValidationSet) bool {
	if validations == nil {
		return false
	}
	// Only check validation types that are actually implemented in the generator
	return validations.Parse != nil ||
		validations.Filter != nil ||
		validations.BuildHierarchy != nil ||
		validations.GetString != nil ||
		validations.GetInt != nil ||
		validations.GetBool != nil ||
		validations.GetFloat != nil
	// Note: Other validation types (CanonicalFormat/PrettyPrint, RoundTrip, etc.) are not implemented
	// and only generate TODO comments, so they don't count as "implemented validations"
}

func formatEntryArray(entries []map[string]string) string {
	var parts []string
	for _, entry := range entries {
		parts = append(parts, fmt.Sprintf(`mock.Entry{Key: %q, Value: %q}`, entry["key"], entry["value"]))
	}
	return fmt.Sprintf("[]mock.Entry{%s}", strings.Join(parts, ", "))
}

func formatStringArray(arr []string) string {
	var parts []string
	for _, s := range arr {
		parts = append(parts, fmt.Sprintf("%q", s))
	}
	return fmt.Sprintf("[]string{%s}", strings.Join(parts, ", "))
}

func formatGoValue(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return "nil"
	case string:
		return fmt.Sprintf("%q", v)
	case int, int64, float64:
		return fmt.Sprintf("%v", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case map[string]interface{}:
		return formatGoMap(v)
	case []interface{}:
		return formatGoSlice(v)
	default:
		if value == nil {
			return "nil"
		}
		return fmt.Sprintf("%#v", v)
	}
}

func formatGoMap(m map[string]interface{}) string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	// Sort keys to ensure deterministic output
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%q: %s", k, formatGoValue(m[k])))
	}
	return fmt.Sprintf("map[string]interface{}{%s}", strings.Join(parts, ", "))
}

func formatGoSlice(s []interface{}) string {
	var parts []string
	for _, v := range s {
		parts = append(parts, formatGoValue(v))
	}
	return fmt.Sprintf("[]interface{}{%s}", strings.Join(parts, ", "))
}

func formatArgs(args []string) string {
	if len(args) == 0 {
		return "[]string{}"
	}
	// Always format as a slice since CCL functions expect []string
	var parts []string
	for _, arg := range args {
		parts = append(parts, fmt.Sprintf("%q", arg))
	}
	return fmt.Sprintf("[]string{%s}", strings.Join(parts, ", "))
}

// Helper functions to determine which variables are needed

func (g *Generator) needsParseResult(validations *types.ValidationSet) bool {
	if validations == nil {
		return false
	}
	return validations.Parse != nil ||
		validations.Filter != nil ||
		validations.BuildHierarchy != nil ||
		validations.GetString != nil ||
		validations.GetInt != nil ||
		validations.GetBool != nil ||
		validations.GetFloat != nil
}

func (g *Generator) needsObjectResult(validations *types.ValidationSet) bool {
	if validations == nil {
		return false
	}
	return validations.BuildHierarchy != nil ||
		validations.GetString != nil ||
		validations.GetInt != nil ||
		validations.GetBool != nil ||
		validations.GetFloat != nil
}

func (g *Generator) needsFilterResult(validations *types.ValidationSet) bool {
	if validations == nil {
		return false
	}
	return validations.Filter != nil
}

// countAssertions counts the total number of assertions for a validation set
func (g *Generator) countAssertions(validations *types.ValidationSet) int {
	if validations == nil {
		return 0
	}
	count := 0

	if validations.Parse != nil {
		count += g.getValidationCount(validations.Parse)
	}
	if validations.ParseIndented != nil {
		count += g.getValidationCount(validations.ParseIndented)
	}
	if validations.Filter != nil {
		count += g.getValidationCount(validations.Filter)
	}
	if validations.Combine != nil {
		count += g.getValidationCount(validations.Combine)
	}
	if validations.ExpandDotted != nil {
		count += g.getValidationCount(validations.ExpandDotted)
	}
	if validations.BuildHierarchy != nil {
		count += g.getValidationCount(validations.BuildHierarchy)
	}
	if validations.GetString != nil {
		count += g.getValidationCount(validations.GetString)
	}
	if validations.GetInt != nil {
		count += g.getValidationCount(validations.GetInt)
	}
	if validations.GetBool != nil {
		count += g.getValidationCount(validations.GetBool)
	}
	if validations.GetFloat != nil {
		count += g.getValidationCount(validations.GetFloat)
	}
	if validations.GetList != nil {
		count += g.getValidationCount(validations.GetList)
	}
	if validations.PrettyPrint != nil {
		count += g.getValidationCount(validations.PrettyPrint)
	}
	if validations.RoundTrip != nil {
		count += g.getValidationCount(validations.RoundTrip)
	}
	if validations.ComposeAssociative != nil {
		count += g.getValidationCount(validations.ComposeAssociative)
	}
	if validations.IdentityLeft != nil {
		count += g.getValidationCount(validations.IdentityLeft)
	}
	if validations.IdentityRight != nil {
		count += g.getValidationCount(validations.IdentityRight)
	}

	return count
}

// getValidationCount extracts the count from a validation, defaulting to 1
func (g *Generator) getValidationCount(validation interface{}) int {
	// Try to parse as validation with count field
	jsonData, err := json.Marshal(validation)
	if err != nil {
		return 1 // Default to 1 assertion
	}

	// Check if this has an explicit count field
	var checkMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &checkMap); err == nil {
		if countVal, hasCount := checkMap["count"]; hasCount {
			// Has explicit count field - use its value
			if countFloat, ok := countVal.(float64); ok {
				return int(countFloat)
			}
		}
	}

	// For typed access validations, check if it's an array (legacy format)
	var arr []interface{}
	if err := json.Unmarshal(jsonData, &arr); err == nil {
		return len(arr) // Each element in array is one assertion
	}

	// Default to 1 assertion for simple validations
	return 1
}
