package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/surrounded_by_newlines-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// surrounded_by_newlines_parse - function:parse (level 0)
func TestSurroundedByNewlinesParse(t *testing.T) {

	ccl := mock.New()
	input := `
key = val
`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}
