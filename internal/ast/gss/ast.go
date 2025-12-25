package gss

type Display struct {
	Outside string
	Inside  string
}

type Border struct {
	Color     any // "inherit", "transparent", color.NRGBA
	Style     any // "inherit", gss.BorderStyle
	Thickness any // "none", dimensional.Dimension
}

type BorderRadiuses struct {
	TopLeft, TopRight, BottomRight, BottomLeft any // "none", "inherit", dimensional.Dimension
}

type Borders struct {
	Top, Right, Bottom, Left Border
}

type Margin struct {
	Top, Right, Bottom, Left any // "inherit", dimensional.Dimension
}

type Padding struct {
	Top, Right, Bottom, Left any // "inherit", dimensional.Dimension
}

type Font struct {
	Family any // "inherit", []string
	Size   any // "inherit", dimensional.Dimension
	Weight any // "inherit", int
}

type Text struct {
	Color         any // "inherit", "transparent", color.NRGBA
	LineHeight    any // "inherit", dimension.Dimensional
	TextAlignment any // "inherit", gss.TextAlignment
}

type Dimensions struct {
	Height any // "auto", "min-content", "max-content", dimensional.Dimension
	Width  any // "auto", "min-content", "max-content", dimensional.Dimension
}

// TODO: handle shorthand syntaxes during parsing
type Styles struct {
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

type Rule struct {
	Selector string
	Styles   *Styles
}

type Gss struct {
	Rules []*Rule
}
