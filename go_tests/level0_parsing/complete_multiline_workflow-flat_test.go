package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/complete_multiline_workflow-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// complete_multiline_workflow_parse - function:parse feature:multiline (level 0)
func TestCompleteMultilineWorkflowParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `description = Welcome to our app
  This is a multi-line description
  With several lines
config =
  settings =
    value1 = one
    value2 = two`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "description", Value: "Welcome to our app\n  This is a multi-line description\n  With several lines"}, mock.Entry{Key: "config", Value: "\n  settings =\n    value1 = one\n    value2 = two"}}
	assert.Equal(t, expected, parseResult)

}


// complete_multiline_workflow_build_hierarchy - function:build_hierarchy feature:multiline (level 0)
func TestCompleteMultilineWorkflowBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


