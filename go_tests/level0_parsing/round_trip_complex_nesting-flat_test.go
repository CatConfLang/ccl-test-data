package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_complex_nesting-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// round_trip_complex_nesting_parse - function:parse feature:empty_keys (level 0)
func TestRoundTripComplexNestingParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `app =
  = item1
  config =
    = nested_item
    db =
      host = localhost
      = db_item
  = item2`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "app", Value: "\n  = item1\n  config =\n    = nested_item\n    db =\n      host = localhost\n      = db_item\n  = item2"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_complex_nesting_round_trip - function:round_trip feature:empty_keys (level 0)
func TestRoundTripComplexNestingRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


