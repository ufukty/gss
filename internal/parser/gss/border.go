package gss

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
	"go.ufukty.com/gss/internal/parser/gss/is"
)

// TODO: check input for [color]
// TODO: check input for [style]
// TODO: check input for [width]
func ParseBorder(ts []css.Token) (*ast.Border, error) {
	b := ast.Border{
		Color:     "#000000",
		Style:     "solid",
		Thickness: "none",
	}
	for ts := range csstokens.Split(ts, css.WhitespaceToken, true) {
		if len(ts) > 1 {
			return nil, fmt.Errorf("unsupported border property value length: %d", len(ts))
		}
		t := ts[0]
		switch {
		case is.Color(t):
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
