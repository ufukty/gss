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

	Text struct {
		Ast      *html.Text
		Parent   Parent
		Min, Max Size
	}
)
