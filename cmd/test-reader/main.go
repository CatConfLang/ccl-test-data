// CCL Test Reader - Interactive viewer for CCL test files
//
// Designed for generated format files (generated_tests/) with split-view layout.
// Source format files (source_tests/) have limited display support.
//
// The test-reader displays CCL test cases with their input, expected output,
// and metadata. Features a split-view TUI with filterable test list and detail pane.
package main

import (
"fmt"
"log"
"os"
)

func main() {
if len(os.Args) < 2 {
fmt.Println("Usage: test-reader <test-file.json|directory> [--static]")
fmt.Println("       test-reader tests/                              # Interactive TUI (default)")
fmt.Println("       test-reader tests/api_essential-parsing.json   # Interactive TUI (default)")
fmt.Println("       test-reader tests/ --static                     # Static CLI output")
fmt.Println("       test-reader tests/api_essential-parsing.json --static")
os.Exit(1)
}

path := os.Args[1]
useStatic := len(os.Args) > 2 && os.Args[2] == "--static"

if info, err := os.Stat(path); err == nil && info.IsDir() {
if useStatic {
runFileSelectionCLI(path)
} else {
runFileSelectionTUI(path)
}
} else {
if useStatic {
if err := processTestFile(path); err != nil {
log.Printf("Error processing %s: %v", path, err)
}
} else {
runTUI(path)
}
}
}
