# GitHub Copilot Agent Onboarding Guide

Welcome! This guide provides comprehensive onboarding instructions for GitHub Copilot agents working on the CCL Test Suite repository.

## Repository Overview

This repository contains a **comprehensive JSON test suite** for CCL (Categorical Configuration Language) implementations across all programming languages.

### Key Statistics
- **452 test assertions** across **167 tests**
- **13 JSON test files** organized in `source_tests/core/` and `source_tests/experimental/`
- **Dual-format architecture**: Source format for maintainability, generated flat format for implementation

### Project Purpose
Enable CCL implementations in any language to validate correctness through:
- Progressive implementation testing (basic parsing → full features)
- Feature-based test selection
- Behavioral variant testing
- Comprehensive edge case coverage

## Essential First Steps

### 1. Initial Setup

```bash
# Navigate to repository
cd /home/runner/work/ccl-test-data/ccl-test-data

# Install dependencies
just deps

# Verify clean state
just reset

# View test statistics
just stats
```

### 2. Understand the Build System

- **Build Tool**: `just` (justfile-based task runner)
- **Language**: Go 1.25.1
- **Key Dependencies**: CLI framework, JSON schema validation, ccl-test-lib

### 3. Learn the Core Commands

| Command | Purpose | When to Use |
|---------|---------|-------------|
| `just lint` | Format and lint Go code | **REQUIRED before every commit** |
| `just reset` | Generate basic tests and verify passing | Pre-commit check, quick verification |
| `just validate` | Validate JSON against schema | After modifying test files |
| `just test` | Run tests | Execute test suite |
| `just generate` | Generate flat JSON then Go test files | After adding/modifying tests |
| `just stats` | Display test statistics | Review coverage and distribution |
| `just dev` | Full development cycle | Comprehensive test suite development |

## Repository Structure

```
ccl-test-data/
├── .github/
│   ├── agents/              # Agent instructions (you are here)
│   └── workflows/           # CI/CD workflows
├── cmd/                     # CLI applications
│   ├── ccl-test-runner/    # Main test runner
│   └── test-reader/        # Interactive test browser
├── internal/               # Internal packages
│   ├── mock/              # Reference CCL implementation
│   ├── generator/         # Test generation logic
│   ├── stats/             # Statistics collection
│   └── config/            # Configuration management
├── source_tests/          # Human-maintainable test source
│   ├── core/             # Core functionality tests
│   └── experimental/     # Experimental features
├── generated_tests/       # Machine-generated flat format
├── go_tests/              # Generated Go test files
├── docs/                  # Comprehensive documentation
├── schemas/               # JSON schema definitions
└── justfile              # Task definitions

**Important**: Always include `go_tests/` directory in commits when test data changes.
```

## Test Architecture

### Dual-Format System

1. **Source Format** (`source_tests/`):
   - Human-maintainable
   - Multiple validations per test
   - Rich metadata (features, behaviors, variants)

2. **Generated Format** (`generated_tests/`):
   - Machine-friendly flat format
   - One test per validation
   - Tags for precise filtering

3. **Go Tests** (`go_tests/`):
   - Generated Go test files
   - Executed by Go test runner
   - Must be committed with source changes

### CCL Function Groups

Tests are organized by function groups:

- **Core Parsing**: `parse`, `parse_indented`, `build_hierarchy`
- **Typed Access**: `get_string`, `get_int`, `get_bool`, `get_float`, `get_list`
- **Processing**: `filter`, `compose`, `merge`
- **Formatting/IO**: `canonical_format`, `load`, `round_trip`

### Test Classification

Tests include metadata for precise selection:

- **`features`**: Optional language features (`comments`, `empty_keys`, `multiline`, `unicode`, `whitespace`)
- **`behaviors`**: Implementation choices (mutually exclusive pairs like `boolean_strict` vs `boolean_lenient`)
- **`variants`**: Specification variants (`proposed_behavior`, `reference_compliant`)
- **`conflicts`**: Mutually exclusive options by category

## Development Workflow

### Standard Workflow for Any Change

```bash
# 1. Make your changes to source files

# 2. Validate changes
just validate      # Check JSON schema compliance

# 3. Generate test files (if test data changed)
just generate      # Creates flat JSON + Go tests

# 4. Lint code (REQUIRED)
just lint          # Format and check Go code

# 5. Verify tests pass
just reset         # Must pass before commit

# 6. Commit changes
git add .
git commit -m "Your descriptive message"
# Note: Include go_tests/ files if test data changed
```

### Pre-Commit Checklist

✅ **MUST DO before every commit:**
1. Run `just lint` (formats and checks code)
2. Run `just reset` (must pass completely)
3. Run `just validate` (if JSON files changed)
4. Include `go_tests/` in commit (if test data changed)

### Progressive Implementation Testing

The test suite supports progressive implementation:

**Phase 1**: Basic Parsing
```bash
just generate --run-only function:parse
just test --levels 1
```

**Phase 2**: Add Object Construction
```bash
just generate --run-only function:parse,function:build-hierarchy
just test
```

**Phase 3**: Add Typed Access
```bash
just generate --run-only function:parse,function:build-hierarchy,function:get-string
just test
```

## Common Tasks

### Adding a New Test

1. Choose appropriate file in `source_tests/core/` based on feature category
2. Add test with proper structure:
   ```json
   {
     "name": "descriptive_test_name",
     "input": "CCL input text",
     "tests": [
       {
         "function": "parse",
         "expect": [{"key": "key", "value": "value"}]
       }
     ],
     "features": ["whitespace"],
     "behaviors": ["tabs_as_whitespace"]
   }
   ```
3. Run validation and generation:
   ```bash
   just validate
   just generate
   just test
   ```

### Modifying Existing Tests

1. Edit the test in `source_tests/`
2. Regenerate derived files:
   ```bash
   just generate
   ```
3. Verify changes:
   ```bash
   just test
   ```

### Updating Mock Implementation

1. Edit `internal/mock/ccl.go`
2. Run tests:
   ```bash
   just lint
   just test
   ```
3. Update documentation if API changed

### Adding CLI Commands

1. Edit `cmd/ccl-test-runner/main.go`
2. Update `docs/CLI_REFERENCE.md`
3. Test the command:
   ```bash
   just build
   ./bin/ccl-test-runner <new-command> --help
   ```

## Quality Standards

### Code Quality
- All Go code must pass `golangci-lint`
- Use consistent formatting (`gofmt`)
- Include error handling
- Add comments for exported functions

### Test Quality
- All tests must include proper `count` fields
- Include descriptive `name` for each test
- Add appropriate metadata (`features`, `behaviors`)
- Validate with schema before committing

### Documentation Quality
- Update docs when changing functionality
- Keep examples accurate and runnable
- Cross-reference related documents
- Follow existing documentation style

## Troubleshooting

### Tests Fail After Changes

1. Check if generated files are up-to-date:
   ```bash
   just generate
   ```

2. Verify mock implementation matches test expectations

3. Check for behavioral conflicts in test metadata

### JSON Validation Fails

1. Review error message for specific issue
2. Check schema at `schemas/source-format.json`
3. Compare with working examples in `source_tests/`

### Build Errors

1. Ensure dependencies are installed:
   ```bash
   just deps
   ```

2. Check Go version (requires 1.25.1)

3. Run clean build:
   ```bash
   just clean
   go build ./...
   ```

## Best Practices for Agents

### DO:
✅ Make minimal, surgical changes
✅ Run `just lint && just reset` before every commit
✅ Include generated `go_tests/` files when test data changes
✅ Update documentation when changing functionality
✅ Use existing patterns and conventions
✅ Test incrementally as you make changes
✅ Ask for clarification if requirements are unclear

### DON'T:
❌ Modify working code unnecessarily
❌ Remove or edit unrelated tests
❌ Skip pre-commit validation steps
❌ Change code style without justification
❌ Add new dependencies without necessity
❌ Commit without running `just reset`
❌ Ignore test failures unrelated to your changes

## Getting Help

### Documentation Resources
- **[DEVELOPER_GUIDE.md](/docs/DEVELOPER_GUIDE.md)** - Detailed development workflows
- **[ARCHITECTURE.md](/docs/ARCHITECTURE.md)** - System architecture
- **[CLI_REFERENCE.md](/docs/CLI_REFERENCE.md)** - Complete CLI documentation
- **[test-selection-guide.md](/docs/test-selection-guide.md)** - Test filtering guide
- **[TROUBLESHOOTING.md](/docs/TROUBLESHOOTING.md)** - Common issues and solutions

### Specialized Agent Guides
- **[test-development.md](test-development.md)** - Adding and modifying tests
- **[code-changes.md](code-changes.md)** - Making code changes
- **[documentation.md](documentation.md)** - Updating documentation

### External Resources
- **CCL Core Concepts**: https://ccl.tylerbutler.com/core-concepts
- **CCL API Guide**: https://ccl.tylerbutler.com/api-guide

## Next Steps

1. **Read this entire guide** to understand the fundamentals
2. **Run the essential commands** to familiarize yourself with the workflow
3. **Review the specialized guides** relevant to your task
4. **Start with small changes** to build confidence
5. **Follow the pre-commit checklist** religiously

Remember: Quality over speed. It's better to make small, correct changes than large, problematic ones.
