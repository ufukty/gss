package dom

import (
	"fmt"
	"image/color"

	"go.ufukty.com/gss/internal/dom/units"
	"go.ufukty.com/gss/internal/tokens/gss"
)

var (
	ErrDivisionByZero    = fmt.Errorf("division by zero is undefined")
	ErrIncompatibleUnits = fmt.Errorf("operands have incompetable units")
)

// Context
type (
	Media struct {
		PrefersColorScheme string
	}
	PixelArea struct {
		Width, Height float64
	}
	Context struct {
		Media     Media
		Container PixelArea
		Viewport  PixelArea
	}
)

// GSSE compliant primitive value types
//
// Values of those types can be assigned to a variable of Expr[T] as leaf node
type (
	Dimension struct {
		Number float64
		Unit   units.Unit
	}
)

func (d Dimension) Color(ctx Context, e Element) (float64, error) {
	if !d.Unit.Compare(units.Parse("px")) {
		panic("implement conversion using context")
	}
	return d.Number, nil
}

func (d Dimension) Compare(t Dimension) bool {
	return d.Number == t.Number && d.Unit.Compare(t.Unit)
}

func (d Dimension) String() string {
	return fmt.Sprintf("%.0f%s", d.Number, d.Unit.String())
}

func (d Dimension) Resolve(ctx Context, e Element) (float64, error) {
	return d.Number, nil
}

// Value types are to be used in instantiating [Expr] types.
type (
	Color        color.RGBA // eg. #FF0000
	Pixels       float64    // eg. 10px
	Angle        float64    // eg. 360deg
	Milliseconds int        // eg. 1000ms
	Image        string     // eg. url()
	FontSet      []string   // eg. "Helvetica", "Helvetica Neue", sans-serif
)

type Expr[Resolving any] interface {
	Resolve(ctx Context, e Element) (Resolving, error)
}

// GSSE functions
type (
	LightDark struct {
		Light, Dark Expr[Color]
	}
	Ident[Resolving any] struct {
		Name string
	}
	Addition[Resolving ~float64] struct {
		Lhs, Rhs any
	}
	Subtraction[Resolving ~float64] struct {
		Lhs, Rhs any
	}
	Multiplication[Resolving ~float64] struct {
		Lhs, Rhs any
	}
	Division[Resolving ~float64] struct {
		Dividend, Divisor any
	}
)

var (
	_ Expr[Pixels] = (*Addition[Pixels])(nil)
	_ Expr[Pixels] = (*Subtraction[Pixels])(nil)
	_ Expr[Pixels] = (*Multiplication[Pixels])(nil)
	_ Expr[Pixels] = (*Division[Pixels])(nil)
)

func subtree[T any](a any, ctx Context, e Element) (T, error) {
	switch a := a.(type) {
	case Expr[T]:
		return a.Resolve(ctx, e)
	case T:
		return a, nil
	default:
		var t T
		return t, fmt.Errorf("value of unknown type: %T", a)
	}
}

func (c LightDark) Resolve(ctx Context, e Element) (Color, error) {
	if ctx.Media.PrefersColorScheme == "dark" {
		return subtree[Color](c.Dark, ctx, e)
	}
	return subtree[Color](c.Light, ctx, e)
}

// FIXME: Fetch identity value from DOM not AST once it is available
func (i Ident[Resolving]) Resolve(ctx Context, e Element) (Resolving, error)

func (a Addition[F]) Resolve(ctx Context, e Element) (F, error) {
	l, err := subtree[F](a.Lhs, ctx, e)
	if err != nil {
		return l, fmt.Errorf("left operand: %w", err)
	}
	r, err := subtree[F](a.Rhs, ctx, e)
	if err != nil {
		return r, fmt.Errorf("right operand: %w", err)
	}
	return l + r, nil
}

func (a Subtraction[F]) Resolve(ctx Context, e Element) (F, error) {
	l, err := subtree[F](a.Lhs, ctx, e)
	if err != nil {
		return l, fmt.Errorf("left operand: %w", err)
	}
	r, err := subtree[F](a.Rhs, ctx, e)
	if err != nil {
		return r, fmt.Errorf("right operand: %w", err)
	}
	return l - r, nil
}

func (a Multiplication[F]) Resolve(ctx Context, e Element) (F, error) {
	l, err := subtree[F](a.Lhs, ctx, e)
	if err != nil {
		return l, fmt.Errorf("left operand: %w", err)
	}
	r, err := subtree[F](a.Rhs, ctx, e)
	if err != nil {
		return r, fmt.Errorf("right operand: %w", err)
	}
	return l * r, nil
}

func (a Division[F]) Resolve(ctx Context, e Element) (F, error) {
	l, err := subtree[F](a.Dividend, ctx, e)
	if err != nil {
		return l, fmt.Errorf("left operand: %w", err)
	}
	r, err := subtree[F](a.Divisor, ctx, e)
	if err != nil {
		return r, fmt.Errorf("right operand: %w", err)
	}
	if r == 0 {
		return r, ErrDivisionByZero
	}
	return l / r, nil
}

// Nodes
type (
	Display struct {
		Outside gss.DisplayOutside
		Inside  gss.DisplayInside
	}

	Border struct {
		Color     Expr[Color]
		Style     string
		Thickness Expr[Pixels]
	}

	BorderRadiuses struct {
		TopLeft, TopRight, BottomRight, BottomLeft Expr[Pixels]
	}

	Borders struct {
		Top, Right, Bottom, Left Border
	}

	Margin struct {
		Top, Right, Bottom, Left Expr[Color]
	}

	Padding struct {
		Top, Right, Bottom, Left Expr[Color]
	}

	Font struct {
		Family    []gss.FontFamily
		Dimension Expr[Color]
		Weight    Expr[Color]
	}

	Text struct {
		Color         Expr[Color]
		LineHeight    Expr[Color]
		TextAlignment gss.TextAlignment
	}

	Dimensions struct {
		Height gss.Height
		Width  gss.Width
	}

	// TODO: handle shorthand syntaxes during parsing
	Declarations struct {
		Dimensions      Dimensions
		Margin          Margin
		Padding         Padding
		Display         Display
		Text            Text
		Font            Font
		Border          Borders
		BorderRadiuses  BorderRadiuses
		BackgroundColor Expr[Color]
	}

	QualifiedRule struct {
		Component    string
		Declarations *Declarations
	}

	Stylesheet struct {
		Rules []*QualifiedRule
	}
)
