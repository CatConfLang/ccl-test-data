package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_basic_float-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// parse_basic_float_parse - function:parse (level 0)
func TestParseBasicFloatParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = 98.6`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "temperature", Value: "98.6"}}
	assert.Equal(t, expected, parseResult)

}


// parse_basic_float_build_hierarchy - function:build_hierarchy (level 0)
func TestParseBasicFloatBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_basic_float_get_float - function:get_float (level 0)
func TestParseBasicFloatGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


