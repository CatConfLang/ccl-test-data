package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/nested_structure_parsing-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// nested_structure_parsing_parse - function:parse (level 0)
func TestNestedStructureParsingParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  host = localhost\n  port = 5432"}}
	assert.Equal(t, expected, parseResult)

}


