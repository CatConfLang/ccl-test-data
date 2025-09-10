# CCL Test Schema Update: Feature-Based Tagging System

## Overview

This update introduces a structured feature-based tagging system to replace the current descriptive tags, enabling better test selection for partial CCL implementations.

## Problem Statement

The current tagging system uses descriptive tags like `"basic"`, `"complex"`, `"whitespace"` that don't clearly indicate:
- Which CCL functions are required for a test
- Which language features are optional vs required  
- Which behavioral choices are mutually exclusive (e.g., CRLF handling)

This makes it difficult for implementers to select appropriate test subsets based on their implementation capabilities.

## Solution: Structured Feature Tags

### New Tag Categories

**Function Tags** (`function:*`)
- Indicate which CCL functions are required for the test
- Examples: `function:parse`, `function:make-objects`, `function:get-string`
- Map directly to CCL implementation levels 1-5

**Feature Tags** (`feature:*`)  
- Indicate optional CCL language features
- Examples: `feature:comments`, `feature:dotted-keys`, `feature:multiline`
- Can be skipped by implementations that don't support them

**Behavior Tags** (`behavior:*`)
- Indicate implementation behavioral choices
- Examples: `behavior:crlf-preserve`, `behavior:tabs-to-spaces`
- Mutually exclusive within behavioral groups

**Variant Tags** (`variant:*`)
- Indicate specification variant choices
- Examples: `variant:proposed-behavior`, `variant:reference-compliant`
- Allow testing against different specification versions

## Schema Changes

### Enhanced Test Metadata

**Before:**
```json
{
  "meta": {
    "tags": ["basic", "whitespace"],
    "level": 1,
    "feature": "parsing"
  }
}
```

**After:**
```json
{
  "meta": {
    "tags": ["function:parse", "feature:whitespace", "behavior:crlf-preserve"],
    "conflicts": ["behavior:crlf-normalize"],
    "level": 1,
    "feature": "parsing"
  }
}
```

### New Fields

**`conflicts` (optional array of strings)**
- Lists tags that are mutually exclusive with this test
- Used for behavioral choices and implementation variants
- Enables automatic test exclusion for conflicting configurations

## Migration Guide

### For Test Data Maintainers

1. **Use the tag migration map** in `docs/tag-migration.json` to convert existing tags
2. **Add conflict declarations** for behavioral tests:
   ```json
   {
     "tags": ["function:parse", "behavior:crlf-preserve"],
     "conflicts": ["behavior:crlf-normalize"]
   }
   ```
3. **Preserve backward compatibility** by keeping old tags alongside new ones during transition

### For Implementation Authors

1. **Declare implementation capabilities:**
   ```javascript
   const supportedFunctions = ["parse", "make-objects", "get-string"];
   const supportedFeatures = ["dotted-keys"];
   const behaviorChoices = {
     "line_endings": "crlf-normalize",
     "whitespace": "loose-spacing"
   };
   ```

2. **Filter tests based on capabilities:**
   ```javascript
   function shouldRunTest(test) {
     // Check required functions
     const requiredFunctions = test.meta.tags
       .filter(tag => tag.startsWith('function:'))
       .map(tag => tag.substring(9));
     
     if (!requiredFunctions.every(f => supportedFunctions.includes(f))) {
       return false;
     }
     
     // Check behavioral conflicts
     const behaviorTags = test.meta.tags.filter(tag => tag.startsWith('behavior:'));
     const conflicts = test.meta.conflicts || [];
     
     for (const choice of Object.values(behaviorChoices)) {
       if (conflicts.includes(`behavior:${choice}`)) {
         return false;
       }
     }
     
     return true;
   }
   ```

### Language-Specific Integration

**Go Testing:**
```go
//go:build ccl_comments
func TestCommentParsing(t *testing.T) { 
  // Tests requiring feature:comments
}

//go:build !ccl_pretty_print  
func TestBasicOnly(t *testing.T) { 
  // Tests not requiring function:pretty-print
}
```

**Rust:**
```rust
#[cfg(feature = "ccl-comments")]
#[test] 
fn test_comment_parsing() { 
  // Tests requiring feature:comments
}
```

**Python pytest:**
```python
@pytest.mark.ccl_function("parse", "make-objects")
@pytest.mark.ccl_feature("comments")  
def test_comment_parsing():
    # Tests requiring specific functions and features
    pass
```

## Benefits

1. **Granular Test Selection** - Skip unsupported functions/features precisely
2. **Conflict Resolution** - Automatic handling of mutually exclusive behaviors  
3. **Progressive Implementation** - Clear path from minimal to full implementation
4. **Language Agnostic** - Works with any testing framework via metadata filtering
5. **Backward Compatible** - Existing tests continue to work during migration

## Implementation Timeline

1. **Phase 1:** Add new tag structure to schema, update docs
2. **Phase 2:** Add new tags alongside existing tags in test files  
3. **Phase 3:** Update mock implementation and test runner to use new tags
4. **Phase 4:** Remove old tags after migration is complete

## Validation

The enhanced schema ensures:
- All `function:*` tags correspond to valid CCL functions
- All `behavior:*` tags have appropriate conflicts declared
- Conflict relationships are symmetric
- Tag categories are mutually consistent