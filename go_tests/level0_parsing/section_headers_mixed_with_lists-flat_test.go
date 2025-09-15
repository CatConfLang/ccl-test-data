package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/section_headers_mixed_with_lists-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// section_headers_mixed_with_lists_parse - function:parse (level 0)
func TestSectionHeadersMixedWithListsParse(t *testing.T) {

	ccl := mock.New()
	input := `== Configuration ==
= item1
= item2
key = value
=== Next Section ===
other = data`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Configuration =="}, mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "", Value: "== Next Section ==="}, mock.Entry{Key: "other", Value: "data"}}
	assert.Equal(t, expected, parseResult)

}
