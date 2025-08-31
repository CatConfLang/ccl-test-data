# CCL Parser Conformance Test Suite

This directory contains a language-agnostic test suite for CCL (Categorical Configuration Language) parsers.

## Files

- `ccl-test-suite.json` - The complete test suite with all test cases
- `ccl-test-suite-schema.json` - JSON Schema for validating test suite format
- `CCL-TEST-SUITE.md` - This documentation file

## Test Suite Structure

The test suite is organized as a JSON file with the following structure:

```json
{
  "test_suite": "CCL Parser Conformance Tests",
  "version": "1.0.0",
  "tests": [
    {
      "name": "unique_test_name",
      "description": "Human readable description",
      "input": "CCL input string (use \\n for newlines)",
      "expected": [
        {"key": "expected_key", "value": "expected_value"}
      ],
      "tags": ["category", "feature"]
    }
  ],
  "error_tests": [
    {
      "name": "error_test_name",
      "description": "Description of error condition",
      "input": "Invalid CCL input",
      "expected_error": true,
      "error_message": "Expected error message",
      "tags": ["error", "category"]
    }
  ]
}
```

## Test Categories (Tags)

### Basic Features
- `basic` - Simple key-value pairs
- `whitespace` - Whitespace handling and trimming
- `multiple` - Multiple key-value pairs

### Advanced Features
- `multiline` - Multi-line values
- `continuation` - Line continuation logic
- `indentation` - Indentation-based parsing
- `nested` - Nested configurations

### Data Types and Content
- `unicode` - Unicode character handling
- `equals` - Equals signs in values
- `empty-key` - Empty key handling
- `empty-value` - Empty value handling
- `lists` - List representations

### Edge Cases
- `edge-cases` - Various edge cases
- `line-endings` - Line ending normalization
- `eof` - End-of-file handling
- `error` - Error conditions

### Source
- `gleam` - Test cases from Gleam implementation
- `ocaml` - Test cases from OCaml reference implementation

## Usage in Different Languages

### General Approach
1. Parse the JSON test suite file
2. For each test case:
   - Parse the `input` string with your CCL parser
   - Compare the result with the `expected` array
   - Verify key-value pairs match exactly
3. For error tests:
   - Verify the input causes a parse error
   - Optionally check error message matches

### Example Pseudocode

```python
import json

def run_test_suite(ccl_parser):
    with open('ccl-test-suite.json') as f:
        suite = json.load(f)
    
    passed = 0
    failed = 0
    
    for test in suite['tests']:
        try:
            result = ccl_parser.parse(test['input'])
            expected = [(entry['key'], entry['value']) for entry in test['expected']]
            
            if result == expected:
                print(f"✓ {test['name']}")
                passed += 1
            else:
                print(f"✗ {test['name']}: Expected {expected}, got {result}")
                failed += 1
        except Exception as e:
            print(f"✗ {test['name']}: Unexpected error: {e}")
            failed += 1
    
    print(f"Results: {passed} passed, {failed} failed")
```

## Test Coverage

The test suite includes **42 test cases** covering:

- Basic key-value parsing
- Whitespace and trimming rules
- Multiline values and continuation
- Comment extension
- Composition stability
- Unicode handling
- Line ending normalization
- Edge cases and error conditions
- List representations (empty keys/values)

Test cases are derived from both the Gleam implementation and the OCaml reference implementation to ensure comprehensive coverage.

## Validation

You can validate the test suite format using the provided JSON Schema:

```bash
# Using ajv-cli
ajv validate -s ccl-test-suite-schema.json -d ccl-test-suite.json

# Using python jsonschema
python -c "
import json
import jsonschema
with open('ccl-test-suite-schema.json') as f: schema = json.load(f)
with open('ccl-test-suite.json') as f: data = json.load(f)
jsonschema.validate(data, schema)
print('Test suite is valid!')
"
```

## Contributing

To add new test cases:

1. Follow the JSON structure defined in the schema
2. Add appropriate tags for categorization
3. Ensure test names are unique
4. Include clear descriptions
5. Test both success and error conditions where relevant

## Version History

- **v1.0.0** - Initial release combining Gleam and OCaml test cases