# CCL Test Runner Justfile (Streamlined)

# Show available commands
default:
    @just --list

# Core aliases
alias t := test
alias gen := generate
alias reset := dev-basic

# === BUILD ===
build:
    go build -o bin/ccl-test-runner ./cmd/ccl-test-runner

install:
    go install ./cmd/ccl-test-runner

# === ESSENTIAL WORKFLOWS ===

# Basic development: generate Level 1 tests and verify they pass
dev-basic:
    just clean
    just generate --level 1
    just lint
    just test --level 1

# Full development: comprehensive test suite
dev:
    just clean
    just generate
    just test

# Production CI: complete validation pipeline
ci:
    just validate
    just generate
    just test
    just docs-check

# === GENERATION ===

# Generate tests (with optional filtering)
generate *ARGS="":
    go run ./cmd/ccl-test-runner generate {{ARGS}}

# === TESTING ===

# Run tests (with optional filtering)  
test *ARGS="":
    #!/usr/bin/env bash
    if [[ "{{ARGS}}" == *"--full"* ]]; then
        just validate
        just docs-check
        just generate
        go run ./cmd/ccl-test-runner test
    else
        go run ./cmd/ccl-test-runner test {{ARGS}}
    fi

# === VALIDATION ===

validate:
    jv schemas/source-format.json source_tests/api_*.json source_tests/property_*.json

docs-check:
    cd scripts && node update-readme-remark.mjs
    git diff --exit-code README.md

# === UTILITIES ===

stats:
    go run ./cmd/ccl-test-runner stats --input tests

list:
    go run ./cmd/ccl-test-runner test --list

clean:
    go run ./cmd/clean go_tests bin

lint:
    go mod tidy
    go fmt ./...
    go vet ./...

deps:
    cd scripts && npm install
    go mod download
    go install github.com/santhosh-tekuri/jsonschema/cmd/jv

# === LEGACY COMMANDS (for compatibility) ===

# Generate only basic tests for mock implementation
generate-mock:
    just generate --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior

# Test with verbose output
test-verbose:
    just test --verbose

# Test specific levels
test-level1:
    just test --levels 1

test-level2:
    just test --levels 2

test-level3:
    just test --levels 3

test-level4:
    just test --levels 4