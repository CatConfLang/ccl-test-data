package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/nested_list_access-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// nested_list_access_parse - function:parse (level 0)
func TestNestedListAccessParse(t *testing.T) {

	ccl := mock.New()
	input := `database =
  hosts = primary
  hosts = secondary
  port = 5432`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  hosts = primary\n  hosts = secondary\n  port = 5432"}}
	assert.Equal(t, expected, parseResult)

}

// nested_list_access_build_hierarchy - function:build_hierarchy (level 0)
func TestNestedListAccessBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_get_list - function:get_list (level 0)
func TestNestedListAccessGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
