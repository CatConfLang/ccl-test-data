package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_string_fallback-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_string_fallback_parse - function:parse (level 0)
func TestParseStringFallbackParse(t *testing.T) {

	ccl := mock.New()
	input := `name = Alice`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}}
	assert.Equal(t, expected, parseResult)

}

// parse_string_fallback_build_hierarchy - function:build_hierarchy (level 0)
func TestParseStringFallbackBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_string_fallback_get_string - function:get_string (level 0)
func TestParseStringFallbackGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
