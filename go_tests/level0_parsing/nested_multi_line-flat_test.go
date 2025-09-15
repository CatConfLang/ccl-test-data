package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/nested_multi_line-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// nested_multi_line_parse - function:parse (level 0)
func TestNestedMultiLineParse(t *testing.T) {

	ccl := mock.New()
	input := `key =
  line1
  line2`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\n  line1\n  line2"}}
	assert.Equal(t, expected, parseResult)

}
