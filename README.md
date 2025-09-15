# CCL Test Suite

> \[!NOTE]
> This is the **official JSON test suite** for CCL implementations across all programming languages, featuring comprehensive **feature-based tagging** for precise test selection and progressive implementation support.

Language-agnostic test suite for the Categorical Configuration Language (CCL) with **feature-based tagging** for precise test selection. Each test specifies which CCL functions to validate and uses structured tags to enable progressive implementation.

## What is CCL?

> \[!TIP]
> New to CCL? Start with the **[Specification Summary](https://ccl.tylerbutler.com/specification-summary)** for a complete overview, then check the **[Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)** for quick implementation guidance.

For comprehensive CCL documentation, see the **[CCL Documentation](https://ccl.tylerbutler.com)** which includes:

- **[Specification Summary](https://ccl.tylerbutler.com/specification-summary)** - Complete language specification
- **[Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)** - Quick syntax guide
- **[Parsing Algorithm](https://ccl.tylerbutler.com/parsing-algorithm)** - Implementation guide
- **[Mathematical Theory](https://ccl.tylerbutler.com/theory)** - Theoretical foundations

### Original Sources

- [CCL Blog Post](https://chshersh.com/blog/2025-01-06-the-most-elegant-configuration-language.html) - Original specification by Dmitrii Kovanikov
- [OCaml Reference Implementation](https://github.com/chshersh/ccl) - Canonical implementation

## Test Suite

This repository contains the **official JSON test suite** for CCL implementations across all programming languages.

### Key Features

> \[!IMPORTANT]
> All tests use a **counted format** with required `count` fields for precise validation verification. Each validation declares exactly how many assertions it represents.

✅ **Dual-format architecture** - Source format for maintainability, generated flat format for implementation\
✅ **Direct API mapping** - Each validation maps to a specific API function\
✅ **Multi-level testing** - Tests declare expected outputs for different parsing levels\
✅ **Conflict resolution** - Automatic handling of mutually exclusive behaviors\
✅ **Progressive implementation** - Clear path from minimal parsing to full features\
✅ **Simple test runners** - Direct iteration over `validations` object keys\
✅ **Assertion counting** - Required explicit counts for validation verification\
✅ **Self-documenting** - Validation names explain what's being tested\
✅ **452 test assertions** - Comprehensive coverage across all CCL features

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

The test suite is organized by feature category:

#### Core Parsing

- **`tests/api-essential-parsing.json`** - Basic parsing functionality for rapid prototyping
- **`tests/api-comprehensive-parsing.json`** - Thorough parsing with edge cases and whitespace variations

#### Advanced Processing

- **`tests/api-processing.json`** - Entry composition, merging, and advanced processing
- **`tests/api-comments.json`** - Comment syntax and filtering functionality

#### Object Construction

- **`tests/api-object-construction.json`** - Converting flat entries to nested objects
- **`tests/api-dotted-keys.json`** - Dotted key expansion and conflict resolution

#### Type System

- **`tests/api-typed-access.json`** - Type-aware value extraction with smart inference

#### Error Handling

- **`tests/api-errors.json`** - Error handling validation

### Using the Test Suite

> \[!IMPORTANT]
> **Counted Format Required**: All validations must include a `count` field that matches the number of expected results. This enables precise assertion counting and self-validating test suites.

#### Source Format Structure (Maintainable)

```json
{
  "name": "basic_multi_level_test",
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
  "meta": {
    "tags": ["function:parse", "function:build_hierarchy", "function:get_string", "feature:dotted_keys"],
    "level": 3,
    "feature": "dotted_keys"
  }
}
```

#### Generated Format Structure (Implementation-Friendly)

```json
{
  "name": "basic_multi_level_test_parse",
  "input": "database.host = localhost",
  "validation": "parse",
  "expected": {
    "count": 1,
    "entries": [{"key": "database.host", "value": "localhost"}]
  },
  "functions": ["parse"],
  "features": ["dotted_keys"],
  "level": 3,
  "source_test": "basic_multi_level_test"
}
```

## Dual-Format Architecture

> \[!TIP]
> **Implementation Strategy**: Use the generated flat format for your test runner implementation. The separate typed fields provide excellent API ergonomics and type safety compared to parsing structured tags.

The test suite uses a **dual-format architecture** optimized for both maintainability and implementation:

### Source Format (Maintainable)

The **source format** maintains readability and ease of authoring:

- Multiple validations per test in a single object
- Structured tags for comprehensive metadata
- Located in `tests/` directory

### Generated Format (Implementation-Friendly)

The **generated format** provides optimal implementation ergonomics:

- One test per validation function (1:N transformation)
- Separate typed fields instead of string parsing
- Type-safe enums with validation
- Direct field access for filtering

### Test Metadata Categories

#### Functions Array - Required CCL functions:

- `parse` - Basic key-value parsing (Level 1)
- `parse_value` - Indentation-aware parsing (Level 2)
- `build_hierarchy` - Object construction from flat entries (Level 3)
- `filter` - Entry filtering (Level 2)
- `combine` - Entry composition (Level 2)
- `expand_dotted` - Dotted key expansion (Level 2, optional)
- `get_string`, `get_int`, `get_bool`, `get_float`, `get_list` - Typed access (Level 4)
- `pretty_print` - Formatting (Level 5)

#### Features Array - Optional language features:

- `comments` - `/=` comment syntax
- `experimental_dotted_keys` - `foo.bar.baz` key syntax
- `empty_keys` - `= value` anonymous list items
- `multiline` - Multi-line value support
- `unicode` - Unicode content handling
- `whitespace` - Complex whitespace preservation

#### Behaviors Array - Implementation choices (mutually exclusive):

- `crlf_preserve_literal` vs `crlf_normalize_to_lf` - Line ending handling
- `tabs_preserve` vs `tabs_to_spaces` - Tab handling
- `strict_spacing` vs `loose_spacing` - Whitespace sensitivity
- `boolean_strict` vs `boolean_lenient` - Boolean parsing
- `list_coercion_enabled` vs `list_coercion_disabled` - List access behavior

#### Variants Array - Specification variants:

- `proposed_behavior` - Proposed specification behavior
- `reference_compliant` - OCaml reference implementation behavior

### Test Filtering Examples

> \[!NOTE]
> **Type-Safe Filtering**: Use direct field access instead of string parsing for better performance and type safety. The generated format provides enum validation and excellent API ergonomics.

#### Minimal Implementation (Parse only)

```javascript
// Filter for basic parsing tests only
const parseTests = flatTests.filter(test => 
  test.functions.includes('parse') && 
  test.functions.length === 1
);
```

#### Basic Implementation (Parse + Objects + Typed Access)

```javascript
// Filter for core functionality with dotted keys
const coreTests = flatTests.filter(test => 
  test.functions.some(f => ['parse', 'build_hierarchy', 'get_string'].includes(f)) &&
  !test.features.some(f => ['comments', 'unicode'].includes(f))
);
```

#### Advanced Implementation (All functions, behavior choices)

```javascript
// Filter based on implementation behavior choices
const advancedTests = flatTests.filter(test => {
  const hasConflictingBehavior = test.conflicts.behaviors?.some(b => 
    ['crlf_preserve_literal', 'boolean_lenient'].includes(b)
  );
  return !hasConflictingBehavior;
});
```

#### Progressive Implementation Strategy

```javascript
// Level-based progressive implementation
const level1Tests = flatTests.filter(test => test.level <= 1);
const level2Tests = flatTests.filter(test => test.level <= 2);
const level3Tests = flatTests.filter(test => test.level <= 3);
```

### Conflict Resolution

> \[!WARNING]
> **Mutually Exclusive Behaviors**: Tests with conflicting behaviors should be filtered based on your implementation choices. The conflicts object provides categorized exclusion lists for precise filtering.

Tests specify conflicting behaviors in categorized structure:

```json
{
  "name": "crlf_preservation_test_parse",
  "behaviors": ["crlf_preserve_literal"],
  "conflicts": {
    "behaviors": ["crlf_normalize_to_lf"]
  }
}
```

If your implementation chooses `crlf_normalize_to_lf`, filter out tests with that value in `conflicts.behaviors`.

### Generating Flat Format Tests

> \[!IMPORTANT]
> **Required Step**: You must generate the flat format tests from the source format before running your test suite. The flat format provides the type-safe, implementation-friendly structure your test runner needs.

Generate flat format tests from the maintainable source format:

```bash
# Generate flat format tests for implementation
just generate-flat

# Validate generated tests against schema
just validate-flat

# Run your test suite against generated tests
# (your implementation reads from generated-tests/ directory)
```

The generator:

- Transforms 1:N (one source test → multiple flat tests)
- Parses structured tags into separate typed fields
- Adds auto-generated function tags
- Provides categorized conflicts structure
- Maintains full traceability to source tests

### Test Runner Implementation Example

```javascript
// Load flat format tests (type-safe with excellent API ergonomics)
const flatTests = loadFlatTests('generated-tests/');

// Filter tests based on implementation capabilities
const supportedTests = flatTests.filter(test => {
  // Check if we support all required functions
  const unsupportedFunctions = test.functions.filter(f => 
    !implementedFunctions.includes(f)
  );
  if (unsupportedFunctions.length > 0) return false;

  // Check if we support all required features  
  const unsupportedFeatures = test.features.filter(f => 
    !implementedFeatures.includes(f)
  );
  if (unsupportedFeatures.length > 0) return false;

  // Check for conflicting behaviors
  const hasConflicts = test.conflicts.behaviors?.some(b => 
    implementationBehaviors.includes(b)
  );
  if (hasConflicts) return false;

  return true;
});

// Run tests with type-safe validation switching
supportedTests.forEach(test => {
  switch (test.validation) {
    case 'parse':
      if (test.expect_error) {
        expect(() => parse(test.input)).toThrow(test.error_type);
      } else {
        const actual = parse(test.input);
        expect(actual).toEqual(test.expected.entries);
        expect(actual.length).toBe(test.expected.count);
      }
      break;
    case 'build_hierarchy':
      const entries = parse(test.input);
      const objects = buildHierarchy(entries);
      expect(objects).toEqual(test.expected.object);
      break;
    case 'get_string':
      const ccl = buildHierarchy(parse(test.input));
      const value = getString(ccl, ...test.args);
      expect(value).toBe(test.expected.value);
      break;
  }
});
```

## Assertion Counting

> \[!IMPORTANT]
> **Self-Validating Tests**: The `count` field enables test runners to verify they're executing the expected number of assertions, preventing silent test failures and ensuring comprehensive coverage.

All validations use the **counted format** with required `count` fields:

### Count Field Guidelines

- **For array results** (`parse`, `filter`, `expand_dotted`): `count` = number of items in `expected` array
- **For object results** (`make_objects`): `count` = typically 1 (single object)
- **For typed access**: `count` = number of test cases in `cases` array
- **For empty results**: `count` = 0 (e.g., empty input parsing)

### Benefits

- **Explicit counting**: Each validation declares exactly how many assertions it represents
- **Self-validating**: Test runners can verify `count` matches actual array lengths
- **Test complexity tracking**: Enables precise measurement of implementation complexity

## Go Test Runner

> \[!TIP]
> **Quick Development Cycle**: Use `just dev-mock` for rapid prototyping or `just reset` to maintain a clean repository state with only passing tests enabled.

This repository includes a comprehensive Go-based test runner for CCL implementations:

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

# Level-specific testing
just test-level1    # Run only Level 1 tests
just test-level2    # Run only Level 2 tests
just test-level3    # Run only Level 3 tests
just test-level4    # Run only Level 4 tests

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

The repository includes a basic Level 1 mock CCL implementation for testing and development:

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
just generate-level1  # Generate only basic Level 1 tests
just test-level1      # Run Level 1 tests (all should pass)
```

**This is the required state for commits and CI.** The `dev-basic` command generates only the most essential tests (basic functions: `parse`, `make-objects`) and skips advanced features that would fail in the current mock implementation. This ensures:

- **Clean commits**: All enabled tests pass before committing
- **Stable CI**: Continuous integration runs pass consistently
- **Development foundation**: Solid base for CCL implementation work

## Documentation

### Test Suite Schema

- **[Schema Technical Reference](docs/generated-schema.md)** - Complete auto-generated field documentation
- **[Schema Implementation Guide](docs/schema-reference.md)** - Practical usage examples and patterns

### Test Suite Architecture

- **[Test Architecture](docs/test-architecture.md)** - How to use this test suite
- **[Test Filtering](docs/test-filtering.md)** - Advanced test filtering patterns

### General Implementation Guidance

- **[Implementation Guide](https://ccl.tylerbutler.com/implementing-ccl)** - Complete CCL implementation guide
- **[Test Architecture](https://ccl.tylerbutler.com/test-architecture)** - General testing concepts

## Contributing

> \[!IMPORTANT]
> **Test Quality Standards**: All new tests must use the counted format, include proper typed fields metadata, and pass JSON schema validation before being accepted.

When adding test cases:

1. **Add to appropriate JSON file** by feature level and category
1. **Include descriptive name and metadata** with typed fields (functions, features, behaviors, variants)
1. **Use counted format** with appropriate `count` values matching result arrays
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
go run cmd/validate-schema/main.go tests/api-*.json

# Generate and run all tests
just dev

# Quick development cycle for basic features
just dev-basic
```

## Test Statistics

> \[!NOTE]
> **Comprehensive Coverage**: The test suite provides **452 assertions** across **167 tests**, ensuring thorough validation of CCL implementations from basic parsing to advanced features.

The test suite provides comprehensive coverage with **452 assertions** across **167 tests**:

```bash
# View detailed statistics
just stats
```

### Current Breakdown

**📊 Overall Statistics:**

- **167 total tests** with **452 assertions** across **10 files**
- **21 mutually exclusive tests** with behavioral/variant conflicts
- **11 CCL functions** from basic parsing to advanced formatting
- **6 language features** (comments, dotted-keys, unicode, etc.)
- **3 behavioral choices** (CRLF, tabs, spacing handling)
- **2 specification variants** (proposed vs reference behavior)

**📚 Level Distribution:**

- **Level 1**: 54 tests (basic parsing)
- **Level 2**: 30 tests (processing)
- **Level 3**: 27 tests (object construction)
- **Level 4**: 56 tests (typed access)

**⚙️ Function Coverage:**

- **parse**: 132 tests (most essential)
- **make-objects**: 66 tests
- **get-string, get-int, get-bool, get-float, get-list**: 38 tests (typed access)
- **pretty-print**: 24 tests
- **compose**: 12 tests
- **Other functions**: 35 tests (filter, expand-dotted, parse-value)

This test suite ensures consistent CCL behavior across all language implementations with precise control over which features to test.
