package parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_reference_compliant.json
// Suite: Flat Format
// Version: 1.0

// single_item_as_list_reference_parse - function:parse behavior:list_coercion_disabled variant:reference_compliant
func TestSingleItemAsListReferenceParse(t *testing.T) {

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

// single_item_as_list_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled variant:reference_compliant
func TestSingleItemAsListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_reference_get_list - function:get_list behavior:list_coercion_disabled variant:reference_compliant
func TestSingleItemAsListReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_reference_parse - function:parse behavior:list_coercion_disabled variant:reference_compliant
func TestMixedDuplicateSingleKeysReferenceParse(t *testing.T) {

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

// mixed_duplicate_single_keys_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled variant:reference_compliant
func TestMixedDuplicateSingleKeysReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_reference_get_list - function:get_list behavior:list_coercion_disabled variant:reference_compliant
func TestMixedDuplicateSingleKeysReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_reference_parse - function:parse behavior:list_coercion_disabled variant:reference_compliant
func TestNestedListAccessReferenceParse(t *testing.T) {

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

// nested_list_access_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled variant:reference_compliant
func TestNestedListAccessReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_reference_get_list - function:get_list behavior:list_coercion_disabled variant:reference_compliant
func TestNestedListAccessReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_reference_parse - function:parse variant:reference_compliant
func TestEmptyListReferenceParse(t *testing.T) {

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

// empty_list_reference_build_hierarchy - function:build_hierarchy variant:reference_compliant
func TestEmptyListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_reference_get_list - function:get_list variant:reference_compliant
func TestEmptyListReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_reference_parse - function:parse variant:reference_compliant
func TestListWithNumbersReferenceParse(t *testing.T) {

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

// list_with_numbers_reference_build_hierarchy - function:build_hierarchy variant:reference_compliant
func TestListWithNumbersReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_reference_get_list - function:get_list variant:reference_compliant
func TestListWithNumbersReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_reference_parse - function:parse variant:reference_compliant
func TestListWithBooleansReferenceParse(t *testing.T) {

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

// list_with_booleans_reference_build_hierarchy - function:build_hierarchy variant:reference_compliant
func TestListWithBooleansReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_reference_get_list - function:get_list variant:reference_compliant
func TestListWithBooleansReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_reference_parse - function:parse feature:whitespace variant:reference_compliant
func TestListWithWhitespaceReferenceParse(t *testing.T) {

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

// list_with_whitespace_reference_build_hierarchy - function:build_hierarchy feature:whitespace variant:reference_compliant
func TestListWithWhitespaceReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_reference_get_list - function:get_list feature:whitespace variant:reference_compliant
func TestListWithWhitespaceReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_reference_parse - function:parse variant:reference_compliant
func TestDeeplyNestedListReferenceParse(t *testing.T) {

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
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  environments =\n    production =\n      servers = web1\n      servers = web2\n      servers = api1"}}
	assert.Equal(t, expected, parseResult)

}

// deeply_nested_list_reference_build_hierarchy - function:build_hierarchy variant:reference_compliant
func TestDeeplyNestedListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_reference_get_list - function:get_list variant:reference_compliant
func TestDeeplyNestedListReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_reference_parse - function:parse feature:unicode variant:reference_compliant
func TestListWithUnicodeReferenceParse(t *testing.T) {

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

// list_with_unicode_reference_build_hierarchy - function:build_hierarchy feature:unicode variant:reference_compliant
func TestListWithUnicodeReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_reference_get_list - function:get_list feature:unicode variant:reference_compliant
func TestListWithUnicodeReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_reference_parse - function:parse variant:reference_compliant
func TestListWithSpecialCharactersReferenceParse(t *testing.T) {

	ccl := mock.New()
	input := `symbols = @#$%
symbols = !^&*()
symbols = []{}|`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "symbols", Value: "@#$%"}, mock.Entry{Key: "symbols", Value: "!^&*()"}, mock.Entry{Key: "symbols", Value: "[]{}|"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_special_characters_reference_build_hierarchy - function:build_hierarchy variant:reference_compliant
func TestListWithSpecialCharactersReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_reference_get_list - function:get_list variant:reference_compliant
func TestListWithSpecialCharactersReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled variant:reference_compliant
func TestComplexMixedListScenariosReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_reference_get_list - function:get_list behavior:list_coercion_disabled variant:reference_compliant
func TestComplexMixedListScenariosReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_reference_parse - function:parse behavior:list_coercion_disabled variant:reference_compliant
func TestListPathTraversalProtectionReferenceParse(t *testing.T) {

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

// list_path_traversal_protection_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled variant:reference_compliant
func TestListPathTraversalProtectionReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_reference_get_list - function:get_list behavior:list_coercion_disabled variant:reference_compliant
func TestListPathTraversalProtectionReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_value_reference_behavior_parse - function:parse variant:reference_compliant
func TestEmptyValueReferenceBehaviorParse(t *testing.T) {

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

// empty_value_reference_behavior_build_hierarchy - function:build_hierarchy variant:reference_compliant
func TestEmptyValueReferenceBehaviorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_empty_values_ocaml_reference_canonical_format - function:canonical_format variant:reference_compliant
func TestCanonicalFormatEmptyValuesOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_tab_preservation_ocaml_reference_canonical_format - function:canonical_format behavior:tabs_preserve variant:reference_compliant
func TestCanonicalFormatTabPreservationOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_unicode_ocaml_reference_canonical_format - function:canonical_format feature:unicode variant:reference_compliant
func TestCanonicalFormatUnicodeOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_line_endings_reference_behavior_parse - function:parse behavior:crlf_preserve_literal variant:reference_compliant
func TestCanonicalFormatLineEndingsReferenceBehaviorParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:crlf_preserve_literal")
}

// canonical_format_line_endings_reference_behavior_canonical_format - function:canonical_format behavior:crlf_preserve_literal variant:reference_compliant
func TestCanonicalFormatLineEndingsReferenceBehaviorCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_consistent_spacing_ocaml_reference_canonical_format - function:canonical_format behavior:strict_spacing variant:reference_compliant
func TestCanonicalFormatConsistentSpacingOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deterministic_output_ocaml_reference_canonical_format - function:canonical_format variant:reference_compliant
func TestDeterministicOutputOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
