#!/bin/bash
# Generate Go types for baseline versions only.
#
# This script:
# 1. Reads baseline versions from manifest.json (created by spec-sync)
# 2. Runs openapi-gen for each baseline version
# 3. Runs version-types-gen to create the version mapping code
#
# Usage:
#   ./scripts/generate-baselines.sh
#
# Prerequisites:
#   manifest.json must exist (run ./bin/spec-sync --bootstrap first)

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
cd "$PROJECT_DIR"

# Configuration
SPECS_DIR="specs"
GENERATED_DIR="generated/go"
MANIFEST_FILE="manifest.json"
MODULE_PATH="github.com/BlackMesaLTD/checkmk-api-spec/generated/go"

# Generate ALL schemas (no filtering)
# Previously used: RESOURCES="host,folder,aux_tag,tag_group,user,contact_group"
# Now generates all schemas from OpenAPI spec for maximum coverage

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --gen-only)
            # Kept for backwards compatibility, no-op now
            shift
            ;;
        --help)
            echo "Usage: $0 [OPTIONS]"
            echo ""
            echo "Generates Go types from OpenAPI specs for all baseline versions."
            echo "Requires manifest.json (created by spec-sync --bootstrap)."
            echo ""
            echo "Options:"
            echo "  --gen-only       Backwards compatibility, no-op"
            echo "  --help           Show this help"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Ensure binaries exist
check_binary() {
    if [ ! -x "bin/$1" ]; then
        echo "Building $1..."
        go build -o "bin/$1" "./cmd/$1/"
    fi
}

check_binary "openapi-gen"
check_binary "version-types-gen"

# Check manifest exists (spec-sync creates this)
if [ ! -f "$MANIFEST_FILE" ]; then
    echo "Error: $MANIFEST_FILE not found."
    echo "Run './bin/spec-sync --bootstrap' to create it from existing specs."
    exit 1
fi

echo "Using manifest: $MANIFEST_FILE"
echo "Baselines: $(jq '.baselines | length' "$MANIFEST_FILE")"

# Step 2: Generate types for each baseline
echo "=================================================="
echo "Step 2: Generating types for baseline versions"
echo "=================================================="

# Extract baselines from manifest
BASELINES=$(jq -r '.baselines[]' "$MANIFEST_FILE")

for baseline in $BASELINES; do
    # Convert version to hierarchical path: 2.4.0p17 -> v2_4_0/p17
    minor="${baseline%p*}"        # e.g., 2.4.0
    patch="${baseline#*p}"        # e.g., 17
    minor_dir="v$(echo "$minor" | tr '.' '_')"  # e.g., v2_4_0
    pkg="p$patch"                 # e.g., p17
    spec_file="$SPECS_DIR/$minor/p$patch.yaml"
    output_dir="$GENERATED_DIR/$minor_dir/$pkg"

    if [ ! -f "$spec_file" ]; then
        echo "Warning: Spec not found for $baseline at $spec_file, skipping"
        continue
    fi

    echo ""
    echo "Generating types for $baseline -> $minor_dir/$pkg"

    # Create output directory
    mkdir -p "$output_dir"

    # Build tag for version-specific compilation (e.g., checkmk_v2_4_0)
    build_tag="checkmk_$minor_dir"

    # Run openapi-gen (generates ALL schemas from the OpenAPI spec)
    ./bin/openapi-gen \
        -spec "$spec_file" \
        -output "$output_dir/" \
        -package "$pkg" \
        -buildtag "$build_tag"

    echo "  Output: $output_dir (build tag: $build_tag)"
done

# Step 3: Generate version_types.go
echo ""
echo "=================================================="
echo "Step 3: Generating version_types.go"
echo "=================================================="

./bin/version-types-gen \
    -baselines "$MANIFEST_FILE" \
    -output "$GENERATED_DIR/version_types.go" \
    -package "types" \
    -module "$MODULE_PATH"

echo ""
echo "=================================================="
echo "Summary"
echo "=================================================="
echo "Manifest file: $MANIFEST_FILE"
echo "Generated packages:"
for baseline in $BASELINES; do
    minor="${baseline%p*}"
    patch="${baseline#*p}"
    minor_dir="v$(echo "$minor" | tr '.' '_')"
    echo "  - $GENERATED_DIR/$minor_dir/p$patch"
done
echo "Version mapping: $GENERATED_DIR/version_types.go"
echo ""
echo "Done!"
