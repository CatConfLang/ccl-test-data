package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/deep_dotted_nesting-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// deep_dotted_nesting_parse - function:parse feature:experimental_dotted_keys (level 0)
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


// deep_dotted_nesting_expand_dotted - function:expand_dotted feature:experimental_dotted_keys (level 0)
func TestDeepDottedNestingExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// deep_dotted_nesting_build_hierarchy - function:build_hierarchy feature:experimental_dotted_keys (level 0)
func TestDeepDottedNestingBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


