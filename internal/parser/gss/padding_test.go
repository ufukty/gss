package gss

import (
	"fmt"
	"testing"

	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/dimensional"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

func TestParsePadding_positionalShorthands(t *testing.T) {
	var (
		p1 = dimensional.New(10, dimensional.Px)
		p2 = dimensional.New(20, dimensional.Px)
		p3 = dimensional.New(30, dimensional.Px)
		p4 = dimensional.New(40, dimensional.Px)
	)
	type tc struct {
		name     string
		input    string
		expected ast.Padding
	}
	tcs := []tc{
		{
			name:     "TRBL (all sides)",
			input:    "10px",
			expected: ast.Padding{Top: p1, Right: p1, Bottom: p1, Left: p1},
		},
		{
			name:     "TB|LR (vertical/horizontal)",
			input:    "10px 20px",
			expected: ast.Padding{Top: p1, Right: p2, Bottom: p1, Left: p2},
		},
		{
			name:     "T|LR|B (top/horizontal/bottom)",
			input:    "10px 20px 30px",
			expected: ast.Padding{Top: p1, Right: p2, Bottom: p3, Left: p2},
		},
		{
			name:     "T|R|B|L (all individual)",
			input:    "10px 20px 30px 40px",
			expected: ast.Padding{Top: p1, Right: p2, Bottom: p3, Left: p4},
		},
		{
			name:     "zero value",
			input:    "0",
			expected: ast.Padding{Top: dimensional.New(0), Right: dimensional.New(0), Bottom: dimensional.New(0), Left: dimensional.New(0)},
		},
		{
			name:     "mixed units",
			input:    "1px 2em 3rem 4pt",
			expected: ast.Padding{Top: dimensional.New(1, dimensional.Px), Right: dimensional.New(2, dimensional.Em), Bottom: dimensional.New(3, dimensional.Rem), Left: dimensional.New(4, dimensional.Pt)},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := csstokens.Tokenize(fmt.Sprintf("padding: %s", tc.input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			got, err := ParsePadding(ts)
			if err != nil {
				t.Errorf("act, unexpected error: %v", err)
			}
			if !tc.expected.IsEqual(got) {
				t.Errorf("assert, expected %#v, got %#v", tc.expected, got)
			}
		})
	}
}

func TestParsePadding_errors(t *testing.T) {
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
			input: "10px invalid",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := csstokens.Tokenize(fmt.Sprintf("padding: %s", tc.input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			_, err = ParsePadding(ts)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
