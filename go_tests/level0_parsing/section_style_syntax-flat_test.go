package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/section_style_syntax-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// section_style_syntax_parse - function:parse (level 0)
func TestSectionStyleSyntaxParse(t *testing.T) {

	ccl := mock.New()
	input := `== Section 2 ==`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Section 2 =="}}
	assert.Equal(t, expected, parseResult)

}
