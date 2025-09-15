package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/empty_key_with_newline-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// empty_key_with_newline_parse - function:parse (level 0)
func TestEmptyKeyWithNewlineParse(t *testing.T) {

	ccl := mock.New()
	input := `
  = val`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}
