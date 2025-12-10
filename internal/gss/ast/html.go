package ast

type Element interface {
	element()
}

// Element
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

	Text struct {
		Content string
	}
)

func (*Div) element()  {}
func (*Html) element() {}
func (*Img) element()  {}
func (*Span) element() {}
func (*Text) element() {}

func (d Div) GetChildren() []Element  { return d.Children }
func (h Html) GetChildren() []Element { return h.Children }
func (s Span) GetChildren() []Element { return s.Children }

type Adopter interface {
	GetChildren() []Element
}

func Visit(e Element, visitor func(Element) bool) {
	c := visitor(e)
	if !c {
		return
	}
	defer visitor(nil)

	if a, ok := e.(Adopter); ok {
		for _, child := range a.GetChildren() {
			Visit(child, visitor)
		}
	}
}
