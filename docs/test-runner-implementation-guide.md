# Test Runner Implementation Guide

This guide provides patterns and strategies for implementing test runners that work with the CCL test suite's separate typed fields architecture. It addresses real-world scenarios like progressive implementation, mixed function requirements, and robust result reporting.

## Test Execution Strategies

When a test contains multiple validations but your implementation doesn't support all required functions, you have several execution strategies to choose from.

### 1. Partial Validation Execution (Recommended)

**Strategy**: Execute only the validations for implemented functions, skip unimplemented ones.

```json
{
  "name": "mixed_requirements_test",
  "input": "database.host = localhost\ndatabase.port = 5432",
  "validations": {
    "parse": {
      "count": 2,
      "expected": [
        {"key": "database.host", "value": "localhost"},
        {"key": "database.port", "value": "5432"}
      ]
    },
    "expand_dotted": {
      "count": 1,
      "expected": [
        {"key": "database", "value": "\n  host = localhost\n  port = 5432"}
      ]
    },
    "build_hierarchy": {
      "count": 1,
      "expected": {
        "database": {"host": "localhost", "port": "5432"}
      }
    }
  },
  "functions": ["parse", "expand_dotted", "build_hierarchy"],
  "features": [],
  "behaviors": [],
  "variants": [],
  "level": 1
}
```

**If `expand_dotted` is unimplemented:**
- ✅ Execute `parse` validation
- ⏭️ Skip `expand_dotted` validation 
- ✅ Execute `build_hierarchy` validation
- 📊 Report: **PARTIAL** (2/3 validations executed)

**Benefits:**
- Tests implemented functionality, providing valuable feedback
- Supports progressive development workflows
- Provides partial progress rather than complete silence

**Trade-offs:**
- May miss integration issues between implemented and unimplemented parts
- Requires more complex result interpretation

### 2. Skip Entire Test (Conservative)

**Strategy**: Skip the entire test if any required function is unimplemented.

**Implementation:**
```pseudocode
function should_skip_test(test, implemented_functions) {
  for function_name in test.functions {
    if function_name not in implemented_functions {
      return true
    }
  }
  return false
}
```

**Benefits:**
- Clean, no partial results to confuse interpretation
- Ensures test integrity - only run tests that can be fully validated
- Simple implementation and clear semantics

**Trade-offs:**
- Loses valuable testing of working functionality
- Not progressive-implementation friendly
- Can result in too many skipped tests during development

### 3. Fail Fast (Strict)

**Strategy**: Fail immediately when encountering an unimplemented function.

**Implementation:**
```pseudocode
function execute_test(test, implementation) {
  for validation_name in test.validations.keys() {
    if not implementation.supports(validation_name) {
      throw UnimplementedFunctionError(validation_name)
    }
    // Execute validation...
  }
}
```

**Benefits:**
- Forces complete implementation before testing
- Clear signal that implementation is incomplete
- Strict enforcement of completeness

**Trade-offs:**
- Not suitable for progressive implementation workflows
- Discourages incremental development
- Can block all testing until everything is implemented

## Progressive Implementation Patterns

### Function-Based Filtering

Use separate typed fields to filter tests based on implemented functions:

```pseudocode
// Pre-filter tests during test discovery
function get_runnable_tests(all_tests, implemented_functions) {
  runnable_tests = []
  
  for test in all_tests {
    required_functions = test.functions
    
    if all_required_functions_implemented(required_functions, implemented_functions) {
      runnable_tests.append(test)
    }
  }
  
  return runnable_tests
}

// Function extraction is no longer needed with separate fields
// Functions are directly available in test.functions array
```

### Runtime Validation Skipping

Skip individual validations at runtime:

```pseudocode
function execute_test_validations(test, implementation) {
  results = []
  
  for validation_name, validation_spec in test.validations {
    if implementation.supports(validation_name) {
      try {
        result = execute_validation(validation_name, validation_spec, test.input)
        results.append(ValidationResult(validation_name, "passed", result))
      } catch (error) {
        results.append(ValidationResult(validation_name, "failed", error))
      }
    } else {
      results.append(ValidationResult(validation_name, "skipped", "function not implemented"))
    }
  }
  
  return determine_overall_result(results)
}

function determine_overall_result(validation_results) {
  passed = validation_results.count(r => r.status == "passed")
  failed = validation_results.count(r => r.status == "failed")
  skipped = validation_results.count(r => r.status == "skipped")
  
  if failed > 0 {
    return "failed"
  } else if passed > 0 and skipped > 0 {
    return "partial"
  } else if passed > 0 {
    return "passed"
  } else {
    return "skipped"
  }
}
```

### Feature-Based Test Organization

Organize test execution based on features and capabilities:

```pseudocode
type TestCapabilities = {
  functions: Set<string>,     // "parse", "build_hierarchy", "get_string", etc.
  features: Set<string>,      // "comments", "dotted_keys", "unicode", etc.
  behaviors: Set<string>      // "crlf_normalize", "tabs_preserve", etc.
}

function filter_tests_by_capabilities(tests, capabilities) {
  compatible_tests = []
  
  for test in tests {
    if test_is_compatible(test, capabilities) {
      compatible_tests.append(test)
    }
  }
  
  return compatible_tests
}

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
  
  // Check variant compatibility
  for variant_name in test.variants {
    if variant_name not in capabilities.variants {
      return false
    }
  }
  
  return true
}
```

## Standard Test Result Schema

### Test Result Types

Define clear result categories for consistent reporting:

```pseudocode
enum TestStatus {
  PASSED,     // All validations executed and passed
  FAILED,     // At least one validation failed
  PARTIAL,    // Some validations passed, others skipped
  SKIPPED     // Entire test skipped due to missing requirements
}

enum ValidationStatus {
  PASSED,          // Validation executed and passed
  FAILED,          // Validation executed but failed
  SKIPPED,         // Validation skipped (function not implemented)
  ERROR            // Validation couldn't execute due to test runner error
}
```

### Detailed Result Structure

```pseudocode
type ValidationResult = {
  name: string,
  status: ValidationStatus,
  message: string,
  expected?: any,
  actual?: any,
  execution_time?: duration
}

type TestResult = {
  test_name: string,
  status: TestStatus,
  validations: ValidationResult[],
  summary: string,
  execution_time: duration,
  metadata: {
    functions: string[],
    features: string[],
    behaviors: string[],
    variants: string[],
    functions_executed: string[],
    functions_skipped: string[]
  }
}
```

### Result Reporting Examples

#### Successful Partial Execution
```json
{
  "test_name": "basic_dotted_key_parsing",
  "status": "PARTIAL",
  "validations": [
    {
      "name": "parse",
      "status": "PASSED",
      "message": "All 2 assertions passed"
    },
    {
      "name": "expand_dotted", 
      "status": "SKIPPED",
      "message": "Function not implemented"
    },
    {
      "name": "build_hierarchy",
      "status": "PASSED", 
      "message": "All 1 assertions passed"
    }
  ],
  "summary": "2/3 validations executed successfully",
  "metadata": {
    "functions": ["parse", "expand_dotted", "build_hierarchy"],
    "features": [],
    "behaviors": [],
    "variants": [],
    "functions_executed": ["parse", "build_hierarchy"],
    "functions_skipped": ["expand_dotted"]
  }
}
```

#### Complete Skip
```json
{
  "test_name": "advanced_typed_access",
  "status": "SKIPPED",
  "validations": [],
  "summary": "Test requires unimplemented functions: get_string, get_int",
  "metadata": {
    "functions": ["parse", "build_hierarchy", "get_string", "get_int"],
    "features": [],
    "behaviors": [],
    "variants": [],
    "functions_executed": [],
    "functions_skipped": ["get_string", "get_int"]
  }
}
```

#### Test Failure
```json
{
  "test_name": "basic_parsing",
  "status": "FAILED",
  "validations": [
    {
      "name": "parse",
      "status": "FAILED",
      "message": "Expected 2 entries, got 1",
      "expected": [
        {"key": "name", "value": "Alice"},
        {"key": "age", "value": "42"}
      ],
      "actual": [
        {"key": "name", "value": "Alice"}
      ]
    }
  ],
  "summary": "Parse validation failed",
  "metadata": {
    "functions": ["parse"],
    "features": [],
    "behaviors": [],
    "variants": [],
    "functions_executed": ["parse"],
    "functions_skipped": []
  }
}
```

## Implementation Examples

### Basic Test Runner Structure

```pseudocode
class CCLTestRunner {
  implementation: CCLImplementation
  capabilities: TestCapabilities
  
  constructor(implementation, capabilities) {
    this.implementation = implementation
    this.capabilities = capabilities
  }
  
  function run_test_suite(test_files) {
    all_results = []
    
    for test_file in test_files {
      tests = load_json(test_file)
      compatible_tests = filter_tests_by_capabilities(tests, this.capabilities)
      
      for test in compatible_tests {
        result = this.execute_test(test)
        all_results.append(result)
      }
    }
    
    return generate_summary_report(all_results)
  }
  
  function execute_test(test) {
    validation_results = []
    start_time = now()
    
    for validation_name, validation_spec in test.validations {
      if this.supports_validation(validation_name) {
        validation_result = this.execute_validation(validation_name, validation_spec, test.input)
        validation_results.append(validation_result)
      } else {
        skip_result = ValidationResult(validation_name, "SKIPPED", "Function not implemented")
        validation_results.append(skip_result)
      }
    }
    
    execution_time = now() - start_time
    overall_status = determine_overall_status(validation_results)
    
    return TestResult(test.name, overall_status, validation_results, execution_time)
  }
  
  function supports_validation(validation_name) {
    return validation_name in this.capabilities.functions
  }
  
  function execute_validation(validation_name, validation_spec, input) {
    try {
      switch validation_name {
        case "parse":
          actual = this.implementation.parse(input)
          assert_equal(actual, validation_spec.expected)
          return ValidationResult("parse", "PASSED", "All assertions passed")
          
        case "build_hierarchy":
          entries = this.implementation.parse(input)
          actual = this.implementation.build_hierarchy(entries)
          assert_equal(actual, validation_spec.expected)
          return ValidationResult("build_hierarchy", "PASSED", "All assertions passed")
          
        case "get_string":
          entries = this.implementation.parse(input)
          obj = this.implementation.build_hierarchy(entries)
          for case in validation_spec.cases {
            actual = this.implementation.get_string(obj, case.args)
            assert_equal(actual, case.expected)
          }
          return ValidationResult("get_string", "PASSED", "All assertions passed")
          
        // ... handle other validation types
      }
    } catch (error) {
      return ValidationResult(validation_name, "FAILED", error.message)
    }
  }
}
```

### Configuration-Driven Test Execution

```pseudocode
type TestConfiguration = {
  enabled_functions: string[],
  enabled_features: string[],
  enabled_behaviors: string[],
  execution_strategy: "partial" | "skip_entire" | "fail_fast",
  
  // Alternative filtering strategies (use one approach, not both)
  filter_mode: "inclusive" | "exclusive",
  
  // Inclusive filtering (only run tests with these)
  only_functions?: string[],
  only_features?: string[],
  
  // Exclusive filtering (run everything except these)  
  skip_functions?: string[],
  skip_features?: string[]
}

function create_test_runner(implementation, config) {
  capabilities = TestCapabilities(
    functions: Set(config.enabled_functions),
    features: Set(config.enabled_features), 
    behaviors: Set(config.enabled_behaviors)
  )
  
  return CCLTestRunner(implementation, capabilities, config.execution_strategy)
}

// Usage examples:

// Minimal implementation - only core parsing
minimal_config = TestConfiguration(
  enabled_functions: ["parse", "build_hierarchy"],
  enabled_features: [],
  enabled_behaviors: [],
  execution_strategy: "partial"
)

// Standard implementation - parsing + typed access
standard_config = TestConfiguration(
  enabled_functions: ["parse", "build_hierarchy", "get_string", "get_int", "get_bool"],
  enabled_features: ["comments"],
  enabled_behaviors: ["crlf_normalize"],
  execution_strategy: "partial"
)

// Experimental implementation - includes experimental features
experimental_config = TestConfiguration(
  enabled_functions: ["parse", "expand_dotted", "build_hierarchy"],
  enabled_features: ["dotted_keys", "comments"],
  enabled_behaviors: ["crlf_normalize"],
  execution_strategy: "partial",
  filter_mode: "inclusive",
  only_features: ["experimental_dotted_keys"]
)
```

## Test Discovery and Filtering

### Tag-Based Test Discovery

```pseudocode
function discover_tests(test_directory, filter_config) {
  all_test_files = find_json_files(test_directory)
  discovered_tests = []
  
  for file in all_test_files {
    tests = load_json(file)
    
    for test in tests {
      if should_include_test(test, filter_config) {
        discovered_tests.append(test)
      }
    }
  }
  
  return discovered_tests
}

function should_include_test(test, filter_config) {
  // Check if all required functions are available
  for function in test.functions {
    if function not in filter_config.enabled_functions {
      return false
    }
  }
  
  // Apply filtering strategy
  if filter_config.filter_mode == "inclusive" {
    // Inclusive: only run tests with specified functions/features
    if filter_config.only_functions {
      has_required_function = false
      for only_function in filter_config.only_functions {
        if only_function in test.functions {
          has_required_function = true
          break
        }
      }
      if not has_required_function {
        return false
      }
    }
    
    if filter_config.only_features {
      has_required_feature = false
      for only_feature in filter_config.only_features {
        if only_feature in test.features {
          has_required_feature = true
          break
        }
      }
      if not has_required_feature {
        return false
      }
    }
  } else if filter_config.filter_mode == "exclusive" {
    // Exclusive: skip tests with specified functions/features
    if filter_config.skip_functions {
      for skip_function in filter_config.skip_functions {
        if skip_function in test.functions {
          return false
        }
      }
    }
    
    if filter_config.skip_features {
      for skip_feature in filter_config.skip_features {
        if skip_feature in test.features {
          return false
        }
      }
    }
  }
  
  return true
}
```

### Smart Test Prioritization

```pseudocode
function prioritize_tests(tests, capabilities) {
  // Categorize tests by implementation readiness
  categories = {
    fully_supported: [],    // All required functions implemented
    partially_supported: [], // Some functions implemented
    unsupported: []         // No functions implemented
  }
  
  for test in tests {
    required_functions = test.functions
    supported_count = count_supported_functions(required_functions, capabilities)
    
    if supported_count == required_functions.length {
      categories.fully_supported.append(test)
    } else if supported_count > 0 {
      categories.partially_supported.append(test)
    } else {
      categories.unsupported.append(test)
    }
  }
  
  // Return in priority order: fully supported first, then partial, then unsupported
  return categories.fully_supported + categories.partially_supported + categories.unsupported
}
```

## Error Handling and Edge Cases

### Graceful Degradation

```pseudocode
function execute_validation_safely(validation_name, validation_spec, input, implementation) {
  try {
    return execute_validation(validation_name, validation_spec, input, implementation)
  } catch (UnimplementedFunctionError e) {
    return ValidationResult(validation_name, "SKIPPED", "Function not implemented: " + e.function_name)
  } catch (ParseError e) {
    return ValidationResult(validation_name, "FAILED", "Parse error: " + e.message)
  } catch (ValidationError e) {
    return ValidationResult(validation_name, "FAILED", "Validation failed: " + e.message)
  } catch (TestRunnerError e) {
    return ValidationResult(validation_name, "ERROR", "Test runner error: " + e.message)
  }
}
```

### Dependency Resolution

Some validations may depend on others. Handle dependencies gracefully:

```pseudocode
function execute_test_with_dependencies(test, implementation) {
  validation_order = resolve_validation_dependencies(test.validations.keys())
  results = {}
  
  for validation_name in validation_order {
    dependencies_met = check_dependencies(validation_name, results)
    
    if not dependencies_met {
      results[validation_name] = ValidationResult(
        validation_name, 
        "SKIPPED", 
        "Dependencies not met"
      )
      continue
    }
    
    if implementation.supports(validation_name) {
      results[validation_name] = execute_validation(validation_name, test.validations[validation_name], test.input)
    } else {
      results[validation_name] = ValidationResult(validation_name, "SKIPPED", "Function not implemented")
    }
  }
  
  return results
}

function resolve_validation_dependencies(validation_names) {
  // Define dependency relationships
  dependencies = {
    "build_hierarchy": ["parse"],
    "get_string": ["parse", "build_hierarchy"],
    "get_int": ["parse", "build_hierarchy"],
    "expand_dotted": ["parse"]
  }
  
  // Return topologically sorted order
  return topological_sort(validation_names, dependencies)
}
```

This guide provides a comprehensive foundation for implementing robust test runners that can handle the complexities of progressive CCL implementation while providing clear, actionable feedback to developers.