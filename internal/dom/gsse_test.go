package dom

import (
	"testing"

	"go.ufukty.com/gss/internal/tokens/gss"
)

func TestSize_Positive(t *testing.T) {
	px := Units(gss.Unit_Px)
	a := Dimension{1, px}
	b := Dimension{2, px}
	expected := Dimension{3, px}
	got, err := a.Add(b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}

func TestSize_Negative1(t *testing.T) {
	a := Dimension{1, Units(gss.Unit_Px)}
	b := Dimension{2, Units(gss.Unit_Em)}
	_, err := a.Add(b)
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestSize_Negative2(t *testing.T) {
	a := Dimension{1, Units(gss.Unit_Px, gss.Unit_Px)}
	b := Dimension{2, Units(gss.Unit_Px)}
	_, err := a.Add(b)
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestMultiply(t *testing.T) {
	var (
		a        = Dimension{1, Units(gss.Unit_Px)}
		b        = Dimension{2, Units(gss.Unit_Em)}
		expected = Dimension{2, Units(gss.Unit_Px, gss.Unit_Em)}
	)
	got, err := a.Mul(b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}

func TestDivide_StripUnit(t *testing.T) {
	var (
		a        = Dimension{10, Units(gss.Unit_Px)}
		b        = Dimension{2, Units(gss.Unit_Px)}
		expected = Dimension{5, Units()}
	)
	got, err := a.Div(b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}

func TestUnit_String(t *testing.T) {
	type tc struct {
		name     string
		input    Unit
		expected string
	}
	tcs := []tc{
		{"", Units(gss.Unit_Em), "em"},
		{"", Units(gss.Unit_Pc), "%"},
		{"", Units(gss.Unit_Em, gss.Unit_Em), "em²"},
		{"", Units(gss.Unit_Em, gss.Unit_Em, gss.Unit_Px), "em²·px"},
		{"", Units(gss.Unit_Em, gss.Unit_Em, gss.Unit_Px, gss.Unit_Px), "em²·px²"},
		{"", Units(gss.Unit_Em, gss.Unit_Pc, gss.Unit_Px, gss.Unit_Px), "%·em·px²"},
	}
	for _, tc := range tcs {
		t.Run(t.Name(), func(t *testing.T) {
			got := tc.input.String()
			if got != tc.expected {
				t.Errorf("assert, expected %q got %q", tc.expected, got)
			}
		})
	}
}
