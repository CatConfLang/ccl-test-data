package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/spaced_equals_not_section_header-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// spaced_equals_not_section_header_parse - function:parse (level 0)
func TestSpacedEqualsNotSectionHeaderParse(t *testing.T) {

	ccl := mock.New()
	input := `= = spaced equals
=  = wide spaces
== Real Header ==
key = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= spaced equals"}, mock.Entry{Key: "", Value: "= wide spaces"}, mock.Entry{Key: "", Value: "= Real Header =="}, mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}
