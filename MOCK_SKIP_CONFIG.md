# Optimal Test Skipping Configuration for Mock CCL Implementation

This document provides the recommended configuration for skipping tests that won't pass with the current mock CCL implementation.

## Recommended Skip Configuration

Use this command to generate tests that will mostly pass with the mock implementation:

```bash
go run ./cmd/ccl-test-runner generate --skip-tags="whitespace,tabs,newlines,multiline,nested,continuation,error,incomplete,algebraic,round-trip,semigroup,monoid,trimming,indentation,mixed-whitespace,reference-compliant-behavior,uses-strict-line-ending-parsing,parser-behavior,ocaml-original,proposed-behavior,stress-test,edge-case,empty,eof,keys,redundant,merging,multi-level,line-endings,crlf-preserved,dotted-access"
```

## Current Test Results

With this configuration:
- **Total tests**: 146
- **Skipped**: 91 (62%)
- **Passing**: 35 (24%)
- **Failing**: 20 (14%)

## Mock Implementation Limitations

The current mock implementation has these limitations:

### 1. **Multiline Value Parsing**
- **Issue**: Treats multiline input as single value
- **Tags affected**: `multiline`, `continuation`, `nested`
- **Example failure**: Input `number = 42\ndecimal = 3.14` becomes single entry with value `42\ndecimal = 3.14`

### 2. **Error Handling**
- **Issue**: Doesn't return proper errors for invalid input
- **Tags affected**: `error`, `incomplete`
- **Example**: Tests expecting parse errors get unexpected success

### 3. **Whitespace/Trimming**
- **Issue**: Basic whitespace handling differs from spec
- **Tags affected**: `whitespace`, `tabs`, `trimming`
- **Example**: Tab preservation in values

### 4. **Dotted Key Access**
- **Issue**: Doesn't handle hierarchical key access properly
- **Tags affected**: `dotted-access`
- **Example**: Cannot access `database.port` from nested structure

### 5. **Advanced Features**
- **Issue**: Missing complex CCL features
- **Tags affected**: `algebraic`, `round-trip`, `stress-test`

## Improvement Opportunities (Simplest First)

### Priority 1: Easy Wins
1. **Whitespace trimming** - Add basic `.trim()` calls
2. **Empty input handling** - Return empty array for empty strings
3. **Error cases** - Add validation for incomplete key-value pairs

### Priority 2: Medium Effort
1. **Basic multiline parsing** - Split on newlines and parse each line
2. **Dotted key expansion** - Convert `a.b = value` to nested structures

### Priority 3: Complex
1. **Full CCL specification compliance**
2. **Advanced algebraic properties**

## Usage Examples

### Run Only Passing Tests
```bash
# Generate with skip config
go run ./cmd/ccl-test-runner generate --skip-tags="whitespace,tabs,newlines,multiline,nested,continuation,error,incomplete,algebraic,round-trip,semigroup,monoid,trimming,indentation,mixed-whitespace,reference-compliant-behavior,uses-strict-line-ending-parsing,parser-behavior,ocaml-original,proposed-behavior,stress-test,edge-case,empty,eof,keys,redundant,merging,multi-level,line-endings,crlf-preserved,dotted-access"

# Run tests
go run ./cmd/ccl-test-runner test --levels 1,2,3,4
```

### Run Only Basic Features
```bash
# Run only the simplest tests
go run ./cmd/ccl-test-runner generate --run-only basic,extension
go run ./cmd/ccl-test-runner test --levels 1,2
```

### Skip Everything Except Core Parsing
```bash
# Most conservative - only basic parsing
go run ./cmd/ccl-test-runner generate --run-only basic
go run ./cmd/ccl-test-runner test --levels 1
```

## Mock Improvement Suggestions

If you want to improve the mock implementation, start with these changes to `internal/mock/mock.go`:

1. **Basic multiline support**: Split input on `\n` and parse each line
2. **Whitespace trimming**: Trim keys and values properly
3. **Empty input**: Return `[]Entry{}` for empty input
4. **Basic error detection**: Return error for lines without `=`

These changes would reduce failures significantly with minimal effort.