# Changelog

All notable changes to the CCL test data will be documented in this file.

## [0.1.0] - 2025-12-31
### Tests
#### Bug Fixes
- Remove incorrect canonical_format tests from proposed_behavior (#22)
- Correct filter function expectations to remove comments (#20)
- Correct tabs_to_spaces behavior to use single space (#28)
- Correct tabs_to_spaces behavior to use single space (#29)

#### Features
- Add bare list indentation tests (#25)
- Add comprehensive behavior tests (#26)

#### Refactor
- Simplify whitespace behavior model (#34)


### Schema
#### Features
- Rename tab/spacing behaviors for clarity (#39)


## [0.0.1] - 2025-12-09
### Tests
#### Bug Fixes
- Add list_coercion_enabled behavior to api_list_access tests (#9)
- Add conflict metadata for tabs_preserve behavior (#15)
- Add list_coercion_enabled behavior to proposed_behavior tests (#17)

#### Features
- Add bare list auto-unwrapping tests for get_list (#8)
- Add reference_compliant test variants with proper conflicts
- Add algebraic property functions and multi-input support
- Unify tests to use inputs field
- Add indented line continuation test cases (#18)


### Schema
#### Features
- Standardize on canonical_format function name
- Implement conditional args field requirements
- Update JSON format to object-based structure with proper $schema usage
- Add array_order_insertion/lexicographic behavior group


### Test Reader
#### Bug Fixes
- Integrate ccl-test-lib for proper type handling


### Cli
#### Documentation
- Document delegation architecture for flat generation



