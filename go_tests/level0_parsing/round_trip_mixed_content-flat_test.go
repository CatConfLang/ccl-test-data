package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_mixed_content-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// round_trip_mixed_content_parse - function:parse (level 0)
func TestRoundTripMixedContentParse(t *testing.T) {

	ccl := mock.New()
	input := `name = Alice
= first item
config =
  port = 3000
= second item
final = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "", Value: "first item"}, mock.Entry{Key: "config", Value: "\n  port = 3000"}, mock.Entry{Key: "", Value: "second item"}, mock.Entry{Key: "final", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_mixed_content_round_trip - function:round_trip (level 0)
func TestRoundTripMixedContentRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
