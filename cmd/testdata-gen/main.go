// Package main implements the testdata-gen tool for generating test fixtures from OpenAPI specs.
//
// This tool generates valid JSON test data based on schema constraints including:
// - Required vs optional fields
// - Enum values
// - Type constraints
//
// Usage:
//
//	testdata-gen -spec openapi.yaml -schema CreateHost -output testdata/create_host.json
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

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

// Generator generates test data from schemas
type Generator struct {
	spec     *OpenAPISpec
	rng      *rand.Rand
	minimal  bool // Only generate required fields
	examples bool // Use example values when available
}

func main() {
	var (
		specPath   = flag.String("spec", "", "Path to OpenAPI spec file")
		schemaName = flag.String("schema", "", "Schema name to generate data for")
		output     = flag.String("output", "", "Output JSON file (prints to stdout if not specified)")
		minimal    = flag.Bool("minimal", false, "Only generate required fields")
		seed       = flag.Int64("seed", 0, "Random seed (0 = use current time)")
		examples   = flag.Bool("examples", true, "Use example values from spec when available")
		list       = flag.Bool("list", false, "List available schemas and exit")
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

	if *list {
		listSchemas(spec)
		return
	}

	if *schemaName == "" {
		log.Fatal("Error: -schema flag is required")
	}

	// Initialize generator
	seedVal := *seed
	if seedVal == 0 {
		seedVal = time.Now().UnixNano()
	}

	gen := &Generator{
		spec:     spec,
		rng:      rand.New(rand.NewSource(seedVal)),
		minimal:  *minimal,
		examples: *examples,
	}

	// Generate data
	data, err := gen.Generate(*schemaName)
	if err != nil {
		log.Fatalf("Failed to generate data: %v", err)
	}

	// Output
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	if *output != "" {
		// Ensure directory exists
		if err := os.MkdirAll(filepath.Dir(*output), 0755); err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
		if err := os.WriteFile(*output, jsonData, 0644); err != nil {
			log.Fatalf("Failed to write output: %v", err)
		}
		fmt.Printf("Generated test data written to: %s\n", *output)
	} else {
		fmt.Println(string(jsonData))
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

func listSchemas(spec *OpenAPISpec) {
	if spec.Components == nil {
		fmt.Println("No schemas found")
		return
	}

	fmt.Printf("Available schemas (%d):\n", len(spec.Components.Schemas))
	for name := range spec.Components.Schemas {
		fmt.Printf("  - %s\n", name)
	}
}

func (g *Generator) Generate(schemaName string) (interface{}, error) {
	if g.spec.Components == nil {
		return nil, fmt.Errorf("no components in spec")
	}

	schema, ok := g.spec.Components.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema %q not found", schemaName)
	}

	return g.generateValue(schema, schemaName)
}

func (g *Generator) generateValue(schema interface{}, context string) (interface{}, error) {
	schemaMap, ok := schema.(map[string]interface{})
	if !ok {
		return nil, nil
	}

	// Handle $ref
	if ref, ok := schemaMap["$ref"].(string); ok {
		return g.resolveAndGenerate(ref, context)
	}

	// Check for example value
	if g.examples {
		if example, ok := schemaMap["example"]; ok {
			return example, nil
		}
	}

	// Handle allOf - merge all schemas
	if allOf, ok := schemaMap["allOf"].([]interface{}); ok {
		result := make(map[string]interface{})
		for _, item := range allOf {
			itemVal, err := g.generateValue(item, context)
			if err != nil {
				continue
			}
			if itemMap, ok := itemVal.(map[string]interface{}); ok {
				for k, v := range itemMap {
					result[k] = v
				}
			}
		}
		return result, nil
	}

	// Handle oneOf/anyOf - pick first option
	if oneOf, ok := schemaMap["oneOf"].([]interface{}); ok && len(oneOf) > 0 {
		return g.generateValue(oneOf[0], context)
	}
	if anyOf, ok := schemaMap["anyOf"].([]interface{}); ok && len(anyOf) > 0 {
		return g.generateValue(anyOf[0], context)
	}

	// Handle by type
	schemaType, _ := schemaMap["type"].(string)

	switch schemaType {
	case "object":
		return g.generateObject(schemaMap, context)
	case "array":
		return g.generateArray(schemaMap, context)
	case "string":
		return g.generateString(schemaMap, context)
	case "integer":
		return g.generateInteger(schemaMap)
	case "number":
		return g.generateNumber(schemaMap)
	case "boolean":
		return g.rng.Intn(2) == 1, nil
	default:
		// Try to infer from properties
		if _, ok := schemaMap["properties"]; ok {
			return g.generateObject(schemaMap, context)
		}
		return nil, nil
	}
}

func (g *Generator) resolveAndGenerate(ref, context string) (interface{}, error) {
	// Parse reference like "#/components/schemas/Host"
	if !strings.HasPrefix(ref, "#/components/schemas/") {
		return nil, fmt.Errorf("unsupported ref: %s", ref)
	}

	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")
	schema, ok := g.spec.Components.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema %s not found", schemaName)
	}

	return g.generateValue(schema, schemaName)
}

func (g *Generator) generateObject(schema map[string]interface{}, context string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	properties, _ := schema["properties"].(map[string]interface{})
	required := getRequiredFields(schema)
	requiredSet := make(map[string]bool)
	for _, r := range required {
		requiredSet[r] = true
	}

	for propName, propSchema := range properties {
		// Skip optional fields if minimal mode
		if g.minimal && !requiredSet[propName] {
			continue
		}

		// Include required fields always, optional fields 70% of the time
		if !requiredSet[propName] && g.rng.Float32() > 0.7 {
			continue
		}

		val, err := g.generateValue(propSchema, context+"."+propName)
		if err != nil {
			continue
		}
		if val != nil {
			result[propName] = val
		}
	}

	return result, nil
}

func (g *Generator) generateArray(schema map[string]interface{}, context string) ([]interface{}, error) {
	items, ok := schema["items"]
	if !ok {
		return []interface{}{}, nil
	}

	// Generate 1-3 items
	count := g.rng.Intn(3) + 1

	result := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		val, err := g.generateValue(items, context+"[]")
		if err != nil {
			continue
		}
		if val != nil {
			result = append(result, val)
		}
	}

	return result, nil
}

func (g *Generator) generateString(schema map[string]interface{}, context string) (string, error) {
	// Handle enum
	if enumVals, ok := schema["enum"].([]interface{}); ok && len(enumVals) > 0 {
		idx := g.rng.Intn(len(enumVals))
		if s, ok := enumVals[idx].(string); ok {
			return s, nil
		}
	}

	// Handle format
	format, _ := schema["format"].(string)
	switch format {
	case "date":
		return time.Now().Format("2006-01-02"), nil
	case "date-time":
		return time.Now().Format(time.RFC3339), nil
	case "email":
		return "test@example.com", nil
	case "uri", "url":
		return "https://example.com/test", nil
	case "ipv4":
		return fmt.Sprintf("%d.%d.%d.%d", g.rng.Intn(256), g.rng.Intn(256), g.rng.Intn(256), g.rng.Intn(256)), nil
	case "ipv6":
		return "2001:db8::1", nil
	}

	// Generate based on field name hints
	contextLower := strings.ToLower(context)
	switch {
	case strings.Contains(contextLower, "host_name") || strings.Contains(contextLower, "hostname"):
		return fmt.Sprintf("testhost-%04d.example.com", g.rng.Intn(10000)), nil
	case strings.Contains(contextLower, "folder"):
		folders := []string{"/", "/Production", "/Test", "/Development", "/Production/Servers", "/Production/Network"}
		return folders[g.rng.Intn(len(folders))], nil
	case strings.Contains(contextLower, "alias"):
		return fmt.Sprintf("Test Alias %d", g.rng.Intn(1000)), nil
	case strings.Contains(contextLower, "title"):
		return fmt.Sprintf("Test Title %d", g.rng.Intn(1000)), nil
	case strings.Contains(contextLower, "description") || strings.Contains(contextLower, "comment"):
		return "Generated test description", nil
	case strings.Contains(contextLower, "site"):
		return "test", nil
	case strings.Contains(contextLower, "ip") || strings.Contains(contextLower, "address"):
		return fmt.Sprintf("192.168.%d.%d", g.rng.Intn(256), g.rng.Intn(256)), nil
	case strings.Contains(contextLower, "name") || strings.Contains(contextLower, "id"):
		return fmt.Sprintf("test_%04d", g.rng.Intn(10000)), nil
	case strings.Contains(contextLower, "tag"):
		return fmt.Sprintf("tag_%d", g.rng.Intn(100)), nil
	default:
		return fmt.Sprintf("test_value_%d", g.rng.Intn(1000)), nil
	}
}

func (g *Generator) generateInteger(schema map[string]interface{}) (int64, error) {
	min := int64(0)
	max := int64(100)

	if minVal, ok := schema["minimum"].(float64); ok {
		min = int64(minVal)
	}
	if maxVal, ok := schema["maximum"].(float64); ok {
		max = int64(maxVal)
	}

	if max <= min {
		max = min + 100
	}

	return min + g.rng.Int63n(max-min+1), nil
}

func (g *Generator) generateNumber(schema map[string]interface{}) (float64, error) {
	min := 0.0
	max := 100.0

	if minVal, ok := schema["minimum"].(float64); ok {
		min = minVal
	}
	if maxVal, ok := schema["maximum"].(float64); ok {
		max = maxVal
	}

	return min + g.rng.Float64()*(max-min), nil
}

func getRequiredFields(schema map[string]interface{}) []string {
	required, ok := schema["required"].([]interface{})
	if !ok {
		return nil
	}

	result := make([]string, 0, len(required))
	for _, r := range required {
		if s, ok := r.(string); ok {
			result = append(result, s)
		}
	}
	return result
}
