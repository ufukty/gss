package dom

import "go.ufukty.com/gss/internal/ast/html"

type Element interface {
	GetAst() html.Element
}

func (d Div) GetAst() html.Element  { return d.Ast }
func (d Html) GetAst() html.Element { return d.Ast }
func (d Img) GetAst() html.Element  { return d.Ast }
func (d Span) GetAst() html.Element { return d.Ast }
func (d Text) GetAst() html.Element { return d.Ast }
