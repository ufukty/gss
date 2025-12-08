package ast

import "go.ufukty.com/gss/internal/gss/tokens"

type FontSize struct {
	Number float64
	Unit   tokens.Unit
}

type Styles struct {
	Display         tokens.Display         `gss:"display"`
	Width           tokens.Width           `gss:"width"`
	Height          tokens.Height          `gss:"height"`
	FontFamily      []tokens.FontFamily    `gss:"font-family"`
	FontSize        FontSize               `gss:"font-size"`
	Color           tokens.Color           `gss:"color"`
	BackgroundColor tokens.BackgroundColor `gss:"background-color"`
}

type Rule struct {
	Selector string
	Styles   Styles
}

type Gss struct {
	Rules []Rule
}
