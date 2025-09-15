package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_float_error-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// parse_float_error_get_float - function:get_float (level 0)
func TestParseFloatErrorGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_float_error_parse - function:parse (level 0)
func TestParseFloatErrorParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = invalid`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "temperature", Value: "invalid"}}
	assert.Equal(t, expected, parseResult)

}


// parse_float_error_build_hierarchy - function:build_hierarchy (level 0)
func TestParseFloatErrorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


