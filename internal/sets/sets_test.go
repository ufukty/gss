package sets

import (
	"fmt"
	"strings"
	"testing"
)

func ExamplePermutations_l3() {
	for p := range Permutations([]string{"a", "b", "c"}) {
		fmt.Println(p)
	}
	// Output:
	// [a b c]
	// [a c b]
	// [b c a]
	// [b a c]
	// [c a b]
	// [c b a]
}

func ExamplePermutations_l4() {
	for p := range Permutations([]string{"a", "b", "c", "d"}) {
		fmt.Println(p)
	}
	// Output:
	// [a b c d]
	// [a b d c]
	// [a c d b]
	// [a c b d]
	// [a d b c]
	// [a d c b]
	// [b d c a]
	// [b d a c]
	// [b c a d]
	// [b c d a]
	// [b a d c]
	// [b a c d]
	// [c a b d]
	// [c a d b]
	// [c b d a]
	// [c b a d]
	// [c d a b]
	// [c d b a]
	// [d c b a]
	// [d c a b]
	// [d b a c]
	// [d b c a]
	// [d a c b]
	// [d a b c]
}

func TestPermutations_Uniqueness(t *testing.T) {
	ps := map[string]int{}
	for p := range Permutations([]string{"a", "b", "c", "d", "e", "h"}) {
		k := strings.Join(p, "")
		if _, ok := ps[k]; !ok {
			ps[k] = 0
		}
		ps[k]++
	}

	for p, c := range ps {
		if c > 1 {
			t.Errorf("unexpected %d duplicates of %s", c, p)
		}
	}
}
