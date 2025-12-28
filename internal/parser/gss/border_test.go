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

func TestParseBorder(t *testing.T) {
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
			ParseBorders(ts)
		})
	}
}
