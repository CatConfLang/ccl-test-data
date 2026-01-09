// Package config provides centralized configuration for the CCL test runner
// with validation for mutually exclusive behavioral choices.
package config

import (
	"fmt"
	"strings"

	"github.com/catconflang/ccl-test-data/config"
)

// RunnerConfig centralizes all behavioral choices, feature selections, and implementation capabilities
// for the CCL test suite. All mutually exclusive choices must be explicitly made.
type RunnerConfig struct {
	Implementation ImplementationSettings `json:"implementation"`
	Behaviors      BehaviorChoices        `json:"behavior_choices"`
	Variant        VariantChoice          `json:"variant_choice"`
	TestFiltering  TestFilteringOptions   `json:"test_filtering"`
}

// ImplementationSettings defines what CCL functions and features are supported
type ImplementationSettings struct {
	Name               string               `json:"name"`
	Version            string               `json:"version"`
	SupportedFunctions []config.CCLFunction `json:"supported_functions"`
	SupportedFeatures  []config.CCLFeature  `json:"supported_features"`
}

// BehaviorChoices contains REQUIRED mutually exclusive behavioral choices
// All fields must be explicitly set - no defaults allowed
type BehaviorChoices struct {
	CRLFHandling   *config.CCLBehavior `json:"crlf_handling"`   // REQUIRED: crlf_normalize_to_lf | crlf_preserve_literal
	TabHandling    *config.CCLBehavior `json:"tab_handling"`    // REQUIRED: tabs_as_content | tabs_as_whitespace
	IndentOutput   *config.CCLBehavior `json:"indent_output"`   // REQUIRED: indent_spaces | indent_tabs
	Boolean        *config.CCLBehavior `json:"boolean"`         // REQUIRED: boolean_strict | boolean_lenient
	ListCoercion   *config.CCLBehavior `json:"list_coercion"`   // REQUIRED: list_coercion_enabled | list_coercion_disabled
	ToplevelIndent *config.CCLBehavior `json:"toplevel_indent"` // REQUIRED: toplevel_indent_strip | toplevel_indent_preserve
}

// VariantChoice contains REQUIRED specification variant choice
type VariantChoice struct {
	Specification *config.CCLVariant `json:"specification"` // REQUIRED: proposed_behavior | reference_compliant
}

// TestFilteringOptions controls which tests are run
type TestFilteringOptions struct {
	RunOnlyFunctions []string `json:"run_only_functions,omitempty"` // Override behavior filtering
	SkipTags         []string `json:"skip_tags,omitempty"`          // Skip tests with these tags
	SkipTestsByName  []string `json:"skip_tests_by_name,omitempty"` // Skip specific tests by name
	SkipDisabled     bool     `json:"skip_disabled"`                // Skip disabled tests
}

// DefaultConfig returns the default configuration for the CCL test runner
// NOTE: This configuration makes explicit behavioral choices for the mock implementation
func DefaultConfig() *RunnerConfig {
	crlf := config.BehaviorCRLFNormalize    // Normalize CRLF to LF for consistent line endings
	tabs := config.BehaviorTabsAsWhitespace // Tabs are whitespace (count for indentation, get trimmed)
	indent := config.BehaviorIndentSpaces   // Use spaces for printed indentation
	boolean := config.BehaviorBooleanLenient
	listCoercion := config.BehaviorListCoercionOff
	toplevelIndent := config.BehaviorToplevelIndentStrip // Strip leading indent at top-level (matches OCaml reference)
	variant := config.VariantProposed

	return &RunnerConfig{
		Implementation: ImplementationSettings{
			Name:    "ccl-test-data-runner",
			Version: "1.0.0",
			SupportedFunctions: []config.CCLFunction{
				config.FunctionParse,
				// Note: BuildHierarchy requires Parse to output flat entries, not multiline
				// config.FunctionBuildHierarchy,
				// Note: Typed functions require BuildHierarchy for object navigation
				// config.FunctionGetString,
				// config.FunctionGetInt,
				// config.FunctionGetBool,
				// config.FunctionGetFloat,
				// config.FunctionGetList,
				// config.FunctionFilter, // Filter function has unused variable issues in test generation
			},
			SupportedFeatures: []config.CCLFeature{
				config.FeatureComments,
				config.FeatureExperimentalDottedKeys,
				config.FeatureUnicode,
			},
		},
		Behaviors: BehaviorChoices{
			CRLFHandling:   &crlf,
			TabHandling:    &tabs,
			IndentOutput:   &indent,
			Boolean:        &boolean,
			ListCoercion:   &listCoercion,
			ToplevelIndent: &toplevelIndent,
		},
		Variant: VariantChoice{
			Specification: &variant,
		},
		TestFiltering: TestFilteringOptions{
			RunOnlyFunctions: []string{"parse", "get-string", "get-int", "get-bool", "get-float", "get-list"}, // Basic functions only
			SkipTags:         []string{"behavior:crlf_normalize_to_lf", "behavior:tabs_as_content"},           // Skip conflicting behaviors
			SkipTestsByName: []string{
				// Indentation-aware parsing (requires multiline value preservation)
				"deep_nested_objects", "nested_duplicate_keys", "round_trip_deeply_nested",
				// CRLF behavior mismatches
				"crlf_normalize_to_lf_proposed", "crlf_normalize_to_lf_indented_proposed",
				// Tab/spacing behavior mismatches
				"canonical_format_consistent_spacing", "canonical_format_line_endings_proposed",
				"canonical_format_tab_preservation", "key_with_tabs",
				// Multiline tests that require advanced parsing
				"multiline_section_header_value", "unindented_multiline_becomes_continuation",
				"multiline_values", "nested_multi_line", "nested_single_line", "complex_multi_newline_whitespace",
				"key_with_newline_before_equals", "round_trip_multiline_values", "round_trip_whitespace_normalization",
				// Complex nested structure tests requiring hierarchy support
				"nested_structure_parsing", "nested_objects_with_lists", "deeply_nested_list_reference",
				"hierarchical_with_expand_dotted_validation", "mixed_dotted_and_regular_keys", "mixed_flat_and_nested",
				// Workflow tests requiring multiple functions
				"complete_lists_workflow", "complete_mixed_workflow", "complete_multiline_workflow",
				"complete_nested_workflow", "real_world_complete_workflow",
				// Round-trip property tests requiring advanced features
				"round_trip_complex_nesting", "round_trip_mixed_content", "round_trip_nested_structures",
				"round_trip_property_complex", "round_trip_property_nested",
				// Algebraic property tests (monoid/semigroup)
				"monoid_left_identity_basic", "monoid_left_identity_nested", "monoid_right_identity_basic",
				"monoid_right_identity_nested", "semigroup_associativity_nested",
				// Advanced list and error handling
				"nested_list_access", "nested_list_access_reference", "list_error_nested_missing_key",
				// Stress tests
				"ocaml_stress_test_original",
			},
			SkipDisabled: true,
		},
	}
}

// Validate ensures all required behavioral choices are made and no conflicts exist
func (rc *RunnerConfig) Validate() error {
	var errors []string

	// Validate required behavioral choices are made
	if rc.Behaviors.CRLFHandling == nil {
		errors = append(errors, "CRLF handling choice is required (crlf_normalize_to_lf | crlf_preserve_literal)")
	}
	if rc.Behaviors.TabHandling == nil {
		errors = append(errors, "Tab handling choice is required (tabs_as_content | tabs_as_whitespace)")
	}
	if rc.Behaviors.IndentOutput == nil {
		errors = append(errors, "Indent output choice is required (indent_spaces | indent_tabs)")
	}
	if rc.Behaviors.Boolean == nil {
		errors = append(errors, "Boolean parsing choice is required (boolean_strict | boolean_lenient)")
	}
	if rc.Behaviors.ListCoercion == nil {
		errors = append(errors, "List coercion choice is required (list_coercion_enabled | list_coercion_disabled)")
	}
	if rc.Behaviors.ToplevelIndent == nil {
		errors = append(errors, "Toplevel indent choice is required (toplevel_indent_strip | toplevel_indent_preserve)")
	}

	// Validate required variant choice is made
	if rc.Variant.Specification == nil {
		errors = append(errors, "Specification variant choice is required (proposed_behavior | reference_compliant)")
	}

	// Validate behavioral choices are from valid conflict groups
	if rc.Behaviors.CRLFHandling != nil {
		if err := rc.validateBehaviorInGroup(*rc.Behaviors.CRLFHandling, "crlf_handling"); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if rc.Behaviors.TabHandling != nil {
		if err := rc.validateBehaviorInGroup(*rc.Behaviors.TabHandling, "tab_handling"); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if rc.Behaviors.IndentOutput != nil {
		if err := rc.validateBehaviorInGroup(*rc.Behaviors.IndentOutput, "indent_output"); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if rc.Behaviors.Boolean != nil {
		if err := rc.validateBehaviorInGroup(*rc.Behaviors.Boolean, "boolean"); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if rc.Behaviors.ListCoercion != nil {
		if err := rc.validateBehaviorInGroup(*rc.Behaviors.ListCoercion, "list_coercion"); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if rc.Behaviors.ToplevelIndent != nil {
		if err := rc.validateBehaviorInGroup(*rc.Behaviors.ToplevelIndent, "toplevel_indent"); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("configuration validation failed:\n  - %s", strings.Join(errors, "\n  - "))
	}

	return nil
}

// validateBehaviorInGroup checks that a behavior belongs to the specified conflict group
func (rc *RunnerConfig) validateBehaviorInGroup(behavior config.CCLBehavior, group string) error {
	conflicts := config.GetBehaviorConflicts()
	groupBehaviors, exists := conflicts[group]
	if !exists {
		return fmt.Errorf("unknown behavior group: %s", group)
	}

	for _, validBehavior := range groupBehaviors {
		if behavior == validBehavior {
			return nil
		}
	}

	var validOptions []string
	for _, b := range groupBehaviors {
		validOptions = append(validOptions, string(b))
	}

	return fmt.Errorf("invalid behavior '%s' for group '%s'. Valid options: %s",
		behavior, group, strings.Join(validOptions, " | "))
}

// ToImplementationConfig converts RunnerConfig to ccl-test-lib ImplementationConfig
func (rc *RunnerConfig) ToImplementationConfig() config.ImplementationConfig {
	var behaviorChoices []config.CCLBehavior

	// Add all selected behaviors
	if rc.Behaviors.CRLFHandling != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.CRLFHandling)
	}
	if rc.Behaviors.TabHandling != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.TabHandling)
	}
	if rc.Behaviors.IndentOutput != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.IndentOutput)
	}
	if rc.Behaviors.Boolean != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.Boolean)
	}
	if rc.Behaviors.ListCoercion != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.ListCoercion)
	}
	if rc.Behaviors.ToplevelIndent != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.ToplevelIndent)
	}

	var variantChoice config.CCLVariant
	if rc.Variant.Specification != nil {
		variantChoice = *rc.Variant.Specification
	}

	return config.ImplementationConfig{
		Name:               rc.Implementation.Name,
		Version:            rc.Implementation.Version,
		SupportedFunctions: rc.Implementation.SupportedFunctions,
		SupportedFeatures:  rc.Implementation.SupportedFeatures,
		BehaviorChoices:    behaviorChoices,
		VariantChoice:      variantChoice,
	}
}

// GetConflictingTags returns tags that should be filtered out based on behavioral choices
func (rc *RunnerConfig) GetConflictingTags() []string {
	var conflictingTags []string
	conflicts := config.GetBehaviorConflicts()

	// For each conflict group, exclude the behaviors we didn't choose
	for _, groupBehaviors := range conflicts {
		for _, behavior := range groupBehaviors {
			// Check if this behavior conflicts with our choices
			isChosen := rc.isBehaviorChosen(behavior)
			if !isChosen {
				conflictingTags = append(conflictingTags, "behavior:"+string(behavior))
			}
		}
	}

	// Add variant conflicts
	if rc.Variant.Specification != nil {
		if *rc.Variant.Specification == config.VariantProposed {
			conflictingTags = append(conflictingTags, "variant:reference_compliant")
		} else {
			conflictingTags = append(conflictingTags, "variant:proposed_behavior")
		}
	}

	return conflictingTags
}

// isBehaviorChosen checks if a specific behavior was chosen in our configuration
func (rc *RunnerConfig) isBehaviorChosen(behavior config.CCLBehavior) bool {
	return (rc.Behaviors.CRLFHandling != nil && *rc.Behaviors.CRLFHandling == behavior) ||
		(rc.Behaviors.TabHandling != nil && *rc.Behaviors.TabHandling == behavior) ||
		(rc.Behaviors.IndentOutput != nil && *rc.Behaviors.IndentOutput == behavior) ||
		(rc.Behaviors.Boolean != nil && *rc.Behaviors.Boolean == behavior) ||
		(rc.Behaviors.ListCoercion != nil && *rc.Behaviors.ListCoercion == behavior) ||
		(rc.Behaviors.ToplevelIndent != nil && *rc.Behaviors.ToplevelIndent == behavior)
}
