package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/section_header_at_end-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// section_header_at_end_parse - function:parse (level 0)
func TestSectionHeaderAtEndParse(t *testing.T) {

	ccl := mock.New()
	input := `key = value
== Final Section ==`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "", Value: "= Final Section =="}}
	assert.Equal(t, expected, parseResult)

}
