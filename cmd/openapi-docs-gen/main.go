// Package main implements the openapi-docs-gen tool for generating API documentation
// from CheckMK OpenAPI specs.
//
// Usage:
//
//	openapi-docs-gen -manifest manifest.json -format all
//	openapi-docs-gen -spec specs/2.4.0/p17.yaml -output docs/v2_4_0/p17 -format markdown
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// ManifestInput matches the manifest.json structure
type ManifestInput struct {
	Baselines []string                `json:"baselines"`
	Mapping   map[string]BaselineInfo `json:"mapping"`
}

// BaselineInfo describes how a version maps to a baseline
type BaselineInfo struct {
	Spec        string `json:"spec"`
	Baseline    string `json:"baseline"`
	Package     string `json:"package"`
	Path        string `json:"path"`
	ImportAlias string `json:"import_alias"`
	IsBaseline  bool   `json:"is_baseline"`
	MaxSeverity string `json:"max_severity"`
}

// OpenAPISpec represents the OpenAPI specification structure
type OpenAPISpec struct {
	OpenAPI    string                 `yaml:"openapi"`
	Info       Info                   `yaml:"info"`
	Paths      map[string]PathItem    `yaml:"paths"`
	Components *Components            `yaml:"components"`
}

// Info contains API metadata
type Info struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

// PathItem represents operations on a path
type PathItem map[string]*Operation

// Operation represents an HTTP method operation
type Operation struct {
	Tags        []string              `yaml:"tags"`
	Summary     string                `yaml:"summary"`
	Description string                `yaml:"description"`
	OperationID string                `yaml:"operationId"`
	Parameters  []Parameter           `yaml:"parameters"`
	RequestBody *RequestBody          `yaml:"requestBody"`
	Responses   map[string]*Response  `yaml:"responses"`
}

// Parameter represents an API parameter
type Parameter struct {
	Name        string  `yaml:"name"`
	In          string  `yaml:"in"`
	Description string  `yaml:"description"`
	Required    bool    `yaml:"required"`
	Schema      *Schema `yaml:"schema"`
}

// RequestBody represents a request body
type RequestBody struct {
	Description string             `yaml:"description"`
	Required    bool               `yaml:"required"`
	Content     map[string]*Media  `yaml:"content"`
}

// Media represents media type content
type Media struct {
	Schema *Schema `yaml:"schema"`
}

// Response represents an API response
type Response struct {
	Description string            `yaml:"description"`
	Content     map[string]*Media `yaml:"content"`
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
	ReadOnly             bool               `yaml:"readOnly"`
	WriteOnly            bool               `yaml:"writeOnly"`
	Deprecated           bool               `yaml:"deprecated"`
	Nullable             bool               `yaml:"nullable"`
	Title                string             `yaml:"title"`
}

// Generator holds the state for documentation generation
type Generator struct {
	manifest     *ManifestInput
	specsDir     string
	docsOutput   string
	redocOutput  string
	title        string
	baseURL      string
}

// VersionInfo holds metadata about a version for templates
type VersionInfo struct {
	Version       string
	Path          string
	Minor         string
	EndpointCount int
	SchemaCount   int
}

func main() {
	var (
		specPath     = flag.String("spec", "", "Single OpenAPI spec file")
		manifestPath = flag.String("manifest", "", "manifest.json for batch processing")
		output       = flag.String("output", "", "Output directory (for single spec)")
		docsOutput   = flag.String("docs-output", "docs", "Markdown output directory")
		redocOutput  = flag.String("redoc-output", "public", "ReDoc HTML output directory")
		format       = flag.String("format", "all", "Output format: markdown, redoc, all")
		title        = flag.String("title", "CheckMK REST API", "Documentation title")
		baseURL      = flag.String("base-url", "/", "Base URL for links")
		specsDir     = flag.String("specs-dir", "specs", "Specs directory")
	)
	flag.Parse()

	if *specPath == "" && *manifestPath == "" {
		log.Fatal("Error: either -spec or -manifest is required")
	}

	gen := &Generator{
		specsDir:    *specsDir,
		docsOutput:  *docsOutput,
		redocOutput: *redocOutput,
		title:       *title,
		baseURL:     *baseURL,
	}

	if *specPath != "" {
		// Single spec mode
		if *output == "" {
			log.Fatal("Error: -output is required when using -spec")
		}
		if err := gen.GenerateSingle(*specPath, *output, *format); err != nil {
			log.Fatalf("Failed to generate: %v", err)
		}
	} else {
		// Batch mode from manifest
		data, err := os.ReadFile(*manifestPath)
		if err != nil {
			log.Fatalf("Failed to read manifest: %v", err)
		}

		gen.manifest = &ManifestInput{}
		if err := json.Unmarshal(data, gen.manifest); err != nil {
			log.Fatalf("Failed to parse manifest: %v", err)
		}

		if err := gen.GenerateAll(*format); err != nil {
			log.Fatalf("Failed to generate: %v", err)
		}
	}

	fmt.Println("Documentation generated successfully")
}

// LoadSpec loads an OpenAPI specification from a YAML file
func LoadSpec(path string) (*OpenAPISpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading spec file: %w", err)
	}

	spec := &OpenAPISpec{}
	if err := yaml.Unmarshal(data, spec); err != nil {
		return nil, fmt.Errorf("parsing YAML: %w", err)
	}

	return spec, nil
}

// GenerateSingle generates documentation for a single spec file
func (g *Generator) GenerateSingle(specPath, outputDir, format string) error {
	spec, err := LoadSpec(specPath)
	if err != nil {
		return err
	}

	// Extract version from path (e.g., "specs/2.4.0/p17.yaml" -> "2.4.0p17")
	version := extractVersionFromPath(specPath)

	switch format {
	case "markdown", "all":
		if err := g.generateMarkdownDocs(spec, version, outputDir); err != nil {
			return fmt.Errorf("generating markdown: %w", err)
		}
	}

	if format == "redoc" || format == "all" {
		redocDir := filepath.Join(g.redocOutput, strings.ReplaceAll(version, ".", "_"))
		if err := g.generateReDocPage(spec, version, redocDir, specPath); err != nil {
			return fmt.Errorf("generating ReDoc: %w", err)
		}
	}

	return nil
}

// GenerateAll generates documentation for all baselines in the manifest
func (g *Generator) GenerateAll(format string) error {
	fmt.Printf("Generating documentation for %d baselines...\n", len(g.manifest.Baselines))

	var versions []VersionInfo

	for _, baseline := range g.manifest.Baselines {
		info := g.manifest.Mapping[baseline]
		specPath := filepath.Join(g.specsDir, info.Spec)

		spec, err := LoadSpec(specPath)
		if err != nil {
			log.Printf("Warning: failed to load %s: %v", baseline, err)
			continue
		}

		// Extract minor version (e.g., "2.4.0p17" -> "2.4")
		minor := extractMinorVersion(baseline)

		vi := VersionInfo{
			Version:       baseline,
			Path:          info.Path,
			Minor:         minor,
			EndpointCount: countEndpoints(spec),
			SchemaCount:   countSchemas(spec),
		}
		versions = append(versions, vi)

		// Generate markdown
		if format == "markdown" || format == "all" {
			mdOutput := filepath.Join(g.docsOutput, info.Path)
			if err := g.generateMarkdownDocs(spec, baseline, mdOutput); err != nil {
				log.Printf("Warning: failed to generate markdown for %s: %v", baseline, err)
			}
		}

		// Generate ReDoc
		if format == "redoc" || format == "all" {
			redocDir := filepath.Join(g.redocOutput, strings.ReplaceAll(baseline, ".", "_"))
			if err := g.generateReDocPage(spec, baseline, redocDir, specPath); err != nil {
				log.Printf("Warning: failed to generate ReDoc for %s: %v", baseline, err)
			}
		}

		fmt.Printf("  Generated: %s (%d endpoints, %d schemas)\n", baseline, vi.EndpointCount, vi.SchemaCount)
	}

	// Generate index files
	if format == "markdown" || format == "all" {
		if err := g.generateMarkdownIndex(versions); err != nil {
			return fmt.Errorf("generating markdown index: %w", err)
		}
	}

	if format == "redoc" || format == "all" {
		if err := g.generateReDocLanding(versions); err != nil {
			return fmt.Errorf("generating ReDoc landing: %w", err)
		}
	}

	return nil
}

// generateMarkdownDocs generates markdown documentation for a single spec
func (g *Generator) generateMarkdownDocs(spec *OpenAPISpec, version, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Generate README.md (version index)
	if err := g.generateVersionReadme(spec, version, outputDir); err != nil {
		return err
	}

	// Generate endpoints.md
	if err := g.generateEndpointsMd(spec, version, outputDir); err != nil {
		return err
	}

	// Generate schemas.md
	if err := g.generateSchemasMd(spec, version, outputDir); err != nil {
		return err
	}

	return nil
}

// generateVersionReadme generates the README.md for a specific version
func (g *Generator) generateVersionReadme(spec *OpenAPISpec, version, outputDir string) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s v%s\n\n", g.title, version))
	sb.WriteString("API documentation generated from OpenAPI specification.\n\n")

	sb.WriteString("## Quick Links\n\n")
	sb.WriteString("- [All Endpoints](endpoints.md)\n")
	sb.WriteString("- [All Schemas](schemas.md)\n\n")

	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- **Version**: %s\n", version))
	sb.WriteString(fmt.Sprintf("- **Endpoints**: %d\n", countEndpoints(spec)))
	sb.WriteString(fmt.Sprintf("- **Schemas**: %d\n\n", countSchemas(spec)))

	// Group endpoints by tag
	tagCounts := make(map[string]int)
	for _, pathItem := range spec.Paths {
		for _, op := range pathItem {
			if op != nil && len(op.Tags) > 0 {
				tagCounts[op.Tags[0]]++
			}
		}
	}

	if len(tagCounts) > 0 {
		sb.WriteString("## Endpoint Categories\n\n")
		sb.WriteString("| Category | Count |\n")
		sb.WriteString("|----------|-------|\n")

		tags := make([]string, 0, len(tagCounts))
		for tag := range tagCounts {
			tags = append(tags, tag)
		}
		sort.Strings(tags)

		for _, tag := range tags {
			anchor := strings.ToLower(strings.ReplaceAll(tag, " ", "-"))
			sb.WriteString(fmt.Sprintf("| [%s](endpoints.md#%s) | %d |\n", tag, anchor, tagCounts[tag]))
		}
	}

	return os.WriteFile(filepath.Join(outputDir, "README.md"), []byte(sb.String()), 0644)
}

// generateEndpointsMd generates the endpoints.md file
func (g *Generator) generateEndpointsMd(spec *OpenAPISpec, version, outputDir string) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# API Endpoints - v%s\n\n", version))

	// Group by tags
	tagEndpoints := make(map[string][]struct {
		Method string
		Path   string
		Op     *Operation
	})

	for path, pathItem := range spec.Paths {
		for method, op := range pathItem {
			if op == nil {
				continue
			}
			tag := "Other"
			if len(op.Tags) > 0 {
				tag = op.Tags[0]
			}
			tagEndpoints[tag] = append(tagEndpoints[tag], struct {
				Method string
				Path   string
				Op     *Operation
			}{strings.ToUpper(method), path, op})
		}
	}

	// Sort tags
	tags := make([]string, 0, len(tagEndpoints))
	for tag := range tagEndpoints {
		tags = append(tags, tag)
	}
	sort.Strings(tags)

	// Table of contents
	sb.WriteString("## Table of Contents\n\n")
	for _, tag := range tags {
		anchor := strings.ToLower(strings.ReplaceAll(tag, " ", "-"))
		sb.WriteString(fmt.Sprintf("- [%s](#%s)\n", tag, anchor))
	}
	sb.WriteString("\n---\n\n")

	// Generate each section
	for _, tag := range tags {
		endpoints := tagEndpoints[tag]

		// Sort endpoints by path
		sort.Slice(endpoints, func(i, j int) bool {
			return endpoints[i].Path < endpoints[j].Path
		})

		sb.WriteString(fmt.Sprintf("## %s\n\n", tag))

		for _, ep := range endpoints {
			sb.WriteString(fmt.Sprintf("### %s %s\n\n", ep.Method, ep.Path))

			if ep.Op.Summary != "" {
				sb.WriteString(fmt.Sprintf("%s\n\n", ep.Op.Summary))
			}

			// Parameters
			if len(ep.Op.Parameters) > 0 {
				sb.WriteString("**Parameters:**\n\n")
				sb.WriteString("| Name | In | Type | Required | Description |\n")
				sb.WriteString("|------|-----|------|----------|-------------|\n")

				for _, param := range ep.Op.Parameters {
					paramType := "string"
					if param.Schema != nil && param.Schema.Type != "" {
						paramType = param.Schema.Type
					}
					required := "No"
					if param.Required {
						required = "Yes"
					}
					desc := cleanDescription(param.Description)
					if len(desc) > 80 {
						desc = desc[:77] + "..."
					}
					sb.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s | %s |\n",
						param.Name, param.In, paramType, required, desc))
				}
				sb.WriteString("\n")
			}

			// Request body
			if ep.Op.RequestBody != nil && ep.Op.RequestBody.Content != nil {
				for _, media := range ep.Op.RequestBody.Content {
					if media.Schema != nil && media.Schema.Ref != "" {
						schemaName := extractSchemaName(media.Schema.Ref)
						sb.WriteString(fmt.Sprintf("**Request Body:** [%s](schemas.md#%s)\n\n",
							schemaName, strings.ToLower(schemaName)))
					}
				}
			}

			// Responses
			if len(ep.Op.Responses) > 0 {
				sb.WriteString("**Responses:**\n\n")
				sb.WriteString("| Code | Description | Schema |\n")
				sb.WriteString("|------|-------------|--------|\n")

				codes := make([]string, 0, len(ep.Op.Responses))
				for code := range ep.Op.Responses {
					codes = append(codes, code)
				}
				sort.Strings(codes)

				for _, code := range codes {
					resp := ep.Op.Responses[code]
					schemaLink := "-"
					if resp.Content != nil {
						for _, media := range resp.Content {
							if media.Schema != nil && media.Schema.Ref != "" {
								schemaName := extractSchemaName(media.Schema.Ref)
								schemaLink = fmt.Sprintf("[%s](schemas.md#%s)", schemaName, strings.ToLower(schemaName))
							}
						}
					}
					desc := cleanDescription(resp.Description)
					sb.WriteString(fmt.Sprintf("| %s | %s | %s |\n", code, desc, schemaLink))
				}
				sb.WriteString("\n")
			}

			sb.WriteString("---\n\n")
		}
	}

	return os.WriteFile(filepath.Join(outputDir, "endpoints.md"), []byte(sb.String()), 0644)
}

// generateSchemasMd generates the schemas.md file
func (g *Generator) generateSchemasMd(spec *OpenAPISpec, version, outputDir string) error {
	if spec.Components == nil || len(spec.Components.Schemas) == 0 {
		return nil
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Schema Definitions - v%s\n\n", version))

	// Sort schema names
	schemaNames := make([]string, 0, len(spec.Components.Schemas))
	for name := range spec.Components.Schemas {
		schemaNames = append(schemaNames, name)
	}
	sort.Strings(schemaNames)

	// Table of contents
	sb.WriteString("## Table of Contents\n\n")
	for _, name := range schemaNames {
		anchor := strings.ToLower(name)
		sb.WriteString(fmt.Sprintf("- [%s](#%s)\n", name, anchor))
	}
	sb.WriteString("\n---\n\n")

	// Generate each schema
	for _, name := range schemaNames {
		schema := spec.Components.Schemas[name]

		sb.WriteString(fmt.Sprintf("## %s\n\n", name))

		if schema.Description != "" {
			sb.WriteString(fmt.Sprintf("%s\n\n", cleanDescription(schema.Description)))
		}

		// Properties table
		if len(schema.Properties) > 0 {
			sb.WriteString("### Properties\n\n")
			sb.WriteString("| Property | Type | Required | Description |\n")
			sb.WriteString("|----------|------|----------|-------------|\n")

			// Sort properties
			propNames := make([]string, 0, len(schema.Properties))
			for propName := range schema.Properties {
				propNames = append(propNames, propName)
			}
			sort.Strings(propNames)

			requiredSet := make(map[string]bool)
			for _, r := range schema.Required {
				requiredSet[r] = true
			}

			for _, propName := range propNames {
				prop := schema.Properties[propName]
				propType := getSchemaType(prop, spec)
				required := "No"
				if requiredSet[propName] {
					required = "Yes"
				}
				desc := cleanDescription(prop.Description)
				if len(desc) > 80 {
					desc = desc[:77] + "..."
				}
				sb.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s |\n",
					propName, propType, required, desc))
			}
			sb.WriteString("\n")
		}

		// Enum values
		if len(schema.Enum) > 0 {
			sb.WriteString("### Enum Values\n\n")
			for _, v := range schema.Enum {
				sb.WriteString(fmt.Sprintf("- `%v`\n", v))
			}
			sb.WriteString("\n")
		}

		sb.WriteString("---\n\n")
	}

	return os.WriteFile(filepath.Join(outputDir, "schemas.md"), []byte(sb.String()), 0644)
}

// generateMarkdownIndex generates the main docs/README.md index
func (g *Generator) generateMarkdownIndex(versions []VersionInfo) error {
	if err := os.MkdirAll(g.docsOutput, 0755); err != nil {
		return err
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s Documentation\n\n", g.title))
	sb.WriteString("Generated API documentation for CheckMK REST API across all baseline versions.\n\n")

	// Group by minor version
	minorVersions := make(map[string][]VersionInfo)
	for _, v := range versions {
		minorVersions[v.Minor] = append(minorVersions[v.Minor], v)
	}

	// Sort minor versions (descending)
	minors := make([]string, 0, len(minorVersions))
	for m := range minorVersions {
		minors = append(minors, m)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(minors)))

	sb.WriteString("## Available Versions\n\n")

	for _, minor := range minors {
		vList := minorVersions[minor]
		// Sort versions within minor (descending by patch)
		sort.Slice(vList, func(i, j int) bool {
			return vList[i].Version > vList[j].Version
		})

		sb.WriteString(fmt.Sprintf("### CheckMK %s.x\n\n", minor))
		sb.WriteString("| Version | Endpoints | Schemas |\n")
		sb.WriteString("|---------|-----------|----------|\n")

		for _, v := range vList {
			sb.WriteString(fmt.Sprintf("| [%s](%s/) | %d | %d |\n",
				v.Version, v.Path, v.EndpointCount, v.SchemaCount))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## See Also\n\n")
	sb.WriteString(fmt.Sprintf("- [Interactive Documentation](%s)\n", g.baseURL))

	return os.WriteFile(filepath.Join(g.docsOutput, "README.md"), []byte(sb.String()), 0644)
}

// generateReDocPage generates a ReDoc HTML page for a version
func (g *Generator) generateReDocPage(spec *OpenAPISpec, version, outputDir, specPath string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Reference spec from /specs/ directory (no copying)
	// specPath is like "specs/2.4.0/p17.yaml" -> "/specs/2.4.0/p17.yaml"
	specURL := "/" + specPath

	// Generate HTML page
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s v%s</title>
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
    <style>
        body { margin: 0; padding: 0; }
        .version-banner {
            background: #1a1a1a;
            color: white;
            padding: 0.5rem 1rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-family: 'Roboto', sans-serif;
        }
        .version-banner a { color: #6cb4ee; text-decoration: none; }
        .version-banner a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <div class="version-banner">
        <span><strong>%s</strong> v%s</span>
        <div>
            <a href="../">All Versions</a>
        </div>
    </div>

    <redoc spec-url="%s"
           expand-responses="200,201"
           hide-hostname
           theme='{
               "colors": { "primary": { "main": "#0366d6" } },
               "typography": { "fontSize": "15px" }
           }'>
    </redoc>

    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>
</body>
</html>`, g.title, version, g.title, version, specURL)

	return os.WriteFile(filepath.Join(outputDir, "index.html"), []byte(html), 0644)
}

// generateReDocLanding generates the landing page with version selector
func (g *Generator) generateReDocLanding(versions []VersionInfo) error {
	if err := os.MkdirAll(g.redocOutput, 0755); err != nil {
		return err
	}

	// Generate versions.json
	versionsJSON, err := json.MarshalIndent(versions, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(g.redocOutput, "versions.json"), versionsJSON, 0644); err != nil {
		return err
	}

	// Generate index.html
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s Documentation</title>
    <style>
        * { box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            margin: 0;
            padding: 0;
            background: #f5f5f5;
        }
        .container { max-width: 1200px; margin: 0 auto; padding: 2rem; }
        h1 { color: #1a1a1a; margin-bottom: 0.5rem; }
        .subtitle { color: #666; margin-bottom: 2rem; }
        .version-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
            gap: 1rem;
        }
        .version-card {
            background: white;
            border: 1px solid #e1e4e8;
            border-radius: 8px;
            padding: 1.25rem;
            transition: all 0.2s;
        }
        .version-card:hover {
            border-color: #0366d6;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
            transform: translateY(-2px);
        }
        .version-card h3 { margin: 0 0 0.5rem; }
        .version-card a {
            color: #0366d6;
            text-decoration: none;
            font-size: 1.1rem;
            font-weight: 600;
        }
        .version-card a:hover { text-decoration: underline; }
        .version-card .stats { color: #666; font-size: 0.9rem; }
        .minor-section { margin-bottom: 2.5rem; }
        .minor-section h2 {
            color: #333;
            border-bottom: 2px solid #0366d6;
            padding-bottom: 0.5rem;
            margin-bottom: 1rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>%s</h1>
        <p class="subtitle">Select a version to view the interactive API documentation.</p>

        <div id="versions"></div>
    </div>

    <script>
        fetch('versions.json')
            .then(r => r.json())
            .then(versions => {
                // Group by minor version
                const grouped = {};
                versions.forEach(v => {
                    if (!grouped[v.Minor]) grouped[v.Minor] = [];
                    grouped[v.Minor].push(v);
                });

                const container = document.getElementById('versions');

                // Sort minor versions descending
                Object.keys(grouped).sort().reverse().forEach(minor => {
                    const section = document.createElement('div');
                    section.className = 'minor-section';
                    section.innerHTML = '<h2>CheckMK ' + minor + '.x</h2>';

                    const grid = document.createElement('div');
                    grid.className = 'version-grid';

                    // Sort versions descending within minor
                    grouped[minor].sort((a, b) => b.Version.localeCompare(a.Version)).forEach(v => {
                        const path = v.Version.replace(/\./g, '_');
                        grid.innerHTML +=
                            '<div class="version-card">' +
                                '<h3><a href="' + path + '/">' + v.Version + '</a></h3>' +
                                '<p class="stats">' + v.EndpointCount + ' endpoints, ' + v.SchemaCount + ' schemas</p>' +
                            '</div>';
                    });

                    section.appendChild(grid);
                    container.appendChild(section);
                });
            });
    </script>
</body>
</html>`, g.title, g.title)

	return os.WriteFile(filepath.Join(g.redocOutput, "index.html"), []byte(html), 0644)
}

// Helper functions

func extractVersionFromPath(path string) string {
	// "specs/2.4.0/p17.yaml" -> "2.4.0p17"
	parts := strings.Split(path, "/")
	if len(parts) >= 3 {
		minor := parts[len(parts)-2]
		patch := strings.TrimSuffix(parts[len(parts)-1], ".yaml")
		return minor + patch
	}
	return "unknown"
}

func extractMinorVersion(version string) string {
	// "2.4.0p17" -> "2.4"
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return version
}

func extractSchemaName(ref string) string {
	// "#/components/schemas/HostConfig" -> "HostConfig"
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

func countEndpoints(spec *OpenAPISpec) int {
	count := 0
	for _, pathItem := range spec.Paths {
		for _, op := range pathItem {
			if op != nil {
				count++
			}
		}
	}
	return count
}

func countSchemas(spec *OpenAPISpec) int {
	if spec.Components == nil {
		return 0
	}
	return len(spec.Components.Schemas)
}

func cleanDescription(desc string) string {
	// Remove newlines and extra whitespace
	desc = strings.ReplaceAll(desc, "\n", " ")
	desc = strings.ReplaceAll(desc, "\r", " ")
	for strings.Contains(desc, "  ") {
		desc = strings.ReplaceAll(desc, "  ", " ")
	}
	return strings.TrimSpace(desc)
}

func getSchemaType(schema *Schema, spec *OpenAPISpec) string {
	if schema == nil {
		return "any"
	}

	if schema.Ref != "" {
		return extractSchemaName(schema.Ref)
	}

	if schema.Type == "array" && schema.Items != nil {
		itemType := getSchemaType(schema.Items, spec)
		return "[]" + itemType
	}

	if schema.Type != "" {
		return schema.Type
	}

	if len(schema.OneOf) > 0 || len(schema.AnyOf) > 0 {
		return "oneOf"
	}

	return "object"
}
