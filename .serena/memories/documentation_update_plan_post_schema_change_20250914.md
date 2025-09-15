# Documentation Update Plan - Post Schema Change (September 14, 2025)

## Context: Major Schema Transformation Completed

The CCL test suite recently underwent a **fundamental architectural shift** from unified tags to separate typed fields, representing the most significant schema change to date. This plan documents the required documentation updates to align with the new architecture.

## Schema Change Summary

### Before (Unified Tags)
```json
"meta": {
  "tags": ["function:parse", "feature:comments", "behavior:crlf-preserve"],
  "level": 1,
  "feature": "parsing"
}
```

### After (Separate Typed Fields)
```json
// Source format (maintainable)
"meta": {
  "tags": ["function:parse", "feature:comments"],  // Still present for source
  "level": 1,
  "feature": "parsing"
}

// Generated format (implementation-friendly)
{
  "functions": ["parse"],
  "features": ["comments"], 
  "behaviors": ["crlf_preserve"],
  "variants": ["reference_compliant"],
  "level": 1,
  "conflicts": {
    "behaviors": ["crlf_normalize_to_lf"]
  }
}
```

### Architecture Benefits
- **Type-safe filtering**: Direct field access vs string parsing
- **Better API ergonomics**: Separate arrays for different metadata types
- **Conflict resolution**: Categorized structure for mutually exclusive behaviors
- **Implementation efficiency**: No tag parsing required in test runners

## Level System Correction: 5-Level → 4-Level

### New 4-Level Architecture
1. **Level 1**: Core CCL (parse + build_hierarchy) - Text to hierarchical objects
2. **Level 2**: Typed Access (get_string, get_int, etc.) - Type-safe value extraction  
3. **Level 3**: Advanced Processing (filter, compose, expand_dotted) - Entry manipulation
4. **Level 4**: Experimental Features (dotted keys, pretty_print) - Implementation-specific

### Old 5-Level System (DEPRECATED)
- Level 5 (Pretty Printing/Formatting) merged into Level 4 (Experimental)

## Required Documentation Updates

### 🔴 CRITICAL - Level System Updates (7 files)

**Files with 5-level references requiring 4-level correction:**

1. **`docs/test-matrix.md`**
   - Update "Level 5: Advanced Features" → "Level 4: Experimental Features"
   - Merge pretty-print into Level 4 sections

2. **`docs/MOCK_IMPLEMENTATION.md`** 
   - Update "Level 5: Pretty Printing" → "Level 4: Pretty Printing"
   - Consolidate formatting functions under Level 4

3. **`docs/CLI_REFERENCE.md`**
   - Update statistics showing "Level 5" → "Level 4" 
   - Recalculate test counts for 4-level system

4. **`docs/test-filtering.md`**
   - Update "pretty-print (Level 5)" → "pretty-print (Level 4)"
   - Update level-based filtering examples

5. **`docs/README.md`**
   - Update "5-level CCL implementation approach" → "4-level"
   - Update architecture descriptions

6. **`docs/DEVELOPER_GUIDE.md`**
   - Update "5-level architecture" → "4-level architecture"
   - Restructure level progression guidance

7. **`docs/API.md`**
   - Update mixed level references throughout
   - Ensure consistent 4-level terminology

### 🟡 IMPORTANT - Schema Structure Updates

**`docs/API.md` TestMeta Structure:**
```go
// CURRENT (outdated)
type TestMeta struct {
    Tags     []string `json:"tags"`
    Level    int      `json:"level"`
    Feature  string   `json:"feature"`
    Conflicts []string `json:"conflicts,omitempty"`
}

// UPDATED (should reflect new separate fields)
type TestMeta struct {
    Functions []string `json:"functions"`
    Features  []string `json:"features"`
    Behaviors []string `json:"behaviors"`
    Variants  []string `json:"variants"`
    Level     int      `json:"level"`
    Conflicts ConflictSpec `json:"conflicts,omitempty"`
}

type ConflictSpec struct {
    Behaviors []string `json:"behaviors,omitempty"`
    Variants  []string `json:"variants,omitempty"`
}
```

### 🟢 VERIFICATION - Implementation Guides

**`docs/test-runner-implementation-guide.md`:**
- Verify examples use separate fields approach
- Update filtering patterns to use direct field access
- Ensure generated format examples are current

**`docs/implementing-ccl.md`:**
- Appears correctly updated to 4-level system
- Verify API examples match new schema
- Check test runner implementation examples

## Files Already Updated ✅

- **`README.md`**: ✅ Completely transformed to dual-format architecture
- **`enhanced-metadata-schema.md`**: ✅ LLM optimization focus (separate from main schema)

## Implementation Strategy

### Phase 1: Level System Corrections
1. Update all 7 files with 5-level → 4-level corrections
2. Ensure consistent level numbering and descriptions
3. Merge Level 5 content into Level 4 appropriately

### Phase 2: Schema Structure Updates  
1. Update API.md with correct TestMeta structure
2. Add separate fields examples and documentation
3. Update type definitions and code examples

### Phase 3: Verification and Validation
1. Verify test-runner-implementation-guide.md alignment
2. Cross-check all schema references for consistency
3. Validate examples against current implementation

### Quality Standards
- **Authoritative tone**: Definitive language, no backward compatibility
- **Technical accuracy**: Schema examples must match implementation
- **Consistency**: Uniform terminology across all documentation
- **Completeness**: No orphaned references to old system

## Success Metrics

✅ **Documentation Alignment**: All docs reflect 4-level system consistently  
✅ **Schema Accuracy**: TestMeta and examples match current implementation  
✅ **Implementation Guidance**: Clear path for new implementers using separate fields  
✅ **Quality Standards**: Professional documentation meeting project standards

## Risk Mitigation

- **Validation**: Cross-reference with working implementation in `internal/mock/`
- **Testing**: Ensure documentation examples can be executed successfully
- **Consistency Checks**: Automated verification of level numbering across files
- **Schema Compliance**: Validate all JSON examples against current schema

## Timeline Considerations

This documentation update is **critical for project integrity** as:
- New implementers will be confused by mixed 4-level/5-level references
- Schema examples that don't match implementation create development friction  
- The separate fields architecture is fundamental to the test suite's value proposition

**Priority**: High - Should be completed before next major release or implementation guidance is provided to external developers.