package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_property_basic-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// round_trip_property_basic_parse - function:parse (level 0)
func TestRoundTripPropertyBasicParse(t *testing.T) {

	ccl := mock.New()
	input := `key = value
another = test`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "another", Value: "test"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_property_basic_round_trip - function:round_trip (level 0)
func TestRoundTripPropertyBasicRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
