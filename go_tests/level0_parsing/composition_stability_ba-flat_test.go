package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/composition_stability_ba-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// composition_stability_ba_parse - function:parse (level 0)
func TestCompositionStabilityBaParse(t *testing.T) {

	ccl := mock.New()
	input := `b = 20
c = 3
a = 1
b = 2`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "b", Value: "20"}, mock.Entry{Key: "c", Value: "3"}, mock.Entry{Key: "a", Value: "1"}, mock.Entry{Key: "b", Value: "2"}}
	assert.Equal(t, expected, parseResult)

}
