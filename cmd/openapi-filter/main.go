// Package main implements the openapi-filter tool for filtering OpenAPI specs to specific resources.
//
// This tool extracts only the paths and schemas relevant to specific resources,
// solving the problem of discriminated union types that cause code generation failures.
//
// Usage:
//
//	openapi-filter -resource host -input full-spec.yaml -output filtered-spec.yaml
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// OpenAPISpec represents the OpenAPI specification structure
type OpenAPISpec struct {
	OpenAPI    string                 `yaml:"openapi"`
	Info       map[string]interface{} `yaml:"info"`
	Servers    []interface{}          `yaml:"servers,omitempty"`
	Paths      map[string]interface{} `yaml:"paths"`
	Components *Components            `yaml:"components"`
	Tags       []interface{}          `yaml:"tags,omitempty"`
}

// Components contains reusable schema definitions
type Components struct {
	Schemas         map[string]interface{} `yaml:"schemas,omitempty"`
	SecuritySchemes map[string]interface{} `yaml:"securitySchemes,omitempty"`
	Parameters      map[string]interface{} `yaml:"parameters,omitempty"`
	Responses       map[string]interface{} `yaml:"responses,omitempty"`
}

// ResourcePathPatterns maps resource names to their API path patterns
var ResourcePathPatterns = map[string][]*regexp.Regexp{
	"host": {
		regexp.MustCompile(`/domain-types/host_config`),
		regexp.MustCompile(`/objects/host_config`),
	},
	"folder": {
		regexp.MustCompile(`/domain-types/folder_config`),
		regexp.MustCompile(`/objects/folder_config`),
	},
	"aux_tag": {
		regexp.MustCompile(`/domain-types/aux_tag`),
		regexp.MustCompile(`/objects/aux_tag`),
	},
	"tag_group": {
		regexp.MustCompile(`/domain-types/host_tag_group`),
		regexp.MustCompile(`/objects/host_tag_group`),
	},
	"user": {
		regexp.MustCompile(`/domain-types/user_config`),
		regexp.MustCompile(`/objects/user_config`),
	},
	"contact_group": {
		regexp.MustCompile(`/domain-types/contact_group_config`),
		regexp.MustCompile(`/objects/contact_group_config`),
	},
	"rule": {
		regexp.MustCompile(`/domain-types/rule`),
		regexp.MustCompile(`/objects/rule`),
	},
	"time_period": {
		regexp.MustCompile(`/domain-types/time_period`),
		regexp.MustCompile(`/objects/time_period`),
	},
	"activation": {
		regexp.MustCompile(`/domain-types/activation_run`),
	},
	"version": {
		regexp.MustCompile(`/version`),
	},
}

func main() {
	var (
		inputPath  = flag.String("input", "", "Path to input OpenAPI YAML spec")
		outputPath = flag.String("output", "", "Path to output filtered YAML spec")
		resources  = flag.String("resources", "", "Comma-separated list of resources to include")
		listPaths  = flag.Bool("list-paths", false, "List all paths in the spec and exit")
	)
	flag.Parse()

	if *inputPath == "" {
		log.Fatal("Error: -input flag is required")
	}

	// Load spec
	data, err := os.ReadFile(*inputPath)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	var spec OpenAPISpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// List paths mode
	if *listPaths {
		listAllPaths(&spec)
		return
	}

	if *outputPath == "" {
		log.Fatal("Error: -output flag is required")
	}

	if *resources == "" {
		log.Fatal("Error: -resources flag is required")
	}

	resourceList := strings.Split(*resources, ",")

	// Filter the spec
	filtered := filterSpec(&spec, resourceList)

	// Write output
	output, err := yaml.Marshal(filtered)
	if err != nil {
		log.Fatalf("Failed to marshal filtered spec: %v", err)
	}

	if err := os.WriteFile(*outputPath, output, 0644); err != nil {
		log.Fatalf("Failed to write output: %v", err)
	}

	// Print stats
	fmt.Printf("Filtered spec written to: %s\n", *outputPath)
	fmt.Printf("  Paths: %d (from %d)\n", len(filtered.Paths), len(spec.Paths))
	if filtered.Components != nil && spec.Components != nil {
		fmt.Printf("  Schemas: %d (from %d)\n", len(filtered.Components.Schemas), len(spec.Components.Schemas))
	}
}

func listAllPaths(spec *OpenAPISpec) {
	paths := make([]string, 0, len(spec.Paths))
	for path := range spec.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	fmt.Printf("Available paths (%d):\n", len(paths))
	for _, path := range paths {
		// Try to identify resource
		resource := identifyResource(path)
		if resource != "" {
			fmt.Printf("  %s [%s]\n", path, resource)
		} else {
			fmt.Printf("  %s\n", path)
		}
	}
}

func identifyResource(path string) string {
	for resource, patterns := range ResourcePathPatterns {
		for _, pattern := range patterns {
			if pattern.MatchString(path) {
				return resource
			}
		}
	}
	return ""
}

func filterSpec(spec *OpenAPISpec, resources []string) *OpenAPISpec {
	filtered := &OpenAPISpec{
		OpenAPI: spec.OpenAPI,
		Info:    spec.Info,
		Servers: spec.Servers,
		Paths:   make(map[string]interface{}),
		Components: &Components{
			Schemas:         make(map[string]interface{}),
			SecuritySchemes: spec.Components.SecuritySchemes,
			Parameters:      make(map[string]interface{}),
			Responses:       make(map[string]interface{}),
		},
		Tags: spec.Tags,
	}

	// Collect patterns for requested resources
	var patterns []*regexp.Regexp
	for _, res := range resources {
		if p, ok := ResourcePathPatterns[res]; ok {
			patterns = append(patterns, p...)
		} else {
			log.Printf("Warning: unknown resource %q", res)
		}
	}

	// Always include version endpoint
	patterns = append(patterns, ResourcePathPatterns["version"]...)

	// Filter paths
	for path, ops := range spec.Paths {
		for _, pattern := range patterns {
			if pattern.MatchString(path) {
				filtered.Paths[path] = ops
				break
			}
		}
	}

	// Collect referenced schemas from filtered paths
	referencedSchemas := collectReferencedSchemas(filtered.Paths)

	// Add referenced schemas with their dependencies
	for schemaName := range referencedSchemas {
		addSchemaWithDeps(filtered.Components.Schemas, spec.Components.Schemas, schemaName)
	}

	return filtered
}

func collectReferencedSchemas(paths map[string]interface{}) map[string]bool {
	schemas := make(map[string]bool)

	var findRefs func(interface{})
	findRefs = func(v interface{}) {
		switch val := v.(type) {
		case map[string]interface{}:
			if ref, ok := val["$ref"].(string); ok {
				if strings.HasPrefix(ref, "#/components/schemas/") {
					schemaName := strings.TrimPrefix(ref, "#/components/schemas/")
					schemas[schemaName] = true
				}
			}
			for _, child := range val {
				findRefs(child)
			}
		case []interface{}:
			for _, item := range val {
				findRefs(item)
			}
		}
	}

	findRefs(paths)
	return schemas
}

func addSchemaWithDeps(target, source map[string]interface{}, schemaName string) {
	if _, exists := target[schemaName]; exists {
		return // Already added
	}

	schema, ok := source[schemaName]
	if !ok {
		return // Schema not found
	}

	target[schemaName] = schema

	// Find and add dependencies
	deps := collectReferencedSchemas(map[string]interface{}{"schema": schema})
	for dep := range deps {
		addSchemaWithDeps(target, source, dep)
	}
}
