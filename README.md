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
- **[API Reference](docs/api-reference.md)** - Proposed API structure and patterns
- **[Implementing CCL](docs/implementing-ccl.md)** - Guide for language authors
- **[Examples](docs/examples/)** - Practical CCL configuration files

## Language-Agnostic Test Suite

This repository contains comprehensive test cases for CCL parsers organized by feature and implementation priority. Use these tests to validate your CCL implementation regardless of programming language.

### Feature-Based Organization

CCL tests are organized by functionality rather than rigid levels, making it easier to implement features incrementally:

### Core Functionality
Essential tests that every CCL implementation needs:

**Essential Parsing** - `tests/core/essential-parsing.json` (18 tests)  
Start here for rapid prototyping and basic CCL support:
- Basic key-value parsing with whitespace handling
- Multiline values and continuation lines  
- Unicode support and line ending normalization
- Empty keys/values, equals-in-values handling

**Comprehensive Parsing** - `tests/core/comprehensive-parsing.json` (30 tests)  
Production-ready validation with edge cases:
- Whitespace variations (tabs, spaces, trimming)
- Line ending handling (Unix, Windows, Mac)
- Edge cases (empty keys/values, multiple equals)
- Stress testing with realistic examples

**Object Construction** - `tests/core/object-construction.json` (8 tests)  
**API**: `make_objects(entries) → CCL`

Essential for hierarchical access:
- Recursive parsing of nested values using fixed-point algorithm
- Duplicate key merging in object construction
- Empty key handling for list-style data
- Complex nested configuration support

### Optional Features
Implement these based on your needs:

**Dotted Key Expansion** - `tests/features/dotted-keys.json` (18 tests)  
Enables dual access patterns (`database.host` ↔ hierarchical):
- Basic dotted key expansion to nested structures
- Deep nesting support (3+ levels)
- Mixed dotted and nested syntax
- Conflict resolution and merging

**Comment Filtering** - `tests/features/comments.json` (3 tests)  
**API**: `filter(entries)` 

Remove documentation keys from configuration:
- Comment syntax (keys starting with `/`)
- Filtering and processing behavior

**Entry Processing** - `tests/features/processing.json` (21 tests)  
**API**: `compose_entries()`, advanced processing

Advanced entry composition and merging:
- Duplicate key handling and composition
- Entry list merging with algebraic properties
- Complex composition scenarios

**Typed Access** - `tests/features/typed-access.json` (17 tests)  
**API**: `get_string()`, `get_int()`, `get_bool()`, etc.

Type-aware extraction with validation:
- Smart type inference (integers, floats, booleans) 
- Configurable parsing options and validation
- Language-specific convenience functions
- Dual access pattern support (dotted + hierarchical)

### Integration & Validation

**Error Handling** - `tests/integration/errors.json` (5 tests)  
Malformed input detection and error reporting

**Pretty Printing** - `tests/pretty-print.json` (15 tests)  
Round-trip testing and canonical formatting

## Implementation Path

**Recommended progression for new implementations:**

1. **Start Simple**: `core/essential-parsing.json` (18 tests)
   - Gets you a working CCL parser quickly
   - Handles 80% of real-world CCL files

2. **Add Hierarchy**: `core/object-construction.json` (8 tests)  
   - Enables nested configuration access
   - Required for practical use

3. **Production Ready**: `core/comprehensive-parsing.json` (30 tests)
   - Handles edge cases and whitespace variations
   - Required for robust production systems

4. **Choose Features**: Select from `features/` based on your needs
   - `dotted-keys.json` - For convenient `database.host` style access
   - `typed-access.json` - For `get_int()`, `get_bool()` convenience
   - `comments.json` - For documentation in config files
   - `processing.json` - For advanced composition features

5. **Validate**: `integration/errors.json` for error handling

## Test Structure

```
tests/
├── core/                     # Essential functionality
│   ├── essential-parsing.json      # 18 tests - start here
│   ├── comprehensive-parsing.json  # 30 tests - production ready  
│   └── object-construction.json    # 8 tests - hierarchy support
├── features/                 # Optional enhancements
│   ├── dotted-keys.json           # 18 tests - dual access patterns
│   ├── comments.json              # 3 tests - comment filtering
│   ├── processing.json            # 21 tests - advanced composition
│   └── typed-access.json          # 17 tests - type-safe getters
├── integration/              # Validation & edge cases
│   └── errors.json               # 5 tests - error handling
└── pretty-print.json        # 15 tests - round-trip formatting
```

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
    "tags": ["basic", "whitespace"],
    "level": 1
  }
}
```

## Test Coverage

The feature-based architecture includes **135 test cases** total:

### Core (56 tests)
- **Essential**: 18 tests for rapid implementation  
- **Comprehensive**: 30 tests for production validation
- **Object Construction**: 8 tests for hierarchy support

### Features (59 tests)  
- **Dotted Keys**: 18 tests for dual access patterns
- **Comments**: 3 tests for filtering functionality
- **Processing**: 21 tests for advanced composition  
- **Typed Access**: 17 tests for type-safe APIs

### Integration (5 tests)
- **Error Handling**: Malformed input detection and reporting

### Utilities (15 tests)
- **Pretty Print**: Round-trip testing and canonical formatting

## Usage

### Running Tests

**Validate test files:**
```bash
npm test                    # Validate all test files
npm run validate:core       # Core validation only
```

**Test Statistics:**
```bash
./scripts/collect-stats.sh  # JSON output  
./scripts/collect-stats.sh --interactive  # Pretty display
```

### Language Integration

Each test provides multiple expected outputs for different implementation approaches:

```json
{
  "input": "key = value",
  "expected": [{"key": "key", "value": "value"}],           // Raw parsing
  "expected_nested": {"key": "value"},                      // Object form
  "expected_typed": {"key": {"type": "string", "value": "value"}}  // Typed
}
```

Choose the output format that matches your implementation level.

## Contributing

This test suite is designed to be comprehensive and implementation-agnostic. When adding tests:

1. **Choose the right category**: Core vs Features vs Integration
2. **Follow naming conventions**: Descriptive test names with consistent tagging
3. **Validate against schema**: All tests must pass `npm run validate`
4. **Update documentation**: Keep this README and docs/ current

## License

This test suite and documentation is provided as a reference for CCL implementations across all programming languages.