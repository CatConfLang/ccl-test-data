package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_deeply_nested-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// round_trip_deeply_nested_parse - function:parse (level 0)
func TestRoundTripDeeplyNestedParse(t *testing.T) {

	ccl := mock.New()
	input := `level1 =
  level2 =
    level3 =
      level4 =
        deep = value
        = deep_item`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "level1", Value: "\n  level2 =\n    level3 =\n      level4 =\n        deep = value\n        = deep_item"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_deeply_nested_round_trip - function:round_trip (level 0)
func TestRoundTripDeeplyNestedRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
