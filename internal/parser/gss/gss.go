package gss

import (
	"io"

	"go.ufukty.com/gss/internal/ast/ast"
)

func silent[T any](t T, _ error) T {
	return t
}

func rule(src io.Reader) (*ast.Rule, error)

func Gss(r io.Reader) (*ast.Gss, error) {
	g := &ast.Gss{
		Rules: []*ast.Rule{},
	}

	return g, nil
}
