package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/array_style_list-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// array_style_list_parse - function:parse feature:empty_keys (level 0)
func TestArrayStyleListParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `1 =
2 =
3 =`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "1", Value: ""}, mock.Entry{Key: "2", Value: ""}, mock.Entry{Key: "3", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


