.PHONY: all build clean generate help baselines sync sync-dry-run sync-bootstrap sync-cleanup union-descriptions

# Default target
all: build

# Build all tools
build:
	@echo "Building tools..."
	go build -o bin/openapi-gen ./cmd/openapi-gen
	go build -o bin/openapi-filter ./cmd/openapi-filter
	go build -o bin/openapi-diff ./cmd/openapi-diff
	go build -o bin/version-types-gen ./cmd/version-types-gen
	go build -o bin/schema-check ./cmd/schema-check
	go build -o bin/testdata-gen ./cmd/testdata-gen
	go build -o bin/spec-sync ./cmd/spec-sync
	go build -o bin/description-union-gen ./cmd/description-union-gen
	@echo "Done."

# Sync specs from Docker Hub (check for new versions)
sync: build
	./bin/spec-sync

# Sync with dry-run (show what would be done)
sync-dry-run: build
	./bin/spec-sync --dry-run

# Bootstrap manifest from existing specs (no Docker)
sync-bootstrap: build
	./bin/spec-sync --bootstrap

# Remove non-baseline spec files
sync-cleanup: build
	./bin/spec-sync --cleanup

# Generate types for all baseline versions
baselines: build
	./scripts/generate-baselines.sh

# Generate version_types.go from manifest.json
version-types: build
	./bin/version-types-gen \
		-baselines manifest.json \
		-output generated/go/version_types.go \
		-package types \
		-module github.com/BlackMesaLTD/checkmk-api-spec/generated/go

# Generate types for a specific version (generates ALL schemas)
# Usage: make generate-version VERSION=2.4.0p17
generate-version: build
	@if [ -z "$(VERSION)" ]; then echo "Usage: make generate-version VERSION=2.4.0p17"; exit 1; fi
	$(eval MINOR := $(shell echo $(VERSION) | sed 's/p.*//' ))
	$(eval PATCH := $(shell echo $(VERSION) | sed 's/.*p/p/' ))
	$(eval MINOR_DIR := v$(subst .,_,$(MINOR)))
	./bin/openapi-gen \
		-spec specs/$(MINOR)/$(PATCH).yaml \
		-output generated/go/$(MINOR_DIR)/$(PATCH)/ \
		-package $(PATCH)
	go fmt ./generated/go/...

# Compare two specific versions
# Usage: make diff OLD=2.2.0p43 NEW=2.3.0p41
diff: build
	@if [ -z "$(OLD)" ] || [ -z "$(NEW)" ]; then \
		echo "Usage: make diff OLD=2.2.0p43 NEW=2.3.0p41"; \
		exit 1; \
	fi
	@mkdir -p diffs
	$(eval OLD_MINOR := $(shell echo $(OLD) | sed 's/p.*//' ))
	$(eval OLD_PATCH := $(shell echo $(OLD) | sed 's/.*p/p/' ))
	$(eval NEW_MINOR := $(shell echo $(NEW) | sed 's/p.*//' ))
	$(eval NEW_PATCH := $(shell echo $(NEW) | sed 's/.*p/p/' ))
	./bin/openapi-diff \
		-old specs/$(OLD_MINOR)/$(OLD_PATCH).yaml \
		-new specs/$(NEW_MINOR)/$(NEW_PATCH).yaml \
		-output diffs/$(OLD)-to-$(NEW).json

# Check schema completeness for a version
# Usage: make check VERSION=2.4.0p17
check: build
	@if [ -z "$(VERSION)" ]; then echo "Usage: make check VERSION=2.4.0p17"; exit 1; fi
	$(eval MINOR := $(shell echo $(VERSION) | sed 's/p.*//' ))
	$(eval PATCH := $(shell echo $(VERSION) | sed 's/.*p/p/' ))
	$(eval PKG := v$(shell echo $(VERSION) | sed 's/\./_/g'))
	./bin/schema-check \
		-spec specs/$(MINOR)/$(PATCH).yaml \
		-types generated/go/$(PKG)/types.gen.go \
		-schema HostConfig

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf tmp/
	rm -rf diffs/

# Clean generated code (careful!)
clean-generated:
	rm -rf generated/go/

# Generate union descriptions from all baselines
union-descriptions: build
	./bin/description-union-gen \
		-manifest manifest.json \
		-output generated/go/union \
		-generated generated/go

# Full pipeline: sync specs, cleanup, generate types
full-pipeline: build
	@echo "=== Full Pipeline ==="
	@echo "Step 1: Syncing specs from Docker Hub..."
	./bin/spec-sync
	@echo ""
	@echo "Step 2: Cleaning up non-baseline specs..."
	./bin/spec-sync --cleanup
	@echo ""
	@echo "Step 3: Generating types for baselines..."
	./scripts/generate-baselines.sh
	@echo ""
	@echo "Pipeline complete!"

# Help
help:
	@echo "CheckMK API Spec - Makefile targets:"
	@echo ""
	@echo "  Build:"
	@echo "    make build              - Build all tools"
	@echo "    make clean              - Remove build artifacts"
	@echo ""
	@echo "  Sync specs (from Docker Hub):"
	@echo "    make sync               - Fetch new versions, update manifest"
	@echo "    make sync-dry-run       - Show what sync would do"
	@echo "    make sync-bootstrap     - Build manifest from existing specs"
	@echo "    make sync-cleanup       - Remove non-baseline spec files"
	@echo ""
	@echo "  Generate types:"
	@echo "    make baselines          - Generate types for all baselines"
	@echo "    make version-types      - Generate version_types.go mapping"
	@echo "    make union-descriptions - Generate merged field descriptions"
	@echo "    make generate-version VERSION=x  - Generate for specific version"
	@echo ""
	@echo "  Utilities:"
	@echo "    make diff OLD=x NEW=y   - Compare two versions"
	@echo "    make check VERSION=x    - Check schema completeness"
	@echo ""
	@echo "  Pipeline:"
	@echo "    make full-pipeline      - Sync + cleanup + generate"
