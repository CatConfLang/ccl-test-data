// validate-config is a simple tool to validate CCL YAML configuration files
package main

import (
	"fmt"
	"os"

	"github.com/ccl-test-data/test-runner/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config.yaml>\n", os.Args[0])
		os.Exit(1)
	}

	configPath := os.Args[1]

	config, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Configuration is valid\n")
	fmt.Printf("Functions: %v\n", config.Functions)
	if len(config.Features) > 0 {
		fmt.Printf("Features: %v\n", config.Features)
	}
	if len(config.Behaviors) > 0 {
		fmt.Printf("Behaviors: %v\n", config.Behaviors)
	}
	if len(config.Variants) > 0 {
		fmt.Printf("Variants: %v\n", config.Variants)
	}
	if len(config.SkipTests) > 0 {
		fmt.Printf("Skip tests: %v\n", config.SkipTests)
	}
}