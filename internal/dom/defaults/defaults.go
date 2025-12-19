package defaults

import (
	"go.ufukty.com/gss/internal/ast/gss"

	"go.ufukty.com/gss/internal/ast/html"
)

var (
	noBorder = gss.Borders{
		Top:    gss.Border{"#000", "none", "0"},
		Right:  gss.Border{"#000", "none", "0"},
		Bottom: gss.Border{"#000", "none", "0"},
		Left:   gss.Border{"#000", "none", "0"},
	}
	noBorderRadius = gss.BorderRadiuses{
		TopLeft:     "0",
		TopRight:    "0",
		BottomRight: "0",
		BottomLeft:  "0",
	}
	autoDimensions = gss.Dimensions{
		Width:  "auto",
		Height: "auto",
	}
	inheritFont = gss.Font{
		Family: "inherit",
		Size:   "1em",
		Weight: "400",
	}
	noMargin = gss.Margin{
		Top:    "0",
		Right:  "0",
		Bottom: "0",
		Left:   "0",
	}
	somePadding = gss.Padding{
		Top:    "10px",
		Right:  "10px",
		Bottom: "10px",
		Left:   "10px",
	}
	block = gss.Display{
		Outside: "block",
		Inside:  "flow",
	}
	inline = gss.Display{
		Outside: "inline",
		Inside:  "flow",
	}
	inheritText = gss.Text{
		Color:         "inherit",
		LineHeight:    "inherit",
		TextAlignment: "inherit",
	}
)

var (
	Div = gss.Styles{
		BackgroundColor: "none",
		Border:          noBorder,
		BorderRadiuses:  noBorderRadius,
		Dimensions:      autoDimensions,
		Display:         block,
		Font:            inheritFont,
		Margin:          noMargin,
		Padding:         somePadding,
		Text:            inheritText,
	}
	Html = gss.Styles{
		BackgroundColor: "#fff",
		Border:          noBorder,
		BorderRadiuses:  noBorderRadius,
		Dimensions:      autoDimensions,
		Display:         block,
		Font:            gss.Font{Family: "sans-serif", Size: "14px", Weight: "400"},
		Margin:          noMargin,
		Padding:         somePadding,
		Text:            gss.Text{Color: "#000", LineHeight: "inherit", TextAlignment: "left"},
	}
	Img = gss.Styles{
		BackgroundColor: "none",
		Border:          noBorder,
		BorderRadiuses:  noBorderRadius,
		Dimensions:      autoDimensions,
		Display:         inline,
		Font:            inheritFont,
		Margin:          noMargin,
		Padding:         somePadding,
		Text:            inheritText,
	}
	Span = gss.Styles{
		BackgroundColor: "none",
		Border:          noBorder,
		BorderRadiuses:  noBorderRadius,
		Dimensions:      autoDimensions,
		Display:         inline,
		Font:            inheritFont,
		Margin:          noMargin,
		Padding:         somePadding,
		Text:            inheritText,
	}
	TextNode = gss.Styles{
		BackgroundColor: "none",
		Border:          noBorder,
		BorderRadiuses:  noBorderRadius,
		Dimensions:      autoDimensions,
		Display:         inline,
		Font:            inheritFont,
		Margin:          noMargin,
		Padding:         somePadding,
		Text:            inheritText,
	}
)

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
	case *html.TextNode:
		return &TextNode
	}
	return nil
}
