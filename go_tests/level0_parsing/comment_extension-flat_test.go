package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/comment_extension-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// comment_extension_parse - function:parse feature:comments (level 0)
func TestCommentExtensionParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/= This is an environment section
port = 8080
serve = index.html
/= Database section
mode = in-memory
connections = 16`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "This is an environment section"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "serve", Value: "index.html"}, mock.Entry{Key: "/", Value: "Database section"}, mock.Entry{Key: "mode", Value: "in-memory"}, mock.Entry{Key: "connections", Value: "16"}}
	assert.Equal(t, expected, parseResult)

}


// comment_extension_filter - function:filter feature:comments (level 0)
func TestCommentExtensionFilter(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


