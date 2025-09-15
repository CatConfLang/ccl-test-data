package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/nested_duplicate_keys-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// nested_duplicate_keys_parse - function:parse (level 0)
func TestNestedDuplicateKeysParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  server = web1
  server = web2
  port = 80`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  server = web1\n  server = web2\n  port = 80"}}
	assert.Equal(t, expected, parseResult)

}


// nested_duplicate_keys_build_hierarchy - function:build_hierarchy (level 0)
func TestNestedDuplicateKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


