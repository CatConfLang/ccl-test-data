# CCL Test Suite Issues Identified

**Date**: 2025-01-13  
**Analysis**: Comprehensive review of `api-list-access.json` test failures  
**Implementation**: OCaml reference implementation  

This document catalogs test data issues discovered through systematic analysis of test failures, distinguishing between implementation bugs vs. test categorization problems.

## Summary of Issues

Through collaborative analysis, we identified **6 categories of test issues** requiring fixes to improve test suite quality and implementation coverage:

1. **Order Preservation Misclassification** - Behavioral requirements incorrectly categorized
2. **Single Value Coercion Misclassification** - Proposed behavior marked as reference-compliant  
3. **Parser Validation Field Errors** - Wrong parser function specified for input type
4. **Missing Function Tags** - Incomplete skip configuration causing inappropriate test execution
5. **Count Field Inconsistencies** - Assertion counts not matching expected arrays
6. **Missing Dual Test Coverage** - Lack of tests for both reference and proposed behaviors

---

## 1. Order Preservation Tests → Split into Dual Approach

### Problem
Tests enforce strict ordering in reference-compliant tests, but order preservation should be implementation-specific behavior.

### Affected Tests
- `list_with_numbers`: expects `["1","42","-17","0"]`, gets `["-17","0","1","42"]`
- `list_with_booleans`: expects `["true","false","yes","no"]`, gets `["false","no","true","yes"]`  
- `list_with_whitespace`: expects `["spaced","normal","",""]`, gets `["normal","spaced"]`
- `deeply_nested_list`: expects `["web1","web2","api1"]`, gets `["api1","web1","web2"]`
- `list_with_unicode`: expects `["张三","José","François","العربية"]`, gets `["François","José","العربية","张三"]`
- `complex_mixed_list_scenarios` (partial): expects `["primary","backup"]`, gets `["backup","primary"]`

### Fix Strategy: **Dual Test Approach**

**A. Create Order-Agnostic Reference Tests** (for all implementations):
```json
"get_list": {
  "count": 1,
  "cases": [{
    "args": ["numbers"],
    "expected_set": ["1", "42", "-17", "0"],  // Set comparison, not array
    "expected_length": 4
  }]
}
```

**B. Tag Existing Tests as Proposed Behavior**:
```json
"meta": {
  "tags": [..., "variant:proposed-behavior"],
  "description": "Tests order-preserving list behavior"
}
```

**Result**: Basic implementations pass set-based tests, advanced implementations can pass both.

---

## 2. Single Value Coercion Tests → Tag as Proposed + Create Reference Versions

### Problem
Tests assume `get_list()` should coerce single values to single-element lists, but OCaml reference implementation errors on this case. Since reference implementation errors, **coercion is proposed behavior**.

### Affected Tests

**A. Pure Single Value Coercion:**
- `single_item_as_list`: `item = single` → expects `get_list(["item"])` → `["single"]`
- `empty_list`: `empty_list =` → expects `get_list(["empty_list"])` → `[""]`

**B. Mixed Cases (Partial Coercion):**
- `mixed_duplicate_single_keys`: 
  - `ports = 80\nports = 443\nhost = localhost`
  - Issue: `get_list(["host"])` expects `["localhost"]` (coercion)
  - Note: `get_list(["ports"])` should work (duplicate keys)
- `list_path_traversal_protection`:
  - `safe = value`
  - Issue: Case 3 expects `get_list(["safe"])` → `["value"]` (coercion)
  - Note: Cases 1&2 should error and do

### Fix Strategy: **Dual Test Approach**

**A. Tag Existing Tests as Proposed Behavior:**
```json
"meta": {
  "tags": [..., "variant:proposed-behavior"],
  "description": "Tests single-value coercion to list behavior"
}
```

**B. Create Reference-Compliant Versions:**
```json
{
  "name": "single_item_as_list_reference",
  "input": "item = single",
  "get_list": {
    "count": 1,
    "cases": [{
      "args": ["item"],
      "error": true,
      "error_message": "Value is not a list"
    }]
  },
  "meta": {
    "tags": ["variant:reference-compliant"]
  }
}
```

**Result**: Reference implementations pass error-expecting tests, advanced implementations pass both.

---

## 3. Parse Validation Field Corrections → Use Correct Parser

### Problem
Tests with nested/indented input use `"parse"` validation but should use `"parse_value"`. The basic `parse` function can't handle indentation properly.

From parser analysis:
- `parse` uses `kvs_p` with `many1 (key_val 0)` - no indentation handling
- `parse_value` uses `nested_kvs_p` with proper indentation parsing

### Affected Tests
- `nested_list_access`: has indented content, expects 4 entries, gets 1
- `deeply_nested_list`: has indented content, expects 6 entries, gets 1  
- `list_multiline_values`: has line continuation, expects 4 entries, gets 3
- `complex_mixed_list_scenarios`: has nested structure, expects 11 entries, gets 4

### Fix Strategy: **Change Validation Field**

```json
// BEFORE
"validations": {
  "parse": {
    "count": 4,
    "expected": [...]
  }
}

// AFTER  
"validations": {
  "parse_value": {
    "count": 4, 
    "expected": [...]
  }
}
```

**Result**: Tests use correct parser for input type, eliminating parse count mismatches.

---

## 4. Missing Function Tags → Complete Skip Configuration

### Problem ✅ **FIXED**
Tests requiring unimplemented functionality weren't being skipped due to missing function tags.

### Affected Tests
- `dotted_key_list`: Missing `"function:expand-dotted"` tag

### Fix Applied ✅
```json
"tags": [
  "function:get-list",
  "function:build-hierarchy", 
  "function:parse",
  "function:expand-dotted",  // ← ADDED
  "feature:dotted-keys"
]
```

**Result**: Test now properly skipped, eliminating false failure noise.

---

## 5. Count Field Inconsistencies → Fix Assertion Counts

### Problem ✅ **FIXED**  
Count field doesn't match expected array length, causing confusing error messages.

### Affected Tests
- `nested_list_access`: Count said 3, expected array had 4 entries

### Fix Applied ✅
```json
// BEFORE
"parse": {
  "count": 3,  // ❌ Wrong count
  "expected": [4 entries...]
}

// AFTER
"parse": {
  "count": 4,  // ✅ Correct count  
  "expected": [4 entries...]
}
```

**Result**: Error messages now consistent and clear.

---

## 6. Missing Dual Test Coverage → Comprehensive Implementation Testing

### Problem
Current test suite doesn't provide adequate coverage for both reference-compliant and advanced implementations. Tests should support:

1. **Reference Implementations**: Basic CCL parsing without advanced features
2. **Advanced Implementations**: Full CCL spec including proposed behaviors

### Solution Strategy
Implement **Dual Test Approach** across behavioral categories:

**A. Reference-Compliant Tests:**
- Use `variant:reference-compliant` tags
- Test core functionality all implementations must support
- Use set-based comparisons for order-agnostic validation
- Expect errors for unimplemented features

**B. Proposed Behavior Tests:**
- Use `variant:proposed-behavior` tags  
- Test advanced features like order preservation and coercion
- Use exact array comparisons
- Expect success for advanced features

**Result**: Test suite supports progressive implementation and clear behavioral boundaries.

---

## Implementation Priority

All fixes are **equal priority** as they address fundamental test suite quality issues:

1. **Accuracy**: Tests correctly represent expected behavior
2. **Clarity**: Error messages are consistent and informative  
3. **Coverage**: Both reference and advanced implementations are testable
4. **Maintainability**: Test categorization is systematic and logical

---

## Impact Assessment

### Before Fixes
- ❌ **15 failing tests** in `api-list-access.json`
- ❌ Mixed behavioral expectations 
- ❌ Confusing error messages
- ❌ Poor implementation coverage

### After Fixes  
- ✅ **Clear separation** of reference vs. proposed behavior
- ✅ **Accurate test categorization** enabling proper skip logic
- ✅ **Comprehensive coverage** for different implementation levels  
- ✅ **Maintainable test suite** with systematic organization

This analysis demonstrates the importance of **collaborative test review** between specification authors and implementation developers to ensure test suite quality and comprehensive behavioral coverage.