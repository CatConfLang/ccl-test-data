package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_edge_cases.json
// Suite: Flat Format
// Version: 1.0

// basic_single_no_spaces_parse - function:parse (level 0)
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

// basic_with_spaces_parse - function:parse (level 0)
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

// indented_key_parsevalue - function:parsevalue (level 0)
func TestIndentedKeyParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// value_trailing_spaces_parse - function:parse (level 0)
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

// key_value_surrounded_spaces_parse - function:parse (level 0)
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

// surrounded_by_newlines_parse - function:parse (level 0)
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

// key_empty_value_parse - function:parse (level 0)
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

// empty_value_with_newline_parse - function:parse (level 0)
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

// empty_value_with_spaces_parse - function:parse (level 0)
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

// empty_key_indented_parsevalue - function:parsevalue (level 0)
func TestEmptyKeyIndentedParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_key_with_newline_parse - function:parse (level 0)
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

// empty_key_value_with_spaces_parse - function:parse (level 0)
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

// equals_in_value_no_spaces_parse - function:parse (level 0)
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

// equals_in_value_with_spaces_parse - function:parse (level 0)
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

// multiple_key_value_pairs_parse - function:parse (level 0)
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

// key_with_tabs_parse - function:parse behavior:tabs_preserve (level 0)
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

// key_with_tabs_ocaml_reference_parse - function:parse behavior:tabs_preserve (level 0)
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

// whitespace_only_value_parse - function:parse (level 0)
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

// spaces_vs_tabs_continuation_parsevalue - function:parsevalue behavior:tabs_preserve (level 0)
func TestSpacesVsTabsContinuationParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// spaces_vs_tabs_continuation_ocaml_reference_parsevalue - function:parsevalue behavior:tabs_preserve (level 0)
func TestSpacesVsTabsContinuationOcamlReferenceParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// multiple_empty_equality_parse - function:parse (level 0)
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

// key_with_newline_before_equals_parse - function:parse (level 0)
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

// complex_multi_newline_whitespace_parse - function:parse (level 0)
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

// empty_value_with_trailing_spaces_newline_parse - function:parse (level 0)
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

// empty_key_value_with_surrounding_newlines_parse - function:parse (level 0)
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

// quotes_treated_as_literal_unquoted_parse - function:parse (level 0)
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

// quotes_treated_as_literal_quoted_parse - function:parse (level 0)
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

// nested_single_line_parse - function:parse (level 0)
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

// nested_multi_line_parse - function:parse (level 0)
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

// nested_with_blank_line_parsevalue - function:parsevalue (level 0)
func TestNestedWithBlankLineParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deep_nested_structure_parsevalue - function:parsevalue (level 0)
func TestDeepNestedStructureParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// realistic_stress_test_parse - function:parse (level 0)
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

// ocaml_stress_test_original_parse - function:parse (level 0)
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

// ocaml_stress_test_original_buildhierarchy - function:buildhierarchy (level 0)
func TestOcamlStressTestOriginalBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// ocaml_stress_test_original_getstring - function:getstring (level 0)
func TestOcamlStressTestOriginalGetstring(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
