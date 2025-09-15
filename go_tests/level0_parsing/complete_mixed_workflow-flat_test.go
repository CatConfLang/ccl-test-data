package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/complete_mixed_workflow-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// complete_mixed_workflow_build_hierarchy - function:build_hierarchy (level 0)
func TestCompleteMixedWorkflowBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// complete_mixed_workflow_parse - function:parse (level 0)
func TestCompleteMixedWorkflowParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app = MyApp
version = 1.0.0
config =
  debug = true
  features =
    feature1 = enabled
    feature2 = disabled`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "app", Value: "MyApp"}, mock.Entry{Key: "version", Value: "1.0.0"}, mock.Entry{Key: "config", Value: "\n  debug = true\n  features =\n    feature1 = enabled\n    feature2 = disabled"}}
	assert.Equal(t, expected, parseResult)

}


