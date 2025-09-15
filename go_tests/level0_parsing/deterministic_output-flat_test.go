package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/deterministic_output-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// deterministic_output_parse - function:parse (level 0)
func TestDeterministicOutputParse(t *testing.T) {

	ccl := mock.New()
	input := `z = last
a = first
m = middle`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "z", Value: "last"}, mock.Entry{Key: "a", Value: "first"}, mock.Entry{Key: "m", Value: "middle"}}
	assert.Equal(t, expected, parseResult)

}

// deterministic_output_canonical_format - function:canonical_format (level 0)
func TestDeterministicOutputCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
