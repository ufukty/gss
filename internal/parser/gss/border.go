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
func ParseBorder(ts []css.Token) (*ast.Border, error) {
	if !csstokens.IsBalanced(ts) {
		return nil, fmt.Errorf("unbalanced parentheses")
	}
	b := ast.Border{
		Color: "#000000",
		Style: "solid",
		Width: "none",
	}
	for ts := range csstokens.Split(ts, css.WhitespaceToken, true) {
		if len(ts) > 1 {
			return nil, fmt.Errorf("unsupported border property value length: %d", len(ts))
		}
		t := ts[0]
		switch {
		case core.IsColor(t):
			c, err := core.ParseColor(t)
			if err != nil {
				return nil, fmt.Errorf("color: %w", err)
			}
			b.Color = c
		}
	}
	return &b, nil
}

func ParseBorders(ts []css.Token) (*ast.Borders, error) {
	if !csstokens.IsBalanced(ts) {
		return nil, fmt.Errorf("unbalanced parentheses")
	}
	ss := make([]*ast.Border, 0, len(ts))
	for ts := range csstokens.Split(ts, css.CommaToken, true) {
		b, err := ParseBorder(ts)
		if err != nil {
			return nil, fmt.Errorf("comma separated values %d: %w", len(ss), err)
		}
		ss = append(ss, b)
	}
	bs := ast.Borders{}
	switch len(ss) {
	case 1:
		bs.Top, bs.Right, bs.Bottom, bs.Left = *ss[0], *ss[0], *ss[0], *ss[0]
	case 2:
		bs.Top, bs.Bottom, bs.Left, bs.Right = *ss[0], *ss[0], *ss[1], *ss[1]
	case 3:
		bs.Top, bs.Left, bs.Right, bs.Bottom = *ss[0], *ss[1], *ss[1], *ss[2]
	case 4:
		bs.Top, bs.Right, bs.Bottom, bs.Left = *ss[0], *ss[1], *ss[2], *ss[3]
	}
	return &bs, nil
}
