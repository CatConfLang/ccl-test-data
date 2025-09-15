package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/unicode_parsing-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// unicode_parsing_parse - function:parse feature:unicode (level 0)
func TestUnicodeParsingParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `emoji = 😀😃😄
配置 = config`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "emoji", Value: "😀😃😄"}, mock.Entry{Key: "配置", Value: "config"}}
	assert.Equal(t, expected, parseResult)

}


