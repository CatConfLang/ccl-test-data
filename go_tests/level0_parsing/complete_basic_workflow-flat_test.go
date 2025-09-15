package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/complete_basic_workflow-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// complete_basic_workflow_parse - function:parse (level 0)
func TestCompleteBasicWorkflowParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice
age = 42`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "age", Value: "42"}}
	assert.Equal(t, expected, parseResult)

}


// complete_basic_workflow_build_hierarchy - function:build_hierarchy (level 0)
func TestCompleteBasicWorkflowBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


