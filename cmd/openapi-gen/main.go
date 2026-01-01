// Package main implements the openapi-gen tool for generating Go types from CheckMK OpenAPI specs.
//
// Usage:
//
//	openapi-gen -spec specs/2.4.0p17/openapi.yaml -output generated/go/v2_4_0p17/ -package v2_4_0p17 -resources host,folder,aux_tag
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

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
	Schemas map[string]*Schema `yaml:"schemas"`
}

// Schema represents an OpenAPI schema definition
type Schema struct {
	Type                 string             `yaml:"type"`
	Format               string             `yaml:"format"`
	Properties           map[string]*Schema `yaml:"properties"`
	Items                *Schema            `yaml:"items"`
	Ref                  string             `yaml:"$ref"`
	Description          string             `yaml:"description"`
	Required             []string           `yaml:"required"`
	AdditionalProperties interface{}        `yaml:"additionalProperties"`
	Enum                 []interface{}      `yaml:"enum"`
	OneOf                []*Schema          `yaml:"oneOf"`
	AllOf                []*Schema          `yaml:"allOf"`
	AnyOf                []*Schema          `yaml:"anyOf"`
	Default              interface{}        `yaml:"default"`
	Example              interface{}        `yaml:"example"`
	Minimum              *float64           `yaml:"minimum"`
	Maximum              *float64           `yaml:"maximum"`
	MinLength            *int               `yaml:"minLength"`
	MaxLength            *int               `yaml:"maxLength"`
	Pattern              string             `yaml:"pattern"`
}

// EnumInfo holds information about an enum type to be generated
type EnumInfo struct {
	TypeName    string
	Description string
	Values      []string
	FieldName   string // Original field name this enum was found on
}

// Generator holds the state for code generation
type Generator struct {
	spec          *OpenAPISpec
	packageName   string
	outputDir     string
	version       string
	resources     []string
	schemasToGen  []string
	excludeFields map[string]bool
	enumsFound    map[string]*EnumInfo // Track enums to generate
	fieldsFound   map[string][]string  // Track field names per schema
}

func main() {
	var (
		specPath    = flag.String("spec", "", "Path to OpenAPI YAML spec file")
		outputDir   = flag.String("output", "", "Output directory for generated files")
		packageName = flag.String("package", "generated", "Go package name")
		version     = flag.String("version", "", "CheckMK version (e.g., 2.4.0p17)")
		resources   = flag.String("resources", "", "Comma-separated list of resources (host,folder,aux_tag)")
		schemas     = flag.String("schemas", "", "Comma-separated list of schemas (alternative to -resources)")
		listSchemas = flag.Bool("list-schemas", false, "List all available schemas and exit")
	)
	flag.Parse()

	if *specPath == "" {
		log.Fatal("Error: -spec flag is required")
	}

	gen := &Generator{
		packageName: *packageName,
		outputDir:   *outputDir,
		version:     *version,
		excludeFields: map[string]bool{
			"update_attributes": true,
			"remove_attributes": true,
		},
		enumsFound:  make(map[string]*EnumInfo),
		fieldsFound: make(map[string][]string),
	}

	if err := gen.LoadSpec(*specPath); err != nil {
		log.Fatalf("Failed to load spec: %v", err)
	}

	// List schemas mode
	if *listSchemas {
		gen.listAllSchemas()
		return
	}

	// Determine which schemas to generate
	if *resources != "" {
		gen.resources = strings.Split(*resources, ",")
		gen.schemasToGen = GetSchemasForResources(gen.resources)
	} else if *schemas != "" {
		gen.schemasToGen = strings.Split(*schemas, ",")
	}

	if *outputDir == "" {
		log.Fatal("Error: -output flag is required")
	}

	if err := gen.Generate(); err != nil {
		log.Fatalf("Failed to generate code: %v", err)
	}

	fmt.Printf("Successfully generated types in %s\n", *outputDir)
}

func (g *Generator) LoadSpec(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading spec file: %w", err)
	}

	g.spec = &OpenAPISpec{}
	if err := yaml.Unmarshal(data, g.spec); err != nil {
		return fmt.Errorf("parsing YAML: %w", err)
	}

	// Extract version from spec if not provided
	if g.version == "" {
		if info, ok := g.spec.Info["version"].(string); ok {
			g.version = info
		}
	}

	return nil
}

func (g *Generator) listAllSchemas() {
	if g.spec.Components == nil {
		fmt.Println("No schemas found")
		return
	}

	schemas := make([]string, 0, len(g.spec.Components.Schemas))
	for name := range g.spec.Components.Schemas {
		schemas = append(schemas, name)
	}
	sort.Strings(schemas)

	fmt.Printf("Available schemas (%d):\n", len(schemas))
	for _, name := range schemas {
		fmt.Printf("  - %s\n", name)
	}
}

func (g *Generator) Generate() error {
	// Ensure output directory exists
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Determine which schemas to generate
	schemas := g.schemasToGen
	if len(schemas) == 0 && g.spec.Components != nil {
		// If no specific schemas requested, generate common ones
		schemas = GetSchemasForResources([]string{"host", "folder", "aux_tag"})
	}

	// Filter to only existing schemas
	var existingSchemas []string
	for _, name := range schemas {
		if g.spec.Components != nil && g.spec.Components.Schemas[name] != nil {
			existingSchemas = append(existingSchemas, name)
		} else {
			log.Printf("Warning: schema %s not found in spec", name)
		}
	}
	sort.Strings(existingSchemas)

	// Generate types.gen.go
	if err := g.generateTypesFile(existingSchemas); err != nil {
		return err
	}

	// Generate enums.gen.go
	if err := g.generateEnumsFile(); err != nil {
		return err
	}

	// Generate fields.gen.go
	if err := g.generateFieldsFile(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateTypesFile(schemas []string) error {
	var buf strings.Builder

	// Write header
	g.writeHeader(&buf, "types.gen.go", "Type definitions for CheckMK REST API")

	// Generate structs for each schema
	for _, schemaName := range schemas {
		schema := g.spec.Components.Schemas[schemaName]
		resolved := g.resolveSchema(schema)

		if err := g.generateStruct(&buf, schemaName, resolved); err != nil {
			return fmt.Errorf("generating struct %s: %w", schemaName, err)
		}
		buf.WriteString("\n")
	}

	// Write to file
	outputPath := filepath.Join(g.outputDir, "types.gen.go")
	if err := os.WriteFile(outputPath, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing types file: %w", err)
	}

	return nil
}

func (g *Generator) generateEnumsFile() error {
	if len(g.enumsFound) == 0 {
		return nil
	}

	var buf strings.Builder

	// Write header
	g.writeHeader(&buf, "enums.gen.go", "Enum constants for CheckMK REST API field values")

	// Sort enum names for consistent output
	enumNames := make([]string, 0, len(g.enumsFound))
	for name := range g.enumsFound {
		enumNames = append(enumNames, name)
	}
	sort.Strings(enumNames)

	// Generate each enum
	for _, typeName := range enumNames {
		info := g.enumsFound[typeName]
		g.generateEnum(&buf, info)
		buf.WriteString("\n")
	}

	// Write to file
	outputPath := filepath.Join(g.outputDir, "enums.gen.go")
	if err := os.WriteFile(outputPath, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing enums file: %w", err)
	}

	return nil
}

func (g *Generator) generateFieldsFile() error {
	if len(g.fieldsFound) == 0 {
		return nil
	}

	var buf strings.Builder

	// Write header
	g.writeHeader(&buf, "fields.gen.go", "Field name lists for validation")

	// Sort schema names for consistent output
	schemaNames := make([]string, 0, len(g.fieldsFound))
	for name := range g.fieldsFound {
		schemaNames = append(schemaNames, name)
	}
	sort.Strings(schemaNames)

	// Generate field lists
	for _, schemaName := range schemaNames {
		fields := g.fieldsFound[schemaName]
		if len(fields) == 0 {
			continue
		}

		sort.Strings(fields)
		typeName := toGoTypeName(schemaName)

		buf.WriteString(fmt.Sprintf("// %sFieldNames lists all valid field names for %s.\n", typeName, schemaName))
		buf.WriteString(fmt.Sprintf("// Use for validation or iteration over available fields.\n"))
		buf.WriteString(fmt.Sprintf("var %sFieldNames = []string{\n", typeName))
		for _, field := range fields {
			buf.WriteString(fmt.Sprintf("\t%q,\n", field))
		}
		buf.WriteString("}\n\n")
	}

	// Write to file
	outputPath := filepath.Join(g.outputDir, "fields.gen.go")
	if err := os.WriteFile(outputPath, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing fields file: %w", err)
	}

	return nil
}

func (g *Generator) writeHeader(buf *strings.Builder, filename, description string) {
	buf.WriteString(fmt.Sprintf("// Code generated by openapi-gen from CheckMK %s. DO NOT EDIT.\n", g.version))
	buf.WriteString("//\n")
	buf.WriteString(fmt.Sprintf("// %s\n", description))
	buf.WriteString("//\n")
	buf.WriteString(fmt.Sprintf("// Source: %s\n", filename))
	if len(g.resources) > 0 {
		buf.WriteString(fmt.Sprintf("// Resources: %s\n", strings.Join(g.resources, ", ")))
	}
	buf.WriteString("\n")
	buf.WriteString(fmt.Sprintf("package %s\n\n", g.packageName))
}

func (g *Generator) resolveSchema(schema *Schema) *Schema {
	if schema == nil {
		return nil
	}

	// If this is a reference, resolve it
	if schema.Ref != "" {
		refSchema := g.resolveRef(schema.Ref)
		if refSchema != nil {
			// Copy description from ref if current is empty
			if schema.Description == "" && refSchema.Description != "" {
				schema.Description = refSchema.Description
			}
			return g.resolveSchema(refSchema)
		}
	}

	// Recursively resolve properties
	if schema.Properties != nil {
		for key, prop := range schema.Properties {
			schema.Properties[key] = g.resolveSchema(prop)
		}
	}

	// Resolve array items
	if schema.Items != nil {
		schema.Items = g.resolveSchema(schema.Items)
	}

	return schema
}

func (g *Generator) resolveRef(ref string) *Schema {
	// Parse reference like "#/components/schemas/Host"
	parts := strings.Split(ref, "/")
	if len(parts) != 4 || parts[0] != "#" || parts[1] != "components" || parts[2] != "schemas" {
		return nil
	}

	schemaName := parts[3]
	if g.spec.Components == nil {
		return nil
	}

	return g.spec.Components.Schemas[schemaName]
}

func (g *Generator) generateStruct(buf *strings.Builder, name string, schema *Schema) error {
	if schema == nil {
		return nil
	}

	typeName := toGoTypeName(name)

	// Write struct documentation
	if schema.Description != "" {
		writeDocComment(buf, typeName, schema.Description, "")
	} else {
		buf.WriteString(fmt.Sprintf("// %s represents a CheckMK API type.\n", typeName))
	}

	buf.WriteString(fmt.Sprintf("type %s struct {\n", typeName))

	// Sort properties for consistent output
	var propNames []string
	for propName := range schema.Properties {
		if g.excludeFields[propName] {
			continue
		}
		propNames = append(propNames, propName)
	}
	sort.Strings(propNames)

	// Track fields for this schema
	g.fieldsFound[name] = propNames

	// Generate fields
	for _, propName := range propNames {
		prop := schema.Properties[propName]
		fieldName := toGoFieldName(propName)
		goType := g.schemaToGoType(prop, name, propName)
		jsonTag := propName

		// Check if field is required
		isRequired := false
		for _, req := range schema.Required {
			if req == propName {
				isRequired = true
				break
			}
		}

		if !isRequired {
			jsonTag += ",omitempty"
		}

		// Write field documentation
		if prop.Description != "" {
			writeFieldDocComment(buf, prop.Description, prop.Example)
		}

		buf.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, goType, jsonTag))
	}

	buf.WriteString("}\n")

	return nil
}

func (g *Generator) schemaToGoType(schema *Schema, parentSchema, fieldName string) string {
	if schema == nil {
		return "interface{}"
	}

	// Handle references
	if schema.Ref != "" {
		refSchema := g.resolveRef(schema.Ref)
		if refSchema != nil {
			return g.schemaToGoType(refSchema, parentSchema, fieldName)
		}
		parts := strings.Split(schema.Ref, "/")
		if len(parts) > 0 {
			return toGoTypeName(parts[len(parts)-1])
		}
	}

	// Handle arrays
	if schema.Type == "array" {
		if schema.Items != nil {
			itemType := g.schemaToGoType(schema.Items, parentSchema, fieldName)
			return "[]" + itemType
		}
		return "[]interface{}"
	}

	// Handle objects
	if schema.Type == "object" {
		return "map[string]interface{}"
	}

	// Handle oneOf/anyOf
	if len(schema.OneOf) > 0 || len(schema.AnyOf) > 0 {
		return "interface{}"
	}

	// Handle string enums - generate enum type
	if schema.Type == "string" && len(schema.Enum) > 0 {
		enumTypeName := g.registerEnum(parentSchema, fieldName, schema)
		return enumTypeName
	}

	// Handle primitive types
	switch schema.Type {
	case "string":
		return "string"
	case "integer":
		if schema.Format == "int64" {
			return "int64"
		}
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	default:
		return "interface{}"
	}
}

func (g *Generator) registerEnum(parentSchema, fieldName string, schema *Schema) string {
	// Create a meaningful enum type name
	typeName := toGoTypeName(parentSchema) + toGoTypeName(fieldName)

	// Convert enum values to strings
	var values []string
	for _, v := range schema.Enum {
		if s, ok := v.(string); ok {
			values = append(values, s)
		}
	}

	if len(values) == 0 {
		return "string"
	}

	g.enumsFound[typeName] = &EnumInfo{
		TypeName:    typeName,
		Description: schema.Description,
		Values:      values,
		FieldName:   fieldName,
	}

	return typeName
}

func (g *Generator) generateEnum(buf *strings.Builder, info *EnumInfo) {
	// Type definition
	if info.Description != "" {
		writeDocComment(buf, info.TypeName, info.Description, "")
	} else {
		buf.WriteString(fmt.Sprintf("// %s represents valid values for the %s field.\n", info.TypeName, info.FieldName))
	}
	buf.WriteString(fmt.Sprintf("type %s string\n\n", info.TypeName))

	// Constants
	buf.WriteString("const (\n")
	for _, value := range info.Values {
		constName := info.TypeName + toGoConstName(value)
		buf.WriteString(fmt.Sprintf("\t// %s represents the %q value.\n", constName, value))
		buf.WriteString(fmt.Sprintf("\t%s %s = %q\n", constName, info.TypeName, value))
	}
	buf.WriteString(")\n\n")

	// Valid values function
	buf.WriteString(fmt.Sprintf("// Valid%sValues returns all valid values for %s.\n", info.TypeName, info.TypeName))
	buf.WriteString(fmt.Sprintf("// Use with Terraform validators: stringvalidator.OneOf(Valid%sValues()...)\n", info.TypeName))
	buf.WriteString(fmt.Sprintf("func Valid%sValues() []string {\n", info.TypeName))
	buf.WriteString("\treturn []string{\n")
	for _, value := range info.Values {
		buf.WriteString(fmt.Sprintf("\t\tstring(%s%s),\n", info.TypeName, toGoConstName(value)))
	}
	buf.WriteString("\t}\n")
	buf.WriteString("}\n")
}

// Helper functions

func toGoTypeName(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == '.'
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			result.WriteString(part[1:])
		}
	}

	return result.String()
}

func toGoFieldName(s string) string {
	return toGoTypeName(s)
}

func toGoConstName(s string) string {
	// Convert value like "cmk-agent" to "CmkAgent"
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			result.WriteString(strings.ToLower(part[1:]))
		}
	}

	name := result.String()
	if name == "" {
		return "Unknown"
	}
	return name
}

func writeDocComment(buf *strings.Builder, typeName, description, extra string) {
	description = sanitizeComment(description)

	// Split into lines for proper formatting
	lines := strings.Split(description, ". ")
	buf.WriteString(fmt.Sprintf("// %s %s\n", typeName, lines[0]))
	for i := 1; i < len(lines); i++ {
		if lines[i] != "" {
			buf.WriteString(fmt.Sprintf("// %s\n", lines[i]))
		}
	}
	if extra != "" {
		buf.WriteString(fmt.Sprintf("// %s\n", extra))
	}
}

func writeFieldDocComment(buf *strings.Builder, description string, example interface{}) {
	description = sanitizeComment(description)
	buf.WriteString(fmt.Sprintf("\t// %s\n", description))
	if example != nil {
		buf.WriteString(fmt.Sprintf("\t// Example: %v\n", example))
	}
}

func sanitizeComment(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")

	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}

	s = strings.TrimSpace(s)

	if len(s) > 300 {
		s = s[:297] + "..."
	}

	return s
}
