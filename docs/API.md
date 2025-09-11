# CCL Test Suite API Documentation

This document provides comprehensive API documentation for the CCL Test Suite Go packages.

## Package Overview

The CCL Test Suite consists of several Go packages that work together to provide a comprehensive testing framework for CCL implementations:

- **`cmd/ccl-test-runner`** - Main CLI application for test generation and execution
- **`internal/mock`** - Reference CCL implementation for testing and development
- **`internal/generator`** - Test file generation from JSON test data
- **`internal/stats`** - Test suite statistics and analytics
- **`internal/benchmark`** - Performance benchmarking utilities
- **`internal/styles`** - Terminal styling and output formatting
- **`internal/types`** - Common data structures and schema definitions

## Core Types

### Entry Structure

```go
type Entry struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```

The fundamental data structure representing a key-value pair from CCL parsing. Used throughout all levels of the CCL implementation hierarchy.

### Test Metadata

```go
type TestMeta struct {
    Tags     []string `json:"tags"`
    Level    int      `json:"level"`
    Feature  string   `json:"feature"`
    Conflicts []string `json:"conflicts,omitempty"`
}
```

Metadata structure for test categorization and progressive implementation support:
- **Tags**: Structured feature tags for precise test selection (`function:*`, `feature:*`, `behavior:*`)
- **Level**: CCL implementation level (1-5)
- **Feature**: Primary feature category being tested
- **Conflicts**: Mutually exclusive behaviors for this test

## Package: internal/mock

The mock package provides a working CCL implementation that serves as both a reference implementation and a development tool for progressive CCL implementation.

### CCL Implementation

```go
type CCL struct{}

func New() *CCL
```

Creates a new mock CCL implementation instance.

### Level 1: Raw Parsing

```go
func (c *CCL) Parse(input string) ([]Entry, error)
```

**Purpose**: Convert CCL text input into flat key-value entries.

**Features**:
- Basic key-value pair parsing
- Comment support using '/=' syntax
- Line-by-line processing with error context
- Whitespace normalization

**Returns**: Array of Entry structures or parsing error with line context.

**Example**:
```go
ccl := mock.New()
entries, err := ccl.Parse("key = value\n/= comment\nother = data")
// Returns: [{"key": "key", "value": "value"}, {"key": "other", "value": "data"}]
```

### Level 2: Entry Processing

```go
func (c *CCL) Filter(entries []Entry, predicate func(Entry) bool) []Entry
func (c *CCL) Compose(entries1, entries2 []Entry) []Entry
func (c *CCL) ExpandDotted(entries []Entry) []Entry
```

**Filter**: Remove entries based on predicate function
**Compose**: Merge two entry arrays with duplicate handling
**ExpandDotted**: Convert dotted keys to nested structure preparation

### Level 3: Object Construction

```go
func (c *CCL) MakeObjects(entries []Entry) map[string]interface{}
```

**Purpose**: Transform flat entries into nested object hierarchies.

**Features**:
- Dotted key expansion (`database.host` → `{"database": {"host": "value"}}`)
- Duplicate key handling (converts to arrays)
- Empty key support for array-style syntax
- Type preservation for nested structures

**Returns**: Nested map structure representing the configuration object.

### Level 4: Typed Access

```go
func (c *CCL) GetString(obj map[string]interface{}, path []string) (string, error)
func (c *CCL) GetInt(obj map[string]interface{}, path []string) (int, error)
func (c *CCL) GetBool(obj map[string]interface{}, path []string) (bool, error)
func (c *CCL) GetFloat(obj map[string]interface{}, path []string) (float64, error)
func (c *CCL) GetList(obj map[string]interface{}, path []string) ([]interface{}, error)
```

**Purpose**: Type-safe value extraction with automatic conversion.

**Path Navigation**: Use string array for nested key access: `["database", "host"]`

**Type Conversion Rules**:
- **String**: Direct string values or string representation of other types
- **Int**: Numeric string parsing with error handling
- **Bool**: "true"/"false" string parsing (case-insensitive)
- **Float**: Decimal number parsing
- **List**: Array values or single values converted to single-item arrays

**Error Handling**: Returns typed errors for missing keys, type conversion failures, and path navigation errors.

## Package: internal/generator

The generator package handles conversion of JSON test data to executable Go test files.

### Generator Configuration

```go
type Options struct {
    SkipDisabled bool
    SkipTags     []string
    RunOnly      []string
}

func NewWithOptions(inputDir, outputDir string, options Options) *Generator
```

**SkipDisabled**: Skip tests with disabled feature tags
**SkipTags**: Additional tags to exclude from generation
**RunOnly**: Generate only tests with specified tags (overrides skip behavior)

### Test Generation

```go
func (g *Generator) GenerateAll() error
func (g *Generator) GetStats() GenerationStats
```

**GenerateAll**: Process all JSON test files and generate corresponding Go test files

**GetStats**: Return statistics about test generation including counts of total tests, active tests, and assertions

### Generation Statistics

```go
type GenerationStats struct {
    TotalTests        int
    SkippedTests      int
    TotalAssertions   int
    SkippedAssertions int
}
```

## Package: internal/stats

The stats package provides comprehensive analytics and reporting for the test suite.

### Enhanced Statistics Collection

```go
func NewEnhancedCollector(testDir string) *EnhancedCollector
func (c *EnhancedCollector) CollectEnhancedStats() (*EnhancedStats, error)
```

Collects detailed statistics including:
- Test count by level and feature
- Assertion distribution analysis
- Tag usage patterns
- Feature coverage metrics

### Statistics Display

```go
func PrintEnhancedStats(stats *EnhancedStats)
```

Provides formatted output with:
- Summary statistics
- Level-by-level breakdown
- Feature coverage analysis
- Tag distribution
- Test file organization

## Package: internal/benchmark

The benchmark package provides performance measurement and regression detection.

### Benchmark Tracking

```go
type Tracker struct{}

func NewTracker() *Tracker
func (t *Tracker) StartBenchmark(name string)
func (t *Tracker) EndBenchmark(name string) BenchmarkResult
```

**Performance Metrics Tracked**:
- Execution duration
- Memory allocation (bytes)
- Peak memory usage
- Operation counts

### Regression Detection

```go
func CompareResults(current, historical []BenchmarkResult, threshold float64) []RegressionAlert
```

Compares current performance against historical results and identifies regressions exceeding the specified threshold percentage.

## Package: internal/styles

The styles package provides consistent terminal output formatting across all CLI commands.

### Status Messages

```go
func Status(icon, message string, args ...interface{})
func Success(message string, args ...interface{})
func Error(message string, args ...interface{})
func Warning(message string, args ...interface{})
```

**Styled Output Functions**:
- **Status**: Progress indicators with icons
- **Success**: Green checkmark success messages
- **Error**: Red error messages with formatting
- **Warning**: Yellow warning messages
- **Info/InfoLite**: Information messages with different emphasis levels

### Command Display

```go
func Command(cmd string)
```

Displays executed commands with consistent formatting for debugging and transparency.

## CLI Integration

### Command Structure

The CLI application provides four main commands:

1. **generate**: Convert JSON test data to Go test files
2. **test**: Execute generated tests with enhanced formatting
3. **stats**: Display comprehensive test suite analytics
4. **benchmark**: Performance measurement and regression detection

### Test Filtering

All commands support filtering by:
- **Levels**: CCL implementation levels (1-5)
- **Features**: Feature categories (parsing, objects, comments, etc.)
- **Tags**: Structured tags for precise selection

### Output Formats

Multiple output formats supported:
- **pretty**: Human-readable formatted output (default)
- **json**: Machine-readable JSON output
- **table**: Tabular format for test results
- **verbose**: Detailed output with full context

## Usage Patterns

### Progressive Implementation

1. **Start with Level 1**: Basic parsing functionality
   ```bash
   ./ccl-test-runner generate --run-only function:parse
   ./ccl-test-runner test --levels 1
   ```

2. **Add Object Construction**: Level 3 functionality
   ```bash
   ./ccl-test-runner generate --run-only function:parse,function:make-objects
   ./ccl-test-runner test --levels 1,3
   ```

3. **Complete Type System**: Add Level 4 typed access
   ```bash
   ./ccl-test-runner generate --run-only function:parse,function:make-objects,function:get-string
   ./ccl-test-runner test --levels 1,3,4
   ```

### Feature-Specific Development

Target specific CCL features:
```bash
# Comments support
./ccl-test-runner generate --run-only feature:comments
./ccl-test-runner test --features comments

# Dotted keys
./ccl-test-runner generate --run-only feature:dotted-keys
./ccl-test-runner test --features dotted-keys
```

### Quality Assurance

```bash
# Full test suite validation
./ccl-test-runner stats --format json > test-metrics.json
./ccl-test-runner benchmark --compare historical-results.json
./ccl-test-runner test --format verbose > test-results.log
```

## Error Handling

### Parse Errors

The mock implementation provides detailed error messages with line context:
```go
// Example error: "line 3: invalid syntax: expected '=' after key 'database.host'"
entries, err := ccl.Parse(invalidInput)
if err != nil {
    log.Printf("Parse error: %v", err)
}
```

### Type Conversion Errors

Typed access functions return specific errors for conversion failures:
```go
value, err := ccl.GetInt(obj, []string{"port"})
if err != nil {
    // Error: "failed to convert 'localhost' to int: invalid syntax"
}
```

### Path Navigation Errors

Missing keys or invalid paths return navigation errors:
```go
value, err := ccl.GetString(obj, []string{"nonexistent", "key"})
if err != nil {
    // Error: "key path 'nonexistent.key' not found"
}
```

## Extension Points

### Adding New Validation Types

To add new validation types to the test suite:

1. Define validation structure in test JSON
2. Add parsing logic to generator templates
3. Implement corresponding mock function
4. Add structured tags for the new functionality

### Custom Test Runners

The JSON test format supports custom test runners:
- Iterate over `validations` object keys
- Check `count` field for assertion verification
- Use structured tags for progressive implementation
- Handle `conflicts` array for mutually exclusive behaviors

### Performance Monitoring

Extend the benchmark package for custom metrics:
- Add new benchmark operations
- Define custom performance thresholds
- Implement historical comparison logic
- Create regression alert systems