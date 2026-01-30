# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Start

This repository contains a comprehensive JSON test suite for CCL (Categorical Configuration Language) implementations with feature-based tagging for precise test selection.

**Test Suite Stats:** Test files organized in `source_tests/core/` and `source_tests/experimental/` (run `just stats` for current counts)

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

### Testing Options
```bash
just test                       # Run tests (basic tests by default)
just test-all                   # Run all tests including failing ones
just test-verbose              # Run tests with verbose output
just test -- <go-test-args>    # Pass custom arguments to go test
```

### Adding New Tests
1. Add to appropriate `source_tests/core/api_*.json` or `source_tests/experimental/` file
2. Each test includes:
   - `name`: Unique test identifier
   - `input`: CCL text to test
   - `tests`: Array of function validations with `function` and `expect` fields
   - `features`: Optional array (`comments`, `empty_keys`, `multiline`, `unicode`, `whitespace`)
   - `behaviors`: Optional array (`boolean_strict`, `crlf_normalize_to_lf`, etc.)
   - `variants`: Optional array (`proposed_behavior`, `reference_compliant`)
3. Run `just generate && just test` to verify

### Test Files Structure
```
source_tests/
├── core/
│   ├── api_comments.json
│   ├── api_core_ccl_hierarchy.json
│   ├── api_core_ccl_integration.json
│   ├── api_core_ccl_parsing.json
│   ├── api_edge_cases.json
│   ├── api_errors.json
│   ├── api_advanced_processing.json
│   ├── api_list_access.json
│   ├── api_proposed_behavior.json
│   ├── api_reference_compliant.json
│   ├── api_typed_access.json
│   ├── api_whitespace_behaviors.json
│   ├── property_round_trip.json
│   └── property_algebraic.json
└── experimental/
    └── api_experimental.json
```

## Command Reference

| Command | Purpose | When to Use |
|---------|---------|-------------|
| `just reset` (alias: `dev-basic`) | Generate core tests and verify they pass | Before commits, quick verification |
| `just dev` | Full development cycle: generate all tests and run | Complete test suite development |
| `just lint` | Format and lint Go code | Before every commit (required) |
| `just validate` | Validate JSON schema compliance | After modifying test files |
| `just stats` | Show detailed test statistics | Review function coverage, feature distribution |
| `just generate` | Generate flat JSON then Go test files | Main test generation (combines generate-flat + generate-go) |
| `just test` | Run tests (basic by default) | Execute test suite |
| `just test-all` | Run all tests including failing ones | Full test validation |

## Test Architecture

### Dual-Format System
- **Source Format** (`source_tests/`): Human-maintainable with multiple validations per test
- **Generated Format** (`generated_tests/`): Machine-friendly flat format (one test per validation)
- **Go Tests** (`go_tests/`): Generated Go test files for execution

### CCL Function Groups (per schema)
- **Core Parsing**: `parse`, `parse_indented`, `build_hierarchy`
- **Typed Access**: `get_string`, `get_int`, `get_bool`, `get_float`, `get_list`
- **Processing**: `filter`, `compose`, `merge`
- **Formatting/IO**: `canonical_format`, `load`, `round_trip`

**Note:** Mock implementation (`internal/mock/ccl.go`) provides: Parse, ParseIndented, Filter, BuildHierarchy, GetString, GetInt, GetBool, GetFloat, GetList, PrettyPrint, ExpandDotted. (Note: mock uses `Combine` method name but schema specifies `compose` function)

**Function Details:**
- **`parse`**: Basic lexical parsing - returns flat entries where values are raw strings
- **`parse_indented`**: Indentation-normalized parsing - calculates common leading whitespace and strips it from all lines (like Python's `textwrap.dedent`)
- **`build_hierarchy`**: Recursively parses entry values to create nested object structure

### Test Metadata
- **`functions`** - Required CCL functions (filter: skip if unsupported)
- **`features`** - Language features exercised (informational only, for reporting gaps)
- **`behaviors`** - Implementation choices (filter via `conflicts` field)
- **`variants`** - Spec interpretation (filter via `conflicts` field)
- **`conflicts`** - Mutually exclusive options (filter: skip if your choice is listed)

## Function-Based Implementation

### Mock Implementation
The `internal/mock/ccl.go` provides a working CCL implementation with core functions.

### Available Functions
Tests are organized by the CCL functions they validate. Implement the functions your library needs:

**Core Parsing:**
- `parse` - Basic lexical parsing to flat entries
- `parse_indented` - Indentation normalization (used by build_hierarchy)
- `build_hierarchy` - Object construction from flat entries

**Typed Access:**
- `get_string`, `get_int`, `get_bool`, `get_float`, `get_list` - Type-safe value extraction

**Processing:**
- `filter`, `compose`, `merge` - Entry manipulation

**Formatting/IO:**
- `canonical_format`, `load`, `round_trip` - Output and validation

## Build System

- **Build tool**: `just` (justfile)
- **Go version**: See go.mod for current version
- **Module**: `github.com/catconflang/ccl-test-data`

## Changelog Generation

This project uses [Python Semantic Release](https://python-semantic-release.readthedocs.io/) for changelog generation.

### Squash Merge Commit Format

When squash merging PRs with multiple logical changes, format the commit body with `*` prefixes for each sub-commit:

```
feat(scope): main PR title (#123)

Brief summary of the changes.

* fix(tests): first logical change

Description of the first change.

* fix(schema): second logical change

Description of the second change.

* feat(cli): third logical change

Description of the third change.
```

**Key points:**
- Main commit title follows conventional commits: `type(scope): description`
- Each sub-entry starts with `* type(scope): description` on its own line
- Include a blank line and description after each sub-entry
- Each sub-entry becomes a separate changelog entry with its body text preserved
- `chore` type entries are excluded from the changelog

### Changelog Commands
```bash
semantic-release changelog          # Generate/update CHANGELOG.md
semantic-release --noop version     # Preview what would happen
semantic-release version --no-push  # Version without pushing
```

## Shared Test Infrastructure

This repository includes reusable Go packages for CCL implementations:

```
config/     # Type-safe capability constants (CCLFunction, CCLFeature, CCLBehavior)
loader/     # Test loading with filtering by capabilities
generator/  # Source-to-flat format transformation
types/      # Unified data structures for test suites
```

### Usage in CCL Implementations

```go
import (
    "github.com/catconflang/ccl-test-data/config"
    "github.com/catconflang/ccl-test-data/loader"
    "github.com/catconflang/ccl-test-data/types"
)

// Declare capabilities
cfg := config.ImplementationConfig{
    SupportedFunctions: []config.CCLFunction{
        config.FunctionParse,
        config.FunctionBuildHierarchy,
    },
    SupportedFeatures: []config.CCLFeature{
        config.FeatureComments,
    },
}

// Load compatible tests
testLoader := loader.NewTestLoader("path/to/ccl-test-data", cfg)
tests, _ := testLoader.LoadAllTests(loader.LoadOptions{
    Format:     loader.FormatFlat,
    FilterMode: loader.FilterCompatible,
})
```

**Note:** These packages were consolidated from the deprecated `ccl-test-lib` repository. For local development, use a `replace` directive in your `go.mod`:
```
replace github.com/catconflang/ccl-test-data => ../ccl-test-data
```

## Before Committing

1. **`just lint`** - Format and check Go code
2. **`just reset`** - Verify repository in clean state
3. **`just validate`** - Check JSON schema compliance
4. **Include generated files** - Commit updated `go_tests/` files

## Test Data Format

```json
{
  "tests": [
    {
      "name": "basic_parsing",
      "input": "key = value",
      "tests": [
        {
          "function": "parse",
          "expect": [
            {"key": "key", "value": "value"}
          ]
        }
      ],
      "features": ["whitespace"],
      "behaviors": ["tabs_as_whitespace"]
    }
  ]
}
```

See `schemas/source-format.json` for complete schema definition.
