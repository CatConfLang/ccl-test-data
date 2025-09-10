#!/usr/bin/env python3
"""
Script to remove legacy descriptive tags from test JSON files, keeping only structured feature-based tags.
"""

import json
import os
import sys
from pathlib import Path

# Valid structured tag prefixes
VALID_TAG_PREFIXES = ["function:", "feature:", "behavior:", "variant:"]

def is_structured_tag(tag):
    """Check if a tag uses the new structured format"""
    return any(tag.startswith(prefix) for prefix in VALID_TAG_PREFIXES)

def clean_tags(tags):
    """Remove legacy tags, keeping only structured feature-based tags"""
    return [tag for tag in tags if is_structured_tag(tag)]

def clean_test_case(test_case):
    """Clean legacy tags from a single test case"""
    meta = test_case.get("meta", {})
    old_tags = meta.get("tags", [])
    
    # Filter to only structured tags
    new_tags = clean_tags(old_tags)
    
    # Update meta with cleaned tags
    meta["tags"] = sorted(new_tags)
    test_case["meta"] = meta
    return test_case

def clean_test_file(file_path):
    """Clean legacy tags from a single test file"""
    print(f"Cleaning {file_path}...")
    
    with open(file_path, 'r') as f:
        data = json.load(f)
    
    # Clean each test case
    for test_case in data.get("tests", []):
        clean_test_case(test_case)
    
    # Write back with pretty formatting
    with open(file_path, 'w') as f:
        json.dump(data, f, indent=2, ensure_ascii=False)
        f.write('\n')  # Add trailing newline

def main():
    """Main cleaning function"""
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
    
    print(f"Found {len(test_files)} test files to clean")
    
    # Clean each file
    for test_file in sorted(test_files):
        try:
            clean_test_file(test_file)
        except Exception as e:
            print(f"Error cleaning {test_file}: {e}")
            sys.exit(1)
    
    print(f"Successfully cleaned {len(test_files)} test files")

if __name__ == "__main__":
    main()