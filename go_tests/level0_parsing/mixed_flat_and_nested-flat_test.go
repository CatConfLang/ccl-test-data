package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/mixed_flat_and_nested-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// mixed_flat_and_nested_parse - function:parse (level 0)
func TestMixedFlatAndNestedParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice
config =
  debug = true
  timeout = 30
version = 1.0`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "config", Value: "\n  debug = true\n  timeout = 30"}, mock.Entry{Key: "version", Value: "1.0"}}
	assert.Equal(t, expected, parseResult)

}


// mixed_flat_and_nested_build_hierarchy - function:build_hierarchy (level 0)
func TestMixedFlatAndNestedBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


