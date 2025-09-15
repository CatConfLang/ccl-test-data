package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/real_world_complete_workflow-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// real_world_complete_workflow_parse - function:parse (level 0)
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


// real_world_complete_workflow_build_hierarchy - function:build_hierarchy (level 0)
func TestRealWorldCompleteWorkflowBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


