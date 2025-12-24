package dimensional

import (
	"fmt"
	"math"
	"strconv"
)

var (
	ErrUnknownUnit                   = fmt.Errorf("unknown unit")
	ErrCannonicalizingContextualUnit = fmt.Errorf("cannonicalizing contextual unit")
	ErrDivisionByZero                = fmt.Errorf("division by zero is undefined")
	ErrIncompatibleUnits             = fmt.Errorf("operands have incompetable units")
)

// GSSE compliant primitive value types
//
// Values of those types can be assigned to a variable of Expr[T] as leaf node
type (
	Dimension struct {
		Value float64
		Unit  Compound
	}
)

func (d Dimension) Compare(t Dimension) bool {
	return d.Value == t.Value && d.Unit.Compare(t.Unit)
}

func round(f float64, preserveDecimals int) float64 {
	p := math.Pow10(preserveDecimals)
	return math.Round(f*p) / p
}

// returns the number of decimals just enough to represent
// the whole precision or the most within the cap.
func minDecimals(f float64, cap int) (int, bool) {
	lv, li := math.Inf(+1), -1
	for i := 0; i <= cap; i++ {
		cur := round(f, i)
		if cur == f {
			return i, false
		}
		if cur != lv {
			lv, li = cur, i
		}
	}
	return li, true
}

func (d Dimension) unitless() string {
	md, tilde := minDecimals(d.Value, 2)
	if tilde {
		return "~" + strconv.FormatFloat(d.Value, 'f', md, 64)
	}
	return strconv.FormatFloat(d.Value, 'f', md, 64)
}

func (d Dimension) String() string {
	return d.unitless() + d.Unit.String()
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

func New(v float64, units ...Unit) Dimension {
	return Dimension{
		Value: v,
		Unit:  parse(units...),
	}
}
