# CCL Test Suite JSON Schema Reference

**Validation-Based Testing Format**

This document provides implementation-focused documentation for the **validation-based JSON schema** used in the CCL test suite files. All test files in the `tests/` directory use the validation format that makes multi-level testing explicit and eliminates confusion.

> **📋 Complete Technical Reference:** For exhaustive field-by-field documentation, see [`generated-schema.md`](generated-schema.md) - automatically generated from the schema with complete type information and field listings.

> **🛠️ Implementation Guide:** This document focuses on practical usage, examples, and implementation patterns for working with the CCL test suite schema.

## Schema Overview

The CCL test suite uses a **validation-based JSON schema** (`tests/schema.json`) that makes API function testing explicit. Instead of confusing multi-level fields (`expected_flat`, `expected_nested`, etc.), each test specifies exactly which API functions to validate.

**Schema Location**: `tests/schema.json`  
**Schema Version**: JSON Schema Draft 07

## Key Benefits of Validation-Based Testing

✅ **Explicit API Testing**: Each validation maps to exact API function  
✅ **No Multi-Level Confusion**: Eliminates `expected_flat` vs `expected_nested` guesswork  
✅ **Easy Test Runners**: Direct iteration over `validations` object keys  
✅ **Clear Intent**: Obvious what functions to test and what results to expect

## Root Object Structure

```json
{
  "suite": "string",
  "version": "string", 
  "description": "string",
  "tests": [...]
}
```

### Root Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `suite` | string | ✓ | Name of the test suite (e.g., "CCL Essential Parsing (Validation Format)") |
| `version` | string | ✓ | Version of the test suite format |
| `description` | string |  | Description of the test suite purpose and scope |
| `tests` | array | ✓ | Array of validation-based test cases |

## Validation-Based Test Case Structure

Each test case uses the explicit `validations` object instead of multiple `expected_*` fields:

```json
{
  "name": "string",
  "input": "string",
  "input1": "string",  // For composition tests
  "input2": "string",  // For composition tests  
  "input3": "string",  // For associativity tests
  "validations": {
    "parse": [...],              // Level 1: Entry parsing
    "filter": [...],             // Level 2: Comment filtering
    "compose": {...},            // Level 2: Entry composition
    "expand_dotted": [...],      // Level 2: Dotted key expansion
    "make_objects": {...},       // Level 3: Object construction
    "get_string": {...},         // Level 4: String access
    "get_int": {...},            // Level 4: Integer access
    "get_bool": {...},           // Level 4: Boolean access
    "get_float": {...},          // Level 4: Float access
    "pretty_print": "string",    // Output formatting
    "round_trip": {...},         // Parse-format-parse identity
    "canonical_format": {...},   // Canonical formatting
    "associativity": {...}       // Composition laws
  },
  "meta": {...}
}
```

### Test Case Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | ✓ | Unique identifier for the test case within the suite |
| `input` | string | ✓ | Primary CCL input string to parse |
| `input1`, `input2`, `input3` | string |  | Additional inputs for composition/associativity tests |
| `validations` | object | ✓ | Object containing API function validations to perform |
| `meta` | object | ✓ | Test metadata including level and categorization |

## Validation Types Reference

### Level 1: Entry Parsing

#### `parse` Validation
Tests the core `parse(text)` API function.

```json
"parse": [
  {"key": "database.host", "value": "localhost"},
  {"key": "database.port", "value": "8080"}
]
```

**Error Format:**
```json
"parse": {
  "error": true,
  "error_type": "ParseError",
  "error_message": "end_of_input"
}
```

### Level 2: Entry Processing  

#### `filter` Validation
Tests the `filter(entries)` API function for comment removal.

```json
"filter": [
  {"key": "key", "value": "value"}
]
```

#### `compose` Validation  
Tests the `compose(left, right)` API function for entry composition.

```json
"compose": {
  "left": [{"key": "key1", "value": "value1"}],
  "right": [{"key": "key2", "value": "value2"}],
  "expected": [
    {"key": "key1", "value": "value1"},
    {"key": "key2", "value": "value2"}
  ]
}
```

#### `expand_dotted` Validation
Tests dotted key expansion (e.g., `database.host` → nested structure).

```json
"expand_dotted": [
  {"key": "database", "value": "\n  host = localhost"}
]
```

### Level 3: Object Construction

#### `make_objects` Validation
Tests the `make_objects(entries)` API function for hierarchical structure creation.

```json
"make_objects": {
  "database": {
    "host": "localhost", 
    "port": "8080"
  }
}
```

### Level 4: Typed Access

#### Typed Access Validations
Test type-safe accessor functions with dual access patterns.

```json
"get_string": {
  "args": ["database.host"],     // Dotted access
  "expected": "localhost"
},
"get_int": {
  "args": ["database", "port"],  // Hierarchical access  
  "expected": 8080
},
"get_bool": {
  "args": ["enabled"],
  "expected": true
}
```

**Error Format:**
```json
"get_string": {
  "args": ["nonexistent.key"],
  "error": true,
  "error_type": "AccessError",
  "error_message": "Path not found"
}
```

### Output & Formatting

#### `pretty_print` Validation
Tests canonical output formatting.

```json
"pretty_print": "key = value\nnested = \n  sub = val"
```

#### `round_trip` Validation  
Tests parse-format-parse identity.

```json
"round_trip": {
  "property": "identity",
  "description": "Parse → Pretty-print → Parse should be identical"
}
```

#### `canonical_format` Validation
Tests standardized formatting output.

```json
"canonical_format": {
  "expected": "key = value\nnested = val",
  "description": "Should produce standardized formatting"
}
```

### Composition Laws

#### `associativity` Validation
Tests algebraic properties of entry composition.

```json
"associativity": {
  "property": "semigroup_associativity",
  "left_assoc": "(A + B) + C", 
  "right_assoc": "A + (B + C)",
  "should_be_equal": true
}
```

## Entry Object Structure

The `Entry` object represents a parsed key-value pair:

```json
{
  "key": "string",
  "value": "string"
}
```

## Test Metadata Structure

Every test case must include a `meta` object with categorization and level information:

```json
{
  "tags": ["string"],
  "level": 1|2|3|4,
  "feature": "string",
  "difficulty": "basic|intermediate|advanced"
}
```

### Feature Categories

| Feature | Description | Files |
|---------|-------------|--------|
| `parsing` | Core parsing functionality | `essential-parsing.json`, `comprehensive-parsing.json` |
| `processing` | Entry composition and merging | `processing.json` |
| `comments` | Comment syntax and filtering | `comments.json` |
| `object-construction` | Flat entries to nested objects | `object-construction.json` |
| `dotted-keys` | Dotted key expansion | `dotted-keys.json` |
| `typed-parsing` | Type-aware value extraction | `typed-access.json` |
| `pretty-printing` | Formatting and round-trip tests | `pretty-print.json` |
| `error-handling` | Error detection and reporting | `errors.json` |

## Usage Examples

### Basic Level 1 Test

```json
{
  "name": "basic_key_value",
  "input": "key = value",
  "validations": {
    "parse": [
      {"key": "key", "value": "value"}
    ]
  },
  "meta": {
    "tags": ["basic"],
    "level": 1,
    "feature": "parsing"
  }
}
```

### Multi-Level Test

```json
{
  "name": "nested_with_typed_access",
  "input": "database.host = localhost\ndatabase.port = 8080",
  "validations": {
    "parse": [
      {"key": "database.host", "value": "localhost"},
      {"key": "database.port", "value": "8080"}
    ],
    "expand_dotted": [
      {"key": "database", "value": "\n  host = localhost\n  port = 8080"}
    ],
    "make_objects": {
      "database": {"host": "localhost", "port": "8080"}
    },
    "get_string": {
      "args": ["database.host"],
      "expected": "localhost"
    },
    "get_int": {
      "args": ["database", "port"],
      "expected": 8080
    }
  },
  "meta": {
    "tags": ["dotted-keys", "typed_parsing"],
    "level": 4,
    "feature": "typed-parsing"
  }
}
```

### Error Test

```json
{
  "name": "incomplete_key_error",
  "input": "key",
  "validations": {
    "parse": {
      "error": true,
      "error_type": "ParseError",
      "error_message": "end_of_input"
    }
  },
  "meta": {
    "tags": ["error", "incomplete"],
    "level": 1,
    "feature": "error-handling"
  }
}
```

## Implementation Guidelines

### Test Runner for Validation Format

```javascript
function runValidationTest(testCase) {
  // Iterate over all validations in the test
  for (const [validationType, expected] of Object.entries(testCase.validations)) {
    switch (validationType) {
      case 'parse':
        if (expected.error) {
          // Test error case
          expect(() => parse(testCase.input)).toThrow(expected.error_message);
        } else {
          // Test success case
          const actual = parse(testCase.input);
          expect(actual).toEqual(expected);
        }
        break;
        
      case 'make_objects':
        const entries = parse(testCase.input);
        const actual = makeObjects(entries);
        expect(actual).toEqual(expected);
        break;
        
      case 'get_string':
        const ccl = makeObjects(parse(testCase.input));
        if (expected.error) {
          expect(() => getString(ccl, ...expected.args)).toThrow();
        } else {
          const actual = getString(ccl, ...expected.args);
          expect(actual).toBe(expected.expected);
        }
        break;
        
      case 'filter':
        const filteredEntries = filter(parse(testCase.input));
        expect(filteredEntries).toEqual(expected);
        break;
        
      case 'round_trip':
        const originalEntries = parse(testCase.input);
        const formatted = prettyPrint(originalEntries);
        const reparsedEntries = parse(formatted);
        expect(reparsedEntries).toEqual(originalEntries);
        break;
        
      // ... handle other validation types
    }
  }
}
```

### Benefits of Validation-Based Testing

1. **🎯 Explicit Intent**: Each test clearly shows which API functions to test
2. **🚀 Easy Implementation**: Test runners iterate over validation keys
3. **📚 Self-Documenting**: Validation names explain what's being tested
4. **🔧 Flexible**: Easy to add additional validation types
5. **✅ No Confusion**: Eliminates guesswork about multiple expected fields
6. **📊 Clear Coverage**: See exactly what APIs are tested per test case

## Schema Validation

Validate all test files against the schema:

```bash
# Using AJV CLI
ajv validate -s tests/schema.json -d "tests/*.json"

# Using npm script
npm run validate
```

The validation-based format provides a foundation for testing CCL implementations with explicit API function testing, eliminating the confusion of multi-level field approaches.