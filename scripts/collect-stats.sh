#!/bin/bash
set -e

# CCL Test Statistics Collector
# Auto-discovers and categorizes test files using their metadata

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

# Discover all test files (excluding schema/meta files)
ALL_TEST_FILES=($(fd -t f "\.json$" tests/ --max-depth 1 -E "schema.json" -E "pretty-print-schema.json"))

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

# Get file description and suite name
get_description() {
    local file="$1"
    jq -r '.description // "No description"' "$file" 2>/dev/null || echo "No description"
}

get_suite_name() {
    local file="$1"
    jq -r '.suite // "Unknown"' "$file" 2>/dev/null || echo "Unknown"
}

# Categorize test file based on suite name and description
categorize_file() {
    local file="$1"
    local suite=$(get_suite_name "$file")
    local desc=$(get_description "$file")
    local basename_file=$(basename "$file" .json)
    
    # Categorize based on suite name and description content
    if [[ "$suite" =~ (Essential|Comprehensive|Object) ]] || 
       [[ "$desc" =~ (essential|comprehensive|object-construction) ]] || 
       [[ "$basename_file" =~ (essential-parsing|comprehensive-parsing|object-construction) ]]; then
        echo "core"
    elif [[ "$suite" =~ (Dotted|Comment|Processing|Typed) ]] || 
         [[ "$desc" =~ (dotted|comment|processing|typed|composition|filtering) ]] ||
         [[ "$basename_file" =~ (dotted-keys|comments|processing|typed-access) ]]; then
        echo "features"
    elif [[ "$suite" =~ (Error) ]] || [[ "$desc" =~ (error|malformed) ]]; then
        echo "integration"
    elif [[ "$suite" =~ (Pretty) ]] || [[ "$desc" =~ (pretty|formatting|round-trip) ]]; then
        echo "utilities"
    else
        # Default fallback based on filename if metadata unclear
        if [[ "$basename_file" =~ essential|comprehensive|object ]]; then
            echo "core"
        elif [[ "$basename_file" =~ dotted|comment|processing|typed ]]; then
            echo "features"
        elif [[ "$basename_file" =~ error ]]; then
            echo "integration"
        elif [[ "$basename_file" =~ pretty ]]; then
            echo "utilities"
        else
            echo "other"
        fi
    fi
}

# Initialize counters
declare -A CATEGORY_COUNTS
declare -A CATEGORY_FILES
declare -A FILE_DESCRIPTIONS
declare -A CATEGORY_TOTALS=(["core"]=0 ["features"]=0 ["integration"]=0 ["utilities"]=0 ["other"]=0)

# Process all test files
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 244 "Categorizing test files by metadata..."
fi

for file in "${ALL_TEST_FILES[@]}"; do
    if [[ -n "$file" && -f "$file" ]]; then
        count=$(count_tests "$file")
        if [[ "$count" -gt 0 ]]; then  # Only process files with actual tests
            basename_file=$(basename "$file" .json)
            category=$(categorize_file "$file")
            
            # Update category totals
            CATEGORY_TOTALS["$category"]=$((${CATEGORY_TOTALS["$category"]} + count))
            
            # Store file info
            CATEGORY_COUNTS["${category}_$basename_file"]=$count
            CATEGORY_FILES["${category}_$basename_file"]="$file"
            FILE_DESCRIPTIONS["${category}_$basename_file"]=$(get_description "$file")
            
            if [[ "$USE_GUM" == "true" ]]; then
                # Capitalize category name for display
                category_display=$(echo "$category" | sed 's/^./\U&/')
                gum style --foreground 244 "  $category_display: $basename_file ($count tests)"
            fi
        fi
    fi
done

# Calculate total
TOTAL_COUNT=0
for total in "${CATEGORY_TOTALS[@]}"; do
    TOTAL_COUNT=$((TOTAL_COUNT + total))
done

# Show interactive summary
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 212 --bold "📊 Test Statistics Summary"
    echo
    gum style --foreground 33 --bold "Feature-Based Structure (Flat):"
    gum style --padding "0 2" --foreground 244 \
        "Core: ${CATEGORY_TOTALS['core']} tests" \
        "Features: ${CATEGORY_TOTALS['features']} tests" \
        "Integration: ${CATEGORY_TOTALS['integration']} tests" \
        "Utilities: ${CATEGORY_TOTALS['utilities']} tests"
    
    if [[ "${CATEGORY_TOTALS['other']}" -gt 0 ]]; then
        gum style --padding "0 2" --foreground 244 "Other: ${CATEGORY_TOTALS['other']} tests"
    fi
    
    echo
    gum style --foreground 33 --bold "Total: $TOTAL_COUNT tests"
    echo
    gum style --foreground 33 --bold "JSON Output:"
fi

# Generate JSON output
cat << EOF
{
  "structure": "flat",
  "categories": {
    "core": {
      "total": ${CATEGORY_TOTALS['core']},
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
      "total": ${CATEGORY_TOTALS['features']},
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
      "total": ${CATEGORY_TOTALS['integration']},
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
      "total": ${CATEGORY_TOTALS['utilities']},
EOF

# Add utility file details  
first=true
for key in "${!CATEGORY_COUNTS[@]}"; do
    if [[ $key == utilities_* ]]; then
        name=${key#utilities_}
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

    }
  },
  "totals": {
    "core": ${CATEGORY_TOTALS['core']},
    "features": ${CATEGORY_TOTALS['features']},
    "integration": ${CATEGORY_TOTALS['integration']},
    "utilities": ${CATEGORY_TOTALS['utilities']},
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