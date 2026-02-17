# Fuzz Test Generator Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build a seeded randomized test generator that produces source-format JSON test cases covering special character combinations in CCL keys and values.

**Architecture:** New `fuzz/` package with generator logic, thin CLI wrapper in `cmd/ccl-test-runner/generate_fuzz.go`, justfile integration. Output goes to `source_tests/fuzz/api_fuzz_special_chars.json` in source format.

**Tech Stack:** Go 1.25, `math/rand/v2`, `encoding/json`, existing `internal/mock` for validation

---

### Task 1: Create `fuzz/generator.go` — Core Generator Package

**Files:**
- Create: `fuzz/generator.go`
- Create: `fuzz/generator_test.go`

**Step 1: Write the failing test**

Create `fuzz/generator_test.go`:

```go
package fuzz

import (
	"testing"
)

func TestGenerateWithDefaultSeed(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{
		Seed:  42,
		Count: 10,
	})
	suite, err := gen.Generate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(suite.Tests) != 10 {
		t.Errorf("expected 10 tests, got %d", len(suite.Tests))
	}
}

func TestGenerateIsDeterministic(t *testing.T) {
	gen1 := NewGenerator(GeneratorOptions{Seed: 42, Count: 5})
	gen2 := NewGenerator(GeneratorOptions{Seed: 42, Count: 5})

	suite1, _ := gen1.Generate()
	suite2, _ := gen2.Generate()

	for i := range suite1.Tests {
		if suite1.Tests[i].Name != suite2.Tests[i].Name {
			t.Errorf("test %d: names differ: %s vs %s", i, suite1.Tests[i].Name, suite2.Tests[i].Name)
		}
	}
}

func TestGenerateDifferentSeedsDifferentOutput(t *testing.T) {
	gen1 := NewGenerator(GeneratorOptions{Seed: 42, Count: 5})
	gen2 := NewGenerator(GeneratorOptions{Seed: 99, Count: 5})

	suite1, _ := gen1.Generate()
	suite2, _ := gen2.Generate()

	allSame := true
	for i := range suite1.Tests {
		if suite1.Tests[i].Name != suite2.Tests[i].Name {
			allSame = false
			break
		}
	}
	if allSame {
		t.Error("expected different seeds to produce different output")
	}
}

func TestAllTestsHaveParseValidation(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 10})
	suite, _ := gen.Generate()

	for _, test := range suite.Tests {
		hasParse := false
		for _, v := range test.Tests {
			if v.Function == "parse" {
				hasParse = true
				break
			}
		}
		if !hasParse {
			t.Errorf("test %s missing parse validation", test.Name)
		}
	}
}

func TestTestNamesAreUnique(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	suite, _ := gen.Generate()

	seen := map[string]bool{}
	for _, test := range suite.Tests {
		if seen[test.Name] {
			t.Errorf("duplicate test name: %s", test.Name)
		}
		seen[test.Name] = true
	}
}

func TestTestNamesMatchPattern(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 10})
	suite, _ := gen.Generate()

	for _, test := range suite.Tests {
		// Names must match ^[a-zA-Z0-9_]+$ per source-format.json schema
		for _, c := range test.Name {
			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
				t.Errorf("test name %q contains invalid character %c", test.Name, c)
				break
			}
		}
	}
}
```

**Step 2: Run test to verify it fails**

Run: `go test ./fuzz/ -v`
Expected: FAIL — package doesn't exist yet

**Step 3: Write the implementation**

Create `fuzz/generator.go`. This is the core of the feature — it defines the special character sets, test category builders, and the source-format output structure.

Key implementation notes:
- Use `math/rand/v2` with `rand.New(rand.NewPCG(uint64(seed), 0))` for deterministic output
- The output structure must match the JSON source format schema: `{"$schema": "...", "tests": [...]}`
- Each test has: `name` (snake_case, `^[a-zA-Z0-9_]+$`), `inputs` (array of CCL strings), `tests` (array of `{function, expect}` objects)
- For `parse` validations, `expect` is an array of `{key, value}` entries
- For `build_hierarchy` validations, `expect` is a nested object
- For `get_string` validations, `expect` is a string with `args` array for the key path
- Use the mock CCL implementation (`internal/mock`) to compute expected values from generated inputs — don't hand-compute them
- Special chars list from design: `/ \ : ; @ # $ % & * + - _ ( ) [ ] { } < > ' " ~ ` ` | ? !`
- **Important:** The `=` character cannot appear in keys (it's the CCL delimiter). Exclude `=` from key character sets.
- **Important:** Newline characters cannot appear in keys. Keys are single-line.
- Test name encoding: use category prefix + index, e.g., `fuzz_single_char_001`, `fuzz_combo_keys_003`, `fuzz_nested_005`

The generator should allocate ~50 tests across 5 categories:
1. `generateSingleCharTests(rng, 15)` — one special char as key
2. `generateComboKeyTests(rng, 10)` — 2-4 random special chars as key
3. `generatePositionalTests(rng, 10)` — special char at start/middle/end of alphanumeric key
4. `generateSpecialValueTests(rng, 5)` — special chars in values
5. `generateNestedTests(rng, 10)` — hierarchical CCL with special char keys

```go
package fuzz

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"

	"github.com/catconflang/ccl-test-data/internal/mock"
)

// Special character categories for fuzz testing
var (
	PathSeparators = []rune{'/', '\\'}
	Punctuation    = []rune{':', ';', '@', '#', '$', '%', '&', '*', '+', '-', '_'}
	Brackets       = []rune{'(', ')', '[', ']', '{', '}', '<', '>'}
	Quotes         = []rune{'\'', '"'}
	Other          = []rune{'~', '`', '|', '?', '!'}

	// AllSpecialChars is the combined set (excludes = and whitespace)
	AllSpecialChars []rune
)

func init() {
	AllSpecialChars = append(AllSpecialChars, PathSeparators...)
	AllSpecialChars = append(AllSpecialChars, Punctuation...)
	AllSpecialChars = append(AllSpecialChars, Brackets...)
	AllSpecialChars = append(AllSpecialChars, Quotes...)
	AllSpecialChars = append(AllSpecialChars, Other...)
}

// SourceTest represents a test in source format
type SourceTest struct {
	Name     string       `json:"name"`
	Tests    []Validation `json:"tests"`
	Inputs   []string     `json:"inputs"`
	Features []string     `json:"features,omitempty"`
}

// Validation represents a function validation
type Validation struct {
	Function string      `json:"function"`
	Expect   interface{} `json:"expect"`
	Args     []string    `json:"args,omitempty"`
}

// SourceFile represents the top-level source format JSON
type SourceFile struct {
	Schema string       `json:"$schema"`
	Tests  []SourceTest `json:"tests"`
}

// GeneratorOptions configures fuzz test generation
type GeneratorOptions struct {
	Seed  int64
	Count int
}

// Generator produces randomized fuzz test cases
type Generator struct {
	opts GeneratorOptions
	rng  *rand.Rand
	ccl  *mock.CCL
}

// NewGenerator creates a new fuzz test generator
func NewGenerator(opts GeneratorOptions) *Generator {
	if opts.Count <= 0 {
		opts.Count = 50
	}
	return &Generator{
		opts: opts,
		rng:  rand.New(rand.NewPCG(uint64(opts.Seed), 0)),
		ccl:  mock.New(),
	}
}

// Generate produces the fuzz test suite
func (g *Generator) Generate() (*SourceFile, error) {
	// Distribute count across categories (proportional to design)
	total := g.opts.Count
	singleCount := max(1, total*15/50)
	comboCount := max(1, total*10/50)
	positionalCount := max(1, total*10/50)
	valueCount := max(1, total*5/50)
	nestedCount := total - singleCount - comboCount - positionalCount - valueCount
	if nestedCount < 1 {
		nestedCount = 1
	}

	var tests []SourceTest
	tests = append(tests, g.generateSingleCharTests(singleCount)...)
	tests = append(tests, g.generateComboKeyTests(comboCount)...)
	tests = append(tests, g.generatePositionalTests(positionalCount)...)
	tests = append(tests, g.generateSpecialValueTests(valueCount)...)
	tests = append(tests, g.generateNestedTests(nestedCount)...)

	return &SourceFile{
		Schema: "../../schemas/source-format.json",
		Tests:  tests,
	}, nil
}

// WriteToFile writes the generated test suite to a JSON file
func (g *Generator) WriteToFile(suite *SourceFile, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	data, err := json.MarshalIndent(suite, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Ensure trailing newline
	data = append(data, '\n')

	outputFile := filepath.Join(outputDir, "api_fuzz_special_chars.json")
	if err := os.WriteFile(outputFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
```

Then implement the 5 `generate*` category methods. Each method:
1. Constructs a CCL input string
2. Runs it through `g.ccl.Parse()` to get expected entries
3. Builds the `SourceTest` with `parse` validation (and `build_hierarchy`/`get_string` where appropriate)

Key helper: `charName(r rune) string` — converts a rune to a safe snake_case name for test naming (e.g., `'/'` → `"slash"`, `'@'` → `"at"`, etc.)

**Step 4: Run test to verify it passes**

Run: `go test ./fuzz/ -v`
Expected: PASS

**Step 5: Commit**

```bash
git add fuzz/
git commit -m "feat(fuzz): add core fuzz test generator package

Seeded randomized generator producing source-format JSON tests
for special character combinations in CCL keys and values.
Covers 5 categories: single char, combo, positional, values, nested."
```

---

### Task 2: Add mock validation to the generator

**Files:**
- Modify: `fuzz/generator.go`
- Modify: `fuzz/generator_test.go`

**Step 1: Write the failing test**

Add to `fuzz/generator_test.go`:

```go
func TestValidateGeneratedTests(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	suite, err := gen.Generate()
	if err != nil {
		t.Fatalf("generation error: %v", err)
	}

	errors := gen.Validate(suite)
	for _, e := range errors {
		t.Errorf("validation error: %s", e)
	}
	if len(errors) > 0 {
		t.Fatalf("%d validation errors found", len(errors))
	}
}
```

**Step 2: Run test to verify it fails**

Run: `go test ./fuzz/ -v -run TestValidateGeneratedTests`
Expected: FAIL — `Validate` method not defined

**Step 3: Write the implementation**

Add a `Validate` method to `Generator` that:
1. For each test, parses the input through `g.ccl.Parse()`
2. Compares the mock output to the `expect` field in the test's `parse` validation
3. For `build_hierarchy` validations, runs `g.ccl.BuildHierarchy()` and compares
4. Returns a list of error strings (empty = all valid)

```go
// Validate checks all generated tests against the mock implementation
func (g *Generator) Validate(suite *SourceFile) []string {
	var errors []string
	for _, test := range suite.Tests {
		if len(test.Inputs) == 0 {
			errors = append(errors, fmt.Sprintf("%s: no inputs", test.Name))
			continue
		}

		entries, err := g.ccl.Parse(test.Inputs[0])
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: parse error: %v", test.Name, err))
			continue
		}

		for _, v := range test.Tests {
			switch v.Function {
			case "parse":
				// Compare entries with expected
				expectedJSON, _ := json.Marshal(v.Expect)
				actualJSON, _ := json.Marshal(entriesToExpect(entries))
				if string(expectedJSON) != string(actualJSON) {
					errors = append(errors, fmt.Sprintf("%s/parse: expected %s, got %s", test.Name, expectedJSON, actualJSON))
				}
			case "build_hierarchy":
				filtered := g.ccl.Filter(entries)
				obj := g.ccl.BuildHierarchy(filtered)
				expectedJSON, _ := json.Marshal(v.Expect)
				actualJSON, _ := json.Marshal(obj)
				if string(expectedJSON) != string(actualJSON) {
					errors = append(errors, fmt.Sprintf("%s/build_hierarchy: expected %s, got %s", test.Name, expectedJSON, actualJSON))
				}
			}
		}
	}
	return errors
}
```

**Step 4: Run test to verify it passes**

Run: `go test ./fuzz/ -v -run TestValidateGeneratedTests`
Expected: PASS

**Step 5: Commit**

```bash
git add fuzz/
git commit -m "feat(fuzz): add mock validation for generated tests

Validates each generated test against the mock CCL implementation
to ensure expected values are correct before committing."
```

---

### Task 3: Add `generate-fuzz` CLI subcommand

**Files:**
- Create: `cmd/ccl-test-runner/generate_fuzz.go`
- Modify: `cmd/ccl-test-runner/main.go`

**Step 1: Create the CLI wrapper**

Create `cmd/ccl-test-runner/generate_fuzz.go`:

```go
package main

import (
	"fmt"

	"github.com/catconflang/ccl-test-data/fuzz"
	"github.com/catconflang/ccl-test-data/internal/styles"
	"github.com/urfave/cli/v2"
)

func generateFuzzAction(ctx *cli.Context) error {
	seed := ctx.Int64("seed")
	count := ctx.Int("count")
	outputDir := ctx.String("output")
	validate := ctx.Bool("validate")

	styles.Status("🎲", fmt.Sprintf("Generating %d fuzz tests (seed: %d)...", count, seed))

	gen := fuzz.NewGenerator(fuzz.GeneratorOptions{
		Seed:  seed,
		Count: count,
	})

	suite, err := gen.Generate()
	if err != nil {
		return fmt.Errorf("error generating fuzz tests: %w", err)
	}

	if validate {
		styles.InfoLite("Validating generated tests against mock implementation...")
		errors := gen.Validate(suite)
		for _, e := range errors {
			styles.Warning("  %s", e)
		}
		if len(errors) > 0 {
			return fmt.Errorf("%d validation errors found", len(errors))
		}
		styles.InfoLite("All %d tests validated successfully", len(suite.Tests))
	}

	if err := gen.WriteToFile(suite, outputDir); err != nil {
		return fmt.Errorf("error writing fuzz tests: %w", err)
	}

	styles.Success("✅ Generated %d fuzz tests in %s", len(suite.Tests), outputDir)
	return nil
}
```

**Step 2: Register the subcommand in main.go**

Add to the `Commands` slice in `cmd/ccl-test-runner/main.go`, after the `generate-flat` command block (around line 231):

```go
{
	Name:    "generate-fuzz",
	Aliases: []string{"fuzz"},
	Usage:   "Generate randomized fuzz tests for special character edge cases",
	Description: `Generate randomized test cases with special character combinations
in keys and values. Uses seeded randomness for reproducible output.

Output is written in source format to source_tests/fuzz/, ready for
the generate-flat pipeline.`,
	Action: generateFuzzAction,
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "seed",
			Value: 42,
			Usage: "Random seed for reproducible generation",
		},
		&cli.IntFlag{
			Name:  "count",
			Value: 50,
			Usage: "Number of test cases to generate",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Value:   "source_tests/fuzz",
			Usage:   "Output directory for generated test files",
		},
		&cli.BoolFlag{
			Name:  "validate",
			Value: true,
			Usage: "Validate generated tests against mock implementation",
		},
	},
},
```

**Step 3: Verify it compiles and runs**

Run: `go build ./cmd/ccl-test-runner && ./bin/ccl-test-runner generate-fuzz --help`
Expected: Shows help with seed, count, output, validate flags

Run: `go run ./cmd/ccl-test-runner generate-fuzz --validate`
Expected: Generates 50 tests in `source_tests/fuzz/`, all pass validation

**Step 4: Commit**

```bash
git add cmd/ccl-test-runner/generate_fuzz.go cmd/ccl-test-runner/main.go
git commit -m "feat(cli): add generate-fuzz subcommand

Thin CLI wrapper for the fuzz package, following the same pattern
as generate-flat. Supports --seed, --count, --output, --validate flags."
```

---

### Task 4: Update justfile and generate-flat pipeline

**Files:**
- Modify: `justfile`

**Step 1: Add generate-fuzz recipe and update generate-flat**

Add after the `generate-flat` recipe (around line 74):

```just
# Generate fuzz test cases (seeded randomized special character tests)
generate-fuzz *ARGS="":
    go run ./cmd/ccl-test-runner generate-fuzz {{ARGS}}
```

Update the `generate-flat` recipe to process both directories:

```just
generate-flat *ARGS="":
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/core --validate {{ARGS}}
    go run ./cmd/ccl-test-runner generate-flat --source ./source_tests/fuzz --validate {{ARGS}}
```

**Step 2: Verify the full pipeline works**

Run: `just generate-fuzz`
Expected: Generates `source_tests/fuzz/api_fuzz_special_chars.json`

Run: `just generate-flat`
Expected: Processes both `core/` and `fuzz/` directories, outputs to `generated_tests/`

Run: `just generate-go`
Expected: Generates Go test files from all flat tests including fuzz

Run: `just test`
Expected: All tests pass (including new fuzz tests)

**Step 3: Commit**

```bash
git add justfile
git commit -m "feat(build): integrate fuzz generator into justfile pipeline

Adds generate-fuzz recipe and updates generate-flat to process
both source_tests/core/ and source_tests/fuzz/ directories."
```

---

### Task 5: Generate and commit the initial fuzz test data

**Files:**
- Create (generated): `source_tests/fuzz/api_fuzz_special_chars.json`

**Step 1: Generate the committed test data**

Run: `just generate-fuzz --validate`
Expected: Creates `source_tests/fuzz/api_fuzz_special_chars.json` with ~50 tests

**Step 2: Run the full pipeline to verify**

Run: `just reset`
Expected: Full clean → build → lint → test cycle passes

Run: `just validate`
Expected: JSON schema validation passes for both source and generated formats

**Step 3: Commit the generated data**

```bash
git add source_tests/fuzz/
git commit -m "feat(tests): add generated fuzz test data for special characters

50 tests covering special character edge cases in CCL keys and values.
Generated with seed 42 for reproducibility.

Categories: single char keys, combo keys, positional chars,
special values, nested structures with special char keys."
```

---

### Task 6: Final verification — full pipeline

**Step 1: Run the full CI pipeline**

Run: `just lint`
Expected: No lint issues

Run: `just reset`
Expected: Full clean build and test cycle passes

Run: `just stats`
Expected: Shows increased test/assertion counts (should be ~278 tests, ~540+ assertions)

Run: `go test ./fuzz/ -v`
Expected: All fuzz package tests pass

**Step 2: Final commit (if any generated files changed)**

```bash
git add generated_tests/ go_tests/
git commit -m "chore: regenerate flat and Go test files with fuzz tests"
```
