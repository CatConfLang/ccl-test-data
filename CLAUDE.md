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
just test                   # Run tests
just test-all               # Run all tests including failing ones
```

> [!IMPORTANT]
> **`just reset` vs `just generate`:** `reset` (and `build`) use filtered generation (`--run-only function:parse --skip-tags ...`). `just generate` is unfiltered and produces a much larger go_tests set. Always use `reset`/`build` unless you specifically need the full set — running `generate` after a small edit produces thousands of lines of unrelated drift in `go_tests/`.

> [!NOTE]
> **OCaml is canonical.** The test suite tracks the OCaml reference-implementation semantics. See <https://catconflang.com/reference/canonical-semantics/> for the normative rule and implications; this test-data repo's fixtures follow it.

### Testing Options
```bash
just test                       # Run tests (basic tests by default)
just test-all                   # Run all tests including failing ones
```

### Adding New Tests
1. Add to appropriate `source_tests/core/api_*.json` or `source_tests/experimental/` file
2. Each test includes:
   - `name`: Unique test identifier
   - `input`: CCL text to test
   - `tests`: Array of function validations with `function` and `expect` fields
   - `features`: Optional array (`comments`, `empty_keys`, `multiline_continuation`, `multiline_keys`, `unicode`, `whitespace`, `tab_in_value_preserved`, `toplevel_indent_strip`)
   - `behaviors`: Optional array (`boolean_strict`, `crlf_normalize_to_lf`, etc.)
   - `variants`: Optional array (`proposed_behavior`, `reference_compliant`)
3. Run `just generate && just test` to verify

### Test Files Structure
```
source_tests/
├── core/                # Stable fixtures grouped by area
│   ├── api_*.json       # Per-area API tests (parsing, hierarchy, model, errors, comments,
│   │                    # typed access, list access, advanced processing, edge cases,
│   │                    # filter predicates, whitespace behaviors, proposed/reference variants)
│   ├── api_fuzz_*.json  # Generated fuzz seeds
│   └── property_*.json  # Algebraic + round-trip properties
└── experimental/        # In-development fixtures (api_experimental.json, etc.)
```
Run `ls source_tests/core/` for the current file list.

## Command Reference

| Command | Purpose | When to Use |
|---------|---------|-------------|
| `just reset` | Clean, build, lint, and test (ensures clean state) | Before commits, quick verification |
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
- **Core Parsing**: `parse`, `parse_indented`, `build_hierarchy`, `build_model`
- **Typed Access**: `get_string`, `get_int`, `get_bool`, `get_float`, `get_list`
- **Processing**: `filter`, `compose`, `expand_dotted`
- **Formatting/IO**: `canonical_format`, `load`, `round_trip`

**Note:** Mock implementation (`internal/mock/ccl.go`) provides: Parse, ParseIndented, Filter, Compose, ExpandDotted, BuildHierarchy, BuildModel, GetString, GetInt, GetBool, GetFloat, GetList, PrettyPrint, Print, RoundTrip, ComposeAssociative, IdentityLeft, IdentityRight.

**Function Details:**
- **`parse`**: Basic lexical parsing - returns flat entries where values are raw strings
- **`parse_indented`**: Indentation-normalized parsing - calculates common leading whitespace and strips it from all lines (like Python's `textwrap.dedent`)
- **`build_hierarchy`**: Recursively parses entry values to create nested object structure
- **`build_model`**: Canonical OCaml-reference model construction — like `build_hierarchy` but follows the reference semantics for key folding and value coercion

> [!IMPORTANT]
> **Mock `Parse` multiline-keys gap:** `internal/mock/ccl.go`'s `Parse` does **not** fold unindented non-`=` lines into the next key (OCaml-canonical multiline-keys behavior). When adding a `parse` validation to a fixture that exercises multiline keys, the matching `*Parse` Go test will fail and must be appended to the basic-only skip list in `cmd/ccl-test-runner/main.go` (alongside the existing `TestKeyWithNewlineBeforeEqualsParse` etc.). This is a known mock limitation, not a bug in the test data.

### Test Metadata
- **`functions`** - Required CCL functions (filter: skip if unsupported)
- **`features`** - Language features exercised (informational only, for reporting gaps)
- **`behaviors`** - Implementation choices (filter via `conflicts` field)
- **`variants`** - Spec interpretation (filter via `conflicts` field)
- **`conflicts`** - Mutually exclusive options (filter: skip if your choice is listed)

> [!NOTE]
> **Features vs behaviors:** Features annotate universal OCaml-canonical rules every conformant impl follows (e.g. `tab_in_value_preserved`, `toplevel_indent_strip`). Behaviors encode genuine implementation *choices* with mutually-exclusive pairs (e.g. `indent_spaces` vs `indent_tabs`, `continuation_tab_to_space` vs `continuation_tab_preserve`). When adding a tag, ask: "could a reasonable impl do otherwise?" — yes → behavior pair, no → feature.

## Function-Based Implementation

Tests are organized by the CCL functions they validate. Implement the functions your library needs — see CCL Function Groups above. The mock in `internal/mock/ccl.go` is the reference implementation used to validate expected values.

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
      "features": ["whitespace", "tab_in_value_preserved"]
    }
  ]
}
```

See `schemas/source-format.json` for complete schema definition.
