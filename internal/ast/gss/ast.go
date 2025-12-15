package gss

import (
	"go.ufukty.com/gss/internal/ast/gsse"
	"go.ufukty.com/gss/internal/tokens/gss"
)

type Display struct {
	Outside gss.DisplayOutside
	Inside  gss.DisplayInside
}

type Styles struct {
	Display         Display              `gss:"display"`
	Width           gss.Width            `gss:"width"`
	Height          gss.Height           `gss:"height"`
	FontFamily      []gss.FontFamily     `gss:"font-family"`
	FontSize        gsse.Expr[gsse.Size] `gss:"font-size"`
	Color           gss.Color            `gss:"color"`
	BackgroundColor gss.BackgroundColor  `gss:"background-color"`
}

type Rule struct {
	Selector string
	Styles   *Styles
}

type Gss struct {
	Rules []*Rule
}
