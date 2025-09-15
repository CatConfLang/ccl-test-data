package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/composition_stability_duplicate_keys-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// composition_stability_duplicate_keys_parse - function:parse (level 0)
func TestCompositionStabilityDuplicateKeysParse(t *testing.T) {

	ccl := mock.New()
	input := `a = 1
b = 2
b = 20
c = 3`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "1"}, mock.Entry{Key: "b", Value: "2"}, mock.Entry{Key: "b", Value: "20"}, mock.Entry{Key: "c", Value: "3"}}
	assert.Equal(t, expected, parseResult)

}
