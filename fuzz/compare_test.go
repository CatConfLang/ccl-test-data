package fuzz

import (
	"testing"
)

func TestCompare(t *testing.T) {
	report, err := Compare(42, 99, 50)
	if err != nil {
		t.Fatalf("Compare() error: %v", err)
	}

	if !report.CategoriesMatch {
		t.Error("expected categories to match across seeds")
	}

	if len(report.SharedNames)+len(report.OnlyInA)+len(report.OnlyInB) != 50+len(report.OnlyInB) {
		// Sanity: shared + onlyA should cover all of A
		if len(report.SharedNames)+len(report.OnlyInA) != 50 {
			t.Errorf("name accounting error: shared(%d) + onlyInA(%d) != 50",
				len(report.SharedNames), len(report.OnlyInA))
		}
	}

	if report.SharedInputs == 50 {
		t.Error("expected different seeds to produce different inputs")
	}
}

func TestCompareSameSeed(t *testing.T) {
	report, err := Compare(42, 42, 10)
	if err != nil {
		t.Fatalf("Compare() error: %v", err)
	}

	if len(report.OnlyInA) != 0 || len(report.OnlyInB) != 0 {
		t.Errorf("same seed should produce identical names: onlyA=%d onlyB=%d",
			len(report.OnlyInA), len(report.OnlyInB))
	}

	if report.SharedInputs != report.TotalInputsA {
		t.Errorf("same seed should have all inputs shared: %d != %d",
			report.SharedInputs, report.TotalInputsA)
	}
}

func TestCompareFormatReport(t *testing.T) {
	report, err := Compare(42, 99, 10)
	if err != nil {
		t.Fatalf("Compare() error: %v", err)
	}

	output := report.FormatReport()
	if len(output) == 0 {
		t.Error("FormatReport() returned empty string")
	}

	// Should contain both seed values
	if !containsStr(output, "seed 42") || !containsStr(output, "seed 99") {
		t.Error("report should mention both seeds")
	}
}

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && searchStr(s, substr)
}

func searchStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
