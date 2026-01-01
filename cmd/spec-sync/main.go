// Package main implements the spec-sync tool for managing OpenAPI spec baselines.
//
// This tool:
//   - Queries Docker Hub for available CheckMK versions
//   - Compares specs to detect API changes
//   - Maintains a manifest mapping versions to their baseline specs
//   - Only stores specs where the API actually changed
//
// Usage:
//
//	spec-sync --bootstrap              # Build manifest from existing specs
//	spec-sync                          # Sync new versions from Docker Hub
//	spec-sync --minor 2.4              # Only process 2.4.x versions
//	spec-sync --dry-run                # Show what would be done
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Config holds the tool configuration
type Config struct {
	SpecsDir     string
	ManifestPath string
	Bootstrap    bool
	Cleanup      bool
	DryRun       bool
	Verbose      bool
	Minor        string
	Force        bool
}

var versionRegex = regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)p(\d+)$`)

func main() {
	cfg := parseFlags()

	if cfg.Verbose {
		log.Printf("Config: specs=%s manifest=%s bootstrap=%v cleanup=%v dry-run=%v minor=%s",
			cfg.SpecsDir, cfg.ManifestPath, cfg.Bootstrap, cfg.Cleanup, cfg.DryRun, cfg.Minor)
	}

	if cfg.Bootstrap {
		if err := runBootstrap(cfg); err != nil {
			log.Fatalf("Bootstrap failed: %v", err)
		}
	} else if cfg.Cleanup {
		if err := runCleanup(cfg); err != nil {
			log.Fatalf("Cleanup failed: %v", err)
		}
	} else {
		if err := runSync(cfg); err != nil {
			log.Fatalf("Sync failed: %v", err)
		}
	}
}

func parseFlags() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.SpecsDir, "specs", "specs", "Directory containing spec files")
	flag.StringVar(&cfg.ManifestPath, "manifest", "manifest.json", "Path to manifest file")
	flag.BoolVar(&cfg.Bootstrap, "bootstrap", false, "Build manifest from existing specs (no Docker)")
	flag.BoolVar(&cfg.Cleanup, "cleanup", false, "Remove non-baseline spec files")
	flag.BoolVar(&cfg.DryRun, "dry-run", false, "Show what would be done without making changes")
	flag.BoolVar(&cfg.Verbose, "v", false, "Verbose output")
	flag.StringVar(&cfg.Minor, "minor", "", "Filter by minor version (e.g., 2.4)")
	flag.BoolVar(&cfg.Force, "force", false, "Re-check versions even if already in manifest")

	flag.Parse()

	return cfg
}

// runBootstrap builds the manifest from existing spec files
func runBootstrap(cfg *Config) error {
	log.Println("=== Bootstrap Mode ===")
	log.Printf("Scanning specs directory: %s", cfg.SpecsDir)

	// Find all existing spec files
	versions, err := findExistingSpecs(cfg.SpecsDir)
	if err != nil {
		return fmt.Errorf("failed to find specs: %w", err)
	}

	// Filter by minor version if specified
	if cfg.Minor != "" {
		versions = filterByMinor(versions, cfg.Minor)
	}

	log.Printf("Found %d spec files", len(versions))

	if len(versions) == 0 {
		log.Println("No specs found, nothing to do")
		return nil
	}

	// Sort versions
	sortVersions(versions)

	// Load or create manifest
	manifest, err := LoadManifest(cfg.ManifestPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load manifest: %w", err)
	}
	if manifest == nil {
		manifest = NewManifest()
	}

	// Group by minor version
	byMinor := groupByMinor(versions)

	// Process each minor version series
	for minor, minorVersions := range byMinor {
		log.Printf("\n--- Processing %s series (%d versions) ---", minor, len(minorVersions))

		if err := processMinorSeries(cfg, manifest, minor, minorVersions); err != nil {
			log.Printf("Warning: failed to process %s: %v", minor, err)
			continue
		}
	}

	// Save manifest
	if cfg.DryRun {
		log.Println("\n[Dry run] Would save manifest:")
		manifest.Print()
		return nil
	}

	if err := manifest.Save(cfg.ManifestPath); err != nil {
		return fmt.Errorf("failed to save manifest: %w", err)
	}

	log.Printf("\nManifest saved to: %s", cfg.ManifestPath)
	manifest.PrintSummary()

	return nil
}

// runCleanup removes non-baseline spec files
func runCleanup(cfg *Config) error {
	log.Println("=== Cleanup Mode ===")

	// Load manifest
	manifest, err := LoadManifest(cfg.ManifestPath)
	if err != nil {
		return fmt.Errorf("failed to load manifest: %w", err)
	}

	// Build set of baseline spec paths
	baselineSpecs := make(map[string]bool)
	for _, entry := range manifest.Versions {
		if entry.IsBaseline {
			baselineSpecs[entry.Spec] = true
		}
	}

	log.Printf("Found %d baselines in manifest", len(baselineSpecs))

	// Find all spec files on disk
	allSpecs, err := findExistingSpecs(cfg.SpecsDir)
	if err != nil {
		return fmt.Errorf("failed to find specs: %w", err)
	}

	log.Printf("Found %d spec files on disk", len(allSpecs))

	// Find non-baseline specs
	var toDelete []string
	for _, version := range allSpecs {
		relPath := relativeSpecPath(version)
		if !baselineSpecs[relPath] {
			toDelete = append(toDelete, version)
		}
	}

	if len(toDelete) == 0 {
		log.Println("No non-baseline specs to delete")
		return nil
	}

	log.Printf("\nWill delete %d non-baseline spec files:", len(toDelete))
	for _, version := range toDelete {
		specPath := versionToSpecPath(cfg.SpecsDir, version)
		log.Printf("  - %s", specPath)
	}

	if cfg.DryRun {
		log.Println("\n[Dry run] No files deleted")
		return nil
	}

	// Delete non-baseline specs
	deleted := 0
	for _, version := range toDelete {
		specPath := versionToSpecPath(cfg.SpecsDir, version)
		if err := os.Remove(specPath); err != nil {
			log.Printf("  Warning: failed to delete %s: %v", specPath, err)
		} else {
			deleted++
		}
	}

	log.Printf("\nDeleted %d files", deleted)

	return nil
}

// processMinorSeries processes all versions in a minor series
func processMinorSeries(cfg *Config, manifest *Manifest, minor string, versions []string) error {
	var currentBaseline string
	var currentBaselineSpec []byte

	for i, version := range versions {
		specPath := versionToSpecPath(cfg.SpecsDir, version)

		// Check if already in manifest (unless force)
		if !cfg.Force {
			if entry, exists := manifest.Versions[version]; exists {
				if cfg.Verbose {
					log.Printf("  %s: already in manifest (baseline: %s)", version, entry.Baseline)
				}
				// Update current baseline if this is one
				if entry.IsBaseline {
					currentBaseline = version
					var err error
					currentBaselineSpec, err = os.ReadFile(specPath)
					if err != nil {
						return fmt.Errorf("failed to read baseline spec %s: %w", specPath, err)
					}
				}
				continue
			}
		}

		// Read spec
		specData, err := os.ReadFile(specPath)
		if err != nil {
			log.Printf("  %s: failed to read spec: %v", version, err)
			continue
		}

		// First version in series is always a baseline
		if i == 0 || currentBaseline == "" {
			log.Printf("  %s: BASELINE (first in series)", version)
			manifest.Versions[version] = VersionEntry{
				Spec:        relativeSpecPath(version),
				Baseline:    version,
				Package:     versionToPackage(version),
				IsBaseline:  true,
				MaxSeverity: "initial",
			}
			currentBaseline = version
			currentBaselineSpec = specData
			continue
		}

		// Compare with current baseline
		diff, err := CompareSpecs(currentBaselineSpec, specData)
		if err != nil {
			log.Printf("  %s: failed to compare: %v", version, err)
			continue
		}

		if SeverityOrder[diff.MaxSeverity] >= SeverityOrder[SeverityMinor] {
			// API changed - new baseline
			log.Printf("  %s: BASELINE (API changed: %s, %d changes)",
				version, diff.MaxSeverity, diff.TotalChanges)
			manifest.Versions[version] = VersionEntry{
				Spec:        relativeSpecPath(version),
				Baseline:    version,
				Package:     versionToPackage(version),
				IsBaseline:  true,
				MaxSeverity: string(diff.MaxSeverity),
			}
			currentBaseline = version
			currentBaselineSpec = specData
		} else {
			// No significant change - point to current baseline
			if cfg.Verbose {
				log.Printf("  %s: points to %s (no API changes)", version, currentBaseline)
			}
			// Get the baseline's package name
			baselineEntry := manifest.Versions[currentBaseline]
			manifest.Versions[version] = VersionEntry{
				Spec:        relativeSpecPath(currentBaseline),
				Baseline:    currentBaseline,
				Package:     baselineEntry.Package,
				IsBaseline:  false,
				MaxSeverity: string(diff.MaxSeverity),
			}
		}
	}

	return nil
}

// runSync syncs new versions from Docker Hub
func runSync(cfg *Config) error {
	log.Println("=== Sync Mode ===")

	// Load existing manifest
	manifest, err := LoadManifest(cfg.ManifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("No existing manifest, starting fresh")
			manifest = NewManifest()
		} else {
			return fmt.Errorf("failed to load manifest: %w", err)
		}
	}

	log.Printf("Loaded manifest with %d versions", len(manifest.Versions))

	// Find missing versions
	log.Println("Checking Docker Hub for new versions...")
	missing, err := FindMissingVersions(manifest, cfg.Minor)
	if err != nil {
		return fmt.Errorf("failed to check Docker Hub: %w", err)
	}

	if len(missing) == 0 {
		log.Println("No new versions found")
		return nil
	}

	log.Printf("Found %d new versions to fetch", len(missing))
	for _, v := range missing {
		log.Printf("  - %s", v)
	}

	if cfg.DryRun {
		log.Println("\n[Dry run] Would fetch and process these versions")
		return nil
	}

	// Process each missing version
	success := 0
	failed := 0

	for _, version := range missing {
		log.Printf("\nProcessing %s...", version)

		// Fetch spec from Docker (kept in memory until we decide if it's a baseline)
		specData, err := FetchSpecFromDocker(version, cfg.SpecsDir, cfg.Verbose)
		if err != nil {
			log.Printf("  Failed to fetch: %v", err)
			failed++
			continue
		}

		// Find baseline for comparison
		minor := extractMinor(version)
		latestBaseline := findLatestBaselineForMinor(manifest, minor)

		var isBaseline bool
		var maxSeverity string

		if latestBaseline == "" {
			// First version for this minor - always a baseline
			isBaseline = true
			maxSeverity = "initial"
			log.Printf("  BASELINE (first in %s series)", minor)
		} else {
			// Compare with latest baseline (in memory, no disk I/O yet)
			baselineSpecPath := versionToSpecPath(cfg.SpecsDir, latestBaseline)
			baselineData, err := os.ReadFile(baselineSpecPath)
			if err != nil {
				log.Printf("  Warning: couldn't read baseline spec: %v", err)
				// Treat as new baseline since we can't compare
				isBaseline = true
				maxSeverity = "unknown"
			} else {
				diff, err := CompareSpecs(baselineData, specData)
				if err != nil {
					log.Printf("  Warning: comparison failed: %v", err)
					failed++
					continue
				}

				maxSeverity = string(diff.MaxSeverity)

				if SeverityOrder[diff.MaxSeverity] >= SeverityOrder[SeverityMinor] {
					isBaseline = true
					log.Printf("  BASELINE (API changed: %s, %d changes)", diff.MaxSeverity, diff.TotalChanges)
				} else {
					isBaseline = false
					log.Printf("  Points to %s (no API changes, severity: %s)", latestBaseline, maxSeverity)
				}
			}
		}

		// Only save spec file if it's a baseline
		if isBaseline {
			if err := SaveSpec(version, cfg.SpecsDir, specData); err != nil {
				log.Printf("  Failed to save: %v", err)
				failed++
				continue
			}
			specPath := versionToSpecPath(cfg.SpecsDir, version)
			log.Printf("  Saved spec to %s", specPath)

			manifest.Versions[version] = VersionEntry{
				Spec:        relativeSpecPath(version),
				Baseline:    version,
				Package:     versionToPackage(version),
				IsBaseline:  true,
				MaxSeverity: maxSeverity,
			}
		} else {
			// Not a baseline - just update manifest to point to existing baseline
			baselineEntry := manifest.Versions[latestBaseline]
			manifest.Versions[version] = VersionEntry{
				Spec:        relativeSpecPath(latestBaseline),
				Baseline:    latestBaseline,
				Package:     baselineEntry.Package,
				IsBaseline:  false,
				MaxSeverity: maxSeverity,
			}
		}

		success++
	}

	// Save manifest
	if err := manifest.Save(cfg.ManifestPath); err != nil {
		return fmt.Errorf("failed to save manifest: %w", err)
	}

	log.Printf("\n=== Summary ===")
	log.Printf("Processed: %d, Failed: %d", success, failed)
	log.Printf("Manifest saved to: %s", cfg.ManifestPath)
	manifest.PrintSummary()

	return nil
}

// findLatestBaselineForMinor finds the latest baseline version for a minor series
func findLatestBaselineForMinor(manifest *Manifest, minor string) string {
	var latestBaseline string
	for version, entry := range manifest.Versions {
		if entry.IsBaseline && strings.HasPrefix(version, minor) {
			if latestBaseline == "" || compareVersions(version, latestBaseline) > 0 {
				latestBaseline = version
			}
		}
	}
	return latestBaseline
}

// findExistingSpecs finds all spec files in the specs directory
func findExistingSpecs(specsDir string) ([]string, error) {
	var versions []string

	err := filepath.Walk(specsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".yml") {
			return nil
		}

		// Extract version from path: specs/2.4.0/p17.yaml -> 2.4.0p17
		version := specPathToVersion(specsDir, path)
		if version != "" {
			versions = append(versions, version)
		}

		return nil
	})

	return versions, err
}

// specPathToVersion converts a spec path to version string
// specs/2.4.0/p17.yaml -> 2.4.0p17
func specPathToVersion(specsDir, path string) string {
	rel, err := filepath.Rel(specsDir, path)
	if err != nil {
		return ""
	}

	// Expected format: 2.4.0/p17.yaml
	parts := strings.Split(rel, string(filepath.Separator))
	if len(parts) != 2 {
		return ""
	}

	minor := parts[0] // 2.4.0
	file := parts[1]  // p17.yaml

	// Extract patch number
	file = strings.TrimSuffix(file, ".yaml")
	file = strings.TrimSuffix(file, ".yml")

	if !strings.HasPrefix(file, "p") {
		return ""
	}

	return minor + file // 2.4.0p17
}

// versionToSpecPath converts version to spec path
// 2.4.0p17 -> specs/2.4.0/p17.yaml
func versionToSpecPath(specsDir, version string) string {
	parts := strings.SplitN(version, "p", 2)
	if len(parts) != 2 {
		return ""
	}
	minor := parts[0]
	patch := parts[1]
	return filepath.Join(specsDir, minor, "p"+patch+".yaml")
}

// relativeSpecPath returns the relative spec path for manifest
// 2.4.0p17 -> 2.4.0/p17.yaml
func relativeSpecPath(version string) string {
	parts := strings.SplitN(version, "p", 2)
	if len(parts) != 2 {
		return version + ".yaml"
	}
	return parts[0] + "/p" + parts[1] + ".yaml"
}

// versionToPackage converts version to Go package name
// 2.4.0p17 -> v2_4_0p17
func versionToPackage(version string) string {
	pkg := "v" + strings.ReplaceAll(version, ".", "_")
	return pkg
}

// filterByMinor filters versions to only those matching the minor version
func filterByMinor(versions []string, minor string) []string {
	var filtered []string
	for _, v := range versions {
		if strings.HasPrefix(v, minor) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// groupByMinor groups versions by their minor version
func groupByMinor(versions []string) map[string][]string {
	groups := make(map[string][]string)
	for _, v := range versions {
		parts := strings.SplitN(v, "p", 2)
		if len(parts) == 2 {
			minor := parts[0]
			groups[minor] = append(groups[minor], v)
		}
	}

	// Sort each group
	for minor := range groups {
		sortVersions(groups[minor])
	}

	return groups
}

// sortVersions sorts version strings in semantic order
func sortVersions(versions []string) {
	sort.Slice(versions, func(i, j int) bool {
		return compareVersions(versions[i], versions[j]) < 0
	})
}

// compareVersions compares two version strings
// Returns -1 if a < b, 0 if a == b, 1 if a > b
func compareVersions(a, b string) int {
	aParts := parseVersion(a)
	bParts := parseVersion(b)

	for i := 0; i < 4; i++ {
		if aParts[i] < bParts[i] {
			return -1
		}
		if aParts[i] > bParts[i] {
			return 1
		}
	}
	return 0
}

// parseVersion parses version string into [major, minor, patch, p-number]
func parseVersion(v string) [4]int {
	matches := versionRegex.FindStringSubmatch(v)
	if len(matches) != 5 {
		return [4]int{}
	}

	var parts [4]int
	fmt.Sscanf(matches[1], "%d", &parts[0])
	fmt.Sscanf(matches[2], "%d", &parts[1])
	fmt.Sscanf(matches[3], "%d", &parts[2])
	fmt.Sscanf(matches[4], "%d", &parts[3])
	return parts
}
