package parsing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tylerbutler/ccl-test-data/internal/mock"
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_basic_integer_get_int - function:get_int feature:optional_typed_accessors
func TestParseBasicIntegerGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_basic_float_get_float - function:get_float feature:optional_typed_accessors
func TestParseBasicFloatGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_true_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict behavior:boolean_lenient
func TestParseBooleanTrueGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_yes_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_lenient
func TestParseBooleanYesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_yes_strict_literal_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseBooleanYesStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_false_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict behavior:boolean_lenient
func TestParseBooleanFalseGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_string_fallback_get_string - function:get_string
func TestParseStringFallbackGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_negative_integer_get_int - function:get_int feature:optional_typed_accessors
func TestParseNegativeIntegerGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_get_int - function:get_int feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_get_bool - function:get_bool feature:empty_keys feature:optional_typed_accessors behavior:boolean_lenient
func TestParseZeroValuesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_get_float - function:get_float feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_int - function:get_int feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_bool - function:get_bool feature:empty_keys feature:optional_typed_accessors behavior:boolean_strict
func TestParseZeroValuesStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_float - function:get_float feature:empty_keys feature:optional_typed_accessors
func TestParseZeroValuesStrictLiteralGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_get_int - function:get_int feature:optional_typed_accessors
func TestParseBooleanVariantsGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_lenient
func TestParseBooleanVariantsGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_strict_literal_get_int - function:get_int feature:optional_typed_accessors
func TestParseBooleanVariantsStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_strict_literal_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseBooleanVariantsStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_string - function:get_string feature:optional_typed_accessors
func TestParseMixedTypesGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_int - function:get_int feature:optional_typed_accessors
func TestParseMixedTypesGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_lenient
func TestParseMixedTypesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_float - function:get_float feature:optional_typed_accessors
func TestParseMixedTypesGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_string - function:get_string feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_int - function:get_int feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseMixedTypesStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_float - function:get_float feature:optional_typed_accessors
func TestParseMixedTypesStrictLiteralGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_whitespace_get_int - function:get_int feature:whitespace feature:optional_typed_accessors
func TestParseWithWhitespaceGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_whitespace_get_bool - function:get_bool feature:whitespace feature:optional_typed_accessors
func TestParseWithWhitespaceGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_conservative_options_get_string - function:get_string feature:optional_typed_accessors
func TestParseWithConservativeOptionsGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_conservative_options_get_int - function:get_int feature:optional_typed_accessors
func TestParseWithConservativeOptionsGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_integer_error_get_int - function:get_int feature:optional_typed_accessors
func TestParseIntegerErrorGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_float_error_get_float - function:get_float feature:optional_typed_accessors
func TestParseFloatErrorGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_error_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestParseBooleanErrorGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_missing_path_error_get_string - function:get_string
func TestParseMissingPathErrorGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// boolean_numeric_one_zero_strict_get_bool - function:get_bool feature:optional_typed_accessors behavior:boolean_strict
func TestBooleanNumericOneZeroStrictGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// boolean_nested_object_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestBooleanNestedObjectBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// type_mismatch_nested_path_build_hierarchy - function:build_hierarchy feature:optional_typed_accessors
func TestTypeMismatchNestedPathBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
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
	t.Skip("Test does not match run-only filter: [function:parse]")
}
