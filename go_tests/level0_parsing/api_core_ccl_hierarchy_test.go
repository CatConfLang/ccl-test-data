package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_core_ccl_hierarchy.json
// Suite: Flat Format
// Version: 1.0

// basic_object_construction_parse - function:parse (level 0)
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

// basic_object_construction_buildhierarchy - function:buildhierarchy (level 0)
func TestBasicObjectConstructionBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deep_nested_objects_parse - function:parse (level 0)
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

// deep_nested_objects_buildhierarchy - function:buildhierarchy (level 0)
func TestDeepNestedObjectsBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// duplicate_keys_to_lists_parse - function:parse (level 0)
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

// duplicate_keys_to_lists_buildhierarchy - function:buildhierarchy (level 0)
func TestDuplicateKeysToListsBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

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

// nested_duplicate_keys_buildhierarchy - function:buildhierarchy (level 0)
func TestNestedDuplicateKeysBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

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

// mixed_flat_and_nested_buildhierarchy - function:buildhierarchy (level 0)
func TestMixedFlatAndNestedBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_objects_with_lists_parse - function:parse (level 0)
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

// nested_objects_with_lists_buildhierarchy - function:buildhierarchy (level 0)
func TestNestedObjectsWithListsBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
