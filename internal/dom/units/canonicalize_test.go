package units

import (
	"math"
	"testing"
)

func TestCanonicalize_Basic(t *testing.T) {
	type tc struct {
		input    Dimension
		expected Dimension
	}
	tcs := map[string]tc{
		"cm>px": {NewDimensional(2.54, Cm), NewDimensional(96, Px)},
		"in>px": {NewDimensional(1, In), NewDimensional(96, Px)},
		"mm>px": {NewDimensional(25.4, Mm), NewDimensional(96, Px)},
		"pc>px": {NewDimensional(1, Pc), NewDimensional(16.0, Px)},
		"pt>px": {NewDimensional(72, Pt), NewDimensional(96, Px)},
		"px>px": {NewDimensional(1, Px), NewDimensional(1, Px)},
		"q>px":  {NewDimensional(40*2.54, Q), NewDimensional(96, Px)},

		"deg>deg":  {NewDimensional(45, Deg), NewDimensional(45.0, Deg)},
		"grad>deg": {NewDimensional(200, Grad), NewDimensional(180, Deg)},
		"rad>deg":  {NewDimensional(2*math.Pi, Rad), NewDimensional(360, Deg)},
		"turn>deg": {NewDimensional(0.5, Turn), NewDimensional(180.0, Deg)},

		"dppx>dppx": {NewDimensional(1, Dppx), NewDimensional(1, Dppx)},
		"dpcm>dppx": {NewDimensional(96, Dpcm), NewDimensional(2.54, Dppx)},
		"dpi>dppx":  {NewDimensional(96, Dpi), NewDimensional(1.0, Dppx)},

		"hz>hz":  {NewDimensional(500, Hz), NewDimensional(500, Hz)},
		"khz>hz": {NewDimensional(2.4, KHz), NewDimensional(2400, Hz)},

		"s>s":  {NewDimensional(1.5, S), NewDimensional(1.5, S)},
		"ms>s": {NewDimensional(1500, Ms), NewDimensional(1.5, S)},

		"zero value":     {NewDimensional(0, In), NewDimensional(0.0, Px)},
		"negative value": {NewDimensional(-1, In), NewDimensional(-PxPerIn, Px)},
	}

	for tn, tc := range tcs {
		t.Run(tn, func(t *testing.T) {
			got, err := Canonicalize(tc.input)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if !tc.expected.Compare(got) {
				t.Errorf("assert, expected %s, got %s", tc.expected, got)
			}
		})
	}
}

func TestCanonicalize_Complex(t *testing.T) {
	type tc struct {
		input    Dimension
		expected Dimension
	}

	tcs := map[string]tc{
		"px square>px square":   {NewDimensional(1, Px, Px), NewDimensional(1, Px, Px)},
		"cm sqaure>px square":   {NewDimensional(2.54*2.54, Cm, Cm), NewDimensional(96.0*96.0, Px, Px)},
		"deg square>deg square": {NewDimensional(10, Pt, Pt), NewDimensional(10*(PxPerIn/PtPerIn)*(PxPerIn/PtPerIn), Px, Px)},
	}

	for tn, tc := range tcs {
		t.Run(tn, func(t *testing.T) {
			got, err := Canonicalize(tc.input)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if !tc.expected.Compare(got) {
				t.Errorf("assert, expected %s, got %s", tc.expected, got)
			}
		})
	}
}
