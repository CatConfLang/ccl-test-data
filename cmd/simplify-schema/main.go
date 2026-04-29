// simplify-schema strips JSON Schema constructs that go-jsonschema (omissis fork)
// can't translate into useful Go types, producing a simplified schema suitable
// for type generation. Runtime validation continues to use the original schema.
//
// Two transformations:
//  1. Drop `properties.tests.items.allOf` (contains an `if/then` conditional that
//     go-jsonschema ignores anyway, but its presence collapses the items type to
//     interface{}).
//  2. Replace `properties.tests.items.properties.features.items` (which uses a
//     `oneOf` of [enum, ^experimental_, ^optional_] patterns) with a plain
//     `{type: string}`. go-jsonschema can't synthesize a single Go type for
//     "enum or two regex patterns", so the simplified version produces a plain
//     []string field — runtime validation still enforces the patterns.
//
// Usage: simplify-schema <input.json> <output.json>
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: simplify-schema <input.json> <output.json>")
		os.Exit(2)
	}

	in, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input: %v\n", err)
		os.Exit(1)
	}

	var schema map[string]any
	if err := json.Unmarshal(in, &schema); err != nil {
		fmt.Fprintf(os.Stderr, "parse input: %v\n", err)
		os.Exit(1)
	}

	items := nav(schema, "properties", "tests", "items")
	if items == nil {
		fmt.Fprintln(os.Stderr, "schema missing properties.tests.items")
		os.Exit(1)
	}
	delete(items, "allOf")

	if featItems := nav(items, "properties", "features", "items"); featItems != nil {
		delete(featItems, "oneOf")
	}

	out, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(os.Args[2], append(out, '\n'), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write output: %v\n", err)
		os.Exit(1)
	}
}

func nav(m map[string]any, keys ...string) map[string]any {
	for _, k := range keys {
		next, ok := m[k].(map[string]any)
		if !ok {
			return nil
		}
		m = next
	}
	return m
}
