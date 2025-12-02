# Changelog

All notable changes to the CCL test data will be documented in this file.

## [v0.0.1] - 2025-12-02

### Bug Fixes

- Correct parsing precedence - integers win overlapping cases
- Resolve all parser test failures and implement complete test suite
- Consolidate level-4 typed parsing tests into unified schema
- Resolve npm audit vulnerabilities with fast-json-patch override
- Allow error cases in typed access validation counted format
- Prevent unused imports in generated tests
- Disallow count:0 in schema and fix assertion counting
- Ensure deterministic test generation by sorting map keys
- Resolve missing conflicts declarations for mutually exclusive tests
- Resolve test data inconsistencies and validation issues
- Resolve JSON schema validation errors in test files
- Enable boolean true/false tests in lenient mode
- Update JSON schema to match recent CCL API changes
- Harmonize function tags to use underscores for schema consistency
- Complete underscore naming harmonization across test suite
- Eliminate remaining kebab-case function names in metadata
- Resolve go vet issues in generated test files
- Omit empty conflicts field from flat JSON files
- Resolve format mismatch between ccl-test-lib and ccl-test-data
- Resolve linting issues and code quality improvements
- **generation**: Resolve schema validation and flat format issues
- **generation**: Standardize null arrays to empty arrays
- **build**: Resolve compilation errors in type references
- Resolve mock CCL implementation test failures
- Resolve test generation issues in ccl-test-data
- **test-reader**: Integrate ccl-test-lib for proper type handling
- Test-reader updates
- **tests**: Add list_coercion_enabled behavior to api_list_access tests (#9)
- Filter behavior tags to only apply to relevant functions

### Documentation

- Spec 1.1.0
- Update test counts to reflect actual test suite size
- Establish CCL central documentation hub
- Refer to main docs
- Remove temporal language and version references
- Update schema to document new feature categories and tag conventions
- Add comprehensive test filtering design pattern documentation
- Update CLAUDE.md with comprehensive development guide
- Update documentation for feature-based tagging system
- Update documentation for new function:parse-value tag
- Add flexible error message testing guidelines
- Add comprehensive documentation suite with enhanced README
- Update to 4-level system and add cross-references to ccl.tylerbutler.com
- Update architecture from 5-level to 4-level system
- Update documentation and settings for current architecture
- **cli**: Document delegation architecture for flat generation
- Add CLI generator approach plan for test filtering
- Replace level concept with function groups and enhance technical documentation
- Comprehensive documentation cleanup and reorganization
- Remove level references from comments and documentation
- Clean up remaining level references across project
- Optimize documentation for efficiency and readability
- Add comprehensive Copilot agent onboarding instructions (#3)
- Standardize snake_case naming and update test counts (#7)
- Add missing behaviors and link to behavior reference

### Features

- Add TDD framework for CCL typed parsing feature
- Add comprehensive polishing improvements
- Implement comprehensive algebraic property testing for CCL
- Implement decorative section headers with minimal API
- Add npm-based schema validation infrastructure
- Implement comprehensive schema validation with inheritance
- Implement README automation with jq/sd and mise dependency management
- Add gitignore and remove node_modules from tracking
- Update local settings
- Update settings after cleanup
- Enhance bash scripts with Charm Gum interactivity
- Reorganize tests from level-based to feature-based structure
- Flatten test directory structure for easier parsing
- Implement feature-based test organization with comprehensive schema documentation
- Implement validation-based test format with remark README updater
- Add optional assertion counting to test suite
- Add proposed/reference compliance tagging for OCaml behavioral differences
- Add OCaml reference tests for parsing edge cases
- Add OCaml reference test for boolean parsing differences
- Add OCaml reference tests for boolean parsing differences
- Add OCaml reference tests for tab handling differences
- Add OCaml reference tests for comprehensive parsing differences
- Add OCaml reference tests for array ordering differences
- Enforce counted format with required count field in all validations
- Enhance justfile with mock development workflow and update README
- Complete migration from package.json scripts to justfile
- Replace rimraf with Go-based clean utility
- Move package.json to scripts/ folder for cleaner project structure
- Replace find/xargs with cross-platform jv JSON schema validator
- Add reset alias and fix test-mock workflow
- Expand reset command to include more passing tests
- Expand reset baseline and fix CRLF handling
- Improve JSON schema validation constraints
- Implement feature-based tagging system with enhanced statistics
- Add test-reader CLI with enhanced styling and TUI support
- Add directory support with file selection to test-reader CLI
- Enhance test-reader UI with improved entry display and navigation
- Add scrollable entry panes for improved large entry list handling
- Add performance benchmarking with object pooling and enhanced error reporting
- Implement comprehensive LLM optimization with enhanced metadata v2.1
- Clarify CRLF behavioral variants and consolidate dotted key features
- Implement separate fields schema for test metadata
- Replace schemas with current format validation
- Fix CCL test suite failures and enhance architecture
- **config**: Centralize behavioral choices with validation
- **config**: Centralize behavioral choices with validation
- **schema**: Standardize on canonical_format function name
- **schema**: Implement conditional args field requirements
- **schema**: Update JSON format to object-based structure with proper $schema usage
- Add Claude Code workflows (#1)
- Add simplified YAML configuration system for CCL test runner
- Restructure source tests to support $schema fields
- Remove level property from JSON schemas and test data
- Remove level-related infrastructure from stats and CLI systems
- Add test exclusion to just reset for progressive implementation
- Refactor justfile with reusable test runner invocation
- Add typed_accessors feature for typed accessor function tests
- Refactor test generation with feature categorization and remove invalid associativity tests
- Rename parse_value to parse_dedented (#5)
- **tests**: Add bare list auto-unwrapping tests for get_list (#8)
- **tests**: Add reference_compliant test variants with proper conflicts
- **schema**: Add array_order_insertion/lexicographic behavior group
- Add array_order behaviors to schema and whitespace behavior tests
- Add release workflow with git-cliff changelog generation

### Refactor

- Reorganize test suite into 4-level architecture
- Migrate shell scripts to ESM for cross-platform compatibility
- Clarify test types with api vs property naming
- Clarify test categorization between optional features and spec ambiguities
- Add self-documenting prefixes to filtering tags for clarity
- Update test generation to use filter-based test selection
- Consolidate CCL test suite variants
- Reorganize test suite for Core CCL level reorientation
- Rename generated_tests to go_tests for clarity
- Standardize all naming to snake_case format
- Streamline justfile from 244 to 85 lines
- Switch to flat format architecture using ccl-test-lib
- Consolidate documentation and improve project organization
- Remove outdated 'counted format' terminology (#2)
- Restructure justfile generation commands
- Reorganize test suite with improved structure
- Consolidate variant tests into single files per type
- Implement optional_ prefix for feature categorization
- Rename parse_dedented to parse_indented (#6)

### Testing

- Add json test data
- Test improvements
- Test cleanup
- Add tests for nested parsing
- Move tests to JSON suite
- Test case schema
- Commit tests from just reset
- Fix assertion counts
- Add reference-compliant behavior tests and update generated tests
- Add missing comment parsong tests in lvl 1
- Regenerate Go test files for new object format
- Regenerate Go test files for new object format
- Regenerate test files after schema updates

### Build

- Justfile

### Chore

- Update Claude settings configuration
- Clean up unused Go dependencies
- Format
- Update .gitignore for Go build artifacts and local settings
- Add git diff to auto-approved commands
- Clean up temporary files and redundant documentation
- Update CLAUDE.md
- Code cleanup
- **release**: Data-vdata-v0.0.1


