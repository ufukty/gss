package gss

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/parser/gss/core"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

// ParsePaddingValue parses a single padding value (dimension or keyword)
func ParsePaddingValue(ts []css.Token) (any, error) {
	if len(ts) != 1 {
		return nil, fmt.Errorf("unexpected padding value token count: %d", len(ts))
	}
	t := ts[0]
	if !core.IsBorderWidth(t) {
		return nil, fmt.Errorf("invalid padding value: %q", t.Data)
	}
	return core.ParseBorderWidth(t)
}

// ParsePadding parses the `padding` shorthand property
// Supports CSS standard space-separated format:
//   - padding: 10px (all sides)
//   - padding: 10px 20px (top/bottom, left/right)
//   - padding: 10px 20px 30px (top, left/right, bottom)
//   - padding: 10px 20px 30px 40px (top, right, bottom, left)
func ParsePadding(ts []css.Token) (ast.Padding, error) {
	if !csstokens.IsBalanced(ts) {
		return ast.Padding{}, fmt.Errorf("unbalanced parentheses")
	}

	values := []any{}
	for ts := range csstokens.Split(ts, css.WhitespaceToken, true) {
		v, err := ParsePaddingValue(ts)
		if err != nil {
			return ast.Padding{}, fmt.Errorf("parsing value %d: %w", len(values), err)
		}
		values = append(values, v)
	}

	if len(values) == 0 || len(values) > 4 {
		return ast.Padding{}, fmt.Errorf("invalid number of padding values: %d", len(values))
	}

	p := ast.Padding{}
	p.Top, p.Right, p.Bottom, p.Left = directionalDemux(values)
	return p, nil
}
