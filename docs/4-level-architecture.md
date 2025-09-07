# CCL 4-Level Architecture

CCL implementations are designed as a layered architecture where each level builds on the previous one. This approach provides clear implementation milestones and allows developers to choose their level of CCL support.

## Architecture Overview

```
Level 4: Typed Parsing      ← get_int(), get_bool(), smart types
Level 3: Object Construction ← make_objects(), nested structure  
Level 2: Entry Processing    ← filter_keys(), comments, composition
Level 1: Entry Parsing       ← parse(), core key-value extraction
```

Each level has specific APIs, test suites, and implementation requirements.

## Level 1: Entry Parsing (Core)

**Purpose:** Convert raw CCL text to flat key-value entries  
**API:** `parse(text) → Result(List(Entry), ParseError)`  
**Status:** Required for all CCL implementations

### Functionality
- Basic key-value parsing with `=` delimiter
- Whitespace handling and normalization  
- Multiline values through indented continuation lines
- Unicode support and line ending normalization
- Empty keys/values and equals-in-values handling
- Core error detection and reporting

### Test Coverage
- **File:** `tests/level-1-parsing.json` (48 tests)
- **Essential tests:** 18 core functionality tests
- **Redundant tests:** 30 additional edge case variations
- **Focus areas:** Basic parsing, whitespace, multiline values, unicode

### Example Implementation
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

## Level 2: Entry Processing (Extensions)

**Purpose:** Process Entry[] to filtered/grouped Entry[]  
**API:** `filter_keys()`, composition functions  
**Status:** Optional but recommended

### Functionality
- Comment filtering (keys starting with `/`)
- Duplicate key handling and composition  
- Entry list merging with algebraic properties
- Decorative section parsing (future feature)

### Test Coverage
- **File:** `tests/level-2-processing.json` (28 tests)
- **Comment filtering:** Remove documentation keys
- **Composition tests:** Duplicate key merging behavior
- **Algebraic properties:** Associativity and order independence

### Example Implementation
```pseudocode
function filter_comments(entries) {
  return entries.filter(entry => !entry.key.starts_with("/"))
}

function compose_entries(entries1, entries2) {
  // Merge with duplicate key handling
  return merge_duplicate_keys(entries1 + entries2)
}
```

## Level 3: Object Construction (Hierarchical)

**Purpose:** Convert flat Entry[] to nested object structures  
**API:** `make_objects(entries) → CCL`  
**Status:** Required for hierarchical access

### Functionality
- Recursive parsing of nested values using fixed-point algorithm
- Duplicate key merging in object construction
- Empty key handling for list-style data  
- Complex nested configuration support

### Test Coverage
- **File:** `tests/level-3-objects.json` (8 tests)
- **Recursive parsing:** Nested value construction
- **Duplicate merging:** Object-level key combination
- **List handling:** Empty keys become arrays
- **Complex examples:** Real-world configuration patterns

### Fixed-Point Algorithm
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

## Level 4: Typed Parsing (Language-Specific)

**Purpose:** Type-aware extraction with validation and inference  
**API:** `get_int()`, `get_bool()`, `get_typed_value()`, etc.  
**Status:** Language-specific conveniences

### Functionality
- Smart type inference (integers, floats, booleans)
- Configurable parsing options and validation
- Language-specific convenience functions  
- Type safety and error handling

### Test Coverage
- **File:** `tests/level-4-typed.json` (17 tests)
- **Type inference:** Automatic string-to-type conversion
- **Boolean parsing:** Multiple boolean representations
- **Error handling:** Type validation and fallbacks
- **Language patterns:** Idiomatic type access

### Example Implementation
```pseudocode
function get_int(ccl_obj, path) {
  value = get_string(ccl_obj, path)
  
  if value matches integer_pattern {
    return parse_int(value)
  } else {
    return TypeError("Expected integer at " + path)
  }
}

function get_bool(ccl_obj, path) {
  value = get_string(ccl_obj, path).toLowerCase()
  
  if value in ["true", "yes", "on", "1"] {
    return true
  } else if value in ["false", "no", "off", "0"] {
    return false  
  } else {
    return TypeError("Expected boolean at " + path)
  }
}
```

## Error Handling (Cross-Level)

**Purpose:** Malformed input detection across all levels  
**Test Coverage:** `tests/errors.json` (5 tests)

### Error Categories
1. **Parse Errors:** Malformed CCL syntax
2. **Type Errors:** Invalid type conversion  
3. **Path Errors:** Nonexistent configuration keys
4. **Validation Errors:** Failed constraint checking

## Implementation Strategy

### Progressive Implementation
1. **Start with Level 1:** Implement core parsing first
2. **Add Level 3:** Enable hierarchical access
3. **Include Level 2:** Add comment filtering
4. **Complete Level 4:** Provide type-safe APIs

### Test-Driven Development
```bash
# Run level-specific tests during development
run_tests("level-1-parsing.json")     # Core functionality
run_tests("level-3-objects.json")     # Object construction  
run_tests("level-2-processing.json")  # Comment processing
run_tests("level-4-typed.json")       # Type safety
```

### API Design Patterns

**Consistent Error Handling:**
```pseudocode
Result<T, Error> pattern for all fallible operations
- Ok(value) for successful operations
- Error(message) for failures with descriptive messages
```

**Composable Functions:**
```pseudocode
parse(text) |> filter_comments |> make_objects |> get_typed_value
```

## Architecture Benefits

1. **Clear Milestones:** Each level provides concrete implementation goals
2. **Incremental Development:** Build and test one level at a time  
3. **Flexible Support:** Implementations can choose their level of support
4. **Language Agnostic:** Architecture works across programming languages
5. **Test Coverage:** Comprehensive test suite for each level
6. **Maintenance:** Isolated concerns make debugging easier

## Implementation Examples

### Minimal Implementation (Level 1 + 3)
```pseudocode
// Basic CCL parser with object support
entries = parse(ccl_text)
objects = make_objects(entries)  
value = get_string(objects, "database.host")
```

### Full Implementation (All Levels)
```pseudocode  
// Complete CCL parser with type safety
entries = parse(ccl_text)
filtered = filter_comments(entries)
objects = make_objects(filtered)
port = get_int(objects, "server.port")  // Type-safe access
```

### Production Implementation
```pseudocode
// Robust configuration loading
try {
  entries = parse(read_file("config.ccl"))
  config = make_objects(filter_comments(entries))
  
  validate_required_keys(config)
  return load_typed_config(config)
} catch (error) {
  log_error("Configuration failed:", error)
  return default_config()
}
```

The 4-level architecture provides a systematic approach to building robust CCL implementations while maintaining simplicity and flexibility for different use cases.