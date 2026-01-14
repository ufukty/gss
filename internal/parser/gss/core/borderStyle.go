package core

import (
	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
	"go.ufukty.com/gss/internal/tokens"
)

var borderStyles = map[tokens.BorderStyle]any{
	tokens.BorderStyleNone:   nil,
	tokens.BorderStyleHidden: nil,
	tokens.BorderStyleSolid:  nil,
	tokens.BorderStyleDashed: nil,
	tokens.BorderStyleDotted: nil,
}

func IsBorderStyle(tk css.Token) bool {
	return tk.TokenType == css.IdentToken && (has(borderStyles, tokens.BorderStyle(tk.Data)) || csstokens.IsGlobal(tk))
}

// returns [tokens.BorderStyle] or globals
func ParseBorderStyle(t css.Token) (any, error) {
	if has(borderStyles, tokens.BorderStyle(t.Data)) {
		return tokens.BorderStyle(t.Data), nil
	}
	if csstokens.IsGlobal(t) {
		return string(t.Data), nil // TODO: consider use of custom type for global keywords
	}
	return nil, nil
}
