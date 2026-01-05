// Package generator provides utilities for transforming source format
// CCL tests to implementation-friendly flat format.
package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// BehaviorMetadata represents the behavior-metadata.json file structure
type BehaviorMetadata struct {
	Schema      string                      `json:"$schema,omitempty"`
	ID          string                      `json:"$id,omitempty"`
	Title       string                      `json:"title,omitempty"`
	Description string                      `json:"description,omitempty"`
	Behaviors   map[string]BehaviorInfo     `json:"behaviors"`
	Defaults    BehaviorMetadataDefaults    `json:"defaults,omitempty"`
}

// BehaviorInfo contains metadata about a single behavior
type BehaviorInfo struct {
	Description         string   `json:"description"`
	AffectedFunctions   []string `json:"affectedFunctions"`
	MutuallyExclusiveWith []string `json:"mutuallyExclusiveWith"`
}

// BehaviorMetadataDefaults contains default handling rules
type BehaviorMetadataDefaults struct {
	UnmappedBehavior string `json:"unmappedBehavior"`
	Description      string `json:"description"`
}

// LoadBehaviorMetadata loads behavior metadata from the schemas directory
func LoadBehaviorMetadata(schemasDir string) (*BehaviorMetadata, error) {
	metadataPath := filepath.Join(schemasDir, "behavior-metadata.json")

	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read behavior metadata: %w", err)
	}

	var metadata BehaviorMetadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse behavior metadata: %w", err)
	}

	return &metadata, nil
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

	return conflicts
}

// ValidateBehavior checks if a behavior is known in the metadata
func (m *BehaviorMetadata) ValidateBehavior(behavior string) bool {
	_, exists := m.Behaviors[behavior]
	return exists
}

// GetAllBehaviors returns all known behavior names
func (m *BehaviorMetadata) GetAllBehaviors() []string {
	behaviors := make([]string, 0, len(m.Behaviors))
	for name := range m.Behaviors {
		behaviors = append(behaviors, name)
	}
	return behaviors
}

// ValidationResult contains results from validating source tests
type ValidationResult struct {
	TestName       string
	Warnings       []string
	Errors         []string
}

// ValidateSourceTest validates a source test against the behavior metadata
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

	// Check for missing conflict declarations
	expectedConflicts := m.GetConflictingBehaviors(behaviors)
	if len(expectedConflicts) > 0 {
		declaredSet := make(map[string]bool)
		for _, c := range declaredConflicts {
			declaredSet[c] = true
		}

		for _, expected := range expectedConflicts {
			if !declaredSet[expected] {
				result.Warnings = append(result.Warnings,
					fmt.Sprintf("missing conflict declaration: behavior '%s' should conflict with '%s'",
						findBehaviorCausingConflict(behaviors, expected, m), expected))
			}
		}
	}

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
