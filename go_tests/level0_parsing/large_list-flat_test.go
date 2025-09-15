package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/large_list-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// large_list_build_hierarchy - function:build_hierarchy (level 0)
func TestLargeListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// large_list_get_list - function:get_list (level 0)
func TestLargeListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// large_list_parse - function:parse (level 0)
func TestLargeListParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `items = item01
items = item02
items = item03
items = item04
items = item05
items = item06
items = item07
items = item08
items = item09
items = item10
items = item11
items = item12
items = item13
items = item14
items = item15
items = item16
items = item17
items = item18
items = item19
items = item20`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "item01"}, mock.Entry{Key: "items", Value: "item02"}, mock.Entry{Key: "items", Value: "item03"}, mock.Entry{Key: "items", Value: "item04"}, mock.Entry{Key: "items", Value: "item05"}, mock.Entry{Key: "items", Value: "item06"}, mock.Entry{Key: "items", Value: "item07"}, mock.Entry{Key: "items", Value: "item08"}, mock.Entry{Key: "items", Value: "item09"}, mock.Entry{Key: "items", Value: "item10"}, mock.Entry{Key: "items", Value: "item11"}, mock.Entry{Key: "items", Value: "item12"}, mock.Entry{Key: "items", Value: "item13"}, mock.Entry{Key: "items", Value: "item14"}, mock.Entry{Key: "items", Value: "item15"}, mock.Entry{Key: "items", Value: "item16"}, mock.Entry{Key: "items", Value: "item17"}, mock.Entry{Key: "items", Value: "item18"}, mock.Entry{Key: "items", Value: "item19"}, mock.Entry{Key: "items", Value: "item20"}}
	assert.Equal(t, expected, parseResult)

}


