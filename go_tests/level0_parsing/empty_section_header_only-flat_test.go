package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/empty_section_header_only-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// empty_section_header_only_parse - function:parse (level 0)
func TestEmptySectionHeaderOnlyParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `== Empty Section ==`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Empty Section =="}}
	assert.Equal(t, expected, parseResult)

}


