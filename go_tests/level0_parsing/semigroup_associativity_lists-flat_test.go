package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/semigroup_associativity_lists-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// semigroup_associativity_lists_associativity - function:associativity (level 0)
func TestSemigroupAssociativityListsAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// semigroup_associativity_lists_parse - function:parse (level 0)
func TestSemigroupAssociativityListsParse(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
= item3`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "", Value: "item3"}}
	assert.Equal(t, expected, parseResult)

}
