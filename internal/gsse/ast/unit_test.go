package ast

import (
	"testing"

	"go.ufukty.com/gss/internal/gss/tokens"
)

func TestUnit_String(t *testing.T) {
	type tc struct {
		name     string
		input    Unit
		expected string
	}
	tcs := []tc{
		{"", Units(tokens.Unit_Em), "em"},
		{"", Units(tokens.Unit_Pc), "%"},
		{"", Units(tokens.Unit_Em, tokens.Unit_Em), "em²"},
		{"", Units(tokens.Unit_Em, tokens.Unit_Em, tokens.Unit_Px), "em²·px"},
		{"", Units(tokens.Unit_Em, tokens.Unit_Em, tokens.Unit_Px, tokens.Unit_Px), "em²·px²"},
		{"", Units(tokens.Unit_Em, tokens.Unit_Pc, tokens.Unit_Px, tokens.Unit_Px), "%·em·px²"},
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
