// Strings are used for quantitive values as they can have multiple
// units or even defined by complex expressions (GSSE) that is not
// handled in this step.
package gss

type Display struct {
	Outside string
	Inside  string
}

type Border struct {
	Color     string
	Style     string
	Thickness string
}

type BorderRadiuses struct {
	TopLeft, TopRight, BottomRight, BottomLeft string
}

type Borders struct {
	Top, Right, Bottom, Left Border
}

type Margin struct {
	Top, Right, Bottom, Left string
}

type Padding struct {
	Top, Right, Bottom, Left string
}

type Font struct {
	Family string
	Size   string
	Weight string
}

type Text struct {
	Color         string
	LineHeight    string
	TextAlignment string
}

type Dimensions struct {
	Height string
	Width  string
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
	BackgroundColor string
}

type Rule struct {
	Selector string
	Styles   *Styles
}

type Gss struct {
	Rules []*Rule
}
