package parsing_test

import (
	"testing"
	
	"github.com/tylerbutler/ccl-test-data/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_typed_access.json
// Suite: Flat Format
// Version: 1.0



// parse_basic_integer_parse - function:parse feature:optional_typed_accessors
func TestParseBasicIntegerParse(t *testing.T) {
	

	ccl := mock.New()
	input := `port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "port", Value: "8080"}}
	assert.Equal(t, expected, parseResult)

}


// parse_basic_integer_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBasicIntegerBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"port": "8080"}
	assert.Equal(t, expected, objectResult)

}


// parse_basic_integer_get_int - function:get_int feature:optional_typed_accessors
func TestParseBasicIntegerGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `port = 8080`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"port"})
	require.NoError(t, err)
	assert.Equal(t, 8080, result)

}


// parse_basic_float_parse - function:parse feature:optional_typed_accessors
func TestParseBasicFloatParse(t *testing.T) {
	

	ccl := mock.New()
	input := `temperature = 98.6`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "temperature", Value: "98.6"}}
	assert.Equal(t, expected, parseResult)

}


// parse_basic_float_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBasicFloatBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `temperature = 98.6`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"temperature": "98.6"}
	assert.Equal(t, expected, objectResult)

}


// parse_basic_float_get_float - function:get_float feature:optional_typed_accessors
func TestParseBasicFloatGetFloat(t *testing.T) {
	

	ccl := mock.New()
	input := `temperature = 98.6`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_float validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetFloat(hierarchy, []string{"temperature"})
	require.NoError(t, err)
	assert.Equal(t, 98.6, result)

}


// parse_boolean_true_parse - function:parse feature:optional_typed_accessors
func TestParseBooleanTrueParse(t *testing.T) {
	

	ccl := mock.New()
	input := `enabled = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "enabled", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_true_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBooleanTrueBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `enabled = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"enabled": "true"}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_true_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict behavior:boolean_lenient
func TestParseBooleanTrueGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `enabled = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"enabled"})
	require.NoError(t, err)
	assert.Equal(t, true, result)

}


// parse_boolean_yes_parse - function:parse feature:optional_typed_accessors
func TestParseBooleanYesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `active = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "active", Value: "yes"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_yes_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBooleanYesBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `active = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"active": "yes"}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_yes_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_lenient
func TestParseBooleanYesGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `active = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"active"})
	require.NoError(t, err)
	assert.Equal(t, true, result)

}


// parse_boolean_yes_strict_literal_parse - function:parse feature:optional_typed_accessors
func TestParseBooleanYesStrictLiteralParse(t *testing.T) {
	

	ccl := mock.New()
	input := `active = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "active", Value: "yes"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_yes_strict_literal_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBooleanYesStrictLiteralBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `active = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"active": "yes"}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_yes_strict_literal_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseBooleanYesStrictLiteralGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `active = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"active"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// parse_boolean_false_parse - function:parse feature:optional_typed_accessors
func TestParseBooleanFalseParse(t *testing.T) {
	

	ccl := mock.New()
	input := `disabled = false`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "disabled", Value: "false"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_false_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBooleanFalseBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `disabled = false`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"disabled": "false"}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_false_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict behavior:boolean_lenient
func TestParseBooleanFalseGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `disabled = false`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"disabled"})
	require.NoError(t, err)
	assert.Equal(t, false, result)

}


// parse_string_fallback_parse - function:parse
func TestParseStringFallbackParse(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}}
	assert.Equal(t, expected, parseResult)

}


// parse_string_fallback_build_hierarchy - function:build_hierarchy
func TestParseStringFallbackBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"name": "Alice"}
	assert.Equal(t, expected, objectResult)

}


// parse_string_fallback_get_string - function:get_string
func TestParseStringFallbackGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `name = Alice`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"name"})
	require.NoError(t, err)
	assert.Equal(t, "Alice", result)

}


// parse_negative_integer_parse - function:parse feature:optional_typed_accessors
func TestParseNegativeIntegerParse(t *testing.T) {
	

	ccl := mock.New()
	input := `offset = -42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "offset", Value: "-42"}}
	assert.Equal(t, expected, parseResult)

}


// parse_negative_integer_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseNegativeIntegerBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `offset = -42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"offset": "-42"}
	assert.Equal(t, expected, objectResult)

}


// parse_negative_integer_get_int - function:get_int feature:optional_typed_accessors
func TestParseNegativeIntegerGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `offset = -42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"offset"})
	require.NoError(t, err)
	assert.Equal(t, -42, result)

}


// parse_zero_values_parse - function:parse feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "count", Value: "0"}, mock.Entry{Key: "distance", Value: "0.0"}, mock.Entry{Key: "disabled", Value: "no"}}
	assert.Equal(t, expected, parseResult)

}


// parse_zero_values_build_hierarchy - function:build_hierarchy feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"count": "0", "disabled": "no", "distance": "0.0"}
	assert.Equal(t, expected, objectResult)

}


// parse_zero_values_get_int - function:get_int feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"count"})
	require.NoError(t, err)
	assert.Equal(t, 0, result)

}


// parse_zero_values_get_bool - function:get_bool feature:empty_keys feature:optional_typed_accessors behavior:boolean_lenient
func TestParseZeroValuesGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"disabled"})
	require.NoError(t, err)
	assert.Equal(t, false, result)

}


// parse_zero_values_get_float - function:get_float feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesGetFloat(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_float validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetFloat(hierarchy, []string{"distance"})
	require.NoError(t, err)
	assert.Equal(t, 0, result)

}


// parse_zero_values_strict_literal_parse - function:parse feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesStrictLiteralParse(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "count", Value: "0"}, mock.Entry{Key: "distance", Value: "0.0"}, mock.Entry{Key: "disabled", Value: "no"}}
	assert.Equal(t, expected, parseResult)

}


// parse_zero_values_strict_literal_build_hierarchy - function:build_hierarchy feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesStrictLiteralBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"count": "0", "disabled": "no", "distance": "0.0"}
	assert.Equal(t, expected, objectResult)

}


// parse_zero_values_strict_literal_get_int - function:get_int feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesStrictLiteralGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"count"})
	require.NoError(t, err)
	assert.Equal(t, 0, result)

}


// parse_zero_values_strict_literal_get_bool - function:get_bool feature:empty_keys feature:optional_typed_accessors behavior:boolean_strict
func TestParseZeroValuesStrictLiteralGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"disabled"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// parse_zero_values_strict_literal_get_float - function:get_float feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesStrictLiteralGetFloat(t *testing.T) {
	

	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_float validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetFloat(hierarchy, []string{"distance"})
	require.NoError(t, err)
	assert.Equal(t, 0, result)

}


// parse_boolean_variants_parse - function:parse feature:optional_typed_accessors
func TestParseBooleanVariantsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "flag1", Value: "yes"}, mock.Entry{Key: "flag2", Value: "on"}, mock.Entry{Key: "flag3", Value: "1"}, mock.Entry{Key: "flag4", Value: "false"}, mock.Entry{Key: "flag5", Value: "no"}, mock.Entry{Key: "flag6", Value: "off"}, mock.Entry{Key: "flag7", Value: "0"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_variants_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBooleanVariantsBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"flag1": "yes", "flag2": "on", "flag3": "1", "flag4": "false", "flag5": "no", "flag6": "off", "flag7": "0"}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_variants_get_int - function:get_int feature:optional_typed_accessors
func TestParseBooleanVariantsGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"flag3"})
	require.NoError(t, err)
	assert.Equal(t, 1, result)

}


// parse_boolean_variants_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_lenient
func TestParseBooleanVariantsGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"flag1"})
	require.NoError(t, err)
	assert.Equal(t, true, result)

}


// parse_boolean_variants_strict_literal_parse - function:parse feature:optional_typed_accessors
func TestParseBooleanVariantsStrictLiteralParse(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "flag1", Value: "yes"}, mock.Entry{Key: "flag2", Value: "on"}, mock.Entry{Key: "flag3", Value: "1"}, mock.Entry{Key: "flag4", Value: "false"}, mock.Entry{Key: "flag5", Value: "no"}, mock.Entry{Key: "flag6", Value: "off"}, mock.Entry{Key: "flag7", Value: "0"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_variants_strict_literal_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBooleanVariantsStrictLiteralBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"flag1": "yes", "flag2": "on", "flag3": "1", "flag4": "false", "flag5": "no", "flag6": "off", "flag7": "0"}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_variants_strict_literal_get_int - function:get_int feature:optional_typed_accessors
func TestParseBooleanVariantsStrictLiteralGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"flag3"})
	require.NoError(t, err)
	assert.Equal(t, 1, result)

}


// parse_boolean_variants_strict_literal_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseBooleanVariantsStrictLiteralGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"flag1"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// parse_mixed_types_parse - function:parse feature:optional_typed_accessors
func TestParseMixedTypesParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "ssl", Value: "true"}, mock.Entry{Key: "timeout", Value: "30.5"}, mock.Entry{Key: "debug", Value: "off"}}
	assert.Equal(t, expected, parseResult)

}


// parse_mixed_types_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseMixedTypesBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"debug": "off", "host": "localhost", "port": "8080", "ssl": "true", "timeout": "30.5"}
	assert.Equal(t, expected, objectResult)

}


// parse_mixed_types_get_string - function:get_string feature:optional_typed_accessors
func TestParseMixedTypesGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"host"})
	require.NoError(t, err)
	assert.Equal(t, "localhost", result)

}


// parse_mixed_types_get_int - function:get_int feature:optional_typed_accessors
func TestParseMixedTypesGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"port"})
	require.NoError(t, err)
	assert.Equal(t, 8080, result)

}


// parse_mixed_types_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_lenient
func TestParseMixedTypesGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"ssl"})
	require.NoError(t, err)
	assert.Equal(t, true, result)

}


// parse_mixed_types_get_float - function:get_float feature:optional_typed_accessors
func TestParseMixedTypesGetFloat(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_float validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetFloat(hierarchy, []string{"timeout"})
	require.NoError(t, err)
	assert.Equal(t, 30.5, result)

}


// parse_mixed_types_strict_literal_parse - function:parse feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralParse(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "ssl", Value: "true"}, mock.Entry{Key: "timeout", Value: "30.5"}, mock.Entry{Key: "debug", Value: "off"}}
	assert.Equal(t, expected, parseResult)

}


// parse_mixed_types_strict_literal_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"debug": "off", "host": "localhost", "port": "8080", "ssl": "true", "timeout": "30.5"}
	assert.Equal(t, expected, objectResult)

}


// parse_mixed_types_strict_literal_get_string - function:get_string feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"host"})
	require.NoError(t, err)
	assert.Equal(t, "localhost", result)

}


// parse_mixed_types_strict_literal_get_int - function:get_int feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"port"})
	require.NoError(t, err)
	assert.Equal(t, 8080, result)

}


// parse_mixed_types_strict_literal_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseMixedTypesStrictLiteralGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"ssl"})
	require.NoError(t, err)
	assert.Equal(t, true, result)

}


// parse_mixed_types_strict_literal_get_float - function:get_float feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralGetFloat(t *testing.T) {
	

	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_float validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetFloat(hierarchy, []string{"timeout"})
	require.NoError(t, err)
	assert.Equal(t, 30.5, result)

}


// parse_with_whitespace_parse - function:parse feature:whitespace feature:optional_typed_accessors
func TestParseWithWhitespaceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "number", Value: "42"}, mock.Entry{Key: "flag", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}


// parse_with_whitespace_build_hierarchy - function:build_hierarchy feature:whitespace feature:optional_typed_accessors
func TestParseWithWhitespaceBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"flag": "true", "number": "42"}
	assert.Equal(t, expected, objectResult)

}


// parse_with_whitespace_get_int - function:get_int feature:whitespace feature:optional_typed_accessors
func TestParseWithWhitespaceGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"number"})
	require.NoError(t, err)
	assert.Equal(t, 42, result)

}


// parse_with_whitespace_get_bool - function:get_bool feature:whitespace feature:optional_typed_accessors
func TestParseWithWhitespaceGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"flag"})
	require.NoError(t, err)
	assert.Equal(t, true, result)

}


// parse_with_conservative_options_parse - function:parse feature:optional_typed_accessors
func TestParseWithConservativeOptionsParse(t *testing.T) {
	

	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "number", Value: "42"}, mock.Entry{Key: "decimal", Value: "3.14"}, mock.Entry{Key: "flag", Value: "true"}, mock.Entry{Key: "text", Value: "hello"}}
	assert.Equal(t, expected, parseResult)

}


// parse_with_conservative_options_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseWithConservativeOptionsBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"decimal": "3.14", "flag": "true", "number": "42", "text": "hello"}
	assert.Equal(t, expected, objectResult)

}


// parse_with_conservative_options_get_string - function:get_string feature:optional_typed_accessors
func TestParseWithConservativeOptionsGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"decimal"})
	require.NoError(t, err)
	assert.Equal(t, "3.14", result)

}


// parse_with_conservative_options_get_int - function:get_int feature:optional_typed_accessors
func TestParseWithConservativeOptionsGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"number"})
	require.NoError(t, err)
	assert.Equal(t, 42, result)

}


// parse_integer_error_parse - function:parse feature:optional_typed_accessors
func TestParseIntegerErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `port = not_a_number`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "port", Value: "not_a_number"}}
	assert.Equal(t, expected, parseResult)

}


// parse_integer_error_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseIntegerErrorBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `port = not_a_number`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"port": "not_a_number"}
	assert.Equal(t, expected, objectResult)

}


// parse_integer_error_get_int - function:get_int feature:optional_typed_accessors
func TestParseIntegerErrorGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `port = not_a_number`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"port"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, 0, result)
	}

}


// parse_float_error_parse - function:parse feature:optional_typed_accessors
func TestParseFloatErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `temperature = invalid`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "temperature", Value: "invalid"}}
	assert.Equal(t, expected, parseResult)

}


// parse_float_error_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseFloatErrorBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `temperature = invalid`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"temperature": "invalid"}
	assert.Equal(t, expected, objectResult)

}


// parse_float_error_get_float - function:get_float feature:optional_typed_accessors
func TestParseFloatErrorGetFloat(t *testing.T) {
	

	ccl := mock.New()
	input := `temperature = invalid`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_float validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetFloat(hierarchy, []string{"temperature"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, 0.0, result)
	}

}


// parse_boolean_error_parse - function:parse feature:optional_typed_accessors
func TestParseBooleanErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `enabled = maybe`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "enabled", Value: "maybe"}}
	assert.Equal(t, expected, parseResult)

}


// parse_boolean_error_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestParseBooleanErrorBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `enabled = maybe`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"enabled": "maybe"}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_error_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseBooleanErrorGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `enabled = maybe`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"enabled"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// parse_missing_path_error_parse - function:parse
func TestParseMissingPathErrorParse(t *testing.T) {
	

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


// parse_missing_path_error_build_hierarchy - function:build_hierarchy
func TestParseMissingPathErrorBuildHierarchy(t *testing.T) {
	

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


// parse_missing_path_error_get_string - function:get_string
func TestParseMissingPathErrorGetString(t *testing.T) {
	

	ccl := mock.New()
	input := `existing = value`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_string validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetString(hierarchy, []string{"missing"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, "", result)
	}

}


// boolean_case_sensitivity_uppercase_parse - function:parse feature:optional_typed_accessors
func TestBooleanCaseSensitivityUppercaseParse(t *testing.T) {
	

	ccl := mock.New()
	input := `upper_true = TRUE
upper_false = FALSE`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "upper_true", Value: "TRUE"}, mock.Entry{Key: "upper_false", Value: "FALSE"}}
	assert.Equal(t, expected, parseResult)

}


// boolean_case_sensitivity_uppercase_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestBooleanCaseSensitivityUppercaseGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `upper_true = TRUE
upper_false = FALSE`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"upper_true"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// boolean_case_sensitivity_mixed_parse - function:parse feature:optional_typed_accessors
func TestBooleanCaseSensitivityMixedParse(t *testing.T) {
	

	ccl := mock.New()
	input := `mixed_true = True
mixed_false = False`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "mixed_true", Value: "True"}, mock.Entry{Key: "mixed_false", Value: "False"}}
	assert.Equal(t, expected, parseResult)

}


// boolean_case_sensitivity_mixed_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestBooleanCaseSensitivityMixedGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `mixed_true = True
mixed_false = False`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"mixed_true"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// boolean_lenient_uppercase_yes_no_parse - function:parse feature:optional_typed_accessors
func TestBooleanLenientUppercaseYesNoParse(t *testing.T) {
	

	ccl := mock.New()
	input := `upper_yes = YES
upper_no = NO`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "upper_yes", Value: "YES"}, mock.Entry{Key: "upper_no", Value: "NO"}}
	assert.Equal(t, expected, parseResult)

}


// boolean_lenient_uppercase_yes_no_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_lenient
func TestBooleanLenientUppercaseYesNoGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `upper_yes = YES
upper_no = NO`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"upper_yes"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// boolean_numeric_one_zero_strict_parse - function:parse feature:optional_typed_accessors
func TestBooleanNumericOneZeroStrictParse(t *testing.T) {
	

	ccl := mock.New()
	input := `one = 1
zero = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "one", Value: "1"}, mock.Entry{Key: "zero", Value: "0"}}
	assert.Equal(t, expected, parseResult)

}


// boolean_numeric_one_zero_strict_get_int - function:get_int feature:optional_typed_accessors
func TestBooleanNumericOneZeroStrictGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `one = 1
zero = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"one"})
	require.NoError(t, err)
	assert.Equal(t, 1, result)

}


// boolean_numeric_one_zero_strict_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestBooleanNumericOneZeroStrictGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `one = 1
zero = 0`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"one"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// boolean_with_whitespace_parse - function:parse feature:optional_typed_accessors feature:whitespace
func TestBooleanWithWhitespaceParse(t *testing.T) {
	

	ccl := mock.New()
	input := `padded =   true   `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "padded", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}


// boolean_with_whitespace_get_bool - function:get_bool feature:optional_typed_accessors feature:whitespace behavior:boolean_strict
func TestBooleanWithWhitespaceGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `padded =   true   `
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"padded"})
	require.NoError(t, err)
	assert.Equal(t, true, result)

}


// boolean_nested_object_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestBooleanNestedObjectBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  debug = true
  verbose = false
  experimental = yes`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"debug": "true", "experimental": "yes", "verbose": "false"}}
	assert.Equal(t, expected, objectResult)

}


// type_mismatch_get_int_on_bool_parse - function:parse feature:optional_typed_accessors
func TestTypeMismatchGetIntOnBoolParse(t *testing.T) {
	

	ccl := mock.New()
	input := `flag = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "flag", Value: "true"}}
	assert.Equal(t, expected, parseResult)

}


// type_mismatch_get_int_on_bool_get_int - function:get_int feature:optional_typed_accessors
func TestTypeMismatchGetIntOnBoolGetInt(t *testing.T) {
	

	ccl := mock.New()
	input := `flag = true`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_int validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetInt(hierarchy, []string{"flag"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, 0, result)
	}

}


// type_mismatch_get_bool_on_int_parse - function:parse feature:optional_typed_accessors
func TestTypeMismatchGetBoolOnIntParse(t *testing.T) {
	

	ccl := mock.New()
	input := `number = 42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "number", Value: "42"}}
	assert.Equal(t, expected, parseResult)

}


// type_mismatch_get_bool_on_int_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestTypeMismatchGetBoolOnIntGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `number = 42`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"number"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


// type_mismatch_get_float_on_bool_parse - function:parse feature:optional_typed_accessors
func TestTypeMismatchGetFloatOnBoolParse(t *testing.T) {
	

	ccl := mock.New()
	input := `flag = false`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "flag", Value: "false"}}
	assert.Equal(t, expected, parseResult)

}


// type_mismatch_get_float_on_bool_get_float - function:get_float feature:optional_typed_accessors
func TestTypeMismatchGetFloatOnBoolGetFloat(t *testing.T) {
	

	ccl := mock.New()
	input := `flag = false`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_float validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetFloat(hierarchy, []string{"flag"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, 0.0, result)
	}

}


// type_mismatch_nested_path_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestTypeMismatchNestedPathBuildHierarchy(t *testing.T) {
	

	ccl := mock.New()
	input := `config =
  name = test
  count = abc`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{"config": map[string]interface{}{"count": "abc", "name": "test"}}
	assert.Equal(t, expected, objectResult)

}


// boolean_empty_value_error_parse - function:parse feature:optional_typed_accessors
func TestBooleanEmptyValueErrorParse(t *testing.T) {
	

	ccl := mock.New()
	input := `empty =`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty", Value: ""}}
	assert.Equal(t, expected, parseResult)

}


// boolean_empty_value_error_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestBooleanEmptyValueErrorGetBool(t *testing.T) {
	

	ccl := mock.New()
	input := `empty =`
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// get_bool validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	hierarchy := ccl.BuildHierarchy(parseResult)
	result, err := ccl.GetBool(hierarchy, []string{"empty"})
	if err != nil {
		require.Error(t, err)
	} else {
		assert.Equal(t, false, result)
	}

}


