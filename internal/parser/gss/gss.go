package gss

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

// TRBL, TB|RL, T|RL|B, T|R|B|L
func directionalDemux[T any](ss []T) (T, T, T, T) {
	switch len(ss) {
	case 1:
		return ss[0], ss[0], ss[0], ss[0]
	case 2:
		return ss[0], ss[1], ss[0], ss[1]
	case 3:
		return ss[0], ss[1], ss[2], ss[1]
	case 4:
		return ss[0], ss[1], ss[2], ss[3]
	}
	panic(fmt.Sprintf("unexpected number of directional shorthand values: %d", len(ss)))
}

// use for TRBL, TB|RL, T|RL|B, T|R|B|L
func byCommas[T any](ts []css.Token, parser func([]css.Token) (T, error)) ([]T, error) {
	if !csstokens.IsBalanced(ts) {
		return nil, fmt.Errorf("unbalanced parentheses")
	}
	ss := make([]T, 0, len(ts))
	for ts := range csstokens.Split(ts, css.CommaToken, true) {
		b, err := parser(ts)
		if err != nil {
			return nil, fmt.Errorf("comma separated values %d: %w", len(ss), err)
		}
		ss = append(ss, b)
	}
	return ss, nil
}
