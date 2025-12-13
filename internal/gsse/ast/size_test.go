package ast

import (
	"testing"

	"go.ufukty.com/gss/internal/gss/tokens"
)

func TestSize_Positive(t *testing.T) {
	px := Units(tokens.Unit_Px)
	a := Size{1, px}
	b := Size{2, px}
	c, err := a.Add(b)
	if err != nil {
		t.Errorf("act: %v", err)
	}
	if c.Number != 3 {
		t.Errorf("assert, expected 3 for value and got %.0f", c.Number)
	}
	if !c.Unit.Compare(px) {
		t.Errorf("assert, expected %q got %q", px, c.Unit)
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
