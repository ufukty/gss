package ast

import (
	"testing"

	"go.ufukty.com/gss/internal/gss/tokens"
)

func TestSize(t *testing.T) {
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
