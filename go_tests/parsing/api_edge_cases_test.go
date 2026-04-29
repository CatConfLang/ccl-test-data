package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_edge_cases.json
// Suite: Flat Format
// Version: 1.0



// basic_single_no_spaces_parse - function:parse
func TestBasicSingleNoSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key=val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// basic_with_spaces_parse - function:parse feature:whitespace
func TestBasicWithSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// indented_key_parse_indented - function:parse_indented feature:whitespace
func TestIndentedKeyParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// value_trailing_spaces_parse - function:parse feature:whitespace
func TestValueTrailingSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key = val  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// key_value_surrounded_spaces_parse - function:parse feature:whitespace
func TestKeyValueSurroundedSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `  key  =  val  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// surrounded_by_newlines_parse - function:parse
func TestSurroundedByNewlinesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `
key = val
`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// key_empty_value_parse - function:parse feature:empty_keys
func TestKeyEmptyValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key =`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_value_with_newline_parse - function:parse feature:empty_keys
func TestEmptyValueWithNewlineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key =
`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_value_with_spaces_parse - function:parse feature:empty_keys feature:whitespace
func TestEmptyValueWithSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key =  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_key_indented_parse_indented - function:parse_indented feature:empty_keys
func TestEmptyKeyIndentedParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// empty_key_with_newline_parse - function:parse feature:empty_keys
func TestEmptyKeyWithNewlineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `
  = val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// empty_key_value_with_spaces_parse - function:parse feature:empty_keys feature:whitespace
func TestEmptyKeyValueWithSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `  =  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// equals_in_value_no_spaces_parse - function:parse
func TestEqualsInValueNoSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `a=b=c`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "b=c"}}
	assert.Equal(t, expected, parseResult)

}


// equals_in_value_with_spaces_parse - function:parse feature:whitespace
func TestEqualsInValueWithSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `a = b = c`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "b = c"}}
	assert.Equal(t, expected, parseResult)

}


// multiple_key_value_pairs_parse - function:parse
func TestMultipleKeyValuePairsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key1 = val1
key2 = val2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "val1"}, mock.Entry{Key: "key2", Value: "val2"}}
	assert.Equal(t, expected, parseResult)

}


// key_with_tabs_ocaml_reference_parse - function:parse feature:whitespace variant:reference_compliant
func TestKeyWithTabsOcamlReferenceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `	key	=	value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}


// whitespace_only_value_parse - function:parse feature:empty_keys feature:whitespace
func TestWhitespaceOnlyValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `onlyspaces =     `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "onlyspaces", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// multiple_empty_equality_parse - function:parse feature:empty_keys feature:whitespace
func TestMultipleEmptyEqualityParse(t *testing.T) {
	

	ccl := mock.New()
	input := ` =  = `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "="}}
	assert.Equal(t, expected, parseResult)

}


// key_with_newline_before_equals_parse - function:parse feature:multiline_keys feature:whitespace
func TestKeyWithNewlineBeforeEqualsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key 
= val
`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// complex_multi_newline_whitespace_parse - function:parse feature:multiline_keys feature:whitespace
func TestComplexMultiNewlineWhitespaceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `  
 key  
=  val  
`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// multiline_key_with_spaces_parse - function:parse feature:multiline_keys feature:whitespace
func TestMultilineKeyWithSpacesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `my
 key
= val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "my key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// multiline_key_three_lines_parse - function:parse feature:multiline_keys
func TestMultilineKeyThreeLinesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `a
 b
 c
= val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a b c", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// multiline_key_empty_value_parse - function:parse feature:multiline_keys feature:empty_keys
func TestMultilineKeyEmptyValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key
=`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// multiline_key_with_regular_entry_parse - function:parse feature:multiline_keys
func TestMultilineKeyWithRegularEntryParse(t *testing.T) {
	

	ccl := mock.New()
	input := `first = val1
key
= val2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "first", Value: "val1"}, mock.Entry{Key: "key", Value: "val2"}}
	assert.Equal(t, expected, parseResult)

}


// multiline_key_blank_lines_between_parse - function:parse feature:multiline_keys feature:whitespace
func TestMultilineKeyBlankLinesBetweenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key

= val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// multiline_key_tabs_in_continuation_parse - function:parse feature:multiline_keys feature:whitespace
func TestMultilineKeyTabsInContinuationParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key
	= val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "val"}}
	assert.Equal(t, expected, parseResult)

}


// empty_value_with_trailing_spaces_newline_parse - function:parse feature:empty_keys feature:whitespace
func TestEmptyValueWithTrailingSpacesNewlineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key =  
`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// empty_key_value_with_surrounding_newlines_parse - function:parse feature:empty_keys feature:whitespace
func TestEmptyKeyValueWithSurroundingNewlinesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `
  =  
`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// quotes_treated_as_literal_unquoted_parse - function:parse
func TestQuotesTreatedAsLiteralUnquotedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}}
	assert.Equal(t, expected, parseResult)

}


// quotes_treated_as_literal_quoted_parse - function:parse
func TestQuotesTreatedAsLiteralQuotedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host = "localhost"`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "\"localhost\""}}
	assert.Equal(t, expected, parseResult)

}


// nested_single_line_parse - function:parse
func TestNestedSingleLineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key =
  val`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\n  val"}}
	assert.Equal(t, expected, parseResult)

}


// nested_multi_line_parse - function:parse feature:multiline_continuation behavior:multiline_values
func TestNestedMultiLineParse(t *testing.T) {
	

	ccl := mock.New()
	input := `key =
  line1
  line2`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\n  line1\n  line2"}}
	assert.Equal(t, expected, parseResult)

}


// nested_with_blank_line_parse_indented - function:parse_indented feature:multiline_continuation behavior:multiline_values
func TestNestedWithBlankLineParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// deep_nested_structure_parse_indented - function:parse_indented
func TestDeepNestedStructureParseIndented(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// realistic_stress_test_parse - function:parse
func TestRealisticStressTestParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Dmitrii Kovanikov
login = chshersh
language = OCaml
date = 2024-05-25`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Dmitrii Kovanikov"}, mock.Entry{Key: "login", Value: "chshersh"}, mock.Entry{Key: "language", Value: "OCaml"}, mock.Entry{Key: "date", Value: "2024-05-25"}}
	assert.Equal(t, expected, parseResult)

}


// ocaml_stress_test_original_parse - function:parse feature:comments feature:empty_keys
func TestOcamlStressTestOriginalParse(t *testing.T) {
	

	ccl := mock.New()
	input := `/= This is a CCL document
title = CCL Example

database =
  enabled = true
  ports =
    = 8000
    = 8001
    = 8002
  limits =
    cpu = 1500mi
    memory = 10Gb

user =
  guestId = 42

user =
  login = chshersh
  createdAt = 2024-12-31`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "This is a CCL document"}, mock.Entry{Key: "title", Value: "CCL Example"}, mock.Entry{Key: "database", Value: "\n  enabled = true\n  ports =\n    = 8000\n    = 8001\n    = 8002\n  limits =\n    cpu = 1500mi\n    memory = 10Gb"}, mock.Entry{Key: "user", Value: "\n  guestId = 42"}, mock.Entry{Key: "user", Value: "\n  login = chshersh\n  createdAt = 2024-12-31"}}
	assert.Equal(t, expected, parseResult)

}


// ocaml_stress_test_original_build_hierarchy - function:parse_indented function:build_hierarchy feature:comments feature:empty_keys
func TestOcamlStressTestOriginalBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// ocaml_stress_test_original_build_model - function:parse_indented function:build_model feature:comments feature:empty_keys
func TestOcamlStressTestOriginalBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// ocaml_stress_test_original_get_string - function:get_string feature:comments feature:empty_keys
func TestOcamlStressTestOriginalGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// forward_slashes_in_map_keys_parse - function:parse
func TestForwardSlashesInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `mappings =
  config/settings.json = .vscode/settings.json
  src/template.env = .env`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "mappings", Value: "\n  config/settings.json = .vscode/settings.json\n  src/template.env = .env"}}
	assert.Equal(t, expected, parseResult)

}


// forward_slashes_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestForwardSlashesInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// forward_slashes_in_map_keys_build_model - function:parse_indented function:build_model
func TestForwardSlashesInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// forward_slashes_in_map_keys_get_string - function:get_string behavior:path_traversal
func TestForwardSlashesInMapKeysGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// backslashes_in_map_keys_parse - function:parse
func TestBackslashesInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `paths =
  C:\Users\config = user_settings
  D:\data\file.txt = backup`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "paths", Value: "\n  C:\\Users\\config = user_settings\n  D:\\data\\file.txt = backup"}}
	assert.Equal(t, expected, parseResult)

}


// backslashes_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestBackslashesInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// backslashes_in_map_keys_build_model - function:parse_indented function:build_model
func TestBackslashesInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// colons_in_map_keys_parse - function:parse
func TestColonsInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `timestamps =
  12:30:45 = morning
  23:59:59 = midnight`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "timestamps", Value: "\n  12:30:45 = morning\n  23:59:59 = midnight"}}
	assert.Equal(t, expected, parseResult)

}


// colons_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestColonsInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// colons_in_map_keys_build_model - function:parse_indented function:build_model
func TestColonsInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// hyphens_in_map_keys_parse - function:parse
func TestHyphensInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `packages =
  my-package-name = 1.0.0
  another-lib = 2.3.4`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "packages", Value: "\n  my-package-name = 1.0.0\n  another-lib = 2.3.4"}}
	assert.Equal(t, expected, parseResult)

}


// hyphens_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestHyphensInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// hyphens_in_map_keys_build_model - function:parse_indented function:build_model
func TestHyphensInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// at_signs_in_map_keys_parse - function:parse
func TestAtSignsInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `emails =
  user@example.com = primary
  admin@test.org = secondary`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "emails", Value: "\n  user@example.com = primary\n  admin@test.org = secondary"}}
	assert.Equal(t, expected, parseResult)

}


// at_signs_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestAtSignsInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// at_signs_in_map_keys_build_model - function:parse_indented function:build_model
func TestAtSignsInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// hash_in_map_keys_parse - function:parse
func TestHashInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `issues =
  issue#123 = open
  bug#456 = closed`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "issues", Value: "\n  issue#123 = open\n  bug#456 = closed"}}
	assert.Equal(t, expected, parseResult)

}


// hash_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestHashInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// hash_in_map_keys_build_model - function:parse_indented function:build_model
func TestHashInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// brackets_in_map_keys_parse - function:parse
func TestBracketsInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `arrays =
  items[0] = first
  items[1] = second`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "arrays", Value: "\n  items[0] = first\n  items[1] = second"}}
	assert.Equal(t, expected, parseResult)

}


// brackets_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestBracketsInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// brackets_in_map_keys_build_model - function:parse_indented function:build_model
func TestBracketsInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parentheses_in_map_keys_parse - function:parse
func TestParenthesesInMapKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `functions =
  init() = setup
  run(args) = execute`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "functions", Value: "\n  init() = setup\n  run(args) = execute"}}
	assert.Equal(t, expected, parseResult)

}


// parentheses_in_map_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestParenthesesInMapKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// parentheses_in_map_keys_build_model - function:parse_indented function:build_model
func TestParenthesesInMapKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_special_chars_in_keys_parse - function:parse
func TestMixedSpecialCharsInKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `complex =
  user@host:8080/api = endpoint
  file#v1.2.3 = release
  path\to\[item] = location`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "complex", Value: "\n  user@host:8080/api = endpoint\n  file#v1.2.3 = release\n  path\\to\\[item] = location"}}
	assert.Equal(t, expected, parseResult)

}


// mixed_special_chars_in_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestMixedSpecialCharsInKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_special_chars_in_keys_build_model - function:parse_indented function:build_model
func TestMixedSpecialCharsInKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_like_keys_parse - function:parse
func TestUrlLikeKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `endpoints =
  https://api.example.com/v1 = production
  http://localhost:3000/test = development`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "endpoints", Value: "\n  https://api.example.com/v1 = production\n  http://localhost:3000/test = development"}}
	assert.Equal(t, expected, parseResult)

}


// url_like_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlLikeKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_like_keys_build_model - function:parse_indented function:build_model
func TestUrlLikeKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_parent_parent_parse - function:parse
func TestRelativePathParentParentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `../.. = up_two_levels`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "../..", Value: "up_two_levels"}}
	assert.Equal(t, expected, parseResult)

}


// relative_path_parent_parent_build_hierarchy - function:parse_indented function:build_hierarchy
func TestRelativePathParentParentBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_parent_parent_build_model - function:parse_indented function:build_model
func TestRelativePathParentParentBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_parent_parse - function:parse
func TestRelativePathParentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `../ = parent_dir`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "../", Value: "parent_dir"}}
	assert.Equal(t, expected, parseResult)

}


// relative_path_parent_build_hierarchy - function:parse_indented function:build_hierarchy
func TestRelativePathParentBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_parent_build_model - function:parse_indented function:build_model
func TestRelativePathParentBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_single_dot_parse - function:parse
func TestRelativePathSingleDotParse(t *testing.T) {
	

	ccl := mock.New()
	input := `. = current_dir`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ".", Value: "current_dir"}}
	assert.Equal(t, expected, parseResult)

}


// relative_path_single_dot_build_hierarchy - function:parse_indented function:build_hierarchy
func TestRelativePathSingleDotBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_single_dot_build_model - function:parse_indented function:build_model
func TestRelativePathSingleDotBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// double_slash_parse - function:parse
func TestDoubleSlashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `// = double_slash_value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "//", Value: "double_slash_value"}}
	assert.Equal(t, expected, parseResult)

}


// double_slash_build_hierarchy - function:parse_indented function:build_hierarchy
func TestDoubleSlashBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// double_slash_build_model - function:parse_indented function:build_model
func TestDoubleSlashBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_in_value_parse - function:parse
func TestRelativePathInValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `path = ../../src`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "path", Value: "../../src"}}
	assert.Equal(t, expected, parseResult)

}


// relative_path_in_value_build_hierarchy - function:parse_indented function:build_hierarchy
func TestRelativePathInValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_in_value_build_model - function:parse_indented function:build_model
func TestRelativePathInValueBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_in_value_get_string - function:get_string
func TestRelativePathInValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_in_nested_value_parse - function:parse
func TestRelativePathInNestedValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `mappings =
  ../foo = ../bar
  ../../config = ../../data`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "mappings", Value: "\n  ../foo = ../bar\n  ../../config = ../../data"}}
	assert.Equal(t, expected, parseResult)

}


// relative_path_in_nested_value_build_hierarchy - function:parse_indented function:build_hierarchy
func TestRelativePathInNestedValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_in_nested_value_build_model - function:parse_indented function:build_model
func TestRelativePathInNestedValueBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_path_in_nested_value_get_string - function:get_string behavior:path_traversal
func TestRelativePathInNestedValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// double_slash_in_nested_parse - function:parse
func TestDoubleSlashInNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `urls =
  //api = //backup
  //server = /root`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "urls", Value: "\n  //api = //backup\n  //server = /root"}}
	assert.Equal(t, expected, parseResult)

}


// double_slash_in_nested_build_hierarchy - function:parse_indented function:build_hierarchy
func TestDoubleSlashInNestedBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// double_slash_in_nested_build_model - function:parse_indented function:build_model
func TestDoubleSlashInNestedBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_paths_deeply_nested_parse - function:parse
func TestRelativePathsDeeplyNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  build = 
    output = ../../dist
    cache = ../.cache`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  build = \n    output = ../../dist\n    cache = ../.cache"}}
	assert.Equal(t, expected, parseResult)

}


// relative_paths_deeply_nested_build_hierarchy - function:parse_indented function:build_hierarchy
func TestRelativePathsDeeplyNestedBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_paths_deeply_nested_build_model - function:parse_indented function:build_model
func TestRelativePathsDeeplyNestedBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_paths_deeply_nested_get_string - function:get_string behavior:path_traversal
func TestRelativePathsDeeplyNestedGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_paths_as_nested_keys_parse - function:parse
func TestRelativePathsAsNestedKeysParse(t *testing.T) {
	

	ccl := mock.New()
	input := `imports =
  .. = 
    main = ../src/main
    test = ../tests`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "imports", Value: "\n  .. = \n    main = ../src/main\n    test = ../tests"}}
	assert.Equal(t, expected, parseResult)

}


// relative_paths_as_nested_keys_build_hierarchy - function:parse_indented function:build_hierarchy
func TestRelativePathsAsNestedKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_paths_as_nested_keys_build_model - function:parse_indented function:build_model
func TestRelativePathsAsNestedKeysBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// relative_paths_as_nested_keys_get_string - function:get_string behavior:path_traversal
func TestRelativePathsAsNestedKeysGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_relative_and_absolute_nested_parse - function:parse
func TestMixedRelativeAndAbsoluteNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `paths =
  relative = 
    up = ../../
    current = .
  absolute = 
    root = /`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "paths", Value: "\n  relative = \n    up = ../../\n    current = .\n  absolute = \n    root = /"}}
	assert.Equal(t, expected, parseResult)

}


// mixed_relative_and_absolute_nested_build_hierarchy - function:parse_indented function:build_hierarchy
func TestMixedRelativeAndAbsoluteNestedBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_relative_and_absolute_nested_build_model - function:parse_indented function:build_model
func TestMixedRelativeAndAbsoluteNestedBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// mixed_relative_and_absolute_nested_get_string - function:get_string behavior:path_traversal
func TestMixedRelativeAndAbsoluteNestedGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// double_slash_deeply_nested_parse - function:parse
func TestDoubleSlashDeeplyNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `servers =
  primary = 
    api = //api.example.com
    cdn = //cdn.example.com
  secondary = 
    //internal = //backup.example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "\n  primary = \n    api = //api.example.com\n    cdn = //cdn.example.com\n  secondary = \n    //internal = //backup.example.com"}}
	assert.Equal(t, expected, parseResult)

}


// double_slash_deeply_nested_build_hierarchy - function:parse_indented function:build_hierarchy
func TestDoubleSlashDeeplyNestedBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// double_slash_deeply_nested_build_model - function:parse_indented function:build_model
func TestDoubleSlashDeeplyNestedBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// double_slash_deeply_nested_get_string - function:get_string behavior:path_traversal
func TestDoubleSlashDeeplyNestedGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_as_key_parse - function:parse
func TestUrlAsKeyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `https://api.example.com = production`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "https://api.example.com", Value: "production"}}
	assert.Equal(t, expected, parseResult)

}


// url_as_key_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlAsKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_as_key_build_model - function:parse_indented function:build_model
func TestUrlAsKeyBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_as_key_get_string - function:get_string
func TestUrlAsKeyGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_port_as_key_parse - function:parse
func TestUrlWithPortAsKeyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `http://localhost:8080 = dev`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "http://localhost:8080", Value: "dev"}}
	assert.Equal(t, expected, parseResult)

}


// url_with_port_as_key_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlWithPortAsKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_port_as_key_build_model - function:parse_indented function:build_model
func TestUrlWithPortAsKeyBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_path_as_key_parse - function:parse
func TestUrlWithPathAsKeyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `https://api.example.com/v1/users = users_endpoint`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "https://api.example.com/v1/users", Value: "users_endpoint"}}
	assert.Equal(t, expected, parseResult)

}


// url_with_path_as_key_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlWithPathAsKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_path_as_key_build_model - function:parse_indented function:build_model
func TestUrlWithPathAsKeyBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_in_value_parse - function:parse
func TestUrlInValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `endpoint = https://api.example.com/v1/data`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "endpoint", Value: "https://api.example.com/v1/data"}}
	assert.Equal(t, expected, parseResult)

}


// url_in_value_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlInValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_in_value_build_model - function:parse_indented function:build_model
func TestUrlInValueBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_in_value_get_string - function:get_string
func TestUrlInValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// urls_as_nested_keys_and_values_parse - function:parse
func TestUrlsAsNestedKeysAndValuesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `mappings =
  https://api.example.com = https://prod.example.com
  https://staging.example.com = https://stage.example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "mappings", Value: "\n  https://api.example.com = https://prod.example.com\n  https://staging.example.com = https://stage.example.com"}}
	assert.Equal(t, expected, parseResult)

}


// urls_as_nested_keys_and_values_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlsAsNestedKeysAndValuesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// urls_as_nested_keys_and_values_build_model - function:parse_indented function:build_model
func TestUrlsAsNestedKeysAndValuesBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// urls_as_nested_keys_and_values_get_string - function:get_string behavior:path_traversal
func TestUrlsAsNestedKeysAndValuesGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// urls_deeply_nested_parse - function:parse
func TestUrlsDeeplyNestedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  services = 
    api = 
      url = https://api.example.com
      backup = https://backup.example.com
    cdn = 
      url = https://cdn.example.com`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  services = \n    api = \n      url = https://api.example.com\n      backup = https://backup.example.com\n    cdn = \n      url = https://cdn.example.com"}}
	assert.Equal(t, expected, parseResult)

}


// urls_deeply_nested_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlsDeeplyNestedBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// urls_deeply_nested_build_model - function:parse_indented function:build_model
func TestUrlsDeeplyNestedBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// urls_deeply_nested_get_string - function:get_string behavior:path_traversal
func TestUrlsDeeplyNestedGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_query_params_as_key_parse - function:parse behavior:delimiter_prefer_spaced
func TestUrlWithQueryParamsAsKeyParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:delimiter_prefer_spaced")
}


// url_with_query_params_as_key_build_hierarchy - function:parse_indented function:build_hierarchy behavior:delimiter_prefer_spaced
func TestUrlWithQueryParamsAsKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_query_params_as_key_build_model - function:parse_indented function:build_model behavior:delimiter_prefer_spaced
func TestUrlWithQueryParamsAsKeyBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_first_url_with_query_params_parse - function:parse behavior:delimiter_first_equals
func TestDelimiterFirstUrlWithQueryParamsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `https://api.example.com/search?q=test&page=1 = search_results`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "https://api.example.com/search?q", Value: "test&page=1 = search_results"}}
	assert.Equal(t, expected, parseResult)

}


// delimiter_first_url_with_query_params_build_hierarchy - function:parse_indented function:build_hierarchy behavior:delimiter_first_equals
func TestDelimiterFirstUrlWithQueryParamsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_first_url_with_query_params_build_model - function:parse_indented function:build_model behavior:delimiter_first_equals
func TestDelimiterFirstUrlWithQueryParamsBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_first_multiple_equals_parse - function:parse behavior:delimiter_first_equals
func TestDelimiterFirstMultipleEqualsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `a=b = c=d`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "b = c=d"}}
	assert.Equal(t, expected, parseResult)

}


// delimiter_first_multiple_equals_build_hierarchy - function:parse_indented function:build_hierarchy behavior:delimiter_first_equals
func TestDelimiterFirstMultipleEqualsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_first_multiple_equals_build_model - function:parse_indented function:build_model behavior:delimiter_first_equals
func TestDelimiterFirstMultipleEqualsBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_first_empty_value_parse - function:parse behavior:delimiter_first_equals
func TestDelimiterFirstEmptyValueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `a=b =`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "b ="}}
	assert.Equal(t, expected, parseResult)

}


// delimiter_first_empty_value_build_hierarchy - function:parse_indented function:build_hierarchy behavior:delimiter_first_equals
func TestDelimiterFirstEmptyValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_first_empty_value_build_model - function:parse_indented function:build_model behavior:delimiter_first_equals
func TestDelimiterFirstEmptyValueBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_spaced_multiple_equals_parse - function:parse behavior:delimiter_prefer_spaced
func TestDelimiterSpacedMultipleEqualsParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:delimiter_prefer_spaced")
}


// delimiter_spaced_multiple_equals_build_hierarchy - function:parse_indented function:build_hierarchy behavior:delimiter_prefer_spaced
func TestDelimiterSpacedMultipleEqualsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_spaced_multiple_equals_build_model - function:parse_indented function:build_model behavior:delimiter_prefer_spaced
func TestDelimiterSpacedMultipleEqualsBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_spaced_fallback_no_space_parse - function:parse behavior:delimiter_prefer_spaced
func TestDelimiterSpacedFallbackNoSpaceParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:delimiter_prefer_spaced")
}


// delimiter_spaced_fallback_no_space_build_hierarchy - function:parse_indented function:build_hierarchy behavior:delimiter_prefer_spaced
func TestDelimiterSpacedFallbackNoSpaceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_spaced_fallback_no_space_build_model - function:parse_indented function:build_model behavior:delimiter_prefer_spaced
func TestDelimiterSpacedFallbackNoSpaceBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_spaced_empty_value_parse - function:parse behavior:delimiter_prefer_spaced
func TestDelimiterSpacedEmptyValueParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:delimiter_prefer_spaced")
}


// delimiter_spaced_empty_value_build_hierarchy - function:parse_indented function:build_hierarchy behavior:delimiter_prefer_spaced
func TestDelimiterSpacedEmptyValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// delimiter_spaced_empty_value_build_model - function:parse_indented function:build_model behavior:delimiter_prefer_spaced
func TestDelimiterSpacedEmptyValueBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_fragment_as_key_parse - function:parse
func TestUrlWithFragmentAsKeyParse(t *testing.T) {
	

	ccl := mock.New()
	input := `https://docs.example.com/guide#section-1 = docs_section_1`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "https://docs.example.com/guide#section-1", Value: "docs_section_1"}}
	assert.Equal(t, expected, parseResult)

}


// url_with_fragment_as_key_build_hierarchy - function:parse_indented function:build_hierarchy
func TestUrlWithFragmentAsKeyBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// url_with_fragment_as_key_build_model - function:parse_indented function:build_model
func TestUrlWithFragmentAsKeyBuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


