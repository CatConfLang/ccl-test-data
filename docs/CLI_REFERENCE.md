# CCL Test Runner CLI Reference

Command-line interface reference for the CCL Test Suite's test runner and utilities.

## CLI Tools

| Tool | Purpose |
|------|---------|
| `ccl-test-runner` | Main test generation and execution |
| `test-reader` | Interactive test browser |
| `validate-schema` | JSON schema validation |
| `clean` | Cleanup and maintenance |

## Installation

```bash
# Build and install
go build -o bin/ccl-test-runner ./cmd/ccl-test-runner
just install                # Install to PATH
```

## Main Command: ccl-test-runner

### Global Options
| Flag | Description |
|------|-------------|
| `--help, -h` | Show help information |
| `--version` | Show version information |

### Command: generate

Generate Go test files from JSON test data.

#### Options
| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--input` | `-i` | `tests` | Input directory containing JSON test files |
| `--output` | `-o` | `go_tests` | Output directory for generated test files |
| `--skip-disabled` | | `true` | Skip tests with disabled feature tags |
| `--skip-tags` | | | Additional tags to skip (comma-separated) |
| `--run-only` | | | Only generate tests with these tags |

#### Examples
```bash
# Basic usage
ccl-test-runner generate
ccl-test-runner generate --output my_tests

# Progressive implementation
ccl-test-runner generate --run-only function:parse
ccl-test-runner generate --run-only function:parse,function:build_hierarchy

# Skip advanced features
ccl-test-runner generate --skip-tags feature:unicode,feature:multiline
```

### Command: test

Run generated tests with enhanced output formatting.

#### Options
| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `pretty` | Output format (pretty, table, verbose, json) |
| `--features` | | | Filter by features (comments, parsing, objects) |
| `--list` | | | List available test packages without running |
| `--verbose` | `-v` | | Verbose output (same as --format verbose) |

#### Output Formats
- **pretty**: Clean, human-readable output with color coding (default)
- **table**: Tabular test results in compact format
- **verbose**: Detailed output with full error messages
- **json**: Machine-readable structured output

#### Examples
```bash
# Basic testing
ccl-test-runner test
ccl-test-runner test --format table

# Filtering
ccl-test-runner test --features comments,parsing

# Pass-through Go test flags
ccl-test-runner test -cover -race
ccl-test-runner test -run TestGenerated.*
```

### Command: stats

Collect and display comprehensive test suite statistics.

#### Options
| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--input` | `-i` | `tests` | Input directory containing JSON test files |
| `--format` | `-f` | `pretty` | Output format (pretty, json) |

#### Output Formats
- **pretty**: Formatted statistics with sections (default)
- **json**: Machine-readable JSON for automation

#### Example Output (Pretty)
```
📊 CCL Test Suite Statistics

📊 SUMMARY
  Total Tests: 157, Total Assertions: 423, Test Files: 11

🏷️ FUNCTION TAGS
  function:parse: 89 tests
  function:build_hierarchy: 45 tests
  function:get_string: 34 tests
```

### Command: benchmark

Run performance benchmarks on core operations.

#### Options
| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--input` | `-i` | `tests` | Input directory containing JSON test files |
| `--output` | `-o` | `go_tests` | Output directory for generated test files |
| `--results` | `-r` | `benchmarks/results.json` | File to save benchmark results |
| `--compare` | `-c` | | Historical results file to compare against |
| `--threshold` | | `10.0` | Regression threshold percentage |

#### Benchmark Operations
- **Test Generation**: JSON parsing, template generation, file I/O
- **Statistics Collection**: Analysis time, memory usage, aggregation

#### Example Output
```
🚀 Performance Benchmarks

🏃 RESULTS
Test Generation: 234ms, 2.3 MB allocated, 157 tests generated
Statistics: 89ms, 890 KB allocated, 11 files analyzed

✅ Results saved to benchmarks/results.json
```

#### Regression Detection
```bash
ccl-test-runner benchmark --compare benchmarks/historical.json --threshold 15.0
```

## Utility Commands

### test-reader
Interactive test browser for exploring the test suite.
- Browse tests by category
- Examine individual test structure and validations
- Search by name, tags, or content
- Export filtered test sets

### validate-schema
Validate JSON test files against the schema.
```bash
validate-schema tests/*.json           # Validate all test files
validate-schema tests/api_parsing.json # Validate specific file
```

### clean
Cleanup and maintenance utility.
- Remove old generated test files
- Clear build artifacts
- Remove old benchmark results

## Integration with Just

Task runner for common operations:

### Essential Commands
| Command | Purpose |
|---------|---------|
| `just lint` | Format and lint code |
| `just reset` | Generate basic tests and verify |
| `just test` | Full test suite execution |
| `just generate` | Generate all tests |
| `just stats` | Display statistics |
| `just validate` | Validate JSON schema |
| `just benchmark` | Run performance benchmarks |

### Feature-Specific Testing
| Command | Purpose |
|---------|---------|
| `just test-parsing` | Test parsing features |
| `just test-objects` | Test object construction |
| `just test-comments` | Test comment features |

## Environment Variables

| Variable | Purpose |
|----------|---------|
| `GO_TEST_FLAGS` | Default flags passed to go test |
| `GOTESTSUM_FORMAT` | Default gotestsum output format |

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | General error |
| `2` | Usage error |
| `125` | Test failures |

## Example Workflows

### Progressive Implementation
```bash
# Start with basic parsing
ccl-test-runner generate --run-only function:parse
ccl-test-runner test

# Add object construction
ccl-test-runner generate --run-only function:parse,function:build_hierarchy
ccl-test-runner test
```

### Continuous Integration
```bash
validate-schema tests/*.json
ccl-test-runner generate
ccl-test-runner test --format json -cover > coverage.json
ccl-test-runner stats --format json > stats.json
```

### Quality Assurance
```bash
just lint          # Format and lint
just reset         # Must pass for commits
just test          # Comprehensive tests
just benchmark     # Check performance
```