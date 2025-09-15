package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_integer_error-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// parse_integer_error_parse - function:parse (level 0)
func TestParseIntegerErrorParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = not_a_number`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "port", Value: "not_a_number"}}
	assert.Equal(t, expected, parseResult)

}


// parse_integer_error_build_hierarchy - function:build_hierarchy (level 0)
func TestParseIntegerErrorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_integer_error_get_int - function:get_int (level 0)
func TestParseIntegerErrorGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


