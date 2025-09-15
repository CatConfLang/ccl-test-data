package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_edge_case_zero_length-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// list_edge_case_zero_length_get_list - function:get_list (level 0)
func TestListEdgeCaseZeroLengthGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_edge_case_zero_length_parse - function:parse (level 0)
func TestListEdgeCaseZeroLengthParse(t *testing.T) {

	ccl := mock.New()
	input := ""

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}

// list_edge_case_zero_length_build_hierarchy - function:build_hierarchy (level 0)
func TestListEdgeCaseZeroLengthBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
