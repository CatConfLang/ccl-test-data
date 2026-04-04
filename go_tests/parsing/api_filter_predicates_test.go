package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_filter_predicates.json
// Suite: Flat Format
// Version: 1.0



// filter_by_key_equality_parse - function:parse
func TestFilterByKeyEqualityParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
mode = debug`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "mode", Value: "debug"}}
	assert.Equal(t, expected, parseResult)

}


// filter_by_key_equality_filter - function:filter
func TestFilterByKeyEqualityFilter(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
mode = debug`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement filter validation with predicate: key == "port"
	// The Go mock only supports comment-exclusion (key != "/")
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// filter_by_value_not_empty_parse - function:parse feature:empty_keys
func TestFilterByValueNotEmptyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = app
description =
version = 1.0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "app"}, mock.Entry{Key: "description", Value: ""}, mock.Entry{Key: "version", Value: "1.0"}}
	assert.Equal(t, expected, parseResult)

}


// filter_by_value_not_empty_filter - function:filter feature:empty_keys
func TestFilterByValueNotEmptyFilter(t *testing.T) {
	

	ccl := mock.New()
	input := `name = app
description =
version = 1.0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement filter validation with predicate: value != ""
	// The Go mock only supports comment-exclusion (key != "/")
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// filter_by_value_equality_parse - function:parse
func TestFilterByValueEqualityParse(t *testing.T) {
	

	ccl := mock.New()
	input := `debug = true
verbose = false
logging = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "debug", Value: "true"}, mock.Entry{Key: "verbose", Value: "false"}, mock.Entry{Key: "logging", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}


// filter_by_value_equality_filter - function:filter
func TestFilterByValueEqualityFilter(t *testing.T) {
	

	ccl := mock.New()
	input := `debug = true
verbose = false
logging = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement filter validation with predicate: value == "true"
	// The Go mock only supports comment-exclusion (key != "/")
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// filter_no_matches_parse - function:parse
func TestFilterNoMatchesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}}
	assert.Equal(t, expected, parseResult)

}


// filter_no_matches_filter - function:filter
func TestFilterNoMatchesFilter(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement filter validation with predicate: key == "nonexistent"
	// The Go mock only supports comment-exclusion (key != "/")
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// filter_keeps_all_parse - function:parse
func TestFilterKeepsAllParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}}
	assert.Equal(t, expected, parseResult)

}


// filter_keeps_all_filter - function:filter
func TestFilterKeepsAllFilter(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement filter validation with predicate: key != "nonexistent"
	// The Go mock only supports comment-exclusion (key != "/")
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


