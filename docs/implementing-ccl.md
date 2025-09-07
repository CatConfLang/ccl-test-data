# Implementing CCL - A Guide for Language Authors

This guide helps you implement a CCL parser in your programming language using the 4-level architecture and comprehensive test suite.

## Quick Start

1. **Study the specification** - Read the [CCL FAQ](ccl_faq.md) and [Getting Started Guide](getting-started.md)
2. **Choose your target level** - Start with Level 1, progress through Level 4
3. **Use the test suite** - Language-agnostic JSON tests validate your implementation
4. **Follow the reference** - OCaml reference implementation at https://github.com/chshersh/ccl

## Implementation Roadmap

### Phase 1: Core Parsing (Level 1)
**Goal:** Parse CCL text into flat key-value entries  
**Test Suite:** `tests/level-1-parsing.json` (48 tests)

#### Essential Algorithm
```pseudocode
function parse(text: string) -> Result<List<Entry>, ParseError> {
  entries = []
  lines = split_lines_with_positions(text)
  
  for line in lines {
    if line.contains("=") {
      (key, value) = split_on_first_equals(line)
      key = trim_key(key)
      value = extract_initial_value(value)
      
      // Handle multiline continuation
      while next_line_is_continuation(lines, current_index) {
        continuation = get_continuation_content(lines, current_index + 1)
        value += "\n" + continuation
        current_index += 1
      }
      
      entries.append(Entry(key, value))
    }
  }
  
  return Ok(entries)
}
```

#### Key Implementation Details

**Line Splitting:**
- Preserve line ending information (for error reporting)
- Handle Unix (`\n`), Windows (`\r\n`), and legacy Mac (`\r`) line endings
- Normalize to consistent internal representation

**Key Extraction:**
- Split on first `=` character only
- Trim whitespace from keys
- Empty keys are valid (used for lists)

**Value Extraction:**
- Preserve leading/trailing whitespace in values
- Handle values containing `=` characters
- Empty values are valid

**Continuation Lines:**
- Lines with indentation greater than parent continue the value
- Preserve relative indentation in multiline values
- Handle mixed tabs and spaces (warn in strict mode)

### Phase 2: Object Construction (Level 3)
**Goal:** Convert flat entries to nested objects  
**Test Suite:** `tests/level-3-objects.json` (8 tests)

#### Fixed-Point Algorithm
```pseudocode
function make_objects(entries: List<Entry>) -> CCL {
  result = {}
  
  for entry in entries {
    if entry.value.contains_ccl_syntax() {
      // Recursively parse nested content
      nested_entries = parse(entry.value)
      nested_object = make_objects(nested_entries)
      result = merge_into_result(result, entry.key, nested_object)
    } else {
      result = merge_into_result(result, entry.key, entry.value)
    }
  }
  
  return result
}

function contains_ccl_syntax(value: string) -> boolean {
  // Check if value looks like CCL (contains "=")
  lines = split_lines(value)
  for line in lines {
    if trim(line).contains("=") {
      return true
    }
  }
  return false
}
```

#### Duplicate Key Handling
```pseudocode
function merge_into_result(result: CCL, key: string, value: any) {
  if key == "" {
    // Empty keys create lists
    if result[""] exists {
      result[""].append(value)
    } else {
      result[""] = [value]
    }
  } else if result[key] exists {
    // Merge duplicate keys
    result[key] = deep_merge(result[key], value)
  } else {
    result[key] = value
  }
}
```

### Phase 3: Entry Processing (Level 2)  
**Goal:** Filter and compose entries  
**Test Suite:** `tests/level-2-processing.json` (28 tests)

#### Comment Filtering
```pseudocode
function filter_comments(entries: List<Entry>) -> List<Entry> {
  return entries.filter(entry -> !entry.key.starts_with("/"))
}

// Alternative: filter by custom comment prefixes
function filter_by_prefixes(entries: List<Entry>, prefixes: List<string>) {
  return entries.filter(entry -> !any(prefixes, prefix -> entry.key.starts_with(prefix)))
}
```

#### Entry Composition
```pseudocode
function compose_entries(left: List<Entry>, right: List<Entry>) -> List<Entry> {
  // Simple concatenation - merging happens at object level
  return left + right
}
```

### Phase 4: Typed Access (Level 4)
**Goal:** Type-safe value extraction  
**Test Suite:** `tests/level-4-typed.json` (17 tests)

#### Type-Safe Accessors
```pseudocode
function get_string(ccl: CCL, path: string) -> Result<string, Error> {
  value = navigate_path(ccl, path)
  
  match value {
    string -> Ok(string)
    _ -> Error("Expected string at " + path)
  }
}

function get_int(ccl: CCL, path: string) -> Result<int, Error> {
  str_result = get_string(ccl, path)
  match str_result {
    Ok(str) -> parse_int(str)
    Error(e) -> Error(e)
  }
}

function get_bool(ccl: CCL, path: string) -> Result<bool, Error> {
  str_result = get_string(ccl, path)
  match str_result {
    Ok(str) -> match str.to_lowercase() {
      "true" | "yes" | "on" | "1" -> Ok(true)
      "false" | "no" | "off" | "0" -> Ok(false)
      _ -> Error("Invalid boolean: " + str)
    }
    Error(e) -> Error(e)
  }
}
```

## Language-Specific Considerations

### Error Handling Patterns

**Rust:**
```rust
type ParseResult<T> = Result<T, ParseError>;
type AccessResult<T> = Result<T, AccessError>;
```

**Go:**
```go
func Parse(text string) ([]Entry, error)
func GetString(ccl CCL, path string) (string, error)
```

**Python:**
```python
def parse(text: str) -> Union[List[Entry], ParseError]
def get_string(ccl: CCL, path: str) -> Union[str, AccessError]
```

**JavaScript:**
```javascript
function parse(text) { /* returns entries or throws ParseError */ }
function getString(ccl, path) { /* returns string or throws AccessError */ }
```

### Data Structure Design

**Immutable Languages (Haskell, OCaml):**
```haskell
data Entry = Entry String String
data CCL = CCLString String | CCLList [String] | CCLObject (Map String CCL)
```

**Object-Oriented Languages (Java, C#):**
```java
class Entry {
  public final String key;
  public final String value;
}

interface CCLValue {}
class CCLString implements CCLValue { String value; }
class CCLList implements CCLValue { List<String> items; }
class CCLObject implements CCLValue { Map<String, CCLValue> fields; }
```

## Testing Your Implementation

### Running the Test Suite

Each level has a dedicated test file with specific format:

```json
{
  "tests": [
    {
      "name": "basic_key_value",
      "input": "key = value",
      "expected": [
        {"key": "key", "value": "value"}
      ],
      "meta": {
        "level": 1,
        "tags": ["basic"]
      }
    }
  ]
}
```

### Test Runner Implementation
```pseudocode
function run_test_suite(test_file: string) {
  test_data = load_json(test_file)
  
  for test in test_data.tests {
    try {
      // Level 1 tests
      if test.expected exists {
        actual = parse(test.input)
        assert_equal(actual, test.expected)
      }
      
      // Level 3 tests  
      if test.expected_nested exists {
        entries = parse(test.input)
        objects = make_objects(entries)
        assert_equal(objects, test.expected_nested)
      }
      
      // Error tests
      if test.expected_error exists {
        result = parse(test.input)
        assert_error(result)
      }
      
      print("✅ " + test.name)
    } catch (error) {
      print("❌ " + test.name + ": " + error)
    }
  }
}
```

### Progressive Testing Strategy

1. **Start with essential tests:**
   ```bash
   # Filter non-redundant Level 1 tests
   jq '.tests[] | select(.meta.tags | contains(["redundant"]) | not)' level-1-parsing.json
   ```

2. **Add comprehensive coverage:**
   ```bash
   # Run all tests for your target level
   run_tests("level-1-parsing.json")   # 48 tests
   run_tests("level-3-objects.json")   # 8 tests
   ```

3. **Validate error handling:**
   ```bash
   run_tests("errors.json")            # 5 tests
   ```

## Performance Considerations

### Parsing Performance
- **Line-by-line processing** is typically fastest for CCL
- **Minimize string allocations** during key/value extraction
- **Lazy evaluation** for nested object construction
- **Streaming parsers** for very large configuration files

### Memory Usage
- **Share string data** between entries when possible
- **Use rope/gap buffer structures** for large multiline values
- **Implement copy-on-write** for object merging operations

### Optimization Strategies
```pseudocode
// Fast path for flat configurations (no nesting)
if !text.contains_indented_lines() {
  return parse_flat_only(text)  // Skip object construction
}

// Lazy object construction
class LazyObject {
  entries: List<Entry>
  constructed: Option<CCL>
  
  function get(key: string) {
    if !constructed {
      constructed = Some(make_objects(entries))
    }
    return constructed.get(key)
  }
}
```

## Common Implementation Pitfalls

1. **Incorrect continuation handling** - Ensure indentation comparison is exact
2. **Wrong equals splitting** - Only split on first `=`, preserve others in value
3. **Unicode issues** - Handle UTF-8 properly, including multi-byte characters
4. **Path navigation errors** - Handle nested access edge cases
5. **Type conversion edge cases** - Boolean parsing, integer overflow
6. **Memory leaks** - In languages without GC, manage string lifetimes carefully

## API Design Guidelines

### Consistent Naming
- `parse()` for Level 1 entry parsing
- `make_objects()` for Level 3 object construction  
- `filter_comments()` for Level 2 comment filtering
- `get_string()`, `get_int()`, `get_bool()` for Level 4 typed access

### Error Messages
Provide helpful error messages with:
- **Line/column numbers** for parse errors
- **Key paths** for access errors  
- **Expected vs actual** for type errors
- **Suggestions** for common mistakes

### Documentation Requirements
- **API documentation** with examples
- **Migration guide** from popular formats
- **Performance characteristics** and limitations
- **Thread safety** guarantees (if applicable)

## Validation and Release

### Pre-Release Checklist
- [ ] All test suites pass (106+ tests total)
- [ ] Error messages are helpful and consistent
- [ ] Performance is acceptable for target use cases
- [ ] Documentation is complete
- [ ] Thread safety is documented/implemented
- [ ] Memory usage is reasonable

### Publishing Guidelines
- **Semantic versioning** with level support clearly indicated
- **Clear API stability** promises
- **Examples and tutorials** for common use cases
- **Contribution guidelines** for community involvement

The CCL specification and test suite provide everything needed to build a robust, compliant implementation in any programming language. Focus on correctness first, then optimize for your language's specific performance characteristics.