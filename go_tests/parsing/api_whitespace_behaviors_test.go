package parsing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tylerbutler/ccl-test-data/internal/mock"
)

// Generated from generated_tests/api_whitespace_behaviors.json
// Suite: Flat Format
// Version: 1.0

// tabs_as_content_in_value_parse - function:parse feature:whitespace behavior:tabs_as_content
func TestTabsAsContentInValueParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_content")
}

// tabs_as_content_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_as_content
func TestTabsAsContentInValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_content_in_value_get_string - function:get_string feature:whitespace behavior:tabs_as_content
func TestTabsAsContentInValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_content_leading_tab_parse - function:parse feature:whitespace behavior:tabs_as_content
func TestTabsAsContentLeadingTabParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_content")
}

// tabs_as_content_leading_tab_get_string - function:get_string feature:whitespace behavior:tabs_as_content
func TestTabsAsContentLeadingTabGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_whitespace_in_value_parse - function:parse feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceInValueParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_whitespace")
}

// tabs_as_whitespace_in_value_build_hierarchy - function:build_hierarchy feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceInValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_whitespace_in_value_get_string - function:get_string feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceInValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_whitespace_leading_tab_parse - function:parse feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceLeadingTabParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_whitespace")
}

// tabs_as_whitespace_leading_tab_get_string - function:get_string feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceLeadingTabGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_whitespace_multiple_tabs_parse - function:parse feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceMultipleTabsParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_whitespace")
}

// tabs_as_content_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_as_content
func TestTabsAsContentMultilineParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_content")
}

// tabs_as_whitespace_multiline_parse - function:parse feature:whitespace feature:multiline behavior:tabs_as_whitespace
func TestTabsAsWhitespaceMultilineParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_whitespace")
}

// tabs_as_whitespace_mixed_indent_parse - function:parse feature:whitespace feature:multiline behavior:tabs_as_whitespace
func TestTabsAsWhitespaceMixedIndentParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_whitespace")
}

// tabs_canonical_format_as_content_canonical_format - function:canonical_format feature:whitespace behavior:tabs_as_content
func TestTabsCanonicalFormatAsContentCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_canonical_format_as_whitespace_canonical_format - function:canonical_format feature:whitespace behavior:tabs_as_whitespace
func TestTabsCanonicalFormatAsWhitespaceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_whitespace_multiline_print_canonical_format - function:canonical_format feature:whitespace feature:multiline behavior:tabs_as_whitespace behavior:indent_spaces
func TestTabsAsWhitespaceMultilinePrintCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// tabs_as_whitespace_round_trip_round_trip - function:round_trip feature:whitespace behavior:tabs_as_whitespace
func TestTabsAsWhitespaceRoundTripRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_bare_list_indentation_canonical_format - function:canonical_format feature:empty_keys feature:whitespace behavior:indent_spaces
func TestNestedBareListIndentationCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_bare_list_indentation_canonical_format - function:canonical_format feature:empty_keys feature:whitespace behavior:indent_spaces
func TestDeeplyNestedBareListIndentationCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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

// crlf_normalize_to_lf_basic_build_hierarchy - function:build_hierarchy feature:whitespace
func TestCrlfNormalizeToLfBasicBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// crlf_preserve_literal_basic_parse - function:parse feature:whitespace behavior:crlf_preserve_literal
func TestCrlfPreserveLiteralBasicParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:crlf_preserve_literal")
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
	t.Skip("Test skipped due to tag filter: behavior:crlf_preserve_literal")
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

// crlf_nested_structure_build_hierarchy - function:build_hierarchy feature:whitespace
func TestCrlfNestedStructureBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// behavior_combo_tabs_and_crlf_parse - function:parse feature:whitespace behavior:tabs_as_whitespace behavior:crlf_normalize_to_lf
func TestBehaviorComboTabsAndCrlfParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_whitespace")
}

// behavior_combo_content_tabs_crlf_parse - function:parse feature:whitespace behavior:tabs_as_content behavior:crlf_normalize_to_lf
func TestBehaviorComboContentTabsCrlfParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_as_content")
}
