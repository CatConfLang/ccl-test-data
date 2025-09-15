package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/whitespace_only_value-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// whitespace_only_value_parse - function:parse feature:empty_keys feature:whitespace (level 0)
func TestWhitespaceOnlyValueParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `onlyspaces =     `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "onlyspaces", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


