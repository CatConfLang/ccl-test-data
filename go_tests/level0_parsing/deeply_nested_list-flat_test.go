package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/deeply_nested_list-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// deeply_nested_list_parse - function:parse (level 0)
func TestDeeplyNestedListParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers = web1
      servers = web2
      servers = api1`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: ""}, mock.Entry{Key: "environments", Value: ""}, mock.Entry{Key: "production", Value: ""}, mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "api1"}}
	assert.Equal(t, expected, parseResult)

}


// deeply_nested_list_build_hierarchy - function:build_hierarchy (level 0)
func TestDeeplyNestedListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// deeply_nested_list_get_list - function:get_list (level 0)
func TestDeeplyNestedListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


