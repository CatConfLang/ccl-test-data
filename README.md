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
✅ **Assertion counting** - Optional explicit counts for validation verification\
✅ **Self-documenting** - Validation names explain what's being tested\
✅ **148+ test assertions** - Comprehensive coverage across all CCL features

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

**Core Parsing** (49 tests)

- **`tests/essential-parsing.json`** (18 tests) - Basic parsing functionality for rapid prototyping
- **`tests/comprehensive-parsing.json`** (31 tests) - Thorough parsing with edge cases, whitespace variations, and OCaml stress test

**Advanced Processing** (24 tests)

- **`tests/processing.json`** (21 tests) - Entry composition, merging, and advanced processing
- **`tests/comments.json`** (3 tests) - Comment syntax and filtering functionality

**Object Construction** (26 tests)

- **`tests/object-construction.json`** (8 tests) - Converting flat entries to nested objects
- **`tests/dotted-keys.json`** (18 tests) - Dotted key expansion and conflict resolution

**Type System** (17 tests)

- **`tests/typed-access.json`** (17 tests) - Type-aware value extraction with smart inference

**Output & Validation** (32 tests)

- **`tests/pretty-print.json`** (15 tests) - Formatting and round-trip tests
- **`tests/algebraic-properties.json`** (12 tests) - Algebraic property validation (associativity, monoid laws)
- **`tests/errors.json`** (5 tests) - Error handling validation

### Using the Test Suite

1. **Load test files** into your language's testing framework
1. **Iterate over validations** - Each test specifies which API functions to validate
1. **Filter by level** to test only the features you implement
1. **Use metadata** to focus on specific functionality

**Test format structure:**

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
for (const [validationType, validationData] of Object.entries(test.validations)) {
  switch (validationType) {
    case 'parse':
      const actual = parse(test.input);
      // Handle both legacy array format and new counted format
      const expected = Array.isArray(validationData) ? validationData : validationData.items;
      expect(actual).toEqual(expected);
      break;
    case 'make_objects':
      const entries = parse(test.input);
      const objects = makeObjects(entries);
      // Handle both legacy object format and new counted format  
      const expectedObj = validationData.result || validationData;
      expect(objects).toEqual(expectedObj);
      break;
    case 'get_string':
      const ccl = makeObjects(parse(test.input));
      // Handle legacy single test, legacy array, and new counted format
      const tests = validationData.cases || (Array.isArray(validationData) ? validationData : [validationData]);
      tests.forEach(test => {
        const value = getString(ccl, ...test.args);
        expect(value).toBe(test.expected);
      });
      break;
  }
}
```

### **Assertion Counting**

The test suite supports **explicit assertion counting** to help implementers validate they're running the correct number of assertions:

**Legacy Format** (backwards compatible):

```json
{
  "validations": {
    "parse": [{"key": "name", "value": "Alice"}],
    "get_string": {"args": ["name"], "expected": "Alice"}
  }
}
```

**New Counted Format** (optional, more explicit):

```json
{
  "validations": {
    "parse": {
      "count": 1,
      "items": [{"key": "name", "value": "Alice"}]
    },
    "get_string": {
      "count": 1, 
      "cases": [{"args": ["name"], "expected": "Alice"}]
    }
  }
}
```

**Benefits:**

- **Explicit counting**: Each validation declares exactly how many assertions it represents
- **Self-validating**: Test runners can verify `count` matches actual array lengths
- **Backwards compatible**: Existing tests continue working unchanged

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
