package gss

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/sets"
)

func tokenize(in string) ([]css.Token, error) {
	p := css.NewParser(parse.NewInputString(in), true)
	gt, _, _ := p.Next()
	if gt == css.ErrorGrammar {
		if err := p.Err(); err != io.EOF {
			return nil, err
		}
	}
	return p.Values(), nil
}

func TestParseBorder_combinations(t *testing.T) {
	var (
		colors = []string{"", "red", "#f00", "#F00", "#f000", "#F000", "#ff0000", "#FF0000", "#ff000000", "#FF000000"}
		styles = []string{"", "solid", "dashed", "dotted"}
		widths = []string{"", "0", "1px", "2pt", "3em", "4rem", "5vh", "6vw"}
	)

	for combination := range sets.Product(colors, styles, widths) {
		input := strings.Join(combination, " ")
		t.Run(input, func(t *testing.T) {
			ts, err := tokenize(fmt.Sprintf("border: %s", input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			ParseBorder(ts)
		})
	}
}

func TestParseBorders_positionalShorthands(t *testing.T) {
	tcs := map[string]string{
		"T|LR|B":  "1px solid #000, 2px solid #000, 3px solid #000",
		"T|R|B|L": "1px solid #000, 2px solid #000, 3px solid #000, 4px solid #000",
		"TB|LR":   "1px solid #000, 2px solid #000",
		"TRBL":    "1px solid #000",
	}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			ts, err := tokenize(fmt.Sprintf("border: %s", input))
			if err != nil {
				t.Fatalf("prep, unexpected error: %v", err)
			}
			ParseBorders(ts)
		})
	}
}
