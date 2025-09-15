package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/canonical_format_tab_preservation-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// canonical_format_tab_preservation_parse - function:parse behavior:tabs_preserve (level 0)
func TestCanonicalFormatTabPreservationParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `value_with_tabs = text		with	tabs	`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "value_with_tabs", Value: "text\t\twith\ttabs\t"}}
	assert.Equal(t, expected, parseResult)

}


// canonical_format_tab_preservation_canonical_format - function:canonical_format behavior:tabs_preserve (level 0)
func TestCanonicalFormatTabPreservationCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


