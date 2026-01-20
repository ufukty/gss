package gss

import (
	"fmt"
	"testing"

	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/dimensional"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

func TestParseMargin_positionalShorthands(t *testing.T) {
	var (
		m1 = dimensional.New(10, dimensional.Px)
		m2 = dimensional.New(20, dimensional.Px)
		m3 = dimensional.New(30, dimensional.Px)
		m4 = dimensional.New(40, dimensional.Px)
	)
	type tc struct {
		name     string
		input    string
		expected ast.Margin
	}
	tcs := []tc{
		{
			name:     "TRBL (all sides)",
			input:    "10px",
			expected: ast.Margin{Top: m1, Right: m1, Bottom: m1, Left: m1},
		},
		{
			name:     "TB|LR (vertical/horizontal)",
			input:    "10px 20px",
			expected: ast.Margin{Top: m1, Right: m2, Bottom: m1, Left: m2},
		},
		{
			name:     "T|LR|B (top/horizontal/bottom)",
			input:    "10px 20px 30px",
			expected: ast.Margin{Top: m1, Right: m2, Bottom: m3, Left: m2},
		},
		{
			name:     "T|R|B|L (all individual)",
			input:    "10px 20px 30px 40px",
			expected: ast.Margin{Top: m1, Right: m2, Bottom: m3, Left: m4},
		},
		{
			name:     "zero value",
			input:    "0",
			expected: ast.Margin{Top: dimensional.New(0), Right: dimensional.New(0), Bottom: dimensional.New(0), Left: dimensional.New(0)},
		},
		{
			name:     "mixed units",
			input:    "1vh 2vw 3em 4rem",
			expected: ast.Margin{Top: dimensional.New(1, dimensional.Vh), Right: dimensional.New(2, dimensional.Vw), Bottom: dimensional.New(3, dimensional.Em), Left: dimensional.New(4, dimensional.Rem)},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := csstokens.Tokenize(fmt.Sprintf("margin: %s", tc.input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			got, err := ParseMargin(ts)
			if err != nil {
				t.Errorf("act, unexpected error: %v", err)
			}
			if !tc.expected.IsEqual(got) {
				t.Errorf("assert, expected %#v, got %#v", tc.expected, got)
			}
		})
	}
}

func TestParseMargin_errors(t *testing.T) {
	type tc struct {
		name  string
		input string
	}
	tcs := []tc{
		{
			name:  "too many values",
			input: "10px 20px 30px 40px 50px",
		},
		{
			name:  "empty",
			input: "",
		},
		{
			name:  "invalid value",
			input: "10px notavalue",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := csstokens.Tokenize(fmt.Sprintf("margin: %s", tc.input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			_, err = ParseMargin(ts)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
