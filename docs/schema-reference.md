# CCL Test Suite JSON Schema Reference

This document provides implementation-focused documentation for the JSON schema used in the CCL test suite files. All test files in the `tests/` directory specify which CCL parsing functions to validate and their expected outputs.

> **📋 Complete Technical Reference:** For exhaustive field-by-field documentation, see [`generated-schema.md`](generated-schema.md) - automatically generated from the schema with complete type information and field listings.

> **🛠️ Implementation Guide:** This document focuses on practical usage, examples, and implementation patterns for working with the CCL test suite schema.

## Schema Overview

The CCL test suite uses a JSON schema (`tests/schema.json`) that makes API function testing explicit. Each test specifies exactly which API functions to validate.

**Schema Location**: `tests/schema.json`  
**Schema Version**: JSON Schema Draft 07

## Test Format Features

✅ **Direct API mapping**: Each validation maps to a specific API function  
✅ **Multi-level testing**: Tests declare expected outputs for different parsing levels
✅ **Simple test runners**: Direct iteration over `validations` object keys  
✅ **Clear intent**: Obvious what functions to test and what results to expect

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
| `tests` | array | ✓ | Array of test cases |

## Test Case Structure

Each test case uses explicit `validations` objects that specify which API functions to test:

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

## Validation Format Requirements

**IMPORTANT: All validations now use the counted format** with a required `count` field that specifies the number of assertions the validation represents.

### Count Field

The `count` field tracks assertion complexity for test generation and indicates how many individual tests this validation represents:

- **For array results** (`parse`, `filter`, `expand_dotted`): `count` = number of items in `expected` array
- **For object results** (`make_objects`): `count` = typically 1 (single object)
- **For typed access**: `count` = number of test cases in `cases` array
- **For empty results**: `count` = 0 (e.g., empty input parsing)

```json
// Examples of count values
"parse": {"count": 3, "expected": [...]},     // 3 entries parsed
"parse": {"count": 0, "expected": []},        // Empty input
"make_objects": {"count": 1, "expected": {...}}, // Single object
"get_bool": {"count": 2, "cases": [...]}      // 2 test cases
```

## Validation Types Reference

### Level 1: Entry Parsing

#### `parse` Validation
Tests the core `parse(text)` API function.

**Standard Format:**
```json
"parse": {
  "count": 2,
  "expected": [
    {"key": "database.host", "value": "localhost"},
    {"key": "database.port", "value": "8080"}
  ]
}
```

**Empty Input Format:**
```json
"parse": {
  "count": 0,
  "expected": []
}
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
"filter": {
  "count": 1,
  "expected": [
    {"key": "key", "value": "value"}
  ]
}
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
"expand_dotted": {
  "count": 1,
  "expected": [
    {"key": "database", "value": "\n  host = localhost"}
  ]
}
```

### Level 3: Object Construction

#### `make_objects` Validation
Tests the `make_objects(entries)` API function for hierarchical structure creation.

```json
"make_objects": {
  "count": 1,
  "expected": {
    "database": {
      "host": "localhost", 
      "port": "8080"
    }
  }
}
```

### Level 4: Typed Access

#### Typed Access Validations
Test type-safe accessor functions with dual access patterns.

**Single Test Case Format:**
```json
"get_string": {
  "count": 1,
  "cases": [
    {
      "args": ["database.host"],     // Dotted access
      "expected": "localhost"
    }
  ]
},
"get_int": {
  "count": 1,
  "cases": [
    {
      "args": ["database", "port"],  // Hierarchical access  
      "expected": 8080
    }
  ]
},
"get_bool": {
  "count": 1,
  "cases": [
    {
      "args": ["enabled"],
      "expected": true
    }
  ]
}
```

**Multiple Test Cases Format:**
```json
"get_bool": {
  "count": 2,
  "cases": [
    {
      "args": ["enabled"],
      "expected": true
    },
    {
      "args": ["disabled"], 
      "expected": false
    }
  ]
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

**Mixed Success/Error Cases:**
```json
"get_bool": {
  "count": 2,
  "cases": [
    {
      "args": ["enabled"],
      "expected": true
    },
    {
      "args": ["invalid_bool"],
      "error": true,
      "error_message": "Value is not a boolean"
    }
  ]
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
| `flexible-boolean-parsing` | Enhanced boolean parsing (yes/no/on/off) | `typed-access.json` |
| `crlf-normalization` | Line ending normalization | `essential-parsing.json` |
| `pretty-printing` | Formatting and round-trip tests | `pretty-print.json` |
| `error-handling` | Error detection and reporting | `errors.json` |

> **📖 Test Filtering Guide**: See [`test-filtering.md`](test-filtering.md) for detailed guidance on filtering tests by feature support and compliance requirements.

## Usage Examples

### Basic Level 1 Test

```json
{
  "name": "basic_key_value",
  "input": "key = value",
  "validations": {
    "parse": {
      "count": 1,
      "expected": [
        {"key": "key", "value": "value"}
      ]
    }
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
    "parse": {
      "count": 2,
      "expected": [
        {"key": "database.host", "value": "localhost"},
        {"key": "database.port", "value": "8080"}
      ]
    },
    "expand_dotted": {
      "count": 1,
      "expected": [
        {"key": "database", "value": "\n  host = localhost\n  port = 8080"}
      ]
    },
    "make_objects": {
      "count": 1,
      "expected": {
        "database": {"host": "localhost", "port": "8080"}
      }
    },
    "get_string": {
      "count": 1,
      "cases": [
        {
          "args": ["database.host"],
          "expected": "localhost"
        }
      ]
    },
    "get_int": {
      "count": 1,
      "cases": [
        {
          "args": ["database", "port"],
          "expected": 8080
        }
      ]
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

## Test Filtering Architecture

The CCL test suite uses a dual filtering system enabling implementations to run different test subsets based on:

1. **Feature capabilities** - What functionality does your implementation support?
2. **Behavioral compliance** - How does your implementation handle spec ambiguities?

### Feature-Based Filtering

Filter tests by `meta.feature` to match your implementation's capabilities:

```javascript
// Basic implementation
const basicFeatures = ["parsing", "typed-parsing", "object-construction"];
const myTests = tests.filter(test => basicFeatures.includes(test.meta.feature));

// Enhanced implementation with optional features  
const enhancedFeatures = [
  "parsing", "typed-parsing", "object-construction",
  "flexible-boolean-parsing", "crlf-normalization"
];
const myTests = tests.filter(test => enhancedFeatures.includes(test.meta.feature));
```

### Compliance-Based Filtering

Filter tests by `meta.tags` to choose your interpretation of spec ambiguities:

```javascript
// CCL proposed behavior (enhanced/flexible)
const proposedTests = tests.filter(test => 
  !test.meta.tags.includes("reference-compliant-behavior")
);

// OCaml reference compliant behavior (strict/baseline)
const referenceTests = tests.filter(test => 
  !test.meta.tags.includes("proposed-behavior") ||
   test.meta.tags.includes("reference-compliant-behavior")
);
```

### Tag-Feature Design Pattern

The filtering system follows consistent patterns:

- **Optional features**: Clear requirement tags (e.g., `needs-flexible-boolean-parsing` tag + `flexible-boolean-parsing` feature)
- **Baseline behavior**: Descriptive behavior tags with base feature category (e.g., `uses-strict-boolean-parsing` tag + `typed-parsing` feature)
- **Spec ambiguities**: `proposed-behavior` vs `reference-compliant-behavior` tags with same base feature

> **📖 Complete Filtering Guide**: See [`test-filtering.md`](test-filtering.md) for comprehensive documentation of the filtering architecture, patterns, and examples.

## Implementation Guidelines

### Test Runner

```javascript
function runValidationTest(testCase) {
  // Iterate over all validations in the test
  for (const [validationType, validation] of Object.entries(testCase.validations)) {
    switch (validationType) {
      case 'parse':
        if (validation.error) {
          // Test error case
          expect(() => parse(testCase.input)).toThrow(validation.error_message);
        } else {
          // Test success case with counted format
          const actual = parse(testCase.input);
          expect(actual).toEqual(validation.expected);
          expect(actual.length).toBe(validation.count); // Verify count matches
        }
        break;
        
      case 'make_objects':
        const entries = parse(testCase.input);
        const actual = makeObjects(entries);
        expect(actual).toEqual(validation.expected);
        // Count typically 1 for make_objects
        break;
        
      case 'get_string':
        const ccl = makeObjects(parse(testCase.input));
        if (validation.error) {
          // Error format
          expect(() => getString(ccl, ...validation.args)).toThrow();
        } else {
          // Counted format with test cases
          validation.cases.forEach(testCase => {
            if (testCase.error) {
              expect(() => getString(ccl, ...testCase.args)).toThrow();
            } else {
              const actual = getString(ccl, ...testCase.args);
              expect(actual).toBe(testCase.expected);
            }
          });
          expect(validation.cases.length).toBe(validation.count); // Verify count
        }
        break;
        
      case 'filter':
        const filteredEntries = filter(parse(testCase.input));
        expect(filteredEntries).toEqual(validation.expected);
        expect(filteredEntries.length).toBe(validation.count);
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

### Test Format Benefits

1. **🎯 Clear intent**: Each test shows which API functions to test
2. **🚀 Simple implementation**: Test runners iterate over validation keys
3. **📚 Self-documenting**: Validation names explain what's being tested
4. **🔧 Extensible**: Easy to add additional validation types
5. **✅ Explicit structure**: Clear mapping between tests and API functions
6. **📊 Comprehensive coverage**: See exactly what APIs are tested per test case
7. **🔢 Assertion tracking**: Count field enables precise test generation and complexity measurement

## Schema Validation

Validate all test files against the schema:

```bash
# Using AJV CLI
ajv validate -s tests/schema.json -d "tests/*.json"

# Using npm script
npm run validate
```

This test format provides a foundation for testing CCL implementations with explicit API function testing and clear structure.