package html

import (
	"io"

	"go.ufukty.com/gss/internal/ast"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func findBody(r *html.Node) *html.Node {
	if r.DataAtom == atom.Body {
		return r
	}
	for c := range r.ChildNodes() {
		if f := findBody(c); f != nil {
			return f
		}
	}
	return nil
}

func Html(src io.Reader) (*ast.Html, error) {
	return nil, nil
}
