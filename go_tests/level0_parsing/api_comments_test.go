package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/api_comments.json
// Suite:
// Version:

// comment_extension_parse - function:parse (level 0)
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

// comment_syntax_slash_equals_parse - function:parse (level 0)
func TestCommentSyntaxSlashEqualsParse(t *testing.T) {

	ccl := mock.New()
	input := `/= this is a comment`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "/", Value: "this is a comment"}}
	assert.Equal(t, expected, parseResult)

}

// comment_syntax_slash_equals_filter - function:filter feature:comments (level 0)
func TestCommentSyntaxSlashEqualsFilter(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// section_headers_with_comments_parse - function:parse (level 0)
func TestSectionHeadersWithCommentsParse(t *testing.T) {

	ccl := mock.New()
	input := `== Database Config ==
/= Connection settings
host = localhost
=== Cache Config ===
/= Redis configuration
port = 6379`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database Config =="}, mock.Entry{Key: "/", Value: "Connection settings"}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "", Value: "== Cache Config ==="}, mock.Entry{Key: "/", Value: "Redis configuration"}, mock.Entry{Key: "port", Value: "6379"}}
	assert.Equal(t, expected, parseResult)

}

// section_headers_with_comments_filter - function:filter feature:comments (level 0)
func TestSectionHeadersWithCommentsFilter(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
