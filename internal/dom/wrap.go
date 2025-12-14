package dom

import (
	"go.ufukty.com/gss/internal/ast/html"
)

func wrap(e html.Element) Element {
	switch e := e.(type) {
	case *html.Div:
		return &Div{Ast: e}
	case *html.Html:
		return &Html{Ast: e}
	case *html.Img:
		return &Img{Ast: e}
	case *html.Span:
		return &Span{Ast: e}
	case *html.Text:
		return &Text{Ast: e}
	}
	return nil
}

// DOM for AST
func Wrap(a html.Element) Element {
	// _a ast
	// _d dom
	// p_ parent
	// c_ child
	d := wrap(a)

	pa, pa_ok := a.(html.Parent)
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
