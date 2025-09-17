# CCL Test Suite Technical Overview

A comprehensive technical overview of the CCL (Categorical Configuration Language) Test Suite for software engineers and implementers.

## What is This Project?

The CCL Test Suite is a **comprehensive language-agnostic testing framework** for CCL implementations. It provides:

- **452 test assertions** across **167 tests** in JSON format
- **Progressive implementation support** via 4-level CCL architecture
- **Feature-based tagging** for precise test selection
- **Go-based tooling** for test generation and validation
- **Reference implementation** demonstrating CCL patterns

## Quick Start for Developers

### Prerequisites
```bash
# Required: Go 1.23+ and just (task runner)
go version                    # Check Go installation
just --version               # Check just installation

# Clone and setup
git clone <repository>
cd ccl-test-data
just deps                    # Install dependencies
```

### Essential Commands
```bash
# Development workflow
just reset                   # Generate basic tests, ensure clean state (REQUIRED for commits)
just test                    # Run comprehensive test suite
just stats                   # Show test coverage statistics

# Progressive development
just test --functions core            # Test core functions only
just test --functions core,typed     # Test core + typed access functions
just generate --run-only function:parse,function:build_hierarchy
```

## Technical Architecture

### System Design
```
Source Tests (JSON) → Generator → [Flat JSON + Go Tests] → Mock CCL → Results
     ↑                  ↓                   ↓               ↓
Maintainable     Type-Safe Format    Reference Impl    Validation
```

### Core Components

| Component | Purpose | Key Files |
|-----------|---------|-----------|
| **Source Tests** | Human-maintainable test definitions | `source_tests/api_*.json` |
| **Generator** | Transform source → implementation formats | `internal/generator/` |
| **Mock CCL** | Reference implementation | `internal/mock/ccl.go` |
| **CLI Tools** | Test runner and analysis | `cmd/ccl-test-runner/` |
| **Schemas** | JSON validation | `schemas/*.json` |

### Data Flow
1. **Source Format**: Multi-validation tests with structured metadata
2. **Generated Format**: Flat, type-safe tests (1:N transformation)
3. **Go Tests**: Generated test functions using mock implementation
4. **Validation**: Execute tests and report results

## Function Groups

The CCL Test Suite supports progressive implementation across function groups:

### Core CCL Functions
```go
// Essential CCL functions - what users expect
Parse(input string) ([]Entry, error)              // "key = value" → entries
ParseValue(input string) ([]Entry, error)         // Indentation-aware parsing
BuildHierarchy([]Entry) map[string]interface{}    // entries → nested objects
```
**Goal**: Complete Core CCL functionality (text → hierarchical configuration)

### Typed Access Functions
```go
// Type-safe value extraction
GetString(obj, path) (string, error)
GetInt(obj, path) (int, error)
GetBool(obj, path) (bool, error)
GetFloat(obj, path) (float64, error)
GetList(obj, path) ([]interface{}, error)
```
**Goal**: Type-safe access to configuration values

### Processing Functions
```go
// Entry manipulation and composition
Filter([]Entry) []Entry                   // Remove entries by criteria
Combine([]Entry) []Entry                  // Merge and compose entries
```
**Goal**: Advanced configuration processing

### Experimental Functions
```go
// Experimental features
ExpandDotted([]Entry) []Entry            // Handle dotted key expansion (experimental)
```
**Goal**: Experimental functionality for testing new features

### Formatting Functions
```go
// Standardized output generation
CanonicalFormat(map[string]interface{}) string    // Generate formatted CCL (also known as PrettyPrint)
```
**Goal**: Round-trip formatting and standardization

## Feature-Based Testing

### Structured Tags
Tests use structured tags for precise categorization:

```json
{
  "meta": {
    "tags": [
      "function:parse",              // Required CCL function
      "feature:comments",            // Optional language feature
      "behavior:strict_spacing"      // Implementation choice
    ],
    "level": 1,
    "conflicts": ["behavior:loose_spacing"]
  }
}
```

### Tag Categories

| Category | Purpose | Examples |
|----------|---------|----------|
| **function:*** | Required CCL functions | `parse`, `build_hierarchy`, `get_string` |
| **feature:*** | Optional language features | `comments`, `experimental_dotted_keys`, `unicode` |
| **behavior:*** | Implementation choices | `crlf_preserve_literal`, `boolean_strict` |
| **variant:*** | Specification variants | `proposed_behavior`, `reference_compliant` |

### Progressive Implementation Strategy
```bash
# 1. Start with core parsing
just generate --run-only function:parse,function:parse_value

# 2. Add object construction (complete core)
just generate --run-only function:parse,function:parse_value,function:build_hierarchy

# 3. Add typed access
just generate --run-only function:parse,function:parse_value,function:build_hierarchy,function:get_string

# 4. Skip experimental features during development
just generate --skip-tags function:expand_dotted
```

## Test Data Format

### Dual Format Architecture

#### Source Format (Maintainable)
```json
{
  "name": "basic_parsing_test",
  "input": "key = value\nother = data",
  "validations": {
    "parse": {
      "count": 2,
      "expected": [
        {"key": "key", "value": "value"},
        {"key": "other", "value": "data"}
      ]
    },
    "get_string": {
      "count": 1,
      "cases": [{"args": ["key"], "expected": "value"}]
    }
  },
  "meta": {
    "tags": ["function:parse", "function:get_string"],
    "level": 1
  }
}
```

#### Generated Format (Implementation-Friendly)
```json
{
  "name": "basic_parsing_test_parse",
  "input": "key = value\nother = data",
  "validation": "parse",
  "expected": {
    "count": 2,
    "entries": [
      {"key": "key", "value": "value"},
      {"key": "other", "value": "data"}
    ]
  },
  "functions": ["parse"],
  "features": [],
  "level": 1
}
```

### Counted Assertions
All tests include **required `count` fields** for self-validation:
- Enables precise assertion counting
- Prevents silent test failures
- Supports test complexity measurement

## Implementation Examples

### Basic CCL Implementation Pattern
```go
type CCL struct{}

func (c *CCL) Parse(input string) ([]Entry, error) {
    var entries []Entry
    lines := strings.Split(input, "\n")

    for lineNum, line := range lines {
        line = strings.TrimSpace(line)

        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "/=") {
            continue
        }

        // Parse key-value pairs
        if parts := strings.SplitN(line, "=", 2); len(parts) == 2 {
            entries = append(entries, Entry{
                Key:   strings.TrimSpace(parts[0]),
                Value: strings.TrimSpace(parts[1]),
            })
        } else {
            return nil, fmt.Errorf("line %d: invalid syntax", lineNum+1)
        }
    }

    return entries, nil
}

func (c *CCL) BuildHierarchy(entries []Entry) map[string]interface{} {
    result := make(map[string]interface{})

    for _, entry := range entries {
        setNestedValue(result, strings.Split(entry.Key, "."), entry.Value)
    }

    return result
}
```

### Test Integration Pattern
```go
func TestCCLImplementation(t *testing.T) {
    // Load flat format tests (type-safe, implementation-friendly)
    tests := loadFlatTests("generated_tests/")

    // Filter for supported functionality
    supportedTests := filterTests(tests, []string{"parse", "build_hierarchy"})

    ccl := NewCCL()

    for _, test := range supportedTests {
        switch test.Validation {
        case "parse":
            entries, err := ccl.Parse(test.Input)
            assert.NoError(t, err)
            assert.Equal(t, test.Expected.Entries, entries)
            assert.Equal(t, test.Expected.Count, len(entries))

        case "build_hierarchy":
            entries, _ := ccl.Parse(test.Input)
            obj := ccl.BuildHierarchy(entries)
            assert.Equal(t, test.Expected.Object, obj)
        }
    }
}
```

## Development Workflow

### Standard Development Cycle
```bash
# 1. Start development session
just reset                   # Ensure clean state

# 2. Make changes to source tests or implementation
# 3. Generate tests and validate
just generate               # Generate Go tests from JSON
just test                   # Run comprehensive test suite

# 4. Before committing
just lint                   # Format and lint (REQUIRED)
just reset                  # Verify clean passing state
```

### Adding New Test Cases
```bash
# 1. Add test to appropriate source_tests/api_*.json file
# 2. Include proper structured tags and count fields
# 3. Validate and generate
just validate               # Check JSON schema compliance
just generate               # Generate flat format and Go tests
just test-generated         # Run generated tests

# 4. Analyze impact
just stats                  # View updated test coverage
```

### Repository State Management
- **`just reset`**: Generates only basic tests that mock implementation can pass
- **Required for commits**: Ensures repository stays in clean, passing state
- **CI/CD Integration**: Automated validation pipeline

## Performance Characteristics

### Scalability
- **Template Caching**: Pre-compiled templates for efficiency
- **Streaming Processing**: Handle large test sets efficiently

## Quality Assurance

### Validation Pipeline
1. **JSON Schema Validation**: Strict structure compliance
2. **Generated Test Validation**: Template output verification
3. **Mock Implementation Testing**: Reference implementation validation
4. **Integration Testing**: End-to-end workflow testing

### Quality Gates
- **Pre-commit**: `just lint && just reset` (both must pass)
- **CI Pipeline**: Full test suite validation
- **Release**: Comprehensive validation + documentation updates

## Integration Strategies

### For CCL Implementers
1. **Use Generated Format**: Implement against flat JSON format for type safety
2. **Progressive Implementation**: Start with core functions, add features incrementally
3. **Feature Selection**: Use structured tags to filter relevant tests
4. **Validation Pattern**: Compare results against expected outputs with count verification

### For Language Bindings
1. **Test Data Reuse**: Use same JSON test data across language implementations
2. **Consistent Behavior**: Ensure identical behavior across implementations
3. **Feature Parity**: Track feature implementation across languages

### For CI/CD Integration
```bash
# Validation pipeline
ccl-test-runner validate     # Schema validation
ccl-test-runner generate     # Test generation
ccl-test-runner test --format json > results.json
ccl-test-runner stats --format json > coverage.json
```

## Documentation Resources

| Document | Purpose | Audience |
|----------|---------|----------|
| **[ARCHITECTURE.md](docs/ARCHITECTURE.md)** | System design and component architecture | Engineers, Architects |
| **[DEVELOPER_GUIDE.md](docs/DEVELOPER_GUIDE.md)** | Development workflow and contribution guidelines | Contributors |
| **[CLI_REFERENCE.md](docs/CLI_REFERENCE.md)** | Complete command-line tool documentation | Users, DevOps |
| **[API.md](docs/API.md)** | Complete API reference | Implementers |

## External Resources

- **[CCL Documentation](https://ccl.tylerbutler.com)** - Complete language specification
- **[Original CCL Blog Post](https://chshersh.com/blog/2025-01-06-the-most-elegant-configuration-language.html)** - Language introduction
- **[OCaml Reference Implementation](https://github.com/chshersh/ccl)** - Canonical implementation

This technical overview provides engineers with everything needed to understand, use, and contribute to the CCL Test Suite.