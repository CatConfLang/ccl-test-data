package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/semigroup_associativity_nested-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// semigroup_associativity_nested_parse - function:parse (level 0)
func TestSemigroupAssociativityNestedParse(t *testing.T) {

	ccl := mock.New()
	input := `config =
  host = localhost
  port = 8080
db =
  name = test`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  host = localhost\n  port = 8080"}, mock.Entry{Key: "db", Value: "\n  name = test"}}
	assert.Equal(t, expected, parseResult)

}

// semigroup_associativity_nested_associativity - function:associativity (level 0)
func TestSemigroupAssociativityNestedAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
