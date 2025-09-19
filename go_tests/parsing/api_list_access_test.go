package parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_list_access.json
// Suite: Flat Format
// Version: 1.0

// basic_list_from_duplicates_parse - function:parse
func TestBasicListFromDuplicatesParse(t *testing.T) {

	ccl := mock.New()
	input := `servers = web1
servers = web2
servers = web3`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "web3"}}
	assert.Equal(t, expected, parseResult)

}

// basic_list_from_duplicates_build_hierarchy - function:build_hierarchy
func TestBasicListFromDuplicatesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// basic_list_from_duplicates_get_list - function:get_list
func TestBasicListFromDuplicatesGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_parse - function:parse behavior:list_coercion_enabled
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

// single_item_as_list_build_hierarchy - function:build_hierarchy behavior:list_coercion_enabled
func TestSingleItemAsListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_get_list - function:get_list behavior:list_coercion_enabled
func TestSingleItemAsListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_reference_parse - function:parse behavior:list_coercion_disabled
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

// single_item_as_list_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled
func TestSingleItemAsListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_reference_get_list - function:get_list behavior:list_coercion_disabled
func TestSingleItemAsListReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_parse - function:parse behavior:list_coercion_enabled
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

// mixed_duplicate_single_keys_build_hierarchy - function:build_hierarchy behavior:list_coercion_enabled
func TestMixedDuplicateSingleKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_get_list - function:get_list behavior:list_coercion_enabled
func TestMixedDuplicateSingleKeysGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_reference_parse - function:parse behavior:list_coercion_disabled
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

// mixed_duplicate_single_keys_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled
func TestMixedDuplicateSingleKeysReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_reference_get_list - function:get_list behavior:list_coercion_disabled
func TestMixedDuplicateSingleKeysReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_parse - function:parse behavior:list_coercion_enabled
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

// nested_list_access_build_hierarchy - function:build_hierarchy behavior:list_coercion_enabled
func TestNestedListAccessBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_get_list - function:get_list behavior:list_coercion_enabled
func TestNestedListAccessGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_reference_parse - function:parse behavior:list_coercion_disabled
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

// nested_list_access_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled
func TestNestedListAccessReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_reference_get_list - function:get_list behavior:list_coercion_disabled
func TestNestedListAccessReferenceGetList(t *testing.T) {
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

// empty_list_build_hierarchy - function:build_hierarchy
func TestEmptyListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_get_list - function:get_list
func TestEmptyListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_reference_parse - function:parse
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

// empty_list_reference_build_hierarchy - function:build_hierarchy
func TestEmptyListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_reference_get_list - function:get_list
func TestEmptyListReferenceGetList(t *testing.T) {
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

// list_with_numbers_build_hierarchy - function:build_hierarchy
func TestListWithNumbersBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_get_list - function:get_list
func TestListWithNumbersGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_reference_parse - function:parse
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

// list_with_numbers_reference_build_hierarchy - function:build_hierarchy
func TestListWithNumbersReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_reference_get_list - function:get_list
func TestListWithNumbersReferenceGetList(t *testing.T) {
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

// list_with_booleans_build_hierarchy - function:build_hierarchy
func TestListWithBooleansBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_get_list - function:get_list
func TestListWithBooleansGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_reference_parse - function:parse
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

// list_with_booleans_reference_build_hierarchy - function:build_hierarchy
func TestListWithBooleansReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_reference_get_list - function:get_list
func TestListWithBooleansReferenceGetList(t *testing.T) {
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

// list_with_whitespace_build_hierarchy - function:build_hierarchy feature:whitespace
func TestListWithWhitespaceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_get_list - function:get_list feature:whitespace
func TestListWithWhitespaceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_reference_parse - function:parse feature:whitespace
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

// list_with_whitespace_reference_build_hierarchy - function:build_hierarchy feature:whitespace
func TestListWithWhitespaceReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_reference_get_list - function:get_list feature:whitespace
func TestListWithWhitespaceReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_parse - function:parse
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

// deeply_nested_list_build_hierarchy - function:build_hierarchy
func TestDeeplyNestedListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_get_list - function:get_list
func TestDeeplyNestedListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_reference_parse - function:parse
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

// deeply_nested_list_reference_build_hierarchy - function:build_hierarchy
func TestDeeplyNestedListReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_reference_get_list - function:get_list
func TestDeeplyNestedListReferenceGetList(t *testing.T) {
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

// list_with_unicode_build_hierarchy - function:build_hierarchy feature:unicode
func TestListWithUnicodeBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_get_list - function:get_list feature:unicode
func TestListWithUnicodeGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_reference_parse - function:parse feature:unicode
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

// list_with_unicode_reference_build_hierarchy - function:build_hierarchy feature:unicode
func TestListWithUnicodeReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_reference_get_list - function:get_list feature:unicode
func TestListWithUnicodeReferenceGetList(t *testing.T) {
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

// list_with_special_characters_build_hierarchy - function:build_hierarchy
func TestListWithSpecialCharactersBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_get_list - function:get_list
func TestListWithSpecialCharactersGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_reference_parse - function:parse
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

// list_with_special_characters_reference_build_hierarchy - function:build_hierarchy
func TestListWithSpecialCharactersReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_reference_get_list - function:get_list
func TestListWithSpecialCharactersReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// large_list_parse - function:parse
func TestLargeListParse(t *testing.T) {

	ccl := mock.New()
	input := `items = item01
items = item02
items = item03
items = item04
items = item05
items = item06
items = item07
items = item08
items = item09
items = item10
items = item11
items = item12
items = item13
items = item14
items = item15
items = item16
items = item17
items = item18
items = item19
items = item20`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "item01"}, mock.Entry{Key: "items", Value: "item02"}, mock.Entry{Key: "items", Value: "item03"}, mock.Entry{Key: "items", Value: "item04"}, mock.Entry{Key: "items", Value: "item05"}, mock.Entry{Key: "items", Value: "item06"}, mock.Entry{Key: "items", Value: "item07"}, mock.Entry{Key: "items", Value: "item08"}, mock.Entry{Key: "items", Value: "item09"}, mock.Entry{Key: "items", Value: "item10"}, mock.Entry{Key: "items", Value: "item11"}, mock.Entry{Key: "items", Value: "item12"}, mock.Entry{Key: "items", Value: "item13"}, mock.Entry{Key: "items", Value: "item14"}, mock.Entry{Key: "items", Value: "item15"}, mock.Entry{Key: "items", Value: "item16"}, mock.Entry{Key: "items", Value: "item17"}, mock.Entry{Key: "items", Value: "item18"}, mock.Entry{Key: "items", Value: "item19"}, mock.Entry{Key: "items", Value: "item20"}}
	assert.Equal(t, expected, parseResult)

}

// large_list_build_hierarchy - function:build_hierarchy
func TestLargeListBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// large_list_get_list - function:get_list
func TestLargeListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_comments_parse - function:parse feature:comments
func TestListWithCommentsParse(t *testing.T) {

	ccl := mock.New()
	input := `servers = web1
/= Production servers
servers = web2
servers = web3
/= End of list`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "/", Value: "Production servers"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "web3"}, mock.Entry{Key: "/", Value: "End of list"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_comments_build_hierarchy - function:build_hierarchy feature:comments
func TestListWithCommentsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_comments_get_list - function:get_list feature:comments
func TestListWithCommentsGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_missing_key_parse - function:parse
func TestListErrorMissingKeyParse(t *testing.T) {

	ccl := mock.New()
	input := `existing = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "existing", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// list_error_missing_key_build_hierarchy - function:build_hierarchy
func TestListErrorMissingKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_missing_key_get_list - function:get_list
func TestListErrorMissingKeyGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_nested_missing_key_parse - function:parse
func TestListErrorNestedMissingKeyParse(t *testing.T) {

	ccl := mock.New()
	input := `config =
  server = web1`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  server = web1"}}
	assert.Equal(t, expected, parseResult)

}

// list_error_nested_missing_key_build_hierarchy - function:build_hierarchy
func TestListErrorNestedMissingKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_nested_missing_key_get_list - function:get_list
func TestListErrorNestedMissingKeyGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_non_object_path_parse - function:parse
func TestListErrorNonObjectPathParse(t *testing.T) {

	ccl := mock.New()
	input := `value = simple`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "value", Value: "simple"}}
	assert.Equal(t, expected, parseResult)

}

// list_error_non_object_path_build_hierarchy - function:build_hierarchy
func TestListErrorNonObjectPathBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_non_object_path_get_list - function:get_list
func TestListErrorNonObjectPathGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_parse_value - function:parse_value feature:multiline
func TestListMultilineValuesParseValue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_build_hierarchy - function:build_hierarchy feature:multiline
func TestListMultilineValuesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_get_list - function:get_list feature:multiline
func TestListMultilineValuesGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_parse_value - function:parse_value behavior:list_coercion_enabled
func TestComplexMixedListScenariosParseValue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_build_hierarchy - function:build_hierarchy behavior:list_coercion_enabled
func TestComplexMixedListScenariosBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_get_list - function:get_list behavior:list_coercion_enabled
func TestComplexMixedListScenariosGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled
func TestComplexMixedListScenariosReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_reference_get_list - function:get_list behavior:list_coercion_disabled
func TestComplexMixedListScenariosReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_edge_case_zero_length_parse - function:parse
func TestListEdgeCaseZeroLengthParse(t *testing.T) {

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

// list_edge_case_zero_length_build_hierarchy - function:build_hierarchy
func TestListEdgeCaseZeroLengthBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_edge_case_zero_length_get_list - function:get_list
func TestListEdgeCaseZeroLengthGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_parse - function:parse behavior:list_coercion_enabled
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

// list_path_traversal_protection_build_hierarchy - function:build_hierarchy behavior:list_coercion_enabled
func TestListPathTraversalProtectionBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_get_list - function:get_list behavior:list_coercion_enabled
func TestListPathTraversalProtectionGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_reference_parse - function:parse behavior:list_coercion_disabled
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

// list_path_traversal_protection_reference_build_hierarchy - function:build_hierarchy behavior:list_coercion_disabled
func TestListPathTraversalProtectionReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_reference_get_list - function:get_list behavior:list_coercion_disabled
func TestListPathTraversalProtectionReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
