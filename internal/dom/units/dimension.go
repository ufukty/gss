package units

import (
	"fmt"
)

var (
	ErrDivisionByZero    = fmt.Errorf("division by zero is undefined")
	ErrIncompatibleUnits = fmt.Errorf("operands have incompetable units")
)

// GSSE compliant primitive value types
//
// Values of those types can be assigned to a variable of Expr[T] as leaf node
type (
	Dimension struct {
		Value float64
		Unit  Complex
	}
)

func (d Dimension) Compare(t Dimension) bool {
	return d.Value == t.Value && d.Unit.Compare(t.Unit)
}

func (d Dimension) String() string {
	return fmt.Sprintf("%.0f%s", d.Value, d.Unit.String())
}

func Add(a, b Dimension) (Dimension, error) {
	if !a.Unit.Compare(b.Unit) {
		return a, ErrIncompatibleUnits
	}
	c := Dimension{
		Value: a.Value + b.Value,
		Unit:  a.Unit,
	}
	return c, nil
}

func Subtract(a, b Dimension) (Dimension, error) {
	if !a.Unit.Compare(b.Unit) {
		return a, ErrIncompatibleUnits
	}
	c := Dimension{
		Value: a.Value - b.Value,
		Unit:  a.Unit,
	}
	return c, nil
}

func Multiply(a, b Dimension) (Dimension, error) {
	c := Dimension{
		Value: a.Value * b.Value,
		Unit:  a.Unit.Multiply(b.Unit),
	}
	return c, nil
}

func Divide(a, b Dimension) (Dimension, error) {
	if b.Value == 0 {
		return a, ErrDivisionByZero
	}
	c := Dimension{
		Value: a.Value / b.Value,
		Unit:  a.Unit.Divide(b.Unit),
	}
	return c, nil
}

func NewDimensional(v float64, units ...Unit) Dimension {
	return Dimension{
		Value: v,
		Unit:  Parse(units...),
	}
}
