package level1_core_ccl_hierarchy_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-core-ccl-hierarchy.json
// Suite: CCL Core Hierarchy - Validation Format
// Version: 2.1
// Description: Core CCL hierarchy construction - validates flat entries → nested objects transformation as part of atomic Core CCL.


// basic_object_construction - function:parse function:expand-dotted function:build-hierarchy (level 1)
func TestBasicObjectConstruction(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice
age = 42`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "age", Value: "42"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"age": "42", "name": "Alice"}
	assert.Equal(t, expectedObjects, objectResult)

}


// single_nested_object - function:parse function:expand-dotted function:build-hierarchy (level 1)
func TestSingleNestedObject(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  enabled = true
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "\n  enabled = true\n  port = 5432"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"enabled": "true", "port": "5432"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// deep_nested_objects - function:parse function:expand-dotted function:build-hierarchy (level 1)
func TestDeepNestedObjects(t *testing.T) {
	
	
	ccl := mock.New()
	input := `server =
  database =
    host = localhost
    port = 5432
  cache =
    enabled = true`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "server", Value: "\n  database =\n    host = localhost\n    port = 5432\n  cache =\n    enabled = true"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"server": map[string]interface{}{"cache": map[string]interface{}{"enabled": "true"}, "database": map[string]interface{}{"host": "localhost", "port": "5432"}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// duplicate_keys_to_lists - function:parse function:expand-dotted function:build-hierarchy (level 1)
func TestDuplicateKeysToLists(t *testing.T) {
	
	
	ccl := mock.New()
	input := `item = first
item = second
item = third`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "item", Value: "first"}, mock.Entry{Key: "item", Value: "second"}, mock.Entry{Key: "item", Value: "third"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"item": []interface{}{"first", "second", "third"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// nested_duplicate_keys - function:parse function:expand-dotted function:build-hierarchy (level 1)
func TestNestedDuplicateKeys(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  server = web1
  server = web2
  port = 80`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: "\n  server = web1\n  server = web2\n  port = 80"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"port": "80", "server": []interface{}{"web1", "web2"}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// mixed_flat_and_nested - function:parse function:expand-dotted function:build-hierarchy (level 1)
func TestMixedFlatAndNested(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice
config =
  debug = true
  timeout = 30
version = 1.0`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "config", Value: "\n  debug = true\n  timeout = 30"}, mock.Entry{Key: "version", Value: "1.0"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"debug": "true", "timeout": "30"}, "name": "Alice", "version": "1.0"}
	assert.Equal(t, expectedObjects, objectResult)

}


// nested_objects_with_lists - function:parse function:expand-dotted function:build-hierarchy (level 1)
func TestNestedObjectsWithLists(t *testing.T) {
	
	
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
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "environments", Value: "\n  prod =\n    server = web1\n    server = web2\n    port = 80\n  dev =\n    server = localhost\n    port = 3000"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"environments": map[string]interface{}{"dev": map[string]interface{}{"port": "3000", "server": "localhost"}, "prod": map[string]interface{}{"port": "80", "server": []interface{}{"web1", "web2"}}}}
	assert.Equal(t, expectedObjects, objectResult)

}


