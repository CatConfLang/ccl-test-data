package parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_edge_cases.json
// Suite: Flat Format
// Version: 1.0



// basic_single_no_spaces_parse - function:parse
func TestBasicSingleNoSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key=val`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// basic_with_spaces_parse - function:parse
func TestBasicWithSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key = val`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// indented_key_parse_value - function:parse_value
func TestIndentedKeyParseValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `  key = val`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement parse_value validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// value_trailing_spaces_parse - function:parse
func TestValueTrailingSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key = val  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// key_value_surrounded_spaces_parse - function:parse
func TestKeyValueSurroundedSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `  key  =  val  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// surrounded_by_newlines_parse - function:parse
func TestSurroundedByNewlinesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `
key = val
`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// key_empty_value_parse - function:parse
func TestKeyEmptyValueParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_value_with_newline_parse - function:parse
func TestEmptyValueWithNewlineParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =
`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_value_with_spaces_parse - function:parse
func TestEmptyValueWithSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_key_indented_parse_value - function:parse_value
func TestEmptyKeyIndentedParseValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `  = val`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement parse_value validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// empty_key_with_newline_parse - function:parse
func TestEmptyKeyWithNewlineParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `
  = val`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// empty_key_value_with_spaces_parse - function:parse
func TestEmptyKeyValueWithSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `  =  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// equals_in_value_no_spaces_parse - function:parse
func TestEqualsInValueNoSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a=b=c`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "b=c"}}
	assert.Equal(t, expected, parseResult)

}


// equals_in_value_with_spaces_parse - function:parse
func TestEqualsInValueWithSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a = b = c`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "b = c"}}
	assert.Equal(t, expected, parseResult)

}


// multiple_key_value_pairs_parse - function:parse
func TestMultipleKeyValuePairsParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key1 = val1
key2 = val2`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "val1"}, mock.Entry{Key: "key2", Value: "val2"}}
	assert.Equal(t, expected, parseResult)

}


// key_with_tabs_parse - function:parse behavior:tabs_preserve
func TestKeyWithTabsParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `	key	=	value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\tvalue"}}
	assert.Equal(t, expected, parseResult)

}


// key_with_tabs_ocaml_reference_parse - function:parse behavior:tabs_preserve
func TestKeyWithTabsOcamlReferenceParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `	key	=	value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// whitespace_only_value_parse - function:parse
func TestWhitespaceOnlyValueParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `onlyspaces =     `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "onlyspaces", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// spaces_vs_tabs_continuation_parse_value - function:parse_value behavior:tabs_preserve
func TestSpacesVsTabsContinuationParseValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `text = First
    four spaces
	tab preserved`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement parse_value validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// spaces_vs_tabs_continuation_ocaml_reference_parse_value - function:parse_value behavior:tabs_preserve
func TestSpacesVsTabsContinuationOcamlReferenceParseValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `text = First
    four spaces
	tab preserved`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement parse_value validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// multiple_empty_equality_parse - function:parse
func TestMultipleEmptyEqualityParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := ` =  = `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "="}}
	assert.Equal(t, expected, parseResult)

}


// key_with_newline_before_equals_parse - function:parse
func TestKeyWithNewlineBeforeEqualsParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key 
= val
`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// complex_multi_newline_whitespace_parse - function:parse
func TestComplexMultiNewlineWhitespaceParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `  
 key  
=  val  
`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// empty_value_with_trailing_spaces_newline_parse - function:parse
func TestEmptyValueWithTrailingSpacesNewlineParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =  
`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_key_value_with_surrounding_newlines_parse - function:parse
func TestEmptyKeyValueWithSurroundingNewlinesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `
  =  
`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// quotes_treated_as_literal_unquoted_parse - function:parse
func TestQuotesTreatedAsLiteralUnquotedParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}}
	assert.Equal(t, expected, parseResult)

}


// quotes_treated_as_literal_quoted_parse - function:parse
func TestQuotesTreatedAsLiteralQuotedParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = "localhost"`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "\"localhost\""}}
	assert.Equal(t, expected, parseResult)

}


// nested_single_line_parse - function:parse
func TestNestedSingleLineParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =
  val`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\n  val"}}
	assert.Equal(t, expected, parseResult)

}


// nested_multi_line_parse - function:parse
func TestNestedMultiLineParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =
  line1
  line2`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\n  line1\n  line2"}}
	assert.Equal(t, expected, parseResult)

}


// nested_with_blank_line_parse_value - function:parse_value
func TestNestedWithBlankLineParseValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =
  line1

  line2`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement parse_value validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// deep_nested_structure_parse_value - function:parse_value
func TestDeepNestedStructureParseValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =
  field1 = value1
  field2 =
    subfield = x
    another = y`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement parse_value validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// realistic_stress_test_parse - function:parse
func TestRealisticStressTestParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Dmitrii Kovanikov
login = chshersh
language = OCaml
date = 2024-05-25`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Dmitrii Kovanikov"}, mock.Entry{Key: "login", Value: "chshersh"}, mock.Entry{Key: "language", Value: "OCaml"}, mock.Entry{Key: "date", Value: "2024-05-25"}}
	assert.Equal(t, expected, parseResult)

}


// ocaml_stress_test_original_parse - function:parse
func TestOcamlStressTestOriginalParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/= This is a CCL document
title = CCL Example

database =
  enabled = true
  ports =
    = 8000
    = 8001
    = 8002
  limits =
    cpu = 1500mi
    memory = 10Gb

user =
  guestId = 42

user =
  login = chshersh
  createdAt = 2024-12-31`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "This is a CCL document"}, mock.Entry{Key: "title", Value: "CCL Example"}, mock.Entry{Key: "database", Value: "\n  enabled = true\n  ports =\n    = 8000\n    = 8001\n    = 8002\n  limits =\n    cpu = 1500mi\n    memory = 10Gb"}, mock.Entry{Key: "user", Value: "\n  guestId = 42"}, mock.Entry{Key: "user", Value: "\n  login = chshersh\n  createdAt = 2024-12-31"}}
	assert.Equal(t, expected, parseResult)

}


// ocaml_stress_test_original_build_hierarchy - function:build_hierarchy
func TestOcamlStressTestOriginalBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/= This is a CCL document
title = CCL Example

database =
  enabled = true
  ports =
    = 8000
    = 8001
    = 8002
  limits =
    cpu = 1500mi
    memory = 10Gb

user =
  guestId = 42

user =
  login = chshersh
  createdAt = 2024-12-31`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// ocaml_stress_test_original_get_string - function:get_string
func TestOcamlStressTestOriginalGetString(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/= This is a CCL document
title = CCL Example

database =
  enabled = true
  ports =
    = 8000
    = 8001
    = 8002
  limits =
    cpu = 1500mi
    memory = 10Gb

user =
  guestId = 42

user =
  login = chshersh
  createdAt = 2024-12-31`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"title"})
	require.NoError(t, err)
	assert.Equal(t, "CCL Example", result)

}


