# Test Development Guide for Copilot Agents

This guide provides specialized instructions for agents working on test suite development in the CCL Test Suite repository.

## Prerequisites

Before working on test development, ensure you've read:
- [onboarding.md](onboarding.md) - General repository understanding
- [test-selection-guide.md](/docs/test-selection-guide.md) - Test filtering and selection

## Test Suite Organization

### Test File Structure

```
source_tests/
├── core/                                  # Core functionality tests
│   ├── api_comments.json                 # Comment syntax
│   ├── api_core_ccl_hierarchy.json       # Object construction
│   ├── api_core_ccl_integration.json     # Integration tests
│   ├── api_core_ccl_parsing.json         # Core parsing
│   ├── api_edge_cases.json               # Edge cases
│   ├── api_errors.json                   # Error handling
│   ├── api_advanced_processing.json      # Advanced features
│   ├── api_list_access.json              # List operations
│   ├── api_proposed_behavior.json        # Proposed behaviors
│   ├── api_reference_compliant.json      # Reference compliance
│   ├── api_typed_access.json             # Typed access functions
│   ├── property_round_trip.json          # Round-trip properties
│   └── property_algebraic.json           # Algebraic properties
└── experimental/
    └── api_experimental.json             # Experimental features
```

### Choosing the Right Test File

| Category | File | Use When |
|----------|------|----------|
| Basic parsing | `api_core_ccl_parsing.json` | Testing fundamental parse functionality |
| Object construction | `api_core_ccl_hierarchy.json` | Testing `build_hierarchy` and nested objects |
| Type access | `api_typed_access.json` | Testing `get_string`, `get_int`, etc. |
| Comments | `api_comments.json` | Testing comment syntax and filtering |
| Error handling | `api_errors.json` | Testing error cases and messages |
| Edge cases | `api_edge_cases.json` | Testing unusual inputs and boundaries |
| Processing | `api_advanced_processing.json` | Testing `filter`, `compose`, `merge` |
| Lists | `api_list_access.json` | Testing list operations |
| Proposed features | `api_proposed_behavior.json` | Testing proposed specification behavior |
| Reference compliance | `api_reference_compliant.json` | Testing OCaml reference compatibility |
| Experimental | `api_experimental.json` | Testing unstable/experimental features |

## Test Structure

### Basic Test Format

```json
{
  "name": "unique_descriptive_name",
  "input": "key = value",
  "tests": [
    {
      "function": "parse",
      "expect": [
        {"key": "key", "value": "value"}
      ]
    }
  ],
  "features": [],
  "behaviors": [],
  "variants": [],
  "conflicts": {}
}
```

### Complete Test Example

```json
{
  "name": "multiline_value_with_comments",
  "input": "key = \"\"\"line1\nline2\"\"\" # comment",
  "tests": [
    {
      "function": "parse",
      "expect": [
        {"key": "key", "value": "line1\nline2"}
      ]
    },
    {
      "function": "get_string",
      "args": ["key"],
      "expect": "line1\nline2"
    },
    {
      "function": "filter",
      "args": ["comments"],
      "expect": [
        {"key": "key", "value": "line1\nline2"}
      ]
    }
  ],
  "features": ["multiline", "comments"],
  "behaviors": ["crlf_normalize_to_lf"],
  "variants": [],
  "conflicts": {
    "behaviors": ["crlf_preserve_literal"]
  }
}
```

## Test Metadata

### Features Array

Optional language features that the test requires:

| Feature | Description | Example Use |
|---------|-------------|-------------|
| `comments` | Requires comment support | Tests with `#` comments |
| `empty_keys` | Requires empty key support | `"" = value` |
| `multiline` | Requires multiline strings | `"""text"""` |
| `unicode` | Requires Unicode support | UTF-8 characters |
| `whitespace` | Tests whitespace handling | Significant spaces/tabs |

### Behaviors Array

Implementation choices (mutually exclusive pairs):

| Category | Options | Description |
|----------|---------|-------------|
| Boolean parsing | `boolean_strict`, `boolean_lenient` | How to parse boolean values |
| Line endings | `crlf_preserve_literal`, `crlf_normalize_to_lf` | CRLF handling |
| Tabs | `tabs_preserve`, `tabs_to_spaces` | Tab character handling |
| Spacing | `strict_spacing`, `loose_spacing` | Whitespace sensitivity |
| List coercion | `list_coercion_enabled`, `list_coercion_disabled` | Single value to list conversion |

### Variants Array

Specification variant choices:

| Variant | Description |
|---------|-------------|
| `proposed_behavior` | Enhanced/flexible interpretation |
| `reference_compliant` | Strict OCaml reference compatibility |

### Conflicts Object

Specify mutually exclusive options:

```json
{
  "conflicts": {
    "behaviors": ["crlf_preserve_literal"],
    "variants": ["reference_compliant"],
    "features": ["experimental_syntax"],
    "functions": ["deprecated_parse"]
  }
}
```

## Adding a New Test

### Step-by-Step Workflow

```bash
# 1. Navigate to repository
cd /home/runner/work/ccl-test-data/ccl-test-data

# 2. Choose appropriate test file
# Edit source_tests/core/api_*.json

# 3. Add your test to the "tests" array

# 4. Validate JSON structure
just validate

# 5. Generate derived files
just generate

# 6. Run tests to verify
just test

# 7. Check statistics
just stats

# 8. Pre-commit validation
just lint
just reset
```

### Example: Adding a Basic Parsing Test

1. Open `source_tests/core/api_core_ccl_parsing.json`

2. Add to the `tests` array:
```json
{
  "name": "simple_key_value",
  "input": "name = John",
  "tests": [
    {
      "function": "parse",
      "expect": [
        {"key": "name", "value": "John"}
      ]
    }
  ],
  "features": [],
  "behaviors": []
}
```

3. Validate and generate:
```bash
just validate
just generate
just test
```

### Example: Adding a Multi-Function Test

```json
{
  "name": "nested_object_access",
  "input": "user.name = Alice\nuser.age = 30",
  "tests": [
    {
      "function": "parse",
      "expect": [
        {"key": "user.name", "value": "Alice"},
        {"key": "user.age", "value": "30"}
      ]
    },
    {
      "function": "build_hierarchy",
      "expect": {
        "user": {
          "name": "Alice",
          "age": "30"
        }
      }
    },
    {
      "function": "get_string",
      "args": ["user.name"],
      "expect": "Alice"
    },
    {
      "function": "get_int",
      "args": ["user.age"],
      "expect": 30
    }
  ],
  "features": [],
  "behaviors": []
}
```

## Test Generation Process

### Understanding the Pipeline

```
source_tests/           → generated_tests/      → go_tests/
(Human-friendly)         (Flat format)           (Go test files)
- Multiple validations   - One test per          - Executable tests
  per test                validation             - Go test framework
- Rich metadata         - Tags for filtering    - Must be committed
```

### Generation Commands

```bash
# Generate everything (recommended)
just generate

# Generate only flat JSON
just generate-flat

# Generate only Go tests
just generate-go

# Generate with filtering
just generate --run-only function:parse
just generate --skip-tags multiline,error
```

### Verifying Generation

After generation, verify:

```bash
# Check generated flat format
ls -l generated_tests/

# Check generated Go tests
ls -l go_tests/

# Run generated tests
just test

# View statistics
just stats
```

## Function-Specific Test Patterns

### Parse Function

Tests the core parsing function:

```json
{
  "function": "parse",
  "expect": [
    {"key": "key1", "value": "value1"},
    {"key": "key2", "value": "value2"}
  ]
}
```

### ParseDedented Function

Tests indentation-normalized parsing (dedenting):

```json
{
  "function": "parse_indented",
  "expect": [
    {"key": "database", "value": ""},
    {"key": "host", "value": "localhost"}
  ]
}
```

This function calculates the common leading whitespace prefix and strips it from all lines, treating the dedented keys as top-level.

### BuildHierarchy Function

Tests object construction:

```json
{
  "function": "build_hierarchy",
  "expect": {
    "section": {
      "key": "value"
    }
  }
}
```

### Typed Access Functions

Test type-safe value extraction:

```json
{
  "function": "get_string",
  "args": ["key"],
  "expect": "value"
}
```

```json
{
  "function": "get_int",
  "args": ["count"],
  "expect": 42
}
```

```json
{
  "function": "get_bool",
  "args": ["enabled"],
  "expect": true
}
```

```json
{
  "function": "get_float",
  "args": ["pi"],
  "expect": 3.14
}
```

```json
{
  "function": "get_list",
  "args": ["items"],
  "expect": ["item1", "item2"]
}
```

### Processing Functions

Test entry manipulation:

```json
{
  "function": "filter",
  "args": ["comments"],
  "expect": [
    {"key": "key", "value": "value"}
  ]
}
```

```json
{
  "function": "compose",
  "args": [/* additional entries */],
  "expect": [/* composed result */]
}
```

### Error Testing

Test error conditions:

```json
{
  "function": "parse",
  "expect_error": "Expected '=' after key"
}
```

## Test Quality Standards

### Naming Conventions

- Use snake_case: `nested_object_with_comments`
- Be descriptive: `multiline_value_with_unicode_chars`
- Include key features: `empty_key_with_strict_spacing`

### Input Guidelines

- Use realistic CCL syntax
- Test edge cases explicitly
- Include comments explaining unusual cases
- Keep inputs minimal but complete

### Expected Output Guidelines

- Match exact output format from mock implementation
- Include all relevant fields
- Test both success and error cases
- Verify type conversions are correct

### Metadata Guidelines

- Include all relevant features
- Specify behavioral choices explicitly
- Document conflicts clearly
- Use appropriate variants

## Common Test Patterns

### Testing Comments

```json
{
  "name": "inline_comment",
  "input": "key = value # this is a comment",
  "tests": [
    {
      "function": "parse",
      "expect": [{"key": "key", "value": "value"}]
    },
    {
      "function": "filter",
      "args": ["comments"],
      "expect": [{"key": "key", "value": "value"}]
    }
  ],
  "features": ["comments"]
}
```

### Testing Multiline Values

```json
{
  "name": "multiline_string",
  "input": "text = \"\"\"line1\nline2\nline3\"\"\"",
  "tests": [
    {
      "function": "parse",
      "expect": [{"key": "text", "value": "line1\nline2\nline3"}]
    }
  ],
  "features": ["multiline"]
}
```

### Testing Unicode

```json
{
  "name": "unicode_characters",
  "input": "greeting = Hello 世界 🌍",
  "tests": [
    {
      "function": "parse",
      "expect": [{"key": "greeting", "value": "Hello 世界 🌍"}]
    }
  ],
  "features": ["unicode"]
}
```

### Testing Nested Objects

```json
{
  "name": "deeply_nested",
  "input": "a.b.c.d = value",
  "tests": [
    {
      "function": "parse",
      "expect": [{"key": "a.b.c.d", "value": "value"}]
    },
    {
      "function": "build_hierarchy",
      "expect": {
        "a": {
          "b": {
            "c": {
              "d": "value"
            }
          }
        }
      }
    }
  ]
}
```

## Troubleshooting Test Development

### Validation Fails

**Error**: JSON schema validation fails

**Solution**:
1. Check schema at `schemas/source-format.json`
2. Verify required fields: `name`, `input`, `tests`
3. Ensure `tests` array is not empty
4. Validate metadata arrays contain valid values

### Generation Produces Unexpected Results

**Error**: Generated tests don't match expectations

**Solution**:
1. Review generated flat JSON in `generated_tests/`
2. Check that function names match schema
3. Verify metadata is correctly formed
4. Run `just generate` again after fixes

### Tests Fail After Adding

**Error**: New tests fail when run

**Solution**:
1. Verify mock implementation supports the function
2. Check that expected output format is correct
3. Test with simpler input first
4. Review behavioral requirements

### Statistics Don't Update

**Error**: `just stats` doesn't show new tests

**Solution**:
1. Ensure generation completed successfully
2. Check that tests are in correct JSON files
3. Verify JSON syntax is valid
4. Run `just generate` again

## Best Practices for Test Development

### DO:
✅ Start with simple tests and add complexity gradually
✅ Test one feature at a time when possible
✅ Include both positive and negative test cases
✅ Use descriptive names that explain what's being tested
✅ Add comments in input CCL to explain edge cases
✅ Verify tests pass before committing
✅ Run `just stats` to see coverage impact

### DON'T:
❌ Add tests without validating JSON schema
❌ Skip the generation step
❌ Forget to specify required metadata
❌ Test multiple unrelated features in one test
❌ Use ambiguous or misleading names
❌ Commit without running `just reset`
❌ Add tests for unimplemented functions without noting them

## Pre-Commit Checklist for Test Development

Before committing test changes:

```bash
# 1. Validate JSON structure
just validate

# 2. Generate all derived files
just generate

# 3. Lint code (if any Go changes)
just lint

# 4. Run tests
just reset

# 5. Check statistics
just stats

# 6. Review changes
git status
git diff source_tests/
git diff go_tests/

# 7. Stage all changes (including generated files)
git add source_tests/ go_tests/ generated_tests/

# 8. Commit with descriptive message
git commit -m "Add tests for [feature description]"
```

## Advanced Topics

### Progressive Test Implementation

When implementing a new CCL feature:

1. **Phase 1**: Add basic tests
   ```bash
   just generate --run-only function:parse
   just test
   ```

2. **Phase 2**: Add intermediate tests
   ```bash
   just generate --run-only function:parse,function:build-hierarchy
   just test
   ```

3. **Phase 3**: Add complete tests
   ```bash
   just generate
   just test-all
   ```

### Test Filtering Examples

```bash
# Run only parsing tests
just test --levels 1

# Run tests for specific features
just test --features comments,parsing

# Skip experimental tests
just generate --skip-tags experimental

# Run only specific functions
just generate --run-only function:parse,function:get-string
```

### Batch Test Addition

When adding multiple related tests:

1. Add all tests to source files
2. Run validation once: `just validate`
3. Generate once: `just generate`
4. Test once: `just test`
5. Commit as a single logical change

## References

- **[test-selection-guide.md](/docs/test-selection-guide.md)** - Test selection and filtering
- **[schema-reference.md](/docs/schema-reference.md)** - Complete schema documentation
- **[test-architecture.md](/docs/test-architecture.md)** - Test suite design
- **Source Schema**: `schemas/source-format.json`
- **Generated Schema**: `schemas/generated-format.json`
