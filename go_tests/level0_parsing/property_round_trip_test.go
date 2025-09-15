package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/property_round_trip.json
// Suite: Flat Format
// Version: 1.0

// round_trip_whitespace_normalization_parse - function:parse (level 0)
func TestRoundTripWhitespaceNormalizationParse(t *testing.T) {

	ccl := mock.New()
	input := `  key  =  value  
  nested  = 
    sub  =  val  `

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "value  \n  nested  = \n    sub  =  val"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_empty_keys_lists_parse - function:parse (level 0)
func TestRoundTripEmptyKeysListsParse(t *testing.T) {

	ccl := mock.New()
	input := `= item1
= item2
regular = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "item1"}, mock.Entry{Key: "", Value: "item2"}, mock.Entry{Key: "regular", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_nested_structures_parse - function:parse (level 0)
func TestRoundTripNestedStructuresParse(t *testing.T) {

	ccl := mock.New()
	input := `config =
  host = localhost
  port = 8080
  db =
    name = mydb
    user = admin`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "config", Value: "\n  host = localhost\n  port = 8080\n  db =\n    name = mydb\n    user = admin"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_multiline_values_parse - function:parse (level 0)
func TestRoundTripMultilineValuesParse(t *testing.T) {

	ccl := mock.New()
	input := `script =
  #!/bin/bash
  echo hello
  exit 0`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "script", Value: "\n  #!/bin/bash\n  echo hello\n  exit 0"}}
	assert.Equal(t, expected, parseResult)

}

// round_trip_mixed_content_parse - function:parse (level 0)
func TestRoundTripMixedContentParse(t *testing.T) {

	ccl := mock.New()
	input := `name = Alice
= first item
config =
  port = 3000
= second item
final = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Alice"}, mock.Entry{Key: "", Value: "first item"}, mock.Entry{Key: "config", Value: "\n  port = 3000"}, mock.Entry{Key: "", Value: "second item"}, mock.Entry{Key: "final", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_empty_values_parse - function:parse (level 0)
func TestCanonicalFormatEmptyValuesParse(t *testing.T) {

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

// canonical_format_empty_values_canonical - function:canonical (level 0)
func TestCanonicalFormatEmptyValuesCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_tab_preservation_parse - function:parse behavior:tabs_preserve (level 0)
func TestCanonicalFormatTabPreservationParse(t *testing.T) {

	ccl := mock.New()
	input := `value_with_tabs = text		with	tabs	`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "value_with_tabs", Value: "text\t\twith\ttabs\t"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_tab_preservation_canonical - function:canonical behavior:tabs_preserve (level 0)
func TestCanonicalFormatTabPreservationCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_unicode_parse - function:parse (level 0)
func TestCanonicalFormatUnicodeParse(t *testing.T) {

	ccl := mock.New()
	input := `unicode = 你好世界
emo = 🌟✨`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "unicode", Value: "你好世界"}, mock.Entry{Key: "emo", Value: "🌟✨"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_unicode_canonical - function:canonical (level 0)
func TestCanonicalFormatUnicodeCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_complex_nesting_parse - function:parse (level 0)
func TestRoundTripComplexNestingParse(t *testing.T) {

	ccl := mock.New()
	input := `app =
  = item1
  config =
    = nested_item
    db =
      host = localhost
      = db_item
  = item2`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "app", Value: "\n  = item1\n  config =\n    = nested_item\n    db =\n      host = localhost\n      = db_item\n  = item2"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_line_endings_proposed_parse - function:parse behavior:crlf_preserve_literal (level 0)
func TestCanonicalFormatLineEndingsProposedParse(t *testing.T) {

	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_line_endings_proposed_canonical - function:canonical behavior:crlf_preserve_literal (level 0)
func TestCanonicalFormatLineEndingsProposedCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// crlf_normalize_to_lf_proposed_parse - function:parse behavior:crlf_normalize_to_lf (level 0)
func TestCrlfNormalizeToLfProposedParse(t *testing.T) {

	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}}
	assert.Equal(t, expected, parseResult)

}

// crlf_normalize_to_lf_proposed_canonical - function:canonical behavior:crlf_normalize_to_lf (level 0)
func TestCrlfNormalizeToLfProposedCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// crlf_normalize_to_lf_indented_proposed_parse - function:parse behavior:crlf_normalize_to_lf (level 0)
func TestCrlfNormalizeToLfIndentedProposedParse(t *testing.T) {

	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}}
	assert.Equal(t, expected, parseResult)

}

// crlf_normalize_to_lf_indented_proposed_canonical - function:canonical behavior:crlf_normalize_to_lf (level 0)
func TestCrlfNormalizeToLfIndentedProposedCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_deeply_nested_parse - function:parse (level 0)
func TestRoundTripDeeplyNestedParse(t *testing.T) {

	ccl := mock.New()
	input := `level1 =
  level2 =
    level3 =
      level4 =
        deep = value
        = deep_item`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "level1", Value: "\n  level2 =\n    level3 =\n      level4 =\n        deep = value\n        = deep_item"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_consistent_spacing_parse - function:parse behavior:strict_spacing (level 0)
func TestCanonicalFormatConsistentSpacingParse(t *testing.T) {

	ccl := mock.New()
	input := `key1=value1
key2  =  value2
key3	=	value3`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1"}, mock.Entry{Key: "key2", Value: "value2"}, mock.Entry{Key: "key3", Value: "\tvalue3"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_consistent_spacing_canonical - function:canonical behavior:strict_spacing (level 0)
func TestCanonicalFormatConsistentSpacingCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_empty_multiline_parse - function:parse (level 0)
func TestRoundTripEmptyMultilineParse(t *testing.T) {

	ccl := mock.New()
	input := `empty_section =

other = value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "empty_section", Value: ""}, mock.Entry{Key: "other", Value: "value"}}
	assert.Equal(t, expected, parseResult)

}

// deterministic_output_parse - function:parse (level 0)
func TestDeterministicOutputParse(t *testing.T) {

	ccl := mock.New()
	input := `z = last
a = first
m = middle`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "z", Value: "last"}, mock.Entry{Key: "a", Value: "first"}, mock.Entry{Key: "m", Value: "middle"}}
	assert.Equal(t, expected, parseResult)

}

// deterministic_output_canonical - function:canonical (level 0)
func TestDeterministicOutputCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_empty_values_ocaml_reference_canonical - function:canonical (level 0)
func TestCanonicalFormatEmptyValuesOcamlReferenceCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_tab_preservation_ocaml_reference_canonical - function:canonical behavior:tabs_preserve (level 0)
func TestCanonicalFormatTabPreservationOcamlReferenceCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_unicode_ocaml_reference_canonical - function:canonical (level 0)
func TestCanonicalFormatUnicodeOcamlReferenceCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_line_endings_reference_behavior_parse - function:parse behavior:crlf_preserve_literal (level 0)
func TestCanonicalFormatLineEndingsReferenceBehaviorParse(t *testing.T) {

	ccl := mock.New()
	input := "key1 = value1\r\nkey2 = value2\r\n"

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key1", Value: "value1\r"}, mock.Entry{Key: "key2", Value: "value2\r"}}
	assert.Equal(t, expected, parseResult)

}

// canonical_format_line_endings_reference_behavior_canonical - function:canonical behavior:crlf_preserve_literal (level 0)
func TestCanonicalFormatLineEndingsReferenceBehaviorCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_consistent_spacing_ocaml_reference_canonical - function:canonical behavior:strict_spacing (level 0)
func TestCanonicalFormatConsistentSpacingOcamlReferenceCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deterministic_output_ocaml_reference_canonical - function:canonical (level 0)
func TestDeterministicOutputOcamlReferenceCanonical(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
