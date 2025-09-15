package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/ocaml_stress_test_original-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// ocaml_stress_test_original_get_string - function:get_string feature:comments feature:empty_keys (level 0)
func TestOcamlStressTestOriginalGetString(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// ocaml_stress_test_original_parse - function:parse feature:comments feature:empty_keys (level 0)
func TestOcamlStressTestOriginalParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/= This is a CCL document
title = CCL Example

database =
  enabled = true
  ports =
    = 8000
    = 8001
    = 8002
  limits =
    cpu = 1500mi
    memory = 10Gb

user =
  guestId = 42

user =
  login = chshersh
  createdAt = 2024-12-31`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "This is a CCL document"}, mock.Entry{Key: "title", Value: "CCL Example"}, mock.Entry{Key: "database", Value: "\n  enabled = true\n  ports =\n    = 8000\n    = 8001\n    = 8002\n  limits =\n    cpu = 1500mi\n    memory = 10Gb"}, mock.Entry{Key: "user", Value: "\n  guestId = 42"}, mock.Entry{Key: "user", Value: "\n  login = chshersh\n  createdAt = 2024-12-31"}}
	assert.Equal(t, expected, parseResult)

}


// ocaml_stress_test_original_build_hierarchy - function:build_hierarchy feature:comments feature:empty_keys (level 0)
func TestOcamlStressTestOriginalBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


