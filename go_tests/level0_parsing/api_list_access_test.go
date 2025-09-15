package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_list_access.json
// Suite: Flat Format
// Version: 1.0

// basic_list_from_duplicates_parse - function:parse (level 0)
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

// basic_list_from_duplicates_buildhierarchy - function:buildhierarchy (level 0)
func TestBasicListFromDuplicatesBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// basic_list_from_duplicates_getlist - function:getlist (level 0)
func TestBasicListFromDuplicatesGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_parse - function:parse behavior:list_coercion_enabled (level 0)
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

// single_item_as_list_buildhierarchy - function:buildhierarchy behavior:list_coercion_enabled (level 0)
func TestSingleItemAsListBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_getlist - function:getlist behavior:list_coercion_enabled (level 0)
func TestSingleItemAsListGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// single_item_as_list_reference_parse - function:parse behavior:list_coercion_disabled (level 0)
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

// single_item_as_list_reference_buildhierarchy - function:buildhierarchy behavior:list_coercion_disabled (level 0)
func TestSingleItemAsListReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_parse - function:parse behavior:list_coercion_enabled (level 0)
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

// mixed_duplicate_single_keys_buildhierarchy - function:buildhierarchy behavior:list_coercion_enabled (level 0)
func TestMixedDuplicateSingleKeysBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_getlist - function:getlist behavior:list_coercion_enabled (level 0)
func TestMixedDuplicateSingleKeysGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_reference_parse - function:parse behavior:list_coercion_disabled (level 0)
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

// mixed_duplicate_single_keys_reference_buildhierarchy - function:buildhierarchy behavior:list_coercion_disabled (level 0)
func TestMixedDuplicateSingleKeysReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_parse - function:parse behavior:list_coercion_enabled (level 0)
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

// nested_list_access_buildhierarchy - function:buildhierarchy behavior:list_coercion_enabled (level 0)
func TestNestedListAccessBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_getlist - function:getlist behavior:list_coercion_enabled (level 0)
func TestNestedListAccessGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// nested_list_access_reference_parse - function:parse behavior:list_coercion_disabled (level 0)
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

// nested_list_access_reference_buildhierarchy - function:buildhierarchy behavior:list_coercion_disabled (level 0)
func TestNestedListAccessReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_parse - function:parse (level 0)
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

// empty_list_buildhierarchy - function:buildhierarchy (level 0)
func TestEmptyListBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_getlist - function:getlist (level 0)
func TestEmptyListGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_list_reference_parse - function:parse (level 0)
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

// empty_list_reference_buildhierarchy - function:buildhierarchy (level 0)
func TestEmptyListReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_parse - function:parse (level 0)
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

// list_with_numbers_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithNumbersBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_getlist - function:getlist (level 0)
func TestListWithNumbersGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_numbers_reference_parse - function:parse (level 0)
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

// list_with_numbers_reference_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithNumbersReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_parse - function:parse (level 0)
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

// list_with_booleans_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithBooleansBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_getlist - function:getlist (level 0)
func TestListWithBooleansGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_reference_parse - function:parse (level 0)
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

// list_with_booleans_reference_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithBooleansReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_parse - function:parse (level 0)
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

// list_with_whitespace_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithWhitespaceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_getlist - function:getlist (level 0)
func TestListWithWhitespaceGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_whitespace_reference_parse - function:parse (level 0)
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

// list_with_whitespace_reference_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithWhitespaceReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_parse - function:parse (level 0)
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

// deeply_nested_list_buildhierarchy - function:buildhierarchy (level 0)
func TestDeeplyNestedListBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_getlist - function:getlist (level 0)
func TestDeeplyNestedListGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deeply_nested_list_reference_parse - function:parse (level 0)
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

// deeply_nested_list_reference_buildhierarchy - function:buildhierarchy (level 0)
func TestDeeplyNestedListReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_parse - function:parse (level 0)
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

// list_with_unicode_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithUnicodeBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_getlist - function:getlist (level 0)
func TestListWithUnicodeGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_unicode_reference_parse - function:parse (level 0)
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

// list_with_unicode_reference_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithUnicodeReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_parse - function:parse (level 0)
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

// list_with_special_characters_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithSpecialCharactersBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_getlist - function:getlist (level 0)
func TestListWithSpecialCharactersGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_special_characters_reference_parse - function:parse (level 0)
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

// list_with_special_characters_reference_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithSpecialCharactersReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// large_list_parse - function:parse (level 0)
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

// large_list_buildhierarchy - function:buildhierarchy (level 0)
func TestLargeListBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// large_list_getlist - function:getlist (level 0)
func TestLargeListGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_comments_parse - function:parse (level 0)
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

// list_with_comments_buildhierarchy - function:buildhierarchy (level 0)
func TestListWithCommentsBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_comments_getlist - function:getlist (level 0)
func TestListWithCommentsGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_missing_key_parse - function:parse (level 0)
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

// list_error_missing_key_buildhierarchy - function:buildhierarchy (level 0)
func TestListErrorMissingKeyBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_nested_missing_key_parse - function:parse (level 0)
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

// list_error_nested_missing_key_buildhierarchy - function:buildhierarchy (level 0)
func TestListErrorNestedMissingKeyBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_error_non_object_path_parse - function:parse (level 0)
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

// list_error_non_object_path_buildhierarchy - function:buildhierarchy (level 0)
func TestListErrorNonObjectPathBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_parsevalue - function:parsevalue (level 0)
func TestListMultilineValuesParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_buildhierarchy - function:buildhierarchy (level 0)
func TestListMultilineValuesBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_multiline_values_getlist - function:getlist (level 0)
func TestListMultilineValuesGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_parsevalue - function:parsevalue behavior:list_coercion_enabled (level 0)
func TestComplexMixedListScenariosParsevalue(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_buildhierarchy - function:buildhierarchy behavior:list_coercion_enabled (level 0)
func TestComplexMixedListScenariosBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_getlist - function:getlist behavior:list_coercion_enabled (level 0)
func TestComplexMixedListScenariosGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// complex_mixed_list_scenarios_reference_buildhierarchy - function:buildhierarchy behavior:list_coercion_disabled (level 0)
func TestComplexMixedListScenariosReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_edge_case_zero_length_parse - function:parse (level 0)
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

// list_edge_case_zero_length_buildhierarchy - function:buildhierarchy (level 0)
func TestListEdgeCaseZeroLengthBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_parse - function:parse behavior:list_coercion_enabled (level 0)
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

// list_path_traversal_protection_buildhierarchy - function:buildhierarchy behavior:list_coercion_enabled (level 0)
func TestListPathTraversalProtectionBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_getlist - function:getlist behavior:list_coercion_enabled (level 0)
func TestListPathTraversalProtectionGetlist(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_path_traversal_protection_reference_parse - function:parse behavior:list_coercion_disabled (level 0)
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

// list_path_traversal_protection_reference_buildhierarchy - function:buildhierarchy behavior:list_coercion_disabled (level 0)
func TestListPathTraversalProtectionReferenceBuildhierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
