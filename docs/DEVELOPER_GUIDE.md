# Developer Guide: Extending the CCL Test Suite

This guide provides comprehensive instructions for developers working on the CCL Test Suite, including adding new tests, extending functionality, and contributing to the codebase.

## Table of Contents

- [Project Architecture](#project-architecture)
- [Development Workflow](#development-workflow)
- [Adding New Tests](#adding-new-tests)
- [Extending the Mock Implementation](#extending-the-mock-implementation)
- [Generator Development](#generator-development)
- [CLI Command Development](#cli-command-development)
- [Testing and Quality Assurance](#testing-and-quality-assurance)
- [Performance Considerations](#performance-considerations)

## Project Architecture

### Multi-Level CCL Implementation

The CCL Test Suite implements a 4-level architecture that allows for progressive implementation:

```
Level 1: Core CCL (Parse + MakeObjects)
    ↓
Level 2: Typed Access (GetString, GetInt, GetBool, GetFloat, GetList)
    ↓
Level 3: Advanced Processing (Filter, Compose, ExpandDotted)
    ↓
Level 4: Experimental Features (PrettyPrint)
```

### Package Structure

```
cmd/
├── ccl-test-runner/    # Main CLI application
└── test-reader/        # Interactive test browser

internal/
├── mock/               # Reference CCL implementation
├── generator/          # Test file generation
├── stats/              # Statistics and analytics
├── benchmark/          # Performance measurement
├── styles/             # Terminal styling
└── types/              # Common data structures

tests/                  # JSON test data files
go_tests/        # Generated Go test files
docs/                   # Documentation
```

### Structured Tagging System

All tests use structured tags for precise categorization:

- **`function:*`** - Required CCL functions (`function:parse`, `function:make-objects`)
- **`feature:*`** - Optional language features (`feature:comments`, `feature:dotted-keys`)
- **`behavior:*`** - Implementation choices (`behavior:strict-spacing`)
- **`variant:*`** - Specification variants (`variant:proposed-behavior`)

## Development Workflow

### Prerequisites

```bash
# Install Go 1.24.0 or later
go version

# Install development tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install just (build tool)
# See: https://github.com/casey/just#installation

# Clone and setup
git clone <repository>
cd ccl-test-data
just deps
```

### Essential Commands

```bash
# Development cycle
just lint                   # Format and lint (REQUIRED before commits)
just reset                  # Generate basic tests and verify passing
just test                   # Full test suite execution

# Test development
just generate               # Generate all Go tests from JSON
just test-generated         # Run generated tests only
just validate               # Validate JSON against schema

# Statistics and analysis
just stats                  # Display comprehensive statistics
just benchmark              # Performance measurement
```

### Pre-Commit Checklist

1. **Format and lint**: `just lint`
2. **Validate clean state**: `just reset` (must pass)
3. **Schema validation**: `just validate`
4. **Include generated files**: Commit updated `go_tests/` files

## Adding New Tests

### Test File Organization

Tests are organized by feature category in the `tests/` directory:

- **`api-essential-parsing.json`** - Basic Level 1 functionality
- **`api-comprehensive-parsing.json`** - Advanced parsing with edge cases
- **`api-comments.json`** - Comment syntax support
- **`api-dotted-keys.json`** - Dotted key expansion
- **`api-object-construction.json`** - Level 1 object building
- **`api-typed-access.json`** - Level 2 type-safe access
- **`api-processing.json`** - Level 3 composition and filtering
- **`api-errors.json`** - Error handling validation

### Test Structure

Each test follows this structure with **required** `count` fields:

```json
{
  "name": "descriptive_test_name",
  "input": "CCL input text",
  "validations": {
    "validation_name": {
      "count": 1,
      "expected": "expected_result"
    },
    "case_validation": {
      "count": 2,
      "cases": [
        {"args": ["arg1"], "expected": "result1"},
        {"args": ["arg2"], "expected": "result2"}
      ]
    }
  },
  "meta": {
    "tags": ["function:parse", "feature:comments"],
    "level": 1,
    "feature": "parsing",
    "conflicts": ["behavior:strict-spacing"]
  }
}
```

### Validation Types

#### Direct Validations

Test a single function call with expected output:

```json
"parse": {
  "count": 1,
  "expected": [{"key": "name", "value": "value"}]
}
```

#### Case-Based Validations

Test multiple function calls with different arguments:

```json
"get_string": {
  "count": 2,
  "cases": [
    {"args": ["key1"], "expected": "value1"},
    {"args": ["key2"], "expected": "value2"}
  ]
}
```

#### Error Validations

Test expected error conditions:

```json
"parse_error": {
  "count": 1,
  "expected_error": "parsing error message pattern"
}
```

### Structured Tags

#### Required Function Tags

Tag tests with the CCL functions they validate:

```json
"tags": [
  "function:parse",           // Level 1: Parse function
  "function:make-objects",    // Level 1: Object construction
  "function:get-string"       // Level 2: String access
]
```

#### Optional Feature Tags

Tag tests with language features they require:

```json
"tags": [
  "feature:comments",         // Comment syntax support
  "feature:dotted-keys",      // Dotted key expansion
  "feature:unicode",          // Unicode character support
  "feature:multiline"         // Multiline value support
]
```

#### Behavior Tags

Tag tests with implementation behavior they expect:

```json
"tags": [
  "behavior:strict-spacing",  // Strict whitespace handling
  "behavior:crlf-preserve"    // Preserve CRLF line endings
]
```

#### Conflict Handling

Specify mutually exclusive behaviors:

```json
"conflicts": [
  "behavior:loose-spacing"    // Conflicts with strict-spacing
]
```

### Adding a New Test

1. **Choose appropriate test file** based on feature category
2. **Write test structure** with proper validation types
3. **Add structured tags** for functions and features
4. **Include count fields** matching array lengths or case counts
5. **Validate JSON schema**: `just validate`
6. **Generate and test**: `just generate && just test-generated`
7. **Check statistics**: `just stats` to see coverage impact

Example new test:

```json
{
  "name": "unicode_key_parsing",
  "input": "配置 = 值\n键名 = 数据",
  "validations": {
    "parse": {
      "count": 2,
      "expected": [
        {"key": "配置", "value": "值"},
        {"key": "键名", "value": "数据"}
      ]
    }
  },
  "meta": {
    "tags": ["function:parse", "feature:unicode"],
    "level": 1,
    "feature": "parsing"
  }
}
```

## Extending the Mock Implementation

### Implementation Levels

The mock implementation in `internal/mock/ccl.go` provides working CCL functionality across multiple levels.

#### Level 1: Raw Parsing

```go
func (c *CCL) Parse(input string) ([]Entry, error) {
    var entries []Entry
    lines := strings.Split(input, "\n")
    
    for lineNum, line := range lines {
        // Trim whitespace
        line = strings.TrimSpace(line)
        
        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "/=") {
            continue
        }
        
        // Parse key-value pairs
        if parts := strings.SplitN(line, "=", 2); len(parts) == 2 {
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            entries = append(entries, Entry{Key: key, Value: value})
        } else {
            return nil, fmt.Errorf("line %d: invalid syntax", lineNum+1)
        }
    }
    
    return entries, nil
}
```

#### Level 3: Object Construction

```go
func (c *CCL) MakeObjects(entries []Entry) map[string]interface{} {
    result := make(map[string]interface{})
    
    for _, entry := range entries {
        setNestedValue(result, strings.Split(entry.Key, "."), entry.Value)
    }
    
    return result
}
```

#### Level 2: Typed Access

```go
func (c *CCL) GetString(obj map[string]interface{}, path []string) (string, error) {
    value, err := c.getValue(obj, path)
    if err != nil {
        return "", err
    }
    
    switch v := value.(type) {
    case string:
        return v, nil
    default:
        return fmt.Sprintf("%v", v), nil
    }
}
```

### Adding New Functions

1. **Define function signature** matching test expectations
2. **Implement core functionality** with proper error handling
3. **Add structured tags** to relevant tests
4. **Update generator templates** if needed
5. **Test implementation**: `just reset && just test-generated`

### Error Handling Patterns

Provide detailed error messages with context:

```go
func (c *CCL) validateInput(input string) error {
    if input == "" {
        return fmt.Errorf("empty input provided")
    }
    
    // Additional validation logic
    return nil
}

func (c *CCL) parseLineWithContext(line string, lineNum int) (Entry, error) {
    if !strings.Contains(line, "=") {
        return Entry{}, fmt.Errorf("line %d: missing '=' separator in '%s'", 
            lineNum+1, line)
    }
    
    // Parse logic with detailed errors
}
```

## Generator Development

### Template System

The generator uses Go templates to create test files from JSON data:

```go
// internal/generator/templates.go
const testTemplate = `package {{.PackageName}}_test

import (
    "testing"
    "github.com/ccl-test-data/test-runner/internal/mock"
)

func TestGenerated{{.TestName}}(t *testing.T) {
    ccl := mock.New()
    
    {{range .Validations}}
    // {{.Name}} validation
    {{.GeneratedCode}}
    {{end}}
}
`
```

### Adding New Validation Types

1. **Define validation structure** in test JSON
2. **Add template generation logic**:

```go
func generateValidation(validation Validation) string {
    switch validation.Type {
    case "new_validation_type":
        return generateNewValidationType(validation)
    default:
        return generateDefaultValidation(validation)
    }
}
```

3. **Update test schema** in `tests/schema.json`
4. **Test generation**: `just generate && just test-generated`

### Filtering and Selection

The generator supports sophisticated filtering:

```go
type Options struct {
    SkipDisabled bool      // Skip disabled feature tags
    SkipTags     []string  // Additional tags to skip
    RunOnly      []string  // Generate only these tags
}
```

### Progressive Implementation

Generate tests for specific implementation levels:

```bash
# Level 1 only
./ccl-test-runner generate --run-only function:parse

# Levels 1 and 3
./ccl-test-runner generate --run-only function:parse,function:make-objects

# Exclude advanced features
./ccl-test-runner generate --skip-tags feature:unicode,feature:multiline
```

## CLI Command Development

### Command Structure

CLI commands follow the urfave/cli/v2 pattern:

```go
&cli.Command{
    Name:        "command-name",
    Aliases:     []string{"alias"},
    Usage:       "Short description",
    Description: "Long description with examples",
    Action:      actionFunction,
    Flags: []cli.Flag{
        &cli.StringFlag{
            Name:    "flag-name",
            Aliases: []string{"f"},
            Value:   "default",
            Usage:   "Flag description",
        },
    },
},
```

### Adding New Commands

1. **Define command structure** in `cmd/ccl-test-runner/main.go`
2. **Implement action function**:

```go
func newCommandAction(ctx *cli.Context) error {
    // Get flag values
    flagValue := ctx.String("flag-name")
    
    // Implement command logic
    styles.Status("🚀", "Starting command execution...")
    
    // Handle errors appropriately
    if err := doWork(flagValue); err != nil {
        return fmt.Errorf("command failed: %w", err)
    }
    
    styles.Success("✅ Command completed successfully")
    return nil
}
```

3. **Add to command list** in main app definition
4. **Test command functionality**
5. **Update documentation**

### Output Formatting

Use the styles package for consistent output:

```go
import "github.com/ccl-test-data/test-runner/internal/styles"

// Status messages
styles.Status("🔍", "Analyzing test files...")
styles.Success("✅ Analysis complete")
styles.Error("❌ Failed to process: %v", err)
styles.Warning("⚠️  Performance regression detected")

// Information
styles.Info("Processing %d test files", count)
styles.InfoLite("Found %d assertions", assertions)

// Command display
styles.Command("gotestsum --format testname ./go_tests/...")
```

### Flag Handling

Support multiple output formats and filtering options:

```go
func processFlags(ctx *cli.Context) error {
    format := ctx.String("format")
    levels := ctx.IntSlice("levels")
    features := ctx.StringSlice("features")
    
    // Validate flag combinations
    if format == "json" && len(levels) > 0 {
        return fmt.Errorf("JSON format doesn't support level filtering")
    }
    
    // Process flags appropriately
    return nil
}
```

## Testing and Quality Assurance

### Test Categories

1. **Unit Tests**: Individual function testing
2. **Integration Tests**: Full workflow validation
3. **Generated Tests**: JSON-based test execution
4. **Performance Tests**: Benchmark validation

### Quality Gates

Before commits:

1. **Linting**: `just lint` (golangci-lint + gofmt)
2. **Schema Validation**: `just validate` (JSON schema compliance)
3. **Clean State**: `just reset` (all enabled tests pass)
4. **Full Suite**: `just test` (comprehensive validation)

### Mock Implementation Testing

The mock implementation should pass all enabled tests:

```bash
# Test current implementation
just generate-mock     # Generate tests for current mock capabilities
just test-mock         # Run mock-specific tests

# Progressive testing
just test-level1       # Test Level 1 functionality only
just test-level3       # Test up to Level 3
```

### Error Scenarios

Test error handling thoroughly:

```json
{
  "name": "invalid_syntax_error",
  "input": "invalid syntax here",
  "validations": {
    "parse_error": {
      "count": 1,
      "expected_error": "line 1: missing '=' separator"
    }
  },
  "meta": {
    "tags": ["function:parse", "feature:error-handling"],
    "level": 1,
    "feature": "error-handling"
  }
}
```

## Performance Considerations

### Benchmark Integration

The benchmark package tracks performance metrics:

```go
// Start benchmark
tracker := benchmark.NewTracker()
tracker.StartBenchmark("test-generation")

// Perform operation
err := generator.GenerateAll()

// End benchmark
result := tracker.EndBenchmark("test-generation")
```

### Optimization Guidelines

1. **Memory Management**: Reuse data structures where possible
2. **I/O Efficiency**: Batch file operations
3. **Parsing Performance**: Optimize hot paths in mock implementation
4. **Test Generation**: Use object pooling for frequent allocations

### Performance Monitoring

```bash
# Run performance benchmarks
just benchmark

# Compare against historical results
just benchmark --compare benchmarks/historical.json --threshold 10.0
```

### Memory Profiling

For detailed performance analysis:

```go
// Add to benchmark code
import _ "net/http/pprof"
import "runtime/pprof"

// Profile memory usage
f, _ := os.Create("memory.prof")
defer f.Close()
pprof.WriteHeapProfile(f)
```

## Best Practices

### Code Organization

1. **Package Separation**: Keep concerns separated by package
2. **Interface Usage**: Define interfaces for testability
3. **Error Wrapping**: Use `fmt.Errorf("context: %w", err)`
4. **Documentation**: Add godoc comments for exported functions

### Test Design

1. **Progressive Complexity**: Start simple, add edge cases
2. **Clear Naming**: Test names should describe what they validate
3. **Proper Tagging**: Use structured tags consistently
4. **Count Accuracy**: Ensure count fields match actual test assertions

### Git Workflow

1. **Feature Branches**: Use descriptive branch names
2. **Atomic Commits**: Each commit should represent one logical change
3. **Generated Files**: Always commit updated go_tests/ files
4. **Clean History**: Rebase to maintain clean commit history

### Documentation

1. **API Documentation**: Keep docs/API.md current
2. **Examples**: Include practical usage examples
3. **Migration Guides**: Document breaking changes
4. **Performance Notes**: Document performance characteristics

## Troubleshooting

### Common Issues

1. **Schema Validation Fails**: Check JSON structure against `tests/schema.json`
2. **Generated Tests Fail**: Verify mock implementation supports required functions
3. **Tag Conflicts**: Ensure mutually exclusive behaviors are properly marked
4. **Performance Regression**: Check benchmark results for optimization opportunities

### Debug Commands

```bash
# Detailed test output
just test-verbose

# JSON validation with details
go run ./cmd/validate-schema tests/*.json

# Statistics with full breakdown
just stats --format json | jq '.'

# Generate specific test subset
./ccl-test-runner generate --run-only function:parse --skip-disabled=false
```

### Development Environment

Recommended development setup:

```bash
# Editor configuration
# .vscode/settings.json
{
  "go.lintTool": "golangci-lint",
  "go.formatTool": "goimports",
  "go.testFlags": ["-v", "-race"]
}

# Git hooks
# .git/hooks/pre-commit
#!/bin/sh
just lint && just reset
```

This guide provides a comprehensive foundation for extending and contributing to the CCL Test Suite. For specific implementation details, refer to the API documentation and existing code examples.