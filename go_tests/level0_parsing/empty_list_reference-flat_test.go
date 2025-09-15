package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/empty_list_reference-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// empty_list_reference_build_hierarchy - function:build_hierarchy (level 0)
func TestEmptyListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// empty_list_reference_get_list - function:get_list (level 0)
func TestEmptyListReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// empty_list_reference_parse - function:parse (level 0)
func TestEmptyListReferenceParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `empty_list =`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_list", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


