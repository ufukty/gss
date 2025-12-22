package units

import (
	"testing"
)

func TestSize_Positive(t *testing.T) {
	px := Parse(Px)
	got, err := Add(Dimension{1, px}, Dimension{2, px})
	if err != nil {
		t.Fatalf("act: %v", err)
	}
	expected := 3.0
	if expected != got {
		t.Errorf("assert, expected %f got %f", expected, got)
	}
}

func TestSize_Negative1(t *testing.T) {
	input := Addition[float64]{
		Dimension{1, Parse(Px)},
		Dimension{2, Parse(Em)},
	}
	_, err := input.Resolve(Context{}, nil)
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestSize_Negative2(t *testing.T) {
	input := Addition[float64]{
		Dimension{1, Parse(Px, Px)},
		Dimension{2, Parse(Em)},
	}
	_, err := input.Resolve(Context{}, nil)
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestMultiply(t *testing.T) {
	var (
		a        = Dimension{1, Parse(gss.Unit_Px)}
		b        = Dimension{2, Parse(gss.Unit_Em)}
		expected = Dimension{2, Parse(Px, Em)}
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
		a        = Dimension{10, Parse(Px)}
		b        = Dimension{2, Parse(Px)}
		expected = Dimension{5, Parse()}
	)
	got, err := a.Div(b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}
