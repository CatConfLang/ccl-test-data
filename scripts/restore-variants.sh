#!/bin/bash

# Script to restore variants fields to source test files based on old test format
# Usage: ./scripts/restore-variants.sh

set -e

echo "🔄 Restoring variants fields to source test files..."

# Extract complete variant mappings from old tests
echo "📊 Extracting variant mappings from historical data..."

# Create temporary mapping files
TEMP_DIR=$(mktemp -d)
API_LIST_VARIANTS="$TEMP_DIR/api_list_variants.txt"
PROPERTY_RT_VARIANTS="$TEMP_DIR/property_rt_variants.txt"
API_ADV_VARIANTS="$TEMP_DIR/api_adv_variants.txt"
API_TYPED_VARIANTS="$TEMP_DIR/api_typed_variants.txt"

# Extract variants from old tests (commit 8a8ebad~1)
echo "Extracting api-list-access.json variants..."
git show 8a8ebad~1:tests/api-list-access.json | jq -r '.tests[] | select(.meta.tags | contains(["variant:proposed_behavior"]) or contains(["variant:reference_compliant"])) | "\(.name):\([.meta.tags[] | select(startswith("variant:"))] | .[0] | sub("variant:"; ""))"' > "$API_LIST_VARIANTS"

echo "Extracting property-round-trip.json variants..."
git show 8a8ebad~1:tests/property-round-trip.json | jq -r '.tests[] | select(.meta.tags | contains(["variant:proposed_behavior"]) or contains(["variant:reference_compliant"])) | "\(.name):\([.meta.tags[] | select(startswith("variant:"))] | .[0] | sub("variant:"; ""))"' > "$PROPERTY_RT_VARIANTS"

echo "Extracting api-advanced-processing.json variants..."
git show 8a8ebad~1:tests/api-advanced-processing.json | jq -r '.tests[] | select(.meta.tags | contains(["variant:proposed_behavior"]) or contains(["variant:reference_compliant"])) | "\(.name):\([.meta.tags[] | select(startswith("variant:"))] | .[0] | sub("variant:"; ""))"' > "$API_ADV_VARIANTS"

echo "Extracting api-typed-access.json variants..."
git show 8a8ebad~1:tests/api-typed-access.json | jq -r '.tests[] | select(.meta.tags | contains(["variant:proposed_behavior"]) or contains(["variant:reference_compliant"])) | "\(.name):\([.meta.tags[] | select(startswith("variant:"))] | .[0] | sub("variant:"; ""))"' > "$API_TYPED_VARIANTS"

# Function to add variants to a test file
add_variants_to_file() {
    local file_path="$1"
    local variant_mapping_file="$2"
    local backup_file="${file_path}.backup"

    echo "🔄 Processing $file_path..."

    # Create backup
    cp "$file_path" "$backup_file"

    # Read the variant mappings and apply them
    while IFS=':' read -r test_name variant_type; do
        if [ -n "$test_name" ] && [ -n "$variant_type" ]; then
            echo "  Adding $variant_type variant to $test_name"

            # Use jq to add variants field to the specific test
            jq --arg test_name "$test_name" --arg variant_type "$variant_type" '
                .tests |= map(
                    if .name == $test_name then
                        . + {"variants": [$variant_type]}
                    else
                        .
                    end
                )
            ' "$file_path" > "${file_path}.tmp" && mv "${file_path}.tmp" "$file_path"
        fi
    done < "$variant_mapping_file"

    # Validate the updated file
    if jv schemas/source-format.json "$file_path" > /dev/null 2>&1; then
        echo "  ✅ $file_path updated and validated successfully"
        rm "$backup_file"
    else
        echo "  ❌ Validation failed for $file_path, restoring backup"
        mv "$backup_file" "$file_path"
        return 1
    fi
}

# Apply variants to each file
echo ""
echo "📝 Applying variants to source test files..."

# Apply to api_list_access.json
if [ -s "$API_LIST_VARIANTS" ]; then
    add_variants_to_file "source_tests/core/api_list_access.json" "$API_LIST_VARIANTS"
fi

# Apply to property_round_trip.json
if [ -s "$PROPERTY_RT_VARIANTS" ]; then
    add_variants_to_file "source_tests/core/property_round_trip.json" "$PROPERTY_RT_VARIANTS"
fi

# Apply to api_advanced_processing.json
if [ -s "$API_ADV_VARIANTS" ]; then
    add_variants_to_file "source_tests/core/api_advanced_processing.json" "$API_ADV_VARIANTS"
fi

# Apply to api_typed_access.json
if [ -s "$API_TYPED_VARIANTS" ]; then
    add_variants_to_file "source_tests/core/api_typed_access.json" "$API_TYPED_VARIANTS"
fi

# Show summary
echo ""
echo "📊 Summary of variants added:"
echo "api_list_access.json: $(wc -l < "$API_LIST_VARIANTS") variants"
echo "property_round_trip.json: $(wc -l < "$PROPERTY_RT_VARIANTS") variants"
echo "api_advanced_processing.json: $(wc -l < "$API_ADV_VARIANTS") variants"
echo "api_typed_access.json: $(wc -l < "$API_TYPED_VARIANTS") variants"

# Final validation of all files
echo ""
echo "🔍 Final validation of all source test files..."
VALIDATION_PASSED=true

for file in source_tests/core/*.json source_tests/experimental/*.json; do
    if [ -f "$file" ]; then
        if jv schemas/source-format.json "$file" > /dev/null 2>&1; then
            echo "  ✅ $(basename "$file") - valid"
        else
            echo "  ❌ $(basename "$file") - validation failed"
            VALIDATION_PASSED=false
        fi
    fi
done

# Cleanup
rm -rf "$TEMP_DIR"

if [ "$VALIDATION_PASSED" = true ]; then
    echo ""
    echo "🎉 All variants restored successfully!"
    echo "📋 Next steps:"
    echo "  1. Review the changes: git diff"
    echo "  2. Run tests: just test"
    echo "  3. Commit changes: git add -A && git commit -m 'feat: restore variants fields from historical test data'"
else
    echo ""
    echo "⚠️  Some files failed validation. Please check the errors above."
    exit 1
fi