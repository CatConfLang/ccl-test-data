# CCL Test Filtering with Typed Fields

This document explains the type-safe filtering approach used in the CCL test suite's generated flat format, enabling implementations to run different subsets of tests based on feature support and compliance requirements.

## Overview

The CCL test suite uses a **typed fields architecture** with direct field access:
1. **Functions filtering** - `test.functions[]` array contains required CCL functions
2. **Features filtering** - `test.features[]` array contains optional language features
3. **Behaviors filtering** - `test.behaviors[]` array contains implementation choices
4. **Variants filtering** - `test.variants[]` array contains specification variants

## Typed Fields Architecture

### Generated Format Structure

The generated flat format provides type-safe filtering through separate arrays:

**Functions Array** (`test.functions[]`) - Required CCL functions:
- `parse` - Basic key-value parsing
- `build-hierarchy` - Object construction from flat entries
- `get-string`, `get-int`, `get-bool`, `get-float`, `get-list` - Typed access
- `filter`, `compose`, `expand-dotted` - Entry processing
- `pretty-print` - Canonical formatting

**Features Array** (`test.features[]`) - Optional language features:
- `comments` - `/=` comment syntax
- `dotted-keys` - `foo.bar.baz` key syntax
- `empty-keys` - `= value` anonymous list items
- `multiline` - Multi-line value support
- `unicode` - Unicode content handling
- `whitespace` - Complex whitespace preservation

**Behaviors Array** (`test.behaviors[]`) - Implementation choices (mutually exclusive):
- `crlf-preserve-literal` vs `crlf-normalize-to-lf`
- `tabs-preserve` vs `tabs-to-spaces`
- `strict-spacing` vs `loose-spacing`
- `boolean-strict` vs `boolean-lenient`

**Variants Array** (`test.variants[]`) - Specification variants:
- `proposed-behavior` - Proposed specification behavior
- `reference-compliant` - OCaml reference implementation behavior

### Type-Safe Filtering Patterns

#### 1. Function Requirements Filtering

Direct field access for checking required CCL functions:

```json
{
  "name": "parse_boolean_yes_get_bool",
  "input": "active = yes",
  "validation": "get_bool",
  "expected": {"value": true, "count": 1},
  "functions": ["parse", "build-hierarchy", "get-bool"],
  "features": ["comments"],
  "behaviors": ["boolean-lenient"],
  "variants": ["proposed-behavior"]
}
```

**Type-safe filtering logic:**
```javascript
// Check if implementation supports all required functions
const supportsAllFunctions = test.functions.every(fn => 
  implementedFunctions.includes(fn)
);
```

#### 2. Feature Requirements Filtering

Optional language features with direct array access:

```json
{
  "name": "comment_syntax_parse",
  "input": "key = value\n/= comment",
  "validation": "parse",
  "expected": {"entries": [{"key": "key", "value": "value"}], "count": 1},
  "functions": ["parse"],
  "features": ["comments"],
  "behaviors": [],
  "variants": []
}
```

**Type-safe filtering logic:**
```javascript
// Check if implementation supports all required features
const supportsAllFeatures = test.features.every(feature => 
  implementedFeatures.includes(feature)
);
```

#### 3. Behavior and Variant Filtering

Mutually exclusive behaviors and specification variants:

```json
{
  "name": "tab_preservation_proposed_parse",
  "input": "key = \tvalue",
  "validation": "parse",
  "functions": ["parse"],
  "features": ["whitespace"],
  "behaviors": ["tabs-preserve"],
  "variants": ["proposed-behavior"],
  "conflicts": {
    "behaviors": ["tabs-to-spaces"],
    "variants": ["reference-compliant"]
  }
}
```

**Type-safe filtering logic:**
```javascript
// Check for behavioral conflicts
const hasConflictingBehavior = test.conflicts.behaviors?.some(behavior => 
  implementationBehaviors.includes(behavior)
);
const hasConflictingVariant = test.conflicts.variants?.some(variant => 
  implementationVariants.includes(variant)
);
const isCompatible = !hasConflictingBehavior && !hasConflictingVariant;
```

## Type-Safe Filtering Examples

### Function-Based Implementation Support

```javascript
// Basic implementation - core functions only
const implementedFunctions = ["parse", "build-hierarchy", "get-string"];
const supportedTests = flatTests.filter(test => 
  test.functions.every(fn => implementedFunctions.includes(fn))
);
```

```javascript
// Enhanced implementation - includes processing functions
const implementedFunctions = [
  "parse", "build-hierarchy", "get-string", "get-int", "get-bool",
  "filter", "compose", "expand-dotted"
];
const supportedTests = flatTests.filter(test => 
  test.functions.every(fn => implementedFunctions.includes(fn))
);
```

### Feature and Behavior Filtering

```javascript
// Implementation with optional features
const implementedFeatures = ["comments", "dotted-keys"];
const featureCompatibleTests = flatTests.filter(test =>
  test.features.every(feature => implementedFeatures.includes(feature))
);

// Implementation behavior choices
const implementationBehaviors = ["crlf-normalize-to-lf", "boolean-strict"];
const implementationVariants = ["reference-compliant"];

const behaviorCompatibleTests = flatTests.filter(test => {
  const hasConflictingBehavior = test.conflicts?.behaviors?.some(b => 
    implementationBehaviors.includes(b)
  );
  const hasConflictingVariant = test.conflicts?.variants?.some(v => 
    implementationVariants.includes(v)
  );
  return !hasConflictingBehavior && !hasConflictingVariant;
});
```

### Complete Implementation Filtering

```javascript
// Conservative implementation - core functions, reference behavior
const conservativeTests = flatTests.filter(test => {
  // Only core functions
  const supportedFunctions = ["parse", "build-hierarchy", "get-string", "get-int"];
  const functionsSupported = test.functions.every(fn => 
    supportedFunctions.includes(fn)
  );
  
  // No optional features
  const noOptionalFeatures = test.features.length === 0;
  
  // Reference-compliant behavior
  const implementationBehaviors = ["crlf-normalize-to-lf", "boolean-strict"];
  const implementationVariants = ["reference-compliant"];
  const hasConflicts = test.conflicts?.behaviors?.some(b => 
    implementationBehaviors.includes(b)
  ) || test.conflicts?.variants?.some(v => 
    implementationVariants.includes(v)
  );
  
  return functionsSupported && noOptionalFeatures && !hasConflicts;
});

// Progressive implementation - enhanced functions, proposed behavior
const progressiveTests = flatTests.filter(test => {
  // Enhanced function set
  const supportedFunctions = [
    "parse", "build-hierarchy", "get-string", "get-int", "get-bool",
    "filter", "compose", "expand-dotted", "pretty-print"
  ];
  const functionsSupported = test.functions.every(fn => 
    supportedFunctions.includes(fn)
  );
  
  // Optional features supported
  const supportedFeatures = ["comments", "dotted-keys", "unicode"];
  const featuresSupported = test.features.every(feature => 
    supportedFeatures.includes(feature)
  );
  
  // Proposed behavior choices
  const implementationBehaviors = ["crlf-preserve-literal", "boolean-lenient"];
  const implementationVariants = ["proposed-behavior"];
  const hasConflicts = test.conflicts?.behaviors?.some(b => 
    implementationBehaviors.includes(b)
  ) || test.conflicts?.variants?.some(v => 
    implementationVariants.includes(v)
  );
  
  return functionsSupported && featuresSupported && !hasConflicts;
});
```

## Typed Fields Architecture Benefits

### 1. **Type-Safe Filtering**
- **Direct field access**: `test.functions[]`, `test.features[]`, `test.behaviors[]`, `test.variants[]`
- **No string parsing**: Use array methods like `.includes()`, `.every()`, `.some()`
- **Enum validation**: JSON schema ensures valid values only

### 2. **Excellent API Ergonomics**
- **Intuitive filtering**: `test.functions.includes('parse')` vs `test.meta.tags.includes('function:parse')`
- **Performance**: Direct array access faster than string parsing
- **Tooling support**: Better IDE autocomplete and type checking

### 3. **Progressive Implementation Support**
- **Minimal implementations**: Filter by core functions only
- **Enhanced implementations**: Add functions and features incrementally
- **Behavioral choices**: Clear conflict resolution patterns

### 4. **Conflict Resolution**
- **Categorized conflicts**: `test.conflicts.behaviors[]`, `test.conflicts.variants[]`
- **Precise filtering**: Filter out specific conflicting behaviors
- **Clear semantics**: Mutually exclusive choices explicitly defined

### 5. **Maintainable and Extensible**
- **Schema validation**: Ensures consistency across test suite
- **Clear semantics**: Each field has specific, well-defined purpose
- **Future-proof**: Easy to add new functions, features, behaviors

## Field Value Conventions

### Functions Array Values
- `parse` - Basic key-value parsing
- `build-hierarchy` - Object construction from flat entries
- `get-string`, `get-int`, `get-bool`, `get-float`, `get-list` - Typed access
- `filter`, `compose`, `expand-dotted` - Entry processing
- `pretty-print` - Canonical formatting

### Features Array Values
- `comments` - `/=` comment syntax support
- `dotted-keys` - `foo.bar.baz` key syntax support
- `empty-keys` - `= value` anonymous list items
- `multiline` - Multi-line value support
- `unicode` - Unicode content handling
- `whitespace` - Complex whitespace preservation

### Behaviors Array Values (Mutually Exclusive)
- Line endings: `crlf-preserve-literal` vs `crlf-normalize-to-lf`
- Tab handling: `tabs-preserve` vs `tabs-to-spaces`
- Spacing: `strict-spacing` vs `loose-spacing`
- Boolean parsing: `boolean-strict` vs `boolean-lenient`
- List access: `list-coercion-enabled` vs `list-coercion-disabled`

### Variants Array Values
- `proposed-behavior` - Proposed CCL specification behavior
- `reference-compliant` - OCaml reference implementation behavior

## Implementation Guidance

### 1. **Start with Function Filtering**
Determine what CCL functions your implementation supports:

```javascript
const implementedFunctions = ["parse", "build-hierarchy", "get-string"];
const functionTests = flatTests.filter(test => 
  test.functions.every(fn => implementedFunctions.includes(fn))
);
```

### 2. **Add Feature Support**
Include optional language features your implementation supports:

```javascript
const implementedFeatures = ["comments", "dotted-keys"];
const featureCompatibleTests = functionTests.filter(test =>
  test.features.every(feature => implementedFeatures.includes(feature))
);
```

### 3. **Configure Behavior Choices**
Choose your implementation's behavior for spec ambiguities:

```javascript
const implementationBehaviors = ["crlf-normalize-to-lf", "boolean-strict"];
const implementationVariants = ["reference-compliant"];

const compatibleTests = featureCompatibleTests.filter(test => {
  const hasConflictingBehavior = test.conflicts?.behaviors?.some(b => 
    implementationBehaviors.includes(b)
  );
  const hasConflictingVariant = test.conflicts?.variants?.some(v => 
    implementationVariants.includes(v)
  );
  return !hasConflictingBehavior && !hasConflictingVariant;
});
```

### 4. **Complete Test Runner Example**

```javascript
function getCompatibleTests(flatTests, capabilities) {
  return flatTests.filter(test => {
    // Check function support
    const functionsSupported = test.functions.every(fn => 
      capabilities.functions.includes(fn)
    );
    
    // Check feature support
    const featuresSupported = test.features.every(feature => 
      capabilities.features.includes(feature)
    );
    
    // Check for behavioral conflicts
    const hasConflictingBehavior = test.conflicts?.behaviors?.some(b => 
      capabilities.behaviors.includes(b)
    );
    const hasConflictingVariant = test.conflicts?.variants?.some(v => 
      capabilities.variants.includes(v)
    );
    
    return functionsSupported && featuresSupported && 
           !hasConflictingBehavior && !hasConflictingVariant;
  });
}

// Usage
const capabilities = {
  functions: ["parse", "build-hierarchy", "get-string", "get-int"],
  features: ["comments"],
  behaviors: ["crlf-normalize-to-lf", "boolean-strict"],
  variants: ["reference-compliant"]
};

const runnableTests = getCompatibleTests(flatTests, capabilities);
```

This typed fields architecture provides type-safe, high-performance filtering with excellent API ergonomics and clear conflict resolution patterns.