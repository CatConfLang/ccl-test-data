# CCL Test Suite

Language-agnostic test suite for the Categorical Configuration Language (CCL). Each test specifies which CCL parsing functions to validate and their expected outputs.

## What is CCL?

For comprehensive CCL documentation, see the **[CCL Documentation](https://ccl.tylerbutler.com)** which includes:

- **[Specification Summary](https://ccl.tylerbutler.com/specification-summary)** - Complete language specification
- **[Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)** - Quick syntax guide
- **[Parsing Algorithm](https://ccl.tylerbutler.com/parsing-algorithm)** - Implementation guide
- **[Mathematical Theory](https://ccl.tylerbutler.com/theory)** - Theoretical foundations

**Original sources:**

- [CCL Blog Post](https://chshersh.com/blog/2025-01-06-the-most-elegant-configuration-language.html) - Original specification by Dmitrii Kovanikov
- [OCaml Reference Implementation](https://github.com/chshersh/ccl) - Canonical implementation

## Test Suite

This repository contains the **official JSON test suite** for CCL implementations across all programming languages.

### Test Format Features

✅ **Direct API mapping** - Each validation maps to a specific API function\
✅ **Multi-level testing** - Tests declare expected outputs for different parsing levels\
✅ **Simple test runners** - Direct iteration over `validations` object keys\
✅ **Assertion counting** - Required explicit counts for validation verification\
✅ **Self-documenting** - Validation names explain what's being tested\
✅ **446 test assertions** - Comprehensive coverage across all CCL features

### Quick Start

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
```

### Test Files

The test suite is organized by feature category:

**Core Parsing**

- **`tests/api-essential-parsing.json`** - Basic parsing functionality for rapid prototyping
- **`tests/api-comprehensive-parsing.json`** - Thorough parsing with edge cases and whitespace variations

**Advanced Processing**

- **`tests/api-processing.json`** - Entry composition, merging, and advanced processing
- **`tests/api-comments.json`** - Comment syntax and filtering functionality

**Object Construction**

- **`tests/api-object-construction.json`** - Converting flat entries to nested objects
- **`tests/api-dotted-keys.json`** - Dotted key expansion and conflict resolution

**Type System**

- **`tests/api-typed-access.json`** - Type-aware value extraction with smart inference

**Error Handling**

- **`tests/api-errors.json`** - Error handling validation

### Using the Test Suite

The test suite now uses a **counted format** with required `count` fields for all validations:

**Test format structure:**

```json
{
  "name": "basic_multi_level_test",
  "input": "database.host = localhost",
  "validations": {
    "parse": {
      "count": 1,
      "expected": [{"key": "database.host", "value": "localhost"}]
    },
    "make_objects": {
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
  "meta": {"tags": ["dotted-keys"], "level": 4}
}
```

**Test runner example:**

```javascript
for (const [validationType, validation] of Object.entries(test.validations)) {
  switch (validationType) {
    case 'parse':
      if (validation.error) {
        expect(() => parse(test.input)).toThrow(validation.error_message);
      } else {
        const actual = parse(test.input);
        expect(actual).toEqual(validation.expected);
        expect(actual.length).toBe(validation.count); // Verify count
      }
      break;
    case 'make_objects':
      const entries = parse(test.input);
      const objects = makeObjects(entries);
      expect(objects).toEqual(validation.expected);
      break;
    case 'get_string':
      const ccl = makeObjects(parse(test.input));
      if (validation.error) {
        expect(() => getString(ccl, ...validation.args)).toThrow();
      } else {
        validation.cases.forEach(testCase => {
          if (testCase.error) {
            expect(() => getString(ccl, ...testCase.args)).toThrow();
          } else {
            const value = getString(ccl, ...testCase.args);
            expect(value).toBe(testCase.expected);
          }
        });
        expect(validation.cases.length).toBe(validation.count); // Verify count
      }
      break;
  }
}
```

### Assertion Counting

All validations now use the **counted format** with required `count` fields:

- **For array results** (`parse`, `filter`, `expand_dotted`): `count` = number of items in `expected` array
- **For object results** (`make_objects`): `count` = typically 1 (single object)
- **For typed access**: `count` = number of test cases in `cases` array
- **For empty results**: `count` = 0 (e.g., empty input parsing)

**Benefits:**

- **Explicit counting**: Each validation declares exactly how many assertions it represents
- **Self-validating**: Test runners can verify `count` matches actual array lengths
- **Test complexity tracking**: Enables precise measurement of implementation complexity

## Go Test Runner

This repository includes a comprehensive Go-based test runner for CCL implementations:

### Commands

```bash
# Basic usage
just build          # Build the test runner binary
just generate       # Generate all Go test files
just test           # Run all tests
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
just validate       # Validate test files against schema
just clean          # Clean generated files
```

### Mock Implementation

The repository includes a basic Level 1 mock CCL implementation for testing and development:

- **Location**: `internal/mock/ccl.go`
- **Features**: Basic key-value parsing, comment handling, empty input support
- **Usage**: Demonstrates test integration patterns

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

When adding test cases:

1. Add to appropriate JSON file by feature level
1. Include descriptive name and metadata
1. Use counted format with appropriate `count` values
1. Validate JSON structure with `just validate`
1. Update test counts in documentation

## Validation

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

### Current Test Statistics

- **161 total tests** with **446 assertions**
- **88 active tests** (307 assertions) for standard implementations
- **73 skipped tests** (139 assertions) for advanced/optional features

This test suite ensures consistent CCL behavior across all language implementations.
