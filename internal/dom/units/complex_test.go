package units

import (
	"testing"
)

func TestUnit_String(t *testing.T) {
	type tc struct {
		name     string
		input    Complex
		expected string
	}
	tcs := []tc{
		{"", Parse(Em), "em"},
		{"", Parse(Pc), "%"},
		{"", Parse(Em, Em), "em²"},
		{"", Parse(Em, Em, Px), "em²·px"},
		{"", Parse(Em, Em, Px, Px), "em²·px²"},
		{"", Parse(Em, Pc, Px, Px), "%·em·px²"},
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
