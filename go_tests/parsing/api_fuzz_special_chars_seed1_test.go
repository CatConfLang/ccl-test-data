package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_fuzz_special_chars_seed1.json
// Suite: Flat Format
// Version: 1.0



// s1_fuzz_single_squote_parse - function:parse
func TestS1FuzzSingleSquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `' = val406`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "'", Value: "val406"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_asterisk_parse - function:parse
func TestS1FuzzSingleAsteriskParse(t *testing.T) {
	

	ccl := mock.New()
	input := `* = val236`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "*", Value: "val236"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_backslash_parse - function:parse
func TestS1FuzzSingleBackslashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `\ = val775`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\\", Value: "val775"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_rbrace_parse - function:parse
func TestS1FuzzSingleRbraceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `} = val221`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}", Value: "val221"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_tilde_parse - function:parse
func TestS1FuzzSingleTildeParse(t *testing.T) {
	

	ccl := mock.New()
	input := `~ = val416`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "~", Value: "val416"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_semicolon_parse - function:parse
func TestS1FuzzSingleSemicolonParse(t *testing.T) {
	

	ccl := mock.New()
	input := `; = val155`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ";", Value: "val155"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_plus_parse - function:parse
func TestS1FuzzSinglePlusParse(t *testing.T) {
	

	ccl := mock.New()
	input := `+ = val794`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "+", Value: "val794"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_dollar_parse - function:parse
func TestS1FuzzSingleDollarParse(t *testing.T) {
	

	ccl := mock.New()
	input := `$ = val665`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$", Value: "val665"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_percent_parse - function:parse
func TestS1FuzzSinglePercentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `% = val973`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%", Value: "val973"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_at_parse - function:parse
func TestS1FuzzSingleAtParse(t *testing.T) {
	

	ccl := mock.New()
	input := `@ = val410`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "@", Value: "val410"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_dquote_parse - function:parse
func TestS1FuzzSingleDquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `" = val797`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"", Value: "val797"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_lparen_parse - function:parse
func TestS1FuzzSingleLparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `( = val364`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "(", Value: "val364"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_lt_parse - function:parse
func TestS1FuzzSingleLtParse(t *testing.T) {
	

	ccl := mock.New()
	input := `< = val789`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "<", Value: "val789"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_lbracket_parse - function:parse
func TestS1FuzzSingleLbracketParse(t *testing.T) {
	

	ccl := mock.New()
	input := `[ = val216`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "[", Value: "val216"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_single_rparen_parse - function:parse
func TestS1FuzzSingleRparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `) = val374`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")", Value: "val374"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_at_lt_slash_parse - function:parse
func TestS1FuzzComboAtLtSlashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `@</ = combo594`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "@</", Value: "combo594"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_slash_lparen_semicolon_backslash_parse - function:parse
func TestS1FuzzComboSlashLparenSemicolonBackslashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `/(;\ = combo220`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/(;\\", Value: "combo220"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_lt_hash_parse - function:parse
func TestS1FuzzComboLtHashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `<# = combo31`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "<#", Value: "combo31"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_lparen_dquote_bang_parse - function:parse
func TestS1FuzzComboLparenDquoteBangParse(t *testing.T) {
	

	ccl := mock.New()
	input := `("! = combo359`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "(\"!", Value: "combo359"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_slash_hash_backslash_dquote_parse - function:parse
func TestS1FuzzComboSlashHashBackslashDquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `/#\" = combo87`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/#\\\"", Value: "combo87"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_slash_percent_pipe_parse - function:parse
func TestS1FuzzComboSlashPercentPipeParse(t *testing.T) {
	

	ccl := mock.New()
	input := `/%| = combo957`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/%|", Value: "combo957"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_lbrace_pipe_backslash_parse - function:parse
func TestS1FuzzComboLbracePipeBackslashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `{|\ = combo828`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "{|\\", Value: "combo828"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_bang_percent_slash_parse - function:parse
func TestS1FuzzComboBangPercentSlashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `!%/ = combo382`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "!%/", Value: "combo382"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_lbrace_pipe_rparen_parse - function:parse
func TestS1FuzzComboLbracePipeRparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `{|) = combo608`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "{|)", Value: "combo608"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_combo_colon_semicolon_squote_lt_parse - function:parse
func TestS1FuzzComboColonSemicolonSquoteLtParse(t *testing.T) {
	

	ccl := mock.New()
	input := `:;'< = combo488`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ":;'<", Value: "combo488"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_start_lbracket_host_parse - function:parse
func TestS1FuzzPosStartLbracketHostParse(t *testing.T) {
	

	ccl := mock.New()
	input := `[host = pos428`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "[host", Value: "pos428"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_middle_dollar_gamma_parse - function:parse
func TestS1FuzzPosMiddleDollarGammaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `ga$mma = pos138`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "ga$mma", Value: "pos138"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_end_pipe_name_parse - function:parse
func TestS1FuzzPosEndPipeNameParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name| = pos691`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name|", Value: "pos691"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_start_asterisk_server_parse - function:parse
func TestS1FuzzPosStartAsteriskServerParse(t *testing.T) {
	

	ccl := mock.New()
	input := `*server = pos50`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "*server", Value: "pos50"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_middle_gt_alpha_parse - function:parse
func TestS1FuzzPosMiddleGtAlphaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `al>pha = pos153`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "al>pha", Value: "pos153"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_end_squote_item_parse - function:parse
func TestS1FuzzPosEndSquoteItemParse(t *testing.T) {
	

	ccl := mock.New()
	input := `item' = pos133`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item'", Value: "pos133"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_start_dquote_mode_parse - function:parse
func TestS1FuzzPosStartDquoteModeParse(t *testing.T) {
	

	ccl := mock.New()
	input := `"mode = pos155`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"mode", Value: "pos155"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_middle_ampersand_port_parse - function:parse
func TestS1FuzzPosMiddleAmpersandPortParse(t *testing.T) {
	

	ccl := mock.New()
	input := `po&rt = pos601`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "po&rt", Value: "pos601"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_end_hash_host_parse - function:parse
func TestS1FuzzPosEndHashHostParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host# = pos976`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host#", Value: "pos976"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_pos_start_dollar_path_parse - function:parse
func TestS1FuzzPosStartDollarPathParse(t *testing.T) {
	

	ccl := mock.New()
	input := `$path = pos556`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$path", Value: "pos556"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_val_name_0_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzValName0Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = data?x77]x88`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "data?x77]x88"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_val_name_0_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzValName0BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_name_0_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzValName0BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_name_0_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzValName0GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_data_1_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzValData1Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `data = data'x23!x22`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "data", Value: "data'x23!x22"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_val_data_1_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzValData1BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_data_1_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzValData1BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_data_1_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzValData1GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_mode_2_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzValMode2Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `mode = data?x30"x15~x73`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "mode", Value: "data?x30\"x15~x73"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_val_mode_2_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzValMode2BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_mode_2_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzValMode2BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_mode_2_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzValMode2GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_name_3_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzValName3Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = data$x23~x81@x41`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "data$x23~x81@x41"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_val_name_3_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzValName3BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_name_3_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzValName3BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_name_3_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzValName3GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_user_4_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzValUser4Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `user = data}x85$x2$x43`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "user", Value: "data}x85$x2$x43"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_val_user_4_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzValUser4BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_user_4_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzValUser4BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_val_user_4_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzValUser4GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_ampersand_rbrace_0_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandRbrace0Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `&alpha = nested297
}gamma = deep715`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&alpha", Value: "nested297"}, mock.Entry{Key: "}gamma", Value: "deep715"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_ampersand_rbrace_0_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandRbrace0BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_ampersand_rbrace_0_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandRbrace0BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_ampersand_rbrace_0_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandRbrace0GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_question_rbracket_1_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedQuestionRbracket1Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `?port = nested105
]name = deep997`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "?port", Value: "nested105"}, mock.Entry{Key: "]name", Value: "deep997"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_question_rbracket_1_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedQuestionRbracket1BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_question_rbracket_1_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedQuestionRbracket1BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_question_rbracket_1_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedQuestionRbracket1GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_plus_slash_2_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedPlusSlash2Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `+config = nested427
/server = deep329`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "+config", Value: "nested427"}, mock.Entry{Key: "/server", Value: "deep329"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_plus_slash_2_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedPlusSlash2BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_plus_slash_2_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedPlusSlash2BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_plus_slash_2_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedPlusSlash2GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_ampersand_dollar_3_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandDollar3Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `&name = nested41
$port = deep486`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&name", Value: "nested41"}, mock.Entry{Key: "$port", Value: "deep486"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_ampersand_dollar_3_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandDollar3BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_ampersand_dollar_3_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandDollar3BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_ampersand_dollar_3_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedAmpersandDollar3GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_tilde_squote_4_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedTildeSquote4Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `~beta = nested332
'beta = deep158`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "~beta", Value: "nested332"}, mock.Entry{Key: "'beta", Value: "deep158"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_tilde_squote_4_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedTildeSquote4BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_tilde_squote_4_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedTildeSquote4BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_tilde_squote_4_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedTildeSquote4GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_hyphen_ampersand_5_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedHyphenAmpersand5Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `-config = nested723
&epsilon = deep891`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "-config", Value: "nested723"}, mock.Entry{Key: "&epsilon", Value: "deep891"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_hyphen_ampersand_5_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedHyphenAmpersand5BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_hyphen_ampersand_5_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedHyphenAmpersand5BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_hyphen_ampersand_5_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedHyphenAmpersand5GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_dollar_asterisk_6_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedDollarAsterisk6Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `$path = nested884
*mode = deep88`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$path", Value: "nested884"}, mock.Entry{Key: "*mode", Value: "deep88"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_dollar_asterisk_6_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedDollarAsterisk6BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_dollar_asterisk_6_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedDollarAsterisk6BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_dollar_asterisk_6_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedDollarAsterisk6GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_underscore_semicolon_7_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedUnderscoreSemicolon7Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `_path = nested515
;alpha = deep519`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "_path", Value: "nested515"}, mock.Entry{Key: ";alpha", Value: "deep519"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_underscore_semicolon_7_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedUnderscoreSemicolon7BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_underscore_semicolon_7_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedUnderscoreSemicolon7BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_underscore_semicolon_7_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedUnderscoreSemicolon7GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_at_rparen_8_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedAtRparen8Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `@server = nested286
)gamma = deep475`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "@server", Value: "nested286"}, mock.Entry{Key: ")gamma", Value: "deep475"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_at_rparen_8_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedAtRparen8BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_at_rparen_8_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedAtRparen8BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_at_rparen_8_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedAtRparen8GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_rbrace_tilde_9_parse - function:parse feature:optional_typed_accessors
func TestS1FuzzNestedRbraceTilde9Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `}beta = nested295
~alpha = deep415`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}beta", Value: "nested295"}, mock.Entry{Key: "~alpha", Value: "deep415"}}
	assert.Equal(t, expected, parseResult)

}


// s1_fuzz_nested_rbrace_tilde_9_build_hierarchy - function:parse_indented function:build_hierarchy feature:optional_typed_accessors
func TestS1FuzzNestedRbraceTilde9BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_rbrace_tilde_9_build_model - function:parse_indented function:build_model feature:optional_typed_accessors
func TestS1FuzzNestedRbraceTilde9BuildModel(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s1_fuzz_nested_rbrace_tilde_9_get_string - function:get_string feature:optional_typed_accessors
func TestS1FuzzNestedRbraceTilde9GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


