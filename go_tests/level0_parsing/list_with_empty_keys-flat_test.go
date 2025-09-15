package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_with_empty_keys-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// list_with_empty_keys_parse - function:parse feature:empty_keys (level 0)
func TestListWithEmptyKeysParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `= 3
= 1
= 2`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "3"}, mock.Entry{Key: "", Value: "1"}, mock.Entry{Key: "", Value: "2"}}
	assert.Equal(t, expected, parseResult)

}


