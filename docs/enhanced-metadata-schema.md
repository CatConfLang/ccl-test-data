# Enhanced Metadata Schema for LLM Optimization

## Current State Analysis

### Existing JSON Structure (v2.0 Schema)
```json
{
  "$schema": "./schema.json",
  "suite": "CCL Essential Parsing (Validation Format)",
  "version": "2.0", 
  "description": "Core parsing functionality - minimum viable CCL implementation",
  "tests": [
    {
      "name": "basic_pairs",
      "input": "key = value",
      "validations": { "parse": { "count": 2, "expected": [...] } },
      "meta": {
        "tags": ["function:parse"],
        "level": 1,
        "feature": "parsing"
      }
    }
  ]
}
```

## Enhanced Schema Design

### File-Level Metadata Enhancement
```json
{
  "$schema": "./schema.json",
  "suite": "CCL Essential Parsing (Validation Format)",
  "version": "2.1",
  "description": "Core parsing functionality - minimum viable CCL implementation",
  "llm_metadata": {
    "llm_description": "Level 1 CCL parsing tests - converts text to flat key-value pairs without object construction",
    "complexity_level": "Level 1",
    "prerequisite_tests": [],
    "related_functions": ["parse", "load"],
    "implementation_notes": "Start here for new CCL implementations. Validates core parsing without advanced features.",
    "test_count": 45,
    "assertion_count": 89,
    "cross_references": {
      "gleam_implementation": "packages/ccl/src/ccl.gleam#parse",
      "go_reference": "internal/mock/ccl.go#Parse",
      "documentation": "https://ccl.tylerbutler.com/api-reference#parsing"
    }
  },
  "tests": [...]
}
```

### Test-Level Metadata Enhancement
```json
{
  "name": "basic_pairs",
  "input": "name = Alice\nage = 42", 
  "validations": {
    "parse": {
      "count": 2,
      "expected": [
        {"key": "name", "value": "Alice"},
        {"key": "age", "value": "42"}
      ]
    }
  },
  "meta": {
    "tags": ["function:parse"],
    "level": 1,
    "feature": "parsing",
    "llm_guidance": {
      "test_purpose": "Validates basic key-value pair parsing with string and numeric values",
      "implementation_focus": "Ensure parse() function correctly splits on '=' and handles multiple lines",
      "common_pitfalls": ["Don't convert '42' to integer during parsing - values remain strings at Level 1"],
      "related_tests": ["equals_in_values", "whitespace_handling"],
      "complexity_indicators": {
        "parsing_difficulty": "trivial",
        "edge_cases": "none", 
        "prerequisites": []
      }
    }
  }
}
```

## Schema Enhancement Strategy

### Phase 1: File-Level Metadata (All 10 JSON Files)

#### api-essential-parsing.json
```json
"llm_metadata": {
  "llm_description": "Level 1 CCL parsing tests - basic key-value pair extraction from text",
  "complexity_level": "Level 1",
  "prerequisite_tests": [],
  "related_functions": ["parse", "load"],
  "implementation_notes": "Start here for new CCL implementations. Foundation parsing without object construction.",
  "test_count": 45,
  "assertion_count": 89,
  "learning_path": "Begin → api-processing.json → api-object-construction.json → api-typed-access.json"
}
```

#### api-comprehensive-parsing.json  
```json
"llm_metadata": {
  "llm_description": "Level 1 CCL parsing edge cases - whitespace, unicode, complex syntax validation",
  "complexity_level": "Level 1+",
  "prerequisite_tests": ["api-essential-parsing.json"],
  "related_functions": ["parse", "load"],
  "implementation_notes": "Advanced Level 1 parsing. Handle edge cases after basic parsing works.",
  "test_count": 23,
  "assertion_count": 47
}
```

#### api-processing.json
```json
"llm_metadata": {
  "llm_description": "Level 2 CCL entry processing - composition, filtering, and advanced text handling",
  "complexity_level": "Level 2", 
  "prerequisite_tests": ["api-essential-parsing.json"],
  "related_functions": ["filter", "compose", "expand-dotted"],
  "implementation_notes": "Level 2 functionality. Requires working Level 1 parsing foundation.",
  "test_count": 67,
  "assertion_count": 134
}
```

#### api-object-construction.json
```json
"llm_metadata": {
  "llm_description": "Level 3 CCL object construction - converts flat entries to nested objects",
  "complexity_level": "Level 3",
  "prerequisite_tests": ["api-essential-parsing.json", "api-processing.json"],
  "related_functions": ["make-objects"],
  "implementation_notes": "Level 3 core functionality. Transforms flat key-value pairs into hierarchical objects.",
  "test_count": 89, 
  "assertion_count": 178
}
```

#### api-typed-access.json
```json
"llm_metadata": {
  "llm_description": "Level 4 CCL typed access - type-aware value extraction with conversion",
  "complexity_level": "Level 4",
  "prerequisite_tests": ["api-essential-parsing.json", "api-object-construction.json"],
  "related_functions": ["get-string", "get-int", "get-bool", "get-float", "get-list"],
  "implementation_notes": "Level 4 completion. Requires Level 3 object construction for navigation.",
  "test_count": 151,
  "assertion_count": 302
}
```

#### api-comments.json
```json
"llm_metadata": {
  "llm_description": "CCL comment syntax handling - '/=' prefix comment processing",
  "complexity_level": "Level 2 Feature",
  "prerequisite_tests": ["api-essential-parsing.json"],
  "related_functions": ["parse", "filter"],
  "implementation_notes": "Optional feature. Add after basic parsing works. Filter '/=' comments.",
  "feature_flag": "feature:comments"
}
```

#### api-dotted-keys.json
```json
"llm_metadata": {
  "llm_description": "CCL dotted key expansion - converts 'foo.bar.baz' to nested object structure",
  "complexity_level": "Level 3 Feature",
  "prerequisite_tests": ["api-essential-parsing.json", "api-object-construction.json"],
  "related_functions": ["make-objects", "expand-dotted"],
  "implementation_notes": "Advanced Level 3 feature. Requires object construction foundation.",
  "feature_flag": "feature:dotted-keys"
}
```

### Phase 2: Test-Level Metadata (Selected High-Value Tests)

Priority tests for detailed LLM guidance:
1. **First test in each file** - Entry point guidance
2. **Complex tests** - Multi-step implementation guidance  
3. **Common failure tests** - Pitfall documentation
4. **Edge case tests** - Boundary condition guidance

#### Implementation Pattern
```json
"llm_guidance": {
  "test_purpose": "One-line description of what this test validates",
  "implementation_focus": "Key implementation details for this specific test",
  "common_pitfalls": ["Array of common mistakes developers make"],
  "debugging_hints": ["Suggestions for when this test fails"],
  "related_tests": ["Other tests that build on or relate to this one"],
  "complexity_indicators": {
    "parsing_difficulty": "trivial|easy|moderate|complex|advanced",
    "edge_cases": "none|few|moderate|many",
    "prerequisites": ["List of concepts/functions that must work first"]
  }
}
```

## Cross-Reference Integration

### Repository Linking Schema
```json
"cross_references": {
  "gleam_implementation": {
    "package": "ccl",
    "file": "src/ccl.gleam", 
    "function": "parse",
    "line_reference": "#L45-L60"
  },
  "go_reference": {
    "file": "internal/mock/ccl.go",
    "function": "Parse",
    "line_reference": "#L23-L45"
  },
  "documentation": {
    "api_reference": "https://ccl.tylerbutler.com/api-reference#parsing",
    "implementation_guide": "https://ccl.tylerbutler.com/implementing-ccl",
    "parsing_algorithm": "https://ccl.tylerbutler.com/parsing-algorithm"
  },
  "related_test_files": [
    "api-comprehensive-parsing.json",
    "api-processing.json"
  ]
}
```

## Implementation Strategy

### Automatic Enhancement Script
```go
// scripts/enhance-metadata.go
func enhanceTestFile(filename string) error {
  // 1. Load existing JSON test file
  // 2. Analyze current metadata and test structure
  // 3. Generate appropriate llm_metadata based on patterns
  // 4. Add cross_references based on repository analysis
  // 5. Validate enhanced schema against updated schema.json
  // 6. Write enhanced file with preserved formatting
}
```

### Validation Integration
```bash
# Add to justfile
enhance-metadata:
  go run scripts/enhance-metadata.go tests/*.json

validate-enhanced:
  go run scripts/validate-schema.go tests/*.json --enhanced

# Integration with existing workflow  
test: validate-enhanced generate test-generated
```

### Schema Versioning
- **Current**: v2.0 (existing structure)
- **Enhanced**: v2.1 (adds llm_metadata, preserves existing)
- **Backward Compatible**: Existing tooling continues to work
- **Progressive Enhancement**: Add llm_metadata incrementally

## Success Metrics

### Quantitative Targets
- ✅ 100% of 10 JSON test files have file-level llm_metadata
- ✅ 50% of high-priority tests have test-level llm_guidance  
- ✅ All files include cross_references to related implementations
- ✅ Enhanced schema validates without errors
- ✅ Existing Go test runner continues to work unchanged

### Quality Standards
- **Clarity**: LLM descriptions are precise and actionable
- **Completeness**: All implementation levels covered with guidance
- **Consistency**: Metadata patterns consistent across files
- **Usefulness**: AI tools can effectively navigate and understand test progression
- **Maintainability**: Enhancement process is automatable and sustainable