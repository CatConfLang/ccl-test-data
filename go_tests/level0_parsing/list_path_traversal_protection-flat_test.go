package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_path_traversal_protection-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// list_path_traversal_protection_get_list - function:get_list (level 0)
func TestListPathTraversalProtectionGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_parse - function:parse (level 0)
func TestListPathTraversalProtectionParse(t *testing.T) {

	ccl := mock.New()
	input := `safe = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "safe", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// list_path_traversal_protection_build_hierarchy - function:build_hierarchy (level 0)
func TestListPathTraversalProtectionBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
