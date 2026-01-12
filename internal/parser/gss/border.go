package gss

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/parser/gss/core"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

// TODO: check input for [style]
// TODO: check input for [width]
func parseOneEdgeBorder(ts []css.Token) (ast.Border, error) {
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
		}
	}
	return b, nil
}

// ParseBorderTop parses `border-top` property values
func ParseBorderTop(ts []css.Token) (ast.Border, error) {
	return parseOneEdgeBorder(ts)
}

// ParseBorderRight parses `border-right` property values
func ParseBorderRight(ts []css.Token) (ast.Border, error) {
	return parseOneEdgeBorder(ts)
}

// ParseBorderBottom parses `border-bottom` property values
func ParseBorderBottom(ts []css.Token) (ast.Border, error) {
	return parseOneEdgeBorder(ts)
}

// ParseBorderLeft parses `border-left` property values
func ParseBorderLeft(ts []css.Token) (ast.Border, error) {
	return parseOneEdgeBorder(ts)
}

// ParseBorder parses `border` property values
func ParseBorder(ts []css.Token) (ast.Borders, error) {
	ss, err := byCommas(ts, parseOneEdgeBorder)
	if err != nil {
		return ast.Borders{}, err
	}
	b := ast.Borders{}
	b.Top, b.Right, b.Bottom, b.Left = directionalDemux(ss)
	return b, nil
}
