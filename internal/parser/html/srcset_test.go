package html

import "testing"

func compare[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}

func TestParseSrcSet(t *testing.T) {
	type tc struct {
		name     string
		input    string
		expected map[float64]string
	}
	tcs := []tc{
		{"", "a.png 2x", map[float64]string{2: "a.png"}},
		{"", "a.png 2x, b.png 3x", map[float64]string{2: "a.png", 3: "b.png"}},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseSrcSet(tc.input)
			if err != nil {
				t.Errorf("act, unexpected error: %v", err)
			}
			if !compare(tc.expected, got) {
				t.Errorf("assert, expected %v got %v", tc.expected, got)
			}
		})
	}
}
