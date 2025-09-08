# CCL Test Suite

Language-agnostic test suite for the Categorical Configuration Language (CCL).

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

### Quick Start

```bash
# Clone the test suite
git clone <this-repo>
cd ccl-test-data

# Run validation script
npm test
```

### Test Files

The test suite is organized by feature category:

**Core Parsing** (48 tests)
- **`tests/essential-parsing.json`** (18 tests) - Basic parsing functionality for rapid prototyping
- **`tests/comprehensive-parsing.json`** (30 tests) - Thorough parsing with edge cases and whitespace variations

**Advanced Processing** (24 tests)
- **`tests/processing.json`** (21 tests) - Entry composition, merging, and advanced processing
- **`tests/comments.json`** (3 tests) - Comment syntax and filtering functionality

**Object Construction** (26 tests)
- **`tests/object-construction.json`** (8 tests) - Converting flat entries to nested objects
- **`tests/dotted-keys.json`** (18 tests) - Dotted key expansion and conflict resolution

**Type System** (17 tests)
- **`tests/typed-access.json`** (17 tests) - Type-aware value extraction with smart inference

**Output & Validation** (20 tests)
- **`tests/pretty-print.json`** (15 tests) - Formatting and round-trip tests
- **`tests/errors.json`** (5 tests) - Error handling validation

### Using the Test Suite

1. **Load test files** into your language's testing framework
2. **Filter by level** to test only the features you implement
3. **Use metadata** to focus on specific functionality
4. **Validate JSON structure** matches expected format

Example test case structure:
```json
{
  "name": "basic_key_value",
  "input": "key = value",
  "expected": [{"key": "key", "value": "value"}],
  "meta": {"tags": ["basic"], "level": 1}
}
```

## Documentation

### Test Suite Schema

- **[Schema Technical Reference](docs/generated-schema.md)** - Complete auto-generated field documentation
- **[Schema Implementation Guide](docs/schema-reference.md)** - Practical usage examples and patterns

### Test Suite Architecture

- **[Test Architecture](docs/test-architecture.md)** - How to use this test suite

### General Implementation Guidance

- **[Implementation Guide](https://ccl.tylerbutler.com/implementing-ccl)** - Complete CCL implementation guide
- **[Test Architecture](https://ccl.tylerbutler.com/test-architecture)** - General testing concepts

### Examples

See `docs/examples/` for sample configurations and expected parsing results.

## Contributing

When adding test cases:

1. Add to appropriate JSON file by feature level
2. Include descriptive name and metadata
3. Validate JSON structure
4. Update test counts in documentation

## Validation

```bash
# Validate test suite structure
npm run validate

# Run specific test category
npm run test:parsing
npm run test:objects
npm run test:types
```

This test suite ensures consistent CCL behavior across all language implementations.