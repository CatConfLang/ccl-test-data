package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_whitespace_normalization-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// round_trip_whitespace_normalization_round_trip - function:round_trip (level 0)
func TestRoundTripWhitespaceNormalizationRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_whitespace_normalization_parse - function:parse (level 0)
func TestRoundTripWhitespaceNormalizationParse(t *testing.T) {

	ccl := mock.New()
	input := `  key  =  value  
  nested  = 
    sub  =  val  `

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value  \n  nested  = \n    sub  =  val"}}
	assert.Equal(t, expected, parseResult)

}
