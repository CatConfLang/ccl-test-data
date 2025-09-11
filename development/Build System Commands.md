---
title: Build System Commands
type: note
permalink: development/build-system-commands
---

# CCL Test Data Build System Reference

## Essential Commands
- `just lint` - Format and lint Go code (REQUIRED before commits)
- `just reset` - Generate Level 1 tests only, ensuring all pass  
- `just dev` - Full development cycle: clean, generate, test
- `just test` - Comprehensive suite: validate + docs + generate + test

## Test Generation
- `just generate` - Generate all Go test files from JSON data
- `just generate-mock` - Generate tests optimized for mock implementation
- `just generate-level1` - Generate basic Level 1 functionality only

## Test Execution  
- `just test-generated` - Run generated Go tests only
- `just test-level1` through `just test-level4` - Run specific CCL levels
- `just test-comments`, `just test-parsing`, `just test-objects` - Feature-specific tests

## Validation & Analysis
- `just validate` - Validate JSON test files against schema
- `just stats` - Show detailed test statistics  
- `just docs-check` - Verify documentation is current

## Development Patterns
1. **Pre-commit**: Always run `just lint` before committing
2. **Clean State**: Use `just reset` to ensure repository is in clean state
3. **Mock Development**: Use `just generate-mock` for progressive implementation
4. **Quality Gates**: All enabled tests should pass before commits