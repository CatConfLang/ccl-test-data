package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/single_dot_key-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// single_dot_key_expand_dotted - function:expand_dotted feature:experimental_dotted_keys feature:empty_keys (level 0)
func TestSingleDotKeyExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// single_dot_key_build_hierarchy - function:build_hierarchy feature:experimental_dotted_keys feature:empty_keys (level 0)
func TestSingleDotKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// single_dot_key_parse - function:parse feature:experimental_dotted_keys feature:empty_keys (level 0)
func TestSingleDotKeyParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a. = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a.", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


