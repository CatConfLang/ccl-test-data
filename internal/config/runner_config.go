// Package config provides centralized configuration for the CCL test runner
// with validation for mutually exclusive behavioral choices.
package config

import (
	"fmt"
	"strings"

	"github.com/tylerbu/ccl-test-lib/config"
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
	Name               string                `json:"name"`
	Version            string                `json:"version"`
	SupportedFunctions []config.CCLFunction  `json:"supported_functions"`
	SupportedFeatures  []config.CCLFeature   `json:"supported_features"`
}

// BehaviorChoices contains REQUIRED mutually exclusive behavioral choices
// All fields must be explicitly set - no defaults allowed
type BehaviorChoices struct {
	CRLFHandling  *config.CCLBehavior `json:"crlf_handling"`   // REQUIRED: crlf_normalize_to_lf | crlf_preserve_literal
	TabHandling   *config.CCLBehavior `json:"tab_handling"`    // REQUIRED: tabs_preserve | tabs_to_spaces  
	Spacing       *config.CCLBehavior `json:"spacing"`         // REQUIRED: strict_spacing | loose_spacing
	Boolean       *config.CCLBehavior `json:"boolean"`         // REQUIRED: boolean_strict | boolean_lenient
	ListCoercion  *config.CCLBehavior `json:"list_coercion"`   // REQUIRED: list_coercion_enabled | list_coercion_disabled
}

// VariantChoice contains REQUIRED specification variant choice
type VariantChoice struct {
	Specification *config.CCLVariant `json:"specification"` // REQUIRED: proposed_behavior | reference_compliant
}

// TestFilteringOptions controls which tests are run
type TestFilteringOptions struct {
	RunOnlyFunctions []string `json:"run_only_functions,omitempty"` // Override behavior filtering
	SkipTags         []string `json:"skip_tags,omitempty"`          // Skip tests with these tags
	SkipDisabled     bool     `json:"skip_disabled"`                // Skip disabled tests
}

// DefaultConfig returns the default configuration for the CCL test runner
// NOTE: This configuration makes explicit behavioral choices for the mock implementation
func DefaultConfig() *RunnerConfig {
	crlf := config.BehaviorCRLFNormalize
	tabs := config.BehaviorTabsPreserve
	spacing := config.BehaviorLooseSpacing
	boolean := config.BehaviorBooleanLenient
	listCoercion := config.BehaviorListCoercionOff
	variant := config.VariantProposed

	return &RunnerConfig{
		Implementation: ImplementationSettings{
			Name:    "ccl-test-data-runner",
			Version: "1.0.0",
			SupportedFunctions: []config.CCLFunction{
				config.FunctionParse,
				config.FunctionBuildHierarchy,
				config.FunctionGetString,
				config.FunctionGetInt,
				config.FunctionGetBool,
				config.FunctionGetFloat,
				config.FunctionGetList,
				config.FunctionFilter,
			},
			SupportedFeatures: []config.CCLFeature{
				config.FeatureComments,
				config.FeatureExperimentalDottedKeys,
				config.FeatureUnicode,
			},
		},
		Behaviors: BehaviorChoices{
			CRLFHandling: &crlf,
			TabHandling:  &tabs,
			Spacing:      &spacing,
			Boolean:      &boolean,
			ListCoercion: &listCoercion,
		},
		Variant: VariantChoice{
			Specification: &variant,
		},
		TestFiltering: TestFilteringOptions{
			RunOnlyFunctions: []string{"parse"}, // Default to Level 1 tests only
			SkipDisabled:     true,
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
		errors = append(errors, "Tab handling choice is required (tabs_preserve | tabs_to_spaces)")
	}
	if rc.Behaviors.Spacing == nil {
		errors = append(errors, "Spacing choice is required (strict_spacing | loose_spacing)")
	}
	if rc.Behaviors.Boolean == nil {
		errors = append(errors, "Boolean parsing choice is required (boolean_strict | boolean_lenient)")
	}
	if rc.Behaviors.ListCoercion == nil {
		errors = append(errors, "List coercion choice is required (list_coercion_enabled | list_coercion_disabled)")
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
	if rc.Behaviors.Spacing != nil {
		if err := rc.validateBehaviorInGroup(*rc.Behaviors.Spacing, "spacing"); err != nil {
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
	if rc.Behaviors.Spacing != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.Spacing)
	}
	if rc.Behaviors.Boolean != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.Boolean)
	}
	if rc.Behaviors.ListCoercion != nil {
		behaviorChoices = append(behaviorChoices, *rc.Behaviors.ListCoercion)
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
		(rc.Behaviors.Spacing != nil && *rc.Behaviors.Spacing == behavior) ||
		(rc.Behaviors.Boolean != nil && *rc.Behaviors.Boolean == behavior) ||
		(rc.Behaviors.ListCoercion != nil && *rc.Behaviors.ListCoercion == behavior)
}