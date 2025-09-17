package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_errors.json
// Suite:
// Version:

// just_key_error_parse - function:parse (level 0)
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

// whitespace_only_error_parse - function:parse (level 0)
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

// whitespace_only_error_ocaml_reference_parse - function:parse (level 0)
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

// just_string_error_parse - function:parse (level 0)
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

// multiline_plain_error_parse - function:parse (level 0)
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

// multiline_plain_nested_error_parse - function:parse (level 0)
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
