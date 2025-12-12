package defaults

import (
	"go.ufukty.com/gss/internal/gss/ast"
	"go.ufukty.com/gss/internal/gss/tokens"
)

var Div = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Block, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        ast.FontSize{1.0, tokens.Unit_Em},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Img = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        ast.FontSize{1.0, tokens.Unit_Em},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Span = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        ast.FontSize{1.0, tokens.Unit_Em},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

func For(e ast.Element) *ast.Styles {
	switch e.(type) {
	case *ast.Div:
		return &Div
	case *ast.Img:
		return &Img
	case *ast.Span:
		return &Span
	}
	return nil
}
