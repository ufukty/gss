package ast

import (
	"testing"

	"go.ufukty.com/gss/internal/gss/tokens"
)

func TestSize_Positive(t *testing.T) {
	px := Units(tokens.Unit_Px)
	a := Size{1, px}
	b := Size{2, px}
	expected := Size{3, px}
	got, err := a.Add(b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}

func TestSize_Negative1(t *testing.T) {
	a := Size{1, Units(tokens.Unit_Px)}
	b := Size{2, Units(tokens.Unit_Em)}
	_, err := a.Add(b)
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestSize_Negative2(t *testing.T) {
	a := Size{1, Units(tokens.Unit_Px, tokens.Unit_Px)}
	b := Size{2, Units(tokens.Unit_Px)}
	_, err := a.Add(b)
	if err == nil {
		t.Errorf("act: unexpected success")
	}
}

func TestMultiply(t *testing.T) {
	a := Size{1, Units(tokens.Unit_Px)}
	b := Size{2, Units(tokens.Unit_Em)}
	expected := Size{2, Units(tokens.Unit_Px, tokens.Unit_Em)}
	got, err := a.Mul(b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if !expected.Compare(got) {
		t.Errorf("assert, expected %q got %q", expected, got)
	}
}
