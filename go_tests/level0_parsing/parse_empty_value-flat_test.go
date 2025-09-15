package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_empty_value-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_empty_value_build_hierarchy - function:build_hierarchy (level 0)
func TestParseEmptyValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_empty_value_get_string - function:get_string (level 0)
func TestParseEmptyValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_empty_value_parse - function:parse (level 0)
func TestParseEmptyValueParse(t *testing.T) {

	ccl := mock.New()
	input := `empty_key =`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}
