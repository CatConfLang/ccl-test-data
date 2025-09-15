package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_boolean_true-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// parse_boolean_true_parse - function:parse behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanTrueParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = true`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "enabled", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_true_build_hierarchy - function:build_hierarchy behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanTrueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_boolean_true_get_bool - function:get_bool behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanTrueGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


