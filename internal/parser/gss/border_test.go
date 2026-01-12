package gss

import (
	"fmt"
	"image/color"
	"strings"
	"testing"

	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/dimensional"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
	"go.ufukty.com/gss/internal/sets"
	"go.ufukty.com/gss/internal/tokens"
)

func TestParseBorder_combinations(t *testing.T) {
	var (
		colors = []string{"", "red", "#f00", "#F00", "#f00f", "#F00F", "#ff0000", "#FF0000", "#ff0000ff", "#FF0000FF"}
		styles = []string{"", "solid", "dashed", "dotted"}
		widths = []string{"", "0", "1px", "2pt", "3em", "4rem", "5vh", "6vw"}
	)

	for combination := range sets.Product(colors, styles, widths) {
		input := strings.Join(combination, " ")
		t.Run(input, func(t *testing.T) {
			ts, err := csstokens.Tokenize(fmt.Sprintf("border: %s", input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			ParseBorder(ts)
		})
	}
}

func TestParseBorders_positionalShorthands(t *testing.T) {
	var (
		black = color.NRGBA{0, 0, 0, 255}
		solid = tokens.BorderStyleSolid
		b1    = ast.Border{Color: black, Style: solid, Width: dimensional.New(1, dimensional.Px)}
		b2    = ast.Border{Color: black, Style: solid, Width: dimensional.New(2, dimensional.Px)}
		b3    = ast.Border{Color: black, Style: solid, Width: dimensional.New(3, dimensional.Px)}
		b4    = ast.Border{Color: black, Style: solid, Width: dimensional.New(4, dimensional.Px)}
	)
	type tc struct {
		name     string
		input    string
		expected ast.Borders
	}
	tcs := []tc{
		{
			name:     "T|LR|B",
			input:    "1px solid #000, 2px solid #000, 3px solid #000",
			expected: ast.Borders{Top: b1, Right: b2, Bottom: b3, Left: b2},
		},
		{
			name:     "T|R|B|L",
			input:    "1px solid #000, 2px solid #000, 3px solid #000, 4px solid #000",
			expected: ast.Borders{Top: b1, Right: b2, Bottom: b3, Left: b4},
		},
		{
			name:     "TB|LR",
			input:    "1px solid #000, 2px solid #000",
			expected: ast.Borders{Top: b1, Right: b2, Bottom: b1, Left: b2},
		},
		{
			name:     "TRBL",
			input:    "1px solid #000",
			expected: ast.Borders{Top: b1, Right: b1, Bottom: b1, Left: b1},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := csstokens.Tokenize(fmt.Sprintf("border: %s", tc.input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			got, err := ParseBorders(ts)
			if err != nil {
				t.Errorf("act, unexpected error: %v", err)
			}
			if tc.expected != *got {
				t.Errorf("assert, expected %v got %v", tc.expected, got)
			}
		})
	}
}
