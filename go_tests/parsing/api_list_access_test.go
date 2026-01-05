package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
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
	

	ccl := mock.New()
	input := `servers = web1
servers = web2
servers = web3`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"servers": []interface{}{"web1", "web2", "web3"}}
	assert.Equal(t, expected, objectResult)

}


// basic_list_from_duplicates_get_list - function:get_list behavior:list_coercion_enabled
func TestBasicListFromDuplicatesGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `servers = web1
servers = web2
servers = web3`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"servers"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"web1", "web2", "web3"}, result)

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
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"items": []interface{}{"item01", "item02", "item03", "item04", "item05", "item06", "item07", "item08", "item09", "item10", "item11", "item12", "item13", "item14", "item15", "item16", "item17", "item18", "item19", "item20"}}
	assert.Equal(t, expected, objectResult)

}


// large_list_get_list - function:get_list behavior:list_coercion_enabled
func TestLargeListGetList(t *testing.T) {
	

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
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"items"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"item01", "item02", "item03", "item04", "item05", "item06", "item07", "item08", "item09", "item10", "item11", "item12", "item13", "item14", "item15", "item16", "item17", "item18", "item19", "item20"}, result)

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


// list_with_comments_build_hierarchy - function:build_hierarchy feature:comments behavior:array_order_insertion
func TestListWithCommentsBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `servers = web1
/= Production servers
servers = web2
servers = web3
/= End of list`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"/": []interface{}{"Production servers", "End of list"}, "servers": []interface{}{"web1", "web2", "web3"}}
	assert.Equal(t, expected, objectResult)

}


// list_with_comments_get_list - function:get_list feature:comments behavior:list_coercion_enabled behavior:array_order_insertion
func TestListWithCommentsGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `servers = web1
/= Production servers
servers = web2
servers = web3
/= End of list`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"servers"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"web1", "web2", "web3"}, result)

}


// list_with_comments_lexicographic_parse - function:parse feature:comments
func TestListWithCommentsLexicographicParse(t *testing.T) {
	

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


// list_with_comments_lexicographic_build_hierarchy - function:build_hierarchy feature:comments behavior:array_order_lexicographic
func TestListWithCommentsLexicographicBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `servers = web1
/= Production servers
servers = web2
servers = web3
/= End of list`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"/": []interface{}{"End of list", "Production servers"}, "servers": []interface{}{"web1", "web2", "web3"}}
	assert.Equal(t, expected, objectResult)

}


// list_with_comments_lexicographic_get_list - function:get_list feature:comments behavior:list_coercion_enabled behavior:array_order_lexicographic
func TestListWithCommentsLexicographicGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `servers = web1
/= Production servers
servers = web2
servers = web3
/= End of list`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"servers"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"web1", "web2", "web3"}, result)

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
	

	ccl := mock.New()
	input := `existing = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"existing": "value"}
	assert.Equal(t, expected, objectResult)

}


// list_error_missing_key_get_list - function:get_list
func TestListErrorMissingKeyGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `existing = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"missing"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Empty(t, result)
	}

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
	

	ccl := mock.New()
	input := `config =
  server = web1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"server": "web1"}}
	assert.Equal(t, expected, objectResult)

}


// list_error_nested_missing_key_get_list - function:get_list
func TestListErrorNestedMissingKeyGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  server = web1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"config", "missing"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Empty(t, result)
	}

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
	

	ccl := mock.New()
	input := `value = simple`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"value": "simple"}
	assert.Equal(t, expected, objectResult)

}


// list_error_non_object_path_get_list - function:get_list
func TestListErrorNonObjectPathGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `value = simple`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"value", "nested"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Empty(t, result)
	}

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
	

	ccl := mock.New()
	input := ""
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// list_edge_case_zero_length_get_list - function:get_list
func TestListEdgeCaseZeroLengthGetList(t *testing.T) {
	

	ccl := mock.New()
	input := ""
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"nonexistent"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Empty(t, result)
	}

}


// bare_list_basic_parse - function:parse feature:empty_keys
func TestBareListBasicParse(t *testing.T) {
	

	ccl := mock.New()
	input := `servers =
  = web1
  = web2
  = web3`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "\n  = web1\n  = web2\n  = web3"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_basic_build_hierarchy - function:build_hierarchy feature:empty_keys
func TestBareListBasicBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `servers =
  = web1
  = web2
  = web3`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"servers": map[string]interface{}{"": []interface{}{"web1", "web2", "web3"}}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_basic_get_list - function:get_list feature:empty_keys
func TestBareListBasicGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `servers =
  = web1
  = web2
  = web3`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"servers"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"web1", "web2", "web3"}, result)

}


// bare_list_nested_parse - function:parse feature:empty_keys
func TestBareListNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `network =
  ports =
    = 80
    = 443
    = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "network", Value: "\n  ports =\n    = 80\n    = 443\n    = 8080"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_build_hierarchy - function:build_hierarchy feature:empty_keys behavior:array_order_insertion
func TestBareListNestedBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `network =
  ports =
    = 80
    = 443
    = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"network": map[string]interface{}{"ports": map[string]interface{}{"": []interface{}{"80", "443", "8080"}}}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_nested_get_list - function:get_list feature:empty_keys behavior:array_order_insertion
func TestBareListNestedGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `network =
  ports =
    = 80
    = 443
    = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"network", "ports"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"80", "443", "8080"}, result)

}


// bare_list_nested_lexicographic_parse - function:parse feature:empty_keys
func TestBareListNestedLexicographicParse(t *testing.T) {
	

	ccl := mock.New()
	input := `network =
  ports =
    = 80
    = 443
    = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "network", Value: "\n  ports =\n    = 80\n    = 443\n    = 8080"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_lexicographic_build_hierarchy - function:build_hierarchy feature:empty_keys behavior:array_order_lexicographic
func TestBareListNestedLexicographicBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `network =
  ports =
    = 80
    = 443
    = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"network": map[string]interface{}{"ports": map[string]interface{}{"": []interface{}{"443", "80", "8080"}}}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_nested_lexicographic_get_list - function:get_list feature:empty_keys behavior:array_order_lexicographic
func TestBareListNestedLexicographicGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `network =
  ports =
    = 80
    = 443
    = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"network", "ports"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"443", "80", "8080"}, result)

}


// bare_list_with_comments_parse - function:parse feature:empty_keys feature:comments
func TestBareListWithCommentsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `allowed_hosts =
  /= Production hosts
  = localhost
  = 127.0.0.1
  = example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "allowed_hosts", Value: "\n  /= Production hosts\n  = localhost\n  = 127.0.0.1\n  = example.com"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_with_comments_build_hierarchy - function:build_hierarchy feature:empty_keys feature:comments behavior:array_order_insertion
func TestBareListWithCommentsBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `allowed_hosts =
  /= Production hosts
  = localhost
  = 127.0.0.1
  = example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"allowed_hosts": map[string]interface{}{"": []interface{}{"localhost", "127.0.0.1", "example.com"}, "/": "Production hosts"}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_with_comments_get_list - function:get_list feature:empty_keys feature:comments behavior:array_order_insertion
func TestBareListWithCommentsGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `allowed_hosts =
  /= Production hosts
  = localhost
  = 127.0.0.1
  = example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"allowed_hosts"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"localhost", "127.0.0.1", "example.com"}, result)

}


// bare_list_with_comments_lexicographic_parse - function:parse feature:empty_keys feature:comments
func TestBareListWithCommentsLexicographicParse(t *testing.T) {
	

	ccl := mock.New()
	input := `allowed_hosts =
  /= Production hosts
  = localhost
  = 127.0.0.1
  = example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "allowed_hosts", Value: "\n  /= Production hosts\n  = localhost\n  = 127.0.0.1\n  = example.com"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_with_comments_lexicographic_build_hierarchy - function:build_hierarchy feature:empty_keys feature:comments behavior:array_order_lexicographic
func TestBareListWithCommentsLexicographicBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `allowed_hosts =
  /= Production hosts
  = localhost
  = 127.0.0.1
  = example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"allowed_hosts": map[string]interface{}{"": []interface{}{"127.0.0.1", "example.com", "localhost"}, "/": "Production hosts"}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_with_comments_lexicographic_get_list - function:get_list feature:empty_keys feature:comments behavior:array_order_lexicographic
func TestBareListWithCommentsLexicographicGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `allowed_hosts =
  /= Production hosts
  = localhost
  = 127.0.0.1
  = example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"allowed_hosts"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"127.0.0.1", "example.com", "localhost"}, result)

}


// bare_list_deeply_nested_parse - function:parse feature:empty_keys
func TestBareListDeeplyNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers =
        = web1
        = web2
        = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  environments =\n    production =\n      servers =\n        = web1\n        = web2\n        = api1"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_deeply_nested_build_hierarchy - function:build_hierarchy feature:empty_keys behavior:array_order_insertion
func TestBareListDeeplyNestedBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers =
        = web1
        = web2
        = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"environments": map[string]interface{}{"production": map[string]interface{}{"servers": map[string]interface{}{"": []interface{}{"web1", "web2", "api1"}}}}}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_deeply_nested_get_list - function:get_list feature:empty_keys behavior:array_order_insertion
func TestBareListDeeplyNestedGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers =
        = web1
        = web2
        = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"config", "environments", "production", "servers"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"web1", "web2", "api1"}, result)

}


// bare_list_deeply_nested_lexicographic_parse - function:parse feature:empty_keys
func TestBareListDeeplyNestedLexicographicParse(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers =
        = web1
        = web2
        = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  environments =\n    production =\n      servers =\n        = web1\n        = web2\n        = api1"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_deeply_nested_lexicographic_build_hierarchy - function:build_hierarchy feature:empty_keys behavior:array_order_lexicographic
func TestBareListDeeplyNestedLexicographicBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers =
        = web1
        = web2
        = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"environments": map[string]interface{}{"production": map[string]interface{}{"servers": map[string]interface{}{"": []interface{}{"api1", "web1", "web2"}}}}}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_deeply_nested_lexicographic_get_list - function:get_list feature:empty_keys behavior:array_order_lexicographic
func TestBareListDeeplyNestedLexicographicGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  environments =
    production =
      servers =
        = web1
        = web2
        = api1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"config", "environments", "production", "servers"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"api1", "web1", "web2"}, result)

}


// bare_list_mixed_with_other_keys_parse - function:parse feature:empty_keys
func TestBareListMixedWithOtherKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432
  replicas =
    = replica1
    = replica2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database", Value: "\n  host = localhost\n  port = 5432\n  replicas =\n    = replica1\n    = replica2"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_mixed_with_other_keys_build_hierarchy - function:build_hierarchy feature:empty_keys
func TestBareListMixedWithOtherKeysBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432
  replicas =
    = replica1
    = replica2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"database": map[string]interface{}{"host": "localhost", "port": "5432", "replicas": map[string]interface{}{"": []interface{}{"replica1", "replica2"}}}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_mixed_with_other_keys_get_list - function:get_list feature:empty_keys
func TestBareListMixedWithOtherKeysGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `database =
  host = localhost
  port = 5432
  replicas =
    = replica1
    = replica2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"database", "replicas"})
	require.NoError(t, err)
	assert.Equal(t, []interface {}{"replica1", "replica2"}, result)

}


// bare_list_error_not_a_list_parse - function:parse
func TestBareListErrorNotAListParse(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  setting = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  setting = value"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_error_not_a_list_build_hierarchy - function:build_hierarchy
func TestBareListErrorNotAListBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  setting = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"setting": "value"}}
	assert.Equal(t, expected, objectResult)

}


// bare_list_error_not_a_list_get_list - function:get_list behavior:list_coercion_disabled
func TestBareListErrorNotAListGetList(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  setting = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_list validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetList(hierarchy, []string{"config", "setting"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Empty(t, result)
	}

}


