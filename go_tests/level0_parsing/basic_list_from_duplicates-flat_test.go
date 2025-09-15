package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/basic_list_from_duplicates-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// basic_list_from_duplicates_parse - function:parse (level 0)
func TestBasicListFromDuplicatesParse(t *testing.T) {

	ccl := mock.New()
	input := `servers = web1
servers = web2
servers = web3`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "web3"}}
	assert.Equal(t, expected, parseResult)

}

// basic_list_from_duplicates_build_hierarchy - function:build_hierarchy (level 0)
func TestBasicListFromDuplicatesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// basic_list_from_duplicates_get_list - function:get_list (level 0)
func TestBasicListFromDuplicatesGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
