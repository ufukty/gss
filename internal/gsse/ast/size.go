package ast

import (
	"fmt"

	"go.ufukty.com/gss/internal/html/ast"
)

type Size struct {
	Number float64
	Unit   Unit
}

func (s Size) String() string {
	return fmt.Sprintf("%.0f%s", s.Number, s.Unit.String())
}

func (a Size) Add(b Size) (Size, error) {
	if !a.Unit.Compare(b.Unit) {
		return Size{}, fmt.Errorf("operands %s and %s have different units", a, b)
	}
	return Size{a.Number + b.Number, a.Unit}, nil
}

func (a Size) Sub(b Size) (Size, error) {
	if !a.Unit.Compare(b.Unit) {
		return Size{}, fmt.Errorf("operands %s and %s have different units", a, b)
	}
	return Size{a.Number - b.Number, a.Unit}, nil
}

func (a Size) Mul(b Size) (Size, error) {
	return Size{a.Number * b.Number, a.Unit.Multiply(b.Unit)}, nil
}

func (a Size) Div(b Size) (Size, error) {
	return Size{a.Number / b.Number, a.Unit.Divide(b.Unit)}, nil
}

func (s Size) Resolve(ctx Context, e ast.Element) (Size, error) {
	return s, nil
}
