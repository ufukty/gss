package ast

import (
	"fmt"

	"go.ufukty.com/gss/internal/files/gss/tokens"
	"go.ufukty.com/gss/internal/html/ast"
)

type (
	Media struct {
		PrefersColorScheme string
	}
	Context struct {
		Media Media
	}
)

// Expr of Final is a GSS expression that can be resolved to
// the type Final when its Resolve method is called with the
// rendering context.
type Expr[Final any] interface {
	Resolve(Context, ast.Element) (Final, error)
}

var _ Expr[tokens.Color] = (*LightDark[tokens.Color])(nil)

func (c LightDark[T]) Resolve(ctx Context, e ast.Element) (T, error) {
	if ctx.Media.PrefersColorScheme == "dark" {
		return c.Dark.Resolve(ctx, e)
	}
	return c.Light.Resolve(ctx, e)
}

// FIXME: Fetch identity value from DOM not AST once it is available
func (i Ident[Final]) Resolve(ctx Context, e ast.Element) (Final, error)

func (a Addition) Resolve(ctx Context, e ast.Element) (Size, error) {
	l, err := a.Lhs.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("lhs: %w", err)
	}
	r, err := a.Rhs.Resolve(ctx, e)
	if err != nil {
		return Size{}, fmt.Errorf("rhs: %w", err)
	}
	return l.Add(r)
}

func (a Subtraction) Resolve(ctx Context, e ast.Element) (Size, error) {
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

func (a Multiplication) Resolve(ctx Context, e ast.Element) (Size, error) {
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

func (a Division) Resolve(ctx Context, e ast.Element) (Size, error) {
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
