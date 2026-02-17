package fuzz

import (
	"fmt"
	"sort"
	"strings"
)

// DatasetStats holds extracted metrics for a single seed's dataset.
type DatasetStats struct {
	Seed         int64
	Names        []string
	SingleChars  []string
	CompoundKeys []string
	Inputs       []string
	Values       []string
}

// CompareReport holds the comparison results across multiple fuzz datasets.
type CompareReport struct {
	Seeds []int64
	Count int

	// Per-seed stats
	Datasets []DatasetStats

	// Categories match across all seeds
	CategoriesMatch bool

	// Aggregate metrics
	NamesSharedByAll  int
	NamesInAnySeed    int
	CharsSharedByAll  int
	CharsInAnySeed    int
	KeysSharedByAll   int
	KeysInAnySeed     int
	InputsSharedByAll int
	InputsInAnySeed   int
	ValuesSharedByAll int
	ValuesInAnySeed   int
}

// Compare generates datasets for each seed and compares them.
func Compare(seeds []int64, count int) (*CompareReport, error) {
	if len(seeds) < 2 {
		return nil, fmt.Errorf("need at least 2 seeds to compare, got %d", len(seeds))
	}

	report := &CompareReport{
		Seeds: seeds,
		Count: count,
	}

	var suites []*SourceFile
	for _, seed := range seeds {
		gen := NewGenerator(GeneratorOptions{Seed: seed, Count: count})
		suite, err := gen.Generate()
		if err != nil {
			return nil, fmt.Errorf("generate seed %d: %w", seed, err)
		}
		suites = append(suites, suite)

		report.Datasets = append(report.Datasets, DatasetStats{
			Seed:         seed,
			Names:        testNames(suite),
			SingleChars:  singleCharKeys(suite),
			CompoundKeys: allKeys(suite),
			Inputs:       allInputs(suite),
			Values:       allValues(suite),
		})
	}

	// Check categories match across all seeds
	report.CategoriesMatch = true
	firstCats := categoryBreakdown(suites[0])
	for i := 1; i < len(suites); i++ {
		if !mapsEqual(firstCats, categoryBreakdown(suites[i])) {
			report.CategoriesMatch = false
			break
		}
	}

	// Compute aggregate set metrics
	report.NamesSharedByAll, report.NamesInAnySeed = multiSetStats(extractField(report.Datasets, func(d DatasetStats) []string { return d.Names }))
	report.CharsSharedByAll, report.CharsInAnySeed = multiSetStats(extractField(report.Datasets, func(d DatasetStats) []string { return d.SingleChars }))
	report.KeysSharedByAll, report.KeysInAnySeed = multiSetStats(extractField(report.Datasets, func(d DatasetStats) []string { return d.CompoundKeys }))
	report.InputsSharedByAll, report.InputsInAnySeed = multiSetStats(extractField(report.Datasets, func(d DatasetStats) []string { return d.Inputs }))
	report.ValuesSharedByAll, report.ValuesInAnySeed = multiSetStats(extractField(report.Datasets, func(d DatasetStats) []string { return d.Values }))

	return report, nil
}

// FormatReport produces a human-readable comparison report.
func (r *CompareReport) FormatReport() string {
	var sb strings.Builder
	n := len(r.Seeds)

	seedStrs := make([]string, n)
	for i, s := range r.Seeds {
		seedStrs[i] = fmt.Sprintf("%d", s)
	}

	sb.WriteString(fmt.Sprintf("Fuzz Dataset Comparison: %d seeds [%s] (%d tests each)\n",
		n, strings.Join(seedStrs, ", "), r.Count))
	sb.WriteString(strings.Repeat("=", 60) + "\n\n")

	// Structure
	sb.WriteString("Structure\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	match := "Yes"
	if !r.CategoriesMatch {
		match = "NO"
	}
	sb.WriteString(fmt.Sprintf("  Categories match:  %s\n\n", match))

	// Per-seed table
	sb.WriteString("Per-Seed Counts\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	sb.WriteString(fmt.Sprintf("  %-12s %6s %6s %6s %6s\n", "Seed", "Names", "Chars", "Keys", "Values"))
	for _, ds := range r.Datasets {
		sb.WriteString(fmt.Sprintf("  %-12d %6d %6d %6d %6d\n",
			ds.Seed, len(ds.Names), len(ds.SingleChars), len(ds.CompoundKeys), len(ds.Values)))
	}
	sb.WriteString("\n")

	// Aggregate coverage
	sb.WriteString("Combined Coverage\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	sb.WriteString(fmt.Sprintf("  %-22s %6s %6s\n", "Metric", "All", "Any"))
	sb.WriteString(fmt.Sprintf("  %-22s %6d %6d\n", "Test names", r.NamesSharedByAll, r.NamesInAnySeed))
	sb.WriteString(fmt.Sprintf("  %-22s %6d %6d\n", "Single-char keys", r.CharsSharedByAll, r.CharsInAnySeed))
	sb.WriteString(fmt.Sprintf("  %-22s %6d %6d\n", "Compound keys", r.KeysSharedByAll, r.KeysInAnySeed))
	sb.WriteString(fmt.Sprintf("  %-22s %6d %6d\n", "Unique inputs", r.InputsSharedByAll, r.InputsInAnySeed))
	sb.WriteString(fmt.Sprintf("  %-22s %6d %6d\n", "Unique values", r.ValuesSharedByAll, r.ValuesInAnySeed))
	sb.WriteString("\n")
	sb.WriteString("  All = present in every seed; Any = present in at least one seed\n")

	return sb.String()
}

// helpers

// multiSetStats computes intersection and union sizes across multiple string sets.
func multiSetStats(sets [][]string) (sharedByAll, inAny int) {
	if len(sets) == 0 {
		return 0, 0
	}

	// Union: everything across all sets
	union := map[string]bool{}
	// Count: how many sets contain each item
	counts := map[string]int{}

	for _, set := range sets {
		seen := map[string]bool{}
		for _, s := range set {
			union[s] = true
			if !seen[s] {
				counts[s]++
				seen[s] = true
			}
		}
	}

	n := len(sets)
	for _, count := range counts {
		if count == n {
			sharedByAll++
		}
	}
	inAny = len(union)
	return
}

// extractField extracts a string slice from each dataset using the given accessor.
func extractField(datasets []DatasetStats, fn func(DatasetStats) []string) [][]string {
	result := make([][]string, len(datasets))
	for i, ds := range datasets {
		result[i] = fn(ds)
	}
	return result
}

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
