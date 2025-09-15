package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_with_numbers_reference-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// list_with_numbers_reference_parse - function:parse (level 0)
func TestListWithNumbersReferenceParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `numbers = 1
numbers = 42
numbers = -17
numbers = 0`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "numbers", Value: "1"}, mock.Entry{Key: "numbers", Value: "42"}, mock.Entry{Key: "numbers", Value: "-17"}, mock.Entry{Key: "numbers", Value: "0"}}
	assert.Equal(t, expected, parseResult)

}


// list_with_numbers_reference_build_hierarchy - function:build_hierarchy (level 0)
func TestListWithNumbersReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_numbers_reference_get_list - function:get_list (level 0)
func TestListWithNumbersReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


