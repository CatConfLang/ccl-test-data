# Developer Guide: CCL Test Suite

Development guide for extending and contributing to the CCL Test Suite.

## Project Architecture

### CCL Function Groups
Progressive implementation approach:
- **Core**: Parse + BuildHierarchy
- **Typed Access**: GetString, GetInt, GetBool, GetFloat, GetList
- **Processing**: Filter, Compose
- **Formatting**: PrettyPrint

### Package Structure
```
cmd/                    # CLI applications
internal/
├── mock/              # Reference CCL implementation
├── generator/         # Test file generation
├── stats/             # Statistics and analytics
└── types/             # Common data structures
tests/                 # JSON test data files
go_tests/              # Generated Go test files
```

### Structured Tagging
- **`function:*`** - Required CCL functions
- **`feature:*`** - Optional language features
- **`behavior:*`** - Implementation choices

## Development Workflow

### Essential Commands
| Command | Purpose |
|---------|---------|
| `just lint` | Format and lint (REQUIRED before commits) |
| `just reset` | Generate basic tests and verify passing |
| `just test` | Full test suite execution |
| `just generate` | Generate all Go tests from JSON |
| `just validate` | Validate JSON against schema |
| `just stats` | Display comprehensive statistics |

### Pre-Commit Checklist
1. `just lint` - Format and lint
2. `just reset` - Validate clean state (must pass)
3. `just validate` - Schema validation
4. Commit updated `go_tests/` files

## Adding New Tests

### Test File Organization
Tests are organized by feature category:
- `api_essential-parsing.json` - Basic core functionality
- `api_object-construction.json` - Core object building
- `api_typed-access.json` - Type-safe access functions
- `api_comments.json` - Comment syntax support
- `api_errors.json` - Error handling validation

### Test Structure
```json
{
  "name": "descriptive_test_name",
  "input": "CCL input text",
  "validations": {
    "parse": {
      "count": 1,
      "expected": [{"key": "name", "value": "value"}]
    },
    "get_string": {
      "count": 2,
      "cases": [
        {"args": ["key1"], "expected": "value1"},
        {"args": ["key2"], "expected": "value2"}
      ]
    }
  },
  "features": ["comments"],
  "behaviors": []
}
```

### Validation Types
- **Direct**: `"parse": {"count": 1, "expected": [...]}`
- **Case-based**: `"get_string": {"count": 2, "cases": [...]}`
- **Error**: `"parse_error": {"count": 1, "expected_error": "..."}`

### Adding a New Test
1. Choose appropriate test file by feature category
2. Write test structure with proper validation types
3. Add structured metadata for functions and features
4. Include count fields matching array lengths
5. Run `just validate && just generate && just test`

## Extending the Mock Implementation

### Implementation in `internal/mock/ccl.go`

The mock implementation provides working CCL functionality across multiple function groups:

#### Core Functions
- `Parse(input string) ([]Entry, error)` - Basic key-value parsing
- `BuildHierarchy(entries []Entry) map[string]interface{}` - Object construction

#### Typed Access Functions
- `GetString(obj, path) (string, error)`
- `GetInt(obj, path) (int, error)`
- `GetBool(obj, path) (bool, error)`

### Adding New Functions
1. Define function signature matching test expectations
2. Implement core functionality with proper error handling
3. Add structured metadata to relevant tests
4. Run `just reset && just test` to verify

### Error Handling Pattern
```go
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
The generator uses Go templates to create test files from JSON data in `internal/generator/templates.go`.

### Adding New Validation Types
1. Define validation structure in test JSON
2. Add template generation logic to handle new validation type
3. Update test schema in `tests/schema.json`
4. Run `just generate && just test`

### Filtering Options
```go
type Options struct {
    SkipDisabled bool      // Skip disabled feature tags
    SkipTags     []string  // Additional metadata to skip
    RunOnly      []string  // Generate only these metadata elements
}
```

### Progressive Implementation
```bash
# Core functions only
./ccl-test-runner generate --run-only function:parse

# Exclude advanced features
./ccl-test-runner generate --skip-tags feature:unicode
```

## CLI Command Development

### Command Structure
CLI commands use the urfave/cli/v2 pattern in `cmd/ccl-test-runner/main.go`.

### Adding New Commands
1. Define command structure with flags
2. Implement action function with error handling
3. Add to command list in main app definition
4. Update documentation

### Output Formatting
Use the styles package for consistent output:
```go
styles.Status("🔍", "Analyzing test files...")
styles.Success("✅ Analysis complete")
styles.Error("❌ Failed to process: %v", err)
```

## Testing and Quality Assurance

### Quality Gates
Before commits:
1. `just lint` - Linting (golangci-lint + gofmt)
2. `just validate` - JSON schema compliance
3. `just reset` - All enabled tests pass
4. `just test` - Comprehensive validation

### Mock Implementation Testing
```bash
just test-core         # Test core functionality only
just test-typed        # Test typed access functions
```

### Error Scenarios
Test error handling with `parse_error` validations that specify expected error patterns.

## Performance Considerations

### Benchmarking
```bash
just benchmark         # Run performance benchmarks
```

The benchmark package tracks performance metrics for test generation and statistics collection.

### Optimization Guidelines
1. **Memory Management**: Reuse data structures
2. **I/O Efficiency**: Batch file operations
3. **Test Generation**: Use object pooling for allocations

## Best Practices

### Code Organization
- Package separation by concerns
- Interface usage for testability
- Error wrapping with context
- Godoc comments for exported functions

### Test Design
- Progressive complexity (simple → edge cases)
- Clear naming describing validation purpose
- Proper structured metadata
- Accurate count fields

### Git Workflow
- Feature branches with descriptive names
- Atomic commits per logical change
- Always commit updated `go_tests/` files
- Clean commit history via rebase

## Making a Release

This project uses [git-cliff](https://git-cliff.org/) for changelog generation
and follows [Conventional Commits](https://www.conventionalcommits.org/).

### Prerequisites

Install git-cliff:
```bash
# macOS
brew install git-cliff

# cargo
cargo install git-cliff

# or use mise/asdf
```

### Release Workflow

1. **Check suggested version** (based on commits since last tag):
   ```bash
   just release-check
   ```
   This analyzes commits and suggests: patch (fix), minor (feat), or major (BREAKING CHANGE).

2. **Preview the changelog**:
   ```bash
   just release-preview
   ```

3. **Create the release** (with your chosen version):
   ```bash
   just release 1.2.0
   ```
   This will:
   - Update CHANGELOG.md with all changes since last release
   - Commit the changelog
   - Create tag `data-v1.2.0`

4. **Push to trigger CI**:
   ```bash
   git push origin main --tags
   ```

   CI will automatically:
   - Validate JSON schemas
   - Rewrite `$schema` URLs to versioned GitHub raw URLs
   - Create GitHub release with individual JSON files and ZIP archive
   - Use changelog content for release notes

### Commit Message Format

Follow conventional commits for automatic changelog generation:

| Prefix | Description | Version Bump |
|--------|-------------|--------------|
| `feat:` | New feature | Minor |
| `fix:` | Bug fix | Patch |
| `docs:` | Documentation | Patch |
| `perf:` | Performance | Patch |
| `refactor:` | Code refactoring | Patch |
| `test:` | Tests | Patch |
| `chore:` | Maintenance | Patch |
| `feat!:` or `BREAKING CHANGE:` | Breaking change | Major |

**Examples:**
- `feat: add new test cases for unicode` → Features section
- `fix: correct expected values in parsing tests` → Bug Fixes section
- `feat(schema): add new behavior field` → Features with scope

## Troubleshooting

### Common Issues
1. **Schema Validation Fails**: Check JSON structure against `tests/schema.json`
2. **Generated Tests Fail**: Verify mock implementation supports required functions
3. **Tag Conflicts**: Ensure mutually exclusive behaviors are properly marked

### Debug Commands
```bash
just test-verbose               # Detailed test output
just stats --format json       # Statistics breakdown
./ccl-test-runner generate --run-only function:parse  # Specific subset
```