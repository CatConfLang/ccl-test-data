package parsing_test

import (
	"testing"
	
	"github.com/catconflang/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_fuzz_special_chars_seed99.json
// Suite: Flat Format
// Version: 1.0



// s99_fuzz_single_underscore_parse - function:parse
func TestS99FuzzSingleUnderscoreParse(t *testing.T) {
	

	ccl := mock.New()
	input := `_ = val35`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "_", Value: "val35"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_backslash_parse - function:parse
func TestS99FuzzSingleBackslashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `\ = val524`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\\", Value: "val524"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_percent_parse - function:parse
func TestS99FuzzSinglePercentParse(t *testing.T) {
	

	ccl := mock.New()
	input := `% = val196`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%", Value: "val196"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_dquote_parse - function:parse
func TestS99FuzzSingleDquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `" = val52`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"", Value: "val52"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_pipe_parse - function:parse
func TestS99FuzzSinglePipeParse(t *testing.T) {
	

	ccl := mock.New()
	input := `| = val653`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "|", Value: "val653"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_question_parse - function:parse
func TestS99FuzzSingleQuestionParse(t *testing.T) {
	

	ccl := mock.New()
	input := `? = val143`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "?", Value: "val143"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_rbrace_parse - function:parse
func TestS99FuzzSingleRbraceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `} = val623`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}", Value: "val623"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_hyphen_parse - function:parse
func TestS99FuzzSingleHyphenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `- = val645`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "-", Value: "val645"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_lparen_parse - function:parse
func TestS99FuzzSingleLparenParse(t *testing.T) {
	

	ccl := mock.New()
	input := `( = val220`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "(", Value: "val220"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_slash_parse - function:parse
func TestS99FuzzSingleSlashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `/ = val115`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "val115"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_bang_parse - function:parse
func TestS99FuzzSingleBangParse(t *testing.T) {
	

	ccl := mock.New()
	input := `! = val103`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "!", Value: "val103"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_hash_parse - function:parse
func TestS99FuzzSingleHashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `# = val874`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "#", Value: "val874"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_ampersand_parse - function:parse
func TestS99FuzzSingleAmpersandParse(t *testing.T) {
	

	ccl := mock.New()
	input := `& = val520`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&", Value: "val520"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_dollar_parse - function:parse
func TestS99FuzzSingleDollarParse(t *testing.T) {
	

	ccl := mock.New()
	input := `$ = val921`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "$", Value: "val921"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_single_colon_parse - function:parse
func TestS99FuzzSingleColonParse(t *testing.T) {
	

	ccl := mock.New()
	input := `: = val671`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ":", Value: "val671"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_lbracket_asterisk_gt_parse - function:parse
func TestS99FuzzComboLbracketAsteriskGtParse(t *testing.T) {
	

	ccl := mock.New()
	input := `[*> = combo828`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "[*>", Value: "combo828"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_percent_semicolon_squote_parse - function:parse
func TestS99FuzzComboPercentSemicolonSquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `%;' = combo983`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%;'", Value: "combo983"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_squote_underscore_rparen_squote_parse - function:parse
func TestS99FuzzComboSquoteUnderscoreRparenSquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `'_)' = combo857`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "'_)'", Value: "combo857"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_underscore_lbracket_lt_backslash_parse - function:parse
func TestS99FuzzComboUnderscoreLbracketLtBackslashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `_[<\ = combo602`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "_[<\\", Value: "combo602"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_ampersand_lparen_dquote_parse - function:parse
func TestS99FuzzComboAmpersandLparenDquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `&(" = combo48`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "&(\"", Value: "combo48"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_rparen_pipe_lparen_hash_parse - function:parse
func TestS99FuzzComboRparenPipeLparenHashParse(t *testing.T) {
	

	ccl := mock.New()
	input := `)|(# = combo531`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")|(#", Value: "combo531"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_tilde_asterisk_at_question_parse - function:parse
func TestS99FuzzComboTildeAsteriskAtQuestionParse(t *testing.T) {
	

	ccl := mock.New()
	input := `~*@? = combo716`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "~*@?", Value: "combo716"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_underscore_asterisk_question_parse - function:parse
func TestS99FuzzComboUnderscoreAsteriskQuestionParse(t *testing.T) {
	

	ccl := mock.New()
	input := `_*? = combo909`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "_*?", Value: "combo909"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_bang_squote_parse - function:parse
func TestS99FuzzComboBangSquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `!' = combo182`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "!'", Value: "combo182"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_combo_question_underscore_dquote_parse - function:parse
func TestS99FuzzComboQuestionUnderscoreDquoteParse(t *testing.T) {
	

	ccl := mock.New()
	input := `?_" = combo211`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "?_\"", Value: "combo211"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_start_percent_item_parse - function:parse
func TestS99FuzzPosStartPercentItemParse(t *testing.T) {
	

	ccl := mock.New()
	input := `%item = pos832`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%item", Value: "pos832"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_middle_rbracket_data_parse - function:parse
func TestS99FuzzPosMiddleRbracketDataParse(t *testing.T) {
	

	ccl := mock.New()
	input := `da]ta = pos418`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "da]ta", Value: "pos418"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_end_gt_delta_parse - function:parse
func TestS99FuzzPosEndGtDeltaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `delta> = pos239`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "delta>", Value: "pos239"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_start_rparen_port_parse - function:parse
func TestS99FuzzPosStartRparenPortParse(t *testing.T) {
	

	ccl := mock.New()
	input := `)port = pos895`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")port", Value: "pos895"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_middle_underscore_alpha_parse - function:parse
func TestS99FuzzPosMiddleUnderscoreAlphaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `al_pha = pos195`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "al_pha", Value: "pos195"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_end_lbrace_epsilon_parse - function:parse
func TestS99FuzzPosEndLbraceEpsilonParse(t *testing.T) {
	

	ccl := mock.New()
	input := `epsilon{ = pos824`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "epsilon{", Value: "pos824"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_start_lt_alpha_parse - function:parse
func TestS99FuzzPosStartLtAlphaParse(t *testing.T) {
	

	ccl := mock.New()
	input := `<alpha = pos633`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "<alpha", Value: "pos633"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_middle_percent_host_parse - function:parse
func TestS99FuzzPosMiddlePercentHostParse(t *testing.T) {
	

	ccl := mock.New()
	input := `ho%st = pos14`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "ho%st", Value: "pos14"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_end_pipe_item_parse - function:parse
func TestS99FuzzPosEndPipeItemParse(t *testing.T) {
	

	ccl := mock.New()
	input := `item| = pos19`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item|", Value: "pos19"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_pos_start_percent_server_parse - function:parse
func TestS99FuzzPosStartPercentServerParse(t *testing.T) {
	

	ccl := mock.New()
	input := `%server = pos981`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%server", Value: "pos981"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_val_path_0_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzValPath0Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `path = data'x66`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "path", Value: "data'x66"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_val_path_0_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzValPath0BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `path = data'x66`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"path": "data'x66"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_val_path_0_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzValPath0GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `path = data'x66`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"path"})
	require.NoError(t, err)
	assert.Equal(t, "data'x66", result)

}


// s99_fuzz_val_item_1_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzValItem1Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `item = data]x41!x78`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "item", Value: "data]x41!x78"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_val_item_1_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzValItem1BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `item = data]x41!x78`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"item": "data]x41!x78"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_val_item_1_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzValItem1GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `item = data]x41!x78`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"item"})
	require.NoError(t, err)
	assert.Equal(t, "data]x41!x78", result)

}


// s99_fuzz_val_server_2_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzValServer2Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `server = data[x73;x54!x24`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "server", Value: "data[x73;x54!x24"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_val_server_2_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzValServer2BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `server = data[x73;x54!x24`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"server": "data[x73;x54!x24"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_val_server_2_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzValServer2GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `server = data[x73;x54!x24`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"server"})
	require.NoError(t, err)
	assert.Equal(t, "data[x73;x54!x24", result)

}


// s99_fuzz_val_path_3_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzValPath3Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `path = data~x4>x82~x46`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "path", Value: "data~x4>x82~x46"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_val_path_3_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzValPath3BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `path = data~x4>x82~x46`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"path": "data~x4>x82~x46"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_val_path_3_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzValPath3GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `path = data~x4>x82~x46`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"path"})
	require.NoError(t, err)
	assert.Equal(t, "data~x4>x82~x46", result)

}


// s99_fuzz_val_alpha_4_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzValAlpha4Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `alpha = data<x11[x43`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "alpha", Value: "data<x11[x43"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_val_alpha_4_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzValAlpha4BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `alpha = data<x11[x43`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"alpha": "data<x11[x43"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_val_alpha_4_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzValAlpha4GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `alpha = data<x11[x43`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"alpha"})
	require.NoError(t, err)
	assert.Equal(t, "data<x11[x43", result)

}


// s99_fuzz_nested_slash_rparen_0_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedSlashRparen0Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `/data = nested929
)item = deep37`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/data", Value: "nested929"}, mock.Entry{Key: ")item", Value: "deep37"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_slash_rparen_0_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedSlashRparen0BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `/data = nested929
)item = deep37`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{")item": "deep37", "/data": "nested929"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_slash_rparen_0_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedSlashRparen0GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `/data = nested929
)item = deep37`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"/data"})
	require.NoError(t, err)
	assert.Equal(t, "nested929", result)

}


// s99_fuzz_nested_lt_percent_1_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedLtPercent1Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `<port = nested722
%config = deep795`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "<port", Value: "nested722"}, mock.Entry{Key: "%config", Value: "deep795"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_lt_percent_1_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedLtPercent1BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `<port = nested722
%config = deep795`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"%config": "deep795", "<port": "nested722"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_lt_percent_1_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedLtPercent1GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `<port = nested722
%config = deep795`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"<port"})
	require.NoError(t, err)
	assert.Equal(t, "nested722", result)

}


// s99_fuzz_nested_slash_lbrace_2_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedSlashLbrace2Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `/epsilon = nested268
{server = deep795`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/epsilon", Value: "nested268"}, mock.Entry{Key: "{server", Value: "deep795"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_slash_lbrace_2_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedSlashLbrace2BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `/epsilon = nested268
{server = deep795`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"/epsilon": "nested268", "{server": "deep795"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_slash_lbrace_2_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedSlashLbrace2GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `/epsilon = nested268
{server = deep795`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"/epsilon"})
	require.NoError(t, err)
	assert.Equal(t, "nested268", result)

}


// s99_fuzz_nested_dquote_lbracket_3_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedDquoteLbracket3Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `"host = nested435
[server = deep479`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"host", Value: "nested435"}, mock.Entry{Key: "[server", Value: "deep479"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_dquote_lbracket_3_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedDquoteLbracket3BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `"host = nested435
[server = deep479`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"\"host": "nested435", "[server": "deep479"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_dquote_lbracket_3_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedDquoteLbracket3GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `"host = nested435
[server = deep479`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"\"host"})
	require.NoError(t, err)
	assert.Equal(t, "nested435", result)

}


// s99_fuzz_nested_backslash_backslash_4_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedBackslashBackslash4Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `\item = nested708
\name = deep133`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\\item", Value: "nested708"}, mock.Entry{Key: "\\name", Value: "deep133"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_backslash_backslash_4_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedBackslashBackslash4BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `\item = nested708
\name = deep133`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"\\item": "nested708", "\\name": "deep133"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_backslash_backslash_4_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedBackslashBackslash4GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `\item = nested708
\name = deep133`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"\\item"})
	require.NoError(t, err)
	assert.Equal(t, "nested708", result)

}


// s99_fuzz_nested_question_dquote_5_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedQuestionDquote5Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `?delta = nested18
"item = deep576`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "?delta", Value: "nested18"}, mock.Entry{Key: "\"item", Value: "deep576"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_question_dquote_5_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedQuestionDquote5BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `?delta = nested18
"item = deep576`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"\"item": "deep576", "?delta": "nested18"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_question_dquote_5_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedQuestionDquote5GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `?delta = nested18
"item = deep576`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"?delta"})
	require.NoError(t, err)
	assert.Equal(t, "nested18", result)

}


// s99_fuzz_nested_dquote_backslash_6_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedDquoteBackslash6Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `"user = nested759
\data = deep718`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "\"user", Value: "nested759"}, mock.Entry{Key: "\\data", Value: "deep718"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_dquote_backslash_6_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedDquoteBackslash6BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `"user = nested759
\data = deep718`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"\"user": "nested759", "\\data": "deep718"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_dquote_backslash_6_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedDquoteBackslash6GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `"user = nested759
\data = deep718`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"\"user"})
	require.NoError(t, err)
	assert.Equal(t, "nested759", result)

}


// s99_fuzz_nested_rparen_backslash_7_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedRparenBackslash7Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `)port = nested292
\server = deep527`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: ")port", Value: "nested292"}, mock.Entry{Key: "\\server", Value: "deep527"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_rparen_backslash_7_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedRparenBackslash7BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `)port = nested292
\server = deep527`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{")port": "nested292", "\\server": "deep527"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_rparen_backslash_7_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedRparenBackslash7GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `)port = nested292
\server = deep527`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{")port"})
	require.NoError(t, err)
	assert.Equal(t, "nested292", result)

}


// s99_fuzz_nested_rbrace_dollar_8_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedRbraceDollar8Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `}delta = nested591
$gamma = deep225`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "}delta", Value: "nested591"}, mock.Entry{Key: "$gamma", Value: "deep225"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_rbrace_dollar_8_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedRbraceDollar8BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `}delta = nested591
$gamma = deep225`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"$gamma": "deep225", "}delta": "nested591"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_rbrace_dollar_8_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedRbraceDollar8GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `}delta = nested591
$gamma = deep225`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"}delta"})
	require.NoError(t, err)
	assert.Equal(t, "nested591", result)

}


// s99_fuzz_nested_percent_ampersand_9_parse - function:parse feature:optional_typed_accessors
func TestS99FuzzNestedPercentAmpersand9Parse(t *testing.T) {
	

	ccl := mock.New()
	input := `%epsilon = nested761
&port = deep211`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "%epsilon", Value: "nested761"}, mock.Entry{Key: "&port", Value: "deep211"}}
	assert.Equal(t, expected, parseResult)

}


// s99_fuzz_nested_percent_ampersand_9_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestS99FuzzNestedPercentAmpersand9BuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `%epsilon = nested761
&port = deep211`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"%epsilon": "nested761", "&port": "deep211"}
	assert.Equal(t, expected, objectResult)

}


// s99_fuzz_nested_percent_ampersand_9_get_string - function:get_string feature:optional_typed_accessors
func TestS99FuzzNestedPercentAmpersand9GetString(t *testing.T) {
	

	ccl := mock.New()
	input := `%epsilon = nested761
&port = deep211`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"%epsilon"})
	require.NoError(t, err)
	assert.Equal(t, "nested761", result)

}


