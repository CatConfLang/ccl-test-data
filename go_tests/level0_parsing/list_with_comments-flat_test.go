package level0_parsing_test

import (
	"testing"

	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_with_comments-flat.json
// Suite: Generated Flat Format
// Version: 1.0

// list_with_comments_parse - function:parse (level 0)
func TestListWithCommentsParse(t *testing.T) {

	ccl := mock.New()
	input := `servers = web1
/= Production servers
servers = web2
servers = web3
/= End of list`

	// Declare variables for reuse across validations

	var err error

	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "servers", Value: "web1"}, mock.Entry{Key: "/", Value: "Production servers"}, mock.Entry{Key: "servers", Value: "web2"}, mock.Entry{Key: "servers", Value: "web3"}, mock.Entry{Key: "/", Value: "End of list"}}
	assert.Equal(t, expected, parseResult)

}

// list_with_comments_build_hierarchy - function:build_hierarchy (level 0)
func TestListWithCommentsBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}

// list_with_comments_get_list - function:get_list (level 0)
func TestListWithCommentsGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}
