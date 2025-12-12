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

var Html = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Block, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_SansSerif},
	FontSize:        ast.FontSize{12.0, tokens.Unit_Px},
	Color:           "#000",
	BackgroundColor: "#fff",
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

var Text = ast.Styles{
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
	case *ast.Html:
		return &Html
	case *ast.Img:
		return &Img
	case *ast.Span:
		return &Span
	case *ast.Text:
		return &Text
	}
	return nil
}
