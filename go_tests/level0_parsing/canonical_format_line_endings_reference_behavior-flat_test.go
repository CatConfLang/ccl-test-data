package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/canonical_format_line_endings_reference_behavior-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// canonical_format_line_endings_reference_behavior_canonical_format - function:canonical_format behavior:crlf_preserve_literal (level 0)
func TestCanonicalFormatLineEndingsReferenceBehaviorCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// canonical_format_line_endings_reference_behavior_parse - function:parse behavior:crlf_preserve_literal (level 0)
func TestCanonicalFormatLineEndingsReferenceBehaviorParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1\r"}, mock.Entry{Key: "key2", Value: "value2\r"}}
	assert.Equal(t, expected, parseResult)

}


