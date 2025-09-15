package level0_parsing_test

import (
	"testing"
	
	"github.com/ccl-test-data/test-runner/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generated from generated_tests/section_headers_with_colons-flat.json
// Suite: Generated Flat Format
// Version: 1.0



// section_headers_with_colons_parse - function:parse feature:empty_keys (level 0)
func TestSectionHeadersWithColonsParse(t *testing.T) {
	
	
	ccl := mock.New()
	input := `== Database: Production ==
host = db.prod.com
=== Cache: Redis Config ===
port = 6379`
	
	
	
	
	// Declare variables for reuse across validations
	
	
	
	var err error
	
	// Parse validation
	parseResult, err := ccl.Parse(input)
	require.NoError(t, err)
	expected := []mock.Entry{mock.Entry{Key: "", Value: "= Database: Production =="}, mock.Entry{Key: "host", Value: "db.prod.com"}, mock.Entry{Key: "", Value: "== Cache: Redis Config ==="}, mock.Entry{Key: "port", Value: "6379"}}
	assert.Equal(t, expected, parseResult)

}


