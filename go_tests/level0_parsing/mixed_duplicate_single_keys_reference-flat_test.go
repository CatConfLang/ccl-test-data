package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/mixed_duplicate_single_keys_reference-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// mixed_duplicate_single_keys_reference_build_hierarchy - function:build_hierarchy (level 0)
func TestMixedDuplicateSingleKeysReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_reference_get_list - function:get_list (level 0)
func TestMixedDuplicateSingleKeysReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// mixed_duplicate_single_keys_reference_parse - function:parse (level 0)
func TestMixedDuplicateSingleKeysReferenceParse(t *testing.T) {

	ccl := mock.New()
	input := `ports = 80
ports = 443
host = localhost`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "ports", Value: "80"}, mock.Entry{Key: "ports", Value: "443"}, mock.Entry{Key: "host", Value: "localhost"}}
	assert.Equal(t, expected, parseResult)

}
