package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/multiple_key_value_pairs-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// multiple_key_value_pairs_parse - function:parse (level 0)
func TestMultipleKeyValuePairsParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `key1 = val1
key2 = val2`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "val1"}, mock.Entry{Key: "key2", Value: "val2"}}
	assert.Equal(t, expected, parseResult)

}


