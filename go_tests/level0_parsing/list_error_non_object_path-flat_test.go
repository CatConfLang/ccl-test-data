package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_error_non_object_path-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// list_error_non_object_path_parse - function:parse (level 0)
func TestListErrorNonObjectPathParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `value = simple`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "value", Value: "simple"}}
	assert.Equal(t, expected, parseResult)

}


// list_error_non_object_path_build_hierarchy - function:build_hierarchy (level 0)
func TestListErrorNonObjectPathBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_error_non_object_path_get_list - function:get_list (level 0)
func TestListErrorNonObjectPathGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


