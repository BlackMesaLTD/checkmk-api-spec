#!/bin/bash
set -e

# Fetch OpenAPI spec from CheckMK instance
#
# Usage:
#   ./fetch-spec.sh -url https://checkmk.example.com -version 2.4.0p17 [-user automation] [-pass PASSWORD]
#   ./fetch-spec.sh --docker  # Fetch from local docker containers

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# Default values
SITE="test"
USER="automation"
PASS=""
URL=""
VERSION=""
DOCKER_MODE=false

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -url|--url)
            URL="$2"
            shift 2
            ;;
        -version|--version)
            VERSION="$2"
            shift 2
            ;;
        -user|--user)
            USER="$2"
            shift 2
            ;;
        -pass|--pass)
            PASS="$2"
            shift 2
            ;;
        -site|--site)
            SITE="$2"
            shift 2
            ;;
        --docker)
            DOCKER_MODE=true
            shift
            ;;
        -h|--help)
            echo "Usage: $0 [OPTIONS]"
            echo ""
            echo "Fetch OpenAPI spec from CheckMK instance"
            echo ""
            echo "Options:"
            echo "  -url URL        CheckMK base URL (e.g., https://checkmk.example.com)"
            echo "  -version VER    CheckMK version (e.g., 2.4.0p17)"
            echo "  -user USER      API username (default: automation)"
            echo "  -pass PASS      API password/secret"
            echo "  -site SITE      CheckMK site name (default: test)"
            echo "  --docker        Fetch from local Docker containers"
            echo "  -h, --help      Show this help"
            exit 0
            ;;
        *)
            echo -e "${RED}Unknown option: $1${NC}"
            exit 1
            ;;
    esac
done

fetch_from_url() {
    local url="$1"
    local version="$2"
    local user="$3"
    local pass="$4"
    local site="$5"

    local spec_url="${url}/${site}/check_mk/api/1.0/openapi-swagger-ui.yaml"
    local output_dir="${PROJECT_ROOT}/specs/${version}"
    local output_file="${output_dir}/openapi.yaml"

    echo -e "${GREEN}Fetching spec from: ${spec_url}${NC}"

    mkdir -p "$output_dir"

    if curl -sf -u "${user}:${pass}" "$spec_url" -o "$output_file"; then
        echo -e "${GREEN}Saved to: ${output_file}${NC}"
        # Show spec size
        local size=$(wc -c < "$output_file")
        echo -e "Spec size: ${size} bytes"
    else
        echo -e "${RED}Failed to fetch spec from ${spec_url}${NC}"
        return 1
    fi
}

fetch_from_docker() {
    echo -e "${GREEN}Fetching specs from Docker containers...${NC}"

    # Container configurations
    declare -A CONTAINERS=(
        ["2.2.0p43"]="checkmk-2.2.0p43:5020"
        ["2.3.0p41"]="checkmk-2.3.0p41:5030"
        ["2.4.0p17"]="checkmk-2.4.0p17:5040"
    )

    # Try to get password from credential service
    local cred_url="http://localhost:5099/credentials"
    local default_pass="testSecret123"

    for version in "${!CONTAINERS[@]}"; do
        local container_info="${CONTAINERS[$version]}"
        local container_name="${container_info%%:*}"
        local port="${container_info##*:}"
        local url="http://localhost:${port}"

        echo -e "\n${YELLOW}Processing ${version}...${NC}"

        # Check if container is running
        if ! docker ps --format '{{.Names}}' | grep -q "^${container_name}$"; then
            echo -e "${YELLOW}Container ${container_name} not running, skipping${NC}"
            continue
        fi

        # Try to get password from credential service
        local pass="$default_pass"
        if curl -sf "${cred_url}/${version:0:3}" > /dev/null 2>&1; then
            local cred_json=$(curl -sf "${cred_url}/${version:0:3}")
            if [ -n "$cred_json" ]; then
                pass=$(echo "$cred_json" | grep -o '"password":"[^"]*"' | cut -d'"' -f4)
                [ -z "$pass" ] && pass="$default_pass"
            fi
        fi

        fetch_from_url "$url" "$version" "$USER" "$pass" "$SITE"
    done
}

# Main
if [ "$DOCKER_MODE" = true ]; then
    fetch_from_docker
elif [ -n "$URL" ] && [ -n "$VERSION" ]; then
    if [ -z "$PASS" ]; then
        echo -e "${RED}Error: Password required (-pass)${NC}"
        exit 1
    fi
    fetch_from_url "$URL" "$VERSION" "$USER" "$PASS" "$SITE"
else
    echo -e "${RED}Error: Either --docker or -url and -version required${NC}"
    echo "Run with -h for help"
    exit 1
fi

echo -e "\n${GREEN}Done!${NC}"
