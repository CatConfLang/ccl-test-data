package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_with_whitespace-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// list_with_whitespace_parse - function:parse (level 0)
func TestListWithWhitespaceParse(t *testing.T) {

	ccl := mock.New()
	input := `items =   spaced   
items = normal
items =
items =   `

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "spaced"}, mock.Entry{Key: "items", Value: "normal"}, mock.Entry{Key: "items", Value: ""}, mock.Entry{Key: "items", Value: ""}}
	assert.Equal(t, expected, parseResult)

}

// list_with_whitespace_build_hierarchy - function:build_hierarchy (level 0)
func TestListWithWhitespaceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_get_list - function:get_list (level 0)
func TestListWithWhitespaceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
