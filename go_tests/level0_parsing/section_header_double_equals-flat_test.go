package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/section_header_double_equals-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// section_header_double_equals_parse - function:parse (level 0)
func TestSectionHeaderDoubleEqualsParse(t *testing.T) {

	ccl := mock.New()
	input := `== Database Config ==
host = localhost
port = 5432`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database Config =="}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "5432"}}
	assert.Equal(t, expected, parseResult)

}
