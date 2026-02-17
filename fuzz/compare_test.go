package fuzz

import (
	"strings"
	"testing"
)

func TestCompareTwoSeeds(t *testing.T) {
	report, err := Compare([]int64{42, 99}, 50)
	if err != nil {
		t.Fatalf("Compare() error: %v", err)
	}

	if !report.CategoriesMatch {
		t.Error("expected categories to match across seeds")
	}

	if report.InputsSharedByAll == report.InputsInAnySeed {
		t.Error("expected different seeds to have some non-overlapping inputs")
	}

	if report.NamesInAnySeed <= 50 {
		t.Error("expected union of names to exceed a single seed's count")
	}
}

func TestCompareThreeSeeds(t *testing.T) {
	report, err := Compare([]int64{42, 99, 7}, 20)
	if err != nil {
		t.Fatalf("Compare() error: %v", err)
	}

	if len(report.Datasets) != 3 {
		t.Errorf("expected 3 datasets, got %d", len(report.Datasets))
	}

	// With 3 seeds, union should be larger than any single seed
	if report.KeysInAnySeed <= len(report.Datasets[0].CompoundKeys) {
		t.Error("expected combined key coverage to exceed a single seed")
	}
}

func TestCompareSameSeed(t *testing.T) {
	report, err := Compare([]int64{42, 42}, 10)
	if err != nil {
		t.Fatalf("Compare() error: %v", err)
	}

	if report.NamesSharedByAll != report.NamesInAnySeed {
		t.Errorf("same seed should have identical names: all=%d any=%d",
			report.NamesSharedByAll, report.NamesInAnySeed)
	}

	if report.InputsSharedByAll != report.InputsInAnySeed {
		t.Errorf("same seed should have identical inputs: all=%d any=%d",
			report.InputsSharedByAll, report.InputsInAnySeed)
	}
}

func TestCompareNeedsTwoSeeds(t *testing.T) {
	_, err := Compare([]int64{42}, 10)
	if err == nil {
		t.Error("expected error with only 1 seed")
	}
}

func TestCompareFormatReport(t *testing.T) {
	report, err := Compare([]int64{42, 99, 7}, 10)
	if err != nil {
		t.Fatalf("Compare() error: %v", err)
	}

	output := report.FormatReport()
	if len(output) == 0 {
		t.Error("FormatReport() returned empty string")
	}

	for _, seed := range []string{"42", "99", "7"} {
		if !strings.Contains(output, seed) {
			t.Errorf("report should mention seed %s", seed)
		}
	}

	if !strings.Contains(output, "3 seeds") {
		t.Error("report should mention number of seeds")
	}
}
