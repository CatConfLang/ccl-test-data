package fuzz

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateWithDefaultSeed(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 10})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() returned error: %v", err)
	}
	if sf == nil {
		t.Fatal("Generate() returned nil SourceFile")
	}
	if len(sf.Tests) < 10 {
		t.Errorf("expected at least 10 tests, got %d", len(sf.Tests))
	}
}

func TestGenerateIsDeterministic(t *testing.T) {
	gen1 := NewGenerator(GeneratorOptions{Seed: 123, Count: 20})
	sf1, err := gen1.Generate()
	if err != nil {
		t.Fatalf("first Generate() error: %v", err)
	}

	gen2 := NewGenerator(GeneratorOptions{Seed: 123, Count: 20})
	sf2, err := gen2.Generate()
	if err != nil {
		t.Fatalf("second Generate() error: %v", err)
	}

	if len(sf1.Tests) != len(sf2.Tests) {
		t.Fatalf("test count mismatch: %d vs %d", len(sf1.Tests), len(sf2.Tests))
	}

	for i := range sf1.Tests {
		if sf1.Tests[i].Name != sf2.Tests[i].Name {
			t.Errorf("test %d name mismatch: %q vs %q", i, sf1.Tests[i].Name, sf2.Tests[i].Name)
		}
	}
}

func TestGenerateDifferentSeedsDifferentOutput(t *testing.T) {
	gen1 := NewGenerator(GeneratorOptions{Seed: 100, Count: 20})
	sf1, err := gen1.Generate()
	if err != nil {
		t.Fatalf("seed 100 Generate() error: %v", err)
	}

	gen2 := NewGenerator(GeneratorOptions{Seed: 999, Count: 20})
	sf2, err := gen2.Generate()
	if err != nil {
		t.Fatalf("seed 999 Generate() error: %v", err)
	}

	// At least some test names should differ
	differentCount := 0
	minLen := min(len(sf1.Tests), len(sf2.Tests))
	for i := 0; i < minLen; i++ {
		if sf1.Tests[i].Name != sf2.Tests[i].Name {
			differentCount++
		}
	}

	if differentCount == 0 {
		t.Error("different seeds produced identical test names")
	}
}

func TestAllTestsHaveParseValidation(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	for _, test := range sf.Tests {
		hasParseValidation := false
		for _, v := range test.Tests {
			if v.Function == "parse" {
				hasParseValidation = true
				break
			}
		}
		if !hasParseValidation {
			t.Errorf("test %q has no parse validation", test.Name)
		}
	}
}

func TestTestNamesAreUnique(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	seen := make(map[string]bool)
	for _, test := range sf.Tests {
		if seen[test.Name] {
			t.Errorf("duplicate test name: %q", test.Name)
		}
		seen[test.Name] = true
	}
}

func TestTestNamesMatchPattern(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	pattern := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	for _, test := range sf.Tests {
		if !pattern.MatchString(test.Name) {
			t.Errorf("test name %q does not match pattern ^[a-zA-Z0-9_]+$", test.Name)
		}
	}
}

func TestValidateGeneratedTests(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	errors := gen.Validate(sf)
	for _, e := range errors {
		t.Errorf("validation error: %s", e)
	}
}

func TestWriteToFile(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 10})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	tmpDir := t.TempDir()
	err = gen.WriteToFile(sf, tmpDir)
	if err != nil {
		t.Fatalf("WriteToFile() error: %v", err)
	}

	outPath := filepath.Join(tmpDir, gen.Filename())
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output file: %v", err)
	}

	// Verify trailing newline
	if len(data) == 0 || data[len(data)-1] != '\n' {
		t.Error("output file missing trailing newline")
	}

	// Verify it starts with JSON object
	if data[0] != '{' {
		t.Error("output file does not start with '{'")
	}
}

func TestSchemaField(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 10})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	expected := "../../schemas/source-format.json"
	if sf.Schema != expected {
		t.Errorf("$schema: expected %q, got %q", expected, sf.Schema)
	}
}

func TestCharName(t *testing.T) {
	cases := []struct {
		r    rune
		want string
	}{
		{'/', "slash"},
		{'\\', "backslash"},
		{'@', "at"},
		{'#', "hash"},
		{'$', "dollar"},
		{'!', "bang"},
		{'*', "asterisk"},
		{'?', "question"},
		{'|', "pipe"},
		{'"', "dquote"},
		{'\'', "squote"},
		{'~', "tilde"},
		{'`', "backtick"},
	}

	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			got := charName(tc.r)
			if got != tc.want {
				t.Errorf("charName(%q) = %q, want %q", tc.r, got, tc.want)
			}
		})
	}
}

func TestAllSpecialCharsExcludesEquals(t *testing.T) {
	for _, ch := range AllSpecialChars {
		if ch == '=' {
			t.Error("AllSpecialChars should not contain '='")
		}
	}
}

func TestNestedTestsHaveBuildHierarchy(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	for _, test := range sf.Tests {
		if !isNestedTest(test.Name) {
			continue
		}
		hasBH := false
		for _, v := range test.Tests {
			if v.Function == "build_hierarchy" {
				hasBH = true
				break
			}
		}
		if !hasBH {
			t.Errorf("nested test %q missing build_hierarchy validation", test.Name)
		}
	}
}

func isNestedTest(name string) bool {
	return strings.Contains(name, "fuzz_nested_")
}

func TestInputsNonEmpty(t *testing.T) {
	gen := NewGenerator(GeneratorOptions{Seed: 42, Count: 50})
	sf, err := gen.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}

	for _, test := range sf.Tests {
		if len(test.Inputs) == 0 {
			t.Errorf("test %q has no inputs", test.Name)
		}
		for i, input := range test.Inputs {
			if input == "" {
				t.Errorf("test %q input[%d] is empty", test.Name, i)
			}
		}
	}
}
