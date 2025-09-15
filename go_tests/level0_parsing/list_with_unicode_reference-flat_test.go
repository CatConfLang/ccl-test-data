package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/list_with_unicode_reference-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// list_with_unicode_reference_get_list - function:get_list feature:unicode (level 0)
func TestListWithUnicodeReferenceGetList(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


// list_with_unicode_reference_parse - function:parse feature:unicode (level 0)
func TestListWithUnicodeReferenceParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `names = 张三
names = José
names = François
names = العربية`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "names", Value: "张三"}, mock.Entry{Key: "names", Value: "José"}, mock.Entry{Key: "names", Value: "François"}, mock.Entry{Key: "names", Value: "العربية"}}
	assert.Equal(t, expected, parseResult)

}


// list_with_unicode_reference_build_hierarchy - function:build_hierarchy feature:unicode (level 0)
func TestListWithUnicodeReferenceBuildHierarchy(t *testing.T) {
	t.Skip("Test does not match run-only filter: [function:parse]")
}


