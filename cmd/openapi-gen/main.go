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
	// Additional OpenAPI properties for enhanced metadata
	ReadOnly   bool   `yaml:"readOnly"`
	WriteOnly  bool   `yaml:"writeOnly"`
	Deprecated bool   `yaml:"deprecated"`
	Nullable   bool   `yaml:"nullable"`
	Title      string `yaml:"title"`
}

// FieldMetadata holds comprehensive metadata about a field
type FieldMetadata struct {
	Name        string
	Type        string // Go type or OpenAPI type
	Description string
	ReadOnly    bool
	WriteOnly   bool
	Required    bool
	Deprecated  bool
	Nullable    bool
	Example     interface{}
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
	spec           *OpenAPISpec
	packageName    string
	outputDir      string
	version        string
	resources      []string
	schemasToGen   []string
	excludeFields  map[string]bool
	enumsFound     map[string]*EnumInfo          // Track enums to generate
	fieldsFound    map[string][]string           // Track field names per schema
	fieldsMeta     map[string][]FieldMetadata    // Track detailed field metadata per schema
	requiredFound   map[string][]string           // Track required fields per schema
	readOnlyFound   map[string][]string           // Track readOnly fields per schema
	deprecatedFound map[string][]string           // Track deprecated fields per schema
	descriptions    map[string]map[string]string  // Track field descriptions per schema
	fieldTypes      map[string]map[string]string  // Track field types per schema
	generatedTypes  map[string]bool               // Track which types were generated
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
		enumsFound:      make(map[string]*EnumInfo),
		fieldsFound:     make(map[string][]string),
		fieldsMeta:      make(map[string][]FieldMetadata),
		requiredFound:   make(map[string][]string),
		readOnlyFound:   make(map[string][]string),
		deprecatedFound: make(map[string][]string),
		descriptions:    make(map[string]map[string]string),
		fieldTypes:      make(map[string]map[string]string),
		generatedTypes:  make(map[string]bool),
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

	// Generate fields.gen.go (includes required and readOnly lists)
	if err := g.generateFieldsFile(); err != nil {
		return err
	}

	// Generate metadata.gen.go (descriptions, types, etc.)
	if err := g.generateMetadataFile(); err != nil {
		return err
	}

	// Generate requests.gen.go (request builders)
	if err := g.generateRequestsFile(); err != nil {
		return err
	}

	// Generate mappings.gen.go (API→Terraform field mappings)
	if err := g.generateMappingsFile(); err != nil {
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

		// Track generated type
		g.generatedTypes[schemaName] = true
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

		// All field names
		buf.WriteString(fmt.Sprintf("// %sFieldNames lists all valid field names for %s.\n", typeName, schemaName))
		buf.WriteString(fmt.Sprintf("// Use for validation or iteration over available fields.\n"))
		buf.WriteString(fmt.Sprintf("var %sFieldNames = []string{\n", typeName))
		for _, field := range fields {
			buf.WriteString(fmt.Sprintf("\t%q,\n", field))
		}
		buf.WriteString("}\n\n")

		// Required field names
		if requiredFields := g.requiredFound[schemaName]; len(requiredFields) > 0 {
			sort.Strings(requiredFields)
			buf.WriteString(fmt.Sprintf("// %sRequiredFieldNames lists required fields for %s.\n", typeName, schemaName))
			buf.WriteString(fmt.Sprintf("// Use for validation before API calls.\n"))
			buf.WriteString(fmt.Sprintf("var %sRequiredFieldNames = []string{\n", typeName))
			for _, field := range requiredFields {
				buf.WriteString(fmt.Sprintf("\t%q,\n", field))
			}
			buf.WriteString("}\n\n")
		}

		// ReadOnly field names (computed fields)
		readOnlyFields := g.readOnlyFound[schemaName]
		if len(readOnlyFields) > 0 {
			sort.Strings(readOnlyFields)
			buf.WriteString(fmt.Sprintf("// %sReadOnlyFieldNames lists read-only (computed) fields for %s.\n", typeName, schemaName))
			buf.WriteString(fmt.Sprintf("// These fields cannot be set in requests and are only returned in responses.\n"))
			buf.WriteString(fmt.Sprintf("var %sReadOnlyFieldNames = []string{\n", typeName))
			for _, field := range readOnlyFields {
				buf.WriteString(fmt.Sprintf("\t%q,\n", field))
			}
			buf.WriteString("}\n\n")
		}

		// Deprecated field names
		deprecatedFields := g.deprecatedFound[schemaName]
		if len(deprecatedFields) > 0 {
			sort.Strings(deprecatedFields)
			buf.WriteString(fmt.Sprintf("// %sDeprecatedFieldNames lists deprecated fields for %s.\n", typeName, schemaName))
			buf.WriteString(fmt.Sprintf("// These fields should not be used and may be removed in future versions.\n"))
			buf.WriteString(fmt.Sprintf("var %sDeprecatedFieldNames = []string{\n", typeName))
			for _, field := range deprecatedFields {
				buf.WriteString(fmt.Sprintf("\t%q,\n", field))
			}
			buf.WriteString("}\n\n")
		}

		// CompareKey field names (all fields minus read-only fields)
		// Only generate for *Attribute schemas used for create/update
		if strings.HasSuffix(schemaName, "Attribute") || strings.HasPrefix(schemaName, "Create") {
			compareFields := make([]string, 0, len(fields))
			readOnlySet := make(map[string]bool)
			for _, f := range readOnlyFields {
				readOnlySet[f] = true
			}
			// Also exclude common metadata/computed fields that shouldn't be compared
			excludedFields := map[string]bool{
				"meta_data":              true,
				"network_scan_result":    true,
				"links":                  true,
				"members":                true,
				"domainType":             true,
				"inventory_failed":       true, // Computed by discovery
				"waiting_for_discovery":  true, // Computed by discovery queue
				"locked_by":              true, // Computed by locking system
				"locked_attributes":      true, // Computed by locking system
				"effective_attributes":   true, // Computed from inheritance
				"is_cluster":             true, // Computed from nodes
				"is_offline":             true, // Computed from monitoring
			}
			for _, f := range fields {
				if !readOnlySet[f] && !excludedFields[f] {
					compareFields = append(compareFields, f)
				}
			}
			if len(compareFields) > 0 {
				sort.Strings(compareFields)
				buf.WriteString(fmt.Sprintf("// %sCompareKeyFields lists fields used for comparison/hashing in %s.\n", typeName, schemaName))
				buf.WriteString(fmt.Sprintf("// Excludes read-only and metadata fields.\n"))
				buf.WriteString(fmt.Sprintf("var %sCompareKeyFields = []string{\n", typeName))
				for _, field := range compareFields {
					buf.WriteString(fmt.Sprintf("\t%q,\n", field))
				}
				buf.WriteString("}\n\n")
			}
		}
	}

	// Write to file
	outputPath := filepath.Join(g.outputDir, "fields.gen.go")
	if err := os.WriteFile(outputPath, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing fields file: %w", err)
	}

	return nil
}

func (g *Generator) generateMetadataFile() error {
	if len(g.descriptions) == 0 && len(g.fieldTypes) == 0 {
		return nil
	}

	var buf strings.Builder

	// Write header
	g.writeHeader(&buf, "metadata.gen.go", "Field metadata including descriptions and types")

	// Sort schema names for consistent output
	schemaNames := make([]string, 0, len(g.descriptions))
	for name := range g.descriptions {
		schemaNames = append(schemaNames, name)
	}
	sort.Strings(schemaNames)

	// Generate descriptions map
	buf.WriteString("// FieldDescriptions maps schema names to field descriptions.\n")
	buf.WriteString("// Use for generating documentation or user-facing error messages.\n")
	buf.WriteString("var FieldDescriptions = map[string]map[string]string{\n")
	for _, schemaName := range schemaNames {
		descriptions := g.descriptions[schemaName]
		if len(descriptions) == 0 {
			continue
		}

		buf.WriteString(fmt.Sprintf("\t%q: {\n", schemaName))

		// Sort field names for consistent output
		fieldNames := make([]string, 0, len(descriptions))
		for name := range descriptions {
			fieldNames = append(fieldNames, name)
		}
		sort.Strings(fieldNames)

		for _, fieldName := range fieldNames {
			desc := sanitizeComment(descriptions[fieldName])
			buf.WriteString(fmt.Sprintf("\t\t%q: %q,\n", fieldName, desc))
		}
		buf.WriteString("\t},\n")
	}
	buf.WriteString("}\n\n")

	// Generate field types map
	buf.WriteString("// FieldTypes maps schema names to field OpenAPI types.\n")
	buf.WriteString("// Use for type coercion and validation.\n")
	buf.WriteString("var FieldTypes = map[string]map[string]string{\n")
	for _, schemaName := range schemaNames {
		types := g.fieldTypes[schemaName]
		if len(types) == 0 {
			continue
		}

		buf.WriteString(fmt.Sprintf("\t%q: {\n", schemaName))

		// Sort field names for consistent output
		fieldNames := make([]string, 0, len(types))
		for name := range types {
			fieldNames = append(fieldNames, name)
		}
		sort.Strings(fieldNames)

		for _, fieldName := range fieldNames {
			fieldType := types[fieldName]
			buf.WriteString(fmt.Sprintf("\t\t%q: %q,\n", fieldName, fieldType))
		}
		buf.WriteString("\t},\n")
	}
	buf.WriteString("}\n\n")

	// Generate helper functions
	buf.WriteString("// GetFieldDescription returns the description for a field in a schema.\n")
	buf.WriteString("// Returns empty string if not found.\n")
	buf.WriteString("func GetFieldDescription(schemaName, fieldName string) string {\n")
	buf.WriteString("\tif fields, ok := FieldDescriptions[schemaName]; ok {\n")
	buf.WriteString("\t\treturn fields[fieldName]\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn \"\"\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// GetFieldType returns the OpenAPI type for a field in a schema.\n")
	buf.WriteString("// Returns empty string if not found.\n")
	buf.WriteString("func GetFieldType(schemaName, fieldName string) string {\n")
	buf.WriteString("\tif fields, ok := FieldTypes[schemaName]; ok {\n")
	buf.WriteString("\t\treturn fields[fieldName]\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn \"\"\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// IsReadOnlyField checks if a field is read-only for a given schema.\n")
	buf.WriteString("func IsReadOnlyField(schemaName, fieldName string) bool {\n")
	buf.WriteString("\tswitch schemaName {\n")
	for _, schemaName := range schemaNames {
		if readOnlyFields := g.readOnlyFound[schemaName]; len(readOnlyFields) > 0 {
			typeName := toGoTypeName(schemaName)
			buf.WriteString(fmt.Sprintf("\tcase %q:\n", schemaName))
			buf.WriteString(fmt.Sprintf("\t\tfor _, f := range %sReadOnlyFieldNames {\n", typeName))
			buf.WriteString("\t\t\tif f == fieldName {\n")
			buf.WriteString("\t\t\t\treturn true\n")
			buf.WriteString("\t\t\t}\n")
			buf.WriteString("\t\t}\n")
		}
	}
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn false\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// IsRequiredField checks if a field is required for a given schema.\n")
	buf.WriteString("func IsRequiredField(schemaName, fieldName string) bool {\n")
	buf.WriteString("\tswitch schemaName {\n")
	for _, schemaName := range schemaNames {
		if requiredFields := g.requiredFound[schemaName]; len(requiredFields) > 0 {
			typeName := toGoTypeName(schemaName)
			buf.WriteString(fmt.Sprintf("\tcase %q:\n", schemaName))
			buf.WriteString(fmt.Sprintf("\t\tfor _, f := range %sRequiredFieldNames {\n", typeName))
			buf.WriteString("\t\t\tif f == fieldName {\n")
			buf.WriteString("\t\t\t\treturn true\n")
			buf.WriteString("\t\t\t}\n")
			buf.WriteString("\t\t}\n")
		}
	}
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn false\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// IsDeprecatedField checks if a field is deprecated for a given schema.\n")
	buf.WriteString("func IsDeprecatedField(schemaName, fieldName string) bool {\n")
	buf.WriteString("\tswitch schemaName {\n")
	for _, schemaName := range schemaNames {
		if deprecatedFields := g.deprecatedFound[schemaName]; len(deprecatedFields) > 0 {
			typeName := toGoTypeName(schemaName)
			buf.WriteString(fmt.Sprintf("\tcase %q:\n", schemaName))
			buf.WriteString(fmt.Sprintf("\t\tfor _, f := range %sDeprecatedFieldNames {\n", typeName))
			buf.WriteString("\t\t\tif f == fieldName {\n")
			buf.WriteString("\t\t\t\treturn true\n")
			buf.WriteString("\t\t\t}\n")
			buf.WriteString("\t\t}\n")
		}
	}
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn false\n")
	buf.WriteString("}\n")

	// Write to file
	outputPath := filepath.Join(g.outputDir, "metadata.gen.go")
	if err := os.WriteFile(outputPath, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing metadata file: %w", err)
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

	// Initialize metadata maps for this schema
	if g.descriptions[name] == nil {
		g.descriptions[name] = make(map[string]string)
	}
	if g.fieldTypes[name] == nil {
		g.fieldTypes[name] = make(map[string]string)
	}

	var requiredFields []string
	var readOnlyFields []string
	var deprecatedFields []string
	var fieldMetadata []FieldMetadata

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

		if isRequired {
			requiredFields = append(requiredFields, propName)
		} else {
			jsonTag += ",omitempty"
		}

		// Track readOnly fields
		if prop.ReadOnly {
			readOnlyFields = append(readOnlyFields, propName)
		}

		// Track deprecated fields
		if prop.Deprecated {
			deprecatedFields = append(deprecatedFields, propName)
		}

		// Track descriptions
		if prop.Description != "" {
			g.descriptions[name][propName] = prop.Description
		}

		// Track field types (OpenAPI type, not Go type)
		openAPIType := prop.Type
		if openAPIType == "" && len(prop.Enum) > 0 {
			openAPIType = "enum"
		}
		if openAPIType == "" && prop.Ref != "" {
			openAPIType = "object"
		}
		g.fieldTypes[name][propName] = openAPIType

		// Collect comprehensive field metadata
		meta := FieldMetadata{
			Name:        propName,
			Type:        openAPIType,
			Description: prop.Description,
			ReadOnly:    prop.ReadOnly,
			WriteOnly:   prop.WriteOnly,
			Required:    isRequired,
			Deprecated:  prop.Deprecated,
			Nullable:    prop.Nullable,
			Example:     prop.Example,
		}
		fieldMetadata = append(fieldMetadata, meta)

		// Write field documentation
		if prop.Description != "" {
			writeFieldDocComment(buf, prop.Description, prop.Example)
		}

		buf.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, goType, jsonTag))
	}

	buf.WriteString("}\n")

	// Store collected metadata
	g.requiredFound[name] = requiredFields
	g.readOnlyFound[name] = readOnlyFields
	g.deprecatedFound[name] = deprecatedFields
	g.fieldsMeta[name] = fieldMetadata

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

func (g *Generator) generateMappingsFile() error {
	var buf strings.Builder

	// Write header
	g.writeHeader(&buf, "mappings.gen.go", "API to Terraform field mappings for import state")

	// Define mappings for common response types
	// These map API response paths to Terraform state field names
	mappings := []struct {
		TypeName    string
		TFField     string
		APIPath     []string
		Description string
	}{
		// Host mappings
		{"HostConfig", "host_name", []string{"id"}, "Host name from API id field"},
		{"HostConfig", "folder", []string{"extensions", "folder"}, "Folder path from extensions"},
		{"HostConfig", "attributes", []string{"extensions", "attributes"}, "Host attributes from extensions"},

		// Folder mappings
		{"Folder", "name", []string{"id"}, "Folder ID/name"},
		{"Folder", "title", []string{"title"}, "Folder title"},
		{"Folder", "parent", []string{"extensions", "path"}, "Parent path from extensions"},
		{"Folder", "attributes", []string{"extensions", "attributes"}, "Folder attributes from extensions"},
	}

	// Group mappings by type
	typeMap := make(map[string][]struct {
		TFField     string
		APIPath     []string
		Description string
	})

	for _, m := range mappings {
		if _, exists := g.generatedTypes[m.TypeName]; !exists {
			continue
		}
		typeMap[m.TypeName] = append(typeMap[m.TypeName], struct {
			TFField     string
			APIPath     []string
			Description string
		}{m.TFField, m.APIPath, m.Description})
	}

	// Generate mapping variables and functions for each type
	for typeName, fields := range typeMap {
		// Generate mapping variable
		buf.WriteString(fmt.Sprintf("// %sFieldMappings maps Terraform field names to API response paths.\n", typeName))
		buf.WriteString(fmt.Sprintf("var %sFieldMappings = map[string][]string{\n", typeName))
		for _, f := range fields {
			pathStr := "\"" + strings.Join(f.APIPath, "\", \"") + "\""
			buf.WriteString(fmt.Sprintf("\t%q: {%s}, // %s\n", f.TFField, pathStr, f.Description))
		}
		buf.WriteString("}\n\n")

		// Generate extraction function
		buf.WriteString(fmt.Sprintf("// Extract%sField extracts a Terraform field value from a %s API response.\n", typeName, typeName))
		buf.WriteString(fmt.Sprintf("// Returns nil if the path doesn't exist.\n"))
		buf.WriteString(fmt.Sprintf("func Extract%sField(response map[string]interface{}, tfField string) interface{} {\n", typeName))
		buf.WriteString(fmt.Sprintf("\tpath, ok := %sFieldMappings[tfField]\n", typeName))
		buf.WriteString("\tif !ok {\n")
		buf.WriteString("\t\treturn nil\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn extractNestedField(response, path)\n")
		buf.WriteString("}\n\n")
	}

	// Generate helper function for nested field extraction
	buf.WriteString("// extractNestedField extracts a value from a nested map using a path.\n")
	buf.WriteString("func extractNestedField(data map[string]interface{}, path []string) interface{} {\n")
	buf.WriteString("\tif len(path) == 0 {\n")
	buf.WriteString("\t\treturn data\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\n")
	buf.WriteString("\tcurrent := data\n")
	buf.WriteString("\tfor i, key := range path {\n")
	buf.WriteString("\t\tval, ok := current[key]\n")
	buf.WriteString("\t\tif !ok {\n")
	buf.WriteString("\t\t\treturn nil\n")
	buf.WriteString("\t\t}\n")
	buf.WriteString("\n")
	buf.WriteString("\t\t// If this is the last key, return the value\n")
	buf.WriteString("\t\tif i == len(path)-1 {\n")
	buf.WriteString("\t\t\treturn val\n")
	buf.WriteString("\t\t}\n")
	buf.WriteString("\n")
	buf.WriteString("\t\t// Otherwise, it must be a map to continue\n")
	buf.WriteString("\t\tif nested, ok := val.(map[string]interface{}); ok {\n")
	buf.WriteString("\t\t\tcurrent = nested\n")
	buf.WriteString("\t\t} else {\n")
	buf.WriteString("\t\t\treturn nil\n")
	buf.WriteString("\t\t}\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn nil\n")
	buf.WriteString("}\n")

	// Write to file
	outputPath := filepath.Join(g.outputDir, "mappings.gen.go")
	if err := os.WriteFile(outputPath, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing mappings file: %w", err)
	}

	return nil
}

func (g *Generator) generateRequestsFile() error {
	var buf strings.Builder

	// Write header
	g.writeHeader(&buf, "requests.gen.go", "Request builder functions for type-safe API calls")

	// Add encoding/json import
	buf.WriteString("import \"encoding/json\"\n\n")

	// Generate request builders for Create* types
	requestTypes := []struct {
		Name       string
		AttrType   string
		AttrSchema string
	}{
		{"CreateHost", "HostCreateAttribute", "HostCreateAttribute"},
		{"CreateClusterHost", "HostCreateAttribute", "HostCreateAttribute"},
		{"CreateFolder", "FolderCreateAttribute", "FolderCreateAttribute"},
		{"UpdateHost", "HostUpdateAttribute", "HostUpdateAttribute"},
		{"UpdateFolder", "FolderUpdateAttribute", "FolderUpdateAttribute"},
	}

	for _, rt := range requestTypes {
		// Check if this type exists in our generated schemas
		if _, exists := g.generatedTypes[rt.Name]; !exists {
			continue
		}

		typeName := toGoTypeName(rt.Name)

		// Build request from map
		buf.WriteString(fmt.Sprintf("// Build%sFromMap creates a %s from a map of attributes.\n", typeName, typeName))
		buf.WriteString(fmt.Sprintf("// This allows converting Terraform attribute maps to typed API requests.\n"))
		buf.WriteString(fmt.Sprintf("func Build%sFromMap(data map[string]interface{}) (*%s, error) {\n", typeName, typeName))
		buf.WriteString("\tjsonData, err := json.Marshal(data)\n")
		buf.WriteString("\tif err != nil {\n")
		buf.WriteString("\t\treturn nil, err\n")
		buf.WriteString("\t}\n")
		buf.WriteString(fmt.Sprintf("\tvar req %s\n", typeName))
		buf.WriteString("\tif err := json.Unmarshal(jsonData, &req); err != nil {\n")
		buf.WriteString("\t\treturn nil, err\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn &req, nil\n")
		buf.WriteString("}\n\n")

		// ToMap converts typed request to map
		buf.WriteString(fmt.Sprintf("// %sToMap converts a %s to a map for API calls.\n", typeName, typeName))
		buf.WriteString(fmt.Sprintf("func %sToMap(req *%s) (map[string]interface{}, error) {\n", typeName, typeName))
		buf.WriteString("\tif req == nil {\n")
		buf.WriteString("\t\treturn nil, nil\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\tjsonData, err := json.Marshal(req)\n")
		buf.WriteString("\tif err != nil {\n")
		buf.WriteString("\t\treturn nil, err\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\tvar result map[string]interface{}\n")
		buf.WriteString("\tif err := json.Unmarshal(jsonData, &result); err != nil {\n")
		buf.WriteString("\t\treturn nil, err\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn result, nil\n")
		buf.WriteString("}\n\n")
	}

	// Generate response parsers for View/Config types
	responseTypes := []string{
		"HostConfig",
		"Folder",
		"HostConfigCollection",
		"FolderCollection",
	}

	for _, rt := range responseTypes {
		if _, exists := g.generatedTypes[rt]; !exists {
			continue
		}

		typeName := toGoTypeName(rt)

		// Parse response from JSON
		buf.WriteString(fmt.Sprintf("// Parse%sFromJSON parses a JSON response into a %s.\n", typeName, typeName))
		buf.WriteString(fmt.Sprintf("func Parse%sFromJSON(data []byte) (*%s, error) {\n", typeName, typeName))
		buf.WriteString(fmt.Sprintf("\tvar resp %s\n", typeName))
		buf.WriteString("\tif err := json.Unmarshal(data, &resp); err != nil {\n")
		buf.WriteString("\t\treturn nil, err\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn &resp, nil\n")
		buf.WriteString("}\n\n")

		// Parse response from map
		buf.WriteString(fmt.Sprintf("// Parse%sFromMap parses a map into a %s.\n", typeName, typeName))
		buf.WriteString(fmt.Sprintf("func Parse%sFromMap(data map[string]interface{}) (*%s, error) {\n", typeName, typeName))
		buf.WriteString("\tjsonData, err := json.Marshal(data)\n")
		buf.WriteString("\tif err != nil {\n")
		buf.WriteString("\t\treturn nil, err\n")
		buf.WriteString("\t}\n")
		buf.WriteString(fmt.Sprintf("\treturn Parse%sFromJSON(jsonData)\n", typeName))
		buf.WriteString("}\n\n")
	}

	// Write to file
	outputPath := filepath.Join(g.outputDir, "requests.gen.go")
	if err := os.WriteFile(outputPath, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing requests file: %w", err)
	}

	return nil
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
