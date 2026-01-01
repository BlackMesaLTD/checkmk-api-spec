// Package types provides version-aware type information for CheckMK API.
//
// Integration tests that verify generated types match live API responses.
// Run with: go test -tags=integration ./generated/go/...
//
//go:build integration

package types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// TestConfig holds test environment configuration
type TestConfig struct {
	URL      string
	Username string
	Password string
	Version  string
}

func getTestConfig(t *testing.T) *TestConfig {
	url := os.Getenv("CHECKMK_URL")
	if url == "" {
		t.Skip("CHECKMK_URL not set, skipping integration test")
	}

	return &TestConfig{
		URL:      url,
		Username: getEnvOrDefault("CHECKMK_USERNAME", "automation"),
		Password: getEnvOrDefault("CHECKMK_PASSWORD", ""),
		Version:  os.Getenv("CHECKMK_VERSION"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// OpenAPISpec represents the relevant parts of the OpenAPI specification
type OpenAPISpec struct {
	Components struct {
		Schemas map[string]SchemaDefinition `yaml:"schemas"`
	} `yaml:"components"`
}

type SchemaDefinition struct {
	Properties map[string]PropertyDefinition `yaml:"properties"`
	AllOf      []SchemaRef                   `yaml:"allOf"`
}

type SchemaRef struct {
	Ref        string                        `yaml:"$ref"`
	Properties map[string]PropertyDefinition `yaml:"properties"`
}

type PropertyDefinition struct {
	Enum []string `yaml:"enum"`
	Type string   `yaml:"type"`
}

func fetchOpenAPISpec(cfg *TestConfig) (*OpenAPISpec, error) {
	url := fmt.Sprintf("%s/check_mk/api/1.0/openapi-swagger-ui.yaml", cfg.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(cfg.Username, cfg.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch OpenAPI spec: %d - %s", resp.StatusCode, string(body))
	}

	var spec OpenAPISpec
	if err := yaml.NewDecoder(resp.Body).Decode(&spec); err != nil {
		return nil, fmt.Errorf("failed to parse OpenAPI spec: %w", err)
	}

	return &spec, nil
}

func fetchVersion(cfg *TestConfig) (string, error) {
	url := fmt.Sprintf("%s/check_mk/api/1.0/version", cfg.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(cfg.Username, cfg.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch version: %d", resp.StatusCode)
	}

	var result struct {
		Versions struct {
			CheckMK string `json:"checkmk"`
		} `json:"versions"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Versions.CheckMK, nil
}

// extractEnumValues extracts enum values from a schema property
func extractEnumValues(spec *OpenAPISpec, schemaName, propertyName string) []string {
	schema, ok := spec.Components.Schemas[schemaName]
	if !ok {
		return nil
	}

	// Check direct properties
	if prop, ok := schema.Properties[propertyName]; ok && len(prop.Enum) > 0 {
		return prop.Enum
	}

	// Check allOf schemas
	for _, ref := range schema.AllOf {
		if prop, ok := ref.Properties[propertyName]; ok && len(prop.Enum) > 0 {
			return prop.Enum
		}

		// Follow $ref
		if ref.Ref != "" {
			refName := strings.TrimPrefix(ref.Ref, "#/components/schemas/")
			if refSchema, ok := spec.Components.Schemas[refName]; ok {
				if prop, ok := refSchema.Properties[propertyName]; ok && len(prop.Enum) > 0 {
					return prop.Enum
				}
			}
		}
	}

	return nil
}

// extractFieldNames extracts all property names from a schema
func extractFieldNames(spec *OpenAPISpec, schemaName string) []string {
	schema, ok := spec.Components.Schemas[schemaName]
	if !ok {
		return nil
	}

	fields := make(map[string]bool)

	// Add direct properties
	for name := range schema.Properties {
		fields[name] = true
	}

	// Add properties from allOf
	for _, ref := range schema.AllOf {
		for name := range ref.Properties {
			fields[name] = true
		}

		// Follow $ref
		if ref.Ref != "" {
			refName := strings.TrimPrefix(ref.Ref, "#/components/schemas/")
			if refSchema, ok := spec.Components.Schemas[refName]; ok {
				for name := range refSchema.Properties {
					fields[name] = true
				}
			}
		}
	}

	result := make([]string, 0, len(fields))
	for name := range fields {
		result = append(result, name)
	}
	return result
}

func TestGeneratedMatchesLiveAPI_TagAgent(t *testing.T) {
	cfg := getTestConfig(t)

	// Get version from API
	version, err := fetchVersion(cfg)
	if err != nil {
		t.Fatalf("Failed to fetch version: %v", err)
	}
	t.Logf("Testing against CheckMK version: %s", version)

	// Look up baseline for this version
	baseline := LookupBaseline(version)
	if baseline == "" {
		t.Skipf("No baseline for version %s", version)
	}
	t.Logf("Using baseline: %s", baseline)

	// Get generated values
	generatedValues := ValidHostTagAgentValues(baseline)
	if len(generatedValues) == 0 {
		t.Fatalf("No generated tag_agent values for baseline %s", baseline)
	}
	t.Logf("Generated tag_agent values: %v", generatedValues)

	// Fetch OpenAPI spec
	spec, err := fetchOpenAPISpec(cfg)
	if err != nil {
		t.Fatalf("Failed to fetch OpenAPI spec: %v", err)
	}

	// Extract tag_agent enum from spec
	apiValues := extractEnumValues(spec, "HostCreateAttribute", "tag_agent")
	if len(apiValues) == 0 {
		// Try alternative schema names
		apiValues = extractEnumValues(spec, "CreateHost", "tag_agent")
	}
	if len(apiValues) == 0 {
		t.Logf("Could not find tag_agent enum in API spec, skipping comparison")
		return
	}
	t.Logf("API tag_agent values: %v", apiValues)

	// Compare
	for _, apiVal := range apiValues {
		if !contains(generatedValues, apiVal) {
			t.Errorf("API value %q not in generated values", apiVal)
		}
	}

	for _, genVal := range generatedValues {
		if !contains(apiValues, genVal) {
			t.Errorf("Generated value %q not in API values", genVal)
		}
	}
}

func TestGeneratedMatchesLiveAPI_HostFields(t *testing.T) {
	cfg := getTestConfig(t)

	version, err := fetchVersion(cfg)
	if err != nil {
		t.Fatalf("Failed to fetch version: %v", err)
	}
	t.Logf("Testing against CheckMK version: %s", version)

	baseline := LookupBaseline(version)
	if baseline == "" {
		t.Skipf("No baseline for version %s", version)
	}

	generatedFields := HostCreateAttributeFieldNames(baseline)
	if len(generatedFields) == 0 {
		t.Fatalf("No generated host fields for baseline %s", baseline)
	}
	t.Logf("Generated host field count: %d", len(generatedFields))

	spec, err := fetchOpenAPISpec(cfg)
	if err != nil {
		t.Fatalf("Failed to fetch OpenAPI spec: %v", err)
	}

	apiFields := extractFieldNames(spec, "HostCreateAttribute")
	if len(apiFields) == 0 {
		t.Logf("Could not find HostCreateAttribute schema, skipping comparison")
		return
	}
	t.Logf("API host field count: %d", len(apiFields))

	// Check that all API fields are in generated (generated may have more)
	missingFromGenerated := []string{}
	for _, apiField := range apiFields {
		if !contains(generatedFields, apiField) {
			missingFromGenerated = append(missingFromGenerated, apiField)
		}
	}

	if len(missingFromGenerated) > 0 {
		t.Errorf("API fields missing from generated: %v", missingFromGenerated)
	}
}

func TestGeneratedMatchesLiveAPI_FolderFields(t *testing.T) {
	cfg := getTestConfig(t)

	version, err := fetchVersion(cfg)
	if err != nil {
		t.Fatalf("Failed to fetch version: %v", err)
	}

	baseline := LookupBaseline(version)
	if baseline == "" {
		t.Skipf("No baseline for version %s", version)
	}

	generatedFields := FolderCreateAttributeFieldNames(baseline)
	if len(generatedFields) == 0 {
		t.Fatalf("No generated folder fields for baseline %s", baseline)
	}
	t.Logf("Generated folder field count: %d", len(generatedFields))

	spec, err := fetchOpenAPISpec(cfg)
	if err != nil {
		t.Fatalf("Failed to fetch OpenAPI spec: %v", err)
	}

	apiFields := extractFieldNames(spec, "FolderCreateAttribute")
	if len(apiFields) == 0 {
		t.Logf("Could not find FolderCreateAttribute schema, skipping comparison")
		return
	}
	t.Logf("API folder field count: %d", len(apiFields))

	missingFromGenerated := []string{}
	for _, apiField := range apiFields {
		if !contains(generatedFields, apiField) {
			missingFromGenerated = append(missingFromGenerated, apiField)
		}
	}

	if len(missingFromGenerated) > 0 {
		t.Errorf("API fields missing from generated: %v", missingFromGenerated)
	}
}

func TestBaselineLookupMatchesRunningVersion(t *testing.T) {
	cfg := getTestConfig(t)

	version, err := fetchVersion(cfg)
	if err != nil {
		t.Fatalf("Failed to fetch version: %v", err)
	}

	baseline := LookupBaseline(version)
	t.Logf("Version %s -> Baseline %s", version, baseline)

	if baseline == "" {
		t.Errorf("LookupBaseline returned empty for running version %s", version)
	}

	// Verify baseline returns valid values
	tagValues := ValidHostTagAgentValues(baseline)
	if len(tagValues) == 0 {
		t.Errorf("ValidHostTagAgentValues returned empty for baseline %s", baseline)
	}

	hostFields := HostCreateAttributeFieldNames(baseline)
	if len(hostFields) == 0 {
		t.Errorf("HostCreateAttributeFieldNames returned empty for baseline %s", baseline)
	}

	folderFields := FolderCreateAttributeFieldNames(baseline)
	if len(folderFields) == 0 {
		t.Errorf("FolderCreateAttributeFieldNames returned empty for baseline %s", baseline)
	}
}
