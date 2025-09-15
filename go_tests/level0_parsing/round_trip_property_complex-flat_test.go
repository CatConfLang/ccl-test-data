package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_property_complex-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// round_trip_property_complex_parse - function:parse (level 0)
func TestRoundTripPropertyComplexParse(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
config =
  nested =
    deep = value
  list =
    = a
    = b
    = c
final = end`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "config", Value: "\n  nested =\n    deep = value\n  list =\n    = a\n    = b\n    = c"}, mock.Entry{Key: "final", Value: "end"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_property_complex_round_trip - function:round_trip (level 0)
func TestRoundTripPropertyComplexRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
