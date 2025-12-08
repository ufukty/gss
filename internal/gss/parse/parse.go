package parse

import (
	"io"

	"go.ufukty.com/gss/internal/gss/ast"
)

func Gss(r io.Reader) (*ast.Element, error) {
	root := &ast.Element{
		Parent:   &ast.Element{},
		Children: []*ast.Element{},
		Tag:      "",
		Id:       "",
		Classes:  []string{},
		Styles:   ast.Styles{},
	}
	return nil, nil
}
