package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_zero_values_strict_literal-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_zero_values_strict_literal_get_int - function:get_int (level 0)
func TestParseZeroValuesStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_float - function:get_float (level 0)
func TestParseZeroValuesStrictLiteralGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_bool - function:get_bool (level 0)
func TestParseZeroValuesStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_parse - function:parse (level 0)
func TestParseZeroValuesStrictLiteralParse(t *testing.T) {

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "count", Value: "0"}, mock.Entry{Key: "distance", Value: "0.0"}, mock.Entry{Key: "disabled", Value: "no"}}
	assert.Equal(t, expected, parseResult)

}

// parse_zero_values_strict_literal_build_hierarchy - function:build_hierarchy (level 0)
func TestParseZeroValuesStrictLiteralBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
