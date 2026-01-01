# CheckMK API Spec

Go types generated from CheckMK REST API OpenAPI specifications.

## Overview

This repository contains:
- **OpenAPI specs** for CheckMK versions 2.2, 2.3, and 2.4
- **Generated Go types** for use in Terraform providers and other Go applications
- **Generation tools** for creating types from specs

## Installation

```go
import (
    v2_4_0p17 "github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_4_0p17"
)
```

## Available Versions

| CheckMK Version | Package |
|-----------------|---------|
| 2.4.0p17 | `github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_4_0p17` |
| 2.3.0p41 | `github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_3_0p41` |
| 2.2.0p43 | `github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_2_0p43` |

## Usage

### Using Pre-Generated Types

```go
package main

import (
    "fmt"
    v24 "github.com/BlackMesaLTD/checkmk-api-spec/generated/go/v2_4_0p17"
)

func main() {
    // Get valid values for tag_agent
    validAgents := v24.ValidHostTagAgentValues()
    fmt.Println("Valid agent types:", validAgents)

    // Get all valid host attribute field names
    fields := v24.HostCreateAttributeFieldNames
    fmt.Println("Available attributes:", fields)
}
```

### Generating Types for Untested Versions

If you're running a CheckMK version not yet in this repository:

1. **Fetch the OpenAPI spec from your instance:**
   ```bash
   ./scripts/fetch-spec.sh -url https://your-checkmk.local -version 2.5.0p1 -user automation -pass YOUR_PASSWORD
   ```

2. **Generate types:**
   ```bash
   make generate-version VERSION=2.5.0p1
   ```

3. **Use in your project with go mod replace:**
   ```
   // go.mod
   replace github.com/BlackMesaLTD/checkmk-api-spec => /path/to/local/checkmk-api-spec
   ```

## Tools

| Tool | Purpose |
|------|---------|
| `openapi-gen` | Generate Go types from OpenAPI specs |
| `openapi-filter` | Filter spec to specific resource endpoints |
| `openapi-diff` | Compare schemas across versions |
| `schema-check` | Validate completeness of generated types |
| `testdata-gen` | Generate test fixtures from spec constraints |

### Building Tools

```bash
make build
```

### Fetching Specs

```bash
# From running CheckMK containers (requires docker)
make fetch-specs

# From a specific instance
./scripts/fetch-spec.sh -url https://checkmk.example.com -version 2.4.0p17
```

### Generating Types

```bash
# Generate for all versions
make generate

# Generate for a specific version
make generate-version VERSION=2.4.0p17
```

## Type Mode in Terraform Provider

The Terraform provider supports two modes:

| Mode | Behavior |
|------|----------|
| `static` | Validate against generated types (default) |
| `hollow` | Accept any attributes, rely on API validation |

```hcl
provider "checkmk" {
  url       = "https://checkmk.example.com"
  username  = "automation"
  password  = var.password
  type_mode = "static"  # or "hollow"
}
```

## Resources Covered

Currently generates types for:
- **Host** - Host configuration and attributes
- **Folder** - Folder structure and properties
- **AuxTag** - Auxiliary tags

## License

Apache 2.0
