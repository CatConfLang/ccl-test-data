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
    CORE_TOTAL=$(gum spin --spinner dot --title "Extracting core count" -- sh -c "echo '$STATS' | jq -r '.totals.core'")
    FEATURES_TOTAL=$(echo "$STATS" | jq -r '.totals.features')
    INTEGRATION_TOTAL=$(echo "$STATS" | jq -r '.totals.integration')
    UTILITIES_TOTAL=$(echo "$STATS" | jq -r '.totals.utilities')
    TOTAL=$(echo "$STATS" | jq -r '.totals.overall')
    
    echo
    gum style --foreground 33 --bold "Current Test Counts:"
    gum style --padding "0 2" --foreground 244 \
        "Core: $CORE_TOTAL" \
        "Features: $FEATURES_TOTAL" \
        "Integration: $INTEGRATION_TOTAL" \
        "Utilities: $UTILITIES_TOTAL" \
        "Total: $TOTAL"
else
    CORE_TOTAL=$(echo "$STATS" | jq -r '.totals.core')
    FEATURES_TOTAL=$(echo "$STATS" | jq -r '.totals.features')
    INTEGRATION_TOTAL=$(echo "$STATS" | jq -r '.totals.integration')
    UTILITIES_TOTAL=$(echo "$STATS" | jq -r '.totals.utilities')
    TOTAL=$(echo "$STATS" | jq -r '.totals.overall')
    
    echo "📊 Current counts: Core:$CORE_TOTAL Features:$FEATURES_TOTAL Integration:$INTEGRATION_TOTAL Utilities:$UTILITIES_TOTAL Total:$TOTAL"
fi

# Store original content for diff comparison
cp README.md README.md.tmp

# Apply updates with progress indicators  
if [[ "$USE_GUM" == "true" ]]; then
    echo
    gum style --foreground 33 --bold "Applying Updates:"
    
    gum spin --spinner dot --title "Updating total count" -- \
        sd "includes \*\*\d+ test cases\*\* total:" "includes **$TOTAL test cases** total:" README.md
    
    gum spin --spinner dot --title "Updating category sections" -- \
        sh -c "sd \"### Core \\(\\d+ tests\\)\" \"### Core ($CORE_TOTAL tests)\" README.md && \
               sd \"### Features \\(\\d+ tests\\)\" \"### Features ($FEATURES_TOTAL tests)\" README.md && \
               sd \"### Integration \\(\\d+ tests\\)\" \"### Integration ($INTEGRATION_TOTAL tests)\" README.md && \
               sd \"### Utilities \\(\\d+ tests\\)\" \"### Utilities ($UTILITIES_TOTAL tests)\" README.md"
else
    # Update total count in main description
    sd "includes \*\*(\d+|null) test cases\*\* total:" "includes **$TOTAL test cases** total:" README.md

    # Update individual category sections
    sd "### Core \(\d+ tests\)" "### Core ($CORE_TOTAL tests)" README.md
    sd "### Features \(\d+ tests\)" "### Features ($FEATURES_TOTAL tests)" README.md  
    sd "### Integration \(\d+ tests\)" "### Integration ($INTEGRATION_TOTAL tests)" README.md
    sd "### Utilities \(\d+ tests\)" "### Utilities ($UTILITIES_TOTAL tests)" README.md
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