package ast

type Element interface {
	elem()
}

type (
	Div struct {
		Id       string
		Classes  []string
		Parent   Element
		Children []Element

		TextContent string
	}

	Span struct {
		Id       string
		Classes  []string
		Parent   Element
		Children []Element

		TextContent string
	}

	Img struct {
		Id       string
		Classes  []string
		Parent   Element
		Children []Element

		Src    string
		SrcSet map[float64]string
	}
)

func (Div) elem()  {}
func (Span) elem() {}
func (Img) elem()  {}

type Html struct {
	Root Element
}
