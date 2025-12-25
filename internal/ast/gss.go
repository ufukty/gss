package ast

// Gss
type (
	Display struct {
		Outside string
		Inside  string
	}
	Border struct {
		Color     any // "inherit", "transparent", color.NRGBA
		Style     any // "inherit", gss.BorderStyle
		Thickness any // "none", dimensional.Dimension
	}
	BorderRadiuses struct {
		TopLeft, TopRight, BottomRight, BottomLeft any // "none", "inherit", dimensional.Dimension
	}
	Borders struct {
		Top, Right, Bottom, Left Border
	}
	Margin struct {
		Top, Right, Bottom, Left any // "inherit", dimensional.Dimension
	}
	Padding struct {
		Top, Right, Bottom, Left any // "inherit", dimensional.Dimension
	}
	Font struct {
		Family any // "inherit", []string
		Size   any // "inherit", dimensional.Dimension
		Weight any // "inherit", int
	}
	Text struct {
		Color         any // "inherit", "transparent", color.NRGBA
		LineHeight    any // "inherit", dimension.Dimensional
		TextAlignment any // "inherit", gss.TextAlignment
	}
	Dimensions struct {
		Height any // "auto", "min-content", "max-content", dimensional.Dimension
		Width  any // "auto", "min-content", "max-content", dimensional.Dimension
	}
	// TODO: handle shorthand syntaxes during parsing
	Styles struct {
		Dimensions      Dimensions
		Margin          Margin
		Padding         Padding
		Display         Display
		Text            Text
		Font            Font
		Border          Borders
		BorderRadiuses  BorderRadiuses
		BackgroundColor any // "inherit", "transparent", color.NRGBA
	}
	Rule struct {
		Selector string
		Styles   *Styles
	}
	Gss struct {
		Rules []*Rule
	}
)
