package gss

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/parser/gss/core"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

// ParseBorderForOneEdge parses values of properties:
//   - border-top
//   - border-right
//   - border-bottom
//   - border-left
func ParseBorderForOneEdge(ts []css.Token) (ast.Border, error) {
	if !csstokens.IsBalanced(ts) {
		return ast.Border{}, fmt.Errorf("unbalanced parentheses")
	}
	b := ast.Border{
		Color: "currentcolor",
		Style: "solid",
		Width: "none",
	}
	for ts := range csstokens.Split(ts, css.WhitespaceToken, true) {
		if len(ts) > 1 {
			return ast.Border{}, fmt.Errorf("unsupported border property value length: %d", len(ts))
		}
		t := ts[0]
		switch {
		case core.IsColor(t):
			c, err := core.ParseColor(t)
			if err != nil {
				return ast.Border{}, fmt.Errorf("color: %w", err)
			}
			b.Color = c
		case core.IsBorderStyle(t):
			bs, err := core.ParseBorderStyle(t)
			if err != nil {
				return ast.Border{}, fmt.Errorf("style: %w", err)
			}
			b.Style = bs
		case core.IsBorderWidth(t):
			bw, err := core.ParseBorderWidth(t)
			if err != nil {
				return ast.Border{}, fmt.Errorf("width: %w", err)
			}
			b.Width = bw
		default:
			return ast.Border{}, fmt.Errorf("invalid border value: %q (type: %s)", string(t.Data), t.TokenType.String())
		}
	}
	return b, nil
}

// ParseBorder parses `border` property values
func ParseBorder(ts []css.Token) (ast.Borders, error) {
	ss, err := byCommas(ts, ParseBorderForOneEdge)
	if err != nil {
		return ast.Borders{}, err
	}
	b := ast.Borders{}
	b.Top, b.Right, b.Bottom, b.Left = directionalDemux(ss)
	return b, nil
}
