# CCL Test Suite Architecture Understanding

## Variant Tagging System Design

### Purpose
The CCL test suite uses a sophisticated tagging system to handle the reality that different implementations (reference vs proposed) may have different behaviors for the same functionality.

### Core Concept
- **Reference Implementation**: The OCaml implementation defines the "correct" behavior
- **Proposed Behaviors**: Extensions or changes that conflict with reference behavior
- **Implementation Flexibility**: Other implementations can choose which behaviors to support

### Tag Structure

#### Variant Tags
- `variant:reference-compliant` - Behavior matches OCaml reference implementation
- `variant:proposed-behavior` - Behavior differs from reference (proposed extension)

#### Function Tags  
- `function:parse`, `function:make-objects`, etc. - Required CCL functions
- Used to skip tests when functions are unimplemented

#### Behavior Tags
- `behavior:crlf-preserve-literal` vs `behavior:crlf-normalize-to-lf`
- `behavior:tabs-to-spaces` vs `behavior:tabs-preserve` 
- `behavior:boolean-strict` vs `behavior:boolean-lenient`
- `behavior:loose-spacing` vs `behavior:strict-spacing`

#### Conflict System
Tests specify `conflicts` array listing incompatible variants/behaviors:
```json
"conflicts": [
  "variant:reference-compliant", 
  "behavior:crlf-normalize-to-lf"
]
```

## Implementation Configuration Pattern

### Reference Implementation (OCaml)
```ocaml
skip_variants = ["proposed-behavior"];  // Skip non-reference behaviors
skip_behaviors = ["boolean-lenient"];   // Skip unsupported behaviors  
skip_functions = ["expand-dotted"];     // Skip unimplemented functions
```

### Alternative Implementation Example
```javascript
// A proposed-behavior implementation might use:
skip_variants = ["reference-compliant"];  // Skip reference-only behaviors
skip_behaviors = ["boolean-strict"];      // Different behavior preferences
```

## Test Filtering Logic

### Multi-Layer Filtering
1. **Explicit Test Names**: Skip known problematic tests by name
2. **Function Requirements**: Skip if required functions not implemented
3. **Behavior Conflicts**: Skip if test requires unsupported behaviors
4. **Variant Conflicts**: Skip if test variant not supported
5. **Legacy Compatibility**: Additional filtering for backward compatibility

### Configuration-Driven Approach
- Each implementation defines its capabilities via configuration
- Test runner automatically filters incompatible tests
- Enables accurate success rate measurement for implemented features

## Benefits

### For Implementation Authors
- Clear understanding of which tests should pass
- Ability to incrementally implement features
- Automatic filtering prevents false failures

### For Test Suite Maintainers  
- Single test suite supports multiple implementation approaches
- Clear specification of reference vs proposed behaviors
- Extensible system for new behaviors/variants

### For CCL Language Evolution
- Safe experimentation with proposed behaviors
- Clear migration path from proposed to reference
- Backward compatibility preservation

## Real-World Example: CRLF Handling

### Reference Behavior (OCaml)
- Uses `behavior:crlf-preserve-literal`
- Preserves original line endings in input
- Tagged as `variant:reference-compliant`

### Proposed Behavior
- Uses `behavior:crlf-normalize-to-lf` 
- Normalizes all line endings to LF
- Tagged as `variant:proposed-behavior`
- Conflicts with reference behavior

### Test Suite Handling
- Both behaviors have separate tests
- Reference implementation skips proposed tests
- Proposed implementation would skip reference tests
- Both can coexist in same test suite

## Architecture Lessons

### Variant Tagging is Essential
- Enables single test suite for multiple implementation approaches
- Prevents confusion about which tests "should" pass
- Supports language evolution and experimentation

### Configuration Over Code
- Behavior differences handled by configuration, not code changes
- Makes implementation capabilities explicit and testable
- Enables automated compatibility validation

### Test Specification Quality Matters
- Incorrect tagging causes false failures/passes
- Requires careful maintenance as language evolves
- Critical for multi-implementation ecosystem health