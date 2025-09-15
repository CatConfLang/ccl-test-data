package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/complete_nested_workflow-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// complete_nested_workflow_parse - function:parse (level 0)
func TestCompleteNestedWorkflowParse(t *testing.T) {

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432
  enabled = true`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  host = localhost\n  port = 5432\n  enabled = true"}}
	assert.Equal(t, expected, parseResult)

}

// complete_nested_workflow_build_hierarchy - function:build_hierarchy (level 0)
func TestCompleteNestedWorkflowBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
