package gsse

import (
	"testing"

	"go.ufukty.com/gss/internal/tokens/gss"
)

func TestUnit_String(t *testing.T) {
	type tc struct {
		name     string
		input    Unit
		expected string
	}
	tcs := []tc{
		{"", Units(gss.Unit_Em), "em"},
		{"", Units(gss.Unit_Pc), "%"},
		{"", Units(gss.Unit_Em, gss.Unit_Em), "em²"},
		{"", Units(gss.Unit_Em, gss.Unit_Em, gss.Unit_Px), "em²·px"},
		{"", Units(gss.Unit_Em, gss.Unit_Em, gss.Unit_Px, gss.Unit_Px), "em²·px²"},
		{"", Units(gss.Unit_Em, gss.Unit_Pc, gss.Unit_Px, gss.Unit_Px), "%·em·px²"},
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
