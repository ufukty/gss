package gss

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/parser/gss/core"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

// ParseMarginValue parses a single margin value (dimension or keyword)
func ParseMarginValue(ts []css.Token) (any, error) {
	if len(ts) != 1 {
		return nil, fmt.Errorf("unexpected margin value token count: %d", len(ts))
	}
	t := ts[0]
	if !core.IsBorderWidth(t) {
		return nil, fmt.Errorf("invalid margin value: %q", t.Data)
	}
	return core.ParseBorderWidth(t)
}

// ParseMargin parses the `margin` shorthand property
// Supports CSS standard space-separated format:
//   - margin: 10px (all sides)
//   - margin: 10px 20px (top/bottom, left/right)
//   - margin: 10px 20px 30px (top, left/right, bottom)
//   - margin: 10px 20px 30px 40px (top, right, bottom, left)
func ParseMargin(ts []css.Token) (ast.Margin, error) {
	if !csstokens.IsBalanced(ts) {
		return ast.Margin{}, fmt.Errorf("unbalanced parentheses")
	}

	values := []any{}
	for ts := range csstokens.Split(ts, css.WhitespaceToken, true) {
		v, err := ParseMarginValue(ts)
		if err != nil {
			return ast.Margin{}, fmt.Errorf("parsing value %d: %w", len(values), err)
		}
		values = append(values, v)
	}

	if len(values) == 0 || len(values) > 4 {
		return ast.Margin{}, fmt.Errorf("invalid number of margin values: %d", len(values))
	}

	m := ast.Margin{}
	m.Top, m.Right, m.Bottom, m.Left = directionalDemux(values)
	return m, nil
}
