package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	dockerHubAPI   = "https://hub.docker.com/v2/repositories/checkmk/check-mk-raw/tags"
	dockerImage    = "checkmk/check-mk-raw"
	siteName       = "cmk"
	containerPort  = "5000"
	apiVersion     = "1.0"
	fetchTimeout   = 120 * time.Second
	startupTimeout = 120 * time.Second
)

var (
	versionTagRegex = regexp.MustCompile(`^2\.[234]\.[0-9]+p[0-9]+$`)
	credentials     = []string{"cmkadmin:test123", "automation:test123"}
)

// DockerHubResponse represents the Docker Hub API response
type DockerHubResponse struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next string `json:"next"`
}

// GetDockerHubVersions fetches all available versions from Docker Hub
func GetDockerHubVersions() ([]string, error) {
	var allVersions []string
	url := dockerHubAPI + "?page_size=100"

	client := &http.Client{Timeout: 30 * time.Second}

	for url != "" {
		resp, err := client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch Docker Hub tags: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response: %w", err)
		}

		var hubResp DockerHubResponse
		if err := json.Unmarshal(body, &hubResp); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}

		for _, result := range hubResp.Results {
			if versionTagRegex.MatchString(result.Name) {
				allVersions = append(allVersions, result.Name)
			}
		}

		url = hubResp.Next
	}

	// Sort versions
	sortVersions(allVersions)
	return allVersions, nil
}

// FetchSpecFromDocker fetches the OpenAPI spec from a Docker container
func FetchSpecFromDocker(version string, specsDir string, verbose bool) ([]byte, error) {
	image := fmt.Sprintf("%s:%s", dockerImage, version)
	containerName := fmt.Sprintf("spec-fetch-%s", strings.ReplaceAll(version, ".", "-"))

	// Cleanup any existing container with same name
	exec.Command("docker", "rm", "-f", containerName).Run()

	if verbose {
		fmt.Printf("  Pulling image %s...\n", image)
	}

	// Pull the image
	pullCmd := exec.Command("docker", "pull", image)
	if err := pullCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to pull image: %w", err)
	}

	if verbose {
		fmt.Printf("  Starting container...\n")
	}

	// Run the container
	runCmd := exec.Command("docker", "run", "-d",
		"--name", containerName,
		"-e", "CMK_SITE_ID="+siteName,
		"-e", "CMK_PASSWORD=test123",
		image)

	output, err := runCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}
	containerID := strings.TrimSpace(string(output))

	// Ensure cleanup
	defer func() {
		if verbose {
			fmt.Printf("  Cleaning up...\n")
		}
		exec.Command("docker", "rm", "-f", containerID).Run()
		exec.Command("docker", "rmi", image).Run()
	}()

	// Wait for container to be healthy
	if err := waitForContainer(containerID, verbose); err != nil {
		return nil, err
	}

	// Extra wait for site initialization
	time.Sleep(10 * time.Second)

	// Fetch the spec
	specURL := fmt.Sprintf("http://localhost:%s/%s/check_mk/api/%s/openapi-swagger-ui.yaml",
		containerPort, siteName, apiVersion)

	var specData []byte
	for _, creds := range credentials {
		curlCmd := exec.Command("docker", "exec", containerID,
			"curl", "-s", "-u", creds, specURL)

		data, err := curlCmd.Output()
		if err == nil && len(data) > 0 && strings.Contains(string(data), "info:") {
			specData = data
			break
		}
	}

	if len(specData) == 0 {
		return nil, fmt.Errorf("failed to fetch spec from container")
	}

	// Validate spec
	if !strings.Contains(string(specData), "info:") {
		return nil, fmt.Errorf("response doesn't look like an OpenAPI spec")
	}

	return specData, nil
}

// waitForContainer waits for the CheckMK container to be ready
func waitForContainer(containerID string, verbose bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), startupTimeout)
	defer cancel()

	healthURL := fmt.Sprintf("http://localhost:%s/%s/check_mk/login.py", containerPort, siteName)

	for {
		select {
		case <-ctx.Done():
			// Get logs for debugging
			logsCmd := exec.Command("docker", "logs", "--tail", "20", containerID)
			logs, _ := logsCmd.Output()
			return fmt.Errorf("timeout waiting for container: %s", string(logs))
		default:
			// Check if container is still running
			inspectCmd := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", containerID)
			output, err := inspectCmd.Output()
			if err != nil || strings.TrimSpace(string(output)) != "true" {
				return fmt.Errorf("container stopped unexpectedly")
			}

			// Check health endpoint
			curlCmd := exec.Command("docker", "exec", containerID,
				"curl", "-s", "-o", "/dev/null", "-w", "%{http_code}", healthURL)
			httpCode, _ := curlCmd.Output()

			code := strings.TrimSpace(string(httpCode))
			if code == "200" || code == "302" {
				if verbose {
					fmt.Printf("  Container ready (HTTP %s)\n", code)
				}
				return nil
			}

			if verbose {
				fmt.Printf("  Waiting for container... (HTTP %s)\n", code)
			}
			time.Sleep(5 * time.Second)
		}
	}
}

// SaveSpec saves spec data to the appropriate file
func SaveSpec(version string, specsDir string, data []byte) error {
	specPath := versionToSpecPath(specsDir, version)

	// Create directory if needed
	dir := filepath.Dir(specPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(specPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write spec: %w", err)
	}

	return nil
}

// FindMissingVersions returns versions that exist on Docker Hub but not in manifest
func FindMissingVersions(manifest *Manifest, minor string) ([]string, error) {
	hubVersions, err := GetDockerHubVersions()
	if err != nil {
		return nil, err
	}

	var missing []string
	for _, v := range hubVersions {
		// Filter by minor if specified
		if minor != "" && !strings.HasPrefix(v, minor) {
			continue
		}

		// Check if already in manifest
		if _, exists := manifest.Versions[v]; !exists {
			missing = append(missing, v)
		}
	}

	sort.Slice(missing, func(i, j int) bool {
		return compareVersions(missing[i], missing[j]) < 0
	})

	return missing, nil
}
