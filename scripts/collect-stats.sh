#!/bin/bash
set -e

# CCL Test Statistics Collector
# Auto-discovers and analyzes test files in feature-based structure

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
                fd)
                    gum style --foreground 244 "Install with:" --padding "1 0 0 0" 
                    gum style --foreground 33 --padding "0 2" "mise install" "brew install fd" "apt install fd-find"
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
                fd)
                    echo "   mise install   # (if using mise)"
                    echo "   brew install fd   # (if using homebrew)"
                    echo "   apt install fd-find   # (if using apt)"
                    ;;
            esac
        fi
        exit 1
    fi
}

check_tool jq
check_tool fd

# Show progress message
if [[ "$USE_GUM" == "true" ]]; then
    gum style --foreground 212 --bold "🔍 CCL Test Statistics Collector"
    gum spin --spinner dot --title "Discovering test files..." -- sleep 0.5
else
    echo "🔍 Collecting CCL test statistics..."
fi

# Auto-discover test files using fd
discover_tests() {
    local category="$1"
    fd -t f "\.json$" "tests/$category/" 2>/dev/null || true
}

# Count tests in a file
count_tests() {
    local file="$1"
    local regular_tests=$(jq '.tests | length' "$file" 2>/dev/null || echo "0")
    local composition_tests=$(jq '.composition_tests | length' "$file" 2>/dev/null || echo "0")
    
    if [[ "$composition_tests" != "0" ]]; then
        echo $((regular_tests + composition_tests))
    else
        echo "$regular_tests"
    fi
}

# Get file description
get_description() {
    local file="$1"
    jq -r '.description // "No description"' "$file" 2>/dev/null || echo "No description"
}

# Initialize counters
declare -A CATEGORY_COUNTS
declare -A CATEGORY_FILES
declare -A FILE_DESCRIPTIONS

# Discover and categorize files
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 244 "Discovering test files by category..."
fi

# Core functionality tests
CORE_FILES=($(discover_tests "core"))
CORE_TOTAL=0
for file in "${CORE_FILES[@]}"; do
    if [[ -n "$file" ]]; then
        count=$(count_tests "$file")
        CORE_TOTAL=$((CORE_TOTAL + count))
        basename_file=$(basename "$file" .json)
        CATEGORY_COUNTS["core_$basename_file"]=$count
        CATEGORY_FILES["core_$basename_file"]="$file"
        FILE_DESCRIPTIONS["core_$basename_file"]=$(get_description "$file")
        
        if [[ "$USE_GUM" == "true" ]]; then
            gum style --foreground 244 "  Core: $basename_file ($count tests)"
        fi
    fi
done

# Feature tests
FEATURE_FILES=($(discover_tests "features"))
FEATURES_TOTAL=0
for file in "${FEATURE_FILES[@]}"; do
    if [[ -n "$file" ]]; then
        count=$(count_tests "$file")
        FEATURES_TOTAL=$((FEATURES_TOTAL + count))
        basename_file=$(basename "$file" .json)
        CATEGORY_COUNTS["features_$basename_file"]=$count
        CATEGORY_FILES["features_$basename_file"]="$file"
        FILE_DESCRIPTIONS["features_$basename_file"]=$(get_description "$file")
        
        if [[ "$USE_GUM" == "true" ]]; then
            gum style --foreground 244 "  Feature: $basename_file ($count tests)"
        fi
    fi
done

# Integration tests
INTEGRATION_FILES=($(discover_tests "integration"))
INTEGRATION_TOTAL=0
for file in "${INTEGRATION_FILES[@]}"; do
    if [[ -n "$file" ]]; then
        count=$(count_tests "$file")
        INTEGRATION_TOTAL=$((INTEGRATION_TOTAL + count))
        basename_file=$(basename "$file" .json)
        CATEGORY_COUNTS["integration_$basename_file"]=$count
        CATEGORY_FILES["integration_$basename_file"]="$file"
        FILE_DESCRIPTIONS["integration_$basename_file"]=$(get_description "$file")
        
        if [[ "$USE_GUM" == "true" ]]; then
            gum style --foreground 244 "  Integration: $basename_file ($count tests)"
        fi
    fi
done

# Pretty print tests (special case)
PRETTY_FILE="tests/pretty-print.json"
PRETTY_COUNT=0
if [[ -f "$PRETTY_FILE" ]]; then
    PRETTY_COUNT=$(count_tests "$PRETTY_FILE")
    CATEGORY_COUNTS["utility_pretty-print"]=$PRETTY_COUNT
    CATEGORY_FILES["utility_pretty-print"]="$PRETTY_FILE"
    FILE_DESCRIPTIONS["utility_pretty-print"]=$(get_description "$PRETTY_FILE")
fi

# Calculate totals
TOTAL_COUNT=$((CORE_TOTAL + FEATURES_TOTAL + INTEGRATION_TOTAL + PRETTY_COUNT))

# Show interactive summary
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 212 --bold "📊 Test Statistics Summary"
    echo
    gum style --foreground 33 --bold "Feature-Based Structure:"
    gum style --padding "0 2" --foreground 244 \
        "Core: $CORE_TOTAL tests" \
        "Features: $FEATURES_TOTAL tests" \
        "Integration: $INTEGRATION_TOTAL tests" \
        "Utilities: $PRETTY_COUNT tests"
    echo
    gum style --foreground 33 --bold "Total: $TOTAL_COUNT tests"
    echo
    gum style --foreground 33 --bold "JSON Output:"
fi

# Generate JSON output
cat << EOF
{
  "structure": "feature-based",
  "categories": {
    "core": {
      "total": $CORE_TOTAL,
EOF

# Add core file details
first=true
for key in "${!CATEGORY_COUNTS[@]}"; do
    if [[ $key == core_* ]]; then
        name=${key#core_}
        count=${CATEGORY_COUNTS[$key]}
        if [[ $first == true ]]; then
            first=false
        else
            echo ","
        fi
        echo -n "      \"$name\": $count"
    fi
done

cat << EOF

    },
    "features": {
      "total": $FEATURES_TOTAL,
EOF

# Add feature file details
first=true
for key in "${!CATEGORY_COUNTS[@]}"; do
    if [[ $key == features_* ]]; then
        name=${key#features_}
        count=${CATEGORY_COUNTS[$key]}
        if [[ $first == true ]]; then
            first=false
        else
            echo ","
        fi
        echo -n "      \"$name\": $count"
    fi
done

cat << EOF

    },
    "integration": {
      "total": $INTEGRATION_TOTAL,
EOF

# Add integration file details  
first=true
for key in "${!CATEGORY_COUNTS[@]}"; do
    if [[ $key == integration_* ]]; then
        name=${key#integration_}
        count=${CATEGORY_COUNTS[$key]}
        if [[ $first == true ]]; then
            first=false
        else
            echo ","
        fi
        echo -n "      \"$name\": $count"
    fi
done

cat << EOF

    },
    "utilities": {
      "pretty-print": $PRETTY_COUNT
    }
  },
  "totals": {
    "core": $CORE_TOTAL,
    "features": $FEATURES_TOTAL,
    "integration": $INTEGRATION_TOTAL,
    "utilities": $PRETTY_COUNT,
    "overall": $TOTAL_COUNT
  },
  "descriptions": {
EOF

# Add descriptions
first=true
for key in "${!FILE_DESCRIPTIONS[@]}"; do
    desc=${FILE_DESCRIPTIONS[$key]}
    # Escape quotes in description
    desc=$(echo "$desc" | sed 's/"/\\"/g')
    if [[ $first == true ]]; then
        first=false
    else
        echo ","
    fi
    echo -n "    \"$key\": \"$desc\""
done

cat << EOF

  }
}
EOF