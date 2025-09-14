package level4_experimental_dotted_keys_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-experimental.json
// Suite: CCL Experimental Features - Validation Format
// Version: 2.1
// Description: Experimental and implementation-specific CCL features - NOT part of standard specification, subject to change.


// basic_dotted_key_expansion - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy (level 4)
func TestBasicDottedKeyExpansion(t *testing.T) {
	
	
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
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"host": "localhost"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// multiple_dotted_keys - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy (level 4)
func TestMultipleDottedKeys(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
app.name = MyApp`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "database.port", Value: "5432"}, mock.Entry{Key: "app.name", Value: "MyApp"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"app": map[string]interface{}{"name": "MyApp"}, "database": map[string]interface{}{"host": "localhost", "port": "5432"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// deep_dotted_nesting - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy (level 4)
func TestDeepDottedNesting(t *testing.T) {
	
	
	ccl := mock.New()
	input := `server.database.credentials.user = admin
server.database.credentials.pass = secret`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "server.database.credentials.user", Value: "admin"}, mock.Entry{Key: "server.database.credentials.pass", Value: "secret"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"server": map[string]interface{}{"database": map[string]interface{}{"credentials": map[string]interface{}{"pass": "secret", "user": "admin"}}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// mixed_dotted_and_regular_keys - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy (level 4)
func TestMixedDottedAndRegularKeys(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app = MyApp
database.host = localhost
config =
  debug = true
logging.level = info`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "app", Value: "MyApp"}, mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "config", Value: "\n  debug = true"}, mock.Entry{Key: "logging.level", Value: "info"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"app": "MyApp", "config": map[string]interface{}{"debug": "true"}, "database": map[string]interface{}{"host": "localhost"}, "logging": map[string]interface{}{"level": "info"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// dotted_key_conflicts_resolution - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy (level 4)
func TestDottedKeyConflictsResolution(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database = old_value
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "old_value"}, mock.Entry{Key: "database.host", Value: "localhost"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"host": "localhost"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// dotted_keys_with_lists - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy (level 4)
func TestDottedKeysWithLists(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers.web = web1
servers.web = web2
servers.api = api1`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "servers.web", Value: "web1"}, mock.Entry{Key: "servers.web", Value: "web2"}, mock.Entry{Key: "servers.api", Value: "api1"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"servers": map[string]interface{}{"api": "api1", "web": []interface{}{"web1", "web2"}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// empty_dotted_key_segments - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy feature:empty-keys (level 4)
func TestEmptyDottedKeySegments(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a..b = value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "a..b", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"a": map[string]interface{}{"": map[string]interface{}{"b": "value"}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// single_dot_key - feature:experimental-dotted-keys function:expand-dotted function:build-hierarchy feature:empty-keys (level 4)
func TestSingleDotKey(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a. = value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "a.", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"a": map[string]interface{}{"": "value"}}
	assert.Equal(t, expectedObjects, objectResult)

}


