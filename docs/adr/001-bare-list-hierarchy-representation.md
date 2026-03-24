# ADR-001: Bare List Hierarchy Representation

## Status

Accepted

## Date

2025-07-18

## Context

CCL supports "bare lists" — entries with empty keys (e.g., `= item1`, `= item2`). When
`build_hierarchy` converts flat parsed entries into a nested object structure, there are
two reasonable ways to represent these lists:

1. **Array of objects** — bare list entries become an array:
   ```json
   {"items": [{"name": "first"}, {"name": "second"}]}
   ```

2. **Nested map with empty key** — bare list entries remain in a map structure:
   ```json
   {"items": {"": {"name": {"first": {}, "second": {}}}}}
   ```

The OCaml reference implementation uses the map approach. This test suite uses the array
approach. Both are valid interpretations — `build_hierarchy` output is
implementation-defined.

Separately, the `get_list` function always returns an array regardless of internal
representation. Its purpose is to provide a consistent list interface.

See: [Issue #93](https://github.com/catconflang/ccl-test-data/issues/93)

## Decision

The test suite uses the **array-of-objects representation** as its canonical form for bare
lists in `build_hierarchy` output. Tests that validate `build_hierarchy` expect arrays, not
nested maps with empty keys.

Implementations that use a different internal representation for bare lists:

- **Should not** declare `build_hierarchy` as a supported function if their output differs
  from the array form.
- **Can still** declare `get_list`, which has well-defined array semantics independent of
  hierarchy representation.
- **Can still** use the test suite for all other functions (`parse`, `get_string`, etc.).

The existing tests are left as-is. No changes to the test suite are needed.

## Consequences

- **Test suite clarity:** The canonical representation is explicitly documented, reducing
  confusion when implementations produce different `build_hierarchy` output.
- **Implementation flexibility:** Implementations are not forced to use arrays internally.
  They simply omit `build_hierarchy` from their supported functions if their representation
  differs.
- **OCaml compatibility:** The OCaml reference implementation can use this test suite for
  everything except `build_hierarchy` tests involving bare lists.
- **Future work:** If a map-based `build_hierarchy` variant is needed, it could be added as
  a separate function or behind a behavior flag.
