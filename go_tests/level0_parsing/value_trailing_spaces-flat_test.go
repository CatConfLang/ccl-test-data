package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/value_trailing_spaces-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// value_trailing_spaces_parse - function:parse feature:whitespace (level 0)
func TestValueTrailingSpacesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key = val  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


