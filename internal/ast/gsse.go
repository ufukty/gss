package ast

// Gsse
//
// Children either are or resolve to values of core units
type (
	LightDark      struct{ Light, Dark any }
	Ident          struct{ Name any }
	Calc           struct{ Expr any }
	Min            struct{ Exprs []any }
	Max            struct{ Exprs []any }
	Addition       struct{ Left, Right any }
	Subtraction    struct{ Left, Right any }
	Multiplication struct{ Left, Right any }
	Division       struct{ Dividend, Divisor any }
)
