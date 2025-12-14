package gss

import (
	"io"

	"go.ufukty.com/gss/internal/ast/gss"
)

func rule(src io.Reader) (*gss.Rule, error)

func Gss(r io.Reader) (*gss.Gss, error) {
	g := &gss.Gss{
		Rules: []*gss.Rule{},
	}

	return g, nil
}
