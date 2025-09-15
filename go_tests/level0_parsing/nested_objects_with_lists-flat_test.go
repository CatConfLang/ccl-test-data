package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/nested_objects_with_lists-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// nested_objects_with_lists_parse - function:parse (level 0)
func TestNestedObjectsWithListsParse(t *testing.T) {

	ccl := mock.New()
	input := `environments =
  prod =
    server = web1
    server = web2
    port = 80
  dev =
    server = localhost
    port = 3000`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "environments", Value: "\n  prod =\n    server = web1\n    server = web2\n    port = 80\n  dev =\n    server = localhost\n    port = 3000"}}
	assert.Equal(t, expected, parseResult)

}

// nested_objects_with_lists_build_hierarchy - function:build_hierarchy (level 0)
func TestNestedObjectsWithListsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
