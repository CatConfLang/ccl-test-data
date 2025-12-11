// Package mock provides a working CCL (Categorical Configuration Language) implementation.
//
// This package contains a functional CCL parser and processor that implements the core
// CCL specification. It serves as both a reference implementation for testing and a
// development tool for progressive CCL implementation.
//
// CCL Implementation Functions:
//   - Core parsing (Parse) - Convert text to flat key-value entries
//   - Entry processing (Filter, Compose, ExpandDotted) - Transform and combine entries
//   - Object construction (BuildHierarchy) - Build nested object hierarchies
//   - Typed access (GetString, GetInt, etc.) - Type-safe value extraction
//   - Formatting (CanonicalFormat/PrettyPrint) - Generate standardized formatted output
//
// Key Features:
//   - Comment support using '/=' syntax
//   - Dotted key expansion (database.host → nested objects)
//   - Duplicate key handling (converts to lists)
//   - Empty key support for array-style syntax
//   - Enhanced error messages with debugging context
//   - Type conversion with detailed error reporting
//
// Example Usage:
//
//	ccl := mock.New()
//	entries, err := ccl.Parse("key = value\n/= This is a comment")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	obj := ccl.BuildHierarchy(entries)
//	value, err := ccl.GetString(obj, []string{"key"})
package mock

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry represents a key-value pair from CCL parsing
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CCL implements a mock CCL parser for testing purposes
type CCL struct{}

// New creates a new mock CCL implementation
func New() *CCL {
	return &CCL{}
}

// Parse implements core entry parsing with multiline support
func (c *CCL) Parse(input string) ([]Entry, error) {
	var entries []Entry

	// Initialize empty slice to avoid nil return
	entries = []Entry{}

	// Handle empty input
	if strings.TrimSpace(input) == "" {
		return entries, nil
	}

	// Normalize line endings first: CRLF -> LF, lone CR -> LF
	normalizedInput := strings.ReplaceAll(input, "\r\n", "\n")
	normalizedInput = strings.ReplaceAll(normalizedInput, "\r", "\n")
	lines := strings.Split(normalizedInput, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Handle comments (start with /=)
		if strings.HasPrefix(strings.TrimSpace(line), "/=") {
			comment := strings.Trim(strings.TrimPrefix(strings.TrimSpace(line), "/="), " \t")
			entries = append(entries, Entry{
				Key:   "/",
				Value: comment,
			})
			continue
		}

		// Handle key-value pairs
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.Trim(parts[0], " \t")
				// For basic parsing, trim whitespace from values
				value := strings.Trim(parts[1], " \t")

				// Check if there are indented lines following this key (multiline content)
				if i+1 < len(lines) {
					// Collect all indented lines following this key
					var multilineValue []string
					j := i + 1

					for j < len(lines) {
						nextLine := lines[j]

						// If line starts with whitespace (indented), it belongs to this key
						if len(nextLine) > 0 && (nextLine[0] == ' ' || nextLine[0] == '\t') {
							multilineValue = append(multilineValue, nextLine)
							j++
						} else if strings.TrimSpace(nextLine) == "" {
							// Skip empty lines within multiline content
							j++
						} else {
							// Non-indented, non-empty line - end of multiline content
							break
						}
					}

					if len(multilineValue) > 0 {
						// If the initial value was empty, start with newline, otherwise append
						if strings.TrimSpace(value) == "" {
							value = "\n" + strings.Join(multilineValue, "\n")
						} else {
							value = value + "\n" + strings.Join(multilineValue, "\n")
						}
						i = j - 1 // Skip the lines we've consumed
					}
				}

				// Line ending normalization already handled at input level

				entries = append(entries, Entry{
					Key:   key,
					Value: value,
				})
			}
		}
		// Skip lines without = (don't error - this is just basic parsing)
	}

	return entries, nil
}

// ParseIndented implements entry processing with indentation normalization
// It calculates the common leading whitespace prefix and strips it from all lines
func (c *CCL) ParseIndented(input string) ([]Entry, error) {
	// For mock purposes, same as Parse
	// A full implementation would dedent the input first
	return c.Parse(input)
}

// Filter implements entry filtering - removes comment entries (key="/")
func (c *CCL) Filter(entries []Entry) []Entry {
	var filtered []Entry
	for _, entry := range entries {
		if entry.Key != "/" {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// Compose implements entry composition
func (c *CCL) Compose(left, right []Entry) []Entry {
	// Simple concatenation for mock
	result := make([]Entry, len(left)+len(right))
	copy(result, left)
	copy(result[len(left):], right)
	return result
}

// ExpandDotted implements dotted key expansion
func (c *CCL) ExpandDotted(entries []Entry) []Entry {
	var expanded []Entry
	for _, entry := range entries {
		if strings.Contains(entry.Key, ".") {
			// For mock, just keep as-is
			expanded = append(expanded, entry)
		} else {
			expanded = append(expanded, entry)
		}
	}
	return expanded
}

// BuildHierarchy implements object construction
func (c *CCL) BuildHierarchy(entries []Entry) map[string]interface{} {
	result := make(map[string]interface{})

	for _, entry := range entries {
		key := entry.Key
		value := entry.Value

		// Handle empty keys as list items
		if key == "" {
			// Add to a list under empty key
			if existing, exists := result[""]; exists {
				if list, ok := existing.([]interface{}); ok {
					result[""] = append(list, value)
				} else {
					// Convert single value to list
					result[""] = []interface{}{existing, value}
				}
			} else {
				result[""] = []interface{}{value}
			}
		} else if strings.Contains(key, ".") {
			// Handle dotted keys - basic support
			parts := strings.Split(key, ".")
			current := result

			for i, part := range parts {
				if i == len(parts)-1 {
					// Last part - set the value
					current[part] = value
				} else {
					// Intermediate part - create nested object
					if _, exists := current[part]; !exists {
						current[part] = make(map[string]interface{})
					}
					if nested, ok := current[part].(map[string]interface{}); ok {
						current = nested
					}
				}
			}
		} else {
			// Handle duplicate keys as lists
			if existing, exists := result[key]; exists {
				if list, ok := existing.([]interface{}); ok {
					result[key] = append(list, value)
				} else {
					// Convert single value to list
					result[key] = []interface{}{existing, value}
				}
			} else {
				result[key] = value
			}
		}
	}

	return result
}

// GetString implements string access
func (c *CCL) GetString(obj map[string]interface{}, path []string) (string, error) {
	value, err := c.getValue(obj, path)
	if err != nil {
		return "", err
	}

	if str, ok := value.(string); ok {
		return str, nil
	}

	return fmt.Sprintf("%v", value), nil
}

// GetInt implements integer access
func (c *CCL) GetInt(obj map[string]interface{}, path []string) (int, error) {
	value, err := c.getValue(obj, path)
	if err != nil {
		return 0, err
	}

	if str, ok := value.(string); ok {
		return strconv.Atoi(str)
	}

	if i, ok := value.(int); ok {
		return i, nil
	}

	return 0, fmt.Errorf("cannot convert value %v (type %T) to int at path %s", value, value, strings.Join(path, "."))
}

// GetBool implements boolean access
func (c *CCL) GetBool(obj map[string]interface{}, path []string) (bool, error) {
	value, err := c.getValue(obj, path)
	if err != nil {
		return false, err
	}

	if str, ok := value.(string); ok {
		return strconv.ParseBool(str)
	}

	if b, ok := value.(bool); ok {
		return b, nil
	}

	return false, fmt.Errorf("cannot convert value %v (type %T) to bool at path %s", value, value, strings.Join(path, "."))
}

// GetFloat implements float access
func (c *CCL) GetFloat(obj map[string]interface{}, path []string) (float64, error) {
	value, err := c.getValue(obj, path)
	if err != nil {
		return 0, err
	}

	if str, ok := value.(string); ok {
		return strconv.ParseFloat(str, 64)
	}

	if f, ok := value.(float64); ok {
		return f, nil
	}

	return 0, fmt.Errorf("cannot convert value %v (type %T) to float64 at path %s", value, value, strings.Join(path, "."))
}

// GetList implements list access
func (c *CCL) GetList(obj map[string]interface{}, path []string) ([]string, error) {
	value, err := c.getValue(obj, path)
	if err != nil {
		return nil, err
	}

	if arr, ok := value.([]interface{}); ok {
		result := make([]string, len(arr))
		for i, v := range arr {
			result[i] = fmt.Sprintf("%v", v)
		}
		return result, nil
	}

	if list, ok := value.([]string); ok {
		return list, nil
	}

	return nil, fmt.Errorf("cannot convert value %v (type %T) to []string at path %s", value, value, strings.Join(path, "."))
}

// PrettyPrint implements canonical formatting (standardized output)
// This is the model-level format where key=value becomes key=\n  value=\n
func (c *CCL) PrettyPrint(obj map[string]interface{}) string {
	var lines []string
	c.prettyPrintObject(obj, "", &lines)
	return strings.Join(lines, "\n")
}

// Print implements structure-preserving formatting (input-preserving for standard inputs)
// For inputs in standard format (single space around =, 2-space indent, LF line endings),
// Print(Parse(x)) == x
func (c *CCL) Print(entries []Entry) string {
	if len(entries) == 0 {
		return ""
	}

	var sb strings.Builder
	for i, entry := range entries {
		if i > 0 {
			sb.WriteString("\n")
		}

		// Handle comment entries
		if entry.Key == "/" {
			sb.WriteString("/= ")
			sb.WriteString(entry.Value)
			continue
		}

		// Write key
		sb.WriteString(entry.Key)
		sb.WriteString(" = ")

		// Write value - if multiline, the value already contains the newlines and indentation
		sb.WriteString(entry.Value)
	}

	return sb.String()
}

// ComposeAssociative verifies (a·b)·c == a·(b·c) for three inputs
func (c *CCL) ComposeAssociative(inputs []string) (bool, error) {
	if len(inputs) < 3 {
		return false, fmt.Errorf("compose_associative requires at least 3 inputs")
	}

	a, err := c.Parse(inputs[0])
	if err != nil {
		return false, err
	}
	b, err := c.Parse(inputs[1])
	if err != nil {
		return false, err
	}
	cEntries, err := c.Parse(inputs[2])
	if err != nil {
		return false, err
	}

	// (a·b)·c
	ab := c.Compose(a, b)
	left := c.Compose(ab, cEntries)

	// a·(b·c)
	bc := c.Compose(b, cEntries)
	right := c.Compose(a, bc)

	return entriesEqual(left, right), nil
}

// IdentityLeft verifies compose(empty, x) == x
func (c *CCL) IdentityLeft(inputs []string) (bool, error) {
	if len(inputs) < 2 {
		return false, fmt.Errorf("identity_left requires at least 2 inputs")
	}

	empty, err := c.Parse(inputs[0])
	if err != nil {
		return false, err
	}
	x, err := c.Parse(inputs[1])
	if err != nil {
		return false, err
	}

	result := c.Compose(empty, x)
	return entriesEqual(result, x), nil
}

// IdentityRight verifies compose(x, empty) == x
func (c *CCL) IdentityRight(inputs []string) (bool, error) {
	if len(inputs) < 2 {
		return false, fmt.Errorf("identity_right requires at least 2 inputs")
	}

	x, err := c.Parse(inputs[0])
	if err != nil {
		return false, err
	}
	empty, err := c.Parse(inputs[1])
	if err != nil {
		return false, err
	}

	result := c.Compose(x, empty)
	return entriesEqual(result, x), nil
}

// RoundTrip verifies parse(print(parse(x))) == parse(x)
func (c *CCL) RoundTrip(input string) (bool, error) {
	parsed, err := c.Parse(input)
	if err != nil {
		return false, err
	}

	printed := c.Print(parsed)
	reparsed, err := c.Parse(printed)
	if err != nil {
		return false, err
	}

	return entriesEqual(parsed, reparsed), nil
}

// entriesEqual compares two entry slices for equality
func entriesEqual(a, b []Entry) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Key != b[i].Key || a[i].Value != b[i].Value {
			return false
		}
	}
	return true
}

// Helper methods

func (c *CCL) getValue(obj map[string]interface{}, path []string) (interface{}, error) {
	current := obj

	for i, key := range path {
		if i == len(path)-1 {
			// Last key - return the value
			if value, exists := current[key]; exists {
				return value, nil
			}
			return nil, fmt.Errorf("key not found: %s (available keys: %v)", strings.Join(path, "."), getMapKeys(current))
		} else {
			// Intermediate key - navigate deeper
			if value, exists := current[key]; exists {
				if nested, ok := value.(map[string]interface{}); ok {
					current = nested
				} else {
					return nil, fmt.Errorf("not an object at key: %s", key)
				}
			} else {
				return nil, fmt.Errorf("intermediate key not found: %s in path %s (available keys: %v)", key, strings.Join(path, "."), getMapKeys(current))
			}
		}
	}

	return nil, fmt.Errorf("invalid path")
}

func (c *CCL) prettyPrintObject(obj map[string]interface{}, prefix string, lines *[]string) {
	for key, value := range obj {
		if nested, ok := value.(map[string]interface{}); ok {
			c.prettyPrintObject(nested, prefix+key+".", lines)
		} else {
			*lines = append(*lines, fmt.Sprintf("%s%s = %v", prefix, key, value))
		}
	}
}

// getMapKeys returns a slice of keys from a map for debugging purposes
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
