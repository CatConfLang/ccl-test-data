package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/equals_in_value_no_spaces-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// equals_in_value_no_spaces_parse - function:parse (level 0)
func TestEqualsInValueNoSpacesParse(t *testing.T) {

	ccl := mock.New()
	input := `a=b=c`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "a", Value: "b=c"}}
	assert.Equal(t, expected, parseResult)

}
