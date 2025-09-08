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

The test suite is organized by CCL feature level:

- **`tests/ccl-entry-parsing.json`** (18 tests) - Core parsing functionality
- **`tests/ccl-entry-processing.json`** (10 tests) - Comment filtering and composition
- **`tests/ccl-object-construction.json`** (8 tests) - Nested object building
- **`tests/ccl-typed-parsing-examples.json`** (12 tests) - Type-safe value access
- **`tests/ccl-pretty-printer.json`** (15 tests) - Formatting and round-trip tests
- **`tests/ccl-errors.json`** (5 tests) - Error handling validation

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

### Test Suite Specific

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