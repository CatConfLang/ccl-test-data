package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_mixed_types-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_mixed_types_parse - function:parse (level 0)
func TestParseMixedTypesParse(t *testing.T) {

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "ssl", Value: "true"}, mock.Entry{Key: "timeout", Value: "30.5"}, mock.Entry{Key: "debug", Value: "off"}}
	assert.Equal(t, expected, parseResult)

}

// parse_mixed_types_build_hierarchy - function:build_hierarchy (level 0)
func TestParseMixedTypesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_string - function:get_string (level 0)
func TestParseMixedTypesGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_int - function:get_int (level 0)
func TestParseMixedTypesGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_bool - function:get_bool (level 0)
func TestParseMixedTypesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_float - function:get_float (level 0)
func TestParseMixedTypesGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
