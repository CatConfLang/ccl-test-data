# CCL Test Runner Justfile (Streamlined)

# Show available commands
default:
    @just --list

# Core aliases
alias t := test
alias gen := generate
alias flat := generate-flat
alias reset := dev-basic
alias view := view-tests
alias vs := view-tests-static

# === BUILD ===

# Build all binaries
build:
    go build -o bin/ccl-test-runner ./cmd/ccl-test-runner
    go build -o bin/test-reader ./cmd/test-reader

# Build individual binaries (used by other tasks)
build-test-runner:
    go build -o bin/ccl-test-runner ./cmd/ccl-test-runner

build-test-reader:
    go build -o bin/test-reader ./cmd/test-reader

# Install all tools to $GOPATH/bin
install:
    go install ./cmd/ccl-test-runner
    go install ./cmd/test-reader

# Install individual tools
install-test-runner:
    go install ./cmd/ccl-test-runner

install-test-reader:
    go install ./cmd/test-reader

# === ESSENTIAL WORKFLOWS ===

# Basic development: generate core tests and verify they pass (excludes known failing edge cases)
dev-basic:
    just clean
    just generate-flat
    just generate-go --run-only function:parse --skip-tags behavior:crlf_preserve_literal,behavior:tabs_preserve,behavior:tabs_to_spaces,behavior:strict_spacing
    just lint
    #!/usr/bin/env bash
    echo "🧪 Running tests..."
    echo "📋 Running basic tests (excluding known failing edge cases):"
    echo "  - TestKeyWithNewlineBeforeEqualsParse: newline within key portion before equals"
    echo "  - TestComplexMultiNewlineWhitespaceParse: complex whitespace with newlines in key"
    echo "  - TestDeeplyNestedListParse: nested structure parsing (expects flat entries)"
    echo "  - TestRoundTripWhitespaceNormalizationParse: whitespace handling inconsistencies"
    just _run-tests --basic-only

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

# Generate all: flat JSON files then Go test files
generate *ARGS="":
    just generate-flat {{ARGS}}
    just generate-go {{ARGS}}

# Generate Go test files from flat JSON files (with optional filtering)
generate-go *ARGS="":
    go run ./cmd/ccl-test-runner generate {{ARGS}}

# Generate flat JSON files from source JSON files (source-to-flat conversion)
# ARCHITECTURE NOTE: This delegates to ccl-test-lib via a thin CLI wrapper
# - Logic lives in: ccl-test-lib/generator.NewFlatGenerator()
# - CLI wrapper: cmd/ccl-test-runner/generate_flat.go
# - Separation: Library contains logic, CLI provides convenience interface
generate-flat *ARGS="":
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/core {{ARGS}}

# === TESTING ===

# Helper function for running tests with the test runner
_run-tests *ARGS="":
    go run ./cmd/ccl-test-runner test {{ARGS}}

# Run tests (with optional filtering)
test *ARGS="":
    #!/usr/bin/env bash
    if [[ "{{ARGS}}" == *"--full"* ]]; then
        just validate
        just docs-check
        just generate
        just _run-tests
    elif [[ "{{ARGS}}" == *"--all"* ]]; then
        # Run all tests including failing ones
        FILTERED_ARGS=$(echo "{{ARGS}}" | sed 's/--all//g' | sed 's/^ *//' | sed 's/ *$//')
        if [[ -z "$FILTERED_ARGS" ]]; then
            just _run-tests
        else
            just _run-tests $FILTERED_ARGS
        fi
    else
        # Default: run only passing tests (basic-only mode)
        if [[ -z "{{ARGS}}" ]]; then
            just _run-tests --basic-only
        else
            just _run-tests --basic-only {{ARGS}}
        fi
    fi

# === VALIDATION ===

validate:
    jv schemas/source-format.json source_tests/**/*.json
    jv schemas/generated-format.json generated_tests/*.json

# Update README.md with current test statistics using remark.js AST processing
docs-check:
    cd scripts && node update-readme-remark.mjs
    git diff --exit-code README.md

# === UTILITIES ===

stats:
    go run ./cmd/ccl-test-runner stats --input source_tests

list:
    go run ./cmd/ccl-test-runner test --list

# Interactive test viewer (TUI-based) - builds test-reader if needed
view-tests PATH="source_tests":
    just build-test-reader
    ./bin/test-reader {{PATH}}

# Static test viewer (CLI output) - builds test-reader if needed
view-tests-static PATH="source_tests":
    just build-test-reader
    ./bin/test-reader {{PATH}} --static

# View specific test file interactively
view-test FILE:
    just build-test-reader
    ./bin/test-reader {{FILE}}

# View specific test file with static output
view-test-static FILE:
    just build-test-reader
    ./bin/test-reader {{FILE}} --static

clean:
    go run ./cmd/clean go_tests bin
    rm -f bin/ccl-test-runner bin/test-reader

lint:
    go mod tidy
    go fmt ./...
    go vet ./...

deps:
    cd scripts && npm install
    go mod download
    go install github.com/santhosh-tekuri/jsonschema/cmd/jv@latest

# === RELEASE ===

# Show suggested next version based on conventional commits
release-check:
    git cliff --bumped-version

# Preview changelog for next release
release-preview:
    git cliff --unreleased

# Create release: updates CHANGELOG.md, commits, and tags
release version:
    git cliff --tag data-v{{version}} -o CHANGELOG.md
    git add CHANGELOG.md
    git commit -m "chore(release): data-v{{version}}"
    git tag data-v{{version}}
    @echo "Release data-v{{version}} created. Push with: git push origin main --tags"

# === CONVENIENCE COMMANDS ===

# Generate only basic tests for mock implementation
generate-mock:
    just generate --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior

# Test with verbose output
test-verbose:
    just test --verbose

# Run all tests including failing ones (overrides default basic-only mode)
test-all:
    just test --all

