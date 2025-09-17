# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

Official JSON test suite for CCL (Categorical Configuration Language) implementations with **feature-based tagging** for precise test selection, comprehensive Go-based test runner, and mock implementation.

### Architecture

**CCL Implementation Functions:**
- **Core Parsing**: Raw parsing (text → flat entries) - `Parse()`
- **Entry Processing**: Indentation, comments, filtering - `Filter()`, `Combine()`, `ExpandDotted()`
- **Object Construction**: Flat → nested objects - `BuildHierarchy()`
- **Typed Access**: Type-safe value extraction - `GetString()`, `GetInt()`, etc.
- **Formatting**: Validation/formatting - `CanonicalFormat()`

**Key Components:**
- `source_tests/api_*.json` - Feature-specific test suites with structured tagging (12 files, 180 tests, 384 assertions)
- `cmd/ccl-test-runner/` - CLI for test generation and execution with enhanced statistics
- `internal/mock/ccl.go` - Working CCL implementation (should pass most basic to advanced tests)
- `internal/generator/` - Go test file generation from JSON data with template-based output
- `internal/stats/enhanced.go` - Feature-based statistics and analysis
- `../ccl-test-lib/` - Shared library for flat test generation and common functionality

## Essential Commands

### Pre-Commit Workflow
```bash
just lint                   # Format and lint Go code (REQUIRED before commits)
just reset                  # Generate basic tests, run them (ensures clean state)
```

### Development Commands
```bash
# Basic development (reset is alias for dev-basic)
just reset                  # Generate core parsing tests only, ensuring all pass
just dev-basic             # Same as reset - generates minimal passing test set

# Full development cycle  
just dev                   # Generate all tests and run them
just test                  # Comprehensive suite: validate + docs + generate + test

# Test generation and execution
just generate              # Generate all Go test files from JSON data
just generate-flat         # Generate flat JSON files from source format
just test                  # Run all tests with optional filtering
```

### Function Group Testing
```bash
# Test by function group
just test --functions parsing    # Basic parsing (Parse function)
just test --functions processing # Entry processing (Filter, Combine, etc.)
just test --functions objects    # Object construction (BuildHierarchy)
just test --functions typed      # Typed access (GetString, GetInt, etc.)

# Test by feature category
just test-parsing          # All parsing-related functionality
just test-objects          # Object construction tests
just test-comments         # Comment handling tests
```

### Quality and Validation
```bash
just validate              # Validate JSON test files against schema
just stats                 # Show detailed test statistics
just docs-check            # Verify documentation is current
```

## Mock Implementation Strategy

The `internal/mock/ccl.go` contains a working CCL implementation that should pass most tests using only basic functions:

**Core Functions Implemented:**
- `Parse()` - Core parsing with comment support (`/=` syntax)
- `BuildHierarchy()` - Object construction with dotted key support
- `GetString()`, `GetInt()`, `GetBool()`, `GetFloat()`, `GetList()` - Typed access functions
- `Filter()`, `Combine()`, `ExpandDotted()` - Entry processing (basic implementations)
- `CanonicalFormat()` - Formatting (basic implementation)

**Repository State Management:**
- `just reset` generates only tests that the mock implementation can pass
- Uses structured tags: `function:parse`, `function:build_hierarchy`, `function:get_string` (skips advanced features)
- All enabled tests should pass before commits to maintain clean CI state

## Feature-Based Tagging System

### Structured Tags

All tests now use structured tags for precise test selection:

**Function Tags** (`function:*`) - Required CCL functions:
- `function:parse`, `function:parse_value`, `function:filter`, `function:expand_dotted`
- `function:build_hierarchy`, `function:get_string`, `function:get_int`, `function:get_bool`, `function:get_float`, `function:get_list`
- `function:canonical_format`, `function:round_trip`, `function:load`, `function:associativity`

**Feature Tags** (`feature:*`) - Optional language features:
- `feature:comments`, `feature:experimental_dotted_keys`, `feature:empty_keys`, `feature:multiline`, `feature:unicode`, `feature:whitespace`

**Behavior Tags** (`behavior:*`) - Implementation choices (mutually exclusive):
- `behavior:crlf_preserve_literal` vs `behavior:crlf_normalize_to_lf`
- `behavior:tabs_preserve` vs `behavior:tabs_to_spaces`
- `behavior:strict_spacing` vs `behavior:loose_spacing`
- `behavior:boolean_strict` vs `behavior:boolean_lenient`
- `behavior:list_coercion_enabled` vs `behavior:list_coercion_disabled`

**Variant Tags** (`variant:*`) - Specification variants:
- `variant:proposed-behavior` vs `variant:reference-compliant`

### Test Selection Strategy

**For Mock Implementation:**
```bash
# Generate tests for basic functions only
just generate --run-only function:parse,function:build_hierarchy,function:get_string

# Skip advanced features
just generate --skip-tags feature:comments,feature:unicode,behavior:crlf_preserve_literal
```

**For Progressive Implementation:**
1. Start with `function:parse` only (core parsing)
2. Add `function:parse_value` for indentation-aware parsing
3. Add `function:build_hierarchy` for object construction
4. Add typed access: `function:get_string`, `function:get_int`, etc.
5. Add processing: `function:filter`, `function:expand_dotted`
6. Add formatting: `function:canonical_format`

## Test Data Format

### Counted Assertions
All tests use required `count` fields for precise validation:

```json
{
  "name": "basic_parsing",
  "input": "key = value",
  "validations": {
    "parse": {
      "count": 1,
      "expected": [{"key": "key", "value": "value"}]
    },
    "get_string": {
      "count": 1, 
      "cases": [{"args": ["key"], "expected": "value"}]
    }
  },
  "meta": {
    "tags": ["function:parse", "function:get-string"],
    "group": "parsing",
    "feature": "parsing"
  }
}
```

### Test Organization
- **Core Parsing**: `api_core_ccl_parsing.json` (basic parsing functionality)
- **Advanced Processing**: `api_advanced_processing.json` (composition and filtering)
- **Hierarchy**: `api_core_ccl_hierarchy.json` (nested object creation)
- **Integration**: `api_core_ccl_integration.json` (cross-function functionality)
- **Typed Access**: `api_typed_access.json` (type-safe access)
- **List Access**: `api_list_access.json` (list-specific operations)
- **Comments**: `api_comments.json` (comment syntax and filtering)
- **Edge Cases**: `api_edge_cases.json` (boundary conditions)
- **Errors**: `api_errors.json` (error handling validation)
- **Experimental**: `api_experimental.json` (experimental features)

## Implementation Guidelines

### Before Committing
1. **Always run `just lint`** - formats and checks Go code
2. **Ensure `just reset` passes** - verifies repository is in clean state
3. **Validate changes with `just validate`** - checks JSON schema compliance
4. **Include generated test files in commits** - changes to `internal/generator/` or test JSON files require committing updated `go_tests/` files

### Adding Tests
1. Add to appropriate `source_tests/api_*.json` file by feature category
2. Include structured tags:
   - Required: `function:*` tags based on validations used
   - Optional: `feature:*` tags for language features required
   - Behavior: `behavior:*` tags for implementation choices (some tests may have multiple behavior tags if they work in all modes)
   - Conflicts: `conflicts` array for mutually exclusive behaviors (only when truly incompatible)
3. Include proper `count` fields matching array lengths or case counts
4. Run `just generate` and `just test` to verify
5. Check `just stats` to see impact on test coverage

### Mock Implementation Development
The mock implementation should handle progressively more CCL features:
- Currently passes basic parsing, object construction, and typed access
- `Parse()` function handles key-value pairs and `/=` comments
- `BuildHierarchy()` supports dotted keys and duplicate key lists
- Typed getters convert string values to appropriate types
- `CanonicalFormat()` provides basic formatted output

### Build System
- **Build tool**: `just` (justfile) for cross-platform automation
- **Go version**: 1.25.1
- **Module**: `github.com/ccl-test-data/test-runner`
- **Key deps**: CLI framework (urfave/cli), JSON schema validation (jsonschema), styling libraries (charmbracelet), ccl-test-lib
- **External dependency**: `../ccl-test-lib/` - shared library for flat test generation and types

## Architecture Overview

### Dual-Format Test System
- **Source Format**: `source_tests/api_*.json` - Human-maintainable test definitions with multiple validations per test
- **Generated Format**: `generated_tests/api_*.json` - Machine-friendly flat format (one test per validation)
- **Go Tests**: `go_tests/` - Generated Go test files for execution

### CLI Commands and Workflow
- `just generate-flat` - Transform source tests to flat format (uses ccl-test-lib)
- `just generate` - Create Go test files from flat format
- `just test` - Run generated Go tests against mock implementation
- `just dev-basic` / `just reset` - Quick development cycle (core parsing tests only)

### Test Statistics (Current)
- **180 tests** across **12 source files** with **384 assertions**
- **Function coverage**: `parse` (163 tests), `build_hierarchy` (77 tests), `get_list` (48 tests)
- **Feature distribution**: `empty_keys` (43 tests), `whitespace` (24 tests), `multiline` (10 tests)
- **Behavioral choices**: boolean parsing, CRLF handling, list coercion, spacing rules

## Common Development Tasks

### Running Specific Test Subsets
```bash
# Test by function group
just test --functions core      # Core functions only
just test --functions typed     # Typed access functions

# Feature-specific testing
just test-comments              # Comment handling tests
just test-parsing              # All parsing functionality

# Generate for specific capabilities
just generate --run-only function:parse,function:build_hierarchy
just generate --skip-tags feature:unicode,behavior:strict_spacing
```

### Debug and Development
```bash
just list                      # Show all available test packages
just stats                     # Detailed statistics and coverage
just validate                  # JSON schema validation
just clean                     # Remove generated files
```