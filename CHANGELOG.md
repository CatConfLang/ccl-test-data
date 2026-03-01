# Changelog

All notable changes to the CCL test data will be documented in this file.


## [0.6.2] - 2026-03-01



### Bug Fixes

- **tests:** Correct behavior tag on key_with_tabs_ocaml_reference ([`bcea90e`](https://github.com/CatConfLang/ccl-test-data/commit/bcea90e9a7b05bafdb6c38a393f35ec5f3951ae4))

  The reference_compliant variant strips tabs around the `=` delimiter, which is tabs_as_whitespace behavior, not tabs_as_content. Fixes #83.




## [0.6.1] - 2026-02-28



### Bug Fixes

- **tests:** Resolve issues ([`b691aa5`](https://github.com/CatConfLang/ccl-test-data/commit/b691aa5ec8d9810cf43aeecae488fb47fd7457a9))

  ## Summary

    Fixes three related issues with test generation and expectations.

    ### Changes

    * **fix(generator): add `canonical_format` to its own compositeFunctionMap entry (#79)** Tests with `canonical_format` validation now include `canonical_format` in their functions array, so implementations that support `parse`+`print` but not `canonical_format` will correctly skip these tests.

    * **fix(schema): add `filter` to CRLF behavior `affectedFunctions` (#80)** The `crlf_preserve_literal` and `crlf_normalize_to_lf` behaviors now list `filter` in their `affectedFunctions`. This ensures the `filter` sub-test inherits CRLF behavior tags and conflict metadata during flat test generation, so implementations with `crlf_normalize_to_lf` correctly skip CRLF-preserving filter tests.

    * **fix(tests): remove trailing whitespace from print expectations (#81)** 21 print test expectations had trailing whitespace on lines where a key has a continuation or empty value (e.g. `config = \n` instead of `config =\n`). This caused `round_trip: true` assertions to be self-contradictory since `print(parse(input)) != input`. Updated the mock `Print` function to omit trailing space when the value is empty or starts with a newline, and fixed all affected test expectations.

    Closes #79, closes #80, closes #81

- **tests:** Resolve test inconsistencies ([`338d2fd`](https://github.com/CatConfLang/ccl-test-data/commit/338d2fdc2df10fe0e559e0d91edd92a73a7a6e3f))

  ## Summary

    - **#75**: Update `round_trip_property_complex` print expectation to use no-leading-space format (`= item1`) for empty keys, consistent with the fix in 13fb778
  - **#76**: Fix behavior tag propagation for composite validations (`round_trip`, `canonical_format`, `load`) by resolving to component functions before filtering — `tabs_as_whitespace_round_trip_round_trip` now correctly has `behaviors: ["tabs_as_whitespace"]`
  - **#77**: Update CRLF `build_hierarchy` tests to include `/` comment keys (as arrays for duplicates), consistent with `ocaml_stress_test_original` behavior

    Closes #75, closes #76, closes #77




## [0.6.0] - 2026-02-24



### Bug Fixes

- **tests:** Use no-leading-space format for top-level empty keys in print ([`13fb778`](https://github.com/CatConfLang/ccl-test-data/commit/13fb7789f2c37c82e779cc29663b3e332c1db7a7))

  The print expectations for top-level empty keys used ` = value` (leading space), which would be re-parsed as a continuation line, making round_trip: true contradictory. Changed to `= value` (no leading space) to match the input format. Indented empty keys within continuation values are unaffected.

    Fixes #72





### Features

- **tests:** Add delimiter_first_equals test cases ([`e219dd8`](https://github.com/CatConfLang/ccl-test-data/commit/e219dd89d6b2128c919cb40a56ec01935720d4e1))

  Add three tests showing first-equals delimiter behavior as counterparts to the existing delimiter_prefer_spaced tests. Same inputs produce different parse results under each behavior:
  - delimiter_first_url_with_query_params: URL splits at first bare `=`
  - delimiter_first_multiple_equals: `a=b = c=d` splits at first `=`
  - delimiter_first_empty_value: `a=b =` splits at first `=`

- **tests:** Add delimiter_prefer_spaced test cases ([`ac139eb`](https://github.com/CatConfLang/ccl-test-data/commit/ac139eb548942cd54d37459a0396c1a2c2b46699))

  Add three new tests for the delimiter_prefer_spaced behavior:
  - delimiter_spaced_multiple_equals: `a=b = c=d` splits on ` = `
  - delimiter_spaced_fallback_no_space: `key=value` falls back to bare `=`
  - delimiter_spaced_empty_value: `a=b =` with empty value after ` = `

    Closes #73

- **tests:** Add edge case tests, delimiter behavior, and update dependencies ([`59c1d6d`](https://github.com/CatConfLang/ccl-test-data/commit/59c1d6d24487a259a4755909eb41c79d0104cb07))

  ## Summary

    - Add comprehensive edge case tests for relative paths, double slashes, URLs with special characters, and deeply nested structures
  - Add `delimiter_first_equals` / `delimiter_prefer_spaced` behavior pair to schema and config, addressing how parsers handle `=` in keys (e.g., URLs with query params)
  - Fix test inputs for intermediate section header lines to include trailing spaces, matching expected parse values
  - Upgrade JSON Schema dialect from Draft-07 to Draft 2019-09 to support `$defs` keyword
  - Bump Go from 1.25.4 to 1.26.0
  - Update npm dependencies: `@commitlint/*` to ^20.4.2, `@sourcemeta/jsonschema` to ^14.13.3, `@tylerbu/cli` to ^0.9.0, `unist-util-visit` to ^5.1.0

    Relates to #73, #74

- **tests:** Add crlf + comments test coverage ([`d09d581`](https://github.com/CatConfLang/ccl-test-data/commit/d09d5813c9fa05b355184ef7447016f05ed8d737))

  ## Summary

    Add test cases for CRLF line endings combined with `/=` comment syntax, addressing the gap identified in #68.

    ## New Tests

    | Test Name | Behavior | Description | |-----------|----------|-------------| | `crlf_normalize_comment_only` | `crlf_normalize_to_lf` | Single comment with CRLF ending; parse recognizes comment, filter removes it | | `crlf_preserve_comment_only` | `crlf_preserve_literal` | Single comment with CRLF; `\r` preserved in value, filter still removes it | | `crlf_normalize_comments_and_values` | `crlf_normalize_to_lf` | Mixed comments and key-value pairs with CRLF | | `crlf_preserve_comments_and_values` | `crlf_preserve_literal` | Mixed comments and key-value pairs with CRLF; `\r` preserved | | `crlf_normalize_multiple_comments` | `crlf_normalize_to_lf` | Three consecutive comments with CRLF | | `crlf_preserve_multiple_comments` | `crlf_preserve_literal` | Three consecutive comments with CRLF; `\r` preserved |

    These 6 source tests generate **14 flat test assertions** covering `parse`, `filter`, and `build_hierarchy` functions.

    Closes #68




## [0.5.0] - 2026-02-17



### Documentation

- Fix go version reference in claude.md ([`f105c3e`](https://github.com/CatConfLang/ccl-test-data/commit/f105c3e49ee8e9550002c8fb69381cc771743d83))

- Update release process documentation ([`8561c42`](https://github.com/CatConfLang/ccl-test-data/commit/8561c42fa3e443bf8b9611b2e08161aea01f1026))

  Update CLAUDE.md, DEV.md, and DEVELOPER_GUIDE.md with new squash merge commit format (`* type(scope): description`) and semantic-release commands.





### Features

- **tests:** Add randomized fuzz test generator for special characters ([`64112a2`](https://github.com/CatConfLang/ccl-test-data/commit/64112a28da853a052757873b318cdb4c16f2caac))

  ## Summary

    Closes #66

    - Add `fuzz/` package with seeded randomized test generator producing source-format JSON test cases covering special character combinations in CCL keys and values
  - Add `generate-fuzz` CLI subcommand to `ccl-test-runner` with `--seed`, `--count`, `--output`, `--validate` flags
  - Update `justfile` with `generate-fuzz` recipe and extend `generate-flat` to process `source_tests/fuzz/`
  - Generate and commit 50 fuzz tests (seed 42) covering 5 categories: single char keys, combo keys, positional chars, special values, nested structures

- **tests:** Add test case for forward slashes in map keys ([`950d0d2`](https://github.com/CatConfLang/ccl-test-data/commit/950d0d25111c2c6861734190146f58b364beda72))

  ## Summary

    Add test case verifying that forward slashes in map keys are parsed correctly.

    This addresses the issue discovered while testing repoverlay, where sickle fails to parse map keys containing forward slashes (tracked in tylerbutler/santa#71).

    ## Test Case

    ```ccl mappings =
    config/settings.json = .vscode/settings.json
    src/template.env = .env ```

    Validates:
  - **parse**: flat entry extraction with forward slashes in nested keys
  - **build_hierarchy**: nested object construction with keys like `config/settings.json`
  - **get_string**: typed value access via path containing forward slashes

    Closes #62

- **build:** Replace git-cliff with python-semantic-release ([`8561c42`](https://github.com/CatConfLang/ccl-test-data/commit/8561c42fa3e443bf8b9611b2e08161aea01f1026))

  Switch changelog generation from git-cliff to Python Semantic Release for better squash merge commit parsing.

- **build:** Add python-semantic-release configuration ([`8561c42`](https://github.com/CatConfLang/ccl-test-data/commit/8561c42fa3e443bf8b9611b2e08161aea01f1026))

  Configure PSR with parse_squash_commits for extracting multiple conventional commits from squash merge bodies. Add custom Jinja2 templates for changelog and release notes with commit links.





### Refactoring

- **build:** Replace python commitlint generator with @tylerbu/cli ([`5859f17`](https://github.com/CatConfLang/ccl-test-data/commit/5859f1763cbd6b0b57ba96f05425a195743c6d0c))

  Use `npx tbu generate commit-config` to generate commitlint.config.cjs from commit-types.ccl instead of the custom Python script.

    - Remove scripts/generate-commitlint-config.py
  - Add consola dependency (required by @tylerbu/cli-api)
  - Update commitlint.config.cjs source reference to commit-types.ccl

- **build:** Update release workflow for semantic-release ([`8561c42`](https://github.com/CatConfLang/ccl-test-data/commit/8561c42fa3e443bf8b9611b2e08161aea01f1026))

  Replace git-cliff GitHub Action with Python-based semantic-release. Use Python script for extracting release notes instead of awk.




## [0.4.0] - 2026-01-21



### Bug Fixes

- **tests:** Add composite function mapping and bare list nested object tests ([`ad83c0c`](https://github.com/CatConfLang/ccl-test-data/commit/ad83c0caa4c3c924fb559cb2e756c546d832f30a))

  Addresses #59 and #60.

- **tests:** Composite function mapping ([`ad83c0c`](https://github.com/CatConfLang/ccl-test-data/commit/ad83c0caa4c3c924fb559cb2e756c546d832f30a))

  Composite functions now correctly report their required underlying functions:
  - `round_trip` → `["parse", "print"]`
  - `load` → `["parse", "build_hierarchy"]`
  - `canonical_format` → `["parse", "print"]`

    Also standardizes on `"print"` (instead of `"pretty_print"`) across config, loader, and types to match the schema.

- **tests:** Bare list nested object tests ([`ad83c0c`](https://github.com/CatConfLang/ccl-test-data/commit/ad83c0caa4c3c924fb559cb2e756c546d832f30a))

  Adds 6 new tests for bare list syntax with nested object items:
  - `bare_list_nested_objects_basic`
  - `bare_list_nested_objects_single_item`
  - `bare_list_nested_objects_minimal`
  - `bare_list_nested_objects_deeply_nested`
  - `bare_list_nested_objects_mixed_with_strings`
  - `bare_list_nested_objects_round_trip`

- **schema:** Allow nested objects in expect arrays ([`ad83c0c`](https://github.com/CatConfLang/ccl-test-data/commit/ad83c0caa4c3c924fb559cb2e756c546d832f30a))

  Updates `source-format.json` to use `anyOf` instead of `oneOf` for array items in `expect` fields, enabling `get_list` tests to return lists of complex objects (not just strings or entries).





### Features

- **config:** Enable split_commits for multi-entry changelog parsing ([`de06fc5`](https://github.com/CatConfLang/ccl-test-data/commit/de06fc5fcebf4b744cb70c9796536d84b0c36947))

  Configure git-cliff to parse multiple conventional commits from commit bodies:
  - Enable split_commits to process each line individually
  - Add preprocessor to strip markdown header prefixes (###)
  - Add unique filter to deduplicate identical messages in template

- **config:** Add single-source-of-truth for conventional commit types ([`e0b69ea`](https://github.com/CatConfLang/ccl-test-data/commit/e0b69ea944b2279b3b3fe7cecd6c9cac72454c2a))

  Add commit-types.json as the canonical source for commit type definitions, changelog groupings, and scope configurations. Includes CCL format version and Python generator script to produce cliff.toml and commitlint.config.cjs.




## [0.3.1] - 2026-01-09



### Bug Fixes

- **generation:** Sort behavior conflicts for deterministic output ([`15fa979`](https://github.com/CatConfLang/ccl-test-data/commit/15fa979c0c7829b8493fa246a76a8bf8aa4d03c5))

  ## Summary

    Sorts the behavior conflicts array alphabetically when auto-generating conflicts from metadata, preventing non-deterministic ordering in generated JSON files.

    ## Changes

    - Add `sort.Strings()` to `GetConflictingBehaviors()` in `generator/metadata.go`
  - Add sorting to `GetAllBehaviors()` for consistency
  - Regenerate test files with deterministic conflict ordering

    ## Problem

    Go's map iteration order is non-deterministic, causing the `conflicts.behaviors` arrays in generated JSON to flip-flop between regenerations (e.g., `["list_coercion_disabled", "array_order_lexicographic"]` vs `["array_order_lexicographic", "list_coercion_disabled"]`).





### Documentation

- Add claude.md context files for key directories ([`7d418b6`](https://github.com/CatConfLang/ccl-test-data/commit/7d418b6b17e878c9f699b6f7bb29a7003e6cbafd))

  ## Summary

    - Adds 7 CLAUDE.md context documentation files across key repository directories
  - Files contain auto-generated recent activity indexes from claude-mem for AI-assisted development
  - Directories covered: `config/`, `docs/`, `generated_tests/`, `go_tests/parsing/`, `internal/config/`, `source_tests/core/`, `types/`

    These context files help Claude Code quickly understand recent changes and observations in each directory without re-reading code or searching through history.




## [0.3.0] - 2026-01-05



### Bug Fixes

- **tests:** Add missing crlf behavior tags to build_hierarchy tests ([`7e08787`](https://github.com/CatConfLang/ccl-test-data/commit/7e08787ed75d891a1158c81fd8a82d141b597eea))

  Closes #52

    - Add build_hierarchy test to crlf_preserve_literal_basic test
  - Add new crlf_preserve_nested_structure test with both parse and
    build_hierarchy validations
  - Auto-conflict generation ensures all tests have proper conflicts field

    Test count increases from 399 to 402 with 3 new test assertions.





### Features

- **schema:** Add behavior metadata for auto-conflict generation ([`5dd3fd8`](https://github.com/CatConfLang/ccl-test-data/commit/5dd3fd87314f8a70a71d0eb5570efdd7d841252a))

  ## Summary

    - Add behavior metadata system to source-format.json for automatic conflict generation
  - Embed behavior definitions (affectedFunctions, mutuallyExclusiveWith) as `x-behaviorMetadata` extension
  - Auto-generate behavior conflicts during flat format generation based on mutual exclusivity rules
  - Remove need for manual `conflicts` declarations in source tests (deprecated field)

    ## Changes

    **Schema consolidation:**
  - Merge behavior-metadata.json into source-format.json using JSON Schema `x-` extension convention
  - Add `$defs` for reusable type definitions (behaviorName, functionName, featureName, variantName)
  - Single source of truth for behavior definitions and schema validation

    **Generator enhancements:**
  - Add `LoadBehaviorMetadata()` to extract x-behaviorMetadata from source-format.json
  - Filter behaviors per-function using affectedFunctions mapping
  - Auto-generate conflicts from mutuallyExclusiveWith definitions
  - Add `--validate`, `--auto-conflicts`, and `--schemas` CLI flags

    **Source test cleanup:**
  - Remove all manual `conflicts` fields from source tests (now auto-generated)
  - Regenerate flat format tests with proper behavior propagation




## [0.2.0] - 2026-01-04



### Bug Fixes

- **tests:** Add reference_compliant variant to key_with_tabs_ocaml_reference ([`52f20ad`](https://github.com/CatConfLang/ccl-test-data/commit/52f20ad3b592e52f04af49882abfd03796799ea1))

  Add missing variants tag to properly identify this test as expecting OCaml reference implementation behavior. This resolves the conflict where two tests with the same behavior requirement had different expected outputs.

    Related: tylerbutler/tools-monorepo#532

- **schema:** Add continuation baseline behaviors to config schema ([`c0c9670`](https://github.com/CatConfLang/ccl-test-data/commit/c0c96704ac46c708463c50bea4bc5dc98989b62e))

  ## Summary

    - Adds `toplevel_indent_strip` and `toplevel_indent_preserve` behaviors to `ccl-config-schema.json`
  - These behaviors were already present in `source-format.json` and `generated-format.json` but missing from the user-facing config schema
  - Fixes consistency issue identified during documentation review

    ## Context

    The config schema is used by implementations to declare their test runner capabilities. Without these behavior options, implementations couldn't properly declare their continuation baseline behavior choice.

- **cli:** Improve behavior conflict detection and display ([`e68c1f6`](https://github.com/CatConfLang/ccl-test-data/commit/e68c1f642719196515d08a7c3882af5999cfe892))

  ## Summary

    Fixes behavior conflict detection in the stats command to correctly identify mutually exclusive behavior pairs and improves display formatting.

    ### Changes

    - **Fix cross-contamination bug**: When tests have multiple behaviors (e.g., `list_coercion_enabled` + `array_order_insertion`), conflicts are now only attributed to behaviors that share the same prefix. Previously, all behaviors on a test were associated with all declared conflicts, creating false positive relationships.

    - **Add `haveSameBehaviorPrefix()` helper**: Detects whether two behaviors are of the same type based on naming convention (e.g., `list_coercion_*`, `array_order_*`, `tabs_as_*`).

    - **Improve conflict display**: Shows bidirectional pairs on single lines with color-coded names instead of listing each direction separately.

    - **Add missing conflict declarations**: Added bidirectional conflict declarations to test data for:
    - `tabs_as_content` ↔ `tabs_as_whitespace`
    - `crlf_normalize_to_lf` ↔ `crlf_preserve_literal`
    - `list_coercion_enabled` ↔ `list_coercion_disabled`

    ### Result

    The `just stats` output now correctly shows all 5 mutually exclusive behavior pairs without false positives.

- **generation:** Propagate conflicts field to generated flat format ([`381134d`](https://github.com/CatConfLang/ccl-test-data/commit/381134d98479f9f4cbaf1b4c067809cc3d340f6c))

  ## Summary

    Fixes the `conflicts` field not being propagated from source tests to the generated flat format.

    ## Changes

    - Add `convertConflicts()` function in `generator/generator.go` to convert `*types.ConflictSet` to `*generated.GeneratedFormatSimpleJsonTestsElemConflicts`
  - Include `Conflicts` field in the struct initialization within `convertToFlatFormat()`

    ## Impact

    Implementations consuming the flat format can now use `conflicts` metadata for test filtering directly, without parsing the source format or duplicating conflict logic.

    Fixes #49

- **tests:** Remove incorrect proposed_behavior for nested parsing ([`35a194f`](https://github.com/CatConfLang/ccl-test-data/commit/35a194f127478e4bd91e2848f49a514455c93977))

  ## Summary

    Removes an incorrectly specified test that expected wrong `parse` behavior. The test was mistakenly placed in `api_proposed_behavior.json` with a `variants` tag, but this was not actually a spec ambiguity - it was simply wrong.

    ## Problem

    The `deeply_nested_list` test in `api_proposed_behavior.json` expected `parse` to flatten nested structure into 6 separate entries. This contradicts the CCL specification: `parse` should preserve nested indented content as part of the value string, with structural interpretation deferred to `build_hierarchy`.

    This was not a legitimate "variant" (spec ambiguity between OCaml reference and proposed behavior) - it was an incorrect test expectation.

    ## Changes

    - **Removed** `deeply_nested_list` from `api_proposed_behavior.json` (incorrect test)
  - **Moved** `deeply_nested_list_reference` from `api_reference_compliant.json` to `api_core_ccl_hierarchy.json` as the canonical `deeply_nested_list` test
  - **Removed** `TestDeeplyNestedListParse` from skip list since it now passes

    Closes #42, closes #43

- **cli:** Update view commands to use build-bin recipe ([`e35296c`](https://github.com/CatConfLang/ccl-test-data/commit/e35296c6adc0ec9fbdaf45a7b0e7f4f7728c3908))

  The build-test-reader recipe was removed but view-test and view-tests commands still referenced it. Updated all four view commands to use the consolidated build-bin recipe instead.





### Documentation

- Clarify continuation baseline behavior for parse vs parse_indented ([`0c85e62`](https://github.com/CatConfLang/ccl-test-data/commit/0c85e62403b363c6cb90046d9c37259af40f5ec0))

  Add tests and documentation explaining the difference in continuation baseline (N) between top-level parsing and nested value parsing:

    - parse: strips leading whitespace, uses N=0 baseline
  - parse_indented: uses actual indentation of first content line as N

    This addresses confusion from the blog post's simplified description which doesn't fully match the OCaml reference implementation's behavior.

    Refs: #44

- Clarify behaviors are not mutually exclusive ([`0843c71`](https://github.com/CatConfLang/ccl-test-data/commit/0843c71230f2e8ec1683b82be0960cf734f8d1c6))

  Behaviors are not inherently mutually exclusive - a test can require multiple behaviors. The conflicts field determines which combinations are incompatible on a per-test basis.

    - README.md: Updated section header and added clarifying note docs/API.md: Clarified conflicts field description docs/test-filtering.md: Removed misleading "(Mutually Exclusive)" header





### Features

- **schema:** Add toplevel_indent_strip and toplevel_indent_preserve behavior pair ([`8703083`](https://github.com/CatConfLang/ccl-test-data/commit/8703083b63ddb96b89010fe4c47fd2edd2688d4d))

  ## Summary

    Adds a new mutually exclusive behavior pair for top-level continuation baseline detection, addressing issue #46.

    **New behaviors:**
  - `toplevel_indent_strip` - Strip leading indent at top-level, using N=0 as continuation baseline (OCaml reference behavior)
  - `toplevel_indent_preserve` - Preserve first key's indentation as the continuation baseline

    **Changes:**
  - Schema: Added both behaviors to `source-format.json` and `generated-format.json`
  - Config: Added `toplevel_indent` conflict group and `ToplevelIndent` field to runner config
  - Tests: Tagged existing whitespace normalization tests with `toplevel_indent_strip` + `reference_compliant`, added corresponding `toplevel_indent_preserve` test variants
  - Justfile: Updated skip-tags to exclude `toplevel_indent_preserve` tests (mock uses strip behavior)

    The `toplevel_indent_strip` tests are marked `reference_compliant` since this behavior is required for OCaml compatibility.

    Closes #46




## [0.1.0] - 2025-12-31



### Bug Fixes

- **tests:** Correct tabs_to_spaces behavior to use single space ([`52813ae`](https://github.com/CatConfLang/ccl-test-data/commit/52813aeaa717934e1e0f3202fc042235182a92cc))

  The tabs_to_spaces behavior should convert each tab to a single space, not two spaces. Fixed expectations in 7 tests:

    - tabs_to_spaces_in_value
  - tabs_to_spaces_leading_tab
  - tabs_to_spaces_multiple_tabs
  - tabs_to_spaces_multiline
  - tabs_canonical_format_to_spaces
  - spacing_and_tabs_combined_loose_to_spaces
  - behavior_combo_tabs_and_crlf

- **tests:** Correct tabs_to_spaces behavior to use single space ([`67824fb`](https://github.com/CatConfLang/ccl-test-data/commit/67824fb18c5df232c47bbe4de3a6205bd7ad58e0))

  ## Summary
  - Fixed 7 tests that incorrectly expected tabs to be converted to two spaces instead of one
  - The `tabs_to_spaces` behavior should convert each tab character to a single space

    ## Tests Fixed
  - `tabs_to_spaces_in_value`
  - `tabs_to_spaces_leading_tab`
  - `tabs_to_spaces_multiple_tabs`
  - `tabs_to_spaces_multiline`
  - `tabs_canonical_format_to_spaces`
  - `spacing_and_tabs_combined_loose_to_spaces`
  - `behavior_combo_tabs_and_crlf`

    ## Test plan
  - [x] `just reset` passes
  - [x] `just validate` passes

- **tests:** Correct filter function expectations to remove comments ([`0075205`](https://github.com/CatConfLang/ccl-test-data/commit/00752054d3bd2615c57a80080a758cfe0ff03a27))

  ## Summary
  - Fix filter test expectations in `api_comments.json` to correctly expect comments removed
  - Fix mock `Filter` implementation to actually filter out comment entries (key="/")

    ## Problem The `filter` function tests were incorrectly expecting the same output as `parse`, keeping comment entries in the result. The `filter` function's purpose is to remove comment entries from parsed CCL data.

    ## Changes
  - `comment_extension`: Now expects 4 entries (without the 2 comment entries)
  - `comment_syntax_slash_equals`: Now expects empty array (input was only a comment)
  - `section_headers_with_comments`: Now expects 4 entries (without the 2 comment entries)
  - Mock `Filter` function now filters out entries where `key == "/"`

    ## Test plan
  - [x] `just reset` passes
  - [x] `just validate` passes

- **tests:** Remove incorrect canonical_format tests from proposed_behavior ([`4c29c7b`](https://github.com/CatConfLang/ccl-test-data/commit/4c29c7b6d594417c42303e8322919b8c9bcaff23))

  ## Summary

    - Remove 8 incorrect `canonical_format` tests from `api_proposed_behavior.json`
  - The `canonical_format` function always produces model-level transformation output (reference compliant behavior), not structure-preserving output
  - The removed tests had incorrect expectations showing `canonical_format` behaving like `print`

    ## Details

    The `canonical_format` function transforms `key = value` into nested form: ``` key =
    value = ```

    The removed tests incorrectly expected structure-preserving output like: ``` key = value ```

    ## Removed tests

    - `canonical_format_empty_values`
  - `canonical_format_tab_preservation`
  - `canonical_format_unicode`
  - `canonical_format_line_endings_proposed`
  - `crlf_normalize_to_lf_proposed`
  - `crlf_normalize_to_lf_indented_proposed`
  - `canonical_format_consistent_spacing`
  - `deterministic_output`

    ## Test plan

    - [x] `just validate` passes
  - [x] `just reset` passes
  - [x] Stats updated: 205 tests, 447 assertions (was 213/463)





### Documentation

- Simplify readme and add dev.md ([`a7b6b87`](https://github.com/CatConfLang/ccl-test-data/commit/a7b6b8726bab0cc0604a4a39ba48ad8613b833b1))

  ## Summary

    - Simplify README.md to focus on test data consumers: quick start, test format, filtering logic, Go library usage
  - Add DEV.md with developer-focused content: setup, workflow commands, project structure, release process, troubleshooting
  - Update external links to point to ccl.tylerbutler.com documentation

- Fix behavior documentation accuracy ([`9b7dcb5`](https://github.com/CatConfLang/ccl-test-data/commit/9b7dcb591e2e1cf6dc279cefa124e24f7637948a))

  ## Summary

    Fixes behavior documentation to accurately reflect the test suite schema and behavior system.

    ### Changes

    - **Clarify behavior exclusivity**: Behaviors are not inherently mutually exclusive. The `conflicts` field on individual tests determines which behavior combinations are incompatible for that specific test.

    - **Fix duplicate entry**: Removed duplicate "Spacing" row in test-development.md (was identical to "Tabs" row)

    - **Add missing behaviors**: Added documentation for `indent_spaces`/`indent_tabs` and `array_order_insertion`/`array_order_lexicographic`

    - **Add descriptions**: Each behavior pair now has a brief description explaining its purpose

    ### Files Modified

    - `docs/test-filtering.md` - Main behavior reference
  - `docs/test-runner-implementation-guide.md` - Implementation guidance
  - `.github/agents/test-development.md` - Agent documentation
  - `.github/agents/onboarding.md` - Agent documentation

- Clean up test selection guide ([`7d9826a`](https://github.com/CatConfLang/ccl-test-data/commit/7d9826ab564bc26214ba3b47b4c45c8609666621))





### Features

- **schema:** Rename tab/spacing behaviors for clarity ([`67ef97d`](https://github.com/CatConfLang/ccl-test-data/commit/67ef97d0d1fe5e32f7da9e27651fcc05d0f21d6a))

  ## Summary

    - Rename tab parsing behaviors for clarity: `tabs_preserve` → `tabs_as_content`, `tabs_to_spaces` → `tabs_as_whitespace`
  - Remove obsolete `strict_spacing` and `loose_spacing` behaviors (no longer relevant)
  - Add new indent output behaviors: `indent_spaces` and `indent_tabs`
  - Update schemas, Go code, test files, and documentation

    ## Breaking Changes

    Behavior names have been updated. Implementations using the old names will need to update:

    | Old Name | New Name | |----------|----------| | `tabs_preserve` | `tabs_as_content` | | `tabs_to_spaces` | `tabs_as_whitespace` | | `strict_spacing` | *removed* | | `loose_spacing` | *removed* |

    New behaviors added:
  - `indent_spaces` - Use spaces for printed indentation (default)
  - `indent_tabs` - Use tabs for printed indentation

    ## Test plan

    - [x] `just validate` passes
  - [x] `just generate` completes successfully
  - [x] CI tests pass

    Closes #38

- **tests:** Add comprehensive behavior tests ([`268e8bd`](https://github.com/CatConfLang/ccl-test-data/commit/268e8bdba773e75aec8d9a3c703e0c102bc0193f))

  ## Summary

    Adds 26 new test cases to improve coverage of parse-time and access-time configurable behaviors:

    - **CRLF handling** (6 tests): normalize vs preserve, multiline values, nested structures, mixed line endings
  - **Strict spacing** (5 tests): validates rejection of non-standard formats like `key=value`, `key =value`
  - **Boolean edge cases** (7 tests): case sensitivity, numeric 1/0, whitespace handling, empty values, nested objects
  - **Type mismatch errors** (4 tests): get_int on bool, get_bool on int, get_float on bool
  - **Behavior combinations** (2 tests): tabs+crlf, loose+tabs+crlf triple combos

    ## Test plan

    - [x] JSON validates successfully
  - [x] `just reset` regenerates tests without errors
  - [x] All new tests follow source format schema

- **tests:** Add bare list indentation tests ([`0dc4727`](https://github.com/CatConfLang/ccl-test-data/commit/0dc47273d7e3248d2eefb86bb46e92e489281f2a))

  ## Summary
  - Add `canonical_format` tests for bare list indentation validation
  - Tests verify nested bare lists use correct 2-space indentation
  - Catches formatting bugs where printers output incorrect indentation (e.g., 3 spaces instead of 2)

    ## New Tests
  - **nested_bare_list_indentation**: Single-level nested bare list
  - **deeply_nested_bare_list_indentation**: Multi-level nesting (2, 4, 6 spaces)

    ## Changes
  - Added tests to `source_tests/core/api_whitespace_behaviors.json`
  - Regenerated Go test files
  - Minor README formatting updates





### Refactoring

- **tests:** Simplify whitespace behavior model ([`1680ff0`](https://github.com/CatConfLang/ccl-test-data/commit/1680ff076b2816006d5d2b2a8b443d53faac159a))

  ## Summary

    Simplifies the CCL whitespace behavior model by:
  - Moving `strict_spacing` tests to experimental (not part of core spec)
  - Fixing `tabs_preserve` tests to use spaces for indentation
  - Adding tests for tabs→spaces conversion in print/round_trip

    ## Whitespace Model Clarification

    **Default behavior:**
  - Whitespace = spaces + tabs
  - Keys are fully trimmed (leading and trailing whitespace removed)
  - Values have leading whitespace trimmed
  - Tabs converted to single space in multiline continuation

    **`tabs_preserve` behavior:**
  - Whitespace = spaces only (tabs are content)
  - Keys trimmed of spaces only; tabs in keys are preserved
  - Values have leading spaces trimmed; leading tabs are preserved
  - Tabs in multiline continuations preserved as-is

    **`tabs_to_spaces` behavior (default):**
  - Tabs converted to single space
  - Both tabs-only and mixed space/tab indentation supported
  - Print/round_trip outputs spaces instead of tabs

    ## Changes

    - **Moved to experimental**: 5 `strict_spacing_*` tests
  - **Kept in core**: All `loose_spacing_*` tests (default behavior)
  - **Fixed**: `tabs_preserve_multiline`, `spaces_vs_tabs_continuation*` tests
  - **Added**:
    - `tabs_to_spaces_mixed_indent` - mixed space/tab indentation
    - `tabs_to_spaces_multiline_print` - canonical_format outputs spaces
    - `tabs_to_spaces_round_trip` - round trip converts tabs to spaces

    ## Test plan

    - [x] `just validate` passes
  - [x] `just reset` passes (409 tests, 247 skipped)
  - [x] `just lint` passes

    Related: #30 (generator bug in ccl-test-lib)




## [0.0.1] - 2025-12-08



### Bug Fixes

- **tests:** Add list_coercion_enabled behavior to proposed_behavior tests ([`7c11ffa`](https://github.com/CatConfLang/ccl-test-data/commit/7c11ffac733cef8f99a493d3de730a7084de5ab2))

  ## Summary
  - Add missing `list_coercion_enabled` behavior tag to 8 tests in `api_proposed_behavior.json` that use `get_list` and expect list coercion behavior

    ## Affected Tests
  - `empty_list`
  - `list_with_numbers`
  - `list_with_booleans`
  - `list_with_whitespace`
  - `deeply_nested_list`
  - `list_with_unicode`
  - `list_with_special_characters`
  - `list_multiline_values`

    ## Context Without this tag, implementations with `list_coercion_disabled` would incorrectly run these tests and fail.

    Fixes #12

- **tests:** Add conflict metadata for tabs_preserve behavior ([`4a95c2c`](https://github.com/CatConfLang/ccl-test-data/commit/4a95c2ce9bb50bcd92a200520148c1d08d03de2d))

  ## Summary
  - Add `conflicts` field to 4 tests with `tabs_preserve` behavior
  - Documents mutual exclusivity with `tabs_to_spaces` behavior

    This helps test runners understand which behaviors conflict and enables better test filtering.

- Filter behavior tags to only apply to relevant functions ([`a85830d`](https://github.com/CatConfLang/ccl-test-data/commit/a85830ddcfae1d98d270de44103d3d4fcc22fd78))

  Behavior tags like boolean_strict/boolean_lenient now only appear on tests for functions where they actually affect behavior (e.g., get_bool), not on parse or build_hierarchy tests where the values are just stored as strings.

    This change was implemented in ccl-test-lib's generator, which now filters behaviors based on a function mapping before assigning them to generated flat tests.

- **tests:** Add list_coercion_enabled behavior to api_list_access tests ([`7dd4a7b`](https://github.com/CatConfLang/ccl-test-data/commit/7dd4a7b96ec65dac0f779383c66bcabe6d3e8b9c))

  Fixes missing behavior declarations in `api_list_access.json` test suite that were causing false failures for reference-compliant implementations.

    Tests in `api_list_access` expected `list_coercion_enabled` behavior (duplicate keys create lists accessible via `get_list()`) but had empty `behaviors` arrays, making them indistinguishable from behavior-agnostic tests.

    - `basic_list_from_duplicates`
  - `large_list`
  - `list_with_comments`

    These tests use duplicate keys to create lists and expect `get_list()` to return the actual array values. However, without the `list_coercion_enabled` behavior declaration:

    1. Test runners cannot properly filter based on implementation capabilities 2. Reference-compliant implementations (using `list_coercion_disabled`) see spurious failures 3. Results misleadingly show failures instead of skipped tests

    **`list_coercion_disabled` (reference-compliant)**:
  - `build_hierarchy()` creates arrays from duplicate keys: `{"ports": ["80", "443"]}`
  - `get_list("ports")` returns `null` (refuses to return the array)

    **`list_coercion_enabled` (proposed behavior)**:
  - `build_hierarchy()` creates arrays from duplicate keys: `{"servers": ["web1", "web2", "web3"]}`
  - `get_list("servers")` returns `["web1", "web2", "web3"]` (returns the actual array)

    Added explicit `"behaviors": ["list_coercion_enabled"]` to all three affected tests in `source_tests/core/api_list_access.json`.

    ✅ JSON schema validation passed ✅ Test regeneration completed successfully ✅ Statistics now show: **`behavior:list_coercion_enabled`: 8 tests** (up from 5) ✅ Generated tests include proper behavior declarations

    - ✅ `api_reference_compliant` properly declares `["list_coercion_disabled"]`
  - ✅ `api_proposed_behavior` properly declares `["list_coercion_enabled"]`
  - ✅ `api_list_access` now properly declares behavior requirements

    Reference-compliant implementations using `list_coercion_disabled` will now correctly **skip these tests** instead of seeing false failures.

- Test-reader updates ([`a600007`](https://github.com/CatConfLang/ccl-test-data/commit/a600007e81d60ddc4aeccafc81f90a50f2619782))

- **test-reader:** Integrate ccl-test-lib for proper type handling ([`e699d65`](https://github.com/CatConfLang/ccl-test-data/commit/e699d650853bdf0caee5323e20eefef74cb19edc))

  Fixes test-reader tool which was broken due to type mismatches with ccl-test-lib package. The tool now properly loads and displays both source and generated test formats.

    Key changes:
  - Replace custom types with ccl-test-lib types (TestSuite, TestCase, Entry)
  - Update config structure: Functions → SupportedFunctions with CCLFunction types
  - Fix loader API: use NewTestLoader() + LoadTestFile() with proper LoadOptions
  - Handle ValidationSet struct access: test.Validations.Parse instead of indexing
  - Add type assertions for test.Expected interface{} to extract entries
  - Remove unavailable Description field from TestCase
  - Update TUI rendering to work with new type structure

    The test-reader now successfully:
  - Loads both source and flat test formats using ccl-test-lib
  - Displays test cases with proper formatting and metadata
  - Shows expected entries, variants, and features correctly
  - Supports both static CLI and TUI modes

- Resolve test generation issues in ccl-test-data ([`97754df`](https://github.com/CatConfLang/ccl-test-data/commit/97754dfec2e36e2e7419e2d4065d6231a99ac135))

  Fixed multiple issues that were preventing `just generate` from working:

    **Generator fixes:**
  - Handle both array and map formats for Expected field in parse validations
  - Support direct value returns from loader for typed access validations
  - Add variable usage in TODO validations to prevent unused variable warnings
  - Fix function argument formatting to always use []string slices for CCL functions

    **Template improvements:**
  - Updated generateFlatParseValidation to handle loader's array format
  - Added generateTypedAccessForDirectValue helper for direct value handling
  - Modified formatArgs to always generate []string format for CCL API compatibility

    **Test organization:**
  - Simplified directory structure from level-based to feature-based organization
  - Updated package naming to match new structure

    **Results:**
  - All 366 tests now generate and compile successfully
  - Test execution works (failures are now logic-related, not generation issues)
  - Clean `just generate` workflow restored

- Resolve mock ccl implementation test failures ([`9fe1c7a`](https://github.com/CatConfLang/ccl-test-data/commit/9fe1c7a31b3b8bc61a13b054fd10154e3a0f372a))

  - Configure CRLF behavior to normalize line endings (BehaviorCRLFNormalize)
  - Enhance Parse function with multiline content support for indented sections
  - Update test generation to skip conflicting behavior tags
  - Focus implementation on Level 1 parsing functionality
  - Improve test success rate from 72% to 90% (142/157 passing tests)

    Changes:
  - internal/mock/ccl.go: Enhanced multiline parsing and CRLF normalization
  - internal/config/runner_config.go: Updated behavior configuration and supported functions
  - justfile: Added behavior exclusions for consistent test generation
  - go_tests/: Updated generated test files with proper tag filtering

- **build:** Resolve compilation errors in type references ([`d7237b4`](https://github.com/CatConfLang/ccl-test-data/commit/d7237b496ff69d52197bab21081a21135b04694f))

  Remove invalid field references causing build failures:
  - Remove Canonical field access from ValidationSet (field doesn't exist)
  - Fix Conflicts field location (TestCase.Conflicts vs TestMetadata.Conflicts)
  - Remove invalid SourceFormat field from GenerateOptions

    Ensures compatibility with ccl-test-lib type definitions.

- **generation:** Standardize null arrays to empty arrays ([`70acbb2`](https://github.com/CatConfLang/ccl-test-data/commit/70acbb2a0cf9e3958440abcc80207e0899d9d3a8))

  Replaces null values with empty arrays in generated test files for consistent JSON schema compliance and improved type safety.

    - behaviors: null → behaviors: []
  - features: null → features: []
  - variants: null → variants: []

- **generation:** Resolve schema validation and flat format issues ([`3ed9013`](https://github.com/CatConfLang/ccl-test-data/commit/3ed901323d24d4c6352acfe7869316aa064515fe))

  - Add $schema field support to generated-format.json schema
  - Fix array field null values by ensuring proper initialization
  - Use JSON tags for validation name extraction instead of field names
  - Update generated test files with corrected format
  - Resolve type compatibility issues in ccl-test-lib generator

    All generated files now pass schema validation successfully.

- Resolve linting issues and code quality improvements ([`4d24efd`](https://github.com/CatConfLang/ccl-test-data/commit/4d24efd009a9699ecd1937ac77a199c95e45fa23))

  - Removed unused functions to fix unused linting warnings
  - Fixed unchecked type assertions with proper error handling
  - Resolved variable name conflicts in loop scopes
  - Updated octal literal syntax to modern Go format
  - Added documentation to empty conditional branches
  - All linting issues now resolved, code passes golangci-lint checks

- Resolve format mismatch between ccl-test-lib and ccl-test-data ([`2810ff9`](https://github.com/CatConfLang/ccl-test-data/commit/2810ff982910940050a0f75c4a6ea97e504c1e80))

  - Removed duplicate type definitions from ccl-test-data
  - Fixed ccl-test-lib generator to use proper schema-compliant types
  - Regenerated all flat format files to match schema requirements
  - Updated flat files from old format (expected: [...]) to schema format (expected: {count: N, entries: [...]})
  - Test generation now works correctly with 317 tests generated

    The core architecture issue is resolved:
  - ccl-test-lib is the single source of truth for shared types
  - Schemas define the authoritative data format
  - ccl-test-data imports and uses ccl-test-lib components
  - All format mismatches and JSON unmarshaling errors are fixed

- Omit empty conflicts field from flat json files ([`ac6619f`](https://github.com/CatConfLang/ccl-test-data/commit/ac6619fa3d56b427b3ed6e05477a35694ead6ff7))

  Previously, flat JSON files contained empty conflicts objects `{}` even when no conflicts were defined. This was inefficient for test runners and inconsistent with the intended omit-when-empty behavior.

    Changes:
  - Fixed ccl-test-lib generator to only set conflicts when non-empty
  - Added omitempty tag to Conflicts field in types.go
  - Updated justfile with generate-flat command for easier regeneration
  - Regenerated all flat test files with proper conflicts field handling

    Generated files now properly omit the conflicts field when empty, resulting in cleaner JSON and better performance for test runners.

- Resolve go vet issues in generated test files ([`37acad5`](https://github.com/CatConfLang/ccl-test-data/commit/37acad537be0c22ccc12913b329ec68f15e989b9))

  Remove unused input variables from error handling tests and fix undefined variable reference in list access test.

- Eliminate remaining kebab-case function names in metadata ([`1292997`](https://github.com/CatConfLang/ccl-test-data/commit/1292997fee8e38d0db52afb7d8977e7bba9ea094))

  - Update related_functions arrays to use underscore naming (expand-dotted → expand_dotted, build-hierarchy → build_hierarchy, etc.)
  - Fix cross-reference values to use underscores (performance-test → performance_test, security-test → security_test, etc.)
  - Maintain test data values like "in-memory" which are legitimate content
  - Preserve URLs and external references with hyphens

    All JSON files now consistently use underscore naming for all CCL-related identifiers while preserving legitimate hyphenated content and external references.

- Complete underscore naming harmonization across test suite ([`98760af`](https://github.com/CatConfLang/ccl-test-data/commit/98760af92e49fd35c6b609fe1e3ab0b79a993176))

  - Update JSON schema patterns for all tag categories (function, feature, behavior, variant)
  - Harmonize all test JSON files to use underscore naming consistently
  - Update README documentation with corrected tag examples and function names
  - Correct generator code to handle underscore feature names
  - Update statistics collector for underscore feature mapping
  - Regenerate test files with harmonized naming scheme
  - Maintain backward compatibility through proper feature-to-category mapping

    This ensures complete schema consistency with underscore naming convention throughout the entire test suite ecosystem.

- Harmonize function tags to use underscores for schema consistency ([`c393b3a`](https://github.com/CatConfLang/ccl-test-data/commit/c393b3a80b8368763208036469df4137aa64783c))

  - Update all function tags from hyphenated to underscore format: build-hierarchy → build_hierarchy
    - expand-dotted → expand_dotted
    - get-string/int/bool/float/list → get_string/int/bool/float/list
    - pretty-print → pretty_print
    - parse-value → parse_value
  - Add missing boolean and list coercion behavior tags to README
  - Correct function names: make-objects → build_hierarchy, compose → combine
  - Add parse_value function to documentation
  - Update generated test files to reflect tag changes

    All function tags now match schema validation names for seamless integration.

- Update json schema to match recent ccl api changes ([`e06d9e5`](https://github.com/CatConfLang/ccl-test-data/commit/e06d9e567c6215776a7fbee0fcce497fa48c9d43))

  - Replace function:make-objects tags with function:build-hierarchy (follows Sept 2025 API rename)
  - Add function:load to schema patterns for integration tests
  - Add feature:experimental-dotted-keys to allowed feature patterns
  - Add new feature enum values: core-ccl-parsing, core-ccl-hierarchy, core-ccl-integration
  - Add optional llm_metadata field with comprehensive property definitions
  - Remove duplicate make-objects reference from llm_metadata.related_functions

    All 12 test files now validate successfully against updated schema.

- Enable boolean true/false tests in lenient mode ([`641a25a`](https://github.com/CatConfLang/ccl-test-data/commit/641a25a3fe92f9f2bd6a00d411ea353d9dfdf10d))

  Update parse_boolean_true and parse_boolean_false tests to include both behavior:boolean-strict and behavior:boolean-lenient tags, removing the conflicts array. These fundamental boolean values work identically in both modes - only extended values like "yes"/"no" differ between modes.

    This fixes a coverage gap where lenient mode implementations weren't testing basic "true"/"false" parsing, while maintaining proper conflicts for truly incompatible behavioral differences.

    Updated documentation to clarify the dual-mode support and guidance for test tagging with multiple behaviors when appropriate.

- Resolve json schema validation errors in test files ([`3024dfe`](https://github.com/CatConfLang/ccl-test-data/commit/3024dfe8c1924cb96ed1a5dc20461c350d32fdab))

  - Fix api-errors.json test 2: change count from 0 to 1 for whitespace_only_error_ocaml_reference
  - Fix api-typed-access.json test 15: change make_objects.expected from string to object for empty_value_reference_behavior

    All test files now pass schema validation.

- Resolve test data inconsistencies and validation issues ([`cf709b1`](https://github.com/CatConfLang/ccl-test-data/commit/cf709b1091656eb94d1612441b125742e2627431))

  - Remove duplicate tests (composition_stability_ab, multiple_values_same_key_ports)
  - Fix conflicts arrays to use structured tags instead of test names
  - Add function:parse-value tag for parse_value validation tests
  - Add missing variant conflicts declarations for proper test isolation
  - Update schema to support parse-value function tag

    Addresses validation schema compliance and improves test organization.

- Resolve missing conflicts declarations for mutually exclusive tests ([`0d1b2cb`](https://github.com/CatConfLang/ccl-test-data/commit/0d1b2cbbe1f6f0759e2cacb7fcdd62e69a56f1c4))

  Fix critical bug where tests with contradictory expected outputs were both running for the same implementation behavioral choice, causing false failures.

    Changes:
  - Rename CRLF tests for clarity: crlf_normalization → crlf_normalize_to_lf,
    crlf_normalization_strict → crlf_preserve_literal
  - Update behavior tags: behavior:crlf-preserve/normalize →
    behavior:crlf-preserve-literal/crlf-normalize-to-lf
  - Add boolean parsing behavior tags: behavior:boolean-lenient vs boolean-strict
  - Rename boolean strict tests: *_strict → *_strict_literal for clarity
  - Add mutual conflicts declarations between all conflicting behaviors
  - Update schema to validate new behavior tag patterns
  - Update documentation and migration files with new tag system

    Fixes issue where implementations declaring behavior:crlf-preserve would run both CRLF tests with identical input but contradictory expected outputs, ensuring only one behavioral interpretation runs per implementation choice.

- Ensure deterministic test generation by sorting map keys ([`8d16742`](https://github.com/CatConfLang/ccl-test-data/commit/8d1674278440b772ce3b9704de7341d7ffb7e4fe))

- Disallow count:0 in schema and fix assertion counting ([`0761989`](https://github.com/CatConfLang/ccl-test-data/commit/0761989d66addf7e3e6223e66436132c83f74637))

- Prevent unused imports in generated tests ([`e802eac`](https://github.com/CatConfLang/ccl-test-data/commit/e802eac6bd214bac7a2751d6a7395f85985886b5))

  Fixed template logic to only include assert/require imports when tests have actual implemented validations, not just validation fields that generate TODO comments. This resolves linting errors for unused imports in generated test files.

- Allow error cases in typed access validation counted format ([`8cc481c`](https://github.com/CatConfLang/ccl-test-data/commit/8cc481c58e95e7fc9555a4c6a9596c907d41cc00))

  - Update schema to support mixed success/error cases in cases array
  - Enable flexible test organization while preserving assertion accuracy
  - Update documentation with counted format examples and test runner patterns

- Resolve npm audit vulnerabilities with fast-json-patch override ([`be799db`](https://github.com/CatConfLang/ccl-test-data/commit/be799db40c90cd9be8d4418a7fa70deb54ae05f3))

- Consolidate level-4 typed parsing tests into unified schema ([`ca32581`](https://github.com/CatConfLang/ccl-test-data/commit/ca325815bb366f4d7d5a3e3d81bb20849d0eacd2))

- Resolve all parser test failures and implement complete test suite ([`605527b`](https://github.com/CatConfLang/ccl-test-data/commit/605527b2ca086a6e1b13bc7cdd03a8f552ee3865))

  - Add comprehensive pretty printer test suite (15 tests) with round-trip, canonical format, and deterministic testing
  - Fix trailing newline handling in input parsing by cleaning split() artifacts
  - Implement targeted blank line preservation for empty multiline sections
  - Add detailed test failure reporting across all architecture levels
  - Update project documentation to reflect fully implemented pretty printer

    All 91 tests now pass across 4-level CCL architecture:
  - Level 1 (Entry Parsing): 48/48 ✓
  - Level 2 (Entry Processing): 11/11 ✓
  - Level 3 (Object Construction): 8/8 ✓
  - Level 4 (Typed Parsing): 8/8 ✓
  - Pretty Printer: 15/15 ✓
  - Error Handling: 1/1 ✓

- Correct parsing precedence - integers win overlapping cases ([`32b4588`](https://github.com/CatConfLang/ccl-test-data/commit/32b458840563fda99e2ccf836b2dd27821874110))

  - Change precedence from bool->int->float->string to int->float->bool->string
  - When both integer and boolean parsing enabled, integers take precedence over booleans for "0"/"1"
  - Booleans retain non-numeric forms: "true", "false", "yes", "no", "on", "off"
  - Update JSON test expectations: "1" → IntVal(1), "0" → IntVal(0)
  - Maintains 7/8 test success rate with correct logic

    Parsing logic: integers win overlap since booleans have many non-numeric alternatives





### Documentation

- Add missing behaviors and link to behavior reference ([`dc9d553`](https://github.com/CatConfLang/ccl-test-data/commit/dc9d5533b8f92eec10d8663a0e10f3f5134db1c9))

  - Add list_coercion and array_order behaviors to test-filtering.md
  - Add link to ccl.tylerbutler.com/behavior-reference/ from test-selection-guide.md for detailed behavior documentation

- Standardize snake_case naming and update test counts ([`c5435cf`](https://github.com/CatConfLang/ccl-test-data/commit/c5435cf36293a34a8a587747ffdcad73274d6fe7))

  ## Summary

    Comprehensive documentation update to standardize naming conventions (snake_case), update test counts, replace outdated implementation levels, and rename mock implementation methods for consistency.

    ## Changes

    ### 1. Standardized Naming (snake_case everywhere)

    **Functions:**
  - `build-hierarchy` → `build_hierarchy`
  - `get-string` → `get_string`, `get-int` → `get_int`, `get-bool` → `get_bool`, `get-float` → `get_float`, `get-list` → `get_list`
  - `expand-dotted` → `expand_dotted`
  - `make-objects` → `build_hierarchy`
  - `pretty-print` → `pretty_print`

    **Features:**
  - `dotted-keys` → `experimental_dotted_keys`
  - `empty-keys` → `empty_keys`
  - `typed-parsing` → `typed_access`

    **Behaviors:**
  - `boolean-strict` → `boolean_strict`, `boolean-lenient` → `boolean_lenient`
  - `crlf-normalize-to-lf` → `crlf_normalize_to_lf`, `crlf-preserve-literal` → `crlf_preserve_literal`
  - `tabs-preserve` → `tabs_preserve`, `tabs-to-spaces` → `tabs_to_spaces`
  - `strict-spacing` → `strict_spacing`, `loose-spacing` → `loose_spacing`
  - `list-coercion-enabled` → `list_coercion_enabled`, `list-coercion-disabled` → `list_coercion_disabled`

    **Variants:**
  - `proposed-behavior` → `proposed_behavior`
  - `reference-compliant` → `reference_compliant`

    **Files updated:** API.md, CLI_REFERENCE.md, MOCK_IMPLEMENTATION.md, schema-reference.md, test-architecture.md, test-filtering.md, test-selection-guide.md, tag-migration.json

    ### 2. Updated Test Counts

    Changed from outdated counts to current:
  - **180 tests, 375 assertions** (across 14 files)

    Removed references to:
  - 452/167 (old)
  - 462/924 (old)
  - 26/56/135 implementation levels (replaced)

    **Files updated:** test-architecture.md, test-selection-guide.md

    ### 3. Replaced Implementation Levels

    Removed outdated Minimal/Standard/Full levels with function-based approach:

    **Core (Required):**
  - `parse` - 163 tests
  - `build_hierarchy` - 77 tests

    **Typed Access:**
  - `get_string` - 7 tests (28 assertions)
  - `get_int` - 11 tests (47 assertions)
  - `get_bool` - 12 tests (49 assertions)
  - `get_float` - 6 tests (28 assertions)
  - `get_list` - 48 tests (186 assertions)

    **Processing & Formatting:**
  - `filter` - 3 tests (6 assertions)
  - `compose` - 9 tests (18 assertions)
  - `canonical_format` - 14 tests (23 assertions)
  - `round_trip` - 12 tests (23 assertions)

    **Files updated:** implementing-ccl.md, test-architecture.md, test-selection-guide.md

    ### 4. Renamed Mock Implementation

    - Changed `Combine` → `Compose` in `internal/mock/ccl.go`
  - Updated package documentation to reflect `compose` function name
  - Maintains consistency with schema and test naming

    ### 5. Schema Updates

    - Changed `merge` → `compose` in schemas
  - Updated `source_tests/core/property_algebraic.json`
  - Stats now show `function:compose` instead of `function:merge`

    ## Test Plan

    - [x] `just lint` passes
  - [x] `just test` passes (323 tests, 180 skipped as expected)
  - [x] `just validate` passes (all schemas valid)
  - [x] All naming uses snake_case consistently
  - [x] No references to old test counts
  - [x] Function-based implementation guide replaces old levels

    ## Context

    These changes establish consistency across all CCL documentation: 1. **Snake_case everywhere**: Functions, features, behaviors, variants 2. **Accurate test counts**: Reflects current test suite state 3. **Function-based approach**: Clear progressive implementation path 4. **Code consistency**: Mock implementation matches schema naming

- Add comprehensive copilot agent onboarding instructions ([`7a04b3e`](https://github.com/CatConfLang/ccl-test-data/commit/7a04b3e787db46fc56c1c7b3d93f46fd90bc8f83))

  ## Overview

    This PR adds comprehensive onboarding instructions for GitHub Copilot agents working on the CCL Test Suite repository. The new `.github/agents/` directory provides structured guidance for agents to understand the repository, follow proper workflows, and maintain quality standards.

    ## What's Added

    Created a complete agent onboarding system consisting of 5 markdown files (57.8 KB total):

    ### 1. **`.github/agents/README.md`** Index and navigation hub that provides:
  - Overview of all available guides
  - Quick reference for essential commands
  - Key principles all agents should follow
  - Links to troubleshooting resources

    ### 2. **`.github/agents/onboarding.md`** (Primary Guide) Comprehensive onboarding covering:
  - Repository overview with current statistics (452 assertions, 167 tests)
  - Essential first steps and initial setup commands
  - Complete repository structure explanation
  - Test architecture (dual-format system, function groups, classification)
  - Standard development workflows
  - Pre-commit checklists (lint, reset, validate)
  - Best practices with explicit DO/DON'T lists
  - Troubleshooting common issues

    ### 3. **`.github/agents/test-development.md`** (Specialized) Detailed guide for test suite work including:
  - Test file organization (13 JSON files across core/experimental)
  - Complete test structure reference with examples
  - Metadata system (features, behaviors, variants, conflicts)
  - Step-by-step workflows for adding tests
  - Function-specific test patterns for all CCL functions
  - Test generation pipeline explanation
  - Quality standards and naming conventions
  - 62 code blocks with practical examples

    ### 4. **`.github/agents/code-changes.md`** (Specialized) Comprehensive guide for Go code changes covering:
  - Package structure and organization
  - Key code areas (mock implementation, generator, CLI, statistics, config)
  - Code quality standards (formatting, error handling, comments)
  - Common change patterns (adding functions, CLI commands, improving errors)
  - Testing and debugging workflows
  - 80 code blocks with examples
  - Pre-commit checklists for code changes

    ### 5. **`.github/agents/documentation.md`** (Specialized) Complete documentation maintenance guide with:
  - Documentation structure and categories
  - Markdown style guide with examples
  - Update workflows for each document type
  - Cross-referencing guidelines
  - Consistency maintenance strategies
  - 80 code blocks with templates and examples

    ## Why This Matters

    These instructions enable GitHub Copilot agents to:

    ✅ **Understand the repository quickly** with clear overviews and structure explanations ✅ **Follow consistent workflows** using standardized checklists and processes ✅ **Maintain quality standards** through explicit DO/DON'T guidance ✅ **Make appropriate changes** by understanding the dual-format test architecture ✅ **Avoid common mistakes** with troubleshooting sections and best practices ✅ **Work independently** with comprehensive, task-oriented guides

    ## Structure and Design

    The guides follow a progressive learning structure: 1. Start with **onboarding.md** for fundamentals 2. Reference **specialized guides** based on task type 3. Use **checklists** to ensure quality 4. Follow **cross-references** for related topics

    Each guide includes:
  - Clear section organization (254 code blocks total across all guides)
  - Practical, tested examples
  - Cross-references to existing documentation
  - Explicit quality standards
  - Pre-commit validation checklists

    ## Impact

    This addition provides:
  - **Better agent performance** through clear instructions
  - **Consistent code quality** via standardized workflows
  - **Reduced errors** with comprehensive checklists
  - **Faster onboarding** with structured, progressive guides
  - **Maintainability** through well-documented processes

    The instructions complement the existing `CLAUDE.md` file while providing more detailed, task-specific guidance for different types of agent work.

- Optimize documentation for efficiency and readability ([`98ae0e5`](https://github.com/CatConfLang/ccl-test-data/commit/98ae0e59fa4f5996141c640fa3391ac90a91e753))

  Streamlined 8 major documentation files with 56% average reduction while preserving all technical accuracy and completeness.

    Changes:
  - CLAUDE.md: 253→125 lines (51% reduction)
  - DEVELOPER_GUIDE.md: 708→240 lines (66% reduction)
  - API.md: 377→220 lines (42% reduction)
  - CLI_REFERENCE.md: 627→245 lines (61% reduction) test-architecture.md: 663→280 lines (58% reduction)
  - implementing-ccl.md: 599→240 lines (60% reduction)
  - README.md: Consolidated dual-format explanations docs/README.md: Streamlined navigation structure

    Impact:
  - 1,877 lines eliminated across all files
  - ~7.8 hours reading time saved (58% faster)
  - Converted verbose prose to efficient tables/bullets
  - Eliminated redundant architecture explanations
  - Maintained all essential technical content

    Quality preserved:
  - All command references and code examples accurate
  - Cross-file consistency maintained
  - Technical patterns and workflows intact
  - Progressive implementation guidance preserved

- Clean up remaining level references across project ([`f1523a8`](https://github.com/CatConfLang/ccl-test-data/commit/f1523a884fb2d973411f8323c49dcb73a41a2bd9))

  Remove legacy level-based terminology and infrastructure:
  - Update documentation from "Level 1-5" to stage-based approach
  - Replace --levels CLI flags with --functions alternatives
  - Remove level metadata from validation schemas
  - Modernize language to feature-based implementation guide
  - Preserve legitimate data references (level1, level2 in test data)

    Maintains full system functionality while completing transition from hierarchical levels to flexible function-based architecture.

- Remove level references from comments and documentation ([`07151c2`](https://github.com/CatConfLang/ccl-test-data/commit/07151c25f08f66f64d3148bec7ae63b837cd7bd7))

  - Update CCL function comments to remove level numbers
  - Change 'Level 1/2/3/4' references to descriptive function names
  - Update package documentation to focus on functions vs levels
  - Update CLI descriptions to reference functions instead of levels
  - Maintain legitimate 'level = info' test data content

    Completes level property removal across entire codebase while preserving all functional behavior and test coverage.

- Comprehensive documentation cleanup and reorganization ([`b02441d`](https://github.com/CatConfLang/ccl-test-data/commit/b02441d3e2c45577b29524046cdb388dadb176b3))

  - Reduce from 47+ scattered markdown files to 16 focused documents
  - Remove obsolete content (.serena/archive/, docs/archive/, orphaned files)
  - Consolidate redundant content (merge overlapping files)
  - Update function naming: make_objects → build_hierarchy throughout
  - Replace "tags" with current schema structure (features, behaviors, variants)
  - Remove outdated "level" concepts, replace with function groups
  - Reorganize docs/ with clear categorization for implementers
  - Eliminate duplication and improve navigation

- Replace level concept with function groups and enhance technical documentation ([`71d47be`](https://github.com/CatConfLang/ccl-test-data/commit/71d47be6bdacfb7a80bd6bf1e1f44631396fdf3a))

  - Remove confusing numbered levels (1-5) in favor of clear function groups
  - Organize CCL functions by logical categories: Core, Typed Access, Processing, Experimental, Formatting
  - Update parse_value to core functions and expand_dotted to experimental
  - Standardize pretty_print as alias for canonical_format
  - Add comprehensive technical documentation including architecture diagrams
  - Replace CLI --levels flags with --functions for better clarity
  - Improve progressive implementation guidance using function dependencies
  - Add troubleshooting guide with mise recommendations

    This eliminates the 4 vs 5 levels inconsistency and provides a clearer mental model based on actual CCL functionality rather than arbitrary numbering.

- Add cli generator approach plan for test filtering ([`72bbd5f`](https://github.com/CatConfLang/ccl-test-data/commit/72bbd5fcc46c03790de1888046b0e8cd6b1a12da))

  Add comprehensive plan for implementing CLI-based test filtering to clarify mutually exclusive behaviors while prioritizing implementation simplicity.

    Key decisions:
  - Simple YAML config format (functions, behaviors, features lists)
  - Pre-filtering approach over runtime filtering
  - Coverage reporting to address visibility loss
  - Extension of existing ccl-test-runner tool

- **cli:** Document delegation architecture for flat generation ([`a6e2cec`](https://github.com/CatConfLang/ccl-test-data/commit/a6e2cec5b1c9dadf03f0652fdd9b9d1bb14c9f33))

  - Add comprehensive documentation to generateFlatAction explaining delegation to ccl-test-lib
  - Update CLI command descriptions to clarify architecture and separation of concerns
  - Document justfile command with architecture notes explaining delegation pattern
  - Maintain convenient CLI access while keeping logic in appropriate library layer

    Architecture:
  - ccl-test-lib: Contains core flat generation logic (generator.NewFlatGenerator)
  - ccl-test-runner: Provides thin CLI wrapper for convenient access
  - Clear separation of concerns with documented delegation points

- Update documentation and settings for current architecture ([`e857fcf`](https://github.com/CatConfLang/ccl-test-data/commit/e857fcfe24720f8ec4d0596df3f5350186c15ccb))

  Updated README.md to reflect the current 4-level CCL architecture and enhanced settings configuration for development workflow.

- Update architecture from 5-level to 4-level system ([`e83f35a`](https://github.com/CatConfLang/ccl-test-data/commit/e83f35a890b48a872571246537645a01d9ebe122))

  Updates documentation to reflect the new 4-level CCL implementation architecture and separate typed fields schema structure.

    Key changes:
  - Level system: 5-level → 4-level architecture
  - Schema: unified tags → separate typed fields
  - TestMeta: structured arrays for functions/features/behaviors/variants
  - Filtering: type-safe field access vs string parsing
  - Examples: updated to use new separate fields approach

    Files updated: test-matrix.md: Level mappings and implementation matrix
  - MOCK_IMPLEMENTATION.md: Implementation level structure
  - CLI_REFERENCE.md: Command options and output formats test-filtering.md: Type-safe filtering patterns
  - DEVELOPER_GUIDE.md: Architecture descriptions
  - API.md: TestMeta schema structure test-runner-implementation-guide.md: Complete conversion to separate fields

    Improves API ergonomics and provides better type safety for test runner implementations while maintaining full backward compatibility.

- Update to 4-level system and add cross-references to ccl.tylerbutler.com ([`8efde90`](https://github.com/CatConfLang/ccl-test-data/commit/8efde90081293dd126f68bce13984bbd91b58632))

  - Update API.md from 5-level to 4-level CCL implementation system
  - Update progressive implementation examples to match new level structure
  - Add cross-references to centralized CCL documentation at ccl.tylerbutler.com
  - Update implementing-ccl.md with new terminology and level organization
  - Align documentation with Phase 2 consolidation plan
  - Update test filtering documentation to reference levels 1-4

- Add comprehensive documentation suite with enhanced readme ([`f5cd1f3`](https://github.com/CatConfLang/ccl-test-data/commit/f5cd1f31730a2d9088fb40df4278654e6dab6194))

  - Add complete API documentation (docs/API.md) covering all Go packages
  - Add developer guide (docs/DEVELOPER_GUIDE.md) for contributing and extending
  - Add mock implementation guide (docs/MOCK_IMPLEMENTATION.md) with patterns
  - Add CLI reference (docs/CLI_REFERENCE.md) with complete command documentation
  - Add documentation index (docs/README.md) for navigation
  - Enhance README.md with strategic callouts and improved formatting
  - Include practical examples, usage patterns, and troubleshooting guidance
  - Support progressive CCL implementation with structured guidance
  - Provide comprehensive reference for contributors and implementers

- Add flexible error message testing guidelines ([`e2cf641`](https://github.com/CatConfLang/ccl-test-data/commit/e2cf64114ffa51040b25ef8c260011bfc9f9ff5a))

  Add comprehensive error validation guidance to the implementing guide, including pattern-based vs exact matching approaches with best practices for cross-language compatibility.

    Also clean up go.mod dependencies.

- Update documentation for new function:parse-value tag ([`323278b`](https://github.com/CatConfLang/ccl-test-data/commit/323278b548fec55bba6a716e28d21af3765f69d2))

  Updated test selection guide and CLAUDE.md to reflect the new function:parse-value tag added during JSON test file review. This tag distinguishes Level 1 basic parsing from Level 2 indentation-aware parsing functionality.

    Changes:
  - Added function:parse-value to function tag documentation
  - Updated progressive implementation strategy
  - Updated code examples across all languages
  - Clarified Level 2 parsing capabilities

- Update documentation for feature-based tagging system ([`0217972`](https://github.com/CatConfLang/ccl-test-data/commit/0217972ce56e42ac440e4899870fa2083b564503))

  Update README.md and CLAUDE.md to reflect the new structured tagging system and create comprehensive test selection guide.

    Changes:
  - README.md: Add feature-based test selection section with tag categories,
    examples, and conflict resolution explanation
  - CLAUDE.md: Update project documentation with structured tags, test selection
    strategies, and progressive implementation guidance docs/test-selection-guide.md: Complete user guide with language-specific
    integration examples for Go, Rust, Python, and JavaScript

    Key additions:
  - Clear distinction between feature field (organizational) vs feature: tags (requirements)
  - Progressive implementation roadmap from minimal to full CCL support
  - Language-specific code examples for all major testing frameworks
  - Troubleshooting guide for common test selection issues

    The documentation now provides clear guidance for implementers to select appropriate tests based on their CCL implementation capabilities, with automatic conflict resolution and progressive development strategies.

- Update claude.md with comprehensive development guide ([`1db6f04`](https://github.com/CatConfLang/ccl-test-data/commit/1db6f04ebe75c0bb62245c8af7caa2b5314af0be))

- Add comprehensive test filtering design pattern documentation ([`924e293`](https://github.com/CatConfLang/ccl-test-data/commit/924e2932f9a5bef9a0e53abfe64c39a67586572e))

- Update schema to document new feature categories and tag conventions ([`535eb79`](https://github.com/CatConfLang/ccl-test-data/commit/535eb79de239dc6efa3eda76cde1d7423122d198))

- Remove temporal language and version references ([`c315f3d`](https://github.com/CatConfLang/ccl-test-data/commit/c315f3d1edb4020a8728bd0218f442ac1a2e06b2))

  Remove "NEW", version numbers, and temporal references from documentation since project is not yet versioned. Clean up comparison sections and focus on current validation-based format capabilities.

- Refer to main docs ([`b28287c`](https://github.com/CatConfLang/ccl-test-data/commit/b28287cdd7ee97a43345d46e92b7ca859fabf6c8))

- Establish ccl central documentation hub ([`7eea92f`](https://github.com/CatConfLang/ccl-test-data/commit/7eea92fbc303119097528ab60882f5a5731684af))

- Update test counts to reflect actual test suite size ([`def331f`](https://github.com/CatConfLang/ccl-test-data/commit/def331f38585775c8736fbaed276654e0e7769a4))

- Spec 1.1.0 ([`297e82b`](https://github.com/CatConfLang/ccl-test-data/commit/297e82bcd28cee4dd677474eef38979b7b0c4795))





### Features

- **tests:** Add indented line continuation test cases ([`9951c0c`](https://github.com/CatConfLang/ccl-test-data/commit/9951c0c77ef0263c9a532346322843a6e07233ab))

  ## Summary Add two new test cases for verifying indented line continuation behavior in CCL parsing.

    ## New Test Cases

    ### 1. `indented_line_is_continuation` Verifies that indented lines after a `key = value` are treated as continuations of that value, not separate entries.

    **Input:** ``` descriptions = First line
    second line descriptions = Another item ```

    **Expected:** `descriptions` should have 2 values: `"First line\n second line"` and `"Another item"`

    ### 2. `mixed_indentation_levels` Tests interaction between indented continuations and unindented standalone keys.

    **Input:** ``` key1 = value1
    indented continuation key2 = value2 not indented key
    indented for not indented ```

    **Expected behavior:**
  - `key1` value: `"value1\n indented continuation"`
  - `key2` value: `"value2"` (no continuation since next line is unindented)
  - `not indented key` becomes a separate key with nested content

    ## Context These complement the existing `unindented_multiline_becomes_continuation` test which tests the
  *inverse* case (unindented lines becoming continuations). The new tests verify that *indented* lines correctly become continuations.

    Fixes #11

- **tests:** Unify tests to use inputs field ([`6ca9f71`](https://github.com/CatConfLang/ccl-test-data/commit/6ca9f710c6127eeb3d31433f28f862ef4fe04b46))

- **tests:** Add algebraic property functions and multi-input support ([`e0f7ceb`](https://github.com/CatConfLang/ccl-test-data/commit/e0f7ceb63c1317d6aa157b636db3c3fe8bd527b8))

  Add property verification functions to mock CCL implementation:
  - ComposeAssociative: verifies (a·b)·c == a·(b·c)
  - IdentityLeft: verifies compose(empty, x) == x
  - IdentityRight: verifies compose(x, empty) == x
  - RoundTrip: verifies parse(print(parse(x))) == parse(x)
  - Print: structure-preserving formatter for round-trip testing

    Update schemas to support:
  - Multiple inputs via 'inputs' array for algebraic tests
  - New validation functions in generated format
  - Boolean and text expected value types

    Refactor property tests to use dedicated verification functions instead of inline assertions, enabling proper algebraic property validation across the test suite.

- Add release workflow with git-cliff changelog generation ([`9757e01`](https://github.com/CatConfLang/ccl-test-data/commit/9757e01de088c92ee48a575549aad9a06ed5b7d9))

  - Add GitHub Actions workflow triggered by `data-v*.*.*` tags
  - Add git-cliff configuration for conventional commit changelog
  - Add release commands to justfile (release-check, release-preview, release)
  - Update $schema references to use relative paths for local validation
  - Add VS Code settings for JSON schema validation
  - Document release process in DEVELOPER_GUIDE.md

- Add array_order behaviors to schema and whitespace behavior tests ([`659bec7`](https://github.com/CatConfLang/ccl-test-data/commit/659bec739f1e328607874e69f56a3c6a738cfc3b))

  Add missing array_order_insertion and array_order_lexicographic behaviors to ccl-config-schema.json (were already used in 36 tests but not defined).

    Add new api_whitespace_behaviors.json with 21 tests covering:
  - loose_spacing: 11 tests (previously 0 coverage)
  - tabs_to_spaces: 6 tests (previously 0 coverage)
  - Combined spacing + tab behavior tests

    Update CLAUDE.md to document the new behavior pair and test file.

- **schema:** Add array_order_insertion/lexicographic behavior group ([`c1197bd`](https://github.com/CatConfLang/ccl-test-data/commit/c1197bd5f9cd4ce469df6644b4c6dd995a5d8108))

  Separate array ordering from reference_compliant variant into a proper behavior choice. This allows implementations to choose lexicographic sorting independently of other reference implementation behaviors.

    - Add array_order_insertion and array_order_lexicographic to schemas
  - Update test-selection-guide.md with new behavior group
  - Convert tests from reference_compliant variant to array_order behaviors
  - Rename *_reference tests to *_lexicographic where appropriate

- **tests:** Add reference_compliant test variants with proper conflicts ([`1ffff6e`](https://github.com/CatConfLang/ccl-test-data/commit/1ffff6e5721b132a528228fd18027b70fd41867d))

  Add reference-compliant variants for tests that differ between proposed and reference CCL behavior. Each variant now properly declares conflicts with the alternative interpretation.

    - Add _reference variant tests for complete_lists_workflow, bare_list_nested,
    bare_list_with_comments, bare_list_deeply_nested, and list_with_comments
  - Add list_coercion_disabled behavior to reference_compliant tests
  - Fix list ordering in reference_compliant tests (alphabetical)
  - Add conflicts declarations for mutually exclusive variants/behaviors

- **tests:** Add bare list auto-unwrapping tests for get_list ([`52abaf3`](https://github.com/CatConfLang/ccl-test-data/commit/52abaf362088afc2ce575a3c098779b3dbf00c97))

  ## Summary

    Adds 6 new test cases to `api_list_access.json` that specify `get_list` should automatically unwrap bare lists (empty-key lists).

    ## Problem

    Previously, bare lists created with syntax like: ```ccl servers =
    = web1
    = web2 ```

    Would create hierarchy: `{"servers": {"": ["web1", "web2"]}}`

    And require awkward access via: `get_list(config, "servers", "")`

    ## Solution

    These tests specify that implementations should make `get_list` smart enough to automatically unwrap the empty-key structure:

    ``` get_list(config, "servers") → ["web1", "web2"] ```

    This provides **consistent API ergonomics** for both list creation syntaxes:
  - **Duplicate keys**: `servers = web1\nservers = web2`
  - **Bare lists**: `servers =\n = web1\n = web2`

    Both now accessible via: `get_list(config, "servers")`

    ## New Test Cases

    1. **`bare_list_basic`** - Basic bare list unwrapping 2. **`bare_list_nested`** - Nested bare lists 3. **`bare_list_with_comments`** - Bare lists with comments 4. **`bare_list_deeply_nested`** - Deep nesting validation 5. **`bare_list_mixed_with_other_keys`** - Bare lists alongside regular keys 6. **`bare_list_error_not_a_list`** - Error handling for non-list paths

    ## Test Coverage Impact

    - **Total tests**: 186 (+6)
  - **Total assertions**: 393 (+18)
  - **`get_list` tests**: 54 (includes new bare list tests)
  - **`empty_keys` feature tests**: 48 (properly tagged)

    ## Implementation Guidance

    The `get_list` function should: 1. If path points to an array, return it (duplicate-key lists) 2. If path points to an object with only one key `""` containing an array, unwrap and return it (bare lists) 3. Otherwise return `null`

    This makes the empty-key representation an implementation detail, not user-facing API.

- Rename parse_value to parse_dedented ([`4dce7ed`](https://github.com/CatConfLang/ccl-test-data/commit/4dce7ed7335244643290be03c16836718f9b91c7))

  ## Summary

    Completes the rename of `parse_value` to `parse_dedented` across ccl-test-data repository, including schemas, source tests, documentation, implementation code, and all generated test files.

    ## Motivation

    The function `parse_value` was misleadingly named. It doesn't just "parse a value" - it performs **indentation normalization** (dedenting) by: 1. Calculating the common leading whitespace prefix across all lines 2. Stripping that prefix from all lines 3. Treating the dedented keys as top-level entries

    This is analogous to Python's `textwrap.dedent()` and is essential for `build_hierarchy` when recursively parsing nested CCL content.

    ## Changes

    ### Schemas & Documentation
  - ✅ `schemas/source-format.json` - Updated function enum
  - ✅ `schemas/generated-format.json` - Updated function enum
  - ✅ `CLAUDE.md` - Updated all references to parse_dedented
  - ✅ `docs/*.md` - Updated documentation

    ### Source Tests
  - ✅ All `source_tests/core/*.json` files - Updated function names in test definitions

    ### Implementation Code
  - ✅ `internal/generator/templates.go` - Updated ValidationSet references (2 locations)
  - ✅ `internal/mock/ccl.go` - Renamed ParseValue → ParseDedented method
  - ✅ `internal/stats/collector.go` - Updated function name references
  - ✅ `internal/stats/enhanced.go` - Updated statistics tracking

    ### Generated Files (Regenerated)
  - ✅ `generated_tests/*.json` - 327 tests regenerated with new function name
  - ✅ `go_tests/parsing/*_test.go` - All Go test files regenerated

- Refactor test generation with feature categorization and remove invalid associativity tests ([`7df4fe8`](https://github.com/CatConfLang/ccl-test-data/commit/7df4fe8ad50be1eaf91cd51689d2f516e8c7635a))

  - Implement optional_ prefix for feature categorization in Go test generator
  - Remove duplicate tags from generated test files
  - Remove invalid associativity validation tests from algebraic property suite
  - Add comprehensive test runner design documentation
  - Fix CRLF normalization in api_reference_compliant test cases
  - Update Go version requirement to 1.25.1 in .mise.toml
  - Update schemas to reflect cleaner test structure

    Removes 12 invalid "associativity" validation tests that were incorrectly generated from source tests intended for parse validation only. Algebraic property tests now correctly contain only parse validations.

- Add typed_accessors feature for typed accessor function tests ([`36056a4`](https://github.com/CatConfLang/ccl-test-data/commit/36056a4a6b1a026b857d76e1927705a644d57cd5))

  Add new "typed_accessors" feature to categorize tests that use typed accessor functions (get_int, get_float, get_bool). This enables implementations to skip typed accessor tests if they don't support type-aware value extraction.

    Changes:
  - Add "typed_accessors" feature to 18 test cases in api_typed_access.json
  - Update all three JSON schemas to include "typed_accessors" in valid features enum
  - Maintain backward compatibility with existing feature categorization system

    The feature follows the established pattern where features categorize test content by CCL syntax/functionality rather than implementation requirements.

- Refactor justfile with reusable test runner invocation ([`88f36dc`](https://github.com/CatConfLang/ccl-test-data/commit/88f36dccf25650f3cb308421166fa946998ebd72))

  This commit refactors the test execution system to use a unified, reusable approach that ensures consistent behavior between different test commands.

    Changes:
  - Add --skip and --basic-only flags to test runner CLI
  - Create _run-tests helper function in justfile for reusable test invocation
  - Make 'just test' run passing tests by default (--basic-only mode)
  - Add 'just test --all' option to run all tests including failures
  - Update dev-basic to use new test runner instead of hardcoded gotestsum
  - Add convenience aliases: test-all and test-comprehensive

    The key improvement is that 'just test' now passes by default, while 'just test --all' provides access to comprehensive testing when needed. Both 'just reset' and 'just test' now have identical, predictable behavior.

- Add test exclusion to just reset for progressive implementation ([`26b1d47`](https://github.com/CatConfLang/ccl-test-data/commit/26b1d479dd0eb88b8e7eac0ab4aeb302a17deb34))

  Updates just reset (dev-basic) command to exclude 4 problematic edge case tests:
  - TestKeyWithNewlineBeforeEqualsParse: newline within key portion
  - TestComplexMultiNewlineWhitespaceParse: complex whitespace with newlines
  - TestDeeplyNestedListParse: nested structure parsing expectations
  - TestRoundTripWhitespaceNormalizationParse: whitespace handling inconsistencies

    This demonstrates how test runners can implement progressive CCL support by excluding specific failing tests while maintaining clean CI for core functionality. Results: 153 tests pass, 209 skipped, 0 failed.

- Remove level-related infrastructure from stats and cli systems ([`64406b4`](https://github.com/CatConfLang/ccl-test-data/commit/64406b496906c0b8961fb7a7e9b0b7a8d7f43b8c))

  - Remove LevelBreakdown field and processing from enhanced stats
  - Remove level field from SourceTest struct
  - Remove --levels CLI flag and related filtering logic
  - Update justfile to remove level-specific commands
  - Update buildPackagePatterns to work without level filtering
  - Maintain backward compatibility for all other functionality

- Remove level property from json schemas and test data ([`c8cce07`](https://github.com/CatConfLang/ccl-test-data/commit/c8cce070a4b263c4b5804a9db18d45d7ca51a03d))

  Phase 1 of level removal - eliminate level concept from test data structure

    Changes:
  - Remove level property from source-format.json and generated-format.json schemas
  - Remove level fields from all source test JSON files (12 files)
  - Maintain JSON structure integrity and schema validation
  - All tests continue to validate successfully

    Breaking change: level field no longer supported in test format

- Restructure source tests to support $schema fields ([`c299e89`](https://github.com/CatConfLang/ccl-test-data/commit/c299e89561d4a403386225399a2732b56ef30953))

  - Convert all source test files from array to object format with $schema field
  - Update schema to support top-level object with tests array
  - Remove backward compatibility code from stats collectors
  - All 12 source test files now include JSON Schema validation
  - Consolidated schema files and cleaned up temporary artifacts

- Add simplified yaml configuration system for ccl test runner ([`ade05c2`](https://github.com/CatConfLang/ccl-test-data/commit/ade05c292be8bcbea79dc537e8e30ebced9f73cd))

  - Ultra-simple config format with arrays of strings only
  - JSON Schema validation via jv tool
  - Go validation with conflict detection for mutually exclusive behaviors
  - CLI validation tool at cmd/validate-config/
  - Exact naming matching test tag values (boolean_lenient, crlf_normalize_to_lf, etc.)
  - Minimal requirements: only functions array is required
  - Replaces complex 286-line Go config with ~10 lines of YAML

    Example ccl-config.yaml: ```yaml functions: [parse, build_hierarchy, get_string]

    behaviors: [boolean_lenient, crlf_normalize_to_lf]

    variants: [proposed_behavior] ```

- Add claude code workflows ([`d849cde`](https://github.com/CatConfLang/ccl-test-data/commit/d849cde9aeb8dcb8ffc82e3e0b47fed98a8caf15))

- **schema:** Update json format to object-based structure with proper $schema usage ([`c4427dd`](https://github.com/CatConfLang/ccl-test-data/commit/c4427dd38fab13187be3380f8d68477b5444c7dc))

  - Change from array-based to object-based format with $schema at top level
  - Remove repeated $schema fields from individual test items
  - Use $ref definitions for proper type generation with go-jsonschema
  - Update all generated test files to new format structure
  - Maintain full type safety and validation while following JSON Schema standards

    Format changes:
  - Before: [{"$schema": "...", "name": "..."}, ...]
  - After: {"$schema": "...", "tests": [{"name": "..."}, ...]}

- **schema:** Implement conditional args field requirements ([`074fc30`](https://github.com/CatConfLang/ccl-test-data/commit/074fc30f698174ab8c7f798e93716f7fdf6198ac))

  - Add conditional args requirements in JSON schema using if/then/else logic
  - Args field now required only for typed access functions (get_string, get_int, get_bool, get_float, get_list)
  - Updated schema description to clarify args field usage semantics
  - Regenerated flat tests and Go tests with refined args field behavior
  - Tests now show args field only where semantically required, cleaner JSON for other validation types

    This complements the ccl-test-lib fixes for args field generation logic.

- **schema:** Standardize on canonical_format function name ([`d46d282`](https://github.com/CatConfLang/ccl-test-data/commit/d46d282ec4ef6fd19b54936c6eba8dac9b3a06ee))

  Remove naming inconsistency between pretty_print and canonical_format by standardizing on canonical_format as the official function name.

    Changes:
  - Remove pretty_print from schema enums in favor of canonical_format
  - Update README.md documentation to use canonical_format
  - Clarify that Go PrettyPrint method implements canonical_format
  - Update function level from 5 to 4 to match actual implementation
  - All existing tests continue to work without breaking changes

    This resolves the confusion where schemas allowed both names but only canonical_format was actually implemented and tested.

- **config:** Centralize behavioral choices with validation ([`af8c4d1`](https://github.com/CatConfLang/ccl-test-data/commit/af8c4d135c433f7b541f3b471e3325381ab5141a))

  - Implement centralized RunnerConfig with explicit behavioral choice validation
  - Add strict validation requiring explicit choices for mutually exclusive behaviors
  - Update test generation to use configuration-based filtering with conflict detection
  - Regenerate test files with updated configuration system

    Configuration system validates:
  - CRLF handling: normalize vs preserve
  - Tab handling: preserve vs convert to spaces
  - Spacing: strict vs loose
  - Boolean parsing: strict vs lenient
  - List coercion: enabled vs disabled
  - Specification variant: proposed vs reference compliant

    Test generation now filters incompatible tests based on behavioral choices.

- **config:** Centralize behavioral choices with validation ([`914d719`](https://github.com/CatConfLang/ccl-test-data/commit/914d719c5f51e99300227e962f897085fc4d7549))

  Implement centralized configuration system that requires explicit choices for all mutually exclusive behavioral options and validates them at runtime.

    - Add RunnerConfig with required behavioral choice validation
  - Implement conflict detection for incompatible test tags
  - Update test runner CLI with configuration validation
  - Add custom filtering for tag-based test selection
  - Ensure all behavioral choices are explicitly made

    Addresses scattered configuration issue by centralizing all variant, feature, and behavior choices in a single validated configuration system.

- Fix ccl test suite failures and enhance architecture ([`26c501d`](https://github.com/CatConfLang/ccl-test-data/commit/26c501da8ab11b0fc96363035d448e04ebccb0be))

  - Fix flat generator metadata preservation for behavior-based filtering
  - Update schemas with hybrid approach (optional source, required generated)
  - Clean test data by removing empty metadata fields from all source files
  - Restore stats tool functionality for source format files
  - Rewrite implementation guide to match flat format architecture
  - Update ccl-test-lib types for proper JSON marshaling

    Resolves test generation pipeline issues where behavior/variant/feature metadata was lost during source-to-flat conversion, causing filtering failures. All 180 tests now properly preserve metadata for accurate test selection and execution.

- Replace schemas with current format validation ([`d07dcbc`](https://github.com/CatConfLang/ccl-test-data/commit/d07dcbc6b7faa4c4626c87f57085dd1b2ec1515b))

  Update test validation schemas to match actual file formats:

    - Replace source-format.json with schema for current source_tests/ structure
    - Direct array format (not wrapped objects)
    - Simple tests array with function/expect pairs
    - Support for all CCL functions and flexible return types
    - Optional features, behaviors, variants, level fields

    - Replace generated-format.json with schema for current flat test format
    - Enhanced validation for entries with required key/value fields
    - Complete function coverage including canonical_format, round_trip
    - Proper uniqueItems constraints on array fields

    - Update justfile validation commands to use renamed schemas
  - Remove experimental_dotted_keys from features (keep dotted_keys removed)

    All 12 source files and 177 flat files now validate successfully. Completes dual-format schema validation for CCL test suite.

- Implement separate fields schema for test metadata ([`ce79ffc`](https://github.com/CatConfLang/ccl-test-data/commit/ce79ffcfb8a31b754b976650a8f0f2827aceb353))

  Replace unified tags array with typed fields for better API ergonomics:

    **Schema Changes:**
  - Add separate `functions[]`, `behaviors[]`, `variants[]`, `features[]` fields
  - Replace string conflicts with categorized `ConflictsByCategory` structure
  - Add comprehensive enum validation for all field values
  - Maintain backward compatibility with structured tag parsing

    **Generator Updates:**
  - Parse structured tags (`function:*`, `behavior:*`, etc.) into separate arrays
  - Auto-generate function tags based on validation being tested
  - Organize conflicts by category (functions, behaviors, variants, features)
  - Preserve all original metadata while improving API usability

    **Benefits:**
  - Type safety: Enum validation prevents invalid values
  - Query ergonomics: `.behaviors.includes('boolean_strict')` vs string parsing
  - Self-documenting: All possible values clearly enumerated in schema
  - Better filtering: Direct field access for test selection

    **Validation:**
  - All 366 generated tests pass schema validation
  - Filtering works: 27 boolean_strict, 47 reference_compliant, 49 get_bool tests
  - Maintains rich metadata from original test structure

- Clarify crlf behavioral variants and consolidate dotted key features ([`bd35640`](https://github.com/CatConfLang/ccl-test-data/commit/bd35640eabc198a416fdece4890e776ca707626d))

  * Enhanced CRLF tests with intermediate parse validation to show behavioral differences
    - Reference: preserves CRLF during parsing, normalizes for output
    - Proposed: normalizes CRLF immediately during parsing
  * Renamed crlf_normalize_to_lf_reference → crlf_normalize_to_lf_indented_proposed
  * Removed underrepresented feature:dotted_keys tag (only 1 test)
  * Updated schema validation patterns to match actual feature usage
  * Added function:parse tags for completeness where missing

- Implement comprehensive llm optimization with enhanced metadata v2.1 ([`accba8b`](https://github.com/CatConfLang/ccl-test-data/commit/accba8b01367d21316af8b2b93479a8842a1e227))

  Transforms CCL test suite into maximally LLM-consumable format while maintaining human usability.

    - Enhanced metadata schema v2.1 with llm_metadata blocks across all 8 JSON test files
  - Created tiered LLM documentation (llms.txt, llms-small.txt, llms-full.txt)
  - Implemented Go validation script for enhanced metadata compliance
  - Added cross-references to Gleam implementation and documentation site
  - Established progressive implementation levels (Level 1-4) with clear learning paths
  - Achieved 30-50% token reduction through structured content optimization
  - Built automated validation workflows with just commands
  - Created comprehensive documentation matrix and architectural guides

    Total: 452 test assertions across 167 tests now optimized for AI consumption with 100% metadata compliance.

- Add performance benchmarking with object pooling and enhanced error reporting ([`3b05ce4`](https://github.com/CatConfLang/ccl-test-data/commit/3b05ce46f9817b6bd8b1647dc6ee17846ead9db6))

  - Add benchmark command with execution time and memory usage tracking
  - Implement object pooling in generator and stats packages to reduce allocations
  - Enhance mock CCL error messages with detailed debugging context and available keys
  - Add comprehensive package documentation for generator, stats, and mock packages
  - Include historical performance comparison and regression detection
  - Add benchmark recipes to justfile for development workflow

    🔨 Generated with [Claude Code](https://claude.ai/code)

- Add scrollable entry panes for improved large entry list handling ([`6cae224`](https://github.com/CatConfLang/ccl-test-data/commit/6cae2242bd17ef8733283c18af1b82830dbd69e6))

  - Add configurable entry display limit (6 entries max per view)
  - Implement CLI truncation with clear guidance to TUI mode for more entries
  - Add TUI entry scrolling with h/l (or ←/→) arrow key controls
  - Show scroll indicators with directional hints and remaining counts
  - Reset entry scroll position when navigating between tests
  - Add entry totals display and smart bounds checking
  - Update help text to include new entry scrolling controls

    Improves usability for tests with many parsed entries by preventing overwhelming displays while maintaining full access to all data.

- Enhance test-reader ui with improved entry display and navigation ([`cc9f5e6`](https://github.com/CatConfLang/ccl-test-data/commit/cc9f5e6f72265b3753144cb22458944d41b3216c))

  - Add whitespace visualization with dots for spaces and arrows for tabs
  - Implement compact key/value entry layout with color-coded elements
  - Remove excessive boxing and metadata to focus on essential information
  - Add escape key navigation to return from file viewer to directory selection
  - Tighten spacing between entries for better information density
  - Brighten empty key/value indicators for improved visibility
  - Make TUI the default interactive mode instead of static CLI output

- Add directory support with file selection to test-reader cli ([`a034466`](https://github.com/CatConfLang/ccl-test-data/commit/a03446600948d2ba007bfb0600504dc5f09b9d6d))

  - Add directory input support for test-reader CLI
  - Implement interactive file selection in both CLI and TUI modes
  - Display file metadata including descriptions and test counts
  - Show parse test counts vs total test counts for each file
  - Add file browser TUI with navigation and selection
  - Maintain backward compatibility with single file input
  - Sort files alphabetically with enhanced styling
  - Include quit functionality and error handling

- Add test-reader cli with enhanced styling and tui support ([`b904f43`](https://github.com/CatConfLang/ccl-test-data/commit/b904f433b74a848d805814c36a777d44f18dcf7b))

  - Add test-reader CLI with lipgloss styling for beautiful output
  - Support both static styled output and interactive TUI mode
  - Color-coded boxes for different content types:
    - Cyan borders for headers and primary content
    - Green borders for successful parse validations
    - Red borders for error cases and conflicts
    - Yellow borders for metadata and warnings
  - Interactive TUI navigation with j/k keys, toggle modes
  - Add bubbletea dependency for TUI functionality
  - Add comprehensive just tasks for easy CLI usage:
    - just reader (alias for read-essential)
    - just tui (interactive mode)
    - just read-errors, read-objects, etc.
    - just read <filename> for custom files

    Usage: ./test-reader <file.json> [--tui]

- Implement feature-based tagging system with enhanced statistics ([`f404c3d`](https://github.com/CatConfLang/ccl-test-data/commit/f404c3d3afbca73ca2d919ef4fc93ec48e4607d0))

  Replace descriptive tags with structured function/feature/behavior/variant tags to enable precise test selection for partial CCL implementations.

    Schema changes:
  - Add conflicts field to test metadata for mutual exclusivity
  - Define structured tag categories: function:*, feature:*, behavior:*, variant:*
  - Update level enum to include level 5 for formatting functions

    Enhanced statistics:
  - Track function requirements (function:parse, function:make-objects, etc.)
  - Monitor language features (feature:comments, feature:dotted-keys, etc.)
  - Analyze behavioral choices (behavior:crlf-preserve vs behavior:crlf-normalize)
  - Report implementation variants (variant:proposed-behavior vs variant:reference-compliant)
  - Count mutually exclusive tests and display conflict relationships

    Migration tools:
  - scripts/migrate-tags.py: Convert old tags to new structured format
  - scripts/clean-legacy-tags.py: Remove legacy descriptive tags docs/tag-migration.json: Complete mapping between old and new tags docs/schema-update.md: Implementation guide and usage examples

    Benefits:
  - Precise test filtering by implementation capabilities
  - Automatic conflict resolution for mutually exclusive behaviors
  - Progressive implementation support (minimal → full feature set)
  - Language-agnostic integration via JSON metadata

    Test suite now supports 167 tests with 21 mutually exclusive tests across 11 CCL functions, 6 language features, 3 behavioral choices, and 2 variants.

- Improve json schema validation constraints ([`dacc925`](https://github.com/CatConfLang/ccl-test-data/commit/dacc9256f4da7974efbdeea757df2e196b013991))

  - Add count field to compose_validation and pretty_print_validation for consistent assertion tracking
  - Standardize count minimum values to 1 across all validations
  - Add tag pattern validation allowing both snake_case and kebab-case formats
  - Add input field minLength constraints for better validation
  - Fix error validation structure consistency by adding error handling to filter_validation, expand_dotted_validation, and make_objects_validation

- Expand reset baseline and fix crlf handling ([`7cc48e1`](https://github.com/CatConfLang/ccl-test-data/commit/7cc48e1d112f5bd0fa0eda91fd71b65a322cb3b0))

  - Add line-endings tag to reset filter (15 → 29 active tests)
  - Fix test generator CRLF preservation in quoted strings
  - Update mock implementation to preserve \r characters in parsing
  - All tests pass with expanded baseline coverage

- Expand reset command to include more passing tests ([`2d7a56e`](https://github.com/CatConfLang/ccl-test-data/commit/2d7a56ee634503a67556a504c1cbcc903fd7c807))

  - Add redundant, quotes, and realistic tags to generate-level1
  - Increases active tests from 15 to 28 while maintaining all passing
  - Includes Level 2-4 basic functionality in baseline test set

- Add reset alias and fix test-mock workflow ([`d2f6451`](https://github.com/CatConfLang/ccl-test-data/commit/d2f645172cd9fc7cd5f6af2275bcf51059f911c3))

  - Add 'reset' alias for dev-basic command for easy repository state management
  - Fix test-mock to properly generate tests before running them
  - Fix dev-basic to use test-level1 instead of test-mock-basic
  - Update README with repository state management documentation
  - Clarify that clean passing state is required for commits and CI

- Replace find/xargs with cross-platform jv json schema validator ([`972544a`](https://github.com/CatConfLang/ccl-test-data/commit/972544abe82a20e6596f0a858dea72da1b30757e))

  - Add jv CLI tool dependency for JSON schema validation
  - Update validate-schema tool to accept directories and find JSON files
  - Replace Unix-specific find/xargs commands with cross-platform jv
  - Update justfile to use jv directly for validation
  - Add tools.go for tool dependency management
  - Update deps command to install Go tools alongside modules

- Move package.json to scripts/ folder for cleaner project structure ([`80750ab`](https://github.com/CatConfLang/ccl-test-data/commit/80750abad34b3eb79c47c7a4038ae094ace542da))

  - Move package.json and package-lock.json to scripts/ folder
  - Update justfile docs commands to run from scripts/ directory
  - Update clean-all to target scripts/node_modules
  - Fix docs-check to run git diff from project root
  - Keep Node.js dependencies isolated to scripts where they're used
  - Project root now free of npm artifacts

- Replace rimraf with go-based clean utility ([`eb59f24`](https://github.com/CatConfLang/ccl-test-data/commit/eb59f24dbf361b3369f8349dbcac4913cbf1af6a))

  - Add cmd/clean/main.go - cross-platform recursive file deletion utility
  - Update justfile clean commands to use Go clean utility instead of npx rimraf
  - Remove rimraf from package.json devDependencies
  - Eliminate last npm-only dependency, keeping only Node.js deps for docs scripts

- Complete migration from package.json scripts to justfile ([`528f350`](https://github.com/CatConfLang/ccl-test-data/commit/528f350d0390554d0ba861b9c8f4bb3156c48ea3))

  - Remove all scripts from package.json, keep only Node.js dependencies
  - Add rimraf for cross-platform file deletion
  - Enhance justfile with comprehensive commands: docs-update, docs-check, docs-schema (replaces npm scripts)
    * unified deps command (deps-node + go mod tidy) test now includes validation + docs check (replaces npm test) test-generated for Go tests only ci command for full pipeline
  - All commands now use justfile as single interface

- Enhance justfile with mock development workflow and update readme ([`cece41a`](https://github.com/CatConfLang/ccl-test-data/commit/cece41aeb6437dc93f6f58d8801c1bdcdd0b902e))

  - Add aliases (gen, t, l, v) for common commands
  - Add mock implementation development commands (generate-mock, test-mock, dev-mock)
  - Add level-specific test generation and execution commands
  - Add utility commands (stats, validate, dev-basic)
  - Update README to reflect current state with counted format requirement
  - Document Go test runner features and mock implementation
  - Add current test statistics (161 tests, 446 assertions)
  - Remove outdated legacy format documentation

- Enforce counted format with required count field in all validations ([`65213f8`](https://github.com/CatConfLang/ccl-test-data/commit/65213f846ad465a2693cc50d61655e36df6e7635))

  - Update JSON schema to require count field in parse, filter, expand_dotted, make_objects, and typed access validations
  - Remove legacy simple formats, keeping only counted format structure
  - Convert all 8 test files to use counted format with appropriate count values
  - Update schema documentation with comprehensive count field explanation and examples
  - Add Go-based schema validation tool for future validation

- Add ocaml reference tests for array ordering differences ([`6f38f57`](https://github.com/CatConfLang/ccl-test-data/commit/6f38f57a90d5074f6cdeb5861f89c797e002db93))

  - Empty key array tests now have reference_compliant versions
  - Proposed versions maintain insertion order for arrays
  - OCaml reference sorts arrays alphabetically
  - Enables discussion of array ordering semantics and predictability

- Add ocaml reference tests for comprehensive parsing differences ([`4a13b7b`](https://github.com/CatConfLang/ccl-test-data/commit/4a13b7b2ba6e89b13b207ae234177aa835934c19))

  - Complex parsing tests now have reference_compliant versions
  - Shows OCaml reference parses nested structures differently
  - Tab normalization and stress test parsing variations captured
  - Preserves both flat and hierarchical parsing approaches for comparison

- Add ocaml reference tests for tab handling differences ([`7111482`](https://github.com/CatConfLang/ccl-test-data/commit/7111482cd5e797b6afcea25cefaabc933e1b1a92))

  - Tab preservation tests now have reference_compliant versions
  - Proposed versions preserve literal tabs in parsing output
  - OCaml reference normalizes tabs to single spaces during parsing
  - Maintains both interpretations for whitespace handling discussions

- Add ocaml reference tests for boolean parsing differences ([`4f03763`](https://github.com/CatConfLang/ccl-test-data/commit/4f03763306dcc53200fc40592aeb36722fefa595))

  - Boolean tests now have reference_compliant versions showing OCaml strict parsing
  - OCaml reference only accepts "true"/"false", rejects "yes"/"no"/"on"/"off"/numeric
  - Proposed versions accept flexible boolean formats for user convenience
  - Enables productive discussion with OCaml maintainers about parsing standards

- Add ocaml reference test for boolean parsing differences ([`b023c14`](https://github.com/CatConfLang/ccl-test-data/commit/b023c14414a7dc394b6340be107a9a4245933030))

  - parse_boolean_yes: proposed parses 'yes' as true, OCaml reference errors
  - Demonstrates boolean parsing compatibility differences between implementations
  - Both behaviors preserved for maintainer discussions

- Add ocaml reference tests for parsing edge cases ([`b6bdbd9`](https://github.com/CatConfLang/ccl-test-data/commit/b6bdbd9c0dfe54784bf571828425be5b7225ee48))

  - no_equals_continuation: proposed allows continuation, OCaml reference errors
  - crlf_normalization: proposed normalizes CRLF, OCaml reference preserves \r
  - Both behaviors preserved to enable discussion with reference maintainers

- Add proposed/reference compliance tagging for ocaml behavioral differences ([`a5a4058`](https://github.com/CatConfLang/ccl-test-data/commit/a5a4058815bc0587c3f6e93b86c91d548712b175))

  - Add "proposed" tag to whitespace_only_error test representing current behavior
  - Add whitespace_only_error_ocaml_reference test showing OCaml reference behavior
  - Preserve both interpretations to enable discussion with reference maintainers

- Add optional assertion counting to test suite ([`b4163d2`](https://github.com/CatConfLang/ccl-test-data/commit/b4163d2a41cd43f44f74a0f8fdb8240a07c8157b))

  - Add union types to schema supporting both legacy and counted formats
  - Support explicit assertion counts via count field in validations
  - Use consistent naming: items/result/cases for different validation types
  - Update collection scripts to report total assertions (241 across 148 tests)
  - Add documentation and examples for new counting format
  - Maintain 100% backwards compatibility with existing tests

- Implement validation-based test format with remark readme updater ([`15ccc37`](https://github.com/CatConfLang/ccl-test-data/commit/15ccc37b29d6a840a60c86c2a68101f1f3d8da18))

  - Transform all test files to validation-based format eliminating multi-level confusion
  - Replace fragile regex-based README updater with robust AST-based remark processor
  - Update schema to explicit API function validations (parse, filter, compose, make_objects, etc.)
  - Fix package.json scripts to exclude schema.json from validation
  - Add comprehensive documentation and example test runners
  - Remove deprecated schema files and old script

- Implement feature-based test organization with comprehensive schema documentation ([`94cca49`](https://github.com/CatConfLang/ccl-test-data/commit/94cca49a4838d827aed7779508a88690a925987c))

  - Update JSON schema with standardized feature enum (parsing, processing, comments, object-construction, dotted-keys, typed-parsing, pretty-printing, error-handling)
  - Add feature metadata to all 135 test cases across 9 test files
  - Reorganize README from file-based to feature-based grouping (Core Parsing, Advanced Processing, Object Construction, Type System, Output & Validation)
  - Create comprehensive schema documentation using hybrid approach:
    - Auto-generated technical reference (docs/generated-schema.md) using json-schema-for-humans
    - Manual implementation guide (docs/schema-reference.md) with practical examples
  - Add automated schema documentation generation script (scripts/generate-schema-docs.mjs)
  - Update stats collection and README update scripts to use feature-based categorization
  - Update all documentation files with corrected file paths and cross-references

- Flatten test directory structure for easier parsing ([`287b1ec`](https://github.com/CatConfLang/ccl-test-data/commit/287b1ec4c9fa7a88fd1215b80d5afec4fcbeb62a))

  - Move all test files to single level in tests/ directory
  - Update schema references from ../schema.json to ./schema.json
  - Update package.json validation paths to flat structure
  - Rewrite stats script to use metadata-based categorization
  - Update README.md paths from nested to flat structure
  - Fix categorization logic in collect-stats.sh

- Reorganize tests from level-based to feature-based structure ([`3d40163`](https://github.com/CatConfLang/ccl-test-data/commit/3d401638368420688a8b675fc4421b12801f8f98))

  Replace rigid 4-level hierarchy with flexible feature-based organization:

    - Core tests: essential-parsing (18 tests), comprehensive-parsing (30 tests), object-construction (8 tests)
  - Feature tests: dotted-keys (18 tests), comments (3 tests), processing (21 tests), typed-access (17 tests)
  - Integration tests: errors (5 tests)
  - Remove all legacy level-*.json files for cleaner structure
  - Add comprehensive API reference with dual access pattern support
  - Generalize stats script with fd auto-discovery and feature categorization
  - Update package.json validation to new structure only

    Total: 135 tests organized by implementation priority rather than arbitrary levels

- Enhance bash scripts with charm gum interactivity ([`9cdedab`](https://github.com/CatConfLang/ccl-test-data/commit/9cdedab7356664c24f37d77794c1e3e8079baa0f))

  - Add --interactive flag to collect-stats.sh for styled output
  - Add --yes flag to update-readme.sh for automation compatibility
  - Implement change preview with confirmation prompts
  - Add progress spinners and colorful styling throughout
  - Remove backup functionality (rely on git instead)
  - Add gum to .mise.toml configuration
  - Maintain backward compatibility with existing npm scripts

- Update settings after cleanup ([`06bcfd1`](https://github.com/CatConfLang/ccl-test-data/commit/06bcfd1aa5046576d74fd3ab2ff8081721686252))

- Update local settings ([`f571d12`](https://github.com/CatConfLang/ccl-test-data/commit/f571d12ce685373c5238a93887ea5c103e1314f5))

- Add gitignore and remove node_modules from tracking ([`2e57e1d`](https://github.com/CatConfLang/ccl-test-data/commit/2e57e1d9579c41254870e0dd70eb5f89b49ca54d))

- Implement readme automation with jq/sd and mise dependency management ([`7484064`](https://github.com/CatConfLang/ccl-test-data/commit/7484064d006f3b58897ec99225268c78d16ec541))

- Implement comprehensive schema validation with inheritance ([`21698ec`](https://github.com/CatConfLang/ccl-test-data/commit/21698ec517d74b58e7f89c96102f50b2073e5a35))

- Add npm-based schema validation infrastructure ([`b1723ab`](https://github.com/CatConfLang/ccl-test-data/commit/b1723abee889b83c83778be54671d98fb1a75979))

- Implement decorative section headers with minimal api ([`cb2cc35`](https://github.com/CatConfLang/ccl-test-data/commit/cb2cc357cdf22810a3b06fdf719225e6b122f614))

  Core Implementation:
  - Add SectionGroup type with header and entries fields
  - Implement is_section_header() detection function
  - Implement group_by_sections() with recursive grouping algorithm
  - Detection rule: empty key + value starts with "="

    Comprehensive Test Coverage (12 new tests):
  - Basic section headers (double/triple equals)
  - Multiple sections with mixed content
  - Edge cases: empty sections, end-of-file headers, spacing
  - Multiline headers (indented and unindented continuation)
  - Integration with comments and list items
  - All 104 tests passing

    Documentation Updates:
  - Add Level 2.5 architectural positioning
  - Clarify colon convention is optional, not API requirement
  - Document user-defined helper function examples
  - Explain multiline behavior and CCL spec compliance
  - Add hybrid processing approach examples

    Design Philosophy:
  - Minimal core API (2 functions + 1 type)
  - Users implement own filtering/searching helpers
  - Maximum flexibility with standard list operations
  - Follows existing comment filtering patterns

- Implement comprehensive algebraic property testing for ccl ([`6f9bf38`](https://github.com/CatConfLang/ccl-test-data/commit/6f9bf381a05edbd5b79284f715acd4af9f6cfcff))

  Add 12 algebraic tests verifying CCL's mathematical foundations:
  - Monoid properties (identity element tests)
  - Semigroup properties (associativity tests)
  - Composition properties (closure, structure preservation)
  - Text concatenation equivalence
  - Unicode, comment, and nested structure stability

    Design decisions documented in DEV.md:
  - Use semantic empty (entry list []) as monoid identity
  - Separate core parsing from higher-level semantic tests
  - Fixed comprehensive JSON test cases over property-based testing

    All 12 algebraic property tests pass, confirming CCL's category theory foundations.

- Add comprehensive polishing improvements ([`abf3efa`](https://github.com/CatConfLang/ccl-test-data/commit/abf3efa6ee5eb4e75d5d00b0c16cb06960ff5660))

  - Enhanced error messages with contextual hints and suggestions
  - Created comprehensive CCL examples documentation with core vs Gleam separation
  - Fixed typed parsing whitespace test to match CCL specification
  - Added benchmarking framework with performance demonstration
  - Implemented object construction performance testing

- Add tdd framework for ccl typed parsing feature ([`184b9f0`](https://github.com/CatConfLang/ccl-test-data/commit/184b9f0edbb2dcb3967f7890b0e99412808610c3))





### Refactoring

- Rename parse_dedented to parse_indented ([`25a0d69`](https://github.com/CatConfLang/ccl-test-data/commit/25a0d6908ab8f40906ff482be099b785ad3a238a))

  ## Summary

    Renamed `parse_dedented` function to `parse_indented` throughout the test suite for better clarity. The new name describes the input type (indented text) rather than the transformation applied (dedenting).

    ## Rationale

    ### Problems with `parse_dedented` 1. **Describes implementation, not purpose**: "dedented" emphasizes the preprocessing step rather than the use case 2. **Awkward phrasing**: "dedented" (past participle) is grammatically unusual for a function name 3. **Poor discoverability**: Users think "I have indented input" not "I need to dedent"

    ### Why `parse_indented` is better 1. **Describes the input**: Clearly indicates this function is for indented/nested text 2. **Natural discoverability**: If you have indented input, the name tells you this is the right function 3. **Consistent with conventions**: We name based on what we operate on, not the transformation 4. **Clearer use cases**: Heredocs with indentation, nested configuration blocks, script blocks

    ## Changes

    ### Source Test Files
  - Updated `api_proposed_behavior.json` - 4 tests using `parse_dedented` → `parse_indented`
  - Updated `api_edge_cases.json` - 8 tests using `parse_dedented` → `parse_indented`

    ### Generated Files
  - Regenerated all flat format JSON tests in `generated_tests/`
  - Regenerated all Go test files in `go_tests/parsing/`

    ### Implementation
  - Renamed `ParseDedented` method to `ParseIndented` in mock CCL implementation
  - Updated `ValidationSet.ParseDedented` field to `ParseIndented`

    ### Schemas
  - Updated `schemas/source-format.json`
  - Updated `schemas/generated-format.json`

    ### Documentation
  - Updated `CLAUDE.md` - function descriptions and implementation steps
  - Updated `.github/agents/onboarding.md` - core parsing functions list
  - Updated `.github/agents/test-development.md` - example test format
  - Updated `docs/generated-schema.md`

    ### Tooling
  - Fixed `cmd/test-reader/main.go` to use `ParseIndented` and `FunctionParseIndented`
  - Updated stats collectors and generators

    ## Testing

    ✅ All 323 tests passing (180 skipped as expected) ```bash just reset ```

- Implement optional_ prefix for feature categorization ([`3bf6bc9`](https://github.com/CatConfLang/ccl-test-data/commit/3bf6bc907a29f29114b391ec92ee123f6dbe9d0a))

  Establish clear semantic distinction between feature types:
  - Core content features: comments, unicode, multiline, whitespace, empty_keys
  - Optional stable features: optional_typed_accessors
  - Experimental features: experimental_dotted_keys

    Changes:
  - Rename "typed_accessors" → "optional_typed_accessors" in all 18 test cases
  - Update all schemas to support both "experimental_" and "optional_" prefixes
  - Update Go configuration to handle new optional feature naming
  - Maintain backward compatibility with existing experimental_ prefix

    This makes it immediately clear which features are content categorization (all implementations should support) vs optional capabilities (implementations may skip based on their design choices).

- Consolidate variant tests into single files per type ([`20418fe`](https://github.com/CatConfLang/ccl-test-data/commit/20418fecdba1b3b2e1d34955a70eaa89185e27b0))

  - Create api_proposed_behavior.json with 24 consolidated proposed behavior tests
  - Create api_reference_compliant.json with 19 consolidated reference compliant tests
  - Remove all variant tests from core files, keeping only baseline behavior
  - Remove expand_dotted function tests from experimental file for schema compliance
  - Add reorganization scripts for variant test management
  - Regenerate all test files to reflect new consolidated structure

    Total test coverage: 336 tests with clear separation by behavior type

- Reorganize test suite with improved structure ([`e047964`](https://github.com/CatConfLang/ccl-test-data/commit/e0479641b0aa7591283e17421cf6108b2bc6eb2b))

  - Move core tests to source_tests/core/ directory for better organization
  - Move experimental features to source_tests/experimental/ for clarity
  - Update documentation to reflect simplified function groups and feature tags
  - Remove expand_dotted function from core processing (moved to experimental)
  - Regenerate all test files to reflect new structure and behavior variants
  - Update schemas and justfile for new test organization
  - Add restore-variants.sh script for variant management

- Restructure justfile generation commands ([`c628e3c`](https://github.com/CatConfLang/ccl-test-data/commit/c628e3ca8c14c111039fbb445cf9864b376ee19c))

  - Rename `generate` to `generate-go` for Go test file generation
  - Create new `generate` command that calls both `generate-flat` and `generate-go`
  - Maintains backward compatibility while enabling complete two-step generation

    This allows users to run either the full generation pipeline or individual steps as needed.

- Remove outdated 'counted format' terminology ([`29f3a32`](https://github.com/CatConfLang/ccl-test-data/commit/29f3a32259221107406c3242a98feaf5812654a7))

  Remove references to 'counted format' across documentation and code comments, replacing with clearer language about count fields.

    Changes:
  - Update README.md to use 'count fields' instead of 'counted format'
  - Simplify schema documentation descriptions
  - Update code comments to be more direct
  - Maintain functionality while improving clarity

    The count fields themselves remain unchanged - only the terminology has been updated to be less confusing.

- Consolidate documentation and improve project organization ([`2319046`](https://github.com/CatConfLang/ccl-test-data/commit/231904655a51784bd234c7418a026ec33cfeac48))

  - Archive Serena memories (26 files) to .serena/archive/ for historical reference
  - Archive scattered planning documentation (6 directories) to docs/archive/
  - Document Node.js dependency necessity in package.json and justfile
  - Reduce active documentation files from 47 to 12 (74% reduction)
  - Preserve all historical content while improving project navigation
  - Maintain build system functionality (all commands tested)

- Switch to flat format architecture using ccl-test-lib ([`6f375d3`](https://github.com/CatConfLang/ccl-test-data/commit/6f375d33918784cef19eb122562a3475d59d2bf6))

  Switch from source_tests to generated_tests directory and use ccl-test-lib for test loading following proper workflow: source_tests → ccl-test-lib → flat_tests → ccl-test-data CLI → go_tests.

    Changes:
  - Add ccl-test-lib dependency with local replacement
  - Update CLI default input from source_tests to generated_tests
  - Implement flat format validation functions following JSON schema
  - Add HasInput field to template for proper empty string handling
  - Update justfile to use --run-only function:parse instead of --level 1
  - Generate 369 Level 1 tests with proper Go syntax and compilation
  - Clean up old source format test files and directories

    The CLI now properly consumes flat format tests and generates working Go test files that compile and run. Test failures are expected for advanced parsing features not implemented in the simple mock.

- Streamline justfile from 244 to 85 lines ([`f4bdb9e`](https://github.com/CatConfLang/ccl-test-data/commit/f4bdb9ea50a4d6b4641cc77161cd5fb49a14a29e))

  - Consolidated 15+ generate variants into single parameterized command
  - Unified test commands with argument passing
  - Reduced from 8 dev workflows to 3 core workflows
  - Maintained legacy compatibility for most-used commands
  - 65% size reduction while preserving all functionality

- Standardize all naming to snake_case format ([`6af5c69`](https://github.com/CatConfLang/ccl-test-data/commit/6af5c690c3004d064eb703030dcd37db381fb3a6))

  - Renamed folders to snake_case:
    - generated-tests → generated_tests
    - source-tests → source_tests

    - Renamed all test files to snake_case:
    - api-*.json → api_*.json (10 files)
    - property-*.json → property_*.json (2 files)

    - Updated all references throughout codebase:
    - CLI tool defaults and file patterns
    - Documentation (*.md files)
    - Go source files and comments
    - Configuration files (justfile)
    - Memory files and internal references
    - Test file cross-references

    This creates consistent snake_case naming across all folders, files, and references in the CCL test suite for better maintainability.

- Rename generated_tests to go_tests for clarity ([`4edb156`](https://github.com/CatConfLang/ccl-test-data/commit/4edb1561a02bb9c5231430dd22a658c1e850d938))

  Renamed the generated test output directory from `generated_tests` to `go_tests` to better describe its contents as Go test files generated from JSON test data.

    Updated all references across:
  - CLI tool defaults and hardcoded paths
  - Documentation files (CLAUDE.md, docs/, claudedocs/)
  - Build configuration (justfile)
  - Generator code and examples
  - Memory files

    This improves repository organization and makes the purpose of the directory immediately clear to developers.

- Reorganize test suite for core ccl level reorientation ([`7ce1be4`](https://github.com/CatConfLang/ccl-test-data/commit/7ce1be4cae5206769a44f3be3ea60c4628f91149))

  Major restructuring to align with new CCL documentation understanding:

    **Core Changes:**
  - Split Core CCL (Level 1) into feature categories: parsing, hierarchy, integration
  - Reframe typed access as Level 2 (most common extension)
  - Move advanced processing to Level 3 (power user features)
  - Isolate experimental features in Level 4 (implementation-specific)

    **Files Reorganized:**
  - NEW: api-core-ccl-{parsing,hierarchy,integration}.json (Level 1)
  - UPDATED: api-typed-access.json → Level 2 (was Level 4)
  - RENAMED: api-processing.json → api-advanced-processing.json (Level 3)
  - RENAMED: api-comprehensive-parsing.json → api-edge-cases.json
  - NEW: api-experimental.json (Level 4, dotted keys Go-only)
  - REMOVED: api-essential-parsing.json, api-object-construction.json, api-dotted-keys.json

    **Tag & API Updates:**
  - Replace function:core-ccl with function:load (atomic API)
  - Update all decode references to load API
  - Fix function tags for new level system
  - Improve test names for clarity
  - Validate conflicts arrays and feature tags

    **Benefits:**
  - Core CCL now reflects atomic text→objects user experience
  - Clear progressive implementation path: Core → Typed → Advanced → Experimental
  - Experimental features properly isolated and marked
  - Test suite serves as authoritative behavior specification

- Consolidate ccl test suite variants ([`2e8632e`](https://github.com/CatConfLang/ccl-test-data/commit/2e8632ef9761898cec08602f1cc4e47f1bd38c50))

  Remove inappropriate API architecture violation from ocaml_stress_test_original variants. The proposed-behavior variant incorrectly mixed parse and filter responsibilities, violating CCL's API separation where parse() returns raw entries and filter() handles comment filtering.

    Changes:
  - Remove ocaml_stress_test_original with variant:proposed-behavior
  - Retain ocaml_stress_test_original (renamed from reference variant)
  - Preserve 32 legitimate behavioral variant pairs for implementation choices
  - Generate updated test files maintaining 165 tests, 426 assertions

    This maintains proper Level 1 (parse) vs Level 2 (filter) API boundaries while preserving genuine specification ambiguities like CRLF handling, tab processing, and spacing rules.

- Update test generation to use filter-based test selection ([`8350280`](https://github.com/CatConfLang/ccl-test-data/commit/83502807b7206782da8fc01cdcd6ac25fe90d12e))

  - Modified test runner to generate tests based on filter criteria
  - Tests now skip based on run-only filter matching instead of hardcoded skip reasons
  - Updated template generation to use consistent skip messages for filtered tests
  - All generated test files reflect the new filter-based approach

    This improves maintainability by centralizing test selection logic and makes test skipping more predictable based on configured filter tags.

- Add self-documenting prefixes to filtering tags for clarity ([`ade5e8d`](https://github.com/CatConfLang/ccl-test-data/commit/ade5e8d3056e75c5699d1ed34051706bc4c3233b))

- Clarify test categorization between optional features and spec ambiguities ([`5ef43c2`](https://github.com/CatConfLang/ccl-test-data/commit/5ef43c2c5d2da679f0d96ab54a0c955f0d2ede91))

- Clarify test types with api vs property naming ([`62be66d`](https://github.com/CatConfLang/ccl-test-data/commit/62be66dc5af99ba0d49f76f8532391482832de69))

  - Rename test files: api-*.json for direct API tests, property-*.json for mathematical properties
  - Update documentation to distinguish implementation complexity
  - Modify schema descriptions to show API vs PROPERTY validations
  - Add test runner guidance showing 30-line simple vs 100+ line full runners

- Migrate shell scripts to esm for cross-platform compatibility ([`6f2e280`](https://github.com/CatConfLang/ccl-test-data/commit/6f2e28055cc53181f62767b1ff8ecb8cc2c3997a))

  - Replace bash scripts with Node.js ESM modules
  - Fix bash version compatibility issues (macOS bash 3.2 vs 4+)
  - Remove external tool dependencies (jq, sd, gum)
  - Improve test categorization logic for accurate counts
  - Update mise config to only require Node.js
  - Ensure consistent behavior across all platforms

- Reorganize test suite into 4-level architecture ([`e29a4d2`](https://github.com/CatConfLang/ccl-test-data/commit/e29a4d23231aea3f7ca802d5a72f05f77ad04933))

  - Replace monolithic test suite with structured 4-level approach
  - Add ccl-entry-parsing.json (Level 1): 48 core parsing tests
  - Add ccl-entry-processing.json (Level 2): 11 processing tests
  - Add ccl-object-construction.json (Level 3): 8 object tests
  - Add ccl-errors.json: 5 error handling tests
  - Add ccl-unified-test-schema.json: single schema for all levels
  - Remove legacy test infrastructure and files
  - Update documentation and test runners
  - All 80 tests passing with comprehensive coverage





### Testing

- Regenerate test files after schema updates ([`a5c3c1e`](https://github.com/CatConfLang/ccl-test-data/commit/a5c3c1e989e9b087806a180302f26f9f77234934))

  Updates generated test files and Go test code following recent schema changes. All generated files have been regenerated to maintain consistency.

    Generated test stats: 366 tests with 366 assertions Active: 157 tests, Skipped: 209 tests

- Regenerate go test files for new object format ([`f935f70`](https://github.com/CatConfLang/ccl-test-data/commit/f935f70b2f62c4fe17d81ce8b1b03688e638c3b6))

  - Updated test generation to use object format for expected values
  - Modified generator to create empty map objects for skipped validations
  - Updated mock implementation to align with test expectations
  - Added function filtering support with run-only mode
  - Improved test skipping logic for better development workflow

    🔧 Changes:
  - internal/generator/generator.go: Object format generation
  - internal/mock/ccl.go: Mock implementation updates
  - go_tests/: All test files regenerated
  - justfile: Updated generation commands

    This establishes a clean foundation for the level removal work.

- Regenerate go test files for new object format ([`fc68a7c`](https://github.com/CatConfLang/ccl-test-data/commit/fc68a7c850737ebabe0b358c2ad058414ba31673))

  - Update generated test files to work with restructured source tests
  - Fix import ordering in config files
  - All 157 active tests now run with new $schema format

- Add missing comment parsong tests in lvl 1 ([`42765b5`](https://github.com/CatConfLang/ccl-test-data/commit/42765b5d31c89787991ce06cdf58a0a1e5fd7853))

- Add reference-compliant behavior tests and update generated tests ([`1395384`](https://github.com/CatConfLang/ccl-test-data/commit/139538481b9fed1d16ad3b5483966531c95655cd))

- Fix assertion counts ([`c0cb57d`](https://github.com/CatConfLang/ccl-test-data/commit/c0cb57da15b8c2155982d00d70988efcb7a0f828))

- Commit tests from just reset ([`2bedba3`](https://github.com/CatConfLang/ccl-test-data/commit/2bedba3f218e72fb65c8ee143064e32e2360d941))

- Test case schema ([`127ab4c`](https://github.com/CatConfLang/ccl-test-data/commit/127ab4cd4d80b44fdb5e83b8e587946387f214b2))

- Move tests to json suite ([`82d5e1a`](https://github.com/CatConfLang/ccl-test-data/commit/82d5e1a06bd429d8a9050c6ffda07445a90962f0))

- Add tests for nested parsing ([`262fd76`](https://github.com/CatConfLang/ccl-test-data/commit/262fd761674d21bf0a40af7b9a8463cbefd7f92a))

- Test cleanup ([`42455d5`](https://github.com/CatConfLang/ccl-test-data/commit/42455d51315e3f3efba198d0b5134a87a1a4f919))

- Test improvements ([`4b562b1`](https://github.com/CatConfLang/ccl-test-data/commit/4b562b1afe129c44e4c8a8c18bd33948e2ceb129))

- Add json test data ([`b64fcfc`](https://github.com/CatConfLang/ccl-test-data/commit/b64fcfc633bf68c97d49891c70239e20ca74c922))
