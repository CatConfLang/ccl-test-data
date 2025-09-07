#!/bin/bash
set -e

# CCL Test Statistics Collector
# Extracts test counts and metadata from JSON test files

cd "$(dirname "$0")/.."

# Check for required tools
if ! command -v jq &> /dev/null; then
    echo "❌ 'jq' command not found. Install with:"
    echo "   mise install   # (if using mise)"
    echo "   brew install jq   # (if using homebrew)"
    echo "   apt install jq   # (if using apt)"
    exit 1
fi

echo "🔍 Collecting CCL test statistics..."

# Extract test counts from each file
LEVEL_1_COUNT=$(jq '.tests | length' tests/level-1-parsing.json)
LEVEL_2_TESTS=$(jq '.tests | length' tests/level-2-processing.json)  
LEVEL_2_COMP=$(jq '.composition_tests | length' tests/level-2-processing.json)
LEVEL_3_COUNT=$(jq '.tests | length' tests/level-3-objects.json)
LEVEL_4_COUNT=$(jq '.tests | length' tests/level-4-typed.json)
ERROR_COUNT=$(jq '.tests | length' tests/errors.json)
PRETTY_COUNT=$(jq '.tests | length' tests/pretty-print.json)

# Calculate totals
LEVEL_2_TOTAL=$((LEVEL_2_TESTS + LEVEL_2_COMP))
TOTAL_COUNT=$((LEVEL_1_COUNT + LEVEL_2_TOTAL + LEVEL_3_COUNT + LEVEL_4_COUNT + ERROR_COUNT))

# Extract metadata
LEVEL_1_DESC=$(jq -r '.description' tests/level-1-parsing.json)
LEVEL_2_DESC=$(jq -r '.description' tests/level-2-processing.json)
LEVEL_3_DESC=$(jq -r '.description' tests/level-3-objects.json)
LEVEL_4_DESC=$(jq -r '.description' tests/level-4-typed.json)
ERROR_DESC=$(jq -r '.description' tests/errors.json)
PRETTY_DESC=$(jq -r '.description' tests/pretty-print.json)

# Output results
cat << EOF
{
  "counts": {
    "level1": $LEVEL_1_COUNT,
    "level2": $LEVEL_2_TOTAL,
    "level2_tests": $LEVEL_2_TESTS,
    "level2_composition": $LEVEL_2_COMP,
    "level3": $LEVEL_3_COUNT,
    "level4": $LEVEL_4_COUNT,
    "errors": $ERROR_COUNT,
    "pretty_print": $PRETTY_COUNT,
    "total": $TOTAL_COUNT
  },
  "descriptions": {
    "level1": "$LEVEL_1_DESC",
    "level2": "$LEVEL_2_DESC", 
    "level3": "$LEVEL_3_DESC",
    "level4": "$LEVEL_4_DESC",
    "errors": "$ERROR_DESC",
    "pretty_print": "$PRETTY_DESC"
  }
}
EOF