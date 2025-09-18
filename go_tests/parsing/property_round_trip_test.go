package parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/property_round_trip.json
// Suite: Flat Format
// Version: 1.0

// round_trip_basic_round_trip - function:round_trip
func TestRoundTripBasicRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_whitespace_normalization_parse - function:parse
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

// round_trip_whitespace_normalization_round_trip - function:round_trip
func TestRoundTripWhitespaceNormalizationRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_empty_keys_lists_parse - function:parse
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

// round_trip_empty_keys_lists_round_trip - function:round_trip
func TestRoundTripEmptyKeysListsRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_nested_structures_parse - function:parse
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

// round_trip_nested_structures_round_trip - function:round_trip
func TestRoundTripNestedStructuresRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_multiline_values_parse - function:parse
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

// round_trip_multiline_values_round_trip - function:round_trip
func TestRoundTripMultilineValuesRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_mixed_content_parse - function:parse
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

// round_trip_mixed_content_round_trip - function:round_trip
func TestRoundTripMixedContentRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_empty_values_parse - function:parse
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

// canonical_format_empty_values_canonical_format - function:canonical_format
func TestCanonicalFormatEmptyValuesCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_tab_preservation_parse - function:parse behavior:tabs_preserve
func TestCanonicalFormatTabPreservationParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:tabs_preserve")
}

// canonical_format_tab_preservation_canonical_format - function:canonical_format behavior:tabs_preserve
func TestCanonicalFormatTabPreservationCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_unicode_parse - function:parse
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

// canonical_format_unicode_canonical_format - function:canonical_format
func TestCanonicalFormatUnicodeCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_complex_nesting_parse - function:parse
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

// round_trip_complex_nesting_round_trip - function:round_trip
func TestRoundTripComplexNestingRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_line_endings_proposed_parse - function:parse behavior:crlf_preserve_literal
func TestCanonicalFormatLineEndingsProposedParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:crlf_preserve_literal")
}

// canonical_format_line_endings_proposed_canonical_format - function:canonical_format behavior:crlf_preserve_literal
func TestCanonicalFormatLineEndingsProposedCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// crlf_normalize_to_lf_proposed_parse - function:parse behavior:crlf_normalize_to_lf
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

// crlf_normalize_to_lf_proposed_canonical_format - function:canonical_format behavior:crlf_normalize_to_lf
func TestCrlfNormalizeToLfProposedCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// crlf_normalize_to_lf_indented_proposed_parse - function:parse behavior:crlf_normalize_to_lf
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

// crlf_normalize_to_lf_indented_proposed_canonical_format - function:canonical_format behavior:crlf_normalize_to_lf
func TestCrlfNormalizeToLfIndentedProposedCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_deeply_nested_parse - function:parse
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

// round_trip_deeply_nested_round_trip - function:round_trip
func TestRoundTripDeeplyNestedRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_consistent_spacing_parse - function:parse behavior:strict_spacing
func TestCanonicalFormatConsistentSpacingParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:strict_spacing")
}

// canonical_format_consistent_spacing_canonical_format - function:canonical_format behavior:strict_spacing
func TestCanonicalFormatConsistentSpacingCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// round_trip_empty_multiline_parse - function:parse
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

// round_trip_empty_multiline_round_trip - function:round_trip
func TestRoundTripEmptyMultilineRoundTrip(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deterministic_output_parse - function:parse
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

// deterministic_output_canonical_format - function:canonical_format
func TestDeterministicOutputCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_empty_values_ocaml_reference_canonical_format - function:canonical_format
func TestCanonicalFormatEmptyValuesOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_tab_preservation_ocaml_reference_canonical_format - function:canonical_format behavior:tabs_preserve
func TestCanonicalFormatTabPreservationOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_unicode_ocaml_reference_canonical_format - function:canonical_format
func TestCanonicalFormatUnicodeOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_line_endings_reference_behavior_parse - function:parse behavior:crlf_preserve_literal
func TestCanonicalFormatLineEndingsReferenceBehaviorParse(t *testing.T) {
	t.Skip("Test skipped due to tag filter: behavior:crlf_preserve_literal")
}

// canonical_format_line_endings_reference_behavior_canonical_format - function:canonical_format behavior:crlf_preserve_literal
func TestCanonicalFormatLineEndingsReferenceBehaviorCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// canonical_format_consistent_spacing_ocaml_reference_canonical_format - function:canonical_format behavior:strict_spacing
func TestCanonicalFormatConsistentSpacingOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// deterministic_output_ocaml_reference_canonical_format - function:canonical_format
func TestDeterministicOutputOcamlReferenceCanonicalFormat(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
