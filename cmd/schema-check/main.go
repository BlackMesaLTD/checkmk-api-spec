// Package main implements the schema-check tool for validating completeness.
//
// This tool compares generated Go types against OpenAPI schemas to find:
// - Missing fields in Go types
// - Extra fields in Go types not in schema
//
// Usage:
//
//	schema-check -spec openapi.yaml -types types.gen.go -schema HostConfig
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// OpenAPISpec represents the OpenAPI specification structure
type OpenAPISpec struct {
	Components *Components `yaml:"components"`
}

// Components contains reusable schema definitions
type Components struct {
	Schemas map[string]interface{} `yaml:"schemas"`
}

// CoverageReport contains the comparison results
type CoverageReport struct {
	GoType        string   `json:"go_type"`
	OpenAPISchema string   `json:"openapi_schema"`
	TotalFields   int      `json:"total_fields"`
	CoveredFields int      `json:"covered_fields"`
	MissingFields []string `json:"missing_fields,omitempty"`
	ExtraFields   []string `json:"extra_fields,omitempty"`
	Coverage      float64  `json:"coverage_percent"`
}

func main() {
	var (
		specPath   = flag.String("spec", "", "Path to OpenAPI spec file")
		typesPath  = flag.String("types", "", "Path to Go types file or directory")
		schemaName = flag.String("schema", "", "OpenAPI schema name to check")
		goType     = flag.String("gotype", "", "Go type name (defaults to schema name)")
		listFields = flag.Bool("list-fields", false, "List all fields from schema and exit")
	)
	flag.Parse()

	if *specPath == "" {
		log.Fatal("Error: -spec flag is required")
	}

	// Load spec
	spec, err := loadSpec(*specPath)
	if err != nil {
		log.Fatalf("Failed to load spec: %v", err)
	}

	if *listFields {
		listSchemaFields(spec)
		return
	}

	if *schemaName == "" {
		log.Fatal("Error: -schema flag is required")
	}

	if *typesPath == "" {
		log.Fatal("Error: -types flag is required")
	}

	goTypeName := *goType
	if goTypeName == "" {
		goTypeName = *schemaName
	}

	// Get schema fields
	schemaFields, err := getSchemaFields(spec, *schemaName)
	if err != nil {
		log.Fatalf("Failed to get schema fields: %v", err)
	}

	// Get Go struct fields
	goFields, err := getGoStructFields(*typesPath, goTypeName)
	if err != nil {
		log.Fatalf("Failed to get Go struct fields: %v", err)
	}

	// Compare
	report := compare(*schemaName, goTypeName, schemaFields, goFields)

	// Print report
	printReport(report)
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

func listSchemaFields(spec *OpenAPISpec) {
	if spec.Components == nil {
		fmt.Println("No components found")
		return
	}

	schemas := make([]string, 0, len(spec.Components.Schemas))
	for name := range spec.Components.Schemas {
		schemas = append(schemas, name)
	}
	sort.Strings(schemas)

	for _, name := range schemas {
		schema := spec.Components.Schemas[name]
		fields, _ := extractProperties(schema)
		fmt.Printf("%s (%d fields):\n", name, len(fields))
		sort.Strings(fields)
		for _, f := range fields {
			fmt.Printf("  - %s\n", f)
		}
		fmt.Println()
	}
}

func getSchemaFields(spec *OpenAPISpec, schemaName string) ([]string, error) {
	if spec.Components == nil {
		return nil, fmt.Errorf("no components in spec")
	}

	schema, ok := spec.Components.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema %q not found", schemaName)
	}

	fields, err := extractProperties(schema)
	if err != nil {
		return nil, err
	}

	return fields, nil
}

func extractProperties(schema interface{}) ([]string, error) {
	schemaMap, ok := schema.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("schema is not an object")
	}

	var fields []string

	// Get direct properties
	if props, ok := schemaMap["properties"].(map[string]interface{}); ok {
		for name := range props {
			fields = append(fields, name)
		}
	}

	// Handle allOf - merge properties from all schemas
	if allOf, ok := schemaMap["allOf"].([]interface{}); ok {
		for _, item := range allOf {
			subFields, _ := extractProperties(item)
			fields = append(fields, subFields...)
		}
	}

	return fields, nil
}

func getGoStructFields(path, typeName string) ([]string, error) {
	// Check if path is a file or directory
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	var files []string
	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".go") {
				files = append(files, filepath.Join(path, e.Name()))
			}
		}
	} else {
		files = []string{path}
	}

	// Parse Go files and find the struct
	for _, file := range files {
		fields, err := parseGoFile(file, typeName)
		if err == nil && len(fields) > 0 {
			return fields, nil
		}
	}

	return nil, fmt.Errorf("type %q not found in %s", typeName, path)
}

func parseGoFile(path, typeName string) ([]string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var fields []string

	ast.Inspect(f, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if typeSpec.Name.Name != typeName {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range structType.Fields.List {
			// Get JSON tag
			if field.Tag != nil {
				tag := field.Tag.Value
				jsonTag := extractJSONTag(tag)
				if jsonTag != "" && jsonTag != "-" {
					// Remove omitempty suffix
					jsonTag = strings.Split(jsonTag, ",")[0]
					fields = append(fields, jsonTag)
				}
			}
		}

		return false
	})

	return fields, nil
}

func extractJSONTag(tag string) string {
	// Tag format: `json:"field_name,omitempty"`
	tag = strings.Trim(tag, "`")
	parts := strings.Split(tag, " ")

	for _, part := range parts {
		if strings.HasPrefix(part, "json:") {
			value := strings.TrimPrefix(part, "json:")
			value = strings.Trim(value, "\"")
			return value
		}
	}

	return ""
}

func compare(schemaName, goTypeName string, schemaFields, goFields []string) *CoverageReport {
	schemaSet := make(map[string]bool)
	goSet := make(map[string]bool)

	for _, f := range schemaFields {
		schemaSet[f] = true
	}
	for _, f := range goFields {
		goSet[f] = true
	}

	report := &CoverageReport{
		GoType:        goTypeName,
		OpenAPISchema: schemaName,
		TotalFields:   len(schemaFields),
	}

	// Find missing (in schema but not in Go)
	for f := range schemaSet {
		if goSet[f] {
			report.CoveredFields++
		} else {
			report.MissingFields = append(report.MissingFields, f)
		}
	}

	// Find extra (in Go but not in schema)
	for f := range goSet {
		if !schemaSet[f] {
			report.ExtraFields = append(report.ExtraFields, f)
		}
	}

	sort.Strings(report.MissingFields)
	sort.Strings(report.ExtraFields)

	if report.TotalFields > 0 {
		report.Coverage = float64(report.CoveredFields) / float64(report.TotalFields) * 100
	}

	return report
}

func printReport(report *CoverageReport) {
	fmt.Printf("Schema Coverage Report\n")
	fmt.Printf("======================\n\n")
	fmt.Printf("Go Type:        %s\n", report.GoType)
	fmt.Printf("OpenAPI Schema: %s\n", report.OpenAPISchema)
	fmt.Printf("Coverage:       %.1f%% (%d/%d fields)\n\n",
		report.Coverage, report.CoveredFields, report.TotalFields)

	if len(report.MissingFields) > 0 {
		fmt.Printf("Missing fields (in schema but not in Go):\n")
		for _, f := range report.MissingFields {
			fmt.Printf("  - %s\n", f)
		}
		fmt.Println()
	}

	if len(report.ExtraFields) > 0 {
		fmt.Printf("Extra fields (in Go but not in schema):\n")
		for _, f := range report.ExtraFields {
			fmt.Printf("  - %s\n", f)
		}
		fmt.Println()
	}

	if len(report.MissingFields) == 0 && len(report.ExtraFields) == 0 {
		fmt.Printf("✓ Perfect match!\n")
	}
}
