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

// Heap's Algorithm (Iterative)
//
// Going left to right, swaps the first element that isn't
// swapped as many times as its index number since one of
// its righthand items gets swapped.
// Repeation ends when each spot have as many swaps as its
// index.
//
// Chart for count array:
//
//	0 0 0 0 0
//	0 1 0 0 0 (after 0th and 1st swapped)
//	0 0 1 0 0 (after 0th and 2nd swapped. 1st set 0 as passed)
//	0 1 1 0 0 (after 0th and 1st swapped. 2nd is untouched)
//	0 0 2 0 0 ...
//	...
//	...
//	...
//	0 1 2 3 4
func Permutations(set []string) iter.Seq[[]string] {
	// It is normal the [i] moves in both directions. It gets
	// resetted after each swap.
	//
	// The [c] values count how many times the [i]th element
	// swapped with one of its lefthand items since the last
	// time one of its righthand items gets swapped.
	return func(yield func([]string) bool) {
		p := slices.Clone(set)
		if !yield(p) {
			return
		}
		c := make([]int, len(set))
		for i := 0; i < len(set); {
			if c[i] < i {
				if i%2 == 0 {
					p[0], p[i] = p[i], p[0]
				} else {
					p[c[i]], p[i] = p[i], p[c[i]]
				}
				if !yield(p) {
					return
				}
				c[i]++
				i = 0
			} else {
				c[i] = 0
				i++
			}
		}
	}
}
