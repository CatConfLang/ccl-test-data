package level4_typed_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-typed-access.json
// Suite: CCL Typed Access - Validation Format
// Version: 2.0
// Description: Type-aware value extraction with validation and smart type inference using new validation-based format - get_string(), get_int(), get_bool(), get_float()


// parse_basic_integer - function:get-int function:make-objects function:parse (level 4)
func TestParseBasicInteger(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = 8080`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "port", Value: "8080"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"port": "8080"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[port] expected:8080]] count:1]

}


// parse_basic_float - function:get-float function:make-objects function:parse (level 4)
func TestParseBasicFloat(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = 98.6`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "temperature", Value: "98.6"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"temperature": "98.6"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetFloat validation
	// Validation data: map[cases:[map[args:[temperature] expected:98.6]] count:1]

}


// parse_boolean_true - function:get-bool function:make-objects function:parse (level 4)
func TestParseBooleanTrue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = true`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "enabled", Value: "true"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"enabled": "true"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[enabled] expected:true]] count:1]

}


// parse_boolean_yes - function:get-bool function:make-objects function:parse (level 4)
func TestParseBooleanYes(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "active", Value: "yes"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"active": "yes"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[active] expected:true]] count:1]

}


// parse_boolean_yes_strict - function:get-bool function:make-objects function:parse (level 4)
func TestParseBooleanYesStrict(t *testing.T) {
	
	
	ccl := mock.New()
	input := `active = yes`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "active", Value: "yes"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"active": "yes"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[active] error:true error_message:Value is not a boolean]] count:1]

}


// parse_boolean_false - function:get-bool function:make-objects function:parse (level 4)
func TestParseBooleanFalse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `disabled = false`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "disabled", Value: "false"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"disabled": "false"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[disabled] expected:false]] count:1]

}


// parse_string_fallback - function:get-string function:make-objects function:parse (level 4)
func TestParseStringFallback(t *testing.T) {
	
	
	ccl := mock.New()
	input := `name = Alice`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"name": "Alice"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetString validation
	// Validation data: map[cases:[map[args:[name] expected:Alice]] count:1]

}


// parse_negative_integer - function:get-int function:make-objects function:parse (level 4)
func TestParseNegativeInteger(t *testing.T) {
	
	
	ccl := mock.New()
	input := `offset = -42`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "offset", Value: "-42"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"offset": "-42"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[offset] expected:-42]] count:1]

}


// parse_zero_values - feature:empty-keys function:get-bool function:get-float function:get-int function:make-objects function:parse (level 4)
func TestParseZeroValues(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "count", Value: "0"}, mock.Entry{Key: "distance", Value: "0.0"}, mock.Entry{Key: "disabled", Value: "no"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"count": "0", "disabled": "no", "distance": "0.0"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[count] expected:0]] count:1]
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[disabled] expected:false]] count:1]
	// TODO: Implement GetFloat validation
	// Validation data: map[cases:[map[args:[distance] expected:0]] count:1]

}


// parse_zero_values_strict - feature:empty-keys function:get-bool function:get-float function:get-int function:make-objects function:parse (level 4)
func TestParseZeroValuesStrict(t *testing.T) {
	
	
	ccl := mock.New()
	input := `count = 0
distance = 0.0
disabled = no`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "count", Value: "0"}, mock.Entry{Key: "distance", Value: "0.0"}, mock.Entry{Key: "disabled", Value: "no"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"count": "0", "disabled": "no", "distance": "0.0"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[count] expected:0]] count:1]
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[disabled] error:true error_message:Value is not a boolean]] count:1]
	// TODO: Implement GetFloat validation
	// Validation data: map[cases:[map[args:[distance] expected:0]] count:1]

}


// parse_boolean_variants - function:get-bool function:get-int function:make-objects function:parse (level 4)
func TestParseBooleanVariants(t *testing.T) {
	
	
	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "flag1", Value: "yes"}, mock.Entry{Key: "flag2", Value: "on"}, mock.Entry{Key: "flag3", Value: "1"}, mock.Entry{Key: "flag4", Value: "false"}, mock.Entry{Key: "flag5", Value: "no"}, mock.Entry{Key: "flag6", Value: "off"}, mock.Entry{Key: "flag7", Value: "0"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"flag1": "yes", "flag2": "on", "flag3": "1", "flag4": "false", "flag5": "no", "flag6": "off", "flag7": "0"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[flag3] expected:1]] count:1]
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[flag1] expected:true]] count:1]

}


// parse_boolean_variants_strict - function:get-bool function:get-int function:make-objects function:parse (level 4)
func TestParseBooleanVariantsStrict(t *testing.T) {
	
	
	ccl := mock.New()
	input := `flag1 = yes
flag2 = on
flag3 = 1
flag4 = false
flag5 = no
flag6 = off
flag7 = 0`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "flag1", Value: "yes"}, mock.Entry{Key: "flag2", Value: "on"}, mock.Entry{Key: "flag3", Value: "1"}, mock.Entry{Key: "flag4", Value: "false"}, mock.Entry{Key: "flag5", Value: "no"}, mock.Entry{Key: "flag6", Value: "off"}, mock.Entry{Key: "flag7", Value: "0"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"flag1": "yes", "flag2": "on", "flag3": "1", "flag4": "false", "flag5": "no", "flag6": "off", "flag7": "0"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[flag3] expected:1]] count:1]
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[flag1] error:true error_message:Value is not a boolean]] count:1]

}


// parse_mixed_types - function:get-bool function:get-float function:get-int function:get-string function:make-objects function:parse (level 4)
func TestParseMixedTypes(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "ssl", Value: "true"}, mock.Entry{Key: "timeout", Value: "30.5"}, mock.Entry{Key: "debug", Value: "off"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"debug": "off", "host": "localhost", "port": "8080", "ssl": "true", "timeout": "30.5"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetString validation
	// Validation data: map[cases:[map[args:[host] expected:localhost]] count:1]
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[port] expected:8080]] count:1]
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[ssl] expected:true]] count:1]
	// TODO: Implement GetFloat validation
	// Validation data: map[cases:[map[args:[timeout] expected:30.5]] count:1]

}


// parse_mixed_types_strict - function:get-bool function:get-float function:get-int function:get-string function:make-objects function:parse (level 4)
func TestParseMixedTypesStrict(t *testing.T) {
	
	
	ccl := mock.New()
	input := `host = localhost
port = 8080
ssl = true
timeout = 30.5
debug = off`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "ssl", Value: "true"}, mock.Entry{Key: "timeout", Value: "30.5"}, mock.Entry{Key: "debug", Value: "off"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"debug": "off", "host": "localhost", "port": "8080", "ssl": "true", "timeout": "30.5"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetString validation
	// Validation data: map[cases:[map[args:[host] expected:localhost]] count:1]
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[port] expected:8080]] count:1]
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[ssl] expected:true]] count:1]
	// TODO: Implement GetFloat validation
	// Validation data: map[cases:[map[args:[timeout] expected:30.5]] count:1]

}


// parse_empty_value - function:get-string function:make-objects function:parse (level 4)
func TestParseEmptyValue(t *testing.T) {
	
	
	ccl := mock.New()
	input := `empty_key =`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "empty_key", Value: ""}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"empty_key": ""}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetString validation
	// Validation data: map[cases:[map[args:[empty_key] expected:]] count:1]

}


// parse_with_whitespace - feature:whitespace function:get-bool function:get-int function:make-objects function:parse (level 4)
func TestParseWithWhitespace(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number =   42   
flag =  true  `
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "number", Value: "42"}, mock.Entry{Key: "flag", Value: "true"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"flag": "true", "number": "42"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[number] expected:42]] count:1]
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[flag] expected:true]] count:1]

}


// parse_with_conservative_options - function:get-int function:get-string function:make-objects function:parse (level 4)
func TestParseWithConservativeOptions(t *testing.T) {
	
	
	ccl := mock.New()
	input := `number = 42
decimal = 3.14
flag = true
text = hello`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "number", Value: "42"}, mock.Entry{Key: "decimal", Value: "3.14"}, mock.Entry{Key: "flag", Value: "true"}, mock.Entry{Key: "text", Value: "hello"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"decimal": "3.14", "flag": "true", "number": "42", "text": "hello"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetString validation
	// Validation data: map[cases:[map[args:[decimal] expected:3.14]] count:1]
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[number] expected:42]] count:1]

}


// parse_integer_error - function:get-int function:make-objects function:parse (level 4)
func TestParseIntegerError(t *testing.T) {
	
	
	ccl := mock.New()
	input := `port = not_a_number`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "port", Value: "not_a_number"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"port": "not_a_number"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetInt validation
	// Validation data: map[cases:[map[args:[port] error:true error_message:Cannot parse 'not_a_number' as integer at path 'port']] count:1]

}


// parse_float_error - function:get-float function:make-objects function:parse (level 4)
func TestParseFloatError(t *testing.T) {
	
	
	ccl := mock.New()
	input := `temperature = invalid`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "temperature", Value: "invalid"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"temperature": "invalid"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetFloat validation
	// Validation data: map[cases:[map[args:[temperature] error:true error_message:Cannot parse 'invalid' as float at path 'temperature']] count:1]

}


// parse_boolean_error - function:get-bool function:make-objects function:parse (level 4)
func TestParseBooleanError(t *testing.T) {
	
	
	ccl := mock.New()
	input := `enabled = maybe`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "enabled", Value: "maybe"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"enabled": "maybe"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetBool validation
	// Validation data: map[cases:[map[args:[enabled] error:true error_message:Cannot parse 'maybe' as boolean at path 'enabled']] count:1]

}


// parse_missing_path_error - function:get-string function:make-objects function:parse (level 4)
func TestParseMissingPathError(t *testing.T) {
	
	
	ccl := mock.New()
	input := `existing = value`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	var objectResult map[string]interface{}
	
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "existing", Value: "value"}}
	assert.Equal(t, expectedParse, parseResult)
	// MakeObjects validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	objectResult = ccl.MakeObjects(parseResult)
	expectedObjects := map[string]interface{}{"existing": "value"}
	assert.Equal(t, expectedObjects, objectResult)
	// TODO: Implement GetString validation
	// Validation data: map[cases:[map[args:[missing] error:true error_message:Path 'missing' not found.]] count:1]

}


