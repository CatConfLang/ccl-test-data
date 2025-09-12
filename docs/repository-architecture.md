# Repository Architecture Analysis

## Cross-Repository Dependencies

### Repository Overview
```
ccl-test-data/        # Official JSON test suite (this repository)
├── tests/*.json      # 10 test files, 452 assertions, 167 tests
├── internal/mock/    # Go reference implementation
├── cmd/ccl-test-runner/ # CLI test runner
└── justfile          # Build automation

ccl_gleam/           # Gleam implementation (Level 1-4 complete)
├── packages/ccl_core/ # Minimal parsing (zero deps)
├── packages/ccl/     # Full-featured library
├── packages/ccl_test_loader/ # JSON test utilities
└── justfile          # Build automation

tools-monorepo/packages/ccl-docs/ # Documentation site (Astro Starlight)
├── src/content/docs/ # Markdown documentation
├── astro.config.mjs  # Astro configuration
├── package.json      # Dependencies and scripts
└── netlify.toml      # Deployment config
```

### Dependency Matrix

| Repository | Depends On | Provides To | Integration Points |
|------------|------------|-------------|-------------------|
| **ccl-test-data** | None (source of truth) | ccl_gleam, ccl-docs | JSON test files, Go reference implementation |
| **ccl_gleam** | ccl-test-data (tests) | ccl-docs (examples) | Test JSON loading, implementation patterns |
| **ccl-docs** | ccl-test-data (spec), ccl_gleam (examples) | Public documentation | API docs, implementation guides |

### Implementation Artifacts Mapping

#### ccl-test-data → ccl_gleam
```
tests/api-essential-parsing.json → packages/ccl/test/level1_tests.gleam
tests/api-object-construction.json → packages/ccl/test/level3_tests.gleam
tests/api-typed-access.json → packages/ccl/test/level4_tests.gleam
internal/mock/ccl.go → Reference patterns for Gleam implementation
```

#### ccl-test-data → ccl-docs
```
tests/*.json metadata → src/content/docs/api-reference.md
docs/levels.md → src/content/docs/implementation-levels.md
internal/mock/ patterns → src/content/docs/implementing-ccl.md examples
```

#### ccl_gleam → ccl-docs
```
packages/ccl/src/ API → src/content/docs/api-reference.md Gleam examples
packages/ccl_core/ architecture → src/content/docs/core-concepts.md
Level 1-4 implementation → src/content/docs/parsing-algorithm.md patterns
```

## Architecture Patterns

### Build Systems
- **ccl-test-data**: `just` + Go 1.23.0 + JSON schema validation
- **ccl_gleam**: `just` + Gleam + BEAM VM target + multi-package workspace  
- **ccl-docs**: pnpm + Astro + Starlight + Netlify deployment

### Test Integration
- **ccl-test-data**: Authoritative JSON test suite with 452 assertions
- **ccl_gleam**: Consumes JSON tests via `ccl_test_loader` package
- **ccl-docs**: References test files for API documentation examples

### Version Coordination
- **Test Data**: JSON schema versioning (currently v2.0)
- **Gleam**: Semantic versioning for packages (ccl_core, ccl, ccl_test_loader)
- **Documentation**: Astro site versioning with git-based deployment

## Current State Analysis

### Assets ✅
- **ccl-test-data**: 452 test assertions with feature-based tagging system
- **ccl_gleam**: Complete Level 1-4 implementation (91% pass rate equivalent)
- **ccl-docs**: Functional Astro Starlight site with comprehensive CCL documentation
- **Build Systems**: Working `just` commands across all repositories
- **Cross-Language Tests**: JSON format enables consistent validation

### Integration Points ✅
- **Test Consumption**: ccl_gleam successfully loads and executes ccl-test-data JSON tests
- **Documentation References**: ccl-docs links to test files and implementation examples
- **API Consistency**: Common function naming across Go mock and Gleam implementations

### Gaps for LLM Optimization ❌
- **No LLMs.txt files**: Missing AI-consumable repository summaries
- **Limited cross-repo linking**: Documentation doesn't deeply link to test files
- **No metadata in test JSON**: Missing LLM-friendly descriptions and guidance
- **Documentation lacks structured frontmatter**: No machine-parseable metadata
- **No unified navigation**: Repositories operate independently for AI tools

## Implementation Architecture

### CCL 4-Level Architecture (Implemented Consistently)
```
Level 1: Entry Parsing
├── Go: Parse() → []Entry
├── Gleam: parse(text) → Result(List(Entry), ParseError)  
└── Tests: api-essential-parsing.json (45 tests)

Level 2: Entry Processing  
├── Go: Filter(), Compose(), ExpandDotted() → []Entry
├── Gleam: Comment filtering, composition functions
└── Tests: api-processing.json (67 tests)

Level 3: Object Construction
├── Go: MakeObjects(entries) → map[string]any
├── Gleam: build_hierarchy(entries) → CCL
└── Tests: api-object-construction.json (89 tests)

Level 4: Typed Access
├── Go: GetString(), GetInt(), GetBool(), GetFloat() → (T, error)
├── Gleam: get_string(), get_int(), get_bool() → Result(T, Error)
└── Tests: api-typed-access.json (151 tests)
```

### Feature Tag Consistency
Both implementations use consistent structured tags:
- `function:parse`, `function:make-objects`, `function:get-string` etc.
- `feature:comments`, `feature:dotted-keys`, `feature:unicode`
- `behavior:*` and `variant:*` for implementation choices

## Resource Allocation for LLM Optimization

### Technical Requirements Met ✅
- ✅ Access to all three repositories
- ✅ Node.js/pnpm for Astro development (ccl-docs)
- ✅ Go development environment (ccl-test-data)
- ✅ Gleam development environment (ccl_gleam)
- ✅ Build automation via `just` across all repos

### Ready for Enhancement
- **starlight-llms-txt**: Added to ccl-docs package.json for LLM file generation
- **JSON Schema**: Existing v2.0 schema ready for metadata enhancement
- **Documentation Site**: Astro Starlight ready for frontmatter optimization
- **Cross-Repository Patterns**: Consistent architecture enables systematic linking

## Next Steps for Phase 1 Completion

1. **Configure starlight-llms-txt** in ccl-docs astro.config.mjs
2. **Design enhanced metadata schema** for JSON test files  
3. **Create repository linking strategy** for AI navigation
4. **Validate foundation setup** across all three repositories