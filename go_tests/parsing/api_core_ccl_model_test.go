package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_core_ccl_model.json
// Suite: Flat Format
// Version: 1.0



// model_single_pair_parse - function:parse
func TestModelSinglePairParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// model_single_pair_build_model - function:parse_indented function:build_model
func TestModelSinglePairBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// model_duplicate_keys_merge_to_multi_leaf_parse - function:parse
func TestModelDuplicateKeysMergeToMultiLeafParse(t *testing.T) {
	

	ccl := mock.New()
	input := `item = first
item = second`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item", Value: "first"}, mock.Entry{Key: "item", Value: "second"}}
	assert.Equal(t, expected, parseResult)

}


// model_duplicate_keys_merge_to_multi_leaf_build_model - function:parse_indented function:build_model
func TestModelDuplicateKeysMergeToMultiLeafBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// model_repeated_value_is_idempotent_parse - function:parse
func TestModelRepeatedValueIsIdempotentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `item = a
item = b
item = a`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item", Value: "a"}, mock.Entry{Key: "item", Value: "b"}, mock.Entry{Key: "item", Value: "a"}}
	assert.Equal(t, expected, parseResult)

}


// model_repeated_value_is_idempotent_build_model - function:parse_indented function:build_model
func TestModelRepeatedValueIsIdempotentBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// model_multiple_distinct_keys_parse - function:parse
func TestModelMultipleDistinctKeysParse(t *testing.T) {
	

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


// model_multiple_distinct_keys_build_model - function:parse_indented function:build_model
func TestModelMultipleDistinctKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// model_empty_input_parse - function:parse
func TestModelEmptyInputParse(t *testing.T) {
	

	ccl := mock.New()
	input := ""
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


// model_empty_input_build_model - function:parse_indented function:build_model
func TestModelEmptyInputBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


