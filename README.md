# CheckMK API Spec

Go types generated from CheckMK REST API OpenAPI specifications, with automatic baseline detection to minimize generated code while maintaining full version coverage.

## Overview

This repository:
- Tracks **all CheckMK patch versions** (2.2.x, 2.3.x, 2.4.x) in a manifest
- Stores **only baseline specs** where the API actually changed
- Generates **Go types for baselines only** (~42 packages instead of ~100)
- Maps any version to its baseline at runtime
- Provides **union descriptions** with version annotations for documentation

## How It Works

```
CheckMK 2.4.0p1  ──┐
CheckMK 2.4.0p2  ──┼── Same API ──► Baseline: 2.4.0p1 ──► v2_4_0/p1 package
CheckMK 2.4.0p3  ──┘
CheckMK 2.4.0p4  ────── API changed ──► Baseline: 2.4.0p4 ──► v2_4_0/p4 package
```

The `manifest.json` maps every known version to its baseline, so your code can use any CheckMK version and get the correct types.

## Installation

```go
import (
    // Import a specific baseline package
    "github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_4_0/p17"

    // Or use the version_types.go mapping
    types "github.com/BlackMesaLTD/checkmk-api-spec/generated/go"
)
```

## Quick Start

### Using a Specific Baseline

```go
package main

import (
    "fmt"
    p17 "github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_4_0/p17"
)

func main() {
    // Get valid values for host tag_agent attribute
    validAgents := p17.ValidHostCreateAttributeTagAgentValues()
    fmt.Println("Valid agent types:", validAgents)

    // Get all valid host attribute field names
    fields := p17.HostCreateAttributeFieldNames
    fmt.Println("Available attributes:", fields)

    // Get field descriptions for documentation
    desc := p17.GetFieldDescription("HostCreateAttribute", "tag_agent")
    fmt.Println("Description:", desc)
}
```

### Using Runtime Version Lookup

```go
package main

import (
    "fmt"
    types "github.com/BlackMesaLTD/checkmk-api-spec/generated/go"
)

func main() {
    // Look up baseline for any CheckMK version
    baseline := types.LookupBaseline("2.4.0p15")
    fmt.Println("Baseline:", baseline) // "v2_4_p14"

    // Get valid values for any version
    validAgents := types.ValidHostTagAgentValues(baseline)
    fmt.Println("Valid agents:", validAgents)
}
```

### Using Union Descriptions (Version-Annotated)

```go
package main

import (
    "fmt"
    "github.com/BlackMesaLTD/checkmk-api-spec/generated/go/union"
)

func main() {
    // Get merged description across all versions
    field := union.GetUnionField("HostCreateAttribute", "tag_agent")
    if field != nil {
        fmt.Println("Description:", field.Description)
        fmt.Println("Available since:", field.MinVersion)
        fmt.Println("Markdown:", field.FormatMarkdown())
    }
}
```

## Package Structure

Packages are organized hierarchically by minor version:

```
generated/go/
├── v2_2_0/
│   ├── p1/           # 2.2.0p1 baseline
│   ├── p3/           # 2.2.0p3 baseline
│   └── ...           # 18 baselines total
├── v2_3_0/
│   ├── p1/           # 2.3.0p1 baseline
│   └── ...           # 17 baselines total
├── v2_4_0/
│   ├── p1/           # 2.4.0p1 baseline
│   ├── p17/          # 2.4.0p17 baseline
│   └── ...           # 7 baselines total
├── union/
│   └── descriptions.gen.go  # Merged descriptions across all versions
└── version_types.go  # Runtime version-to-baseline mapping
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
    "2.4.0p17": {
      "baseline": "2.4.0p17",
      "package": "p17",
      "path": "v2_4_0/p17",
      "import_alias": "v2_4_p17",
      "is_baseline": true
    },
    "2.4.0p15": {
      "baseline": "2.4.0p14",
      "package": "p14",
      "path": "v2_4_0/p14",
      "import_alias": "v2_4_p14",
      "is_baseline": false
    }
  }
}
```

## Generated Files Per Baseline

Each baseline package contains:

| File | Purpose |
|------|---------|
| `types.gen.go` | Go struct types for API schemas |
| `enums.gen.go` | Enum types and validator functions |
| `fields.gen.go` | Field name lists and compare keys |
| `metadata.gen.go` | Field descriptions and types |
| `requests.gen.go` | Request builder functions |
| `mappings.gen.go` | API response to Terraform field mappings |

## Tools

| Tool | Purpose |
|------|---------|
| `spec-sync` | Sync specs from Docker Hub, manage manifest, cleanup |
| `openapi-gen` | Generate Go types from OpenAPI specs |
| `openapi-diff` | Compare specs with severity classification |
| `openapi-filter` | Filter specs to specific endpoints |
| `version-types-gen` | Generate version_types.go mapping |
| `description-union-gen` | Generate merged descriptions across all baselines |
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

# Generate version_types.go mapping
make version-types

# Generate union descriptions from all baselines
make union-descriptions

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
