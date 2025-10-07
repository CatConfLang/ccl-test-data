# GitHub Copilot Agent Instructions

This directory contains specialized instructions for GitHub Copilot agents working on the CCL Test Suite repository.

## Available Agent Guides

### General Purpose
- **[onboarding.md](onboarding.md)** - Primary onboarding guide for all agents. Start here to understand the repository structure, workflow, and essential commands.

### Specialized Guides
- **[test-development.md](test-development.md)** - Guide for agents working on test suite development, including adding new tests and test generation.
- **[code-changes.md](code-changes.md)** - Guide for agents making changes to Go code, including the mock implementation and CLI tools.
- **[documentation.md](documentation.md)** - Guide for agents updating documentation files and ensuring consistency across docs.

## How to Use These Guides

1. **Always start with [onboarding.md](onboarding.md)** to understand the project fundamentals
2. **Reference specialized guides** based on your specific task
3. **Follow the workflows** exactly as described to maintain quality standards
4. **Use the checklists** in each guide to ensure you don't miss critical steps

## Key Principles

All agents working on this repository should follow these principles:

- **Minimal Changes**: Make the smallest possible changes to achieve the goal
- **Test Before Commit**: Always run `just lint && just reset` before committing
- **Include Generated Files**: Commit updated `go_tests/` files when modifying test data
- **Documentation Sync**: Update documentation when changing functionality
- **Schema Compliance**: Validate JSON files with `just validate`

## Quick Reference

Essential commands for all agents:
```bash
just deps      # Install dependencies (first time setup)
just reset     # Generate basic tests and verify (pre-commit check)
just lint      # Format and lint Go code (required before commits)
just validate  # Validate JSON test files
just stats     # View test coverage statistics
```

## Getting Help

If you encounter issues or need clarification:
- Check the [TROUBLESHOOTING.md](/docs/TROUBLESHOOTING.md) guide
- Review the [DEVELOPER_GUIDE.md](/docs/DEVELOPER_GUIDE.md) for detailed workflows
- Refer to the [CLI_REFERENCE.md](/docs/CLI_REFERENCE.md) for command details
