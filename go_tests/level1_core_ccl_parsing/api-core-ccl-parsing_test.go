package level1_core_ccl_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api_core-ccl-parsing.json
// Suite: CCL Core Parsing - Validation Format
// Version: 2.1
// Description: Core CCL text parsing fundamentals - validates correct text → entry conversion as foundation for hierarchy building.

// basic_key_value_pairs - function:parse (level 1)
func TestBasicKeyValuePairs(t *testing.T) {

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
	input := `msg = k=v pairs work fine
path = /bin/app=prod`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "msg", Value: "k=v pairs work fine"}, mock.Entry{Key: "path", Value: "/bin/app=prod"}}
	assert.Equal(t, expectedParse, parseResult)

}

// whitespace_trimming - function:parse feature:whitespace (level 1)
func TestWhitespaceTrimming(t *testing.T) {

	ccl := mock.New()
	input := `  key   =    value with spaces   
other = normal`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key", Value: "value with spaces"}, mock.Entry{Key: "other", Value: "normal"}}
	assert.Equal(t, expectedParse, parseResult)

}

// multiline_values - function:parse feature:multiline (level 1)
func TestMultilineValues(t *testing.T) {

	ccl := mock.New()
	input := `description = First line
  Second line
  Third line
done = yes`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "description", Value: "First line\n  Second line\n  Third line"}, mock.Entry{Key: "done", Value: "yes"}}
	assert.Equal(t, expectedParse, parseResult)

}

// empty_values - function:parse feature:empty_keys (level 1)
func TestEmptyValues(t *testing.T) {

	ccl := mock.New()
	input := `empty =
other = value`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "empty", Value: ""}, mock.Entry{Key: "other", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)

}

// nested_structure_parsing - function:parse (level 1)
func TestNestedStructureParsing(t *testing.T) {

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "\n  host = localhost\n  port = 5432"}}
	assert.Equal(t, expectedParse, parseResult)

}

// unicode_parsing - function:parse feature:unicode (level 1)
func TestUnicodeParsing(t *testing.T) {

	ccl := mock.New()
	input := `emoji = 😀😃😄
配置 = config`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "emoji", Value: "😀😃😄"}, mock.Entry{Key: "配置", Value: "config"}}
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
