package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/multiline_values-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// multiline_values_parse - function:parse (level 0)
func TestMultilineValuesParse(t *testing.T) {

	ccl := mock.New()
	input := `description = First line
  Second line
  Third line
done = yes`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "description", Value: "First line\n  Second line\n  Third line"}, mock.Entry{Key: "done", Value: "yes"}}
	assert.Equal(t, expected, parseResult)

}
