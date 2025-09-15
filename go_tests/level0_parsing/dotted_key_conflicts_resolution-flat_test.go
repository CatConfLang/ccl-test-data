package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/dotted_key_conflicts_resolution-flat.json
// Suite: Generated Flat Format
// Version: 1.0

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

// dotted_key_conflicts_resolution_expand_dotted - function:expand_dotted (level 0)
func TestDottedKeyConflictsResolutionExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// dotted_key_conflicts_resolution_build_hierarchy - function:build_hierarchy (level 0)
func TestDottedKeyConflictsResolutionBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
