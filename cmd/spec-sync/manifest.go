package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// Manifest tracks all known versions and their baseline mappings
type Manifest struct {
	Baselines   []string                `json:"baselines"`    // List of baseline versions
	Mapping     map[string]VersionEntry `json:"mapping"`      // Version -> entry mapping
	LastChecked time.Time               `json:"last_checked"`

	// Keep internal map for backwards compat during transition
	Versions map[string]VersionEntry `json:"-"`
}

// VersionEntry maps a version to its baseline spec
type VersionEntry struct {
	Spec        string `json:"spec"`         // Relative path: "2.4.0/p1.yaml"
	Baseline    string `json:"baseline"`     // Baseline version: "2.4.0p1"
	Package     string `json:"package"`      // Go package name: "v2_4_0p1"
	IsBaseline  bool   `json:"is_baseline"`  // True if this version IS a baseline
	MaxSeverity string `json:"max_severity"` // Severity that triggered baseline: "initial", "breaking", "minor"
}

// NewManifest creates a new empty manifest
func NewManifest() *Manifest {
	return &Manifest{
		Baselines:   []string{},
		Mapping:     make(map[string]VersionEntry),
		Versions:    make(map[string]VersionEntry), // internal use
		LastChecked: time.Now(),
	}
}

// LoadManifest loads a manifest from file
func LoadManifest(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	// Initialize maps
	if manifest.Mapping == nil {
		manifest.Mapping = make(map[string]VersionEntry)
	}
	if manifest.Versions == nil {
		manifest.Versions = make(map[string]VersionEntry)
	}

	// Populate internal Versions map from Mapping for backwards compat
	for k, v := range manifest.Mapping {
		manifest.Versions[k] = v
	}

	return &manifest, nil
}

// Save writes the manifest to file with versions sorted numerically
func (m *Manifest) Save(path string) error {
	m.LastChecked = time.Now()

	// Sync Mapping from Versions (internal map)
	m.Mapping = make(map[string]VersionEntry)
	for k, v := range m.Versions {
		m.Mapping[k] = v
	}

	// Get sorted version keys
	versions := m.GetVersions()

	// Calculate baselines list
	m.Baselines = m.GetBaselines()

	// Build ordered output manually for proper version sorting
	var buf strings.Builder
	buf.WriteString("{\n")

	// Write baselines array
	buf.WriteString("  \"baselines\": [\n")
	for i, b := range m.Baselines {
		buf.WriteString(fmt.Sprintf("    %q", b))
		if i < len(m.Baselines)-1 {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	buf.WriteString("  ],\n")

	// Write mapping
	buf.WriteString("  \"mapping\": {\n")
	for i, v := range versions {
		entry := m.Versions[v]
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal entry for %s: %w", v, err)
		}

		buf.WriteString(fmt.Sprintf("    %q: %s", v, string(entryJSON)))
		if i < len(versions)-1 {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	buf.WriteString("  },\n")

	buf.WriteString(fmt.Sprintf("  \"last_checked\": %q\n", m.LastChecked.Format(time.RFC3339)))
	buf.WriteString("}\n")

	if err := os.WriteFile(path, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("failed to write manifest: %w", err)
	}

	return nil
}

// GetBaselines returns all baseline versions sorted
func (m *Manifest) GetBaselines() []string {
	var baselines []string
	for version, entry := range m.Versions {
		if entry.IsBaseline {
			baselines = append(baselines, version)
		}
	}
	sortVersions(baselines)
	return baselines
}

// GetVersions returns all versions sorted
func (m *Manifest) GetVersions() []string {
	var versions []string
	for version := range m.Versions {
		versions = append(versions, version)
	}
	sortVersions(versions)
	return versions
}

// PrintSummary prints a summary of the manifest
func (m *Manifest) PrintSummary() {
	baselines := m.GetBaselines()
	versions := m.GetVersions()

	log.Printf("\nManifest Summary:")
	log.Printf("  Total versions: %d", len(versions))
	log.Printf("  Baselines: %d", len(baselines))

	// Group by minor
	byMinor := make(map[string]struct {
		total     int
		baselines int
	})

	for _, v := range versions {
		minor := extractMinor(v)
		entry := byMinor[minor]
		entry.total++
		if m.Versions[v].IsBaseline {
			entry.baselines++
		}
		byMinor[minor] = entry
	}

	// Print by minor version
	var minors []string
	for minor := range byMinor {
		minors = append(minors, minor)
	}
	sort.Strings(minors)

	log.Printf("\n  By minor version:")
	for _, minor := range minors {
		entry := byMinor[minor]
		log.Printf("    %s: %d versions, %d baselines", minor, entry.total, entry.baselines)
	}
}

// Print prints the full manifest (for dry-run)
func (m *Manifest) Print() {
	baselines := m.GetBaselines()

	log.Printf("\nBaselines (%d):", len(baselines))
	for _, b := range baselines {
		log.Printf("  - %s", b)
	}

	log.Printf("\nVersion mappings (%d):", len(m.Versions))

	// Group by minor for cleaner output
	byMinor := make(map[string][]string)
	for v := range m.Versions {
		minor := extractMinor(v)
		byMinor[minor] = append(byMinor[minor], v)
	}

	var minors []string
	for minor := range byMinor {
		minors = append(minors, minor)
	}
	sort.Strings(minors)

	for _, minor := range minors {
		versions := byMinor[minor]
		sortVersions(versions)
		log.Printf("\n  %s series:", minor)
		for _, v := range versions {
			entry := m.Versions[v]
			if entry.IsBaseline {
				log.Printf("    %s: BASELINE", v)
			} else {
				log.Printf("    %s: -> %s", v, entry.Baseline)
			}
		}
	}
}

// extractMinor extracts the minor version from a full version
// 2.4.0p17 -> 2.4.0
func extractMinor(version string) string {
	for i := len(version) - 1; i >= 0; i-- {
		if version[i] == 'p' {
			return version[:i]
		}
	}
	return version
}
