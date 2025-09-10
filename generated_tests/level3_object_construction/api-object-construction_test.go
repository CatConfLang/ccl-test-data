package level3_object_construction_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-object-construction.json
// Suite: CCL Object Construction - Validation Format
// Version: 2.0
// Description: Converting flat Entry[] to nested objects via make_objects() - essential for hierarchical access


// recursive_nested_single - recursive nested basic (level 3)
func TestRecursiveNestedSingle(t *testing.T) {
	t.Skip("Complex parsing not implemented in mock CCL")
}


// recursive_nested_multiple - recursive nested multiple (level 3)
func TestRecursiveNestedMultiple(t *testing.T) {
	t.Skip("Complex parsing not implemented in mock CCL")
}


// deep_recursive_nesting - recursive nested deep (level 3)
func TestDeepRecursiveNesting(t *testing.T) {
	t.Skip("Complex parsing not implemented in mock CCL")
}


// multiple_keys_same_name_merge - merge duplicate-keys object-construction (level 3)
func TestMultipleKeysSameNameMerge(t *testing.T) {
	
	
	ccl := mock.New()
	input := `user =
  name = alice
user =
  age = 25`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "user", Value: "\n  name = alice"}, mock.Entry{Key: "user", Value: "\n  age = 25"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"user": map[string]interface{}{"age": "25", "name": "alice"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// empty_key_list_style - empty-key lists array-style (level 3)
func TestEmptyKeyListStyle(t *testing.T) {
	
	
	ccl := mock.New()
	input := `ports =
  = 8000
  = 8001
  = 8002`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "ports", Value: "\n  = 8000\n  = 8001\n  = 8002"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"ports": map[string]interface{}{"": []interface{}{"8000", "8001", "8002"}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// empty_key_top_level - empty-key top-level merge proposed-behavior (level 3)
func TestEmptyKeyTopLevel(t *testing.T) {
	
	
	ccl := mock.New()
	input := `= root_value
key = normal
= another_root`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "root_value"}, mock.Entry{Key: "key", Value: "normal"}, mock.Entry{Key: "", Value: "another_root"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"": []interface{}{"root_value", "another_root"}, "key": "normal"}
	assert.Equal(t, expectedObjects, objectResult)

}


// empty_key_top_level_ocaml_reference - empty-key top-level merge reference-compliant-behavior (level 3)
func TestEmptyKeyTopLevelOcamlReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `= root_value
key = normal
= another_root`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "root_value"}, mock.Entry{Key: "key", Value: "normal"}, mock.Entry{Key: "", Value: "another_root"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"": []interface{}{"another_root", "root_value"}, "key": "normal"}
	assert.Equal(t, expectedObjects, objectResult)

}


// mixed_flat_and_nested - mixed flat nested (level 3)
func TestMixedFlatAndNested(t *testing.T) {
	t.Skip("Complex parsing not implemented in mock CCL")
}


// stress_test_complex_nesting - stress complex merge empty-key comments (level 3)
func TestStressTestComplexNesting(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/ = This is a CCL document
title = CCL Example

database =
  enabled = true
  ports =
    = 8000
    = 8001
    = 8002
  limits =
    cpu = 1500mi
    memory = 10Gb

user =
  guestId = 42

user =
  login = chshersh
  createdAt = 2024-12-31`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "/", Value: "This is a CCL document"}, mock.Entry{Key: "title", Value: "CCL Example"}, mock.Entry{Key: "database", Value: "\n  enabled = true\n  ports =\n    = 8000\n    = 8001\n    = 8002\n  limits =\n    cpu = 1500mi\n    memory = 10Gb"}, mock.Entry{Key: "user", Value: "\n  guestId = 42"}, mock.Entry{Key: "user", Value: "\n  login = chshersh\n  createdAt = 2024-12-31"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"/": "This is a CCL document", "database": map[string]interface{}{"enabled": "true", "limits": map[string]interface{}{"cpu": "1500mi", "memory": "10Gb"}, "ports": map[string]interface{}{"": []interface{}{"8000", "8001", "8002"}}}, "title": "CCL Example", "user": map[string]interface{}{"createdAt": "2024-12-31", "guestId": "42", "login": "chshersh"}}
	assert.Equal(t, expectedObjects, objectResult)

}


