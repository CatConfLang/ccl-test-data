#!/bin/bash
set -e

# CCL README Updater
# Updates README.md with current test counts and file information

cd "$(dirname "$0")/.."

# Parse arguments
AUTO_YES=false
for arg in "$@"; do
    case $arg in
        --yes|-y)
            AUTO_YES=true
            ;;
    esac
done

# Check if we should use interactive mode (not piped and not auto-yes)
if [[ -t 1 ]] && [[ "$AUTO_YES" == "false" ]]; then
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
                sd)
                    gum style --foreground 244 "Install with:" --padding "1 0 0 0"
                    gum style --foreground 33 --padding "0 2" "mise install" "cargo install sd" "brew install sd"
                    ;;
            esac
        else
            echo "❌ '$tool' command not found. Install with:"
            case "$tool" in
                sd)
                    echo "   mise install   # (if using mise)"
                    echo "   cargo install sd   # (if using cargo)"  
                    echo "   brew install sd   # (if using homebrew)"
                    ;;
            esac
        fi
        exit 1
    fi
}

check_tool sd

# Show initial message
if [[ "$USE_GUM" == "true" ]]; then
    gum style --foreground 212 --bold "📝 CCL README Updater"
    gum style --foreground 244 "Updating README.md with current test statistics"
else
    echo "📝 Updating README.md with current test statistics..."
fi

# Collect current stats
if [[ "$USE_GUM" == "true" ]]; then
    echo
    STATS=$(gum spin --spinner dot --title "Collecting current test statistics" -- sh -c './scripts/collect-stats.sh | tail -n +2')
else
    STATS=$(./scripts/collect-stats.sh | tail -n +2)
fi

# Extract values using jq  
if [[ "$USE_GUM" == "true" ]]; then
    LEVEL_1=$(gum spin --spinner dot --title "Extracting level 1 count" -- sh -c "echo '$STATS' | jq -r '.counts.level1'")
    LEVEL_2_TOTAL=$(echo "$STATS" | jq -r '.counts.level2')
    LEVEL_2_TESTS=$(echo "$STATS" | jq -r '.counts.level2_tests')
    LEVEL_2_COMP=$(echo "$STATS" | jq -r '.counts.level2_composition')
    LEVEL_3=$(echo "$STATS" | jq -r '.counts.level3')
    LEVEL_4=$(echo "$STATS" | jq -r '.counts.level4')
    ERROR_COUNT=$(echo "$STATS" | jq -r '.counts.errors')
    PRETTY_COUNT=$(echo "$STATS" | jq -r '.counts.pretty_print')
    TOTAL=$(echo "$STATS" | jq -r '.counts.total')
    
    echo
    gum style --foreground 33 --bold "Current Test Counts:"
    gum style --padding "0 2" --foreground 244 \
        "Level 1: $LEVEL_1" \
        "Level 2: $LEVEL_2_TOTAL ($LEVEL_2_TESTS + $LEVEL_2_COMP composition)" \
        "Level 3: $LEVEL_3" \
        "Level 4: $LEVEL_4" \
        "Errors: $ERROR_COUNT" \
        "Pretty Print: $PRETTY_COUNT" \
        "Total: $TOTAL"
else
    LEVEL_1=$(echo "$STATS" | jq -r '.counts.level1')
    LEVEL_2_TOTAL=$(echo "$STATS" | jq -r '.counts.level2')
    LEVEL_2_TESTS=$(echo "$STATS" | jq -r '.counts.level2_tests')
    LEVEL_2_COMP=$(echo "$STATS" | jq -r '.counts.level2_composition')
    LEVEL_3=$(echo "$STATS" | jq -r '.counts.level3')
    LEVEL_4=$(echo "$STATS" | jq -r '.counts.level4')
    ERROR_COUNT=$(echo "$STATS" | jq -r '.counts.errors')
    PRETTY_COUNT=$(echo "$STATS" | jq -r '.counts.pretty_print')
    TOTAL=$(echo "$STATS" | jq -r '.counts.total')
    
    echo "📊 Current counts: L1:$LEVEL_1 L2:$LEVEL_2_TOTAL L3:$LEVEL_3 L4:$LEVEL_4 Err:$ERROR_COUNT Pretty:$PRETTY_COUNT Total:$TOTAL"
fi

# Store original content for diff comparison
cp README.md README.md.tmp

# Apply updates with progress indicators
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 33 --bold "Applying Updates:"
    
    gum spin --spinner dot --title "Updating file references" -- \
        sh -c "sd \"level-1-parsing\.json\\\` \\(\\d+ tests\\)\" \"level-1-parsing.json\\\` ($LEVEL_1 tests)\" README.md && \
               sd \"level-2-processing\.json\\\` \\(\\d+ tests \\+ \\d+ composition tests\\)\" \"level-2-processing.json\\\` ($LEVEL_2_TESTS tests + $LEVEL_2_COMP composition tests)\" README.md && \
               sd \"level-3-objects\.json\\\` \\(\\d+ tests\\)\" \"level-3-objects.json\\\` ($LEVEL_3 tests)\" README.md && \
               sd \"level-4-typed\.json\\\` \\(\\d+ tests\\)\" \"level-4-typed.json\\\` ($LEVEL_4 tests)\" README.md"
    
    gum spin --spinner dot --title "Updating total count" -- \
        sd "includes \*\*\d+ test cases\*\* total:" "includes **$TOTAL test cases** total:" README.md
    
    gum spin --spinner dot --title "Updating level sections" -- \
        sh -c "sd \"### Level 1 \\(Core Parsing\\) - \\d+ tests\" \"### Level 1 (Core Parsing) - $LEVEL_1 tests\" README.md && \
               sd \"### Level 2 \\(Processing\\) - \\d+ tests\" \"### Level 2 (Processing) - $LEVEL_2_TOTAL tests\" README.md && \
               sd \"### Level 3 \\(Objects\\) - \\d+ tests\" \"### Level 3 (Objects) - $LEVEL_3 tests\" README.md && \
               sd \"### Level 4 \\(Typed\\) - \\d+ tests\" \"### Level 4 (Typed) - $LEVEL_4 tests\" README.md && \
               sd \"### Error Handling - \\d+ tests\" \"### Error Handling - $ERROR_COUNT tests\" README.md"
    
    gum spin --spinner dot --title "Updating file tree section" -- \
        sh -c "sd \"level-1-parsing\.json.*# Level 1: Core parsing tests \\(\\d+ tests\\)\" \"level-1-parsing.json      # Level 1: Core parsing tests ($LEVEL_1 tests)\" README.md && \
               sd \"level-2-processing\.json.*# Level 2: Processing and composition tests \\(\\d+ tests\\)\" \"level-2-processing.json   # Level 2: Processing and composition tests ($LEVEL_2_TOTAL tests)\" README.md && \
               sd \"level-3-objects\.json.*# Level 3: Object construction tests \\(\\d+ tests\\)\" \"level-3-objects.json      # Level 3: Object construction tests ($LEVEL_3 tests)\" README.md && \
               sd \"level-4-typed\.json.*# Level 4: Typed parsing tests \\(\\d+ tests\\)\" \"level-4-typed.json        # Level 4: Typed parsing tests ($LEVEL_4 tests)\" README.md && \
               sd \"errors\.json.*# Error handling tests \\(\\d+ tests\\)\" \"errors.json               # Error handling tests ($ERROR_COUNT tests)\" README.md"
else
    # Update file references in header sections (sd has much cleaner syntax!)
    sd "level-1-parsing\.json\` \(\d+ tests\)" "level-1-parsing.json\` ($LEVEL_1 tests)" README.md
    sd "level-2-processing\.json\` \(\d+ tests \+ \d+ composition tests\)" "level-2-processing.json\` ($LEVEL_2_TESTS tests + $LEVEL_2_COMP composition tests)" README.md  
    sd "level-3-objects\.json\` \(\d+ tests\)" "level-3-objects.json\` ($LEVEL_3 tests)" README.md
    sd "level-4-typed\.json\` \(\d+ tests\)" "level-4-typed.json\` ($LEVEL_4 tests)" README.md

    # Update total count in main description
    sd "includes \*\*\d+ test cases\*\* total:" "includes **$TOTAL test cases** total:" README.md

    # Update individual level sections
    sd "### Level 1 \(Core Parsing\) - \d+ tests" "### Level 1 (Core Parsing) - $LEVEL_1 tests" README.md
    sd "### Level 2 \(Processing\) - \d+ tests" "### Level 2 (Processing) - $LEVEL_2_TOTAL tests" README.md
    sd "### Level 3 \(Objects\) - \d+ tests" "### Level 3 (Objects) - $LEVEL_3 tests" README.md
    sd "### Level 4 \(Typed\) - \d+ tests" "### Level 4 (Typed) - $LEVEL_4 tests" README.md
    sd "### Error Handling - \d+ tests" "### Error Handling - $ERROR_COUNT tests" README.md

    # Update file tree section  
    sd "level-1-parsing\.json.*# Level 1: Core parsing tests \(\d+ tests\)" "level-1-parsing.json      # Level 1: Core parsing tests ($LEVEL_1 tests)" README.md
    sd "level-2-processing\.json.*# Level 2: Processing and composition tests \(\d+ tests\)" "level-2-processing.json   # Level 2: Processing and composition tests ($LEVEL_2_TOTAL tests)" README.md
    sd "level-3-objects\.json.*# Level 3: Object construction tests \(\d+ tests\)" "level-3-objects.json      # Level 3: Object construction tests ($LEVEL_3 tests)" README.md
    sd "level-4-typed\.json.*# Level 4: Typed parsing tests \(\d+ tests\)" "level-4-typed.json        # Level 4: Typed parsing tests ($LEVEL_4 tests)" README.md
    sd "errors\.json.*# Error handling tests \(\d+ tests\)" "errors.json               # Error handling tests ($ERROR_COUNT tests)" README.md
fi

# Check if anything actually changed
if cmp -s README.md README.md.tmp; then
    if [[ "$USE_GUM" == "true" ]]; then
        echo
        gum style --foreground 46 "✅ README.md is already up to date"
    else
        echo "✅ README.md is already up to date"
    fi
    rm README.md.tmp
else
    # Show preview and get confirmation in interactive mode
    if [[ "$USE_GUM" == "true" ]] && [[ "$AUTO_YES" == "false" ]]; then
        echo
        gum style --foreground 33 --bold "🔄 Preview of Changes:"
        
        # Create a styled diff preview
        DIFF_OUTPUT=$(diff -u README.md.tmp README.md || true)
        if [[ -n "$DIFF_OUTPUT" ]]; then
            echo "$DIFF_OUTPUT" | head -20 | gum style --border normal --padding "1 2" --border-foreground 244
            
            if [[ $(echo "$DIFF_OUTPUT" | wc -l) -gt 20 ]]; then
                gum style --foreground 244 "(showing first 20 lines of diff)"
            fi
        fi
        
        echo
        if gum confirm "Apply these changes to README.md?"; then
            gum style --foreground 46 "✅ README.md updated successfully"
            rm README.md.tmp
        else
            # Restore original file
            mv README.md.tmp README.md
            gum style --foreground 208 "❌ Changes cancelled - README.md restored"
            exit 1
        fi
    else
        # Non-interactive mode - just apply changes and show summary
        if [[ "$USE_GUM" == "true" ]]; then
            echo
            gum style --foreground 46 "✅ README.md updated successfully"
            
            # Show what changed
            gum style --foreground 33 --bold "🔄 Changes Applied:"
            diff -u README.md.tmp README.md | head -10 | gum style --border normal --padding "1 2" --border-foreground 244 || true
        else
            echo "✅ README.md updated successfully"
            
            # Show what changed
            echo "🔄 Changes made:"
            diff -u README.md.tmp README.md | head -10 || true
        fi
        rm README.md.tmp
    fi
fi