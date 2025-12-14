package defaults

import (
	gss "go.ufukty.com/gss/internal/files/gss/ast"
	"go.ufukty.com/gss/internal/files/gss/tokens"
	gsse "go.ufukty.com/gss/internal/gsse/ast"
	html "go.ufukty.com/gss/internal/html/ast"
)

var Div = gss.Styles{
	Display:         gss.Display{tokens.DisplayOutside_Block, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(tokens.Unit_Em)},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Html = gss.Styles{
	Display:         gss.Display{tokens.DisplayOutside_Block, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_SansSerif},
	FontSize:        gsse.Size{12.0, gsse.Units(tokens.Unit_Px)},
	Color:           "#000",
	BackgroundColor: "#fff",
}

var Img = gss.Styles{
	Display:         gss.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(tokens.Unit_Em)},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Span = gss.Styles{
	Display:         gss.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(tokens.Unit_Em)},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

var Text = gss.Styles{
	Display:         gss.Display{tokens.DisplayOutside_Inline, tokens.DisplayInside_Flow},
	Width:           tokens.Width_Auto,
	Height:          tokens.Height_Auto,
	FontFamily:      []tokens.FontFamily{tokens.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(tokens.Unit_Em)},
	Color:           tokens.Color_Inherit,
	BackgroundColor: tokens.BackgroundColor_Inherit,
}

func For(e html.Element) *gss.Styles {
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
