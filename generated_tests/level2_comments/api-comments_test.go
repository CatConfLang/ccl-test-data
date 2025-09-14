package level2_comments_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from tests/api-comments.json
// Suite: CCL Comment Filtering - Validation Format
// Version: 2.1
// Description: Comment syntax and filtering functionality - remove documentation keys from configuration


// comment_extension - function:parse feature:comments function:filter (level 2)
func TestCommentExtension(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/= This is an environment section
port = 8080
serve = index.html
/= Database section
mode = in-memory
connections = 16`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	var filterResult []mock.Entry
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "/", Value: "This is an environment section"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "serve", Value: "index.html"}, mock.Entry{Key: "/", Value: "Database section"}, mock.Entry{Key: "mode", Value: "in-memory"}, mock.Entry{Key: "connections", Value: "16"}}
	assert.Equal(t, expectedParse, parseResult)
	// Filter validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	filterResult = ccl.Filter(parseResult)
	expectedFilter := []mock.Entry{mock.Entry{Key: "/", Value: "This is an environment section"}, mock.Entry{Key: "port", Value: "8080"}, mock.Entry{Key: "serve", Value: "index.html"}, mock.Entry{Key: "/", Value: "Database section"}, mock.Entry{Key: "mode", Value: "in-memory"}, mock.Entry{Key: "connections", Value: "16"}}
	assert.Equal(t, expectedFilter, filterResult)

}


// comment_syntax_slash_equals - function:parse feature:comments function:filter (level 2)
func TestCommentSyntaxSlashEquals(t *testing.T) {
	
	
	ccl := mock.New()
	input := `/= this is a comment`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	var filterResult []mock.Entry
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "/", Value: "this is a comment"}}
	assert.Equal(t, expectedParse, parseResult)
	// Filter validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	filterResult = ccl.Filter(parseResult)
	expectedFilter := []mock.Entry{mock.Entry{Key: "/", Value: "this is a comment"}}
	assert.Equal(t, expectedFilter, filterResult)

}


// section_headers_with_comments - function:parse feature:comments feature:empty-keys function:filter (level 2)
func TestSectionHeadersWithComments(t *testing.T) {
	
	
	ccl := mock.New()
	input := `== Database Config ==
/= Connection settings
host = localhost
=== Cache Config ===
/= Redis configuration
port = 6379`
	
	
	
	
	// Declare variables for reuse across validations
	var parseResult []mock.Entry
	
	var filterResult []mock.Entry
	var err error
	
	// Parse validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	expectedParse := []mock.Entry{mock.Entry{Key: "", Value: "= Database Config =="}, mock.Entry{Key: "/", Value: "Connection settings"}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "", Value: "== Cache Config ==="}, mock.Entry{Key: "/", Value: "Redis configuration"}, mock.Entry{Key: "port", Value: "6379"}}
	assert.Equal(t, expectedParse, parseResult)
	// Filter validation
	parseResult, err = ccl.Parse(input)
	require.NoError(t, err)
	filterResult = ccl.Filter(parseResult)
	expectedFilter := []mock.Entry{mock.Entry{Key: "", Value: "= Database Config =="}, mock.Entry{Key: "/", Value: "Connection settings"}, mock.Entry{Key: "host", Value: "localhost"}, mock.Entry{Key: "", Value: "== Cache Config ==="}, mock.Entry{Key: "/", Value: "Redis configuration"}, mock.Entry{Key: "port", Value: "6379"}}
	assert.Equal(t, expectedFilter, filterResult)

}


