# CCL Test Suite Schema Reference

This document provides comprehensive documentation for both the **source format** (maintainable) and **generated flat format** (implementation-friendly) used in the CCL test suite's dual-format architecture.

> **📋 Complete Technical Reference:** For exhaustive field-by-field documentation, see [`generated-schema.md`](generated-schema.md) - automatically generated from the schema with complete type information and field listings.

> **🛠️ Implementation Guide:** This document focuses on practical usage, examples, and implementation patterns for working with both schema formats.

## Dual-Format Architecture Overview

The CCL test suite uses a **dual-format architecture** optimized for both maintainability and implementation:

- **Source Format** (`tests/`): Maintainable format with grouped validations per test
- **Generated Flat Format** (`generated-tests/`): Implementation-friendly with separate typed fields

**Schema Locations**: 
- Source: `tests/schema.json` (JSON Schema Draft 07)
- Generated: `generated-tests/flat-test-schema.json` (JSON Schema Draft 07)

## Architecture Benefits

### Source Format (Maintainable)
✅ **Grouped validations**: Multiple validations per test reduce input duplication  
✅ **Clear structure**: Logical organization with direct API mapping  
✅ **Readable format**: Easy to author and maintain test cases

### Generated Flat Format (Implementation-Friendly)
✅ **Type-safe filtering**: Direct field access with `test.functions[]`, `test.features[]`  
✅ **Excellent API ergonomics**: No string parsing, better performance
✅ **Simple test runners**: One test per validation function (1:N transformation)
✅ **Schema validation**: Enum constraints ensure consistent values

## Source Format Structure

The source format in `tests/` directory maintains readability for test authoring:

```json
{
  "suite": "string",
  "version": "string", 
  "description": "string",
  "tests": [...]
}
```

### Source Format Root Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `suite` | string | ✓ | Name of the test suite (e.g., "CCL Essential Parsing (Validation Format)") |
| `version` | string | ✓ | Version of the test suite format |
| `description` | string |  | Description of the test suite purpose and scope |
| `tests` | array | ✓ | Array of source test cases with grouped validations |

## Generated Flat Format Structure

The generated flat format in `generated-tests/` directory provides optimal implementation ergonomics:

```json
{
  "name": "basic_parsing_workflow_parse",
  "input": "database.host = localhost",
  "validation": "parse",
  "expected": {
    "count": 1,
    "entries": [{"key": "database.host", "value": "localhost"}]
  },
  "functions": ["parse"],
  "features": ["dotted-keys"],
  "behaviors": ["crlf-normalize-to-lf"],
  "variants": ["reference-compliant"],
  "level": 3,
  "source_test": "basic_parsing_workflow",
  "conflicts": {
    "behaviors": ["crlf-preserve-literal"],
    "variants": ["proposed-behavior"]
  }
}
```

### Generated Format Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | ✓ | Unique test name (source_test + validation type) |
| `input` | string | ✓ | CCL input string to parse |
| `validation` | string | ✓ | Single validation type (parse, make-objects, etc.) |
| `expected` | object | ✓ | Expected result with count field |
| `functions` | array | ✓ | Required CCL functions |
| `features` | array | ✓ | Required optional language features |
| `behaviors` | array | ✓ | Implementation behavior choices |
| `variants` | array | ✓ | Specification variant choices |
| `level` | number | ✓ | CCL implementation level (1-5) |
| `source_test` | string | ✓ | Original source test name |
| `conflicts` | object |  | Mutually exclusive behaviors/variants |

## Source Test Case Structure

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

## Type-Safe Test Filtering

The generated flat format enables type-safe filtering through direct field access:

### Function-Based Filtering

Filter tests by required CCL functions using `test.functions[]` array:

```javascript
// Basic implementation - core functions only
const implementedFunctions = ["parse", "make-objects", "get-string"];
const supportedTests = flatTests.filter(test => 
  test.functions.every(fn => implementedFunctions.includes(fn))
);

// Enhanced implementation - includes processing functions
const implementedFunctions = [
  "parse", "make-objects", "get-string", "get-int", "get-bool",
  "filter", "compose", "expand-dotted"
];
const supportedTests = flatTests.filter(test => 
  test.functions.every(fn => implementedFunctions.includes(fn))
);
```

### Feature and Behavior Filtering

Filter by optional features and implementation choices:

```javascript
// Implementation with optional features
const implementedFeatures = ["comments", "dotted-keys"];
const implementationBehaviors = ["crlf-normalize-to-lf", "boolean-strict"];
const implementationVariants = ["reference-compliant"];

const compatibleTests = flatTests.filter(test => {
  // Check feature support
  const featuresSupported = test.features.every(feature => 
    implementedFeatures.includes(feature)
  );
  
  // Check for behavioral conflicts
  const hasConflictingBehavior = test.conflicts?.behaviors?.some(b => 
    implementationBehaviors.includes(b)
  );
  const hasConflictingVariant = test.conflicts?.variants?.some(v => 
    implementationVariants.includes(v)
  );
  
  return featuresSupported && !hasConflictingBehavior && !hasConflictingVariant;
});
```

### Complete Implementation Example

```javascript
function getCompatibleTests(flatTests, capabilities) {
  return flatTests.filter(test => {
    // Check function support
    const functionsSupported = test.functions.every(fn => 
      capabilities.functions.includes(fn)
    );
    
    // Check feature support
    const featuresSupported = test.features.every(feature => 
      capabilities.features.includes(feature)
    );
    
    // Check for behavioral conflicts
    const hasConflictingBehavior = test.conflicts?.behaviors?.some(b => 
      capabilities.behaviors.includes(b)
    );
    const hasConflictingVariant = test.conflicts?.variants?.some(v => 
      capabilities.variants.includes(v)
    );
    
    return functionsSupported && featuresSupported && 
           !hasConflictingBehavior && !hasConflictingVariant;
  });
}

// Usage
const capabilities = {
  functions: ["parse", "make-objects", "get-string", "get-int"],
  features: ["comments"],
  behaviors: ["crlf-normalize-to-lf", "boolean-strict"],
  variants: ["reference-compliant"]
};

const runnableTests = getCompatibleTests(flatTests, capabilities);
```

### Benefits of Typed Fields Architecture

- **Type-safe filtering**: Direct array access vs string parsing
- **Better performance**: `test.functions.includes()` vs `test.meta.tags.includes('function:*')`
- **Excellent API ergonomics**: Intuitive filtering patterns
- **Schema validation**: Enum constraints ensure consistent values
- **Conflict resolution**: Categorized conflicts structure

> **📖 Complete Filtering Guide**: See [`test-filtering.md`](test-filtering.md) for comprehensive documentation of the typed fields filtering architecture, patterns, and examples.

## Implementation Guidelines

### Test Runners

#### Source Format Test Runner (Multi-validation)

```javascript
function runSourceTest(testCase) {
  // Iterate over all validations in the source test
  for (const [validationType, validation] of Object.entries(testCase.validations)) {
    switch (validationType) {
      case 'parse':
        if (validation.error) {
          expect(() => parse(testCase.input)).toThrow(validation.error_message);
        } else {
          const actual = parse(testCase.input);
          expect(actual).toEqual(validation.expected);
          expect(actual.length).toBe(validation.count);
        }
        break;
        
      case 'make_objects':
        const entries = parse(testCase.input);
        const actual = makeObjects(entries);
        expect(actual).toEqual(validation.expected);
        break;
        
      // ... handle other validation types
    }
  }
}
```

#### Generated Flat Format Test Runner (Single-validation)

```javascript
function runFlatTest(flatTest) {
  // Each flat test has exactly one validation
  switch (flatTest.validation) {
    case 'parse':
      if (flatTest.expect_error) {
        expect(() => parse(flatTest.input)).toThrow(flatTest.error_type);
      } else {
        const actual = parse(flatTest.input);
        expect(actual).toEqual(flatTest.expected.entries);
        expect(actual.length).toBe(flatTest.expected.count);
      }
      break;
      
    case 'make_objects':
      const entries = parse(flatTest.input);
      const actual = makeObjects(entries);
      expect(actual).toEqual(flatTest.expected.object);
      break;
      
    case 'get_string':
      const ccl = makeObjects(parse(flatTest.input));
      if (flatTest.expect_error) {
        expect(() => getString(ccl, ...flatTest.args)).toThrow();
      } else {
        const actual = getString(ccl, ...flatTest.args);
        expect(actual).toBe(flatTest.expected.value);
      }
      break;
      
    // ... handle other validation types
  }
}

// Usage with type-safe filtering
const compatibleTests = getCompatibleTests(flatTests, capabilities);
compatibleTests.forEach(test => runFlatTest(test));
```

### Dual-Format Benefits

#### Source Format Benefits
1. **📝 Maintainable**: Grouped validations reduce input duplication
2. **📚 Readable**: Clear structure for test authoring
3. **🔧 Extensible**: Easy to add validation types per test
4. **🎯 Clear intent**: Direct mapping of test to API functions

#### Generated Flat Format Benefits
1. **⚡ Type-safe filtering**: Direct field access vs string parsing
2. **🚀 Excellent performance**: Array methods faster than tag parsing
3. **🛠️ API ergonomics**: Intuitive filtering patterns
4. **🔍 Simple test runners**: One validation per test (no complex logic)
5. **📋 Schema validation**: Enum constraints ensure consistency
6. **🔢 Assertion tracking**: Count field enables precise measurement
7. **⚙️ Conflict resolution**: Categorized conflicts structure

## Schema Validation

Validate all test files against the schema:

```bash
# Using AJV CLI
ajv validate -s tests/schema.json -d "tests/*.json"

# Using npm script
npm run validate
```

This test format provides a foundation for testing CCL implementations with explicit API function testing and clear structure.