package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_errors.json
// Suite: Flat Format
// Version: 1.0



// just_key_error_parse - function:parse
func TestJustKeyErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


// whitespace_only_error_parse - function:parse feature:whitespace
func TestWhitespaceOnlyErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `   `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


// whitespace_only_error_ocaml_reference_parse - function:parse feature:whitespace
func TestWhitespaceOnlyErrorOcamlReferenceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `   `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


// just_string_error_parse - function:parse
func TestJustStringErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


// multiline_plain_error_parse - function:parse feature:multiline
func TestMultilinePlainErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `val
  next`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


// multiline_plain_nested_error_parse - function:parse feature:multiline
func TestMultilinePlainNestedErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `
val
  next`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


