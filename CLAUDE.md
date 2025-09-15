# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

Official JSON test suite for CCL (Categorical Configuration Language) implementations with **feature-based tagging** for precise test selection, comprehensive Go-based test runner, and mock implementation.

### Architecture

**Multi-Level CCL Implementation:**
- **Level 1**: Raw parsing (text → flat entries) - `Parse()`
- **Level 2**: Entry processing (indentation, comments) - `Filter()`, `Compose()`, `ExpandDotted()`  
- **Level 3**: Object construction (flat → nested objects) - `MakeObjects()`
- **Level 4**: Typed access (type-safe value extraction) - `GetString()`, `GetInt()`, etc.
- **Level 5**: Validation/formatting - `PrettyPrint()`

**Key Components:**
- `tests/api-*.json` - Feature-specific test suites with structured tagging and counted assertions
- `cmd/ccl-test-runner/` - CLI for test generation and execution with enhanced statistics
- `internal/mock/ccl.go` - Working CCL implementation (should pass most Level 1-4 tests)
- `internal/generator/` - Go test file generation from JSON data
- `internal/stats/enhanced.go` - Feature-based statistics and analysis
- `docs/tag-migration.json` - Tag migration mapping and implementation guidance

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
- Uses structured tags: `function:parse`, `function:make-objects`, `function:get-string` (skips advanced features)
- All enabled tests should pass before commits to maintain clean CI state

## Feature-Based Tagging System

### Structured Tags

All tests now use structured tags for precise test selection:

**Function Tags** (`function:*`) - Required CCL functions:
- `function:parse`, `function:parse-value`, `function:filter`, `function:compose`, `function:expand-dotted`
- `function:make-objects`, `function:get-string`, `function:get-int`, `function:get-bool`, `function:get-float`, `function:get-list`
- `function:pretty-print`

**Feature Tags** (`feature:*`) - Optional language features:
- `feature:comments`, `feature:dotted-keys`, `feature:empty-keys`, `feature:multiline`, `feature:unicode`, `feature:whitespace`

**Behavior Tags** (`behavior:*`) - Implementation choices (mutually exclusive):
- `behavior:crlf-preserve` vs `behavior:crlf-normalize`
- `behavior:tabs-preserve` vs `behavior:tabs-to-spaces`
- `behavior:strict-spacing` vs `behavior:loose-spacing`

**Variant Tags** (`variant:*`) - Specification variants:
- `variant:proposed-behavior` vs `variant:reference-compliant`

### Test Selection Strategy

**For Mock Implementation:**
```bash
# Generate tests for basic functions only
just generate --run-only function:parse,function:make-objects,function:get-string

# Skip advanced features
just generate --skip-tags feature:comments,feature:unicode,variant:proposed-behavior
```

**For Progressive Implementation:**
1. Start with `function:parse` only (Level 1)
2. Add `function:parse-value` for indentation-aware parsing (Level 2)
3. Add `function:make-objects` for object construction (Level 3)
4. Add typed access: `function:get-string`, `function:get-int`, etc. (Level 4)
5. Add processing: `function:filter`, `function:compose`, `function:expand-dotted` (Level 2)
6. Add formatting: `function:pretty-print` (Level 5)

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
    "level": 1,
    "feature": "parsing"
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
4. **Include generated test files in commits** - changes to `internal/generator/` or test JSON files require committing updated `go_tests/` files

### Adding Tests
1. Add to appropriate `tests/api-*.json` file by feature level
2. Include structured tags:
   - Required: `function:*` tags based on validations used
   - Optional: `feature:*` tags for language features required
   - Behavior: `behavior:*` tags for implementation choices (some tests may have multiple behavior tags if they work in all modes)
   - Conflicts: `conflicts` array for mutually exclusive behaviors (only when truly incompatible)
3. Include proper `count` fields matching array lengths or case counts
4. Run `just generate` and `just test-generated` to verify
5. Check `just stats` to see impact on test coverage

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