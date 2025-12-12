package dom

import "go.ufukty.com/gss/internal/html/ast"

type Size struct {
	Width, Height float64
}

type (
	Div struct {
		Ast      *ast.Div
		Parent   Parent
		Children []Child
		Min, Max Size
	}

	Html struct {
		Ast      *ast.Html
		Children []Child
		Min, Max Size
	}

	Img struct {
		Ast      *ast.Img
		Parent   Parent
		Min, Max Size
	}

	Span struct {
		Ast      *ast.Span
		Parent   Parent
		Children []Child
		Min, Max Size
	}

	Text struct {
		Ast      *ast.Text
		Parent   Parent
		Min, Max Size
	}
)
