package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/deep_nested_objects-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// deep_nested_objects_parse - function:parse (level 0)
func TestDeepNestedObjectsParse(t *testing.T) {

	ccl := mock.New()
	input := `server =
  database =
    host = localhost
    port = 5432
  cache =
    enabled = true`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "server", Value: "\n  database =\n    host = localhost\n    port = 5432\n  cache =\n    enabled = true"}}
	assert.Equal(t, expected, parseResult)

}

// deep_nested_objects_build_hierarchy - function:build_hierarchy (level 0)
func TestDeepNestedObjectsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deep_nested_objects_load - function:load (level 0)
func TestDeepNestedObjectsLoad(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
