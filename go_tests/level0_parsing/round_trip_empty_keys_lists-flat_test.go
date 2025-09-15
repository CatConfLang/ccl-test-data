package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_empty_keys_lists-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// round_trip_empty_keys_lists_parse - function:parse (level 0)
func TestRoundTripEmptyKeysListsParse(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
regular = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "regular", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_empty_keys_lists_round_trip - function:round_trip (level 0)
func TestRoundTripEmptyKeysListsRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
