package gsse

type (
	LightDark[Final any] struct {
		Light, Dark string
	}
	Ident[Final any] struct {
		Name string
	}
	Addition struct {
		Lhs, Rhs string
	}
	Subtraction struct {
		Lhs, Rhs string
	}
	Multiplication struct {
		Lhs, Rhs string
	}
	Division struct {
		Dividend, Divisor string
	}
)
