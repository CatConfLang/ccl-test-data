#!/usr/bin/env python3
"""
Script to migrate test JSON files from old descriptive tags to new feature-based tags.
"""

import json
import os
import sys
from pathlib import Path

# Tag migration mapping from docs/tag-migration.json
TAG_MIGRATION = {
    "basic": ["function:parse"],
    "essential-parsing": ["function:parse"], 
    "empty": [],
    "empty-value": ["feature:empty-keys"],
    "empty-key": ["feature:empty-keys"],
    "empty-multiline": ["feature:multiline", "feature:empty-keys"],
    
    "lists": ["feature:empty-keys"],
    "nested": ["function:make-objects"],
    "multiline": ["feature:multiline"],
    "whitespace": ["feature:whitespace"],
    "unicode": ["feature:unicode"],
    
    "complex": ["function:make-objects"],
    "mixed": ["function:make-objects", "feature:empty-keys"],
    "deep": ["function:make-objects"],
    
    "crlf": ["behavior:crlf-preserve"],
    "tabs": ["behavior:tabs-preserve"], 
    "spacing": ["behavior:strict-spacing"],
    
    "round-trip": ["function:pretty-print"],
    "canonical": ["function:pretty-print"],
    "deterministic": ["function:pretty-print"],
    "order": ["function:pretty-print"],
    
    "algebraic": ["function:compose"],
    "associativity": ["function:compose"],
    "semigroup": ["function:compose"],
    "monoid": ["function:compose"],
    "identity": ["function:compose"],
    "left": ["function:compose"],
    "right": ["function:compose"],
    
    "proposed-behavior": ["variant:proposed-behavior"],
    "reference-compliant-behavior": ["variant:reference-compliant"]
}

# Automatic conflicts based on behavioral choices
AUTOMATIC_CONFLICTS = {
    "behavior:crlf-preserve": ["behavior:crlf-normalize"],
    "behavior:tabs-preserve": ["behavior:tabs-to-spaces"],
    "behavior:strict-spacing": ["behavior:loose-spacing"],
    "variant:proposed-behavior": ["variant:reference-compliant"]
}

def infer_function_tags_from_validations(validations):
    """Infer required function tags from validation structure"""
    function_tags = set()
    
    for validation_name in validations.keys():
        if validation_name == "parse":
            function_tags.add("function:parse")
        elif validation_name == "parse_indented":
            function_tags.add("function:parse")
        elif validation_name == "filter":
            function_tags.add("function:filter")
        elif validation_name == "compose":
            function_tags.add("function:compose")
        elif validation_name == "expand_dotted":
            function_tags.add("function:expand-dotted")
        elif validation_name == "make_objects":
            function_tags.add("function:make-objects")
        elif validation_name == "get_string":
            function_tags.add("function:get-string")
        elif validation_name == "get_int":
            function_tags.add("function:get-int")
        elif validation_name == "get_bool":
            function_tags.add("function:get-bool")
        elif validation_name == "get_float":
            function_tags.add("function:get-float")
        elif validation_name == "get_list":
            function_tags.add("function:get-list")
        elif validation_name == "pretty_print":
            function_tags.add("function:pretty-print")
    
    return list(function_tags)

def infer_feature_tags_from_input(input_text):
    """Infer feature tags from test input content"""
    feature_tags = set()
    
    if "/=" in input_text:
        feature_tags.add("feature:comments")
    
    if "." in input_text and "=" in input_text:
        # Check for dotted keys (not just dots in values)
        lines = input_text.split('\n')
        for line in lines:
            if '=' in line:
                key = line.split('=', 1)[0].strip()
                if '.' in key:
                    feature_tags.add("feature:dotted-keys")
                    break
    
    if input_text.count('\n') > 0:
        # Check for empty keys (anonymous list items)
        lines = input_text.split('\n')
        for line in lines:
            if line.strip().startswith('='):
                feature_tags.add("feature:empty-keys")
                break
    
    if '\r' in input_text:
        feature_tags.add("behavior:crlf-preserve")
    
    return list(feature_tags)

def migrate_tags(old_tags):
    """Convert old tags to new feature-based tags"""
    new_tags = set()
    
    for old_tag in old_tags:
        if old_tag in TAG_MIGRATION:
            new_tags.update(TAG_MIGRATION[old_tag])
        else:
            # Keep unknown tags as-is (legacy)
            new_tags.add(old_tag)
    
    return list(new_tags)

def add_conflicts(tags):
    """Add conflict declarations for behavioral tags"""
    conflicts = set()
    
    for tag in tags:
        if tag in AUTOMATIC_CONFLICTS:
            conflicts.update(AUTOMATIC_CONFLICTS[tag])
    
    return list(conflicts) if conflicts else None

def migrate_test_case(test_case):
    """Migrate a single test case to new schema"""
    meta = test_case.get("meta", {})
    old_tags = meta.get("tags", [])
    validations = test_case.get("validations", {})
    input_text = test_case.get("input", "")
    
    # Start with migrated old tags
    new_tags = set(migrate_tags(old_tags))
    
    # Add function tags based on validations
    function_tags = infer_function_tags_from_validations(validations)
    new_tags.update(function_tags)
    
    # Add feature tags based on input content
    feature_tags = infer_feature_tags_from_input(input_text)
    new_tags.update(feature_tags)
    
    # Update meta with new tags
    meta["tags"] = sorted(list(new_tags))
    
    # Add conflicts if any behavioral tags present
    conflicts = add_conflicts(new_tags)
    if conflicts:
        meta["conflicts"] = sorted(conflicts)
    
    test_case["meta"] = meta
    return test_case

def migrate_test_file(file_path):
    """Migrate a single test file"""
    print(f"Migrating {file_path}...")
    
    with open(file_path, 'r') as f:
        data = json.load(f)
    
    # Migrate each test case
    for test_case in data.get("tests", []):
        migrate_test_case(test_case)
    
    # Write back with pretty formatting
    with open(file_path, 'w') as f:
        json.dump(data, f, indent=2, ensure_ascii=False)
        f.write('\n')  # Add trailing newline

def main():
    """Main migration function"""
    tests_dir = Path("tests")
    
    if not tests_dir.exists():
        print("Error: tests/ directory not found")
        sys.exit(1)
    
    # Find all JSON test files (exclude schema.json)
    test_files = []
    for json_file in tests_dir.glob("*.json"):
        if json_file.name != "schema.json":
            test_files.append(json_file)
    
    if not test_files:
        print("No test files found in tests/ directory")
        sys.exit(1)
    
    print(f"Found {len(test_files)} test files to migrate")
    
    # Migrate each file
    for test_file in sorted(test_files):
        try:
            migrate_test_file(test_file)
        except Exception as e:
            print(f"Error migrating {test_file}: {e}")
            sys.exit(1)
    
    print(f"Successfully migrated {len(test_files)} test files")

if __name__ == "__main__":
    main()