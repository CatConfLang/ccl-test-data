package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/consecutive_section_headers-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// consecutive_section_headers_parse - function:parse (level 0)
func TestConsecutiveSectionHeadersParse(t *testing.T) {

	ccl := mock.New()
	input := `== First Section ==
=== Nested Section ===
==== Deep Section ====
key = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= First Section =="}, mock.Entry{Key: "", Value: "== Nested Section ==="}, mock.Entry{Key: "", Value: "=== Deep Section ===="}, mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}
