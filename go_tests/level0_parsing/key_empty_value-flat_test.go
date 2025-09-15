package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/key_empty_value-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// key_empty_value_parse - function:parse feature:empty_keys (level 0)
func TestKeyEmptyValueParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key =`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


