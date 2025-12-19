package dom

import (
	"fmt"
	"image/color"

	"go.ufukty.com/gss/internal/dom/unit"
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
		Unit   unit.Unit
	}
)

func (d Dimension) Color(ctx Context, e Element) (float64, error) {
	if !d.Unit.Compare(Units("px")) {
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

func (d Dimension) Add(b Dimension) (Dimension, error) {
	if !d.Unit.Compare(b.Unit) {
		return Dimension{}, fmt.Errorf("%s + %s: %w", d, b, ErrIncompatibleUnits)
	}
	return Dimension{d.Number + b.Number, d.Unit}, nil
}

func (d Dimension) Sub(b Dimension) (Dimension, error) {
	if !d.Unit.Compare(b.Unit) {
		return Dimension{}, fmt.Errorf("%s - %s: %w", d, b, ErrIncompatibleUnits)
	}
	return Dimension{d.Number - b.Number, d.Unit}, nil
}

func (d Dimension) Mul(b Dimension) (Dimension, error) {
	return Dimension{d.Number * b.Number, d.Unit.Multiply(b.Unit)}, nil
}

func (d Dimension) Div(b Dimension) (Dimension, error) {
	if b.Number == 0 {
		return Dimension{}, fmt.Errorf("%s / %s: %w", d, b, ErrDivisionByZero)
	}
	return Dimension{d.Number / b.Number, d.Unit.Divide(b.Unit)}, nil
}

func (d Dimension) Resolve(ctx Context, e Element) (Dimension, error) {
	return d, nil
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
	Addition struct {
		Lhs, Rhs Expr[Dimension]
	}
	Subtraction struct {
		Lhs, Rhs Expr[Dimension]
	}
	Multiplication struct {
		Lhs, Rhs Expr[Dimension]
	}
	Division struct {
		Dividend, Divisor Expr[Dimension]
	}
)

var (
	_ Expr[Dimension] = (*Addition)(nil)
	_ Expr[Dimension] = (*Subtraction)(nil)
	_ Expr[Dimension] = (*Multiplication)(nil)
	_ Expr[Dimension] = (*Division)(nil)
)

func (c LightDark) Resolve(ctx Context, e Element) (Color, error) {
	if ctx.Media.PrefersColorScheme == "dark" {
		return c.Dark.Resolve(ctx, e)
	}
	return c.Light.Resolve(ctx, e)
}

// FIXME: Fetch identity value from DOM not AST once it is available
func (i Ident[Resolving]) Resolve(ctx Context, e Element) (Resolving, error)

func (a Addition) Resolve(ctx Context, e Element) (Dimension, error) {
	l, err := a.Lhs.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Rhs.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Add(r)
}

func (a Subtraction) Resolve(ctx Context, e Element) (Dimension, error) {
	l, err := a.Lhs.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Rhs.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Sub(r)
}

func (a Multiplication) Resolve(ctx Context, e Element) (Dimension, error) {
	l, err := a.Lhs.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Rhs.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Mul(r)
}

func (a Division) Resolve(ctx Context, e Element) (Dimension, error) {
	l, err := a.Dividend.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Divisor.Resolve(ctx, e)
	if err != nil {
		return Dimension{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Div(r)
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
		Thickness Expr[Color]
	}

	BorderRadiuses struct {
		TopLeft, TopRight, BottomRight, BottomLeft Expr[Color]
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
		Family    []gss.FontFamily `gss:"font-family"`
		Dimension Expr[Color]      `gss:"font-size"`
		Weight    Expr[Color]      `gss:"font-weight"`
	}

	Text struct {
		Color         Expr[gss.Color]   `gss:"color"`
		LineHeight    Expr[Color]       `gss:"line-height"`
		TextAlignment gss.TextAlignment `gss:"text-alignment"`
	}

	Dimensions struct {
		Height gss.Height `gss:"height"`
		Width  gss.Width  `gss:"width"`
	}

	// TODO: handle shorthand syntaxes during parsing
	Declarations struct {
		Dimensions      Dimensions
		Margin          Margin  `gss:"margin"`
		Padding         Padding `gss:"padding"`
		Display         Display `gss:"display"`
		Text            Text
		Font            Font
		Border          Borders         `gss:"border"`
		BorderRadiuses  BorderRadiuses  `gss:"border-radius"`
		BackgroundColor Expr[gss.Color] `gss:"background-color"`
	}

	QualifiedRule struct {
		Component    string
		Declarations *Declarations
	}

	Stylesheet struct {
		Rules []*QualifiedRule
	}
)
