# CCL Mock Implementation (Go)

The mock at `internal/mock/ccl.go` is a **Go reference implementation** used to generate expected values for the structured test suite. It is not a full production-quality CCL implementation.

**For canonical CCL language semantics — function signatures, error conditions, edge cases — see the CCL documentation site:**

- [Functions Reference](https://catconflang.com/reference/functions/)
- [Features Reference](https://catconflang.com/reference/features/)
- [Behavior Reference](https://catconflang.com/behavior-reference/)
- [Parsing Algorithm](https://catconflang.com/parsing-algorithm)
- [Continuation Lines](https://catconflang.com/continuation-lines)
- [Decisions](https://catconflang.com/reference/decisions/bare-list-hierarchy/) (bare-list, CRLF)

This document covers the **Go mock's package layout and design decisions** — what's useful for someone reading or extending the mock itself. It does not redefine CCL semantics.

## Package Layout

```
internal/mock/
  ccl.go        # CCL struct with Parse, BuildHierarchy, typed accessors, etc.
```

Exposed functions on the `CCL` receiver (see `ccl.go`):

| Function | Purpose |
| --- | --- |
| `Parse` | Flat key-value parsing from CCL text |
| `ParseIndented` | Nested-value parsing with dynamic baseline |
| `BuildHierarchy` | Flat entries → nested `map[string]any` |
| `GetString` / `GetInt` / `GetBool` / `GetFloat` / `GetList` | Typed accessors over a hierarchy |
| `Filter` | Entry filtering (pass-through in mock) |
| `Compose` | Entry-list concatenation |
| `ExpandDotted` | Dotted-key expansion (mock delegates to BuildHierarchy) |
| `PrettyPrint` / `Print` | Format hierarchy back to text |
| `RoundTrip` | Parse + PrettyPrint composition |
| `ComposeAssociative` / `IdentityLeft` / `IdentityRight` | Algebraic-property validations used by `property_algebraic.json` |

## Design Decisions Specific to the Go Mock

These are choices **this mock** makes — they are implementation-level decisions, not CCL spec requirements.

### Representation

- `Entry` is `struct { Key, Value string }` with JSON tags for direct marshaling to/from the test JSON format.
- Hierarchy is `map[string]any` (`map[string]interface{}`). Values are either strings, nested maps, or `[]any` for duplicate-key accumulation.
- This uniform `any` representation is convenient for equality-checking against expected JSON test values; a production implementation would typically pick a tagged union or generics.

### Error handling

- Detailed error messages with full path context and available-sibling hints. Example: `key not found: a.b.c (available keys: [a.b.d, a.b.e])`.
- Type-conversion errors include both the Go type seen and the path where the mismatch occurred.

### CRLF preservation

- Line splitting is `strings.Split(input, "\n")`. If the original line ended in `\r`, the mock re-attaches `\r` to the value after trimming. This is the direct approach; the canonical rule (see the [CRLF nested-structure decision](https://catconflang.com/reference/decisions/crlf-nested/)) says implementations should treat `\r\n` as a single line terminator so indentation measurement isn't corrupted.

### Deliberate simplifications

- `Filter` is pass-through in the mock. The canonical semantics are "strip comment entries"; the mock leaves comments in because test fixtures validate comment round-tripping.
- `ExpandDotted` is pass-through. `BuildHierarchy` in the mock recognizes dotted keys directly and expands them there, which is simpler than maintaining a separate entry-transform pass.

## Extending the Mock

If you add a new CCL function to the canonical taxonomy (`config/CCLFunction`):

1. Add a method on `CCL` in `internal/mock/ccl.go` producing the canonical expected value.
2. Update the generator in `internal/generator/` so it emits test cases exercising the new function.
3. Add a function section on the [website](https://catconflang.com/reference/functions/) and an entry to `src/data/tags.ts` in the website repo so `just validate-tags` resolves the new tag.

## Testing the Mock

```bash
just reset              # clean, regenerate with filters, run
just test               # run the basic test suite
just test-all           # including known-failing tests
just validate-tags      # ensure every test tag resolves to a website URL
```
