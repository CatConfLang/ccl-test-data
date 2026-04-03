# Fuzz Test Generator Design

**Issue:** [#66 — Add randomized test case generator for special character fuzzing](https://github.com/CatConfLang/ccl-test-data/issues/66)
**Date:** 2026-02-17

## Summary

A seeded randomized test generator that produces source-format JSON test cases covering special character combinations in keys and values. Output is committed to the repo for reproducibility.

## Architecture

New `fuzz` package at the library level with a thin CLI wrapper as a `generate-fuzz` subcommand in `ccl-test-runner`. Follows the existing pattern: `generator/generator.go` (library) + `cmd/ccl-test-runner/generate_flat.go` (CLI wrapper).

### New Files

- `fuzz/generator.go` — Core generation logic (seeded RNG, test case builders, source-format output)
- `cmd/ccl-test-runner/generate_fuzz.go` — Thin CLI wrapper

### Modified Files

- `cmd/ccl-test-runner/main.go` — Register `generate-fuzz` subcommand
- `justfile` — Add `generate-fuzz` recipe, update `generate-flat` to process `source_tests/fuzz/`

### Output

Source-format JSON files in `source_tests/fuzz/api_fuzz_special_chars.json`. These flow through the existing pipeline: `generate-flat` → `generate-go` → `go test`.

## Test Categories (~50 tests)

1. **Single special character keys** (~15) — One char from each category as a standalone key
2. **Multi-character combination keys** (~10) — Random combos of 2-4 special chars
3. **Positional special chars** (~10) — Special chars at start/middle/end of alphanumeric keys
4. **Values with special characters** (~5) — Special chars in values
5. **Nested structures with special char keys** (~10) — Hierarchical CCL with special char keys

### Special Characters

Path separators: `/`, `\`
Punctuation: `:`, `;`, `@`, `#`, `$`, `%`, `&`, `*`, `+`, `-`, `_`
Brackets: `(`, `)`, `[`, `]`, `{`, `}`, `<`, `>`
Quotes: `'`, `"`
Other: `~`, `` ` ``, `|`, `?`, `!`

## Validation Functions Per Test

- **`parse`** — Always present
- **`build_hierarchy`** — For nested structure tests
- **`get_string`** — For select tests with special char key paths

## Generator Details

- Seeded RNG (`int64`, default `42`) for reproducibility
- Same seed produces identical output
- After generating, validates each test through the mock CCL implementation

## CLI Interface

```
ccl-test-runner generate-fuzz [flags]
  --seed INT       Random seed (default: 42)
  --count INT      Number of tests to generate (default: 50)
  --output DIR     Output directory (default: source_tests/fuzz)
  --validate       Validate generated tests against mock implementation
```

## Justfile Integration

```just
generate-fuzz *ARGS="":
    go run ./cmd/ccl-test-runner generate-fuzz {{ARGS}}

generate-flat *ARGS="":
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/core --validate {{ARGS}}
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/fuzz --validate {{ARGS}}
```

## Decisions

- **Committed output**: Generated tests are committed for reproducibility across implementations
- **Output location**: `source_tests/fuzz/` — separate from hand-written tests
- **Scope**: Keys + values + nesting (not dotted keys or comment patterns)
- **Volume**: ~50 tests (adds ~22% to current suite)
- **No schema changes**: Uses existing valid values only
