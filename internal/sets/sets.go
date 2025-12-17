package sets

import (
	"iter"
	"slices"
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

// This version of Heap's algorithm chooses to stabilize
// the items on left. Meaning, items at the right are more
// volatile.
func Permutations(set []string) iter.Seq[[]string] {
	// Going right to left, swaps the first element that isn't
	// swapped as many times as its reverse index number since
	// one of its lefthand items swapped.
	// Repeation ends when each spot have as many swaps as its
	// reverse index. eg. C = [5 4 3 2 1 0] s(set)=6
	return func(yield func([]string) bool) {
		p := slices.Clone(set)
		if !yield(p) {
			return
		}
		c := make([]int, len(set))
		last := len(set) - 1
		for d := last; d >= 0; {
			if c[d] < last-d {
				if (last-d)%2 == 0 {
					p[last], p[d] = p[d], p[last]
				} else {
					p[last-c[d]], p[d] = p[d], p[last-c[d]]
				}
				if !yield(p) {
					return
				}
				c[d]++
				d = last
			} else {
				c[d] = 0
				d--
			}
		}
	}
}
