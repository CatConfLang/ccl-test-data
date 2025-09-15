package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_error_nested_missing_key-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// list_error_nested_missing_key_parse - function:parse (level 0)
func TestListErrorNestedMissingKeyParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  server = web1`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  server = web1"}}
	assert.Equal(t, expected, parseResult)

}


// list_error_nested_missing_key_build_hierarchy - function:build_hierarchy (level 0)
func TestListErrorNestedMissingKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_error_nested_missing_key_get_list - function:get_list (level 0)
func TestListErrorNestedMissingKeyGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


