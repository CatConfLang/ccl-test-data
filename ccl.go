// Package ccl_test_data provides shared CCL test infrastructure
// for loading, filtering, and generating test suites across implementations.
package ccl_test_data

import (
	"github.com/tylerbutler/ccl-test-data/config"
	"github.com/tylerbutler/ccl-test-data/loader"
	"github.com/tylerbutler/ccl-test-data/types"
)

// Version of the package
const Version = "v0.1.0"

// NewLoader creates a test loader with sensible defaults
func NewLoader(testDataPath string, cfg config.ImplementationConfig) *loader.TestLoader {
	return loader.NewTestLoader(testDataPath, cfg)
}

// LoadCompatibleTests is a convenience function for the most common use case
func LoadCompatibleTests(testDataPath string, cfg config.ImplementationConfig) ([]types.TestCase, error) {
	testLoader := NewLoader(testDataPath, cfg)
	return testLoader.LoadAllTests(loader.LoadOptions{
		Format:     loader.FormatFlat,
		FilterMode: loader.FilterCompatible,
	})
}

// GetTestStats provides quick statistics for a test set
func GetTestStats(testDataPath string, cfg config.ImplementationConfig) (types.TestStatistics, error) {
	testLoader := NewLoader(testDataPath, cfg)
	tests, err := testLoader.LoadAllTests(loader.LoadOptions{
		Format:     loader.FormatFlat,
		FilterMode: loader.FilterAll,
	})
	if err != nil {
		return types.TestStatistics{}, err
	}
	return testLoader.GetTestStatistics(tests), nil
}
