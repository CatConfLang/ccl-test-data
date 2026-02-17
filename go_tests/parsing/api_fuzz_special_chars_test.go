package parsing_test

import (
	"testing"

	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_fuzz_special_chars.json
// Suite: Flat Format
// Version: 1.0

// fuzz_single_lparen_parse - function:parse
func TestFuzzSingleLparenParse(t *testing.T) {

	ccl := mock.New()
	input := `( = val620`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "(", Value: "val620"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_hyphen_parse - function:parse
func TestFuzzSingleHyphenParse(t *testing.T) {

	ccl := mock.New()
	input := `- = val619`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "-", Value: "val619"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_rbracket_parse - function:parse
func TestFuzzSingleRbracketParse(t *testing.T) {

	ccl := mock.New()
	input := `] = val334`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "]", Value: "val334"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_backslash_parse - function:parse
func TestFuzzSingleBackslashParse(t *testing.T) {

	ccl := mock.New()
	input := `\ = val595`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\\", Value: "val595"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_slash_parse - function:parse
func TestFuzzSingleSlashParse(t *testing.T) {

	ccl := mock.New()
	input := `/ = val509`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "val509"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_dollar_parse - function:parse
func TestFuzzSingleDollarParse(t *testing.T) {

	ccl := mock.New()
	input := `$ = val877`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$", Value: "val877"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_rbrace_parse - function:parse
func TestFuzzSingleRbraceParse(t *testing.T) {

	ccl := mock.New()
	input := `} = val644`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}", Value: "val644"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_plus_parse - function:parse
func TestFuzzSinglePlusParse(t *testing.T) {

	ccl := mock.New()
	input := `+ = val412`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "+", Value: "val412"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_dquote_parse - function:parse
func TestFuzzSingleDquoteParse(t *testing.T) {

	ccl := mock.New()
	input := `" = val691`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"", Value: "val691"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_bang_parse - function:parse
func TestFuzzSingleBangParse(t *testing.T) {

	ccl := mock.New()
	input := `! = val528`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "!", Value: "val528"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_hash_parse - function:parse
func TestFuzzSingleHashParse(t *testing.T) {

	ccl := mock.New()
	input := `# = val238`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "#", Value: "val238"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_gt_parse - function:parse
func TestFuzzSingleGtParse(t *testing.T) {

	ccl := mock.New()
	input := `> = val318`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ">", Value: "val318"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_lbracket_parse - function:parse
func TestFuzzSingleLbracketParse(t *testing.T) {

	ccl := mock.New()
	input := `[ = val72`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "[", Value: "val72"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_rparen_parse - function:parse
func TestFuzzSingleRparenParse(t *testing.T) {

	ccl := mock.New()
	input := `) = val688`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")", Value: "val688"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_single_question_parse - function:parse
func TestFuzzSingleQuestionParse(t *testing.T) {

	ccl := mock.New()
	input := `? = val825`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "?", Value: "val825"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_pipe_lt_semicolon_parse - function:parse
func TestFuzzComboPipeLtSemicolonParse(t *testing.T) {

	ccl := mock.New()
	input := `|<; = combo808`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "|<;", Value: "combo808"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_rbracket_dquote_asterisk_parse - function:parse
func TestFuzzComboRbracketDquoteAsteriskParse(t *testing.T) {

	ccl := mock.New()
	input := `]"* = combo591`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "]\"*", Value: "combo591"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_rbrace_squote_plus_parse - function:parse
func TestFuzzComboRbraceSquotePlusParse(t *testing.T) {

	ccl := mock.New()
	input := `}'+ = combo440`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}'+", Value: "combo440"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_ampersand_dquote_rparen_parse - function:parse
func TestFuzzComboAmpersandDquoteRparenParse(t *testing.T) {

	ccl := mock.New()
	input := `&") = combo488`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&\")", Value: "combo488"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_hyphen_tilde_parse - function:parse
func TestFuzzComboHyphenTildeParse(t *testing.T) {

	ccl := mock.New()
	input := `-~ = combo135`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "-~", Value: "combo135"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_at_lt_semicolon_lbracket_parse - function:parse
func TestFuzzComboAtLtSemicolonLbracketParse(t *testing.T) {

	ccl := mock.New()
	input := `@<;[ = combo491`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "@<;[", Value: "combo491"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_squote_underscore_parse - function:parse
func TestFuzzComboSquoteUnderscoreParse(t *testing.T) {

	ccl := mock.New()
	input := `'_ = combo768`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "'_", Value: "combo768"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_lt_colon_parse - function:parse
func TestFuzzComboLtColonParse(t *testing.T) {

	ccl := mock.New()
	input := `<: = combo590`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "<:", Value: "combo590"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_at_slash_asterisk_parse - function:parse
func TestFuzzComboAtSlashAsteriskParse(t *testing.T) {

	ccl := mock.New()
	input := `@/* = combo681`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "@/*", Value: "combo681"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_combo_dquote_pipe_slash_parse - function:parse
func TestFuzzComboDquotePipeSlashParse(t *testing.T) {

	ccl := mock.New()
	input := `"|/ = combo76`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"|/", Value: "combo76"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_start_bang_port_parse - function:parse
func TestFuzzPosStartBangPortParse(t *testing.T) {

	ccl := mock.New()
	input := `!port = pos235`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "!port", Value: "pos235"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_middle_squote_data_parse - function:parse
func TestFuzzPosMiddleSquoteDataParse(t *testing.T) {

	ccl := mock.New()
	input := `da'ta = pos941`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "da'ta", Value: "pos941"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_end_gt_data_parse - function:parse
func TestFuzzPosEndGtDataParse(t *testing.T) {

	ccl := mock.New()
	input := `data> = pos187`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "data>", Value: "pos187"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_start_slash_host_parse - function:parse
func TestFuzzPosStartSlashHostParse(t *testing.T) {

	ccl := mock.New()
	input := `/host = pos306`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/host", Value: "pos306"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_middle_hash_name_parse - function:parse
func TestFuzzPosMiddleHashNameParse(t *testing.T) {

	ccl := mock.New()
	input := `na#me = pos189`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "na#me", Value: "pos189"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_end_plus_delta_parse - function:parse
func TestFuzzPosEndPlusDeltaParse(t *testing.T) {

	ccl := mock.New()
	input := `delta+ = pos397`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "delta+", Value: "pos397"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_start_squote_path_parse - function:parse
func TestFuzzPosStartSquotePathParse(t *testing.T) {

	ccl := mock.New()
	input := `'path = pos238`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "'path", Value: "pos238"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_middle_lparen_port_parse - function:parse
func TestFuzzPosMiddleLparenPortParse(t *testing.T) {

	ccl := mock.New()
	input := `po(rt = pos139`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "po(rt", Value: "pos139"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_end_semicolon_gamma_parse - function:parse
func TestFuzzPosEndSemicolonGammaParse(t *testing.T) {

	ccl := mock.New()
	input := `gamma; = pos404`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "gamma;", Value: "pos404"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_pos_start_gt_alpha_parse - function:parse
func TestFuzzPosStartGtAlphaParse(t *testing.T) {

	ccl := mock.New()
	input := `>alpha = pos196`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ">alpha", Value: "pos196"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_val_data_0_parse - function:parse feature:optional_typed_accessors
func TestFuzzValData0Parse(t *testing.T) {

	ccl := mock.New()
	input := `data = data|x58`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "data", Value: "data|x58"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_val_data_0_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzValData0BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_data_0_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzValData0GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_beta_1_parse - function:parse feature:optional_typed_accessors
func TestFuzzValBeta1Parse(t *testing.T) {

	ccl := mock.New()
	input := `beta = data\x12`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "beta", Value: "data\\x12"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_val_beta_1_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzValBeta1BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_beta_1_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzValBeta1GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_name_2_parse - function:parse feature:optional_typed_accessors
func TestFuzzValName2Parse(t *testing.T) {

	ccl := mock.New()
	input := `name = data\x77"x95_x85`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "data\\x77\"x95_x85"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_val_name_2_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzValName2BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_name_2_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzValName2GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_path_3_parse - function:parse feature:optional_typed_accessors
func TestFuzzValPath3Parse(t *testing.T) {

	ccl := mock.New()
	input := `path = data-x45`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "path", Value: "data-x45"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_val_path_3_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzValPath3BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_path_3_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzValPath3GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_name_4_parse - function:parse feature:optional_typed_accessors
func TestFuzzValName4Parse(t *testing.T) {

	ccl := mock.New()
	input := `name = data{x49"x55`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "data{x49\"x55"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_val_name_4_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzValName4BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_val_name_4_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzValName4GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_rbracket_0_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedLbraceRbracket0Parse(t *testing.T) {

	ccl := mock.New()
	input := `{host = nested259
]name = deep985`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "{host", Value: "nested259"}, mock.Entry{Key: "]name", Value: "deep985"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_lbrace_rbracket_0_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedLbraceRbracket0BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_rbracket_0_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedLbraceRbracket0GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_rbrace_1_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedLbraceRbrace1Parse(t *testing.T) {

	ccl := mock.New()
	input := `{port = nested169
}mode = deep270`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "{port", Value: "nested169"}, mock.Entry{Key: "}mode", Value: "deep270"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_lbrace_rbrace_1_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedLbraceRbrace1BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_rbrace_1_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedLbraceRbrace1GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_ampersand_percent_2_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedAmpersandPercent2Parse(t *testing.T) {

	ccl := mock.New()
	input := `&delta = nested56
%data = deep952`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&delta", Value: "nested56"}, mock.Entry{Key: "%data", Value: "deep952"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_ampersand_percent_2_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedAmpersandPercent2BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_ampersand_percent_2_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedAmpersandPercent2GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_dquote_gt_3_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedDquoteGt3Parse(t *testing.T) {

	ccl := mock.New()
	input := `"epsilon = nested885
>user = deep168`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"epsilon", Value: "nested885"}, mock.Entry{Key: ">user", Value: "deep168"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_dquote_gt_3_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedDquoteGt3BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_dquote_gt_3_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedDquoteGt3GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_percent_dollar_4_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedPercentDollar4Parse(t *testing.T) {

	ccl := mock.New()
	input := `%alpha = nested438
$path = deep65`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%alpha", Value: "nested438"}, mock.Entry{Key: "$path", Value: "deep65"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_percent_dollar_4_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedPercentDollar4BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_percent_dollar_4_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedPercentDollar4GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_dquote_rbracket_5_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedDquoteRbracket5Parse(t *testing.T) {

	ccl := mock.New()
	input := `"gamma = nested524
]mode = deep57`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"gamma", Value: "nested524"}, mock.Entry{Key: "]mode", Value: "deep57"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_dquote_rbracket_5_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedDquoteRbracket5BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_dquote_rbracket_5_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedDquoteRbracket5GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_ampersand_6_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedLbraceAmpersand6Parse(t *testing.T) {

	ccl := mock.New()
	input := `{data = nested474
&port = deep22`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "{data", Value: "nested474"}, mock.Entry{Key: "&port", Value: "deep22"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_lbrace_ampersand_6_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedLbraceAmpersand6BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_ampersand_6_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedLbraceAmpersand6GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_semicolon_hyphen_7_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedSemicolonHyphen7Parse(t *testing.T) {

	ccl := mock.New()
	input := `;data = nested239
-server = deep284`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ";data", Value: "nested239"}, mock.Entry{Key: "-server", Value: "deep284"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_semicolon_hyphen_7_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedSemicolonHyphen7BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_semicolon_hyphen_7_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedSemicolonHyphen7GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_rbrace_8_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedLbraceRbrace8Parse(t *testing.T) {

	ccl := mock.New()
	input := `{host = nested125
}user = deep615`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "{host", Value: "nested125"}, mock.Entry{Key: "}user", Value: "deep615"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_lbrace_rbrace_8_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedLbraceRbrace8BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_lbrace_rbrace_8_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedLbraceRbrace8GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_tilde_ampersand_9_parse - function:parse feature:optional_typed_accessors
func TestFuzzNestedTildeAmpersand9Parse(t *testing.T) {

	ccl := mock.New()
	input := `~beta = nested835
&gamma = deep966`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "~beta", Value: "nested835"}, mock.Entry{Key: "&gamma", Value: "deep966"}}
	assert.Equal(t, expected, parseResult)

}

// fuzz_nested_tilde_ampersand_9_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestFuzzNestedTildeAmpersand9BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// fuzz_nested_tilde_ampersand_9_get_string - function:get_string feature:optional_typed_accessors
func TestFuzzNestedTildeAmpersand9GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
