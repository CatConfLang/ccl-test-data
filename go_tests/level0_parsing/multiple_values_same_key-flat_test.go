package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/multiple_values_same_key-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// multiple_values_same_key_parse - function:parse (level 0)
func TestMultipleValuesSameKeyParse(t *testing.T) {

	ccl := mock.New()
	input := `ports = 8000
ports = 8001
ports = 8002`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "ports", Value: "8000"}, mock.Entry{Key: "ports", Value: "8001"}, mock.Entry{Key: "ports", Value: "8002"}}
	assert.Equal(t, expected, parseResult)

}
