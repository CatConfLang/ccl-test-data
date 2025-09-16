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


// parse_basic_integer_buildhierarchy - function:buildhierarchy (level 0)
func TestParseBasicIntegerBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = 8080`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_basic_integer_getint - function:getint (level 0)
func TestParseBasicIntegerGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = 8080`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

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


// parse_basic_float_buildhierarchy - function:buildhierarchy (level 0)
func TestParseBasicFloatBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = 98.6`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_basic_float_getfloat - function:getfloat (level 0)
func TestParseBasicFloatGetfloat(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = 98.6`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getfloat validation

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


// parse_boolean_true_buildhierarchy - function:buildhierarchy behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanTrueBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = true`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_boolean_true_getbool - function:getbool behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanTrueGetbool(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = true`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getbool validation

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


// parse_boolean_yes_buildhierarchy - function:buildhierarchy behavior:boolean_lenient (level 0)
func TestParseBooleanYesBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_boolean_yes_getbool - function:getbool behavior:boolean_lenient (level 0)
func TestParseBooleanYesGetbool(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getbool validation

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


// parse_boolean_yes_strict_literal_buildhierarchy - function:buildhierarchy behavior:boolean_strict (level 0)
func TestParseBooleanYesStrictLiteralBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

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


// parse_boolean_false_buildhierarchy - function:buildhierarchy behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanFalseBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `disabled = false`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_boolean_false_getbool - function:getbool behavior:boolean_strict behavior:boolean_lenient (level 0)
func TestParseBooleanFalseGetbool(t *testing.T) {
	
	
	ccl := mock.New()
	input := `disabled = false`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getbool validation

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


// parse_string_fallback_buildhierarchy - function:buildhierarchy (level 0)
func TestParseStringFallbackBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_string_fallback_getstring - function:getstring (level 0)
func TestParseStringFallbackGetstring(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getstring validation

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


// parse_negative_integer_buildhierarchy - function:buildhierarchy (level 0)
func TestParseNegativeIntegerBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `offset = -42`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_negative_integer_getint - function:getint (level 0)
func TestParseNegativeIntegerGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `offset = -42`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

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


// parse_zero_values_buildhierarchy - function:buildhierarchy behavior:boolean_lenient (level 0)
func TestParseZeroValuesBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_zero_values_getint - function:getint behavior:boolean_lenient (level 0)
func TestParseZeroValuesGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

}


// parse_zero_values_getbool - function:getbool behavior:boolean_lenient (level 0)
func TestParseZeroValuesGetbool(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getbool validation

}


// parse_zero_values_getfloat - function:getfloat behavior:boolean_lenient (level 0)
func TestParseZeroValuesGetfloat(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getfloat validation

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


// parse_zero_values_strict_literal_buildhierarchy - function:buildhierarchy behavior:boolean_strict (level 0)
func TestParseZeroValuesStrictLiteralBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_zero_values_strict_literal_getint - function:getint behavior:boolean_strict (level 0)
func TestParseZeroValuesStrictLiteralGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

}


// parse_zero_values_strict_literal_getfloat - function:getfloat behavior:boolean_strict (level 0)
func TestParseZeroValuesStrictLiteralGetfloat(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getfloat validation

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


// parse_boolean_variants_buildhierarchy - function:buildhierarchy behavior:boolean_lenient (level 0)
func TestParseBooleanVariantsBuildhierarchy(t *testing.T) {
	
	
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
	
	// TODO: Implement buildhierarchy validation

}


// parse_boolean_variants_getint - function:getint behavior:boolean_lenient (level 0)
func TestParseBooleanVariantsGetint(t *testing.T) {
	
	
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
	
	// TODO: Implement getint validation

}


// parse_boolean_variants_getbool - function:getbool behavior:boolean_lenient (level 0)
func TestParseBooleanVariantsGetbool(t *testing.T) {
	
	
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
	
	// TODO: Implement getbool validation

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


// parse_boolean_variants_strict_literal_buildhierarchy - function:buildhierarchy behavior:boolean_strict (level 0)
func TestParseBooleanVariantsStrictLiteralBuildhierarchy(t *testing.T) {
	
	
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
	
	// TODO: Implement buildhierarchy validation

}


// parse_boolean_variants_strict_literal_getint - function:getint behavior:boolean_strict (level 0)
func TestParseBooleanVariantsStrictLiteralGetint(t *testing.T) {
	
	
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
	
	// TODO: Implement getint validation

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


// parse_mixed_types_buildhierarchy - function:buildhierarchy behavior:boolean_lenient (level 0)
func TestParseMixedTypesBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_mixed_types_getstring - function:getstring behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetstring(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getstring validation

}


// parse_mixed_types_getint - function:getint behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

}


// parse_mixed_types_getbool - function:getbool behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetbool(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getbool validation

}


// parse_mixed_types_getfloat - function:getfloat behavior:boolean_lenient (level 0)
func TestParseMixedTypesGetfloat(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getfloat validation

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


// parse_mixed_types_strict_literal_buildhierarchy - function:buildhierarchy behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_mixed_types_strict_literal_getstring - function:getstring behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetstring(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getstring validation

}


// parse_mixed_types_strict_literal_getint - function:getint behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

}


// parse_mixed_types_strict_literal_getbool - function:getbool behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetbool(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getbool validation

}


// parse_mixed_types_strict_literal_getfloat - function:getfloat behavior:boolean_strict (level 0)
func TestParseMixedTypesStrictLiteralGetfloat(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getfloat validation

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


// parse_empty_value_buildhierarchy - function:buildhierarchy (level 0)
func TestParseEmptyValueBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `empty_key =`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_empty_value_getstring - function:getstring (level 0)
func TestParseEmptyValueGetstring(t *testing.T) {
	
	
	ccl := mock.New()
	input := `empty_key =`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getstring validation

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


// empty_value_reference_behavior_buildhierarchy - function:buildhierarchy (level 0)
func TestEmptyValueReferenceBehaviorBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `empty_key =`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

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


// parse_with_whitespace_buildhierarchy - function:buildhierarchy (level 0)
func TestParseWithWhitespaceBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_with_whitespace_getint - function:getint (level 0)
func TestParseWithWhitespaceGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

}


// parse_with_whitespace_getbool - function:getbool (level 0)
func TestParseWithWhitespaceGetbool(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getbool validation

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


// parse_with_conservative_options_buildhierarchy - function:buildhierarchy (level 0)
func TestParseWithConservativeOptionsBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


// parse_with_conservative_options_getstring - function:getstring (level 0)
func TestParseWithConservativeOptionsGetstring(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getstring validation

}


// parse_with_conservative_options_getint - function:getint (level 0)
func TestParseWithConservativeOptionsGetint(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement getint validation

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


// parse_integer_error_buildhierarchy - function:buildhierarchy (level 0)
func TestParseIntegerErrorBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = not_a_number`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

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


// parse_float_error_buildhierarchy - function:buildhierarchy (level 0)
func TestParseFloatErrorBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = invalid`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

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


// parse_boolean_error_buildhierarchy - function:buildhierarchy behavior:boolean_strict (level 0)
func TestParseBooleanErrorBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = maybe`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

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


// parse_missing_path_error_buildhierarchy - function:buildhierarchy (level 0)
func TestParseMissingPathErrorBuildhierarchy(t *testing.T) {
	
	
	ccl := mock.New()
	input := `existing = value`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// TODO: Implement buildhierarchy validation

}


