package ast

type (
	LightDark[Final any] struct {
		Light, Dark Expr[Final]
	}
	Ident[Final any] struct {
		Name string
	}
	Addition struct {
		Lhs, Rhs Expr[Size]
	}
	Subtraction struct {
		Lhs, Rhs Expr[Size]
	}
	Multiplication struct {
		Lhs, Rhs Expr[Size]
	}
	Division struct {
		Dividend, Divisor Expr[Size]
	}
)
