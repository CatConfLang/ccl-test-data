package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/quotes_treated_as_literal_unquoted-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// quotes_treated_as_literal_unquoted_parse - function:parse (level 0)
func TestQuotesTreatedAsLiteralUnquotedParse(t *testing.T) {

	ccl := mock.New()
	input := `host = localhost`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}}
	assert.Equal(t, expected, parseResult)

}
