package level4_algebraic_properties_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/property-algebraic.json
// Suite: CCL Algebraic Properties
// Version: 2.0
// Description: Algebraic property tests including semigroup associativity, monoid identity laws, and round-trip properties that ensure mathematical correctness of CCL operations

// semigroup_associativity_basic - function:parse function:combine (level 4)
func TestSemigroupAssociativityBasic(t *testing.T) {

	ccl := mock.New()
	input := `a = 1
b = 2
c = 3`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "a", Value: "1"}, mock.Entry{Key: "b", Value: "2"}, mock.Entry{Key: "c", Value: "3"}}
	assert.Equal(t, expectedParse, parseResult)

}

// semigroup_associativity_nested - function:parse function:combine function:build_hierarchy (level 4)
func TestSemigroupAssociativityNested(t *testing.T) {

	ccl := mock.New()
	input := `config =
  host = localhost
  port = 8080
db =
  name = test`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: "\n  host = localhost\n  port = 8080"}, mock.Entry{Key: "db", Value: "\n  name = test"}}
	assert.Equal(t, expectedParse, parseResult)

}

// semigroup_associativity_lists - function:parse feature:empty_keys function:combine (level 4)
func TestSemigroupAssociativityLists(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
= item3`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "", Value: "item3"}}
	assert.Equal(t, expectedParse, parseResult)

}

// monoid_left_identity_basic - function:parse function:combine (level 4)
func TestMonoidLeftIdentityBasic(t *testing.T) {

	ccl := mock.New()
	input := `key = value
nested =
  sub = val`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "nested", Value: "\n  sub = val"}}
	assert.Equal(t, expectedParse, parseResult)

}

// monoid_right_identity_basic - function:parse function:combine (level 4)
func TestMonoidRightIdentityBasic(t *testing.T) {

	ccl := mock.New()
	input := `key = value
nested =
  sub = val`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "nested", Value: "\n  sub = val"}}
	assert.Equal(t, expectedParse, parseResult)

}

// monoid_left_identity_nested - function:parse function:combine function:build_hierarchy (level 4)
func TestMonoidLeftIdentityNested(t *testing.T) {

	ccl := mock.New()
	input := `config =
  database =
    host = localhost
    port = 5432
  cache =
    redis = true`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: "\n  database =\n    host = localhost\n    port = 5432\n  cache =\n    redis = true"}}
	assert.Equal(t, expectedParse, parseResult)

}

// monoid_right_identity_nested - function:parse function:combine function:build_hierarchy (level 4)
func TestMonoidRightIdentityNested(t *testing.T) {

	ccl := mock.New()
	input := `config =
  database =
    host = localhost
    port = 5432
  cache =
    redis = true`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: "\n  database =\n    host = localhost\n    port = 5432\n  cache =\n    redis = true"}}
	assert.Equal(t, expectedParse, parseResult)

}

// monoid_left_identity_lists - function:parse feature:empty_keys function:combine (level 4)
func TestMonoidLeftIdentityLists(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
= item3`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "", Value: "item3"}}
	assert.Equal(t, expectedParse, parseResult)

}

// monoid_right_identity_lists - function:parse feature:empty_keys function:combine (level 4)
func TestMonoidRightIdentityLists(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
= item3`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "", Value: "item3"}}
	assert.Equal(t, expectedParse, parseResult)

}

// round_trip_property_basic - function:parse function:combine function:pretty_print (level 4)
func TestRoundTripPropertyBasic(t *testing.T) {

	ccl := mock.New()
	input := `key = value
another = test`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "another", Value: "test"}}
	assert.Equal(t, expectedParse, parseResult)

}

// round_trip_property_nested - function:parse function:combine function:build_hierarchy function:pretty_print (level 4)
func TestRoundTripPropertyNested(t *testing.T) {

	ccl := mock.New()
	input := `config =
  host = localhost
  port = 8080
  db =
    name = mydb
    user = admin`

	// Declare variables for reuse across validations
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: "\n  host = localhost\n  port = 8080\n  db =\n    name = mydb\n    user = admin"}}
	assert.Equal(t, expectedParse, parseResult)

}

// round_trip_property_complex - function:parse feature:empty_keys function:combine function:build_hierarchy function:pretty_print (level 4)
func TestRoundTripPropertyComplex(t *testing.T) {

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
	var parseResult []mock.Entry

	var err error

	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "config", Value: "\n  nested =\n    deep = value\n  list =\n    = a\n    = b\n    = c"}, mock.Entry{Key: "final", Value: "end"}}
	assert.Equal(t, expectedParse, parseResult)

}
