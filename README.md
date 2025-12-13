# CCL Test Suite

> \[!NOTE]
> This is a **comprehensive JSON test suite** for CCL implementations across all programming languages, featuring **feature-based classification** for precise test selection and function-based implementation support.

Language-agnostic test suite for the Categorical Configuration Language (CCL) with **feature-based classification** for precise test selection. Each test specifies which CCL functions to validate and uses structured metadata to enable function-based implementation.

> \[!TIP]
> **New to this project?** Check the **[Developer Guide](docs/DEVELOPER_GUIDE.md)** for development workflow and **[Architecture](docs/ARCHITECTURE.md)** for system design details.

## What is CCL?

> \[!TIP]
> New to CCL? Start with the **[Getting Started Guide](https://ccl.tylerbutler.com/getting-started/)** for an overview, then check the **[Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)** for quick implementation guidance.

For comprehensive CCL documentation, see the **[CCL Documentation](https://ccl.tylerbutler.com)** which includes:

- **[Getting Started](https://ccl.tylerbutler.com/getting-started/)** - Language overview
- **[Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)** - Quick syntax guide
- **[Parsing Algorithm](https://ccl.tylerbutler.com/parsing-algorithm)** - Implementation details
- **[Implementing CCL](https://ccl.tylerbutler.com/implementing-ccl/)** - Build your own parser

### Original Sources

- [CCL Blog Post](https://chshersh.com/blog/2025-01-06-the-most-elegant-configuration-language.html) - Original specification by Dmitrii Kovanikov
- [OCaml Reference Implementation](https://github.com/chshersh/ccl) - Canonical implementation

## Test Suite

This repository contains a **comprehensive JSON test suite** for CCL implementations across all programming languages.

### Key Features

> \[!IMPORTANT]
> All tests include required `count` fields for precise validation verification. Each validation declares exactly how many assertions it represents.

✅ **Dual-format architecture** - Source format for maintainability, generated flat format for implementation\
✅ **Direct API mapping** - Each validation maps to a specific API function\
✅ **Multi-stage testing** - Tests declare expected outputs for different parsing stages\
✅ **Conflict resolution** - Automatic handling of mutually exclusive behaviors\
✅ **Function-based implementation** - Independent capabilities can be implemented in any order\
✅ **Simple test runners** - Direct iteration over `validations` object keys\
✅ **Assertion counting** - Required explicit counts for validation verification\
✅ **Self-documenting** - Validation names explain what's being tested

### Quick Start

> \[!TIP]
> Use `just reset` before committing to ensure all enabled tests pass. This maintains repository in a clean, stable state for CI and development.

```bash
# Clone the test suite
git clone <this-repo>
cd ccl-test-data

# Install dependencies and run tests
just deps
just test

# Generate tests for mock implementation development
just generate-mock
just test-mock

# Set repository to clean, passing state (required for commits)
just reset  # alias for dev-basic
```

### Test Files

Test files are organized in `source_tests/`:

```
source_tests/
├── core/
│   ├── api_core_ccl_parsing.json       # Basic parsing
│   ├── api_core_ccl_hierarchy.json     # Object construction
│   ├── api_core_ccl_integration.json   # Full pipeline tests
│   ├── api_typed_access.json           # Type-aware value extraction
│   ├── api_comments.json               # Comment syntax
│   ├── api_errors.json                 # Error handling
│   ├── api_edge_cases.json             # Edge cases
│   ├── api_list_access.json            # List operations
│   ├── api_whitespace_behaviors.json   # Whitespace handling
│   ├── api_advanced_processing.json    # Entry composition/merging
│   ├── api_proposed_behavior.json      # Proposed spec behaviors
│   ├── api_reference_compliant.json    # OCaml-compatible behaviors
│   ├── property_round_trip.json        # Round-trip validation
│   └── property_algebraic.json         # Algebraic properties
└── experimental/
    └── api_experimental.json           # Experimental features
```

### Using the Test Suite

> \[!IMPORTANT]
> **Count Fields Required**: All validations must include a `count` field that matches the number of expected results. This enables precise assertion counting and self-validating test suites.

#### Source Format Structure (Maintainable)

```json
{
  "name": "basic_multi_stage_test",
  "input": "database.host = localhost",
  "validations": {
    "parse": {
      "count": 1,
      "expected": [{"key": "database.host", "value": "localhost"}]
    },
    "build_hierarchy": {
      "count": 1,
      "expected": {"database": {"host": "localhost"}}
    },
    "get_string": {
      "count": 1,
      "cases": [
        {
          "args": ["database.host"],
          "expected": "localhost"
        }
      ]
    }
  },
  "features": ["dotted_keys"],
  "behaviors": [],
  "variants": []
}
```

#### Generated Format Structure (Implementation-Friendly)

```json
{
  "name": "basic_multi_stage_test_parse",
  "input": "database.host = localhost",
  "validation": "parse",
  "expected": {
    "count": 1,
    "entries": [{"key": "database.host", "value": "localhost"}]
  },
  "functions": ["parse"],
  "features": ["dotted_keys"],
  "source_test": "basic_multi_stage_test"
}
```

## Dual-Format Architecture

- **Source Format** (`source_tests/`): Multiple validations per test, human-maintainable
- **Generated Format** (`generated_tests/`): One test per validation, implementation-friendly

### Test Metadata

| Field | Purpose | Filter? |
|-------|---------|---------|
| `functions` | Required CCL functions | Yes - skip if unsupported |
| `features` | Language features exercised | No - informational for reporting |
| `behaviors` | Implementation choices | Via `conflicts` field |
| `variants` | Spec interpretation (temporary) | Via `conflicts` field |
| `conflicts` | Mutually exclusive options | Yes - skip if your choice is listed |

### Filtering Logic

```javascript
function shouldSkip(test, my) {
  if (!test.functions.every(f => my.functions.includes(f))) return true;
  if (test.conflicts?.behaviors?.some(b => my.behaviors.includes(b))) return true;
  if (test.conflicts?.variants?.some(v => my.variants.includes(v))) return true;
  return false;
}
```

See [Test Selection Guide](docs/test-selection-guide.md) for complete filtering documentation.

## Go Test Runner

### Available Commands

```bash
# Flat format generation (recommended for implementations)
just generate-flat   # Generate implementation-friendly flat format tests
just validate-flat   # Validate generated flat tests against schema

# Go mock implementation development  
just generate       # Generate Go test files for mock implementation
just test           # Run all Go tests
just list           # List available test packages

# Mock implementation development
just generate-mock  # Generate tests for mock implementation
just test-mock      # Run tests suitable for mock implementation
just dev-mock       # Full development cycle for mock

# Function group testing
just test --functions core            # Run core function tests
just test --functions typed           # Run typed access tests
just test --functions processing      # Run processing function tests
just test --functions formatting      # Run formatting function tests

# Feature-specific testing
just test-comments  # Run comment-related tests
just test-parsing   # Run parsing tests
just test-objects   # Run object construction tests

# Utilities
just stats          # Show test generation statistics
just validate       # Validate source test files against schema
just validate-all   # Validate both source and generated formats
just clean          # Clean generated files
```

### Mock Implementation

> \[!NOTE]
> **Learning Resource**: The mock implementation serves as both a working example and a foundation for development. It demonstrates proper test integration patterns and API structure.

The repository includes a basic mock CCL implementation for testing and development:

- **Location**: `internal/mock/ccl.go`
- **Features**: Basic key-value parsing, comment handling, empty input support
- **Usage**: Demonstrates test integration patterns and API structure

### Repository State Management

> \[!WARNING]
> **Critical for CI/CD**: The repository **must** be in a clean, passing state before commits. Use `just reset` to ensure all enabled tests pass and maintain stable CI builds.

The repository should be maintained in a clean, passing state. Use these commands to ensure all enabled tests pass:

```bash
# Standard repository state (all tests should pass)
just reset  # alias for dev-basic

# Or run the steps manually:
just generate --functions core  # Generate only basic core function tests
just test --functions core      # Run core function tests (all should pass)
```

**This is the required state for commits and CI.** The `dev-basic` command generates only the most essential tests (basic functions: `parse`, `build-hierarchy`) and skips advanced features that would fail in the current mock implementation. This ensures:

- **Clean commits**: All enabled tests pass before committing
- **Stable CI**: Continuous integration runs pass consistently
- **Development foundation**: Solid base for CCL implementation work

## Documentation

### Test Suite Documentation

- **[Test Architecture](docs/test-architecture.md)** - How to use this test suite
- **[Test Filtering](docs/test-filtering.md)** - Advanced test filtering patterns
- **[Schema Reference](docs/schema-reference.md)** - Practical usage examples
- **[Generated Schema](docs/generated-schema.md)** - Auto-generated field documentation

### CCL Implementation

- **[Implementing CCL](https://ccl.tylerbutler.com/implementing-ccl/)** - Complete CCL implementation guide

## Contributing

> \[!IMPORTANT]
> **Test Quality Standards**: All new tests must include proper count fields and typed fields metadata, and pass JSON schema validation before being accepted.

When adding test cases:

1. **Add to appropriate JSON file** by feature category
1. **Include descriptive name and metadata** with typed fields (functions, features, behaviors, variants)
1. **Include count fields** with appropriate `count` values matching result arrays
1. **Validate JSON structure** with `just validate` before submitting
1. **Generate flat format** with `just generate-flat` and ensure tests pass
1. **Update test counts** in documentation and ensure `just stats` reflects changes

## Validation

> \[!TIP]
> **Development Workflow**: Run `just validate` before committing changes to catch JSON schema violations early. Use `just dev-basic` for rapid iteration during development.

```bash
# Validate test suite structure
just validate

# Run schema validation
go run cmd/validate-schema/main.go tests/api_*.json

# Generate and run all tests
just dev

# Quick development cycle for basic features
just dev-basic
```

## Test Statistics

```bash
# View current test statistics
just stats
```

The test suite ensures consistent CCL behavior across all language implementations with precise control over which features to test.
