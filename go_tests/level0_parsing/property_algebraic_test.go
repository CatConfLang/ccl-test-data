package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/property_algebraic.json
// Suite: Flat Format
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

// semigroup_associativity_lists_associativity - function:associativity (level 0)
func TestSemigroupAssociativityListsAssociativity(t *testing.T) {
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

// monoid_left_identity_basic_associativity - function:associativity (level 0)
func TestMonoidLeftIdentityBasicAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// monoid_right_identity_basic_parse - function:parse (level 0)
func TestMonoidRightIdentityBasicParse(t *testing.T) {

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

// monoid_right_identity_basic_associativity - function:associativity (level 0)
func TestMonoidRightIdentityBasicAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// monoid_left_identity_nested_parse - function:parse (level 0)
func TestMonoidLeftIdentityNestedParse(t *testing.T) {

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

// monoid_left_identity_nested_associativity - function:associativity (level 0)
func TestMonoidLeftIdentityNestedAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

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

// monoid_left_identity_lists_parse - function:parse (level 0)
func TestMonoidLeftIdentityListsParse(t *testing.T) {

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

// monoid_left_identity_lists_associativity - function:associativity (level 0)
func TestMonoidLeftIdentityListsAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// monoid_right_identity_lists_parse - function:parse (level 0)
func TestMonoidRightIdentityListsParse(t *testing.T) {

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

// monoid_right_identity_lists_associativity - function:associativity (level 0)
func TestMonoidRightIdentityListsAssociativity(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_property_basic_parse - function:parse (level 0)
func TestRoundTripPropertyBasicParse(t *testing.T) {

	ccl := mock.New()
	input := `key = value
another = test`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "another", Value: "test"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_property_basic_round_trip - function:round_trip (level 0)
func TestRoundTripPropertyBasicRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_property_nested_parse - function:parse (level 0)
func TestRoundTripPropertyNestedParse(t *testing.T) {

	ccl := mock.New()
	input := `config =
  host = localhost
  port = 8080
  db =
    name = mydb
    user = admin`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  host = localhost\n  port = 8080\n  db =\n    name = mydb\n    user = admin"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_property_nested_round_trip - function:round_trip (level 0)
func TestRoundTripPropertyNestedRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_property_complex_parse - function:parse (level 0)
func TestRoundTripPropertyComplexParse(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
config =
  nested =
    deep = value
  list =
    = a
    = b
    = c
final = end`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "config", Value: "\n  nested =\n    deep = value\n  list =\n    = a\n    = b\n    = c"}, mock.Entry{Key: "final", Value: "end"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_property_complex_round_trip - function:round_trip (level 0)
func TestRoundTripPropertyComplexRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
