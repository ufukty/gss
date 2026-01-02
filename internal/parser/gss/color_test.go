package gss

import (
	"image/color"
	"testing"
)

func Test_ParseColor(t *testing.T) {
	type tc struct {
		name     string
		input    string
		expected color.NRGBA
	}
	red := color.NRGBA{255, 0, 0, 255}
	tcs := []tc{
		{"rgb", "#f00", red},
		{"RGB", "#F00", red},
		{"rgba", "#f00f", red},
		{"RGBA", "#F00F", red},
		{"rrggbb", "#ff0000", red},
		{"RRGGBB", "#FF0000", red},
		{"rrggbbaa", "#ff0000ff", red},
		{"RRGGBBAA", "#FF0000FF", red},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseColor(tc.input)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if tc.expected != got {
				t.Errorf("assert, expected %v, got %v", tc.expected, got)
			}
		})
	}
}
