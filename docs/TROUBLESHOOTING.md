# Troubleshooting Guide: CCL Test Suite

This guide provides solutions for common issues when working with the CCL Test Suite, including development problems, test failures, and performance issues.

## Tool Management Recommendation

> [!TIP]
> **Use mise for tool management**: The project works best with [mise](https://mise.jdx.dev/) for managing Go, Node.js, and other tool versions. Create a `.mise.toml` file in your project root with the required versions for consistent development environments.

## Common Issues and Solutions

### Development Environment Issues

#### Issue: `just` command not found
```bash
Error: command not found: just
```

**Solution:**
```bash
# Install just (task runner)
# On macOS
brew install just

# On Linux
curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to ~/bin

# On Windows
cargo install just

# Verify installation
just --version
```

#### Issue: Go version compatibility
```bash
Error: Go version 1.20 is not supported. Requires Go 1.23+
```

**Solution:**
```bash
# Recommended: Use mise for version management
mise install go@1.23.0
mise use go@1.23.0

# Or update Go manually
# Download from https://golang.org/dl/

# Verify version
go version
```

#### Issue: Missing dependencies
```bash
Error: package "github.com/santhosh-tekuri/jsonschema/cmd/jv" not found
```

**Solution:**
```bash
# Install all dependencies
just deps

# Or install manually
go install github.com/santhosh-tekuri/jsonschema/cmd/jv
cd scripts && npm install
```

### Test Generation Issues

#### Issue: Schema validation failures
```bash
Error: JSON schema validation failed
tests/api_parsing.json:15: missing required field 'count'
```

**Solution:**
```bash
# Check JSON structure against schema
just validate

# Validate specific file
jv schemas/source-format.json tests/api_parsing.json

# Common fixes:
# 1. Add missing 'count' fields to validations
# 2. Ensure 'meta.tags' includes required function tags
# 3. Verify JSON syntax (trailing commas, quotes)
```

#### Issue: Test generation produces no output
```bash
Generated 0 tests with 0 total assertions
```

**Solution:**
```bash
# Check if filters are too restrictive
ccl-test-runner generate --run-only function:parse --verbose

# Verify source test files exist
ls -la source_tests/

# Check for tag mismatches
ccl-test-runner stats --verbose
```

#### Issue: Template compilation errors
```bash
Error: template compilation failed: unexpected token
```

**Solution:**
```bash
# Clean and regenerate
just clean
just generate

# Check template syntax in internal/generator/templates.go
# Verify template data structure matches expectations
```

### Test Execution Issues

#### Issue: Mock implementation test failures
```bash
FAIL: TestGeneratedParsing (0.00s)
    Expected: [{"key":"name","value":"John"}]
    Actual:   []
```

**Solution:**
```bash
# Use basic test set that mock implementation supports
just reset  # Generate only passing tests

# Debug mock implementation
go test -v ./internal/mock

# Check if test requires unsupported features
ccl-test-runner generate --run-only function:parse --skip-tags feature:unicode
```

#### Issue: Performance test failures
```bash
FAIL: BenchmarkTestGeneration
    Performance regression: 150% slower than baseline
```

**Solution:**
```bash
# Run detailed benchmark analysis
ccl-test-runner benchmark --verbose

# Check for memory leaks
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

# Reset performance baseline if intentional
ccl-test-runner benchmark --save-baseline benchmarks/new-baseline.json
```

### JSON Schema Issues

#### Issue: Structured tag validation errors
```bash
Error: invalid tag format 'function-parse' (expected 'function:parse')
```

**Solution:**
```json
// Correct tag format
"tags": [
  "function:parse",           // ✅ Correct
  "feature:comments"          // ✅ Correct
]

// Incorrect formats
"tags": [
  "function-parse",           // ❌ Wrong separator
  "parse",                    // ❌ Missing category
  "function:parse:extra"      // ❌ Too many colons
]
```

#### Issue: Count field mismatches
```bash
Error: validation 'parse' declares count=2 but expected array has 3 items
```

**Solution:**
```json
// Fix count to match array length
{
  "validations": {
    "parse": {
      "count": 3,  // ✅ Matches array length
      "expected": [
        {"key": "a", "value": "1"},
        {"key": "b", "value": "2"},
        {"key": "c", "value": "3"}
      ]
    }
  }
}
```

### CLI Tool Issues

#### Issue: CLI commands not found
```bash
Error: ccl-test-runner: command not found
```

**Solution:**
```bash
# Build and install CLI
just build
just install

# Or run directly
go run ./cmd/ccl-test-runner generate

# Add to PATH
export PATH=$PATH:$(pwd)/bin
```

#### Issue: CLI flag parsing errors
```bash
Error: unknown flag: --run-only
```

**Solution:**
```bash
# Check available flags
ccl-test-runner generate --help

# Verify flag syntax
ccl-test-runner generate --run-only="function:parse"  # With equals
ccl-test-runner generate --run-only function:parse   # Without equals
```

### Git and Repository Issues

#### Issue: Repository not in clean state
```bash
Error: Repository has uncommitted changes
Please commit or stash changes before running tests
```

**Solution:**
```bash
# Check git status
git status

# Commit generated files
git add go_tests/
git commit -m "feat: update generated tests"

# Or use reset to clean state
just reset
```

#### Issue: Generated files out of sync
```bash
Warning: Generated files are older than source files
Please regenerate tests
```

**Solution:**
```bash
# Regenerate all files
just clean
just generate

# Always commit generated files with source changes
git add source_tests/ go_tests/ generated_tests/
git commit -m "feat: add new test cases with generated files"
```

## Debugging Strategies

### Verbose Output

#### Enable verbose test generation
```bash
ccl-test-runner generate --verbose
ccl-test-runner test --format verbose
```

#### Debug specific test issues
```bash
# Test single function
ccl-test-runner generate --run-only function:parse
ccl-test-runner test --functions parse

# Debug specific feature
ccl-test-runner generate --run-only feature:comments --verbose
```

### Mock Implementation Debugging

#### Test mock implementation directly
```go
// Add debug prints to internal/mock/ccl.go
func (c *CCL) Parse(input string) ([]Entry, error) {
    fmt.Printf("DEBUG: Parsing input: %q\n", input)
    // ... implementation
    fmt.Printf("DEBUG: Parsed entries: %+v\n", entries)
    return entries, nil
}
```

#### Use Go debugging tools
```bash
# Run with race detection
go test -race ./internal/mock

# Run with memory sanitizer
go test -msan ./internal/mock

# Use delve debugger
dlv test ./internal/mock
```

### Performance Debugging

#### Profile test generation
```bash
go test -cpuprofile=cpu.prof -bench=BenchmarkGeneration ./internal/generator
go tool pprof cpu.prof
```

#### Monitor memory usage
```bash
go test -memprofile=mem.prof -bench=. ./...
go tool pprof mem.prof
```

#### Trace execution
```bash
go test -trace=trace.out -bench=. ./...
go tool trace trace.out
```

## Error Message Reference

### Common Error Patterns

#### Schema Validation Errors
| Error Message | Cause | Solution |
|---------------|--------|----------|
| `missing required field 'count'` | Validation missing count field | Add `"count": N` to validation |
| `invalid tag format` | Wrong tag structure | Use `category:value` format |
| `conflicting behaviors` | Mutually exclusive tags | Remove conflicting tags or use conflicts field |

#### Generation Errors
| Error Message | Cause | Solution |
|---------------|--------|----------|
| `no tests generated` | Filters too restrictive | Check `--run-only` and `--skip-tags` filters |
| `template compilation failed` | Template syntax error | Check template files in `internal/generator/` |
| `file write permission denied` | Permission issue | Check directory permissions for output |

#### Test Execution Errors
| Error Message | Cause | Solution |
|---------------|--------|----------|
| `mock implementation failed` | Feature not implemented | Use `--skip-tags` or implement feature |
| `assertion count mismatch` | Wrong count field | Fix count to match expected results |
| `timeout exceeded` | Performance issue | Check for infinite loops or optimize |

## Recovery Procedures

### Reset to Clean State
```bash
# Nuclear option: reset everything
just clean
git checkout -- go_tests/ generated_tests/
just reset
```

### Rebuild from Scratch
```bash
# Complete rebuild
rm -rf go_tests/ generated_tests/ bin/
just deps
just build
just reset
```

### Fix Corrupted Test Data
```bash
# Validate all test files
find source_tests/ -name "*.json" -exec jv schemas/source-format.json {} \;

# Fix common issues automatically
# (This would be a script to fix common formatting issues)
./scripts/fix-test-format.sh source_tests/
```

## Getting Help

### Self-Service Debugging
1. **Check Documentation**: Review relevant docs in `docs/` directory
2. **Run Diagnostics**: Use `just stats` and `just validate` for health checks
3. **Check Logs**: Enable verbose output for detailed error information
4. **Compare with Working State**: Use `git bisect` to find breaking changes

### Community Support
1. **GitHub Issues**: Search existing issues or create new one
2. **Documentation**: Reference implementation patterns in `docs/`
3. **Code Examples**: Study mock implementation for reference patterns

### Contributing Fixes
1. **Reproduce Issue**: Create minimal reproduction case
2. **Write Test**: Add test case that demonstrates the issue
3. **Implement Fix**: Fix the issue with appropriate tests
4. **Submit PR**: Include test case and fix in pull request

## Prevention Strategies

### Development Best Practices
1. **Always run `just lint`** before committing
2. **Use `just reset`** to ensure clean state
3. **Validate JSON** with `just validate` after changes
4. **Test incrementally** with function-based testing

### Quality Gates
```bash
# Pre-commit checklist
just lint                    # ✅ Code formatting and style
just validate                # ✅ JSON schema compliance
just reset                   # ✅ Basic tests pass
just test                    # ✅ Full test suite
```

### Continuous Integration
```yaml
# Example CI workflow
name: Quality Check
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
      - name: Lint and Test
        run: |
          just deps
          just lint
          just validate
          just test
          just benchmark
```

This troubleshooting guide covers the most common issues and provides systematic approaches to diagnosing and fixing problems in the CCL Test Suite.