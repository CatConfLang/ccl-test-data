# CLI Generator Approach for Test Filtering

**Goal**: Simplify implementer experience by providing pre-filtered tests that match their implementation capabilities, eliminating the need to understand complex mutual exclusion logic.

## Core Principle

**Implementation simplicity is most important.** Rather than enhance the schema with complex behavioral metadata, provide a CLI tool that accepts simple configuration and outputs only compatible tests.

## Approach Overview

1. **Simple Config Format** - Implementers specify what they support
2. **CLI Generator** - Filters tests based on config
3. **Standard Output** - Generated tests use existing JSON format
4. **Coverage Reporting** - Address visibility loss with optional reporting

## Simple Configuration Format

```yaml
# ccl-impl.yml - Dead simple configuration
functions:
  - parse
  - make_objects
  - get_string
  - get_int
  - get_bool

behaviors:
  - crlf_normalize_to_lf
  - boolean_lenient
  - loose_spacing

features:
  - comments
  - unicode
```

**Design principles:**
- No complex groupings or exclusion logic in config
- Just list what you support
- Self-documenting through simplicity

## CLI Interface

### Basic Commands
```bash
# Basic usage
ccl-test-gen --config my-impl.yml

# Inline config (no file needed)
ccl-test-gen --functions parse,get_string --behaviors crlf_normalize_to_lf

# With coverage info to address visibility loss
ccl-test-gen --config my-impl.yml --show-excluded

# Integration with existing justfile
just generate-filtered my-impl.yml
```

### Extended Justfile Commands
```makefile
# Add to existing justfile
generate-filtered CONFIG:
    ./bin/ccl-test-runner generate-filtered --config {{CONFIG}} --output filtered_tests.json

coverage-report CONFIG:
    ./bin/ccl-test-runner generate-filtered --config {{CONFIG}} --coverage-report --no-output

# Predefined common configs
generate-minimal:
    just generate-filtered configs/minimal.yml

generate-basic:
    just generate-filtered configs/basic.yml
```

## Filtering Logic

**Simple inclusion rules - include test if:**

1. All required `functions` are in your config
2. All required `features` are in your config
3. ALL `behaviors` in test are in your config (or test has no behavior requirements)
4. No `conflicts` with your chosen behaviors

```go
func includeTest(test GeneratedTest, config Config) bool {
    // Check functions
    for _, fn := range test.Functions {
        if !contains(config.Functions, fn) {
            return false
        }
    }

    // Check features
    for _, feature := range test.Features {
        if !contains(config.Features, feature) {
            return false
        }
    }

    // Check behaviors - test's behaviors must be subset of config
    for _, behavior := range test.Behaviors {
        if !contains(config.Behaviors, behavior) {
            return false
        }
    }

    // Check conflicts
    for _, conflict := range test.Conflicts.Behaviors {
        if contains(config.Behaviors, conflict) {
            return false
        }
    }

    return true
}
```

## Coverage Reporting

Address the visibility loss drawback with clear reporting:

```
Implementation Coverage Report
=============================
✅ Included Tests: 342/452 (75.6%)
❌ Excluded Tests: 110/452 (24.4%)

Missing Functions (would unlock 45 tests):
  • function:filter (23 tests)
  • function:pretty_print (22 tests)

Missing Features (would unlock 38 tests):
  • feature:dotted_keys (25 tests)
  • feature:multiline (13 tests)

Behavioral Conflicts (27 tests excluded):
  • behavior:crlf_preserve_literal conflicts with your choice of crlf_normalize_to_lf (15 tests)
  • behavior:boolean_strict conflicts with your choice of boolean_lenient (12 tests)

Quick Wins:
  • Add feature:dotted_keys → +25 tests (80.1% coverage)
  • Add function:filter → +23 tests (85.6% coverage)
```

## File Structure

Minimal changes to existing structure:

```
ccl-test-data/
├── configs/
│   ├── minimal.yml      # parse only
│   ├── basic.yml        # current mock level
│   └── reference.yml    # full reference implementation
├── generated_tests/     # existing full format (unchanged)
├── plans/               # this document
└── bin/ccl-test-runner  # extend with generate-filtered command
```

## Benefits

✅ **Implementer Simplicity**: Only see relevant tests, standard JSON format
✅ **No Schema Changes**: Keeps existing test format unchanged
✅ **Performance**: Pre-filtered tests, smaller files, faster execution
✅ **Clean Separation**: Configuration logic isolated in CLI tool
✅ **Coverage Visibility**: Optional reporting addresses information loss
✅ **Easy Integration**: Extends existing Go toolchain naturally

## Implementation Plan

### Phase 1: Core CLI Generator
1. Extend `ccl-test-runner` with `generate-filtered` subcommand
2. Implement simple YAML config parsing
3. Add filtering logic based on functions/features/behaviors/conflicts
4. Generate standard JSON output format

### Phase 2: Coverage Reporting
1. Add `--coverage-report` flag
2. Implement exclusion analysis and reporting
3. Show "quick wins" for adding features/functions

### Phase 3: Predefined Configs
1. Create common configuration files (minimal, basic, reference)
2. Add justfile integration commands
3. Documentation and examples

### Phase 4: Advanced Features (Optional)
1. `--what-if` flag for exploring feature additions
2. `--dry-run` mode for testing configurations
3. Config validation and helpful error messages

## Decision Rationale

**Why CLI Generator over Schema Enhancement:**
- Implementation simplicity is the top priority
- Implementers don't need to understand complex exclusion logic
- Maintains familiar JSON test format
- Lower risk approach (no schema changes)
- Can be extended later if needed

**Why not Schema Enhancement:**
- Adds complexity for implementers
- Requires learning new schema concepts
- More maintenance overhead
- Higher adoption barrier

This approach prioritizes getting implementers productive quickly with minimal cognitive overhead while still providing visibility into what they're missing through optional coverage reporting.