package level1_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-essential-parsing.json
// Suite: CCL Essential Parsing (Validation Format)
// Version: 2.0
// Description: Core parsing functionality - minimum viable CCL implementation. Start here for rapid prototyping.


// basic_pairs - function:parse (level 1)
func TestBasicPairs(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice
age = 42`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "age", Value: "42"}}
	assert.Equal(t, expectedParse, parseResult)

}


// equals_in_values - function:parse (level 1)
func TestEqualsInValues(t *testing.T) {
	
	
	ccl := mock.New()
	input := `msg = k=v pairs live happily here
more = a=b=c=d`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "msg", Value: "k=v pairs live happily here"}, mock.Entry{Key: "more", Value: "a=b=c=d"}}
	assert.Equal(t, expectedParse, parseResult)

}


// trimming_rules - feature:whitespace function:parse (level 1)
func TestTrimmingRules(t *testing.T) {
	
	
	ccl := mock.New()
	input := `  spaces around key   =    value with leading spaces removed and trailing tabs kept? 		`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "spaces around key", Value: "value with leading spaces removed and trailing tabs kept?"}}
	assert.Equal(t, expectedParse, parseResult)

}


// multiline_values - feature:multiline function:parse (level 1)
func TestMultilineValues(t *testing.T) {
	
	
	ccl := mock.New()
	input := `description = First
  Second line
  Third line
done = yes`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "description", Value: "First\n  Second line\n  Third line"}, mock.Entry{Key: "done", Value: "yes"}}
	assert.Equal(t, expectedParse, parseResult)

}


// blank_lines_in_values - feature:multiline function:parse (level 1)
func TestBlankLinesInValues(t *testing.T) {
	
	
	ccl := mock.New()
	input := `body = Line one

  Line three after a blank line`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "body", Value: "Line one\n\n  Line three after a blank line"}}
	assert.Equal(t, expectedParse, parseResult)

}


// no_equals_continuation - feature:multiline function:parse variant:proposed-behavior (level 1)
func TestNoEqualsContinuation(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key = first line
second line without equals
third line`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key", Value: "first line\nsecond line without equals\nthird line"}}
	assert.Equal(t, expectedParse, parseResult)

}


// no_equals_continuation_ocaml_reference - feature:multiline function:parse variant:reference-compliant (level 1)
func TestNoEqualsContinuationOcamlReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key = first line
second line without equals
third line`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}


// key_only_empty_value - feature:empty-keys function:parse (level 1)
func TestKeyOnlyEmptyValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `emptykey =
other = value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "emptykey", Value: ""}, mock.Entry{Key: "other", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)

}


// empty_key_with_value - feature:empty-keys function:parse (level 1)
func TestEmptyKeyWithValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `= val`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "val"}}
	assert.Equal(t, expectedParse, parseResult)

}


// empty_key_and_value - feature:empty-keys function:parse (level 1)
func TestEmptyKeyAndValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `=`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: ""}}
	assert.Equal(t, expectedParse, parseResult)

}


// unicode_graphemes - feature:unicode function:parse (level 1)
func TestUnicodeGraphemes(t *testing.T) {
	
	
	ccl := mock.New()
	input := `emoji = 😀😃😄`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "emoji", Value: "😀😃😄"}}
	assert.Equal(t, expectedParse, parseResult)

}


// unicode_keys - feature:unicode function:parse (level 1)
func TestUnicodeKeys(t *testing.T) {
	
	
	ccl := mock.New()
	input := `配置 = config in Chinese
ключ = key in Russian`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "配置", Value: "config in Chinese"}, mock.Entry{Key: "ключ", Value: "key in Russian"}}
	assert.Equal(t, expectedParse, parseResult)

}


// crlf_normalization - behavior:crlf-preserve function:parse (level 1)
func TestCrlfNormalization(t *testing.T) {
	
	
	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}}
	assert.Equal(t, expectedParse, parseResult)

}


// crlf_normalization_strict - behavior:crlf-preserve function:parse (level 1)
func TestCrlfNormalizationStrict(t *testing.T) {
	
	
	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key1", Value: "value1\r"}, mock.Entry{Key: "key2", Value: "value2\r"}}
	assert.Equal(t, expectedParse, parseResult)

}


// empty_input - function:parse (level 1)
func TestEmptyInput(t *testing.T) {
	
	
	ccl := mock.New()
	
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}


// eof_without_newline - function:parse (level 1)
func TestEofWithoutNewline(t *testing.T) {
	
	
	ccl := mock.New()
	input := `last = final value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "last", Value: "final value"}}
	assert.Equal(t, expectedParse, parseResult)

}


// tab_preservation_in_values - behavior:tabs-preserve feature:whitespace function:parse variant:proposed-behavior (level 1)
func TestTabPreservationInValues(t *testing.T) {
	
	
	ccl := mock.New()
	input := `mixed =    	value with leading spaces and tabs	`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "mixed", Value: "\tvalue with leading spaces and tabs"}}
	assert.Equal(t, expectedParse, parseResult)

}


// tab_preservation_in_values_ocaml_reference - behavior:tabs-preserve feature:whitespace function:parse variant:reference-compliant (level 1)
func TestTabPreservationInValuesOcamlReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `mixed =    	value with leading spaces and tabs	`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "mixed", Value: "value with leading spaces and tabs"}}
	assert.Equal(t, expectedParse, parseResult)

}


// mixed_indentation_continuation - feature:multiline function:parse variant:proposed-behavior (level 1)
func TestMixedIndentationContinuation(t *testing.T) {
	
	
	ccl := mock.New()
	input := `description = First line
    spaces indented
	tab indented
  	mixed indent`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "description", Value: "First line\n    spaces indented\n\ttab indented\n  \tmixed indent"}}
	assert.Equal(t, expectedParse, parseResult)

}


// mixed_indentation_continuation_ocaml_reference - feature:multiline function:parse variant:reference-compliant (level 1)
func TestMixedIndentationContinuationOcamlReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `description = First line
    spaces indented
	tab indented
  	mixed indent`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "description", Value: "First line\n    spaces indented\n tab indented\n   mixed indent"}}
	assert.Equal(t, expectedParse, parseResult)

}


// indented_equals_continuation - function:parse (level 1)
func TestIndentedEqualsContinuation(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key1 = val1
  inner = some
key2 = val2`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key1", Value: "val1\n  inner = some"}, mock.Entry{Key: "key2", Value: "val2"}}
	assert.Equal(t, expectedParse, parseResult)

}


// nested_key_value_pairs - function:make-objects function:parse (level 1)
func TestNestedKeyValuePairs(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =
  field1 = value1
  field2 = value2`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key", Value: "\n  field1 = value1\n  field2 = value2"}}
	assert.Equal(t, expectedParse, parseResult)

}


