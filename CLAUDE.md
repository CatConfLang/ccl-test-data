# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Start

This repository contains a comprehensive JSON test suite for CCL (Categorical Configuration Language) implementations with feature-based tagging for precise test selection.

**Essential first commands:**
```bash
just deps                   # Install dependencies
just reset                  # Generate and run basic tests (ensures clean state)
just stats                  # View test coverage and statistics
```

## Development Workflow

### Standard Commands
```bash
# Pre-commit workflow (REQUIRED)
just lint                   # Format and lint Go code
just reset                  # Generate basic tests, verify they pass
just validate               # Validate JSON test files

# Development cycle
just dev                    # Generate all tests and run them
just test                   # Run tests with optional filtering
```

### Testing by Category
```bash
just test --functions core      # Core parsing functions
just test --functions typed     # Typed access functions
just test-parsing              # All parsing-related functionality
just test-comments             # Comment handling tests
```

### Adding New Tests
1. Add to appropriate `source_tests/api_*.json` file
2. Include structured tags: `function:*`, `feature:*`, `behavior:*`
3. Include proper `count` fields matching array lengths
4. Run `just generate && just test` to verify

## Command Reference

| Command | Purpose | When to Use |
|---------|---------|-------------|
| `just reset` | Generate core parsing tests only | Before commits, quick verification |
| `just dev` | Generate all tests and run them | Full development cycle |
| `just lint` | Format and lint Go code | Before every commit (required) |
| `just validate` | Validate JSON schema compliance | After modifying test files |
| `just stats` | Show detailed test statistics | Function coverage, feature distribution |
| `just generate` | Generate Go test files from flat JSON | Main test generation |
| `just test` | Run tests with optional filtering | `--functions`, `--features`, `--behaviors` |

## Test Architecture

### Dual-Format System
- **Source Format** (`source_tests/`): Human-maintainable with multiple validations per test
- **Generated Format** (`generated_tests/`): Machine-friendly flat format (one test per validation)
- **Go Tests** (`go_tests/`): Generated Go test files for execution

### CCL Function Groups
- **Core**: `Parse()`, `BuildHierarchy()` - Convert text to hierarchical objects
- **Typed Access**: `GetString()`, `GetInt()`, `GetBool()`, `GetFloat()`, `GetList()`
- **Processing**: `Filter()`, `Combine()`
- **Formatting**: `CanonicalFormat()`

### Tagging System
- **`function:*`** - Required CCL functions (`parse`, `build_hierarchy`, `get_string`)
- **`feature:*`** - Optional features (`comments`, `unicode`)
- **`behavior:*`** - Implementation choices (`crlf_preserve`, `boolean_strict`)

## Progressive Implementation

### Mock Implementation
The `internal/mock/ccl.go` provides a working CCL implementation with core functions.

### Implementation Steps
1. Start with `function:parse` (core parsing)
2. Add `function:build_hierarchy` (object construction)
3. Add typed access: `function:get_string`, `function:get_int`, etc.
4. Add processing: `function:filter`

### Test Selection
```bash
# Basic functions only
just generate --run-only function:parse,function:build_hierarchy

# Skip advanced features
just generate --skip-tags feature:comments,feature:unicode
```

## Build System

- **Build tool**: `just` (justfile)
- **Go version**: 1.25.1
- **Key deps**: CLI framework, JSON schema validation, ccl-test-lib

## Before Committing

1. **`just lint`** - Format and check Go code
2. **`just reset`** - Verify repository in clean state
3. **`just validate`** - Check JSON schema compliance
4. **Include generated files** - Commit updated `go_tests/` files

## Test Data Format

```json
{
  "name": "basic_parsing",
  "input": "key = value",
  "validations": {
    "parse": {
      "count": 1,
      "expected": [{"key": "key", "value": "value"}]
    }
  },
  "meta": {
    "tags": ["function:parse"],
    "feature": "parsing"
  }
}
```
