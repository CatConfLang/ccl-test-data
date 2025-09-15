package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/complete_lists_workflow-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// complete_lists_workflow_parse - function:parse (level 0)
func TestCompleteListsWorkflowParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers =
  server = web1
  server = web2
  server = web3
ports =
  port = 80
  port = 443`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "\n  server = web1\n  server = web2\n  server = web3"}, mock.Entry{Key: "ports", Value: "\n  port = 80\n  port = 443"}}
	assert.Equal(t, expected, parseResult)

}


// complete_lists_workflow_build_hierarchy - function:build_hierarchy (level 0)
func TestCompleteListsWorkflowBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


