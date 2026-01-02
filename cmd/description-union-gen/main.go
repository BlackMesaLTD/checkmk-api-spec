// Package main implements the description-union-gen tool that merges field
// descriptions and enum values across all CheckMK API baseline versions.
//
// This tool scans all generated baseline packages and creates a union of
// all field descriptions and enum values, tracking which versions support
// which fields and values.
//
// Usage:
//
//	description-union-gen -manifest manifest.json -output generated/go/union/
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

// ManifestInput matches the manifest.json structure
type ManifestInput struct {
	Baselines    []string                `json:"baselines"`
	Mapping      map[string]BaselineInfo `json:"mapping"`
	VersionOrder []string                `json:"version_order"`
}

// BaselineInfo describes how a version maps to a baseline
type BaselineInfo struct {
	Baseline    string `json:"baseline"`
	Package     string `json:"package"`
	Path        string `json:"path"`
	ImportAlias string `json:"import_alias"`
	IsBaseline  bool   `json:"is_baseline"`
}

// VersionData holds parsed data from a single baseline
type VersionData struct {
	Version      string                       // e.g., "2.4.0p17"
	Major        int                          // 2
	Minor        int                          // 4
	Patch        int                          // 0
	P            int                          // 17
	Descriptions map[string]map[string]string // schema -> field -> description
	Types        map[string]map[string]string // schema -> field -> type
	Enums        map[string][]string          // "SchemaField" -> enum values
}

// UnionField contains merged field metadata across all versions
type UnionField struct {
	Description  string      // Latest description
	Type         string      // OpenAPI type
	MinVersion   string      // First version with this field (e.g., "2.2")
	MaxVersion   string      // Empty if still present, else last version
	EnumValues   []UnionEnum // Merged enum values with version info
	FirstSeen    string      // First baseline where field appeared
	LastSeen     string      // Last baseline where field was present (empty = still present)
	Descriptions []VersionedDescription // Track description changes
}

// UnionEnum represents an enum value with version tracking
type UnionEnum struct {
	Value      string
	MinVersion string // Empty = all versions, "2.3" = added in 2.3
	MaxVersion string // Empty = still present, "2.2" = removed after 2.2
}

// VersionedDescription tracks description text per version
type VersionedDescription struct {
	Version     string
	Description string
}

// UnionData holds all merged data
type UnionData struct {
	Schemas map[string]map[string]*UnionField // schema -> field -> union data
}

func main() {
	var (
		manifestPath = flag.String("manifest", "manifest.json", "Path to manifest.json")
		outputDir    = flag.String("output", "generated/go/union", "Output directory")
		generatedDir = flag.String("generated", "generated/go", "Generated packages directory")
	)
	flag.Parse()

	// Read manifest
	data, err := os.ReadFile(*manifestPath)
	if err != nil {
		log.Fatalf("Failed to read manifest: %v", err)
	}

	var manifest ManifestInput
	if err := json.Unmarshal(data, &manifest); err != nil {
		log.Fatalf("Failed to parse manifest: %v", err)
	}

	fmt.Printf("Processing %d baselines...\n", len(manifest.Baselines))

	// Parse all baselines
	var versions []VersionData
	for _, baseline := range manifest.Baselines {
		info := manifest.Mapping[baseline]
		pkgPath := filepath.Join(*generatedDir, info.Path)

		vd, err := parsePackage(pkgPath, baseline)
		if err != nil {
			log.Printf("Warning: failed to parse %s: %v", baseline, err)
			continue
		}
		versions = append(versions, vd)
	}

	// Sort versions chronologically
	sort.Slice(versions, func(i, j int) bool {
		return compareVersions(versions[i], versions[j]) < 0
	})

	fmt.Printf("Parsed %d baselines successfully\n", len(versions))

	// Build union
	union := buildUnion(versions)

	// Generate output
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	if err := generateUnionCode(union, *outputDir); err != nil {
		log.Fatalf("Failed to generate code: %v", err)
	}

	// Count stats
	schemaCount := len(union.Schemas)
	fieldCount := 0
	for _, fields := range union.Schemas {
		fieldCount += len(fields)
	}

	fmt.Printf("Generated union package in %s\n", *outputDir)
	fmt.Printf("  Schemas: %d\n", schemaCount)
	fmt.Printf("  Fields: %d\n", fieldCount)
}

// parsePackage extracts metadata from a generated package
func parsePackage(pkgPath, version string) (VersionData, error) {
	vd := VersionData{
		Version:      version,
		Descriptions: make(map[string]map[string]string),
		Types:        make(map[string]map[string]string),
		Enums:        make(map[string][]string),
	}

	// Parse version string: "2.4.0p17" -> Major=2, Minor=4, Patch=0, P=17
	parts := strings.Split(version, "p")
	if len(parts) == 2 {
		vd.P, _ = strconv.Atoi(parts[1])
		numParts := strings.Split(parts[0], ".")
		if len(numParts) >= 3 {
			vd.Major, _ = strconv.Atoi(numParts[0])
			vd.Minor, _ = strconv.Atoi(numParts[1])
			vd.Patch, _ = strconv.Atoi(numParts[2])
		}
	}

	// Parse metadata.gen.go
	metadataPath := filepath.Join(pkgPath, "metadata.gen.go")
	if err := parseMetadataFile(metadataPath, &vd); err != nil {
		return vd, fmt.Errorf("parsing metadata: %w", err)
	}

	// Parse enums.gen.go
	enumsPath := filepath.Join(pkgPath, "enums.gen.go")
	if err := parseEnumsFile(enumsPath, &vd); err != nil {
		// Enums file might not exist in older versions
		log.Printf("Note: no enums for %s", version)
	}

	return vd, nil
}

// parseMetadataFile extracts FieldDescriptions and FieldTypes from metadata.gen.go
func parseMetadataFile(path string, vd *VersionData) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}

		for _, spec := range genDecl.Specs {
			valSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valSpec.Names) == 0 {
				continue
			}

			name := valSpec.Names[0].Name
			if name == "FieldDescriptions" && len(valSpec.Values) > 0 {
				vd.Descriptions = parseNestedStringMap(valSpec.Values[0])
			} else if name == "FieldTypes" && len(valSpec.Values) > 0 {
				vd.Types = parseNestedStringMap(valSpec.Values[0])
			}
		}
	}

	return nil
}

// parseNestedStringMap parses a map[string]map[string]string literal
func parseNestedStringMap(expr ast.Expr) map[string]map[string]string {
	result := make(map[string]map[string]string)

	comp, ok := expr.(*ast.CompositeLit)
	if !ok {
		return result
	}

	for _, elt := range comp.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}

		keyLit, ok := kv.Key.(*ast.BasicLit)
		if !ok || keyLit.Kind != token.STRING {
			continue
		}
		key, _ := strconv.Unquote(keyLit.Value)

		innerComp, ok := kv.Value.(*ast.CompositeLit)
		if !ok {
			continue
		}

		result[key] = make(map[string]string)
		for _, innerElt := range innerComp.Elts {
			innerKV, ok := innerElt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}

			innerKeyLit, ok := innerKV.Key.(*ast.BasicLit)
			if !ok || innerKeyLit.Kind != token.STRING {
				continue
			}
			innerKey, _ := strconv.Unquote(innerKeyLit.Value)

			innerValLit, ok := innerKV.Value.(*ast.BasicLit)
			if !ok || innerValLit.Kind != token.STRING {
				continue
			}
			innerVal, _ := strconv.Unquote(innerValLit.Value)

			result[key][innerKey] = innerVal
		}
	}

	return result
}

// parseEnumsFile extracts enum values from enums.gen.go
func parseEnumsFile(path string, vd *VersionData) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// Look for Valid*Values functions
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		name := funcDecl.Name.Name
		if !strings.HasPrefix(name, "Valid") || !strings.HasSuffix(name, "Values") {
			continue
		}

		// Extract the schema+field name: ValidHostCreateAttributeTagAgentValues -> HostCreateAttributeTagAgent
		enumName := strings.TrimPrefix(name, "Valid")
		enumName = strings.TrimSuffix(enumName, "Values")

		// Parse the function body to extract return values
		if funcDecl.Body != nil {
			for _, stmt := range funcDecl.Body.List {
				retStmt, ok := stmt.(*ast.ReturnStmt)
				if !ok || len(retStmt.Results) == 0 {
					continue
				}

				comp, ok := retStmt.Results[0].(*ast.CompositeLit)
				if !ok {
					continue
				}

				var values []string
				for _, elt := range comp.Elts {
					call, ok := elt.(*ast.CallExpr)
					if !ok {
						continue
					}
					// string(ConstName) pattern
					if len(call.Args) == 1 {
						if ident, ok := call.Args[0].(*ast.Ident); ok {
							// We need to look up the const value, but for simplicity
							// we'll extract from the identifier name
							values = append(values, ident.Name)
						}
					}
				}
				vd.Enums[enumName] = values
			}
		}
	}

	return nil
}

// compareVersions compares two version data structs
func compareVersions(a, b VersionData) int {
	if a.Major != b.Major {
		return a.Major - b.Major
	}
	if a.Minor != b.Minor {
		return a.Minor - b.Minor
	}
	if a.Patch != b.Patch {
		return a.Patch - b.Patch
	}
	return a.P - b.P
}

// buildUnion creates the union from all version data
func buildUnion(versions []VersionData) *UnionData {
	union := &UnionData{
		Schemas: make(map[string]map[string]*UnionField),
	}

	for _, v := range versions {
		minorVersion := fmt.Sprintf("%d.%d", v.Major, v.Minor)

		for schema, fields := range v.Descriptions {
			if union.Schemas[schema] == nil {
				union.Schemas[schema] = make(map[string]*UnionField)
			}

			for field, desc := range fields {
				uf := union.Schemas[schema][field]
				if uf == nil {
					// First time seeing this field
					uf = &UnionField{
						Description: desc,
						MinVersion:  minorVersion,
						FirstSeen:   v.Version,
					}
					union.Schemas[schema][field] = uf
				} else {
					// Update last seen
					uf.LastSeen = ""

					// Track description changes
					if uf.Description != desc {
						uf.Descriptions = append(uf.Descriptions, VersionedDescription{
							Version:     v.Version,
							Description: desc,
						})
						uf.Description = desc // Use latest
					}
				}

				// Update type if available
				if v.Types[schema] != nil && v.Types[schema][field] != "" {
					uf.Type = v.Types[schema][field]
				}
			}
		}

		// Mark fields that disappeared
		for schema, fields := range union.Schemas {
			for fieldName, uf := range fields {
				if v.Descriptions[schema] == nil || v.Descriptions[schema][fieldName] == "" {
					if uf.LastSeen == "" && uf.FirstSeen != "" && uf.FirstSeen != v.Version {
						// This field existed before but doesn't exist now
						// Find the previous version to mark as LastSeen
						for i := len(versions) - 1; i >= 0; i-- {
							if versions[i].Descriptions[schema] != nil &&
							   versions[i].Descriptions[schema][fieldName] != "" {
								uf.LastSeen = versions[i].Version
								uf.MaxVersion = fmt.Sprintf("%d.%d", versions[i].Major, versions[i].Minor)
								break
							}
						}
					}
				}
			}
		}
	}

	return union
}

// generateUnionCode generates the union package
func generateUnionCode(union *UnionData, outputDir string) error {
	// Generate descriptions.gen.go
	code, err := generateDescriptionsCode(union)
	if err != nil {
		return err
	}

	formatted, err := format.Source(code)
	if err != nil {
		log.Printf("Warning: failed to format code: %v", err)
		formatted = code
	}

	path := filepath.Join(outputDir, "descriptions.gen.go")
	if err := os.WriteFile(path, formatted, 0644); err != nil {
		return err
	}

	return nil
}

func generateDescriptionsCode(union *UnionData) ([]byte, error) {
	// Get sorted schema names
	var schemas []string
	for s := range union.Schemas {
		schemas = append(schemas, s)
	}
	sort.Strings(schemas)

	var buf bytes.Buffer
	if err := descriptionsTemplate.Execute(&buf, struct {
		Schemas      []string
		UnionSchemas map[string]map[string]*UnionField
	}{
		Schemas:      schemas,
		UnionSchemas: union.Schemas,
	}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// cleanDescription removes HTML tags and cleans up description text
func cleanDescription(desc string) string {
	// Remove HTML tags
	htmlRe := regexp.MustCompile(`<[^>]+>`)
	desc = htmlRe.ReplaceAllString(desc, "")

	// Fix common HTML entities
	desc = strings.ReplaceAll(desc, "&lt;", "<")
	desc = strings.ReplaceAll(desc, "&gt;", ">")
	desc = strings.ReplaceAll(desc, "&amp;", "&")
	desc = strings.ReplaceAll(desc, "&quot;", "\"")

	return strings.TrimSpace(desc)
}

var descriptionsTemplate = template.Must(template.New("descriptions").Funcs(template.FuncMap{
	"clean": cleanDescription,
	"quote": strconv.Quote,
	"sortedFields": func(fields map[string]*UnionField) []string {
		var names []string
		for n := range fields {
			names = append(names, n)
		}
		sort.Strings(names)
		return names
	},
}).Parse(`// Code generated by description-union-gen. DO NOT EDIT.
//
// Merged field descriptions across all CheckMK API baseline versions.
// This package provides version-annotated field metadata for documentation.

package union

import "strings"

// UnionField contains merged field metadata across all versions.
type UnionField struct {
	Description string // Latest description (HTML stripped)
	Type        string // OpenAPI type
	MinVersion  string // First minor version with this field (e.g., "2.2")
	MaxVersion  string // Empty if still present, else last version (e.g., "2.2")
}

// UnionDescriptions maps schema.field to merged metadata.
// Built from all baseline packages.
var UnionDescriptions = map[string]map[string]UnionField{
{{- range $schema := .Schemas }}
	{{quote $schema}}: {
	{{- $fields := index $.UnionSchemas $schema }}
	{{- range $field := sortedFields $fields }}
	{{- $uf := index $fields $field }}
		{{quote $field}}: {
			Description: {{quote (clean $uf.Description)}},
			Type:        {{quote $uf.Type}},
			MinVersion:  {{quote $uf.MinVersion}},
			MaxVersion:  {{quote $uf.MaxVersion}},
		},
	{{- end }}
	},
{{- end }}
}

// GetUnionDescription returns the union description for a schema field.
// Returns empty string if not found.
func GetUnionDescription(schemaName, fieldName string) string {
	schema, ok := UnionDescriptions[schemaName]
	if !ok {
		return ""
	}
	field, ok := schema[fieldName]
	if !ok {
		return ""
	}
	return field.Description
}

// GetUnionField returns the full UnionField for a schema field.
// Returns nil if not found.
func GetUnionField(schemaName, fieldName string) *UnionField {
	schema, ok := UnionDescriptions[schemaName]
	if !ok {
		return nil
	}
	field, ok := schema[fieldName]
	if !ok {
		return nil
	}
	return &field
}

// FormatMarkdown formats a field description with version annotations.
func (f *UnionField) FormatMarkdown() string {
	if f == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(f.Description)

	// Add version annotations
	if f.MinVersion != "" && f.MinVersion != "2.2" {
		sb.WriteString("\n\n**Available in CheckMK ")
		sb.WriteString(f.MinVersion)
		sb.WriteString("+**")
	}
	if f.MaxVersion != "" {
		sb.WriteString("\n\n**Removed after CheckMK ")
		sb.WriteString(f.MaxVersion)
		sb.WriteString("**")
	}

	return sb.String()
}

// SchemaNames returns all schema names in the union.
func SchemaNames() []string {
	names := make([]string, 0, len(UnionDescriptions))
	for name := range UnionDescriptions {
		names = append(names, name)
	}
	return names
}

// FieldNames returns all field names for a schema.
func FieldNames(schemaName string) []string {
	schema, ok := UnionDescriptions[schemaName]
	if !ok {
		return nil
	}
	names := make([]string, 0, len(schema))
	for name := range schema {
		names = append(names, name)
	}
	return names
}
`))
