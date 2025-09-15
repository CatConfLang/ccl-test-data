package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/basic_key_value_pairs-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// basic_key_value_pairs_parse - function:parse (level 0)
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
