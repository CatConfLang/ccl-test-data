package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_basic_integer-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// parse_basic_integer_parse - function:parse (level 0)
func TestParseBasicIntegerParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = 8080`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "port", Value: "8080"}}
	assert.Equal(t, expected, parseResult)

}


// parse_basic_integer_build_hierarchy - function:build_hierarchy (level 0)
func TestParseBasicIntegerBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_basic_integer_get_int - function:get_int (level 0)
func TestParseBasicIntegerGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


