// Package generator provides utilities for transforming source format
// CCL tests to implementation-friendly flat format.
package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// BehaviorMetadata represents the x-behaviorMetadata section in source-format.json
type BehaviorMetadata struct {
	Behaviors map[string]BehaviorInfo  `json:"behaviors"`
	Defaults  BehaviorMetadataDefaults `json:"defaults,omitempty"`
}

// BehaviorInfo contains metadata about a single behavior
type BehaviorInfo struct {
	Description           string   `json:"description"`
	AffectedFunctions     []string `json:"affectedFunctions"`
	MutuallyExclusiveWith []string `json:"mutuallyExclusiveWith"`
}

// BehaviorMetadataDefaults contains default handling rules
type BehaviorMetadataDefaults struct {
	UnmappedBehavior string `json:"unmappedBehavior"`
	Description      string `json:"description"`
}

// sourceFormatSchema represents the structure needed to extract x-behaviorMetadata
type sourceFormatSchema struct {
	BehaviorMetadata *BehaviorMetadata `json:"x-behaviorMetadata"`
}

// LoadBehaviorMetadata loads behavior metadata from source-format.json in the schemas directory
func LoadBehaviorMetadata(schemasDir string) (*BehaviorMetadata, error) {
	schemaPath := filepath.Join(schemasDir, "source-format.json")

	data, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read source-format.json: %w", err)
	}

	var schema sourceFormatSchema
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse source-format.json: %w", err)
	}

	if schema.BehaviorMetadata == nil {
		return nil, fmt.Errorf("x-behaviorMetadata section not found in source-format.json")
	}

	return schema.BehaviorMetadata, nil
}

// FilterBehaviorsForFunction filters behaviors to only include those that affect
// the given function, using the behavior metadata mapping.
func (m *BehaviorMetadata) FilterBehaviorsForFunction(behaviors []string, function string) []string {
	if behaviors == nil {
		return make([]string, 0)
	}

	filtered := make([]string, 0, len(behaviors))
	for _, behavior := range behaviors {
		info, exists := m.Behaviors[behavior]
		if !exists {
			// Unmapped behavior = global, always include
			filtered = append(filtered, behavior)
			continue
		}

		// Check if this function is in the affected list
		for _, fn := range info.AffectedFunctions {
			if fn == function {
				filtered = append(filtered, behavior)
				break
			}
		}
	}

	return filtered
}

// GetConflictingBehaviors returns all behaviors that conflict with the given behaviors
// based on the mutuallyExclusiveWith definitions.
func (m *BehaviorMetadata) GetConflictingBehaviors(behaviors []string) []string {
	if behaviors == nil {
		return nil
	}

	conflictSet := make(map[string]bool)
	for _, behavior := range behaviors {
		info, exists := m.Behaviors[behavior]
		if !exists {
			continue
		}

		for _, conflict := range info.MutuallyExclusiveWith {
			conflictSet[conflict] = true
		}
	}

	// Convert map to slice
	if len(conflictSet) == 0 {
		return nil
	}

	conflicts := make([]string, 0, len(conflictSet))
	for conflict := range conflictSet {
		conflicts = append(conflicts, conflict)
	}

	// Sort for deterministic output
	sort.Strings(conflicts)

	return conflicts
}

// ValidateBehavior checks if a behavior is known in the metadata
func (m *BehaviorMetadata) ValidateBehavior(behavior string) bool {
	_, exists := m.Behaviors[behavior]
	return exists
}

// GetAllBehaviors returns all known behavior names in sorted order
func (m *BehaviorMetadata) GetAllBehaviors() []string {
	behaviors := make([]string, 0, len(m.Behaviors))
	for name := range m.Behaviors {
		behaviors = append(behaviors, name)
	}
	sort.Strings(behaviors)
	return behaviors
}

// ValidationResult contains results from validating source tests
type ValidationResult struct {
	TestName string
	Warnings []string
	Errors   []string
}

// ValidateSourceTest validates a source test against the behavior metadata
// Note: Missing conflict declarations are not warned about since auto-generation
// from behavior metadata handles conflicts automatically (--auto-conflicts=true by default)
func (m *BehaviorMetadata) ValidateSourceTest(testName string, behaviors []string, declaredConflicts []string) ValidationResult {
	result := ValidationResult{
		TestName: testName,
		Warnings: make([]string, 0),
		Errors:   make([]string, 0),
	}

	// Check for unknown behaviors
	for _, behavior := range behaviors {
		if !m.ValidateBehavior(behavior) {
			result.Warnings = append(result.Warnings,
				fmt.Sprintf("unknown behavior '%s' (not in metadata, treating as global)", behavior))
		}
	}

	// Conflicts are auto-generated from behavior metadata, so we don't warn about
	// missing declarations in source tests. The x-behaviorMetadata section in source-format.json is the
	// single source of truth for mutually exclusive behaviors.

	return result
}

// findBehaviorCausingConflict finds which behavior from the list causes the conflict
func findBehaviorCausingConflict(behaviors []string, conflictingBehavior string, m *BehaviorMetadata) string {
	for _, behavior := range behaviors {
		info, exists := m.Behaviors[behavior]
		if !exists {
			continue
		}
		for _, conflict := range info.MutuallyExclusiveWith {
			if conflict == conflictingBehavior {
				return behavior
			}
		}
	}
	return "unknown"
}
