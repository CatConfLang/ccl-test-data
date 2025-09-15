# CCL Test Runner Justfile

# Show available commands
default:
    @just --list

# Common aliases
alias gen := generate
alias t := test
alias l := list
alias v := test-verbose
alias reset := dev-basic
alias reader := read-dir
alias tui := read-dir-tui

# Build the ccl-test-runner binary
build:
    go build -o bin/ccl-test-runner ./cmd/ccl-test-runner

# Build the test-reader CLI
build-reader:
    go build -o bin/test-reader ./cmd/test-reader

# Install ccl-test-runner to $GOPATH/bin
install:
    go install ./cmd/ccl-test-runner

# Generate Go test files from JSON test data
generate:
    go run ./cmd/ccl-test-runner generate --input tests

# Generate implementation-friendly flat tests from source format
generate-flat:
    go run ./cmd/ccl-test-runner generate-flat

# Generate tests optimized for mock implementation (skip advanced features)
generate-mock:
    go run ./cmd/ccl-test-runner generate --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior

# Generate tests for basic Level 1 functionality only
generate-level1:
    go run ./cmd/ccl-test-runner generate --run-only basic,essential-parsing,empty,redundant,quotes,realistic,line-endings

# Generate tests for mock implementation development (Level 1 + comments)
generate-mock-dev:
    go run ./cmd/ccl-test-runner generate --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior,reference-compliant-behavior --run-only basic,essential-parsing,empty,comments

# Run comprehensive test suite including validation and docs (replaces npm test)
test:
    just validate
    just docs-check
    just generate
    just test-generated

# Run generated Go tests only
test-generated:
    go run ./cmd/ccl-test-runner test

# Run tests suitable for mock implementation
test-mock:
    just generate-mock
    go run ./cmd/ccl-test-runner test

# Run only basic tests for mock development
test-mock-basic:
    just generate-mock-dev
    go run ./cmd/ccl-test-runner test

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

# Show comprehensive test suite statistics
stats:
    go run ./cmd/ccl-test-runner stats --input tests

# Show test statistics in JSON format
stats-json:
    go run ./cmd/ccl-test-runner stats --input tests --format json

# Build and run test-reader with directory selection (CLI mode)
read-dir:
    just build-reader
    ./bin/test-reader tests/

# Build and run test-reader with directory selection (TUI mode)
read-dir-tui:
    just build-reader
    ./bin/test-reader tests/ --tui

# Run test-reader on a specific file (pass filename as argument)
read FILE:
    just build-reader
    ./bin/test-reader {{FILE}}

# Validate all test files against schema
validate:
    jv tests/schema.json tests/api-*.json tests/property-*.json

# Validate generated flat format tests against generated schema  
validate-flat:
    jv schemas/generated-format.json generated-tests/*.json

# Validate both source and generated formats
validate-all:
    just validate
    just validate-flat

# Validate enhanced LLM metadata in test files
validate-metadata:
    @echo "🤖 Validating enhanced LLM metadata in test files..."
    @go run scripts/validate-enhanced-metadata.go "tests/api-*.json"

# Clean up generated files (cross-platform)
clean:
    go run ./cmd/clean go_tests bin

# Clean everything including node_modules (cross-platform)
clean-all:
    go run ./cmd/clean go_tests bin scripts/node_modules

# Update documentation with current stats (replaces npm run docs:update)
docs-update:
    cd scripts && node update-readme-remark.mjs

# Check if documentation is up to date (replaces npm run docs:check)
docs-check:
    cd scripts && node update-readme-remark.mjs
    git diff --exit-code README.md

# Generate schema documentation (replaces npm run docs:schema)
docs-schema:
    cd scripts && node generate-schema-docs.mjs

# Run go mod tidy and format code
check:
    go mod tidy
    go fmt ./...
    
# Lint and format all code
lint:
    go mod tidy
    go fmt ./...
    go vet ./...

# Install Node.js dependencies (in scripts folder)
deps-node:
    cd scripts && npm install

# Install all dependencies (Node.js + Go modules + tools)
deps: deps-node
    go mod download
    go install github.com/santhosh-tekuri/jsonschema/cmd/jv

# Full development cycle: clean, generate, and test
dev:
    just clean
    just generate
    just test-generated

# Mock development cycle: clean, generate mock tests, and run them
dev-mock:
    just clean
    just generate-mock
    just test-mock

# Full development workflow with docs
dev-full:
    just clean
    just generate
    just test

# Run performance benchmarks
benchmark:
    go run ./cmd/ccl-test-runner benchmark

# Run benchmarks with comparison to historical results
benchmark-compare HISTORICAL_FILE:
    go run ./cmd/ccl-test-runner benchmark --compare {{HISTORICAL_FILE}}

# Run benchmarks and save as baseline
benchmark-baseline:
    mkdir -p benchmarks
    go run ./cmd/ccl-test-runner benchmark --results benchmarks/baseline.json

# CI/CD pipeline with benchmarking
ci:
    just validate
    just generate
    just test-generated
    just docs-check
    just benchmark-baseline

# Enhanced CI with regression detection
ci-benchmark BASELINE:
    just validate
    just generate
    just test-generated
    just docs-check
    just benchmark-compare {{BASELINE}}

# Quick development cycle for basic functionality
dev-basic:
    just clean
    just generate-level1
    just lint
    just test-level1
