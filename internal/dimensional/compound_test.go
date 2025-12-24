package dimensional

import (
	"testing"
)

func TestCompound_String(t *testing.T) {
	type tc struct {
		name     string
		input    Compound
		expected string
	}
	tcs := []tc{
		{"", parse(Em), "em"},
		{"", parse(Pc), "pc"},
		{"", parse(Em, Em), "em²"},
		{"", parse(Em, Em, Px), "em²·px"},
		{"", parse(Em, Em, Px, Px), "em²·px²"},
		{"", parse(Em, Pc, Px, Px), "em·pc·px²"},
	}
	for _, tc := range tcs {
		t.Run(t.Name(), func(t *testing.T) {
			got := tc.input.String()
			if got != tc.expected {
				t.Errorf("assert, expected %q got %q", tc.expected, got)
			}
		})
	}
}
