package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/multiple_dotted_keys-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// multiple_dotted_keys_parse - function:parse feature:experimental_dotted_keys (level 0)
func TestMultipleDottedKeysParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `database.host = localhost
database.port = 5432
app.name = MyApp`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "database.host", Value: "localhost"}, mock.Entry{Key: "database.port", Value: "5432"}, mock.Entry{Key: "app.name", Value: "MyApp"}}
	assert.Equal(t, expected, parseResult)

}


// multiple_dotted_keys_expand_dotted - function:expand_dotted feature:experimental_dotted_keys (level 0)
func TestMultipleDottedKeysExpandDotted(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// multiple_dotted_keys_build_hierarchy - function:build_hierarchy feature:experimental_dotted_keys (level 0)
func TestMultipleDottedKeysBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


