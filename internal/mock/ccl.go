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

// Compose implements Level 2: Entry composition
func (c *CCL) Compose(left, right []Entry) []Entry {
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

// MakeObjects implements Level 3: Object construction
func (c *CCL) MakeObjects(entries []Entry) map[string]interface{} {
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

	return 0, fmt.Errorf("cannot convert %v to int", value)
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

	return false, fmt.Errorf("cannot convert %v to bool", value)
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

	return 0, fmt.Errorf("cannot convert %v to float64", value)
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

	return nil, fmt.Errorf("cannot convert %v to []string", value)
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
			return nil, fmt.Errorf("key not found: %s", strings.Join(path, "."))
		} else {
			// Intermediate key - navigate deeper
			if value, exists := current[key]; exists {
				if nested, ok := value.(map[string]interface{}); ok {
					current = nested
				} else {
					return nil, fmt.Errorf("not an object at key: %s", key)
				}
			} else {
				return nil, fmt.Errorf("key not found: %s", key)
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
