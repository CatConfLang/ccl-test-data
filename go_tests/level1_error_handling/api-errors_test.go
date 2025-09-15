package level1_error_handling_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-errors.json
// Suite: CCL Error Cases (Validation Format)
// Version: 2.1
// Description: Error handling tests - malformed input detection across all levels using validation format

// just_key_error - function:parse (level 1)
func TestJustKeyError(t *testing.T) {

	ccl := mock.New()
	input := `key`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}

// whitespace_only_error - feature:whitespace function:parse variant:proposed_behavior (level 1)
func TestWhitespaceOnlyError(t *testing.T) {

	ccl := mock.New()
	input := `   `

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}

// whitespace_only_error_ocaml_reference - feature:whitespace function:parse variant:reference_compliant (level 1)
func TestWhitespaceOnlyErrorOcamlReference(t *testing.T) {

	ccl := mock.New()
	input := `   `

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}

// just_string_error - function:parse (level 1)
func TestJustStringError(t *testing.T) {

	ccl := mock.New()
	input := `val`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}

// multiline_plain_error - feature:multiline function:parse (level 1)
func TestMultilinePlainError(t *testing.T) {

	ccl := mock.New()
	input := `val
  next`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}

// multiline_plain_nested_error - feature:multiline function:build_hierarchy function:parse (level 1)
func TestMultilinePlainNestedError(t *testing.T) {

	ccl := mock.New()
	input := `
val
  next`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)

}
