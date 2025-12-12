package defaults

import (
	"go.ufukty.com/gss/internal/gss/ast"
	"go.ufukty.com/gss/internal/gss/tokens"
	html "go.ufukty.com/gss/internal/html/ast"
)

var Div = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Block, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        ast.Size{1.0, tokens.Unit_Em},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Html = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Block, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_SansSerif},
	FontSize:        ast.Size{12.0, tokens.Unit_Px},
	Color:           "#000",
	BackgroundColor: "#fff",
}

var Img = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        ast.Size{1.0, tokens.Unit_Em},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Span = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        ast.Size{1.0, tokens.Unit_Em},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Text = ast.Styles{
	Display:         ast.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        ast.Size{1.0, tokens.Unit_Em},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

func For(e html.Element) *ast.Styles {
	switch e.(type) {
	case *html.Div:
		return &Div
	case *html.Html:
		return &Html
	case *html.Img:
		return &Img
	case *html.Span:
		return &Span
	case *html.Text:
		return &Text
	}
	return nil
}
