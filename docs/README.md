# CCL Test Suite Documentation

This directory contains comprehensive documentation for the CCL (Categorical Configuration Language) Test Suite.

## Documentation Index

### Core Documentation

- **[ARCHITECTURE.md](ARCHITECTURE.md)** - System architecture, component design, and data flow
- **[API.md](API.md)** - Complete API reference for all Go packages and functions
- **[DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md)** - Development workflow and contribution guide
- **[MOCK_IMPLEMENTATION.md](MOCK_IMPLEMENTATION.md)** - Reference CCL implementation patterns
- **[CLI_REFERENCE.md](CLI_REFERENCE.md)** - Command-line interface reference
- **[TROUBLESHOOTING.md](TROUBLESHOOTING.md)** - Common issues and debugging strategies

### Test Suite Documentation

- **[test-architecture.md](test-architecture.md)** - Test suite design and file organization
- **[test-filtering.md](test-filtering.md)** - Advanced test filtering patterns
- **[test-selection-guide.md](test-selection-guide.md)** - Guide for selecting relevant tests
- **[test-runner-implementation-guide.md](test-runner-implementation-guide.md)** - Building test runners

### Implementation Guides

- **[implementing-ccl.md](implementing-ccl.md)** - Complete CCL implementation guide
- **[schema-reference.md](schema-reference.md)** - Test data schema reference
- **[generated-schema.md](generated-schema.md)** - Auto-generated schema documentation

### Project Overview

The CCL Test Suite is a comprehensive, language-agnostic testing framework for CCL implementations. It provides:

- **Structured Test Data** - JSON-based test definitions with feature-based tagging
- **Progressive Implementation** - Support for 5-level CCL implementation approach  
- **Reference Implementation** - Working mock CCL parser for development and testing
- **Advanced Tooling** - CLI tools for test generation, execution, and analysis

### Quick Navigation

**For New Contributors:**
1. Begin with [ARCHITECTURE.md](ARCHITECTURE.md) for system design overview
2. Follow with [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md) for development workflow and contribution guidelines
3. Review [API.md](API.md) for detailed package documentation
4. Use [CLI_REFERENCE.md](CLI_REFERENCE.md) for command-line tool usage

**For CCL Implementers:**
1. Study [MOCK_IMPLEMENTATION.md](MOCK_IMPLEMENTATION.md) for implementation patterns and strategies
2. Reference [API.md](API.md) for expected function signatures and behavior
3. Use [CLI_REFERENCE.md](CLI_REFERENCE.md) for progressive testing approaches

**For Tool Users:**
1. Reference [CLI_REFERENCE.md](CLI_REFERENCE.md) for complete command documentation
2. Check [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md) for workflow examples
3. Use [API.md](API.md) for understanding output formats and data structures

### Documentation Standards

All documentation in this directory follows these standards:

- **Comprehensive Coverage** - Complete information for each topic
- **Practical Examples** - Real-world usage examples and code samples
- **Progressive Structure** - Information organized from basic to advanced
- **Cross-References** - Links between related concepts across documents
- **Maintenance** - Updated with code changes and feature additions

### Contributing to Documentation

When contributing to the project, please ensure documentation is updated:

1. **API Changes** - Update API.md with new functions, parameters, or return types
2. **New Features** - Add examples and usage patterns to DEVELOPER_GUIDE.md
3. **CLI Modifications** - Update CLI_REFERENCE.md with new commands or options
4. **Implementation Patterns** - Document new patterns in MOCK_IMPLEMENTATION.md

### External References

For conceptual explanations of Core CCL, see [Core Concepts](https://ccl.tylerbutler.com/core-concepts)
For API compatibility across implementations, see [Higher-Level APIs](https://ccl.tylerbutler.com/higher-level-apis)

### Getting Help

For questions about the documentation:

1. **General Questions** - Open an issue in the project repository
2. **Implementation Help** - Reference MOCK_IMPLEMENTATION.md and API.md
3. **Usage Questions** - Check CLI_REFERENCE.md and DEVELOPER_GUIDE.md
4. **Contributing** - Follow DEVELOPER_GUIDE.md workflow and quality standards