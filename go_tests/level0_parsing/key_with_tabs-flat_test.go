package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/key_with_tabs-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// key_with_tabs_parse - function:parse (level 0)
func TestKeyWithTabsParse(t *testing.T) {

	ccl := mock.New()
	input := `	key	=	value`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "key", Value: "\tvalue"}}
	assert.Equal(t, expected, parseResult)

}
