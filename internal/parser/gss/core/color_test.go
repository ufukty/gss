package core

import (
	"fmt"
	"testing"

	"github.com/tdewolff/parse/v2/css"
)

func TestTrimSpaces(t *testing.T) {
	input := []css.Token{
		{TokenType: css.WhitespaceToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.WhitespaceToken},
	}
	trimmed := trimSpaces(input)
	got := len(trimmed)
	if got != 4 {
		t.Errorf("expected 4 got %d", got)
	}
	for _, tk := range trimmed {
		if tk.TokenType == css.WhitespaceToken {
			t.Errorf("unexpected whitespace token in trimmed array")
		}
	}
}

func TestIsColor_positive(t *testing.T) {
	tcs := []string{
		"#fff",
		"#ffff",
		"#ffffff",
		"#ffffffff",

		"red",
		"ReD",
		"transparent",
		"currentColor",
	}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			toks, _ := tokenize(fmt.Sprintf("border-color: %s", tc))
			if !IsColor(toks[0]) {
				t.Fatal("unexpectedly false")
			}
		})
	}
}
