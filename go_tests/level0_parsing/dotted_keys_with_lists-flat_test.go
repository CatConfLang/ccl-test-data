package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/dotted_keys_with_lists-flat.json
// Suite: Generated Flat Format
// Version: 1.0

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

// dotted_keys_with_lists_expand_dotted - function:expand_dotted (level 0)
func TestDottedKeysWithListsExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// dotted_keys_with_lists_build_hierarchy - function:build_hierarchy (level 0)
func TestDottedKeysWithListsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
