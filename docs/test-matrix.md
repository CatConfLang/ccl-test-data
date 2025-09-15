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
| **property-round-trip.json** | 4 | 12 | 24 | parse, pretty-print | round-trip | Consistency |
| **property-algebraic.json** | 4 | 8 | 16 | All functions | algebraic | Properties |

**Total**: 462 tests, 924 assertions across 10 files

## Progressive Implementation Matrix

### Level 1: Core CCL (parse + build_hierarchy)
**Required Files**: api-essential-parsing.json, api-object-construction.json  
**Optional Files**: api-comprehensive-parsing.json, api-dotted-keys.json  
**Functions**: `parse()`, `load()`, `build-hierarchy()`  
**Coverage**: 191 tests, 382 assertions

### Level 2: Typed Access
**Required Files**: api-typed-access.json  
**Prerequisites**: Level 1 complete  
**Functions**: `get-string()`, `get-int()`, `get-bool()`, `get-float()`, `get-list()`  
**Coverage**: 151 tests, 302 assertions

### Level 3: Advanced Processing
**Required Files**: api-processing.json  
**Optional Files**: api-comments.json  
**Prerequisites**: Level 1-2 complete  
**Functions**: `filter()`, `compose()`, `expand-dotted()`  
**Coverage**: 82 tests, 164 assertions

### Level 4: Experimental Features
**Required Files**: property-round-trip.json, property-algebraic.json  
**Optional Files**: api-errors.json  
**Prerequisites**: Level 1-3 complete  
**Functions**: `pretty-print()`, error handling  
**Coverage**: 38 tests, 76 assertions

## Function Implementation Matrix

| Function | Level | Required Files | Test Count | Implementation Priority |
|----------|-------|----------------|------------|------------------------|
| **parse()** | 1 | api-essential-parsing.json | 45 | 🔴 Critical - Start Here |
| **load()** | 1 | api-essential-parsing.json | 12 | 🔴 Critical |
| **filter()** | 3 | api-processing.json, api-comments.json | 25 | 🟡 Important |
| **compose()** | 3 | api-processing.json | 18 | 🟡 Important |
| **expand-dotted()** | 3 | api-processing.json, api-dotted-keys.json | 24 | 🟢 Optional |
| **make-objects()** | 1 | api-object-construction.json | 89 | 🔴 Critical |
| **build-hierarchy()** | 1 | api-object-construction.json | 89 | 🔴 Critical |
| **get-string()** | 2 | api-typed-access.json | 45 | 🟡 Important |
| **get-int()** | 2 | api-typed-access.json | 38 | 🟡 Important |
| **get-bool()** | 2 | api-typed-access.json | 32 | 🟡 Important |
| **get-float()** | 2 | api-typed-access.json | 26 | 🟡 Important |
| **get-list()** | 2 | api-typed-access.json | 10 | 🟢 Optional |
| **pretty-print()** | 4 | property-round-trip.json | 12 | 🟢 Optional |

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
- ✅ Level 1: Parse, Load, BuildHierarchy (100% pass rate)
- ✅ Level 2: GetString, GetInt, GetBool, GetFloat (85% pass rate)
- ✅ Level 3: Filter, Compose, ExpandDotted (95% pass rate)  
- ⚠️ Level 4: PrettyPrint (70% pass rate - edge cases)

### ccl_gleam (Gleam Implementation)  
- ✅ Level 1: parse, load, build_hierarchy (100% pass rate)
- ✅ Level 2: get_string, get_int, get_bool, get_float (100% pass rate)
- ✅ Level 3: Comment filtering (95% pass rate)
- ❌ Level 3: Decorative sections (not implemented)
- ✅ Level 4: Pretty printer (100% pass rate)

## Test Execution Strategy

### Minimal Implementation (MVP)
```bash
# ccl-test-data repository
just generate --run-only function:parse,function:make-objects,function:get-string
just test-level1 && just test-level3 && just test-basic-typed
```

### Progressive Implementation
```bash
# Level 1 (core CCL)
just test-level1
# Level 2 (typed access)  
just test-level2
# Level 3 (advanced processing)
just test-level3
# Level 4 (experimental features)
just test-level4
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

### Gate 2: Types (Level 2)
- ✅ api-typed-access.json: 85%+ pass rate
- ✅ Type conversion and validation work
- ✅ Error handling for invalid types

### Gate 3: Processing (Level 3)  
- ✅ api-processing.json: 90%+ pass rate
- ✅ Comment filtering works correctly
- ✅ Entry composition handles indentation

### Gate 4: Production (Level 4)
- ✅ property-round-trip.json: 100% pass rate
- ✅ All error conditions handled gracefully
- ✅ Performance meets requirements

## Implementation Guidance

### Start Here (New Implementation)
1. **Choose Level 1**: Start with api-essential-parsing.json and api-object-construction.json
2. **Implement parse() + build-hierarchy()**: Core CCL functionality
3. **Validate Early**: Run tests frequently during development
4. **Add Level 2**: Implement typed access for full functionality
5. **Optional Level 3**: Add processing functions as needed

### Common Pitfalls
- **Skipping Object Construction**: Level 2 typed access requires hierarchical objects
- **Overengineering Level 3**: Entry processing is optional for MVP
- **Ignoring Edge Cases**: api-comprehensive-parsing.json catches boundary conditions
- **Poor Error Handling**: api-errors.json validates graceful failure modes

### Success Metrics
- **Functional**: 85%+ pass rate on required test files
- **Progressive**: Each level builds correctly on previous levels  
- **Compatible**: Cross-language compatibility via JSON test format
- **Maintainable**: Test suite guides refactoring and feature additions