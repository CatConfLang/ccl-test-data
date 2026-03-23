package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_fuzz_special_chars_seed49.json
// Suite: Flat Format
// Version: 1.0



// s49_fuzz_single_lparen_parse - function:parse
func TestS49FuzzSingleLparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `( = val911`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "(", Value: "val911"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_pipe_parse - function:parse
func TestS49FuzzSinglePipeParse(t *testing.T) {
	

	ccl := mock.New()
	input := `| = val79`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "|", Value: "val79"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_percent_parse - function:parse
func TestS49FuzzSinglePercentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `% = val935`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%", Value: "val935"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_asterisk_parse - function:parse
func TestS49FuzzSingleAsteriskParse(t *testing.T) {
	

	ccl := mock.New()
	input := `* = val545`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "*", Value: "val545"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_backslash_parse - function:parse
func TestS49FuzzSingleBackslashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `\ = val748`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\\", Value: "val748"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_slash_parse - function:parse
func TestS49FuzzSingleSlashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `/ = val718`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "val718"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_rbracket_parse - function:parse
func TestS49FuzzSingleRbracketParse(t *testing.T) {
	

	ccl := mock.New()
	input := `] = val858`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "]", Value: "val858"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_ampersand_parse - function:parse
func TestS49FuzzSingleAmpersandParse(t *testing.T) {
	

	ccl := mock.New()
	input := `& = val116`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&", Value: "val116"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_at_parse - function:parse
func TestS49FuzzSingleAtParse(t *testing.T) {
	

	ccl := mock.New()
	input := `@ = val518`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "@", Value: "val518"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_dquote_parse - function:parse
func TestS49FuzzSingleDquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `" = val757`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"", Value: "val757"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_dollar_parse - function:parse
func TestS49FuzzSingleDollarParse(t *testing.T) {
	

	ccl := mock.New()
	input := `$ = val874`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$", Value: "val874"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_bang_parse - function:parse
func TestS49FuzzSingleBangParse(t *testing.T) {
	

	ccl := mock.New()
	input := `! = val221`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "!", Value: "val221"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_rparen_parse - function:parse
func TestS49FuzzSingleRparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `) = val331`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")", Value: "val331"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_rbrace_parse - function:parse
func TestS49FuzzSingleRbraceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `} = val9`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}", Value: "val9"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_single_tilde_parse - function:parse
func TestS49FuzzSingleTildeParse(t *testing.T) {
	

	ccl := mock.New()
	input := `~ = val714`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "~", Value: "val714"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_lparen_hash_asterisk_parse - function:parse
func TestS49FuzzComboLparenHashAsteriskParse(t *testing.T) {
	

	ccl := mock.New()
	input := `(#* = combo926`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "(#*", Value: "combo926"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_rbrace_slash_underscore_rbrace_parse - function:parse
func TestS49FuzzComboRbraceSlashUnderscoreRbraceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `}/_} = combo173`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}/_}", Value: "combo173"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_rbracket_backslash_percent_at_parse - function:parse
func TestS49FuzzComboRbracketBackslashPercentAtParse(t *testing.T) {
	

	ccl := mock.New()
	input := `]\%@ = combo764`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "]\\%@", Value: "combo764"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_rbracket_lbrace_squote_lbracket_parse - function:parse
func TestS49FuzzComboRbracketLbraceSquoteLbracketParse(t *testing.T) {
	

	ccl := mock.New()
	input := `]{'[ = combo827`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "]{'[", Value: "combo827"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_colon_rparen_parse - function:parse
func TestS49FuzzComboColonRparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `:) = combo140`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ":)", Value: "combo140"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_gt_tilde_rparen_parse - function:parse
func TestS49FuzzComboGtTildeRparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `>~) = combo274`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ">~)", Value: "combo274"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_semicolon_gt_colon_parse - function:parse
func TestS49FuzzComboSemicolonGtColonParse(t *testing.T) {
	

	ccl := mock.New()
	input := `;>: = combo111`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ";>:", Value: "combo111"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_question_plus_backslash_parse - function:parse
func TestS49FuzzComboQuestionPlusBackslashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `?+\ = combo711`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "?+\\", Value: "combo711"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_ampersand_ampersand_parse - function:parse
func TestS49FuzzComboAmpersandAmpersandParse(t *testing.T) {
	

	ccl := mock.New()
	input := `&& = combo102`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&&", Value: "combo102"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_combo_tilde_ampersand_tilde_parse - function:parse
func TestS49FuzzComboTildeAmpersandTildeParse(t *testing.T) {
	

	ccl := mock.New()
	input := `~&~ = combo549`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "~&~", Value: "combo549"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_start_hash_host_parse - function:parse
func TestS49FuzzPosStartHashHostParse(t *testing.T) {
	

	ccl := mock.New()
	input := `#host = pos568`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "#host", Value: "pos568"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_middle_slash_beta_parse - function:parse
func TestS49FuzzPosMiddleSlashBetaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `be/ta = pos975`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "be/ta", Value: "pos975"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_end_lparen_config_parse - function:parse
func TestS49FuzzPosEndLparenConfigParse(t *testing.T) {
	

	ccl := mock.New()
	input := `config( = pos258`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config(", Value: "pos258"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_start_squote_data_parse - function:parse
func TestS49FuzzPosStartSquoteDataParse(t *testing.T) {
	

	ccl := mock.New()
	input := `'data = pos985`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "'data", Value: "pos985"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_middle_rbracket_alpha_parse - function:parse
func TestS49FuzzPosMiddleRbracketAlphaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `al]pha = pos844`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "al]pha", Value: "pos844"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_end_percent_name_parse - function:parse
func TestS49FuzzPosEndPercentNameParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name% = pos392`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name%", Value: "pos392"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_start_rparen_user_parse - function:parse
func TestS49FuzzPosStartRparenUserParse(t *testing.T) {
	

	ccl := mock.New()
	input := `)user = pos217`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")user", Value: "pos217"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_middle_lbracket_host_parse - function:parse
func TestS49FuzzPosMiddleLbracketHostParse(t *testing.T) {
	

	ccl := mock.New()
	input := `ho[st = pos439`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "ho[st", Value: "pos439"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_end_tilde_beta_parse - function:parse
func TestS49FuzzPosEndTildeBetaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `beta~ = pos557`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "beta~", Value: "pos557"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_pos_start_dollar_config_parse - function:parse
func TestS49FuzzPosStartDollarConfigParse(t *testing.T) {
	

	ccl := mock.New()
	input := `$config = pos46`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$config", Value: "pos46"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_val_epsilon_0_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzValEpsilon0Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `epsilon = data;x5&x96?x73`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "epsilon", Value: "data;x5&x96?x73"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_val_epsilon_0_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzValEpsilon0BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_epsilon_0_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzValEpsilon0GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_config_1_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzValConfig1Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `config = data'x33`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "data'x33"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_val_config_1_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzValConfig1BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_config_1_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzValConfig1GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_name_2_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzValName2Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = data<x71`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "data<x71"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_val_name_2_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzValName2BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_name_2_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzValName2GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_mode_3_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzValMode3Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `mode = data+x19`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "mode", Value: "data+x19"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_val_mode_3_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzValMode3BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_mode_3_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzValMode3GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_server_4_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzValServer4Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `server = data#x70-x10`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "server", Value: "data#x70-x10"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_val_server_4_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzValServer4BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_val_server_4_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzValServer4GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_squote_lbrace_0_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedSquoteLbrace0Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `'name = nested330
{user = deep34`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "'name", Value: "nested330"}, mock.Entry{Key: "{user", Value: "deep34"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_squote_lbrace_0_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedSquoteLbrace0BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_squote_lbrace_0_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedSquoteLbrace0GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_at_underscore_1_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedAtUnderscore1Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `@name = nested517
_user = deep650`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "@name", Value: "nested517"}, mock.Entry{Key: "_user", Value: "deep650"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_at_underscore_1_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedAtUnderscore1BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_at_underscore_1_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedAtUnderscore1GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_rparen_backslash_2_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedRparenBackslash2Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `)epsilon = nested451
\delta = deep648`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")epsilon", Value: "nested451"}, mock.Entry{Key: "\\delta", Value: "deep648"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_rparen_backslash_2_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedRparenBackslash2BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_rparen_backslash_2_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedRparenBackslash2GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_dollar_gt_3_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedDollarGt3Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `$port = nested76
>host = deep444`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$port", Value: "nested76"}, mock.Entry{Key: ">host", Value: "deep444"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_dollar_gt_3_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedDollarGt3BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_dollar_gt_3_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedDollarGt3GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_dquote_lbrace_4_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedDquoteLbrace4Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `"host = nested587
{gamma = deep774`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"host", Value: "nested587"}, mock.Entry{Key: "{gamma", Value: "deep774"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_dquote_lbrace_4_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedDquoteLbrace4BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_dquote_lbrace_4_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedDquoteLbrace4GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_semicolon_asterisk_5_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedSemicolonAsterisk5Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `;host = nested246
*item = deep424`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ";host", Value: "nested246"}, mock.Entry{Key: "*item", Value: "deep424"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_semicolon_asterisk_5_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedSemicolonAsterisk5BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_semicolon_asterisk_5_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedSemicolonAsterisk5GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_question_asterisk_6_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedQuestionAsterisk6Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `?epsilon = nested334
*delta = deep689`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "?epsilon", Value: "nested334"}, mock.Entry{Key: "*delta", Value: "deep689"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_question_asterisk_6_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedQuestionAsterisk6BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_question_asterisk_6_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedQuestionAsterisk6GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_plus_ampersand_7_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedPlusAmpersand7Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `+server = nested496
&gamma = deep440`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "+server", Value: "nested496"}, mock.Entry{Key: "&gamma", Value: "deep440"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_plus_ampersand_7_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedPlusAmpersand7BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_plus_ampersand_7_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedPlusAmpersand7GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_asterisk_dquote_8_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedAsteriskDquote8Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `*item = nested578
"delta = deep232`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "*item", Value: "nested578"}, mock.Entry{Key: "\"delta", Value: "deep232"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_asterisk_dquote_8_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedAsteriskDquote8BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_asterisk_dquote_8_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedAsteriskDquote8GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_rbrace_slash_9_parse - function:parse feature:optional_typed_accessors
func TestS49FuzzNestedRbraceSlash9Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `}host = nested668
/delta = deep756`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}host", Value: "nested668"}, mock.Entry{Key: "/delta", Value: "deep756"}}
	assert.Equal(t, expected, parseResult)

}


// s49_fuzz_nested_rbrace_slash_9_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS49FuzzNestedRbraceSlash9BuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// s49_fuzz_nested_rbrace_slash_9_get_string - function:get_string feature:optional_typed_accessors
func TestS49FuzzNestedRbraceSlash9GetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


