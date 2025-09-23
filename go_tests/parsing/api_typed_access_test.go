package parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_typed_access.json
// Suite: Flat Format
// Version: 1.0



// parse_basic_integer_parse - function:parse
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


// parse_basic_integer_build_hierarchy - function:build_hierarchy
func TestParseBasicIntegerBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = 8080`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_basic_integer_get_int - function:get_int
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


// parse_basic_float_parse - function:parse
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


// parse_basic_float_build_hierarchy - function:build_hierarchy
func TestParseBasicFloatBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = 98.6`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_basic_float_get_float - function:get_float
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


// parse_boolean_true_parse - function:parse behavior:boolean_strict behavior:boolean_lenient
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


// parse_boolean_true_build_hierarchy - function:build_hierarchy behavior:boolean_strict behavior:boolean_lenient
func TestParseBooleanTrueBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = true`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_true_get_bool - function:get_bool behavior:boolean_strict behavior:boolean_lenient
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


// parse_boolean_yes_parse - function:parse behavior:boolean_lenient
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


// parse_boolean_yes_build_hierarchy - function:build_hierarchy behavior:boolean_lenient
func TestParseBooleanYesBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_yes_get_bool - function:get_bool behavior:boolean_lenient
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


// parse_boolean_yes_strict_literal_parse - function:parse behavior:boolean_strict
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


// parse_boolean_yes_strict_literal_build_hierarchy - function:build_hierarchy behavior:boolean_strict
func TestParseBooleanYesStrictLiteralBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_yes_strict_literal_get_bool - function:get_bool behavior:boolean_strict
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


// parse_boolean_false_parse - function:parse behavior:boolean_strict behavior:boolean_lenient
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


// parse_boolean_false_build_hierarchy - function:build_hierarchy behavior:boolean_strict behavior:boolean_lenient
func TestParseBooleanFalseBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `disabled = false`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_false_get_bool - function:get_bool behavior:boolean_strict behavior:boolean_lenient
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
	expected := map[string]interface{}{}
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


// parse_negative_integer_parse - function:parse
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


// parse_negative_integer_build_hierarchy - function:build_hierarchy
func TestParseNegativeIntegerBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `offset = -42`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_negative_integer_get_int - function:get_int
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


// parse_zero_values_parse - function:parse feature:empty_keys behavior:boolean_lenient
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


// parse_zero_values_build_hierarchy - function:build_hierarchy feature:empty_keys behavior:boolean_lenient
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_zero_values_get_int - function:get_int feature:empty_keys behavior:boolean_lenient
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


// parse_zero_values_get_bool - function:get_bool feature:empty_keys behavior:boolean_lenient
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


// parse_zero_values_get_float - function:get_float feature:empty_keys behavior:boolean_lenient
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


// parse_zero_values_strict_literal_parse - function:parse feature:empty_keys behavior:boolean_strict
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


// parse_zero_values_strict_literal_build_hierarchy - function:build_hierarchy feature:empty_keys behavior:boolean_strict
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_zero_values_strict_literal_get_int - function:get_int feature:empty_keys behavior:boolean_strict
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


// parse_zero_values_strict_literal_get_bool - function:get_bool feature:empty_keys behavior:boolean_strict
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


// parse_zero_values_strict_literal_get_float - function:get_float feature:empty_keys behavior:boolean_strict
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


// parse_boolean_variants_parse - function:parse behavior:boolean_lenient
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


// parse_boolean_variants_build_hierarchy - function:build_hierarchy behavior:boolean_lenient
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_variants_get_int - function:get_int behavior:boolean_lenient
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


// parse_boolean_variants_get_bool - function:get_bool behavior:boolean_lenient
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


// parse_boolean_variants_strict_literal_parse - function:parse behavior:boolean_strict
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


// parse_boolean_variants_strict_literal_build_hierarchy - function:build_hierarchy behavior:boolean_strict
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_variants_strict_literal_get_int - function:get_int behavior:boolean_strict
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


// parse_boolean_variants_strict_literal_get_bool - function:get_bool behavior:boolean_strict
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


// parse_mixed_types_parse - function:parse behavior:boolean_lenient
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


// parse_mixed_types_build_hierarchy - function:build_hierarchy behavior:boolean_lenient
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_mixed_types_get_string - function:get_string behavior:boolean_lenient
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


// parse_mixed_types_get_int - function:get_int behavior:boolean_lenient
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


// parse_mixed_types_get_bool - function:get_bool behavior:boolean_lenient
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


// parse_mixed_types_get_float - function:get_float behavior:boolean_lenient
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


// parse_mixed_types_strict_literal_parse - function:parse behavior:boolean_strict
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


// parse_mixed_types_strict_literal_build_hierarchy - function:build_hierarchy behavior:boolean_strict
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_mixed_types_strict_literal_get_string - function:get_string behavior:boolean_strict
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


// parse_mixed_types_strict_literal_get_int - function:get_int behavior:boolean_strict
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


// parse_mixed_types_strict_literal_get_bool - function:get_bool behavior:boolean_strict
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


// parse_mixed_types_strict_literal_get_float - function:get_float behavior:boolean_strict
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


// parse_with_whitespace_parse - function:parse feature:whitespace
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


// parse_with_whitespace_build_hierarchy - function:build_hierarchy feature:whitespace
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_with_whitespace_get_int - function:get_int feature:whitespace
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


// parse_with_whitespace_get_bool - function:get_bool feature:whitespace
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


// parse_with_conservative_options_parse - function:parse
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


// parse_with_conservative_options_build_hierarchy - function:build_hierarchy
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
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_with_conservative_options_get_string - function:get_string
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


// parse_with_conservative_options_get_int - function:get_int
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


// parse_integer_error_parse - function:parse
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


// parse_integer_error_build_hierarchy - function:build_hierarchy
func TestParseIntegerErrorBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = not_a_number`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_integer_error_get_int - function:get_int
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


// parse_float_error_parse - function:parse
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


// parse_float_error_build_hierarchy - function:build_hierarchy
func TestParseFloatErrorBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = invalid`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_float_error_get_float - function:get_float
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


// parse_boolean_error_parse - function:parse behavior:boolean_strict
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


// parse_boolean_error_build_hierarchy - function:build_hierarchy behavior:boolean_strict
func TestParseBooleanErrorBuildHierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = maybe`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// BuildHierarchy validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	objectResult := ccl.BuildHierarchy(parseResult)
	expected := map[string]interface{}{}
	assert.Equal(t, expected, objectResult)

}


// parse_boolean_error_get_bool - function:get_bool behavior:boolean_strict
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
	expected := map[string]interface{}{}
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


