# CheckMK API Spec

Go types generated from CheckMK REST API OpenAPI specifications, with automatic baseline detection to minimize generated code while maintaining full version coverage.

## Overview

This repository:
- Tracks **all CheckMK patch versions** (2.2.x, 2.3.x, 2.4.x) in a manifest
- Stores **only baseline specs** where the API actually changed
- Generates **Go types for baselines only** (~42 packages instead of ~100)
- Maps any version to its baseline at runtime

## How It Works

```
CheckMK 2.4.0p1  ──┐
CheckMK 2.4.0p2  ──┼── Same API ──► Baseline: 2.4.0p1 ──► v2_4_0p1 package
CheckMK 2.4.0p3  ──┘
CheckMK 2.4.0p4  ────── API changed ──► Baseline: 2.4.0p4 ──► v2_4_0p4 package
```

The `manifest.json` maps every known version to its baseline, so your code can use any CheckMK version and get the correct types.

## Installation

```go
import (
    "github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_4_0p17"
)
```

## Quick Start

```go
package main

import (
    "fmt"
    v24 "github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_4_0p17"
)

func main() {
    // Get valid values for host tag_agent attribute
    validAgents := v24.ValidHostTagAgentValues()
    fmt.Println("Valid agent types:", validAgents)

    // Get all valid host attribute field names
    fields := v24.HostCreateAttributeFieldNames
    fmt.Println("Available attributes:", fields)
}
```

## Available Baselines

See `manifest.json` for the complete mapping. Current baseline counts:
- **2.2.0**: 18 baselines (p1, p3, p4, p5, p8, ...)
- **2.3.0**: 17 baselines (p1, p3, p5, p7, ...)
- **2.4.0**: 7 baselines (p1, p6, p11, p14, p16, p17, p18)

## Manifest Structure

```json
{
  "baselines": ["2.2.0p1", "2.2.0p3", ...],
  "mapping": {
    "2.4.0p1": {
      "spec": "2.4.0/p1.yaml",
      "baseline": "2.4.0p1",
      "package": "v2_4_0p1",
      "is_baseline": true,
      "max_severity": "initial"
    },
    "2.4.0p2": {
      "spec": "2.4.0/p1.yaml",
      "baseline": "2.4.0p1",
      "package": "v2_4_0p1",
      "is_baseline": false,
      "max_severity": "docs"
    }
  }
}
```

## Tools

| Tool | Purpose |
|------|---------|
| `spec-sync` | Sync specs from Docker Hub, manage manifest, cleanup |
| `openapi-gen` | Generate Go types from OpenAPI specs |
| `openapi-diff` | Compare specs with severity classification |
| `openapi-filter` | Filter specs to specific endpoints |
| `version-types-gen` | Generate version_types.go mapping |
| `schema-check` | Validate generated types against specs |
| `testdata-gen` | Generate test fixtures from spec constraints |

## Common Commands

```bash
# Build all tools
make build

# Check for new CheckMK versions on Docker Hub
make sync-dry-run

# Fetch new versions and update manifest
make sync

# Remove non-baseline spec files
make sync-cleanup

# Generate Go types for all baselines
make baselines

# Full pipeline: sync + cleanup + generate
make full-pipeline

# Compare two versions
make diff OLD=2.3.0p41 NEW=2.4.0p17
```

## Automated Updates

A GitHub Actions workflow runs weekly to:
1. Check Docker Hub for new CheckMK versions
2. Fetch specs and compare against latest baseline
3. Update manifest (new baselines or version mappings)
4. Generate types for new baselines
5. Create a PR with changes

See `.github/workflows/update-specs.yml`.

## Fetching from a Specific Instance

To fetch a spec from your own CheckMK instance:

```bash
./scripts/fetch-spec.sh \
  -url https://your-checkmk.local \
  -version 2.5.0p1 \
  -user automation \
  -pass YOUR_PASSWORD
```

## Resources Covered

Generated types include:
- **Host** - Host configuration and attributes
- **Folder** - Folder structure and properties
- **AuxTag** - Auxiliary tags
- **TagGroup** - Tag groups
- **User** - User accounts
- **ContactGroup** - Contact groups

## Severity Levels

When comparing specs, changes are classified:

| Severity | Triggers New Baseline | Examples |
|----------|----------------------|----------|
| `breaking` | Yes | Field removed, type changed, required added |
| `minor` | Yes | Field added, endpoint added |
| `deprecated` | No | Field marked deprecated |
| `docs` | No | Description changed only |
| `none` | No | No changes detected |

## License

Apache 2.0
