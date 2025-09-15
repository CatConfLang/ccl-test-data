# CCL Test Runner CLI Reference

Comprehensive command-line interface reference for the CCL Test Suite's test runner and related utilities.

## Overview

The CCL Test Suite provides several CLI tools for test generation, execution, and analysis:

- **`ccl-test-runner`** - Main test generation and execution tool
- **`test-reader`** - Interactive test browser and exploration tool
- **`validate-schema`** - JSON schema validation utility
- **`clean`** - Cleanup and maintenance utility

## Installation

### Build from Source

```bash
# Build main CLI
go build -o bin/ccl-test-runner ./cmd/ccl-test-runner

# Build all utilities
just build        # Main runner
just build-reader # Test browser
```

### Install to PATH

```bash
# Install to $GOPATH/bin
go install ./cmd/ccl-test-runner

# Or use just
just install
```

## Main Command: ccl-test-runner

The primary CLI application for working with the CCL test suite.

### Global Options

```bash
ccl-test-runner [GLOBAL_OPTIONS] COMMAND [COMMAND_OPTIONS]
```

**Global Flags:**
- `--help, -h` - Show help information
- `--version` - Show version information

**Version Information:**
- Shows build version, git commit, and build date
- Format: `version (commit: hash, built: timestamp)`

### Command: generate

Generate Go test files from JSON test data.

#### Usage

```bash
ccl-test-runner generate [OPTIONS]
```

#### Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--input` | `-i` | `tests` | Input directory containing JSON test files |
| `--output` | `-o` | `go_tests` | Output directory for generated test files |
| `--skip-disabled` | | `true` | Skip tests with disabled feature tags |
| `--skip-tags` | | | Additional tags to skip (comma-separated) |
| `--run-only` | | | Only generate tests with these tags (overrides skip behavior) |

#### Examples

**Basic Generation:**
```bash
# Generate all tests
ccl-test-runner generate

# Generate to custom directory
ccl-test-runner generate --output my_tests
```

**Progressive Implementation:**
```bash
# Generate Level 1 tests only
ccl-test-runner generate --run-only function:parse

# Generate basic functionality
ccl-test-runner generate --run-only function:parse,function:make-objects,function:get-string

# Skip advanced features
ccl-test-runner generate --skip-tags feature:unicode,feature:multiline
```

**Feature-Specific Generation:**
```bash
# Generate comment tests only
ccl-test-runner generate --run-only feature:comments

# Generate parsing and object tests
ccl-test-runner generate --run-only function:parse,function:make-objects

# Skip experimental features
ccl-test-runner generate --skip-tags variant:proposed-behavior
```

#### Output

The generate command provides detailed feedback:

```
🚀 Generating test files...
Generated 157 tests with 423 total assertions
Active tests: 134 (with 367 assertions)
Skipped tests: 23 (with 56 assertions)
✅ Test generation completed successfully
```

### Command: test

Run generated tests with enhanced output formatting.

#### Usage

```bash
ccl-test-runner test [OPTIONS] [GO_TEST_FLAGS...]
```

#### Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `pretty` | Output format (pretty, table, verbose, json) |
| `--levels` | `-l` | | Filter by CCL levels (1,2,3,4) |
| `--features` | | | Filter by features (comments, parsing, objects, etc) |
| `--tags` | | | Filter by tags (not yet implemented) |
| `--list` | | | List available test packages without running |
| `--verbose` | `-v` | | Verbose output (same as --format verbose) |

#### Output Formats

**Pretty Format (Default):**
```bash
ccl-test-runner test --format pretty
```
- Clean, human-readable output
- Color-coded results
- Progress indicators

**Table Format:**
```bash
ccl-test-runner test --format table
```
- Tabular test results
- Compact overview format

**Verbose Format:**
```bash
ccl-test-runner test --format verbose
```
- Detailed test output
- Full error messages and stack traces
- Same as `--verbose` flag

**JSON Format:**
```bash
ccl-test-runner test --format json
```
- Machine-readable output
- Structured test results for automation

#### Filtering Examples

**By Level:**
```bash
# Test Level 1 only (basic parsing)
ccl-test-runner test --levels 1

# Test Levels 1 and 3 (parsing and objects)
ccl-test-runner test --levels 1,3

# Test all levels
ccl-test-runner test --levels 1,2,3,4
```

**By Feature:**
```bash
# Test comment functionality
ccl-test-runner test --features comments

# Test parsing and objects
ccl-test-runner test --features parsing,objects

# Test dotted key support
ccl-test-runner test --features dotted-keys
```

**List Available Tests:**
```bash
# List all test packages
ccl-test-runner test --list

# List filtered packages
ccl-test-runner test --levels 1,2 --list
```

#### Pass-through Options

Additional Go test flags can be passed directly:

```bash
# Run with coverage
ccl-test-runner test -cover

# Run specific test pattern
ccl-test-runner test -run TestGenerated.*

# Run with race detection
ccl-test-runner test -race

# Multiple flags
ccl-test-runner test --format verbose -cover -race -timeout 30s
```

### Command: stats

Collect and display comprehensive test suite statistics.

#### Usage

```bash
ccl-test-runner stats [OPTIONS]
```

#### Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--input` | `-i` | `tests` | Input directory containing JSON test files |
| `--format` | `-f` | `pretty` | Output format (pretty, json) |

#### Output Formats

**Pretty Format (Default):**
```bash
ccl-test-runner stats
```

Produces formatted output:
```
📊 Collecting test statistics...

════════════════════════════════════════════════════════════════════
                          CCL Test Suite Statistics
════════════════════════════════════════════════════════════════════

📊 SUMMARY
  Total Tests:        157
  Total Assertions:   423
  Test Files:         11
  Average Assertions: 2.7 per test

🏗️  IMPLEMENTATION LEVELS
  Level 1 (Core CCL):              81 tests (234 assertions)
  Level 2 (Typed Access):         41 tests (98 assertions)
  Level 3 (Advanced Processing):  23 tests (67 assertions)
  Level 4 (Experimental):         12 tests (24 assertions)

🏷️  STRUCTURED TAGS
  Function Tags:
    function:parse              89 tests
    function:make-objects       45 tests
    function:get-string         34 tests
    function:get-int           12 tests
  
  Feature Tags:
    feature:comments           23 tests
    feature:dotted-keys        18 tests
    feature:unicode            8 tests
    feature:multiline          12 tests

📁 FILE BREAKDOWN
  api_essential-parsing.json:     34 tests (89 assertions)
  api_comprehensive-parsing.json: 23 tests (67 assertions)
  api_comments.json:             18 tests (45 assertions)
  api_object-construction.json:   25 tests (76 assertions)
  api_typed-access.json:         31 tests (84 assertions)
  api_dotted-keys.json:          15 tests (38 assertions)
  api_processing.json:           11 tests (24 assertions)
```

**JSON Format:**
```bash
ccl-test-runner stats --format json
```

Produces machine-readable JSON output suitable for automation and further processing:

```json
{
  "summary": {
    "totalTests": 157,
    "totalAssertions": 423,
    "testFiles": 11,
    "averageAssertions": 2.7
  },
  "levels": {
    "level1": {"tests": 81, "assertions": 234},
    "level2": {"tests": 41, "assertions": 98},
    "level3": {"tests": 23, "assertions": 67},
    "level4": {"tests": 12, "assertions": 24}
  },
  "tags": {
    "function": {
      "function:parse": 89,
      "function:make-objects": 45,
      "function:get-string": 34
    },
    "feature": {
      "feature:comments": 23,
      "feature:dotted-keys": 18
    }
  }
}
```

### Command: benchmark

Run performance benchmarks on core operations.

#### Usage

```bash
ccl-test-runner benchmark [OPTIONS]
```

#### Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--input` | `-i` | `tests` | Input directory containing JSON test files |
| `--output` | `-o` | `go_tests` | Output directory for generated test files |
| `--results` | `-r` | `benchmarks/results.json` | File to save benchmark results |
| `--compare` | `-c` | | Historical results file to compare against |
| `--threshold` | | `10.0` | Regression threshold percentage |

#### Benchmark Operations

The benchmark command measures:

1. **Test Generation Performance**
   - Time to parse JSON test files
   - Template generation speed
   - File I/O performance

2. **Statistics Collection Performance**
   - JSON parsing and analysis time
   - Memory usage during statistics collection
   - Aggregation operation performance

#### Output

```bash
ccl-test-runner benchmark
```

Produces performance metrics:

```
🚀 Running performance benchmarks...

════════════════════════════════════════════════════════════════════
                           Performance Benchmarks
════════════════════════════════════════════════════════════════════

🏃 BENCHMARK RESULTS

Test Generation:
  Duration:        234ms
  Memory Allocated: 2.3 MB
  Peak Memory:     4.1 MB
  Operations:      157 tests generated

Statistics Collection:
  Duration:        89ms  
  Memory Allocated: 890 KB
  Peak Memory:     1.2 MB
  Operations:      11 files analyzed

✅ Benchmark results saved to benchmarks/results.json
```

#### Regression Detection

```bash
# Compare against historical results
ccl-test-runner benchmark --compare benchmarks/historical.json --threshold 15.0
```

If performance regressions exceed the threshold:

```
⚠️  Performance regressions detected!

REGRESSION ALERTS:
  Test Generation: 25% slower (234ms vs 187ms)
  Memory Usage: 18% increase (2.3MB vs 1.9MB)

Performance regression threshold exceeded
```

## Utility Commands

### test-reader

Interactive test browser for exploring the test suite.

#### Usage

```bash
test-reader [OPTIONS]
```

#### Features

- **Interactive Navigation**: Browse tests by category and level
- **Test Detail View**: Examine individual test structure and validations
- **Search Functionality**: Find tests by name, tags, or content
- **Export Options**: Save filtered test sets for development

### validate-schema

Validate JSON test files against the schema.

#### Usage

```bash
validate-schema [FILES...]
```

#### Examples

```bash
# Validate all test files
validate-schema tests/*.json

# Validate specific file
validate-schema tests/api_parsing.json

# Used internally by 'just validate'
```

### clean

Cleanup and maintenance utility.

#### Usage

```bash
clean [OPTIONS]
```

#### Features

- **Generated File Cleanup**: Remove old generated test files
- **Temporary File Cleanup**: Clear build artifacts
- **Benchmark Cleanup**: Remove old benchmark results

## Integration with Just

The project uses `just` as a task runner for common operations:

### Essential Commands

```bash
# Development workflow
just lint                   # Format and lint code
just reset                  # Generate basic tests and verify
just test                   # Full test suite execution

# Generation commands
just generate               # Generate all tests
just generate-mock          # Generate tests for mock implementation

# Testing commands  
just test-generated         # Run generated tests only
just test-level1            # Test Level 1 only
just test-verbose           # Verbose test output

# Analysis commands
just stats                  # Display statistics
just validate               # Validate JSON schema
just benchmark              # Run performance benchmarks
```

### Advanced Commands

```bash
# Level-specific testing
just test-level2            # Test Level 2 functionality
just test-level3            # Test Level 3 functionality
just test-level4            # Test Level 4 functionality

# Feature-specific testing
just test-parsing           # Test parsing features
just test-objects           # Test object construction
just test-comments          # Test comment features

# Development utilities
just build                  # Build main CLI
just build-reader           # Build test browser
just install               # Install to PATH
just clean                 # Clean build artifacts
```

## Environment Variables

### Configuration

- `GO_TEST_FLAGS` - Default flags passed to go test
- `GOTESTSUM_FORMAT` - Default gotestsum output format  
- `CCL_TEST_DEBUG` - Enable debug output (not implemented)

### Usage Examples

```bash
# Set default test format
export GOTESTSUM_FORMAT=testname

# Add default coverage
export GO_TEST_FLAGS="-cover -race"

# Use in CI/CD
CCL_TEST_DEBUG=1 ccl-test-runner test --format json > results.json
```

## Exit Codes

The CLI uses standard exit codes:

- `0` - Success
- `1` - General error (parsing, file not found, etc.)
- `2` - Usage error (invalid flags, missing arguments)
- `125` - Test failures (when running tests)

## Examples and Workflows

### Progressive Implementation Workflow

```bash
# 1. Start with Level 1 parsing
ccl-test-runner generate --run-only function:parse
ccl-test-runner test --levels 1

# 2. Add object construction (Level 1 core)
ccl-test-runner generate --run-only function:parse,function:make-objects
ccl-test-runner test --levels 1

# 3. Add typed access (Level 2)
ccl-test-runner generate --run-only function:parse,function:make-objects,function:get-string
ccl-test-runner test --levels 1,2

# 4. Add remaining typed access functions
ccl-test-runner generate --run-only function:parse,function:make-objects,function:get-string,function:get-int,function:get-bool
ccl-test-runner test --levels 1,2
```

### Feature Development Workflow

```bash
# 1. Analyze current test coverage
ccl-test-runner stats

# 2. Generate tests for new feature
ccl-test-runner generate --run-only feature:new-feature

# 3. Run tests to identify failures
ccl-test-runner test --features new-feature --format verbose

# 4. Implement feature and retest
ccl-test-runner test --features new-feature

# 5. Run full suite to ensure no regressions
ccl-test-runner test
```

### Continuous Integration Workflow

```bash
# 1. Validate test data
validate-schema tests/*.json

# 2. Generate all tests
ccl-test-runner generate

# 3. Run tests with coverage
ccl-test-runner test --format json -cover > coverage.json

# 4. Collect statistics
ccl-test-runner stats --format json > stats.json

# 5. Run performance benchmarks
ccl-test-runner benchmark --results benchmarks/ci-results.json
```

### Quality Assurance Workflow

```bash
# 1. Format and lint
just lint

# 2. Validate clean state  
just reset  # Must pass for commits

# 3. Run comprehensive tests
just test

# 4. Check performance
just benchmark

# 5. Analyze coverage
ccl-test-runner stats
```

This CLI reference provides comprehensive documentation for all available commands, options, and workflows in the CCL Test Suite toolkit.