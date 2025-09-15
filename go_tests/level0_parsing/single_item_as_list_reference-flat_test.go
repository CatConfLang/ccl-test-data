package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/single_item_as_list_reference-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// single_item_as_list_reference_parse - function:parse (level 0)
func TestSingleItemAsListReferenceParse(t *testing.T) {

	ccl := mock.New()
	input := `item = single`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item", Value: "single"}}
	assert.Equal(t, expected, parseResult)

}

// single_item_as_list_reference_build_hierarchy - function:build_hierarchy (level 0)
func TestSingleItemAsListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_reference_get_list - function:get_list (level 0)
func TestSingleItemAsListReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
