# CCL Test Suite

**Validation-Based Testing Format**

Language-agnostic test suite for the Categorical Configuration Language (CCL) with a **validation-based format** that makes multi-level testing explicit and eliminates implementation confusion.

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

### Validation-Based Format Benefits

✅ **Crystal clear API testing** - Each validation maps to exact API function\
✅ **No multi-level confusion** - Eliminates `expected_flat` vs `expected_nested` guesswork\
✅ **Easy test runners** - Direct iteration over `validations` object keys\
✅ **Self-documenting** - Validation names explain what's being tested\
✅ **135+ test assertions** - Same 135 test inputs now validate multiple API functions

### Quick Start

```bash
# Clone the test suite
git clone <this-repo>
cd ccl-test-data

# Run validation script
npm test
```

### Test Files

The test suite is organized by feature category using the **validation-based format**:

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

### Using the Validation-Based Test Suite

1. **Load test files** into your language's testing framework
1. **Iterate over validations** - Each test specifies which API functions to validate
1. **Filter by level** to test only the features you implement
1. **Use metadata** to focus on specific functionality

**Validation format structure:**

```json
{
  "name": "basic_multi_level_test",
  "input": "database.host = localhost",
  "validations": {
    "parse": [{"key": "database.host", "value": "localhost"}],
    "make_objects": {"database": {"host": "localhost"}},
    "get_string": {
      "args": ["database.host"],
      "expected": "localhost"
    }
  },
  "meta": {"tags": ["dotted-keys"], "level": 4}
}
```

**Test runner example:**

```javascript
for (const [validationType, expected] of Object.entries(test.validations)) {
  switch (validationType) {
    case 'parse':
      const actual = parse(test.input);
      expect(actual).toEqual(expected);
      break;
    case 'make_objects':
      const entries = parse(test.input);
      const objects = makeObjects(entries);
      expect(objects).toEqual(expected);
      break;
    case 'get_string':
      const ccl = makeObjects(parse(test.input));
      const value = getString(ccl, ...expected.args);
      expect(value).toBe(expected.expected);
      break;
  }
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
1. Include descriptive name and metadata
1. Validate JSON structure
1. Update test counts in documentation

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
