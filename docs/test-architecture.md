# CCL Test Suite Architecture

Dual-format test architecture with progressive implementation paths and type-safe filtering.

## Implementation Paths

**Core**: Essential → Comprehensive → Object Construction (56 tests)
**Features**: Comments → Dotted Keys → Processing → Typed Access (59 tests)
**Integration**: Error Handling → Pretty Printing (20 tests)

## Test File Overview

| File | Tests | Assertions | Functions | Features | Purpose |
|------|-------|------------|-----------|----------|---------|
| **api_essential-parsing.json** | 45 | 89 | parse, load | - | Foundation parsing |
| **api_comprehensive-parsing.json** | 23 | 47 | parse, load | unicode, whitespace | Advanced parsing |
| **api_processing.json** | 67 | 134 | filter, compose, expand-dotted | comments | Entry processing |
| **api_comments.json** | 15 | 30 | parse, filter | comments | Comment handling |
| **api_object-construction.json** | 89 | 178 | build-hierarchy, build-hierarchy | dotted-keys | Object building |
| **api_dotted-keys.json** | 34 | 68 | build-hierarchy, expand-dotted | dotted-keys | Key expansion |
| **api_typed-access.json** | 151 | 302 | get-string, get-int, get-bool, get-float | - | Type access |
| **api_errors.json** | 18 | 36 | All functions | error-handling | Error validation |
| **property_round-trip.json** | 12 | 24 | parse, pretty-print | round-trip | Consistency |
| **property_algebraic.json** | 8 | 16 | All functions | algebraic | Properties |

**Total**: 462 tests, 924 assertions across 10 files

## Dual-Format Architecture

### Source Format (Maintainable)
Multiple validations per test, easy authoring:
```json
{
  "validations": {
    "parse": {"count": 1, "expected": [...]},
    "make_objects": {"count": 1, "expected": {...}}
  },
  "meta": {"tags": ["function:parse", "function:build-hierarchy"]}
}
```

### Generated Format (Type-Safe)
One validation per test, direct filtering:
```json
{
  "validation": "parse",
  "functions": ["parse"],
  "features": ["dotted-keys"]
}
```

### Validation Types
**API Functions**: `parse`, `filter`, `compose`, `expand_dotted`, `make_objects`, `get_*`
**Properties**: `round_trip`, `associativity`, `canonical_format` (requires custom logic)

## Test Filtering

```javascript
// Filter by capabilities
const supportedTests = flatTests.filter(test =>
  test.functions.every(fn => capabilities.functions.includes(fn)) &&
  test.features.every(f => capabilities.features.includes(f)) &&
  !test.conflicts?.behaviors?.some(b => capabilities.behaviors.includes(b))
);
```

### Benefits
✅ **Source**: Maintainable, readable, extensible
✅ **Generated**: Type-safe filtering, simple runners, performance
✅ **Architecture**: Progressive adoption, clear separation
⚠️ **Property Tests**: Custom logic required

## Core Functions

### Essential Parsing (18 tests)
**API**: `parse(text) → Entry[]`
- Key-value parsing with `=` delimiter
- Whitespace/unicode handling, multiline values
- Empty keys/values, basic error detection

### Comprehensive Parsing (30 tests)
**Purpose**: Production edge cases
- Complex whitespace variations, line endings
- Realistic configuration patterns

### Object Construction (8 tests)
**API**: `make_objects(entries) → CCL`
- Recursive parsing via fixed-point algorithm
- Duplicate key merging, nested structures

## Optional Features

### Dotted Keys (18 tests)
**Purpose**: Dual access patterns
- Expand `database.host = localhost` to nested structures
- Support both `get(obj, "database.host")` and `get(obj, "database", "host")`

### Comments (3 tests)
**API**: `filter(entries)`
- Remove `/` prefixed documentation keys

### Processing (21 tests)
**API**: `compose()`, advanced operations
- Entry composition, merging, algebraic properties

### Typed Access (17 tests)
**API**: `get_string()`, `get_int()`, `get_bool()`, etc.
- Type-safe extraction with smart inference
- Dual access pattern support

## Integration

### Error Handling (5 tests)
**Categories**: Parse errors, type errors, path errors, validation errors

## Implementation Strategy

### Progressive Paths
1. **Minimal** (26 tests): Essential parsing + Object construction
2. **Standard** (56 tests): Core + Dotted keys + Error handling
3. **Full** (135 tests): All features + Properties

### API Patterns
- **Error handling**: `Result<T, Error>` pattern
- **Navigation**: Shared path logic across all getters
- **Reusability**: Core functions composed into higher-level APIs

## Test Runners

### Source Format Runner
Iterates over multiple validations per test case, requires complex switch logic.

### Generated Format Runner (Recommended)
```pseudocode
function run_flat_test(test) {
  switch test.validation {
    case "parse": actual = parse(test.input); assert_equal(actual, test.expected)
    case "get_string": actual = get_string(ccl, ...test.args); assert_equal(actual, test.expected)
    // ... other validations
  }
}
```

### Implementation Complexity
- ✅ **API Tests**: ~30 lines, straightforward mapping
- ⚠️ **Property Tests**: ~100+ lines, custom logic required

## Architecture Benefits

- **Flexible**: Choose features based on actual needs
- **Progressive**: Start simple, add incrementally
- **Language agnostic**: Works across programming languages
- **Comprehensive**: 135 tests cover all CCL functionality
- **Clear testing**: Direct validation-to-API mapping

## Implementation Examples

**Minimal** (26 tests):
```
entries = parse(text); objects = make_objects(entries)
```

**Standard** (56 tests):
```
host = get_string(objects, "database.host")  // Dotted access
port = get_int(objects, "database", "port")   // Hierarchical access
```

**Production** (135 tests):
```
entries = parse(file) → filter() → make_objects() → typed_access()
```