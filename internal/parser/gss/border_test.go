package gss

import (
	"strings"
	"testing"

	"go.ufukty.com/gss/internal/sets"
)

func TestParseBorder(t *testing.T) {
	var (
		colors      = []string{"", "red", "#f00", "#F00", "#f000", "#F000", "#ff0000", "#FF0000", "#ff000000", "#FF000000"}
		styles      = []string{"", "solid", "dashed", "dotted"}
		thicknesses = []string{"", "0", "1px", "2pt", "3em", "4rem", "5vh", "6vw"}
	)

	for c1 := range sets.Combinations(colors, styles, thicknesses) {
		t.Run(strings.Join(c1, " "), func(t *testing.T) {

		})
	}
}
