package parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_core_ccl_integration.json
// Suite: Flat Format
// Version: 1.0

// complete_basic_workflow_parse - function:parse
func TestCompleteBasicWorkflowParse(t *testing.T) {

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

// complete_basic_workflow_build_hierarchy - function:build_hierarchy
func TestCompleteBasicWorkflowBuildHierarchy(t *testing.T) {

	ccl := mock.New()
	input := `name = Alice
age = 42`

	// Declare variables for reuse across validations

	var err error

	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}

// complete_nested_workflow_parse - function:parse
func TestCompleteNestedWorkflowParse(t *testing.T) {

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432
  enabled = true`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  host = localhost\n  port = 5432\n  enabled = true"}}
	assert.Equal(t, expected, parseResult)

}

// complete_nested_workflow_build_hierarchy - function:build_hierarchy
func TestCompleteNestedWorkflowBuildHierarchy(t *testing.T) {

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432
  enabled = true`

	// Declare variables for reuse across validations

	var err error

	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}

// complete_mixed_workflow_parse - function:parse
func TestCompleteMixedWorkflowParse(t *testing.T) {

	ccl := mock.New()
	input := `app = MyApp
version = 1.0.0
config =
  debug = true
  features =
    feature1 = enabled
    feature2 = disabled`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "app", Value: "MyApp"}, mock.Entry{Key: "version", Value: "1.0.0"}, mock.Entry{Key: "config", Value: "\n  debug = true\n  features =\n    feature1 = enabled\n    feature2 = disabled"}}
	assert.Equal(t, expected, parseResult)

}

// complete_mixed_workflow_build_hierarchy - function:build_hierarchy
func TestCompleteMixedWorkflowBuildHierarchy(t *testing.T) {

	ccl := mock.New()
	input := `app = MyApp
version = 1.0.0
config =
  debug = true
  features =
    feature1 = enabled
    feature2 = disabled`

	// Declare variables for reuse across validations

	var err error

	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}

// complete_lists_workflow_parse - function:parse
func TestCompleteListsWorkflowParse(t *testing.T) {

	ccl := mock.New()
	input := `servers =
  server = web1
  server = web2
  server = web3
ports =
  port = 80
  port = 443`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "\n  server = web1\n  server = web2\n  server = web3"}, mock.Entry{Key: "ports", Value: "\n  port = 80\n  port = 443"}}
	assert.Equal(t, expected, parseResult)

}

// complete_lists_workflow_build_hierarchy - function:build_hierarchy
func TestCompleteListsWorkflowBuildHierarchy(t *testing.T) {

	ccl := mock.New()
	input := `servers =
  server = web1
  server = web2
  server = web3
ports =
  port = 80
  port = 443`

	// Declare variables for reuse across validations

	var err error

	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}

// complete_multiline_workflow_parse - function:parse feature:multiline
func TestCompleteMultilineWorkflowParse(t *testing.T) {

	ccl := mock.New()
	input := `description = Welcome to our app
  This is a multi-line description
  With several lines
config =
  settings =
    value1 = one
    value2 = two`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "description", Value: "Welcome to our app\n  This is a multi-line description\n  With several lines"}, mock.Entry{Key: "config", Value: "\n  settings =\n    value1 = one\n    value2 = two"}}
	assert.Equal(t, expected, parseResult)

}

// complete_multiline_workflow_build_hierarchy - function:build_hierarchy feature:multiline
func TestCompleteMultilineWorkflowBuildHierarchy(t *testing.T) {

	ccl := mock.New()
	input := `description = Welcome to our app
  This is a multi-line description
  With several lines
config =
  settings =
    value1 = one
    value2 = two`

	// Declare variables for reuse across validations

	var err error

	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}

// real_world_complete_workflow_parse - function:parse
func TestRealWorldCompleteWorkflowParse(t *testing.T) {

	ccl := mock.New()
	input := `service = MyMicroservice
version = 2.1.0
database =
  host = db.example.com
  port = 5432
  credentials =
    user = service_user
    password = secret123
  pools =
    read = 5
    write = 2
logging =
  level = info
  outputs =
    output = console
    output = file
    output = syslog
features =
  feature_a = enabled
  feature_b = disabled
  feature_c = experimental`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "service", Value: "MyMicroservice"}, mock.Entry{Key: "version", Value: "2.1.0"}, mock.Entry{Key: "database", Value: "\n  host = db.example.com\n  port = 5432\n  credentials =\n    user = service_user\n    password = secret123\n  pools =\n    read = 5\n    write = 2"}, mock.Entry{Key: "logging", Value: "\n  level = info\n  outputs =\n    output = console\n    output = file\n    output = syslog"}, mock.Entry{Key: "features", Value: "\n  feature_a = enabled\n  feature_b = disabled\n  feature_c = experimental"}}
	assert.Equal(t, expected, parseResult)

}

// real_world_complete_workflow_build_hierarchy - function:build_hierarchy
func TestRealWorldCompleteWorkflowBuildHierarchy(t *testing.T) {

	ccl := mock.New()
	input := `service = MyMicroservice
version = 2.1.0
database =
  host = db.example.com
  port = 5432
  credentials =
    user = service_user
    password = secret123
  pools =
    read = 5
    write = 2
logging =
  level = info
  outputs =
    output = console
    output = file
    output = syslog
features =
  feature_a = enabled
  feature_b = disabled
  feature_c = experimental`

	// Declare variables for reuse across validations

	var err error

	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}
