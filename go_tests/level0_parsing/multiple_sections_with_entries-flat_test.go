package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/multiple_sections_with_entries-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// multiple_sections_with_entries_parse - function:parse (level 0)
func TestMultipleSectionsWithEntriesParse(t *testing.T) {

	ccl := mock.New()
	input := `== Database ==
host = localhost

=== Cache ===
redis = enabled

== Logging ==
level = info`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database =="}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "", Value: "== Cache ==="}, mock.Entry{Key: "redis", Value: "enabled"}, mock.Entry{Key: "", Value: "= Logging =="}, mock.Entry{Key: "level", Value: "info"}}
	assert.Equal(t, expected, parseResult)

}
