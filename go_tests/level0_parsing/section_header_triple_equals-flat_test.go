package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/section_header_triple_equals-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// section_header_triple_equals_parse - function:parse (level 0)
func TestSectionHeaderTripleEqualsParse(t *testing.T) {

	ccl := mock.New()
	input := `=== Server Settings ===
host = 0.0.0.0
ssl = true`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "== Server Settings ==="}, mock.Entry{Key: "host", Value: "0.0.0.0"}, mock.Entry{Key: "ssl", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}
