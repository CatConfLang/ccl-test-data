package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/canonical_format_consistent_spacing-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// canonical_format_consistent_spacing_canonical_format - function:canonical_format (level 0)
func TestCanonicalFormatConsistentSpacingCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_consistent_spacing_parse - function:parse (level 0)
func TestCanonicalFormatConsistentSpacingParse(t *testing.T) {

	ccl := mock.New()
	input := `key1=value1
key2  =  value2
key3	=	value3`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}, mock.Entry{Key: "key3", Value: "\tvalue3"}}
	assert.Equal(t, expected, parseResult)

}
