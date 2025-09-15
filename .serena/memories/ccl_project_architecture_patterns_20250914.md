# CCL Project Architecture Patterns - Cross-Implementation Insights

## Discovered Patterns from ccl-test-lib Implementation

### 1. Dual-Format Architecture Pattern
**Problem**: Need both human-maintainable and implementation-friendly test formats
**Solution**: Source format (TestSuite) for authoring, flat format (TestCase[]) for execution
**Key Insight**: Loader can intelligently detect and handle both formats transparently

```go
// Source format: Multi-validation, human-maintainable
{
  "name": "test",
  "input": "key = value", 
  "validations": {
    "parse": [...],
    "get_string": {...}
  }
}

// Flat format: Single-validation, implementation-friendly  
{
  "name": "test_parse",
  "input": "key = value",
  "validation": "parse", 
  "expected": [...]
}
```

### 2. Type-Safe Capability Declaration Pattern
**Problem**: String-based tag parsing is error-prone and hard to validate
**Solution**: Enum-based capability system with compile-time safety

```go
// Replace error-prone string parsing
tags := parseTagString("function:parse,feature:comments")

// With type-safe enums
impl := ImplementationConfig{
    SupportedFunctions: []CCLFunction{FunctionParse},
    SupportedFeatures: []CCLFeature{FeatureComments},
}
```

**Benefits**: 
- Compile-time validation
- IDE autocompletion  
- Refactoring safety
- Clear capability contracts

### 3. Progressive Implementation Support Pattern
**Problem**: Implementations need clear path from minimal to full CCL support
**Solution**: Level-based organization with capability-driven filtering

```go
// Level 1: Basic parsing
impl.SupportedFunctions = []CCLFunction{FunctionParse}

// Level 2: Add processing  
impl.SupportedFunctions = append(impl.SupportedFunctions, 
    FunctionFilter, FunctionCombine)

// Level 3: Add object construction
impl.SupportedFunctions = append(impl.SupportedFunctions,
    FunctionBuildHierarchy)
```

### 4. Backward Compatibility Bridge Pattern
**Problem**: Existing formats and workflows must continue working during migration
**Solution**: Intelligent format detection with graceful degradation

```go
// Try new format first, fall back to legacy
if err := json.Unmarshal(data, &newFormat); err != nil {
    json.Unmarshal(data, &legacyFormat)
    newFormat = convertLegacy(legacyFormat)
}
```

## Cross-Project Architecture Insights

### Workspace Structure Analysis
From examining `/home/tylerbu/code/claude-workspace/`:

1. **ccl-test-data/**: Official test suite with Go-based infrastructure
2. **ccl_gleam/**: Multi-package Gleam implementation (functional approach)
3. **ccl-test-lib/**: Shared Go infrastructure (new)
4. **tools-monorepo/**: TypeScript ecosystem tooling

**Pattern**: Each implementation approach (imperative Go, functional Gleam, web TypeScript) has distinct architectural needs but shares common test requirements.

### Test Infrastructure Evolution
1. **Phase 1**: Individual implementations with custom test loading
2. **Phase 2**: ccl-test-data with centralized JSON test suite  
3. **Phase 3**: ccl-test-lib with shared loading infrastructure (current)
4. **Phase 4**: Type-safe test generation and enhanced metadata (future)

### Language-Agnostic Patterns Discovered
1. **Capability Declaration**: Every implementation needs to declare what it supports
2. **Progressive Testing**: Clear levels allow incremental implementation
3. **Format Evolution**: Need both stable legacy support and modern enhancements
4. **Metadata Standardization**: Type-safe metadata is superior to string tags

## Implementation Quality Patterns

### Error Handling Architecture
```go
// Consistent error wrapping with context
if err != nil {
    return fmt.Errorf("failed to load %s: %w", filename, err)
}
```

### Configuration Validation
```go
// Self-validating configuration with clear error messages
func (c ImplementationConfig) IsValid() error {
    // Check for conflicting behavior choices
    // Validate required fields
    // Return structured errors
}
```

### Performance Considerations
- **Direct field access** instead of reflection where possible
- **Lazy loading** for large test suites
- **Efficient filtering** using capability maps rather than iteration

## Future Architecture Recommendations

### 1. Test Format Evolution
- Move toward enhanced TestSuite format with structured metadata
- Maintain backward compatibility with current flat arrays
- Add validation metadata (count fields, constraint checking)

### 2. Implementation Registration
```go
// Future: Implementation registry pattern
registry := ccl.NewImplementationRegistry()
registry.Register("ccl-go", cclGoConfig)
registry.Register("ccl-gleam", cclGleamConfig)
```

### 3. Cross-Language Support
- JSON schema for capability declaration
- Language-specific binding generation
- Common test result formats

### 4. Enhanced Filtering
- Complex capability queries (AND/OR logic)
- Performance profiling integration
- Compatibility matrix generation

## Lessons for Future CCL Development

1. **Start with capability declaration** - Know what you support before building
2. **Embrace dual formats** - Author-friendly and execution-friendly can coexist
3. **Type safety pays dividends** - Compile-time validation prevents runtime errors
4. **Progressive complexity** - Clear levels enable incremental adoption
5. **Backward compatibility is crucial** - Existing workflows must continue working

These patterns should guide future CCL implementation development and ecosystem evolution.