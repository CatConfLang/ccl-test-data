package parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_whitespace_behaviors.json
// Suite: Flat Format
// Version: 1.0



// spacing_strict_standard_format_parse - function:parse feature:whitespace behavior:strict_spacing behavior:loose_spacing
func TestSpacingStrictStandardFormatParse(t *testing.T) {
	

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


// spacing_loose_no_spaces_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingLooseNoSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key=value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_loose_left_space_only_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingLooseLeftSpaceOnlyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key =value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_loose_right_space_only_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingLooseRightSpaceOnlyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key= value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_loose_multiple_spaces_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingLooseMultipleSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key  =  value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_loose_tabs_around_equals_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingLooseTabsAroundEqualsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key	=	value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_loose_mixed_whitespace_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingLooseMixedWhitespaceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key 	 = 	 value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_loose_multiline_various_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingLooseMultilineVariousParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key1=val1
key2 = val2
key3  =  val3
key4	=	val4`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "val1"}, mock.Entry{Key: "key2", Value: "val2"}, mock.Entry{Key: "key3", Value: "val3"}, mock.Entry{Key: "key4", Value: "val4"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_loose_multiline_various_build_hierarchy - function:build_hierarchy feature:whitespace
func TestSpacingLooseMultilineVariousBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `key1=val1
key2 = val2
key3  =  val3
key4	=	val4`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// spacing_canonical_format_normalizes_loose_parse - function:parse feature:whitespace behavior:loose_spacing
func TestSpacingCanonicalFormatNormalizesLooseParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key1=val1
key2  =  val2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "val1"}, mock.Entry{Key: "key2", Value: "val2"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_canonical_format_normalizes_loose_canonical_format - function:canonical_format feature:whitespace
func TestSpacingCanonicalFormatNormalizesLooseCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `key1=val1
key2  =  val2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// tabs_preserve_in_value_parse - function:parse feature:whitespace behavior:tabs_preserve
func TestTabsPreserveInValueParse(t *testing.T) {
	

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


// tabs_preserve_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_preserve
func TestTabsPreserveInValueBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// tabs_preserve_in_value_get_string - function:get_string feature:whitespace
func TestTabsPreserveInValueGetString(t *testing.T) {
	

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


// tabs_preserve_leading_tab_parse - function:parse feature:whitespace behavior:tabs_preserve
func TestTabsPreserveLeadingTabParse(t *testing.T) {
	

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


// tabs_preserve_leading_tab_get_string - function:get_string feature:whitespace
func TestTabsPreserveLeadingTabGetString(t *testing.T) {
	

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


// tabs_to_spaces_in_value_parse - function:parse feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesInValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "  value  with  tabs"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_to_spaces_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesInValueBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// tabs_to_spaces_in_value_get_string - function:get_string feature:whitespace
func TestTabsToSpacesInValueGetString(t *testing.T) {
	

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
	assert.Equal(t, "  value  with  tabs", result)

}


// tabs_to_spaces_leading_tab_parse - function:parse feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesLeadingTabParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	indented`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "  indented"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_to_spaces_leading_tab_get_string - function:get_string feature:whitespace
func TestTabsToSpacesLeadingTabGetString(t *testing.T) {
	

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
	assert.Equal(t, "  indented", result)

}


// tabs_to_spaces_multiple_tabs_parse - function:parse feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesMultipleTabsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 			three_tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "      three_tabs"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_preserve_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_preserve
func TestTabsPreserveMultilineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `section =
		indented_with_tabs
		another_line`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "section", Value: "\n\t\tindented_with_tabs\n\t\tanother_line"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_to_spaces_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_to_spaces
func TestTabsToSpacesMultilineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `section =
		indented_with_tabs
		another_line`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "section", Value: "\n    indented_with_tabs\n    another_line"}}
	assert.Equal(t, expected, parseResult)

}


// tabs_canonical_format_preserve_canonical_format - function:canonical_format feature:whitespace behavior:tabs_preserve
func TestTabsCanonicalFormatPreserveCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// tabs_canonical_format_to_spaces_canonical_format - function:canonical_format feature:whitespace behavior:tabs_to_spaces
func TestTabsCanonicalFormatToSpacesCanonicalFormat(t *testing.T) {
	

	ccl := mock.New()
	input := `key = 	value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement canonical_format validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


// spacing_and_tabs_combined_loose_preserve_parse - function:parse feature:whitespace behavior:loose_spacing behavior:tabs_preserve
func TestSpacingAndTabsCombinedLoosePreserveParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key	=		value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\tvalue\twith\ttabs"}}
	assert.Equal(t, expected, parseResult)

}


// spacing_and_tabs_combined_loose_to_spaces_parse - function:parse feature:whitespace behavior:loose_spacing behavior:tabs_to_spaces
func TestSpacingAndTabsCombinedLooseToSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key	=		value	with	tabs`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "  value  with  tabs"}}
	assert.Equal(t, expected, parseResult)

}


