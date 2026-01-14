package core

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/dimensional"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

func IsBorderWidth(t css.Token) bool {
	return t.TokenType == css.DimensionToken || t.TokenType == css.NumberToken || csstokens.IsGlobal(t)
}

var regexDimension = regexp.MustCompile(`([0-9]+(?:\.[0-9]+)?)(px|pt|em|rem|vw|vh)`)

func ParseBorderWidth(t css.Token) (any, error) {
	switch {
	case t.TokenType == css.DimensionToken:
		if ms := regexDimension.FindStringSubmatch(string(t.Data)); len(ms) > 0 {
			f, err := strconv.ParseFloat(ms[1], 64)
			if err != nil {
				return nil, fmt.Errorf("parsing the value of dimensional: %w", err)
			}
			return dimensional.New(f, dimensional.Unit(ms[2])), nil
		}

	case t.TokenType == css.NumberToken:
		if string(t.Data) == "0" {
			return dimensional.New(0), nil
		}

	case csstokens.IsGlobal(t):
		return string(t.Data), nil

	}
	return nil, fmt.Errorf("unknown token type: %q", t.TokenType.String())
}
