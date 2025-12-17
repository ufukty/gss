package sets

import (
	"fmt"
	"strings"
	"testing"
)

func ExamplePermutations() {
	for p := range Permutations([]string{"a", "b", "c"}) {
		fmt.Println(p)
	}
	// Output:
	// [a b c]
	// [b a c]
	// [c a b]
	// [a c b]
	// [b c a]
	// [c b a]
}

func TestPermutations_Uniqueness(t *testing.T) {
	ps := map[string]int{}
	for p := range Permutations([]string{"a", "b", "c", "d"}) {
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
