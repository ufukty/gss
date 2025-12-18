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
	Family string `gss:"font-family"`
	Size   string `gss:"font-size"`
	Weight string `gss:"font-weight"`
}

type Text struct {
	Color         string `gss:"color"`
	LineHeight    string `gss:"line-height"`
	TextAlignment string `gss:"text-alignment"`
}

type Dimensions struct {
	Height string `gss:"height"`
	Width  string `gss:"width"`
}

// TODO: handle shorthand syntaxes during parsing
type Styles struct {
	Dimensions      Dimensions
	Margin          Margin  `gss:"margin"`
	Padding         Padding `gss:"padding"`
	Display         Display `gss:"display"`
	Text            Text
	Font            Font
	Border          Borders        `gss:"border"`
	BorderRadiuses  BorderRadiuses `gss:"border-radius"`
	BackgroundColor string         `gss:"background-color"`
}

type Rule struct {
	Selector string
	Styles   *Styles
}

type Gss struct {
	Rules []*Rule
}
