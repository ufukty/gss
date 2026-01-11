package csstokens

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

func Tokenize(in string) ([]css.Token, error) {
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

func IsBalanced(ts []css.Token) bool {
	b := 0
	for _, t := range ts {
		switch t.TokenType {
		case css.LeftParenthesisToken:
			b++
		case css.RightParenthesisToken:
			b--
			if b < 0 {
				return false
			}
		}
	}
	return b == 0
}

// Tokens MUST be balanced. Separators inside matching parentheses
// are treated regularly when scoped.
func Split(ts []css.Token, sep css.TokenType, scoped bool) iter.Seq[[]css.Token] {
	ts = append(ts, css.Token{TokenType: sep}) // for the last split
	return func(yield func([]css.Token) bool) {
		prev := 0
		baln := 0
		for cur, t := range ts {
			switch t.TokenType {
			case css.LeftParenthesisToken:
				baln++
			case css.RightParenthesisToken:
				baln--
			case sep:
				if !scoped || baln == 0 {
					if cur > prev && !yield(ts[prev:cur]) {
						return
					}
					prev = cur + 1
				}
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
