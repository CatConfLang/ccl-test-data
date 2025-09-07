#!/bin/bash
set -e

# CCL README Updater
# Updates README.md with current test counts and file information

cd "$(dirname "$0")/.."

# Check for required tools
if ! command -v sd &> /dev/null; then
    echo "❌ 'sd' command not found. Install with:"
    echo "   mise install   # (if using mise)"
    echo "   cargo install sd   # (if using cargo)"  
    echo "   brew install sd   # (if using homebrew)"
    exit 1
fi

echo "📝 Updating README.md with current test statistics..."

# Collect current stats (skip the first line which is a message)
STATS=$(./scripts/collect-stats.sh | tail -n +2)

# Extract values using jq
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

# Create backup
cp README.md README.md.bak

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

# Check if anything actually changed
if cmp -s README.md README.md.bak; then
    echo "✅ README.md is already up to date"
    rm README.md.bak
else
    echo "✅ README.md updated successfully"
    echo "💾 Backup saved as README.md.bak"
    
    # Show what changed
    echo "🔄 Changes made:"
    diff -u README.md.bak README.md | head -10 || true
fi