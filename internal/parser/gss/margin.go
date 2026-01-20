package gss

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/parser/gss/core"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

// isMarginValue checks if a token is a valid margin value
func isMarginValue(t css.Token) bool {
	// Margin accepts: dimensions, 0, "auto", or global keywords
	if core.IsBorderWidth(t) {
		return true
	}
	// Check for "auto" keyword
	if t.TokenType == css.IdentToken && string(t.Data) == "auto" {
		return true
	}
	return false
}

// parseMarginValue parses a single margin value token
func parseMarginValue(t css.Token) (any, error) {
	// Handle "auto" keyword
	if t.TokenType == css.IdentToken && string(t.Data) == "auto" {
		return "auto", nil
	}

	// Handle dimensional values and global keywords using border-width parser
	return core.ParseBorderWidth(t)
}

// ParseMarginForOneSide parses a single margin value
func ParseMarginForOneSide(ts []css.Token) (any, error) {
	if !csstokens.IsBalanced(ts) {
		return nil, fmt.Errorf("unbalanced parentheses")
	}

	if len(ts) != 1 {
		return nil, fmt.Errorf("expected single value, got %d tokens", len(ts))
	}

	t := ts[0]
	if !isMarginValue(t) {
		return nil, fmt.Errorf("invalid margin value: %q (type: %s)", string(t.Data), t.TokenType.String())
	}

	return parseMarginValue(t)
}

// ParseMargin parses `margin` property values with directional shorthand support
// Supports: margin: <value>; (all sides)
//           margin: <vertical> <horizontal>;
//           margin: <top> <horizontal> <bottom>;
//           margin: <top> <right> <bottom> <left>;
func ParseMargin(ts []css.Token) (ast.Margin, error) {
	ss, err := byWhitespace(ts, ParseMarginForOneSide)
	if err != nil {
		return ast.Margin{}, err
	}

	if len(ss) == 0 || len(ss) > 4 {
		return ast.Margin{}, fmt.Errorf("invalid number of margin values: %d (expected 1-4)", len(ss))
	}

	m := ast.Margin{}
	m.Top, m.Right, m.Bottom, m.Left = directionalDemux(ss)
	return m, nil
}

// byWhitespace splits tokens by whitespace and parses each group
func byWhitespace(ts []css.Token, parser func([]css.Token) (any, error)) ([]any, error) {
	if !csstokens.IsBalanced(ts) {
		return nil, fmt.Errorf("unbalanced parentheses")
	}

	ss := make([]any, 0, len(ts))
	for ts := range csstokens.Split(ts, css.WhitespaceToken, true) {
		s, err := parser(ts)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}
