# CCL Test Suite

This directory contains language-agnostic test cases for CCL (Categorical Configuration Language) parsers organized by architectural level.

## 4-Level Architecture

CCL implementations should progress through these levels:

### Level 1: Entry Parsing (Core)
**File**: `ccl-entry-parsing.json` (48 tests)  
**API**: `parse(text) → Result(List(Entry), ParseError)`

Essential parsing functionality that every CCL implementation must support:
- Basic key-value parsing with whitespace handling
- Multiline values and continuation lines  
- Unicode support and line ending normalization
- Empty keys/values, equals-in-values handling
- Edge cases (EOF, mixed whitespace, etc.)

Many tests tagged `"redundant"` are variations of core functionality for comprehensive coverage.

### Level 2: Entry Processing (Extensions)
**File**: `ccl-entry-processing.json` (11 tests + composition tests)  
**API**: `filter_keys()`, entry composition functions

Processing operations on parsed entries:
- Comment filtering (keys starting with `/`)
- Duplicate key handling and composition
- Entry list merging and algebraic properties
- Decorative section parsing (future)

### Level 3: Object Construction (Hierarchical)
**File**: `ccl-object-construction.json` (8 tests)  
**API**: `make_objects(entries) → CCL`

Converting flat entry lists to nested object structures:
- Recursive parsing of nested values
- Duplicate key merging in object construction
- Empty key handling for list-style data
- Complex nested configuration examples

### Level 4: Typed Parsing (Language-Specific)
**File**: `ccl-typed-parsing-examples.json` (8 tests)  
**API**: `get_int()`, `get_bool()`, `get_typed_value()`, etc.

Type-aware extraction with validation:
- Smart type inference (integers, floats, booleans)
- Configurable parsing options
- Language-specific convenience functions
- Type safety and validation

### Error Handling (All Levels)
**File**: `ccl-errors.json` (5 tests)

Malformed input detection across all levels.

## Schema

All test files use the unified schema defined in `ccl-unified-test-schema.json`:

```json
{
  "name": "test_name",
  "input": "ccl input string",  
  "expected": [...],             // Level 1-2
  "expected_flat": [...],        // Level 3-4 flat parsing
  "expected_nested": {...},      // Level 3 object construction
  "expected_typed": {...},       // Level 4 typed values
  "expected_error": true,        // Error tests
  "meta": {
    "tags": ["basic", "redundant"],
    "level": 1,
    "feature": "parsing"
  }
}
```

## Usage for Implementers

### Quick Start (Level 1 Only)
Focus on `ccl-entry-parsing.json` tests without `"redundant"` tag:
```bash
# Filter out redundant tests for initial implementation
jq '.tests[] | select(.meta.tags | contains(["redundant"]) | not)' ccl-entry-parsing.json
```

### Full Compliance
Implement all levels progressively:
1. Level 1 core tests → basic CCL parsing
2. Level 2 processing → comments and composition  
3. Level 3 objects → hierarchical structures
4. Level 4 typing → language-specific conveniences

### Test Implementation Pattern
```python
# Example in Python
import json

def run_ccl_tests():
    # Level 1: Core parsing
    with open('ccl-entry-parsing.json') as f:
        level1_tests = json.load(f)['tests']
    
    for test in level1_tests:
        entries = parse(test['input'])
        expected = [Entry(e['key'], e['value']) for e in test['expected']]
        assert entries == expected, f"Failed: {test['name']}"
```

## Implementation Status

Implementations can declare their support level:

- **Level 1**: Basic CCL parser ✅
- **Level 2**: Comments + composition ✅ 
- **Level 3**: Object construction ✅
- **Level 4**: Typed parsing ✅

## Files

### Current (4-Level Architecture)
- `ccl-entry-parsing.json` - Level 1 core parsing tests
- `ccl-entry-processing.json` - Level 2 processing tests
- `ccl-object-construction.json` - Level 3 object tests
- `ccl-typed-parsing-examples.json` - Level 4 typed tests
- `ccl-errors.json` - Error handling tests
- `ccl-unified-test-schema.json` - JSON schema for all test files

### Legacy (Archived)
- `legacy-ccl-test-suite.json.backup` - Original monolithic test suite
- `legacy-ccl-test-suite-schema.json.backup` - Original schema

The legacy files are kept as backup but all tests have been migrated to the new 4-level structure with improved organization and comprehensive coverage.

## Test Coverage

The new 4-level architecture includes **80 test cases** total:

### Level 1 (Core Parsing) - 48 tests
- 18 essential tests covering core functionality
- 30 additional tests marked as `"redundant"` for comprehensive edge case coverage
- Includes basic parsing, whitespace handling, multiline values, unicode support

### Level 2 (Processing) - 11 tests  
- Comment filtering and composition behavior
- Duplicate key handling and algebraic properties
- Entry list operations and merging

### Level 3 (Objects) - 8 tests
- Nested object construction from flat entries
- Complex hierarchical configurations
- List handling with empty keys

### Level 4 (Typed) - 8 tests
- Type inference and validation
- Smart parsing with configurable options
- Language-specific type extraction

### Error Handling - 5 tests
- Malformed input detection
- Error message validation

## Version History

- **v2.0.0** - 4-level architecture with 80 comprehensive test cases
- **v1.0.0** - Legacy monolithic test suite (archived)