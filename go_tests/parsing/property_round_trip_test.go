package parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/property_round_trip.json
// Suite: Flat Format
// Version: 1.0



// round_trip_basic_parse - function:parse
func TestRoundTripBasicParse(t *testing.T) {
	

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


// round_trip_basic_round_trip - function:round_trip
func TestRoundTripBasicRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `key = value
nested =
  sub = val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_whitespace_normalization_parse - function:parse feature:whitespace
func TestRoundTripWhitespaceNormalizationParse(t *testing.T) {
	

	ccl := mock.New()
	input := `  key  =  value  
  nested  = 
    sub  =  val  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value  \n  nested  = \n    sub  =  val"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_whitespace_normalization_round_trip - function:round_trip feature:whitespace
func TestRoundTripWhitespaceNormalizationRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `  key  =  value  
  nested  = 
    sub  =  val  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_empty_keys_lists_parse - function:parse feature:empty_keys
func TestRoundTripEmptyKeysListsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `= item1
= item2
regular = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "regular", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_empty_keys_lists_round_trip - function:round_trip feature:empty_keys
func TestRoundTripEmptyKeysListsRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `= item1
= item2
regular = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_nested_structures_parse - function:parse
func TestRoundTripNestedStructuresParse(t *testing.T) {
	

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


// round_trip_nested_structures_round_trip - function:round_trip
func TestRoundTripNestedStructuresRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  host = localhost
  port = 8080
  db =
    name = mydb
    user = admin`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_multiline_values_parse - function:parse feature:multiline
func TestRoundTripMultilineValuesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `script =
  #!/bin/bash
  echo hello
  exit 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "script", Value: "\n  #!/bin/bash\n  echo hello\n  exit 0"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_multiline_values_round_trip - function:round_trip feature:multiline
func TestRoundTripMultilineValuesRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `script =
  #!/bin/bash
  echo hello
  exit 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_mixed_content_parse - function:parse feature:empty_keys
func TestRoundTripMixedContentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice
= first item
config =
  port = 3000
= second item
final = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "", Value: "first item"}, mock.Entry{Key: "config", Value: "\n  port = 3000"}, mock.Entry{Key: "", Value: "second item"}, mock.Entry{Key: "final", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_mixed_content_round_trip - function:round_trip feature:empty_keys
func TestRoundTripMixedContentRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice
= first item
config =
  port = 3000
= second item
final = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_complex_nesting_parse - function:parse feature:empty_keys
func TestRoundTripComplexNestingParse(t *testing.T) {
	

	ccl := mock.New()
	input := `app =
  = item1
  config =
    = nested_item
    db =
      host = localhost
      = db_item
  = item2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "app", Value: "\n  = item1\n  config =\n    = nested_item\n    db =\n      host = localhost\n      = db_item\n  = item2"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_complex_nesting_round_trip - function:round_trip feature:empty_keys
func TestRoundTripComplexNestingRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `app =
  = item1
  config =
    = nested_item
    db =
      host = localhost
      = db_item
  = item2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_deeply_nested_parse - function:parse feature:empty_keys
func TestRoundTripDeeplyNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `level1 =
  level2 =
    level3 =
      level4 =
        deep = value
        = deep_item`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "level1", Value: "\n  level2 =\n    level3 =\n      level4 =\n        deep = value\n        = deep_item"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_deeply_nested_round_trip - function:round_trip feature:empty_keys
func TestRoundTripDeeplyNestedRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `level1 =
  level2 =
    level3 =
      level4 =
        deep = value
        = deep_item`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// round_trip_empty_multiline_parse - function:parse feature:empty_keys feature:multiline
func TestRoundTripEmptyMultilineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `empty_section =

other = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_section", Value: ""}, mock.Entry{Key: "other", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// round_trip_empty_multiline_round_trip - function:round_trip feature:empty_keys feature:multiline
func TestRoundTripEmptyMultilineRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `empty_section =

other = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


