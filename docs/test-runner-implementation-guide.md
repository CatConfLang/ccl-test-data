# CCL Test Runner Implementation Guide

This guide explains how to implement test runners for the CCL flat test format. The flat format simplifies test execution by having one test per validation, making implementation straightforward.

## Flat Test Format Overview

Each test file contains an array of independent tests. Each test validates exactly one CCL function:

```json
[
  {
    "name": "basic_parsing_parse",
    "input": "key = value",
    "validation": "parse",
    "expected": {
      "count": 1,
      "entries": [{"key": "key", "value": "value"}]
    },
    "functions": ["parse"],
    "behaviors": [],
    "variants": [],
    "features": [],
    "level": 1
  }
]
```

## Test Selection Strategy

Since each test validates exactly one function, test execution is simple: **run tests for implemented functions, skip tests for unimplemented functions**.

### Function-Based Filtering

Filter tests based on your implementation's capabilities:

```pseudocode
function get_runnable_tests(all_tests, implemented_functions) {
  runnable_tests = []
  
  for test in all_tests {
    required_functions = test.functions
    
    if all_functions_implemented(required_functions, implemented_functions) {
      runnable_tests.append(test)
    }
  }
  
  return runnable_tests
}

function all_functions_implemented(required, implemented) {
  for function_name in required {
    if function_name not in implemented {
      return false
    }
  }
  return true
}
```

### Behavior and Feature Filtering

Filter tests based on implementation behaviors and supported features:

```pseudocode
function test_is_compatible(test, capabilities) {
  // Check function requirements
  for function_name in test.functions {
    if function_name not in capabilities.functions {
      return false
    }
  }
  
  // Check feature requirements  
  for feature_name in test.features {
    if feature_name not in capabilities.features {
      return false
    }
  }
  
  // Check behavior compatibility
  for behavior_name in test.behaviors {
    if behavior_name not in capabilities.behaviors {
      return false
    }
  }
  
  return true
}
```

## Behavior Conflicts and Resolution

### Understanding Behavior Conflicts

Some behaviors are mutually exclusive - implementations must choose one approach:

```pseudocode
// Conflicting behavior groups
conflicts = {
  "line_endings": ["crlf_preserve_literal", "crlf_normalize_to_lf"],
  "boolean_parsing": ["boolean_lenient", "boolean_strict"], 
  "spacing": ["strict_spacing", "loose_spacing"],
  "tab_handling": ["tabs_preserve", "tabs_to_spaces"],
  "list_coercion": ["list_coercion_enabled", "list_coercion_disabled"]
}
```

### Automatic Conflict Resolution

When generating tests, skip conflicting behaviors your implementation doesn't support:

```pseudocode
function resolve_behavior_conflicts(desired_behaviors, implementation_choices) {
  skip_behaviors = []
  
  // Skip CRLF preservation if implementation normalizes to LF
  if "crlf_normalize_to_lf" in implementation_choices {
    skip_behaviors.append("crlf_preserve_literal")
  }
  
  // Skip strict spacing if implementation uses loose spacing
  if "loose_spacing" in implementation_choices {
    skip_behaviors.append("strict_spacing")
  }
  
  // Skip tab preservation if implementation doesn't handle tabs
  if "tabs_normalize" in implementation_choices {
    skip_behaviors.append("tabs_preserve")
  }
  
  return skip_behaviors
}
```

### CLI Conflict Resolution Examples

```bash
# Implementation that normalizes CRLF to LF - skip literal preservation
ccl-test-runner generate \
  --run-only function:parse \
  --skip-tags behavior:crlf_preserve_literal,behavior:strict_spacing

# Implementation with loose spacing - skip strict spacing tests  
ccl-test-runner generate \
  --run-only function:parse,function:make_objects \
  --skip-tags behavior:strict_spacing,behavior:tabs_preserve

# Implementation with specific boolean handling - skip conflicting approach
ccl-test-runner generate \
  --run-only function:get_bool \
  --skip-tags behavior:boolean_strict  # Use lenient parsing
```

## Implementation Examples by Level

### Level 1: Parse Only
```pseudocode
capabilities = {
  functions: ["parse"],
  features: [],
  behaviors: ["crlf_normalize_to_lf", "boolean_lenient"]
}

// Filters to only function:parse tests
// Skips tests requiring behaviors like tabs_preserve, strict_spacing
```

### Level 2: Parse + Processing  
```pseudocode
capabilities = {
  functions: ["parse", "filter", "compose", "expand_dotted"],
  features: ["comments"],
  behaviors: ["crlf_normalize_to_lf", "boolean_lenient"]
}
```

### Level 3: Parse + Objects
```pseudocode
capabilities = {
  functions: ["parse", "make_objects"],
  features: ["dotted_keys", "empty_keys"],
  behaviors: ["crlf_normalize_to_lf", "boolean_lenient"]
}
```

### Level 4: Parse + Objects + Typed Access
```pseudocode
capabilities = {
  functions: ["parse", "make_objects", "get_string", "get_int", "get_bool", "get_float", "get_list"],
  features: ["dotted_keys", "empty_keys", "comments"],
  behaviors: ["crlf_normalize_to_lf", "boolean_lenient"]
}
```

## Test Execution

### Simple Test Runner Structure

```pseudocode
class CCLTestRunner {
  implementation: CCLImplementation
  capabilities: TestCapabilities
  
  function run_test_suite(test_files) {
    all_results = []
    
    for test_file in test_files {
      tests = load_json(test_file)
      compatible_tests = filter_compatible_tests(tests, this.capabilities)
      
      for test in compatible_tests {
        result = this.execute_test(test)
        all_results.append(result)
      }
    }
    
    return generate_summary_report(all_results)
  }
  
  function execute_test(test) {
    start_time = now()
    
    try {
      result = this.execute_validation(test.validation, test.expected, test.input)
      status = "PASSED"
      message = "All assertions passed"
    } catch (error) {
      status = "FAILED" 
      message = error.message
      result = null
    }
    
    execution_time = now() - start_time
    return TestResult(test.name, status, result, message, execution_time)
  }
  
  function execute_validation(validation_name, expected, input) {
    switch validation_name {
      case "parse":
        actual = this.implementation.parse(input)
        assert_entries_equal(actual, expected.entries, expected.count)
        return actual
        
      case "make_objects":
        entries = this.implementation.parse(input)
        actual = this.implementation.make_objects(entries)
        assert_equal(actual, expected.object)
        return actual
        
      case "get_string":
        entries = this.implementation.parse(input)
        obj = this.implementation.make_objects(entries)
        actual = this.implementation.get_string(obj, expected.args)
        assert_equal(actual, expected.value)
        return actual
        
      // ... handle other validation types
    }
  }
}
```

## CLI Integration Patterns

### Generation Filtering

Use CLI flags to generate only compatible tests:

```bash
# Level 1: Parse only
ccl-test-runner generate --run-only function:parse

# Level 3: Parse + Objects, skip problematic behaviors  
ccl-test-runner generate \
  --run-only function:parse,function:make_objects \
  --skip-tags behavior:strict_spacing,behavior:tabs_preserve,behavior:crlf_preserve_literal

# Level 4: Full typed access
ccl-test-runner generate \
  --run-only function:parse,function:make_objects,function:get_string,function:get_int,function:get_bool \
  --skip-tags behavior:strict_spacing,behavior:tabs_preserve
```

### Test Execution

Run the generated compatible tests:

```bash
# Run all generated tests
ccl-test-runner test

# Run specific levels
ccl-test-runner test --levels 1,3

# Run with verbose output
ccl-test-runner test --format verbose
```

## Configuration-Driven Approach

### Test Configuration

```json
{
  "implementation_name": "my_ccl_parser",
  "capabilities": {
    "functions": ["parse", "make_objects", "get_string", "get_int"],
    "features": ["dotted_keys", "empty_keys"],
    "behaviors": ["crlf_normalize_to_lf", "boolean_lenient"]
  },
  "test_selection": {
    "skip_behaviors": ["strict_spacing", "tabs_preserve", "crlf_preserve_literal"],
    "skip_features": ["unicode", "multiline"],
    "skip_variants": ["proposed_behavior"]
  }
}
```

### Usage

```pseudocode
function create_test_runner_from_config(config_file) {
  config = load_json(config_file)
  capabilities = config.capabilities
  
  // Generate CLI args for test generation
  run_only_tags = []
  run_only_tags.extend("function:" + f for f in capabilities.functions)
  run_only_tags.extend("feature:" + f for f in capabilities.features) 
  run_only_tags.extend("behavior:" + b for b in capabilities.behaviors)
  
  skip_tags = []
  skip_tags.extend("behavior:" + b for b in config.test_selection.skip_behaviors)
  skip_tags.extend("feature:" + f for f in config.test_selection.skip_features)
  
  return TestRunner(capabilities, run_only_tags, skip_tags)
}
```

## Result Reporting

### Test Result Types

```pseudocode
enum TestStatus {
  PASSED,   // Test executed and all assertions passed
  FAILED,   // Test executed but assertions failed
  SKIPPED   // Test skipped due to missing requirements
}

type TestResult = {
  test_name: string,
  status: TestStatus,
  message: string,
  execution_time: duration,
  metadata: {
    validation: string,
    functions: string[],
    features: string[],
    behaviors: string[],
    level: int
  }
}
```

### Summary Reporting

```pseudocode
function generate_summary_report(results) {
  passed = results.count(r => r.status == "PASSED")
  failed = results.count(r => r.status == "FAILED") 
  skipped = results.count(r => r.status == "SKIPPED")
  total = results.length
  
  return {
    total_tests: total,
    passed: passed,
    failed: failed,
    skipped: skipped,
    success_rate: passed / (passed + failed) * 100,
    coverage_rate: (passed + failed) / total * 100,
    
    by_level: group_by_level(results),
    by_function: group_by_function(results),
    
    failed_tests: results.filter(r => r.status == "FAILED"),
    execution_time: sum(r.execution_time for r in results)
  }
}
```

## Progressive Implementation Workflow

### 1. Start with Level 1
```bash
# Generate and run basic parsing tests
ccl-test-runner generate --run-only function:parse
ccl-test-runner test --levels 1
```

### 2. Add Object Construction  
```bash
# Add Level 3 functionality
ccl-test-runner generate --run-only function:parse,function:make_objects
ccl-test-runner test --levels 1,3
```

### 3. Add Typed Access
```bash
# Add Level 4 functionality
ccl-test-runner generate --run-only function:parse,function:make_objects,function:get_string,function:get_int
ccl-test-runner test --levels 1,3,4
```

### 4. Handle Behaviors
```bash
# Skip conflicting behaviors as you encounter them
ccl-test-runner generate \
  --run-only function:parse,function:make_objects,function:get_string \
  --skip-tags behavior:strict_spacing,behavior:tabs_preserve
```

## Common Patterns

### Mock Implementation Testing
```bash
# Generate tests that match current mock capabilities
ccl-test-runner generate \
  --run-only function:parse,function:make_objects,function:get_string,function:get_int,function:get_bool \
  --skip-tags behavior:strict_spacing,behavior:tabs_preserve,behavior:crlf_preserve_literal
```

### Production Implementation Testing  
```bash
# Generate comprehensive test suite
ccl-test-runner generate
ccl-test-runner test --format json > results.json
```

### Feature Development Testing
```bash
# Test specific features during development
ccl-test-runner generate --run-only feature:comments
ccl-test-runner test --features comments
```

This simplified approach eliminates the complexity of partial validation execution while providing clear, predictable test behavior that matches the flat format's design.