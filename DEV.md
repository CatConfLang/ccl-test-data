# Development Guide

Development setup and contribution guide for the CCL Test Data repository.

## Prerequisites

- Go 1.25+
- Node.js (for schema validation and README generation)
- [just](https://github.com/casey/just) (command runner)
- [git-cliff](https://git-cliff.org/) (for releases)

## Setup

```bash
# Install dependencies
just deps

# Verify setup - generates tests and runs them
just reset
```

## Development Workflow

### Essential Commands

| Command | Purpose |
|---------|---------|
| `just reset` | Clean state: generate + lint + test (use before commits) |
| `just dev` | Full development cycle: generate all tests and run |
| `just test` | Run tests (basic tests by default) |
| `just lint` | Format and lint Go code |
| `just validate` | Validate JSON schemas |
| `just stats` | Show test statistics |

### Pre-Commit Checklist

1. `just lint` - Format and lint Go code
2. `just reset` - Ensure clean, passing state
3. `just validate` - Schema compliance check
4. Commit updated `go_tests/` files

## Project Structure

```
source_tests/           # Human-maintainable test definitions
├── core/               # Core CCL tests
└── experimental/       # Experimental features

generated_tests/        # Machine-friendly flat format (generated)
go_tests/               # Generated Go test files

config/                 # Type-safe capability constants
loader/                 # Test loading with capability filtering
generator/              # Source-to-flat format transformation
types/                  # Shared data structures

internal/
├── mock/               # Reference CCL implementation
├── generator/          # Go test generation
└── stats/              # Statistics collection

cmd/                    # CLI applications
├── ccl-test-runner/    # Main test runner CLI
└── test-reader/        # Interactive test viewer

schemas/                # JSON schemas
docs/                   # Documentation
```

## Adding Tests

1. Add to appropriate `source_tests/core/api_*.json` file
2. Include required fields:
   - `name` - Unique identifier
   - `input` - CCL text
   - `validations` - Function validations with `count` fields
   - `functions`, `features`, `behaviors` - Metadata arrays
3. Run `just validate && just generate && just test`

See [docs/test-architecture.md](docs/test-architecture.md) for test structure details.

## Documentation

| Document | Purpose |
|----------|---------|
| [ARCHITECTURE.md](docs/ARCHITECTURE.md) | System design and components |
| [DEVELOPER_GUIDE.md](docs/DEVELOPER_GUIDE.md) | Extended development guide |
| [test-selection-guide.md](docs/test-selection-guide.md) | Test filtering documentation |
| [test-architecture.md](docs/test-architecture.md) | Test suite design |
| [schema-reference.md](docs/schema-reference.md) | Schema field reference |

## Release Process

This project uses [git-cliff](https://git-cliff.org/) for changelog generation and [Conventional Commits](https://www.conventionalcommits.org/).

### Commit Message Format

| Prefix | Version Bump | Example |
|--------|--------------|---------|
| `feat:` | Minor | `feat: add unicode test cases` |
| `fix:` | Patch | `fix: correct expected values` |
| `docs:` | Patch | `docs: update README` |
| `refactor:` | Patch | `refactor: simplify parser` |
| `feat!:` or `BREAKING CHANGE:` | Major | Breaking changes |

### Creating a Release

```bash
# 1. Check suggested version
just release-check

# 2. Preview changelog
just release-preview

# 3. Create release (updates CHANGELOG.md, commits, tags)
just release 1.2.0

# 4. Push to trigger CI
git push origin main --tags
```

CI automatically:
- Validates JSON schemas
- Rewrites `$schema` URLs to versioned GitHub raw URLs
- Creates GitHub release with JSON files and ZIP archive
- Uses changelog content for release notes

## CI Pipeline

The `just ci` command runs the full validation pipeline:

```bash
just validate   # Schema validation
just build      # Generate test files
just lint       # Go formatting and linting
just test       # Run test suite
just build-readme  # Update README stats
```

## Troubleshooting

### Schema Validation Fails
- Check JSON structure against `schemas/source-format.json`
- Ensure `count` fields match array lengths

### Generated Tests Fail
- Verify mock implementation supports required functions
- Check behavior conflicts in test metadata

### Tag Conflicts
- Ensure mutually exclusive behaviors are properly marked in `conflicts` field

See [docs/TROUBLESHOOTING.md](docs/TROUBLESHOOTING.md) for more debugging help.
