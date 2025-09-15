package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/section_headers_with_comments-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// section_headers_with_comments_parse - function:parse feature:comments feature:empty_keys (level 0)
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


// section_headers_with_comments_filter - function:filter feature:comments feature:empty_keys (level 0)
func TestSectionHeadersWithCommentsFilter(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


