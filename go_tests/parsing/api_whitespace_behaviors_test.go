package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_whitespace_behaviors.json
// Suite: Flat Format
// Version: 1.0



// tabs_as_content_in_value_parse - function:parse feature:whitespace behavior:tabs_as_content
func TestTabsAsContentInValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\tvalue\twith\ttabs"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_as_content_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_as_content
func TestTabsAsContentInValueBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"key": "\tvalue\twith\ttabs"}
	assert.Equal(t, expected, objectResult)

}


// tabs_as_content_in_value_get_string - function:get_string feature:whitespace behavior:tabs_as_content
func TestTabsAsContentInValueGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"key"})
	require.NoError(t, err)
	assert.Equal(t, "\tvalue\twith\ttabs", result)

}


// tabs_as_content_leading_tab_parse - function:parse feature:whitespace behavior:tabs_as_content
func TestTabsAsContentLeadingTabParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	indented`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\tindented"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_as_content_leading_tab_get_string - function:get_string feature:whitespace behavior:tabs_as_content
func TestTabsAsContentLeadingTabGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	indented`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"key"})
	require.NoError(t, err)
	assert.Equal(t, "\tindented", result)

}


// tabs_as_whitespace_in_value_parse - function:parse feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceInValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value with tabs"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_as_whitespace_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceInValueBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"key": "value with tabs"}
	assert.Equal(t, expected, objectResult)

}


// tabs_as_whitespace_in_value_get_string - function:get_string feature:whitespace
func TestTabsAsWhitespaceInValueGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"key"})
	require.NoError(t, err)
	assert.Equal(t, "value with tabs", result)

}


// tabs_as_whitespace_leading_tab_parse - function:parse feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceLeadingTabParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	indented`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "indented"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_as_whitespace_leading_tab_get_string - function:get_string feature:whitespace
func TestTabsAsWhitespaceLeadingTabGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	indented`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"key"})
	require.NoError(t, err)
	assert.Equal(t, "indented", result)

}


// tabs_as_whitespace_multiple_tabs_parse - function:parse feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceMultipleTabsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 			three_tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "three_tabs"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_as_content_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_as_content
func TestTabsAsContentMultilineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `section =
 	indented_with_tabs
 	another_line`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "section", Value: "\n \tindented_with_tabs\n \tanother_line"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_as_whitespace_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_as_whitespace
func TestTabsAsWhitespaceMultilineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `section =
		indented_with_tabs
		another_line`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "section", Value: "\nindented_with_tabs\nanother_line"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_as_whitespace_mixed_indent_parse - function:parse feature:whitespace feature:multiline behavior:tabs_as_whitespace
func TestTabsAsWhitespaceMixedIndentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `section =
 	mixed_indent
	 another_line`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "section", Value: "\nmixed_indent\nanother_line"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_canonical_format_as_content_canonical_format - function:canonical_format feature:whitespace behavior:tabs_as_content
func TestTabsCanonicalFormatAsContentCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// tabs_canonical_format_as_whitespace_canonical_format - function:canonical_format feature:whitespace behavior:tabs_as_whitespace
func TestTabsCanonicalFormatAsWhitespaceCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// tabs_as_whitespace_multiline_print_canonical_format - function:canonical_format feature:whitespace feature:multiline behavior:tabs_as_whitespace behavior:indent_spaces
func TestTabsAsWhitespaceMultilinePrintCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `section =
		indented
		another`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// tabs_as_whitespace_round_trip_round_trip - function:round_trip feature:whitespace
func TestTabsAsWhitespaceRoundTripRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// nested_bare_list_indentation_canonical_format - function:canonical_format feature:empty_keys feature:whitespace behavior:indent_spaces
func TestNestedBareListIndentationCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `package =
  = brew
  = scoop
  = nix`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// deeply_nested_bare_list_indentation_canonical_format - function:canonical_format feature:empty_keys feature:whitespace behavior:indent_spaces
func TestDeeplyNestedBareListIndentationCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `app =
  = item1
  config =
    = nested1
    = nested2
    deep =
      = level3a
      = level3b
  = item2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// crlf_normalize_to_lf_basic_parse - function:parse feature:whitespace behavior:crlf_normalize_to_lf
func TestCrlfNormalizeToLfBasicParse(t *testing.T) {
	

	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}}
	assert.Equal(t, expected, parseResult)

}


// crlf_normalize_to_lf_basic_build_hierarchy - function:build_hierarchy feature:whitespace behavior:crlf_normalize_to_lf
func TestCrlfNormalizeToLfBasicBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"key1": "value1", "key2": "value2"}
	assert.Equal(t, expected, objectResult)

}


// crlf_preserve_literal_basic_parse - function:parse feature:whitespace behavior:crlf_preserve_literal
func TestCrlfPreserveLiteralBasicParse(t *testing.T) {
	

	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1\r"}, mock.Entry{Key: "key2", Value: "value2\r"}}
	assert.Equal(t, expected, parseResult)

}


// crlf_normalize_multiline_value_parse - function:parse feature:whitespace feature:multiline behavior:crlf_normalize_to_lf
func TestCrlfNormalizeMultilineValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := "multiline =\r\n  line1\r\n  line2"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "multiline", Value: "\n  line1\n  line2"}}
	assert.Equal(t, expected, parseResult)

}


// crlf_preserve_multiline_value_parse - function:parse feature:whitespace feature:multiline behavior:crlf_preserve_literal
func TestCrlfPreserveMultilineValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := "multiline =\r\n  line1\r\n  line2"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "multiline", Value: "\r\n  line1\r\n  line2"}}
	assert.Equal(t, expected, parseResult)

}


// crlf_mixed_line_endings_parse - function:parse feature:whitespace behavior:crlf_normalize_to_lf
func TestCrlfMixedLineEndingsParse(t *testing.T) {
	

	ccl := mock.New()
	input := "lf_line = value1\ncrlf_line = value2\r\nlf_again = value3\n"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "lf_line", Value: "value1"}, mock.Entry{Key: "crlf_line", Value: "value2"}, mock.Entry{Key: "lf_again", Value: "value3"}}
	assert.Equal(t, expected, parseResult)

}


// crlf_nested_structure_parse - function:parse feature:whitespace behavior:crlf_normalize_to_lf
func TestCrlfNestedStructureParse(t *testing.T) {
	

	ccl := mock.New()
	input := "config =\r\n  host = localhost\r\n  port = 8080"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  host = localhost\n  port = 8080"}}
	assert.Equal(t, expected, parseResult)

}


// crlf_nested_structure_build_hierarchy - function:build_hierarchy feature:whitespace behavior:crlf_normalize_to_lf
func TestCrlfNestedStructureBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := "config =\r\n  host = localhost\r\n  port = 8080"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"host": "localhost", "port": "8080"}}
	assert.Equal(t, expected, objectResult)

}


// behavior_combo_tabs_and_crlf_parse - function:parse feature:whitespace behavior:tabs_as_whitespace behavior:crlf_normalize_to_lf
func TestBehaviorComboTabsAndCrlfParse(t *testing.T) {
	

	ccl := mock.New()
	input := "key = \tvalue\twith\ttabs\r\n"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value with tabs"}}
	assert.Equal(t, expected, parseResult)

}


// behavior_combo_content_tabs_crlf_parse - function:parse feature:whitespace behavior:tabs_as_content behavior:crlf_normalize_to_lf
func TestBehaviorComboContentTabsCrlfParse(t *testing.T) {
	

	ccl := mock.New()
	input := "key1 = \tvalue1\r\nkey2 = \tvalue2\r\n"
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "\tvalue1"}, mock.Entry{Key: "key2", Value: "\tvalue2"}}
	assert.Equal(t, expected, parseResult)

}


