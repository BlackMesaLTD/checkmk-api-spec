// Package types provides version-aware type information for CheckMK API.
package types

import (
	"testing"
)

func TestLookupBaseline(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected BaselinePackage
	}{
		// Exact matches
		{"exact 2.2.0p1", "2.2.0p1", BaselineV2_2_0p1},
		{"exact 2.3.0p41", "2.3.0p41", BaselineV2_3_0p41},
		{"exact 2.4.0p17", "2.4.0p17", BaselineV2_4_0p17},

		// Versions that should map to a baseline (direct map entries)
		{"2.2.0p2 maps to p1", "2.2.0p2", BaselineV2_2_0p1},
		{"2.2.0p42 maps to p33", "2.2.0p42", BaselineV2_2_0p33},
		{"2.3.0p2 maps to p1", "2.3.0p2", BaselineV2_3_0p1},
		{"2.4.0p5 maps to p1", "2.4.0p5", BaselineV2_4_0p1},

		// Versions NOT in map fall back to latest for minor version
		{"2.2.0p10 not in map, falls back to latest", "2.2.0p10", BaselineV2_2_0p44},

		// Latest versions in each major
		{"2.2.0p45", "2.2.0p45", BaselineV2_2_0p44},
		{"2.3.0p42", "2.3.0p42", BaselineV2_3_0p41},
		{"2.4.0p20", "2.4.0p20", BaselineV2_4_0p18},

		// Unknown versions should return empty
		{"unknown 1.0.0", "1.0.0", ""},
		{"unknown 3.0.0", "3.0.0", ""},
		{"invalid format", "invalid", ""},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LookupBaseline(tt.version)
			if result != tt.expected {
				t.Errorf("LookupBaseline(%q) = %q, want %q", tt.version, result, tt.expected)
			}
		})
	}
}

func TestValidHostTagAgentValues(t *testing.T) {
	tests := []struct {
		name           string
		baseline       BaselinePackage
		expectContains []string
		expectNotEmpty bool
	}{
		{
			name:           "2.2.0p43 has expected values",
			baseline:       BaselineV2_2_0p43,
			expectContains: []string{"cmk-agent", "no-agent"},
			expectNotEmpty: true,
		},
		{
			name:           "2.3.0p41 has expected values",
			baseline:       BaselineV2_3_0p41,
			expectContains: []string{"cmk-agent", "no-agent", "all-agents", "special-agents"},
			expectNotEmpty: true,
		},
		{
			name:           "2.4.0p17 has expected values",
			baseline:       BaselineV2_4_0p17,
			expectContains: []string{"cmk-agent", "no-agent", "all-agents", "special-agents"},
			expectNotEmpty: true,
		},
		{
			name:           "empty baseline returns nil",
			baseline:       "",
			expectNotEmpty: false,
		},
		{
			name:           "invalid baseline returns nil",
			baseline:       "invalid",
			expectNotEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidHostTagAgentValues(tt.baseline)

			if tt.expectNotEmpty && len(result) == 0 {
				t.Errorf("ValidHostTagAgentValues(%q) returned empty, expected values", tt.baseline)
			}

			if !tt.expectNotEmpty && len(result) != 0 {
				t.Errorf("ValidHostTagAgentValues(%q) = %v, expected empty", tt.baseline, result)
			}

			for _, expected := range tt.expectContains {
				if !contains(result, expected) {
					t.Errorf("ValidHostTagAgentValues(%q) missing %q, got %v", tt.baseline, expected, result)
				}
			}
		})
	}
}

func TestHostCreateAttributeFieldNames(t *testing.T) {
	tests := []struct {
		name           string
		baseline       BaselinePackage
		expectContains []string
		expectNotEmpty bool
	}{
		{
			name:           "2.2.0p43 has expected fields",
			baseline:       BaselineV2_2_0p43,
			expectContains: []string{"alias", "ipaddress", "site", "tag_agent"},
			expectNotEmpty: true,
		},
		{
			name:           "2.3.0p41 has expected fields",
			baseline:       BaselineV2_3_0p41,
			expectContains: []string{"alias", "ipaddress", "site", "tag_agent", "labels"},
			expectNotEmpty: true,
		},
		{
			name:           "2.4.0p17 has expected fields",
			baseline:       BaselineV2_4_0p17,
			expectContains: []string{"alias", "ipaddress", "site", "tag_agent", "labels"},
			expectNotEmpty: true,
		},
		{
			name:           "empty baseline returns nil",
			baseline:       "",
			expectNotEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HostCreateAttributeFieldNames(tt.baseline)

			if tt.expectNotEmpty && len(result) == 0 {
				t.Errorf("HostCreateAttributeFieldNames(%q) returned empty, expected values", tt.baseline)
			}

			if !tt.expectNotEmpty && len(result) != 0 {
				t.Errorf("HostCreateAttributeFieldNames(%q) = %v, expected empty", tt.baseline, result)
			}

			for _, expected := range tt.expectContains {
				if !contains(result, expected) {
					t.Errorf("HostCreateAttributeFieldNames(%q) missing %q", tt.baseline, expected)
				}
			}
		})
	}
}

func TestFolderCreateAttributeFieldNames(t *testing.T) {
	tests := []struct {
		name           string
		baseline       BaselinePackage
		expectContains []string
		expectNotEmpty bool
	}{
		{
			name:           "2.2.0p43 has expected fields",
			baseline:       BaselineV2_2_0p43,
			expectContains: []string{"tag_agent", "tag_criticality"},
			expectNotEmpty: true,
		},
		{
			name:           "2.3.0p41 has expected fields",
			baseline:       BaselineV2_3_0p41,
			expectContains: []string{"tag_agent", "tag_criticality"},
			expectNotEmpty: true,
		},
		{
			name:           "2.4.0p17 has expected fields",
			baseline:       BaselineV2_4_0p17,
			expectContains: []string{"tag_agent", "tag_criticality"},
			expectNotEmpty: true,
		},
		{
			name:           "empty baseline returns nil",
			baseline:       "",
			expectNotEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FolderCreateAttributeFieldNames(tt.baseline)

			if tt.expectNotEmpty && len(result) == 0 {
				t.Errorf("FolderCreateAttributeFieldNames(%q) returned empty, expected values", tt.baseline)
			}

			if !tt.expectNotEmpty && len(result) != 0 {
				t.Errorf("FolderCreateAttributeFieldNames(%q) = %v, expected empty", tt.baseline, result)
			}

			for _, expected := range tt.expectContains {
				if !contains(result, expected) {
					t.Errorf("FolderCreateAttributeFieldNames(%q) missing %q", tt.baseline, expected)
				}
			}
		})
	}
}

func TestAllBaselinesReturnValues(t *testing.T) {
	// Test that every defined baseline returns non-nil values
	baselines := []BaselinePackage{
		BaselineV2_2_0p1, BaselineV2_2_0p3, BaselineV2_2_0p4, BaselineV2_2_0p5,
		BaselineV2_2_0p8, BaselineV2_2_0p9, BaselineV2_2_0p11, BaselineV2_2_0p12,
		BaselineV2_2_0p14, BaselineV2_2_0p18, BaselineV2_2_0p21, BaselineV2_2_0p22,
		BaselineV2_2_0p23, BaselineV2_2_0p26, BaselineV2_2_0p32, BaselineV2_2_0p33,
		BaselineV2_2_0p43, BaselineV2_2_0p44,
		BaselineV2_3_0p1, BaselineV2_3_0p3, BaselineV2_3_0p5, BaselineV2_3_0p7,
		BaselineV2_3_0p11, BaselineV2_3_0p14, BaselineV2_3_0p22, BaselineV2_3_0p23,
		BaselineV2_3_0p26, BaselineV2_3_0p27, BaselineV2_3_0p31, BaselineV2_3_0p33,
		BaselineV2_3_0p36, BaselineV2_3_0p37, BaselineV2_3_0p39, BaselineV2_3_0p40,
		BaselineV2_3_0p41,
		BaselineV2_4_0p1, BaselineV2_4_0p6, BaselineV2_4_0p11, BaselineV2_4_0p14,
		BaselineV2_4_0p16, BaselineV2_4_0p17, BaselineV2_4_0p18,
	}

	for _, baseline := range baselines {
		t.Run(string(baseline), func(t *testing.T) {
			tagValues := ValidHostTagAgentValues(baseline)
			if len(tagValues) == 0 {
				t.Errorf("ValidHostTagAgentValues(%s) returned empty", baseline)
			}

			hostFields := HostCreateAttributeFieldNames(baseline)
			if len(hostFields) == 0 {
				t.Errorf("HostCreateAttributeFieldNames(%s) returned empty", baseline)
			}

			folderFields := FolderCreateAttributeFieldNames(baseline)
			if len(folderFields) == 0 {
				t.Errorf("FolderCreateAttributeFieldNames(%s) returned empty", baseline)
			}
		})
	}
}

// Helper function
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
