# CCL Test Suite Documentation

This directory documents the **CCL test-suite infrastructure**: how tests are organized, how the runner filters them, how to extend the mock, and how to contribute fixtures.

> [!IMPORTANT]
> **CCL language documentation lives on the website**, not here.
>
> - Functions, features, behaviors, decisions: <https://ccl.tylerbutler.com>
> - Test-tag to URL index (machine-readable): <https://ccl.tylerbutler.com/tag-index.json>
>
> If you're implementing CCL, start there. This directory is for test-infrastructure contributors.

## Scope

| Doc | Scope |
| --- | --- |
| **[ARCHITECTURE.md](ARCHITECTURE.md)** | Dual-format test design and component interactions |
| **[API.md](API.md)** | Go package reference (loader, config, generator, types) |
| **[CLI_REFERENCE.md](CLI_REFERENCE.md)** | Command reference for `ccl-test-runner`, `test-reader`, `validate-schema`, `validate-tags` |
| **[DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md)** | Contributor workflow |
| **[TROUBLESHOOTING.md](TROUBLESHOOTING.md)** | Common issues |
| **[MOCK_IMPLEMENTATION.md](MOCK_IMPLEMENTATION.md)** | Go mock-package layout and design decisions (not language semantics) |
| **[test-architecture.md](test-architecture.md)** | Test file organization, statistics |
| **[test-runner-design-principles.md](test-runner-design-principles.md)** | Configuration-validation rules for test runners |
| **[test-runner-implementation-guide.md](test-runner-implementation-guide.md)** | Building a test runner |
| **[test-selection-guide.md](test-selection-guide.md)** | Filtering tests by capability (see also the website's [Test Suite Guide](https://ccl.tylerbutler.com/test-suite-guide) for the conceptual side) |
| **[test-filtering.md](test-filtering.md)** | Type-safe filtering internals |
| **[schema-reference.md](schema-reference.md)** | Practical source/generated format reference |
| **[generated-schema.md](generated-schema.md)** | Auto-generated per-field schema reference |
| **[consuming-results.md](consuming-results.md)** | Deriving summaries from test-results output |
| **[implementing-ccl.md](implementing-ccl.md)** | Pointer to the canonical implementer guide on the website |
| **[tag-index.json](tag-index.json)** | Cached copy of the website's tag-URL index (fallback for `just validate-tags`) |
| **[adr/](adr/)** | Historical architectural decision records; normative rules are now on the website |
| **[plans/](plans/)** | Test-infrastructure design plans (fuzz generator, etc.) |

## Quick nav

- **New contributor?** → [ARCHITECTURE.md](ARCHITECTURE.md), then [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md).
- **Building a CCL parser?** → <https://ccl.tylerbutler.com/implementing-ccl>.
- **Building a test runner?** → [test-runner-implementation-guide.md](test-runner-implementation-guide.md).
- **Extending the mock?** → [MOCK_IMPLEMENTATION.md](MOCK_IMPLEMENTATION.md).

## Contributing

When modifying this repo:

1. **Taxonomy changes** (new function/feature/behavior/variant): update `config/config.go`, add to `src/data/tags.ts` in the website repo, and refresh `docs/tag-index.json` here. `just validate-tags` enforces the contract.
2. **New tests**: follow the schema in `schemas/source-format.json`; tag features/behaviors/variants with values declared in `config/config.go`.
3. **Tool changes**: update [CLI_REFERENCE.md](CLI_REFERENCE.md).
4. **Language-level docs** (what CCL *is*, how functions behave): contribute to the website, not this directory.
