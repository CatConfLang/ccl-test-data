package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path1> [path2] [...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Recursively removes files and directories (cross-platform rm -rf)\n")
		os.Exit(1)
	}

	for _, path := range os.Args[1:] {
		if err := os.RemoveAll(path); err != nil {
			fmt.Fprintf(os.Stderr, "Error removing %s: %v\n", path, err)
			os.Exit(1)
		}
	}
}