package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/single_item_as_list-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// single_item_as_list_build_hierarchy - function:build_hierarchy behavior:list_coercion_enabled (level 0)
func TestSingleItemAsListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// single_item_as_list_get_list - function:get_list behavior:list_coercion_enabled (level 0)
func TestSingleItemAsListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// single_item_as_list_parse - function:parse behavior:list_coercion_enabled (level 0)
func TestSingleItemAsListParse(t *testing.T) {
	
	
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


