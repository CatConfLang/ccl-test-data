package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/mixed_keys_with_duplicates-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// mixed_keys_with_duplicates_parse - function:parse feature:empty_keys (level 0)
func TestMixedKeysWithDuplicatesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = app
ports = 8000
name = service
ports = 8001`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "app"}, mock.Entry{Key: "ports", Value: "8000"}, mock.Entry{Key: "name", Value: "service"}, mock.Entry{Key: "ports", Value: "8001"}}
	assert.Equal(t, expected, parseResult)

}


