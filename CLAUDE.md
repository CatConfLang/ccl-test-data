# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Start

This repository contains a comprehehensive JSON test suite for CCL (Categorical Configuration Language) implementations. It provides comprehensive language-agnostic tests with feature-based tagging for precise test selection.

**Essential first commands:**
```bash
just deps                   # Install dependencies
just reset                  # Generate and run basic tests (ensures clean state)
just stats                  # View test coverage and statistics
```

**What this gives you:**
- Verified working repository state
- Understanding of test scope (180 tests, 384 assertions across 12 files)
- Ready environment for development

## Common Development Tasks

### Standard Development Workflow
```bash
# Pre-commit workflow (REQUIRED before commits)
just lint                   # Format and lint Go code
just reset                  # Generate basic tests, verify they pass
just validate               # Validate JSON test files

# Full development cycle
just dev                    # Generate all tests and run them
just test                   # Run tests with optional filtering
```

### Testing Specific Features
```bash
# Test by function group
just test --functions core      # Core parsing functions
just test --functions typed     # Typed access functions (GetString, GetInt, etc.)
just test --functions processing # Entry processing (Filter, Combine, etc.)

# Test by feature category
just test-parsing              # All parsing-related functionality
just test-comments             # Comment handling tests
just test-objects              # Object construction tests
```

### Adding New Tests
1. Add to appropriate `source_tests/api_*.json` file by feature category
2. Include structured tags:
   - **Function tags**: `function:parse`, `function:build_hierarchy`, `function:get_string`
   - **Feature tags**: `feature:comments`, `feature:experimental_dotted_keys`
   - **Behavior tags**: `behavior:boolean_strict`, `behavior:crlf_preserve_literal`
3. Include proper `count` fields matching array lengths
4. Run `just generate && just test` to verify
5. Check `just stats` to see impact on test coverage

### Debug and Analysis
```bash
just stats                     # Detailed statistics and coverage
just list                      # Show all available test packages
just validate                  # JSON schema validation
just clean                     # Remove generated files
just generate-flat             # Generate flat format from source tests
```

## Command Reference

### Essential Commands
| Command | Purpose | When to Use |
|---------|---------|-------------|
| `just reset` | Generate core parsing tests only | Before commits, quick verification |
| `just dev` | Generate all tests and run them | Full development cycle |
| `just lint` | Format and lint Go code | Before every commit (required) |
| `just validate` | Validate JSON schema compliance | After modifying test files |

### Test Generation
| Command | Purpose | Details |
|---------|---------|---------|
| `just generate` | Generate Go test files from flat JSON | Main test generation |
| `just generate-flat` | Transform source tests to flat format | Uses ccl-test-lib |
| `just clean` | Remove generated files | Clean slate |

### Test Execution
| Command | Purpose | Filtering Options |
|---------|---------|------------------|
| `just test` | Run tests with optional filtering | `--functions`, `--features`, `--behaviors` |
| `just test --functions core` | Core functions only | `parse`, `build_hierarchy` |
| `just test --functions typed` | Typed access functions | `get_string`, `get_int`, etc. |

### Analysis and Validation
| Command | Purpose | Output |
|---------|---------|--------|
| `just stats` | Show detailed test statistics | Function coverage, feature distribution |
| `just list` | Show available test packages | All generated test categories |
| `just docs-check` | Verify documentation is current | README statistics validation |

## Architecture and Test System

### Dual-Format Test Architecture
The repository uses a dual-format system optimized for both maintainability and implementation:

- **Source Format** (`source_tests/api_*.json`): Human-maintainable test definitions with multiple validations per test
- **Generated Format** (`generated_tests/api_*.json`): Machine-friendly flat format (one test per validation)
- **Go Tests** (`go_tests/`): Generated Go test files for execution

**Why two formats?**
- Source format: Easy to write and maintain, multiple validations per test
- Generated format: Type-safe implementation with better API ergonomics
- Transformation happens via `just generate-flat` using the ccl-test-lib

### CCL Function Groups
Tests are organized around progressive CCL implementation:

- **Core Parsing**: `Parse()` - Convert text to flat key-value entries
- **Entry Processing**: `Filter()`, `Combine()`, `ExpandDotted()` - Transform and combine entries
- **Object Construction**: `BuildHierarchy()` - Build nested object hierarchies from flat entries
- **Typed Access**: `GetString()`, `GetInt()`, `GetBool()`, `GetFloat()`, `GetList()` - Type-safe value extraction
- **Formatting**: `CanonicalFormat()` - Generate standardized formatted output

### Feature-Based Tagging System

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

## Implementation Details

### Mock Implementation Strategy
The `internal/mock/ccl.go` contains a working CCL implementation for testing and development:

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

### Progressive Implementation Guide
For building your own CCL implementation:

1. Start with `function:parse` only (core parsing)
2. Add `function:parse_value` for indentation-aware parsing
3. Add `function:build_hierarchy` for object construction
4. Add typed access: `function:get_string`, `function:get_int`, etc.
5. Add processing: `function:filter`, `function:expand_dotted`
6. Add formatting: `function:canonical_format`

**Test Selection Examples:**
```bash
# Generate tests for basic functions only
just generate --run-only function:parse,function:build_hierarchy,function:get_string

# Skip advanced features
just generate --skip-tags feature:comments,feature:unicode,behavior:crlf_preserve_literal
```

### Current Test Statistics
- **180 tests** across **12 source files** with **384 assertions**
- **Top functions**: `parse` (163 tests), `build_hierarchy` (77 tests), `get_list` (48 tests)
- **Top features**: `empty_keys` (43 tests), `whitespace` (24 tests), `multiline` (10 tests)
- **Behavioral choices**: boolean parsing, CRLF handling, list coercion, spacing rules

### Test File Organization
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

## Build System and Dependencies

- **Build tool**: `just` (justfile) for cross-platform automation
- **Go version**: 1.25.1
- **Module**: `github.com/ccl-test-data/test-runner`
- **Key deps**: CLI framework (urfave/cli), JSON schema validation (jsonschema), styling libraries (charmbracelet), ccl-test-lib
- **External dependency**: `../ccl-test-lib/` - shared library for flat test generation and types

## Before Committing

1. **Always run `just lint`** - formats and checks Go code
2. **Ensure `just reset` passes** - verifies repository is in clean state
3. **Validate changes with `just validate`** - checks JSON schema compliance
4. **Include generated test files in commits** - changes to `internal/generator/` or test JSON files require committing updated `go_tests/` files

## Test Data Format Example

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

## Troubleshooting

### Tests Failing After Changes
```bash
just clean && just reset    # Reset to known good state
just validate              # Check JSON schema compliance
just lint                  # Fix formatting issues
```

### Understanding Test Coverage
```bash
just stats                 # See function and feature coverage
just list                  # See all test packages
just test --functions core # Test only core functions
```

### Repository Structure Issues
```bash
just deps                  # Reinstall dependencies
go mod tidy               # Clean Go module dependencies
just clean && just generate # Regenerate all test files
```
