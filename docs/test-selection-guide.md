# CCL Test Selection Guide

This guide shows how to use the feature-based tagging system to select appropriate tests for your CCL implementation.

## Overview

The CCL test suite uses **structured tags** to enable precise test selection based on implementation capabilities. This allows you to:

- Run only tests your implementation supports
- Skip advanced features you haven't implemented yet
- Avoid conflicting behavioral tests
- Progress incrementally from basic to full CCL support

## Tag Categories

### Function Tags (`function:*`)

These indicate which CCL functions are required for a test to run:

| Tag | Description | Level | Example Usage |
|-----|-------------|--------|---------------|
| `function:parse` | Basic key-value parsing | 1 | `Parse("key = value")` |
| `function:filter` | Entry filtering | 2 | `Filter(entries, predicate)` |
| `function:compose` | Entry composition | 2 | `Compose(left, right)` |
| `function:expand-dotted` | Dotted key expansion | 2 | `ExpandDotted(entries)` |
| `function:make-objects` | Object construction | 3 | `MakeObjects(entries)` |
| `function:get-string` | String value access | 4 | `GetString(obj, "key")` |
| `function:get-int` | Integer value access | 4 | `GetInt(obj, "count")` |
| `function:get-bool` | Boolean value access | 4 | `GetBool(obj, "enabled")` |
| `function:get-float` | Float value access | 4 | `GetFloat(obj, "rate")` |
| `function:get-list` | List value access | 4 | `GetList(obj, "items")` |
| `function:pretty-print` | Formatting output | 5 | `PrettyPrint(obj)` |

### Feature Tags (`feature:*`)

These indicate optional language features that may not be supported by all implementations:

| Tag | Description | Example |
|-----|-------------|---------|
| `feature:comments` | `/=` comment syntax | `/= This is a comment` |
| `feature:dotted-keys` | Hierarchical key syntax | `database.host = localhost` |
| `feature:empty-keys` | Anonymous list items | `= item1\n= item2` |
| `feature:multiline` | Multi-line values | `description = Line 1\nLine 2` |
| `feature:unicode` | Unicode content | `name = José` |
| `feature:whitespace` | Complex whitespace handling | Preserving tabs, spaces |

### Behavior Tags (`behavior:*`)

These indicate implementation choices that are mutually exclusive:

| Tag Group | Options | Description |
|-----------|---------|-------------|
| Line Endings | `behavior:crlf-preserve` vs `behavior:crlf-normalize` | How to handle `\r\n` sequences |
| Tab Handling | `behavior:tabs-preserve` vs `behavior:tabs-to-spaces` | Tab character processing |
| Whitespace | `behavior:strict-spacing` vs `behavior:loose-spacing` | Whitespace sensitivity |

### Variant Tags (`variant:*`)

These indicate different specification interpretations:

| Tag | Description |
|-----|-------------|
| `variant:proposed-behavior` | Follows proposed CCL specification |
| `variant:reference-compliant` | Matches OCaml reference implementation |

## Implementation Strategies

### 1. Minimal Implementation (Parse Only)

For rapid prototyping or minimal CCL support:

```json
{
  "supported_functions": ["function:parse"],
  "skip_all_features": true,
  "behavior_choices": {
    "line_endings": "behavior:crlf-normalize",
    "whitespace": "behavior:loose-spacing"
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
    "function:make-objects", 
    "function:get-string",
    "function:get-int",
    "function:get-bool"
  ],
  "supported_features": ["feature:dotted-keys"],
  "behavior_choices": {
    "line_endings": "behavior:crlf-normalize",
    "tabs": "behavior:tabs-to-spaces",
    "whitespace": "behavior:loose-spacing"
  }
}
```

**Tests to run:** ~120 tests covering parsing, object construction, and basic typed access
**Use case:** Configuration libraries, application settings

### 3. Processing Implementation (Add Level 2 Functions)

For advanced configuration manipulation:

```json
{
  "supported_functions": [
    "function:parse",
    "function:filter",
    "function:compose", 
    "function:expand-dotted",
    "function:make-objects",
    "function:get-string",
    "function:get-int",
    "function:get-bool",
    "function:get-float"
  ],
  "supported_features": [
    "feature:dotted-keys",
    "feature:comments"
  ]
}
```

**Tests to run:** ~150 tests including entry processing and composition
**Use case:** Configuration builders, template systems

### 4. Full Implementation (All Functions + Features)

For complete CCL implementations:

```json
{
  "supported_functions": ["function:*"],
  "supported_features": [
    "feature:comments",
    "feature:dotted-keys", 
    "feature:empty-keys",
    "feature:multiline",
    "feature:unicode"
  ],
  "variant_choice": "variant:proposed-behavior"
}
```

**Tests to run:** All 167 tests with chosen behavioral variants
**Use case:** Complete CCL libraries, specification-compliant implementations

## Language-Specific Integration

### Go Testing

```go
func TestCCLImplementation(t *testing.T) {
    // Define implementation capabilities
    supportedFunctions := []string{
        "function:parse", 
        "function:make-objects", 
        "function:get-string",
    }
    supportedFeatures := []string{"feature:dotted-keys"}
    behaviorChoices := map[string]string{
        "line_endings": "behavior:crlf-normalize",
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
            "function:make-objects",
            "function:get-string",
        ];
        
        let supported_features = vec!["feature:dotted-keys"];
        
        let behavior_choices = HashMap::from([
            ("line_endings", "behavior:crlf-normalize"),
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
    "function:make-objects", 
    "function:get-string"
]

SUPPORTED_FEATURES = ["feature:dotted-keys"]

BEHAVIOR_CHOICES = {
    "line_endings": "behavior:crlf-normalize"
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
@pytest.mark.ccl_function("parse", "make-objects")
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
    'function:make-objects',
    'function:get-string'
  ];
  
  const supportedFeatures = ['feature:dotted-keys'];
  
  const behaviorChoices = {
    line_endings: 'behavior:crlf-normalize'
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

### Phase 2: Object Construction  
- **Target**: Add `function:make-objects`
- **Tests**: ~80 tests
- **Goal**: Convert flat entries to nested objects
- **Features**: Consider adding `feature:dotted-keys`
- **Time**: 2-3 days

### Phase 3: Typed Access
- **Target**: Add `function:get-string`, `function:get-int`, `function:get-bool`
- **Tests**: ~120 tests  
- **Goal**: Type-safe value extraction
- **Time**: 1-2 days

### Phase 4: Processing Functions
- **Target**: Add `function:filter`, `function:compose`, `function:expand-dotted`
- **Tests**: ~150 tests
- **Goal**: Advanced entry manipulation
- **Features**: Consider adding `feature:comments`
- **Time**: 3-4 days

### Phase 5: Full Implementation
- **Target**: Add `function:pretty-print`, remaining features
- **Tests**: All 167 tests
- **Goal**: Complete CCL specification compliance
- **Time**: 2-3 days

## Troubleshooting

### High Test Failure Rate
- Check that your `supported_functions` list matches your actual implementation
- Verify behavioral choices match your implementation's behavior
- Use `just stats` to see test distribution

### Conflicting Test Results
- Review `conflicts` arrays in failing tests
- Ensure behavioral choices are consistent across your implementation
- Check for variant tag conflicts (`variant:proposed-behavior` vs `variant:reference-compliant`)

### Missing Features
- Use feature tags to identify what's needed: `feature:comments`, `feature:unicode`, etc.
- Implement features incrementally, updating supported lists as you go
- Check the tag migration guide in `docs/tag-migration.json` for implementation hints

## Reference

- **Schema**: `tests/schema.json` - Complete tag validation rules
- **Migration Guide**: `docs/tag-migration.json` - Tag mapping and examples  
- **Statistics**: Run `just stats` for current test breakdown
- **Examples**: See language-specific implementation patterns above