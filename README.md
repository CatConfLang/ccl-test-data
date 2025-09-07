# CCL (Categorical Configuration Language)

Central repository for CCL language specification, documentation, and test suite.

## What is CCL?

CCL is a minimal, human-readable configuration language based on simple key-value pairs with elegant composition and nesting capabilities. It's designed around mathematical principles from Category Theory, providing predictable parsing and powerful structural features.

```ccl
/= Application Configuration
app.name = MyApplication
app.debug = true

database =
  host = localhost
  port = 5432
  
servers =
  = web-1.example.com
  = web-2.example.com
```

## Documentation

### Language Guide
- **[Getting Started](docs/getting-started.md)** - Learn CCL syntax and basic concepts
- **[CCL FAQ](docs/ccl_faq.md)** - Common questions and gotchas
- **[Format Comparison](docs/format-comparison.md)** - CCL vs JSON, YAML, TOML, etc.
- **[Glossary](docs/glossary.md)** - Technical terms and definitions

### Implementation Guide  
- **[4-Level Architecture](docs/4-level-architecture.md)** - CCL's systematic implementation approach
- **[Implementing CCL](docs/implementing-ccl.md)** - Guide for language authors
- **[Examples](docs/examples/)** - Practical CCL configuration files

## Language-Agnostic Test Suite

This repository contains comprehensive test cases for CCL parsers organized by architectural level. Use these tests to validate your CCL implementation regardless of programming language.

### 4-Level Architecture

CCL implementations should progress through these levels:

### Level 1: Entry Parsing (Core)
**File**: `tests/level-1-parsing.json` (48 tests)  
**API**: `parse(text) → Result(List(Entry), ParseError)`

Essential parsing functionality that every CCL implementation must support:
- Basic key-value parsing with whitespace handling
- Multiline values and continuation lines  
- Unicode support and line ending normalization
- Empty keys/values, equals-in-values handling
- Edge cases (EOF, mixed whitespace, etc.)

Many tests tagged `"redundant"` are variations of core functionality for comprehensive coverage.

### Level 2: Entry Processing (Extensions)
**File**: `tests/level-2-processing.json` (24 tests + 4 composition tests)  
**API**: `filter_keys()`, entry composition functions

Processing operations on parsed entries:
- Comment filtering (keys starting with `/`)
- Duplicate key handling and composition
- Entry list merging and algebraic properties
- Decorative section parsing (future)

### Level 3: Object Construction (Hierarchical)
**File**: `tests/level-3-objects.json` (8 tests)  
**API**: `make_objects(entries) → CCL`

Converting flat entry lists to nested object structures:
- Recursive parsing of nested values
- Duplicate key merging in object construction
- Empty key handling for list-style data
- Complex nested configuration examples

### Level 4: Typed Parsing (Language-Specific)
**File**: `tests/level-4-typed.json` (17 tests)  
**API**: `get_int()`, `get_bool()`, `get_typed_value()`, etc.

Type-aware extraction with validation:
- Smart type inference (integers, floats, booleans)
- Configurable parsing options
- Language-specific convenience functions
- Type safety and validation

### Error Handling (All Levels)
**File**: `tests/errors.json` (5 tests)

Malformed input detection across all levels.

## Schema

All test files use the unified schema defined in `tests/schema.json`:

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
Focus on `tests/level-1-parsing.json` tests without `"redundant"` tag:
```bash
# Filter out redundant tests for initial implementation
jq '.tests[] | select(.meta.tags | contains(["redundant"]) | not)' tests/level-1-parsing.json
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
    with open('tests/level-1-parsing.json') as f:
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

```
tests/
├── level-1-parsing.json      # Level 1: Core parsing tests (48 tests)
├── level-2-processing.json   # Level 2: Processing and composition tests (28 tests)
├── level-3-objects.json      # Level 3: Object construction tests (8 tests)
├── level-4-typed.json        # Level 4: Typed parsing tests (17 tests)
├── errors.json               # Error handling tests (5 tests)
├── utilities.json            # Additional tools and utilities
└── schema.json               # Test schema definitions
```

This simple structure makes it easy to:
- Quickly locate test files by architecture level
- Run targeted test suites for specific implementation phases
- Navigate and understand the test organization at a glance

## Test Coverage

The new 4-level architecture includes **106 test cases** total:

### Level 1 (Core Parsing) - 48 tests
- 18 essential tests covering core functionality
- 30 additional tests marked as `"redundant"` for comprehensive edge case coverage
- Includes basic parsing, whitespace handling, multiline values, unicode support

### Level 2 (Processing) - 28 tests  
- Comment filtering and composition behavior
- Duplicate key handling and algebraic properties
- Entry list operations and merging

### Level 3 (Objects) - 8 tests
- Nested object construction from flat entries
- Complex hierarchical configurations
- List handling with empty keys

### Level 4 (Typed) - 17 tests
- Type inference and validation
- Smart parsing with configurable options
- Language-specific type extraction

### Error Handling - 5 tests
- Malformed input detection
- Error message validation

## Development

### Prerequisites

This project uses [mise](https://mise.jdx.dev/) for dependency management:

```bash
# Install mise (if not already installed)
curl https://mise.run | sh

# Install project dependencies
mise install
```

Required tools (automatically installed by mise):
- `jq` - JSON processor for extracting test statistics
- `sd` - Fast find-and-replace tool for README automation

### Available Scripts

```bash
npm test              # Validate schemas + check README sync
npm run validate      # Validate all test files against schemas
npm run docs:stats    # Show current test counts
npm run docs:update   # Update README with latest counts
npm run docs:check    # Verify README is up to date
```

## CCL Implementations

### Available Implementations

- **[Gleam](../ccl_gleam/)** - Full 4-level implementation with type-safe parsing
- **[OCaml](https://github.com/chshersh/ccl)** - Reference implementation by CCL creator
- **Go** - Coming soon
- **Rust** - Coming soon

### Implementation Status

Implementations can declare their support level:

- **Level 1**: Basic CCL parser ✅ (required)
- **Level 2**: Comments + composition ✅ 
- **Level 3**: Object construction ✅ (recommended)
- **Level 4**: Typed parsing ✅ (language-specific)

## External Resources

- **[Original Blog Post](https://chshersh.com/blog/2025-01-06-the-most-elegant-configuration-language.html)** - CCL specification and rationale
- **[Reference Implementation](https://github.com/chshersh/ccl)** - OCaml implementation by CCL creator
- **[Community Discussions](https://github.com/chshersh/ccl/discussions)** - Implementation questions and language design

## Contributing

This repository is the central hub for CCL language documentation and testing. Contributions welcome:

1. **Documentation improvements** - Clarify language features and usage patterns
2. **Test case additions** - Expand coverage for edge cases and new features  
3. **Implementation guides** - Help for new language implementations
4. **Example configurations** - Real-world CCL usage patterns

## Version History

- **v2.0.0** - Comprehensive language documentation + 4-level test architecture (106 test cases)
- **v1.0.0** - Legacy monolithic test suite (archived)