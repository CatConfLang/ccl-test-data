# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

Official JSON test suite for CCL (Categorical Configuration Language) implementations with a comprehensive Go-based test runner and mock implementation.

### Architecture

**Multi-Level CCL Implementation:**
- **Level 1**: Raw parsing (text → flat entries) - `Parse()`
- **Level 2**: Entry processing (indentation, comments) - `Filter()`, `Compose()`, `ExpandDotted()`  
- **Level 3**: Object construction (flat → nested objects) - `MakeObjects()`
- **Level 4**: Typed access (type-safe value extraction) - `GetString()`, `GetInt()`, etc.
- **Level 5**: Validation/formatting - `PrettyPrint()`

**Key Components:**
- `tests/api-*.json` - Feature-specific test suites with counted assertions
- `cmd/ccl-test-runner/` - CLI for test generation and execution
- `internal/mock/ccl.go` - Working CCL implementation (should pass most Level 1-4 tests)
- `internal/generator/` - Go test file generation from JSON data

## Essential Commands

### Pre-Commit Workflow
```bash
just lint                   # Format and lint Go code (REQUIRED before commits)
just reset                  # Generate basic tests, run them (ensures clean state)
```

### Development Commands
```bash
# Basic development (reset is alias for dev-basic)
just reset                  # Generate Level 1 tests only, ensuring all pass
just dev-basic             # Same as reset - generates minimal passing test set

# Full development cycle  
just dev                   # Generate all tests and run them
just test                  # Comprehensive suite: validate + docs + generate + test

# Test generation and execution
just generate              # Generate all Go test files from JSON data
just test-generated        # Run generated Go tests only
```

### Level and Feature Testing
```bash
# Test by CCL implementation level
just test-level1           # Basic parsing (Parse function)
just test-level2           # Entry processing (Filter, Compose, etc.)
just test-level3           # Object construction (MakeObjects)
just test-level4           # Typed access (GetString, GetInt, etc.)

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
- `Parse()` - Level 1 parsing with comment support (`/=` syntax)
- `MakeObjects()` - Level 3 object construction with dotted key support
- `GetString()`, `GetInt()`, `GetBool()`, `GetFloat()`, `GetList()` - Level 4 typed access
- `Filter()`, `Compose()`, `ExpandDotted()` - Level 2 processing (basic implementations)

**Repository State Management:**
- `just reset` generates only tests that the mock implementation can pass
- Uses tags: `basic`, `essential-parsing`, `empty` (skips advanced features)
- All enabled tests should pass before commits to maintain clean CI state

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
  }
}
```

### Test Organization
- **Essential**: `api-essential-parsing.json` (basic Level 1 functionality)
- **Comprehensive**: `api-comprehensive-parsing.json` (edge cases, whitespace)
- **Processing**: `api-processing.json` (Level 2 composition and filtering)
- **Objects**: `api-object-construction.json` (Level 3 nested object creation)
- **Typed**: `api-typed-access.json` (Level 4 type-safe access)
- **Comments**: `api-comments.json` (comment syntax and filtering)

## Implementation Guidelines

### Before Committing
1. **Always run `just lint`** - formats and checks Go code
2. **Ensure `just reset` passes** - verifies repository is in clean state
3. **Validate changes with `just validate`** - checks JSON schema compliance
4. **Include generated test files in commits** - changes to `internal/generator/` or test JSON files require committing updated `generated_tests/` files

### Adding Tests
1. Add to appropriate `tests/api-*.json` file by feature level
2. Include proper `count` fields matching array lengths or case counts
3. Run `just generate` and `just test-generated` to verify

### Mock Implementation Development
The mock implementation should handle progressively more CCL features:
- Currently passes basic parsing, object construction, and typed access
- Level 1 `Parse()` function handles key-value pairs and `/=` comments
- Level 3 `MakeObjects()` supports dotted keys and duplicate key lists
- Level 4 typed getters convert string values to appropriate types

### Build System
- **Build tool**: `just` (justfile) for cross-platform automation
- **Go version**: 1.23.0 
- **Module**: `github.com/ccl-test-data/test-runner`
- **Key deps**: CLI framework, JSON schema validation, styling libraries