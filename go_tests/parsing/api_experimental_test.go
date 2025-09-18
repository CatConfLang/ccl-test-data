package parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_experimental.json
// Suite: Flat Format
// Version: 1.0



// basic_dotted_key_expansion_parse - function:parse
func TestBasicDottedKeyExpansionParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}}
	assert.Equal(t, expected, parseResult)

}


// basic_dotted_key_expansion_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestBasicDottedKeyExpansionExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// basic_dotted_key_expansion_build_hierarchy - function:build_hierarchy
func TestBasicDottedKeyExpansionBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// multiple_dotted_keys_parse - function:parse
func TestMultipleDottedKeysParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
app.name = MyApp`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "database.port", Value: "5432"}, mock.Entry{Key: "app.name", Value: "MyApp"}}
	assert.Equal(t, expected, parseResult)

}


// multiple_dotted_keys_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestMultipleDottedKeysExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
app.name = MyApp`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// multiple_dotted_keys_build_hierarchy - function:build_hierarchy
func TestMultipleDottedKeysBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
app.name = MyApp`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// deep_dotted_nesting_parse - function:parse
func TestDeepDottedNestingParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `server.database.credentials.user = admin
server.database.credentials.pass = secret`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "server.database.credentials.user", Value: "admin"}, mock.Entry{Key: "server.database.credentials.pass", Value: "secret"}}
	assert.Equal(t, expected, parseResult)

}


// deep_dotted_nesting_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestDeepDottedNestingExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `server.database.credentials.user = admin
server.database.credentials.pass = secret`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// deep_dotted_nesting_build_hierarchy - function:build_hierarchy
func TestDeepDottedNestingBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `server.database.credentials.user = admin
server.database.credentials.pass = secret`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// mixed_dotted_and_regular_keys_parse - function:parse
func TestMixedDottedAndRegularKeysParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app = MyApp
database.host = localhost
config =
  debug = true
logging.level = info`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "app", Value: "MyApp"}, mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "config", Value: "\n  debug = true"}, mock.Entry{Key: "logging.level", Value: "info"}}
	assert.Equal(t, expected, parseResult)

}


// mixed_dotted_and_regular_keys_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestMixedDottedAndRegularKeysExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app = MyApp
database.host = localhost
config =
  debug = true
logging.level = info`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// mixed_dotted_and_regular_keys_build_hierarchy - function:build_hierarchy
func TestMixedDottedAndRegularKeysBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app = MyApp
database.host = localhost
config =
  debug = true
logging.level = info`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// dotted_key_conflicts_resolution_parse - function:parse
func TestDottedKeyConflictsResolutionParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database = old_value
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "old_value"}, mock.Entry{Key: "database.host", Value: "localhost"}}
	assert.Equal(t, expected, parseResult)

}


// dotted_key_conflicts_resolution_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestDottedKeyConflictsResolutionExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database = old_value
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// dotted_key_conflicts_resolution_build_hierarchy - function:build_hierarchy
func TestDottedKeyConflictsResolutionBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database = old_value
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// dotted_keys_with_lists_parse - function:parse
func TestDottedKeysWithListsParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers.web = web1
servers.web = web2
servers.api = api1`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers.web", Value: "web1"}, mock.Entry{Key: "servers.web", Value: "web2"}, mock.Entry{Key: "servers.api", Value: "api1"}}
	assert.Equal(t, expected, parseResult)

}


// dotted_keys_with_lists_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestDottedKeysWithListsExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers.web = web1
servers.web = web2
servers.api = api1`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// dotted_keys_with_lists_build_hierarchy - function:build_hierarchy
func TestDottedKeysWithListsBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers.web = web1
servers.web = web2
servers.api = api1`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// empty_dotted_key_segments_parse - function:parse
func TestEmptyDottedKeySegmentsParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a..b = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a..b", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// empty_dotted_key_segments_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestEmptyDottedKeySegmentsExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a..b = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// empty_dotted_key_segments_build_hierarchy - function:build_hierarchy
func TestEmptyDottedKeySegmentsBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a..b = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// single_dot_key_parse - function:parse
func TestSingleDotKeyParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a. = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a.", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// single_dot_key_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestSingleDotKeyExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a. = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// single_dot_key_build_hierarchy - function:build_hierarchy
func TestSingleDotKeyBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a. = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// hierarchical_with_expand_dotted_validation_parse - function:parse
func TestHierarchicalWithExpandDottedValidationParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  enabled = true
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  enabled = true\n  port = 5432"}}
	assert.Equal(t, expected, parseResult)

}


// hierarchical_with_expand_dotted_validation_expand_dotted - function:expand_dotted feature:experimental_dotted_keys
func TestHierarchicalWithExpandDottedValidationExpandDotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  enabled = true
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expand_dotted validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// hierarchical_with_expand_dotted_validation_build_hierarchy - function:build_hierarchy
func TestHierarchicalWithExpandDottedValidationBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  enabled = true
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// dotted_key_list_access_parse - function:parse
func TestDottedKeyListAccessParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.hosts = primary
database.hosts = secondary
database.port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database.hosts", Value: "primary"}, mock.Entry{Key: "database.hosts", Value: "secondary"}, mock.Entry{Key: "database.port", Value: "5432"}}
	assert.Equal(t, expected, parseResult)

}


// dotted_key_list_access_build_hierarchy - function:build_hierarchy
func TestDottedKeyListAccessBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.hosts = primary
database.hosts = secondary
database.port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// dotted_key_list_access_get_list - function:get_list
func TestDottedKeyListAccessGetList(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.hosts = primary
database.hosts = secondary
database.port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"database", "port"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"5432"}, result)

}


