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

func visit(children []Element, visitor func(Element) bool) {
	for _, child := range children {
		Visit(child, visitor)
	}
}

func Visit(e Element, visitor func(Element) bool) {
	c := visitor(e)
	if !c {
		return
	}
	defer visitor(nil)

	switch e := e.(type) {
	case *Div:
		visit(e.Children, visitor)
	case *Html:
		visit(e.Children, visitor)
	case *Span:
		visit(e.Children, visitor)
	}
}
