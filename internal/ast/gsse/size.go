package gsse

import (
	"fmt"

	"go.ufukty.com/gss/internal/ast/html"
)

type Size struct {
	Number float64
	Unit   Unit
}

func (s Size) Compare(t Size) bool {
	return s.Number == t.Number && s.Unit.Compare(t.Unit)
}

func (s Size) String() string {
	return fmt.Sprintf("%.0f%s", s.Number, s.Unit.String())
}

func (a Size) Add(b Size) (Size, error) {
	if !a.Unit.Compare(b.Unit) {
		return Size{}, fmt.Errorf("%s + %s: %w", a, b, ErrIncompatibleUnits)
	}
	return Size{a.Number + b.Number, a.Unit}, nil
}

func (a Size) Sub(b Size) (Size, error) {
	if !a.Unit.Compare(b.Unit) {
		return Size{}, fmt.Errorf("%s - %s: %w", a, b, ErrIncompatibleUnits)
	}
	return Size{a.Number - b.Number, a.Unit}, nil
}

func (a Size) Mul(b Size) (Size, error) {
	return Size{a.Number * b.Number, a.Unit.Multiply(b.Unit)}, nil
}

func (a Size) Div(b Size) (Size, error) {
	if b.Number == 0 {
		return Size{}, fmt.Errorf("%s / %s: %w", a, b, ErrDivisionByZero)
	}
	return Size{a.Number / b.Number, a.Unit.Divide(b.Unit)}, nil
}

func (s Size) Resolve(ctx Context, e html.Element) (Size, error) {
	return s, nil
}
