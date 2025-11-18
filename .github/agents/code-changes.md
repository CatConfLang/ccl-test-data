# Code Changes Guide for Copilot Agents

This guide provides specialized instructions for agents making code changes to the CCL Test Suite repository.

## Prerequisites

Before making code changes, ensure you've read:
- [onboarding.md](onboarding.md) - General repository understanding

## Repository Code Structure

### Go Package Organization

```
ccl-test-data/
├── cmd/                            # Command-line applications
│   ├── ccl-test-runner/           # Main test runner CLI
│   │   └── main.go
│   └── test-reader/               # Interactive test browser
│       └── main.go
├── internal/                       # Internal packages
│   ├── mock/                      # Reference CCL implementation
│   │   └── ccl.go
│   ├── generator/                 # Test generation logic
│   │   ├── flat.go               # Flat format generation
│   │   └── go_tests.go           # Go test generation
│   ├── stats/                     # Statistics collection
│   │   └── stats.go
│   ├── config/                    # Configuration management
│   │   └── runner_config.go
│   └── types/                     # Common data structures
│       └── types.go
└── go.mod                         # Go module definition
```

## Making Code Changes

### Standard Workflow

```bash
# 1. Make your code changes

# 2. Format code
go fmt ./...

# 3. Lint code (REQUIRED)
just lint

# 4. Build to check for compile errors
go build ./...

# 5. Run tests
just test

# 6. Verify clean state
just reset

# 7. Commit changes
git add .
git commit -m "Descriptive commit message"
```

## Key Code Areas

### 1. Mock CCL Implementation (`internal/mock/ccl.go`)

**Purpose**: Reference implementation of CCL for testing

**Key Functions**:
- `Parse(input string) ([]Entry, error)` - Parse CCL text to flat entries (basic lexical parsing)
- `ParseDedented(input string) ([]Entry, error)` - Parse with indentation normalization (strips common leading whitespace)
- `BuildHierarchy(entries []Entry) (map[string]any, error)` - Build nested objects (calls ParseDedented on values)
- `GetString(data map[string]any, key string) (string, error)` - Get string value
- `GetInt(data map[string]any, key string) (int, error)` - Get integer value
- `GetBool(data map[string]any, key string) (bool, error)` - Get boolean value
- `GetFloat(data map[string]any, key string) (float64, error)` - Get float value
- `GetList(data map[string]any, key string) ([]any, error)` - Get list value
- `Filter(entries []Entry, filterType string) ([]Entry, error)` - Filter entries
- `PrettyPrint(entries []Entry) string` - Format entries as CCL

**When to modify**:
- Adding new CCL function support
- Fixing parsing bugs
- Improving error messages
- Adding new features to CCL

**Modification checklist**:
```bash
# 1. Edit internal/mock/ccl.go

# 2. Update tests if needed
# (Tests are in go_tests/ but generated from source_tests/)

# 3. Lint and test
just lint
just test

# 4. Update documentation if API changed
# Edit docs/MOCK_IMPLEMENTATION.md
# Edit docs/API.md if public API changed
```

### 2. Test Generator (`internal/generator/`)

**Purpose**: Generate flat JSON and Go test files from source tests

**Key Files**:
- `flat.go` - Generates flat format from source format
- `go_tests.go` - Generates Go test files from flat format

**When to modify**:
- Changing test generation logic
- Adding new test formats
- Improving test organization
- Adding test filtering capabilities

**Modification checklist**:
```bash
# 1. Edit generator files

# 2. Test generation pipeline
just generate-flat
just generate-go

# 3. Verify generated tests
just test

# 4. Check generated files are valid
ls -l generated_tests/
ls -l go_tests/
```

### 3. CLI Tools (`cmd/`)

**Purpose**: Command-line interfaces for test management

**ccl-test-runner commands**:
- `generate` - Generate test files
- `test` - Run tests
- `stats` - Show statistics
- `validate` - Validate JSON

**test-reader commands**:
- `view` - Interactive test browser
- `search` - Search tests

**When to modify**:
- Adding new CLI commands
- Changing command behavior
- Improving user output
- Adding new flags/options

**Modification checklist**:
```bash
# 1. Edit cmd/ccl-test-runner/main.go or cmd/test-reader/main.go

# 2. Build and test
go build -o bin/ccl-test-runner ./cmd/ccl-test-runner
./bin/ccl-test-runner --help

# 3. Update CLI documentation
# Edit docs/CLI_REFERENCE.md

# 4. Test all commands
just test
```

### 4. Statistics Package (`internal/stats/`)

**Purpose**: Collect and display test suite statistics

**Key Functions**:
- Analyze test distribution
- Count assertions
- Generate coverage reports

**When to modify**:
- Adding new statistics
- Improving analysis
- Changing output format

**Modification checklist**:
```bash
# 1. Edit internal/stats/stats.go

# 2. Test statistics collection
just stats

# 3. Verify output format
just stats --format json
```

### 5. Configuration (`internal/config/`)

**Purpose**: Manage test runner configuration

**Key Structures**:
- `RunnerConfig` - Main configuration
- `BehaviorChoices` - Behavioral choices
- `VariantChoice` - Specification variants
- `TestFilteringOptions` - Test filtering

**When to modify**:
- Adding new configuration options
- Changing behavioral choices
- Adding new filtering capabilities

**Modification checklist**:
```bash
# 1. Edit internal/config/runner_config.go

# 2. Update schema if needed
# Edit ccl-config-schema.json

# 3. Update example config
# Edit ccl-config.yaml

# 4. Test configuration loading
just test
```

## Code Quality Standards

### Go Code Style

**Formatting**:
```bash
# Format all code
go fmt ./...

# Or use just
just lint
```

**Naming Conventions**:
- Exported functions: `PascalCase`
- Unexported functions: `camelCase`
- Constants: `PascalCase` or `SCREAMING_SNAKE_CASE`
- Packages: lowercase, single word

**Error Handling**:
```go
// Always check errors
result, err := someFunction()
if err != nil {
    return nil, fmt.Errorf("failed to do something: %w", err)
}

// Use descriptive error messages
if value == "" {
    return nil, errors.New("value cannot be empty")
}
```

**Comments**:
```go
// Export functions must have doc comments
// Parse converts CCL text into a list of key-value entries.
// It returns an error if the input is malformed.
func Parse(input string) ([]Entry, error) {
    // Implementation
}

// Internal comments for complex logic
// Remove comments by filtering out lines starting with #
lines := strings.Split(input, "\n")
```

### Testing Guidelines

**Running Tests**:
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestParse ./...

# Run tests with coverage
go test -cover ./...
```

**Writing Tests** (if adding new non-generated tests):
```go
func TestNewFeature(t *testing.T) {
    // Arrange
    input := "test input"
    expected := "expected output"
    
    // Act
    result, err := NewFeature(input)
    
    // Assert
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("expected %q, got %q", expected, result)
    }
}
```

### Linting

**Required before commit**:
```bash
just lint
```

This runs:
- `gofmt` - Format code
- `golangci-lint` - Comprehensive linting
- Additional checks

**Common linting issues**:
- Unused variables/imports
- Missing error checks
- Exported functions without comments
- Inefficient code patterns

## Common Code Change Patterns

### Adding a New CCL Function

1. **Add to mock implementation**:
```go
// internal/mock/ccl.go
func (m *Mock) NewFunction(input string) (result string, error) {
    // Implementation
    return result, nil
}
```

2. **Add tests** in source_tests/:
```json
{
  "name": "test_new_function",
  "input": "test input",
  "tests": [
    {
      "function": "new_function",
      "expect": "expected result"
    }
  ]
}
```

3. **Update documentation**:
```bash
# Edit docs/MOCK_IMPLEMENTATION.md
# Edit docs/API.md
```

4. **Validate**:
```bash
just validate
just generate
just test
```

### Modifying Existing Function Behavior

1. **Understand current behavior**:
```bash
# Find all tests for the function
grep -r "\"function\": \"parse\"" source_tests/
```

2. **Make targeted change**:
```go
// Make minimal change to fix issue
// Add comments explaining the change
```

3. **Update affected tests if needed**:
```bash
# Regenerate tests
just generate
```

4. **Verify no regressions**:
```bash
just test
```

### Adding a CLI Command

1. **Define command structure**:
```go
// cmd/ccl-test-runner/main.go
&cli.Command{
    Name:  "newcommand",
    Usage: "Description of new command",
    Flags: []cli.Flag{
        &cli.StringFlag{
            Name:  "option",
            Usage: "Description of option",
        },
    },
    Action: func(c *cli.Context) error {
        // Implementation
        return nil
    },
}
```

2. **Test the command**:
```bash
go build -o bin/ccl-test-runner ./cmd/ccl-test-runner
./bin/ccl-test-runner newcommand --help
```

3. **Update documentation**:
```bash
# Edit docs/CLI_REFERENCE.md
```

### Improving Error Messages

1. **Identify unclear errors**:
```go
// Before
return nil, errors.New("invalid input")
```

2. **Make descriptive**:
```go
// After
return nil, fmt.Errorf("invalid input on line %d: expected '=' but found '%s'", 
    lineNum, token)
```

3. **Test error cases**:
```bash
just test
```

## Debugging

### Running with Debugger

Using Delve:
```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug test
dlv test ./internal/mock -- -test.run TestParse

# Debug CLI
dlv exec ./bin/ccl-test-runner -- generate
```

### Adding Debug Output

```go
import "log"

// Temporary debug output
log.Printf("DEBUG: value = %+v\n", value)

// Remove before committing
```

### Tracing Test Generation

```bash
# Verbose output
just generate-flat --verbose
just generate-go --verbose

# Check intermediate files
cat generated_tests/api_core_ccl_parsing.json | jq
```

## Dependencies

### Adding a New Dependency

```bash
# Add dependency
go get github.com/example/package

# Tidy up
go mod tidy

# Verify
go mod verify

# Update docs if user-facing
# Edit README.md or docs/DEVELOPER_GUIDE.md
```

### Updating Dependencies

```bash
# Update specific dependency
go get -u github.com/example/package

# Update all dependencies
go get -u ./...

# Clean up
go mod tidy

# Test after update
just test
```

## Performance Considerations

### Profiling

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.

# Memory profiling
go test -memprofile=mem.prof -bench=.

# View profile
go tool pprof cpu.prof
```

### Optimization Guidelines

- Profile before optimizing
- Focus on hot paths
- Use benchmarks to verify improvements
- Don't sacrifice readability for minor gains

## Pre-Commit Checklist for Code Changes

```bash
# 1. Format code
go fmt ./...

# 2. Lint (REQUIRED)
just lint

# 3. Build
go build ./...

# 4. Run tests
just test

# 5. Verify clean state
just reset

# 6. Update documentation (if API changed)
# Edit relevant docs in docs/

# 7. Review changes
git diff

# 8. Stage changes
git add .

# 9. Commit with descriptive message
git commit -m "Brief description

Detailed explanation of changes:
- What was changed
- Why it was changed
- Any side effects or breaking changes"
```

## Common Issues and Solutions

### Build Errors

**Issue**: Import cycle
```
package example imports example/internal imports example
```

**Solution**: Reorganize imports to break the cycle

---

**Issue**: Undefined function
```
undefined: SomeFunction
```

**Solution**: Check function is exported (starts with capital letter)

### Test Failures

**Issue**: Tests fail after mock changes

**Solution**:
1. Check if test expectations need updating
2. Regenerate tests: `just generate`
3. Verify behavioral changes are intentional

---

**Issue**: Generated tests are incorrect

**Solution**:
1. Check source test format: `just validate`
2. Review generator logic
3. Regenerate: `just generate`

### Linting Errors

**Issue**: "exported function should have comment"

**Solution**: Add doc comment:
```go
// FunctionName does something useful.
func FunctionName() {}
```

---

**Issue**: "error return value not checked"

**Solution**: Always check errors:
```go
// Before
SomeFunction()

// After
if err := SomeFunction(); err != nil {
    return err
}
```

## Best Practices for Code Changes

### DO:
✅ Make minimal, surgical changes
✅ Add comments for complex logic
✅ Check errors explicitly
✅ Use descriptive variable names
✅ Write tests for new functionality
✅ Update documentation when changing APIs
✅ Run `just lint` before every commit
✅ Test incrementally as you code

### DON'T:
❌ Change unrelated code
❌ Ignore linting errors
❌ Skip error handling
❌ Use magic numbers without constants
❌ Add dependencies unnecessarily
❌ Make breaking changes without discussion
❌ Commit without running `just reset`
❌ Use abbreviations in public APIs

## Advanced Topics

### Extending the Generator

When modifying test generation:

1. Understand the pipeline:
   - Source JSON → Flat JSON → Go tests
   
2. Make changes in stages:
   - First modify flat generation
   - Then modify Go test generation
   - Test each stage separately

3. Verify with real tests:
   ```bash
   just generate-flat
   cat generated_tests/sample.json | jq
   just generate-go
   cat go_tests/sample_test.go
   ```

### Working with the Type System

Understanding CCL types:
- All values start as strings
- Type inference happens at access time
- Type conversion errors should be clear

Example:
```go
func GetInt(data map[string]any, key string) (int, error) {
    value, ok := data[key]
    if !ok {
        return 0, fmt.Errorf("key %q not found", key)
    }
    
    // Handle multiple types
    switch v := value.(type) {
    case int:
        return v, nil
    case string:
        return strconv.Atoi(v)
    default:
        return 0, fmt.Errorf("cannot convert %T to int", value)
    }
}
```

### Maintaining Backward Compatibility

When changing public APIs:
1. Add new functions instead of modifying
2. Deprecate old functions with comments
3. Provide migration path in docs
4. Test both old and new code paths

## References

- **[DEVELOPER_GUIDE.md](/docs/DEVELOPER_GUIDE.md)** - Development workflows
- **[API.md](/docs/API.md)** - Complete API documentation
- **[MOCK_IMPLEMENTATION.md](/docs/MOCK_IMPLEMENTATION.md)** - Implementation patterns
- **[ARCHITECTURE.md](/docs/ARCHITECTURE.md)** - System design
- **Go Documentation**: https://golang.org/doc/
- **CLI Framework**: https://github.com/urfave/cli
