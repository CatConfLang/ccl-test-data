package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/monoid_right_identity_nested-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// monoid_right_identity_nested_parse - function:parse (level 0)
func TestMonoidRightIdentityNestedParse(t *testing.T) {

	ccl := mock.New()
	input := `config =
  database =
    host = localhost
    port = 5432
  cache =
    redis = true`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  database =\n    host = localhost\n    port = 5432\n  cache =\n    redis = true"}}
	assert.Equal(t, expected, parseResult)

}

// monoid_right_identity_nested_associativity - function:associativity (level 0)
func TestMonoidRightIdentityNestedAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
