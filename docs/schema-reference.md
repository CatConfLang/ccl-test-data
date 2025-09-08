# CCL Test Suite JSON Schema Reference

This document provides implementation-focused documentation for the JSON schema used in the CCL test suite files. All test files in the `tests/` directory conform to this unified schema structure.

> **📋 Complete Technical Reference:** For exhaustive field-by-field documentation, see [`generated-schema.md`](generated-schema.md) - automatically generated from the schema with complete type information and field listings.

> **🛠️ Implementation Guide:** This document focuses on practical usage, examples, and implementation patterns for working with the CCL test suite schema.

## Schema Overview

The CCL test suite uses a unified JSON schema (`tests/schema.json`) that supports all 4 CCL architecture levels with consistent structure and metadata. This allows for language-agnostic testing across different CCL implementations.

**Schema Location**: `tests/schema.json`  
**Schema Version**: JSON Schema Draft 07  
**Test Suite Version**: 1.0

## Root Object Structure

```json
{
  "suite": "string",
  "version": "string", 
  "description": "string",
  "tests": [...],
  "composition_tests": [...] // Optional
}
```

### Root Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `suite` | string | ✓ | Name of the test suite (e.g., "CCL Essential Parsing", "CCL Object Construction") |
| `version` | string | ✓ | Version of the test suite format (currently "1.0") |
| `description` | string |  | Description of the test suite purpose and scope |
| `tests` | array | ✓ | Array of test cases conforming to the unified test case schema |
| `composition_tests` | array |  | Optional composition tests for Level 2 (entry processing) algebraic properties |

## Test Case Structure

Each test case in the `tests` array follows this unified structure that supports all CCL levels:

```json
{
  "name": "string",
  "input": "string",
  "expected": [...],        // For Level 1-2 tests
  "expected_flat": [...],   // For Level 3+ tests
  "expected_nested": {...}, // For Level 3 tests
  "expected_typed": {...},  // For Level 4 tests
  "expected_error": boolean,// For error tests
  "error_message": "string",// Optional error description
  "parse_options": {...},   // For Level 4 typed tests
  "api_calls": [...],       // For Level 4 tests
  "meta": {...}             // Metadata object (required)
}
```

### Test Case Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | ✓ | Unique identifier for the test case within the suite |
| `input` | string | ✓ | CCL input string to parse |
| `expected` | Entry[] |  | Expected entries for Level 1-2 tests |
| `expected_flat` | Entry[] |  | Expected flat parsing result for Level 3+ tests |
| `expected_nested` | object |  | Expected nested object structure for Level 3 tests |
| `expected_typed` | object |  | Expected typed values for Level 4 tests |
| `expected_error` | boolean |  | True if this test should produce an error |
| `error_message` | string |  | Expected error message pattern |
| `parse_options` | object |  | Parsing options for Level 4 typed tests |
| `api_calls` | string[] |  | API calls being tested for Level 4 |
| `meta` | object | ✓ | Test metadata including level and categorization |

### Entry Object Structure

The `Entry` object represents a parsed key-value pair:

```json
{
  "key": "string",
  "value": "string"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `key` | string | ✓ | The key part of the entry |
| `value` | string | ✓ | The value part of the entry |

### Typed Value Structure

For Level 4 tests, `expected_typed` contains typed value objects:

```json
{
  "type": "StringVal|IntVal|FloatVal|BoolVal|EmptyVal",
  "value": "string|number|boolean|null"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | enum | ✓ | The inferred type: `StringVal`, `IntVal`, `FloatVal`, `BoolVal`, or `EmptyVal` |
| `value` | mixed | ✓ | The typed value (string, number, boolean, or null for EmptyVal) |

### Parse Options

For Level 4 typed parsing tests, optional parsing configuration:

```json
{
  "parse_integers": true,
  "parse_floats": true,
  "parse_booleans": true
}
```

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `parse_integers` | boolean | true | Enable integer parsing |
| `parse_floats` | boolean | true | Enable float parsing |
| `parse_booleans` | boolean | true | Enable boolean parsing |

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

### Metadata Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `tags` | string[] | ✓ | Tags for categorizing and filtering tests |
| `level` | integer | ✓ | CCL architecture level (1=Entry Parsing, 2=Processing, 3=Objects, 4=Typed) |
| `feature` | enum | ✓ | Feature category for organization |
| `difficulty` | enum |  | Test difficulty level |

### Feature Categories

The `feature` field must be one of these standardized categories:

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

### Common Tags

Tags provide fine-grained categorization within features. Common tags include:

**Parsing Related:**
- `basic` - Fundamental functionality
- `whitespace` - Whitespace handling
- `multiline` - Multi-line values
- `continuation` - Indented continuation lines
- `unicode` - Unicode character support
- `line-endings` - Line ending normalization

**Structure Related:**
- `nested` - Nested structures
- `dotted-keys` - Dotted key syntax
- `lists` - List-style entries (empty keys)
- `empty-key` - Empty key handling
- `empty-value` - Empty value handling

**Advanced Features:**
- `typed_parsing` - Type inference
- `comments` - Comment handling
- `composition` - Entry composition
- `duplicate-keys` - Duplicate key handling
- `merge` - Object merging

**Quality Related:**
- `edge-case` - Edge case testing
- `error` - Error condition testing
- `round-trip` - Round-trip validation

## Composition Tests (Level 2)

Level 2 processing tests may include a `composition_tests` array for testing algebraic properties:

```json
{
  "name": "string",
  "property": "monoid_identity_left|monoid_identity_right|semigroup_associativity|closure|concatenation_equivalence",
  "input1": "string",
  "input2": "string", 
  "input3": "string",    // For associativity tests
  "expected_combined": [...],
  "expected_left_assoc": [...],   // For associativity tests
  "expected_right_assoc": [...],  // For associativity tests
  "expected_text_concat": "string",
  "meta": {...}
}
```

### Composition Test Properties

| Property | Description |
|----------|-------------|
| `monoid_identity_left` | Tests left identity: `empty + x = x` |
| `monoid_identity_right` | Tests right identity: `x + empty = x` |
| `semigroup_associativity` | Tests associativity: `(a + b) + c = a + (b + c)` |
| `closure` | Tests that operations remain within the set |
| `concatenation_equivalence` | Tests text concatenation equivalence |

## Usage Examples

### Basic Level 1 Test

```json
{
  "name": "basic_key_value",
  "input": "key = value",
  "expected": [
    {"key": "key", "value": "value"}
  ],
  "meta": {
    "tags": ["basic"],
    "level": 1,
    "feature": "parsing"
  }
}
```

### Level 3 Object Construction Test

```json
{
  "name": "nested_objects",
  "input": "database =\n  host = localhost\n  port = 5432",
  "expected_flat": [
    {"key": "database", "value": "\n  host = localhost\n  port = 5432"}
  ],
  "expected_nested": {
    "database": {
      "host": "localhost",
      "port": "5432"
    }
  },
  "meta": {
    "tags": ["nested", "object-construction"],
    "level": 3,
    "feature": "object-construction"
  }
}
```

### Level 4 Typed Parsing Test

```json
{
  "name": "parse_integer",
  "input": "port = 8080",
  "expected_flat": [
    {"key": "port", "value": "8080"}
  ],
  "expected_nested": {
    "port": "8080"
  },
  "expected_typed": {
    "port": {"type": "IntVal", "value": 8080}
  },
  "meta": {
    "tags": ["typed_parsing", "integer"],
    "level": 4,
    "feature": "typed-parsing"
  }
}
```

### Error Test

```json
{
  "name": "incomplete_key",
  "input": "key",
  "expected_error": true,
  "error_message": "end_of_input",
  "meta": {
    "tags": ["error", "incomplete"],
    "level": 1,
    "feature": "error-handling"
  }
}
```

## Schema Validation

All test files should validate against the schema:

```bash
# Using AJV CLI
ajv validate -s tests/schema.json -d "tests/*.json"

# Using Node.js script
npm run validate
```

## Implementation Guidelines

When implementing CCL parsers, use this schema structure to:

1. **Load test cases** programmatically from JSON files
2. **Filter by level** to test only implemented features  
3. **Filter by feature** to focus on specific functionality
4. **Filter by tags** for targeted testing scenarios
5. **Validate outputs** against expected results

### Test Runner Example

```javascript
function runTestSuite(testFile, options = {}) {
  const suite = JSON.parse(fs.readFileSync(testFile));
  
  for (const test of suite.tests) {
    // Filter by level if specified
    if (options.level && test.meta.level !== options.level) {
      continue;
    }
    
    // Filter by feature if specified  
    if (options.feature && test.meta.feature !== options.feature) {
      continue;
    }
    
    // Run appropriate test based on level
    if (test.expected) {
      // Level 1-2 test
      const result = parse(test.input);
      assert.deepEqual(result, test.expected);
    } else if (test.expected_nested) {
      // Level 3 test
      const entries = parse(test.input);
      const objects = makeObjects(entries);
      assert.deepEqual(objects, test.expected_nested);
    }
    // ... handle other test types
  }
}
```

This schema provides a comprehensive, flexible foundation for testing CCL implementations across all architecture levels while maintaining consistency and language-agnostic compatibility.