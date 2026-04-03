package main

import (
	"fmt"
	"sort"
)

// renderObjectValue renders a single value within an object (string, nested object, or array).
// isLast indicates if this is the last item in the parent container.
// prefix is the accumulated branch characters from parent levels.
func renderObjectValue(key string, value interface{}, prefix string, isLast bool, lines *[]string) {
	branch := "├─ "
	childPrefix := prefix + "│  "
	if isLast {
		branch = "└─ "
		childPrefix = prefix + "   "
	}

	switch v := value.(type) {
	case string:
		line := prefix + objectBranchStyle.Render(branch) +
			objectKeyStyle.Render(key) + " " +
			entryEqualsStyle.Render("=") + " " +
			objectValueStyle.Render(`"`+v+`"`)
		*lines = append(*lines, line)

	case map[string]interface{}:
		line := prefix + objectBranchStyle.Render(branch) + objectKeyStyle.Render(key)
		*lines = append(*lines, line)
		renderObjectLines(v, childPrefix, lines)

	case []interface{}:
		line := prefix + objectBranchStyle.Render(branch) + objectKeyStyle.Render(key)
		*lines = append(*lines, line)
		renderArrayLines(v, childPrefix, lines)

	default:
		line := prefix + objectBranchStyle.Render(branch) +
			objectKeyStyle.Render(key) + " " +
			entryEqualsStyle.Render("=") + " " +
			objectValueStyle.Render(fmt.Sprintf("%v", v))
		*lines = append(*lines, line)
	}
}

// renderArrayLines renders array items with index indicators
func renderArrayLines(arr []interface{}, prefix string, lines *[]string) {
	for i, item := range arr {
		isLast := i == len(arr)-1
		branch := "├─ "
		childPrefix := prefix + "│  "
		if isLast {
			branch = "└─ "
			childPrefix = prefix + "   "
		}

		indexStr := arrayIndexStyle.Render(fmt.Sprintf("[%d]", i))

		switch v := item.(type) {
		case string:
			line := prefix + objectBranchStyle.Render(branch) + indexStr + " " +
				objectValueStyle.Render(`"`+v+`"`)
			*lines = append(*lines, line)

		case map[string]interface{}:
			line := prefix + objectBranchStyle.Render(branch) + indexStr
			*lines = append(*lines, line)
			renderObjectLines(v, childPrefix, lines)

		default:
			line := prefix + objectBranchStyle.Render(branch) + indexStr + " " +
				objectValueStyle.Render(fmt.Sprintf("%v", v))
			*lines = append(*lines, line)
		}
	}
}

// renderObjectLines renders an object's key-value pairs with tree structure
func renderObjectLines(obj map[string]interface{}, prefix string, lines *[]string) {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, key := range keys {
		isLast := i == len(keys)-1
		renderObjectValue(key, obj[key], prefix, isLast, lines)
	}
}

// renderObject renders a complete object as a tree structure.
// Returns a slice of lines that can be displayed with scrolling.
func renderObject(obj map[string]interface{}) []string {
	var lines []string
	renderObjectLines(obj, "", &lines)
	return lines
}
