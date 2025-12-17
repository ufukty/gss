package gsse

import (
	"image/color"

	"go.ufukty.com/gss/internal/ast/html"
)

// Value types are to be used in instantiating [Expr] types.
type (
	Color        color.RGBA // eg. #FF0000
	Pixels       float64    // eg. 10px
	Angle        float64    // eg. 360deg
	Milliseconds int        // eg. 1000ms
	Image        string     // eg. url()
	FontSet      []string   // eg. "Helvetica", "Helvetica Neue", sans-serif
)

// Value types should also be usable directly.
var (
	_ Expr[Color]        = (*Color)(nil)
	_ Expr[Pixels]       = (*Pixels)(nil)
	_ Expr[Angle]        = (*Angle)(nil)
	_ Expr[Milliseconds] = (*Milliseconds)(nil)
	_ Expr[Image]        = (*Image)(nil)
	_ Expr[FontSet]      = (*FontSet)(nil)
)

func (f Color) Resolve(ctx Context, e html.Element) (Color, error) {
	return f, nil
}

func (f Pixels) Resolve(ctx Context, e html.Element) (Pixels, error) {
	return f, nil
}

func (f Angle) Resolve(ctx Context, e html.Element) (Angle, error) {
	return f, nil
}

func (f Milliseconds) Resolve(ctx Context, e html.Element) (Milliseconds, error) {
	return f, nil
}

func (f Image) Resolve(ctx Context, e html.Element) (Image, error) {
	return f, nil
}

func (f FontSet) Resolve(ctx Context, e html.Element) (FontSet, error) {
	return f, nil
}
