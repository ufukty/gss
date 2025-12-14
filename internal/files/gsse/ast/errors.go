package ast

import "fmt"

var (
	ErrDivisionByZero    = fmt.Errorf("division by zero is undefined")
	ErrIncompatibleUnits = fmt.Errorf("operands have incompetable units")
)
