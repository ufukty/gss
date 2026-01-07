package gss

import (
	"io"
	"iter"
	"slices"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

var (
	inherit     = css.Token{TokenType: css.IdentToken, Data: ([]byte)("inherit")}
	initial     = css.Token{TokenType: css.IdentToken, Data: ([]byte)("initial")}
	unset       = css.Token{TokenType: css.IdentToken, Data: ([]byte)("unset")}
	revert      = css.Token{TokenType: css.IdentToken, Data: ([]byte)("revert")}
	revertLayer = css.Token{TokenType: css.IdentToken, Data: ([]byte)("revert-layer")}
)

var globals = []css.Token{inherit, initial, unset, revert, revertLayer}

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

func split(ts []css.Token, sep css.TokenType) iter.Seq[[]css.Token] {
	return func(yield func([]css.Token) bool) {
		prev := 0
		for cur, t := range ts {
			if t.TokenType == sep {
				if cur > prev && !yield(ts[prev:cur]) {
					return
				}
				prev = cur + 1
			}
		}
	}
}

func isGlobal(t css.Token) bool {
	for i := range len(globals) {
		if compare(globals[i], t) {
			return true
		}
	}
	return false
}
