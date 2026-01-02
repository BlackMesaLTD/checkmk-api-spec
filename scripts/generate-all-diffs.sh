#!/bin/bash
# Generate diffs for all consecutive baselines and create changelog.json

set -e

MANIFEST="manifest.json"
DIFFS_DIR="diffs"
CHANGELOG="changelog.json"

# Check dependencies
if ! command -v jq &> /dev/null; then
    echo "Error: jq is required but not installed"
    exit 1
fi

if [ ! -f "$MANIFEST" ]; then
    echo "Error: $MANIFEST not found"
    exit 1
fi

if [ ! -f "./bin/openapi-diff" ]; then
    echo "Error: openapi-diff not built. Run 'make build' first"
    exit 1
fi

# Create diffs directory
mkdir -p "$DIFFS_DIR"

# Get baselines sorted by version
BASELINES=$(jq -r '.baselines[]' "$MANIFEST" | sort -V)
BASELINES_ARRAY=($BASELINES)
TOTAL=${#BASELINES_ARRAY[@]}

echo "Found $TOTAL baselines"
echo ""

# Initialize changelog array file
echo "[]" > "$DIFFS_DIR/_changelog_temp.json"

# Process first version (initial)
FIRST_VERSION="${BASELINES_ARRAY[0]}"
echo "[0/$((TOTAL-1))] $FIRST_VERSION (initial)"
echo '[{"version":"'"$FIRST_VERSION"'","previous":null,"diff":{"max_severity":"initial","summary":{"note":"Initial baseline version"}}}]' > "$DIFFS_DIR/_changelog_temp.json"

# Generate individual diffs for consecutive versions
for ((i=1; i<TOTAL; i++)); do
    OLD_VERSION="${BASELINES_ARRAY[$i-1]}"
    NEW_VERSION="${BASELINES_ARRAY[$i]}"

    # Extract minor/patch for file paths
    OLD_MINOR=$(echo "$OLD_VERSION" | sed 's/p.*//')
    OLD_PATCH=$(echo "$OLD_VERSION" | sed 's/.*\(p[0-9]*\)/\1/')
    NEW_MINOR=$(echo "$NEW_VERSION" | sed 's/p.*//')
    NEW_PATCH=$(echo "$NEW_VERSION" | sed 's/.*\(p[0-9]*\)/\1/')

    OLD_SPEC="specs/$OLD_MINOR/$OLD_PATCH.yaml"
    NEW_SPEC="specs/$NEW_MINOR/$NEW_PATCH.yaml"
    DIFF_FILE="$DIFFS_DIR/${NEW_VERSION}.json"

    if [ ! -f "$OLD_SPEC" ]; then
        echo "  Skipping: $OLD_SPEC not found"
        continue
    fi
    if [ ! -f "$NEW_SPEC" ]; then
        echo "  Skipping: $NEW_SPEC not found"
        continue
    fi

    echo "[$i/$((TOTAL-1))] $OLD_VERSION -> $NEW_VERSION"

    # Run diff (output to version-named file for easier lookup)
    ./bin/openapi-diff \
        -old "$OLD_SPEC" \
        -new "$NEW_SPEC" \
        -output "$DIFF_FILE" 2>/dev/null || true

    if [ -f "$DIFF_FILE" ]; then
        # Create entry and append to changelog
        jq --arg v "$NEW_VERSION" --arg prev "$OLD_VERSION" '{
            version: $v,
            previous: $prev,
            diff: {
                max_severity: .max_severity,
                summary: .summary,
                category_summary: .category_summary,
                path_changes: {
                    added: (.path_changes.added // [] | length),
                    removed: (.path_changes.removed // [] | length)
                },
                schema_changes: (.schema_changes // [] | length)
            },
            provider_relevant: (
                (.category_summary.setup.paths_added // 0) > 0 or
                (.category_summary.setup.paths_removed // 0) > 0 or
                (.category_summary.setup.breaking // 0) > 0 or
                (.summary.breaking_changes // 0) > 0 or
                (.summary.schemas_added // 0) > 0 or
                (.summary.schemas_removed // 0) > 0
            )
        }' "$DIFF_FILE" > "$DIFFS_DIR/_entry_temp.json"

        # Append to changelog
        jq --slurpfile entry "$DIFFS_DIR/_entry_temp.json" '. + $entry' "$DIFFS_DIR/_changelog_temp.json" > "$DIFFS_DIR/_changelog_temp2.json"
        mv "$DIFFS_DIR/_changelog_temp2.json" "$DIFFS_DIR/_changelog_temp.json"
    fi
done

# Finalize changelog.json with metadata
jq '{
    generated: (now | strftime("%Y-%m-%dT%H:%M:%SZ")),
    total_versions: (. | length),
    changelog: (. | sort_by(.version) | reverse)
}' "$DIFFS_DIR/_changelog_temp.json" > "$CHANGELOG"

# Cleanup temp files
rm -f "$DIFFS_DIR/_changelog_temp.json" "$DIFFS_DIR/_entry_temp.json"

echo ""
echo "Generated $CHANGELOG with $TOTAL versions"
echo "Individual diffs saved to $DIFFS_DIR/"
