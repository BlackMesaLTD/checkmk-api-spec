package main

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestSchemaUnmarshalTypeForms(t *testing.T) {
	cases := []struct {
		name         string
		yaml         string
		wantType     string
		wantNullable bool
	}{
		{"3.0 scalar", "type: string\nnullable: true\n", "string", true},
		{"3.0 scalar no null", "type: integer\n", "integer", false},
		{"3.1 array nullable", "type:\n  - string\n  - 'null'\n", "string", true},
		{"3.1 array single", "type:\n  - boolean\n", "boolean", false},
		{"no type", "format: int32\n", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s Schema
			if err := yaml.Unmarshal([]byte(tc.yaml), &s); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if s.Type != tc.wantType {
				t.Errorf("Type = %q, want %q", s.Type, tc.wantType)
			}
			if s.Nullable != tc.wantNullable {
				t.Errorf("Nullable = %v, want %v", s.Nullable, tc.wantNullable)
			}
		})
	}
}
