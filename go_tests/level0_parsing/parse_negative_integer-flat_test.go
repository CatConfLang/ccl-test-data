package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_negative_integer-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_negative_integer_parse - function:parse (level 0)
func TestParseNegativeIntegerParse(t *testing.T) {

	ccl := mock.New()
	input := `offset = -42`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "offset", Value: "-42"}}
	assert.Equal(t, expected, parseResult)

}

// parse_negative_integer_build_hierarchy - function:build_hierarchy (level 0)
func TestParseNegativeIntegerBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_negative_integer_get_int - function:get_int (level 0)
func TestParseNegativeIntegerGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
