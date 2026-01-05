# CCL Test Runner Justfile (Streamlined)

# Aliases
alias t := test
alias gen := generate
alias flat := generate-flat
alias reset := dev-basic
alias view := view-tests
alias vs := view-tests-static
alias pr := ci

# Show available commands
default:
    @just --list

# === BUILD ===

# Build: generate test files from source JSON
build:
    just generate-flat
    just generate-go --run-only function:parse --skip-tags behavior:crlf_preserve_literal,behavior:tabs_as_content,behavior:tabs_as_whitespace,behavior:toplevel_indent_preserve

# Build Go binaries
build-bin:
    go build -o bin/ccl-test-runner ./cmd/ccl-test-runner
    go build -o bin/test-reader ./cmd/test-reader

# Install tools to $GOPATH/bin
install:
    go install ./cmd/ccl-test-runner
    go install ./cmd/test-reader

# === ESSENTIAL WORKFLOWS ===

# Basic development: clean, build, lint, test
dev-basic:
    just clean
    just build
    just lint
    just test

# Full development: comprehensive test suite
dev:
    just clean
    just generate
    just test-all

# Production CI: complete validation pipeline
ci:
    just validate
    just build
    just lint
    just test
    just build-readme

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
# Uses behavior-metadata.json for function-specific filtering and auto-conflicts
generate-flat *ARGS="":
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/core --validate {{ARGS}}

# === TESTING ===

# Run tests
test *ARGS="":
    go run ./cmd/ccl-test-runner test --basic-only {{ARGS}}

# Run all tests including known failing ones
test-all *ARGS="":
    go run ./cmd/ccl-test-runner test {{ARGS}}

# === VALIDATION ===

validate:
    npx @sourcemeta/jsonschema validate schemas/source-format.json source_tests/
    npx @sourcemeta/jsonschema validate schemas/generated-format.json generated_tests/

# Update README.md with current test statistics using remark.js AST processing
build-readme:
    node scripts/update-readme-remark.mjs
    just _check-readme-unchanged

# Check if README.md has uncommitted changes
_check-readme-unchanged:
    git diff --exit-code README.md

# === UTILITIES ===

stats:
    go run ./cmd/ccl-test-runner stats --input source_tests

list:
    go run ./cmd/ccl-test-runner test --list

# Interactive test viewer (TUI-based) - builds test-reader if needed
view-tests PATH="source_tests/core":
    just build-bin
    ./bin/test-reader {{PATH}}

# Static test viewer (CLI output) - builds test-reader if needed
view-tests-static PATH="source_tests/core":
    just build-bin
    ./bin/test-reader {{PATH}} --static

# View specific test file interactively
view-test FILE:
    just build-bin
    ./bin/test-reader {{FILE}}

# View specific test file with static output
view-test-static FILE:
    just build-bin
    ./bin/test-reader {{FILE}} --static

clean:
    go run ./cmd/clean go_tests bin
    rm -f bin/ccl-test-runner bin/test-reader

lint:
    go mod tidy
    go fmt ./...
    go vet ./...

format:
    go fmt ./...

deps:
    npm install
    go mod download

# Configure go.mod for CI (removes local replace directive)
deps-ci:
    go mod edit -dropreplace=github.com/CatConfLang/ccl-test-lib
    go mod tidy

# === RELEASE ===

# Show suggested next version based on conventional commits
release-check:
    git cliff --bumped-version

# Preview changelog for next release
release-preview:
    git cliff --unreleased

# Create release: updates CHANGELOG.md, commits, and tags
release version:
    git cliff --tag v{{version}} -o CHANGELOG.md
    git add CHANGELOG.md
    git commit -m "chore(release): v{{version}}"
    git tag v{{version}}
    @echo "Release v{{version}} created. Push with: git push origin main --tags"

# === CONVENIENCE COMMANDS ===

# Generate only basic tests for mock implementation
generate-mock:
    just generate --skip-tags multiline,error,flexible-boolean-parsing,crlf-normalization,proposed-behavior

# Test with verbose output
test-verbose:
    just test --verbose
