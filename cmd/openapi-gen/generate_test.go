package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// generateTypes runs the full generator over an in-memory spec and returns the
// contents of the generated types.gen.go file.
func generateTypes(t *testing.T, spec string) string {
	t.Helper()
	dir := t.TempDir()
	specPath := filepath.Join(dir, "spec.yaml")
	if err := os.WriteFile(specPath, []byte(spec), 0644); err != nil {
		t.Fatalf("write spec: %v", err)
	}
	outDir := filepath.Join(dir, "out")

	gen := &Generator{
		packageName:     "gen",
		outputDir:       outDir,
		excludeFields:   map[string]bool{},
		enumsFound:      make(map[string]*EnumInfo),
		fieldsFound:     make(map[string][]string),
		fieldsMeta:      make(map[string][]FieldMetadata),
		requiredFound:   make(map[string][]string),
		readOnlyFound:   make(map[string][]string),
		deprecatedFound: make(map[string][]string),
		descriptions:    make(map[string]map[string]string),
		fieldTypes:      make(map[string]map[string]string),
		generatedTypes:  make(map[string]bool),
	}
	if err := gen.LoadSpec(specPath); err != nil {
		t.Fatalf("load spec: %v", err)
	}
	if err := gen.Generate(); err != nil {
		t.Fatalf("generate: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(outDir, "types.gen.go"))
	if err != nil {
		t.Fatalf("read generated types: %v", err)
	}
	return string(data)
}

// TestGenerate31TypeForms verifies OpenAPI 3.1 array-form types
// (["string","null"]) generate concrete Go types, including inside array
// items, rather than failing or falling back to interface{}.
func TestGenerate31TypeForms(t *testing.T) {
	spec := `openapi: 3.1.1
info:
  title: t
  version: v1
paths: {}
components:
  schemas:
    Thing:
      type: object
      properties:
        top_nullable:
          type: [string, "null"]
        multi:
          type: [string, integer, "null"]
        list:
          type: array
          items:
            type: [string, "null"]
`
	out := generateTypes(t, spec)
	for _, want := range []string{
		"TopNullable string",
		"Multi string", // first non-null member wins
		"List []string",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("generated types missing %q\n---\n%s", want, out)
		}
	}
}

// TestGenerate30And31Equivalent is the 3.0 regression guard: a 3.0 schema
// (scalar type + nullable: true) and the equivalent 3.1 schema (type array
// with "null") must generate the same Go field types.
func TestGenerate30And31Equivalent(t *testing.T) {
	v30 := `openapi: 3.0.2
info: {title: t, version: v1}
paths: {}
components:
  schemas:
    Thing:
      type: object
      properties:
        name:
          type: string
          nullable: true
        count:
          type: integer
`
	v31 := `openapi: 3.1.1
info: {title: t, version: v1}
paths: {}
components:
  schemas:
    Thing:
      type: object
      properties:
        name:
          type: [string, "null"]
        count:
          type: integer
`
	out30 := generateTypes(t, v30)
	out31 := generateTypes(t, v31)
	for _, want := range []string{"Name string", "Count int"} {
		if !strings.Contains(out30, want) {
			t.Errorf("3.0 output missing %q", want)
		}
		if !strings.Contains(out31, want) {
			t.Errorf("3.1 output missing %q", want)
		}
	}
}
