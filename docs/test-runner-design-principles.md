# Test Runner Configuration Design: Avoiding Hidden Dependencies and Masked Conflicts

**A Guide for Test Suite Implementers**

## Executive Summary

Multi-layered test filtering systems can mask logical configuration errors, leading to false confidence in test results. This document provides design principles and JSON schema recommendations to prevent configuration inconsistencies that can hide test failures.

## The Problem: Masked Configuration Conflicts

### Example Scenario
```json
{
  "behaviors": ["crlf_normalize_to_lf", "crlf_preserve_literal"],
  "variants": ["reference_compliant"],
  "functions": ["parse", "round_trip"]
}
```

**What happened:**
- Configuration claimed to support mutually exclusive behaviors
- Tests for `crlf_normalize_to_lf` had `variant: "proposed_behavior"` (disabled)
- Tests for `crlf_preserve_literal` had `variant: "reference_compliant"` (enabled)
- All tests passed because conflicting tests were filtered by variant, not behavior
- Behavior conflict was completely hidden

**Result:** False confidence in implementation correctness due to masked logical inconsistency.

## Core Design Principles

### 1. Fail-Fast Configuration Validation

**Principle:** Invalid configurations should be impossible or immediately obvious.

**Implementation:**
```typescript
interface ConfigValidationRule {
  check: (config: TestConfig) => ValidationResult;
  message: string;
  severity: 'error' | 'warning';
}

const MUTUAL_EXCLUSION_RULES: ConfigValidationRule[] = [
  {
    check: (config) => !(hasAll(config.behaviors, ['crlf_normalize_to_lf', 'crlf_preserve_literal'])),
    message: "CRLF behaviors are mutually exclusive",
    severity: 'error'
  },
  {
    check: (config) => !(hasAll(config.behaviors, ['boolean_strict', 'boolean_lenient'])),
    message: "Boolean parsing behaviors are mutually exclusive",
    severity: 'error'
  }
];

function validateConfig(config: TestConfig): void {
  const errors = MUTUAL_EXCLUSION_RULES
    .filter(rule => !rule.check(config))
    .filter(rule => rule.severity === 'error');

  if (errors.length > 0) {
    throw new ConfigurationError(errors.map(e => e.message).join('; '));
  }
}
```

### 2. Single Source of Truth

**Principle:** Use one authoritative decision function instead of layered filtering.

**Anti-pattern:**
```typescript
// DON'T: Multiple filter stages that can mask each other
const variantFiltered = tests.filter(t => variantMatches(t, config));
const behaviorFiltered = variantFiltered.filter(t => behaviorMatches(t, config));
const functionFiltered = behaviorFiltered.filter(t => functionMatches(t, config));
```

**Better approach:**
```typescript
// DO: Single decision function with explicit logic
function shouldRunTest(test: TestCase, config: TestConfig): TestDecision {
  const reasons: SkipReason[] = [];

  if (!variantMatches(test, config)) {
    reasons.push({ type: 'variant_mismatch', details: test.variants });
  }

  if (!behaviorMatches(test, config)) {
    reasons.push({ type: 'behavior_conflict', details: findConflicts(test, config) });
  }

  if (!functionMatches(test, config)) {
    reasons.push({ type: 'missing_function', details: getMissingFunctions(test, config) });
  }

  return reasons.length === 0
    ? { run: true }
    : { run: false, reasons };
}
```

### 3. Explicit Dependency Modeling

**Principle:** Make conflicting options impossible to specify simultaneously.

**Implementation:**
```typescript
// Use discriminated unions for mutually exclusive choices
type CrlfBehavior =
  | { type: 'normalize_to_lf' }
  | { type: 'preserve_literal' };

type BooleanBehavior =
  | { type: 'strict' }
  | { type: 'lenient' };

interface ValidatedConfig {
  functions: string[];
  features: string[];
  crlf: CrlfBehavior;           // Impossible to have both
  boolean: BooleanBehavior;     // Impossible to have both
  variants: string[];
}
```

### 4. Configuration Coverage Analysis

**Principle:** Detect when configuration choices are not being tested.

**Implementation:**
```typescript
function analyzeConfigurationCoverage(tests: TestCase[], config: TestConfig): CoverageReport {
  const runTests = tests.filter(t => shouldRunTest(t, config).run);
  const skippedTests = tests.filter(t => !shouldRunTest(t, config).run);

  const skipReasons = skippedTests.map(t => shouldRunTest(t, config).reasons).flat();
  const reasonsByType = groupBy(skipReasons, r => r.type);

  // Detect potentially masked conflicts
  const allVariantSkips = reasonsByType['variant_mismatch']?.length === skippedTests.length;
  const hasBehaviorConflicts = reasonsByType['behavior_conflict']?.length > 0;

  if (allVariantSkips && hasBehaviorConflicts) {
    warnings.push("Behavior conflicts may be masked by variant filtering");
  }

  return { runTests, skippedTests, warnings };
}
```

### 5. Hierarchical Validation with Explicit Precedence

**Principle:** When multiple filters must coexist, make precedence explicit and documented.

**Implementation:**
```typescript
enum FilterPrecedence {
  VARIANT = 1,      // Highest precedence - architectural choice
  BEHAVIOR = 2,     // Implementation capabilities
  FUNCTION = 3,     // Feature completeness
  FEATURE = 4       // Lowest precedence - optional functionality
}

interface FilterStage {
  precedence: FilterPrecedence;
  name: string;
  apply: (test: TestCase, config: TestConfig) => FilterResult;
}

const FILTER_PIPELINE: FilterStage[] = [
  {
    precedence: FilterPrecedence.VARIANT,
    name: 'variant_filter',
    apply: (test, config) => variantFilter(test, config)
  },
  // ... other filters in precedence order
];
```

## JSON Schema Design Recommendations

### 1. Enforce Mutual Exclusion with `oneOf`

```json
{
  "type": "object",
  "properties": {
    "behaviors": {
      "type": "array",
      "items": {
        "anyOf": [
          {
            "type": "object",
            "properties": {
              "crlf": {
                "oneOf": [
                  { "const": "normalize_to_lf" },
                  { "const": "preserve_literal" }
                ]
              }
            }
          },
          {
            "type": "object",
            "properties": {
              "boolean_parsing": {
                "oneOf": [
                  { "const": "strict" },
                  { "const": "lenient" }
                ]
              }
            }
          }
        ]
      }
    }
  }
}
```

### 2. Use Custom Validation for Complex Constraints

```json
{
  "type": "object",
  "properties": {
    "behaviors": {
      "type": "array",
      "items": { "type": "string" }
    }
  },
  "not": {
    "properties": {
      "behaviors": {
        "allOf": [
          { "contains": { "const": "crlf_normalize_to_lf" } },
          { "contains": { "const": "crlf_preserve_literal" } }
        ]
      }
    }
  },
  "errorMessage": {
    "not": "Cannot specify both crlf_normalize_to_lf and crlf_preserve_literal"
  }
}
```

### 3. Conditional Dependencies with `if/then/else`

```json
{
  "if": {
    "properties": {
      "behaviors": { "contains": { "const": "crlf_normalize_to_lf" } }
    }
  },
  "then": {
    "properties": {
      "variants": {
        "contains": { "const": "proposed_behavior" }
      }
    },
    "required": ["variants"]
  },
  "else": {
    "if": {
      "properties": {
        "behaviors": { "contains": { "const": "crlf_preserve_literal" } }
      }
    },
    "then": {
      "properties": {
        "variants": {
          "contains": { "const": "reference_compliant" }
        }
      },
      "required": ["variants"]
    }
  }
}
```

### 4. Structured Configuration Schema

Instead of flat arrays, use structured objects that prevent invalid combinations:

```json
{
  "type": "object",
  "properties": {
    "implementation_profile": {
      "type": "object",
      "properties": {
        "functions": {
          "type": "array",
          "items": { "enum": ["parse", "build_hierarchy", "get_string"] }
        },
        "line_ending_behavior": {
          "oneOf": [
            { "const": "normalize_crlf_to_lf" },
            { "const": "preserve_literal_crlf" }
          ]
        },
        "boolean_parsing": {
          "oneOf": [
            { "const": "strict" },
            { "const": "lenient" }
          ]
        },
        "supported_variants": {
          "type": "array",
          "items": { "enum": ["reference_compliant", "proposed_behavior"] }
        }
      },
      "required": ["line_ending_behavior", "boolean_parsing"]
    }
  }
}
```

### 5. Cross-Reference Validation

Use `$ref` and custom keywords to validate cross-references:

```json
{
  "$defs": {
    "behavior_sets": {
      "crlf_normalize": {
        "required_behaviors": ["crlf_normalize_to_lf"],
        "required_variants": ["proposed_behavior"],
        "conflicts_with": ["crlf_preserve_literal"]
      },
      "crlf_preserve": {
        "required_behaviors": ["crlf_preserve_literal"],
        "required_variants": ["reference_compliant"],
        "conflicts_with": ["crlf_normalize_to_lf"]
      }
    }
  },
  "properties": {
    "tests": {
      "type": "array",
      "items": {
        "if": {
          "properties": {
            "behaviors": { "contains": { "const": "crlf_normalize_to_lf" } }
          }
        },
        "then": {
          "properties": {
            "variants": { "contains": { "const": "proposed_behavior" } }
          }
        }
      }
    }
  }
}
```

## Implementation Checklist

### At Design Time
- [ ] Identify mutually exclusive configuration options
- [ ] Model conflicts explicitly in types/schemas
- [ ] Define clear precedence rules for multi-layer filtering
- [ ] Design single decision functions over layered filters

### At Startup Time
- [ ] Validate configuration consistency before running tests
- [ ] Generate compatibility matrix for behavior combinations
- [ ] Check that all test categories have runnable tests
- [ ] Warn about potentially masked conflicts

### At Runtime
- [ ] Track skip reasons for all filtered tests
- [ ] Report configuration coverage statistics
- [ ] Detect when variant filtering is masking behavior conflicts
- [ ] Generate actionable error messages for configuration issues

### Schema Design
- [ ] Use `oneOf` for mutually exclusive options
- [ ] Implement cross-field validation with `if/then/else`
- [ ] Structure configuration as objects rather than flat arrays
- [ ] Add custom validation for complex business rules
- [ ] Provide clear error messages for constraint violations

## Conclusion

Configuration complexity grows exponentially with the number of options and their interactions. By applying these principles—especially fail-fast validation and explicit dependency modeling—test runner implementers can prevent the hidden inconsistencies that lead to false confidence in test results.

The key insight is that **configuration errors should be impossible or immediately obvious**, not discovered through careful analysis of test execution patterns.

---

*This document is based on lessons learned from debugging masked conflicts in the CCL test suite, where mutually exclusive behaviors were enabled simultaneously but conflicts were hidden by variant filtering.*