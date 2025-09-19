# Implementing CCL - A Guide for Language Authors

Implementation guide for CCL parsers using the comprehensive test suite.

## Quick Start

1. **Core first**: Essential parsing + Object construction (26 tests)
2. **Add features**: Typed access, processing, formatting as needed
3. **Test-driven**: Each validation maps to specific API functions
4. **Reference**: OCaml implementation at https://github.com/chshersh/ccl

## Core Functions (Required)

### Essential Parsing (18 tests)
**API**: `parse(text) → Entry[]`
- Split on first `=`, handle multiline values via indentation
- Preserve whitespace in values, trim keys
- Support Unix/Windows/Mac line endings

### Object Construction (8 tests)
**API**: `build_hierarchy(entries) → CCL`
- Fixed-point algorithm: recursively parse nested values
- Merge duplicate keys, handle empty keys as lists
- Creates hierarchical objects from flat entries

### Typed Access (17 tests)
**Functions**: `get_string()`, `get_int()`, `get_bool()`, `get_float()`, `get_list()`
**Dual access**: Support both `get_string(ccl, "database", "host")` and `get_string(ccl, "database.host")`

### Processing Functions (21 tests)
**Functions**: `filter()`, `compose()`, `expand_dotted()`
- **Filter**: Remove comment keys (starting with `/`)
- **Compose**: Concatenate entry lists, merge at object level
- **Expand dotted**: Convert `database.host` to nested structures

### Implementation Pattern

**Reusable Navigation**:
```pseudocode
// Unified path parsing for both access patterns
function parse_path(...args) -> string[] {
  return args.length == 1 && args[0].contains(".") ?
    args[0].split(".") : args
}

// Shared navigation logic
function get_raw_value(ccl, ...path) {
  segments = parse_path(...path)
  return navigate_path(ccl, segments)
}

// All getters reuse navigation
function get_int(ccl, ...path) {
  value = get_raw_value(ccl, ...path)
  return parse_int(value)
}
```

## Language-Specific Patterns

**Error Handling**:
- **Rust**: `Result<T, Error>`
- **Go**: `(T, error)`
- **Python**: `Union[T, Error]`
- **JavaScript**: Throw exceptions

**Data Structures**:
- **Immutable**: `Entry(key, value)`, `CCL` with string/list/object variants
- **OOP**: `Entry` class, `CCLValue` interface with implementations

## Testing Strategy

### Test Runner Pattern
```pseudocode
for test in test_data {
  for (validation_type, expected) in test.validations {
    switch validation_type {
      case "parse": actual = parse(test.input); assert_equal(actual, expected)
      case "get_string": actual = get_string(ccl, ...args); assert_equal(actual, expected)
      // ... other validations
    }
  }
}
```

### Progressive Implementation
1. **Essential** (18 tests): Basic parsing
2. **Core** (26 tests): + Object construction
3. **Standard** (56 tests): + Comprehensive parsing
4. **Full** (135 tests): + All features

## Performance Tips

- **Line-by-line processing** for parsing
- **Minimize allocations** during key/value extraction
- **Lazy object construction** for large configs
- **Fast path** for flat configurations (no indentation)
- **String sharing** between entries

## Common Pitfalls

1. **Continuation handling**: Exact indentation comparison required
2. **Equals splitting**: Only split on first `=`
3. **Unicode**: Proper UTF-8 handling
4. **Path navigation**: Handle missing keys gracefully
5. **Type conversion**: Boolean parsing edge cases
6. **Memory management**: String lifetime in non-GC languages

## API Guidelines

### Function Naming
- `parse()`, `build_hierarchy()`, `filter_comments()`
- `get_string()`, `get_int()`, `get_bool()`

### Error Messages
- **Parse errors**: Line/column numbers
- **Access errors**: Key paths
- **Type errors**: Expected vs actual
- **Test validation**: Flexible regex patterns preferred over exact strings

### Documentation
- API examples, migration guides
- Performance characteristics, thread safety

## Release Checklist

- [ ] All test suites pass (135+ tests)
- [ ] Error messages are helpful
- [ ] Performance acceptable
- [ ] Documentation complete
- [ ] Thread safety documented

### Publishing
- Semantic versioning with feature support indicated
- API stability promises
- Examples and contribution guidelines