package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_multiline_values-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// round_trip_multiline_values_parse - function:parse (level 0)
func TestRoundTripMultilineValuesParse(t *testing.T) {

	ccl := mock.New()
	input := `script =
  #!/bin/bash
  echo hello
  exit 0`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "script", Value: "\n  #!/bin/bash\n  echo hello\n  exit 0"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_multiline_values_round_trip - function:round_trip (level 0)
func TestRoundTripMultilineValuesRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
