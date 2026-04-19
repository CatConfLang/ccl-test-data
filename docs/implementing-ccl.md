# Implementing CCL

The implementer's guide has moved.

**Canonical location:** <https://ccl.tylerbutler.com/implementing-ccl>

Related CCL language references on the website:

- [Parsing Algorithm](https://ccl.tylerbutler.com/parsing-algorithm)
- [Continuation Lines](https://ccl.tylerbutler.com/continuation-lines) — `parse` vs `parse_indented`, baseline N rules
- [Behavior Reference](https://ccl.tylerbutler.com/behavior-reference) — implementation choices (indentation output, continuation tab handling, CRLF, booleans, etc.)
- [Library Features](https://ccl.tylerbutler.com/library-features) — typed access, processing, formatting
- [Syntax Reference](https://ccl.tylerbutler.com/syntax-reference)

This repository (`ccl-test-data`) is the test-infrastructure source of truth — test schema, runner design, contributor workflow — not the language spec. For test-runner and schema details, see `test-architecture.md`, `test-runner-implementation-guide.md`, and `schema-reference.md` in this directory.
