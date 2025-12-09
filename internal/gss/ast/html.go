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
func (*Img) element()  {}
func (*Span) element() {}
func (*Text) element() {}

type Html struct {
	Root Element
}

func Visit(e Element, visitor func(Element) bool) {
	c := visitor(e)
	if !c {
		return
	}
	defer visitor(nil)

	switch root := e.(type) {
	case *Div:
		for _, child := range root.Children {
			Visit(child, visitor)
		}
	case *Span:
		for _, child := range root.Children {
			Visit(child, visitor)
		}
	}
}
