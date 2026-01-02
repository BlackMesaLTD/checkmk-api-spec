// Package main implements the openapi-diff tool for comparing OpenAPI specs across versions.
//
// This tool compares two OpenAPI specs and produces a detailed diff report with
// 4-level severity classification:
//   - BREAKING: Field removed, type changed incompatibly, required field added to request
//   - DEPRECATED: Field marked deprecated
//   - MINOR: Optional field added, new endpoint added
//   - DOCS: Description-only changes, no functional impact
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

// Severity levels for API changes
type Severity string

const (
	SeverityBreaking   Severity = "breaking"   // Field removed, type change, required added
	SeverityDeprecated Severity = "deprecated" // Field marked deprecated
	SeverityMinor      Severity = "minor"      // Field/endpoint added
	SeverityDocs       Severity = "docs"       // Description only
	SeverityNone       Severity = "none"       // No changes
)

// Category classifies API endpoints by their purpose
type Category string

const (
	CategorySetup      Category = "setup"      // Configuration endpoints (Terraform provider)
	CategoryMonitoring Category = "monitoring" // Operational/live actions
	CategoryInternal   Category = "internal"   // Internal-only endpoints
)

// Tag sets for category classification
var monitoringTags = map[string]bool{
	"Acknowledge problems": true,
	"Comments":             true,
	"Downtimes":            true,
	"Service status":       true,
	"Metrics":              true,
	"Parent scan":          true,
	"Background Jobs":      true,
}

var internalTags = map[string]bool{
	"Autocomplete (internal)": true,
	"Hosts (internal)":        true,
	"Miscellaneous":           true,
}

// SeverityOrder defines the priority of severities (higher = more severe)
var SeverityOrder = map[Severity]int{
	SeverityNone:       0,
	SeverityDocs:       1,
	SeverityMinor:      2,
	SeverityDeprecated: 3,
	SeverityBreaking:   4,
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

// DiffReport contains the differences between two specs
type DiffReport struct {
	OldVersion      string           `json:"old_version"`
	NewVersion      string           `json:"new_version"`
	MaxSeverity     Severity         `json:"max_severity"`
	PathChanges     *PathChanges     `json:"path_changes"`
	SchemaChanges   []*SchemaChange  `json:"schema_changes"`
	Summary         *Summary         `json:"summary"`
	CategorySummary *CategorySummary `json:"category_summary,omitempty"`
}

// CategorySummary provides breakdown by API category
type CategorySummary struct {
	Setup      *CategoryStats `json:"setup"`
	Monitoring *CategoryStats `json:"monitoring"`
	Internal   *CategoryStats `json:"internal"`
}

// CategoryStats tracks changes within a category
type CategoryStats struct {
	PathsAdded   int `json:"paths_added"`
	PathsRemoved int `json:"paths_removed"`
	Breaking     int `json:"breaking"`
	Minor        int `json:"minor"`
}

// PathChanges tracks path-level differences
type PathChanges struct {
	Added   []*PathChange `json:"added,omitempty"`
	Removed []*PathChange `json:"removed,omitempty"`
}

// PathChange represents a change to an API path
type PathChange struct {
	Path     string   `json:"path"`
	Methods  []string `json:"methods,omitempty"`
	Severity Severity `json:"severity"`
	Tags     []string `json:"tags,omitempty"`
	Category Category `json:"category,omitempty"`
}

// SchemaChange tracks changes to a specific schema
type SchemaChange struct {
	SchemaName    string       `json:"schema_name"`
	MaxSeverity   Severity     `json:"max_severity"`
	AddedFields   []*FieldDiff `json:"added_fields,omitempty"`
	RemovedFields []*FieldDiff `json:"removed_fields,omitempty"`
	ChangedFields []*FieldDiff `json:"changed_fields,omitempty"`
}

// FieldDiff describes a field difference
type FieldDiff struct {
	Path          string   `json:"path"`
	Severity      Severity `json:"severity"`
	OldType       string   `json:"old_type,omitempty"`
	NewType       string   `json:"new_type,omitempty"`
	Description   string   `json:"description,omitempty"`
	WasRequired   bool     `json:"was_required,omitempty"`
	IsRequired    bool     `json:"is_required,omitempty"`
	WasDeprecated bool     `json:"was_deprecated,omitempty"`
	IsDeprecated  bool     `json:"is_deprecated,omitempty"`
	ChangeReason  string   `json:"change_reason,omitempty"`
}

// Summary provides high-level statistics
type Summary struct {
	TotalSchemasOld    int `json:"total_schemas_old"`
	TotalSchemasNew    int `json:"total_schemas_new"`
	SchemasAdded       int `json:"schemas_added"`
	SchemasRemoved     int `json:"schemas_removed"`
	SchemasWithChanges int `json:"schemas_with_changes"`
	TotalFieldsAdded   int `json:"total_fields_added"`
	TotalFieldsRemoved int `json:"total_fields_removed"`
	BreakingChanges    int `json:"breaking_changes"`
	DeprecatedChanges  int `json:"deprecated_changes"`
	MinorChanges       int `json:"minor_changes"`
	DocsChanges        int `json:"docs_changes"`
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
		fmt.Printf("  Max Severity: %s\n", report.MaxSeverity)
		fmt.Printf("  Schemas: %d → %d\n", report.Summary.TotalSchemasOld, report.Summary.TotalSchemasNew)
		fmt.Printf("  Added: %d, Removed: %d, Changed: %d\n",
			report.Summary.SchemasAdded, report.Summary.SchemasRemoved, report.Summary.SchemasWithChanges)
		fmt.Printf("  Fields added: %d, removed: %d\n",
			report.Summary.TotalFieldsAdded, report.Summary.TotalFieldsRemoved)
		fmt.Printf("  By severity: breaking=%d, deprecated=%d, minor=%d, docs=%d\n",
			report.Summary.BreakingChanges, report.Summary.DeprecatedChanges,
			report.Summary.MinorChanges, report.Summary.DocsChanges)
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
		MaxSeverity: SeverityNone,
		PathChanges: &PathChanges{},
		Summary:     &Summary{},
	}

	// Compare paths
	comparePaths(oldSpec, newSpec, report)

	// Compare schemas
	if oldSpec.Components != nil && newSpec.Components != nil {
		compareAllSchemas(oldSpec.Components.Schemas, newSpec.Components.Schemas, schemasToCompare, report, verbose)
	}

	// Calculate max severity
	report.MaxSeverity = calculateMaxSeverity(report)

	return report
}

func comparePaths(oldSpec, newSpec *OpenAPISpec, report *DiffReport) {
	oldPaths := make(map[string]interface{})
	newPaths := make(map[string]interface{})

	for path, methods := range oldSpec.Paths {
		oldPaths[path] = methods
	}
	for path, methods := range newSpec.Paths {
		newPaths[path] = methods
	}

	// Initialize category summary
	report.CategorySummary = &CategorySummary{
		Setup:      &CategoryStats{},
		Monitoring: &CategoryStats{},
		Internal:   &CategoryStats{},
	}

	// Find added paths (minor severity)
	for path, methods := range newPaths {
		if _, exists := oldPaths[path]; !exists {
			tags := extractTags(methods)
			category := classifyCategory(tags)
			change := &PathChange{
				Path:     path,
				Severity: SeverityMinor,
				Methods:  extractMethods(methods),
				Tags:     tags,
				Category: category,
			}
			report.PathChanges.Added = append(report.PathChanges.Added, change)
			report.Summary.MinorChanges++

			// Update category stats
			stats := getCategoryStats(report.CategorySummary, category)
			stats.PathsAdded++
			stats.Minor++
		}
	}

	// Find removed paths (breaking severity)
	for path, methods := range oldPaths {
		if _, exists := newPaths[path]; !exists {
			tags := extractTags(methods)
			category := classifyCategory(tags)
			change := &PathChange{
				Path:     path,
				Severity: SeverityBreaking,
				Methods:  extractMethods(methods),
				Tags:     tags,
				Category: category,
			}
			report.PathChanges.Removed = append(report.PathChanges.Removed, change)
			report.Summary.BreakingChanges++

			// Update category stats
			stats := getCategoryStats(report.CategorySummary, category)
			stats.PathsRemoved++
			stats.Breaking++
		}
	}

	// Sort for consistent output
	sort.Slice(report.PathChanges.Added, func(i, j int) bool {
		return report.PathChanges.Added[i].Path < report.PathChanges.Added[j].Path
	})
	sort.Slice(report.PathChanges.Removed, func(i, j int) bool {
		return report.PathChanges.Removed[i].Path < report.PathChanges.Removed[j].Path
	})
}

// extractTags gets tags from a path item's methods
func extractTags(pathItem interface{}) []string {
	tagSet := make(map[string]bool)
	if m, ok := pathItem.(map[string]interface{}); ok {
		for key, val := range m {
			switch key {
			case "get", "post", "put", "delete", "patch", "head", "options":
				if method, ok := val.(map[string]interface{}); ok {
					if tags, ok := method["tags"].([]interface{}); ok {
						for _, t := range tags {
							if tag, ok := t.(string); ok {
								tagSet[tag] = true
							}
						}
					}
				}
			}
		}
	}

	tags := make([]string, 0, len(tagSet))
	for t := range tagSet {
		tags = append(tags, t)
	}
	sort.Strings(tags)
	return tags
}

// classifyCategory determines the category based on tags
func classifyCategory(tags []string) Category {
	for _, tag := range tags {
		if monitoringTags[tag] {
			return CategoryMonitoring
		}
		if internalTags[tag] || strings.Contains(tag, "(internal)") {
			return CategoryInternal
		}
	}
	return CategorySetup
}

// getCategoryStats returns the stats struct for a given category
func getCategoryStats(cs *CategorySummary, cat Category) *CategoryStats {
	switch cat {
	case CategoryMonitoring:
		return cs.Monitoring
	case CategoryInternal:
		return cs.Internal
	default:
		return cs.Setup
	}
}

func extractMethods(pathItem interface{}) []string {
	methods := []string{}
	if m, ok := pathItem.(map[string]interface{}); ok {
		for key := range m {
			switch key {
			case "get", "post", "put", "delete", "patch", "head", "options":
				methods = append(methods, strings.ToUpper(key))
			}
		}
	}
	sort.Strings(methods)
	return methods
}

func compareAllSchemas(oldSchemas, newSchemas map[string]interface{}, schemasToCompare []string, report *DiffReport, verbose bool) {
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
			report.Summary.MinorChanges++
			continue
		}
		if oldExists && !newExists {
			report.Summary.SchemasRemoved++
			report.Summary.BreakingChanges++
			continue
		}

		// Both exist - compare fields
		change := compareSchemas(schemaName, oldSchema, newSchema, verbose)
		if change != nil {
			report.SchemaChanges = append(report.SchemaChanges, change)
			report.Summary.SchemasWithChanges++
			report.Summary.TotalFieldsAdded += len(change.AddedFields)
			report.Summary.TotalFieldsRemoved += len(change.RemovedFields)

			// Count by severity
			for _, f := range change.AddedFields {
				incrementSeverityCount(f.Severity, report)
			}
			for _, f := range change.RemovedFields {
				incrementSeverityCount(f.Severity, report)
			}
			for _, f := range change.ChangedFields {
				incrementSeverityCount(f.Severity, report)
			}
		}
	}
}

func incrementSeverityCount(sev Severity, report *DiffReport) {
	switch sev {
	case SeverityBreaking:
		report.Summary.BreakingChanges++
	case SeverityDeprecated:
		report.Summary.DeprecatedChanges++
	case SeverityMinor:
		report.Summary.MinorChanges++
	case SeverityDocs:
		report.Summary.DocsChanges++
	}
}

func compareSchemas(name string, old, new interface{}, verbose bool) *SchemaChange {
	oldMap, oldOk := old.(map[string]interface{})
	newMap, newOk := new.(map[string]interface{})

	if !oldOk || !newOk {
		return nil
	}

	change := &SchemaChange{
		SchemaName:  name,
		MaxSeverity: SeverityNone,
	}

	// Get properties and required fields
	oldProps := getProperties(oldMap)
	newProps := getProperties(newMap)
	oldRequired := getRequiredFields(oldMap)
	newRequired := getRequiredFields(newMap)

	// Find added fields
	for propName, propVal := range newProps {
		if _, exists := oldProps[propName]; !exists {
			isRequired := newRequired[propName]
			severity := SeverityMinor // Adding optional field
			reason := "New optional field"

			if isRequired {
				severity = SeverityBreaking // Adding required field is breaking
				reason = "New required field"
			}

			change.AddedFields = append(change.AddedFields, &FieldDiff{
				Path:         propName,
				Severity:     severity,
				NewType:      getType(propVal),
				Description:  getDescription(propVal),
				IsRequired:   isRequired,
				IsDeprecated: isDeprecated(propVal),
				ChangeReason: reason,
			})
			change.MaxSeverity = maxSeverity(change.MaxSeverity, severity)
		}
	}

	// Find removed fields
	for propName, propVal := range oldProps {
		if _, exists := newProps[propName]; !exists {
			wasRequired := oldRequired[propName]
			severity := SeverityBreaking // Removing any field is breaking
			reason := "Field removed"

			change.RemovedFields = append(change.RemovedFields, &FieldDiff{
				Path:          propName,
				Severity:      severity,
				OldType:       getType(propVal),
				WasRequired:   wasRequired,
				WasDeprecated: isDeprecated(propVal),
				ChangeReason:  reason,
			})
			change.MaxSeverity = maxSeverity(change.MaxSeverity, severity)
		}
	}

	// Find changed fields
	for propName, newVal := range newProps {
		if oldVal, exists := oldProps[propName]; exists {
			fieldChanges := compareField(propName, oldVal, newVal, oldRequired[propName], newRequired[propName])
			if fieldChanges != nil {
				change.ChangedFields = append(change.ChangedFields, fieldChanges...)
				for _, fc := range fieldChanges {
					change.MaxSeverity = maxSeverity(change.MaxSeverity, fc.Severity)
				}
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
	sort.Slice(change.ChangedFields, func(i, j int) bool {
		return change.ChangedFields[i].Path < change.ChangedFields[j].Path
	})

	// Return nil if no changes
	if len(change.AddedFields) == 0 && len(change.RemovedFields) == 0 && len(change.ChangedFields) == 0 {
		return nil
	}

	return change
}

func compareField(path string, oldVal, newVal interface{}, wasRequired, isRequired bool) []*FieldDiff {
	var changes []*FieldDiff

	oldMap, oldOk := oldVal.(map[string]interface{})
	newMap, newOk := newVal.(map[string]interface{})

	if !oldOk || !newOk {
		return nil
	}

	oldType := getType(oldVal)
	newType := getType(newVal)
	oldDeprecated := isDeprecated(oldVal)
	newDeprecated := isDeprecated(newVal)
	oldDesc := getDescription(oldVal)
	newDesc := getDescription(newVal)

	// Type change
	if oldType != newType {
		severity := SeverityBreaking
		reason := fmt.Sprintf("Type changed: %s → %s", oldType, newType)
		changes = append(changes, &FieldDiff{
			Path:         path,
			Severity:     severity,
			OldType:      oldType,
			NewType:      newType,
			ChangeReason: reason,
		})
	}

	// Required status change
	if wasRequired != isRequired {
		if isRequired && !wasRequired {
			// Optional → Required is breaking
			changes = append(changes, &FieldDiff{
				Path:         path,
				Severity:     SeverityBreaking,
				WasRequired:  wasRequired,
				IsRequired:   isRequired,
				ChangeReason: "Field changed from optional to required",
			})
		} else {
			// Required → Optional is minor (relaxing constraint)
			changes = append(changes, &FieldDiff{
				Path:         path,
				Severity:     SeverityMinor,
				WasRequired:  wasRequired,
				IsRequired:   isRequired,
				ChangeReason: "Field changed from required to optional",
			})
		}
	}

	// Deprecated status change
	if !oldDeprecated && newDeprecated {
		changes = append(changes, &FieldDiff{
			Path:          path,
			Severity:      SeverityDeprecated,
			WasDeprecated: oldDeprecated,
			IsDeprecated:  newDeprecated,
			ChangeReason:  "Field marked as deprecated",
		})
	}

	// Description-only change
	if oldDesc != newDesc && len(changes) == 0 {
		// Only if no other changes detected
		changes = append(changes, &FieldDiff{
			Path:         path,
			Severity:     SeverityDocs,
			ChangeReason: "Description changed",
		})
	}

	// Check for enum changes
	oldEnum := getEnumValues(oldMap)
	newEnum := getEnumValues(newMap)
	if len(oldEnum) > 0 || len(newEnum) > 0 {
		enumChanges := compareEnums(path, oldEnum, newEnum)
		changes = append(changes, enumChanges...)
	}

	return changes
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

func compareEnums(path string, oldEnum, newEnum []string) []*FieldDiff {
	var changes []*FieldDiff

	oldSet := make(map[string]bool)
	newSet := make(map[string]bool)

	for _, v := range oldEnum {
		oldSet[v] = true
	}
	for _, v := range newEnum {
		newSet[v] = true
	}

	// Check for removed enum values (breaking)
	for v := range oldSet {
		if !newSet[v] {
			changes = append(changes, &FieldDiff{
				Path:         path,
				Severity:     SeverityBreaking,
				ChangeReason: fmt.Sprintf("Enum value removed: %q", v),
			})
		}
	}

	// Check for added enum values (minor)
	for v := range newSet {
		if !oldSet[v] {
			changes = append(changes, &FieldDiff{
				Path:         path,
				Severity:     SeverityMinor,
				ChangeReason: fmt.Sprintf("Enum value added: %q", v),
			})
		}
	}

	return changes
}

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

func getType(prop interface{}) string {
	if propMap, ok := prop.(map[string]interface{}); ok {
		if t, ok := propMap["type"].(string); ok {
			return t
		}
		if ref, ok := propMap["$ref"].(string); ok {
			// Extract schema name from ref
			parts := strings.Split(ref, "/")
			if len(parts) > 0 {
				return "ref:" + parts[len(parts)-1]
			}
			return "ref"
		}
		if _, ok := propMap["oneOf"]; ok {
			return "oneOf"
		}
		if _, ok := propMap["anyOf"]; ok {
			return "anyOf"
		}
		if _, ok := propMap["allOf"]; ok {
			return "allOf"
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

func isDeprecated(prop interface{}) bool {
	if propMap, ok := prop.(map[string]interface{}); ok {
		if deprecated, ok := propMap["deprecated"].(bool); ok {
			return deprecated
		}
	}
	return false
}

func maxSeverity(a, b Severity) Severity {
	if SeverityOrder[a] >= SeverityOrder[b] {
		return a
	}
	return b
}

func calculateMaxSeverity(report *DiffReport) Severity {
	max := SeverityNone

	// Check path changes
	for _, p := range report.PathChanges.Added {
		max = maxSeverity(max, p.Severity)
	}
	for _, p := range report.PathChanges.Removed {
		max = maxSeverity(max, p.Severity)
	}

	// Check schema changes
	for _, sc := range report.SchemaChanges {
		max = maxSeverity(max, sc.MaxSeverity)
	}

	return max
}
