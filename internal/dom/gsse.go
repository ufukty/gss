package dom

import (
	"go.ufukty.com/gss/internal/tokens/tokens"
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

// func subtree[T any](a any, ctx Context, e Element) (T, error) {
// 	switch a := a.(type) {
// 	case any // T
// 		return a.Resolve(ctx, e)
// 	case T:
// 		return a, nil
// 	default:
// 		var t T
// 		return t, fmt.Errorf("value of unknown type: %T", a)
// 	}
// }

// func (c LightDark) Resolve(ctx Context, e Element) (Color, error) {
// 	if ctx.Media.PrefersColorScheme == "dark" {
// 		return subtree[Color](c.Dark, ctx, e)
// 	}
// 	return subtree[Color](c.Light, ctx, e)
// }

// // FIXME: Fetch identity value from DOM not AST once it is available
// func (i Ident[Resolving]) Resolve(ctx Context, e Element) (Resolving, error) {
// 	return *new(Resolving), nil
// }

// func (a Addition[F]) Resolve(ctx Context, e Element) (F, error) {
// 	l, err := subtree[F](a.Lhs, ctx, e)
// 	if err != nil {
// 		return l, fmt.Errorf("left operand: %w", err)
// 	}
// 	r, err := subtree[F](a.Rhs, ctx, e)
// 	if err != nil {
// 		return r, fmt.Errorf("right operand: %w", err)
// 	}
// 	return l + r, nil
// }

// func (a Subtraction[F]) Resolve(ctx Context, e Element) (F, error) {
// 	l, err := subtree[F](a.Lhs, ctx, e)
// 	if err != nil {
// 		return l, fmt.Errorf("left operand: %w", err)
// 	}
// 	r, err := subtree[F](a.Rhs, ctx, e)
// 	if err != nil {
// 		return r, fmt.Errorf("right operand: %w", err)
// 	}
// 	return l - r, nil
// }

// func (a Multiplication[F]) Resolve(ctx Context, e Element) (F, error) {
// 	l, err := subtree[F](a.Lhs, ctx, e)
// 	if err != nil {
// 		return l, fmt.Errorf("left operand: %w", err)
// 	}
// 	r, err := subtree[F](a.Rhs, ctx, e)
// 	if err != nil {
// 		return r, fmt.Errorf("right operand: %w", err)
// 	}
// 	return l * r, nil
// }

// func (a Division[F]) Resolve(ctx Context, e Element) (F, error) {
// 	l, err := subtree[F](a.Dividend, ctx, e)
// 	if err != nil {
// 		return l, fmt.Errorf("left operand: %w", err)
// 	}
// 	r, err := subtree[F](a.Divisor, ctx, e)
// 	if err != nil {
// 		return r, fmt.Errorf("right operand: %w", err)
// 	}
// 	if r == 0 {
// 		return r, ErrDivisionByZero
// 	}
// 	return l / r, nil
// }

// Nodes
type (
	Display struct {
		Outside tokens.DisplayOutside
		Inside  tokens.DisplayInside
	}

	Border struct {
		Color     any // Color
		Style     string
		Thickness any // Pixels
	}

	BorderRadiuses struct {
		TopLeft, TopRight, BottomRight, BottomLeft any // Pixels
	}

	Borders struct {
		Top, Right, Bottom, Left Border
	}

	Margin struct {
		Top, Right, Bottom, Left any // Pixels
	}

	Padding struct {
		Top, Right, Bottom, Left any // Pixels
	}

	Font struct {
		Family    []tokens.FontFamily
		Dimension any // Color
		Weight    any // Color
	}

	Text struct {
		Color         any // Color
		LineHeight    any // Color
		TextAlignment tokens.TextAlignment
	}

	Dimensions struct {
		Height tokens.Height
		Width  tokens.Width
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
		BackgroundColor any // Color
	}

	QualifiedRule struct {
		Component    string
		Declarations *Declarations
	}

	Stylesheet struct {
		Rules []*QualifiedRule
	}
)
