# CCL Test Suite Architecture

This document provides a comprehensive technical overview of the CCL Test Suite architecture, including system design, data flow, component interactions, and cross-repository dependencies.

## System Overview

The CCL Test Suite is a multi-component testing framework designed to support progressive CCL implementation across programming languages. It uses a dual-format architecture with JSON-based test definitions and Go-based tooling.

### Core Design Principles

- **Language Agnostic**: Test data format independent of implementation language
- **Progressive Implementation**: Support for gradual CCL feature adoption
- **Dual Format Architecture**: Maintainable source format + implementation-friendly generated format
- **Feature-Based Testing: Structured metadata for precise test selection
- **Counted Assertions**: Self-validating tests with explicit assertion counts

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                    CCL Test Suite                           │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────────┐    ┌──────────────────┐              │
│  │  Source Tests   │───▶│   Generator      │              │
│  │  (JSON)         │    │                  │              │
│  │                 │    │ ┌──────────────┐ │              │
│  │ • Maintainable  │    │ │ Flat Format  │ │              │
│  │ • Multi-test    │    │ │ Generator    │ │              │
│  │ • Structured    │    │ └──────────────┘ │              │
│  │   Tags          │    │                  │              │
│  └─────────────────┘    │ ┌──────────────┐ │              │
│                         │ │ Go Test      │ │              │
│                         │ │ Generator    │ │              │
│                         │ └──────────────┘ │              │
│                         └──────────────────┘              │
│                                  │                         │
│                                  ▼                         │
│  ┌─────────────────┐    ┌──────────────────┐              │
│  │  Generated      │    │    Mock CCL      │              │
│  │  Tests          │───▶│  Implementation  │              │
│  │                 │    │                  │              │
│  │ • Flat Format   │    │ • Reference      │              │
│  │ • Type-Safe     │    │ • Reference      │              │
│  │ • 1:N Transform │    │ • Development    │              │
│  └─────────────────┘    └──────────────────┘              │
│                                  │                         │
│                                  ▼                         │
│  ┌─────────────────┐    ┌──────────────────┐              │
│  │  Test Runner    │    │   Statistics     │              │
│  │                 │    │                  │              │
│  │ • CLI Interface │    │ • Coverage       │              │
│  │ • Filtering     │    │ • Analysis       │              │
│  │ • Reporting     │    │ • Benchmarks     │              │
│  └─────────────────┘    └──────────────────┘              │
└─────────────────────────────────────────────────────────────┘
```

## Component Architecture

### 1. Data Layer

#### Source Test Format
- **Location**: `source_tests/api_*.json`
- **Purpose**: Human-maintainable test definitions
- **Structure**: Multi-validation tests with structured metadata
- **Schema**: `schemas/source-format.json`

```json
{
  "name": "test_name",
  "input": "CCL input text",
  "validations": {
    "validation_type": {
      "count": 1,
      "expected": "result"
    }
  },
  "meta": {
    "tags": ["function:parse", "feature:comments"],
    "feature": "parsing"
  }
}
```

#### Generated Test Format
- **Location**: `generated_tests/api_*.json`
- **Purpose**: Implementation-friendly flat format
- **Structure**: One test per validation (1:N transformation)
- **Schema**: `schemas/generated-format.json`

```json
{
  "name": "test_name_validation_type",
  "input": "CCL input text",
  "validation": "validation_type",
  "expected": {
    "count": 1,
    "result": "data"
  },
  "functions": ["parse"],
  "features": ["comments"]
}
```

### 2. Generation Layer

#### Test Generator (`internal/generator/`)
- **Primary Function**: Transform source format to generated formats
- **Key Components**:
  - `generator.go`: Core generation logic
  - `templates.go`: Go test template system
  - `pool.go`: Object pooling for efficiency

**Generation Process**:
```go
Source JSON → Parse Metadata → Apply Filters → Generate Go Tests
             ↓
     Extract Tags → Validate Conflicts → Output Flat JSON
```

#### Template System
- **Go Templates**: Generate test functions from JSON data
- **Type Safety**: Strongly-typed template data structures
- **Performance**: Object pooling for high-frequency allocations

### 3. Execution Layer

#### Mock CCL Implementation (`internal/mock/ccl.go`)
- **Purpose**: Reference implementation for test validation
- **Functions Implemented**: Core, Typed Access, Processing, Formatting
- **Features**: Comments, dotted keys, error handling

**CCL Function Groups**:
```
Core Functions
├── Parse(input) → []Entry
└── BuildHierarchy([]Entry) → map[string]interface{}

Typed Access Functions
├── GetString(obj, path) → string
├── GetInt(obj, path) → int
├── GetBool(obj, path) → bool
├── GetFloat(obj, path) → float64
└── GetList(obj, path) → []interface{}

Processing Functions
├── Filter([]Entry) → []Entry
├── Combine([]Entry) → []Entry
└── ExpandDotted([]Entry) → []Entry

Formatting Functions
└── CanonicalFormat(map[string]interface{}) → string
```

#### Test Runner (`cmd/ccl-test-runner/`)
- **CLI Interface**: Command-line tool for test operations
- **Filtering**: Function and feature-based test selection
- **Output Formats**: Pretty, table, verbose, JSON
- **Integration**: Works with Go testing framework

### 4. Analysis Layer

#### Statistics System (`internal/stats/`)
- **Collection**: Test coverage and distribution analysis
- **Metrics**: Function coverage, feature usage distribution
- **Output**: Human-readable and machine-readable formats

## Data Flow

### Test Generation Flow
```
1. Read Source Tests (JSON)
   ↓
2. Parse Metadata & Tags
   ↓
3. Apply Filtering Rules
   ↓
4. Generate Flat Format Tests
   ↓
5. Generate Go Test Files
   ↓
6. Execute Tests via Mock Implementation
```

### Test Execution Flow
```
1. CLI Command (ccl-test-runner test)
   ↓
2. Load Generated Go Tests
   ↓
3. Apply Runtime Filters (feature, function)
   ↓
4. Execute via Go Testing Framework
   ↓
5. Collect Results & Format Output
```

### Statistics Collection Flow
```
1. Read Source Test Files
   ↓
2. Extract Metadata & Count Assertions
   ↓
3. Aggregate by Function/Feature
   ↓
4. Generate Reports (Pretty/JSON)
```

## Progressive Implementation Strategy

### Function Group Implementation
The architecture supports progressive CCL implementation through function groups:

1. **Core Functions**: Basic parsing + object construction
2. **Typed Access Functions**: Type-safe value extraction
3. **Processing Functions**: Advanced entry manipulation
4. **Formatting Functions**: Standardized output generation

### Feature-Based Testing
Tests are tagged with required features, enabling:
- **Minimal Implementation**: Skip optional features
- **Progressive Feature Addition**: Add features incrementally
- **Behavior Choices**: Select implementation-specific behaviors

## Efficiency Architecture

### Object Pooling
- **Generator Pool**: Reuse heavy objects during generation
- **Template Pool**: Cache compiled templates
- **Entry Pool**: Reuse Entry structs during parsing

### Memory Management
- **Streaming Processing**: Process large test sets without full memory load
- **Lazy Loading**: Load test data on demand
- **Garbage Collection**: Explicit cleanup of temporary objects

## Extension Points

### Adding New Validation Types
1. **Define JSON Structure**: Add to source format schema
2. **Update Generator**: Add template generation logic
3. **Mock Implementation**: Add corresponding function
4. **Test Coverage**: Ensure comprehensive test coverage

### Adding New Features
1. **Feature Tags**: Add structured feature tags
2. **Test Cases**: Create feature-specific test cases
3. **Mock Support**: Implement in mock CCL
4. **Documentation**: Update feature documentation

### Custom Implementations
1. **Test Data**: Use generated flat format JSON
2. **Filtering**: Apply function/feature filters
3. **Validation**: Compare against expected results
4. **Integration**: Use standard test patterns

## Security Considerations

### Input Validation
- **JSON Schema**: Strict validation of test data structure
- **Template Safety**: Secure template generation (no code injection)
- **Path Traversal**: Safe file operations with path validation

### Test Isolation
- **Process Isolation**: Tests run in separate processes
- **State Management**: No shared state between tests
- **Resource Cleanup**: Automatic cleanup of temporary resources

## Quality Assurance

### Validation Pipeline
1. **Schema Validation**: JSON structure compliance
2. **Generator Testing**: Template output validation
3. **Mock Testing**: Reference implementation validation
4. **Integration Testing**: End-to-end workflow testing

### Quality Gates
- **Pre-commit**: Format, lint, basic test validation
- **CI Pipeline**: Full test suite validation
- **Release**: Comprehensive validation, documentation updates

This architecture supports robust, scalable testing of CCL implementations while maintaining simplicity and extensibility for future enhancements.