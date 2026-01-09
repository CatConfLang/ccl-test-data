# CCL Test Data

Language-agnostic JSON test suite for [CCL (Categorical Configuration Language)](https://ccl.tylerbutler.com) implementations.

**218 tests | 470 assertions | 15 files**

## Quick Start

Download test files from [GitHub Releases](https://github.com/catconflang/ccl-test-data/releases) or use the generated flat format directly.

```bash
# Clone for development
git clone https://github.com/catconflang/ccl-test-data.git
cd ccl-test-data
```

## Test Format

Tests are in `generated_tests/*.json` (flat format, one assertion per test):

```json
{
  "name": "basic_parsing_parse",
  "input": "key = value",
  "validation": "parse",
  "expected": {
    "count": 1,
    "entries": [{"key": "key", "value": "value"}]
  },
  "functions": ["parse"],
  "features": [],
  "behaviors": [],
  "conflicts": {}
}
```

## Test Filtering

Filter tests based on your implementation's capabilities:

```javascript
function shouldRun(test, myCapabilities) {
  // Must support all required functions
  if (!test.functions.every(f => myCapabilities.functions.includes(f))) return false;

  // Skip tests that conflict with your behavior choices
  if (test.conflicts?.behaviors?.some(b => myCapabilities.behaviors.includes(b))) return false;

  return true;
}
```

### Functions

| Function                                                     | Description                      |
| ------------------------------------------------------------ | -------------------------------- |
| `parse`                                                      | Basic key-value parsing          |
| `build_hierarchy`                                            | Object construction from entries |
| `get_string`, `get_int`, `get_bool`, `get_float`, `get_list` | Typed value access               |
| `filter`, `compose`                                          | Entry processing                 |
| `canonical_format`, `round_trip`                             | Formatting and validation        |

### Behaviors

Implementation choices that affect test compatibility:

| Group           | Options                                           |
| --------------- | ------------------------------------------------- |
| CRLF handling   | `crlf_normalize_to_lf`, `crlf_preserve_literal`   |
| Tab handling    | `tabs_as_whitespace`, `tabs_as_content`           |
| Boolean parsing | `boolean_strict`, `boolean_lenient`               |
| List coercion   | `list_coercion_enabled`, `list_coercion_disabled` |

> **Note:** Behaviors are not inherently mutually exclusive. A test can require multiple behaviors. Use the `conflicts` field to determine incompatible combinations per-test.

See [test-selection-guide.md](docs/test-selection-guide.md) for complete filtering documentation.

## Go Library

For Go implementations, import the test infrastructure directly:

```go
import (
    "github.com/catconflang/ccl-test-data/config"
    "github.com/catconflang/ccl-test-data/loader"
)

cfg := config.ImplementationConfig{
    SupportedFunctions: []config.CCLFunction{
        config.FunctionParse,
        config.FunctionBuildHierarchy,
    },
}

testLoader := loader.NewTestLoader("path/to/ccl-test-data", cfg)
tests, _ := testLoader.LoadAllTests(loader.LoadOptions{
    Format:     loader.FormatFlat,
    FilterMode: loader.FilterCompatible,
})
```

## Resources

- [CCL Documentation](https://ccl.tylerbutler.com) - Language specification and guides
- [Implementing CCL](https://ccl.tylerbutler.com/implementing-ccl/) - Parser implementation guide
- [AI Quickstart](https://ccl.tylerbutler.com/ai-quickstart/) - LLM-friendly CCL reference
- [Original Specification](https://chshersh.com/blog/2025-01-06-the-most-elegant-configuration-language.html) - CCL blog post by Dmitrii Kovanikov

## Development

See [DEV.md](DEV.md) for development setup, contributing guidelines, and release process.

## Schema

JSON schemas for validation:

- `schemas/source-format.json` - Human-maintainable source format
- `schemas/generated-format.json` - Machine-friendly generated format
