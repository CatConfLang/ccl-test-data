package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_errors.json
// Suite: Flat Format
// Version: 1.0



// whitespace_only_error_ocaml_reference_parse - function:parse (level 0)
func TestWhitespaceOnlyErrorOcamlReferenceParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `   `
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{}
	assert.Equal(t, expected, parseResult)

}


