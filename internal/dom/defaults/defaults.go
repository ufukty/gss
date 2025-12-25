package defaults

import (
	"go.ufukty.com/gss/internal/ast"
)

var (
	noBorder = ast.Borders{
		Top:    ast.Border{"#000", "none", "0"},
		Right:  ast.Border{"#000", "none", "0"},
		Bottom: ast.Border{"#000", "none", "0"},
		Left:   ast.Border{"#000", "none", "0"},
	}
	noBorderRadius = ast.BorderRadiuses{
		TopLeft:     "0",
		TopRight:    "0",
		BottomRight: "0",
		BottomLeft:  "0",
	}
	autoDimensions = ast.Dimensions{
		Width:  "auto",
		Height: "auto",
	}
	inheritFont = ast.Font{
		Family: "inherit",
		Size:   "1em",
		Weight: "400",
	}
	noMargin = ast.Margin{
		Top:    "0",
		Right:  "0",
		Bottom: "0",
		Left:   "0",
	}
	somePadding = ast.Padding{
		Top:    "10px",
		Right:  "10px",
		Bottom: "10px",
		Left:   "10px",
	}
	block = ast.Display{
		Outside: "block",
		Inside:  "flow",
	}
	inline = ast.Display{
		Outside: "inline",
		Inside:  "flow",
	}
	inheritText = ast.Text{
		Color:         "inherit",
		LineHeight:    "inherit",
		TextAlignment: "inherit",
	}
)

var (
	Div = ast.Styles{
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
	Html = ast.Styles{
		BackgroundColor: "#fff",
		Border:          noBorder,
		BorderRadiuses:  noBorderRadius,
		Dimensions:      autoDimensions,
		Display:         block,
		Font:            ast.Font{Family: "sans-serif", Size: "14px", Weight: "400"},
		Margin:          noMargin,
		Padding:         somePadding,
		Text:            ast.Text{Color: "#000", LineHeight: "inherit", TextAlignment: "left"},
	}
	Img = ast.Styles{
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
	Span = ast.Styles{
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
	TextNode = ast.Styles{
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
	case *ast.TextNode:
		return &TextNode
	}
	return nil
}
