package dimensional

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
		"cm>px": {New(2.54, Cm), New(96, Px)},
		"in>px": {New(1, In), New(96, Px)},
		"mm>px": {New(25.4, Mm), New(96, Px)},
		"pc>px": {New(1, Pc), New(16.0, Px)},
		"pt>px": {New(72, Pt), New(96, Px)},
		"px>px": {New(1, Px), New(1, Px)},
		"q>px":  {New(40*2.54, Q), New(96, Px)},

		"deg>deg":  {New(45, Deg), New(45.0, Deg)},
		"grad>deg": {New(200, Grad), New(180, Deg)},
		"rad>deg":  {New(2*math.Pi, Rad), New(360, Deg)},
		"turn>deg": {New(0.5, Turn), New(180.0, Deg)},

		"dppx>dppx": {New(1, Dppx), New(1, Dppx)},
		"dpcm>dppx": {New(96, Dpcm), New(2.54, Dppx)},
		"dpi>dppx":  {New(96, Dpi), New(1.0, Dppx)},

		"hz>hz":  {New(500, Hz), New(500, Hz)},
		"khz>hz": {New(2.4, KHz), New(2400, Hz)},

		"s>s":  {New(1.5, S), New(1.5, S)},
		"ms>s": {New(1500, Ms), New(1.5, S)},

		"zero value":     {New(0, In), New(0.0, Px)},
		"negative value": {New(-1, In), New(-pxPerIn, Px)},
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
		"px square>px square":   {New(1, Px, Px), New(1, Px, Px)},
		"cm sqaure>px square":   {New(2.54*2.54, Cm, Cm), New(96.0*96.0, Px, Px)},
		"deg square>deg square": {New(10, Pt, Pt), New(10*(pxPerIn/ptPerIn)*(pxPerIn/ptPerIn), Px, Px)},
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
