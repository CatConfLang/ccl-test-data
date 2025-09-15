package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/monoid_left_identity_basic-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// monoid_left_identity_basic_associativity - function:associativity (level 0)
func TestMonoidLeftIdentityBasicAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// monoid_left_identity_basic_parse - function:parse (level 0)
func TestMonoidLeftIdentityBasicParse(t *testing.T) {

	ccl := mock.New()
	input := `key = value
nested =
  sub = val`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "nested", Value: "\n  sub = val"}}
	assert.Equal(t, expected, parseResult)

}
