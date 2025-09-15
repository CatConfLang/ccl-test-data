package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_boolean_yes-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// parse_boolean_yes_get_bool - function:get_bool behavior:boolean_lenient (level 0)
func TestParseBooleanYesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_boolean_yes_parse - function:parse behavior:boolean_lenient (level 0)
func TestParseBooleanYesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "active", Value: "yes"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_yes_build_hierarchy - function:build_hierarchy behavior:boolean_lenient (level 0)
func TestParseBooleanYesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


