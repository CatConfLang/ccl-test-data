package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_path_traversal_protection_reference-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// list_path_traversal_protection_reference_get_list - function:get_list behavior:list_coercion_disabled (level 0)
func TestListPathTraversalProtectionReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_path_traversal_protection_reference_parse - function:parse behavior:list_coercion_disabled (level 0)
func TestListPathTraversalProtectionReferenceParse(t *testing.T) {
	
	
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


// list_path_traversal_protection_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled (level 0)
func TestListPathTraversalProtectionReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


