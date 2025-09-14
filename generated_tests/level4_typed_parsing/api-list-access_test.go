package level4_typed_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-list-access.json
// Suite: CCL List Access - get_list() Function
// Version: 2.1
// Description: Level 4 CCL list access - type-aware list extraction with validation and conversion support using get_list() function


// basic_list_from_duplicates - function:get-list function:build-hierarchy function:parse (level 4)
func TestBasicListFromDuplicates(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers = web1
servers = web2
servers = web3`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "web3"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"servers": []interface{}{"web1", "web2", "web3"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// single_item_as_list - function:get-list function:build-hierarchy function:parse behavior:list-coercion-enabled variant:proposed-behavior (level 4)
func TestSingleItemAsList(t *testing.T) {
	
	
	ccl := mock.New()
	input := `item = single`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "item", Value: "single"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"item": "single"}
	assert.Equal(t, expectedObjects, objectResult)

}


// single_item_as_list_reference - function:get-list function:build-hierarchy function:parse behavior:list-coercion-disabled variant:reference-compliant (level 4)
func TestSingleItemAsListReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `item = single`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "item", Value: "single"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"item": "single"}
	assert.Equal(t, expectedObjects, objectResult)

}


// mixed_duplicate_single_keys - function:get-list function:build-hierarchy function:parse behavior:list-coercion-enabled variant:proposed-behavior (level 4)
func TestMixedDuplicateSingleKeys(t *testing.T) {
	
	
	ccl := mock.New()
	input := `ports = 80
ports = 443
host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "ports", Value: "80"}, mock.Entry{Key: "ports", Value: "443"}, mock.Entry{Key: "host", Value: "localhost"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"host": "localhost", "ports": []interface{}{"80", "443"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// mixed_duplicate_single_keys_reference - function:get-list function:build-hierarchy function:parse behavior:list-coercion-disabled variant:reference-compliant (level 4)
func TestMixedDuplicateSingleKeysReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `ports = 80
ports = 443
host = localhost`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "ports", Value: "80"}, mock.Entry{Key: "ports", Value: "443"}, mock.Entry{Key: "host", Value: "localhost"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"host": "localhost", "ports": []interface{}{"443", "80"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// nested_list_access - function:get-list function:build-hierarchy function:parse-value behavior:list-coercion-enabled variant:proposed-behavior (level 4)
func TestNestedListAccess(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  hosts = primary
  hosts = secondary
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "\n  hosts = primary\n  hosts = secondary\n  port = 5432"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"hosts": []interface{}{"primary", "secondary"}, "port": "5432"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// nested_list_access_reference - function:get-list function:build-hierarchy function:parse-value behavior:list-coercion-disabled variant:reference-compliant (level 4)
func TestNestedListAccessReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database =
  hosts = primary
  hosts = secondary
  port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database", Value: "\n  hosts = primary\n  hosts = secondary\n  port = 5432"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"hosts": []interface{}{"primary", "secondary"}, "port": "5432"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// empty_list - function:get-list function:build-hierarchy function:parse variant:proposed-behavior (level 4)
func TestEmptyList(t *testing.T) {
	
	
	ccl := mock.New()
	input := `empty_list =`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "empty_list", Value: ""}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"empty_list": ""}
	assert.Equal(t, expectedObjects, objectResult)

}


// empty_list_reference - function:get-list function:build-hierarchy function:parse variant:reference-compliant (level 4)
func TestEmptyListReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `empty_list =`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "empty_list", Value: ""}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"empty_list": ""}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_numbers - function:get-list function:build-hierarchy function:parse variant:proposed-behavior (level 4)
func TestListWithNumbers(t *testing.T) {
	
	
	ccl := mock.New()
	input := `numbers = 1
numbers = 42
numbers = -17
numbers = 0`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "numbers", Value: "1"}, mock.Entry{Key: "numbers", Value: "42"}, mock.Entry{Key: "numbers", Value: "-17"}, mock.Entry{Key: "numbers", Value: "0"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"numbers": []interface{}{"1", "42", "-17", "0"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_numbers_reference - function:get-list function:build-hierarchy function:parse variant:reference-compliant (level 4)
func TestListWithNumbersReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `numbers = 1
numbers = 42
numbers = -17
numbers = 0`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "numbers", Value: "1"}, mock.Entry{Key: "numbers", Value: "42"}, mock.Entry{Key: "numbers", Value: "-17"}, mock.Entry{Key: "numbers", Value: "0"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"numbers": []interface{}{"1", "42", "-17", "0"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_booleans - function:get-list function:build-hierarchy function:parse variant:proposed-behavior (level 4)
func TestListWithBooleans(t *testing.T) {
	
	
	ccl := mock.New()
	input := `flags = true
flags = false
flags = yes
flags = no`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "flags", Value: "true"}, mock.Entry{Key: "flags", Value: "false"}, mock.Entry{Key: "flags", Value: "yes"}, mock.Entry{Key: "flags", Value: "no"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"flags": []interface{}{"true", "false", "yes", "no"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_booleans_reference - function:get-list function:build-hierarchy function:parse variant:reference-compliant (level 4)
func TestListWithBooleansReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `flags = true
flags = false
flags = yes
flags = no`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "flags", Value: "true"}, mock.Entry{Key: "flags", Value: "false"}, mock.Entry{Key: "flags", Value: "yes"}, mock.Entry{Key: "flags", Value: "no"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"flags": []interface{}{"true", "false", "yes", "no"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_whitespace - function:get-list function:build-hierarchy function:parse feature:whitespace variant:proposed-behavior (level 4)
func TestListWithWhitespace(t *testing.T) {
	
	
	ccl := mock.New()
	input := `items =   spaced   
items = normal
items =
items =   `
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "items", Value: "spaced"}, mock.Entry{Key: "items", Value: "normal"}, mock.Entry{Key: "items", Value: ""}, mock.Entry{Key: "items", Value: ""}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"items": []interface{}{"spaced", "normal", "", ""}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_whitespace_reference - function:get-list function:build-hierarchy function:parse feature:whitespace variant:reference-compliant (level 4)
func TestListWithWhitespaceReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `items =   spaced   
items = normal
items =
items =   `
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "items", Value: "spaced"}, mock.Entry{Key: "items", Value: "normal"}, mock.Entry{Key: "items", Value: ""}, mock.Entry{Key: "items", Value: ""}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"items": []interface{}{"spaced", "normal"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// deeply_nested_list - function:get-list function:build-hierarchy function:parse variant:proposed-behavior (level 4)
func TestDeeplyNestedList(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers = web1
      servers = web2
      servers = api1`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: ""}, mock.Entry{Key: "environments", Value: ""}, mock.Entry{Key: "production", Value: ""}, mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "api1"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"environments": map[string]interface{}{"production": map[string]interface{}{"servers": []interface{}{"web1", "web2", "api1"}}}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// deeply_nested_list_reference - function:get-list function:build-hierarchy function:parse-value variant:reference-compliant (level 4)
func TestDeeplyNestedListReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers = web1
      servers = web2
      servers = api1`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: "\n  environments =\n    production =\n      servers = web1\n      servers = web2\n      servers = api1"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"environments": map[string]interface{}{"production": map[string]interface{}{"servers": []interface{}{"web1", "web2", "api1"}}}}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_unicode - function:get-list function:build-hierarchy function:parse feature:unicode variant:proposed-behavior (level 4)
func TestListWithUnicode(t *testing.T) {
	
	
	ccl := mock.New()
	input := `names = 张三
names = José
names = François
names = العربية`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "names", Value: "张三"}, mock.Entry{Key: "names", Value: "José"}, mock.Entry{Key: "names", Value: "François"}, mock.Entry{Key: "names", Value: "العربية"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"names": []interface{}{"张三", "José", "François", "العربية"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_unicode_reference - function:get-list function:build-hierarchy function:parse feature:unicode variant:reference-compliant (level 4)
func TestListWithUnicodeReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `names = 张三
names = José
names = François
names = العربية`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "names", Value: "张三"}, mock.Entry{Key: "names", Value: "José"}, mock.Entry{Key: "names", Value: "François"}, mock.Entry{Key: "names", Value: "العربية"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"names": []interface{}{"张三", "José", "François", "العربية"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_special_characters - function:get-list function:build-hierarchy function:parse variant:proposed-behavior (level 4)
func TestListWithSpecialCharacters(t *testing.T) {
	
	
	ccl := mock.New()
	input := `symbols = @#$%
symbols = !^&*()
symbols = []{}|
symbols = <>=+`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "symbols", Value: "@#$%"}, mock.Entry{Key: "symbols", Value: "!^&*()"}, mock.Entry{Key: "symbols", Value: "[]{}|"}, mock.Entry{Key: "symbols", Value: "<>=+"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"symbols": []interface{}{"@#$%", "!^&*()", "[]{}|", "<>=+"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_special_characters_reference - function:get-list function:build-hierarchy function:parse variant:reference-compliant (level 4)
func TestListWithSpecialCharactersReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `symbols = @#$%
symbols = !^&*()
symbols = []{}|`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "symbols", Value: "@#$%"}, mock.Entry{Key: "symbols", Value: "!^&*()"}, mock.Entry{Key: "symbols", Value: "[]{}|"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"symbols": []interface{}{"@#$%", "!^&*()", "[]{}|"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// large_list - function:get-list function:build-hierarchy function:parse performance-test (level 4)
func TestLargeList(t *testing.T) {
	
	
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
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "items", Value: "item01"}, mock.Entry{Key: "items", Value: "item02"}, mock.Entry{Key: "items", Value: "item03"}, mock.Entry{Key: "items", Value: "item04"}, mock.Entry{Key: "items", Value: "item05"}, mock.Entry{Key: "items", Value: "item06"}, mock.Entry{Key: "items", Value: "item07"}, mock.Entry{Key: "items", Value: "item08"}, mock.Entry{Key: "items", Value: "item09"}, mock.Entry{Key: "items", Value: "item10"}, mock.Entry{Key: "items", Value: "item11"}, mock.Entry{Key: "items", Value: "item12"}, mock.Entry{Key: "items", Value: "item13"}, mock.Entry{Key: "items", Value: "item14"}, mock.Entry{Key: "items", Value: "item15"}, mock.Entry{Key: "items", Value: "item16"}, mock.Entry{Key: "items", Value: "item17"}, mock.Entry{Key: "items", Value: "item18"}, mock.Entry{Key: "items", Value: "item19"}, mock.Entry{Key: "items", Value: "item20"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"items": []interface{}{"item01", "item02", "item03", "item04", "item05", "item06", "item07", "item08", "item09", "item10", "item11", "item12", "item13", "item14", "item15", "item16", "item17", "item18", "item19", "item20"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_with_comments - function:get-list function:build-hierarchy function:parse feature:comments (level 4)
func TestListWithComments(t *testing.T) {
	
	
	ccl := mock.New()
	input := `servers = web1
/= Production servers
servers = web2
servers = web3
/= End of list`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "/", Value: "Production servers"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "web3"}, mock.Entry{Key: "/", Value: "End of list"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"/": []interface{}{"Production servers", "End of list"}, "servers": []interface{}{"web1", "web2", "web3"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// dotted_key_list - function:get-list function:build-hierarchy function:parse function:expand-dotted feature:dotted-keys (level 4)
func TestDottedKeyList(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.hosts = primary
database.hosts = secondary
database.port = 5432`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "database.hosts", Value: "primary"}, mock.Entry{Key: "database.hosts", Value: "secondary"}, mock.Entry{Key: "database.port", Value: "5432"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"database": map[string]interface{}{"hosts": []interface{}{"primary", "secondary"}, "port": "5432"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_error_missing_key - function:get-list function:build-hierarchy function:parse (level 4)
func TestListErrorMissingKey(t *testing.T) {
	
	
	ccl := mock.New()
	input := `existing = value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "existing", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"existing": "value"}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_error_nested_missing_key - function:get-list function:build-hierarchy function:parse (level 4)
func TestListErrorNestedMissingKey(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  server = web1`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "config", Value: "\n  server = web1"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"server": "web1"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_error_non_object_path - function:get-list function:build-hierarchy function:parse (level 4)
func TestListErrorNonObjectPath(t *testing.T) {
	
	
	ccl := mock.New()
	input := `value = simple`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "value", Value: "simple"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"value": "simple"}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_multiline_values - function:get-list function:build-hierarchy function:parse-value feature:multiline variant:proposed-behavior (level 4)
func TestListMultilineValues(t *testing.T) {
	
	
	ccl := mock.New()
	input := `descriptions = First line
second line
descriptions = Another item
descriptions = Third item`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"descriptions": []interface{}{"First line", "Another item", "Third item"}, "second line": ""}
	assert.Equal(t, expectedObjects, objectResult)

}


// complex_mixed_list_scenarios - function:get-list function:build-hierarchy function:parse-value behavior:list-coercion-enabled variant:proposed-behavior (level 4)
func TestComplexMixedListScenarios(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  servers = web1
  servers = web2
  database =
    hosts = primary
    hosts = backup
    port = 5432
  cache = redis
features = auth
features = api
features = ui`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"cache": "redis", "database": map[string]interface{}{"hosts": []interface{}{"primary", "backup"}, "port": "5432"}, "servers": []interface{}{"web1", "web2"}}, "features": []interface{}{"auth", "api", "ui"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// complex_mixed_list_scenarios_reference - function:get-list function:build-hierarchy function:parse-value behavior:list-coercion-disabled variant:reference-compliant (level 4)
func TestComplexMixedListScenariosReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `config =
  servers = web1
  servers = web2
  database =
    hosts = primary
    hosts = backup
    port = 5432
  cache = redis
features = auth
features = api
features = ui`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"config": map[string]interface{}{"cache": "redis", "database": map[string]interface{}{"hosts": []interface{}{"backup", "primary"}, "port": "5432"}, "servers": []interface{}{"web2", "web1"}}, "features": []interface{}{"api", "auth", "ui"}}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_edge_case_zero_length - function:get-list function:build-hierarchy function:parse empty-input-test (level 4)
func TestListEdgeCaseZeroLength(t *testing.T) {
	
	
	ccl := mock.New()
	
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse("")
	require.NoError(t, err)
	expectedParse := []mock.Entry{}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_path_traversal_protection - function:get-list function:build-hierarchy function:parse behavior:list-coercion-enabled security-test variant:proposed-behavior (level 4)
func TestListPathTraversalProtection(t *testing.T) {
	
	
	ccl := mock.New()
	input := `safe = value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "safe", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"safe": "value"}
	assert.Equal(t, expectedObjects, objectResult)

}


// list_path_traversal_protection_reference - function:get-list function:build-hierarchy function:parse behavior:list-coercion-disabled security-test variant:reference-compliant (level 4)
func TestListPathTraversalProtectionReference(t *testing.T) {
	
	
	ccl := mock.New()
	input := `safe = value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "safe", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)
	// BuildHierarchy validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.BuildHierarchy(parseResult)
	expectedObjects := map[string]interface{}{"safe": "value"}
	assert.Equal(t, expectedObjects, objectResult)

}


