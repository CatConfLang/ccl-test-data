# ADR-002: CRLF Handling in Nested Structures

## Status

Accepted

## Date

2026-03-24

## Context

Two tests with the `crlf_preserve_literal` behavior fail when the input contains nested structures (issue [#95](https://github.com/catconflang/ccl-test-data/issues/95)):

- **`crlf_preserve_nested_structure`** — The `\r` before `\n` breaks indentation detection during hierarchy construction, producing incorrect nesting.
- **`crlf_preserve_comments_and_values`** — Comment ordering and values differ from expected output when `\r\n` line endings are present.

Flat CRLF handling works correctly in all implementations tested. The failures occur specifically when CRLF line endings interact with indentation-based nesting logic (i.e., `build_hierarchy` and related functions).

The root cause in affected implementations is that line-splitting logic does not account for `\r\n` as a single line terminator. When lines are split on `\n` alone, a trailing `\r` remains attached to each line. This residual `\r` corrupts whitespace measurement, causing indentation detection to miscalculate nesting depth.

## Decision

Leave the failing tests as-is. The `crlf_preserve_literal` behavior applies uniformly to both flat and nested structures. Implementations that declare this behavior **must** handle `\r\n` line endings correctly during:

1. **Line splitting** — Treat `\r\n` as a single line terminator so that `\r` does not leak into line content.
2. **Indentation detection** — Measure leading whitespace only after stripping line terminators, ensuring `\r` does not inflate whitespace counts.
3. **Hierarchy construction** — Produce the same nested structure regardless of whether input uses `\n` or `\r\n` line endings.

Implementations failing these tests have a line-splitting bug, not a legitimate behavior difference. No new behavior tag or variant is warranted.

## Consequences

- **Test suite unchanged** — The two nested CRLF tests remain with their current expected output and `crlf_preserve_literal` behavior tag.
- **Implementation guidance** — Parser authors should ensure their line-splitting step handles `\r\n` before any indentation analysis. A common fix is to strip `\r` from line endings early in the parsing pipeline.
- **No new behaviors or variants** — CRLF preservation is a single, indivisible behavior. There is no "flat-only" CRLF mode.
- **Backward compatible** — Implementations that already pass these tests are unaffected.
