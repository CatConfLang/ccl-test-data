package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_advanced_processing.json
// Suite: Flat Format
// Version: 1.0



// composition_stability_duplicate_keys_parse - function:parse
func TestCompositionStabilityDuplicateKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `a = 1
b = 2
b = 20
c = 3`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "1"}, mock.Entry{Key: "b", Value: "2"}, mock.Entry{Key: "b", Value: "20"}, mock.Entry{Key: "c", Value: "3"}}
	assert.Equal(t, expected, parseResult)

}


// multiple_values_same_key_parse - function:parse
func TestMultipleValuesSameKeyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `ports = 8000
ports = 8001
ports = 8002`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "ports", Value: "8000"}, mock.Entry{Key: "ports", Value: "8001"}, mock.Entry{Key: "ports", Value: "8002"}}
	assert.Equal(t, expected, parseResult)

}


// list_with_empty_keys_parse - function:parse feature:empty_keys
func TestListWithEmptyKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `= 3
= 1
= 2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "3"}, mock.Entry{Key: "", Value: "1"}, mock.Entry{Key: "", Value: "2"}}
	assert.Equal(t, expected, parseResult)

}


// section_style_syntax_parse - function:parse feature:empty_keys
func TestSectionStyleSyntaxParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== Section 2 ==`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Section 2 =="}}
	assert.Equal(t, expected, parseResult)

}


// composition_stability_ba_parse - function:parse
func TestCompositionStabilityBaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `b = 20
c = 3
a = 1
b = 2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "b", Value: "20"}, mock.Entry{Key: "c", Value: "3"}, mock.Entry{Key: "a", Value: "1"}, mock.Entry{Key: "b", Value: "2"}}
	assert.Equal(t, expected, parseResult)

}


// mixed_keys_with_duplicates_parse - function:parse feature:empty_keys
func TestMixedKeysWithDuplicatesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = app
ports = 8000
name = service
ports = 8001`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "app"}, mock.Entry{Key: "ports", Value: "8000"}, mock.Entry{Key: "name", Value: "service"}, mock.Entry{Key: "ports", Value: "8001"}}
	assert.Equal(t, expected, parseResult)

}


// array_style_list_parse - function:parse feature:empty_keys
func TestArrayStyleListParse(t *testing.T) {
	

	ccl := mock.New()
	input := `1 =
2 =
3 =`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "1", Value: ""}, mock.Entry{Key: "2", Value: ""}, mock.Entry{Key: "3", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// section_header_double_equals_parse - function:parse feature:empty_keys
func TestSectionHeaderDoubleEqualsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== Database Config ==
host = localhost
port = 5432`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database Config =="}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "5432"}}
	assert.Equal(t, expected, parseResult)

}


// section_header_triple_equals_parse - function:parse feature:empty_keys
func TestSectionHeaderTripleEqualsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `=== Server Settings ===
host = 0.0.0.0
ssl = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "== Server Settings ==="}, mock.Entry{Key: "host", Value: "0.0.0.0"}, mock.Entry{Key: "ssl", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}


// multiple_sections_with_entries_parse - function:parse feature:empty_keys
func TestMultipleSectionsWithEntriesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== Database ==
host = localhost

=== Cache ===
redis = enabled

== Logging ==
level = info`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database =="}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "", Value: "== Cache ==="}, mock.Entry{Key: "redis", Value: "enabled"}, mock.Entry{Key: "", Value: "= Logging =="}, mock.Entry{Key: "level", Value: "info"}}
	assert.Equal(t, expected, parseResult)

}


// section_headers_mixed_with_lists_parse - function:parse feature:empty_keys
func TestSectionHeadersMixedWithListsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== Configuration ==
= item1
= item2
key = value
=== Next Section ===
other = data`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Configuration =="}, mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "", Value: "== Next Section ==="}, mock.Entry{Key: "other", Value: "data"}}
	assert.Equal(t, expected, parseResult)

}


// empty_section_header_only_parse - function:parse
func TestEmptySectionHeaderOnlyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== Empty Section ==`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Empty Section =="}}
	assert.Equal(t, expected, parseResult)

}


// section_header_at_end_parse - function:parse feature:empty_keys
func TestSectionHeaderAtEndParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = value
== Final Section ==`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}, mock.Entry{Key: "", Value: "= Final Section =="}}
	assert.Equal(t, expected, parseResult)

}


// section_headers_no_trailing_equals_parse - function:parse feature:empty_keys
func TestSectionHeadersNoTrailingEqualsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== Database Config
host = localhost
=== Server Settings
port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database Config"}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "", Value: "== Server Settings"}, mock.Entry{Key: "port", Value: "8080"}}
	assert.Equal(t, expected, parseResult)

}


// section_headers_with_colons_parse - function:parse feature:empty_keys
func TestSectionHeadersWithColonsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== Database: Production ==
host = db.prod.com
=== Cache: Redis Config ===
port = 6379`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database: Production =="}, mock.Entry{Key: "host", Value: "db.prod.com"}, mock.Entry{Key: "", Value: "== Cache: Redis Config ==="}, mock.Entry{Key: "port", Value: "6379"}}
	assert.Equal(t, expected, parseResult)

}


// spaced_equals_not_section_header_parse - function:parse feature:empty_keys
func TestSpacedEqualsNotSectionHeaderParse(t *testing.T) {
	

	ccl := mock.New()
	input := `= = spaced equals
=  = wide spaces
== Real Header ==
key = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= spaced equals"}, mock.Entry{Key: "", Value: "= wide spaces"}, mock.Entry{Key: "", Value: "= Real Header =="}, mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// consecutive_section_headers_parse - function:parse feature:empty_keys
func TestConsecutiveSectionHeadersParse(t *testing.T) {
	

	ccl := mock.New()
	input := `== First Section ==
=== Nested Section ===
==== Deep Section ====
key = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= First Section =="}, mock.Entry{Key: "", Value: "== Nested Section ==="}, mock.Entry{Key: "", Value: "=== Deep Section ===="}, mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


