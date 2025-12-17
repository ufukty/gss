package gss

import (
	"iter"
	"strings"
	"testing"
)

func combinations(sets ...[]string) iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		idx := make([]int, len(sets))
		cur := make([]string, len(sets))
		for k := 1; k >= 0; {
			for i := range sets {
				cur[i] = sets[i][idx[i]]
			}
			if !yield(cur) {
				return
			}
			for k = len(sets) - 1; k >= 0; k-- { // 000n -> 0010
				idx[k]++
				if idx[k] < len(sets[k]) {
					break
				}
				idx[k] = 0
			}
		}
	}
}

func TestParseBorder(t *testing.T) {
	var (
		colors      = []string{"", "red", "#f00", "#F00", "#f000", "#F000", "#ff0000", "#FF0000", "#ff000000", "#FF000000"}
		styles      = []string{"", "solid", "dashed", "dotted"}
		thicknesses = []string{"", "0", "1px", "2pt", "3em", "4rem", "5vh", "6vw"}
	)

	for c1 := range combinations(colors, styles, thicknesses) {
		t.Run(strings.Join(c1, " "), func(t *testing.T) {

		})
	}
}
