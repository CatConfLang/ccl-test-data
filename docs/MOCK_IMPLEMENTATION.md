# CCL Mock Implementation Guide

This document provides comprehensive documentation of the CCL mock implementation patterns, architectural decisions, and implementation strategies used in the test suite's reference implementation.

## Overview

The mock implementation in `internal/mock/ccl.go` serves as both a reference implementation and a progressive development tool for CCL (Categorical Configuration Language). It demonstrates how to build a complete CCL parser across multiple implementation levels while maintaining compatibility with the structured test suite.

## Implementation Philosophy

### Progressive Implementation Strategy

The mock follows a **level-by-level implementation approach** that allows developers to:

1. **Start Simple**: Implement basic parsing (Level 1) first
2. **Add Complexity Gradually**: Build additional features incrementally
3. **Validate Continuously**: Use the test suite to verify each level
4. **Maintain Compatibility**: Ensure backward compatibility across levels

### Design Principles

- **Minimal Dependencies**: Uses only Go standard library
- **Clear Error Messages**: Provides detailed context for debugging
- **Extensible Architecture**: Easy to add new features and validations
- **Test-Driven**: Designed to pass the structured test suite
- **Educational**: Code serves as learning reference for CCL implementers

## Implementation Levels

### Level 1: Raw Parsing (`Parse`)

The foundation level converts CCL text input into flat key-value pairs.

#### Core Implementation

```go
func (c *CCL) Parse(input string) ([]Entry, error) {
    var entries []Entry
    
    // Handle empty input gracefully
    if strings.TrimSpace(input) == "" {
        return []Entry{}, nil
    }
    
    lines := strings.Split(input, "\n")
    
    for _, line := range lines {
        // Preserve \r characters - only trim spaces and tabs
        originalLine := line
        line = strings.Trim(line, " \t")
        
        if line == "" {
            continue  // Skip empty lines
        }
        
        // Handle comments (start with /=)
        if strings.HasPrefix(line, "/=") {
            comment := strings.Trim(strings.TrimPrefix(line, "/="), " \t")
            entries = append(entries, Entry{
                Key:   "/",      // Special key for comments
                Value: comment,
            })
            continue
        }
        
        // Handle key-value pairs
        if strings.Contains(line, "=") {
            parts := strings.SplitN(line, "=", 2)
            if len(parts) == 2 {
                key := strings.Trim(parts[0], " \t")
                value := strings.Trim(parts[1], " \t")
                
                // Preserve CRLF if present in original line
                if strings.HasSuffix(originalLine, "\r") && !strings.HasSuffix(value, "\r") {
                    value = value + "\r"
                }
                
                entries = append(entries, Entry{
                    Key:   key,
                    Value: value,
                })
            }
        }
        // Note: Lines without '=' are skipped in Level 1 (not errors)
    }
    
    return entries, nil
}
```

#### Key Design Decisions

**Whitespace Handling**:
- Preserve `\r` characters for CRLF compatibility
- Trim only spaces and tabs, not all whitespace
- Handle edge cases with mixed line endings

**Comment Processing**:
- Use `/` as special key for comments
- Preserve comment text after `/=` prefix
- Allow comments to be filtered or processed separately

**Error Strategy**:
- Level 1 is permissive - skip invalid lines rather than error
- Focus on extracting valid key-value pairs
- Detailed error messages for debugging when needed

### Level 1: Object Construction (`MakeObjects`)

The most complex level transforms flat entries into nested object hierarchies.

#### Core Object Construction

```go
func (c *CCL) MakeObjects(entries []Entry) map[string]interface{} {
    result := make(map[string]interface{})
    
    for _, entry := range entries {
        key := entry.Key
        value := entry.Value
        
        if key == "" {
            // Handle empty keys as list items
            handleEmptyKey(result, value)
        } else if strings.Contains(key, ".") {
            // Handle dotted keys - create nested structure
            handleDottedKey(result, key, value)
        } else {
            // Handle regular keys - support duplicate key lists
            handleRegularKey(result, key, value)
        }
    }
    
    return result
}
```

#### Empty Key Handling

```go
func handleEmptyKey(result map[string]interface{}, value string) {
    if existing, exists := result[""]; exists {
        if list, ok := existing.([]interface{}); ok {
            result[""] = append(list, value)
        } else {
            // Convert single value to list
            result[""] = []interface{}{existing, value}
        }
    } else {
        result[""] = []interface{}{value}
    }
}
```

**Empty Key Strategy**: Empty keys create arrays under the empty string key.

#### Dotted Key Expansion

```go
func handleDottedKey(result map[string]interface{}, key, value string) {
    parts := strings.Split(key, ".")
    current := result
    
    for i, part := range parts {
        if i == len(parts)-1 {
            // Last part - set the value
            current[part] = value
        } else {
            // Intermediate part - create nested object
            if _, exists := current[part]; !exists {
                current[part] = make(map[string]interface{})
            }
            if nested, ok := current[part].(map[string]interface{}); ok {
                current = nested
            }
        }
    }
}
```

**Nested Object Creation**: Automatically creates intermediate objects for dotted paths.

#### Duplicate Key Handling

```go
func handleRegularKey(result map[string]interface{}, key, value string) {
    if existing, exists := result[key]; exists {
        if list, ok := existing.([]interface{}); ok {
            result[key] = append(list, value)
        } else {
            // Convert single value to list
            result[key] = []interface{}{existing, value}
        }
    } else {
        result[key] = value
    }
}
```

**List Conversion**: Duplicate keys automatically become arrays.

### Level 2: Typed Access

Type-safe value extraction with automatic conversion.

#### Generic Value Navigation

```go
func (c *CCL) getValue(obj map[string]interface{}, path []string) (interface{}, error) {
    current := obj
    
    for i, key := range path {
        if i == len(path)-1 {
            // Last key - return the value
            if value, exists := current[key]; exists {
                return value, nil
            }
            return nil, fmt.Errorf("key not found: %s (available keys: %v)", 
                strings.Join(path, "."), getMapKeys(current))
        } else {
            // Intermediate key - navigate deeper
            if value, exists := current[key]; exists {
                if nested, ok := value.(map[string]interface{}); ok {
                    current = nested
                } else {
                    return nil, fmt.Errorf("not an object at key: %s", key)
                }
            } else {
                return nil, fmt.Errorf("intermediate key not found: %s in path %s", 
                    key, strings.Join(path, "."))
            }
        }
    }
    
    return nil, fmt.Errorf("invalid path")
}
```

#### Type-Specific Accessors

**String Access**:
```go
func (c *CCL) GetString(obj map[string]interface{}, path []string) (string, error) {
    value, err := c.getValue(obj, path)
    if err != nil {
        return "", err
    }
    
    if str, ok := value.(string); ok {
        return str, nil
    }
    
    return fmt.Sprintf("%v", value), nil  // Convert any type to string
}
```

**Integer Access**:
```go
func (c *CCL) GetInt(obj map[string]interface{}, path []string) (int, error) {
    value, err := c.getValue(obj, path)
    if err != nil {
        return 0, err
    }
    
    if str, ok := value.(string); ok {
        return strconv.Atoi(str)  // Parse string to int
    }
    
    if i, ok := value.(int); ok {
        return i, nil  // Return native int
    }
    
    return 0, fmt.Errorf("cannot convert value %v (type %T) to int at path %s", 
        value, value, strings.Join(path, "."))
}
```

**Boolean Access**:
```go
func (c *CCL) GetBool(obj map[string]interface{}, path []string) (bool, error) {
    value, err := c.getValue(obj, path)
    if err != nil {
        return false, err
    }
    
    if str, ok := value.(string); ok {
        return strconv.ParseBool(str)  // Parse "true"/"false"
    }
    
    if b, ok := value.(bool); ok {
        return b, nil  // Return native bool
    }
    
    return false, fmt.Errorf("cannot convert value %v (type %T) to bool at path %s", 
        value, value, strings.Join(path, "."))
}
```

**List Access**:
```go
func (c *CCL) GetList(obj map[string]interface{}, path []string) ([]string, error) {
    value, err := c.getValue(obj, path)
    if err != nil {
        return nil, err
    }
    
    if arr, ok := value.([]interface{}); ok {
        result := make([]string, len(arr))
        for i, v := range arr {
            result[i] = fmt.Sprintf("%v", v)  // Convert each element to string
        }
        return result, nil
    }
    
    if list, ok := value.([]string); ok {
        return list, nil  // Return native string slice
    }
    
    return nil, fmt.Errorf("cannot convert value %v (type %T) to []string at path %s", 
        value, value, strings.Join(path, "."))
}
```

### Level 3: Advanced Processing

Entry processing functions transform and combine parsed entries.

#### Filter Implementation

```go
func (c *CCL) Filter(entries []Entry) []Entry {
    // Mock implementation preserves all entries
    // Real implementation might filter based on criteria
    return entries
}
```

**Design Pattern**: The mock keeps filtering simple but extensible.

#### Compose Implementation

```go
func (c *CCL) Compose(left, right []Entry) []Entry {
    // Simple concatenation for mock
    result := make([]Entry, len(left)+len(right))
    copy(result, left)
    copy(result[len(left):], right)
    return result
}
```

**Merge Strategy**: Concatenation allows duplicate keys to be handled at Level 1.

#### ExpandDotted Implementation

```go
func (c *CCL) ExpandDotted(entries []Entry) []Entry {
    // Mock implementation passes through - expansion happens in MakeObjects
    return entries
}
```

**Deferred Processing**: Dotted key expansion is handled during object construction for simplicity.

### Level 4: Pretty Printing (`PrettyPrint`)

Generate formatted CCL output from object structures.

```go
func (c *CCL) PrettyPrint(obj map[string]interface{}) string {
    var lines []string
    c.prettyPrintObject(obj, "", &lines)
    return strings.Join(lines, "\n")
}

func (c *CCL) prettyPrintObject(obj map[string]interface{}, prefix string, lines *[]string) {
    for key, value := range obj {
        if nested, ok := value.(map[string]interface{}); ok {
            // Recursively print nested objects with dotted keys
            c.prettyPrintObject(nested, prefix+key+".", lines)
        } else {
            // Print flat key-value pair
            *lines = append(*lines, fmt.Sprintf("%s%s = %v", prefix, key, value))
        }
    }
}
```

## Implementation Patterns

### Error Handling Strategy

The mock implementation follows consistent error handling patterns:

#### Detailed Error Messages

```go
return nil, fmt.Errorf("key not found: %s (available keys: %v)", 
    strings.Join(path, "."), getMapKeys(current))
```

**Context Provision**: Errors include available alternatives for debugging.

#### Type Information

```go
return 0, fmt.Errorf("cannot convert value %v (type %T) to int at path %s", 
    value, value, strings.Join(path, "."))
```

**Type Awareness**: Errors show actual and expected types.

#### Path Context

```go
return nil, fmt.Errorf("intermediate key not found: %s in path %s", 
    key, strings.Join(path, "."))
```

**Navigation Context**: Errors show full path context for nested lookups.

### Data Structure Patterns

#### Interface{} Usage

The mock uses `map[string]interface{}` and `[]interface{}` for maximum flexibility:

```go
type ObjectMap = map[string]interface{}
type ValueList = []interface{}
```

**Flexibility**: Supports any JSON-compatible data structure.
**Simplicity**: Avoids complex type hierarchies.
**Testability**: Easy to compare with expected test results.

#### Entry Representation

```go
type Entry struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```

**Simplicity**: Flat structure matches test JSON format.
**Serialization**: JSON tags support direct marshaling.
**Extensibility**: Easy to add metadata fields.

### Performance Considerations

#### Memory Management

```go
result := make([]Entry, len(left)+len(right))
copy(result, left)
copy(result[len(left):], right)
```

**Pre-allocation**: Avoids repeated allocations during composition.

#### String Operations

```go
line = strings.Trim(line, " \t")  // Specific characters only
```

**Targeted Trimming**: Avoids unnecessary allocations from `TrimSpace`.

#### Map Lookups

```go
if value, exists := current[key]; exists {
    // Process value
}
```

**Single Lookup**: Check existence and retrieve value in one operation.

## Testing Integration

### Mock Capability Declaration

The mock implementation is designed to pass specific test categories:

```go
// Supported functions (Level 1, 2)
var supportedTags = []string{
    "function:parse",
    "function:make-objects", 
    "function:get-string",
    "function:get-int",
    "function:get-bool",
    "function:get-float",
    "function:get-list",
}

// Skipped advanced features
var skippedFeatures = []string{
    "feature:unicode",     // Complex character handling
    "feature:multiline",   // Advanced parsing modes
    "variant:proposed-behavior", // Experimental features
}
```

### Test Generation Strategy

The mock works with the generator's filtering system:

```bash
# Generate tests that mock can pass
just generate-mock  # Uses --run-only with supported tags

# Generate basic implementation tests  
just generate --run-only function:parse,function:make-objects,function:get-string
```

### Progressive Testing

The mock supports incremental testing:

```bash
# Test Level 1 only
just test-level1

# Test Levels 1 and 2
just test --levels 1,2

# Test specific features
just test --features parsing,objects
```

## Extension Patterns

### Adding New Functions

To add a new CCL function:

1. **Define Function Signature**:
```go
func (c *CCL) NewFunction(param ParamType) (ReturnType, error) {
    // Implementation
}
```

2. **Add Error Handling**:
```go
if err := validateInput(param); err != nil {
    return nil, fmt.Errorf("new function failed: %w", err)
}
```

3. **Update Test Tags**:
```json
"tags": ["function:new-function", "level:X"]
```

4. **Add Generator Support**:
```go
case "new_function":
    return generateNewFunctionTest(validation)
```

### Adding New Features

To add CCL language features:

1. **Extend Parser Logic**:
```go
if hasNewFeature(line) {
    return handleNewFeature(line, lineNum)
}
```

2. **Add Feature Tags**:
```json
"tags": ["feature:new-feature"]
```

3. **Update Documentation**:
- Add to API.md function list
- Include examples in DEVELOPER_GUIDE.md

### Behavior Variants

To support different implementation behaviors:

1. **Add Behavior Flag**:
```go
type Options struct {
    StrictMode bool
    // Other options
}
```

2. **Conditional Logic**:
```go
if c.options.StrictMode {
    return handleStrict(input)
}
return handleLenient(input)
```

3. **Tag Conflicts**:
```json
"conflicts": ["behavior:lenient-mode"]
```

## Common Implementation Challenges

### Whitespace Handling

**Challenge**: Different systems expect different whitespace behavior.

**Solution**: Preserve significant whitespace, trim only spaces/tabs:

```go
line = strings.Trim(line, " \t")  // Not TrimSpace()
```

### Line Ending Compatibility

**Challenge**: CRLF vs LF line endings across platforms.

**Solution**: Preserve original line ending characters:

```go
if strings.HasSuffix(originalLine, "\r") && !strings.HasSuffix(value, "\r") {
    value = value + "\r"
}
```

### Type Conversion Edge Cases

**Challenge**: Converting between string and native types safely.

**Solution**: Explicit type checking with detailed errors:

```go
if str, ok := value.(string); ok {
    return strconv.ParseBool(str)
}
return false, fmt.Errorf("type conversion failed: %T to bool", value)
```

### Nested Object Navigation

**Challenge**: Safe navigation through nested object hierarchies.

**Solution**: Step-by-step validation with context:

```go
for i, key := range path {
    if value, exists := current[key]; exists {
        // Validate type before proceeding
        if nested, ok := value.(map[string]interface{}); ok {
            current = nested
        } else {
            return nil, fmt.Errorf("not an object at key: %s", key)
        }
    }
}
```

## Performance Optimization

### Hot Path Optimization

Optimize the most frequently called functions:

1. **Parse Function**: Use efficient string operations
2. **getValue**: Minimize map lookups  
3. **MakeObjects**: Pre-allocate data structures

### Memory Usage

Minimize allocations in loops:

```go
// Pre-allocate with known capacity
result := make([]Entry, 0, len(lines))

// Reuse string slices where possible  
parts := strings.SplitN(line, "=", 2)
```

### Benchmarking Integration

The mock integrates with the benchmark system:

```bash
just benchmark  # Measures mock implementation performance
```

This comprehensive guide provides the foundation for understanding and extending the CCL mock implementation. The patterns and strategies documented here can be applied to create production CCL implementations in any programming language.