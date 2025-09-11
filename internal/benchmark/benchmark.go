// Package benchmark provides performance tracking and monitoring for CCL test operations.
//
// This package implements benchmark tracking for test generation, statistics collection,
// and CCL parsing operations. It supports both one-time measurements and historical
// performance tracking with regression detection.
//
// Key Features:
//   - Test generation performance benchmarks
//   - Statistics collection timing
//   - Memory allocation tracking
//   - Performance regression detection
//   - JSON output for CI/CD integration
//   - Historical performance comparison
//
// Example Usage:
//
//	tracker := benchmark.NewTracker()
//	tracker.StartBenchmark("test-generation")
//	// ... perform operation
//	result := tracker.EndBenchmark("test-generation")
//	tracker.SaveResults("benchmarks/results.json")
package benchmark

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
)

// BenchmarkResult represents a single benchmark measurement
type BenchmarkResult struct {
	Name          string        `json:"name"`
	Duration      time.Duration `json:"duration"`
	MemAllocBytes int64         `json:"memAllocBytes"`
	MemAllocObjs  int64         `json:"memAllocObjs"`
	Timestamp     time.Time     `json:"timestamp"`
	GitCommit     string        `json:"gitCommit,omitempty"`
	GoVersion     string        `json:"goVersion"`
}

// BenchmarkTracker manages performance measurements
type BenchmarkTracker struct {
	results   map[string]*BenchmarkResult
	startMem  map[string]runtime.MemStats
	startTime map[string]time.Time
}

// NewTracker creates a new benchmark tracker
func NewTracker() *BenchmarkTracker {
	return &BenchmarkTracker{
		results:   make(map[string]*BenchmarkResult),
		startMem:  make(map[string]runtime.MemStats),
		startTime: make(map[string]time.Time),
	}
}

// StartBenchmark begins tracking a named operation
func (bt *BenchmarkTracker) StartBenchmark(name string) {
	// Force garbage collection for accurate memory measurements
	runtime.GC()
	runtime.GC() // Call twice for better accuracy

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	bt.startMem[name] = m
	bt.startTime[name] = time.Now()
}

// EndBenchmark completes tracking and records results
func (bt *BenchmarkTracker) EndBenchmark(name string) *BenchmarkResult {
	endTime := time.Now()

	// Get ending memory stats
	var endMem runtime.MemStats
	runtime.ReadMemStats(&endMem)

	startMem, hasStart := bt.startMem[name]
	startTime, hasTime := bt.startTime[name]

	if !hasStart || !hasTime {
		return nil // Benchmark was never started
	}

	result := &BenchmarkResult{
		Name:          name,
		Duration:      endTime.Sub(startTime),
		MemAllocBytes: int64(endMem.TotalAlloc - startMem.TotalAlloc),
		MemAllocObjs:  int64(endMem.Mallocs - startMem.Mallocs),
		Timestamp:     endTime,
		GoVersion:     runtime.Version(),
	}

	// Try to get git commit hash
	if commit := getGitCommit(); commit != "" {
		result.GitCommit = commit
	}

	bt.results[name] = result

	// Clean up tracking data
	delete(bt.startMem, name)
	delete(bt.startTime, name)

	return result
}

// GetResult returns the result for a named benchmark
func (bt *BenchmarkTracker) GetResult(name string) *BenchmarkResult {
	return bt.results[name]
}

// GetAllResults returns all benchmark results
func (bt *BenchmarkTracker) GetAllResults() map[string]*BenchmarkResult {
	// Return a copy to prevent modification
	results := make(map[string]*BenchmarkResult)
	for k, v := range bt.results {
		results[k] = v
	}
	return results
}

// SaveResults saves benchmark results to a JSON file
func (bt *BenchmarkTracker) SaveResults(filepath string) error {
	data, err := json.MarshalIndent(bt.results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal benchmark results: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write benchmark results to %s: %w", filepath, err)
	}

	return nil
}

// LoadResults loads benchmark results from a JSON file
func LoadResults(filepath string) (map[string]*BenchmarkResult, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read benchmark results from %s: %w", filepath, err)
	}

	var results map[string]*BenchmarkResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, fmt.Errorf("failed to unmarshal benchmark results: %w", err)
	}

	return results, nil
}

// CompareResults compares current results with historical results and detects regressions
func CompareResults(current, historical map[string]*BenchmarkResult, thresholdPct float64) []RegressionAlert {
	var alerts []RegressionAlert

	for name, currentResult := range current {
		historicalResult, exists := historical[name]
		if !exists {
			continue // No historical data to compare
		}

		durationChange := float64(currentResult.Duration-historicalResult.Duration) / float64(historicalResult.Duration) * 100
		memChange := float64(currentResult.MemAllocBytes-historicalResult.MemAllocBytes) / float64(historicalResult.MemAllocBytes) * 100

		if durationChange > thresholdPct {
			alerts = append(alerts, RegressionAlert{
				BenchmarkName:   name,
				Metric:          "duration",
				ChangePercent:   durationChange,
				CurrentValue:    currentResult.Duration.String(),
				HistoricalValue: historicalResult.Duration.String(),
			})
		}

		if memChange > thresholdPct {
			alerts = append(alerts, RegressionAlert{
				BenchmarkName:   name,
				Metric:          "memory",
				ChangePercent:   memChange,
				CurrentValue:    fmt.Sprintf("%d bytes", currentResult.MemAllocBytes),
				HistoricalValue: fmt.Sprintf("%d bytes", historicalResult.MemAllocBytes),
			})
		}
	}

	return alerts
}

// RegressionAlert represents a performance regression detection
type RegressionAlert struct {
	BenchmarkName   string  `json:"benchmarkName"`
	Metric          string  `json:"metric"`
	ChangePercent   float64 `json:"changePercent"`
	CurrentValue    string  `json:"currentValue"`
	HistoricalValue string  `json:"historicalValue"`
}

// getGitCommit attempts to get the current git commit hash
func getGitCommit() string {
	// This would typically run git rev-parse HEAD
	// For simplicity, returning empty string - can be enhanced
	return ""
}

// PrintResults prints benchmark results in a human-readable format
func PrintResults(results map[string]*BenchmarkResult) {
	fmt.Printf("🚀 Benchmark Results\n\n")

	for name, result := range results {
		fmt.Printf("📊 %s:\n", name)
		fmt.Printf("  Duration: %v\n", result.Duration)
		fmt.Printf("  Memory Allocated: %d bytes (%d objects)\n", result.MemAllocBytes, result.MemAllocObjs)
		fmt.Printf("  Timestamp: %s\n", result.Timestamp.Format("2006-01-02 15:04:05"))
		if result.GitCommit != "" {
			fmt.Printf("  Git Commit: %s\n", result.GitCommit)
		}
		fmt.Printf("  Go Version: %s\n\n", result.GoVersion)
	}
}

// PrintRegressionAlerts prints performance regression alerts
func PrintRegressionAlerts(alerts []RegressionAlert) {
	if len(alerts) == 0 {
		fmt.Printf("✅ No performance regressions detected\n")
		return
	}

	fmt.Printf("⚠️  Performance Regression Alerts (%d)\n\n", len(alerts))

	for _, alert := range alerts {
		fmt.Printf("🔴 %s (%s):\n", alert.BenchmarkName, alert.Metric)
		fmt.Printf("  Change: %.1f%%\n", alert.ChangePercent)
		fmt.Printf("  Current: %s\n", alert.CurrentValue)
		fmt.Printf("  Historical: %s\n\n", alert.HistoricalValue)
	}
}
