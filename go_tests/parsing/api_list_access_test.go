package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
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


// basic_list_from_duplicates_get_list - function:get_list behavior:list_coercion_enabled
func TestBasicListFromDuplicatesGetList(t *testing.T) {
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


// large_list_get_list - function:get_list behavior:list_coercion_enabled
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


// list_with_comments_build_hierarchy - function:build_hierarchy feature:comments behavior:array_order_insertion
func TestListWithCommentsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_comments_get_list - function:get_list feature:comments behavior:list_coercion_enabled behavior:array_order_insertion
func TestListWithCommentsGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_comments_lexicographic_get_list - function:get_list feature:comments behavior:list_coercion_enabled behavior:array_order_lexicographic
func TestListWithCommentsLexicographicGetList(t *testing.T) {
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_basic_get_list - function:get_list feature:empty_keys
func TestBareListBasicGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_get_list - function:get_list feature:empty_keys behavior:array_order_insertion
func TestBareListNestedGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_lexicographic_get_list - function:get_list feature:empty_keys behavior:array_order_lexicographic
func TestBareListNestedLexicographicGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_with_comments_get_list - function:get_list feature:empty_keys feature:comments behavior:array_order_insertion
func TestBareListWithCommentsGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_with_comments_lexicographic_get_list - function:get_list feature:empty_keys feature:comments behavior:array_order_lexicographic
func TestBareListWithCommentsLexicographicGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_deeply_nested_get_list - function:get_list feature:empty_keys behavior:array_order_insertion
func TestBareListDeeplyNestedGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_deeply_nested_lexicographic_get_list - function:get_list feature:empty_keys behavior:array_order_lexicographic
func TestBareListDeeplyNestedLexicographicGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_mixed_with_other_keys_get_list - function:get_list feature:empty_keys
func TestBareListMixedWithOtherKeysGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_error_not_a_list_get_list - function:get_list behavior:list_coercion_disabled
func TestBareListErrorNotAListGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_basic_parse - function:parse feature:empty_keys
func TestBareListNestedObjectsBasicParse(t *testing.T) {
	

	ccl := mock.New()
	input := `items =
  =
    name = first
    value = 1
  =
    name = second
    value = 2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "\n  =\n    name = first\n    value = 1\n  =\n    name = second\n    value = 2"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_objects_basic_build_hierarchy - function:build_hierarchy feature:empty_keys
func TestBareListNestedObjectsBasicBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_basic_get_list - function:get_list feature:empty_keys
func TestBareListNestedObjectsBasicGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_single_item_parse - function:parse feature:empty_keys
func TestBareListNestedObjectsSingleItemParse(t *testing.T) {
	

	ccl := mock.New()
	input := `items =
  =
    name = only
    value = 42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "\n  =\n    name = only\n    value = 42"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_objects_single_item_build_hierarchy - function:build_hierarchy feature:empty_keys
func TestBareListNestedObjectsSingleItemBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_single_item_get_list - function:get_list feature:empty_keys
func TestBareListNestedObjectsSingleItemGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_minimal_parse - function:parse feature:empty_keys
func TestBareListNestedObjectsMinimalParse(t *testing.T) {
	

	ccl := mock.New()
	input := `items =
  =
    name = first
  =
    name = second`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "\n  =\n    name = first\n  =\n    name = second"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_objects_minimal_build_hierarchy - function:build_hierarchy feature:empty_keys
func TestBareListNestedObjectsMinimalBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_deeply_nested_parse - function:parse feature:empty_keys
func TestBareListNestedObjectsDeeplyNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `items =
  =
    config =
      host = localhost
      port = 8080
  =
    config =
      host = example.com
      port = 443`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "\n  =\n    config =\n      host = localhost\n      port = 8080\n  =\n    config =\n      host = example.com\n      port = 443"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_objects_deeply_nested_build_hierarchy - function:build_hierarchy feature:empty_keys
func TestBareListNestedObjectsDeeplyNestedBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_mixed_with_strings_parse - function:parse feature:empty_keys
func TestBareListNestedObjectsMixedWithStringsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `items =
  = simple_string
  =
    name = nested
    value = obj
  = another_string`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "\n  = simple_string\n  =\n    name = nested\n    value = obj\n  = another_string"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_objects_mixed_with_strings_build_hierarchy - function:build_hierarchy feature:empty_keys behavior:array_order_insertion
func TestBareListNestedObjectsMixedWithStringsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_round_trip_parse - function:parse feature:empty_keys
func TestBareListNestedObjectsRoundTripParse(t *testing.T) {
	

	ccl := mock.New()
	input := `items =
  =
    name = first
    value = 1
  =
    name = second
    value = 2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "items", Value: "\n  =\n    name = first\n    value = 1\n  =\n    name = second\n    value = 2"}}
	assert.Equal(t, expected, parseResult)

}


// bare_list_nested_objects_round_trip_print - function:print feature:empty_keys
func TestBareListNestedObjectsRoundTripPrint(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// bare_list_nested_objects_round_trip_round_trip - function:parse function:print feature:empty_keys
func TestBareListNestedObjectsRoundTripRoundTrip(t *testing.T) {
	

	ccl := mock.New()
	input := `items =
  =
    name = first
    value = 1
  =
    name = second
    value = 2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement round_trip validation
	_ = ccl // Prevent unused variable warning
	_ = input // Prevent unused variable warning
	_ = err // Prevent unused variable warning

}


