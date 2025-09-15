package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/realistic_stress_test-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// realistic_stress_test_parse - function:parse (level 0)
func TestRealisticStressTestParse(t *testing.T) {

	ccl := mock.New()
	input := `name = Dmitrii Kovanikov
login = chshersh
language = OCaml
date = 2024-05-25`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "name", Value: "Dmitrii Kovanikov"}, mock.Entry{Key: "login", Value: "chshersh"}, mock.Entry{Key: "language", Value: "OCaml"}, mock.Entry{Key: "date", Value: "2024-05-25"}}
	assert.Equal(t, expected, parseResult)

}
