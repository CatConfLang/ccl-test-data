package fuzz

import (
	"fmt"
	"sort"
	"strings"
)

// CompareReport holds the comparison results between two fuzz datasets.
type CompareReport struct {
	SeedA int64
	SeedB int64
	Count int

	// Category counts (should be identical across seeds)
	CategoriesMatch bool

	// Name analysis
	SharedNames  []string
	OnlyInA      []string
	OnlyInB      []string

	// Key analysis
	SingleCharsA    []string
	SingleCharsB    []string
	SharedChars     []string
	UniqueCharsA    []string
	UniqueCharsB    []string
	CompoundKeysA   int
	CompoundKeysB   int
	SharedCompound  int

	// Input/value overlap
	SharedInputs int
	SharedValues int
	TotalInputsA int
	TotalInputsB int
	UniqueValuesA int
	UniqueValuesB int
}

// Compare generates two datasets with different seeds and compares them.
func Compare(seedA, seedB int64, count int) (*CompareReport, error) {
	genA := NewGenerator(GeneratorOptions{Seed: seedA, Count: count})
	genB := NewGenerator(GeneratorOptions{Seed: seedB, Count: count})

	suiteA, err := genA.Generate()
	if err != nil {
		return nil, fmt.Errorf("generate seed %d: %w", seedA, err)
	}
	suiteB, err := genB.Generate()
	if err != nil {
		return nil, fmt.Errorf("generate seed %d: %w", seedB, err)
	}

	report := &CompareReport{
		SeedA: seedA,
		SeedB: seedB,
		Count: count,
	}

	// Compare categories
	catsA := categoryBreakdown(suiteA)
	catsB := categoryBreakdown(suiteB)
	report.CategoriesMatch = mapsEqual(catsA, catsB)

	// Compare names
	namesA := testNames(suiteA)
	namesB := testNames(suiteB)
	report.SharedNames, report.OnlyInA, report.OnlyInB = setDiff(namesA, namesB)

	// Compare single-char keys (from fuzz_single_* tests)
	report.SingleCharsA = singleCharKeys(suiteA)
	report.SingleCharsB = singleCharKeys(suiteB)
	report.SharedChars, report.UniqueCharsA, report.UniqueCharsB = setDiff(report.SingleCharsA, report.SingleCharsB)

	// Compare compound keys
	compA := allKeys(suiteA)
	compB := allKeys(suiteB)
	sharedComp, _, _ := setDiff(compA, compB)
	report.CompoundKeysA = len(compA)
	report.CompoundKeysB = len(compB)
	report.SharedCompound = len(sharedComp)

	// Compare inputs
	inputsA := allInputs(suiteA)
	inputsB := allInputs(suiteB)
	sharedInputs, _, _ := setDiff(inputsA, inputsB)
	report.SharedInputs = len(sharedInputs)
	report.TotalInputsA = len(inputsA)
	report.TotalInputsB = len(inputsB)

	// Compare values
	valsA := allValues(suiteA)
	valsB := allValues(suiteB)
	sharedVals, _, _ := setDiff(valsA, valsB)
	report.SharedValues = len(sharedVals)
	report.UniqueValuesA = len(valsA)
	report.UniqueValuesB = len(valsB)

	return report, nil
}

// FormatReport produces a human-readable comparison report.
func (r *CompareReport) FormatReport() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Fuzz Dataset Comparison: seed %d vs seed %d (%d tests each)\n", r.SeedA, r.SeedB, r.Count))
	sb.WriteString(strings.Repeat("=", 60) + "\n\n")

	// Structure
	sb.WriteString("Structure\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	match := "Yes"
	if !r.CategoriesMatch {
		match = "NO"
	}
	sb.WriteString(fmt.Sprintf("  Categories match:  %s\n\n", match))

	// Names
	sb.WriteString("Test Names\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	sb.WriteString(fmt.Sprintf("  Shared:            %d\n", len(r.SharedNames)))
	sb.WriteString(fmt.Sprintf("  Only in seed %d:   %d\n", r.SeedA, len(r.OnlyInA)))
	sb.WriteString(fmt.Sprintf("  Only in seed %d:   %d\n", r.SeedB, len(r.OnlyInB)))
	sb.WriteString("\n")

	// Single-char coverage
	sb.WriteString("Single-Character Key Coverage\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	sb.WriteString(fmt.Sprintf("  seed %d:            %d chars\n", r.SeedA, len(r.SingleCharsA)))
	sb.WriteString(fmt.Sprintf("  seed %d:            %d chars\n", r.SeedB, len(r.SingleCharsB)))
	sb.WriteString(fmt.Sprintf("  Shared:            %d\n", len(r.SharedChars)))
	sb.WriteString(fmt.Sprintf("  Combined unique:   %d\n", len(r.SharedChars)+len(r.UniqueCharsA)+len(r.UniqueCharsB)))
	if len(r.UniqueCharsA) > 0 {
		sb.WriteString(fmt.Sprintf("  Only in seed %d:   %s\n", r.SeedA, strings.Join(r.UniqueCharsA, " ")))
	}
	if len(r.UniqueCharsB) > 0 {
		sb.WriteString(fmt.Sprintf("  Only in seed %d:   %s\n", r.SeedB, strings.Join(r.UniqueCharsB, " ")))
	}
	sb.WriteString("\n")

	// Compound keys
	sb.WriteString("Compound Keys\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	sb.WriteString(fmt.Sprintf("  seed %d:            %d keys\n", r.SeedA, r.CompoundKeysA))
	sb.WriteString(fmt.Sprintf("  seed %d:            %d keys\n", r.SeedB, r.CompoundKeysB))
	sb.WriteString(fmt.Sprintf("  Shared:            %d\n", r.SharedCompound))
	sb.WriteString(fmt.Sprintf("  Combined unique:   %d\n", r.CompoundKeysA+r.CompoundKeysB-r.SharedCompound))
	sb.WriteString("\n")

	// Inputs and values
	sb.WriteString("Input/Value Overlap\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	sb.WriteString(fmt.Sprintf("  Shared inputs:     %d / %d\n", r.SharedInputs, r.TotalInputsA))
	sb.WriteString(fmt.Sprintf("  Shared values:     %d / %d\n", r.SharedValues, r.UniqueValuesA))

	return sb.String()
}

// helpers

func categoryBreakdown(sf *SourceFile) map[string]int {
	counts := map[string]int{}
	for _, t := range sf.Tests {
		prefix := "unknown"
		for _, p := range []string{"fuzz_single_", "fuzz_combo_", "fuzz_pos_", "fuzz_val_", "fuzz_nested_"} {
			if strings.HasPrefix(t.Name, p) {
				prefix = p
				break
			}
		}
		counts[prefix]++
	}
	return counts
}

func testNames(sf *SourceFile) []string {
	names := make([]string, len(sf.Tests))
	for i, t := range sf.Tests {
		names[i] = t.Name
	}
	return names
}

func singleCharKeys(sf *SourceFile) []string {
	var keys []string
	for _, t := range sf.Tests {
		if strings.HasPrefix(t.Name, "fuzz_single_") && len(t.Inputs) > 0 {
			// Extract the key from "key = value"
			input := t.Inputs[0]
			if idx := strings.Index(input, " = "); idx > 0 {
				keys = append(keys, input[:idx])
			}
		}
	}
	sort.Strings(keys)
	return keys
}

func allKeys(sf *SourceFile) []string {
	seen := map[string]bool{}
	for _, t := range sf.Tests {
		if len(t.Inputs) == 0 {
			continue
		}
		input := t.Inputs[0]
		for _, line := range strings.Split(input, "\n") {
			line = strings.TrimSpace(line)
			if idx := strings.Index(line, " = "); idx > 0 {
				seen[line[:idx]] = true
			} else if idx := strings.Index(line, "="); idx > 0 {
				seen[strings.TrimSpace(line[:idx])] = true
			}
		}
	}
	result := make([]string, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

func allInputs(sf *SourceFile) []string {
	seen := map[string]bool{}
	for _, t := range sf.Tests {
		for _, input := range t.Inputs {
			seen[input] = true
		}
	}
	result := make([]string, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	return result
}

func allValues(sf *SourceFile) []string {
	seen := map[string]bool{}
	for _, t := range sf.Tests {
		if len(t.Inputs) == 0 {
			continue
		}
		for _, line := range strings.Split(t.Inputs[0], "\n") {
			line = strings.TrimSpace(line)
			if idx := strings.Index(line, " = "); idx >= 0 {
				seen[strings.TrimSpace(line[idx+3:])] = true
			}
		}
	}
	result := make([]string, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	return result
}

func setDiff(a, b []string) (shared, onlyA, onlyB []string) {
	setA := map[string]bool{}
	setB := map[string]bool{}
	for _, s := range a {
		setA[s] = true
	}
	for _, s := range b {
		setB[s] = true
	}
	for _, s := range a {
		if setB[s] {
			shared = append(shared, s)
		} else {
			onlyA = append(onlyA, s)
		}
	}
	for _, s := range b {
		if !setA[s] {
			onlyB = append(onlyB, s)
		}
	}
	sort.Strings(shared)
	sort.Strings(onlyA)
	sort.Strings(onlyB)
	return
}

func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
