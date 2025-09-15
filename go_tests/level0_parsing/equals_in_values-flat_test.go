package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/equals_in_values-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// equals_in_values_parse - function:parse (level 0)
func TestEqualsInValuesParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `msg = k=v pairs work fine
path = /bin/app=prod`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "msg", Value: "k=v pairs work fine"}, mock.Entry{Key: "path", Value: "/bin/app=prod"}}
	assert.Equal(t, expected, parseResult)

}


