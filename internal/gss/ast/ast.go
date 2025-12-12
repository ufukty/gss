package ast

import (
	"go.ufukty.com/gss/internal/gss/tokens"
	gsse "go.ufukty.com/gss/internal/gsse/ast"
)

type Display struct {
	Outside tokens.DisplayOutside
	Inside  tokens.DisplayInside
}

type Styles struct {
	Display         Display                `gss:"display"`
	Width           tokens.Width           `gss:"width"`
	Height          tokens.Height          `gss:"height"`
	FontFamily      []tokens.FontFamily    `gss:"font-family"`
	FontSize        gsse.Size              `gss:"font-size"`
	Color           tokens.Color           `gss:"color"`
	BackgroundColor tokens.BackgroundColor `gss:"background-color"`
}

type Rule struct {
	Selector string
	Styles   *Styles
}

type Gss struct {
	Rules []*Rule
}
