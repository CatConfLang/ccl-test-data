package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_proposed_behavior.json
// Suite: Flat Format
// Version: 1.0



// multiline_section_header_value_parse_indented - function:parse_indented feature:empty_keys feature:multiline
func TestMultilineSectionHeaderValueParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// unindented_multiline_becomes_continuation_parse_indented - function:parse_indented feature:empty_keys
func TestUnindentedMultilineBecomesContinuationParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// indented_line_is_continuation_parse_indented - function:parse_indented feature:multiline
func TestIndentedLineIsContinuationParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// indented_line_is_continuation_build_hierarchy - function:build_hierarchy feature:multiline behavior:array_order_insertion
func TestIndentedLineIsContinuationBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// indented_line_is_continuation_get_list - function:get_list feature:multiline behavior:list_coercion_enabled behavior:array_order_insertion
func TestIndentedLineIsContinuationGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_indentation_levels_parse_indented - function:parse_indented feature:multiline feature:empty_keys
func TestMixedIndentationLevelsParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_indentation_levels_build_hierarchy - function:build_hierarchy feature:multiline feature:empty_keys
func TestMixedIndentationLevelsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// single_item_as_list_parse - function:parse
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


// single_item_as_list_build_hierarchy - function:build_hierarchy
func TestSingleItemAsListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// single_item_as_list_get_list - function:get_list behavior:list_coercion_enabled
func TestSingleItemAsListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_duplicate_single_keys_parse - function:parse
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


// mixed_duplicate_single_keys_build_hierarchy - function:build_hierarchy behavior:array_order_insertion
func TestMixedDuplicateSingleKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_duplicate_single_keys_get_list - function:get_list behavior:list_coercion_enabled behavior:array_order_insertion
func TestMixedDuplicateSingleKeysGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// nested_list_access_parse - function:parse feature:path_traversal
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


// nested_list_access_build_hierarchy - function:build_hierarchy feature:path_traversal
func TestNestedListAccessBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// nested_list_access_get_list - function:get_list feature:path_traversal behavior:list_coercion_enabled
func TestNestedListAccessGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// empty_list_parse - function:parse
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


// empty_list_build_hierarchy - function:build_hierarchy behavior:array_order_insertion
func TestEmptyListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// empty_list_get_list - function:get_list behavior:list_coercion_enabled behavior:array_order_insertion
func TestEmptyListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_numbers_parse - function:parse
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


// list_with_numbers_build_hierarchy - function:build_hierarchy behavior:array_order_insertion
func TestListWithNumbersBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_numbers_get_list - function:get_list behavior:list_coercion_enabled behavior:array_order_insertion
func TestListWithNumbersGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_booleans_parse - function:parse
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


// list_with_booleans_build_hierarchy - function:build_hierarchy behavior:array_order_insertion
func TestListWithBooleansBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_booleans_get_list - function:get_list behavior:list_coercion_enabled behavior:array_order_insertion
func TestListWithBooleansGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_whitespace_parse - function:parse feature:whitespace
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


// list_with_whitespace_build_hierarchy - function:build_hierarchy feature:whitespace behavior:array_order_insertion
func TestListWithWhitespaceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_whitespace_get_list - function:get_list feature:whitespace behavior:list_coercion_enabled behavior:array_order_insertion
func TestListWithWhitespaceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_unicode_parse - function:parse feature:unicode
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


// list_with_unicode_build_hierarchy - function:build_hierarchy feature:unicode behavior:array_order_insertion
func TestListWithUnicodeBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_unicode_get_list - function:get_list feature:unicode behavior:list_coercion_enabled behavior:array_order_insertion
func TestListWithUnicodeGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_special_characters_parse - function:parse
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


// list_with_special_characters_build_hierarchy - function:build_hierarchy behavior:array_order_insertion
func TestListWithSpecialCharactersBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_special_characters_get_list - function:get_list behavior:list_coercion_enabled behavior:array_order_insertion
func TestListWithSpecialCharactersGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_multiline_values_parse_indented - function:parse_indented feature:multiline
func TestListMultilineValuesParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_multiline_values_build_hierarchy - function:build_hierarchy feature:multiline behavior:array_order_insertion
func TestListMultilineValuesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_multiline_values_get_list - function:get_list feature:multiline behavior:list_coercion_enabled behavior:array_order_insertion
func TestListMultilineValuesGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// complex_mixed_list_scenarios_parse_indented - function:parse_indented feature:path_traversal
func TestComplexMixedListScenariosParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// complex_mixed_list_scenarios_build_hierarchy - function:build_hierarchy feature:path_traversal behavior:array_order_insertion
func TestComplexMixedListScenariosBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// complex_mixed_list_scenarios_get_list - function:get_list feature:path_traversal behavior:list_coercion_enabled behavior:array_order_insertion
func TestComplexMixedListScenariosGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_path_traversal_protection_parse - function:parse
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


// list_path_traversal_protection_build_hierarchy - function:build_hierarchy
func TestListPathTraversalProtectionBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_path_traversal_protection_get_list - function:get_list behavior:list_coercion_enabled
func TestListPathTraversalProtectionGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_empty_value_parse - function:parse
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


// parse_empty_value_build_hierarchy - function:build_hierarchy
func TestParseEmptyValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parse_empty_value_get_string - function:get_string
func TestParseEmptyValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


