package gss

import (
	"go.ufukty.com/gss/internal/ast/gsse"
	"go.ufukty.com/gss/internal/tokens/gss"
)

type Display struct {
	Outside gss.DisplayOutside
	Inside  gss.DisplayInside
}

type Border struct {
	Color     gsse.Expr[gsse.Color]
	Style     string
	Thickness gsse.Expr[gsse.Pixels]
}

type BorderRadiuses struct {
	TopLeft, TopRight, BottomRight, BottomLeft gsse.Expr[gsse.Pixels]
}

type Borders struct {
	Top, Right, Bottom, Left Border
}

type Margin struct {
	Top, Right, Bottom, Left gsse.Expr[gsse.Pixels]
}

type Padding struct {
	Top, Right, Bottom, Left gsse.Expr[gsse.Pixels]
}

type Font struct {
	Family []gss.FontFamily       `gss:"font-family"`
	Size   gsse.Expr[gsse.Pixels] `gss:"font-size"`
	Weight gsse.Expr[gsse.Pixels] `gss:"font-weight"`
}

type Text struct {
	Color         gsse.Expr[gss.Color]   `gss:"color"`
	LineHeight    gsse.Expr[gsse.Pixels] `gss:"line-height"`
	TextAlignment gss.TextAlignment      `gss:"text-alignment"`
}

type Dimensions struct {
	Height gss.Height `gss:"height"`
	Width  gss.Width  `gss:"width"`
}

// TODO: handle shorthand syntaxes during parsing
type Styles struct {
	Dimensions      Dimensions
	Margin          Margin  `gss:"margin"`
	Padding         Padding `gss:"padding"`
	Display         Display `gss:"display"`
	Text            Text
	Font            Font
	Border          Borders              `gss:"border"`
	BorderRadiuses  BorderRadiuses       `gss:"border-radius"`
	BackgroundColor gsse.Expr[gss.Color] `gss:"background-color"`
}

type Rule struct {
	Selector string
	Styles   *Styles
}

type Gss struct {
	Rules []*Rule
}
