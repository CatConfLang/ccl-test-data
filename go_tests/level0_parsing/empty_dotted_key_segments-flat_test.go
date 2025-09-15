package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/empty_dotted_key_segments-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// empty_dotted_key_segments_build_hierarchy - function:build_hierarchy (level 0)
func TestEmptyDottedKeySegmentsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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

// empty_dotted_key_segments_expand_dotted - function:expand_dotted (level 0)
func TestEmptyDottedKeySegmentsExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
