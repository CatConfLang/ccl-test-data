package main

import (
	"fmt"
	"strings"
)

// visualizeWhitespaceInline shows spaces as dots and tabs as arrows
func visualizeWhitespaceInline(s string) string {
	result := strings.ReplaceAll(s, " ", "·")
	result = strings.ReplaceAll(result, "\t", "→")
	return result
}

// formatKey handles empty keys with special indicator
func formatKey(key string) string {
	if key == "" {
		return emptyKeyStyle.Render("(empty-key)")
	}
	return entryKeyStyle.Render(visualizeWhitespaceInline(key))
}

// formatValue handles empty values with special indicator
func formatValue(value string) string {
	if value == "" {
		return emptyValueStyle.Render("(empty)")
	}
	return entryValueStyle.Render(visualizeWhitespaceInline(value))
}

// formatInputContent shows input with whitespace indicators
func formatInputContent(input string) string {
	if input == "" {
		return "(empty)"
	}
	return visualizeWhitespaceInline(input)
}

// isObjectValidation returns true if the validation type expects an object (not entries)
func isObjectValidation(validation string) bool {
	switch validation {
	case "build_hierarchy", "load":
		return true
	default:
		return false
	}
}

// isListValidation returns true if the validation type expects a list
func isListValidation(validation string) bool {
	return validation == "get_list"
}

// ExpectedContentResult holds the rendered expected output content
type ExpectedContentResult struct {
	Header       string   // e.g. "✅ EXPECTED: Object" or "✅ EXPECTED: Entries"
	IsError      bool     // true if expecting an error
	IsObject     bool     // true if object visualization, false for entries
	IsList       bool     // true if list visualization (get_list)
	ListKey      string   // the key argument for list lookups
	Lines        []string // rendered lines (object tree lines or entry box strings)
	TotalItems   int      // total number of items (for count display)
	StartIdx     int      // actual start index after bounds checking
	EndIdx       int      // actual end index after bounds checking
	HasMoreAbove bool     // true if there are items above the visible range
	HasMoreBelow bool     // true if there are items below the visible range
}

// extractEntries extracts Entry slice from test.Expected, handling both
// []interface{} (loader-transformed arrays) and map[string]interface{}
// (structured format with count/entries).
func extractEntries(expected interface{}) []Entry {
	var entries []Entry
	if expected == nil {
		return entries
	}

	switch exp := expected.(type) {
	case []interface{}:
		for _, item := range exp {
			if entryMap, ok := item.(map[string]interface{}); ok {
				key, _ := entryMap["key"].(string)
				value, _ := entryMap["value"].(string)
				entries = append(entries, Entry{Key: key, Value: value})
			}
		}
	case map[string]interface{}:
		if entriesArray, ok := exp["entries"].([]interface{}); ok {
			for _, item := range entriesArray {
				if entryMap, ok := item.(map[string]interface{}); ok {
					key, _ := entryMap["key"].(string)
					value, _ := entryMap["value"].(string)
					entries = append(entries, Entry{Key: key, Value: value})
				}
			}
		}
	}
	return entries
}

// renderEntryLine renders a single entry as a styled string
func renderEntryLine(entry Entry) string {
	if strings.Contains(entry.Value, "\n") {
		keyLine := fmt.Sprintf("%s %s", formatKey(entry.Key), entryEqualsStyle.Render("="))
		valueLine := formatValue(entry.Value)
		return entryBoxStyle.Render(fmt.Sprintf("%s\n%s", keyLine, valueLine))
	}
	entryContent := fmt.Sprintf("%s %s %s", formatKey(entry.Key), entryEqualsStyle.Render("="), formatValue(entry.Value))
	return entryBoxStyle.Render(entryContent)
}

// extractList extracts a string slice from test.Expected for get_list validation.
// Handles both the loader-transformed format (bare []interface{}) and the
// structured format ({"count": N, "list": ["a", "b", ...]}).
func extractList(expected interface{}) []string {
	if expected == nil {
		return nil
	}

	// Handle already-transformed format (loader extracts just the array)
	if arr, ok := expected.([]interface{}); ok {
		items := make([]string, 0, len(arr))
		for _, item := range arr {
			items = append(items, fmt.Sprintf("%v", item))
		}
		return items
	}

	// Handle structured format with "list" field
	exp, ok := expected.(map[string]interface{})
	if !ok {
		return nil
	}
	listRaw, ok := exp["list"]
	if !ok {
		return nil
	}
	arr, ok := listRaw.([]interface{})
	if !ok {
		return nil
	}
	items := make([]string, 0, len(arr))
	for _, item := range arr {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return items
}

// renderListItemLine renders a single list item as a styled string with its index
func renderListItemLine(index int, value string) string {
	indexStr := listIndexStyle.Render(fmt.Sprintf("[%d]", index))
	valueStr := listValueStyle.Render(fmt.Sprintf("%q", value))
	return fmt.Sprintf("  %s %s", indexStr, valueStr)
}

// countExpectedItems returns the total number of scrollable items for a test's
// expected output without performing full styled rendering.
func countExpectedItems(test TestCase) int {
	if test.ExpectError {
		return 0
	}
	if isListValidation(test.Validation) {
		return len(extractList(test.Expected))
	}
	if isObjectValidation(test.Validation) {
		if obj, ok := test.Expected.(map[string]interface{}); ok {
			return len(renderObject(obj))
		}
		return 0
	}
	return len(extractEntries(test.Expected))
}

// renderExpectedContent generates the expected output content for both static and TUI modes.
// scrollOffset is the starting index for display (0 for static mode).
// maxDisplay is the maximum items to show.
func renderExpectedContent(test TestCase, scrollOffset int, maxDisplay int) ExpectedContentResult {
	result := ExpectedContentResult{}

	if test.ExpectError {
		result.Header = "❌ EXPECTED: Error"
		result.IsError = true
		return result
	}

	if isListValidation(test.Validation) {
		result.IsList = true
		if len(test.Args) > 0 {
			result.ListKey = test.Args[0]
		}
		items := extractList(test.Expected)
		result.TotalItems = len(items)
		if result.ListKey != "" {
			result.Header = fmt.Sprintf("✅ EXPECTED: List (key: %s)", listKeyArgStyle.Render(result.ListKey))
		} else {
			result.Header = "✅ EXPECTED: List"
		}
		for i, item := range items {
			result.Lines = append(result.Lines, renderListItemLine(i, item))
		}
		applyScrollBounds(&result, scrollOffset, maxDisplay)
		return result
	}

	if isObjectValidation(test.Validation) {
		if obj, ok := test.Expected.(map[string]interface{}); ok {
			result.Header = "✅ EXPECTED: Object"
			result.IsObject = true
			result.Lines = renderObject(obj)
			result.TotalItems = len(result.Lines)
			applyScrollBounds(&result, scrollOffset, maxDisplay)
			return result
		}
	}

	result.Header = "✅ EXPECTED: Entries"
	result.IsObject = false

	entries := extractEntries(test.Expected)
	result.TotalItems = len(entries)
	for _, entry := range entries {
		result.Lines = append(result.Lines, renderEntryLine(entry))
	}

	applyScrollBounds(&result, scrollOffset, maxDisplay)
	return result
}

// applyScrollBounds calculates the visible range with bounds checking
func applyScrollBounds(result *ExpectedContentResult, scrollOffset int, maxDisplay int) {
	result.StartIdx = scrollOffset
	if result.StartIdx >= result.TotalItems {
		result.StartIdx = result.TotalItems - 1
		if result.StartIdx < 0 {
			result.StartIdx = 0
		}
	}

	result.EndIdx = result.StartIdx + maxDisplay
	if result.EndIdx > result.TotalItems {
		result.EndIdx = result.TotalItems
	}

	result.HasMoreAbove = result.StartIdx > 0
	result.HasMoreBelow = result.EndIdx < result.TotalItems
}
