package gss

import (
	"io"

	"go.ufukty.com/gss/internal/files/gss/ast"
)

func rule(src io.Reader) (*ast.Rule, error)

func Gss(r io.Reader) (*ast.Gss, error) {
	g := &ast.Gss{
		Rules: []*ast.Rule{},
	}

	return g, nil
}
