package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_with_booleans_reference-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// list_with_booleans_reference_get_list - function:get_list (level 0)
func TestListWithBooleansReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_booleans_reference_parse - function:parse (level 0)
func TestListWithBooleansReferenceParse(t *testing.T) {

	ccl := mock.New()
	input := `flags = true
flags = false
flags = yes
flags = no`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "flags", Value: "true"}, mock.Entry{Key: "flags", Value: "false"}, mock.Entry{Key: "flags", Value: "yes"}, mock.Entry{Key: "flags", Value: "no"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_booleans_reference_build_hierarchy - function:build_hierarchy (level 0)
func TestListWithBooleansReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
