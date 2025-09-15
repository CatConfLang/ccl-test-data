# Test Runner Implementation Patterns - Comprehensive Guide

## Overview
Comprehensive documentation created for test runner implementation addressing real-world scenarios with mixed implemented/unimplemented function requirements.

## Core Problem Addressed
**Question**: When a test has multiple validations but implementation doesn't support all required functions, what should the test runner do?

**Context**: Progressive CCL implementation where developers incrementally add function support but tests may require multiple functions.

## Solution Architecture

### 1. Test Execution Strategies

#### Partial Validation Execution (Recommended)
- **Execute**: Implemented validations  
- **Skip**: Unimplemented validations
- **Report**: PARTIAL status with clear breakdown
- **Benefits**: Provides feedback on working functionality, supports progressive development
- **Trade-offs**: May miss integration issues between implemented/unimplemented parts

#### Skip Entire Test (Conservative)
- **Execute**: Nothing if any function missing
- **Report**: SKIPPED status with reason
- **Benefits**: Clean results, ensures test integrity
- **Trade-offs**: Loses valuable testing feedback

#### Fail Fast (Strict)
- **Execute**: Fail immediately on unimplemented function
- **Report**: FAILED status with clear error
- **Benefits**: Forces complete implementation
- **Trade-offs**: Blocks progressive development

### 2. Progressive Implementation Patterns

#### Function-Based Filtering
```pseudocode
// Pre-filter tests during discovery
required_functions = extract_function_tags(test.meta.tags)
if all_required_functions_implemented(required_functions, capabilities) {
  include_test(test)
}
```

#### Runtime Validation Skipping
```pseudocode
for validation_name, validation_spec in test.validations {
  if implementation.supports(validation_name) {
    execute_validation(validation_name, validation_spec, test.input)
  } else {
    skip_validation(validation_name, "Function not implemented")
  }
}
```

#### Tag-Based Capabilities
```pseudocode
type TestCapabilities = {
  functions: Set<string>,     // "parse", "build_hierarchy", etc.
  features: Set<string>,      // "comments", "dotted_keys", etc.
  behaviors: Set<string>      // "crlf_normalize", etc.
}
```

### 3. Standard Test Result Schema

#### Result Types
- **PASSED**: All validations executed and passed
- **FAILED**: At least one validation failed  
- **PARTIAL**: Some validations passed, others skipped
- **SKIPPED**: Entire test skipped due to missing requirements

#### Validation Types
- **PASSED**: Validation executed and passed
- **FAILED**: Validation executed but failed
- **SKIPPED**: Validation skipped (function not implemented)
- **ERROR**: Validation couldn't execute due to test runner error

#### Detailed Result Structure
```json
{
  "test_name": "basic_dotted_key_parsing",
  "status": "PARTIAL", 
  "validations": [
    {"name": "parse", "status": "PASSED", "message": "All 2 assertions passed"},
    {"name": "expand_dotted", "status": "SKIPPED", "message": "Function not implemented"},
    {"name": "build_hierarchy", "status": "PASSED", "message": "All 1 assertions passed"}
  ],
  "summary": "2/3 validations executed successfully",
  "metadata": {
    "functions_required": ["parse", "expand_dotted", "build_hierarchy"],
    "functions_executed": ["parse", "build_hierarchy"], 
    "functions_skipped": ["expand_dotted"]
  }
}
```

## Implementation Examples

### Configuration-Driven Test Execution
```pseudocode
type TestConfiguration = {
  enabled_functions: string[],
  enabled_features: string[],
  enabled_behaviors: string[],
  execution_strategy: "partial" | "skip_entire" | "fail_fast",
  skip_tags: string[],
  only_tags: string[]
}

// Usage examples:
minimal_config = TestConfiguration(
  enabled_functions: ["parse", "build_hierarchy"],
  execution_strategy: "partial"
)

experimental_config = TestConfiguration(
  enabled_functions: ["parse", "expand_dotted", "build_hierarchy"],
  enabled_features: ["dotted_keys"],
  only_tags: ["feature:experimental_dotted_keys"],
  execution_strategy: "partial"
)
```

### Smart Test Prioritization
```pseudocode
function prioritize_tests(tests, capabilities) {
  categories = {
    fully_supported: [],    // All required functions implemented
    partially_supported: [], // Some functions implemented  
    unsupported: []         // No functions implemented
  }
  // Prioritize fully supported tests first for faster feedback
}
```

## Advanced Patterns

### Dependency Resolution
- Some validations depend on others (e.g., `get_string` depends on `parse` + `build_hierarchy`)
- Handle dependencies gracefully with topological sorting
- Skip dependent validations if prerequisites fail

### Error Handling
- Graceful degradation with multiple error types
- Clear distinction between implementation gaps vs bugs
- Comprehensive error reporting for debugging

### Test Discovery and Filtering
- Tag-based test discovery with complex filtering rules
- Smart inclusion/exclusion based on capabilities
- Support for experimental feature isolation

## Key Design Principles

### 1. Progressive Implementation Friendly
- Support incremental function implementation
- Provide meaningful feedback at every stage
- Clear guidance on next implementation steps

### 2. Tag-Driven Architecture
- Use modern structured tagging (`function:*`, `feature:*`, `behavior:*`)
- Avoid outdated level-based systems
- Enable precise test selection and filtering

### 3. Clear Result Semantics
- Distinguish between "not tested" vs "tested and failed"
- Comprehensive metadata for debugging and planning
- Standard schema for tool interoperability

### 4. Configuration Flexibility
- Support different execution strategies
- Enable feature-specific test runs
- Accommodate various implementation approaches

## Documentation Impact

### Fills Critical Gap
Previous documentation covered:
- Test architecture and organization
- CCL implementation guidance
- Basic test runner examples

**Missing**: Practical patterns for handling mixed implementation scenarios

### Provides Concrete Solutions
- Real-world test execution strategies
- Progressive implementation workflows
- Standard result reporting schemas
- Language-agnostic implementation patterns

## Use Cases Addressed

### During Development
- **Early stages**: Run only basic parsing tests
- **Incremental development**: Add features one at a time
- **Integration testing**: Verify interaction between implemented functions

### Production Readiness
- **Feature validation**: Confirm implemented features work correctly
- **Gap analysis**: Identify missing functionality
- **Quality assurance**: Comprehensive testing of supported features

### Team Collaboration
- **Progress tracking**: Clear visibility into implementation status
- **Issue reporting**: Standardized result format for bug reports
- **Planning**: Data-driven decisions about next features to implement

This comprehensive guide transforms test runner implementation from ad-hoc solutions to systematic, well-defined patterns that support real-world CCL development workflows.