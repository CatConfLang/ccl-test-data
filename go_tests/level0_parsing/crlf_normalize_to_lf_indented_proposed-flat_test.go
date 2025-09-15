package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/crlf_normalize_to_lf_indented_proposed-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// crlf_normalize_to_lf_indented_proposed_canonical_format - function:canonical_format behavior:crlf_normalize_to_lf (level 0)
func TestCrlfNormalizeToLfIndentedProposedCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// crlf_normalize_to_lf_indented_proposed_parse - function:parse behavior:crlf_normalize_to_lf (level 0)
func TestCrlfNormalizeToLfIndentedProposedParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}}
	assert.Equal(t, expected, parseResult)

}


