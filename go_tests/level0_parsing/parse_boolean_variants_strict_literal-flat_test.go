package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/parse_boolean_variants_strict_literal-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// parse_boolean_variants_strict_literal_get_int - function:get_int (level 0)
func TestParseBooleanVariantsStrictLiteralGetInt(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_strict_literal_parse - function:parse (level 0)
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

// parse_boolean_variants_strict_literal_build_hierarchy - function:build_hierarchy (level 0)
func TestParseBooleanVariantsStrictLiteralBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// parse_boolean_variants_strict_literal_get_bool - function:get_bool (level 0)
func TestParseBooleanVariantsStrictLiteralGetBool(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
