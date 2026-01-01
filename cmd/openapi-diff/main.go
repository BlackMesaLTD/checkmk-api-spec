// Package main implements the openapi-diff tool for comparing OpenAPI specs across versions.
//
// Usage:
//
//	openapi-diff -old specs/2.3.0p41/openapi.yaml -new specs/2.4.0p17/openapi.yaml -output diff.json
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

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

// DiffReport contains the differences between two specs
type DiffReport struct {
	OldVersion   string          `json:"old_version"`
	NewVersion   string          `json:"new_version"`
	PathChanges  *PathChanges    `json:"path_changes"`
	SchemaChanges []*SchemaChange `json:"schema_changes"`
	Summary      *Summary        `json:"summary"`
}

// PathChanges tracks path-level differences
type PathChanges struct {
	Added   []string `json:"added,omitempty"`
	Removed []string `json:"removed,omitempty"`
}

// SchemaChange tracks changes to a specific schema
type SchemaChange struct {
	SchemaName    string        `json:"schema_name"`
	AddedFields   []*FieldDiff  `json:"added_fields,omitempty"`
	RemovedFields []*FieldDiff  `json:"removed_fields,omitempty"`
	ChangedFields []*FieldDiff  `json:"changed_fields,omitempty"`
}

// FieldDiff describes a field difference
type FieldDiff struct {
	Path        string `json:"path"`
	OldType     string `json:"old_type,omitempty"`
	NewType     string `json:"new_type,omitempty"`
	Description string `json:"description,omitempty"`
}

// Summary provides high-level statistics
type Summary struct {
	TotalSchemasOld     int `json:"total_schemas_old"`
	TotalSchemasNew     int `json:"total_schemas_new"`
	SchemasAdded        int `json:"schemas_added"`
	SchemasRemoved      int `json:"schemas_removed"`
	SchemasWithChanges  int `json:"schemas_with_changes"`
	TotalFieldsAdded    int `json:"total_fields_added"`
	TotalFieldsRemoved  int `json:"total_fields_removed"`
}

func main() {
	var (
		oldPath   = flag.String("old", "", "Path to old OpenAPI spec")
		newPath   = flag.String("new", "", "Path to new OpenAPI spec")
		output    = flag.String("output", "", "Output JSON file (optional, prints to stdout if not specified)")
		resources = flag.String("resources", "", "Comma-separated list of resources to compare (optional)")
		verbose   = flag.Bool("verbose", false, "Print detailed output")
	)
	flag.Parse()

	if *oldPath == "" || *newPath == "" {
		log.Fatal("Error: both -old and -new flags are required")
	}

	// Load specs
	oldSpec, err := loadSpec(*oldPath)
	if err != nil {
		log.Fatalf("Failed to load old spec: %v", err)
	}

	newSpec, err := loadSpec(*newPath)
	if err != nil {
		log.Fatalf("Failed to load new spec: %v", err)
	}

	// Get versions
	oldVersion := getVersion(oldSpec, *oldPath)
	newVersion := getVersion(newSpec, *newPath)

	// Filter schemas if resources specified
	var schemasToCompare []string
	if *resources != "" {
		resourceList := strings.Split(*resources, ",")
		schemasToCompare = getSchemasForResources(resourceList)
	}

	// Generate diff
	report := generateDiff(oldSpec, newSpec, oldVersion, newVersion, schemasToCompare, *verbose)

	// Output
	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal report: %v", err)
	}

	if *output != "" {
		if err := os.WriteFile(*output, jsonData, 0644); err != nil {
			log.Fatalf("Failed to write output: %v", err)
		}
		fmt.Printf("Diff report written to: %s\n", *output)
	} else {
		fmt.Println(string(jsonData))
	}

	// Print summary
	if report.Summary != nil {
		fmt.Printf("\nSummary:\n")
		fmt.Printf("  Schemas: %d → %d\n", report.Summary.TotalSchemasOld, report.Summary.TotalSchemasNew)
		fmt.Printf("  Added: %d, Removed: %d, Changed: %d\n",
			report.Summary.SchemasAdded, report.Summary.SchemasRemoved, report.Summary.SchemasWithChanges)
		fmt.Printf("  Fields added: %d, removed: %d\n",
			report.Summary.TotalFieldsAdded, report.Summary.TotalFieldsRemoved)
	}
}

func loadSpec(path string) (*OpenAPISpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var spec OpenAPISpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, err
	}

	return &spec, nil
}

func getVersion(spec *OpenAPISpec, path string) string {
	if spec.Info != nil {
		if version, ok := spec.Info["version"].(string); ok {
			return version
		}
	}
	// Try to extract from filename
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.Contains(part, "checkmk") || strings.HasPrefix(part, "2.") {
			return part
		}
	}
	return "unknown"
}

func getSchemasForResources(resources []string) []string {
	// Resource to schema mapping (simplified)
	resourceSchemas := map[string][]string{
		"host":          {"HostConfig", "CreateHost", "UpdateHost", "HostExtensions", "HostCreateAttribute"},
		"folder":        {"FolderConfig", "CreateFolder", "UpdateFolder", "FolderExtensions"},
		"aux_tag":       {"AuxTag", "CreateAuxTag", "UpdateAuxTag", "AuxTagExtensions"},
		"tag_group":     {"TagGroup", "CreateTagGroup", "UpdateTagGroup"},
		"user":          {"UserConfig", "CreateUser", "UpdateUser"},
		"contact_group": {"ContactGroup", "CreateContactGroup", "UpdateContactGroup"},
	}

	schemaSet := make(map[string]bool)
	for _, res := range resources {
		if schemas, ok := resourceSchemas[res]; ok {
			for _, s := range schemas {
				schemaSet[s] = true
			}
		}
	}

	result := make([]string, 0, len(schemaSet))
	for s := range schemaSet {
		result = append(result, s)
	}
	sort.Strings(result)
	return result
}

func generateDiff(oldSpec, newSpec *OpenAPISpec, oldVersion, newVersion string, schemasToCompare []string, verbose bool) *DiffReport {
	report := &DiffReport{
		OldVersion:  oldVersion,
		NewVersion:  newVersion,
		PathChanges: &PathChanges{},
		Summary:     &Summary{},
	}

	// Compare paths
	oldPaths := make(map[string]bool)
	newPaths := make(map[string]bool)

	for path := range oldSpec.Paths {
		oldPaths[path] = true
	}
	for path := range newSpec.Paths {
		newPaths[path] = true
	}

	for path := range newPaths {
		if !oldPaths[path] {
			report.PathChanges.Added = append(report.PathChanges.Added, path)
		}
	}
	for path := range oldPaths {
		if !newPaths[path] {
			report.PathChanges.Removed = append(report.PathChanges.Removed, path)
		}
	}
	sort.Strings(report.PathChanges.Added)
	sort.Strings(report.PathChanges.Removed)

	// Compare schemas
	if oldSpec.Components == nil || newSpec.Components == nil {
		return report
	}

	oldSchemas := oldSpec.Components.Schemas
	newSchemas := newSpec.Components.Schemas

	report.Summary.TotalSchemasOld = len(oldSchemas)
	report.Summary.TotalSchemasNew = len(newSchemas)

	// Determine which schemas to compare
	var schemaNames []string
	if len(schemasToCompare) > 0 {
		schemaNames = schemasToCompare
	} else {
		schemaSet := make(map[string]bool)
		for name := range oldSchemas {
			schemaSet[name] = true
		}
		for name := range newSchemas {
			schemaSet[name] = true
		}
		for name := range schemaSet {
			schemaNames = append(schemaNames, name)
		}
		sort.Strings(schemaNames)
	}

	// Compare each schema
	for _, schemaName := range schemaNames {
		oldSchema, oldExists := oldSchemas[schemaName]
		newSchema, newExists := newSchemas[schemaName]

		if !oldExists && newExists {
			report.Summary.SchemasAdded++
			continue
		}
		if oldExists && !newExists {
			report.Summary.SchemasRemoved++
			continue
		}

		// Both exist - compare fields
		change := compareSchemas(schemaName, oldSchema, newSchema)
		if change != nil {
			report.SchemaChanges = append(report.SchemaChanges, change)
			report.Summary.SchemasWithChanges++
			report.Summary.TotalFieldsAdded += len(change.AddedFields)
			report.Summary.TotalFieldsRemoved += len(change.RemovedFields)
		}
	}

	return report
}

func compareSchemas(name string, old, new interface{}) *SchemaChange {
	oldMap, oldOk := old.(map[string]interface{})
	newMap, newOk := new.(map[string]interface{})

	if !oldOk || !newOk {
		return nil
	}

	change := &SchemaChange{
		SchemaName: name,
	}

	// Get properties
	oldProps := getProperties(oldMap)
	newProps := getProperties(newMap)

	// Find added fields
	for propName, propVal := range newProps {
		if _, exists := oldProps[propName]; !exists {
			change.AddedFields = append(change.AddedFields, &FieldDiff{
				Path:        propName,
				NewType:     getType(propVal),
				Description: getDescription(propVal),
			})
		}
	}

	// Find removed fields
	for propName, propVal := range oldProps {
		if _, exists := newProps[propName]; !exists {
			change.RemovedFields = append(change.RemovedFields, &FieldDiff{
				Path:    propName,
				OldType: getType(propVal),
			})
		}
	}

	// Find changed fields
	for propName, newVal := range newProps {
		if oldVal, exists := oldProps[propName]; exists {
			oldType := getType(oldVal)
			newType := getType(newVal)
			if oldType != newType {
				change.ChangedFields = append(change.ChangedFields, &FieldDiff{
					Path:    propName,
					OldType: oldType,
					NewType: newType,
				})
			}
		}
	}

	// Sort for consistent output
	sort.Slice(change.AddedFields, func(i, j int) bool {
		return change.AddedFields[i].Path < change.AddedFields[j].Path
	})
	sort.Slice(change.RemovedFields, func(i, j int) bool {
		return change.RemovedFields[i].Path < change.RemovedFields[j].Path
	})

	// Return nil if no changes
	if len(change.AddedFields) == 0 && len(change.RemovedFields) == 0 && len(change.ChangedFields) == 0 {
		return nil
	}

	return change
}

func getProperties(schema map[string]interface{}) map[string]interface{} {
	if props, ok := schema["properties"].(map[string]interface{}); ok {
		return props
	}
	return make(map[string]interface{})
}

func getType(prop interface{}) string {
	if propMap, ok := prop.(map[string]interface{}); ok {
		if t, ok := propMap["type"].(string); ok {
			return t
		}
		if _, ok := propMap["$ref"].(string); ok {
			return "ref"
		}
		if _, ok := propMap["oneOf"]; ok {
			return "oneOf"
		}
		if _, ok := propMap["anyOf"]; ok {
			return "anyOf"
		}
	}
	return "unknown"
}

func getDescription(prop interface{}) string {
	if propMap, ok := prop.(map[string]interface{}); ok {
		if desc, ok := propMap["description"].(string); ok {
			if len(desc) > 100 {
				return desc[:100] + "..."
			}
			return desc
		}
	}
	return ""
}
