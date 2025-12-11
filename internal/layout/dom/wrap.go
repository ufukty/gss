package dom

import "go.ufukty.com/gss/internal/gss/ast"

func wrap(e ast.Element) Element {
	switch e := e.(type) {
	case *ast.Div:
		return &Div{Ast: e}
	case *ast.Html:
		return &Html{Ast: e}
	case *ast.Img:
		return &Img{Ast: e}
	case *ast.Span:
		return &Span{Ast: e}
	case *ast.Text:
		return &Text{Ast: e}
	}
	return nil
}

func Wrap(e ast.Element) Element {
	w := wrap(e)
	if p, ok := e.(ast.Adopter); ok {
		if pw, ok := w.(Adopter); ok {
			for _, child := range p.GetChildren() {
				pw.AppendChild(wrap(child))
			}
		}
	}
	return nil
}
