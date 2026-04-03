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

**Informational tags** describing which language features a test exercises. Features are for reporting, not filtering.

| Feature | Description | Example |
|---------|-------------|---------|
| `comments` | `/=` comment syntax | `/= This is a comment` |
| `empty_keys` | Anonymous list items | `= item1\n= item2` |
| `multiline` | Multi-line values | `description = Line 1\nLine 2` |
| `unicode` | Unicode content | `name = José` |
| `whitespace` | Complex whitespace handling | Preserving tabs, spaces |

**Use for reporting:**
```javascript
// Find failing tests by feature to identify gaps
const unicodeGaps = failingTests.filter(t => t.features.includes('unicode'));
console.log(`Could support unicode by fixing ${unicodeGaps.length} tests`);
```

### Behaviors Array (`test.behaviors[]`)

**Implementation choices** describing how a test expects certain operations to behave. Behaviors fall into two categories:

**Parsing behaviors** (affect core parsing):
| Behavior | Description |
|----------|-------------|
| `crlf_preserve_literal` | Preserve `\r` characters in values (default) |
| `crlf_normalize_to_lf` | Normalize CRLF to LF during parsing |
| `tabs_as_content` | Tabs are content characters (never indentation, never trimmed) |
| `tabs_as_whitespace` | Tabs are whitespace (count for indentation, get trimmed) |
| `indent_spaces` | Use spaces for printed indentation |
| `indent_tabs` | Use tabs for printed indentation |

**API-specific behaviors** (affect individual functions like `get_bool`):
| Behavior | Description |
|----------|-------------|
| `boolean_strict` | Only accept "true"/"false" |
| `boolean_lenient` | Also accept "yes"/"no", "1"/"0" (superset of strict) |
| `list_coercion_enabled` | Single values coerce to single-item lists |
| `list_coercion_disabled` | Require explicit list syntax |
| `array_order_insertion` | Preserve insertion order |
| `array_order_lexicographic` | Sort elements lexicographically |

> **Note:** Only behaviors with populated `conflicts` fields are mutually exclusive. Some behaviors are supersets (e.g., `boolean_lenient` accepts everything `boolean_strict` does, plus more).

### Variants Array (`test.variants[]`) - Temporary Disambiguation

Direct array containing **specification variant interpretations** for areas where the CCL specification is currently ambiguous or evolving. These exist to handle cases where the spec doesn't clearly define behavior, allowing both the OCaml reference implementation approach and proposed enhanced approaches to coexist during spec evolution.

| Variant | Description | Status |
|---------|-------------|--------|
| `proposed_behavior` | Enhanced/flexible interpretation of ambiguous spec areas | Proposed for future spec |
| `reference_compliant` | Strict compatibility with OCaml reference implementation | Current baseline |

**Filtering:**
```javascript
// Skip tests that conflict with your variant choice
if (test.conflicts?.variants?.some(v => v === myVariant)) skipTest(test);
```

**Important:** These variant values are **temporary disambiguation mechanisms**. Once the CCL specification owners clarify these ambiguities, the variant system will be eliminated and these choices will be converted to specific behaviors (e.g., `multiline-flexible` vs `multiline-strict`).

## Behavior vs Variant

**Behaviors** are stable implementation choices:
- Technical decisions (CRLF handling, boolean parsing)
- Won't change based on spec evolution

**Variants** are temporary spec disambiguation:
- Where OCaml reference and proposed behavior differ
- Will be eliminated once spec is clarified

### Example Test Structure

```json
{
  "behaviors": ["crlf_normalize_to_lf"],
  "conflicts": { "behaviors": ["crlf_preserve_literal"] }
}
```

## Implementation Strategies

### Filtering Logic

1. **Filter by functions** - Skip tests requiring unimplemented functions
2. **Filter by conflicts** - Skip tests whose `conflicts.behaviors` or `conflicts.variants` include your choices
3. **Report by features** - Use feature tags to identify gaps

### Example

```javascript
const my = {
  functions: ['parse', 'build_hierarchy', 'get_string'],
  behaviors: ['crlf_normalize_to_lf', 'boolean_lenient'],
  variants: ['proposed_behavior']
};

tests.filter(test => {
  // Must support all required functions
  if (!test.functions.every(f => my.functions.includes(f))) return false;
  // Must not conflict with our choices
  if (test.conflicts?.behaviors?.some(b => my.behaviors.includes(b))) return false;
  if (test.conflicts?.variants?.some(v => my.variants.includes(v))) return false;
  return true;
});
```

## Language-Specific Examples

### Go

```go
func shouldSkip(test Test, myFuncs, myBehaviors []string) bool {
    // Skip if we don't support a required function
    for _, f := range test.Functions {
        if !contains(myFuncs, f) {
            return true
        }
    }
    // Skip if test conflicts with our behaviors
    for _, b := range test.Conflicts.Behaviors {
        if contains(myBehaviors, b) {
            return true
        }
    }
    return false
}
```

### JavaScript

```javascript
function shouldSkip(test, myFuncs, myBehaviors) {
  if (!test.functions.every(f => myFuncs.includes(f))) return true;
  if (test.conflicts?.behaviors?.some(b => myBehaviors.includes(b))) return true;
  return false;
}
```

## Available Functions

**Core Parsing:**
- `parse` - Basic key-value parsing
- `parse_indented` - Indentation-aware parsing
- `build_hierarchy` - Object construction from flat entries

**Typed Access:**
- `get_string`, `get_int`, `get_bool`, `get_float`, `get_list` - Type-safe value extraction

**Processing:**
- `filter`, `compose`, `expand_dotted` - Entry manipulation

**Formatting/IO:**
- `canonical_format`, `load`, `round_trip` - Output and validation

## Troubleshooting

### High Test Failure Rate
- Check that your supported functions list matches your implementation
- Verify your behavior choices match your implementation
- Use `just stats` to see test distribution

### Conflicting Test Results
- Review `conflicts` arrays in failing tests
- Ensure behavioral choices are consistent

### Reporting on Features
Use feature tags to identify gaps:
```javascript
const gaps = failingTests.filter(t => t.features.includes('unicode'));
console.log(`Could support unicode by fixing ${gaps.length} tests`);
```

## Reference

- **Schema**: `schemas/source-format.json`
- **Statistics**: `just stats`