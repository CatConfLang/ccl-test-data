// Package config provides simplified YAML-based configuration for the CCL test runner.
package config

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/catconflang/ccl-test-data/config"
	"gopkg.in/yaml.v3"
)

// SimpleConfig represents the simplified YAML configuration format.
// All fields are arrays of strings for maximum simplicity.
type SimpleConfig struct {
	Functions []string `yaml:"functions" json:"functions"`
	Features  []string `yaml:"features,omitempty" json:"features,omitempty"`
	Behaviors []string `yaml:"behaviors,omitempty" json:"behaviors,omitempty"`
	Variants  []string `yaml:"variants,omitempty" json:"variants,omitempty"`
	SkipTests []string `yaml:"skip_tests,omitempty" json:"skip_tests,omitempty"`
}

// ValidFunctions defines all supported CCL functions
var ValidFunctions = []string{
	"parse",
	"build_hierarchy",
	"get_string",
	"get_int",
	"get_bool",
	"get_float",
	"get_list",
	"filter",
	"expand_dotted",
	"canonical_format",
}

// ValidFeatures defines all supported CCL features
var ValidFeatures = []string{
	"comments",
	"experimental_dotted_keys",
	"unicode",
	"multiline",
	"whitespace",
	"empty_keys",
	"optional_typed_accessors",
}

// ValidBehaviors defines all supported behavioral choices
var ValidBehaviors = []string{
	"boolean_strict",
	"boolean_lenient",
	"tabs_as_content",
	"tabs_as_whitespace",
	"crlf_preserve_literal",
	"crlf_normalize_to_lf",
	"indent_spaces",
	"indent_tabs",
	"list_coercion_enabled",
	"list_coercion_disabled",
	"array_order_insertion",
	"array_order_lexicographic",
}

// ValidVariants defines all supported specification variants
var ValidVariants = []string{
	"proposed_behavior",
	"reference_compliant",
}

// ConflictingBehaviors defines mutually exclusive behavioral choices
var ConflictingBehaviors = [][]string{
	{"boolean_strict", "boolean_lenient"},
	{"tabs_as_content", "tabs_as_whitespace"},
	{"crlf_preserve_literal", "crlf_normalize_to_lf"},
	{"indent_spaces", "indent_tabs"},
	{"list_coercion_enabled", "list_coercion_disabled"},
	{"array_order_insertion", "array_order_lexicographic"},
}

// LoadConfig loads and validates a YAML configuration file
func LoadConfig(filePath string) (*SimpleConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config SimpleConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML config: %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

// Validate ensures the configuration is valid
func (c *SimpleConfig) Validate() error {
	var errors []string

	// Functions are required and must be valid
	if len(c.Functions) == 0 {
		errors = append(errors, "at least one function must be specified")
	}

	for _, fn := range c.Functions {
		if !slices.Contains(ValidFunctions, fn) {
			errors = append(errors, fmt.Sprintf("invalid function: %s (valid: %s)", fn, strings.Join(ValidFunctions, ", ")))
		}
	}

	// Validate features
	for _, feature := range c.Features {
		if !slices.Contains(ValidFeatures, feature) {
			errors = append(errors, fmt.Sprintf("invalid feature: %s (valid: %s)", feature, strings.Join(ValidFeatures, ", ")))
		}
	}

	// Validate behaviors
	for _, behavior := range c.Behaviors {
		if !slices.Contains(ValidBehaviors, behavior) {
			errors = append(errors, fmt.Sprintf("invalid behavior: %s (valid: %s)", behavior, strings.Join(ValidBehaviors, ", ")))
		}
	}

	// Check for conflicting behaviors
	for _, conflictGroup := range ConflictingBehaviors {
		var found []string
		for _, behavior := range c.Behaviors {
			if slices.Contains(conflictGroup, behavior) {
				found = append(found, behavior)
			}
		}
		if len(found) > 1 {
			errors = append(errors, fmt.Sprintf("conflicting behaviors: %s (pick only one)", strings.Join(found, ", ")))
		}
	}

	// Validate variants
	if len(c.Variants) > 1 {
		errors = append(errors, fmt.Sprintf("only one variant allowed, got: %s", strings.Join(c.Variants, ", ")))
	}
	for _, variant := range c.Variants {
		if !slices.Contains(ValidVariants, variant) {
			errors = append(errors, fmt.Sprintf("invalid variant: %s (valid: %s)", variant, strings.Join(ValidVariants, ", ")))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation errors:\n  - %s", strings.Join(errors, "\n  - "))
	}

	return nil
}

// ToRunnerConfig converts SimpleConfig to the existing RunnerConfig format
func (c *SimpleConfig) ToRunnerConfig() (*RunnerConfig, error) {
	// Convert function strings to CCLFunction enums
	var supportedFunctions []config.CCLFunction
	for _, fn := range c.Functions {
		switch fn {
		case "parse":
			supportedFunctions = append(supportedFunctions, config.FunctionParse)
		case "build_hierarchy":
			supportedFunctions = append(supportedFunctions, config.FunctionBuildHierarchy)
		case "get_string":
			supportedFunctions = append(supportedFunctions, config.FunctionGetString)
		case "get_int":
			supportedFunctions = append(supportedFunctions, config.FunctionGetInt)
		case "get_bool":
			supportedFunctions = append(supportedFunctions, config.FunctionGetBool)
		case "get_float":
			supportedFunctions = append(supportedFunctions, config.FunctionGetFloat)
		case "get_list":
			supportedFunctions = append(supportedFunctions, config.FunctionGetList)
		case "filter":
			supportedFunctions = append(supportedFunctions, config.FunctionFilter)
		case "expand_dotted":
			supportedFunctions = append(supportedFunctions, config.FunctionExpandDotted)
		case "canonical_format":
			// Note: FunctionCanonicalFormat may not be available in current ccl-test-lib version
			// supportedFunctions = append(supportedFunctions, config.FunctionCanonicalFormat)
		}
	}

	// Convert feature strings to CCLFeature enums
	var supportedFeatures []config.CCLFeature
	for _, feature := range c.Features {
		switch feature {
		case "comments":
			supportedFeatures = append(supportedFeatures, config.FeatureComments)
		case "experimental_dotted_keys":
			supportedFeatures = append(supportedFeatures, config.FeatureExperimentalDottedKeys)
		case "unicode":
			supportedFeatures = append(supportedFeatures, config.FeatureUnicode)
		case "multiline":
			supportedFeatures = append(supportedFeatures, config.FeatureMultiline)
		case "whitespace":
			supportedFeatures = append(supportedFeatures, config.FeatureWhitespace)
		case "empty_keys":
			supportedFeatures = append(supportedFeatures, config.FeatureEmptyKeys)
		case "optional_typed_accessors":
			// Handle optional typed accessors feature - this might need a new enum value
			// For now, we'll need to check if this enum exists in ccl-test-lib
		}
	}

	// Build behavior choices with defaults
	behaviors := BehaviorChoices{}

	// Set behavior choices from config, with defaults for unspecified
	for _, behavior := range c.Behaviors {
		switch behavior {
		case "boolean_strict":
			strict := config.BehaviorBooleanStrict
			behaviors.Boolean = &strict
		case "boolean_lenient":
			lenient := config.BehaviorBooleanLenient
			behaviors.Boolean = &lenient
		case "tabs_as_content":
			asContent := config.BehaviorTabsAsContent
			behaviors.TabHandling = &asContent
		case "tabs_as_whitespace":
			asWhitespace := config.BehaviorTabsAsWhitespace
			behaviors.TabHandling = &asWhitespace
		case "crlf_preserve_literal":
			preserve := config.BehaviorCRLFPreserve
			behaviors.CRLFHandling = &preserve
		case "crlf_normalize_to_lf":
			normalize := config.BehaviorCRLFNormalize
			behaviors.CRLFHandling = &normalize
		case "indent_spaces":
			spaces := config.BehaviorIndentSpaces
			behaviors.IndentOutput = &spaces
		case "indent_tabs":
			tabs := config.BehaviorIndentTabs
			behaviors.IndentOutput = &tabs
		case "list_coercion_enabled":
			enabled := config.BehaviorListCoercionOn
			behaviors.ListCoercion = &enabled
		case "list_coercion_disabled":
			disabled := config.BehaviorListCoercionOff
			behaviors.ListCoercion = &disabled
		}
	}

	// Set defaults for unspecified behaviors
	if behaviors.Boolean == nil {
		lenient := config.BehaviorBooleanLenient
		behaviors.Boolean = &lenient
	}
	if behaviors.TabHandling == nil {
		asWhitespace := config.BehaviorTabsAsWhitespace
		behaviors.TabHandling = &asWhitespace
	}
	if behaviors.CRLFHandling == nil {
		normalize := config.BehaviorCRLFNormalize
		behaviors.CRLFHandling = &normalize
	}
	if behaviors.IndentOutput == nil {
		spaces := config.BehaviorIndentSpaces
		behaviors.IndentOutput = &spaces
	}
	if behaviors.ListCoercion == nil {
		disabled := config.BehaviorListCoercionOff
		behaviors.ListCoercion = &disabled
	}

	// Set variant choice with default
	variant := VariantChoice{}
	if len(c.Variants) > 0 {
		switch c.Variants[0] {
		case "proposed_behavior":
			proposed := config.VariantProposed
			variant.Specification = &proposed
		case "reference_compliant":
			// Note: Using the same constant as proposed for now
			// TODO: Update when VariantReferenceCompliant is available
			reference := config.VariantProposed
			variant.Specification = &reference
		}
	} else {
		// Default to proposed behavior
		proposed := config.VariantProposed
		variant.Specification = &proposed
	}

	return &RunnerConfig{
		Implementation: ImplementationSettings{
			Name:               "ccl-test-runner",
			Version:            "1.0.0",
			SupportedFunctions: supportedFunctions,
			SupportedFeatures:  supportedFeatures,
		},
		Behaviors: behaviors,
		Variant:   variant,
		TestFiltering: TestFilteringOptions{
			SkipTestsByName: c.SkipTests,
			SkipDisabled:    true,
		},
	}, nil
}
