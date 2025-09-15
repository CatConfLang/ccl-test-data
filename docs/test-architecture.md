# CCL Test Suite Architecture

The CCL test suite uses a **dual-format architecture** optimized for both maintainability and implementation, with type-safe filtering through separate typed fields.

The architecture provides clear implementation milestones while enabling developers to choose their level of CCL support based on actual needs rather than artificial constraints.

## Architecture Overview

```
Core Functionality     ← Essential for any CCL implementation
├── Essential Parsing     (18 tests) - Start here
├── Comprehensive Parsing (30 tests) - Production ready  
└── Object Construction   (8 tests) - Hierarchical access

Optional Features      ← Choose based on your needs
├── Dotted Keys          (18 tests) - Dual access patterns
├── Comments             (3 tests) - Documentation support
├── Processing           (21 tests) - Advanced composition  
└── Typed Access         (17 tests) - Type-safe APIs

Integration           ← Validation & edge cases
├── Pretty Printing      (15 tests) - Canonical formatting
└── Error Handling       (5 tests) - Robust error reporting
```

Each category has specific APIs, test suites, and implementation requirements.

## Dual-Format Architecture

### Source Format (Maintainable)
Files: `tests/api_*.json` - Grouped validations for easy test authoring

```json
{
  "name": "basic_object_construction", 
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
      "cases": [{
        "args": ["database.host"],
        "expected": "localhost"
      }]
    }
  },
  "meta": {
    "tags": ["function:parse", "function:make-objects", "function:get-string"],
    "level": 4,
    "feature": "dotted-keys"
  }
}
```

### Generated Flat Format (Implementation-Friendly)
Files: `generated_tests/` - One test per validation with typed fields

```json
// Generated from above source test
{
  "name": "basic_object_construction_parse",
  "input": "database.host = localhost",
  "validation": "parse",
  "expected": {
    "count": 1,
    "entries": [{"key": "database.host", "value": "localhost"}]
  },
  "functions": ["parse"],
  "features": ["dotted-keys"],
  "behaviors": [],
  "variants": [],
  "level": 4,
  "source_test": "basic_object_construction"
}
```

### Property Tests (Mathematical Properties)  
Files: `tests/property_*.json` - Custom test runner logic required

**Source Format:**
```json
{
  "name": "round_trip_basic",
  "input": "key = value\nnested =\n  sub = val", 
  "validations": {
    "round_trip": {
      "property": "identity"
    }
  },
  "meta": {
    "tags": ["function:parse", "function:pretty-print"],
    "level": 5,
    "feature": "pretty-printing"
  }
}
```

**Generated Flat Format:**
```json
{
  "name": "round_trip_basic_round_trip",
  "input": "key = value\nnested =\n  sub = val",
  "validation": "round_trip",
  "expected": {"property": "identity"},
  "functions": ["parse", "pretty-print"],
  "features": [],
  "behaviors": [],
  "variants": [],
  "level": 5,
  "source_test": "round_trip_basic"
}
```

### API Validation Types (Direct Mapping)
- **`parse`** - Level 1: Entry parsing validation
- **`filter`** - Level 2: Comment filtering validation  
- **`compose`** - Level 2: Entry composition validation
- **`expand_dotted`** - Level 2: Dotted key expansion validation
- **`make_objects`** - Level 3: Object construction validation
- **`get_string`**, **`get_int`**, **`get_bool`**, **`get_float`** - Level 4: Typed access validation

### Property Validation Types (Custom Logic)
- **`round_trip`** - Parse-format-parse identity validation
- **`associativity`** - Algebraic composition properties
- **`canonical_format`** - Canonical formatting validation

## Type-Safe Test Filtering

The generated flat format enables efficient, type-safe test selection:

### Function-Based Filtering
```javascript
// Basic implementation - core functions only
const implementedFunctions = ["parse", "make-objects", "get-string"];
const supportedTests = flatTests.filter(test => 
  test.functions.every(fn => implementedFunctions.includes(fn))
);

// Advanced implementation - includes processing
const implementedFunctions = [
  "parse", "make-objects", "get-string", "get-int", "get-bool",
  "filter", "compose", "expand-dotted", "pretty-print"
];
```

### Feature and Behavior Filtering
```javascript
// Implementation capabilities
const capabilities = {
  functions: ["parse", "make-objects", "get-string"],
  features: ["comments", "dotted-keys"],
  behaviors: ["crlf-normalize-to-lf", "boolean-strict"],
  variants: ["reference-compliant"]
};

// Type-safe filtering with conflict resolution
const compatibleTests = flatTests.filter(test => {
  const functionsSupported = test.functions.every(fn => 
    capabilities.functions.includes(fn)
  );
  const featuresSupported = test.features.every(feature => 
    capabilities.features.includes(feature)
  );
  const hasConflicts = test.conflicts?.behaviors?.some(b => 
    capabilities.behaviors.includes(b)
  ) || test.conflicts?.variants?.some(v => 
    capabilities.variants.includes(v)
  );
  
  return functionsSupported && featuresSupported && !hasConflicts;
});
```

### Benefits
- **Type-safe**: Direct field access vs string parsing
- **Performance**: Array methods faster than tag matching
- **API ergonomics**: Intuitive filtering patterns
- **Schema validation**: Enum constraints ensure consistency

### Implementation Benefits

**Source Format:**
✅ **Maintainable** - Grouped validations reduce input duplication  
✅ **Readable** - Clear structure for test authoring  
✅ **Extensible** - Easy to add validation types per test

**Generated Flat Format:**
✅ **Type-safe filtering** - Direct field access with `test.functions[]`, `test.features[]`  
✅ **Simple test runners** - One validation per test (no complex iteration)  
✅ **Excellent performance** - Array methods faster than string parsing  
✅ **API ergonomics** - Intuitive filtering patterns

**Architecture:**
✅ **Progressive adoption** - Start with core functions, add features incrementally  
✅ **Clear separation** - API tests vs property tests clearly distinguished  
⚠️ **Property Tests** - Requires custom mathematical property implementations

## Core Functionality (Required)

### Essential Parsing
**Files**: `tests/api_essential-parsing.json` (18 tests)  
**API:** `parse(text) → Result<Entry[], ParseError>`  
**Status:** Required for all CCL implementations

#### Functionality
- Basic key-value parsing with `=` delimiter
- Whitespace handling and normalization  
- Multiline values through indented continuation lines
- Unicode support and line ending normalization
- Empty keys/values and equals-in-values handling
- Core error detection and reporting

#### Test Coverage
- **Focus areas:** Basic parsing, whitespace, multiline values, unicode
- **Essential tests:** 18 core functionality tests for rapid prototyping
- Covers 80% of real-world CCL usage scenarios

#### Example Implementation
```pseudocode
function parse(text) {
  entries = []
  lines = split_lines(text)
  
  for each line {
    if contains("=") {
      key = extract_key(line)
      value = extract_value(line)
      
      // Handle continuation lines
      while next_line_indented() {
        value += "\n" + continuation_content()
      }
      
      entries.append(Entry(key, value))
    }
  }
  
  return entries
}
```

### Comprehensive Parsing  
**Files**: `tests/api_comprehensive-parsing.json` (30 tests)
**Purpose:** Production-ready validation with comprehensive edge cases
**Status:** Recommended for production systems

#### Functionality
- Whitespace variations (tabs, spaces, mixed indentation)
- Line ending handling (Unix, Windows, Mac)
- Edge cases (empty keys/values, multiple equals, quotes)
- Stress testing with realistic configuration examples
- Unicode edge cases and normalization

#### Test Coverage
- **Edge cases:** Comprehensive whitespace and formatting variations
- **Production scenarios:** Complex real-world configuration patterns
- **Robustness:** Handles malformed input gracefully

### Object Construction
**Files**: `tests/api_object-construction.json` (8 tests)  
**API:** `make_objects(entries) → CCL`  
**Status:** Required for hierarchical access

#### Functionality
- Recursive parsing of nested values using fixed-point algorithm
- Duplicate key merging in object construction
- Empty key handling for list-style data  
- Complex nested configuration support

#### Fixed-Point Algorithm
```pseudocode
function make_objects(entries) {
  objects = {}
  
  for each entry in entries {
    if entry.value contains CCL syntax {
      // Recursively parse nested content
      nested_entries = parse(entry.value)
      objects[entry.key] = make_objects(nested_entries)
    } else {
      objects[entry.key] = entry.value
    }
  }
  
  return merge_duplicate_keys(objects)
}
```

## Optional Features (Choose Based on Needs)

### Dotted Key Expansion
**Files**: `tests/api_dotted-keys.json` (18 tests)  
**Purpose:** Enable dual access patterns for convenience
**Status:** Recommended for user-friendly APIs

#### Functionality
- Expand `database.host = localhost` to nested structures
- Support deep nesting (3+ levels)
- Handle mixed dotted and nested syntax
- Resolve conflicts between access patterns

#### Benefits
- Users can write `database.host = localhost` 
- APIs support both `get(obj, "database.host")` and `get(obj, "database", "host")`
- Flexible configuration authoring

### Comment Filtering
**Files**: `tests/api_comments.json` (3 tests)  
**API:** `filter(entries)`  
**Purpose:** Remove documentation keys from configuration

#### Functionality  
- Filter keys starting with `/` (comment syntax)
- Configurable comment prefixes
- Preserve structure while removing documentation

### Entry Processing
**Files**: `tests/api_processing.json` (21 tests)
**API:** `compose()`, advanced processing
**Purpose:** Advanced composition and merging capabilities

#### Functionality
- Duplicate key handling and composition  
- Entry list merging with algebraic properties
- Complex composition scenarios
- Associative and commutative operations

### Typed Access
**Files**: `tests/api_typed-access.json` (17 tests)  
**API:** `get_string()`, `get_int()`, `get_bool()`, etc.  
**Purpose:** Type-safe value extraction with validation

#### Functionality
- Smart type inference (integers, floats, booleans)
- Configurable parsing options and validation
- Language-specific convenience functions  
- Dual access pattern support (dotted + hierarchical)
- Type safety and error handling

#### Example Implementation
```pseudocode
function get_int(ccl_obj, ...path) {
  value = get_raw_value(ccl_obj, ...path)
  
  if value matches integer_pattern {
    return parse_int(value)
  } else {
    return TypeError("Expected integer at " + path)
  }
}

function get_bool(ccl_obj, ...path) {
  value = get_raw_value(ccl_obj, ...path).toLowerCase()
  
  if value in ["true", "yes", "on", "1"] {
    return true
  } else if value in ["false", "no", "off", "0"] {
    return false  
  } else {
    return TypeError("Expected boolean at " + path)
  }
}
```

## Integration & Validation

### Error Handling
**Files**: `tests/api_errors.json` (5 tests)  
**Purpose:** Malformed input detection across all functionality
**Status:** Recommended for robust implementations

#### Error Categories
1. **Parse Errors:** Malformed CCL syntax
2. **Type Errors:** Invalid type conversion  
3. **Path Errors:** Nonexistent configuration keys
4. **Validation Errors:** Failed constraint checking

## Implementation Strategy

### Progressive Implementation
1. **Start with Core**: Essential parsing (18 tests) → Object construction (8 tests)
2. **Add Production Readiness**: Comprehensive parsing (30 tests)
3. **Choose Features**: Select from features/ based on your use case
4. **Validate**: Error handling for robustness

### Implementation Priorities

#### Minimal CCL (Essential - 26 tests)
- `core/essential-parsing.json` (18 tests)
- `core/object-construction.json` (8 tests)
- Basic CCL support for simple configurations

#### Standard CCL (Recommended - 56 tests)  
- All Core functionality (56 tests)
- `features/dotted-keys.json` (18 tests)
- `integration/errors.json` (5 tests)
- Production-ready with convenient access patterns

#### Full CCL (Complete - 135 tests)
- All tests across all categories  
- Maximum feature support and robustness
- Includes pretty-printing and round-trip validation

### Test-Driven Development
```bash
# API Tests (Direct function mapping)
npm run validate:api_essential-parsing       # Essential tests (18)
npm run validate:api_object-construction     # Object construction (8)
npm run validate:api_comprehensive-parsing   # Comprehensive tests (30)
npm run validate:api_dotted-keys            # Dotted key support (18)
npm run validate:api_typed-access           # Type-safe APIs (17)
npm run validate:api_comments               # Comment filtering (3)
npm run validate:api_processing             # Entry processing (21)
npm run validate:api_errors                 # Error handling (5)

# Property Tests (Custom logic required)  
npm run validate:property_round-trip        # Round-trip validation (15)
npm run validate:property_algebraic         # Algebraic properties (5)

# Full validation  
npm test                                     # All 135 tests
```

### API Design Patterns

**Consistent Error Handling:**
```pseudocode
Result<T, Error> pattern for all fallible operations
- Ok(value) for successful operations
- Error(message) for failures with descriptive messages
```

**Reusable Implementation:**
```pseudocode
// Core navigation logic shared across all getters
parse_path(...args) → segments
navigate_path(obj, segments) → value
get_raw_value(obj, ...path) → string

// All typed getters reuse the same navigation
get_int(obj, ...path) → int
get_bool(obj, ...path) → bool
get_string(obj, ...path) → string
```

## Test Runner Implementation

### Source Format Test Runner
**For files:** `tests/api_*.json` - Multi-validation test cases

```pseudocode
function run_source_test(test_case) {
  // Iterate over multiple validations per test
  for (validation_type, validation_spec) in test_case.validations {
    switch validation_type {
      case "parse":
        actual = parse(test_case.input)
        assert_equal(actual, validation_spec.expected)
        assert_equal(actual.length, validation_spec.count)
        
      case "make_objects":
        entries = parse(test_case.input)
        actual = make_objects(entries)
        assert_equal(actual, validation_spec.expected)
        
      case "get_string":
        entries = parse(test_case.input)
        ccl = make_objects(entries)
        for case in validation_spec.cases {
          actual = get_string(ccl, ...case.args)
          assert_equal(actual, case.expected)
        }
        assert_equal(validation_spec.cases.length, validation_spec.count)
        
      // ... other API validations
    }
  }
}
```

### Generated Flat Format Test Runner (Recommended)
**For files:** `generated_tests/` - Type-safe single-validation tests

```pseudocode
function run_flat_test(flat_test) {
  // Each flat test has exactly one validation
  switch flat_test.validation {
    case "parse":
      if flat_test.expect_error {
        assert_throws(() => parse(flat_test.input))
      } else {
        actual = parse(flat_test.input)
        assert_equal(actual, flat_test.expected.entries)
        assert_equal(actual.length, flat_test.expected.count)
      }
      
    case "make_objects":
      entries = parse(flat_test.input)
      actual = make_objects(entries)
      assert_equal(actual, flat_test.expected.object)
      
    case "get_string":
      entries = parse(flat_test.input)
      ccl = make_objects(entries)
      if flat_test.expect_error {
        assert_throws(() => get_string(ccl, ...flat_test.args))
      } else {
        actual = get_string(ccl, ...flat_test.args)
        assert_equal(actual, flat_test.expected.value)
      }
      
    // ... other validations
  }
}

// Usage with type-safe filtering
function run_compatible_tests(flat_tests, capabilities) {
  compatible_tests = flat_tests.filter(test => {
    // Type-safe filtering
    functions_supported = test.functions.every(fn => 
      capabilities.functions.includes(fn)
    )
    features_supported = test.features.every(feature => 
      capabilities.features.includes(feature)
    )
    no_conflicts = !test.conflicts?.behaviors?.some(b => 
      capabilities.behaviors.includes(b)
    )
    
    return functions_supported && features_supported && no_conflicts
  })
  
  compatible_tests.forEach(test => run_flat_test(test))
}
```

### Full Property Test Runner (100+ lines) 
**For files:** `property_*.json` - Requires custom mathematical property logic

```pseudocode
function run_property_validation_test(test_case) {
  // Handle both API validations + complex properties
  for (validation_type, expected) in test_case.validations {
    switch validation_type {
      // All the API validations from above, PLUS:
      
      case "round_trip":
        // Multi-step identity property
        entries = parse(test_case.input)
        formatted = pretty_print(entries)
        reparsed = parse(formatted)
        assert_equal(entries, reparsed)
        
      case "associativity":
        // Complex algebraic property testing
        if expected.property == "semigroup_associativity" {
          // Test (a ⊕ b) ⊕ c == a ⊕ (b ⊕ c)
          a = parse(test_case.input1)
          b = parse(test_case.input2) 
          c = parse(test_case.input3)
          
          left_assoc = compose(compose(a, b), c)
          right_assoc = compose(a, compose(b, c))
          assert_equal(left_assoc, right_assoc)
        }
        
      case "canonical_format":
        // Custom formatting validation logic
        entries = parse(test_case.input)
        canonical = canonical_format(entries)
        assert_matches_canonical_rules(canonical)
    }
  }
}
```

### Implementation Guidance

**Start Here (Simple):**
```bash
# Only run API tests - direct function mapping
validate_tests("tests/api_*.json")           # ~90 tests
```

**Advanced (Full Features):**
```bash  
# Add property tests - custom logic required
validate_tests("tests/property_*.json")      # ~45 tests  
validate_tests("tests/*.json")               # All 135 tests
```

**Implementation Burden:**
- ✅ **API Tests**: Straightforward switch statement over validation types
- ⚠️ **Property Tests**: Custom mathematical property implementations required
- 📊 **Lines of Code**: Simple runner ~30 lines, Full runner ~100+ lines

## Architecture Benefits

1. **Flexible Implementation**: Choose features based on actual needs
2. **Clear Progression**: Start simple, add features incrementally  
3. **Implementation Guidance**: Test counts show relative complexity
4. **Language Agnostic**: Architecture works across programming languages
5. **Comprehensive Coverage**: 135 tests cover all CCL functionality
6. **Maintainable**: Feature-based organization scales with additional features
7. **Clear Testing**: Test format provides explicit validation mapping
8. **Simple Test Runners**: Direct mapping from validations to API functions

## Implementation Examples

### Minimal Implementation (Essential + Objects)
```pseudocode
// Basic CCL parser with hierarchy
entries = parse(ccl_text)              // 18 essential parsing tests
objects = make_objects(entries)        // 8 object construction tests
value = get_string(objects, "database", "host")
```

### Standard Implementation (Core + Dotted Keys)
```pseudocode  
// Convenient CCL parser with dual access
entries = parse(ccl_text)              // All core tests (56)
objects = make_objects(entries)
host = get_string(objects, "database.host")      // Dotted access
port = get_int(objects, "database", "port")      // Hierarchical access
```

### Production Implementation (Full Features)
```pseudocode
// Robust configuration loading
try {
  entries = parse(read_file("config.ccl"))      // Comprehensive parsing
  filtered = filter(entries)                    // Comment filtering  
  config = make_objects(filtered)               // Object construction
  
  validate_required_keys(config)
  return load_typed_config(config)              // Type-safe access
} catch (error) {
  log_error("Configuration failed:", error)     // Error handling
  return default_config()
}
```

The feature-based architecture provides a systematic approach to building CCL implementations while maintaining simplicity and flexibility for different use cases and requirements.