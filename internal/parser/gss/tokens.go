package gss

import (
	"io"
	"slices"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

var inherit = css.Token{TokenType: css.IdentToken, Data: ([]byte)("inherit")}

func tokenize(in string) ([]css.Token, error) {
	p := css.NewParser(parse.NewInputString(in), true)
	gt, _, _ := p.Next()
	if gt == css.ErrorGrammar {
		if err := p.Err(); err != io.EOF {
			return nil, err
		}
	}
	return p.Values(), nil
}

func compare(a, b css.Token) bool {
	return a.TokenType == b.TokenType && slices.Compare(a.Data, b.Data) == 0
}
