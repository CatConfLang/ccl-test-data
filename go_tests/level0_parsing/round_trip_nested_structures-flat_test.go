package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/round_trip_nested_structures-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// round_trip_nested_structures_parse - function:parse (level 0)
func TestRoundTripNestedStructuresParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  host = localhost
  port = 8080
  db =
    name = mydb
    user = admin`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  host = localhost\n  port = 8080\n  db =\n    name = mydb\n    user = admin"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_nested_structures_round_trip - function:round_trip (level 0)
func TestRoundTripNestedStructuresRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


