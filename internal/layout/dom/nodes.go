package dom

import "go.ufukty.com/gss/internal/gss/ast"

type Size struct {
	Width, Height float64
}

type (
	Div struct {
		Ast      *ast.Div
		Parent   Element
		Children []Element
		Min, Max Size
	}

	Html struct {
		Ast      *ast.Html
		Children []Element
		Min, Max Size
	}

	Img struct {
		Ast      *ast.Img
		Parent   Element
		Min, Max Size
	}

	Span struct {
		Ast      *ast.Span
		Parent   Element
		Children []Element
		Min, Max Size
	}

	Text struct {
		Ast      *ast.Text
		Parent   Element
		Min, Max Size
	}
)
