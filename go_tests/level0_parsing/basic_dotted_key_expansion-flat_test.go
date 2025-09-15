package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/basic_dotted_key_expansion-flat.json
// Suite: Generated Flat Format
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

// basic_dotted_key_expansion_expand_dotted - function:expand_dotted (level 0)
func TestBasicDottedKeyExpansionExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// basic_dotted_key_expansion_build_hierarchy - function:build_hierarchy (level 0)
func TestBasicDottedKeyExpansionBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
