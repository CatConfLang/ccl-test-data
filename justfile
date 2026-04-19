# CCL Test Data - JSON test suite for CCL implementations

# === ALIASES ===
alias b := build
alias t := test
alias f := format
alias l := lint
alias c := clean

# Default recipe
default:
    @just --list

# === STANDARD RECIPES ===

# Generate test files from source JSON
build:
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/core --validate
    go run ./cmd/ccl-test-runner generate --run-only function:parse --skip-tags behavior:crlf_preserve_literal,behavior:indent_tabs,behavior:delimiter_prefer_spaced

# Run tests
test:
    go run ./cmd/ccl-test-runner test --basic-only

# Format Go code
format:
    go fmt ./...

# Run linter
lint:
    go mod tidy
    go vet ./...

# Remove build artifacts
clean:
    go run ./cmd/clean go_tests bin
    rm -f bin/ccl-test-runner bin/test-reader

# Full validation workflow
ci: format lint test build
    just validate

alias pr := ci

# === CI ===

# CI-specific dependency setup (drops local replace directives)
deps-ci:
    go mod edit -dropreplace=github.com/CatConfLang/ccl-test-lib
    go mod tidy

# Build README and verify no uncommitted changes
build-readme:
    node scripts/update-readme-remark.mjs
    just _check-readme-unchanged

# Check if README.md has uncommitted changes
_check-readme-unchanged:
    #!/usr/bin/env bash
    if ! git diff --quiet HEAD -- README.md; then
        echo "ERROR: README.md has uncommitted changes. Run 'just build-readme' locally and commit."
        exit 1
    fi

# === UTILITIES ===

# Install dependencies
deps:
    npm install
    go mod download
    uv tool install python-semantic-release

# Validate JSON schema compliance
validate:
    npx @sourcemeta/jsonschema validate schemas/source-format.json source_tests/
    npx @sourcemeta/jsonschema validate schemas/generated-format.json generated_tests/

# Validate that every tag used in source tests has a canonical documentation URL
validate-tags:
    go run ./cmd/validate-tags

# Same, but skip the network fetch and use docs/tag-index.json directly
validate-tags-offline:
    go run ./cmd/validate-tags --offline

# Show test statistics
stats:
    go run ./cmd/ccl-test-runner stats --input source_tests

# Clean, build, lint, and test (ensures clean state)
reset:
    just clean
    just build
    just lint
    just test

# Run all tests including known failing ones
test-all:
    go run ./cmd/ccl-test-runner test

# === GENERATION ===

# Generate all test files (flat JSON + Go tests)
generate:
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/core --validate
    go run ./cmd/ccl-test-runner generate

# Generate flat JSON files from source
generate-flat:
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/core --validate

# Generate Go test files from flat JSON
generate-go:
    go run ./cmd/ccl-test-runner generate

# Install CLI tools to $GOPATH/bin
install:
    go install ./cmd/ccl-test-runner
    go install ./cmd/test-reader

# === RELEASE ===

# Show suggested next version based on conventional commits
release-check:
    semantic-release version --print

# Preview changelog for next release (dry-run)
release-preview:
    semantic-release --noop version

# Generate changelog without version bump
release-changelog:
    semantic-release changelog

# Create release: updates CHANGELOG.md, commits, and tags
release:
    semantic-release version --no-push
    @echo "Release created. Push with: git push origin main --tags"
