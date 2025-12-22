package units

import (
	"testing"

	"go.ufukty.com/gss/internal/tokens/gss"
)

func TestUnit_String(t *testing.T) {
	type tc struct {
		name     string
		input    Complex
		expected string
	}
	tcs := []tc{
		{"", Parse(gss.Unit_Em), "em"},
		{"", Parse(gss.Unit_Pc), "%"},
		{"", Parse(gss.Unit_Em, gss.Unit_Em), "em²"},
		{"", Parse(gss.Unit_Em, gss.Unit_Em, gss.Unit_Px), "em²·px"},
		{"", Parse(gss.Unit_Em, gss.Unit_Em, gss.Unit_Px, gss.Unit_Px), "em²·px²"},
		{"", Parse(gss.Unit_Em, gss.Unit_Pc, gss.Unit_Px, gss.Unit_Px), "%·em·px²"},
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
