package types

import "encoding/json"

// TestSuite represents the root structure of a CCL test file
type TestSuite struct {
	Suite       string     `json:"suite"`
	Version     string     `json:"version"`
	Description string     `json:"description,omitempty"`
	Tests       []TestCase `json:"tests"`
}

// TestCase represents a single test case
type TestCase struct {
	Name        string        `json:"name"`
	Input       string        `json:"input,omitempty"`
	Input1      string        `json:"input1,omitempty"`
	Input2      string        `json:"input2,omitempty"`
	Input3      string        `json:"input3,omitempty"`
	Validations ValidationSet `json:"validations"`
	Meta        TestMetadata  `json:"meta"`
}

// ValidationSet contains all possible validation types
// Using interface{} to handle the flexible oneOf schema structures
type ValidationSet struct {
	Parse          interface{} `json:"parse,omitempty"`
	ParseValue     interface{} `json:"parse_value,omitempty"`
	Filter         interface{} `json:"filter,omitempty"`
	Combine        interface{} `json:"combine,omitempty"`
	ExpandDotted   interface{} `json:"expand_dotted,omitempty"`
	BuildHierarchy interface{} `json:"build_hierarchy,omitempty"`
	GetString      interface{} `json:"get_string,omitempty"`
	GetInt         interface{} `json:"get_int,omitempty"`
	GetBool        interface{} `json:"get_bool,omitempty"`
	GetFloat       interface{} `json:"get_float,omitempty"`
	GetList        interface{} `json:"get_list,omitempty"`
	PrettyPrint    interface{} `json:"pretty_print,omitempty"`
	RoundTrip      interface{} `json:"round_trip,omitempty"`
	Canonical      interface{} `json:"canonical_format,omitempty"`
	Associativity  interface{} `json:"associativity,omitempty"`
}

// TestMetadata contains test categorization information
type TestMetadata struct {
	Tags       []string `json:"tags"`
	Conflicts  []string `json:"conflicts,omitempty"`
	Level      int      `json:"level"`
	Feature    string   `json:"feature,omitempty"`
	Difficulty string   `json:"difficulty,omitempty"`
}

// Entry represents a key-value pair from CCL parsing
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Validation types with support for both legacy and counted formats
// Using json.RawMessage to handle flexible schema structures

type ParseValidation json.RawMessage
type FilterValidation json.RawMessage
type ExpandDottedValidation json.RawMessage
type CombineValidation json.RawMessage
type BuildHierarchyValidation json.RawMessage
type TypedAccessValidation json.RawMessage
type PrettyPrintValidation json.RawMessage
type RoundTripValidation json.RawMessage
type CanonicalValidation json.RawMessage
type AssociativityValidation json.RawMessage
