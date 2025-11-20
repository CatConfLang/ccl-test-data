# CCL Test Suite API Documentation

API documentation for the CCL Test Suite Go packages.

## Package Overview

- **`cmd/ccl-test-runner`** - Main CLI application
- **`internal/mock`** - Reference CCL implementation
- **`internal/generator`** - Test file generation from JSON
- **`internal/stats`** - Test suite statistics
- **`internal/benchmark`** - Performance benchmarking
- **`internal/styles`** - Terminal styling
- **`internal/types`** - Common data structures

## Core Types

### Entry Structure
```go
type Entry struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```
Fundamental key-value pair from CCL parsing.

### Test Metadata
```go
type TestMeta struct {
    Functions []string `json:"functions"`
    Features  []string `json:"features"`
    Behaviors []string `json:"behaviors"`
    Variants  []string `json:"variants"`
    Conflicts ConflictSpec `json:"conflicts,omitempty"`
}

type ConflictSpec struct {
    Behaviors []string `json:"behaviors,omitempty"`
    Variants  []string `json:"variants,omitempty"`
}
```
Metadata for test categorization and progressive implementation:
- **Functions**: Required CCL functions
- **Features**: Optional language features
- **Behaviors**: Implementation choices
- **Conflicts**: Mutually exclusive behaviors

## Package: internal/mock

Working CCL implementation for testing and development.

### CCL Implementation
```go
type CCL struct{}
func New() *CCL
```

### Core Functions
```go
func (c *CCL) Parse(input string) ([]Entry, error)
func (c *CCL) build_hierarchy(entries []Entry) map[string]interface{}
```
- Basic key-value parsing with comment support (`/=` syntax)
- Hierarchical object construction via fixpoint algorithm

### Typed Access Functions
```go
func (c *CCL) get_string(obj map[string]interface{}, path []string) (string, error)
func (c *CCL) get_int(obj map[string]interface{}, path []string) (int, error)
func (c *CCL) get_bool(obj map[string]interface{}, path []string) (bool, error)
func (c *CCL) get_float(obj map[string]interface{}, path []string) (float64, error)
func (c *CCL) get_list(obj map[string]interface{}, path []string) ([]interface{}, error)
```
Type-safe value extraction with automatic conversion. Path navigation uses string arrays: `["database", "host"]`.

## Package: internal/generator

Converts JSON test data to executable Go test files.

### Generator Configuration
```go
type Options struct {
    SkipDisabled bool      // Skip tests with disabled feature tags
    SkipTags     []string  // Additional tags to exclude
    RunOnly      []string  // Generate only specified tags
}

func NewWithOptions(inputDir, outputDir string, options Options) *Generator
```

### Test Generation
```go
func (g *Generator) GenerateAll() error
func (g *Generator) GetStats() GenerationStats

type GenerationStats struct {
    TotalTests        int
    SkippedTests      int
    TotalAssertions   int
    SkippedAssertions int
}
```

## Package: internal/stats

Test suite analytics and reporting.

### Statistics Collection
```go
func NewEnhancedCollector(testDir string) *EnhancedCollector
func (c *EnhancedCollector) CollectEnhancedStats() (*EnhancedStats, error)
func PrintEnhancedStats(stats *EnhancedStats)
```
Collects test counts, assertion distribution, tag usage patterns, and feature coverage.

## Package: internal/benchmark

Performance measurement and regression detection.

### Benchmark Tracking
```go
type Tracker struct{}
func NewTracker() *Tracker
func (t *Tracker) StartBenchmark(name string)
func (t *Tracker) EndBenchmark(name string) BenchmarkResult

func CompareResults(current, historical []BenchmarkResult, threshold float64) []RegressionAlert
```
Tracks execution duration, memory allocation, and identifies performance regressions.

## Package: internal/styles

Consistent terminal output formatting.

### Status Messages
```go
func Status(icon, message string, args ...interface{})
func Success(message string, args ...interface{})
func Error(message string, args ...interface{})
func Warning(message string, args ...interface{})
func Command(cmd string)
```
Styled output functions with icons and color coding.

## CLI Integration

### Command Structure
The CLI provides four main commands:
1. **generate**: Convert JSON test data to Go test files
2. **test**: Execute generated tests with enhanced formatting
3. **stats**: Display comprehensive test suite analytics
4. **benchmark**: Performance measurement and regression detection

### Test Filtering
Commands support filtering by:
- **Functions**: CCL function categories
- **Features**: Feature categories (parsing, objects, comments)
- **Tags**: Structured tags for precise selection

### Output Formats
- **pretty**: Human-readable formatted output (default)
- **json**: Machine-readable JSON output
- **table**: Tabular format for test results
- **verbose**: Detailed output with full context

## Usage Patterns

### Progressive Implementation
```bash
# Start with core functions
./ccl-test-runner generate --run-only function:parse,function:build_hierarchy
./ccl-test-runner test --functions parse

# Add typed access
./ccl-test-runner generate --run-only function:get_string
./ccl-test-runner test --functions parse,build_hierarchy,get_string
```

### Feature-Specific Development
```bash
# Comments support
./ccl-test-runner generate --run-only feature:comments
./ccl-test-runner test --features comments
```

## Error Handling

### Error Types
- **Parse Errors**: Detailed messages with line context
- **Type Conversion Errors**: Specific conversion failure messages
- **Path Navigation Errors**: Missing key path notifications

```go
// Parse error example
entries, err := ccl.Parse(invalidInput)
// Error: "line 3: invalid syntax: expected '=' after key 'database.host'"

// Type conversion error
value, err := ccl.GetInt(obj, []string{"port"})
// Error: "failed to convert 'localhost' to int: invalid syntax"

// Path navigation error
value, err := ccl.GetString(obj, []string{"nonexistent", "key"})
// Error: "key path 'nonexistent.key' not found"
```

## Extension Points

### Adding New Validation Types
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