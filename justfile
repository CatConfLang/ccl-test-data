# CCL Test Runner Justfile

# Show available commands
default:
    @just --list

# Common aliases
alias gen := generate
alias t := test
alias l := list
alias v := test-verbose

# Build the ccl-test-runner binary
build:
    go build -o bin/ccl-test-runner ./cmd/ccl-test-runner

# Install ccl-test-runner to $GOPATH/bin
install:
    go install ./cmd/ccl-test-runner

# Generate Go test files from JSON test data
generate:
    go run ./cmd/ccl-test-runner generate

# Generate tests optimized for mock implementation (skip advanced features)
generate-mock:
    go run ./cmd/ccl-test-runner generate --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior

# Generate tests for basic Level 1 functionality only
generate-level1:
    go run ./cmd/ccl-test-runner generate --run-only basic,essential-parsing,empty

# Generate tests for mock implementation development (Level 1 + comments)
generate-mock-dev:
    go run ./cmd/ccl-test-runner generate --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior,reference-compliant-behavior --run-only basic,essential-parsing,empty,comments

# Run all generated tests
test:
    go run ./cmd/ccl-test-runner test

# Run tests suitable for mock implementation
test-mock:
    go run ./cmd/ccl-test-runner test --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior

# Run only basic tests for mock development
test-mock-basic:
    go run ./cmd/ccl-test-runner test --run-only basic,essential-parsing,empty,comments

# Run only level 1 tests
test-level1:
    go run ./cmd/ccl-test-runner test --levels 1

# Run only level 2 tests  
test-level2:
    go run ./cmd/ccl-test-runner test --levels 2

# Run only level 3 tests
test-level3:
    go run ./cmd/ccl-test-runner test --levels 3

# Run only level 4 tests
test-level4:
    go run ./cmd/ccl-test-runner test --levels 4

# Run only comment-related tests
test-comments:
    go run ./cmd/ccl-test-runner test --features comments

# Run only parsing tests
test-parsing:
    go run ./cmd/ccl-test-runner test --features parsing

# Run only object construction tests
test-objects:
    go run ./cmd/ccl-test-runner test --features object

# Run all tests with verbose output
test-verbose:
    go run ./cmd/ccl-test-runner test --verbose

# Run all tests with table format
test-table:
    go run ./cmd/ccl-test-runner test --format table

# List all available test packages
list:
    go run ./cmd/ccl-test-runner test --list

# Show test generation statistics
stats:
    go run ./cmd/ccl-test-runner generate 2>&1 | tail -5

# Validate all test files against schema
validate:
    find tests/ -name "api-*.json" -not -name "schema.json" | xargs go run cmd/validate-schema/main.go

# Clean up generated files
clean:
    rm -rf generated_tests/ bin/

# Run go mod tidy and format code
check:
    go mod tidy
    go fmt ./...

# Install dependencies
deps:
    go get gotest.tools/gotestsum
    go get github.com/xeipuuv/gojsonschema

# Full development cycle: clean, generate, and test
dev:
    just clean
    just generate
    just test

# Mock development cycle: clean, generate mock tests, and run them
dev-mock:
    just clean
    just generate-mock
    just test-mock

# Quick development cycle for basic functionality
dev-basic:
    just clean
    just generate-level1
    just test-mock-basic