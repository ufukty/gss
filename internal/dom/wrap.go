package dom

import (
	"go.ufukty.com/gss/internal/html/ast"
)

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

// DOM for AST
func Wrap(a ast.Element) Element {
	// _a ast
	// _d dom
	// p_ parent
	// c_ child
	d := wrap(a)

	pa, pa_ok := a.(ast.Parent)
	pd, pd_ok := d.(Parent)
	if pa_ok && pd_ok {
		for _, ca := range pa.GetChildren() {
			cd := Wrap(ca).(Child)
			pd.AppendChild(cd)
			cd.SetParent(pd)
		}
	}

	return d
}
