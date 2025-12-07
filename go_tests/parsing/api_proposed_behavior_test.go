package parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_proposed_behavior.json
// Suite: Flat Format
// Version: 1.0

// multiline_section_header_value_parse_indented - function:parse_indented feature:empty_keys feature:multiline variant:proposed_behavior
func TestMultilineSectionHeaderValueParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// unindented_multiline_becomes_continuation_parse_indented - function:parse_indented feature:empty_keys variant:proposed_behavior
func TestUnindentedMultilineBecomesContinuationParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_parse - function:parse variant:proposed_behavior
func TestSingleItemAsListParse(t *testing.T) {

	ccl := mock.New()
	input := `item = single`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item", Value: "single"}}
	assert.Equal(t, expected, parseResult)

}

// single_item_as_list_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestSingleItemAsListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestSingleItemAsListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_parse - function:parse variant:proposed_behavior
func TestMixedDuplicateSingleKeysParse(t *testing.T) {

	ccl := mock.New()
	input := `ports = 80
ports = 443
host = localhost`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "ports", Value: "80"}, mock.Entry{Key: "ports", Value: "443"}, mock.Entry{Key: "host", Value: "localhost"}}
	assert.Equal(t, expected, parseResult)

}

// mixed_duplicate_single_keys_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestMixedDuplicateSingleKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestMixedDuplicateSingleKeysGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_parse - function:parse variant:proposed_behavior
func TestNestedListAccessParse(t *testing.T) {

	ccl := mock.New()
	input := `database =
  hosts = primary
  hosts = secondary
  port = 5432`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  hosts = primary\n  hosts = secondary\n  port = 5432"}}
	assert.Equal(t, expected, parseResult)

}

// nested_list_access_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestNestedListAccessBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestNestedListAccessGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_parse - function:parse variant:proposed_behavior
func TestEmptyListParse(t *testing.T) {

	ccl := mock.New()
	input := `empty_list =`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_list", Value: ""}}
	assert.Equal(t, expected, parseResult)

}

// empty_list_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestEmptyListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestEmptyListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_parse - function:parse variant:proposed_behavior
func TestListWithNumbersParse(t *testing.T) {

	ccl := mock.New()
	input := `numbers = 1
numbers = 42
numbers = -17
numbers = 0`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "numbers", Value: "1"}, mock.Entry{Key: "numbers", Value: "42"}, mock.Entry{Key: "numbers", Value: "-17"}, mock.Entry{Key: "numbers", Value: "0"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_numbers_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestListWithNumbersBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestListWithNumbersGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_parse - function:parse variant:proposed_behavior
func TestListWithBooleansParse(t *testing.T) {

	ccl := mock.New()
	input := `flags = true
flags = false
flags = yes
flags = no`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "flags", Value: "true"}, mock.Entry{Key: "flags", Value: "false"}, mock.Entry{Key: "flags", Value: "yes"}, mock.Entry{Key: "flags", Value: "no"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_booleans_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestListWithBooleansBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestListWithBooleansGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_parse - function:parse feature:whitespace variant:proposed_behavior
func TestListWithWhitespaceParse(t *testing.T) {

	ccl := mock.New()
	input := `items =   spaced   
items = normal
items =
items =   `

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "spaced"}, mock.Entry{Key: "items", Value: "normal"}, mock.Entry{Key: "items", Value: ""}, mock.Entry{Key: "items", Value: ""}}
	assert.Equal(t, expected, parseResult)

}

// list_with_whitespace_build_hierarchy - function:build_hierarchy feature:whitespace variant:proposed_behavior
func TestListWithWhitespaceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_get_list - function:get_list feature:whitespace behavior:list_coercion_enabled variant:proposed_behavior
func TestListWithWhitespaceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_parse - function:parse variant:proposed_behavior
func TestDeeplyNestedListParse(t *testing.T) {

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers = web1
      servers = web2
      servers = api1`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: ""}, mock.Entry{Key: "environments", Value: ""}, mock.Entry{Key: "production", Value: ""}, mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "api1"}}
	assert.Equal(t, expected, parseResult)

}

// deeply_nested_list_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestDeeplyNestedListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestDeeplyNestedListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_parse - function:parse feature:unicode variant:proposed_behavior
func TestListWithUnicodeParse(t *testing.T) {

	ccl := mock.New()
	input := `names = 张三
names = José
names = François
names = العربية`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "names", Value: "张三"}, mock.Entry{Key: "names", Value: "José"}, mock.Entry{Key: "names", Value: "François"}, mock.Entry{Key: "names", Value: "العربية"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_unicode_build_hierarchy - function:build_hierarchy feature:unicode variant:proposed_behavior
func TestListWithUnicodeBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_get_list - function:get_list feature:unicode behavior:list_coercion_enabled variant:proposed_behavior
func TestListWithUnicodeGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_parse - function:parse variant:proposed_behavior
func TestListWithSpecialCharactersParse(t *testing.T) {

	ccl := mock.New()
	input := `symbols = @#$%
symbols = !^&*()
symbols = []{}|
symbols = <>=+`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "symbols", Value: "@#$%"}, mock.Entry{Key: "symbols", Value: "!^&*()"}, mock.Entry{Key: "symbols", Value: "[]{}|"}, mock.Entry{Key: "symbols", Value: "<>=+"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_special_characters_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestListWithSpecialCharactersBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestListWithSpecialCharactersGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_parse_indented - function:parse_indented feature:multiline variant:proposed_behavior
func TestListMultilineValuesParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_build_hierarchy - function:build_hierarchy feature:multiline variant:proposed_behavior
func TestListMultilineValuesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_get_list - function:get_list feature:multiline behavior:list_coercion_enabled variant:proposed_behavior
func TestListMultilineValuesGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_parse_indented - function:parse_indented variant:proposed_behavior
func TestComplexMixedListScenariosParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestComplexMixedListScenariosBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestComplexMixedListScenariosGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_parse - function:parse variant:proposed_behavior
func TestListPathTraversalProtectionParse(t *testing.T) {

	ccl := mock.New()
	input := `safe = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "safe", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// list_path_traversal_protection_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestListPathTraversalProtectionBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_get_list - function:get_list behavior:list_coercion_enabled variant:proposed_behavior
func TestListPathTraversalProtectionGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_empty_value_parse - function:parse variant:proposed_behavior
func TestParseEmptyValueParse(t *testing.T) {

	ccl := mock.New()
	input := `empty_key =`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}

// parse_empty_value_build_hierarchy - function:build_hierarchy variant:proposed_behavior
func TestParseEmptyValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_empty_value_get_string - function:get_string variant:proposed_behavior
func TestParseEmptyValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_empty_values_parse - function:parse variant:proposed_behavior
func TestCanonicalFormatEmptyValuesParse(t *testing.T) {

	ccl := mock.New()
	input := `empty_key =`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_empty_values_canonical_format - function:canonical_format variant:proposed_behavior
func TestCanonicalFormatEmptyValuesCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_tab_preservation_parse - function:parse behavior:tabs_preserve variant:proposed_behavior
func TestCanonicalFormatTabPreservationParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_preserve")
}

// canonical_format_tab_preservation_canonical_format - function:canonical_format behavior:tabs_preserve variant:proposed_behavior
func TestCanonicalFormatTabPreservationCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_unicode_parse - function:parse feature:unicode variant:proposed_behavior
func TestCanonicalFormatUnicodeParse(t *testing.T) {

	ccl := mock.New()
	input := `unicode = 你好世界
emo = 🌟✨`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "unicode", Value: "你好世界"}, mock.Entry{Key: "emo", Value: "🌟✨"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_unicode_canonical_format - function:canonical_format feature:unicode variant:proposed_behavior
func TestCanonicalFormatUnicodeCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_line_endings_proposed_parse - function:parse behavior:crlf_preserve_literal variant:proposed_behavior
func TestCanonicalFormatLineEndingsProposedParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:crlf_preserve_literal")
}

// canonical_format_line_endings_proposed_canonical_format - function:canonical_format behavior:crlf_preserve_literal variant:proposed_behavior
func TestCanonicalFormatLineEndingsProposedCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// crlf_normalize_to_lf_proposed_parse - function:parse behavior:crlf_normalize_to_lf variant:proposed_behavior
func TestCrlfNormalizeToLfProposedParse(t *testing.T) {

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

// crlf_normalize_to_lf_proposed_canonical_format - function:canonical_format behavior:crlf_normalize_to_lf variant:proposed_behavior
func TestCrlfNormalizeToLfProposedCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// crlf_normalize_to_lf_indented_proposed_parse - function:parse behavior:crlf_normalize_to_lf variant:proposed_behavior
func TestCrlfNormalizeToLfIndentedProposedParse(t *testing.T) {

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

// crlf_normalize_to_lf_indented_proposed_canonical_format - function:canonical_format behavior:crlf_normalize_to_lf variant:proposed_behavior
func TestCrlfNormalizeToLfIndentedProposedCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_consistent_spacing_parse - function:parse behavior:strict_spacing variant:proposed_behavior
func TestCanonicalFormatConsistentSpacingParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:strict_spacing")
}

// canonical_format_consistent_spacing_canonical_format - function:canonical_format variant:proposed_behavior
func TestCanonicalFormatConsistentSpacingCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deterministic_output_parse - function:parse variant:proposed_behavior
func TestDeterministicOutputParse(t *testing.T) {

	ccl := mock.New()
	input := `z = last
a = first
m = middle`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "z", Value: "last"}, mock.Entry{Key: "a", Value: "first"}, mock.Entry{Key: "m", Value: "middle"}}
	assert.Equal(t, expected, parseResult)

}

// deterministic_output_canonical_format - function:canonical_format variant:proposed_behavior
func TestDeterministicOutputCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
