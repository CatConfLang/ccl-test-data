package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/semigroup_associativity_basic-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// semigroup_associativity_basic_parse - function:parse (level 0)
func TestSemigroupAssociativityBasicParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `a = 1
b = 2
c = 3`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "1"}, mock.Entry{Key: "b", Value: "2"}, mock.Entry{Key: "c", Value: "3"}}
	assert.Equal(t, expected, parseResult)

}


// semigroup_associativity_basic_associativity - function:associativity (level 0)
func TestSemigroupAssociativityBasicAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


