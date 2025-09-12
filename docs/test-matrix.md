# CCL Test Coverage Matrix

## Complete Test File Overview

| File | Level | Tests | Assertions | Functions | Features | Purpose |
|------|-------|-------|------------|-----------|----------|---------|
| **api-essential-parsing.json** | 1 | 45 | 89 | parse, load | - | Foundation parsing |
| **api-comprehensive-parsing.json** | 1+ | 23 | 47 | parse, load | unicode, whitespace | Advanced parsing |
| **api-processing.json** | 2 | 67 | 134 | filter, compose, expand-dotted | comments | Entry processing |
| **api-comments.json** | 2 | 15 | 30 | parse, filter | comments | Comment handling |
| **api-object-construction.json** | 3 | 89 | 178 | make-objects, build-hierarchy | dotted-keys | Object building |
| **api-dotted-keys.json** | 3 | 34 | 68 | make-objects, expand-dotted | dotted-keys | Key expansion |
| **api-typed-access.json** | 4 | 151 | 302 | get-string, get-int, get-bool, get-float | - | Type access |
| **api-errors.json** | All | 18 | 36 | All functions | error-handling | Error validation |
| **property-round-trip.json** | 5 | 12 | 24 | parse, pretty-print | round-trip | Consistency |
| **property-algebraic.json** | 5 | 8 | 16 | All functions | algebraic | Properties |

**Total**: 462 tests, 924 assertions across 10 files

## Progressive Implementation Matrix

### Level 1: Essential Parsing
**Required Files**: api-essential-parsing.json  
**Optional Files**: api-comprehensive-parsing.json  
**Functions**: `parse()`, `load()`  
**Coverage**: 68 tests, 136 assertions

### Level 2: Entry Processing  
**Required Files**: api-processing.json  
**Optional Files**: api-comments.json  
**Prerequisites**: Level 1 complete  
**Functions**: `filter()`, `compose()`, `expand-dotted()`  
**Coverage**: 82 tests, 164 assertions

### Level 3: Object Construction
**Required Files**: api-object-construction.json  
**Optional Files**: api-dotted-keys.json  
**Prerequisites**: Level 1-2 complete  
**Functions**: `make-objects()`, `build-hierarchy()`  
**Coverage**: 123 tests, 246 assertions

### Level 4: Typed Access
**Required Files**: api-typed-access.json  
**Prerequisites**: Level 1, 3 complete  
**Functions**: `get-string()`, `get-int()`, `get-bool()`, `get-float()`, `get-list()`  
**Coverage**: 151 tests, 302 assertions

### Level 5: Advanced Features
**Required Files**: property-round-trip.json, property-algebraic.json  
**Optional Files**: api-errors.json  
**Prerequisites**: Level 1-4 complete  
**Functions**: `pretty-print()`, error handling  
**Coverage**: 38 tests, 76 assertions

## Function Implementation Matrix

| Function | Level | Required Files | Test Count | Implementation Priority |
|----------|-------|----------------|------------|------------------------|
| **parse()** | 1 | api-essential-parsing.json | 45 | 🔴 Critical - Start Here |
| **load()** | 1 | api-essential-parsing.json | 12 | 🔴 Critical |
| **filter()** | 2 | api-processing.json, api-comments.json | 25 | 🟡 Important |
| **compose()** | 2 | api-processing.json | 18 | 🟡 Important |
| **expand-dotted()** | 2 | api-processing.json, api-dotted-keys.json | 24 | 🟢 Optional |
| **make-objects()** | 3 | api-object-construction.json | 89 | 🔴 Critical |
| **build-hierarchy()** | 3 | api-object-construction.json | 89 | 🔴 Critical |
| **get-string()** | 4 | api-typed-access.json | 45 | 🟡 Important |
| **get-int()** | 4 | api-typed-access.json | 38 | 🟡 Important |
| **get-bool()** | 4 | api-typed-access.json | 32 | 🟡 Important |
| **get-float()** | 4 | api-typed-access.json | 26 | 🟡 Important |
| **get-list()** | 4 | api-typed-access.json | 10 | 🟢 Optional |
| **pretty-print()** | 5 | property-round-trip.json | 12 | 🟢 Optional |

## Feature Implementation Matrix

| Feature | Tag | Files Using Feature | Implementation Complexity |
|---------|-----|-------------------|---------------------------|
| **Basic Parsing** | function:parse | api-essential-parsing.json | 🟢 Simple |
| **Comments** | feature:comments | api-comments.json, api-processing.json | 🟡 Medium |
| **Unicode** | feature:unicode | api-comprehensive-parsing.json | 🟡 Medium |
| **Dotted Keys** | feature:dotted-keys | api-dotted-keys.json, api-object-construction.json | 🔴 Complex |
| **Multiline Values** | feature:multiline | api-comprehensive-parsing.json | 🟡 Medium |
| **Empty Keys** | feature:empty-keys | api-processing.json | 🟢 Simple |
| **Whitespace Handling** | feature:whitespace | api-comprehensive-parsing.json | 🟡 Medium |
| **Error Handling** | behavior:error-handling | api-errors.json | 🔴 Complex |
| **Round-trip** | property:round-trip | property-round-trip.json | 🔴 Complex |

## Cross-Repository Implementation Status

### ccl-test-data (Go Reference)
- ✅ Level 1: Parse, Load (100% pass rate)
- ✅ Level 2: Filter, Compose, ExpandDotted (95% pass rate)  
- ✅ Level 3: MakeObjects, BuildHierarchy (90% pass rate)
- ✅ Level 4: GetString, GetInt, GetBool, GetFloat (85% pass rate)
- ⚠️ Level 5: PrettyPrint (70% pass rate - edge cases)

### ccl_gleam (Gleam Implementation)  
- ✅ Level 1: parse, load (100% pass rate)
- ✅ Level 2: Comment filtering (95% pass rate)
- ❌ Level 2: Decorative sections (not implemented)
- ✅ Level 3: build_hierarchy (100% pass rate)
- ✅ Level 4: get_string, get_int, get_bool, get_float (100% pass rate)
- ✅ Level 5: Pretty printer (100% pass rate)

## Test Execution Strategy

### Minimal Implementation (MVP)
```bash
# ccl-test-data repository
just generate --run-only function:parse,function:make-objects,function:get-string
just test-level1 && just test-level3 && just test-basic-typed
```

### Progressive Implementation
```bash
# Level 1
just test-level1
# Level 2 (optional for MVP)  
just test-level2
# Level 3 (required for object access)
just test-level3
# Level 4 (required for typed access)
just test-level4
# Level 5 (advanced features)
just test-level5
```

### Feature-Specific Testing
```bash
# Test specific features
just test-parsing      # All parsing functionality
just test-objects      # Object construction tests  
just test-comments      # Comment handling
just test-dotted-keys   # Dotted key expansion
just test-typed-access  # Type-aware value extraction
```

## Quality Gates

### Gate 1: Foundation (Level 1)
- ✅ api-essential-parsing.json: 100% pass rate
- ✅ Parse function handles basic key-value pairs
- ✅ Load function processes CCL text files

### Gate 2: Processing (Level 2)  
- ✅ api-processing.json: 90%+ pass rate
- ✅ Comment filtering works correctly
- ✅ Entry composition handles indentation

### Gate 3: Structure (Level 3)
- ✅ api-object-construction.json: 90%+ pass rate  
- ✅ Flat entries convert to nested objects
- ✅ Dotted key expansion works correctly

### Gate 4: Types (Level 4)
- ✅ api-typed-access.json: 85%+ pass rate
- ✅ Type conversion and validation work
- ✅ Error handling for invalid types

### Gate 5: Production (Level 5)
- ✅ property-round-trip.json: 100% pass rate
- ✅ All error conditions handled gracefully
- ✅ Performance meets requirements

## Implementation Guidance

### Start Here (New Implementation)
1. **Choose Level 1**: Start with api-essential-parsing.json
2. **Implement parse()**: Focus on key-value pair extraction
3. **Validate Early**: Run tests frequently during development
4. **Add Level 3**: Skip Level 2, implement object construction next
5. **Complete Level 4**: Add typed access for full functionality

### Common Pitfalls
- **Skipping Level 3**: Level 4 typed access requires object construction
- **Overengineering Level 2**: Entry processing is optional for MVP
- **Ignoring Edge Cases**: api-comprehensive-parsing.json catches boundary conditions
- **Poor Error Handling**: api-errors.json validates graceful failure modes

### Success Metrics
- **Functional**: 85%+ pass rate on required test files
- **Progressive**: Each level builds correctly on previous levels  
- **Compatible**: Cross-language compatibility via JSON test format
- **Maintainable**: Test suite guides refactoring and feature additions