---
title: Project Overview
type: note
permalink: project/project-overview
---

# CCL Test Data Project Context

## Project Summary
Official JSON test suite for CCL (Categorical Configuration Language) implementations with comprehensive Go-based test runner and mock implementation.

## Architecture Overview
- **Multi-Level CCL Implementation**: 5 levels from raw parsing to validation/formatting
- **Test Organization**: Feature-based tagging system with structured test selection
- **Build System**: `just` (justfile) for cross-platform automation
- **Dependencies**: Go 1.24.0, CLI frameworks, JSON schema validation

## Key Components
- `cmd/ccl-test-runner/` - Main CLI for test generation and execution
- `cmd/test-reader/` - TUI for test data exploration  
- `internal/mock/ccl.go` - Working CCL implementation for testing
- `internal/generator/` - Go test file generation from JSON data
- `internal/stats/` - Enhanced statistics and analysis
- `tests/api-*.json` - Feature-specific test suites with structured tagging

## Development Workflow
```bash
just lint                   # Format and lint Go code (REQUIRED before commits)
just reset                  # Generate basic tests, ensuring all pass
just dev                   # Full development cycle
just test                  # Comprehensive test suite
just stats                 # Show detailed test statistics
```

## Analysis Results (2025-09-11)
- **Overall Rating**: 🟢 EXCELLENT (8.2/10)
- **Security**: No vulnerabilities found
- **Architecture**: Clean modular design with clear boundaries
- **Performance**: Adequate with optimization opportunities
- **Code Quality**: Professional Go code following best practices

## Mock Implementation Strategy
The `internal/mock/ccl.go` contains working CCL implementation supporting:
- Level 1: `Parse()` - Basic parsing with comment support
- Level 3: `MakeObjects()` - Object construction with dotted key support  
- Level 4: `GetString()`, `GetInt()`, etc. - Typed access functions
- Repository maintains clean state with `just reset` generating only passing tests