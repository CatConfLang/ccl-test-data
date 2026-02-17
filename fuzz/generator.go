// Package fuzz provides a generator for CCL fuzz test cases with special characters
// in keys and values. Tests are output in source-format JSON compatible with the
// ccl-test-data pipeline.
package fuzz

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/catconflang/ccl-test-data/internal/mock"
)

// Special character categories for CCL keys and values.
var (
	PathSeparators = []rune{'/', '\\'}
	Punctuation    = []rune{':', ';', '@', '#', '$', '%', '&', '*', '+', '-', '_'}
	Brackets       = []rune{'(', ')', '[', ']', '{', '}', '<', '>'}
	Quotes         = []rune{'\'', '"'}
	// Note: backtick (`) is excluded because the Go test generator uses
	// backtick-delimited raw strings, and backticks cannot be escaped in them.
	Other = []rune{'~', '|', '?', '!'}

	// AllSpecialChars is the combined set of all special characters.
	AllSpecialChars []rune
)

func init() {
	AllSpecialChars = make([]rune, 0,
		len(PathSeparators)+len(Punctuation)+len(Brackets)+len(Quotes)+len(Other))
	AllSpecialChars = append(AllSpecialChars, PathSeparators...)
	AllSpecialChars = append(AllSpecialChars, Punctuation...)
	AllSpecialChars = append(AllSpecialChars, Brackets...)
	AllSpecialChars = append(AllSpecialChars, Quotes...)
	AllSpecialChars = append(AllSpecialChars, Other...)
}

// SourceFile represents a source-format JSON test file.
type SourceFile struct {
	Schema string       `json:"$schema"`
	Tests  []SourceTest `json:"tests"`
}

// SourceTest represents a single test case in source format.
type SourceTest struct {
	Name     string       `json:"name"`
	Tests    []Validation `json:"tests"`
	Inputs   []string     `json:"inputs"`
	Features []string     `json:"features,omitempty"`
}

// Validation represents a function validation within a test.
type Validation struct {
	Function string      `json:"function"`
	Expect   interface{} `json:"expect"`
	Args     []string    `json:"args,omitempty"`
}

// GeneratorOptions configures the fuzz test generator.
type GeneratorOptions struct {
	Seed  int64
	Count int
}

// Generator produces fuzz test cases with special characters.
type Generator struct {
	opts GeneratorOptions
	rng  *rand.Rand
	ccl  *mock.CCL
}

// NewGenerator creates a new fuzz test generator with the given options.
func NewGenerator(opts GeneratorOptions) *Generator {
	return &Generator{
		opts: opts,
		rng:  rand.New(rand.NewPCG(uint64(opts.Seed), 0)),
		ccl:  mock.New(),
	}
}

// Generate produces a SourceFile with fuzz test cases using special characters.
// The total number of tests is distributed proportionally across 5 categories.
func (g *Generator) Generate() (*SourceFile, error) {
	count := g.opts.Count
	if count <= 0 {
		count = 50
	}

	// Distribute count proportionally: ~30% single, ~20% combo, ~20% positional, ~10% value, ~20% nested
	singleCount := max(1, count*30/100)
	comboCount := max(1, count*20/100)
	positionalCount := max(1, count*20/100)
	valueCount := max(1, count*10/100)
	nestedCount := count - singleCount - comboCount - positionalCount - valueCount
	if nestedCount < 1 {
		nestedCount = 1
	}

	var tests []SourceTest

	singleTests, err := g.generateSingleCharTests(singleCount)
	if err != nil {
		return nil, fmt.Errorf("single char tests: %w", err)
	}
	tests = append(tests, singleTests...)

	comboTests, err := g.generateComboKeyTests(comboCount)
	if err != nil {
		return nil, fmt.Errorf("combo key tests: %w", err)
	}
	tests = append(tests, comboTests...)

	positionalTests, err := g.generatePositionalTests(positionalCount)
	if err != nil {
		return nil, fmt.Errorf("positional tests: %w", err)
	}
	tests = append(tests, positionalTests...)

	valueTests, err := g.generateSpecialValueTests(valueCount)
	if err != nil {
		return nil, fmt.Errorf("special value tests: %w", err)
	}
	tests = append(tests, valueTests...)

	nestedTests, err := g.generateNestedTests(nestedCount)
	if err != nil {
		return nil, fmt.Errorf("nested tests: %w", err)
	}
	tests = append(tests, nestedTests...)

	// Deduplicate test names by appending suffix if needed
	tests = deduplicateNames(tests)

	return &SourceFile{
		Schema: "../../schemas/source-format.json",
		Tests:  tests,
	}, nil
}

// WriteToFile writes the generated tests to outputDir/api_fuzz_special_chars.json.
func (g *Generator) WriteToFile(sf *SourceFile, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(sf, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}
	data = append(data, '\n')

	outPath := filepath.Join(outputDir, "api_fuzz_special_chars.json")
	return os.WriteFile(outPath, data, 0o644)
}

// Validate checks all generated tests against the mock CCL implementation.
// Returns a slice of error messages (empty means all tests are valid).
func (g *Generator) Validate(sf *SourceFile) []string {
	var errors []string

	for _, test := range sf.Tests {
		if len(test.Inputs) == 0 {
			errors = append(errors, fmt.Sprintf("test %q: no inputs", test.Name))
			continue
		}
		input := test.Inputs[0]

		// Parse the input through mock
		parsed, err := g.ccl.Parse(input)
		if err != nil {
			errors = append(errors, fmt.Sprintf("test %q: parse error: %v", test.Name, err))
			continue
		}

		for _, v := range test.Tests {
			switch v.Function {
			case "parse":
				expected, ok := v.Expect.([]interface{})
				if !ok {
					errors = append(errors, fmt.Sprintf("test %q: parse expect is not []interface{}, got %T", test.Name, v.Expect))
					continue
				}
				if len(expected) != len(parsed) {
					errors = append(errors, fmt.Sprintf("test %q: parse entry count mismatch: expected %d, got %d", test.Name, len(expected), len(parsed)))
					continue
				}
				for i, exp := range expected {
					entry, ok := exp.(map[string]interface{})
					if !ok {
						errors = append(errors, fmt.Sprintf("test %q: parse expect[%d] is not map", test.Name, i))
						continue
					}
					expKey, _ := entry["key"].(string)
					expVal, _ := entry["value"].(string)
					if parsed[i].Key != expKey {
						errors = append(errors, fmt.Sprintf("test %q: parse entry[%d] key mismatch: expected %q, got %q", test.Name, i, expKey, parsed[i].Key))
					}
					if parsed[i].Value != expVal {
						errors = append(errors, fmt.Sprintf("test %q: parse entry[%d] value mismatch: expected %q, got %q", test.Name, i, expVal, parsed[i].Value))
					}
				}

			case "build_hierarchy":
				filtered := g.ccl.Filter(parsed)
				hierarchy := g.ccl.BuildHierarchy(filtered)
				if !reflect.DeepEqual(normalizeForJSON(hierarchy), normalizeForJSON(v.Expect)) {
					errors = append(errors, fmt.Sprintf("test %q: build_hierarchy mismatch: expected %v, got %v", test.Name, v.Expect, hierarchy))
				}

			case "get_string":
				if len(v.Args) == 0 {
					errors = append(errors, fmt.Sprintf("test %q: get_string has no args", test.Name))
					continue
				}
				filtered := g.ccl.Filter(parsed)
				hierarchy := g.ccl.BuildHierarchy(filtered)
				actual, err := g.ccl.GetString(hierarchy, v.Args)
				if err != nil {
					errors = append(errors, fmt.Sprintf("test %q: get_string error: %v", test.Name, err))
					continue
				}
				expStr, ok := v.Expect.(string)
				if !ok {
					errors = append(errors, fmt.Sprintf("test %q: get_string expect is not string", test.Name))
					continue
				}
				if actual != expStr {
					errors = append(errors, fmt.Sprintf("test %q: get_string mismatch: expected %q, got %q", test.Name, expStr, actual))
				}
			}
		}
	}

	return errors
}

// generateSingleCharTests creates tests with each special char as a standalone key.
func (g *Generator) generateSingleCharTests(count int) ([]SourceTest, error) {
	tests := make([]SourceTest, 0, count)

	// Use as many distinct chars as we can, then loop with random values
	chars := g.shuffledChars()
	for i := 0; i < count; i++ {
		ch := chars[i%len(chars)]

		// Exclude = from keys (CCL delimiter)
		if ch == '=' {
			continue
		}

		key := string(ch)
		value := fmt.Sprintf("val%d", g.rng.IntN(1000))
		input := key + " = " + value

		entries, err := g.ccl.Parse(input)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", input, err)
		}

		validations := []Validation{
			{
				Function: "parse",
				Expect:   entriesToExpect(entries),
			},
		}

		name := fmt.Sprintf("fuzz_single_%s", charName(ch))
		if i >= len(chars) {
			name = fmt.Sprintf("fuzz_single_%s_%d", charName(ch), i)
		}

		tests = append(tests, SourceTest{
			Name:   name,
			Tests:  validations,
			Inputs: []string{input},
		})
	}

	return tests, nil
}

// generateComboKeyTests creates tests with 2-4 special chars combined as a key.
func (g *Generator) generateComboKeyTests(count int) ([]SourceTest, error) {
	tests := make([]SourceTest, 0, count)

	for i := 0; i < count; i++ {
		numChars := 2 + g.rng.IntN(3) // 2-4 chars
		var keyRunes []rune
		var nameParts []string
		for j := 0; j < numChars; j++ {
			ch := g.randomSpecialChar()
			keyRunes = append(keyRunes, ch)
			nameParts = append(nameParts, charName(ch))
		}
		key := string(keyRunes)
		value := fmt.Sprintf("combo%d", g.rng.IntN(1000))
		input := key + " = " + value

		entries, err := g.ccl.Parse(input)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", input, err)
		}

		validations := []Validation{
			{
				Function: "parse",
				Expect:   entriesToExpect(entries),
			},
		}

		name := fmt.Sprintf("fuzz_combo_%s", strings.Join(nameParts, "_"))

		tests = append(tests, SourceTest{
			Name:   name,
			Tests:  validations,
			Inputs: []string{input},
		})
	}

	return tests, nil
}

// generatePositionalTests creates tests with special chars at start/middle/end of keys.
func (g *Generator) generatePositionalTests(count int) ([]SourceTest, error) {
	tests := make([]SourceTest, 0, count)
	positions := []string{"start", "middle", "end"}

	for i := 0; i < count; i++ {
		ch := g.randomSpecialChar()
		pos := positions[i%len(positions)]
		value := fmt.Sprintf("pos%d", g.rng.IntN(1000))

		var key string
		baseWord := g.randomWord()
		switch pos {
		case "start":
			key = string(ch) + baseWord
		case "middle":
			half := len(baseWord) / 2
			if half == 0 {
				half = 1
			}
			key = baseWord[:half] + string(ch) + baseWord[half:]
		case "end":
			key = baseWord + string(ch)
		}

		input := key + " = " + value

		entries, err := g.ccl.Parse(input)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", input, err)
		}

		validations := []Validation{
			{
				Function: "parse",
				Expect:   entriesToExpect(entries),
			},
		}

		name := fmt.Sprintf("fuzz_pos_%s_%s_%s", pos, charName(ch), baseWord)

		tests = append(tests, SourceTest{
			Name:   name,
			Tests:  validations,
			Inputs: []string{input},
		})
	}

	return tests, nil
}

// generateSpecialValueTests creates tests with normal keys but special chars in values.
func (g *Generator) generateSpecialValueTests(count int) ([]SourceTest, error) {
	tests := make([]SourceTest, 0, count)

	for i := 0; i < count; i++ {
		key := g.randomWord()
		// Build a value with 1-3 special chars mixed with alphanumerics
		numSpecial := 1 + g.rng.IntN(3)
		var valueParts []string
		valueParts = append(valueParts, "data")
		for j := 0; j < numSpecial; j++ {
			ch := AllSpecialChars[g.rng.IntN(len(AllSpecialChars))]
			valueParts = append(valueParts, string(ch))
			valueParts = append(valueParts, fmt.Sprintf("x%d", g.rng.IntN(100)))
		}
		value := strings.Join(valueParts, "")

		input := key + " = " + value

		entries, err := g.ccl.Parse(input)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", input, err)
		}

		// Include both parse and get_string validations
		filtered := g.ccl.Filter(entries)
		hierarchy := g.ccl.BuildHierarchy(filtered)
		strVal, err := g.ccl.GetString(hierarchy, []string{key})
		if err != nil {
			return nil, fmt.Errorf("get_string %q: %w", key, err)
		}

		validations := []Validation{
			{
				Function: "parse",
				Expect:   entriesToExpect(entries),
			},
			{
				Function: "build_hierarchy",
				Expect:   hierarchy,
			},
			{
				Function: "get_string",
				Expect:   strVal,
				Args:     []string{key},
			},
		}

		name := fmt.Sprintf("fuzz_val_%s_%d", key, i)

		tests = append(tests, SourceTest{
			Name:     name,
			Tests:    validations,
			Inputs:   []string{input},
			Features: []string{"optional_typed_accessors"},
		})
	}

	return tests, nil
}

// generateNestedTests creates multi-entry CCL where keys contain special chars,
// tested via both parse and build_hierarchy. Each test has 2-3 entries with
// special chars in keys, producing a flat hierarchy object.
func (g *Generator) generateNestedTests(count int) ([]SourceTest, error) {
	tests := make([]SourceTest, 0, count)

	for i := 0; i < count; i++ {
		// Pick special chars for two keys
		char1 := g.randomSpecialChar()
		char2 := g.randomSpecialChar()

		key1 := fmt.Sprintf("%s%s", string(char1), g.randomWord())
		key2 := fmt.Sprintf("%s%s", string(char2), g.randomWord())
		val1 := fmt.Sprintf("nested%d", g.rng.IntN(1000))
		val2 := fmt.Sprintf("deep%d", g.rng.IntN(1000))

		// Build multi-entry CCL: key1 = val1\nkey2 = val2
		input := key1 + " = " + val1 + "\n" + key2 + " = " + val2

		entries, err := g.ccl.Parse(input)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", input, err)
		}

		filtered := g.ccl.Filter(entries)
		hierarchy := g.ccl.BuildHierarchy(filtered)

		// Get values via get_string for the first key
		strVal, err := g.ccl.GetString(hierarchy, []string{key1})
		if err != nil {
			return nil, fmt.Errorf("get_string %q: %w", key1, err)
		}

		validations := []Validation{
			{
				Function: "parse",
				Expect:   entriesToExpect(entries),
			},
			{
				Function: "build_hierarchy",
				Expect:   hierarchy,
			},
			{
				Function: "get_string",
				Expect:   strVal,
				Args:     []string{key1},
			},
		}

		name := fmt.Sprintf("fuzz_nested_%s_%s_%d", charName(char1), charName(char2), i)

		tests = append(tests, SourceTest{
			Name:     name,
			Tests:    validations,
			Inputs:   []string{input},
			Features: []string{"optional_typed_accessors"},
		})
	}

	return tests, nil
}

// charName converts a special character rune to a safe alphanumeric name.
func charName(r rune) string {
	names := map[rune]string{
		'/':  "slash",
		'\\': "backslash",
		':':  "colon",
		';':  "semicolon",
		'@':  "at",
		'#':  "hash",
		'$':  "dollar",
		'%':  "percent",
		'&':  "ampersand",
		'*':  "asterisk",
		'+':  "plus",
		'-':  "hyphen",
		'_':  "underscore",
		'(':  "lparen",
		')':  "rparen",
		'[':  "lbracket",
		']':  "rbracket",
		'{':  "lbrace",
		'}':  "rbrace",
		'<':  "lt",
		'>':  "gt",
		'\'': "squote",
		'"':  "dquote",
		'~':  "tilde",
		'`':  "backtick",
		'|':  "pipe",
		'?':  "question",
		'!':  "bang",
	}
	if name, ok := names[r]; ok {
		return name
	}
	return fmt.Sprintf("u%04x", r)
}

// Helper methods

// shuffledChars returns AllSpecialChars in a random order (excluding '=').
func (g *Generator) shuffledChars() []rune {
	chars := make([]rune, 0, len(AllSpecialChars))
	for _, ch := range AllSpecialChars {
		if ch != '=' {
			chars = append(chars, ch)
		}
	}
	g.rng.Shuffle(len(chars), func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	})
	return chars
}

// randomSpecialChar returns a random special char (excluding '=').
func (g *Generator) randomSpecialChar() rune {
	for {
		ch := AllSpecialChars[g.rng.IntN(len(AllSpecialChars))]
		if ch != '=' {
			return ch
		}
	}
}

// randomWord returns a random simple alphanumeric word from a fixed set.
func (g *Generator) randomWord() string {
	words := []string{
		"alpha", "beta", "gamma", "delta", "epsilon",
		"config", "server", "host", "port", "name",
		"user", "path", "mode", "data", "item",
	}
	return words[g.rng.IntN(len(words))]
}

// entriesToExpect converts mock.Entry slices to the JSON-compatible expect format.
func entriesToExpect(entries []mock.Entry) interface{} {
	result := make([]interface{}, len(entries))
	for i, e := range entries {
		result[i] = map[string]interface{}{
			"key":   e.Key,
			"value": e.Value,
		}
	}
	return result
}

// normalizeForJSON round-trips a value through JSON to normalize types for comparison.
func normalizeForJSON(v interface{}) interface{} {
	data, err := json.Marshal(v)
	if err != nil {
		return v
	}
	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return v
	}
	return result
}

// deduplicateNames ensures all test names are unique by appending suffixes.
func deduplicateNames(tests []SourceTest) []SourceTest {
	seen := make(map[string]int)
	for i := range tests {
		name := tests[i].Name
		candidate := name
		if n, exists := seen[candidate]; exists {
			for {
				n++
				candidate = fmt.Sprintf("%s_%d", name, n)
				if _, exists := seen[candidate]; !exists {
					break
				}
			}
			seen[name] = n
		}
		tests[i].Name = candidate
		seen[candidate] = 0
	}
	return tests
}
