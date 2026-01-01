package main

import (
	"fmt"
	"sort"

	"gopkg.in/yaml.v3"
)

// Severity levels for API changes
type Severity string

const (
	SeverityBreaking   Severity = "breaking"   // Field removed, type change, required added
	SeverityDeprecated Severity = "deprecated" // Field marked deprecated
	SeverityMinor      Severity = "minor"      // Field/endpoint added
	SeverityDocs       Severity = "docs"       // Description only
	SeverityNone       Severity = "none"       // No changes
)

// SeverityOrder defines the priority of severities (higher = more severe)
var SeverityOrder = map[Severity]int{
	SeverityNone:       0,
	SeverityDocs:       1,
	SeverityMinor:      2,
	SeverityDeprecated: 3,
	SeverityBreaking:   4,
}

// DiffResult contains the comparison result
type DiffResult struct {
	MaxSeverity  Severity
	TotalChanges int
	PathsAdded   int
	PathsRemoved int
	FieldsAdded  int
	FieldsRemoved int
}

// OpenAPISpec represents the OpenAPI specification structure
type OpenAPISpec struct {
	OpenAPI    string                 `yaml:"openapi"`
	Info       map[string]interface{} `yaml:"info"`
	Paths      map[string]interface{} `yaml:"paths"`
	Components *Components            `yaml:"components"`
}

// Components contains reusable schema definitions
type Components struct {
	Schemas map[string]interface{} `yaml:"schemas"`
}

// CompareSpecs compares two OpenAPI specs and returns the diff result
func CompareSpecs(oldData, newData []byte) (*DiffResult, error) {
	var oldSpec, newSpec OpenAPISpec

	if err := yaml.Unmarshal(oldData, &oldSpec); err != nil {
		return nil, fmt.Errorf("failed to parse old spec: %w", err)
	}

	if err := yaml.Unmarshal(newData, &newSpec); err != nil {
		return nil, fmt.Errorf("failed to parse new spec: %w", err)
	}

	result := &DiffResult{
		MaxSeverity: SeverityNone,
	}

	// Compare paths
	comparePaths(&oldSpec, &newSpec, result)

	// Compare schemas
	if oldSpec.Components != nil && newSpec.Components != nil {
		compareSchemas(oldSpec.Components.Schemas, newSpec.Components.Schemas, result)
	}

	return result, nil
}

// comparePaths compares API paths between specs
func comparePaths(oldSpec, newSpec *OpenAPISpec, result *DiffResult) {
	oldPaths := make(map[string]bool)
	newPaths := make(map[string]bool)

	for path := range oldSpec.Paths {
		oldPaths[path] = true
	}
	for path := range newSpec.Paths {
		newPaths[path] = true
	}

	// Find added paths (minor severity)
	for path := range newPaths {
		if !oldPaths[path] {
			result.PathsAdded++
			result.TotalChanges++
			updateMaxSeverity(result, SeverityMinor)
		}
	}

	// Find removed paths (breaking severity)
	for path := range oldPaths {
		if !newPaths[path] {
			result.PathsRemoved++
			result.TotalChanges++
			updateMaxSeverity(result, SeverityBreaking)
		}
	}
}

// compareSchemas compares all schemas between specs
func compareSchemas(oldSchemas, newSchemas map[string]interface{}, result *DiffResult) {
	// Get all schema names
	schemaSet := make(map[string]bool)
	for name := range oldSchemas {
		schemaSet[name] = true
	}
	for name := range newSchemas {
		schemaSet[name] = true
	}

	var schemaNames []string
	for name := range schemaSet {
		schemaNames = append(schemaNames, name)
	}
	sort.Strings(schemaNames)

	// Compare each schema
	for _, name := range schemaNames {
		oldSchema, oldExists := oldSchemas[name]
		newSchema, newExists := newSchemas[name]

		if !oldExists && newExists {
			// Schema added - minor
			result.TotalChanges++
			updateMaxSeverity(result, SeverityMinor)
			continue
		}

		if oldExists && !newExists {
			// Schema removed - breaking
			result.TotalChanges++
			updateMaxSeverity(result, SeverityBreaking)
			continue
		}

		// Both exist - compare fields
		compareSchemaFields(oldSchema, newSchema, result)
	}
}

// compareSchemaFields compares fields within a schema
func compareSchemaFields(oldSchema, newSchema interface{}, result *DiffResult) {
	oldMap, oldOk := oldSchema.(map[string]interface{})
	newMap, newOk := newSchema.(map[string]interface{})

	if !oldOk || !newOk {
		return
	}

	oldProps := getProperties(oldMap)
	newProps := getProperties(newMap)
	oldRequired := getRequiredFields(oldMap)
	newRequired := getRequiredFields(newMap)

	// Find added fields
	for propName := range newProps {
		if _, exists := oldProps[propName]; !exists {
			result.FieldsAdded++
			result.TotalChanges++

			// Adding required field is breaking, optional is minor
			if newRequired[propName] {
				updateMaxSeverity(result, SeverityBreaking)
			} else {
				updateMaxSeverity(result, SeverityMinor)
			}
		}
	}

	// Find removed fields (always breaking)
	for propName := range oldProps {
		if _, exists := newProps[propName]; !exists {
			result.FieldsRemoved++
			result.TotalChanges++
			updateMaxSeverity(result, SeverityBreaking)
		}
	}

	// Find changed fields
	for propName, newProp := range newProps {
		if oldProp, exists := oldProps[propName]; exists {
			compareField(propName, oldProp, newProp, oldRequired[propName], newRequired[propName], result)
		}
	}
}

// compareField compares a single field
func compareField(name string, oldProp, newProp interface{}, wasRequired, isRequired bool, result *DiffResult) {
	oldMap, oldOk := oldProp.(map[string]interface{})
	newMap, newOk := newProp.(map[string]interface{})

	if !oldOk || !newOk {
		return
	}

	// Type change
	oldType := getType(oldMap)
	newType := getType(newMap)
	if oldType != newType {
		result.TotalChanges++
		updateMaxSeverity(result, SeverityBreaking)
	}

	// Required status change
	if wasRequired != isRequired {
		result.TotalChanges++
		if isRequired && !wasRequired {
			// Optional -> Required is breaking
			updateMaxSeverity(result, SeverityBreaking)
		} else {
			// Required -> Optional is minor (relaxing)
			updateMaxSeverity(result, SeverityMinor)
		}
	}

	// Deprecated status change
	oldDeprecated := getBool(oldMap, "deprecated")
	newDeprecated := getBool(newMap, "deprecated")
	if !oldDeprecated && newDeprecated {
		result.TotalChanges++
		updateMaxSeverity(result, SeverityDeprecated)
	}

	// Enum changes
	oldEnum := getEnumValues(oldMap)
	newEnum := getEnumValues(newMap)
	if len(oldEnum) > 0 || len(newEnum) > 0 {
		compareEnums(oldEnum, newEnum, result)
	}
}

// compareEnums compares enum values
func compareEnums(oldEnum, newEnum []string, result *DiffResult) {
	oldSet := make(map[string]bool)
	newSet := make(map[string]bool)

	for _, v := range oldEnum {
		oldSet[v] = true
	}
	for _, v := range newEnum {
		newSet[v] = true
	}

	// Removed enum values (breaking)
	for v := range oldSet {
		if !newSet[v] {
			result.TotalChanges++
			updateMaxSeverity(result, SeverityBreaking)
		}
	}

	// Added enum values (minor)
	for v := range newSet {
		if !oldSet[v] {
			result.TotalChanges++
			updateMaxSeverity(result, SeverityMinor)
		}
	}
}

// Helper functions

func getProperties(schema map[string]interface{}) map[string]interface{} {
	if props, ok := schema["properties"].(map[string]interface{}); ok {
		return props
	}
	return make(map[string]interface{})
}

func getRequiredFields(schema map[string]interface{}) map[string]bool {
	required := make(map[string]bool)
	if req, ok := schema["required"].([]interface{}); ok {
		for _, r := range req {
			if s, ok := r.(string); ok {
				required[s] = true
			}
		}
	}
	return required
}

func getType(prop map[string]interface{}) string {
	if t, ok := prop["type"].(string); ok {
		return t
	}
	if ref, ok := prop["$ref"].(string); ok {
		return "ref:" + ref
	}
	if _, ok := prop["oneOf"]; ok {
		return "oneOf"
	}
	if _, ok := prop["anyOf"]; ok {
		return "anyOf"
	}
	if _, ok := prop["allOf"]; ok {
		return "allOf"
	}
	return "unknown"
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

func getEnumValues(prop map[string]interface{}) []string {
	if enum, ok := prop["enum"].([]interface{}); ok {
		values := make([]string, 0, len(enum))
		for _, v := range enum {
			if s, ok := v.(string); ok {
				values = append(values, s)
			}
		}
		return values
	}
	return nil
}

func updateMaxSeverity(result *DiffResult, sev Severity) {
	if SeverityOrder[sev] > SeverityOrder[result.MaxSeverity] {
		result.MaxSeverity = sev
	}
}
