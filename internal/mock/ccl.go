// Package mock provides a working CCL (Categorical Configuration Language) implementation.
//
// This package contains a functional CCL parser and processor that implements the core
// CCL specification across multiple levels. It serves as both a reference implementation
// for testing and a development tool for progressive CCL implementation.
//
// CCL Implementation Levels:
//   - Level 1: Raw parsing (Parse) - Convert text to flat key-value entries
//   - Level 2: Entry processing (Filter, Combine, ExpandDotted) - Transform and combine entries
//   - Level 3: Object construction (BuildHierarchy) - Build nested object hierarchies
//   - Level 4: Typed access (GetString, GetInt, etc.) - Type-safe value extraction
//   - Level 5: Formatting (PrettyPrint) - Generate formatted output
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

// Parse implements Level 1: Raw entry parsing
func (c *CCL) Parse(input string) ([]Entry, error) {
	var entries []Entry

	// Handle empty input
	if strings.TrimSpace(input) == "" {
		return []Entry{}, nil
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Preserve \r characters when trimming - only trim leading/trailing spaces and tabs
		originalLine := line
		line = strings.Trim(line, " \t")
		if line == "" {
			continue
		}

		// Handle comments (start with /=)
		if strings.HasPrefix(line, "/=") {
			comment := strings.Trim(strings.TrimPrefix(line, "/="), " \t")
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
				value := strings.Trim(parts[1], " \t")
				// If original line had \r at the end, preserve it in the value
				if strings.HasSuffix(originalLine, "\r") && !strings.HasSuffix(value, "\r") {
					value = value + "\r"
				}
				entries = append(entries, Entry{
					Key:   key,
					Value: value,
				})
			}
		}
		// Skip lines without = (don't error - Level 1 is just basic parsing)
	}

	return entries, nil
}

// ParseValue implements Level 2: Entry processing with indentation awareness
func (c *CCL) ParseValue(input string) ([]Entry, error) {
	// For mock purposes, same as Parse
	return c.Parse(input)
}

// Filter implements Level 2: Entry filtering (for comment tests, this preserves all entries)
func (c *CCL) Filter(entries []Entry) []Entry {
	// For the mock implementation, filter just returns all entries
	// In a real implementation, this might filter based on certain criteria
	return entries
}

// Combine implements Level 2: Entry composition
func (c *CCL) Combine(left, right []Entry) []Entry {
	// Simple concatenation for mock
	result := make([]Entry, len(left)+len(right))
	copy(result, left)
	copy(result[len(left):], right)
	return result
}

// ExpandDotted implements Level 2: Dotted key expansion
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

// BuildHierarchy implements Level 3: Object construction
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

// GetString implements Level 4: String access
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

// GetInt implements Level 4: Integer access
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

// GetBool implements Level 4: Boolean access
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

// GetFloat implements Level 4: Float access
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

// GetList implements Level 4: List access
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

// PrettyPrint implements Level 5: Pretty printing
func (c *CCL) PrettyPrint(obj map[string]interface{}) string {
	var lines []string
	c.prettyPrintObject(obj, "", &lines)
	return strings.Join(lines, "\n")
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
