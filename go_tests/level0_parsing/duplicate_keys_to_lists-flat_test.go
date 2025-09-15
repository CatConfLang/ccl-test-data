package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/duplicate_keys_to_lists-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// duplicate_keys_to_lists_build_hierarchy - function:build_hierarchy (level 0)
func TestDuplicateKeysToListsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// duplicate_keys_to_lists_load - function:load (level 0)
func TestDuplicateKeysToListsLoad(t *testing.T) {
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
