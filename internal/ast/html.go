package ast

// Html
type (
	Div struct {
		Id       string
		Classes  []string
		Parent   Element
		Children []Element
	}
	Html struct {
		Children []Element
	}
	Img struct {
		Id      string
		Classes []string
		Parent  Element
		Src     string
		SrcSet  map[float64]string
	}
	Span struct {
		Id       string
		Classes  []string
		Parent   Element
		Children []Element
	}
	TextNode struct {
		Content string
	}
)

type Element interface {
	element()
}

func (*Div) element()      {}
func (*Html) element()     {}
func (*Img) element()      {}
func (*Span) element()     {}
func (*TextNode) element() {}

type Parent interface {
	GetChildren() []Element
}

func (d Div) GetChildren() []Element  { return d.Children }
func (h Html) GetChildren() []Element { return h.Children }
func (s Span) GetChildren() []Element { return s.Children }
