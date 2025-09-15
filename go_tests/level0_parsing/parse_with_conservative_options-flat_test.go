package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_with_conservative_options-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_with_conservative_options_parse - function:parse (level 0)
func TestParseWithConservativeOptionsParse(t *testing.T) {

	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "number", Value: "42"}, mock.Entry{Key: "decimal", Value: "3.14"}, mock.Entry{Key: "flag", Value: "true"}, mock.Entry{Key: "text", Value: "hello"}}
	assert.Equal(t, expected, parseResult)

}

// parse_with_conservative_options_build_hierarchy - function:build_hierarchy (level 0)
func TestParseWithConservativeOptionsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_conservative_options_get_int - function:get_int (level 0)
func TestParseWithConservativeOptionsGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_conservative_options_get_string - function:get_string (level 0)
func TestParseWithConservativeOptionsGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
