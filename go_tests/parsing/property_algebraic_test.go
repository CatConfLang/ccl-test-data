package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/property_algebraic.json
// Suite: Flat Format
// Version: 1.0



// semigroup_associativity_basic_compose_associative - function:compose_associative
func TestSemigroupAssociativityBasicComposeAssociative(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// semigroup_associativity_nested_compose_associative - function:compose_associative
func TestSemigroupAssociativityNestedComposeAssociative(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// semigroup_associativity_lists_compose_associative - function:compose_associative feature:empty_keys
func TestSemigroupAssociativityListsComposeAssociative(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// monoid_left_identity_basic_identity_left - function:identity_left
func TestMonoidLeftIdentityBasicIdentityLeft(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// monoid_right_identity_basic_identity_right - function:identity_right
func TestMonoidRightIdentityBasicIdentityRight(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// monoid_left_identity_nested_identity_left - function:identity_left
func TestMonoidLeftIdentityNestedIdentityLeft(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// monoid_right_identity_nested_identity_right - function:identity_right
func TestMonoidRightIdentityNestedIdentityRight(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// monoid_left_identity_lists_identity_left - function:identity_left feature:empty_keys
func TestMonoidLeftIdentityListsIdentityLeft(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// monoid_right_identity_lists_identity_right - function:identity_right feature:empty_keys
func TestMonoidRightIdentityListsIdentityRight(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// round_trip_property_basic_parse - function:parse
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


// round_trip_property_basic_round_trip - function:round_trip
func TestRoundTripPropertyBasicRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// round_trip_property_nested_parse - function:parse
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


// round_trip_property_nested_round_trip - function:round_trip
func TestRoundTripPropertyNestedRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// round_trip_property_complex_parse - function:parse feature:empty_keys
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


// round_trip_property_complex_round_trip - function:round_trip feature:empty_keys
func TestRoundTripPropertyComplexRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


