package dom

import "go.ufukty.com/gss/internal/gss/ast"

type Element interface {
	GetAst() ast.Element
}

func (d Div) GetAst() ast.Element  { return d.Ast }
func (d Html) GetAst() ast.Element { return d.Ast }
func (d Img) GetAst() ast.Element  { return d.Ast }
func (d Span) GetAst() ast.Element { return d.Ast }
func (d Text) GetAst() ast.Element { return d.Ast }
