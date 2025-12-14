package defaults

import (
	gss "go.ufukty.com/gss/internal/ast/gss"
	gsse "go.ufukty.com/gss/internal/ast/gsse"
	html "go.ufukty.com/gss/internal/ast/html"
	gss1 "go.ufukty.com/gss/internal/tokens/gss"
)

var Div = gss.Styles{
	Display:         gss.Display{gss1.DisplayOutside_Block, gss1.DisplayInside_Flow},
	Width:           gss1.Width_Auto,
	Height:          gss1.Height_Auto,
	FontFamily:      []gss1.FontFamily{gss1.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(gss1.Unit_Em)},
	Color:           gss1.Color_Inherit,
	BackgroundColor: gss1.BackgroundColor_Inherit,
}

var Html = gss.Styles{
	Display:         gss.Display{gss1.DisplayOutside_Block, gss1.DisplayInside_Flow},
	Width:           gss1.Width_Auto,
	Height:          gss1.Height_Auto,
	FontFamily:      []gss1.FontFamily{gss1.FontFamily_SansSerif},
	FontSize:        gsse.Size{12.0, gsse.Units(gss1.Unit_Px)},
	Color:           "#000",
	BackgroundColor: "#fff",
}

var Img = gss.Styles{
	Display:         gss.Display{gss1.DisplayOutside_Inline, gss1.DisplayInside_Flow},
	Width:           gss1.Width_Auto,
	Height:          gss1.Height_Auto,
	FontFamily:      []gss1.FontFamily{gss1.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(gss1.Unit_Em)},
	Color:           gss1.Color_Inherit,
	BackgroundColor: gss1.BackgroundColor_Inherit,
}

var Span = gss.Styles{
	Display:         gss.Display{gss1.DisplayOutside_Inline, gss1.DisplayInside_Flow},
	Width:           gss1.Width_Auto,
	Height:          gss1.Height_Auto,
	FontFamily:      []gss1.FontFamily{gss1.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(gss1.Unit_Em)},
	Color:           gss1.Color_Inherit,
	BackgroundColor: gss1.BackgroundColor_Inherit,
}

var Text = gss.Styles{
	Display:         gss.Display{gss1.DisplayOutside_Inline, gss1.DisplayInside_Flow},
	Width:           gss1.Width_Auto,
	Height:          gss1.Height_Auto,
	FontFamily:      []gss1.FontFamily{gss1.FontFamily_Inherit},
	FontSize:        gsse.Size{1.0, gsse.Units(gss1.Unit_Em)},
	Color:           gss1.Color_Inherit,
	BackgroundColor: gss1.BackgroundColor_Inherit,
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
