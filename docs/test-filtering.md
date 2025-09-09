# CCL Test Filtering Design Pattern

This document explains the architectural design pattern used for test filtering in the CCL test suite, enabling implementations to run different subsets of tests based on feature support and compliance requirements.

## Overview

The CCL test suite uses a dual filtering system:
1. **Feature-based filtering** - What capabilities does an implementation support?
2. **Behavioral filtering** - Which interpretation of spec ambiguities does an implementation follow?

## Design Pattern

### Feature Categories vs Tags

**Feature categories** (`meta.feature`) represent functional groupings:
- `parsing` - Basic CCL parsing functionality
- `typed-parsing` - Type-aware value extraction  
- `flexible-boolean-parsing` - Optional enhanced boolean parsing
- `crlf-normalization` - Optional line ending normalization
- `object-construction` - Hierarchical structure creation
- `dotted-keys` - Dotted key expansion support
- etc.

**Tags** (`meta.tags`) provide detailed filtering and behavioral modifiers:
- Content tags: `basic`, `multiline`, `whitespace`, `tabs`, etc.
- Compliance tags: `proposed`, `reference_compliant` 
- Feature requirement tags: `flexible-boolean-parsing`, `strict-boolean-parsing`

### Filtering Patterns

#### 1. Optional Feature Pattern

When a tag name matches a feature category name, it indicates the test requires that optional feature:

```json
{
  "name": "parse_boolean_yes",
  "validations": {"get_bool": {"args": ["active"], "expected": true}},
  "meta": {
    "tags": ["typed_parsing", "boolean", "needs-flexible-boolean-parsing"],
    "feature": "flexible-boolean-parsing"
  }
}
```

**Filtering logic:**
- **Capability filtering**: `feature == "flexible-boolean-parsing"` 
- **Test runners**: Only run if implementation supports flexible boolean parsing
- **Clear intent**: The `needs-flexible-boolean-parsing` tag makes the requirement obvious

#### 2. Baseline Feature Pattern  

Tests using baseline behavior have descriptive tags but keep the base feature category:

```json
{
  "name": "parse_boolean_yes_strict", 
  "validations": {"get_bool": {"args": ["active"], "error": true}},
  "meta": {
    "tags": ["typed_parsing", "boolean", "uses-strict-boolean-parsing"],
    "feature": "typed-parsing"
  }
}
```

**Filtering logic:**
- **Capability filtering**: `feature == "typed-parsing"` (baseline requirement)
- **Behavioral filtering**: `tags.includes("uses-strict-boolean-parsing")` 
- **Clear intent**: The `uses-` prefix clarifies this describes behavior, not requirements

#### 3. Spec Ambiguity Pattern

For spec ambiguities, both interpretations are valid:

```json
{
  "name": "tab_preservation_proposed",
  "meta": {
    "tags": ["whitespace", "tabs", "proposed-behavior"],
    "feature": "parsing"
  }
},
{
  "name": "tab_preservation_ocaml_reference", 
  "meta": {
    "tags": ["whitespace", "tabs", "reference-compliant-behavior"],
    "feature": "parsing"  
  }
}
```

**Filtering logic:**
- **Capability filtering**: Both need `feature == "parsing"` (same baseline)
- **Compliance filtering**: Choose `proposed` vs `reference_compliant` based on implementation approach

## Filtering Examples

### Implementation Feature Support

```javascript
// Basic implementation - only baseline features
const basicFeatures = ["parsing", "typed-parsing", "object-construction"];
const supportedTests = tests.filter(test => 
  basicFeatures.includes(test.meta.feature)
);
```

```javascript
// Enhanced implementation - includes optional features
const enhancedFeatures = [
  "parsing", "typed-parsing", "object-construction",
  "flexible-boolean-parsing", "crlf-normalization"  
];
const supportedTests = tests.filter(test => 
  enhancedFeatures.includes(test.meta.feature)
);
```

### Compliance-Based Filtering

```javascript
// CCL proposed behavior
const proposedTests = tests.filter(test =>
  !test.meta.tags.includes("reference-compliant-behavior")
);

// OCaml reference compliant behavior  
const referenceTests = tests.filter(test =>
  !test.meta.tags.includes("proposed-behavior") || 
   test.meta.tags.includes("reference-compliant-behavior")
);
```

### Combined Filtering

```javascript
// Strict baseline implementation following OCaml reference
const strictTests = tests.filter(test => {
  const hasBasicFeature = ["parsing", "typed-parsing", "object-construction"]
    .includes(test.meta.feature);
  const isReference = test.meta.tags.includes("reference-compliant-behavior") ||
    !test.meta.tags.includes("proposed-behavior");
  
  return hasBasicFeature && isReference;
});

// Flexible implementation with proposed enhancements
const flexibleTests = tests.filter(test => {
  const supportedFeatures = [
    "parsing", "typed-parsing", "object-construction",
    "flexible-boolean-parsing", "crlf-normalization"
  ];
  const isProposed = test.meta.tags.includes("proposed-behavior") ||
    !test.meta.tags.includes("reference-compliant-behavior");
    
  return supportedFeatures.includes(test.meta.feature) && isProposed;
});
```

## Architecture Benefits

### 1. **Clear Separation of Concerns**
- **Features** = "What can this implementation do?"
- **Tags** = "How does this implementation behave?"

### 2. **Intuitive Filtering**
- Filter by `feature` for capability requirements
- Filter by `tags` for behavioral variants
- Combine filters for precise test selection

### 3. **Flexible Implementation Support**
- Minimal implementations: Run baseline tests only
- Enhanced implementations: Add optional features incrementally  
- Spec compliance: Choose interpretation of ambiguous cases

### 4. **Maintainable Test Organization**
- Optional features clearly identified by tag-feature matching
- Baseline behavior uses descriptive tags with base features
- Spec ambiguities use compliance tags with same base features

### 5. **Future-Proof Design**
- Easy to add new optional features
- Clear pattern for handling spec evolution
- Supports multiple compliance interpretations

## Tag Conventions

### Compliance Tags
- `proposed-behavior` - CCL proposed behavior for spec ambiguities
- `reference-compliant-behavior` - OCaml reference implementation behavior  

### Optional Feature Requirements Tags
- `needs-flexible-boolean-parsing` - Requires flexible boolean parsing feature
- `needs-crlf-normalization` - Requires CRLF normalization feature
- `needs-{feature}` pattern for future optional features

### Baseline Behavior Description Tags
- `uses-strict-boolean-parsing` - Uses baseline boolean parsing (true/false only)
- `uses-strict-line-ending-parsing` - Uses baseline line ending parsing (preserve as-is)
- `uses-{behavior}` pattern for baseline behavior descriptions

### Content Tags
- `basic`, `multiline`, `whitespace`, `tabs`, `unicode`, etc.
- Standard descriptive tags for test content categorization

## Implementation Guidance

### 1. **Start with Feature Filtering**
Determine what capabilities your implementation supports:

```javascript
const myFeatures = ["parsing", "object-construction", "typed-parsing"];
const myTests = tests.filter(t => myFeatures.includes(t.meta.feature));
```

### 2. **Add Compliance Filtering**  
Choose your interpretation of spec ambiguities:

```javascript
// For proposed CCL behavior
const myTests = tests.filter(t => 
  myFeatures.includes(t.meta.feature) &&
  !t.meta.tags.includes("reference-compliant-behavior")
);

// For OCaml reference compliance  
const myTests = tests.filter(t =>
  myFeatures.includes(t.meta.feature) &&
  !t.meta.tags.includes("proposed-behavior")
);
```

### 3. **Consider Optional Features**
Add enhanced capabilities incrementally:

```javascript
// Add flexible boolean parsing support
myFeatures.push("flexible-boolean-parsing");

// Now tests with needs-flexible-boolean-parsing will be included
const enhancedTests = tests.filter(t => myFeatures.includes(t.meta.feature));
```

This filtering design enables precise test selection while maintaining clear architectural boundaries between capability requirements and behavioral choices.