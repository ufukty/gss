package dom

import "go.ufukty.com/gss/internal/ast/html"

type Size struct {
	Width, Height float64
}

type (
	Div struct {
		Ast      *html.Div
		Parent   Parent
		Children []Child
		Min, Max Size
	}

	Html struct {
		Ast      *html.Html
		Children []Child
		Min, Max Size
	}

	Img struct {
		Ast      *html.Img
		Parent   Parent
		Min, Max Size
	}

	Span struct {
		Ast      *html.Span
		Parent   Parent
		Children []Child
		Min, Max Size
	}

	TextNode struct {
		Ast      *html.TextNode
		Parent   Parent
		Min, Max Size
	}
)

type Element interface {
	GetAst() html.Element
}

func (d Div) GetAst() html.Element      { return d.Ast }
func (d Html) GetAst() html.Element     { return d.Ast }
func (d Img) GetAst() html.Element      { return d.Ast }
func (d Span) GetAst() html.Element     { return d.Ast }
func (d TextNode) GetAst() html.Element { return d.Ast }

// TODO: update when Go allows interface defining on common fields
type Parent interface {
	Element
	AppendChild(e Child)
	GetChildren() []Child
}

func (d *Div) AppendChild(e Child)  { d.Children = append(d.Children, e) }
func (h *Html) AppendChild(e Child) { h.Children = append(h.Children, e) }
func (s *Span) AppendChild(e Child) { s.Children = append(s.Children, e) }

func (d Div) GetChildren() []Child  { return d.Children }
func (h Html) GetChildren() []Child { return h.Children }
func (s Span) GetChildren() []Child { return s.Children }

var (
	_ Parent = (*Div)(nil)
	_ Parent = (*Html)(nil)
	_ Parent = (*Span)(nil)
)

type Child interface {
	Element
	GetParent() Parent
	SetParent(e Parent)
}

func (d Div) GetParent() Parent      { return d.Parent }
func (i Img) GetParent() Parent      { return i.Parent }
func (s Span) GetParent() Parent     { return s.Parent }
func (t TextNode) GetParent() Parent { return t.Parent }

func (d *Div) SetParent(p Parent)      { d.Parent = p }
func (i *Img) SetParent(p Parent)      { i.Parent = p }
func (s *Span) SetParent(p Parent)     { s.Parent = p }
func (t *TextNode) SetParent(p Parent) { t.Parent = p }

var (
	_ Child = (*Div)(nil)
	_ Child = (*Img)(nil)
	_ Child = (*Span)(nil)
	_ Child = (*TextNode)(nil)
)
