# CCL Test Suite

> [!NOTE]
> This is the **official JSON test suite** for CCL implementations across all programming languages, featuring comprehensive **feature-based tagging** for precise test selection and progressive implementation support.

Language-agnostic test suite for the Categorical Configuration Language (CCL) with **feature-based tagging** for precise test selection. Each test specifies which CCL functions to validate and uses structured tags to enable progressive implementation.

## What is CCL?

> [!TIP]
> New to CCL? Start with the **[Specification Summary](https://ccl.tylerbutler.com/specification-summary)** for a complete overview, then check the **[Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)** for quick implementation guidance.

For comprehensive CCL documentation, see the **[CCL Documentation](https://ccl.tylerbutler.com)** which includes:

- **[Specification Summary](https://ccl.tylerbutler.com/specification-summary)** - Complete language specification
- **[Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)** - Quick syntax guide
- **[Parsing Algorithm](https://ccl.tylerbutler.com/parsing-algorithm)** - Implementation guide
- **[Mathematical Theory](https://ccl.tylerbutler.com/theory)** - Theoretical foundations

### Original Sources

- [CCL Blog Post](https://chshersh.com/blog/2025-01-06-the-most-elegant-configuration-language.html) - Original specification by Dmitrii Kovanikov
- [OCaml Reference Implementation](https://github.com/chshersh/ccl) - Canonical implementation

## Test Suite

This repository contains the **official JSON test suite** for CCL implementations across all programming languages.

### Key Features

> [!IMPORTANT]
> All tests use a **counted format** with required `count` fields for precise validation verification. Each validation declares exactly how many assertions it represents.

✅ **Feature-based tagging** - Structured tags for precise test selection (`function:*`, `feature:*`, `behavior:*`, `variant:*`)  
✅ **Direct API mapping** - Each validation maps to a specific API function  
✅ **Multi-level testing** - Tests declare expected outputs for different parsing levels  
✅ **Conflict resolution** - Automatic handling of mutually exclusive behaviors  
✅ **Progressive implementation** - Clear path from minimal parsing to full features  
✅ **Simple test runners** - Direct iteration over `validations` object keys  
✅ **Assertion counting** - Required explicit counts for validation verification  
✅ **Self-documenting** - Validation names explain what's being tested  
✅ **452 test assertions** - Comprehensive coverage across all CCL features

### Quick Start

> [!TIP]
> Use `just reset` before committing to ensure all enabled tests pass. This maintains repository in a clean, stable state for CI and development.

```bash
# Clone the test suite
git clone <this-repo>
cd ccl-test-data

# Install dependencies and run tests
just deps
just test

# Generate tests for mock implementation development
just generate-mock
just test-mock

# Set repository to clean, passing state (required for commits)
just reset  # alias for dev-basic
```

### Test Files

The test suite is organized by feature category:

#### Core Parsing
- **`tests/api-essential-parsing.json`** - Basic parsing functionality for rapid prototyping
- **`tests/api-comprehensive-parsing.json`** - Thorough parsing with edge cases and whitespace variations

#### Advanced Processing
- **`tests/api-processing.json`** - Entry composition, merging, and advanced processing
- **`tests/api-comments.json`** - Comment syntax and filtering functionality

#### Object Construction
- **`tests/api-object-construction.json`** - Converting flat entries to nested objects
- **`tests/api-dotted-keys.json`** - Dotted key expansion and conflict resolution

#### Type System
- **`tests/api-typed-access.json`** - Type-aware value extraction with smart inference

#### Error Handling
- **`tests/api-errors.json`** - Error handling validation

### Using the Test Suite

> [!IMPORTANT]
> **Counted Format Required**: All validations must include a `count` field that matches the number of expected results. This enables precise assertion counting and self-validating test suites.

#### Test Format Structure

```json
{
  "name": "basic_multi_level_test",
  "input": "database.host = localhost",
  "validations": {
    "parse": {
      "count": 1,
      "expected": [{"key": "database.host", "value": "localhost"}]
    },
    "make_objects": {
      "count": 1,
      "expected": {"database": {"host": "localhost"}}
    },
    "get_string": {
      "count": 1,
      "cases": [
        {
          "args": ["database.host"],
          "expected": "localhost"
        }
      ]
    }
  },
  "meta": {
    "tags": ["function:parse", "function:make-objects", "function:get-string", "feature:dotted-keys"],
    "level": 3,
    "feature": "dotted-keys"
  }
}
```

## Feature-Based Test Selection

> [!TIP]
> **Progressive Implementation**: Start with `function:parse` only, then gradually add `function:make-objects`, typed access functions (`function:get-string`), and advanced features (`feature:comments`, `feature:dotted-keys`) as your implementation matures.

The test suite uses **structured tags** to enable precise test selection based on implementation capabilities:

### Tag Categories

#### Function Tags (`function:*`) - Required CCL functions:
- `function:parse` - Basic key-value parsing (Level 1)
- `function:filter` - Entry filtering (Level 2) 
- `function:compose` - Entry composition (Level 2)
- `function:expand-dotted` - Dotted key expansion (Level 2)
- `function:make-objects` - Object construction (Level 3)
- `function:get-string`, `function:get-int`, `function:get-bool`, `function:get-float`, `function:get-list` - Typed access (Level 4)
- `function:pretty-print` - Formatting (Level 5)

#### Feature Tags (`feature:*`) - Optional language features:
- `feature:comments` - `/=` comment syntax
- `feature:dotted-keys` - `foo.bar.baz` key syntax
- `feature:empty-keys` - `= value` anonymous list items
- `feature:multiline` - Multi-line value support
- `feature:unicode` - Unicode content handling
- `feature:whitespace` - Complex whitespace preservation

#### Behavior Tags (`behavior:*`) - Implementation choices (mutually exclusive):
- `behavior:crlf-preserve` vs `behavior:crlf-normalize` - Line ending handling
- `behavior:tabs-preserve` vs `behavior:tabs-to-spaces` - Tab handling
- `behavior:strict-spacing` vs `behavior:loose-spacing` - Whitespace sensitivity

#### Variant Tags (`variant:*`) - Specification variants:
- `variant:proposed-behavior` - Proposed specification behavior
- `variant:reference-compliant` - OCaml reference implementation behavior

### Test Selection Examples

> [!NOTE]
> **Implementation Strategy**: Use these examples as templates for your test runner configuration. Start minimal and expand gradually as features are implemented.

#### Minimal Implementation (Parse only)
```json
{"supported_tags": ["function:parse"]}
```

#### Basic Implementation (Parse + Objects + Typed Access)
```json
{
  "supported_functions": ["function:parse", "function:make-objects", "function:get-string"],
  "supported_features": ["feature:dotted-keys"],
  "behavior_choices": {"line_endings": "behavior:crlf-normalize"}
}
```

#### Advanced Implementation (All functions, optional features)
```json
{
  "supported_functions": ["function:*"],
  "optional_features": ["feature:comments", "feature:unicode"],
  "skip_variants": ["variant:proposed-behavior"]
}
```

### Conflict Resolution

> [!WARNING]
> **Mutually Exclusive Behaviors**: Tests with conflicting behavioral tags are automatically excluded based on your implementation choices. Choose one behavior per category to avoid test conflicts.

Tests with conflicting behaviors are automatically excluded:

```json
{
  "name": "crlf_preservation_test",
  "meta": {
    "tags": ["function:parse", "behavior:crlf-preserve"],
    "conflicts": ["behavior:crlf-normalize"]
  }
}
```

If your implementation chooses `behavior:crlf-normalize`, tests tagged with `behavior:crlf-preserve` are automatically skipped.

### Feature Field vs Feature Tags

> [!NOTE]
> **Two Types of Features**: Don't confuse the organizational `feature` field (for test suite grouping) with requirement-based `feature:*` tags (for implementation filtering).

There are two distinct "feature" concepts that serve different purposes:

#### `feature` field (organizational) - File/suite categorization:
- `"feature": "parsing"` - This test belongs to parsing test suite
- `"feature": "dotted-keys"` - This test belongs to dotted-keys test suite  
- `"feature": "object-construction"` - This test belongs to object construction suite

#### `feature:*` tags (requirement-based) - Implementation requirements:
- `"feature:dotted-keys"` - This test requires dotted key support to run
- `"feature:comments"` - This test requires comment parsing support
- `"feature:unicode"` - This test requires Unicode handling

**Example:** A test in the "parsing" suite (`"feature": "parsing"`) might still have `"feature:comments"` tag if it tests comment parsing, indicating that implementations without comment support should skip it.

#### Usage Guidelines:
- **File organization**: Use `feature` field for test suite grouping
- **Test filtering**: Use `feature:*` tags for implementation-based filtering

### Test Runner Implementation Example

```javascript
for (const [validationType, validation] of Object.entries(test.validations)) {
  switch (validationType) {
    case 'parse':
      if (validation.error) {
        expect(() => parse(test.input)).toThrow(validation.error_message);
      } else {
        const actual = parse(test.input);
        expect(actual).toEqual(validation.expected);
        expect(actual.length).toBe(validation.count); // Verify count
      }
      break;
    case 'make_objects':
      const entries = parse(test.input);
      const objects = makeObjects(entries);
      expect(objects).toEqual(validation.expected);
      break;
    case 'get_string':
      const ccl = makeObjects(parse(test.input));
      if (validation.error) {
        expect(() => getString(ccl, ...validation.args)).toThrow();
      } else {
        validation.cases.forEach(testCase => {
          if (testCase.error) {
            expect(() => getString(ccl, ...testCase.args)).toThrow();
          } else {
            const value = getString(ccl, ...testCase.args);
            expect(value).toBe(testCase.expected);
          }
        });
        expect(validation.cases.length).toBe(validation.count); // Verify count
      }
      break;
  }
}
```

## Assertion Counting

> [!IMPORTANT]
> **Self-Validating Tests**: The `count` field enables test runners to verify they're executing the expected number of assertions, preventing silent test failures and ensuring comprehensive coverage.

All validations use the **counted format** with required `count` fields:

### Count Field Guidelines
- **For array results** (`parse`, `filter`, `expand_dotted`): `count` = number of items in `expected` array
- **For object results** (`make_objects`): `count` = typically 1 (single object)
- **For typed access**: `count` = number of test cases in `cases` array
- **For empty results**: `count` = 0 (e.g., empty input parsing)

### Benefits
- **Explicit counting**: Each validation declares exactly how many assertions it represents
- **Self-validating**: Test runners can verify `count` matches actual array lengths
- **Test complexity tracking**: Enables precise measurement of implementation complexity

## Go Test Runner

> [!TIP]
> **Quick Development Cycle**: Use `just dev-mock` for rapid prototyping or `just reset` to maintain a clean repository state with only passing tests enabled.

This repository includes a comprehensive Go-based test runner for CCL implementations:

### Available Commands

```bash
# Basic usage
just build          # Build the test runner binary
just generate       # Generate all Go test files
just test           # Run all tests
just list           # List available test packages

# Mock implementation development
just generate-mock  # Generate tests for mock implementation
just test-mock      # Run tests suitable for mock implementation
just dev-mock       # Full development cycle for mock

# Level-specific testing
just test-level1    # Run only Level 1 tests
just test-level2    # Run only Level 2 tests
just test-level3    # Run only Level 3 tests
just test-level4    # Run only Level 4 tests

# Feature-specific testing
just test-comments  # Run comment-related tests
just test-parsing   # Run parsing tests
just test-objects   # Run object construction tests

# Utilities
just stats          # Show test generation statistics
just validate       # Validate test files against schema
just clean          # Clean generated files
```

### Mock Implementation

> [!NOTE]
> **Learning Resource**: The mock implementation serves as both a working example and a foundation for development. It demonstrates proper test integration patterns and API structure.

The repository includes a basic Level 1 mock CCL implementation for testing and development:

- **Location**: `internal/mock/ccl.go`
- **Features**: Basic key-value parsing, comment handling, empty input support
- **Usage**: Demonstrates test integration patterns and API structure

### Repository State Management

> [!WARNING]
> **Critical for CI/CD**: The repository **must** be in a clean, passing state before commits. Use `just reset` to ensure all enabled tests pass and maintain stable CI builds.

The repository should be maintained in a clean, passing state. Use these commands to ensure all enabled tests pass:

```bash
# Standard repository state (all tests should pass)
just reset  # alias for dev-basic

# Or run the steps manually:
just generate-level1  # Generate only basic Level 1 tests
just test-level1      # Run Level 1 tests (all should pass)
```

**This is the required state for commits and CI.** The `dev-basic` command generates only the most essential tests (`basic`, `essential-parsing`, `empty` tags) and skips advanced features that would fail in the current mock implementation. This ensures:

- **Clean commits**: All enabled tests pass before committing
- **Stable CI**: Continuous integration runs pass consistently
- **Development foundation**: Solid base for CCL implementation work

## Documentation

### Test Suite Schema
- **[Schema Technical Reference](docs/generated-schema.md)** - Complete auto-generated field documentation
- **[Schema Implementation Guide](docs/schema-reference.md)** - Practical usage examples and patterns

### Test Suite Architecture
- **[Test Architecture](docs/test-architecture.md)** - How to use this test suite
- **[Test Filtering](docs/test-filtering.md)** - Advanced test filtering patterns

### General Implementation Guidance
- **[Implementation Guide](https://ccl.tylerbutler.com/implementing-ccl)** - Complete CCL implementation guide
- **[Test Architecture](https://ccl.tylerbutler.com/test-architecture)** - General testing concepts

## Contributing

> [!IMPORTANT]
> **Test Quality Standards**: All new tests must use the counted format, include proper metadata tags, and pass JSON schema validation before being accepted.

When adding test cases:

1. **Add to appropriate JSON file** by feature level and category
2. **Include descriptive name and metadata** with proper tag classification
3. **Use counted format** with appropriate `count` values matching result arrays
4. **Validate JSON structure** with `just validate` before submitting
5. **Update test counts** in documentation and ensure `just stats` reflects changes

## Validation

> [!TIP]
> **Development Workflow**: Run `just validate` before committing changes to catch JSON schema violations early. Use `just dev-basic` for rapid iteration during development.

```bash
# Validate test suite structure
just validate

# Run schema validation
go run cmd/validate-schema/main.go tests/api-*.json

# Generate and run all tests
just dev

# Quick development cycle for basic features
just dev-basic
```

## Test Statistics

> [!NOTE]
> **Comprehensive Coverage**: The test suite provides **452 assertions** across **167 tests**, ensuring thorough validation of CCL implementations from basic parsing to advanced features.

The test suite provides comprehensive coverage with **452 assertions** across **167 tests**:

```bash
# View detailed statistics
just stats
```

### Current Breakdown

**📊 Overall Statistics:**
- **167 total tests** with **452 assertions** across **10 files**
- **21 mutually exclusive tests** with behavioral/variant conflicts
- **11 CCL functions** from basic parsing to advanced formatting
- **6 language features** (comments, dotted-keys, unicode, etc.)
- **3 behavioral choices** (CRLF, tabs, spacing handling)
- **2 specification variants** (proposed vs reference behavior)

**📚 Level Distribution:**
- **Level 1**: 54 tests (basic parsing)
- **Level 2**: 30 tests (processing)
- **Level 3**: 27 tests (object construction)  
- **Level 4**: 56 tests (typed access)

**⚙️ Function Coverage:**
- `function:parse`: 132 tests (most essential)
- `function:make-objects`: 66 tests
- `function:get-*`: 38 tests (typed access)
- `function:pretty-print`: 24 tests
- `function:compose`: 12 tests
- Other functions: 35 tests

This test suite ensures consistent CCL behavior across all language implementations with precise control over which features to test.
