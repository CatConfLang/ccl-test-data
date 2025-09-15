package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_with_whitespace-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_with_whitespace_parse - function:parse (level 0)
func TestParseWithWhitespaceParse(t *testing.T) {

	ccl := mock.New()
	input := `number =   42   
flag =  true  `

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "number", Value: "42"}, mock.Entry{Key: "flag", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}

// parse_with_whitespace_build_hierarchy - function:build_hierarchy (level 0)
func TestParseWithWhitespaceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_whitespace_get_int - function:get_int (level 0)
func TestParseWithWhitespaceGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_whitespace_get_bool - function:get_bool (level 0)
func TestParseWithWhitespaceGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
