package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/mixed_dotted_and_regular_keys-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// mixed_dotted_and_regular_keys_parse - function:parse (level 0)
func TestMixedDottedAndRegularKeysParse(t *testing.T) {

	ccl := mock.New()
	input := `app = MyApp
database.host = localhost
config =
  debug = true
logging.level = info`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "app", Value: "MyApp"}, mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "config", Value: "\n  debug = true"}, mock.Entry{Key: "logging.level", Value: "info"}}
	assert.Equal(t, expected, parseResult)

}

// mixed_dotted_and_regular_keys_expand_dotted - function:expand_dotted (level 0)
func TestMixedDottedAndRegularKeysExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_dotted_and_regular_keys_build_hierarchy - function:build_hierarchy (level 0)
func TestMixedDottedAndRegularKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
