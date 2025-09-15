package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/whitespace_trimming-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// whitespace_trimming_parse - function:parse feature:whitespace (level 0)
func TestWhitespaceTrimmingParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `  key   =    value with spaces   
other = normal`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value with spaces"}, mock.Entry{Key: "other", Value: "normal"}}
	assert.Equal(t, expected, parseResult)

}


