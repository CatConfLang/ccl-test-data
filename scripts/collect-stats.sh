#!/bin/bash
set -e

# CCL Test Statistics Collector
# Extracts test counts and metadata from JSON test files

cd "$(dirname "$0")/.."

# Parse arguments
INTERACTIVE=false
if [[ "$1" == "--interactive" ]]; then
    INTERACTIVE=true
fi

# Check if output is being piped (not interactive)
if [[ -t 1 ]] && [[ "$INTERACTIVE" == "true" ]]; then
    USE_GUM=true
else
    USE_GUM=false
fi

# Check for required tools
check_tool() {
    local tool="$1"
    if ! command -v "$tool" &> /dev/null; then
        if [[ "$USE_GUM" == "true" ]] && command -v gum &> /dev/null; then
            gum style --foreground 196 "❌ '$tool' command not found."
            case "$tool" in
                jq)
                    gum style --foreground 244 "Install with:" --padding "1 0 0 0"
                    gum style --foreground 33 --padding "0 2" "mise install" "brew install jq" "apt install jq"
                    ;;
            esac
        else
            echo "❌ '$tool' command not found. Install with:"
            case "$tool" in
                jq)
                    echo "   mise install   # (if using mise)"
                    echo "   brew install jq   # (if using homebrew)"
                    echo "   apt install jq   # (if using apt)"
                    ;;
            esac
        fi
        exit 1
    fi
}

check_tool jq

# Show progress message
if [[ "$USE_GUM" == "true" ]]; then
    gum style --foreground 212 --bold "🔍 CCL Test Statistics Collector"
    gum spin --spinner dot --title "Analyzing test files..." -- sleep 0.5
else
    echo "🔍 Collecting CCL test statistics..."
fi

# Extract test counts from each file
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 244 "Extracting test counts..."
    LEVEL_1_COUNT=$(gum spin --spinner meter --title "Level 1 parsing tests" -- jq '.tests | length' tests/level-1-parsing.json)
    LEVEL_2_TESTS=$(gum spin --spinner meter --title "Level 2 processing tests" -- jq '.tests | length' tests/level-2-processing.json)
    LEVEL_2_COMP=$(gum spin --spinner meter --title "Level 2 composition tests" -- jq '.composition_tests | length' tests/level-2-processing.json)
    LEVEL_3_COUNT=$(gum spin --spinner meter --title "Level 3 object tests" -- jq '.tests | length' tests/level-3-objects.json)
    LEVEL_4_COUNT=$(gum spin --spinner meter --title "Level 4 typed tests" -- jq '.tests | length' tests/level-4-typed.json)
    ERROR_COUNT=$(gum spin --spinner meter --title "Error handling tests" -- jq '.tests | length' tests/errors.json)
    PRETTY_COUNT=$(gum spin --spinner meter --title "Pretty print tests" -- jq '.tests | length' tests/pretty-print.json)
else
    LEVEL_1_COUNT=$(jq '.tests | length' tests/level-1-parsing.json)
    LEVEL_2_TESTS=$(jq '.tests | length' tests/level-2-processing.json)  
    LEVEL_2_COMP=$(jq '.composition_tests | length' tests/level-2-processing.json)
    LEVEL_3_COUNT=$(jq '.tests | length' tests/level-3-objects.json)
    LEVEL_4_COUNT=$(jq '.tests | length' tests/level-4-typed.json)
    ERROR_COUNT=$(jq '.tests | length' tests/errors.json)
    PRETTY_COUNT=$(jq '.tests | length' tests/pretty-print.json)
fi

# Calculate totals
LEVEL_2_TOTAL=$((LEVEL_2_TESTS + LEVEL_2_COMP))
TOTAL_COUNT=$((LEVEL_1_COUNT + LEVEL_2_TOTAL + LEVEL_3_COUNT + LEVEL_4_COUNT + ERROR_COUNT))

# Extract metadata
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 244 "Extracting metadata..."
    LEVEL_1_DESC=$(gum spin --spinner dot --title "Level 1 description" -- jq -r '.description' tests/level-1-parsing.json)
    LEVEL_2_DESC=$(gum spin --spinner dot --title "Level 2 description" -- jq -r '.description' tests/level-2-processing.json)
    LEVEL_3_DESC=$(gum spin --spinner dot --title "Level 3 description" -- jq -r '.description' tests/level-3-objects.json)
    LEVEL_4_DESC=$(gum spin --spinner dot --title "Level 4 description" -- jq -r '.description' tests/level-4-typed.json)
    ERROR_DESC=$(gum spin --spinner dot --title "Error description" -- jq -r '.description' tests/errors.json)
    PRETTY_DESC=$(gum spin --spinner dot --title "Pretty print description" -- jq -r '.description' tests/pretty-print.json)
else
    LEVEL_1_DESC=$(jq -r '.description' tests/level-1-parsing.json)
    LEVEL_2_DESC=$(jq -r '.description' tests/level-2-processing.json)
    LEVEL_3_DESC=$(jq -r '.description' tests/level-3-objects.json)
    LEVEL_4_DESC=$(jq -r '.description' tests/level-4-typed.json)
    ERROR_DESC=$(jq -r '.description' tests/errors.json)
    PRETTY_DESC=$(jq -r '.description' tests/pretty-print.json)
fi

# Show interactive summary before JSON output
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 212 --bold "📊 Test Statistics Summary"
    echo
    gum style --foreground 33 --bold "Test Counts:"
    gum style --padding "0 2" --foreground 244 \
        "Level 1: $LEVEL_1_COUNT tests" \
        "Level 2: $LEVEL_2_TOTAL tests ($LEVEL_2_TESTS + $LEVEL_2_COMP composition)" \
        "Level 3: $LEVEL_3_COUNT tests" \
        "Level 4: $LEVEL_4_COUNT tests" \
        "Errors: $ERROR_COUNT tests" \
        "Pretty Print: $PRETTY_COUNT tests" \
        "" \
        "Total: $TOTAL_COUNT tests"
    echo
    gum style --foreground 33 --bold "JSON Output:"
fi

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