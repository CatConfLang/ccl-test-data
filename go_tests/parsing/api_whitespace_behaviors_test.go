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
	t.Skip("Test skipped due to tag filter: behavior:strict_spacing")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_preserve_in_value_parse - function:parse feature:whitespace behavior:tabs_preserve
func TestTabsPreserveInValueParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_preserve")
}

// tabs_preserve_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_preserve
func TestTabsPreserveInValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_preserve_in_value_get_string - function:get_string feature:whitespace
func TestTabsPreserveInValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_preserve_leading_tab_parse - function:parse feature:whitespace behavior:tabs_preserve
func TestTabsPreserveLeadingTabParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_preserve")
}

// tabs_preserve_leading_tab_get_string - function:get_string feature:whitespace
func TestTabsPreserveLeadingTabGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_to_spaces_in_value_parse - function:parse feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesInValueParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_to_spaces")
}

// tabs_to_spaces_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesInValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_to_spaces_in_value_get_string - function:get_string feature:whitespace
func TestTabsToSpacesInValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_to_spaces_leading_tab_parse - function:parse feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesLeadingTabParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_to_spaces")
}

// tabs_to_spaces_leading_tab_get_string - function:get_string feature:whitespace
func TestTabsToSpacesLeadingTabGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_to_spaces_multiple_tabs_parse - function:parse feature:whitespace behavior:tabs_to_spaces
func TestTabsToSpacesMultipleTabsParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_to_spaces")
}

// tabs_preserve_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_preserve
func TestTabsPreserveMultilineParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_preserve")
}

// tabs_to_spaces_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_to_spaces
func TestTabsToSpacesMultilineParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_to_spaces")
}

// tabs_canonical_format_preserve_canonical_format - function:canonical_format feature:whitespace behavior:tabs_preserve
func TestTabsCanonicalFormatPreserveCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_canonical_format_to_spaces_canonical_format - function:canonical_format feature:whitespace behavior:tabs_to_spaces
func TestTabsCanonicalFormatToSpacesCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// spacing_and_tabs_combined_loose_preserve_parse - function:parse feature:whitespace behavior:loose_spacing behavior:tabs_preserve
func TestSpacingAndTabsCombinedLoosePreserveParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_preserve")
}

// spacing_and_tabs_combined_loose_to_spaces_parse - function:parse feature:whitespace behavior:loose_spacing behavior:tabs_to_spaces
func TestSpacingAndTabsCombinedLooseToSpacesParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_to_spaces")
}

// nested_bare_list_indentation_canonical_format - function:canonical_format feature:empty_keys feature:whitespace
func TestNestedBareListIndentationCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_bare_list_indentation_canonical_format - function:canonical_format feature:empty_keys feature:whitespace
func TestDeeplyNestedBareListIndentationCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
