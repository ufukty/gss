package gsse

type Expr interface {
	expr()
}

type (
	LightDark      struct{ Light, Dark Expr }
	Ident          struct{ Name Expr }
	Addition       struct{ Left, Right Expr }
	Subtraction    struct{ Left, Right Expr }
	Multiplication struct{ Left, Right Expr }
	Division       struct{ Dividend, Divisor Expr }
)

func (LightDark) expr()      {}
func (Ident) expr()          {}
func (Addition) expr()       {}
func (Subtraction) expr()    {}
func (Multiplication) expr() {}
func (Division) expr()       {}
