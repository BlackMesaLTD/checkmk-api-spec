package main

import "testing"

func TestRepoForVersion(t *testing.T) {
	cases := map[string]string{
		"2.2.0p1":  rawImage,
		"2.3.0p41": rawImage,
		"2.4.0p32": rawImage,
		"2.5.0p1":  communityImage,
		"2.5.0p6":  communityImage,
		"2.6.0p1":  communityImage,
		"3.0.0p1":  communityImage,
	}
	for version, want := range cases {
		if got := repoForVersion(version); got != want {
			t.Errorf("repoForVersion(%q) = %q, want %q", version, got, want)
		}
	}
}

func TestGetTypeArrayForm(t *testing.T) {
	// Scalar (3.0) form.
	if got := getType(map[string]interface{}{"type": "string"}); got != "string" {
		t.Errorf("scalar type = %q, want \"string\"", got)
	}
	// Array (3.1) nullable form.
	prop := map[string]interface{}{"type": []interface{}{"string", "null"}}
	if got := getType(prop); got != "null|string" {
		t.Errorf("array type = %q, want \"null|string\"", got)
	}
}
