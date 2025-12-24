package gsse

// Children are either another expression or a value of core types
type (
	LightDark      struct{ Light, Dark any }
	Ident          struct{ Name any }
	Addition       struct{ Left, Right any }
	Subtraction    struct{ Left, Right any }
	Multiplication struct{ Left, Right any }
	Division       struct{ Dividend, Divisor any }
)
