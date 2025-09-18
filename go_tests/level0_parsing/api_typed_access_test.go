package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_typed_access.json
// Suite: Flat Format
// Version: 1.0

// parse_basic_integer_parse - function:parse (level 0)
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

// parse_basic_integer_build_hierarchy - function:build_hierarchy (level 0)
func TestParseBasicIntegerBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_basic_integer_get_int - function:get_int (level 0)
func TestParseBasicIntegerGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_basic_float_parse - function:parse (level 0)
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

// parse_basic_float_build_hierarchy - function:build_hierarchy (level 0)
func TestParseBasicFloatBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_basic_float_get_float - function:get_float (level 0)
func TestParseBasicFloatGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_true_parse - function:parse behavior:boolean_strict behavior:boolean_lenient (level 0)
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

// parse_boolean_true_build_hierarchy - function:build_hierarchy behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanTrueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_true_get_bool - function:get_bool behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanTrueGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_yes_parse - function:parse behavior:boolean_lenient (level 0)
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

// parse_boolean_yes_build_hierarchy - function:build_hierarchy behavior:boolean_lenient (level 0)
func TestParseBooleanYesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_yes_get_bool - function:get_bool behavior:boolean_lenient (level 0)
func TestParseBooleanYesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_yes_strict_literal_parse - function:parse behavior:boolean_strict (level 0)
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

// parse_boolean_yes_strict_literal_build_hierarchy - function:build_hierarchy behavior:boolean_strict (level 0)
func TestParseBooleanYesStrictLiteralBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_yes_strict_literal_get_bool - function:get_bool behavior:boolean_strict (level 0)
func TestParseBooleanYesStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_false_parse - function:parse behavior:boolean_strict behavior:boolean_lenient (level 0)
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

// parse_boolean_false_build_hierarchy - function:build_hierarchy behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanFalseBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_false_get_bool - function:get_bool behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanFalseGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_string_fallback_parse - function:parse (level 0)
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

// parse_string_fallback_build_hierarchy - function:build_hierarchy (level 0)
func TestParseStringFallbackBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_string_fallback_get_string - function:get_string (level 0)
func TestParseStringFallbackGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_negative_integer_parse - function:parse (level 0)
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

// parse_negative_integer_build_hierarchy - function:build_hierarchy (level 0)
func TestParseNegativeIntegerBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_negative_integer_get_int - function:get_int (level 0)
func TestParseNegativeIntegerGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_parse - function:parse behavior:boolean_lenient (level 0)
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

// parse_zero_values_build_hierarchy - function:build_hierarchy behavior:boolean_lenient (level 0)
func TestParseZeroValuesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_get_int - function:get_int behavior:boolean_lenient (level 0)
func TestParseZeroValuesGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_get_bool - function:get_bool behavior:boolean_lenient (level 0)
func TestParseZeroValuesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_get_float - function:get_float behavior:boolean_lenient (level 0)
func TestParseZeroValuesGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_parse - function:parse behavior:boolean_strict (level 0)
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

// parse_zero_values_strict_literal_build_hierarchy - function:build_hierarchy behavior:boolean_strict (level 0)
func TestParseZeroValuesStrictLiteralBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_int - function:get_int behavior:boolean_strict (level 0)
func TestParseZeroValuesStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_bool - function:get_bool behavior:boolean_strict (level 0)
func TestParseZeroValuesStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_zero_values_strict_literal_get_float - function:get_float behavior:boolean_strict (level 0)
func TestParseZeroValuesStrictLiteralGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_parse - function:parse behavior:boolean_lenient (level 0)
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

// parse_boolean_variants_build_hierarchy - function:build_hierarchy behavior:boolean_lenient (level 0)
func TestParseBooleanVariantsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_get_int - function:get_int behavior:boolean_lenient (level 0)
func TestParseBooleanVariantsGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_get_bool - function:get_bool behavior:boolean_lenient (level 0)
func TestParseBooleanVariantsGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_strict_literal_parse - function:parse behavior:boolean_strict (level 0)
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

// parse_boolean_variants_strict_literal_build_hierarchy - function:build_hierarchy behavior:boolean_strict (level 0)
func TestParseBooleanVariantsStrictLiteralBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_strict_literal_get_int - function:get_int behavior:boolean_strict (level 0)
func TestParseBooleanVariantsStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_strict_literal_get_bool - function:get_bool behavior:boolean_strict (level 0)
func TestParseBooleanVariantsStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_parse - function:parse behavior:boolean_lenient (level 0)
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

// parse_mixed_types_build_hierarchy - function:build_hierarchy behavior:boolean_lenient (level 0)
func TestParseMixedTypesBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_string - function:get_string behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_int - function:get_int behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_bool - function:get_bool behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_get_float - function:get_float behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_parse - function:parse behavior:boolean_strict (level 0)
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

// parse_mixed_types_strict_literal_build_hierarchy - function:build_hierarchy behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_string - function:get_string behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_int - function:get_int behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_bool - function:get_bool behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_mixed_types_strict_literal_get_float - function:get_float behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_empty_value_parse - function:parse (level 0)
func TestParseEmptyValueParse(t *testing.T) {

	ccl := mock.New()
	input := `empty_key =`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}

// parse_empty_value_build_hierarchy - function:build_hierarchy (level 0)
func TestParseEmptyValueBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_empty_value_get_string - function:get_string (level 0)
func TestParseEmptyValueGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// empty_value_reference_behavior_parse - function:parse (level 0)
func TestEmptyValueReferenceBehaviorParse(t *testing.T) {

	ccl := mock.New()
	input := `empty_key =`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_key", Value: ""}}
	assert.Equal(t, expected, parseResult)

}

// empty_value_reference_behavior_build_hierarchy - function:build_hierarchy (level 0)
func TestEmptyValueReferenceBehaviorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_whitespace_parse - function:parse (level 0)
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

// parse_with_whitespace_build_hierarchy - function:build_hierarchy (level 0)
func TestParseWithWhitespaceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_whitespace_get_int - function:get_int (level 0)
func TestParseWithWhitespaceGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_whitespace_get_bool - function:get_bool (level 0)
func TestParseWithWhitespaceGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_conservative_options_parse - function:parse (level 0)
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

// parse_with_conservative_options_build_hierarchy - function:build_hierarchy (level 0)
func TestParseWithConservativeOptionsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_conservative_options_get_string - function:get_string (level 0)
func TestParseWithConservativeOptionsGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_with_conservative_options_get_int - function:get_int (level 0)
func TestParseWithConservativeOptionsGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_integer_error_parse - function:parse (level 0)
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

// parse_integer_error_build_hierarchy - function:build_hierarchy (level 0)
func TestParseIntegerErrorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_integer_error_get_int - function:get_int (level 0)
func TestParseIntegerErrorGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_float_error_parse - function:parse (level 0)
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

// parse_float_error_build_hierarchy - function:build_hierarchy (level 0)
func TestParseFloatErrorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_float_error_get_float - function:get_float (level 0)
func TestParseFloatErrorGetFloat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_error_parse - function:parse behavior:boolean_strict (level 0)
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

// parse_boolean_error_build_hierarchy - function:build_hierarchy behavior:boolean_strict (level 0)
func TestParseBooleanErrorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_error_get_bool - function:get_bool behavior:boolean_strict (level 0)
func TestParseBooleanErrorGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_missing_path_error_parse - function:parse (level 0)
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

// parse_missing_path_error_build_hierarchy - function:build_hierarchy (level 0)
func TestParseMissingPathErrorBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_missing_path_error_get_string - function:get_string (level 0)
func TestParseMissingPathErrorGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
