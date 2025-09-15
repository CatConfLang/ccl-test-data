# CCL Test Runner Justfile (Streamlined)

# Show available commands
default:
    @just --list

# Core aliases
alias t := test
alias gen := generate
alias flat := generate-flat
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
    just generate --run-only function:parse
    just lint
    just test --levels 1

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

# Generate Go test files from flat JSON files (with optional filtering)
generate *ARGS="":
    go run ./cmd/ccl-test-runner generate {{ARGS}}

# Generate flat JSON files from source JSON files (source-to-flat conversion)
# ARCHITECTURE NOTE: This delegates to ccl-test-lib via a thin CLI wrapper
# - Logic lives in: ccl-test-lib/generator.NewFlatGenerator()
# - CLI wrapper: cmd/ccl-test-runner/generate_flat.go
# - Separation: Library contains logic, CLI provides convenience interface
generate-flat *ARGS="":
    go run ./cmd/ccl-test-runner generate-flat {{ARGS}}

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
    go run ./cmd/ccl-test-runner stats --input source_tests

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