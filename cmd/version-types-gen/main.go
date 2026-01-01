// Package main implements the version-types-gen tool for generating Go code
// that maps CheckMK versions to their baseline type packages.
//
// This tool reads manifest.json and generates version_types.go which the
// Terraform provider imports to resolve versions to the correct type package.
//
// Usage:
//
//	version-types-gen -baselines manifest.json -output version_types.go -package client
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

// BaselinesInput matches the manifest.json structure from spec-sync
type BaselinesInput struct {
	Baselines    []string                `json:"baselines"`
	Mapping      map[string]BaselineInfo `json:"mapping"`
	VersionOrder []string                `json:"version_order"`
}

// BaselineInfo describes how a version maps to a baseline
type BaselineInfo struct {
	Baseline   string `json:"baseline"`
	Package    string `json:"package"`
	IsBaseline bool   `json:"is_baseline"`
}

// TemplateData is passed to the code template
type TemplateData struct {
	Package         string
	ModulePath      string
	Baselines       []string
	BaselineImports []BaselineImport
	VersionMapping  []VersionMapEntry
	MinorBaselines  []MinorBaselineEntry
}

// BaselineImport represents an import statement
type BaselineImport struct {
	Alias   string
	Package string
}

// VersionMapEntry represents a version to baseline mapping
type VersionMapEntry struct {
	Version  string
	Baseline string
	Package  string
}

// MinorBaselineEntry maps minor versions to their latest baseline
type MinorBaselineEntry struct {
	Minor    string
	Baseline string
	Package  string
}

func main() {
	var (
		baselinesPath = flag.String("baselines", "manifest.json", "Path to manifest.json")
		outputPath    = flag.String("output", "version_types.go", "Output Go file path")
		packageName   = flag.String("package", "client", "Go package name")
		modulePath    = flag.String("module", "github.com/BlackMesaLTD/checkmk-api-spec/generated/go", "Module path for imports")
	)
	flag.Parse()

	// Read baselines.json
	data, err := os.ReadFile(*baselinesPath)
	if err != nil {
		log.Fatalf("Failed to read baselines file: %v", err)
	}

	var baselines BaselinesInput
	if err := json.Unmarshal(data, &baselines); err != nil {
		log.Fatalf("Failed to parse baselines file: %v", err)
	}

	// Generate code
	code, err := generateCode(&baselines, *packageName, *modulePath)
	if err != nil {
		log.Fatalf("Failed to generate code: %v", err)
	}

	// Format code
	formatted, err := format.Source(code)
	if err != nil {
		log.Printf("Warning: failed to format code: %v", err)
		formatted = code
	}

	// Write output
	if err := os.WriteFile(*outputPath, formatted, 0644); err != nil {
		log.Fatalf("Failed to write output: %v", err)
	}

	fmt.Printf("Generated %s\n", *outputPath)
	fmt.Printf("  Package: %s\n", *packageName)
	fmt.Printf("  Baselines: %d\n", len(baselines.Baselines))
	fmt.Printf("  Total versions mapped: %d\n", len(baselines.Mapping))
}

func generateCode(baselines *BaselinesInput, packageName, modulePath string) ([]byte, error) {
	tmplData := &TemplateData{
		Package:    packageName,
		ModulePath: modulePath,
		Baselines:  baselines.Baselines,
	}

	// Build imports
	seen := make(map[string]bool)
	for _, baseline := range baselines.Baselines {
		info := baselines.Mapping[baseline]
		if !seen[info.Package] {
			tmplData.BaselineImports = append(tmplData.BaselineImports, BaselineImport{
				Alias:   info.Package,
				Package: modulePath + "/" + info.Package,
			})
			seen[info.Package] = true
		}
	}

	// Build version mapping (sorted)
	versions := make([]string, 0, len(baselines.Mapping))
	for v := range baselines.Mapping {
		versions = append(versions, v)
	}
	sort.Slice(versions, func(i, j int) bool {
		return compareVersions(versions[i], versions[j]) < 0
	})

	for _, v := range versions {
		info := baselines.Mapping[v]
		tmplData.VersionMapping = append(tmplData.VersionMapping, VersionMapEntry{
			Version:  v,
			Baseline: info.Baseline,
			Package:  info.Package,
		})
	}

	// Build minor version to latest baseline mapping
	minorLatest := make(map[string]string)
	for _, v := range versions {
		minor := getMinorVersion(v)
		info := baselines.Mapping[v]
		if info.IsBaseline {
			minorLatest[minor] = v
		}
	}

	var minors []string
	for m := range minorLatest {
		minors = append(minors, m)
	}
	sort.Strings(minors)

	for _, minor := range minors {
		baseline := minorLatest[minor]
		info := baselines.Mapping[baseline]
		tmplData.MinorBaselines = append(tmplData.MinorBaselines, MinorBaselineEntry{
			Minor:    minor,
			Baseline: baseline,
			Package:  info.Package,
		})
	}

	// Execute template
	var buf bytes.Buffer
	if err := codeTemplate.Execute(&buf, tmplData); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getMinorVersion(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return version
}

func compareVersions(a, b string) int {
	partsA := parseVersion(a)
	partsB := parseVersion(b)

	for i := 0; i < len(partsA) && i < len(partsB); i++ {
		if partsA[i] < partsB[i] {
			return -1
		}
		if partsA[i] > partsB[i] {
			return 1
		}
	}

	if len(partsA) < len(partsB) {
		return -1
	}
	if len(partsA) > len(partsB) {
		return 1
	}
	return 0
}

func parseVersion(v string) []int {
	v = strings.ReplaceAll(v, "p", ".")
	parts := strings.Split(v, ".")
	result := make([]int, len(parts))
	for i, p := range parts {
		fmt.Sscanf(p, "%d", &result[i])
	}
	return result
}

// toTitle converts a package name to a title case identifier
// e.g., "v2_2_0p43" -> "V2_2_0p43"
func toTitle(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

var codeTemplate = template.Must(template.New("code").Funcs(template.FuncMap{
	"title": toTitle,
}).Parse(`// Code generated by version-types-gen. DO NOT EDIT.
//
// This file maps CheckMK versions to their baseline type packages.
// A baseline is a version where the API changed significantly enough
// to require separate type definitions.

package {{.Package}}

import "strings"
{{if .BaselineImports}}
import (
{{range .BaselineImports}}	{{.Alias}} "{{.Package}}"
{{end}})
{{end}}

// BaselinePackage represents a baseline type package identifier.
type BaselinePackage string

// Known baseline packages.
const (
{{range .BaselineImports}}	Baseline{{.Alias | title}} BaselinePackage = "{{.Alias}}"
{{end}})

// VersionToBaseline maps CheckMK versions to their baseline package.
// Use LookupBaseline() instead of accessing this directly.
var VersionToBaseline = map[string]BaselinePackage{
{{range .VersionMapping}}	"{{.Version}}": Baseline{{.Package | title}},
{{end}}}

// MinorToLatestBaseline maps minor versions to their latest baseline.
// Used for unknown patch versions within a known minor.
var MinorToLatestBaseline = map[string]BaselinePackage{
{{range .MinorBaselines}}	"{{.Minor}}": Baseline{{.Package | title}},
{{end}}}

// LookupBaseline returns the baseline package for a given CheckMK version.
// Returns empty string if the version is unknown.
func LookupBaseline(version string) BaselinePackage {
	// Direct match
	if pkg, ok := VersionToBaseline[version]; ok {
		return pkg
	}

	// Try minor version match (for future patch versions)
	minor := extractMinor(version)
	if pkg, ok := MinorToLatestBaseline[minor]; ok {
		return pkg
	}

	return ""
}

// extractMinor extracts "2.4" from "2.4.0p17"
func extractMinor(version string) string {
	parts := splitVersion(version)
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return ""
}

// splitVersion splits a version string into parts
func splitVersion(version string) []string {
	// Handle both "2.4.0p17" and "2.4.0.17" formats
	version = strings.ReplaceAll(version, "p", ".")
	return strings.Split(version, ".")
}

// ValidHostTagAgentValues returns valid tag_agent values for the given baseline.
func ValidHostTagAgentValues(pkg BaselinePackage) []string {
	switch pkg {
{{range .BaselineImports}}	case Baseline{{.Alias | title}}:
		return {{.Alias}}.ValidHostCreateAttributeTagAgentValues()
{{end}}	}
	return nil
}

// HostCreateAttributeFieldNames returns valid host attribute field names for the given baseline.
func HostCreateAttributeFieldNames(pkg BaselinePackage) []string {
	switch pkg {
{{range .BaselineImports}}	case Baseline{{.Alias | title}}:
		return {{.Alias}}.HostCreateAttributeFieldNames
{{end}}	}
	return nil
}

// FolderCreateAttributeFieldNames returns valid folder attribute field names for the given baseline.
func FolderCreateAttributeFieldNames(pkg BaselinePackage) []string {
	switch pkg {
{{range .BaselineImports}}	case Baseline{{.Alias | title}}:
		return {{.Alias}}.FolderCreateAttributeFieldNames
{{end}}	}
	return nil
}
`))
