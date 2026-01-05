package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_core_ccl_hierarchy.json
// Suite: Flat Format
// Version: 1.0



// basic_object_construction_parse - function:parse
func TestBasicObjectConstructionParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice
age = 42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "age", Value: "42"}}
	assert.Equal(t, expected, parseResult)

}


// basic_object_construction_build_hierarchy - function:build_hierarchy
func TestBasicObjectConstructionBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice
age = 42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"age": "42", "name": "Alice"}
	assert.Equal(t, expected, objectResult)

}


// deep_nested_objects_parse - function:parse
func TestDeepNestedObjectsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `server =
  database =
    host = localhost
    port = 5432
  cache =
    enabled = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "server", Value: "\n  database =\n    host = localhost\n    port = 5432\n  cache =\n    enabled = true"}}
	assert.Equal(t, expected, parseResult)

}


// deep_nested_objects_build_hierarchy - function:build_hierarchy
func TestDeepNestedObjectsBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `server =
  database =
    host = localhost
    port = 5432
  cache =
    enabled = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"server": map[string]interface{}{"cache": map[string]interface{}{"enabled": "true"}, "database": map[string]interface{}{"host": "localhost", "port": "5432"}}}
	assert.Equal(t, expected, objectResult)

}


// duplicate_keys_to_lists_parse - function:parse
func TestDuplicateKeysToListsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `item = first
item = second
item = third`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item", Value: "first"}, mock.Entry{Key: "item", Value: "second"}, mock.Entry{Key: "item", Value: "third"}}
	assert.Equal(t, expected, parseResult)

}


// duplicate_keys_to_lists_build_hierarchy - function:build_hierarchy
func TestDuplicateKeysToListsBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `item = first
item = second
item = third`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"item": []interface{}{"first", "second", "third"}}
	assert.Equal(t, expected, objectResult)

}


// nested_duplicate_keys_parse - function:parse
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


// nested_duplicate_keys_build_hierarchy - function:build_hierarchy
func TestNestedDuplicateKeysBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  server = web1
  server = web2
  port = 80`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"port": "80", "server": []interface{}{"web1", "web2"}}}
	assert.Equal(t, expected, objectResult)

}


// mixed_flat_and_nested_parse - function:parse
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


// mixed_flat_and_nested_build_hierarchy - function:build_hierarchy
func TestMixedFlatAndNestedBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice
config =
  debug = true
  timeout = 30
version = 1.0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"debug": "true", "timeout": "30"}, "name": "Alice", "version": "1.0"}
	assert.Equal(t, expected, objectResult)

}


// nested_objects_with_lists_parse - function:parse
func TestNestedObjectsWithListsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `environments =
  prod =
    server = web1
    server = web2
    port = 80
  dev =
    server = localhost
    port = 3000`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "environments", Value: "\n  prod =\n    server = web1\n    server = web2\n    port = 80\n  dev =\n    server = localhost\n    port = 3000"}}
	assert.Equal(t, expected, parseResult)

}


// nested_objects_with_lists_build_hierarchy - function:build_hierarchy
func TestNestedObjectsWithListsBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `environments =
  prod =
    server = web1
    server = web2
    port = 80
  dev =
    server = localhost
    port = 3000`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"environments": map[string]interface{}{"dev": map[string]interface{}{"port": "3000", "server": "localhost"}, "prod": map[string]interface{}{"port": "80", "server": []interface{}{"web1", "web2"}}}}
	assert.Equal(t, expected, objectResult)

}


// deeply_nested_list_parse - function:parse behavior:list_coercion_disabled behavior:array_order_lexicographic
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
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  environments =\n    production =\n      servers = web1\n      servers = web2\n      servers = api1"}}
	assert.Equal(t, expected, parseResult)

}


// deeply_nested_list_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled behavior:array_order_lexicographic
func TestDeeplyNestedListBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers = web1
      servers = web2
      servers = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"environments": map[string]interface{}{"production": map[string]interface{}{"servers": []interface{}{"api1", "web1", "web2"}}}}}
	assert.Equal(t, expected, objectResult)

}


// deeply_nested_list_get_list - function:get_list behavior:list_coercion_disabled behavior:array_order_lexicographic
func TestDeeplyNestedListGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers = web1
      servers = web2
      servers = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"config", "environments", "production", "servers"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Empty(t, result)
	}

}


