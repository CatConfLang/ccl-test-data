package level3_dotted_keys_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-dotted-keys.json
// Suite: CCL Dotted Key Expansion - Validation Format
// Version: 2.0
// Description: Expanding dotted keys (e.g., 'database.host') into hierarchical structures - enables dual access patterns


// basic_single_dotted_key - basic dotted-keys single (level 3)
func TestBasicSingleDottedKey(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"host": "localhost"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// basic_multiple_dotted_keys_same_parent - basic dotted-keys multiple same-parent (level 3)
func TestBasicMultipleDottedKeysSameParent(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
database.name = mydb`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "database.port", Value: "5432"}, mock.Entry{Key: "database.name", Value: "mydb"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"host": "localhost", "name": "mydb", "port": "5432"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// multiple_dotted_keys_different_parents - basic dotted-keys multiple different-parents (level 3)
func TestMultipleDottedKeysDifferentParents(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
server.port = 8080
cache.enabled = true`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "server.port", Value: "8080"}, mock.Entry{Key: "cache.enabled", Value: "true"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"cache": map[string]interface{}{"enabled": "true"}, "database": map[string]interface{}{"host": "localhost"}, "server": map[string]interface{}{"port": "8080"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// deep_dotted_nesting_three_levels - deep-nesting dotted-keys three-levels (level 3)
func TestDeepDottedNestingThreeLevels(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app.database.host = localhost
app.database.port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "app.database.host", Value: "localhost"}, mock.Entry{Key: "app.database.port", Value: "5432"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"app": map[string]interface{}{"database": map[string]interface{}{"host": "localhost", "port": "5432"}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// deep_dotted_nesting_four_levels - deep-nesting dotted-keys four-levels extreme (level 3)
func TestDeepDottedNestingFourLevels(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config.app.database.connection.host = localhost
config.app.database.connection.timeout = 30s`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config.app.database.connection.host", Value: "localhost"}, mock.Entry{Key: "config.app.database.connection.timeout", Value: "30s"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"app": map[string]interface{}{"database": map[string]interface{}{"connection": map[string]interface{}{"host": "localhost", "timeout": "30s"}}}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// mixed_dotted_and_regular_keys - mixed dotted-keys regular-keys (level 3)
func TestMixedDottedAndRegularKeys(t *testing.T) {
	
	
	ccl := mock.New()
	input := `simple = value
database.host = localhost
another = test
server.port = 8080`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "simple", Value: "value"}, mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "another", Value: "test"}, mock.Entry{Key: "server.port", Value: "8080"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"another": "test", "database": map[string]interface{}{"host": "localhost"}, "server": map[string]interface{}{"port": "8080"}, "simple": "value"}
	assert.Equal(t, expectedObjects, objectResult)

}


// conflict_dotted_vs_nested_dotted_wins - conflict-resolution dotted-wins merge (level 3)
func TestConflictDottedVsNestedDottedWins(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  port = 3306
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "\n  port = 3306"}, mock.Entry{Key: "database.host", Value: "localhost"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"host": "localhost", "port": "3306"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// conflict_nested_vs_dotted_merge - conflict-resolution merge order-independent (level 3)
func TestConflictNestedVsDottedMerge(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database =
  port = 3306
  name = mydb`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "database", Value: "\n  port = 3306\n  name = mydb"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"host": "localhost", "name": "mydb", "port": "3306"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// complex_conflict_multiple_sources - complex-conflict multiple-sources merge (level 3)
func TestComplexConflictMultipleSources(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app.database.host = localhost
app =
  name = myapp
app.database.port = 5432
app =
  version = 1.0`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "app.database.host", Value: "localhost"}, mock.Entry{Key: "app", Value: "\n  name = myapp"}, mock.Entry{Key: "app.database.port", Value: "5432"}, mock.Entry{Key: "app", Value: "\n  version = 1.0"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"app": map[string]interface{}{"database": map[string]interface{}{"host": "localhost", "port": "5432"}, "name": "myapp", "version": "1.0"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// dotted_keys_with_empty_key_parent - empty-key dotted-keys edge-case (level 3)
func TestDottedKeysWithEmptyKeyParent(t *testing.T) {
	
	
	ccl := mock.New()
	input := `.host = localhost
.port = 8080`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: ".host", Value: "localhost"}, mock.Entry{Key: ".port", Value: "8080"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"": map[string]interface{}{"host": "localhost", "port": "8080"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// dotted_keys_with_empty_key_child - empty-key dotted-keys trailing-dot (level 3)
func TestDottedKeysWithEmptyKeyChild(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database. = default_value
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database.", Value: "default_value"}, mock.Entry{Key: "database.host", Value: "localhost"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"": "default_value", "host": "localhost"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// edge_case_single_dot - edge-case single-dot empty-key (level 3)
func TestEdgeCaseSingleDot(t *testing.T) {
	
	
	ccl := mock.New()
	input := `. = dot_value
key = normal`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: ".", Value: "dot_value"}, mock.Entry{Key: "key", Value: "normal"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"": map[string]interface{}{"": "dot_value"}, "key": "normal"}
	assert.Equal(t, expectedObjects, objectResult)

}


// edge_case_multiple_consecutive_dots - edge-case consecutive-dots empty-segments (level 3)
func TestEdgeCaseMultipleConsecutiveDots(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app..config = value
db...host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "app..config", Value: "value"}, mock.Entry{Key: "db...host", Value: "localhost"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"app": map[string]interface{}{"": map[string]interface{}{"config": "value"}}, "db": map[string]interface{}{"": map[string]interface{}{"": map[string]interface{}{"host": "localhost"}}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// dotted_keys_with_list_style - list-style dotted-keys duplicate-keys (level 3)
func TestDottedKeysWithListStyle(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers.production = server1
servers.production = server2
servers.staging = staging1`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "servers.production", Value: "server1"}, mock.Entry{Key: "servers.production", Value: "server2"}, mock.Entry{Key: "servers.staging", Value: "staging1"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"servers": map[string]interface{}{"production": []interface{}{"server1", "server2"}, "staging": "staging1"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// dotted_keys_with_empty_key_lists - empty-key list-style dotted-keys mixed (level 3)
func TestDottedKeysWithEmptyKeyLists(t *testing.T) {
	
	
	ccl := mock.New()
	input := `ports. = 8000
ports. = 8001
ports.admin = 9000`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "ports.", Value: "8000"}, mock.Entry{Key: "ports.", Value: "8001"}, mock.Entry{Key: "ports.admin", Value: "9000"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"ports": map[string]interface{}{"": []interface{}{"8000", "8001"}, "admin": "9000"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// backward_compatibility_no_dots - backward-compatibility no-dots unchanged (level 3)
func TestBackwardCompatibilityNoDots(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database = value
host = localhost
port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "value"}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "5432"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"database": "value", "host": "localhost", "port": "5432"}
	assert.Equal(t, expectedObjects, objectResult)

}


// backward_compatibility_nested_syntax - backward-compatibility nested-syntax unchanged (level 3)
func TestBackwardCompatibilityNestedSyntax(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "\n  host = localhost\n  port = 5432"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"host": "localhost", "port": "5432"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// stress_test_complex_mixed_scenario - stress-test complex mixed real-world (level 3)
func TestStressTestComplexMixedScenario(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/ = CCL with dotted keys
app.name = MyApp
app.version = 1.0

database =
  enabled = true

database.connection.host = localhost
database.connection.port = 5432
database.pool.size = 10

servers.web = server1
servers.web = server2
servers.api.endpoint = /api/v1

config.debug = true
config =
  env = production`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "/", Value: "CCL with dotted keys"}, mock.Entry{Key: "app.name", Value: "MyApp"}, mock.Entry{Key: "app.version", Value: "1.0"}, mock.Entry{Key: "database", Value: "\n  enabled = true"}, mock.Entry{Key: "database.connection.host", Value: "localhost"}, mock.Entry{Key: "database.connection.port", Value: "5432"}, mock.Entry{Key: "database.pool.size", Value: "10"}, mock.Entry{Key: "servers.web", Value: "server1"}, mock.Entry{Key: "servers.web", Value: "server2"}, mock.Entry{Key: "servers.api.endpoint", Value: "/api/v1"}, mock.Entry{Key: "config.debug", Value: "true"}, mock.Entry{Key: "config", Value: "\n  env = production"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"/": "CCL with dotted keys", "app": map[string]interface{}{"name": "MyApp", "version": "1.0"}, "config": map[string]interface{}{"debug": "true", "env": "production"}, "database": map[string]interface{}{"connection": map[string]interface{}{"host": "localhost", "port": "5432"}, "enabled": "true", "pool": map[string]interface{}{"size": "10"}}, "servers": map[string]interface{}{"api": map[string]interface{}{"endpoint": "/api/v1"}, "web": []interface{}{"server1", "server2"}}}
	assert.Equal(t, expectedObjects, objectResult)

}


