package sets

import (
	"iter"
)

// Cartesian product.
// Excludes the <nil>
func Product(sets ...[]string) iter.Seq[[]string] {
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
			for k = len(sets) - 1; k >= 0; k-- { // 099 => 100
				idx[k]++
				if idx[k] < len(sets[k]) {
					break
				}
				idx[k] = 0
			}
		}
	}
}
