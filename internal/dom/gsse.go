package dom

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"go.ufukty.com/gss/internal/tokens/gss"
)

var (
	ErrDivisionByZero    = fmt.Errorf("division by zero is undefined")
	ErrIncompatibleUnits = fmt.Errorf("operands have incompetable units")
)

type PixelArea struct {
	Width, Height float64
}

// Context
type (
	Media struct {
		PrefersColorScheme string
	}
	Context struct {
		Media     Media
		Container PixelArea
		Viewport  PixelArea
	}
)

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
		Top, Right, Bottom, Left Expr[Pixels]
	}

	Padding struct {
		Top, Right, Bottom, Left Expr[Pixels]
	}

	Font struct {
		Family []gss.FontFamily `gss:"font-family"`
		Size   Expr[Pixels]     `gss:"font-size"`
		Weight Expr[Pixels]     `gss:"font-weight"`
	}

	Text struct {
		Color         Expr[gss.Color]   `gss:"color"`
		LineHeight    Expr[Pixels]      `gss:"line-height"`
		TextAlignment gss.TextAlignment `gss:"text-alignment"`
	}

	Dimensions struct {
		Height gss.Height `gss:"height"`
		Width  gss.Width  `gss:"width"`
	}

	// TODO: handle shorthand syntaxes during parsing
	Styles struct {
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

	Rule struct {
		Selector string
		Styles   *Styles
	}

	Gss struct {
		Rules []*Rule
	}
)

var (
	_ Expr[gss.Color] = (*LightDark[gss.Color])(nil)
)

func (c LightDark[T]) Resolve(ctx Context, e Element) (T, error) {
	if ctx.Media.PrefersColorScheme == "dark" {
		return c.Dark.Resolve(ctx, e)
	}
	return c.Light.Resolve(ctx, e)
}

// FIXME: Fetch identity value from DOM not AST once it is available
func (i Ident[Final]) Resolve(ctx Context, e Element) (Final, error)

func (a Addition) Resolve(ctx Context, e Element) (Size, error) {
	l, err := a.Operands.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Rhs.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Add(r)
}

func (a Neglect) Resolve(ctx Context, e Element) (Size, error) {
	l, err := a.Lhs.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Rhs.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Sub(r)
}

func (a Multiplication) Resolve(ctx Context, e Element) (Size, error) {
	l, err := a.Lhs.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Rhs.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Mul(r)
}

func (a Division) Resolve(ctx Context, e Element) (Size, error) {
	l, err := a.Dividend.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Divisor.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Div(r)
}

type Size struct {
	Number float64
	Unit   Unit
}

func (s Size) Pixels(ctx Context, e Element) (float64, error) {
	return -1, nil
}

type Pixeler interface {
	Pixels(ctx Context, e Element) (float64, error)
}

type Durationer interface {
	Duration(ctx Context, e Element) (float64, error)
}
type Angler interface {
	Angle(ctx Context, e Element) (float64, error)
}

func (s Size) Compare(t Size) bool {
	return s.Number == t.Number && s.Unit.Compare(t.Unit)
}

func (s Size) String() string {
	return fmt.Sprintf("%.0f%s", s.Number, s.Unit.String())
}

func (a Size) Add(b Size) (Size, error) {
	if !a.Unit.Compare(b.Unit) {
		return Size{}, fmt.Errorf("%s + %s: %w", a, b, ErrIncompatibleUnits)
	}
	return Size{a.Number + b.Number, a.Unit}, nil
}

func (a Size) Sub(b Size) (Size, error) {
	if !a.Unit.Compare(b.Unit) {
		return Size{}, fmt.Errorf("%s - %s: %w", a, b, ErrIncompatibleUnits)
	}
	return Size{a.Number - b.Number, a.Unit}, nil
}

func (a Size) Mul(b Size) (Size, error) {
	return Size{a.Number * b.Number, a.Unit.Multiply(b.Unit)}, nil
}

func (a Size) Div(b Size) (Size, error) {
	if b.Number == 0 {
		return Size{}, fmt.Errorf("%s / %s: %w", a, b, ErrDivisionByZero)
	}
	return Size{a.Number / b.Number, a.Unit.Divide(b.Unit)}, nil
}

func (s Size) Resolve(ctx Context, e Element) (Pixels, error) {
	return s, nil
}

type Unit map[gss.Unit]int // eg. px^2/em

func (a Unit) Compare(b Unit) bool {
	if len(a) != len(b) {
		return false
	}
	for unit := range a {
		if a[unit] != b[unit] {
			return false
		}
	}
	return true
}

func (a Unit) Multiply(b Unit) Unit {
	c := maps.Clone(a)
	for u, p := range b {
		c[u] += p
	}
	return c
}

func (a *Unit) clean() {
	for u, p := range *a {
		if p == 0 {
			delete(*a, u)
		}
	}
}

func (a Unit) Divide(b Unit) Unit {
	c := maps.Clone(a)
	for u, p := range b {
		c[u] -= p
	}
	c.clean()
	return c
}

var superscript = []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}

func positiveDigits(i int) []int {
	if i == 0 {
		return []int{0}
	}
	ds := []int{}
	for ; i > 0; i /= 10 {
		ds = append(ds, i%10)
	}
	slices.Reverse(ds)
	return ds
}

func super(i int) string {
	s := []string{}
	if i < 0 {
		i *= -1
		s = append(s, "⁻")
	}
	for _, d := range positiveDigits(i) {
		s = append(s, superscript[d])
	}
	return strings.Join(s, "")
}

// power in superscript unless ^1
func power(p int) string {
	if p == 1 {
		return ""
	}
	return super(p)
}

func (u Unit) String() string {
	us := []string{}
	for _, c := range slices.Sorted(maps.Keys(u)) {
		us = append(us, fmt.Sprintf("%s%s", c, power(u[c])))
	}
	return strings.Join(us, "·")
}

func Units(us ...gss.Unit) Unit {
	m := map[gss.Unit]int{}
	for _, u := range us {
		m[u] += 1
	}
	return m
}
