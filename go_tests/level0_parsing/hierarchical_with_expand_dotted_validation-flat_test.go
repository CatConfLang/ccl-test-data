package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/hierarchical_with_expand_dotted_validation-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// hierarchical_with_expand_dotted_validation_parse - function:parse feature:experimental_dotted_keys (level 0)
func TestHierarchicalWithExpandDottedValidationParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  enabled = true
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  enabled = true\n  port = 5432"}}
	assert.Equal(t, expected, parseResult)

}


// hierarchical_with_expand_dotted_validation_expand_dotted - function:expand_dotted feature:experimental_dotted_keys (level 0)
func TestHierarchicalWithExpandDottedValidationExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// hierarchical_with_expand_dotted_validation_build_hierarchy - function:build_hierarchy feature:experimental_dotted_keys (level 0)
func TestHierarchicalWithExpandDottedValidationBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


