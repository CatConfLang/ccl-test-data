package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_boolean_false-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// parse_boolean_false_parse - function:parse behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanFalseParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `disabled = false`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "disabled", Value: "false"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_false_build_hierarchy - function:build_hierarchy behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanFalseBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_boolean_false_get_bool - function:get_bool behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanFalseGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


