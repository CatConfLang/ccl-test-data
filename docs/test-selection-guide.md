# CCL Test Selection Guide

This guide shows how to use the feature-based tagging system to select appropriate tests for your CCL implementation.

## Overview

The CCL test suite uses **structured tags** to enable precise test selection based on implementation capabilities. This allows you to:

- Run only tests your implementation supports
- Skip advanced features you haven't implemented yet
- Avoid conflicting behavioral tests
- Progress incrementally from basic to full CCL support

## Typed Fields Architecture

### Functions Array (`test.functions[]`)

Direct array containing CCL functions required for a test to run:

| Function | Description | Example Usage |
|----------|-------------|---------------|
| `parse` | Basic key-value parsing | `Parse("key = value")` |
| `parse_indented` | Parse with indentation normalization | `ParseIndented("  key = val\n  sub")` |
| `filter` | Entry filtering | `Filter(entries, predicate)` |
| `compose` | Entry composition | `Compose(left, right)` |
| `build_hierarchy` | Object construction | `BuildHierarchy(entries)` |
| `get_string` | String value access | `GetString(obj, "key")` |
| `get_int` | Integer value access | `GetInt(obj, "count")` |
| `get_bool` | Boolean value access | `GetBool(obj, "enabled")` |
| `get_float` | Float value access | `GetFloat(obj, "rate")` |
| `get_list` | List value access | `GetList(obj, "items")` |
| `pretty_print` | Formatting output | `PrettyPrint(obj)` |

**Type-safe filtering:**
```javascript
// Check if implementation supports all required functions
const functionsSupported = test.functions.every(fn => 
  implementedFunctions.includes(fn)
);
```

### Features Array (`test.features[]`)

Direct array containing optional language features that may not be supported by all implementations:

| Feature | Description | Example |
|---------|-------------|---------|
| `comments` | `/=` comment syntax | `/= This is a comment` |
| `empty_keys` | Anonymous list items | `= item1\n= item2` |
| `multiline` | Multi-line values | `description = Line 1\nLine 2` |
| `unicode` | Unicode content | `name = José` |
| `whitespace` | Complex whitespace handling | Preserving tabs, spaces |

**Type-safe filtering:**
```javascript
// Check if implementation supports all required features
const featuresSupported = test.features.every(feature => 
  implementedFeatures.includes(feature)
);
```

### Behaviors Array (`test.behaviors[]`)

Direct array containing **implementation choices** - technical decisions about how to handle specific parsing details. These are stable, well-defined choices that implementations must make regardless of spec interpretation:

| Behavior Group | Options | Description |
|----------------|---------|-------------|
| Line Endings | `crlf_preserve_literal` vs `crlf_normalize_to_lf` | CRLF handling: preserve `\r` chars vs normalize to LF |
| Boolean Parsing | `boolean_lenient` vs `boolean_strict` | Boolean values: accept "yes"/"no" vs only "true"/"false". Note: "true"/"false" work in both modes |
| Tab Handling | `tabs_preserve` vs `tabs_to_spaces` | Tab character processing |
| Whitespace | `strict_spacing` vs `loose_spacing` | Whitespace sensitivity |
| List Access | `list_coercion_enabled` vs `list_coercion_disabled` | List access behavior |

### Variants Array (`test.variants[]`) - Temporary Disambiguation

Direct array containing **specification variant interpretations** for areas where the CCL specification is currently ambiguous or evolving. These exist to handle cases where the spec doesn't clearly define behavior, allowing both the OCaml reference implementation approach and proposed enhanced approaches to coexist during spec evolution.

| Variant | Description | Status |
|---------|-------------|--------|
| `proposed_behavior` | Enhanced/flexible interpretation of ambiguous spec areas | Proposed for future spec |
| `reference_compliant` | Strict compatibility with OCaml reference implementation | Current baseline |

**Type-safe filtering:**
```javascript
// Check for conflicting implementation choices
const hasConflictingBehavior = test.conflicts?.behaviors?.some(behavior => 
  implementationBehaviors.includes(behavior)
);
const hasConflictingVariant = test.conflicts?.variants?.some(variant => 
  implementationVariants.includes(variant)
);
const isCompatible = !hasConflictingBehavior && !hasConflictingVariant;
```

**Important:** These variant values are **temporary disambiguation mechanisms**. Once the CCL specification owners clarify these ambiguities, the variant system will be eliminated and these choices will be converted to specific behaviors (e.g., `multiline-flexible` vs `multiline-strict`).

## Behavior vs Variant: When to Use Which

### Use `behavior:*` tags for:
- **Technical implementation details** (CRLF handling, tab processing, boolean parsing)
- **Choices that are implementation-specific** regardless of spec interpretation
- **Stable decisions** that won't change based on spec evolution

### Use `variant:*` tags for:
- **Specification ambiguities** where the "correct" behavior is unclear
- **Cases where OCaml reference and proposed spec differ**
- **Temporary disambiguation** pending spec clarification

### Examples:

**Implementation Choice (behavior):**
```json
// Technical decision: How to handle line endings
"tags": ["behavior:crlf_normalize_to_lf"]
"conflicts": ["behavior:crlf_preserve_literal"]
```

**Specification Ambiguity (variant):**
```json
// Spec unclear: Should multiline values work without explicit continuation?
"tags": ["variant:proposed_behavior"]  // Allow implicit continuation
"conflicts": ["variant:reference_compliant"]  // Require explicit syntax
```

### Future Evolution

When CCL spec owners resolve ambiguities, tests will be migrated:
- If "reference is canonical" → Remove `variant:proposed_behavior` tests
- If "proposed is canonical" → Remove `variant:reference_compliant` tests
- Remaining variant behaviors become specific `behavior:*` tags

## Implementation Strategies

### 1. Core Functions Only

Basic CCL parsing and object construction:

```json
{
  "supported_functions": ["function:parse"],
  "skip_all_features": true,
  "behavior_choices": {
    "line_endings": "behavior:crlf_normalize_to_lf",
    "boolean_parsing": "behavior:boolean_lenient",
    "whitespace": "behavior:loose_spacing"
  }
}
```

**Tests to run:** ~54 tests focusing only on basic parsing
**Use case:** Configuration file readers, simple parsers

### 2. Basic Implementation (Parse + Objects + Typed Access)

For most CCL use cases:

```json
{
  "supported_functions": [
    "function:parse",
    "function:parse_value",
    "function:build_hierarchy",
    "function:get_string",
    "function:get_int",
    "function:get_bool"
  ],
  "supported_features": ["feature:experimental_dotted_keys"],
  "behavior_choices": {
    "line_endings": "behavior:crlf_normalize_to_lf",
    "tabs": "behavior:tabs_to_spaces",
    "whitespace": "behavior:loose_spacing"
  }
}
```

**Tests to run:** ~120 tests covering parsing, object construction, and basic typed access
**Use case:** Configuration libraries, application settings

### 3. Processing Implementation (Add Advanced Functions)

For advanced configuration manipulation:

```json
{
  "supported_functions": [
    "function:parse",
    "function:parse_value",
    "function:filter",
    "function:compose",
    "function:expand_dotted",
    "function:build_hierarchy",
    "function:get_string",
    "function:get_int",
    "function:get_bool",
    "function:get_float"
  ],
  "supported_features": [
    "feature:experimental_dotted_keys",
    "feature:comments"
  ]
}
```

**Tests to run:** ~150 tests including entry processing and composition
**Use case:** Configuration builders, template systems

### 4. Complete Implementation

All functions and features:

```json
{
  "supported_functions": ["function:*"],
  "supported_features": [
    "feature:comments",
    "feature:experimental_dotted_keys",
    "feature:empty_keys",
    "feature:multiline",
    "feature:unicode"
  ],
  "variant_choice": "variant:proposed_behavior"
}
```

**Tests to run:** All 180 tests with chosen behavioral variants
**Use case:** Complete CCL libraries, specification-compliant implementations

## Language-Specific Integration

### Go Testing

```go
func TestCCLImplementation(t *testing.T) {
    // Define implementation capabilities
    supportedFunctions := []string{
        "function:parse",
        "function:parse_value",
        "function:build_hierarchy",
        "function:get_string",
    }
    supportedFeatures := []string{"feature:experimental_dotted_keys"}
    behaviorChoices := map[string]string{
        "line_endings": "behavior:crlf_normalize_to_lf",
    }
    
    // Load and filter tests
    for _, testFile := range loadTestFiles() {
        for _, test := range testFile.Tests {
            if shouldSkipTest(test, supportedFunctions, supportedFeatures, behaviorChoices) {
                t.Skipf("Skipping %s: unsupported features", test.Name)
                continue
            }
            
            // Run test validations
            runTestValidations(t, test)
        }
    }
}

func shouldSkipTest(test Test, supportedFunctions, supportedFeatures []string, behaviorChoices map[string]string) bool {
    // Check required functions
    for _, tag := range test.Meta.Tags {
        if strings.HasPrefix(tag, "function:") {
            function := strings.TrimPrefix(tag, "function:")
            if !contains(supportedFunctions, "function:"+function) {
                return true
            }
        }
    }
    
    // Check optional features
    for _, tag := range test.Meta.Tags {
        if strings.HasPrefix(tag, "feature:") {
            if !contains(supportedFeatures, tag) {
                return true
            }
        }
    }
    
    // Check behavioral conflicts
    for _, conflict := range test.Meta.Conflicts {
        for _, choice := range behaviorChoices {
            if conflict == choice {
                return true
            }
        }
    }
    
    return false
}
```

### Rust Testing

```rust
#[cfg(test)]
mod ccl_tests {
    use super::*;

    #[test]
    fn test_ccl_implementation() {
        let supported_functions = vec![
            "function:parse",
            "function:parse_value",
            "function:build_hierarchy",
            "function:get_string",
        ];

        let supported_features = vec!["feature:experimental_dotted_keys"];

        let behavior_choices = HashMap::from([
            ("line_endings", "behavior:crlf_normalize_to_lf"),
        ]);
        
        for test_file in load_test_files() {
            for test in test_file.tests {
                if should_skip_test(&test, &supported_functions, &supported_features, &behavior_choices) {
                    continue;
                }
                
                run_test_validations(&test);
            }
        }
    }
}

// Use Rust features for conditional compilation
#[cfg(feature = "ccl-comments")]
#[test]
fn test_comment_parsing() {
    // Tests requiring feature:comments
}

#[cfg(not(feature = "ccl-unicode"))]
#[test] 
fn test_basic_parsing_only() {
    // Tests that don't require feature:unicode
}
```

### Python Testing (pytest)

```python
import pytest

# Define implementation capabilities
SUPPORTED_FUNCTIONS = [
    "function:parse",
    "function:parse_value",
    "function:build_hierarchy",
    "function:get_string"
]

SUPPORTED_FEATURES = ["feature:experimental_dotted_keys"]

BEHAVIOR_CHOICES = {
    "line_endings": "behavior:crlf_normalize_to_lf"
}

def should_skip_test(test):
    # Check function requirements
    for tag in test.meta.tags:
        if tag.startswith("function:") and tag not in SUPPORTED_FUNCTIONS:
            return True
            
    # Check feature requirements  
    for tag in test.meta.tags:
        if tag.startswith("feature:") and tag not in SUPPORTED_FEATURES:
            return True
            
    # Check behavioral conflicts
    for conflict in test.meta.get("conflicts", []):
        if conflict in BEHAVIOR_CHOICES.values():
            return True
            
    return False

@pytest.mark.parametrize("test", load_tests())
def test_ccl_implementation(test):
    if should_skip_test(test):
        pytest.skip(f"Unsupported features: {test.name}")
    
    run_test_validations(test)

# Use pytest markers for feature-based filtering
@pytest.mark.ccl_function("parse", "build_hierarchy")
@pytest.mark.ccl_feature("comments")
def test_comment_parsing():
    # Tests requiring specific functions and features
    pass

# Command line usage:
# pytest -m "ccl_function:parse"  # Only basic parsing tests
# pytest -m "not ccl_feature:unicode"  # Skip unicode tests
```

### JavaScript Testing (Jest)

```javascript
describe('CCL Implementation', () => {
  const supportedFunctions = [
    'function:parse',
    'function:parse_value',
    'function:build_hierarchy',
    'function:get_string'
  ];

  const supportedFeatures = ['feature:experimental_dotted_keys'];

  const behaviorChoices = {
    line_endings: 'behavior:crlf_normalize_to_lf'
  };
  
  function shouldSkipTest(test) {
    // Check function requirements
    for (const tag of test.meta.tags) {
      if (tag.startsWith('function:') && !supportedFunctions.includes(tag)) {
        return true;
      }
    }
    
    // Check feature requirements
    for (const tag of test.meta.tags) {
      if (tag.startsWith('feature:') && !supportedFeatures.includes(tag)) {
        return true;
      }
    }
    
    // Check behavioral conflicts
    for (const conflict of test.meta.conflicts || []) {
      if (Object.values(behaviorChoices).includes(conflict)) {
        return true;
      }
    }
    
    return false;
  }
  
  const testFiles = loadTestFiles();
  
  testFiles.forEach(testFile => {
    testFile.tests.forEach(test => {
      const testName = `${testFile.suite}: ${test.name}`;
      
      if (shouldSkipTest(test)) {
        test.skip(testName, () => {
          // Test skipped due to unsupported features
        });
      } else {
        test(testName, () => {
          runTestValidations(test);
        });
      }
    });
  });
});

// Use describe.skip for feature-based exclusion
describe.skipIf(!SUPPORTS_COMMENTS)('Comment parsing tests', () => {
  // Tests requiring feature:comments
});

describe.skipIf(!SUPPORTS_UNICODE)('Unicode handling tests', () => {
  // Tests requiring feature:unicode  
});
```

## Progressive Implementation Guide

### Phase 1: Basic Parsing
- **Target**: `function:parse` only
- **Tests**: ~54 tests
- **Goal**: Parse key-value pairs, handle empty input
- **Time**: 1-2 days

### Phase 2: Enhanced Parsing
- **Target**: Add `function:parse_value`
- **Tests**: ~65 tests
- **Goal**: Handle indentation-aware parsing
- **Time**: 1-2 days

### Phase 3: Object Construction
- **Target**: Add `function:build_hierarchy`
- **Tests**: ~80 tests
- **Goal**: Convert flat entries to nested objects
- **Features**: Consider adding `feature:experimental_dotted_keys`
- **Time**: 2-3 days

### Phase 4: Typed Access
- **Target**: Add `function:get_string`, `function:get_int`, `function:get_bool`
- **Tests**: ~120 tests
- **Goal**: Type-safe value extraction
- **Time**: 1-2 days

### Phase 5: Processing Functions
- **Target**: Add `function:filter`, `function:compose`, `function:expand_dotted`
- **Tests**: ~150 tests
- **Goal**: Advanced entry manipulation
- **Features**: Consider adding `feature:comments`
- **Time**: 3-4 days

### Phase 6: Complete Implementation
- **Target**: Add remaining functions and features
- **Tests**: All 180 tests
- **Goal**: Complete CCL specification support

## Troubleshooting

### High Test Failure Rate
- Check that your `supported_functions` list matches your actual implementation
- Verify behavioral choices match your implementation's behavior
- Use `just stats` to see test distribution

### Conflicting Test Results
- Review `conflicts` arrays in failing tests
- Ensure behavioral choices are consistent across your implementation
- Check for variant tag conflicts (`variant:proposed_behavior` vs `variant:reference_compliant`)

### Missing Features
- Use feature tags to identify what's needed: `feature:comments`, `feature:unicode`, etc.
- Implement features incrementally, updating supported lists as you go
- Check the tag migration guide in `docs/tag-migration.json` for implementation hints

## Reference

- **Schema**: `tests/schema.json` - Complete tag validation rules
- **Migration Guide**: `docs/tag-migration.json` - Tag mapping and examples  
- **Statistics**: Run `just stats` for current test breakdown
- **Examples**: See language-specific implementation patterns above