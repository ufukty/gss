package csstokens

import (
	"fmt"
	"slices"
	"testing"

	"github.com/tdewolff/parse/v2/css"
)

func Example_tokenize_borderStyle() {
	tcs := []string{
		"none", "hidden", "solid", "dotted", "dashed", "double",
		"groove", "ridge", "inset", "outset",
		"inherit", "initial", "unset", "revert", "revert-layer",
	}
	for _, tc := range tcs {
		ts, err := Tokenize(fmt.Sprintf("border-style: %s", tc))
		if err != nil {
			panic(fmt.Errorf("prep, tokenize: %v", err))
		}
		for _, t := range ts {
			fmt.Printf("%v\n", t)
		}
	}
	// Output:
	// Ident('none')
	// Ident('hidden')
	// Ident('solid')
	// Ident('dotted')
	// Ident('dashed')
	// Ident('double')
	// Ident('groove')
	// Ident('ridge')
	// Ident('inset')
	// Ident('outset')
	// Ident('inherit')
	// Ident('initial')
	// Ident('unset')
	// Ident('revert')
	// Ident('revert-layer')
}

func Example_tokenize_borderWidth() {
	tcs := []string{
		"thin", "medium", "thick",
		"0",
		"1px", "4px", "0.1em", "0.5em", "1rem", "1pt", "1cm", "1mm", "1in", "1ch", "1ex", "1vw", "1vh",
		"inherit", "initial", "unset", "revert", "revert-layer",
	}
	for _, tc := range tcs {
		ts, err := Tokenize(fmt.Sprintf("border-width: %s", tc))
		if err != nil {
			panic(fmt.Errorf("prep, tokenize: %v", err))
		}
		for _, t := range ts {
			fmt.Printf("%v\n", t)
		}
	}
	// Output:
	// Ident('thin')
	// Ident('medium')
	// Ident('thick')
	// Number('0')
	// Dimension('1px')
	// Dimension('4px')
	// Dimension('0.1em')
	// Dimension('0.5em')
	// Dimension('1rem')
	// Dimension('1pt')
	// Dimension('1cm')
	// Dimension('1mm')
	// Dimension('1in')
	// Dimension('1ch')
	// Dimension('1ex')
	// Dimension('1vw')
	// Dimension('1vh')
	// Ident('inherit')
	// Ident('initial')
	// Ident('unset')
	// Ident('revert')
	// Ident('revert-layer')
}

func Example_compare() {
	ts, err := Tokenize("border: inherit")
	if err != nil {
		panic(fmt.Errorf("prep, Tokenize: %v", err))
	}
	fmt.Println(compare(ts[0], inherit)) // Output: true
}

func Example_isGlobal() {
	ts, err := Tokenize("border: inherit")
	if err != nil {
		panic(fmt.Errorf("prep, Tokenize: %v", err))
	}
	fmt.Println(isGlobal(ts[0])) // Output: true
}

func compareSplits(t *testing.T, expected, got [][]css.Token) {
	if len(expected) != len(got) {
		t.Errorf("number of splits don't match: expected %d, got %d", len(expected), len(got))
	}
	for i := 0; i < min(len(expected), len(got)); i++ {
		t.Run(fmt.Sprintf("split %d", i), func(t *testing.T) {
			e, g := expected[i], got[i]
			le, lg := len(e), len(g)
			if le != lg {
				t.Errorf("split lengths don't match: expected %d, got %d", le, lg)
			}
			for j := 0; j < min(le, lg); j++ {
				if !compare(e[j], g[j]) {
					t.Errorf("splits differ at the index %d; expected %s, got %s", j, e[j].String(), g[j].String())
				}
			}
		})
	}
}

func tokens(types ...css.TokenType) []css.Token {
	ts := make([]css.Token, 0, len(types))
	for _, t := range types {
		ts = append(ts, css.Token{TokenType: t})
	}
	return ts
}

func TestSplit(t *testing.T) {
	var (
		input = tokens(
			css.WhitespaceToken,
			css.IdentToken,
			css.ColumnToken,
			css.CommentToken,
			css.WhitespaceToken,
			css.WhitespaceToken,
			css.IdentToken,
			css.WhitespaceToken,
			css.IdentToken,
			css.LeftBracketToken,
		)
		expected = [][]css.Token{
			tokens(
				css.IdentToken,
				css.ColumnToken,
				css.CommentToken,
			),
			tokens(
				css.IdentToken,
			),
			tokens(
				css.IdentToken,
				css.LeftBracketToken,
			),
		}
	)
	got := slices.Collect(Split(input, css.WhitespaceToken))
	compareSplits(t, expected, got)
}
