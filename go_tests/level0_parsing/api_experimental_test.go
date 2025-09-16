package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_experimental.json
// Suite: Flat Format
// Version: 1.0



// basic_dotted_key_expansion_parse - function:parse (level 0)
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


// basic_dotted_key_expansion_expanddotted - function:expanddotted (level 0)
func TestBasicDottedKeyExpansionExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// basic_dotted_key_expansion_buildhierarchy - function:buildhierarchy (level 0)
func TestBasicDottedKeyExpansionBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// multiple_dotted_keys_parse - function:parse (level 0)
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


// multiple_dotted_keys_expanddotted - function:expanddotted (level 0)
func TestMultipleDottedKeysExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
app.name = MyApp`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// multiple_dotted_keys_buildhierarchy - function:buildhierarchy (level 0)
func TestMultipleDottedKeysBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
app.name = MyApp`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// deep_dotted_nesting_parse - function:parse (level 0)
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


// deep_dotted_nesting_expanddotted - function:expanddotted (level 0)
func TestDeepDottedNestingExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `server.database.credentials.user = admin
server.database.credentials.pass = secret`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// deep_dotted_nesting_buildhierarchy - function:buildhierarchy (level 0)
func TestDeepDottedNestingBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `server.database.credentials.user = admin
server.database.credentials.pass = secret`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// mixed_dotted_and_regular_keys_parse - function:parse (level 0)
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


// mixed_dotted_and_regular_keys_expanddotted - function:expanddotted (level 0)
func TestMixedDottedAndRegularKeysExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app = MyApp
database.host = localhost
config =
  debug = true
logging.level = info`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// mixed_dotted_and_regular_keys_buildhierarchy - function:buildhierarchy (level 0)
func TestMixedDottedAndRegularKeysBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app = MyApp
database.host = localhost
config =
  debug = true
logging.level = info`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// dotted_key_conflicts_resolution_parse - function:parse (level 0)
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


// dotted_key_conflicts_resolution_expanddotted - function:expanddotted (level 0)
func TestDottedKeyConflictsResolutionExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database = old_value
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// dotted_key_conflicts_resolution_buildhierarchy - function:buildhierarchy (level 0)
func TestDottedKeyConflictsResolutionBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database = old_value
database.host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// dotted_keys_with_lists_parse - function:parse (level 0)
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


// dotted_keys_with_lists_expanddotted - function:expanddotted (level 0)
func TestDottedKeysWithListsExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers.web = web1
servers.web = web2
servers.api = api1`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// dotted_keys_with_lists_buildhierarchy - function:buildhierarchy (level 0)
func TestDottedKeysWithListsBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers.web = web1
servers.web = web2
servers.api = api1`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// empty_dotted_key_segments_parse - function:parse (level 0)
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


// empty_dotted_key_segments_expanddotted - function:expanddotted (level 0)
func TestEmptyDottedKeySegmentsExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a..b = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// empty_dotted_key_segments_buildhierarchy - function:buildhierarchy (level 0)
func TestEmptyDottedKeySegmentsBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a..b = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// single_dot_key_parse - function:parse (level 0)
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


// single_dot_key_expanddotted - function:expanddotted (level 0)
func TestSingleDotKeyExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a. = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// single_dot_key_buildhierarchy - function:buildhierarchy (level 0)
func TestSingleDotKeyBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a. = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// hierarchical_with_expand_dotted_validation_parse - function:parse (level 0)
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


// hierarchical_with_expand_dotted_validation_expanddotted - function:expanddotted (level 0)
func TestHierarchicalWithExpandDottedValidationExpanddotted(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  enabled = true
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement expanddotted validation

}


// hierarchical_with_expand_dotted_validation_buildhierarchy - function:buildhierarchy (level 0)
func TestHierarchicalWithExpandDottedValidationBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  enabled = true
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// dotted_key_list_access_parse - function:parse (level 0)
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


// dotted_key_list_access_buildhierarchy - function:buildhierarchy (level 0)
func TestDottedKeyListAccessBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.hosts = primary
database.hosts = secondary
database.port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// dotted_key_list_access_getlist - function:getlist (level 0)
func TestDottedKeyListAccessGetlist(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.hosts = primary
database.hosts = secondary
database.port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getlist validation

}


