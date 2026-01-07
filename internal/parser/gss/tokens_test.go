package gss

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
		ts, err := tokenize(fmt.Sprintf("border-style: %s", tc))
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
		ts, err := tokenize(fmt.Sprintf("border-width: %s", tc))
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
	ts, err := tokenize("border: inherit")
	if err != nil {
		panic(fmt.Errorf("prep, tokenize: %v", err))
	}
	fmt.Println(compare(ts[0], inherit)) // Output: true
}

func Example_isGlobal() {
	ts, err := tokenize("border: inherit")
	if err != nil {
		panic(fmt.Errorf("prep, tokenize: %v", err))
	}
	fmt.Println(isGlobal(ts[0])) // Output: true
}

func TestSplit(t *testing.T) {
	input := []css.Token{
		{TokenType: css.WhitespaceToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.IdentToken},
		{TokenType: css.WhitespaceToken},
	}
	ss := slices.Collect(split(input, css.WhitespaceToken))
	if len(ss) != 2 {
		t.Fatalf("assert, number of splits: expected 2, got %d", len(ss))
	}
	if len(ss[0]) != 3 {
		t.Errorf("assert, length of first split: expected 3, got %d", len(ss[0]))
	}
	if len(ss[1]) != 1 {
		t.Errorf("assert, length of first split: expected 1, got %d", len(ss[1]))
	}
}
