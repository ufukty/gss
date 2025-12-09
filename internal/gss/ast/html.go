package ast

import "go.ufukty.com/gss/internal/gss/tokens"

type Element interface {
	elem()
}

type (
	Div struct {
		Tag      tokens.Tag
		Id       string
		Classes  []string
		Parent   Element
		Children []Element

		TextContent string
	}

	Span struct {
		Tag      tokens.Tag
		Id       string
		Classes  []string
		Parent   Element
		Children []Element

		TextContent string
	}

	Img struct {
		Tag      tokens.Tag
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
