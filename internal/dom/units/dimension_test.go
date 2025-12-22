package units

import "testing"

func TestAdd_Positive(t *testing.T) {
	got, err := Add(NewDimensional(1, Px), NewDimensional(2, Px))
	if err != nil {
		t.Fatalf("act: %v", err)
	}
	expected := NewDimensional(3, Px)
	if !expected.Compare(got) {
		t.Errorf("assert, expected %s got %s", expected, got)
	}
}

func TestAdd_Negative1(t *testing.T) {
	_, err := Add(NewDimensional(1, Px), NewDimensional(2, Em))
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestAdd_Negative2(t *testing.T) {
	_, err := Add(NewDimensional(1, Px), NewDimensional(2, Px, Px))
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestMultiply(t *testing.T) {
	var (
		a        = NewDimensional(1, Px)
		b        = NewDimensional(2, Em)
		expected = NewDimensional(2, Px, Em)
	)
	got, err := Multiply(a, b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}

func TestDivide_StripUnit(t *testing.T) {
	var (
		a        = NewDimensional(10, Px)
		b        = NewDimensional(2, Px)
		expected = NewDimensional(5)
	)
	got, err := Divide(a, b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}
