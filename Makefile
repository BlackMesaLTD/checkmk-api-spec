.PHONY: all build clean generate fetch-specs help

# Default target
all: build

# Build all tools
build:
	@echo "Building tools..."
	go build -o bin/openapi-gen ./cmd/openapi-gen
	go build -o bin/openapi-filter ./cmd/openapi-filter
	go build -o bin/openapi-diff ./cmd/openapi-diff
	go build -o bin/schema-check ./cmd/schema-check
	go build -o bin/testdata-gen ./cmd/testdata-gen
	@echo "Done."

# Fetch OpenAPI specs from running CheckMK containers
fetch-specs:
	@echo "Fetching OpenAPI specs..."
	./scripts/fetch-spec.sh

# Generate types for a specific version
# Usage: make generate-version VERSION=2.4.0p17
generate-version:
	@if [ -z "$(VERSION)" ]; then echo "Usage: make generate-version VERSION=2.4.0p17"; exit 1; fi
	@echo "Generating types for version $(VERSION)..."
	./bin/openapi-gen \
		-spec specs/$(VERSION)/openapi.yaml \
		-output generated/go/v$(subst .,_,$(subst p,p,$(VERSION)))/ \
		-package v$(subst .,_,$(subst p,p,$(VERSION))) \
		-resources host,folder,aux_tag
	go fmt ./generated/go/...

# Generate types for all versions
generate: build
	@echo "Generating types for all versions..."
	@for version in 2.2.0p43 2.3.0p41 2.4.0p17; do \
		if [ -f "specs/$$version/openapi.yaml" ]; then \
			echo "Generating for $$version..."; \
			pkg_name=$$(echo "v$${version}" | sed 's/\./_/g'); \
			./bin/openapi-gen \
				-spec "specs/$$version/openapi.yaml" \
				-output "generated/go/$$pkg_name/" \
				-package "$$pkg_name" \
				-resources host,folder,aux_tag; \
		else \
			echo "Skipping $$version (no spec found)"; \
		fi \
	done
	go fmt ./generated/go/...
	@echo "Done."

# Compare versions
diff:
	@echo "Comparing OpenAPI specs..."
	./bin/openapi-diff \
		-old specs/2.3.0p41/openapi.yaml \
		-new specs/2.4.0p17/openapi.yaml \
		-resources host,folder,aux_tag \
		-output docs/version-diff.json

# Check schema completeness
check:
	@echo "Checking schema completeness..."
	@for version in 2.2.0p43 2.3.0p41 2.4.0p17; do \
		if [ -f "specs/$$version/openapi.yaml" ]; then \
			pkg_name=$$(echo "v$${version}" | sed 's/\./_/g'); \
			echo "Checking $$version..."; \
			./bin/schema-check \
				-spec "specs/$$version/openapi.yaml" \
				-types "generated/go/$$pkg_name/types.gen.go" \
				-schema HostConfig || true; \
		fi \
	done

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf tmp/

# Help
help:
	@echo "CheckMK API Spec - Makefile targets:"
	@echo ""
	@echo "  make build           - Build all tools"
	@echo "  make fetch-specs     - Fetch OpenAPI specs from CheckMK containers"
	@echo "  make generate        - Generate types for all versions"
	@echo "  make generate-version VERSION=2.4.0p17 - Generate for specific version"
	@echo "  make diff            - Compare specs across versions"
	@echo "  make check           - Check schema completeness"
	@echo "  make clean           - Remove build artifacts"
	@echo "  make help            - Show this help"
