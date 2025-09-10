package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/ccl-test-data/test-runner/internal/types"
)

const testFileTemplate = `package {{.PackageName}}_test

import (
	"testing"
	{{if .HasActiveTests}}
	"github.com/ccl-test-data/test-runner/internal/mock"{{if .HasAssertions}}
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

const testCaseTemplate = `// {{.Name}} - {{.TagsString}} (level {{.Level}})
func Test{{.TestFuncName}}(t *testing.T) {
	{{if .ShouldSkip}}t.Skip("{{.SkipReason}}"){{else}}
	
	ccl := mock.New()
	{{if .Input}}input := {{.InputString}}{{end}}
	{{if .Input1}}input1 := {{.Input1String}}{{end}}
	{{if .Input2}}input2 := {{.Input2String}}{{end}}
	{{if .Input3}}input3 := {{.Input3String}}{{end}}
	
	{{if .HasValidations}}// Declare variables for reuse across validations
	{{if .NeedsParseResult}}var parseResult []mock.Entry{{end}}
	{{if .NeedsObjectResult}}var objectResult map[string]interface{}{{end}}
	{{if .NeedsFilterResult}}var filterResult []mock.Entry{{end}}
	var err error
	{{end}}
{{range .Validations}}	{{.}}
{{end}}{{if not .HasValidations}}	// TODO: Implement test validations
	_ = ccl // Prevent unused variable warning
{{if .Input}}	_ = input // Prevent unused variable warning
{{end}}{{end}}{{end}}
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
	Level             int
	ShouldSkip        bool
	SkipReason        string
	Input             string
	InputString       string
	Input1            string
	Input1String      string
	Input2            string
	Input2String      string
	Input3            string
	Input3String      string
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
		
		// Count assertions and track statistics
		assertionCount := g.countAssertions(test.Validations)
		g.stats.TestCounts[test.Name] = assertionCount
		g.stats.TotalTests++
		
		// Check if this test is not skipped using generator options
		isSkipped := g.shouldSkipTest(test.Meta.Tags)
		if isSkipped {
			g.stats.SkippedTests++
			g.stats.SkippedAssertions += assertionCount
		} else {
			hasActiveTests = true
			g.stats.TotalAssertions += assertionCount
			// Check if this test has assertions (validations)
			if hasValidations(test.Validations) {
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
	data := TestCaseData{
		Name:         test.Name,
		TestFuncName: toPascalCase(test.Name),
		TagsString:   strings.Join(test.Meta.Tags, " "),
		Level:        test.Meta.Level,
		Input:        test.Input,
		InputString:  escapeGoString(test.Input),
		Input1:       test.Input1,
		Input1String: escapeGoString(test.Input1),
		Input2:       test.Input2,
		Input2String: escapeGoString(test.Input2),
		Input3:       test.Input3,
		Input3String: escapeGoString(test.Input3),
	}

	// Check if test should be skipped using generator options
	if g.shouldSkipTest(test.Meta.Tags) {
		data.ShouldSkip = true
		data.SkipReason = g.getSkipReason(test.Meta.Tags)
	}

	// Generate validation assertions
	validations, err := g.generateValidations(test.Validations)
	if err != nil {
		return "", fmt.Errorf("failed to generate validations: %w", err)
	}
	data.Validations = validations
	
	// Check if there are actual implemented validations (not just TODOs)
	hasImplementedValidations := false
	for _, validation := range validations {
		if !strings.Contains(validation, "// TODO: Implement") {
			hasImplementedValidations = true
			break
		}
	}
	
	data.HasValidations = hasImplementedValidations
	
	// Determine which variables are needed (only if there are implemented validations)
	if hasImplementedValidations {
		data.NeedsParseResult = g.needsParseResult(test.Validations)
		data.NeedsObjectResult = g.needsObjectResult(test.Validations)
		data.NeedsFilterResult = g.needsFilterResult(test.Validations)
	}

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
func (g *Generator) generateValidations(validations types.ValidationSet) ([]string, error) {
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

	if validations.MakeObjects != nil {
		assertion, err := g.generateMakeObjectsValidation(validations.MakeObjects)
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

	// Try to parse as counted format with expected field
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

	// Try to parse as counted format or error format
	return g.generateComplexValidation("Parse", validation)
}

// generateFilterValidation creates assertion for filter validation
func (g *Generator) generateFilterValidation(validation interface{}) (string, error) {
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

// generateMakeObjectsValidation creates assertion for make_objects validation
func (g *Generator) generateMakeObjectsValidation(validation interface{}) (string, error) {
	// Try to parse as counted format first
	if countedValidation, ok := g.parseAsCountedObjectValidation(validation); ok {
		return fmt.Sprintf(`// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := %s
	assert.Equal(t, expectedObjects, objectResult)`, formatGoValue(countedValidation.Expected)), nil
	}

	// Try to parse as direct object (legacy format)
	if obj, ok := validation.(map[string]interface{}); ok {
		// Check if it has count field (which would make it counted format we missed)
		if _, hasCount := obj["count"]; hasCount {
			return g.generateComplexValidation("MakeObjects", validation)
		}
		return fmt.Sprintf(`// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := %s
	assert.Equal(t, expectedObjects, objectResult)`, formatGoValue(obj)), nil
	}

	return g.generateComplexValidation("MakeObjects", validation)
}

// generateTypedAccessValidation creates assertion for typed access validation
func (g *Generator) generateTypedAccessValidation(method string, validation interface{}) (string, error) {
	// Try to parse as single test case
	if testCase, ok := g.parseAsTypedAccessCase(validation); ok {
		return fmt.Sprintf(`// %s validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	%sResult, err := ccl.%s(objectResult, %s)
	require.NoError(t, err)
	assert.Equal(t, %s, %sResult)`, 
			method, strings.ToLower(method), method, formatStringArray(testCase.Args), formatGoValue(testCase.Expected), strings.ToLower(method)), nil
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
	Args     []string
	Expected interface{}
}

// CountedValidation represents a validation with count field
type CountedValidation struct {
	Count    int                   `json:"count"`
	Expected []map[string]string  `json:"expected"`
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
		Args     []string    `json:"args"`
		Expected interface{} `json:"expected"`
	}
	if err := json.Unmarshal(jsonData, &testCase); err != nil {
		return nil, false
	}

	if len(testCase.Args) == 0 {
		return nil, false
	}

	return &TypedAccessCase{
		Args:     testCase.Args,
		Expected: testCase.Expected,
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
	return fmt.Sprintf("`%s`", s)
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
			// Skip algebraic and property tests (level 4+ features)
			if tag == "algebraic" || tag == "round-trip" || tag == "semigroup" || tag == "monoid" {
				return true
			}
		}
	}
	return false
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
			return "Level 4+ algebraic properties not implemented in mock CCL"
		}
	}
	return "Skipped test"
}

// hasValidations checks if the ValidationSet has any non-nil validations
func hasValidations(validations types.ValidationSet) bool {
	return validations.Parse != nil ||
		validations.ParseValue != nil ||
		validations.Filter != nil ||
		validations.Compose != nil ||
		validations.ExpandDotted != nil ||
		validations.MakeObjects != nil ||
		validations.GetString != nil ||
		validations.GetInt != nil ||
		validations.GetBool != nil ||
		validations.GetFloat != nil ||
		validations.GetList != nil ||
		validations.PrettyPrint != nil ||
		validations.RoundTrip != nil ||
		validations.Canonical != nil ||
		validations.Associativity != nil
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
	var parts []string
	for k, v := range m {
		parts = append(parts, fmt.Sprintf("%q: %s", k, formatGoValue(v)))
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

// Helper functions to determine which variables are needed

func (g *Generator) needsParseResult(validations types.ValidationSet) bool {
	return validations.Parse != nil || 
		   validations.Filter != nil || 
		   validations.MakeObjects != nil || 
		   validations.GetString != nil ||
		   validations.GetInt != nil ||
		   validations.GetBool != nil ||
		   validations.GetFloat != nil
}

func (g *Generator) needsObjectResult(validations types.ValidationSet) bool {
	return validations.MakeObjects != nil || 
		   validations.GetString != nil ||
		   validations.GetInt != nil ||
		   validations.GetBool != nil ||
		   validations.GetFloat != nil
}

func (g *Generator) needsFilterResult(validations types.ValidationSet) bool {
	return validations.Filter != nil
}

// countAssertions counts the total number of assertions for a validation set
func (g *Generator) countAssertions(validations types.ValidationSet) int {
	count := 0
	
	if validations.Parse != nil {
		count += g.getValidationCount(validations.Parse)
	}
	if validations.ParseValue != nil {
		count += g.getValidationCount(validations.ParseValue)
	}
	if validations.Filter != nil {
		count += g.getValidationCount(validations.Filter)
	}
	if validations.Compose != nil {
		count += g.getValidationCount(validations.Compose)
	}
	if validations.ExpandDotted != nil {
		count += g.getValidationCount(validations.ExpandDotted)
	}
	if validations.MakeObjects != nil {
		count += g.getValidationCount(validations.MakeObjects)
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
	if validations.Canonical != nil {
		count += g.getValidationCount(validations.Canonical)
	}
	if validations.Associativity != nil {
		count += g.getValidationCount(validations.Associativity)
	}
	
	return count
}

// getValidationCount extracts the count from a validation, defaulting to 1
func (g *Generator) getValidationCount(validation interface{}) int {
	// Try to parse as counted format
	jsonData, err := json.Marshal(validation)
	if err != nil {
		return 1 // Default to 1 assertion
	}

	// Try to parse as a structure with count field
	var countStruct struct {
		Count int `json:"count"`
	}
	if err := json.Unmarshal(jsonData, &countStruct); err == nil && countStruct.Count > 0 {
		return countStruct.Count
	}

	// For typed access validations, check if it's an array (legacy format)
	var arr []interface{}
	if err := json.Unmarshal(jsonData, &arr); err == nil {
		return len(arr) // Each element in array is one assertion
	}

	// Default to 1 assertion for simple validations
	return 1
}