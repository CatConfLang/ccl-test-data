package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/canonical_format_unicode-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// canonical_format_unicode_parse - function:parse feature:unicode (level 0)
func TestCanonicalFormatUnicodeParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `unicode = 你好世界
emo = 🌟✨`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "unicode", Value: "你好世界"}, mock.Entry{Key: "emo", Value: "🌟✨"}}
	assert.Equal(t, expected, parseResult)

}


// canonical_format_unicode_canonical_format - function:canonical_format feature:unicode (level 0)
func TestCanonicalFormatUnicodeCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


