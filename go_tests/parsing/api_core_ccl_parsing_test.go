package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_core_ccl_parsing.json
// Suite: Flat Format
// Version: 1.0



// basic_key_value_pairs_parse - function:parse
func TestBasicKeyValuePairsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice
age = 42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "age", Value: "42"}}
	assert.Equal(t, expected, parseResult)

}


// equals_in_values_parse - function:parse
func TestEqualsInValuesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `msg = k=v pairs work fine
path = /bin/app=prod`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "msg", Value: "k=v pairs work fine"}, mock.Entry{Key: "path", Value: "/bin/app=prod"}}
	assert.Equal(t, expected, parseResult)

}


// whitespace_trimming_parse - function:parse feature:whitespace
func TestWhitespaceTrimmingParse(t *testing.T) {
	

	ccl := mock.New()
	input := `  key   =    value with spaces   
other = normal`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value with spaces"}, mock.Entry{Key: "other", Value: "normal"}}
	assert.Equal(t, expected, parseResult)

}


// multiline_values_parse - function:parse feature:multiline
func TestMultilineValuesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `description = First line
  Second line
  Third line
done = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "description", Value: "First line\n  Second line\n  Third line"}, mock.Entry{Key: "done", Value: "yes"}}
	assert.Equal(t, expected, parseResult)

}


// empty_values_parse - function:parse feature:empty_keys
func TestEmptyValuesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `empty =
other = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty", Value: ""}, mock.Entry{Key: "other", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// nested_structure_parsing_parse - function:parse
func TestNestedStructureParsingParse(t *testing.T) {
	

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  host = localhost\n  port = 5432"}}
	assert.Equal(t, expected, parseResult)

}


// unicode_parsing_parse - function:parse feature:unicode
func TestUnicodeParsingParse(t *testing.T) {
	

	ccl := mock.New()
	input := `emoji = 😀😃😄
配置 = config`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "emoji", Value: "😀😃😄"}, mock.Entry{Key: "配置", Value: "config"}}
	assert.Equal(t, expected, parseResult)

}


// empty_input_parse - function:parse
func TestEmptyInputParse(t *testing.T) {
	

	ccl := mock.New()
	input := ""
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


