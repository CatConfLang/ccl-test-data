# CCL Test Suite Behavioral Patterns and Architecture

## CRLF Handling Variants

### Reference Behavior (`variant:reference_compliant`)
- **Processing**: Preserves CRLF during intermediate parsing steps
- **Parse Output**: `{"key": "key1", "value": "value1\r"}` (CRLF preserved)
- **Final Output**: Normalizes to LF in canonical format
- **Tag**: `behavior:crlf_preserve_literal`

### Proposed Behavior (`variant:proposed_behavior`) 
- **Processing**: Normalizes CRLF immediately during parsing
- **Parse Output**: `{"key": "key1", "value": "value1"}` (CRLF removed)
- **Final Output**: Already normalized for canonical format
- **Tag**: `behavior:crlf_normalize_to_lf`

## Canonical Format Types

### Compact Canonical Format
```
Input: "key1 = value1\r\nkey2 = value2\r\n"
Output: "key1 = value1\nkey2 = value2"
```
- Simple key=value preservation
- Used in basic formatting tests

### Indented Canonical Format
```
Input: "key1 = value1\r\nkey2 = value2\r\n"  
Output: "key1 =\n  value1 =\nkey2 =\n  value2 =\n"
```
- Hierarchical structure where values become nested entries
- Uses 2-space indentation
- Used in pretty-printing tests

## Dotted Key Feature Architecture

### Function-Based Tagging
- **`function:expand_dotted`**: 16 tests total
  - 7 core hierarchy tests (standard object construction)
  - 8 experimental tests (dotted key expansion)
  - 1 list access test (literal dotted key parsing)

### Feature-Based Tagging  
- **`feature:experimental_dotted_keys`**: 8 tests (Go implementation only)
  - Expands `database.host = value` into `{database: {host: "value"}}`
  - Subject to change, not standardized

### Standard vs Experimental Distinction
- **Standard**: Treats `database.host` as literal key string
- **Experimental**: Expands `database.host` into nested object hierarchy

## Test Validation Patterns

### Intermediate Validation Strategy
- Use `parse` validation to show processing differences
- Makes behavioral variants self-evident through observable outputs
- Example: CRLF preservation vs normalization in parse results

### Tag Representation Principles
- Feature tags should have >1 test to justify category existence
- Insufficient representation indicates consolidation opportunity
- Function tags can provide sufficient indication for single-test features

## Schema Organization

### Tag Categories
- **`function:*`**: Required CCL API functions (parse, get_string, etc.)
- **`feature:*`**: Optional language features (comments, unicode, etc.)
- **`behavior:*`**: Implementation choices (CRLF handling, spacing, etc.)
- **`variant:*`**: Specification variants (reference_compliant, proposed_behavior)

### Validation Hierarchy
- Function tags: Required capabilities
- Feature tags: Optional extensions
- Behavior tags: Processing choices (can be mutually exclusive)
- Variant tags: Specification compliance levels