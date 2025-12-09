package ast

import "go.ufukty.com/gss/internal/gss/tokens"

type Element interface {
	elem()
}

type element struct {
	Tag        tokens.Tag
	Id         string
	Classes    []string
	Attributes map[string]string
	Parent     Element
	Children   []Element
}

type (
	Div struct {
		element
		TextContent string
	}

	Span struct {
		element
		TextContent string
	}

	Img struct {
		element
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
