# CCL Test Schema Simplification Analysis

## Context
Analyzed the test-runner implementation guide to identify simplifications for making test runner implementations easier. Focused on schema complexity reduction while maintaining functionality.

## Key Findings

### Current Schema Complexity Issues
1. **Multi-validation per test**: Complex partial execution strategies needed
2. **Redundant tagging**: Function tags duplicate validation keys 
3. **Mixed return types**: Arrays, objects, primitives require different handling
4. **Complex dependency resolution**: Between validations within tests
5. **Nested structure**: Multiple levels of nesting complicate parsing

### Maintenance Analysis
- **Current duplication**: Only ~21% (38/181 tests have duplicate inputs)
- **Low maintenance burden**: Most tests have unique inputs for specific edge cases
- **Controlled duplication**: Occurs mainly for fundamental test cases across validation types

## Recommended Solution: Dual Format Approach

### Source Format (Maintainable)
```json
{
  "name": "basic_parsing_workflow",
  "input": "name = Alice\nage = 42",
  "tests": [
    {"function": "parse", "expect": [...]},
    {"function": "build_hierarchy", "expect": {...}},
    {"function": "get_string", "args": ["name"], "expect": "Alice"}
  ],
  "features": ["comments"],
  "level": 3
}
```

### Generated Format (Implementation-Friendly)
```json
{
  "name": "basic_parsing_workflow_parse",
  "input": "name = Alice\nage = 42",
  "validation": "parse",
  "features": ["comments"],
  "level": 3,
  "expected": {"entries": [...], "count": 2}
}
```

## Benefits
1. **Maintainers**: Work with intuitive source format, no input duplication
2. **Implementers**: Get simple flat format, no complex execution logic
3. **Migration**: Gradual transition from current complex format
4. **Validation**: Cross-format consistency checks

## Implementation Strategy
- Generator transforms source → flat tests (1:N relationship)
- Auto-infer dependencies based on CCL level hierarchy
- Standardize all expected results to uniform `{value/entries/object, count}` format
- Build system integration with `source/` and `generated/` directories

## Impact
- Test runner complexity reduced from 500+ lines to ~100 lines
- Eliminates need for partial execution strategies from implementation guide
- Maintains current test coverage while dramatically simplifying implementation
- Creates clear migration path for existing test runners